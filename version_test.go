package loginradius

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

// The SDK version reaches customers through three channels that are edited in
// different places: version.go (rendered from the manifest), the CHANGELOG
// heading, and the git tag. They drifted once already — version.go said
// 12.0.0-rc.1 while the CHANGELOG still announced v12.0.0 — which is the kind
// of mismatch nobody notices until a release is cut against the wrong notes.
//
// These assert the two the SDK can see. The tag is checked at release time.

var semverRE = regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-[0-9A-Za-z.-]+)?(\+[0-9A-Za-z.-]+)?$`)

func TestVersionIsProxyResolvableSemver(t *testing.T) {
	if strings.HasPrefix(Version, "v") {
		t.Fatalf("Version = %q: must not carry a leading v — the git TAG does (v%s), the constant does not", Version, Version)
	}
	if !semverRE.MatchString(Version) {
		t.Fatalf("Version = %q is not valid semver; the Go module proxy will not serve it", Version)
	}
}

// The module path ends in /v12, so major 12 is the only one the proxy accepts.
func TestVersionMajorMatchesModulePath(t *testing.T) {
	if !strings.HasPrefix(Version, "12.") {
		t.Fatalf("Version = %q but the module path is .../v12: a mismatched major is unresolvable", Version)
	}
}

func TestVersionMatchesChangelogHeading(t *testing.T) {
	b, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		t.Fatalf("reading CHANGELOG.md: %v", err)
	}
	// First "## v<version>" heading is the release being described.
	m := regexp.MustCompile(`(?m)^## v(\S+)`).FindSubmatch(b)
	if m == nil {
		t.Fatal("CHANGELOG.md has no '## v<version>' heading")
	}
	if got := string(m[1]); got != Version {
		t.Fatalf("CHANGELOG announces v%s but Version = %q — release notes and artifact disagree", got, Version)
	}
}
