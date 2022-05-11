package concept

import (
"errors"
)

var (
	ErrMissingArgs   = errors.New("FirstName and LastName are mandatory arguments")
	ErrNoPersonFound = errors.New("No person found")
)

type Searcher interface {
	Search(people []*Person, firstName string, lastName string) *Person
}

func (p *Phonebook) Search(people []*Person, firstName string, lastName string) *Person {
	for _, ps := range people {
		if ps.FirstName == firstName && ps.LastName == lastName {
			return ps
		}
	}

	return &Person{}
}

type Person struct {
	FirstName string
	LastName  string
	Phone     string
}

type Phonebook struct {
	People []*Person
}

func (p *Phonebook) Find(searcher Searcher, firstName, lastName string) (string, error) {
	if firstName == "" || lastName == "" {
		return "", ErrMissingArgs
	}

	person := searcher.Search(p.People, firstName, lastName)

	if person == nil {
		return "", ErrNoPersonFound
	}

	return person.Phone, nil
}

/*
Source : https://ieftimov.com/posts/testing-in-go-test-doubles-by-example/

https://martinfowler.com/bliki/TestDouble.html
- Dummy objects are passed around but never actually used. Usually they are just used to fill parameter lists.
- Fake objects actually have working implementations, but usually take some shortcut which makes them not suitable for production (an InMemoryTestDatabase is a good example).
- Stubs provide canned answers to calls made during the test, usually not responding at all to anything outside what's programmed in for the test.
- Spies are stubs that also record some information based on how they were called. One form of this might be an email service that records how many messages it was sent.
- Mocks are pre-programmed with expectations which form a specification of the calls they are expected to receive. They can throw an exception if they receive a call they don't expect and are checked during verification to ensure they got all the calls they were expecting.
*/