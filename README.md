# GO-TAMBOON ไปทำบุญ

This is a small challenge project to see how good you are with Go. Included in this
repository is a CSV list of Song-pah-pa (ซองผ้าป่า). It is a white envelope with money
inside and the donor name printed on the front. The idea is that your donation amount
should be kept secret les the activity becomes an act of flaunting your wealth.

But we're a payment gateway, we can do better than that. The envelope will contain,
intead, a valid CC and a small piece of paper with the desired donation amount. The entire
list is also encrypted using NSA-proof variant of the [Caesar Cipher][1].

### CONTENTS

* `data/fng.csv.rot255` - A ROT-255 encrypted CSV file.
* `cipher/rot255.go` - Sample ROT-255 encrypt-/decrypter.

### EXERCISE

Write a GO command-line module that, when given the CSV list, calls the [Omise API][0] to
make donations for each row in the file and produce a summary at the end.

Example:

```sh
$ cd $GOPATH/omise/go-tamboon
$ go install -v .

$ $GOPATH/bin/go-tamboon test.csv

performing donations...
done.

        total received: THB  210,000.00
  successfully donated: THB  200,000.00
       faulty donation: THB   10,000.00

    average per person: THB      534.23
             top donor: Obi-wan Kenobi
                        Luke Skywalker
                        Kylo Ren
```

**Requirements:***

* Decrypt the file using a simple [ROT-255][2] algorithm.
* Make donations for each row in the decrypted CSV.
* Produce a summary at the end.
* Handle errors gracefully, without stopping the entire process.

**Bonus:**

* Be a good internet citizen and throttles the API call if we hit rate limit.
* Run as fast as possible on a multi-core CPU.
* Allocate as little memory as possible.
* Complete the entire process without leaving large trace of Credit Card numbers
  in memory.

 [0]: https://www.omise.co/charges-api
 [1]: https://en.wikipedia.org/wiki/Caesar_cipher
 [2]: https://play.golang.org/p/7YMTYOkUmaE
