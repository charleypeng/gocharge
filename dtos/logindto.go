package dtos

type LoginDto struct {
	UserName *string `json:"firstName"`
	Age      *string `json:"age,omitempty"`
	Makers   []Maker `json:"makers,omitempty"`
	Lover    *string `json:"lover,omitempty"`
}

type Maker struct {
	Lob  *string
	Waka []string
}
