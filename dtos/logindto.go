package dtos

type LoginDto struct {
	UserName string
	Age      string
	Makers   []Maker
	lover    string
}

type Maker struct {
	Lob  string
	Waka []string
}
