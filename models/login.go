// package models is a const lib for gocharge models
//
// For information, please see readme
package models

import "github.com/charleypeng/gocharge/dtos"

type Login struct {
	UserName string
	Age      string
	Makers   []dtos.Maker
}
