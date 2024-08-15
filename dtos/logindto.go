// package dtos is a const lib for gocharge dtos
//
// For information, please see readme
package dtos

type LoginDto struct {
	UserName string  `json:"firstName"`
	Age      int     `json:"age,omitempty"`
	Makers   []Maker `json:"makers,omitempty"`
	Lover    string  `json:"lover,omitempty"`
}

type Maker struct {
	Lob  *string
	Waka []string
}
