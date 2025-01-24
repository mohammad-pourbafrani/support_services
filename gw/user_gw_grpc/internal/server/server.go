package server

import (
	"appconfigs"
	"appstates"
)

func RunServer() {

	var (
		listenAddress               = appconfigs.String("listen-address", "server listen address")
		authenticationServerAddress = appconfigs.String("authentication-server-address", "authentication server address")
	)

	if err := appconfigs.Parse(); err != nil {
		appstates.PanicMissingEnvParams(err.Error())
	} else {
		appstates.DoneServerLaunch("master server " + *listenAddress + "and authentication server " + *authenticationServerAddress + "is run")
	}

}
