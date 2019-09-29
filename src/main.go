package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
)

func echoResource(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	v, _ := event.ResourceProperties["Echo"].(string)

	log.Println("Echoing...")
	data = map[string]interface{}{
		"Echo":  v,
		"Echo2": v + "2",
	}

	return
}

func main() {
	lambda.Start(cfn.LambdaWrap(echoResource))
}
