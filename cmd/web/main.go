package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	commonConfig "common/config"

	"web/routing"
)

func main() {
	var (
		configFile = flag.String("config,c", "config.yml", "Path to configuration file to use.")
		errc       = make(chan error)
	)

	if err := commonConfig.SetConfigurationFile(*configFile); err != nil {
		log.Fatal(err)
	}

	config, err := commonConfig.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	router, err := routing.NewRouter(config)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    config.HTTP.Address(false),
		Handler: router,
	}

	go func() {
		if config.HTTP.IsSecure && config.HTTP.SecureHTTP != nil {
			keyFile := config.HTTP.SecureHTTP.KeyFilePath
			certFile := config.HTTP.SecureHTTP.CertFilePath
			if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
				errc <- err
			}
		} else {
			if err := server.ListenAndServe(); err != nil {
				errc <- err
			}
		}
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		errc <- fmt.Errorf("%v", <-c)
	}()

	if err := <-errc; err != nil {
		log.Printf("error: %v", err)
		if err := server.Close(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
		os.Exit(1)
	}
}
