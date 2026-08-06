package version

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	serviceerrors "github.com/elum2b/services/errors"
)

const (
	defaultRegistryURL = "https://ghcr.io"
	versionLabel       = "io.elum2b.platform.version"
	ociVersionLabel    = "org.opencontainers.image.version"
)

var current = "dev"

// Client retrieves the latest published platform image version.
type Client struct {
	httpClient  *http.Client
	registryURL string
	repository  string
	cacheTTL    time.Duration

	mu        sync.Mutex
	latest    string
	expiresAt time.Time
}

// Current returns the version embedded into the running binary at build time.
func Current() string {
	return current
}

// NewClient creates a client for a public OCI registry repository.
func NewClient(
	httpClient *http.Client,
	registryURL string,
	repository string,
	cacheTTL time.Duration,
) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	if registryURL == "" {
		registryURL = defaultRegistryURL
	}

	return &Client{
		httpClient:  httpClient,
		registryURL: strings.TrimRight(registryURL, "/"),
		repository:  strings.Trim(repository, "/"),
		cacheTTL:    cacheTTL,
	}
}

// Latest returns the latest published version, refreshing an expired cache.
func (client *Client) Latest(ctx context.Context) (string, error) {
	client.mu.Lock()
	defer client.mu.Unlock()

	if client.latest != "" && time.Now().Before(client.expiresAt) {
		return client.latest, nil
	}

	latest, err := client.fetchLatest(ctx)
	if err != nil {
		if client.latest != "" {
			return client.latest, nil
		}

		return "", serviceerrors.Wrap(
			serviceerrors.CodeUnavailable,
			"latest platform version is unavailable",
			err,
		)
	}

	client.latest = latest
	client.expiresAt = time.Now().Add(client.cacheTTL)

	return latest, nil
}

func (client *Client) fetchLatest(ctx context.Context) (string, error) {
	token, err := client.token(ctx)
	if err != nil {
		return "", err
	}

	manifest, err := client.manifest(ctx, token, "latest")
	if err != nil {
		return "", err
	}

	if manifest.Config.Digest == "" && len(manifest.Manifests) != 0 {
		manifest, err = client.manifest(
			ctx,
			token,
			manifest.Manifests[0].Digest,
		)
		if err != nil {
			return "", err
		}
	}

	if version := imageVersion(manifest.Annotations); version != "" {
		return version, nil
	}

	if manifest.Config.Digest == "" {
		return "", fmt.Errorf("latest image manifest does not contain config")
	}

	config, err := client.imageConfig(ctx, token, manifest.Config.Digest)
	if err != nil {
		return "", err
	}

	if version := imageVersion(config.Config.Labels); version != "" {
		return version, nil
	}

	return "", fmt.Errorf("latest image does not declare a version")
}

func (client *Client) token(ctx context.Context) (string, error) {
	endpoint, err := url.Parse(client.registryURL + "/token")
	if err != nil {
		return "", fmt.Errorf("parse registry token URL: %w", err)
	}

	query := endpoint.Query()
	query.Set("service", endpoint.Host)
	query.Set("scope", "repository:"+client.repository+":pull")

	endpoint.RawQuery = query.Encode()

	var response struct {
		AccessToken string `json:"access_token"`
		Token       string `json:"token"`
	}

	if err := client.getJSON(
		ctx,
		endpoint.String(),
		"",
		&response,
	); err != nil {
		return "", err
	}

	if response.Token != "" {
		return response.Token, nil
	}

	if response.AccessToken != "" {
		return response.AccessToken, nil
	}

	return "", fmt.Errorf("registry token response does not contain a token")
}

func (client *Client) manifest(
	ctx context.Context,
	token string,
	reference string,
) (imageManifest, error) {
	var response imageManifest

	url := client.registryURL + "/v2/" + client.repository + "/manifests/" + reference

	if err := client.getJSON(ctx, url, token, &response); err != nil {
		return imageManifest{}, err
	}

	return response, nil
}

func (client *Client) imageConfig(
	ctx context.Context,
	token string,
	digest string,
) (imageConfig, error) {
	var response imageConfig

	url := client.registryURL + "/v2/" + client.repository + "/blobs/" + digest

	if err := client.getJSON(ctx, url, token, &response); err != nil {
		return imageConfig{}, err
	}

	return response, nil
}

func (client *Client) getJSON(
	ctx context.Context,
	url string,
	token string,
	target any,
) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create registry request: %w", err)
	}

	request.Header.Set(
		"Accept",
		"application/vnd.oci.image.manifest.v1+json, "+
			"application/vnd.oci.image.index.v1+json, "+
			"application/vnd.docker.distribution.manifest.v2+json",
	)

	if token != "" {
		request.Header.Set("Authorization", "Bearer "+token)
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("request registry: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK ||
		response.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("registry returned status %d", response.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(response.Body, 1<<20))
	if err != nil {
		return fmt.Errorf("read registry response: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("decode registry response: %w", err)
	}

	return nil
}

func imageVersion(labels map[string]string) string {
	if labels == nil {
		return ""
	}

	if version := strings.TrimSpace(labels[versionLabel]); version != "" {
		return version
	}

	return strings.TrimSpace(labels[ociVersionLabel])
}

type imageManifest struct {
	Config struct {
		Digest string `json:"digest"`
	} `json:"config"`
	Annotations map[string]string `json:"annotations"`
	Manifests   []struct {
		Digest string `json:"digest"`
	} `json:"manifests"`
}

type imageConfig struct {
	Config struct {
		//nolint:tagliatelle // OCI image config uses an uppercase Labels field.
		Labels map[string]string `json:"Labels"`
	} `json:"config"`
}
