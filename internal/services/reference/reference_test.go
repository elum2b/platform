package reference

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/elum2b/services/reference/storage"
)

func TestResourceStorageConfig(t *testing.T) {
	tests := map[string]struct {
		options storage.Config
		want    storage.Config
	}{
		"local directory fallback": {
			want: storage.Config{
				Directory: filepath.Join("/var/lib/platform", "reference"),
			},
		},
		"configured S3 storage": {
			options: storage.Config{
				Directory:    "/mnt/reference",
				Endpoint:     "minio:9000",
				Bucket:       "reference",
				AccessKey:    "access-key",
				SecretKey:    "secret-key",
				SessionToken: "session-token",
				Region:       "us-east-1",
				Secure:       false,
				UsePathStyle: true,
			},
			want: storage.Config{
				Directory:    "/mnt/reference",
				Endpoint:     "minio:9000",
				Bucket:       "reference",
				AccessKey:    "access-key",
				SecretKey:    "secret-key",
				SessionToken: "session-token",
				Region:       "us-east-1",
				Secure:       false,
				UsePathStyle: true,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := resourceStorageConfig("/var/lib/platform", test.options)
			if !reflect.DeepEqual(actual, test.want) {
				t.Errorf(
					"resourceStorageConfig() = %#v, want %#v",
					actual,
					test.want,
				)
			}
		})
	}
}
