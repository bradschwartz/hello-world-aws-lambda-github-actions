package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Name string `json:"name"`
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event Event) (string, error) {
	return fmt.Sprintf("Received request from: %s", event.Name), nil
}
