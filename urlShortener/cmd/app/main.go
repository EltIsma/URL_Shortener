package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"urlShortener/urlShortener/config"
	"urlShortener/urlShortener/internal/adapters"
	"urlShortener/urlShortener/internal/adapters/localrepo"
	"urlShortener/urlShortener/internal/adapters/pgrepo"
	"urlShortener/urlShortener/internal/app/usecase"
	gohttp "urlShortener/urlShortener/internal/ports/goHTTP"
	"urlShortener/urlShortener/internal/server"
	"urlShortener/urlShortener/pkg/postgres"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.InfoLevel)
	logger.SetFormatter(&log.TextFormatter{})
	config, err := config.Parse()
	if err != nil {
		log.Fatalf("could parse: %s", err)
	}
	log.Println(config)
	useDatabase := flag.Bool("d", false, "Use PostgreSQL to store data")
	flag.Parse()
    var rep adapters.URLRepo
	if *useDatabase {
		dbConn, err := postgres.NewDatabase(config.DbConnString)
		if err != nil {
			log.Fatalf("could not initialize database connection: %s", err)
		}
		rep = pgrepo.New(dbConn.GetDB())
	} else {
		rep = localrepo.New()
	}

	serv := usecase.New(rep)
	handler := gohttp.New(serv)
	sigQuit := make(chan os.Signal, 2)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg, _ := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		s := <-sigQuit
		return fmt.Errorf("captured signal: %v", s)
	})
	server := new(server.Server)
	go func() {
		if err := server.New(gohttp.NewRouter(handler), config); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	fmt.Println("Server Started")
	if err := eg.Wait(); err != nil {
		logger.Infof("gracefully shutting down the server: %v", err)
	}

	if err := server.Shutdown(context.Background()); err != nil {
		_ = fmt.Errorf("error occured on server shutting down: %s", err.Error())
	}

}
