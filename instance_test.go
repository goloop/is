package is

import (
	"testing"
)

// TestNil tests the Nil function.
func TestNil(t *testing.T) {
	t.Run("Nil pointer", func(t *testing.T) {
		var a *int
		if !Nil(a) {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("Nil slice", func(t *testing.T) {
		var b []int
		if !Nil(b) {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("Non-nil channel", func(t *testing.T) {
		c := make(chan int)
		if Nil(c) {
			t.Error("Expected false, but got true")
		}
	})

	t.Run("Nil func", func(t *testing.T) {
		var d func()
		if !Nil(d) {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("Nil interface", func(t *testing.T) {
		var e interface{}
		if !Nil(e) {
			t.Error("Expected true, but got false")
		}
	})

	t.Run("Non-nil interface", func(t *testing.T) {
		var e interface{} = 123
		if Nil(e) {
			t.Error("Expected false, but got true")
		}
	})
}

// TestInstanceSimpleValue tests the Instance function.
func TestInstanceSimpleValue(t *testing.T) {
	t.Parallel()

	t.Run("var a int", func(t *testing.T) {
		t.Parallel()
		var a int
		if !Instance[int](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*int](a) || Instance[uint](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a int8", func(t *testing.T) {
		t.Parallel()
		var a int8
		if !Instance[int8](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*int8](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a int16", func(t *testing.T) {
		t.Parallel()
		var a int16
		if !Instance[int16](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*int16](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a int32", func(t *testing.T) {
		t.Parallel()
		var a int32
		if !Instance[int32](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*int32](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a int64", func(t *testing.T) {
		t.Parallel()
		var a int64
		if !Instance[int64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*int64](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a uint", func(t *testing.T) {
		t.Parallel()
		var a uint
		if !Instance[uint](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*uint](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a uint8", func(t *testing.T) {
		t.Parallel()
		var a uint8
		if !Instance[uint8](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*uint8](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a uint16", func(t *testing.T) {
		t.Parallel()
		var a uint16
		if !Instance[uint16](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*uint16](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a uint32", func(t *testing.T) {
		t.Parallel()
		var a uint32
		if !Instance[uint32](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*uint32](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a uint64", func(t *testing.T) {
		t.Parallel()
		var a uint64
		if !Instance[uint64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*uint64](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a uintptr", func(t *testing.T) {
		t.Parallel()
		var a uintptr
		if !Instance[uintptr](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*uintptr](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a float32", func(t *testing.T) {
		t.Parallel()
		var a float32
		if !Instance[float32](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*float32](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a float64", func(t *testing.T) {
		t.Parallel()
		var a float64
		if !Instance[float64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*float64](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a complex64", func(t *testing.T) {
		t.Parallel()
		var a complex64
		if !Instance[complex64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*complex64](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a complex128", func(t *testing.T) {
		t.Parallel()
		var a complex128
		if !Instance[complex128](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[*complex128](a) || Instance[int](a) || Instance[bool](a) {
			t.Errorf("expected false, but got true")
		}
	})
}

// TestInstanceSimplePointer tests the Instance function.
func TestInstanceSimplePointer(t *testing.T) {
	t.Parallel()

	t.Run("var a *int", func(t *testing.T) {
		t.Parallel()
		var a *int
		if !Instance[*int](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[int](a) || Instance[*uint](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *int8", func(t *testing.T) {
		t.Parallel()
		var a *int8
		if !Instance[*int8](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[int8](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *int16", func(t *testing.T) {
		t.Parallel()
		var a *int16
		if !Instance[*int16](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[int16](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *int32", func(t *testing.T) {
		t.Parallel()
		var a *int32
		if !Instance[*int32](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[int32](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *int64", func(t *testing.T) {
		t.Parallel()
		var a *int64
		if !Instance[*int64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[int64](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *uint", func(t *testing.T) {
		t.Parallel()
		var a *uint
		if !Instance[*uint](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[uint](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *uint8", func(t *testing.T) {
		t.Parallel()
		var a *uint8
		if !Instance[*uint8](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[uint8](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *uint16", func(t *testing.T) {
		t.Parallel()
		var a *uint16
		if !Instance[*uint16](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[uint16](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *uint32", func(t *testing.T) {
		t.Parallel()
		var a *uint32
		if !Instance[*uint32](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[uint32](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *uint64", func(t *testing.T) {
		t.Parallel()
		var a *uint64
		if !Instance[*uint64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[uint64](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *uintptr", func(t *testing.T) {
		t.Parallel()
		var a *uintptr
		if !Instance[*uintptr](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[uintptr](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *float32", func(t *testing.T) {
		t.Parallel()
		var a *float32
		if !Instance[*float32](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[float32](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *float64", func(t *testing.T) {
		t.Parallel()
		var a *float64
		if !Instance[*float64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[float64](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *complex64", func(t *testing.T) {
		t.Parallel()
		var a *complex64
		if !Instance[*complex64](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[complex64](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a *complex128", func(t *testing.T) {
		t.Parallel()
		var a *complex128
		if !Instance[*complex128](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[complex128](a) || Instance[*int](a) || Instance[*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

}

// TestInstanceContainerValue tests the Instance function.
func TestInstanceContainerValue(t *testing.T) {
	t.Parallel()

	t.Run("var a []int", func(t *testing.T) {
		t.Parallel()
		var a []int
		if !Instance[[]int](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[[]*int](a) || Instance[[]uint](a) || Instance[[]bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a [3]int", func(t *testing.T) {
		t.Parallel()
		var a [3]int
		if !Instance[[3]int](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[[3]*int](a) || Instance[[3]uint](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a map[string]int", func(t *testing.T) {
		t.Parallel()
		var a map[string]int
		if !Instance[map[string]int](a) {
			t.Errorf("expected true, but got false")
		}

		// Flexible mode.
		// The type of key or value does not matter.
		// Just checks if the object is a map.
		if Instance[map[string]uint](a) {
			t.Errorf("expected false, but got true")
		}

		// Strict mode.
		// Checks for type matching too.
		if Instance[map[string]uint](a, true) {
			t.Errorf("expected false, but got true")
		}
	})

}

// TestInstanceContainerPointer tests the Instance function.
func TestInstanceContainerPointer(t *testing.T) {
	t.Parallel()

	t.Run("var a []*int", func(t *testing.T) {
		t.Parallel()
		var a []*int
		if !Instance[[]*int](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[[]int](a) || Instance[[]*uint](a) || Instance[[]*bool](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a [3]*int", func(t *testing.T) {
		t.Parallel()
		var a [3]*int
		if !Instance[[3]*int](a) {
			t.Errorf("expected true, but got false")
		}

		if Instance[[3]int](a) || Instance[[3]*uint](a) {
			t.Errorf("expected false, but got true")
		}
	})

	t.Run("var a map[string]int", func(t *testing.T) {
		t.Parallel()
		var a map[string]*int
		if !Instance[map[string]*int](a) {
			t.Errorf("expected true, but got false")
		}

		// Flexible mode.
		// The type of key or value does not matter.
		// Just checks if the object is a map.
		if Instance[map[string]*uint](a) {
			t.Errorf("expected false, but got true")
		}

		// Strict mode.
		// Checks for type matching too.
		if Instance[map[string]*uint](a, true) {
			t.Errorf("expected false, but got true")
		}
	})

}
