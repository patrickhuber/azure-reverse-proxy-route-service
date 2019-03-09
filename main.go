package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	DEFAULT_PORT              = "8080"
	X_FORWARDED_HOST          = "X-Forwarded-Host"
	X_ORIGINAL_HOST           = "X-Original-Host"
	CF_FORWARDED_URL_HEADER   = "X-Cf-Forwarded-Url"
	CF_PROXY_SIGNATURE_HEADER = "X-Cf-Proxy-Signature"
)

func main() {
	var (
		port              string
		err               error
	)

	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}
		
	skipSslValidation, _ := strconv.ParseBool(os.Getenv("SKIP_SSL_VALIDATION"))
		
	log.SetOutput(os.Stdout)
	log.Printf("Starting on port %s", port)

	tlsValidationMessage := "enabled"
	if skipSslValidation {
		tlsValidationMessage = "disabled"
	}
	log.Printf("Starting with tls validation %s", tlsValidationMessage)

	director := func(req *http.Request) {
		originalHost := req.Header.Get("X-Original-Host")
		if strings.TrimSpace(originalHost) != "" {
			req.Header.Add("X-Forwarded-Host", originalHost)
		}

		forwardedURL := req.Header.Get(CF_FORWARDED_URL_HEADER)
		// Note that url.Parse is decoding any url-encoded characters.
		url, err := url.Parse(forwardedURL)
		if err != nil {
			log.Fatalln(err.Error())
		}
		req.URL = url
		req.Host = url.Host
	}

	proxy := &httputil.ReverseProxy{Director: director}

	log.Fatal(http.ListenAndServe(":"+port, proxy))
}
