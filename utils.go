package arrayutils

import "reflect"

func Include(slice interface{}, searchedItem interface{}) bool {
  values := reflect.ValueOf(slice)
  arr := make([]interface{}, values.Len())

  for i := 0; i < values.Len(); i++ {
    arr[i] = values.Index(i).Interface()

    if arr[i] == searchedItem {
      return true
    }
  }

  return false
}

func Any(slice interface{}, condition func(interface{}) bool) bool {
  values := reflect.ValueOf(slice)
  arr := make([]interface{}, values.Len())

  for i := 0; i < values.Len(); i++ {
    arr[i] = values.Index(i).Interface()

    if condition(arr[i]) {
      return true
    }
  }

  return false
}

func None(slice interface{}, condition func(interface{}) bool) bool {
  return !Any(slice, condition)
}

func Collect(slice []interface{}, condition func(interface{}) bool) interface{} {
  values := reflect.ValueOf(slice)
  arr    := make([]interface{}, values.Len())
  var result []interface{}

  for i := 0; i < values.Len(); i++ {
    arr[i] = values.Index(i).Interface()

    if condition(arr[i]) {
      result = append(result, arr[i])
    }
  }

  return result
}

func Compact(slice interface{}) interface{} {
  values := reflect.ValueOf(slice)
  arr := make([]interface{}, values.Len())
  var result []interface{}

  for i := 0; i < values.Len(); i++ {
    arr[i] = values.Index(i).Interface()

    if arr[i] != nil {
      result = append(result, arr[i])
    }
  }

  return result
}