package assert

import (
	"testing"
	"time"
)

type wrappedLogf func(format string, args ...interface{})

func (f wrappedLogf) Errorf(format string, args ...interface{}) {
	f(format, args...)
}

// WrappedAssertions provides assertion methods against an 'actual' value. It
// exposes the methods at on WrappedAssertions but also under the Must field:
// which will ensure that if the assertion fails no more assertions will run for
// that test.
type WrappedAssertions struct {
	Wrapped
	Must Wrapped
}

// Wrapped provides assertion methods against an 'actual' value, reporting to
// the wrapped 't'.
type Wrapped struct {
	t      wrappedLogf
	actual interface{}
}

// Wrap provides a function which will then allow you to assert properties of
// the 'actual' value used.
func Wrap(t *testing.T) func(actual interface{}) *WrappedAssertions {
	return func(actual interface{}) *WrappedAssertions {
		return &WrappedAssertions{
			Wrapped: Wrapped{
				t:      wrappedLogf(t.Errorf),
				actual: actual,
			},
			Must: Wrapped{
				t:      wrappedLogf(t.Fatalf),
				actual: actual,
			},
		}
	}
}

// Fail marks the test as a failure, using the 'actual' value as the failure message.
//
//  assert := assert.Wrap(t)
//  assert("the test failed").Fail()
//
func (w *Wrapped) Fail(msgAndArgs ...interface{}) bool {
	value, _ := w.actual.(string)
	return Fail(w.t, value, msgAndArgs...)
}

// Condition uses the Comparison provided to 'actual' to assert a complex condition.
//
//   assert := assert.Wrap(t)
//   assert(func() bool { return true  }).Condition()
//
func (w *Wrapped) Condition(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(Comparison)
	if !ok {
		return Fail(w.t, "Condition called against a non-Comparison")
	}

	return Condition(w.t, value, msgAndArgs...)
}

// Contains asserts that the specified string contains the specified substring.
//
//    assert("Hello World").Contains("World", "But 'Hello World' does contain 'World'")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Contains(expected interface{}, msgAndArgs ...interface{}) bool {
	return Contains(w.t, w.actual, expected, msgAndArgs...)
}

// Empty asserts that the specified object is empty: i.e. nil, "", false, 0 or a
// slice with len == 0.
//
//   assert(obj).Empty()
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Empty(msgAndArgs ...interface{}) bool {
	return Empty(w.t, w.actual, msgAndArgs...)
}

// Equal asserts that two objects are equal.
//
//    assert(123).Equal(123, "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Equal(expected interface{}, msgAndArgs ...interface{}) bool {
	return Equal(w.t, expected, w.actual, msgAndArgs...)
}

// Equivalent asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert(int32(123)).Equivalent(uint32(123), "123 and 123 should be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Equivalent(expected interface{}, msgAndArgs ...interface{}) bool {
	return Equivalent(w.t, expected, w.actual, msgAndArgs...)
}

// Exactly asserts that two objects are equal is value and type.
//
//    assert(int64(123)).Exactly(int32(123), "123 and 123 should NOT be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Exactly(expected interface{}, msgAndArgs ...interface{}) bool {
	return Exactly(w.t, expected, w.actual, msgAndArgs...)
}

// False asserts that the specified value is true.
//
//    assert(myBool).False("myBool should be false")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) False(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(bool)
	if !ok {
		return Fail(w.t, "False called against a non-bool")
	}

	return False(w.t, value, msgAndArgs...)
}

// Implements asserts that an object is implemented by the specified interface.
//
//    assert(new(MyObject)).Implements((*MyInterface)(nil), "MyObject")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Implements(iface interface{}, msgAndArgs ...interface{}) bool {
	return Implements(w.t, iface, w.actual, msgAndArgs...)
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 assert(22/7.0).InDelta(math.Pi, 0.01)
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) InDelta(expected interface{}, delta float64, msgAndArgs ...interface{}) bool {
	return InDelta(w.t, expected, w.actual, delta, msgAndArgs...)
}

