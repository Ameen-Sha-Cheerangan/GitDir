// Package version provides build-time version metadata.
// Values are injected via -ldflags at build time.
package version

import "fmt"

// These variables are set at build time using:
//
//	go build -ldflags "-X github.com/ameen/gitdir/internal/version.tag=v1.0.0"
var (
	tag    = "dev"
	commit = "none"
)

// String returns the human-readable version string.
func String() string {
	return fmt.Sprintf("%s (commit: %s)", tag, commit)
}
