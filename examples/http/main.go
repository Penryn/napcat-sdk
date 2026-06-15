package main

import (
	"context"
	"fmt"
	"os"

	napcat "github.com/zjutjh/napcat-sdk"
	"github.com/zjutjh/napcat-sdk/api"
	"github.com/zjutjh/napcat-sdk/message"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("NAPCAT_TOKEN")
	userID := os.Getenv("NAPCAT_TARGET_USER_ID")
	if userID == "" {
		panic("请设置 NAPCAT_TARGET_USER_ID")
	}

	client, err := napcat.NewHTTPClient("http://127.0.0.1:3000", napcat.WithToken(token))
	if err != nil {
		panic(err)
	}

	login, err := client.API().GetLoginInfo(ctx, api.GetLoginInfoRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("当前账号:", login.UserID)

	_, err = client.API().SendPrivateMsg(ctx, api.SendPrivateMsgRequest{
		UserID:  userID,
		Message: message.Text("来自 NapCat Go SDK 的 HTTP 消息"),
	})
	if err != nil {
		panic(err)
	}
}
