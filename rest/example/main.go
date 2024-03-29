package main

import (
	"context"
	"fmt"

	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/auth/secretsmanagerauth"
	"github.com/SKF/go-utility/v2/stages"

	hierarchy "github.com/SKF/go-hierarchy-v2-client/rest"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	dd_http "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

const (
	awsProfile  = "example_profile"
	clientID    = "2e51739d-61e8-4c34-bdbc-de0cd4b53dfe" //random uuid in example
	serviceName = "example-service"
)

var awsRegion = "eu-west-1"

func main() {
	ctx := context.Background()

	sess, err := session.NewSessionWithOptions(
		session.Options{
			Config:  aws.Config{Region: aws.String(awsRegion)},
			Profile: awsProfile,
		},
	)
	if err != nil {
		panic(err)
	}

	client := hierarchy.NewClient(
		hierarchy.WithStage(stages.StageSandbox),
		hierarchy.WithClientID(clientID),
		client.WithDatadogTracing(dd_http.RTWithServiceName(serviceName)),
		client.WithTokenProvider(&auth.SecretsManagerTokenProvider{
			Config: secretsmanagerauth.Config{
				WithDatadogTracing:       true,
				ServiceName:              serviceName,
				AWSSession:               sess,
				AWSSecretsManagerAccount: "633888256817",
				AWSSecretsManagerRegion:  awsRegion,
				SecretKey:                "user-credentials/" + serviceName,
				Stage:                    stages.StageSandbox,
			},
		}),
	)

	ancestors, err := client.GetAncestors(ctx, "a2b0645b-62e0-43c7-9d35-1f02fd245f67", 0, "company", "site")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ancestors)
}