// InEpsilon asserts that expected and actual have a relative error less than
// epsilon
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) InEpsilon(expected interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	return InEpsilon(w.t, expected, w.actual, epsilon, msgAndArgs...)
}

// IsType asserts that the specified objects are of the same type.
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) IsType(expected interface{}, msgAndArgs ...interface{}) bool {
	return IsType(w.t, expected, w.actual, msgAndArgs...)
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    assert(mySlice).Len(3, "The size of slice is not 3")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Len(length int, msgAndArgs ...interface{}) bool {
	return Len(w.t, w.actual, length, msgAndArgs...)
}

// Nil asserts that the specified object is nil.
//
//    assert(err).Nil("err should be nothing")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Nil(msgAndArgs ...interface{}) bool {
	return Nil(w.t, w.actual, msgAndArgs...)
}

// NotContains asserts that the specified string does NOT contain the specified substring.
//
//    assert("Earth").NotContains("Hello World", "But 'Hello World' does NOT contain 'Earth'")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) NotContains(expected interface{}, msgAndArgs ...interface{}) bool {
	return NotContains(w.t, w.actual, expected, msgAndArgs...)
}

// NotEmpty asserts that the specified object is NOT empty: i.e. not nil, "",
// false, 0 or a slice with len == 0.
//
//   if assert(obj).NotEmpty() {
//     assert(obj[1]).Equal("two")
//   }
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) NotEmpty(msgAndArgs ...interface{}) bool {
	return NotEmpty(w.t, w.actual, msgAndArgs...)
}

// NotEqual asserts that the specified values are NOT equal.
//
//    assert(obj2).NotEqual(obj1, "two objects shouldn't be equal")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) NotEqual(expected interface{}, msgAndArgs ...interface{}) bool {
	return NotEqual(w.t, expected, w.actual, msgAndArgs...)
}

// NotNil asserts that the specified object is not nil.
//
//    assert(err).NotNil("err should be something")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) NotNil(msgAndArgs ...interface{}) bool {
	return NotNil(w.t, w.actual, msgAndArgs...)
}

// NotPanics asserts that the code inside the specified func does NOT panic.
//
//   assert(func(){ RemainCalm() }).NotPanics("Calling RemainCalm() should NOT panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) NotPanics(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(func())
	if !ok {
		return Fail(w.t, "NotPanics called against a non-func() ")
	}

	return NotPanics(w.t, value, msgAndArgs...)
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//   assert("it's starting").NotRegexp(regexp.MustCompile("starts"))
//   assert("it's not starting").NotRegexp("^start")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) NotRegexp(regex interface{}, msgAndArgs ...interface{}) bool {
	return NotRegexp(w.t, regex, w.actual, msgAndArgs...)
}

// Panics asserts that the code inside the specified func panics.
//
//   assert(func(){ GoCrazy() }).Panics("Calling GoCrazy() should panic")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Panics(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(func())
	if !ok {
		return Fail(w.t, "Panics called against a non-func() ")
	}

	return Panics(w.t, value, msgAndArgs...)
}

// Regexp asserts that a specified regexp matches a string.
//
//   assert("it's starting").Regexp(regexp.MustCompile("start"))
//   assert("it's not starting").Regexp("start...$")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) Regexp(regex interface{}, msgAndArgs ...interface{}) bool {
	return Regexp(w.t, regex, w.actual, msgAndArgs...)
}

// True asserts that the specified value is true.
//
//    assert(myBool).True("myBool should be true")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) True(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(bool)
	if !ok {
		return Fail(w.t, "True called against a non-bool")
	}

	return True(w.t, value, msgAndArgs...)
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   assert(time.Now()).WithinDuration(time.Now(), 10*time.Second, "The difference should not be more than 10s")
//
// Returns whether the assertion was successful (true) or not (false).
func (w *Wrapped) WithinDuration(expected time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(time.Time)
	if !ok {
		return Fail(w.t, "WithinDuration called against a non-time.Time")
	}

	return WithinDuration(w.t, expected, value, delta, msgAndArgs...)
}
