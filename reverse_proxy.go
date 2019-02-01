package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	X_FORWARDED_HOST          = "X-Forwarded-Host"
	X_ORIGINAL_HOST           = "X-Original-Host"
	CF_FORWARDED_URL_HEADER   = "X-Cf-Forwarded-Url"
	CF_PROXY_SIGNATURE_HEADER = "X-Cf-Proxy-Signature"
)

func NewReverseProxy(transport http.RoundTripper) http.Handler {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			forwardedURL := req.Header.Get(CF_FORWARDED_URL_HEADER)

			originalHost := req.Header.Get(X_ORIGINAL_HOST)
			if originalHost != "" {
				req.Header.Set(X_FORWARDED_HOST, originalHost)
			}

			// Note that url.Parse is decoding any url-encoded characters.
			url, err := url.Parse(forwardedURL)
			if err != nil {
				log.Fatalln(err.Error())
			}

			req.URL = url
		},
		Transport: transport,
	}
}
