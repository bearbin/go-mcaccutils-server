package main

import (
	"github.com/PuerkitoBio/throttled"
	"github.com/PuerkitoBio/throttled/store"
	"github.com/coopernurse/gorp"
	"github.com/pmylund/go-cache"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"math/rand"
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
	dbmap = initDatabase()
	// Initialise the cache.
	dataCache = cache.New(6*time.Hour, 20*time.Second)
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
	goji.Handle("/by/*", api)
	api.Use(PlainJSON)
	st := store.NewMemStore(1000)
	api.Use(throttled.RateLimit(throttled.PerHour(config.RateLimitHour), &throttled.VaryBy{Custom: CustomKeyGenerator}, st).Throttle)
	api.Handle("/by/uuid/:uuid", ByUUID)
	api.Handle("/by/name/:username", ByName)

	goji.Get("/", SingleHandler(config.IndexLocation))

	// Serve using the magic of Goji!
	goji.Serve()
}
