package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Salaton/generics/reqres"
)

// The function below only takes integers as the arguments --> Convert it to generic
// Sum returns the sum of the provided arguments.
// func Sum(args ...int) int {
// 	var sum int
// 	for i := 0; i < len(args); i++ {
// 		sum += args[i]
// 	}
// 	return sum
// }

// The generic version (The args can be any type really)
// [T int] is how we represent the generic
// ---> T is the symbol used to represent the generic type
// ---> int is the constraint that shows the concrete type that can be used

// In this example `Sum[T int](value ...T) T`, T is constrained to be only an int and therefore other
// types wont work

// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | ~int64 |
		float32 | float64 |
		complex64 | complex128
}

func Sum[T Numeric](value ...T) T {
	var sum T
	for _, val := range value {
		sum += val
	}
	return sum
}

type id int64

func main() {
	fmt.Println(Sum([]int{1, 2, 3}...))
	fmt.Println(Sum([]int8{1, 2, 3}...))
	fmt.Println(Sum([]id{1, 2, 3}...))

	ctx := context.Background()
	timeout := 30 * time.Second

	requestContext, _ := context.WithTimeout(ctx, timeout)
	m, err := reqres.Get[reqres.RequestObject](requestContext, "https://reqres.in/api/users?page=2")
	if err != nil {
		log.Fatal(err)
	}

	user := reqres.User{
		Name: "Salaton Doe",
		Job:  "Software Engineer",
	}
	newUser, err := reqres.Post[reqres.User](requestContext, "https://reqres.in/api/users", user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newUser)
	fmt.Println(m.Data)
}
