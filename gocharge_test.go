package gocharge_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charleypeng/gocharge"
	"github.com/charleypeng/gocharge/dtos"
)

func TestVersion(t *testing.T) {
	//Version()
	fmt.Println("testing")
	dto := []dtos.LoginDto{
		{
			UserName: "data",
			Age:      432,
		},
		{
			UserName: "mama",
			Age:      42,
		},
	}
	lst := gocharge.NewList(dto)

	result := lst.Where(func(u dtos.LoginDto) bool {
		return strings.Contains(u.UserName, "da")
	})
	fmt.Println(result)
	//myfmt.(&l)
}
