package main

import (
	"math/rand"
	net_url "net/url"
	"strings"

	"github.com/brianvoe/gofakeit"
)

// RandResourceURI generates a random resource URI
func RandResourceURI() string {
	var url string
	num := gofakeit.Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = net_url.QueryEscape(gofakeit.BS())
	}
	url += "/" + strings.ToLower(strings.Join(slug, "/"))
	return url
}

// RandAuthUserID generates a random auth user id
func RandAuthUserID() string {
	candidates := []string{"-", strings.ToLower(gofakeit.Username())}
	return candidates[rand.Intn(2)]
}
