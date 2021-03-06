package aocutil

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
    si, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return si
}

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

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
