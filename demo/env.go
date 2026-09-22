package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// loadDotEnv reads the first .env file it finds and copies the entries into
// the process environment. Stdlib only, on purpose: the demo ships inside the
// SDK module, so a dotenv dependency here would land in every customer's
// go.mod for code they never build.
//
// Real environment variables always win — an exported LR_API_KEY is not
// clobbered by a stale value in a checked-out .env.
//
// DEMO ONLY. A .env holds live tenant credentials, so it is gitignored and
// must never be committed. Production deployments inject secrets through the
// platform's secret manager, not a file on disk.
func loadDotEnv() {
	// demo/ is its own module, so it is always the working directory — there is
	// no outer root to fall back to. Looking for demo/.env from here found
	// nothing and printed a warning on every start, which read like a failure.
	for _, path := range []string{".env"} {
		loaded, err := loadEnvFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: %s: %v\n", path, err)
			continue
		}
		if loaded {
			log.Printf("loaded env from %s", path)
			return
		}
	}
}

// loadEnvFile applies one file. Reports false (with no error) when the file
// does not exist, so a missing .env stays a non-event.
func loadEnvFile(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for lineNo := 1; sc.Scan(); lineNo++ {
		key, value, ok := parseEnvLine(sc.Text())
		if !ok {
			continue
		}
		if key == "" {
			fmt.Fprintf(os.Stderr, "warning: %s:%d: malformed line, skipped\n", path, lineNo)
			continue
		}
		// Never overwrite the real environment: an explicitly exported value
		// is a deliberate override of whatever the file says.
		if _, set := os.LookupEnv(key); set {
			continue
		}
		if err := os.Setenv(key, value); err != nil {
			return false, fmt.Errorf("set %s: %w", key, err)
		}
	}
	if err := sc.Err(); err != nil {
		return false, err
	}
	return true, nil
}

// parseEnvLine splits one KEY=VALUE line. ok is false for blanks and comments
// (nothing to do); ok is true with an empty key for a line that looked like an
// assignment but wasn't valid, so the caller can warn about it.
func parseEnvLine(line string) (key, value string, ok bool) {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "#") {
		return "", "", false
	}
	line = strings.TrimPrefix(line, "export ")

	name, raw, found := strings.Cut(line, "=")
	if !found {
		return "", "", true
	}
	name = strings.TrimSpace(name)
	if !validEnvKey(name) {
		return "", "", true
	}
	return name, unquote(strings.TrimSpace(raw)), true
}

// unquote strips one matching pair of surrounding quotes. Unquoted values also
// drop a trailing `# comment`, which a quoted value keeps verbatim.
func unquote(v string) string {
	if len(v) >= 2 && (v[0] == '"' || v[0] == '\'') && v[len(v)-1] == v[0] {
		return v[1 : len(v)-1]
	}
	if i := strings.Index(v, " #"); i >= 0 {
		v = v[:i]
	}
	return strings.TrimSpace(v)
}

func validEnvKey(k string) bool {
	if k == "" {
		return false
	}
	for i, r := range k {
		switch {
		case r >= 'A' && r <= 'Z', r >= 'a' && r <= 'z', r == '_':
		case r >= '0' && r <= '9' && i > 0:
		default:
			return false
		}
	}
	return true
}
