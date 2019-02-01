package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	var (
		port              string
		skipSslValidation bool
		err               error
	)

	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
	if skipSslValidation, err = strconv.ParseBool(os.Getenv("SKIP_SSL_VALIDATION")); err != nil {
		skipSslValidation = true
	}
	log.SetOutput(os.Stdout)
	log.Printf("Starting on port %s", port)

	tlsValidationMessage := "enabled"
	if skipSslValidation {
		tlsValidationMessage = "disabled"
	}
	log.Printf("Starting with tls validation %s", tlsValidationMessage)

	transport := NewDefaultTransport(skipSslValidation)
	proxy := NewReverseProxy(transport)

	log.Fatal(http.ListenAndServe(":"+port, proxy))
}
