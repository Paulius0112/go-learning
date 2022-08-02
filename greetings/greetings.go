package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// If a function name starts with a capital letter can be
// called by a function not in the same package (exported name/function)
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// := is a shortcut for declaring and initialising variable in a single line
	// it uses the value on the right to determine variable's type
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"great to see you, %v",
		"Hail, %v. Well met",
	}

	return formats[rand.Intn(len(formats))]
}