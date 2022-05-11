package concept

import (
	"reflect"
	"testing"
)

type DummySearcher struct{}

func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: "harry",
		LastName:  "potter",
		Phone:     "00000",
	}
}

func TestFindReturnsError(t *testing.T) {
	phonebook := &Phonebook{}

	want := ErrMissingArgs
	_, err := phonebook.Find(DummySearcher{}, "", "")

	if err != want {
		t.Errorf("Want '%s', got '%s'", want, err)
	}
}

func TestFindReturnsSuccessDummy(t *testing.T) {
	phonebook := &Phonebook{}

	expectPhone := "00000"
	gotPhone, err := phonebook.Find(DummySearcher{}, "harry", "potter")

	if !reflect.DeepEqual(gotPhone, expectPhone){
		t.Errorf("expected (%v), got (%v)", expectPhone, gotPhone)
	}

	if err != nil {
		t.Errorf("Want 'nil', got '%s'", err)
	}
}

