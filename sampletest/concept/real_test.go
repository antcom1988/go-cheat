package concept

import (
	"reflect"
	"testing"
)
/*
Test real implementation
*/
func TestFindReturnsSuccessReal(t *testing.T) {
	phonebook := &Phonebook{
		People: []*Person{
			{
				FirstName: "harry",
				LastName:  "potter",
				Phone:     "1111",
			},
			{
				FirstName: "marry",
				LastName:  "potter",
				Phone:     "00000",
			},
		},
	}

	expectPhone := "1111"
	gotPhone, err := phonebook.Find(&Phonebook{}, "harry", "potter")

	if !reflect.DeepEqual(gotPhone, expectPhone){
		t.Errorf("expected (%v), got (%v)", expectPhone, gotPhone)
	}

	if err != nil {
		t.Errorf("Want 'nil', got '%s'", err)
	}
}
