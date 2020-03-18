# webhooksite

> Go library for [webhook.site](https://webhook.site/)

[![Build Status](https://github.com/rubenv/webhooksite/workflows/Test/badge.svg)](https://github.com/rubenv/webhooksite/actions) [![GoDoc](https://godoc.org/github.com/rubenv/webhooksite?status.png)](https://godoc.org/github.com/rubenv/webhooksite)

Simple library to test webhook calls, ideal for testing callbacks in a CI
server that doesn't accept calls from the outside world.

Features:

* Create tokens, inspect results
* Can be configured to use a self-hosted webhook.site instance

## Usage

```go
// Create a client
c := webhooksite.New()

// Fetch a token
token, err := c.CreateToken()

// Do a webhook call
// curl https://webhook.site/{{ token.UUID }}

// Fetch the requests
reqs, err := c.GetRequests(token.UUID)
```

## License

This library is distributed under the [MIT](LICENSE) license.
