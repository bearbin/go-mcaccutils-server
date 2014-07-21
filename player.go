package main

import (
	"errors"
	"strings"
	"time"
)

type Player struct {
	UUID          string
	LastKnownName string       `json:"last-known-name"`
	Skin          string       `json:"-"`
	Names         []NameRecord `db:"-" json:"name-history"`
	Bans          []BanRecord  `db:"-" json:"ban-history"`
	Created       time.Time    `json:"created"`
	Updated       time.Time    `json:"updated"`
}

func (p *Player) PopulateNames() error {
	p.Names = nil
	_, err := dbmap.Select(&(p.Names), "select * from names where UUID=? order by ID", p.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) PopulateBans() error {
	p.Bans = nil
	_, err := dbmap.Select(&(p.Bans), "select * from bans where UUIDFor=? order by ID", p.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) Populate() error {
	err := p.PopulateNames()
	if err != nil {
		return err
	}
	err = p.PopulateBans()
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) SetCreated() {
	now := time.Now().UTC()
	p.Created = now
	p.Updated = now
}

func (p *Player) SetUpdated() {
	p.Updated = time.Now().UTC()
}

func (p *Player) Save() error {
	// Check to see if the player already exists in the database.
	op, err := dbmap.Get(&Player{}, p.UUID)
	if err != nil {
		return err
	}
	if op != nil {
		p.SetUpdated()
		_, err = dbmap.Update(p)
		if err != nil {
			return err
		}
		return nil
	}
	// The player did not exist in the database, insert them.
	p.SetCreated()
	err = dbmap.Insert(p)
	if err != nil {
		return err
	}
	return nil
}

// Gets a player from the database, cache or web service by their name.
// Automatically adds the player back to the cache if they're not already there.
func GetPlayerByName(name string) (*Player, error) {
	name = strings.ToLower(name)
	// Try and get the player from cache.
	player, found := dataCache.Get(name)
	if found {
		return player.(*Player), nil
	}
	// Try and get the user from database.
	var players []Player
	_, err := dbmap.Select(&players, "SELECT * FROM players WHERE LastKnownName=? COLLATE NOCASE", name)
	if err != nil {
		return nil, err
	}
	if len(players) < 1 {
		// Now the player does not exist in database, get them from the HTTP services.
		uuid, newname, err := getUUID(name)
		if err != nil {
			return nil, err
		}
		// Can we get this UUID from the database?
		databasePlayer, err := dbmap.Get(&Player{}, uuid)
		if err != nil {
			return nil, err
		}
		if databasePlayer != nil {
			// The player exists in database, but their lastknownname is different.
			nrecord := &NameRecord{
				UUID:     uuid,
				Username: newname,
			}
			nrecord.Save()
			databasePlayer.(*Player).LastKnownName = newname
			err = databasePlayer.(*Player).Save()
			if err != nil {
				return nil, err
			}
			dataCache.Add(strings.ToLower(newname), databasePlayer.(*Player), 0)
			dataCache.Add(uuid, databasePlayer.(*Player), 0)
			databasePlayer.(*Player).Populate()
			return databasePlayer.(*Player), nil
		}
		// We have to make a new player now.
		// For now, don't bother collecting bans or anything, just make a new player.
		nrecord := &NameRecord{
			UUID:     uuid,
			Username: newname,
		}
		nrecord.Save()
		nnp := &Player{
			UUID:          uuid,
			LastKnownName: newname,
		}
		nnp.Save()
		dataCache.Add(strings.ToLower(nnp.LastKnownName), nnp, 0)
		dataCache.Add(nnp.UUID, nnp, 0)
		nnp.Populate()
		return nnp, nil
	}
	if len(players) > 1 {
		return nil, errors.New("Failed to find unique player from database by name!")
	}
	np := players[0]
	// Add the player back to cache.
	dataCache.Add(strings.ToLower(np.LastKnownName), &np, 0)
	dataCache.Add(np.UUID, &np, 0)
	(&np).Populate()
	return &np, nil
}

// Gets a player from the database or cache by their name.
// Automatically adds the player back to the cache if they're not already there.
func GetPlayerByUUID(UUID string) (*Player, error) {
	// Try and get player from cache.
	player, found := dataCache.Get(UUID)
	if found {
		println("cache found!")
		return player.(*Player), nil
	}
	// Try and get user from database.
	player, err := dbmap.Get(&Player{}, UUID)
	if err != nil {
		return nil, err
	}
	if player == nil {
		return nil, errors.New("Failed to find player from database by UUID!")
	}
	np := player.(*Player)
	np.Populate()
	// Add the player back to cache.
	dataCache.Add(strings.ToLower(np.LastKnownName), np, 0)
	dataCache.Add(np.UUID, np, 0)
	return np, nil
}
