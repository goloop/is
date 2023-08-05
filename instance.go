package is

import (
	"reflect"
)

/*
Instance checks the type of the value.

Symantic:

func Instance[T any](obj interface{}, mode int) bool {}

where T is a type to determine the match
and obj is a value to check,
and mode is a mode of check:
    -1 and mode < 0 - flexibility mode;
                  0 - normal (default) mode;
     1 and mode > 0 - strict mode.

Example:

var a int

// Default mode, type definition by logical group.
Instance[int](a, 0) == true
Instance[int32](a, 0) == true
Instance[uint](a, 0) == false
Instance[string](a, 0) == false

// Flexibility mode, defining a type by its general form.
Instance[int](a, -1) == true
Instance[int32](a, -1) == true
Instance[uint](a, -1) == true
Instance[uint32](a, -1) == true
Instance[string](a, -1) == false

// Strict mode, determines the type and form of the type.
Instance[int](a, 1) == true
Instance[int32](a, 1) == false
Instance[uint](a, 1) == false
Instance[string](a, 1) == false


1. SIMPLE TYPES
===============
int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
float32, float64, complex64, complex128, bool, byte, rune, string,
*int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32,
*uint64, *float32, *float64, *complex64, *complex128, *bool, *byte,
*rune, *string

Flexibility group:
1. int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
float32, float64, complex64, complex128, byte, rune;
2. string;
3. bool;
4. *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32,
*uint64, *float32, *float64, *complex64, *complex128, *byte, *rune;
5. *string;
6. *bool.

Default mode:
1. int, int8, int16, int32, int64;
2. uint, uint8, uint16, uint32, uint64;
3. float32, float64;
4. complex64, complex128;
5. byte, rune;
6. string;
7. bool;
8. *int, *int8, *int16, *int32, *int64;
9. *uint, *uint8, *uint16, *uint32, *uint64;
10. *float32, *float64;
11. *complex64, *complex128;
12. *byte, *rune;
13. *string;
14. *bool.

Strict mode:
- Strict type definition!

1.1. Value
```
var a int

// Flexibility mode.
// Logical group.
Instance[int](a, -1) && Instance[int8](a, -1) && Instance[int16](a, -1) &&
Instance[uint32](a, -1) && Instance[float64](a, -1) &&
Instance[complex128](a, -1) && Instance[byte](a, -1) &&
Instance[rune](a, -1) == true

// False for different types.
Instance[string](a, -1) || Instance[bool](a, -1) == false

// And false for pointers.
Instance[*int](a, -1) == false

// Normal mode.
// Default group.
// True for all int* types.
Instance[int](a, 0) && Instance[int8](a, 0) && Instance[int16](a, 0) &&
Instance[int32](a, 0) && Instance[int64](a, 0) == true

// False for different types.
Instance[uint](a, 0) || Instance[float32](a, 0) ||
Instance[complex64](a, 0) == false

// And false for pointers.
Instance[*int](a, 0) == false

// Strict mode.
// True for specific type only.
Instance[int](a, 1) == true

// False for different types.
Instance[uint](a, 1) || Instance[float32](a, 1) ||
Instance[complex64](a, 1) == false

// And false for pointers.
Instance[*int](a, 1) == false
```

1.2. Pointer
```
var a *int

// Flexibility mode.
Instance[*int](a, -1) && Instance[*int8](a, -1) && Instance[*int16](a, -1) &&
Instance[*float32](a, -1) && Instance[*complex64](a, -1) &&
Instance[*byte](a, -1) && Instance[*rune](a, -1) == true

// False for different group of types.
Instance[*bool](a, -1) || Instance[*string](a, -1) == false

// And false for non-pointers.
Instance[int](a, -1) == false


// Normal mode.
// True for all *int* types.
Instance[*int](a, 0) && Instance[*int8](a, 0) && Instance[*int16](a, 0) &&
Instance[*int32](a, 0) && Instance[*int64](a, 0) == true

// False for different types.
Instance[*uint](a, 0) || Instance[*float32](a, 0) ||
Instance[*complex64](a, 0) == false

// And false for non-pointers.
Instance[int](a, 0) == false

// Strict mode.
// True for specific type only.
Instance[*int](a, 1) == true

// False for different types.
Instance[*uint](a, 1) || Instance[*float32](a, 1) ||
Instance[*complex64](a, 1) == false

// And false for non-pointers.
Instance[int](a, 1) == false
```

2. COMPLEX TYPES
================

Any other types are complex, because they consist of two or more subtypes.


2.1. Arrays

2.1.1. Value as value.
```
var a [3]int

// Flexibility mode.
// Any size array of int* types.
Instance[[3]int](a, -1) && Instance[[4]int](a, -1) &&
Instance[[0]int32](a, -1) == true

// False for different types.
Instance[[3]bool](a, -1) || Instance[[3]string](a, -1) == false

// And false for pointers.
Instance[[3]*int](a, -1) == false

// Normal mode.
// Any size array of int* types.
Instance[[3]int](a, 0) && Instance[[4]int](a, 0) &&
Instance[[0]int32](a, 0) == true

// False for different types.
Instance[[3]uint](a, 0) || Instance[[3]float32](a, 0)
Instance[[3]bool](a, 0) || Instance[[3]string](a, 0) == false

// And false for pointers.
Instance[[3]*int](a, 0) == false

// Strict mode.
// Strict size array of strict types.
Instance[[3]int](a, 1) == true

// False for different types or sizes.
Instance[[3]uint](a, 1) || Instance[[3]float32](a, 1) ||
Instance[[0]int](a, 1) == false

// And false for pointers.
Instance[[3]*int](a, 1) == false
```

2.1.2. Value as pointer.
```
var a [3]*int

// Flexibility mode.
Instance[[3]*int](a, -1) && Instance[[4]*uint](a, -1) &&
Instance[[0]*float32](a, -1) == true


// False for different types.
Instance[[3]*bool](a, -1) || Instance[[3]*string](a, -1) == false

// And false for non-pointers.
Instance[[3]int](a, -1) == false

// Normal mode.
// Any size array of int* types.
Instance[[3]*int](a, 0) && Instance[[4]*int8](a, 0) &&
Instance[[0]*int64](a, 0) == true

// False for different types.
Instance[[3]*uint](a, 0) || Instance[[3]*float32](a, 0) ||
Instance[[0]*complex64](a, 0) == false

// And false for non-pointers.
Instance[[3]int](a, 0) == false


// Strict mode.
// Strict size array of strict types.
Instance[[3]*int](a, 1) == true

// False for different types or sizes.
Instance[[3]*uint](a, 1) || Instance[[3]*float32](a, 1) ||
Instance[[0]*int](a, 1) == false

// And false for non-pointers.
Instance[[3]int](a, 1) == false
```


2.2. Slices
Exactly the same as for arrays.
But note that slices don't have reflect.Len()

2.3. Maps

2.3.1. Value as value.
```
var a map[int]int

// Flexibility mode.
// Just map the specific form.
Instance[map[int]int](a, 0) && Instance[map[int8]int](a, 0) &&
Instance[map[int16]uint](a, 0) && Instance[map[int32]false32](a, 0) == true

// False for different types.
Instance[map[uint]string](a, 0) || Instance[map[uint8]bool](a, 0) ||
Instance[map[string]uint](a, 0) || Instance[map[string]bool](a, 0) == false

// And false for pointers.
Instance[map[int]*int](a, 0) == false

// Normal mode.
// Just map the specific form.
Instance[map[int]int](a, 0) && Instance[map[int8]int](a, 0) &&
Instance[map[int16]uint](a, 0) && Instance[map[int32]false32](a, 0) == true

// False for different types.
Instance[map[uint]string](a, 0) || Instance[map[uint8]bool](a, 0) ||
Instance[map[string]uint](a, 0) || Instance[map[string]bool](a, 0) == false


// And false for pointers.
Instance[map[int]*int](a, 0) == false
````

2.3.2. Value as pointer.
```
var a map[int]*int

// Flexibility mode.
Instance[map[int]*int](a, -1) && Instance[map[int8]*int](a, -1) &&
Instance[map[int16]*int](a, -1) && Instance[map[int32]*int8](a, -1) == true

// False for different types.
Instance[map[string]*int](a, -1) || Instance[map[uint8]*bool](a, -1)  == false

// And false for non-pointers.
Instance[map[int]int](a, -1) == false

// Normal mode.
Instance[map[int]*int](a, 0) && Instance[map[int8]*int](a, 0) &&
Instance[map[int16]*int](a, 0) && Instance[map[int32]*int8](a, 0) == true

// False for different types.
Instance[map[uint]*int](a, 0) || Instance[map[int]*float32](a, 0)  == false

// And false for non-pointers.
Instance[map[int]int](a, 0) == false


// Strict mode.
Instance[map[int]*int](a, 1) == true

// False for different types or sizes.
Instance[map[int]*int8](a, 1) || Instance[map[int]*float32](a, 1) ||
Instance[map[uint]*int](a, 1) == false

// And false for non-pointers.
Instance[map[int]int](a, 0) == false
```


2.4. Functions
```
sum := func(a int, b int) int {return a + b}

// Flexibility mode.
// Any function, or function with any arguments and return value
// from the flexibility group. If test type has inputs/outputs it mast
// has full signature.
Instance[func()](sum, -1) && Instance[func(int, int) int](sum, -1) &&
Instance[func(float32, int) uint](sum, -1) == true

// Not full signature.
Instance[func(int) int](sum, -1) || Instance[func(int, int)](sum, -1) == false

// Normal mode.
// Any function, or function with any arguments and return value
// from the normal group. If test type has inputs/outputs it mast
// has full signature.
Instance[func()](sum, 0) && Instance[func(int, int) int](sum, 0) &&
Instance[func(uint, int) uint](sum, 0) == true


// Strict mode.
// Strict function, and function with strict arguments.
Instance[func(int, int) int](sum, 1) == true

// False for different types.
Instance[func()](sum, 1) || Instance[func(int, int) uint](sum, 1) ||
Instance[func(uint, int) int](sum, 1) == false
```

2.5. Structs

2.5.1. Value as value.
```
// One is a test struct.
type One struct {
    Name string
    Age int
}

// Two is a test struct with fields like One.
type Two struct {
    Name string
    Age int
}

// Three is a test struct with fields like One and other.
type Three struct {
    Name string
    Age int
    Weight float32
}

var a A

// Flexibility mode.
// Any struct, or struct with some fields.
Instance[struct{}](a, -1) && Instance[One](a, -1) &&
Instance[Two](a, -1) == true

// False for different types.
Instance[Three](a, -1) || Instance[struct{Name string}](a, -1) == false

// And false for pointers.
Instance[*One](a, -1) == false

// Normal mode.
// Any struct, or struct with this type.
Instance[struct{}](a, 0) && Instance[One](a, 0) == true

// False for different types.
Instance[Two](a, 0) || Instance[struct{Name string}](a, 0) == false

// And false for pointers.
Instance[*One](a, 0) == false

// Strict mode.
// Strict struct, and struct with strict fields.
Instance[One](a, 1) == true

// False for different types.
Instance[Two](a, 1) || Instance[struct{Name string}](a, 1) ||
Instance[struct{}](a, -1) == false

// And false for pointers.
Instance[*One](a, 1) == false
```


2.5.2. Value as pointer.
```
// One is a test struct.
type One struct {
    Name string
    Age int
}

// Two is a test struct with fields like One.
type Two struct {
    Name string
    Age int
}

// Three is a test struct with fields like One and other.
type Three struct {
    Name string
    Age int
    Weight float32
}

var a *A

// Flexibility mode.
// Any struct, or struct with some fields.
Instance[*struct{}](a, -1) && Instance[*One](a, -1) &&
Instance[*Two](a, -1) == true

// False for different types.
Instance[*Three](a, -1) || Instance[*struct{Name string}](a, -1) == false

// And false for non-pointers.
Instance[One](a, -1) == false

// Normal mode.
// Any struct, or struct with this type.
Instance[*struct{}](a, 0) && Instance[*One](a, 0) == true

// False for different types.
Instance[*Two](a, 0) || Instance[*struct{Name string}](a, 0) == false

// And false for non-pointers.
Instance[One](a, 0) == false

// Strict mode.
// Strict struct, and struct with strict fields.
Instance[*One](a, 1) == true

// False for different types.
Instance[*Two](a, 1) || Instance[*struct{Name string}](a, 1) ||
Instance[*struct{}](a, -1) == false

// And false for pointers.
Instance[One](a, 1) == false
```

2.6. Channels

The same logic.

3. OVERCOMPLICATED TYPES
========================

For example:
```
// Slice of map with int key and slice of string as values.
var a []map[int][]string

// Flexibility mode, determines only the form of the type.
Instance[[]map[int][]string](a, -1) &&
Instance[[]map[uint8][]string](a, -1) == true

// False for different types.
Instance[[]map[string][]string](a, -1) ||
Instance[[]map[int][]bool](a, -1) == false

// And false for pointers.
Instance[[]map[int][]*string](a, -1) == false

// Normal mode, determines only the form of the type.
Instance[[]map[int][]string](a, 0) &&
Instance[[]map[int8][]string](a, 0) == true

// False for different types.
Instance[[]map[uint][]string](a, 0) &&
Instance[[]map[uint8][]string](a, 0) == true

// And false for pointers.
Instance[[]map[int][]*string](a, 0) == false

// Strict mode.
// Strict type, and type with strict arguments.
Instance[[]map[int][]string](a, 1) == true
```
*/

