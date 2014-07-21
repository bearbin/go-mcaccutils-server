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

func (r *NameRecord) Save() error {
	// Check to see if the record already exists in the database.
	or, err := dbmap.Get(&NameRecord{}, r.UUID, r.Username)
	if err != nil {
		return err
	}
	if or != nil {
		r.SetUpdated()
		_, err = dbmap.Update(r)
		if err != nil {
			return err
		}
		return nil
	}
	// The record did not exist in the database, insert it.
	r.SetCreated()
	err = dbmap.Insert(r)
	if err != nil {
		return err
	}
	return nil
}
