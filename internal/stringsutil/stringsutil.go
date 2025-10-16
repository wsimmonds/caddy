package stringsutil

import "strings"

// TrimFQDNTrailingDot removes a single trailing dot from a hostname, treating
// fully-qualified domain names with a trailing dot as equivalent to their
// non-dotted form. The root label (".") is preserved.
func TrimFQDNTrailingDot(host string) string {
	if host == "" || host == "." {
		return host
	}
	return strings.TrimSuffix(host, ".")
}
