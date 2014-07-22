package main

import (
	"flag"
)

type configuration struct {
	DataLocation  string
	RateLimitHour int
}

func (c *configuration) Populate() {
	flag.StringVar(&c.DataLocation, "data-path", "data/", "the path for storing persistent data.")
	flag.IntVar(&c.RateLimitHour, "rate-limit-hour", 100, "the number of requests allowed per IP hour")
}
