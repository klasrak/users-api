package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func checksum(ds []int64) int64 {
	var s int64
	for i, n := range ds {
		s += n * int64(len(ds)+1-i)
	}
	r := 11 - (s % 11)
	if r == 10 {
		return 0
	}
	return r
}

//IsBrazilianCPFValid checks if a CPF is valid
func IsBrazilianCPFValid(n string) bool {
	u := removeNonDigits(n)

	if len(u) != 11 {
		return false
	}

	ds := make([]int64, 11)
	s := make(map[int64]struct{})

	for i, v := range strings.Split(u, "") {
		c, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return false
		}
		ds[i] = c
		s[c] = struct{}{}
	}

	if len(s) == 1 {
		return false
	}

	return checksum(ds[:9]) == ds[9] && checksum(ds[:10]) == ds[10]
}

//removeNonDigits removes any non-digit from brazilian CPF number
func removeNonDigits(n string) string {
	return regexp.MustCompile(`\D`).ReplaceAllString(n, "")
}
