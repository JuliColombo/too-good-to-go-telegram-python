package main

import (
    "fmt"
    "os/exec"
)
import (
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    commandGetCredentials := exec.Command("python3", "../../../get_credentials.py")

    if errGetCredentials := commandGetCredentials.Start(); errGetCredentials != nil {
        fmt.Println("Error get_credentials: ", errGetCredentials)
    }

    return &events.APIGatewayProxyResponse{
        StatusCode:        200,
        Body:              "Get credentials: ok",
      }, nil
}

func main() {
  // Make the handler available for Remote Procedure Call
  lambda.Start(handler)
}
