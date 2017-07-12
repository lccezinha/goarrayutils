package sliceutils

import (
	"reflect"
	"testing"
)

func TestIncludeItemOnArrayOfInt(t *testing.T) {
	var elements []int
	elements = append(elements, 1, 2, 3)

	item := 3
	result := Include(elements, item)

	if !result {
		t.Fatalf("slice %v must contain item %d", elements, item)
	}
}

func TestIncludeItemOnArrayOfString(t *testing.T) {
	var elements []string
	elements = append(elements, "Luiz", "Cezer", "Cezinha")

	item := "Luiz"
	result := Include(elements, item)

	if !result {
		t.Fatalf("slice %v must contain item %s", elements, item)
	}
}

func TestAnyItemOnArrayOfIntInCallback(t *testing.T) {
	var elements []int
	elements = append(elements, 10, 11, 5)

	callback := func(item interface{}) bool {
		return item.(int) >= 10
	}

	result := Any(elements, callback)

	if !result {
		t.Fatalf("slice %v must have any item in this callback", elements)
	}
}

func TestAnyItemOnArrayOfStringsInCallback(t *testing.T) {
	var elements []string
	elements = append(elements, "Luiz", "Cezer", "Cezinha")

	callback := func(item interface{}) bool {
		return item.(string) == "Luiz"
	}

	result := Any(elements, callback)

	if !result {
		t.Fatalf("slice %v must have some item in this callback", elements)
	}
}

func TestAnyItemOnArrayOfIntNotInCallback(t *testing.T) {
	var elements []int
	elements = append(elements, 10, 11, 5)

	callback := func(item interface{}) bool {
		return item.(int) >= 30
	}

	result := Any(elements, callback)

	if result {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestAnyItemOnArrayOfStringsNotInCallback(t *testing.T) {
	var elements []string
	elements = append(elements, "Luiz", "Cezer", "Cezinha")

	callback := func(item interface{}) bool {
		return item.(string) == "Xunda"
	}

	result := Any(elements, callback)

	if result {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestNoneItemOnArrayInCallback(t *testing.T) {
	var elements []int
	elements = append(elements, 10, 11, 5)

	callback := func(item interface{}) bool {
		return item.(int) >= 10
	}

	result := None(elements, callback)

	if result {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestNoneItemOnArrayOfStringInCallback(t *testing.T) {
	var elements []string
	elements = append(elements, "Cezinha", "Cezer", "Luiz")

	callback := func(item interface{}) bool {
		return item.(string) != "Filho"
	}

	result := None(elements, callback)

	if result {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestCollectReturnNewArray(t *testing.T) {
	var elements, expected []interface{}
	elements = append(elements, 10, 11, 5)
	expected = append(expected, 10, 11)

	callback := func(item interface{}) bool {
		return item.(int) >= 10
	}

	result := Collect(elements, callback)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestCompact(t *testing.T) {
	var elements, expected []interface{}
	elements = append(elements, 10, 11, nil, 5)
	expected = append(expected, 10, 11, 5)

	result := Compact(elements)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestAllWithAllValidValues(t *testing.T) {
	var elements []interface{}
	elements = append(elements, 1, 2, 3)

	callback := func(item interface{}) bool {
		return item.(int) > 0
	}

	result := All(elements, callback)

	if !result {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestAllWithAnyInvalidValue(t *testing.T) {
	var elements []interface{}
	elements = append(elements, 1, 2, 3)

	callback := func(item interface{}) bool {
		return item.(int) > 2
	}

	result := All(elements, callback)

	if result {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}

func TestReject(t *testing.T) {
	var elements, expected []interface{}
	elements = append(elements, 1, 2, 3)
	expected = append(expected, 1)

	callback := func(item interface{}) bool {
		return item.(int) >= 2
	}

	result := Reject(elements, callback)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("slice %v must not have any item in this callback", elements)
	}
}