// Nil returns true if the object is nil. This function checks for
// nil on various types including pointers, slices, maps, channels,
// functions, and interfaces.
//
// Example usage:
//
//	var a *int
//	fmt.Println(Nil(a)) // Outputs: true
//
//	var b []int
//	fmt.Println(Nil(b)) // Outputs: true
//
//	c := make(chan int)
//	fmt.Println(Nil(c)) // Outputs: false
//
//	var d func()
//	fmt.Println(Nil(d)) // Outputs: true
//
//	var e interface{}
//	fmt.Println(Nil(e)) // Outputs: true
//
//	e = 123
//	fmt.Println(Nil(e)) // Outputs: false
func Nil(obj interface{}) bool {
	if obj == nil {
		return true
	}

	rv := reflect.ValueOf(obj)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Func:
		fallthrough
	case reflect.Slice, reflect.Map, reflect.Chan:
		return rv.IsNil()
	}

	return false
}

// // Instance returns true if the instance is the same type as the base.
// func Instance[T any](obj interface{}, strict ...bool) bool {
// 	// Flexibility mode if all flex is true.
// 	flex := len(strict) == 0 // default is true
// 	for _, s := range strict {
// 		if !s {
// 			flex = true
// 			break
// 		}
// 	}

// 	// If the same type, return true.
// 	var base T
// 	brv, irv := reflect.ValueOf(base), reflect.ValueOf(obj)
// 	if !flex && brv.Kind() != irv.Kind() {
// 		return false
// 	}

