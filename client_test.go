package client_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	client "github.com/0ffsideCompass/api-football-go-client"
)

func TestNew(t *testing.T) {
	type args struct {
		key           string
		name          string
		expectedError error
	}

	tests := []args{
		{
			name:          "should return an error if the key is empty",
			key:           "",
			expectedError: errors.New("missing key"),
		},
		{
			name:          "should return a client if the key is not empty",
			key:           "123",
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := client.New(test.key)
			assert.Equal(t, test.expectedError, err)
		})
	}

	fmt.Println(tests)
}
