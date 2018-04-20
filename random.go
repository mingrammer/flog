package main

import (
	"github.com/brianvoe/gofakeit"
	"strings"
)

func RandResourceURI() string {
	var url string
	num := gofakeit.Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = gofakeit.BS()
	}
	url += "/" + strings.ToLower(strings.Join(slug, "/"))
	return url
}
