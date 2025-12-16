package main

import (
	"3x-ui-aggregator/Api"
	"3x-ui-aggregator/Api/Requests"
	"3x-ui-aggregator/Utils"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

func collectSubscriptions() string {
	fullSubscriptionText := ""

	for _, client := range Api.ClientStorage {
		loginReq := &Requests.LoginRequest{
			LoginBody: Requests.LoginPostBody{
				Username: client.Username,
				Password: client.Password,
			},
		}

		resp, err := loginReq.SendRequest(client)
		if err != nil {
			panic(err)
		}

		if !strings.Contains(string(resp), "true") {
			panic(resp)
		}
		getReq := &Requests.GetInboundRequest{
			ID: "1",
		}

		resp, err = getReq.SendRequest(client)
		if err != nil {
			panic(err)
		}

		subId, err := getReq.GetSubID(resp)
		if err != nil {
			panic(err)
		}

		subReq := &Requests.GetSubscriptionRequest{
			SubscriptionId: subId,
		}

		resp, err = subReq.SendRequest(client)
		if err != nil {
			panic(err)
		}

		decodedString, err := Utils.DecodeBase64(string(resp))
		if err != nil {
			panic(err)
		}

		fullSubscriptionText += decodedString
	}

	return Utils.EncodeBase64(fullSubscriptionText)
}

func main() {

	clients, err := Api.LoadClientsFromFile("client.cfg")
	if err != nil {
		return
	}

	Api.ClientStorage = clients

	http.HandleFunc("/free_vpn", HandleSubscription)

	certPath := filepath.Join("/root", "cert", "clawclouds.cc")

	err = http.ListenAndServeTLS(":443", certPath+"/fullchain.pem", certPath+"/privkey.pem", nil)
	if err != nil {
		panic(err)
	}
}

func HandleSubscription(w http.ResponseWriter, r *http.Request) {
	subscriptions := collectSubscriptions()

	_, err := fmt.Fprintf(w, subscriptions)
	if err != nil {
		return
	}
}
