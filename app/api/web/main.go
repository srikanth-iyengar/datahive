package web

import (
	"net/http"

	"datahive.io/api/internal/model"
	router "datahive.io/api/internal/router"
	"datahive.io/api/pkg/monitor"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Start() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	model.InitSchema()
	log.Info().Msg("Webserver startup initiated for port 3000")
	monitor.MonitorAll()
	r := router.GetRouter()
	http.Handle("/", r)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
