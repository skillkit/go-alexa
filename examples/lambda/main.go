package main

import (
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/skillkit/go-alexa"
)

func main() {

}

func Handle(request *alexa.Request, ctx *runtime.Context) (interface{}, error) {
	app := alexa.NewApp(&alexa.Options{
		ApplicationID: "amzn1.ask.skill.aaa",
	})

	app.OnIntent(func(w alexa.ResponseWriter, r *alexa.Request) error {
		w.OutputSpeech("Hello, world!")
		return nil
	})

	return app.Process(request)
}