// 	// Try detect the same type.
// 	brt := reflect.TypeOf(base)
// 	irt := reflect.TypeOf(obj)
// 	return isSameType(brt, irt, flex, 0)
// }

// // isSameType returns true if the instance is the same type as the val.
// func isSameType(brt, irt reflect.Type, flex bool, level int) bool {
// 	// if level == 0 && brt.Kind() != irt.Kind() {
// 	// 	return false
// 	// } else if level > 0 && flex {
// 	// 	return true
// 	// }

// 	if brt.Kind() != irt.Kind() {
// 		return false
// 	}

// 	switch brt.Kind() {
// 	case reflect.Slice, reflect.Array, reflect.Map:
// 		return isContainer(brt, irt, flex, level)
// 	case reflect.Ptr:
// 		return isPointer(brt, irt, flex, level)
// 	case reflect.Struct:
// 		return isStruct(brt, irt, flex, level)
// 	case reflect.Func:
// 		return isFunc(brt, irt, flex, level)
// 	case reflect.Chan:
// 		return isChan(brt, irt, flex, level)
// 	}

// 	return true
// }

// // isContainer returns true if the instance is the same type as the val.
// func isContainer(brt, irt reflect.Type, flex bool, level int) bool {
// 	if flex && brt.Kind() == irt.Kind() {
// 		// Check base type if it pointer.
// 		if brt.Elem().Kind() == reflect.Ptr {
// 			if irt.Elem().Kind() != reflect.Ptr {
// 				return false
// 			}
// 		} else if brt.Elem().Kind() != reflect.Ptr {
// 			if irt.Elem().Kind() == reflect.Ptr {
// 				return true
// 			}
// 		}

