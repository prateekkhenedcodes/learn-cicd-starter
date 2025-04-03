package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey mykey"},
			},
			wantAPIKey: "mykey",
			wantErr:    false,
		},
		{
			name:       "missing Authorization header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name: "malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"Invalid mykey"},
			},
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name: "malformed Authorization header - missing space",
			headers: http.Header{
				"Authorization": []string{"ApiKeymykey"},
			},
			wantAPIKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() = %v, want %v", gotAPIKey, tt.wantAPIKey)
			}
		})
	}
}
