# prime

      Usage
      =====
      A Web Service to return a list of prime numbers from a value passed in as a parameter in the url

      For example  http://your.host.com/primes/segmented/15 will return a JSON document list the primes to 15 -> {"initial":"15","primes":[2,3,5,7,11,13]}
      You can choose the method of calculating the Prime numbers ; either the "Sieve of Aitkin" or the "Sieve of Eratosthenes (Segmented)
      To Choose Aitkin the url format is http://your.host.com/primes/aitkin/15
      To Choose Eratosthenes the url format is http://your.host.com/primes/segmented/15
      The output Can also be represented as XML;
      The URL for XML will be http://your.host.com/primes/xml/aitkin/15

      Build - this will only work if you have go installed. (Tested on go version go1.7rc1 darwin/amd64)
      =====

      Clone repository.

      git clone git@github.com:nerfmiester/prime.git

      cd prime

      Choose your OS (It has only been tested on MAC)

      for linux
      GOOS="linux" go build prime.go

      for windows
      GOOS="windows" go build prime.go


      To run the tests

      go tests -v -cover

      Execution
      ============

      Either if you have go installed

      go run prime.go

      Or execute binary
      _________________________

      linux
      Check permissions
      If the execution bit is not set

      chmod +x linuxprime

      ./linuxprime
      _________________________

      Mac

      ./prime
      _________________________
