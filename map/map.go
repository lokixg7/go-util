package maps

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Map2List writes the values of m to refRet.
func Map2List(m interface{}, refRet interface{}) error {
	// make sure an empty map yields an empty list (nil slice would marshal to null)
	ifSlice := make([]interface{}, 0)

	s := reflect.ValueOf(m)
	if s.Kind() != reflect.Map {
		return fmt.Errorf("Map2List: m must be a map, got %T", m)
	}

	for _, key := range s.MapKeys() {
		ifSlice = append(ifSlice, s.MapIndex(key).Interface())
	}

	jsonStr, _ := json.Marshal(ifSlice)
	err := json.Unmarshal(jsonStr, &refRet)

	return err
}
