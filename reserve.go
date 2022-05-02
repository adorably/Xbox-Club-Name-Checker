package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

func reserve(clubName string, token string) {
	check.ClubName = clubName

	b, _ := json.Marshal(check)
	req := fasthttp.AcquireRequest()
	req.SetBody(b)
	req.Header.SetMethod("POST")
	req.Header.Set("Authorization", token)
	req.Header.Set("X-XBL-Contract-Version", "1")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Accept", "application/json")
	req.SetRequestURI("https://clubaccounts.xboxlive.com/clubs/reserve")

	res := fasthttp.AcquireResponse()

	start := time.Now()

	if err := fasthttp.Do(req, res); err != nil {
		errors++
	}

	fasthttp.ReleaseRequest(req)

	if res.StatusCode() == 201 || res.StatusCode() == 200 {
		duration := time.Since(start)
		reserved++
		fmt.Printf("\r %s[%s*%s] Requests: (%d) | Reserved: (%d) | Taken: (%d) | Rate Limits: (%d) | Errors: (%d) %s", DEFAULT, CYAN, DEFAULT, requests, reserved, taken, rateLimits, errors, FLUSH)
		fmt.Printf("\n\n %s[%s+%s] Reserved Club Name '%s%s%s' (Took %d.%dms) %s\n\n", DEFAULT, CYAN, DEFAULT, CYAN, clubName, DEFAULT, duration.Milliseconds(), duration.Microseconds(), FLUSH)
		err := writeLines("results/reserved.txt", clubName)
		if err != nil {
			fmt.Printf(" %s[%s!%s] Failed to write results to file %s\n", DEFAULT, RED, DEFAULT, FLUSH)
		}
	} else if res.StatusCode() == 409 {
		taken++
		fmt.Printf("\r %s[%s*%s] Requests: (%d) | Reserved: (%d) | Taken: (%d) | Rate Limits: (%d) | Errors: (%d) %s", DEFAULT, CYAN, DEFAULT, requests, reserved, taken, rateLimits, errors, FLUSH)
	} else if res.StatusCode() == 429 {
		rateLimits++
		fmt.Printf("\r %s[%s*%s] Requests: (%d) | Reserved: (%d) | Taken: (%d) | Rate Limits: (%d) | Errors: (%d) %s", DEFAULT, CYAN, DEFAULT, requests, reserved, taken, rateLimits, errors, FLUSH)
	} else {
		errors++
		fmt.Printf("\r %s[%s*%s] Requests: (%d) | Reserved: (%d) | Taken: (%d) | Rate Limits: (%d) | Errors: (%d) %s", DEFAULT, CYAN, DEFAULT, requests, reserved, taken, rateLimits, errors, FLUSH)
	}
}
