package main

import "sync"

var (
	check checkJson
)

var (
	authorizedTokens []string
	clubNames        []string
)

var (
	WaitGroup = sync.WaitGroup{}
)

var (
	tokenCounter int
	requestDelay int
	requests     int
	taken        int
	errors       int
	rateLimits   int
	reserved     int
)
