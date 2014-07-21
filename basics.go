package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type mojangIDResponse struct {
	Name string `json: name`
}

// Name produces the most recent username for a player of the given UUID.
func getName(uuid string) (username string, err error) {
	uuid = strings.Replace(uuid, "-", "", -1)
	// Fetch the account info API for this player UUID.
	resp, err := http.Get("https://sessionserver.mojang.com/session/minecraft/profile/" + uuid)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Read out the body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// Decode the JSON
	decResp := mojangIDResponse{}
	err = json.Unmarshal(body, &decResp)
	if err != nil {
		return "", err
	}
	// Return the decoded name.
	return decResp.Name, nil
}

type mojangNameResponse struct {
	Profiles []mojangNameResponseProfile `json:"profiles"`
	Count    int                         `json:"size"`
}

type mojangNameResponseProfile struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}

// UUID takes the player name and returns the UUID of that player.
// It returns a UUID which may or may not contain dashes (-).
func getUUID(username string) (uuid string, name string, err error) {
	// Hit the API and wait for a response.
	reqBody := strings.NewReader("{\"name\":\"" + username + "\", \"agent\": \"minecraft\"}")
	resp, err := http.Post("https://api.mojang.com/profiles/page/1", "application/json", reqBody)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	// Read out the body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	// Decode the JSON
	decResp := mojangNameResponse{}
	err = json.Unmarshal(body, &decResp)
	if err != nil {
		return "", "", err
	}
	// Make sure the lookup was a success on fishbans' side.
	if decResp.Count < 1 {
		return "", "", errors.New("No Name Found")
	}
	if decResp.Count > 1 {
		return "", "", errors.New("Ambiguous Query")
	}
	return decResp.Profiles[0].UUID, decResp.Profiles[0].Name, nil
}
