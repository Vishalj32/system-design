package models

import "github.com/google/uuid"

type Player struct {
	Name string
	Id   string
}

func NewPlayer(name string) Player {
	return Player{
		Name: name,
		Id:   uuid.New().String(),
	}
}
