package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers        http.Header
		expectedAPIKey string
		wantErr        bool
	}{
		"no authorization header": {
			headers: http.Header{
				"Authorization": []string{
					"",
				},
			},
			wantErr: true,
		},
		"malformed authorization header": {
			headers: http.Header{
				"Authorization": []string{
					"Bearer",
				},
			},
			wantErr: true,
		},
		"valid authorization header": {
			headers: http.Header{
				"Authorization": []string{
					"ApiKey my-api-key",
				},
			},
			expectedAPIKey: "my-api-key",
			wantErr:        false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if tc.wantErr {
				if err == nil {
					t.Error("expected an error, but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if apiKey != tc.expectedAPIKey {
				t.Errorf("expected API key to be %s, but got %s", tc.expectedAPIKey, apiKey)
			}
		})
	}
}
