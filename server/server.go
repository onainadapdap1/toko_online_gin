package server

import (
	"fmt"
	"toko_online_gin/driver"
	"toko_online_gin/router"

	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
	g "github.com/incubus8/go/pkg/gin"
)

func init() {
	gotenv.Load()
}

func StartServer() {
	addr := driver.Config.ServiceHost + ":" + driver.Config.ServicePort
	fmt.Println(driver.Config.ServicePort,"<<<<<<")
	conf := g.Config{
		ListenAddr: addr,
		Handler: router.Router(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at "+ addr)
		},
	}

	g.Run(conf)
}