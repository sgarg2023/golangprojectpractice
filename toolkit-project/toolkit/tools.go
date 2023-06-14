package toolkit

import (
	"crypto/rand"
	"fmt"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+-*"

type Tools struct{}

func (t *Tools) randomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	//fmt.Println(s)
	//fmt.Println("\n----------------------------\n")
	//fmt.Println(r)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		fmt.Println(i, p)
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}