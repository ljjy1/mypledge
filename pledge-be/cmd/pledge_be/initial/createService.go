package initial

import (
	"strconv"

	"pledge-be/internal/config"
	"pledge-be/internal/schedule"
	"pledge-be/internal/server"

	"github.com/go-dev-frame/sponge/pkg/app"
)

// CreateServices create http service
func CreateServices() []app.IServer {
	var cfg = config.Get()
	var servers []app.IServer

	// create a http service
	httpAddr := ":" + strconv.Itoa(cfg.HTTP.Port)
	httpServer := server.NewHTTPServer(httpAddr,
		server.WithHTTPIsProd(cfg.App.Env == "prod"),
		server.WithHTTPTLS(cfg.HTTP.TLS),
	)
	servers = append(servers, httpServer)

	// create a schedule service (asynq scheduler + server)
	servers = append(servers, schedule.NewScheduleServer())

	return servers
}
