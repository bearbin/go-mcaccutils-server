package main

import (
	"encoding/json"
	"github.com/zenazn/goji/web"
	"net/http"
)

func ByName(c web.C, w http.ResponseWriter, r *http.Request) {
	player, err := GetPlayerByName(c.URLParams["username"])
	if err != nil {
		panic(err)
	}
	json, err := json.MarshalIndent(player, "", "    ")
	if err != nil {
		panic(err)
	}
	w.Write(json)
}

func ByUUID(c web.C, w http.ResponseWriter, r *http.Request) {
	player, err := GetPlayerByUUID(c.URLParams["uuid"])
	if err != nil {
		panic(err)
	}
	json, err := json.MarshalIndent(player, "", "    ")
	if err != nil {
		panic(err)
	}
	w.Write(json)
}
