package arrayutils

func Include(arr []interface{}, searchedItem interface{}) bool {
  for _, item := range arr {
    if item == searchedItem {
      return true
    }
  }

  return false
}

func Any(arr []int, condition func(int) bool) bool{
  for _, item := range arr {
    if condition(item) {
      return true
    }
  }

  return false
}