package main

import (
	"encoding/json"

	"github.com/apex/go-apex"
	"github.com/skillkit/go-alexa"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var request *alexa.Request

		if err := json.Unmarshal(event, &request); err != nil {
			return nil, err
		}

		app := alexa.NewApp(&alexa.Options{
			ApplicationID: "amzn1.ask.skill.aaa",
		})

		app.OnIntent(func(w alexa.ResponseWriter, r *alexa.Request) error {
			w.OutputSpeech("Hello, world!")
			return nil
		})

		return app.Process(request)
	})
}
