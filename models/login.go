package models

import "gocharge/dtos"

type Login struct {
	UserName string
	Age      string
	Makers   []dtos.Maker
}
