package main

import (
	"flag"
)

type configuration struct {
	DataLocation  string
	IndexLocation string
	RateLimitHour int
}

func (c *configuration) Populate() {
	flag.StringVar(&c.DataLocation, "data-path", "data/", "the path for storing persistent data.")
	flag.StringVar(&c.IndexLocation, "index-path", "index.html", "the path for the index to serve. useful to provide information about your server.")
	flag.IntVar(&c.RateLimitHour, "rate-limit-hour", 100, "the number of requests allowed per IP hour")
	flag.Parse()
}
