package main

import (
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func main() {
	helper()
	log.Printf("Hello, world!")
}
