package services

import (
	"context"

	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
)

func getAccesToken(ctx context.Context, app *firebase.App) string {
	// get token dari flutter app
	client, err := app.Auth(ctx)

	if err != nil {
		fmt.Printf("error getting auth client: %v\n", err)
		os.Exit(1)
	}

	token, err := client.CustomToken(ctx, "your token")

	if err != nil {
		fmt.Printf("error getting auth client: %v\n", err)
		os.Exit(1)
	}

	return token

}

// func SendMessageTo(ctx context.Context, token string, title string, body string, data map[string]string) error {

// }
