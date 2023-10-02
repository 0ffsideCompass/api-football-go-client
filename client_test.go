package client_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestNew(t *testing.T) {
	type args struct {
		key           string
		name          string
		client        client.HttpClient
		expectedError error
	}

	tests := []args{
		{
			name:          "should return an error if the key is empty",
			key:           "",
			client:        &http.Client{},
			expectedError: errors.New("missing key"),
		},
		{
			name:          "should return an error if the client is nil",
			key:           "123",
			client:        nil,
			expectedError: errors.New("missing http client"),
		},
		{
			name:          "should return a client if the key is not empty",
			client:        &http.Client{},
			key:           "123",
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := client.New(test.key, test.client)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
