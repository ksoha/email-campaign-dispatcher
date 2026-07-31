package main

import (
	"fmt"
	"log"

	"github.com/ksoha/email-dispatcher/internal/auth"
)

func main() {
	hash, err := auth.HashPassword("lazylad@00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hashed password:", hash)
}
