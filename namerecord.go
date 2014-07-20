package main

import (
	"time"
)

type NameRecord struct {
	ID       int `json:"-"`
	UUID     string
	Username string    `json:"username"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func (r *NameRecord) SetCreated() {
	now := time.Now().UTC()
	r.Created = now
	r.Updated = now
}

func (r *NameRecord) SetUpdated() {
	r.Updated = time.Now().UTC()
}
