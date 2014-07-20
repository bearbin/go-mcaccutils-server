package main

import (
	"github.com/coopernurse/gorp"
	"github.com/pmylund/go-cache"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	config    = &configuration{}
	dbmap     *gorp.DbMap
	dataCache *cache.Cache
)

func init() {
	// Seed the RNG. Only needs doing once at startup.
	rand.Seed(time.Now().UTC().UnixNano())
	// Populate the configuration.
	config.Populate()
	// Initialise the database mappings.
	dbmap = databaseInit()
	// Initialise the cache.
	dataCache = cache.New(1*time.Hour, 20*time.Second)
}

func main() {
	// Start name collection process, and other things of that nature.

	// Default Middleware
	goji.Use(middleware.EnvInit)
	goji.Use(middleware.Recoverer)
	goji.Use(GZIPHandler)

	// Set up API middleware.
	// TODO: Set up killing the people who don't want json.
	api := web.New()
	goji.Handle("/*", api)
	api.Use(PlainJSON)
	api.Handle("/uuid/:uuid", ByUUID)
	api.Handle("/name/:username", ByName)

	goji.Get("/*", http.StripPrefix("/", http.FileServer(http.Dir("static/"))))

	// Serve using the magic of Goji!
	goji.Serve()
}
