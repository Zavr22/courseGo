package repository

import "github.com/Zavr22/courseGo"

type Authorization interface {
	CreateUser(user courseGo.User) (int, error)
	Login(username, password string) (courseGo.User, error)
}

type Projector interface {
}

type VideoWall interface {
}

type Monitor interface {
}

type Mount interface {
}

type Category interface {
}

type Repository struct {
	Authorization
	VideoWall
	Mount
	Monitor
	Category
}
