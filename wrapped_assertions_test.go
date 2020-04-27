package assert

import (
	"errors"
	"regexp"
	"testing"
	"time"
)

func TestWrappedImplements(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(new(AssertionTesterConformingObject)).Implements((*AssertionTesterInterface)(nil)) {
		t.Error("Implements method should return true: AssertionTesterConformingObject implements AssertionTesterInterface")
	}
	if assert(new(AssertionTesterNonConformingObject)).Implements((*AssertionTesterInterface)(nil)) {
		t.Error("Implements method should return false: AssertionTesterNonConformingObject does not implements AssertionTesterInterface")
	}
}

func TestWrappedIsType(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(new(AssertionTesterConformingObject)).IsType(new(AssertionTesterConformingObject)) {
		t.Error("IsType should return true: AssertionTesterConformingObject is the same type as AssertionTesterConformingObject")
	}
	if assert(new(AssertionTesterNonConformingObject)).IsType(new(AssertionTesterConformingObject)) {
		t.Error("IsType should return false: AssertionTesterConformingObject is not the same type as AssertionTesterNonConformingObject")
	}
}

func TestWrappedEqual(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert("Hello World").Equal("Hello World") {
		t.Error("Equal should return true")
	}
	if !assert(123).Equal(123) {
		t.Error("Equal should return true")
	}
	if !assert(123.5).Equal(123.5) {
		t.Error("Equal should return true")
	}
	if !assert([]byte("Hello World")).Equal([]byte("Hello World")) {
		t.Error("Equal should return true")
	}
	if !assert(nil).Equal(nil) {
		t.Error("Equal should return true")
	}
}

func TestWrappedEquivalent(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(int32(10)).Equivalent(uint32(10)) {
		t.Error("Equivalent should return true")
	}
}

func TestWrappedNotNil(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(new(AssertionTesterConformingObject)).NotNil() {
		t.Error("NotNil should return true: object is not nil")
	}
	if assert(nil).NotNil() {
		t.Error("NotNil should return false: object is nil")
	}
}

func TestWrappedNil(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(nil).Nil() {
		t.Error("Nil should return true: object is nil")
	}
	if assert(new(AssertionTesterConformingObject)).Nil() {
		t.Error("Nil should return false: object is not nil")
	}
}

func TestWrappedTrue(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(true).True() {
		t.Error("True should return true")
	}
	if assert(false).True() {
		t.Error("True should return false")
	}
}

func TestWrappedFalse(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(false).False() {
		t.Error("False should return true")
	}
	if assert(true).False() {
		t.Error("False should return false")
	}
}

func TestWrappedExactly(t *testing.T) {
	assert := Wrap(new(testing.T))

	a := float32(1)
	b := float64(1)
	c := float32(1)
	d := float32(2)

	if assert(b).Exactly(a) {
		t.Error("Exactly should return false")
	}
	if assert(d).Exactly(a) {
		t.Error("Exactly should return false")
	}
	if !assert(c).Exactly(a) {
		t.Error("Exactly should return true")
	}

	if assert(a).Exactly(nil) {
		t.Error("Exactly should return false")
	}
	if assert(nil).Exactly(a) {
		t.Error("Exactly should return false")
	}
}

func TestWrappedNotEqual(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert("Hello World!").NotEqual("Hello World") {
		t.Error("NotEqual should return true")
	}
	if !assert(1234).NotEqual(123) {
		t.Error("NotEqual should return true")
	}
	if !assert(123.55).NotEqual(123.5) {
		t.Error("NotEqual should return true")
	}
	if !assert([]byte("Hello World!")).NotEqual([]byte("Hello World")) {
		t.Error("NotEqual should return true")
	}
	if !assert(new(AssertionTesterConformingObject)).NotEqual(nil) {
		t.Error("NotEqual should return true")
	}
}

func TestWrappedContains(t *testing.T) {
	assert := Wrap(new(testing.T))
	list := []string{"Foo", "Bar"}

	if !assert("Hello World").Contains("Hello") {
		t.Error("Contains should return true: \"Hello World\" contains \"Hello\"")
	}
	if assert("Hello World").Contains("Salut") {
		t.Error("Contains should return false: \"Hello World\" does not contain \"Salut\"")
	}

	if !assert(list).Contains("Foo") {
		t.Error("Contains should return true: \"[\"Foo\", \"Bar\"]\" contains \"Foo\"")
	}
	if assert(list).Contains("Salut") {
		t.Error("Contains should return false: \"[\"Foo\", \"Bar\"]\" does not contain \"Salut\"")
	}
}
func TestWrappedNotContains(t *testing.T) {
	assert := Wrap(new(testing.T))
	list := []string{"Foo", "Bar"}

	if !assert("Hello World").NotContains("Hello!") {
		t.Error("NotContains should return true: \"Hello World\" does not contain \"Hello!\"")
	}
	if assert("Hello World").NotContains("Hello") {
		t.Error("NotContains should return false: \"Hello World\" contains \"Hello\"")
	}

	if !assert(list).NotContains("Foo!") {
		t.Error("NotContains should return true: \"[\"Foo\", \"Bar\"]\" does not contain \"Foo!\"")
	}
	if assert(list).NotContains("Foo") {
		t.Error("NotContains should return false: \"[\"Foo\", \"Bar\"]\" contains \"Foo\"")
	}
}

