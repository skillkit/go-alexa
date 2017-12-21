package alexa

import (
	"encoding/json"
	"testing"
)

const (
	requestJSON = `{"session":{"new":false,"sessionId":"amzn1.echo-api.session.leoe","attributes":{},"user":{"userId":"amzn1.ask.account.bbb"},"application":{"applicationId":"amzn1.ask.skill.aaa"}},"version":"1.0","request":{"locale":"en-US","timestamp":"2016-10-27T21:06:28Z","type":"IntentRequest","requestId":"amzn1.echo-api.request.234","intent":{"slots":{"Color":{"name":"Color","value":"blue"}},"name":"MyColorIsIntent"}},"context":{"AudioPlayer":{"playerActivity":"IDLE"},"System":{"device":{"supportedInterfaces":{"AudioPlayer":{}}},"application":{"applicationId":"amzn1.ask.skill.423"},"user":{"userId":"amzn1.ask.account.303"}}}}`
)

func TestApp(t *testing.T) {
	app := NewApp(&Options{
		ApplicationID:   "amzn1.ask.skill.aaa",
		IgnoreTimestamp: true,
	})

	app.OnIntent(func(w ResponseWriter, r *Request) error {
		slot, err := r.Slot("Color")
		if err != nil {
			return err
		}

		if slot.Name != "Color" && slot.Value != "blue" {
			t.Errorf("Expected slot name 'Color' and/or value 'blue' got %s and %s", slot.Name, slot.Value)
		}

		w.OutputSpeech("Hello, world!")
		w.SimpleCard("Test", "Hello")
		return nil
	})

	var r *Request
	if err := json.Unmarshal([]byte(requestJSON), &r); err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	w, err := app.Process(r)
	if err != nil {
		t.Errorf("Expected nil got error: %v", err)
	}

	if w.Response.OutputSpeech.Text != "Hello, world!" {
		t.Errorf("Expected 'Hello, world!' got %s", w.Response.OutputSpeech.Text)
	}

	if w.Response.Card.Title != "Test" || w.Response.Card.Content != "Hello" {
		t.Errorf("Expected card title 'Test' and/or content 'Hello' got %s and %s", w.Response.Card.Title, w.Response.Card.Content)
	}
}
