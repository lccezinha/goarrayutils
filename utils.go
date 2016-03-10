package arrayutils

func Include(arr []interface{}, searchedItem interface{}) bool {
  for _, item := range arr {
    if item == searchedItem {
      return true
    }
  }

  return false
}

func Any(arr []int, condition func(int) bool) bool {
  for _, item := range arr {
    if condition(item) {
      return true
    }
  }

  return false
}

func None(arr []int, condition func(int) bool) bool {
  return !Any(arr, condition)
}

func Collect(arr []int, condition func(int) bool) []int {
  var result []int

  for _, item := range arr {
    if condition(item) {
      result = append(result, item)
    }
  }

  return result
}