func TestWrappedPanics(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(func() { panic("Panic!") }).Panics() {
		t.Error("Panics should return true")
	}

	if assert(func() {}).Panics() {
		t.Error("Panics should return false")
	}
}

func TestWrappedNotPanics(t *testing.T) {
	assert := Wrap(new(testing.T))

	if !assert(func() {}).NotPanics() {
		t.Error("NotPanics should return true")
	}

	if assert(func() { panic("Panic!") }).NotPanics() {
		t.Error("NotPanics should return false")
	}
}

func TestWrappedEmpty(t *testing.T) {
	assert := New(t)
	mockAssert := Wrap(new(testing.T))

	assert.True(mockAssert("").Empty(), "Empty string is empty")
	assert.True(mockAssert(nil).Empty(), "Nil is empty")
	assert.True(mockAssert([]string{}).Empty(), "Empty string array is empty")
	assert.True(mockAssert(0).Empty(), "Zero int value is empty")
	assert.True(mockAssert(false).Empty(), "False value is empty")

	assert.False(mockAssert("something").Empty(), "Non Empty string is not empty")
	assert.False(mockAssert(errors.New("something")).Empty(), "Non nil object is not empty")
	assert.False(mockAssert([]string{"something"}).Empty(), "Non empty string array is not empty")
	assert.False(mockAssert(1).Empty(), "Non-zero int value is not empty")
	assert.False(mockAssert(true).Empty(), "True value is not empty")
}

func TestWrappedNotEmpty(t *testing.T) {
	assert := New(t)
	mockAssert := Wrap(new(testing.T))

	assert.False(mockAssert("").NotEmpty(), "Empty string is empty")
	assert.False(mockAssert(nil).NotEmpty(), "Nil is empty")
	assert.False(mockAssert([]string{}).NotEmpty(), "Empty string array is empty")
	assert.False(mockAssert(0).NotEmpty(), "Zero int value is empty")
	assert.False(mockAssert(false).NotEmpty(), "False value is empty")

	assert.True(mockAssert("something").NotEmpty(), "Non Empty string is not empty")
	assert.True(mockAssert(errors.New("something")).NotEmpty(), "Non nil object is not empty")
	assert.True(mockAssert([]string{"something"}).NotEmpty(), "Non empty string array is not empty")
	assert.True(mockAssert(1).NotEmpty(), "Non-zero int value is not empty")
	assert.True(mockAssert(true).NotEmpty(), "True value is not empty")
}

func TestWrappedLen(t *testing.T) {
	assert := New(t)
	mockAssert := Wrap(new(testing.T))

	assert.False(mockAssert(nil).Len(0), "nil does not have length")
	assert.False(mockAssert(0).Len(0), "int does not have length")
	assert.False(mockAssert(true).Len(0), "true does not have length")
	assert.False(mockAssert(false).Len(0), "false does not have length")
	assert.False(mockAssert('A').Len(0), "Rune does not have length")
	assert.False(mockAssert(struct{}{}).Len(0), "Struct does not have length")

	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3

	cases := []struct {
		v interface{}
		l int
	}{
		{[]int{1, 2, 3}, 3},
		{[...]int{1, 2, 3}, 3},
		{"ABC", 3},
		{map[int]int{1: 2, 2: 4, 3: 6}, 3},
		{ch, 3},

		{[]int{}, 0},
		{map[int]int{}, 0},
		{make(chan int), 0},

		{[]int(nil), 0},
		{map[int]int(nil), 0},
		{(chan int)(nil), 0},
	}

	for _, c := range cases {
		assert.True(mockAssert(c.v).Len(c.l), "%#v have %d items", c.v, c.l)
	}
}

