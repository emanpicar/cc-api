package main

import (
	"fmt"
	"net/http"

	"github.com/emanpicar/cc-api/routes"

	"github.com/emanpicar/cc-api/auth"

	"github.com/emanpicar/cc-api/card"
	"github.com/emanpicar/cc-api/logger"
	"github.com/emanpicar/cc-api/luhnalg"
	"github.com/emanpicar/cc-api/settings"
)

func main() {
	logger.Init(settings.GetLogLevel())
	logger.Log.Infoln("Initializing CC API")

	luhnManager := luhnalg.New()
	cardManager := card.New(luhnManager)
	authManager := auth.New()

	logger.Log.Fatal(http.ListenAndServeTLS(
		fmt.Sprintf("%v:%v", settings.GetServerHost(), settings.GetServerPort()),
		settings.GetServerPublicKey(),
		settings.GetServerPrivateKey(),
		routes.New(cardManager, authManager),
	))
}
