package sliceutils

import "reflect"

func Include(slice interface{}, searchedItem interface{}) bool {
	values := reflect.ValueOf(slice)
	resultSlice := make([]interface{}, values.Len())

	for i := 0; i < values.Len(); i++ {
		resultSlice[i] = values.Index(i).Interface()

		if resultSlice[i] == searchedItem {
			return true
		}
	}

	return false
}

func Any(slice interface{}, callback func(interface{}) bool) bool {
	values := reflect.ValueOf(slice)
	resultSlice := make([]interface{}, values.Len())

	for i := 0; i < values.Len(); i++ {
		resultSlice[i] = values.Index(i).Interface()

		if callback(resultSlice[i]) {
			return true
		}
	}

	return false
}

func None(slice interface{}, callback func(interface{}) bool) bool {
	return !Any(slice, callback)
}

func Collect(slice []interface{}, callback func(interface{}) bool) interface{} {
	values := reflect.ValueOf(slice)
	resultSlice := make([]interface{}, values.Len())
	var result []interface{}

	for i := 0; i < values.Len(); i++ {
		resultSlice[i] = values.Index(i).Interface()

		if callback(resultSlice[i]) {
			result = append(result, resultSlice[i])
		}
	}

	return result
}

func Compact(slice interface{}) interface{} {
	values := reflect.ValueOf(slice)
	resultSlice := make([]interface{}, values.Len())
	var result []interface{}

	for i := 0; i < values.Len(); i++ {
		resultSlice[i] = values.Index(i).Interface()

		if resultSlice[i] != nil {
			result = append(result, resultSlice[i])
		}
	}

	return result
}

func All(slice interface{}, callback func(interface{}) bool) bool {
	values := reflect.ValueOf(slice)
	resultSlice := make([]interface{}, values.Len())
	returnValue := true

	for i := 0; i < values.Len(); i++ {
		resultSlice[i] = values.Index(i).Interface()

		if callback(resultSlice[i]) == false {
			returnValue = false
		}
	}

	return returnValue
}

func Reject(slice []interface{}, callback func(interface{}) bool) interface{} {
	values := reflect.ValueOf(slice)
	resultSlice := make([]interface{}, values.Len())
	var result []interface{}

	for i := 0; i < values.Len(); i++ {
		resultSlice[i] = values.Index(i).Interface()

		if !callback(resultSlice[i]) {
			result = append(result, resultSlice[i])
		}
	}

	return result
}
