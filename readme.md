# alexa [![Build Status](https://travis-ci.org/skillkit/go-alexa.svg?branch=master)](https://travis-ci.org/skillkit/go-alexa) [![GoDoc](https://godoc.org/github.com/skillkit/go-alexa?status.svg)](https://godoc.org/github.com/skillkit/go-alexa) [![Go Report Card](https://goreportcard.com/badge/github.com/skillkit/go-alexa)](https://goreportcard.com/report/github.com/skillkit/go-alexa)

Go package for creating Amazon Alexa skills. Not a hundred percent package for all Alexa features. The most basic features exists.

## Installation

```
go get -u github.com/skillkit/go-alexa
```

## Example with HTTP

```go
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
```

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)
