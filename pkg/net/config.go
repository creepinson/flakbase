package net

import (
	"log"
	"net/http"
	"time"

	"github.com/creepinson/flakbase/pkg/store"
	"github.com/gorilla/websocket"
)

// DefaultPort defines the default port for Flakbase.
const DefaultPort = ":9527"

// Config defines the args for a Flakbase server.
type Config struct {
	Host  string
	Rule  string
	Mongo string
}

// Run establishes a http server to handle websocket and rest api.
func Run(config *Config) {
	// initiate the websocket upgrader.
	upgrader := websocket.Upgrader{
		ReadBufferSize:   16384,
		WriteBufferSize:  16384,
		HandshakeTimeout: time.Second * 10,
		CheckOrigin: func(r *http.Request) bool {
			// TODO: check cors
			return true
		},
	}

	// create the datastore handler.
	datastore, err := store.NewHandler(&store.Config{
		Mongo: config.Mongo,
		Rule:  config.Rule,
	})
	if err != nil {
		log.Fatalf("failed to new store handler: %v", err)
	}

	// generate the handler with config.
	s := &handler{
		Config:    config,
		upgrader:  upgrader,
		datastore: datastore,
	}

	// serve the http handler at root.
	http.Handle("/", s)
	if err := http.ListenAndServe(DefaultPort, nil); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
