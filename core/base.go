package core

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/likexian/whois-go"
)

// Foo returns bar. Always.
func Foo() string {
	return "bar"
}

// Random between 0 and 1
func Random() string {
	i := rand.Float64()
	iStr := fmt.Sprintf("%f", i)
	return iStr
}

// WhoIS checks via OS
func WhoIS(ip string) string {
	log.Printf("checking whos is for %v", ip)
	result, err := whois.Whois(ip)
	if err == nil {
		return result
	}
	return "oops"
}
