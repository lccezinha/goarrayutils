package arrayutils

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestIncludeItemOnArrayOfInt(t *testing.T) {
  var elements []int
  elements = append(elements, 1, 2, 3)

  item := 3
  result := Include(elements, item)
  assert.True(t, result, fmt.Sprintf("array %v must contain item %d", elements, item))
}

func TestIncludeItemOnArrayOfString(t *testing.T) {
  var elements []string
  elements = append(elements, "Luiz", "Cezer", "Cezinha")

  item := "Luiz"
  result := Include(elements, item)
  assert.True(t, result, fmt.Sprintf("array %v must contain item %s", elements, item))
}

func TestAnyItemOnArrayOfIntInCondition(t *testing.T) {
  var elements []int
  elements = append(elements, 10, 11, 5)

  condition := func(item interface{}) bool {
    return item.(int) >= 10
  }

  result := Any(elements, condition)
  assert.True(t, result, fmt.Sprintf("array %v must have any item in this condition", elements))
}

func TestAnyItemOnArrayOfStringsInCondition(t *testing.T) {
  var elements []string
  elements = append(elements, "Luiz", "Cezer", "Cezinha")

  condition := func(item interface{}) bool {
    return item.(string) == "Luiz"
  }

  result := Any(elements, condition)
  assert.True(t, result, fmt.Sprintf("array %v must have any item in this condition", elements))
}

func TestAnyItemOnArrayOfIntNotInCondition(t *testing.T) {
  var elements []int
  elements = append(elements, 10, 11, 5)

  condition := func(item interface{}) bool {
    return item.(int) >= 30
  }

  result := Any(elements, condition)
  assert.False(t, result, fmt.Sprintf("array %v must not have any item in this condition", elements))
}

func TestAnyItemOnArrayOfStringsNotInCondition(t *testing.T) {
  var elements []string
  elements = append(elements, "Luiz", "Cezer", "Cezinha")

  condition := func(item interface{}) bool {
    return item.(string) == "Xunda"
  }

  result := Any(elements, condition)
  assert.False(t, result, fmt.Sprintf("array %v must have any item in this condition", elements))
}

func TestNoneItemOnArrayInCondition(t *testing.T) {
  var elements []int
  elements = append(elements, 10, 11, 5)

  condition := func(item interface{}) bool {
    return item.(int) >= 10
  }

  result := None(elements, condition)
  assert.False(t, result, fmt.Sprintf("array %v must not have any item in this condition", elements))
}

func TestNoneItemOnArrayOfStringInCondition(t *testing.T) {
  var elements []string
  elements = append(elements, "Cezinha", "Cezer", "Luiz")

  condition := func(item interface{}) bool {
    return item.(string) != "Filho"
  }

  result := None(elements, condition)
  assert.False(t, result, fmt.Sprintf("array %v must not have any item in this condition", elements))
}

func TestCollectReturnNewArray(t *testing.T) {
  var elements, expected []interface{}
  elements = append(elements, 10, 11, 5)
  expected = append(expected, 10, 11)

  condition := func(item interface{}) bool {
    return item.(int) >= 10
  }

  result := Collect(elements, condition)
  assert.Equal(t, result, expected)
}

func TestCompact(t *testing.T) {
  var elements, expected []interface{}
  elements = append(elements, 10, 11, nil, 5)
  expected = append(expected, 10, 11, 5)

  result := Compact(elements)
  assert.Equal(t, result, expected)
}