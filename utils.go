package arrayutils

func Include(arr []interface{}, searchedItem interface{}) bool {
  for _, item := range arr {
    if item == searchedItem {
      return true
    }
  }

  return false
}

func Any(arr []interface{}, condition func(interface{}) bool) bool {
  for _, item := range arr {
    if condition(item) {
      return true
    }
  }

  return false
}

func None(arr []interface{}, condition func(interface{}) bool) bool {
  return !Any(arr, condition)
}

func Collect(arr []interface{}, condition func(interface{}) bool) interface{} {
  var result []interface{}

  for _, item := range arr {
    if condition(item) {
      result = append(result, item)
    }
  }

  return result
}

func Compact(arr []interface{}) interface{} {
  var result []interface{}

  for _, item := range arr {
    if item != nil {
      result = append(result, item)
    }
  }

  return result
}