func TestWrappedWithinDuration(t *testing.T) {
	assert := New(t)
	mockAssert := Wrap(new(testing.T))
	a := time.Now()
	b := a.Add(10 * time.Second)

	assert.True(mockAssert(b).WithinDuration(a, 10*time.Second), "A 10s difference is within a 10s time difference")
	assert.True(mockAssert(a).WithinDuration(a, 10*time.Second), "A 10s difference is within a 10s time difference")

	assert.False(mockAssert(b).WithinDuration(a, 9*time.Second), "A 10s difference is not within a 9s time difference")
	assert.False(mockAssert(a).WithinDuration(b, 9*time.Second), "A 10s difference is not within a 9s time difference")

	assert.False(mockAssert(b).WithinDuration(a, -9*time.Second), "A 10s difference is not within a 9s time difference")
	assert.False(mockAssert(a).WithinDuration(b, -9*time.Second), "A 10s difference is not within a 9s time difference")

	assert.False(mockAssert(b).WithinDuration(a, -11*time.Second), "A 10s difference is not within a 9s time difference")
	assert.False(mockAssert(a).WithinDuration(b, -11*time.Second), "A 10s difference is not within a 9s time difference")
}

func TestWrappedInDelta(t *testing.T) {
	assert := Wrap(new(testing.T))

	True(t, assert(1).InDelta(1.001, 0.01), "|1.001 - 1| <= 0.01")
	True(t, assert(1.001).InDelta(1, 0.01), "|1 - 1.001| <= 0.01")
	True(t, assert(2).InDelta(1, 1), "|1 - 2| <= 1")
	False(t, assert(2).InDelta(1, 0.5), "Expected |1 - 2| <= 0.5 to fail")
	False(t, assert(1).InDelta(2, 0.5), "Expected |2 - 1| <= 0.5 to fail")
	False(t, assert(nil).InDelta("", 1), "Expected non numerals to fail")

	cases := []struct {
		a, b  interface{}
		delta float64
	}{
		{uint8(2), uint8(1), 1},
		{uint16(2), uint16(1), 1},
		{uint32(2), uint32(1), 1},
		{uint64(2), uint64(1), 1},

		{int(2), int(1), 1},
		{int8(2), int8(1), 1},
		{int16(2), int16(1), 1},
		{int32(2), int32(1), 1},
		{int64(2), int64(1), 1},

		{float32(2), float32(1), 1},
		{float64(2), float64(1), 1},
	}

	for _, tc := range cases {
		True(t, assert(tc.b).InDelta(tc.a, tc.delta), "Expected |%V - %V| <= %v", tc.a, tc.b, tc.delta)
	}
}

func TestWrappedInEpsilon(t *testing.T) {
	assert := Wrap(new(testing.T))

	cases := []struct {
		a, b    interface{}
		epsilon float64
	}{
		{uint8(2), uint16(2), .001},
		{2.1, 2.2, 0.1},
		{2.2, 2.1, 0.1},
		{-2.1, -2.2, 0.1},
		{-2.2, -2.1, 0.1},
		{uint64(100), uint8(101), 0.01},
		{0.1, -0.1, 2},
	}

	for _, tc := range cases {
		True(t, assert(tc.b).InEpsilon(tc.a, tc.epsilon, "Expected %V and %V to have a relative difference of %v", tc.a, tc.b, tc.epsilon))
	}

	cases = []struct {
		a, b    interface{}
		epsilon float64
	}{
		{uint8(2), int16(-2), .001},
		{uint64(100), uint8(102), 0.01},
		{2.1, 2.2, 0.001},
		{2.2, 2.1, 0.001},
		{2.1, -2.2, 1},
		{2.1, "bla-bla", 0},
		{0.1, -0.1, 1.99},
	}

	for _, tc := range cases {
		False(t, assert(tc.b).InEpsilon(tc.a, tc.epsilon, "Expected %V and %V to have a relative difference of %v", tc.a, tc.b, tc.epsilon))
	}
}

func TestWrappedRegexp(t *testing.T) {
	assert := Wrap(new(testing.T))

	cases := []struct {
		rx, str string
	}{
		{"^start", "start of the line"},
		{"end$", "in the end"},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12.34"},
	}

	for _, tc := range cases {
		True(t, assert(tc.str).Regexp(tc.rx))
		True(t, assert(tc.str).Regexp(regexp.MustCompile(tc.rx)))
		False(t, assert(tc.str).NotRegexp(tc.rx))
		False(t, assert(tc.str).NotRegexp(regexp.MustCompile(tc.rx)))
	}

	cases = []struct {
		rx, str string
	}{
		{"^asdfastart", "Not the start of the line"},
		{"end$", "in the end."},
		{"[0-9]{3}[.-]?[0-9]{2}[.-]?[0-9]{2}", "My phone number is 650.12a.34"},
	}

	for _, tc := range cases {
		False(t, assert(tc.str).Regexp(tc.rx), "Expected \"%s\" to not match \"%s\"", tc.rx, tc.str)
		False(t, assert(tc.str).Regexp(regexp.MustCompile(tc.rx)))
		True(t, assert(tc.str).NotRegexp(tc.rx))
		True(t, assert(tc.str).NotRegexp(regexp.MustCompile(tc.rx)))
	}
}
