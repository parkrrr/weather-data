package main_test

import (
	"context"
	//"errors"
	"testing"

	"github.com/parkrrr/weather-data/lambda/get-observations"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: `{ "location":"KORD", "datetime": "2011-10-05T14:48:00.000Z" }`},
			expect:  "KORD",
			err:     nil,
		},
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{Body: `{ "datetime": "2011-10-05T14:48:00.000Z" }`},
			expect:  "KIND",
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handle(context.TODO(), test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
