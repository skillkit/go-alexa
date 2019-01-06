package alexa

import (
	"encoding/json"
	"net/http"
	"strings"
)

// HandleFunc represents the handler function.
type HandleFunc func(ResponseWriter, *Request) error

// Alexa represents the Alexa app.
type Alexa struct {
	ApplicationID    string
	IgnoreTimestamp  bool
	IgnoreCertVerify bool
	onSessionStarted HandleFunc
	onAudioPlayer    HandleFunc
	onLaunch         HandleFunc
	onIntent         HandleFunc
	onSessionEnded   HandleFunc
}

// Options represents Alexa app options.
type Options struct {
	ApplicationID    string
	IgnoreTimestamp  bool
	IgnoreCertVerify bool
}

// NewApp creates a new Alexa app.
func NewApp(opt *Options) *Alexa {
	return &Alexa{
		ApplicationID:    opt.ApplicationID,
		IgnoreTimestamp:  opt.IgnoreTimestamp,
		IgnoreCertVerify: opt.IgnoreCertVerify,
	}
}

// OnSessionStarted sets the session started handler.
func (a *Alexa) OnSessionStarted(h HandleFunc) {
	a.onSessionStarted = h
}

// OnAudioPlayer sets the audio player handler.
func (a *Alexa) OnAudioPlayer(h HandleFunc) {
	a.onAudioPlayer = h
}

// OnLaunch sets the launch handler.
func (a *Alexa) OnLaunch(h HandleFunc) {
	a.onLaunch = h
}

// OnIntent sets the intent handler.
func (a *Alexa) OnIntent(h HandleFunc) {
	a.onIntent = h
}

// OnSessionEnded sets the session ended handler.
func (a *Alexa) OnSessionEnded(h HandleFunc) {
	a.onSessionEnded = h
}

// Process handles a request passed from Alexa.
func (a *Alexa) Process(r *Request) (*Response, error) {
	w := NewResponse()

	if err := a.verifyApplicationID(r); err != nil {
		return nil, err
	}

	if err := a.verifyTimestamp(r); err != nil {
		return w, err
	}

	if r.Session.New && a.onSessionStarted != nil {
		if err := a.onSessionStarted(w, r); err != nil {
			return nil, err
		}
	}

	switch r.Request.Type {
	case "LaunchRequest":
		if a.onLaunch != nil {
			if err := a.onLaunch(w, r); err != nil {
				return nil, err
			}
		}
	case "IntentRequest":
		if a.onIntent != nil {
			if err := a.onIntent(w, r); err != nil {
				return nil, err
			}
		}
	case "SessionEndedRequest":
		if a.onSessionEnded != nil {
			if err := a.onSessionEnded(w, r); err != nil {
				return nil, err
			}
		}
	default:
		if strings.HasPrefix("AudioPlayer", r.Request.Type) {
			if a.onAudioPlayer != nil {
				if err := a.onAudioPlayer(w, r); err != nil {
					return nil, err
				}
			}
		}
		break
	}

	return w, nil
}

// Handler returns a http handler to hook Alexa app into a http server.
func (a *Alexa) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req *Request

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Bail if POST method.
		if strings.ToUpper(r.Method) != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Verify Alexa request.
		if err := a.verifyAlexaRequest(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Bail if JSON decode failes.
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := a.Process(req)

		// Bail if process request failes.
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(resp)
	})
}
