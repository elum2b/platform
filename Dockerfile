FROM golang:1.27-alpine AS build

ARG VERSION=dev

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -trimpath \
    -ldflags="-s -w -X github.com/elum2b/platform/internal/utils/version.current=${VERSION}" \
    -o /platform \
    .

RUN mkdir /data

FROM gcr.io/distroless/static-debian12:nonroot

ARG VERSION=dev

LABEL io.elum2b.platform.version=$VERSION

COPY --from=build /platform /platform
COPY --from=build --chown=nonroot:nonroot /data /var/lib/platform

USER nonroot:nonroot

ENTRYPOINT ["/platform"]
