package main

import (
	"log"
	"testing"
)

func TestStandings(t *testing.T) {
	r := GetWinners("./test-comps.yaml")
	log.Printf("%v", r)
	s := CalcStandings(r)
	log.Printf("%v", s)
}
