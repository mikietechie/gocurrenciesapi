package main

import (
	"log"

	"github.com/google/uuid"
)

func main() {
	str := uuid.NewString()
	log.Println(str)
	log.Println(len(str))

}
