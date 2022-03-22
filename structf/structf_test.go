package structf_test

import (
	"go2/structf"
	"reflect"
	"testing"
)

func TestPrime(t *testing.T) {
	have := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
		Slice       []int
		Object      struct {
			NestedField int
		}
	}{
		FieldString: "stroka",
		FieldInt:    107,
		Slice:       []int{112, 107, 207},
	}
	want := struct {
		FieldString string `json:"field_string"`
		FieldInt    int
		Slice       []int
		Object      struct {
			NestedField int
		}
	}{
		FieldString: "hello",
		FieldInt:    99,
		Slice:       []int{0, 0, 1},
	}
	m := map[string]interface{}{
		"FieldInt":    99,
		"FieldString": "hello",
		"Slice":       []int{0, 0, 1},
	}

	structf.UpdateStruct(&have, m)

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("Expect: %v, recieved: %v", want, have)
	}
}
