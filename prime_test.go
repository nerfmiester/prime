package main

import (
	"fmt"
	"testing"
)

func TestWorkerAitkin(t *testing.T) {

	resp := workerAitkin(15)
	if resp.Initial != "15" {
		t.Error("Invalid response code for prime number expected 15 but got", resp.Initial)
	}

	if resp.Primes[3] != 7 {
		t.Error("Invalid prime number expected 7 but got", resp.Primes[3])
	}

	sumPrime := uint64(0)
	for _, x := range resp.Primes {
		sumPrime += x
	}

	if sumPrime != 41 {
		t.Error("Invalid sum of prime numbers expected 41 but got", sumPrime)
	}

	if sumPrime < 0 {
		t.Error("Invalid sum of prime numbers, must be a positive integer, but got ", sumPrime)
	}

}

func TestWorkerSegmented(t *testing.T) {

	resp := workerSegmented(15)
	if resp.Initial != "15" {
		t.Error("Invalid response code for prime number expected 15 but got", resp.Initial)
	}

	if resp.Primes[3] != 7 {
		t.Error("Invalid prime number expected 7 but got", resp.Primes[3])
	}

	sumPrime := uint64(0)
	for _, x := range resp.Primes {
		sumPrime += x
	}

	if sumPrime != 41 {
		t.Error("Invalid sum of prime numbers expected 41 but got", sumPrime)
	}

	if sumPrime < 0 {
		t.Error("Invalid sum of prime numbers, must be a positive integer, but got ", sumPrime)
	}
}
func TestUsage(t *testing.T) {
	usage()
}

func TestLoadCache(t *testing.T) {

	mapToPrimes = map[uint64]Primers{}

	loadCache(15)

	fmt.Println("Length of mapToPrimes should be 15")

	if len(mapToPrimes) != 15 {
		t.Error("The mapToPrimes slice should be 15 long but was", len(mapToPrimes))
	}

	fmt.Println("The last value should also be a list of values below 15")

	resp := mapToPrimes[15]

	if resp.Initial != "15" {
		t.Error("Invalid response code for prime number expected 15 but got", resp.Initial)
	}

	if resp.Primes[3] != 7 {
		t.Error("Invalid prime number expected 7 but got", resp.Primes[3])
	}

	sumPrime := uint64(0)
	for _, x := range resp.Primes {
		sumPrime += x
	}

	if sumPrime != 41 {
		t.Error("Invalid sum of prime numbers expected 41 but got", sumPrime)
	}

	if sumPrime < 0 {
		t.Error("Invalid sum of prime numbers, must be a positive integer, but got ", sumPrime)
	}

}

func TestMain(t *testing.T) {
	usageBool = true
	main()
}
