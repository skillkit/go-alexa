package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/skillkit/go-alexa"
)

func main() {
	lambda.Start(Handler)
}

func Handler(request *alexa.Request) (*alexa.Response, error) {
	app := alexa.NewApp(&alexa.Options{
		ApplicationID: "amzn1.ask.skill.aaa",
	})

	app.OnIntent(func(w alexa.ResponseWriter, r *alexa.Request) error {
		w.OutputSpeech("Hello, world!")
		return nil
	})

	return app.Process(request)
}
