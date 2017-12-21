package alexa

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrEmptyApplicationID        = errors.New("Application ID should not be empty")
	ErrEmptyRequestApplicationID = errors.New("Request Application ID should not be empty")
	ErrNotSameApplicationID      = errors.New("Application ID and Request Application ID is not the same")
	ErrInvalidTimestamp          = errors.New("Invalid timestamp")
	ErrChainURL                  = errors.New("Unable to find certificate chain header")
	ErrBadChainURL               = errors.New("Chain URL is not valid")
	ErrDownloadCert              = errors.New("Certificate could not be downloaded")
	ErrReadCert                  = errors.New("Could not read certification content")
	ErrParsePem                  = errors.New("Failed to parse certificate PEM")
	ErrBadCertDate               = errors.New("Certificate is not valid")
)

// readCert reads a certification.
func readCert(url string) ([]byte, error) {
	cert, err := http.Get(url)
	if err != nil {
		return []byte{}, ErrDownloadCert
	}

	defer cert.Body.Close()
	c, err := ioutil.ReadAll(cert.Body)

	if err != nil {
		return []byte{}, ErrReadCert
	}

	return c, nil
}

// verifyAlexaRequest verifies that the request actual is from Amazon and a Alexa.
func (a *Alexa) verifyAlexaRequest(r *http.Request) error {
	if a.IgnoreCertVerify {
		return nil
	}

	chainURL := r.Header.Get("SignatureCertChainUrl")

	if chainURL == "" {
		return ErrChainURL
	}

	u, err := url.Parse(chainURL)
	if err != nil {
		return err
	}

	if u.Scheme != "https" || u.Host != "s3.amazonaws.com" || !strings.HasPrefix(u.Path, "/echo.api/") {
		return ErrBadChainURL
	}

	buf, err := readCert(chainURL)
	if err != nil {
		return err
	}

	block, rem := pem.Decode(buf)
	if block == nil {
		return ErrParsePem
	}

	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(rem)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	if time.Now().Unix() < cert.NotBefore.Unix() || time.Now().Unix() > cert.NotAfter.Unix() {
		return ErrBadCertDate
	}

	opts := x509.VerifyOptions{
		DNSName: "echo-api.amazon.com",
		Roots:   roots,
	}

	_, err = cert.Verify(opts)
	return err
}

// verifyApplicationId verifies that the configured Application ID is the same as in the payload.
func (a *Alexa) verifyApplicationID(r *Request) error {
	if len(a.ApplicationID) == 0 {
		return ErrEmptyApplicationID
	}

	if len(r.Session.Application.ApplicationID) == 0 {
		return ErrEmptyRequestApplicationID
	}

	if a.ApplicationID != r.Session.Application.ApplicationID {
		return ErrNotSameApplicationID
	}

	return nil
}

// verifyTimestamp compares the request timestamp to the current timestamp
// and returns an error if they are too far apart.
func (a *Alexa) verifyTimestamp(r *Request) error {
	if a.IgnoreTimestamp {
		return nil
	}

	rt, err := time.Parse(time.RFC3339, r.Request.Timestamp)
	if err != nil {
		return err
	}

	if time.Since(rt) < time.Duration(150)*time.Second {
		return nil
	}

	return ErrInvalidTimestamp
}
