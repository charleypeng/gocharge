package models

import "github.com/charleypeng/gocharge/dtos"

type Login struct {
	UserName string
	Age      string
	Makers   []dtos.Maker
}