// 		return true
// 	}

// 	// Check the Kind of the reflect.Type before calling Kind.
// 	if brt.Kind() != irt.Kind() {
// 		return false
// 	}

// 	// Pointer state.
// 	if brt.Elem().Kind() == reflect.Ptr {
// 		if irt.Elem().Kind() != reflect.Ptr {
// 			return false
// 		}
// 	}

// 	// Check the Kind of the reflect.Type before calling Len.
// 	if brt.Kind() == reflect.Array {
// 		if brt.Len() != irt.Len() {
// 			return false
// 		}

// 		if brt.Elem() != irt.Elem() {
// 			return false
// 		}
// 	}

// 	// Check the Kind of the reflect.Type before calling Elem.
// 	if brt.Kind() == reflect.Slice {
// 		if brt.Elem() != irt.Elem() {
// 			return false
// 		}
// 	}

// 	// Check the Kind of the reflect.Type before calling Key.
// 	if brt.Kind() == reflect.Map {
// 		if brt.Key() != irt.Key() {
// 			return false
// 		}
// 	}

// 	// For interface{} we can't check the Elem() type.
// 	if brt.Elem().Kind() == reflect.Interface {
// 		return true
// 	}

// 	return isSameType(brt.Elem(), irt.Elem(), flex, level+1)
// }

