package main

import (
	"context"
	"os"
	"strconv"

	napcat "github.com/zjutjh/napcat-sdk"
	"github.com/zjutjh/napcat-sdk/api"
	"github.com/zjutjh/napcat-sdk/event"
	"github.com/zjutjh/napcat-sdk/message"
)

func main() {
	ctx := context.Background()
	wsURL := os.Getenv("NAPCAT_WS_URL")
	if wsURL == "" {
		wsURL = "ws://127.0.0.1:3001"
	}

	client, err := napcat.DialWebSocket(ctx, wsURL, napcat.WithToken(os.Getenv("NAPCAT_TOKEN")))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for ev := range client.Events() {
		switch e := ev.(type) {
		case *event.PrivateMessage:
			if e.Message.Text() == "/ping" {
				_, _ = client.API().SendPrivateMsg(ctx, api.SendPrivateMsgRequest{
					UserID:  strconv.FormatInt(e.UserID, 10),
					Message: message.Text("pong"),
				})
			}
		}
	}
}
