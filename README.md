# contextio-go
[Context.IO](https://context.io/) API Golang Library

This library is currently in BETA, and as such we make no promises; the use of this library is provided without warranty.

## Installation

```bash
# For the LITE api
go get github.com/contextio/contextio-go/ciolite

# 2.0 api coming soon...
```

## CIO Lite Usage
```go
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/contextio/contextio-go/ciolite"
)

func main() {
	// Key and Secret
	cioKey := os.Getenv("CONTEXTIO_API_KEY")
	cioSecret := os.Getenv("CONTEXTIO_API_SECRET")

	// Client Instance
	cioLiteClient := ciolite.NewCioLite(cioKey, cioSecret)
	// Can also use with a standard or custom logger:
	// ciolite.NewCioLiteWithLogger(cioKey, cioSecret, logrus.StandardLogger())

	// Discovery Call Parameters
	discoveryParams := ciolite.GetDiscoveryParams{Email: "test@gmail.com"}

	// Actual Discovery Call
	discoveryResp, err := cioLiteClient.GetDiscovery(discoveryParams)
	if err != nil {
		log.Fatal("Error calling ContextIO: " + err.Error())
	}

	// Responses are simple structs, all fields accessible. The following line prints:
	// {Email:test@gmail.com Type:gmail Documentation:[] Found:true
	// IMAP:{Server:imap.gmail.com Username:test@gmail.com UseSSL:true OAuth:true Port:993}}
	fmt.Printf("%+v", discoveryResp)

	// Get a slice of users
	users, _ := cioLiteClient.GetUsers(ciolite.GetUsersParams{})

	// Get a slice of emails in the Inbox of the first users's first email account
	fmt.Println(cioLiteClient.GetUserEmailAccountsFolderMessages(
		users[0].ID,
		users[0].EmailAccounts[0].Label,
		"Inbox",
		ciolite.GetUserEmailAccountsFolderMessageParams{},
	))
}
```

## Support
If you want to open an issue or PR for this library - go ahead! We'd love to hear your feedback.

For API support please consult our [support site](http://support.context.io) and feel free to drop a line to [support@context.io](mailto:support@context.io).
