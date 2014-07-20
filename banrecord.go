package main

import (
	time
)

type BanRecord struct {
	ID int `json:"-"`
	UUIDFor string `json:"UUID-for"`
	UUIDBy string `json:"UUID-by"`
	Reason string `json:"reason"`
	Service string `json:"service"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (r *BanRecord) SetCreated() {
	now := time.Now().UTC()
	r.Created = now
	r.Updated = now
}

func (r *BanRecord) SetUpdated() {
	r.Updated = time.Now().UTC()
}
