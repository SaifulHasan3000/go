package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s : %s", g.LanguageName(), g.Greet(name))
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf(`Ciao %s!`, name)
}

type Portugese struct{}

func (i Portugese) LanguageName() string {
	return "Portugese"
}

func (i Portugese) Greet(name string) string {
	return fmt.Sprintf(`Ola %s!`, name)
}
