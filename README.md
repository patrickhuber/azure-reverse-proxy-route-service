# Azure Reverse Proxy Route Service

A cloud foundry route service that acts as a simple reverse proxy for applications in cloud foundry.

## Route Service Overview

This example route service uses the new headers/features that have been added to the GoRouter. For example:

- `X-CF-Forwarded-Url`: A header that contains the original URL that the GoRouter received.
- `X-CF-Proxy-Signature`: A header that the GoRouter uses to determine if a request has gone through the route service.
- `X-ORIGINAL-HOST`: a header injected by azure application gateway
- `X-FORWARDED-HOST`: a header injected by this route service which contains the same value as `X-ORIGINAL-HOST`

## Getting Started

- Download this repository and `cf push` to your chosen CF deployment.
- Push your app which will be associated with the route service.
- Create a user-provided route service ([see docs](http://docs.cloudfoundry.org/services/route-services.html#user-provided))
- Bind the route service to the route (domain/hostname)
- Tail the logs of this route service in order to verify that requests to your app go through the route service. 

## Environment Variables

### SKIP_SSL_VALIDATION

If you set this environment variable to false, the route service
will validate SSL certificates. By default the route service skips SSL validation.

Example:

```sh
cf set-env logging-route-service SKIP_SSL_VALIDATION false
cf restart logging-route-service
```

### PORT

The default port for starting this application (this is set by cloud foundry)