// // isPointer returns true if the instance is the same type as the val.
// func isPointer(brt, irt reflect.Type, flex bool, level int) bool {
// 	if irt.Kind() != reflect.Ptr {
// 		return false
// 	}

// 	return isSameType(brt.Elem(), irt.Elem(), flex, level+1)
// }

// // isStruct returns true if the instance is the same type as the val.
// func isStruct(brt, irt reflect.Type, flex bool, level int) bool {
// 	// In flex mode, we can compare an object with an empty struct{},
// 	// or with a structure of the same type only.
// 	// Therefore, we check that brt has no fields.
// 	if flex && brt.NumField() == 0 && brt.Kind() == reflect.Struct {
// 		return true
// 	}

// 	// Checking if the name of the types are the same.
// 	if brt.Name() != irt.Name() {
// 		return false
// 	}

// 	if brt.NumField() != irt.NumField() {
// 		return false
// 	}

// 	// Checking if the fields of the types are the same.
// 	for i := 0; i < brt.NumField(); i++ {
// 		if brt.Field(i).Name != irt.Field(i).Name {
// 			return false
// 		}

// 		if !isSameType(brt.Field(i).Type, irt.Field(i).Type, false, level+1) {
// 			return false
// 		}
// 	}

// 	return true
// }

// // isFunc returns true if the instance is the same type as the val.
// func isFunc(brt, irt reflect.Type, flex bool, level int) bool {
// 	if flex && irt.Kind() == reflect.Func {
// 		return true
// 	}

