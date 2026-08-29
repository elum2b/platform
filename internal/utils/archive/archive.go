// Package archive contains archive job helpers.
package archive

import (
	"fmt"
	"time"
)

// FileName returns a unique ZIP archive name for a platform service.
func FileName(service string) string {
	return fileName(service, time.Now().UTC())
}

func fileName(service string, now time.Time) string {
	return fmt.Sprintf("%s_%d.zip", service, now.UnixMilli())
}
