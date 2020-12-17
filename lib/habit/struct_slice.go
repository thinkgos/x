package habit

import (
	"reflect"
	"strconv"
)

// StructsIntSlice returns a slice of int. For more info refer to Slice types StructIntSlice() method.
func StructsIntSlice(s interface{}, fieldName string) []int {
	return New(s).StructIntSlice(fieldName)
}

// StructsUintSlice returns a slice of int. For more info refer to Slice types v() method.
func StructsUintSlice(s interface{}, fieldName string) []uint {
	return New(s).StructUintSlice(fieldName)
}

// StructsInt64Slice returns a slice of int64. For more info refer to Slice types StructInt64Slice() method.
func StructsInt64Slice(s interface{}, fieldName string) []int64 {
	return New(s).StructInt64Slice(fieldName)
}

// StructsUint64Slice returns a slice of int64. For more info refer to Slice types StructUint64Slice() method.
func StructsUint64Slice(s interface{}, fieldName string) []uint64 {
	return New(s).StructUint64Slice(fieldName)
}

// StructStringSlice returns a slice of int64. For more info refer to Slice types StructStringSlice() method.
func StructStringSlice(s interface{}, fieldName string) []string {
	return New(s).StructStringSlice(fieldName)
}

// Slice hold a slice reflect.value
type Slice struct {
	value reflect.Value
}

// New returns a new *Slice with the slice s. It panics if the s's kind is not slice.
func New(s interface{}) *Slice {
	v := reflect.Indirect(reflect.ValueOf(s))
	if v.Kind() != reflect.Slice {
		panic("New: not slice")
	}
	return &Slice{v}
}

// StructIntSlice extracts the given s slice's every element, which is struct, to []int by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not signed integer.
func (sf *Slice) StructIntSlice(fieldName string) []int {
	length := sf.value.Len()
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		v := sf.structFieldVal(i, fieldName)
		slice[i] = int(valueInteger(v))
	}

	return slice
}

// StructUintSlice extracts the given s slice's every element, which is struct, to []uint by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not signed integer.
func (sf *Slice) StructUintSlice(fieldName string) []uint {
	length := sf.value.Len()
	slice := make([]uint, length)

	for i := 0; i < length; i++ {
		v := sf.structFieldVal(i, fieldName)
		slice[i] = uint(valueInteger(v))
	}

	return slice
}

// StructInt64Slice extracts the given s slice's every element, which is struct, to []int64 by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not signed integer.
func (sf *Slice) StructInt64Slice(fieldName string) []int64 {
	length := sf.value.Len()
	slice := make([]int64, length)

	for i := 0; i < length; i++ {
		v := sf.structFieldVal(i, fieldName)
		slice[i] = int64(valueInteger(v))
	}

	return slice
}

// StructUint64Slice extracts the given s slice's every element, which is struct, to []int64 by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not unsigned integer.
func (sf *Slice) StructUint64Slice(fieldName string) []uint64 {
	length := sf.value.Len()
	slice := make([]uint64, length)

	for i := 0; i < length; i++ {
		v := sf.structFieldVal(i, fieldName)
		slice[i] = valueInteger(v)
	}

	return slice
}

// StructStringSlice extracts the given s slice's every element, which is struct, to []string by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer or string.
func (sf *Slice) StructStringSlice(fieldName string) []string {
	length := sf.value.Len()
	slice := make([]string, length)

	for i := 0; i < length; i++ {
		v := sf.structFieldVal(i, fieldName)
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = strconv.FormatUint(v.Uint(), 10)
		case reflect.String:
			slice[i] = v.String()
		case reflect.Float32:
			slice[i] = strconv.FormatFloat(v.Float(), 'f', -1, 32)
		case reflect.Float64:
			slice[i] = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		default:
			panic("StructStringSlice: the value of field is not integer or float or string.")
		}
	}
	return slice
}

func (sf *Slice) structFieldVal(i int, fieldName string) reflect.Value {
	val := sf.value.Index(i)
	if !sf.isStruct(val) {
		panic("structFieldVal: the slice's element is not struct or pointer of struct!")
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	v := val.FieldByName(fieldName)
	if !v.IsValid() {
		panic("structFieldVal: the struct of slice's element has not the field:" + fieldName)
	}
	return v
}

func (sf *Slice) isStruct(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// uninitialized zero value of a struct
	if v.Kind() == reflect.Invalid {
		return false
	}
	return v.Kind() == reflect.Struct
}

// Name returns the slice's type name within its package. For more info refer
// to Name() function.
func (sf *Slice) Name() string {
	return sf.value.Type().Name()
}

func valueInteger(v reflect.Value) uint64 {
	switch v.Kind() { // nolint: exhaustive
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(v.Int())
	case reflect.Float32, reflect.Float64:
		return uint64(v.Float())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	default:
		panic("StructSlice: the value of field is not integer or float.")
	}
}
