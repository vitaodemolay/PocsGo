package main

import (
	"fmt"
	"reflect"
)

const tagName = "validate_tst"

type User struct {
	Id    int    `json:"id" validate_tst:"-"`
	Name  string `json:"name" validate_tst:"required"`
	Email string `json:"email" validate_tst:"email"`
}

func main() {
	// Example usage
	user := User{
		Id:    1,
		Name:  "John Doe",
		Email: "john@test.com",
	}

	t := reflect.TypeOf(user)

	fmt.Println("Type: ", t.Name())
	fmt.Println("Kind: ", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(tagName)

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
