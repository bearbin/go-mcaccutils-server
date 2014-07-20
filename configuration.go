package main

import (
	"flag"
)

type configuration struct {
	DataLocation  string
	RateLimitHour int
	RateLimitDay  int
}

func (c *configuration) Populate() {
	flag.StringVar(&c.DataPath, "data-path", "data/", "the path for storing persistent data.")
	flag.IntVar(&c.RateLimitDay, "rate-limit-day", "10000", "the number of requests allowed per IP day")
	flag.IntVar(&c.RateLimitHour, "rate-limit-hour", "1000", "the number of requests allowed per IP hour")
}