// 	// Check if the function has the same number of arguments.
// 	if brt.NumIn() != irt.NumIn() {
// 		return false
// 	}

// 	// Check if the function has the same number of outputs.
// 	if brt.NumOut() != irt.NumOut() {
// 		return false
// 	}

// 	// Check if the function has the same variadic property.
// 	if brt.IsVariadic() != irt.IsVariadic() {
// 		return false
// 	}

// 	// Check if the function has the same arguments.
// 	for i := 0; i < brt.NumIn(); i++ {
// 		if !isSameType(brt.In(i), irt.In(i), flex, level+1) {
// 			return false
// 		}
// 	}

// 	// Check if the function has the same outputs.
// 	for i := 0; i < brt.NumOut(); i++ {
// 		if !isSameType(brt.Out(i), irt.Out(i), flex, level+1) {
// 			return false
// 		}
// 	}

// 	return true
// }

// // isChan returns true if the instance is the same type as the val.
// func isChan(brt, irt reflect.Type, flex bool, level int) bool {
// 	if flex && irt.Kind() == reflect.Chan {
// 		return true
// 	}

// 	// Check if the channel has the same direction.
// 	if brt.ChanDir() != irt.ChanDir() {
// 		return false
// 	}

// 	// Check if the channel has the same type.
// 	if !isSameType(brt.Elem(), irt.Elem(), flex, level+1) {
// 		return false
// 	}

// 	// Check if the channel has the same buffer size.
// 	if brt.ChanDir() == reflect.BothDir && brt.ChanDir() == irt.ChanDir() {
// 		if brt.ChanDir() != irt.ChanDir() {
// 			return false
// 		}
// 	}

// 	return true
// }

type kindGroup int

const (
	nilGroup kindGroup = iota
	strictGroup
	numberGroup
	stringGroup
	boolGroup
	intGroup
	uintGroup
	floatGroup
	complexGroup
	arrayGroup
	sliceGroup
	mapGroup
	structGroup
	funcGroup
	chanGroup
)

type kindHolder struct {
	brt  kindGroup
	irt  kindGroup
	prev *kindHolder
	next *kindHolder
}

// Instance returns true if the instance is the same type as the base.
func Instance[T any](obj interface{}, mode int) bool {
	var base T
	brt, irt := reflect.TypeOf(base), reflect.TypeOf(obj)

	holder := &kindHolder{}
	return isSameType(holder, brt, irt, mode, 0)
}

func typeOf(rt reflect.Type, mode int) (kindGroup, bool) {
	switch rt.Kind() {
	case reflect.String:
		return stringGroup
	case reflect.Bool:
		return boolGroup
	case reflect.Array:
		return arrayGroup
	case reflect.Slice:
		return sliceGroup
	case reflect.Map:
		return mapGroup
	case reflect.Struct:
		return structGroup
	case reflect.Func:
		return funcGroup
	case reflect.Chan:
		return chanGroup
	default:
		if mode < 0 {
			switch rt.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
				reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr,
				reflect.Float32, reflect.Float64, reflect.Complex64,
				reflect.Complex128:
				return numberGroup
			}
		} else if mode == 0 {
			switch rt.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
				reflect.Int64:
				return intGroup
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
				reflect.Uint64, reflect.Uintptr:
				return uintGroup
			case reflect.Float32, reflect.Float64:
				return floatGroup
			case reflect.Complex64, reflect.Complex128:
				return complexGroup
			}
		} else if mode > 0 {
			return strictGroup
		}
	}

	return -1
}

//func isSameGroup(rt reflect.Type, mode int) bool {
//
//}

