//	Author: Jean Haley
// 	Date: 2021-10-13
// 	Purpose: This program will take a domain name as input and return the number of days until the domain expires.

package main

import (
	"fmt"
	"strings"
	"time"

	whois "github.com/undiabler/golang-whois"
)

func main() {
	whoisStr, err := whois.GetWhois("google.com")
	if err != nil {
		panic(err)
	}

	lineList := strings.Split(whoisStr, "\n")

	for _, s := range lineList {
		term := strings.TrimSpace(strings.Split(s, ":")[0])
		if strings.Contains(term, "Registry Expiry Date") {
			// grabs expiry date as string
			expString := strings.TrimSpace(strings.Split(s, ":")[1]) + ":00:00"
			// converts string to time.Time
			expTime, err := timeConv(expString)
			if err != nil {
				panic(err)
			}
			// print days until expiry
			fmt.Printf("Time until Expiration: %.02f Days\n", time.Until(expTime).Hours()/24)

		}
	}
}

// timeConv converts string to time.Time
func timeConv(s string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, s)
	if err != nil {
		return t, err
	}
	return t, nil
}
