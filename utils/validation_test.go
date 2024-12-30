package utils_test

import (
	"testing"

	"github.com/thesayedirfan/weather/utils"
)

func TestIsValidIP(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		expected bool
	}{
		{"Valid IPv4 address", "192.168.1.1", true},
		{"Valid IPv6 address", "2001:db8::ff00:42:8329", true},
		{"Invalid IP address", "999.999.999.999", false},
		{"Empty string", "", false},
		{"Random string", "invalid_ip", false},
		{"IPv4 with invalid segment", "192.168.1.999", false},
		{"IPv6 malformed", "2001:db8:::1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.IsValidIP(tt.ip)
			if result != tt.expected {
				t.Errorf("IsValidIP(%q) = %v; want %v", tt.ip, result, tt.expected)
			}
		})
	}
}