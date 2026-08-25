package array

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// InArray reports whether target is present in arr.
func InArray(target interface{}, arr interface{}) bool {
	s := reflect.ValueOf(arr)

	for i := 0; i < s.Len(); i++ {
		if target == s.Index(i).Interface() {
			return true
		}
	}

	return false
}

// Intersect writes the common elements of a and b to refRet.
func Intersect(a, b, refRet interface{}) error {
	var (
		ifSlice []interface{}
	)

	m := make(map[interface{}]bool)
	s1 := reflect.ValueOf(a)
	for i := 0; i < s1.Len(); i++ {
		m[s1.Index(i).Interface()] = true
	}

	s2 := reflect.ValueOf(b)

	for i := 0; i < s2.Len(); i++ {
		if _, ok := m[s2.Index(i).Interface()]; ok {
			ifSlice = append(ifSlice, s2.Index(i).Interface())
		}
	}

	jsonStr, _ := json.Marshal(ifSlice)
	err := json.Unmarshal(jsonStr, &refRet)

	return err
}

// Diff writes the elements in X that are not matched in Y to refRet.
func Diff(X, Y, refRet interface{}) error {
	var (
		ifSlice []interface{}
		m       = make(map[interface{}]int)
	)

	s1 := reflect.ValueOf(Y)
	for i := 0; i < s1.Len(); i++ {
		m[s1.Index(i).Interface()]++
	}

	s2 := reflect.ValueOf(X)
	for i := 0; i < s2.Len(); i++ {
		y := s2.Index(i).Interface()
		if m[y] > 0 {
			m[y]--
			continue
		}
		ifSlice = append(ifSlice, y)
	}

	jsonStr, _ := json.Marshal(ifSlice)
	err := json.Unmarshal(jsonStr, &refRet)

	return err
}

// Unique writes the distinct elements of slice to refRet in first-occurrence order.
func Unique(slice interface{}, refRet interface{}) error {
	var (
		ifSlice []interface{}
	)

	keys := make(map[interface{}]bool)
	refVal := reflect.ValueOf(slice)

	for i := 0; i < refVal.Len(); i++ {
		entry := refVal.Index(i).Interface()
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			ifSlice = append(ifSlice, entry)
		}
	}

	jsonStr, _ := json.Marshal(ifSlice)
	err := json.Unmarshal(jsonStr, &refRet)

	return err
}

// Explode joins the elements of array with delimiter.
func Explode(delimiter string, array interface{}) string {
	var (
		joinStr string
	)

	if reflect.TypeOf(array).Kind() != reflect.Slice {
		return ""
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return ""
	}

	for i := 0; i < s.Len(); i++ {
		joinStr += fmt.Sprintf("%v", s.Index(i).Interface()) + delimiter
	}

	return joinStr[0 : len(joinStr)-1]
}
