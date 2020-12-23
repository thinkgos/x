package univ

import (
	"reflect"
	"strconv"
)

// StructsIntSlice returns a slice of int. For more info refer to Slice types StructIntSlice() method.
func StructsIntSlice(s interface{}, fieldName string) []int {
	return NewSlice(s).StructIntSlice(fieldName)
}

// StructsUintSlice returns a slice of int. For more info refer to Slice types v() method.
func StructsUintSlice(s interface{}, fieldName string) []uint {
	return NewSlice(s).StructUintSlice(fieldName)
}

// StructsInt64Slice returns a slice of int64. For more info refer to Slice types StructInt64Slice() method.
func StructsInt64Slice(s interface{}, fieldName string) []int64 {
	return NewSlice(s).StructInt64Slice(fieldName)
}

// StructsUint64Slice returns a slice of int64. For more info refer to Slice types StructUint64Slice() method.
func StructsUint64Slice(s interface{}, fieldName string) []uint64 {
	return NewSlice(s).StructUint64Slice(fieldName)
}

// StructStringSlice returns a slice of int64. For more info refer to Slice types StructStringSlice() method.
func StructStringSlice(s interface{}, fieldName string) []string {
	return NewSlice(s).StructStringSlice(fieldName)
}

// IntSlice returns a slice of int. For more info refer to Slice types IntSlice() method.
func IntSlice(s interface{}) []int {
	return NewSlice(s).IntSlice()
}

// UintSlice returns a slice of uint. For more info refer to Slice types UintSlice() method.
func UintSlice(s interface{}) []uint {
	return NewSlice(s).UintSlice()
}

// Int8Slice returns a slice of int8. For more info refer to Slice types Int8Slice() method.
func Int8Slice(s interface{}) []int8 {
	return NewSlice(s).Int8Slice()
}

// Uint8Slice returns a slice of uint8. For more info refer to Slice types Uint8Slice() method.
func Uint8Slice(s interface{}) []uint8 {
	return NewSlice(s).Uint8Slice()
}

// Int16Slice returns a slice of int16. For more info refer to Slice types Int16Slice() method.
func Int16Slice(s interface{}) []int16 {
	return NewSlice(s).Int16Slice()
}

// Uint16Slice returns a slice of uint16. For more info refer to Slice types Uint16Slice() method.
func Uint16Slice(s interface{}) []uint16 {
	return NewSlice(s).Uint16Slice()
}

// Int32Slice returns a slice of int32. For more info refer to Slice types Int32Slice() method.
func Int32Slice(s interface{}) []int32 {
	return NewSlice(s).Int32Slice()
}

// Uint32Slice returns a slice of uint32. For more info refer to Slice types Uint32Slice() method.
func Uint32Slice(s interface{}) []uint32 {
	return NewSlice(s).Uint32Slice()
}

// Int64Slice returns a slice of int64. For more info refer to Slice types Int64Slice() method.
func Int64Slice(s interface{}) []int64 {
	return NewSlice(s).Int64Slice()
}

// Uint64Slice returns a slice of uint64. For more info refer to Slice types Uint64Slice() method.
func Uint64Slice(s interface{}) []uint64 {
	return NewSlice(s).Uint64Slice()
}

// Slice hold a slice reflect.value
type Slice struct {
	value reflect.Value
}

// NewSlice returns a new *Slice with the slice s. It panics if the s's kind is not slice.
func NewSlice(s interface{}) *Slice {
	v := reflect.Indirect(reflect.ValueOf(s))
	if v.Kind() != reflect.Slice {
		panic("NewSlice: not slice")
	}
	return &Slice{v}
}

// StructIntSlice extracts the given s slice's every element, which is struct, to []int by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
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
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
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
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
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
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
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

// IntSlice extracts the given s slice's every element, which is integer or float, to []int by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) IntSlice() []int {
	length := sf.value.Len()
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// UintSlice extracts the given s slice's every element, which is integer or float, to []uint by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) UintSlice() []uint {
	length := sf.value.Len()
	slice := make([]uint, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int8Slice extracts the given s slice's every element, which is integer or float, to []int8 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Int8Slice() []int8 {
	length := sf.value.Len()
	slice := make([]int8, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int8(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int8(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int8(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint8Slice extracts the given s slice's every element, which is integer or float, to []uint8 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Uint8Slice() []uint8 {
	length := sf.value.Len()
	slice := make([]uint8, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint8(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint8(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint8(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int16Slice extracts the given s slice's every element, which is integer or float, to []int16 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Int16Slice() []int16 {
	length := sf.value.Len()
	slice := make([]int16, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int16(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int16(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int16(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint16Slice extracts the given s slice's every element, which is integer or float, to []uint16 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Uint16Slice() []uint16 {
	length := sf.value.Len()
	slice := make([]uint16, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint16(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint16(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint16(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int32Slice extracts the given s slice's every element, which is integer or float, to []int32 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Int32Slice() []int32 {
	length := sf.value.Len()
	slice := make([]int32, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int32(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int32(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int32(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint32Slice extracts the given s slice's every element, which is integer or float, to []uint32 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Uint32Slice() []uint32 {
	length := sf.value.Len()
	slice := make([]uint32, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint32(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint32(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint32(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int64Slice extracts the given s slice's every element, which is integer or float, to []int64 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Int64Slice() []int64 {
	length := sf.value.Len()
	slice := make([]int64, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = v.Int()
		case reflect.Float32, reflect.Float64:
			slice[i] = int64(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int64(v.Uint())
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint64Slice extracts the given s slice's every element, which is integer or float, to []uint64 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (sf *Slice) Uint64Slice() []uint64 {
	length := sf.value.Len()
	slice := make([]uint64, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(sf.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint64(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint64(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = v.Uint()
		default:
			panic("StructSlice: the value of field is not integer or float.")
		}
	}
	return slice
}

func (sf *Slice) structFieldVal(i int, fieldName string) reflect.Value {
	val := sf.value.Index(i)
	val = reflect.Indirect(val)

	// check is struct
	if !(val.Kind() != reflect.Invalid && val.Kind() == reflect.Struct) {
		panic("structFieldVal: the slice's element is not struct or pointer of struct!")
	}

	v := val.FieldByName(fieldName)
	if !v.IsValid() {
		panic("structFieldVal: the struct of slice's element has not the field:" + fieldName)
	}
	return v
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
