package main

import (
	"encoding/json"
	"github.com/bearbin/go-mcaccutils"
	"github.com/zenazn/goji/web"
	"net/http"
)

func ByName(c web.C, w http.ResponseWriter, r *http.Request) {
	username := c.URLParams["username"]
	// Try and get the user from cache.
	player, found := dataCache.Get("username")
	if found {
		json, err := json.Marshal(player.(*Player))
		if err != nil {
			panic(err)
		}
		w.Write(json)
		return
	}
	// Try and get the user from database.
	player, err := dbmap.Get(&Player{}, username)
	if err != nil {
		panic(err)
	}
	// Check if the player was found.
	if player != nil {
		json, err := json.Marshal(player.(*Player))
		if err != nil {
			panic(err)
		}
		w.Write(json)
		return
	}
	// The player wasn't found, let's see if we can find him ourselves.
	uuid, err := mcaccutils.UUID(username)
	if err != nil {
		panic(err)
	}
	// For now, don't bother collecting bans or anything, just make a new player.
	nrecord := NameRecord{
		UUID: uuid,
		Name: username,
	}
	nrecord.SetCreated()
	dbmap.Insert(nrecord)
	player = Player{
		UUID:          uuid,
		LastKnownName: username,
	}
	player.SetCreated()
	dbmap.Insert(player)
	json, err := json.Marshal(player)
	if err != nil {
		panic(err)
	}
	w.Write(json)
	// TODO - add player to cache.
}
