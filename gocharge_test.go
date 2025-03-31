package gocharge_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/charleypeng/gocharge"
	"github.com/charleypeng/gocharge/dtos"
	"github.com/google/uuid"
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
	for i := 0; i < 1000000; i++ {
		id := uuid.New()
		lst.Add(dtos.LoginDto{UserName: id.String(), Age: i})
	}

	t2 := time.Now()
	// result2 := gocharge.WhereT(&lst.Items, func(u dtos.LoginDto) bool {
	// 	return strings.Contains(u.UserName, "ma") || strings.Contains(u.UserName, "0b")
	// })
	result := lst.Where(func(u dtos.LoginDto) bool {
		return strings.Contains(u.UserName, "ma") || strings.Contains(u.UserName, "9e0")
	})
	fmt.Println("====data test===============\n", "used:", time.Since(t2))
	fmt.Println(result)
	//fmt.Println(result2)
	//myfmt.(&l)
	lsd := gocharge.MyString("hello world")
	fmt.Println("test for mystring", lsd.CompareTo("world"))
}
