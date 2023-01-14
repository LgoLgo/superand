![superand](img/superand.png)

High-performance Rand library, far more efficient than the official library. Especially when concurrent.

## Installation

To install this package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed, then you can use the below Go command to install superand.

```sh
go get -u github.com/LgoLgo/superand
```

2. Import it in your code:

```go
import "github.com/LgoLgo/superand"
```

## Example

It's a simple guessing game.

```go
package main

import (
	"fmt"

	"github.com/LgoLgo/superand"
)

func main() {
	maxNum := 100
	rand := superand.New()
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	for {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
```

## Benchmark

```go
BenchmarkStandardRand
BenchmarkStandardRand-8                 	79792984	        14.71 ns/op

Benchmarksuperand
Benchmarksuperand-8                      	54015529	        20.42 ns/op

BenchmarkUnsafesuperand
BenchmarkUnsafesuperand-8                	275500960	        4.288 ns/op

BenchmarkStandardRandWithConcurrent
BenchmarkStandardRandWithConcurrent-8   	13186624	        91.46 ns/op

BenchmarksuperandWithConcurrent
BenchmarksuperandWithConcurrent-8        	309059797	        3.238 ns/op
```

## License

This project is under the Apache License 2.0. See the LICENSE file for the full license text.