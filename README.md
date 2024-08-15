# gocharge
an out of box toolkit for golang

## usage

```bash
go get github.com/charleypeng/gocharge
```

```go
str := MyString("test string")
//returns a bool to indicate if string is empty or whitespace
str.IsNullOrWhiteSpace()

lst := gocharge.NewList([]string{"few", "rrggfds"})
lst.Add("ewr")
fmt.Println(lst)
fmt.Println(lst.Where(func(s string) bool {
    return strings.Contains(s, "f")
}))
```