package server

import(
	"blog.com/packages/cmd/internal/store";
	"blog.com/packages/cmd/internal/database";
	"blog.com/packages/cmd/internal/conf";
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/rs/zerolog/log"
)

const InternalServerError = "Something went wrong!"

func Start(cfg conf.Config) {

	jwtSetup(cfg)
	store.SetDBConnection(database.NewDBOptions(cfg))
	r := setRouter()

	server := &http.Server{
		Addr:	cfg.Host + ":" + cfg.Port,
		Handler: r,
	}

	//Initializing the server in a goroutine so that
	//it wont clock the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msg("Server ListenAndServe error")
		}
	}()

	//Wait for interrupt signal to gracefully shutdown the server with
	//a time of 5 seconds
	quit := make(chan os.Signal)
	//kill (no pparam)  default send syscall.SIGTERM
	//kill -2 is syscall.SIGINT
	//kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	//the context s used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx ,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")
}
