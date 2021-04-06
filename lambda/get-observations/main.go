package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type WeatherRequest struct {
	Location string    `json:"location,omitempty"`
	Time     time.Time `json:"datetime"`
}

func main() {
	lambda.Start(Handle)
}

func Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	body, err := parseRequest(request.Body)
	if err != nil {
		panic(err)
	}

	if body.Location == "" {
		body.Location = "KIND"
	}

	log.Printf("Using location: %s", body.Location)

	return events.APIGatewayProxyResponse{Body: body.Location, StatusCode: 201}, nil
}

func parseRequest(body string) (WeatherRequest, error) {
	var data WeatherRequest
	err := json.Unmarshal([]byte(body), &data)
	return data, err
}
