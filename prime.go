package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	x, y, n, intPrime int
	z                 string
	algorithm         string
	primeSlice        = []int{}
	wg                = &sync.WaitGroup{}
	keeps             = make([]int, 5000000)
	usageBool         bool
	mapToPrimes       = map[int]Primers{}
)

// Primers is a list of Primes
type Primers struct {
	Initial string `json:"initial"`
	Primes  []int  `json:"primes"`
}

func init() {
	fmt.Println("-->>Init<<--")
	flag.StringVar(&algorithm, "algorithm", "a", "Which algorithm to Use Spanish language.\n\n\t\t a = \"aitkin\"")
	flag.BoolVar(&usageBool, "u", false, "Show the usage parameters.") //#3
}

// Filter Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

const sizeToCache = 5000000

func main() {

	if usageBool {
		usage()
		os.Exit(0)
	}

	r := mux.NewRouter()
	r.HandleFunc("/primes/{prime}", PrimeHandler)
	fmt.Println("Allright geezer")
	// Preload array with up to 5 million in background
	go loadCache(sizeToCache)
	http.ListenAndServe(":8081", r)

}

// PrimeHandler = the function will Orchistrate prime number creation.
func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	val := Primers{}
	err := errors.New("")
	prime := mux.Vars(r)["prime"]
	if intPrime, err = strconv.Atoi(prime); err != nil {
		fmt.Println(err)
	}
	if intPrime <= sizeToCache {
		val = mapToPrimes[intPrime]
		fmt.Println("In cache")
		fmt.Println("Length of mapToPrimes = ", len(mapToPrimes))
	} else {
		val = workerAitkin(intPrime, false)
	}
	j, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Printf("-->>Prime = %s <<--", j)

	w.Write([]byte(j))

}

func loadCache(size int) {
	for i := 1; i <= size; i++ {
		//fmt.Println("i=", i)
		mapToPrimes[i] = workerAitkin(i, false)
	}
}

func workerAitkin(toPrime int, save bool) Primers {

	var x, y, n int
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

	primes := make([]int, 0, 1270606)
	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}

	// primes is now a slice that contains all the
	// primes numbers up to isPrime

	// if in save mode then keep output

	if save {

		keeps = primes
		return Primers{}

	}

	primers := Primers{}
	primers.Initial = strconv.Itoa(toPrime)
	primers.Primes = primes

	return primers

}

func usage() {

	fmt.Printf("\n\tUsage\n\t=====\n\ta webservice to return a list of prime numbers from a value passed in as a parameter in the url\n\n\tFor example  http://your.host.com/primes/15 will return a JSON document  \n\tTodo : Add functionality to select output method\t")

}
