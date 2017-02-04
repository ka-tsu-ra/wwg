package animals

import (
	"reflect"
	"testing"
)

// params have to be a reference to testing.T
func TestSetHobbiesSetsTheHobbies(t *testing.T) {
	hobbies := []string{"Barking"}
	dog := Dog{}

	dog.SetHobbies(hobbies)

	if !reflect.DeepEqual(dog.GetHobbies(), hobbies) {
		t.Fail()
	}
}
