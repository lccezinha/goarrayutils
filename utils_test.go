package arrayutils

import (
  "github.com/stretchr/testify/assert"
  "testing"
  "fmt"
)

func TestIncludeItemOnArrayOfInt(t *testing.T) {
  var elements []interface{}
  elements = append(elements, 1, 2, 3)

  item := 3
  result := Include(elements, item)
  assert.True(t, result, fmt.Sprintf("array %v must contain item %d", elements, item))
}

func TestIncludeItemOnArrayOfString(t *testing.T) {
  var elements []interface{}
  elements = append(elements, "Luiz", "Cezer", "Cezinha")

  item := "Luiz"
  result := Include(elements, item)
  assert.True(t, result, fmt.Sprintf("array %v must contain item %s", elements, item))
}

func TestIncludeItemOnArrayOfFloat(t *testing.T) {
  var elements []interface{}
  elements = append(elements, 10.0, 11.2, 5.2)

  item := 11.2
  result := Include(elements, item)
  assert.True(t, result, fmt.Sprintf("array %v must contain item %.1f", elements, item))
}

func TestAnyItemOnArrayInCondition(t *testing.T) {
  var elements []int
  elements = append(elements, 10, 11, 5)

  condition := func(item int) bool {
    return item >= 10
  }

  result := Any(elements, condition)
  assert.True(t, result, fmt.Sprintf("array %v must have any item in this condition", elements))
}

func TestAnyItemOnArrayNotInCondition(t *testing.T) {
  var elements []int
  elements = append(elements, 10, 11, 5)

  condition := func(item int) bool {
    return item >= 30
  }

  result := Any(elements, condition)
  assert.False(t, result, fmt.Sprintf("array %v must not have any item in this condition", elements))
}