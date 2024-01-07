package reflect_test

import (
	"reflect"
	"testing"
)

type Dog struct {
	name   string
	height int
	weight int
}

func TestStruct(t *testing.T) {
	dog1 := Dog{
		name:   "tama",
		height: 30,
		weight: 10,
	}

	dog2 := Dog{
		name:   "tama",
		height: 30,
		weight: 10,
	}

	if !reflect.DeepEqual(dog1, dog2) {
		t.Errorf("want %v got %v", dog1, dog2)
	}

}
