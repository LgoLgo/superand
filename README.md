# LgoRand

High-performance Rand library, far more efficient than the official library. Especially when concurrent.

## Installation

To install this package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed, then you can use the below Go command to install Lgopool.

```sh
go get -u github.com/LgoLgo/LgoRand
```

2. Import it in your code:

```go
import "github.com/LgoLgo/LgoRand"
```

## Example

It's a simple guessing game.

```go
package main

import (
	"fmt"

	"github.com/LgoLgo/LgoRand"
)

func main() {
	maxNum := 100
	rand := LgoRand.New()
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
