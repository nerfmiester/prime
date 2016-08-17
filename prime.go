package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/caleblloyd/primesieve"
	"github.com/gorilla/mux"
)

var (
	x, y, n, intPrime uint64
	z                 string
	algorithm         string
	writer            io.Writer
	primeSlice        = []uint64{}
	wg                = &sync.WaitGroup{}
	keeps             = make([]uint64, 5000000)
	usageBool         bool
	ok                bool
	usetoml           bool
	mapToPrimes       = map[uint64]Primers{}
)

// Primers is a list of Primes
type Primers struct {
	Initial string   `json:"initial"`
	Primes  []uint64 `json:"primes"`
}

func init() {

	flag.BoolVar(&usageBool, "u", false, "Show the usage parameters.") //#3
}

// Filter Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan uint64, out chan<- uint64, prime uint64) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

const sizeToCache = 5000000

func main() {

	flag.Parse()

	if usageBool {
		usage()
		os.Exit(0)
	}
	fmt.Println("-->>Server Starting.....<<--")
	r := mux.NewRouter()
	r.HandleFunc("/primes/{algorithm}/{prime}", PrimeHandler)
	r.HandleFunc("/primes/xml/{algorithm}/{prime}", PrimeXMLHandler)
	// Preload array with up to 5 million in background
	go loadCache(sizeToCache)
	http.ListenAndServe(":8081", r)

}

// PrimeHandler = the function will Orchistrate prime number creation and return JSON.
func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	val := Primers{}
	err := errors.New("")
	prime := mux.Vars(r)["prime"]
	algorithm := mux.Vars(r)["algorithm"]
	if intPrime, err = strconv.ParseUint(prime, 10, 64); err != nil {
		fmt.Println(err)
	}
	if intPrime <= sizeToCache {
		if val, ok = mapToPrimes[intPrime]; ok {
		} else {
			if algorithm == "segmented" {
				val = workerSegmented(intPrime)
			} else {
				val = workerAitkin(intPrime)
			}
		}
	} else {
		val = workerAitkin(intPrime)
	}
	j, err := json.Marshal(val)

	if err != nil {
		fmt.Println(err)
	}
	w.Write([]byte(j))

}

// PrimeXMLHandler = the function will Orchistrate prime number creation and return XML notation.
func PrimeXMLHandler(w http.ResponseWriter, r *http.Request) {
	val := Primers{}
	err := errors.New("")
	prime := mux.Vars(r)["prime"]
	if intPrime, err = strconv.ParseUint(prime, 10, 64); err != nil {
		fmt.Println(err)
	}
	if intPrime <= sizeToCache {
		if val, ok = mapToPrimes[intPrime]; ok {
		} else {
			val = workerAitkin(intPrime)
		}
	} else {
		val = workerAitkin(intPrime)
	}
	fmt.Println("Implement XML", val)
	// TODO : Implement XML code here

}

func loadCache(size uint64) {
	for i := uint64(1); i <= size; i++ {
		mapToPrimes[i] = workerAitkin(i)
	}
}

func workerSegmented(toPrime uint64) Primers {
	primers := Primers{}
	primers.Initial = strconv.FormatUint(toPrime, 10)
	primers.Primes = primesieve.ListMax(uint64(toPrime))
	return primers
}

func workerAitkin(toPrime uint64) Primers {

	// Many thanks to gofool for this implementation
	// https://raw.githubusercontent.com/agis-/gofool/master/atkin.go

	var x, y, n uint64
	nsqrt := math.Sqrt(float64(toPrime))

	isPrime := make([]bool, (sizeToCache + 1))
	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= toPrime && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= toPrime && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= toPrime && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < toPrime; y += n * n {
				isPrime[y] = false
			}
		}
	}
	//fmt.Println("len of isPrime = ", len(isPrime))
	isPrime[2] = true
	isPrime[3] = true

	primes := make([]uint64, 0, 1270606)
	for x = 0; x < uint64(len(isPrime))-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}

	// primes is now a slice that contains all the
	// primes numbers up to isPrime

	primers := Primers{}
	primers.Initial = strconv.FormatUint(toPrime, 10)
	primers.Primes = primes

	return primers

}

func usage() {

	fmt.Printf("\n\tUsage\n\t=====\n\ta webservice to return a list of prime numbers from a value passed in as a parameter in the url\n\n\tFor example  http://your.host.com/primes/15 will return a JSON document  \n\tTodo : Add functionality to select output method\t")

}
