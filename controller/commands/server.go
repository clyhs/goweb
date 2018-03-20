package commands

import (
	log "github.com/Sirupsen/logrus"
	"github.com/clyhs/goweb/controller/api"
	"github.com/codegangsta/cli"
)

func CmdServer(c *cli.Context) {

	listenAddr := c.String("listen")

	apiConfig := api.ApiConfig{
		ListenAddr: listenAddr,
	}

	shipyardApi, err := api.NewApi(apiConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := shipyardApi.Run(); err != nil {
		log.Fatal(err)
	}
}
