# prime

      Usage
      =====
      A Web Service to return a list of prime numbers from a value passed in as a parameter in the url

      For example  http://your.host.com/primes/segmented/15 will return a JSON document list the primes to 15 -> {"initial":"15","primes":[2,3,5,7,11,13]}
      You can choose the mthod of calculating the Prime numbers ; either the "Sieve of Aitkin" or the "Sieve of Eratosthenes (Segmented)
      To Choose Aitkin the url format is http://your.host.com/primes/aitkin/15
      To Choose Eratosthenes the url format is http://your.host.com/primes/segmented/15
      The output Can also be represented as XML;
      The URL for XML will be http://your.host.com/primes/xml/aitkin/15

      Installation
      ============

      Either ig go is installed

      go run prime.go

      Or

      ./prime
      
