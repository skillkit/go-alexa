package alexa

import (
	"errors"
)

var (
	ErrNoSlotFound = errors.New("No slot found")
)

// Request represents the payload object from Alexa.
type Request struct {
	Context struct {
		AudioPlayer struct {
			OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
			PlayerActivity       string `json:"playerActivity"`
			Token                string `json:"token"`
		} `json:"AudioPlayer"`
		System struct {
			APIAccessToken string `json:"apiAccessToken"`
			APIEndpoint    string `json:"apiEndpoint"`
			Application    struct {
				ApplicationID string `json:"applicationId"`
			} `json:"application"`
			Device struct {
				DeviceID            string `json:"deviceId"`
				SupportedInterfaces struct {
					AudioPlayer struct{} `json:"AudioPlayer"`
				} `json:"supportedInterfaces"`
			} `json:"device"`
			User struct {
				AccessToken string `json:"accessToken"`
				Permissions struct {
					ConsentToken string `json:"consentToken"`
				} `json:"permissions"`
				UserID string `json:"userId"`
			} `json:"user"`
		} `json:"System"`
	} `json:"context"`
	Session struct {
		New        bool   `json:"new"`
		SessionID  string `json:"sessionId"`
		Attributes struct {
		} `json:"attributes"`
		User struct {
			AccessToken string `json:"accessToken"`
			UserID      string `json:"userId"`
		} `json:"user"`
		Application struct {
			ApplicationID string `json:"applicationId"`
		} `json:"application"`
	} `json:"session"`
	Request struct {
		Locale    string  `json:"locale"`
		Timestamp string  `json:"timestamp"`
		Type      string  `json:"type"`
		RequestID string  `json:"requestId"`
		Intent    *Intent `json:"intent"`
	} `json:"request"`
	Version string `json:"version"`
}

// Intent represents a intent object from Alexa.
type Intent struct {
	Name  string           `json:"name"`
	Slots map[string]*Slot `json:"slots"`
}

// Slot represents a slot object from Alexa.
type Slot struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// IntentName returns the intent name.
func (r *Request) IntentName() string {
	return r.Request.Intent.Name
}

// SessionID returns the session id.
func (r *Request) SessionID() string {
	return r.Session.SessionID
}

// Slot will find a slot by name if it exists or return a error.
func (r *Request) Slot(name string) (*Slot, error) {
	slot, ok := r.Request.Intent.Slots[name]
	if !ok {
		return nil, ErrNoSlotFound
	}

	return slot, nil
}

// UserID returns the user id.
func (r *Request) UserID() string {
	return r.Session.User.UserID
}
