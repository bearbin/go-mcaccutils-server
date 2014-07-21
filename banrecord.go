package main

import (
	"time"
)

type BanRecord struct {
	ID         int       `json:"-"`
	UUIDFor    string    `json:"UUID-for"`
	UUIDBy     string    `json:"UUID-by"`
	Reason     string    `json:"reason"`
	ServerName string    `json:"server-name"`
	ServerIP   string    `json:"server-ip"`
	Service    string    `json:"service"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

func (r *BanRecord) SetCreated() {
	now := time.Now().UTC()
	r.Created = now
	r.Updated = now
}

func (r *BanRecord) SetUpdated() {
	r.Updated = time.Now().UTC()
}

func (r *BanRecord) Save() error {
	// Check to see if the record already exists in the database.
	or, err := dbmap.Get(&BanRecord{}, r.ID)
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
