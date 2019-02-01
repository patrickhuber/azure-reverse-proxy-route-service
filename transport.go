package main

import (
	"crypto/tls"
	"net/http"
)

func NewDefaultTransport(skipTlsValidation bool) http.RoundTripper {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipTlsValidation},
	}
}
