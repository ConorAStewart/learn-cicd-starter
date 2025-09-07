package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	validHeaderValue := "ApiKey asdfgghgfsaf123445436"
	validAPIKey := "asdfgghgfsaf123445436"
	tests := []struct {
		name    string
		key     string
		value   string
		apiKey  string
		wantErr bool
	}{
		{
			name:    "Test functionality",
			key:     "Authorization",
			value:   validHeaderValue,
			apiKey:  validAPIKey,
			wantErr: false,
		},
		{
			name:    "Missing Key",
			key:     "",
			value:   "",
			apiKey:  "",
			wantErr: true,
		},
		{
			name:    "Wrong prefix",
			key:     "Authorization",
			value:   "Bearer " + validAPIKey,
			apiKey:  "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := http.Header{}
			if tt.key != "" {
				header.Set(tt.key, tt.value)
			}
			returned_key, err := GetAPIKey(header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if returned_key != tt.apiKey {
				t.Errorf("GetAPIKey() error, returned key = %v, actual key = %v", returned_key, tt.apiKey)
			}

		})
	}
}