func isSameType(holder *kindHolder, brt, irt reflect.Type, mode, level int) bool {
	holder.brt = getGroup(brt, mode)
	holder.irt = getGroup(irt, mode)

	// switch {
	// case mode < 0:
	// 	switch brt.Kind() {
	// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
	// 		reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16,
	// 		reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32,
	// 		reflect.Float64, reflect.Complex64, reflect.Complex128:
	// 		// If irt isn't a number.

	// 	case reflect.String:
	// 		// pass
	// 	case reflect.Bool:
	// 		// pass
	// 	default:
	// 		// Complex types.
	// 	}
	// case mode == 0:
	// 	switch brt.Kind() {
	// 	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
	// 		reflect.Int64:
	// 		// pass
	// 	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
	// 		reflect.Uint64, reflect.Uintptr:
	// 		// pass
	// 	case reflect.Float32, reflect.Float64:
	// 		// pass
	// 	case reflect.Complex64, reflect.Complex128:
	// 		// pass
	// 	case reflect.String:
	// 		// pass
	// 	case reflect.Bool:
	// 		// pass
	// 	default:
	// 		// Complex types.
	// 	}
	// case mode > 0:
	// 	if brt.Kind() != irt.Kind() {
	// 		return false
	// 	}
	// }

	// switch brt.Kind() {
	// case reflect.Slice, reflect.Array, reflect.Map:
	// 	return isContainer(brt, irt, mode)
	// case reflect.Ptr:
	// 	return isPointer(brt, irt, mode)
	// case reflect.Struct:
	// 	return isStruct(brt, irt, mode)
	// case reflect.Func:
	// 	return isFunc(brt, irt, mode)
	// case reflect.Chan:
	// 	return isChan(brt, irt, mode)
	// }

	return true
}

// func isContainer(brt, irt reflect.Type, strict int) bool {
// 	if strict == -1 && brt.Kind() == irt.Kind() {
// 		return true
// 	}

// 	if brt.Kind() == reflect.Array && (brt.Len() != irt.Len() || brt.Elem() != irt.Elem()) {
// 		return false
// 	}

// 	if brt.Kind() == reflect.Slice && brt.Elem() != irt.Elem() {
// 		return false
// 	}

// 	if brt.Kind() == reflect.Map && (brt.Key() != irt.Key() || brt.Elem() != irt.Elem()) {
// 		return false
// 	}

// 	return true
// }

// func isPointer(brt, irt reflect.Type, strict int) bool {
// 	return isSameType(brt.Elem(), irt.Elem(), strict)
// }

// func isStruct(brt, irt reflect.Type, strict int) bool {
// 	if strict == -1 && brt.NumField() == 0 && brt.Kind() == reflect.Struct {
// 		return true
// 	}

// 	if brt.Name() != irt.Name() || brt.NumField() != irt.NumField() {
// 		return false
// 	}

// 	for i := 0; i < brt.NumField(); i++ {
// 		if brt.Field(i).Name != irt.Field(i).Name || !isSameType(brt.Field(i).Type, irt.Field(i).Type, strict) {
// 			return false
// 		}
// 	}

// 	return true
// }

// func isFunc(brt, irt reflect.Type, strict int) bool {
// 	if strict == -1 && irt.Kind() == reflect.Func {
// 		return true
// 	}

// 	if brt.NumIn() != irt.NumIn() || brt.NumOut() != irt.NumOut() || brt.IsVariadic() != irt.IsVariadic() {
// 		return false
// 	}

// 	for i := 0; i < brt.NumIn(); i++ {
// 		if !isSameType(brt.In(i), irt.In(i), strict) {
// 			return false
// 		}
// 	}

// 	for i := 0; i < brt.NumOut(); i++ {
// 		if !isSameType(brt.Out(i), irt.Out(i), strict) {
// 			return false
// 		}
// 	}

// 	return true
// }

// func isChan(brt, irt reflect.Type, strict int) bool {
// 	if strict == -1 && irt.Kind() == reflect.Chan {
// 		return true
// 	}

// 	return brt.ChanDir() == irt.ChanDir() && isSameType(brt.Elem(), irt.Elem(), strict)
// }
