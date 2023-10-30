package gocharge

import (
	"fmt"
	"github.com/devfeel/mapper"
	"gocharge/dtos"
	"gocharge/models"
	"gocharge/myfmt"
	"gocharge/types"
)

func Version() {
	fmt.Printf("v0.0.1\n")

	m := mapper.NewMapper()

	loginDto := dtos.LoginDto{UserName: "fwjke", Age: "33", Makers: []dtos.Maker{{Lob: "wdf"}, {Lob: "24ffe"}}}

	loginModel := models.Login{}

	m.Mapper(&loginDto, &loginModel)

	fmt.Println(loginModel)

	fmt.Println("---------end--------")

	var ws = types.MyString("     ")

	fmt.Println(ws.IsNullOrEmpty(true))

	var sw = myfmt.IsNullOrEmpty("      k ")
	fmt.Println(sw)
	myfmt.LogError(nil)
}
