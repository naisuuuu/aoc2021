package aocutil

import (
	"fmt"
	"strconv"
	"strings"
)

func AtoiS(s string, sep string) []int {
	var is []int
	for _, ss := range strings.Split(s, sep) {
		// There are sometimes leading and consecutive spaces in our input.
		// This makes sure we don't append them to our int slice.
		if len(ss) == 0 {
			continue
		}
		si, err := strconv.Atoi(ss)
		if err != nil {
			panic(fmt.Sprintf("Atois('%s','%s'): %v", s, sep, err))
		}
		is = append(is, si)
	}
	return is
}
