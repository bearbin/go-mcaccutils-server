package main

import (
	"time"
)

type Player struct {
	UUID          string
	LastKnownName string       `json:"last-known-name"`
	Skin          string       `json:"-"`
	Names         []NameRecord `db:"-" json:"name-history"`
	Bans          []BanRecord  `db:"-" json"ban-history"`
	Created       time.Time    `json:"created"`
	Updated       time.Time    `json:"updated"`
}

func (p *Player) PopulateNames() error {
	_, err := dbmap.Select(&(p.Names), "select * from names where UUID=? order by ID", p.UUID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Player) PopulateBans() error {
	_, err := dbmap.Select(&(p.Bans), "select * from bans where UUID=? order by ID", p.UUID)
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
