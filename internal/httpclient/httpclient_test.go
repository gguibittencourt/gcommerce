package httpclient

import (
	"context"
	"github.com/gguibittencourt/gcommerce/internal/httpclient/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type postTestCase struct {
	name              string
	url               string
	body              any
	expected          Response
	expectedErr       error
	requesterResponse *http.Response
	requesterErr      error
}

func TestClient_Post(t *testing.T) {
	tests := []postTestCase{
		{
			name: "should return error when body is not a valid json",
			url:  "http://localhost:8080/123",
			body: nil,
			expected: Response{
				StatusCode: 0,
				Result:     nil,
			},
			expectedErr:       nil,
			requesterResponse: nil,
			requesterErr:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			requester := setupHTTPClientMock(t, test)
			client := NewClient(requester)
			result, err := client.Post(context.Background(), test.url, test.body)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func setupHTTPClientMock(t *testing.T, testCase postTestCase) *mocks.Requester {
	client := mocks.NewRequester(t)
	client.EXPECT().
		Post(testCase.url, "application/json", mock.Anything).
		Return(testCase.requesterResponse, testCase.requesterErr)
	return client
}
