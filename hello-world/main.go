package main

import (
	"context"
	"fmt"
    "encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


func helloHandler(username string) (events.ALBTargetGroupResponse, error) {
    res, err := json.Marshal(username)

    if err != nil {
        return events.ALBTargetGroupResponse {
        StatusCode: 500,
            Body: fmt.Sprintf("%s", err),
        }, err
    }

    return events.ALBTargetGroupResponse {
    Headers: map[string]string{
        "content-type": "application/json",
    },
    Body: fmt.Sprintf("Hello %s", string(res)),
    }, nil
}

func errorHandler() (events.ALBTargetGroupResponse, error) {
    return events.ALBTargetGroupResponse {
    StatusCode: 500,
    Headers: map[string]string{
        "content-type": "text/plain; charset=utf-8",
    },
    Body: fmt.Sprintf("500 internal server error\n"),
    }, nil
}

func notFoundHandler() (events.ALBTargetGroupResponse, error) {
    return events.ALBTargetGroupResponse {
    StatusCode: 404,
    Headers: map[string]string{
        "content-type": "text/plain; charset=utf-8",
    },
    Body: fmt.Sprintf("404 not found\n"),
    }, nil
}

func handleRequest(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {

	fmt.Printf("Processing request data for traceId %s.\n", request.Headers["x-amzn-trace-id"])
	fmt.Printf("Body size = %d.\n", len(request.Body))
	fmt.Printf("context = %s.\n", ctx)
    username := request.QueryStringParameters["username"]

    for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}
    switch request.Path {
        case "/hello": return helloHandler(username)
        case "/fuga": return errorHandler()
    }
    return notFoundHandler()
    
}


func Handler(ctx context.Context, request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	// If no name is provided in the HTTP request body, throw an error

	return events.ALBTargetGroupResponse{Body: request.Body, StatusCode: 200, StatusDescription: "200 OK", IsBase64Encoded: false, Headers: map[string]string{}}, nil
}

func main() {
	lambda.Start(handleRequest)
}