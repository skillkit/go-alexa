package main

import (
	"net/http"

	"github.com/skillkit/go-alexa"
)

func main() {
	app := alexa.NewApp(&alexa.Options{
		ApplicationID: "amzn1.ask.skill.aaa",
	})

	app.OnIntent(func(w alexa.ResponseWriter, r *alexa.Request) error {
		w.OutputSpeech("Hello, world!")
		return nil
	})

	http.Handle("/", app.Handler())
	http.ListenAndServe(":3000", nil)
}
