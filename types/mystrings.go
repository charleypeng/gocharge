package types

import "strings"

type MyStrings []string

func GetMyStrings(myStr *[]string) *MyStrings {
	if myStr == nil {
		return nil
	}
	var m MyStrings
	for _, s := range *myStr {
		m = append(m, s)
	}
	return &m
}

func (s *MyStrings) ToStringLine(seperator ...string) string {
	if s == nil {
		return ""
	}

	sep := ""

	if len(seperator) == 1 {
		sep = seperator[0]
	}
	var builder strings.Builder
	strs := *s
	len := len(strs)
	for i := 0; i < len; i++ {
		if i+1 == len {
			builder.WriteString(strs[i])
		} else {
			builder.WriteString(strs[i] + sep)
		}
	}
	return builder.String()
}
