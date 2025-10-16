package stringsutil

import "testing"

func TestTrimFQDNTrailingDot(t *testing.T) {
	tests := []struct {
		name string
		host string
		want string
	}{
		{
			name: "empty",
			host: "",
			want: "",
		},
		{
			name: "root label",
			host: ".",
			want: ".",
		},
		{
			name: "no trailing dot",
			host: "example.com",
			want: "example.com",
		},
		{
			name: "single trailing dot",
			host: "example.com.",
			want: "example.com",
		},
		{
			name: "unicode hostname",
			host: "täst.de.",
			want: "täst.de",
		},
		{
			name: "multiple trailing dots",
			host: "example.com..",
			want: "example.com.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TrimFQDNTrailingDot(tc.host)
			if got != tc.want {
				t.Fatalf("TrimFQDNTrailingDot(%q) = %q; want %q", tc.host, got, tc.want)
			}
		})
	}
}
