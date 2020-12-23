package univ

import (
	"reflect"
)

// KeysOfMap 获取map m的key的切片,要求m为map且key为字符串(string).
func KeysOfMap(m interface{}) []string {
	rv := reflect.Indirect(reflect.ValueOf(m))
	if rv.Kind() == reflect.Invalid {
		return []string{}
	}
	if rv.Kind() != reflect.Map {
		panic("KeysOfMap: require a map")
	}

	keys := rv.MapKeys()
	ss := make([]string, 0, len(keys))
	for _, key := range keys {
		key = reflect.Indirect(key)
		if key.Kind() != reflect.String {
			panic("KeysOfMap: require string type of map key")
		}
		ss = append(ss, key.String())
	}
	return ss
}

// KeysIntOfMap 获取map m的key的切片,要求m为map且key为数值
// (int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64).
func KeysIntOfMap(m interface{}) []int64 {
	rv := reflect.Indirect(reflect.ValueOf(m))
	if rv.Kind() == reflect.Invalid {
		return []int64{}
	}
	if rv.Kind() != reflect.Map {
		panic("KeysIntOfMap: require a map")
	}

	keys := rv.MapKeys()
	ss := make([]int64, 0, len(keys))
	for _, key := range keys {
		key = reflect.Indirect(key)
		switch key.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		default:
			panic("KeysIntOfMap: require integer type of map key")
		}
		ss = append(ss, key.Int())
	}
	return ss
}
