package is

import "reflect"

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

// // Nil returns true if the object is nil.
// func Nil(obj interface{}) bool {
// 	if obj == nil {
// 		return true
// 	}

// 	rv := reflect.ValueOf(obj)
// 	if rv.Kind() == reflect.Ptr {
// 		return rv.IsNil()
// 	}

// 	return false
// }

// Instance returns true if the instance is the same type as the base.
//
// Example usage:
//
//	var a []int
//	is.Instance[[]int](a) // true
//
//	var a *int
//	is.Instance[*int](a) // true
//	is.Instance[int](a)  // false
//	is.Nil(a)            // true
//
//	var a *int = new(int)
//	is.Instance[*int](a) // true
//	is.Instance[int](a)  // false
//	is.Nil(a)            // false
//
//	type A struct{
//	  N string
//	}
//	type B struct{
//	  N string
//	}
//
//	var a A
//	is.Instance[A](a)              // true
//	is.Instance[B](a)              // false
//	is.Instance[struct{}](a)       // true (flex mode)
//	is.Instance[struct{}](a, true) // false (strict mode)
//	is.Nil(a)                      // false
//
//	var a *A
//	is.Instance[*A](a)              // true
//	is.Instance[*B](a)              // false
//	is.Instance[*struct{}](a)       // true (flex mode)
//	is.Instance[*struct{}](a, true) // false (strict mode)
//	is.Nil(a)                       // true
func Instance[T any](obj interface{}, strict ...bool) bool {
	// Flexibility mode if all flex is true.
	flex := len(strict) == 0 // default is true
	for _, s := range strict {
		if !s {
			flex = true
			break
		}
	}

	// If the same type, return true.
	var base T
	brv, irv := reflect.ValueOf(base), reflect.ValueOf(obj)
	if !flex && brv.Kind() != irv.Kind() {
		return false
	}

	// Try detect the same type.
	brt := reflect.TypeOf(base)
	irt := reflect.TypeOf(obj)
	return isSameType(brt, irt, flex)
}

// isSameType returns true if the instance is the same type as the val.
func isSameType(brt, irt reflect.Type, flex bool) bool {
	if brt.Kind() != irt.Kind() {
		return false
	}

	switch brt.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		return isContainer(brt, irt, flex)
	case reflect.Ptr:
		return isPointer(brt, irt, flex)
	case reflect.Struct:
		return isStruct(brt, irt, flex)
	}

	return true
}

// // isContainer returns true if the instance is the same type as the val.
// func isContainer(brt, irt reflect.Type, flex bool) bool {
// 	if !flex && brt.Len() != irt.Len() {
// 		return false
// 	}
// 	if brt.Elem().Kind() == reflect.Interface {
// 		return true
// 	}
// 	return isSameType(brt.Elem(), irt.Elem(), flex)
// }

// isContainer returns true if the instance is the same type as the val.
func isContainer(brt, irt reflect.Type, flex bool) bool {
	// Check the Kind of the reflect.Type before calling Len
	if brt.Kind() == reflect.Array || brt.Kind() == reflect.Slice {
		if !flex && brt.Len() != irt.Len() {
			return false
		}
	}

	if brt.Elem().Kind() == reflect.Interface {
		return true
	}

	return isSameType(brt.Elem(), irt.Elem(), flex)
}

func isPointer(brt, irt reflect.Type, flex bool) bool {
	if irt.Kind() != reflect.Ptr {
		return false
	}

	return isSameType(brt.Elem(), irt.Elem(), flex)
}

// isStruct returns true if the instance is the same type as the val.
func isStruct(brt, irt reflect.Type, flex bool) bool {
	if flex && brt.Kind() == reflect.Struct {
		return true
	}

	// Checking if the name of the types are the same.
	if brt.Name() != irt.Name() {
		return false
	}

	if brt.NumField() != irt.NumField() {
		return false
	}

	for i := 0; i < brt.NumField(); i++ {
		if brt.Field(i).Name != irt.Field(i).Name {
			return false
		}

		if !isSameType(brt.Field(i).Type, irt.Field(i).Type, false) {
			return false
		}
	}

	return true
}
