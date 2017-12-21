package alexa

import "encoding/json"

// ResponseWriter represents the interface that handle a response to Alexa.
type ResponseWriter interface {
	EndSession(bool)
	LinkAccountCard()
	OutputSpeech(string)
	OutputSpeechSSML(string)
	RepromptSpeech(string)
	RepromptSSML(string)
	SimpleCard(string, string)
	StandardCard(string, string, string, string)
	String() (string, error)
}

// Response represents the top level response object to Alexa.
type Response struct {
	Response          *ResponseObject        `json:"response"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Version           string                 `json:"version"`
}

// NewResponse creates a new default response.
func NewResponse() *Response {
	return &Response{
		Version: "1.0.0",
		Response: &ResponseObject{
			ShouldSessionEnd: true,
		},
		SessionAttributes: make(map[string]interface{}),
	}
}

// ResponseObject represents the response object in the response to Alexa.
type ResponseObject struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       *[]Directive  `json:"directives,omitempty"`
	ShouldSessionEnd bool          `json:"shouldEndSession"`
}

// OutputSpeech represents the output speech object in the response to Alexa.
type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

// Card represents the card object in the response to Alexa.
type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Text    string `json:"text,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

// Image represents the image object in the response to alexa.
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// Reprompt represents the reprompt object in the response to Alexa.
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

// Directive represents the directive object in the response to Alexa.
type Directive struct {
	Type         string `json:"type"`
	PlayBehavior string `json:"playBehavior,omitempty"`
	AudioItem    *struct {
		Stream *Stream `json:"stream,omitempty"`
	} `json:"audioItem,omitempty"`
}

// Stream represents the stream object in the response to Alexa.
type Stream struct {
	Token                string `json:"token"`
	URL                  string `json:"url"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}

// EndSession will set the should session end value.
func (r *Response) EndSession(end bool) {
	r.Response.ShouldSessionEnd = end
}

// LinkAccountCard will set the response card as a link account card.
func (r *Response) LinkAccountCard() {
	r.Response.Card = &Card{
		Type: "LinkAccount",
	}
}

// OutputSpeech will set the output speech as plain text.
func (r *Response) OutputSpeech(content string) {
	r.Response.OutputSpeech = &OutputSpeech{
		Type: "PlainText",
		Text: content,
	}
}

// OutputSpeechSSML will set the output speech as SSML.
func (r *Response) OutputSpeechSSML(content string) {
	r.Response.OutputSpeech = &OutputSpeech{
		Type: "SSML",
		SSML: content,
	}
}

// RepromptSpeech will set the reprompt output speech as plain text.
func (r *Response) RepromptSpeech(content string) {
	r.Response.Reprompt = &Reprompt{
		OutputSpeech: &OutputSpeech{
			Type: "PlainText",
			Text: content,
		},
	}
}

// RepromptSSML will set the reprompt output speech as SSML.
func (r *Response) RepromptSSML(content string) {
	r.Response.Reprompt = &Reprompt{
		OutputSpeech: &OutputSpeech{
			Type: "SSML",
			SSML: content,
		},
	}
}

// SimpleCard will set the response card as a simple card.
func (r *Response) SimpleCard(title, content string) {
	r.Response.Card = &Card{
		Type:    "Simple",
		Title:   title,
		Content: content,
	}
}

// StandardCard will set the response card as a standard card with images.
func (r *Response) StandardCard(title, content, smallImageURL, largeImageURL string) {
	r.Response.Card = &Card{
		Type:    "Standard",
		Title:   title,
		Content: content,
		Image: &Image{
			LargeImageURL: largeImageURL,
			SmallImageURL: smallImageURL,
		},
	}
}

// String will return the response as JSON.
func (r *Response) String() (string, error) {
	buf, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
