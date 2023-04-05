package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Language interface {
	LanguageName() string
	Greet(name string) string
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

type Portuguese struct{}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

func SayHello(name string, language Language) string {
	prefixText := "I can speak "
	return fmt.Sprintf(prefixText+"%s: %s", language.LanguageName(), language.Greet(name))
}
