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
  commandMain := exec.Command("python3", "../../../main.py")
     if errMain := commandMain.Run(); errMain != nil {
        fmt.Println("Error main: ", errMain)
    }

    return &events.APIGatewayProxyResponse{
        StatusCode:        200,
        Body:              "Check new offers: ok",
      }, nil
}

func main() {
  // Make the handler available for Remote Procedure Call
  lambda.Start(handler)
}
