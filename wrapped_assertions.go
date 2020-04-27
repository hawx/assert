package assert

import (
	"time"
)

type WrappedAssertions struct {
	Wrapped
	Must Wrapped
}

type Wrapped struct {
	t      TestingT
	actual interface{}
	fails  bool
}

func Wrap(t TestingT) func(actual interface{}) *WrappedAssertions {
	return func(actual interface{}) *WrappedAssertions {
		return &WrappedAssertions{
			Wrapped: Wrapped{
				t:      t,
				actual: actual,
				fails:  false,
			},
			Must: Wrapped{
				t:      t,
				actual: actual,
				fails:  true,
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
	Fail(w.t, value, msgAndArgs...)
	if w.fails {
		w.t.FailNow()
	}
	return false
}

// Condition uses the Comparison provided to 'actual' to assert a complex condition.
//
//   assert := assert.Wrap(t)
//   assert(func() bool { return true  }).Condition()
//
func (w *Wrapped) Condition(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(Comparison)
	if !ok {
		Fail(w.t, "Condition called against a non-Comparison")
		if w.fails {
			w.t.FailNow()
		}
		return false
	}

	ok = Condition(w.t, value, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Contains(expected interface{}, msgAndArgs ...interface{}) bool {
	ok := Contains(w.t, w.actual, expected, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Empty(msgAndArgs ...interface{}) bool {
	ok := Empty(w.t, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Equal(expected interface{}, msgAndArgs ...interface{}) bool {
	ok := Equal(w.t, expected, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Equivalent(expected interface{}, msgAndArgs ...interface{}) bool {
	ok := Equivalent(w.t, expected, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Exactly(expected interface{}, msgAndArgs ...interface{}) bool {
	ok := Exactly(w.t, expected, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) False(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(bool)
	if !ok {
		Fail(w.t, "False called against a non-bool")
		if w.fails {
			w.t.FailNow()
		}
		return false
	}

	ok = False(w.t, value, msgAndArgs...)

	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Implements(iface interface{}, msgAndArgs ...interface{}) bool {
	ok := Implements(w.t, iface, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) InDelta(expected interface{}, delta float64, msgAndArgs ...interface{}) bool {
	ok := InDelta(w.t, expected, w.actual, delta, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) InEpsilon(expected interface{}, epsilon float64, msgAndArgs ...interface{}) bool {
	ok := InEpsilon(w.t, expected, w.actual, epsilon, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) IsType(expected interface{}, msgAndArgs ...interface{}) bool {
	ok := IsType(w.t, expected, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Len(length int, msgAndArgs ...interface{}) bool {
	ok := Len(w.t, w.actual, length, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Nil(msgAndArgs ...interface{}) bool {
	ok := Nil(w.t, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) NotContains(expected interface{}, msgAndArgs ...interface{}) bool {
	ok := NotContains(w.t, w.actual, expected, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) NotEmpty(msgAndArgs ...interface{}) bool {
	ok := NotEmpty(w.t, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) NotEqual(expected interface{}, msgAndArgs ...interface{}) bool {
	ok := NotEqual(w.t, expected, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) NotNil(msgAndArgs ...interface{}) bool {
	ok := NotNil(w.t, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) NotPanics(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(func())
	if !ok {
		Fail(w.t, "NotPanics called against a non-func() ")
		if w.fails {
			w.t.FailNow()
		}
		return false
	}

	ok = NotPanics(w.t, value, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) NotRegexp(regex interface{}, msgAndArgs ...interface{}) bool {
	ok := NotRegexp(w.t, regex, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Panics(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(func())
	if !ok {
		Fail(w.t, "Panics called against a non-func() ")
		if w.fails {
			w.t.FailNow()
		}
		return false
	}

	ok = Panics(w.t, value, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) Regexp(regex interface{}, msgAndArgs ...interface{}) bool {
	ok := Regexp(w.t, regex, w.actual, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) True(msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(bool)
	if !ok {
		Fail(w.t, "True called against a non-bool")
		if w.fails {
			w.t.FailNow()
		}
		return false
	}

	ok = True(w.t, value, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}

func (w *Wrapped) WithinDuration(expected time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {
	value, ok := w.actual.(time.Time)
	if !ok {
		Fail(w.t, "WithinDuration called against a non-time.Time")
		if w.fails {
			w.t.FailNow()
		}
		return false
	}

	ok = WithinDuration(w.t, expected, value, delta, msgAndArgs...)
	if !ok && w.fails {
		w.t.FailNow()
	}
	return ok
}
