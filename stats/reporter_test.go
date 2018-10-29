package stats

import "testing"

func TestAddTagsToName(t *testing.T) {
	tests := []struct {
		name     string
		tags     map[string]string
		expected string
	}{
		{
			name:     "recvd",
			tags:     nil,
			expected: "recvd.no-os.no-browser",
		},
		{
			name: "recvd",
			tags: map[string]string{
				"os":       "Linux",
				"browser":  "Chrome",
			},
			expected: "recvd.Linux.Chrome",
		},
		{
			name: "r.call",
			tags: map[string]string{
				"host":     "my-host-name",
				"os":       "Linu{}/\tx",
				"browser":  "Chro\\:me",
			},
			expected: "r.call.my-host-name.Linu----x.Chro--me",
		},
	}

	for _, tt := range tests {
		got := addTagsToName(tt.name, tt.tags)
		if got != tt.expected {
			t.Errorf("addTagsToName(%v, %v) got %v, expected %v",
				tt.name, tt.tags, got, tt.expected)
		}
	}
}
