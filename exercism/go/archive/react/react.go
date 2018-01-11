// Package react implements a set of interfaces to perform spreadsheet fun
package react

const testVersion = 5

type writeFunc func(int) int
type writeFuncDual func(int, int) int

// React is our default struct
type React struct{}

type cell struct {
	val int
	cb  []func()
	wf  writeFunc
	wfd writeFuncDual
	a   *cell
	b   *cell
}

// Cancel struct enables callback cancellation
type Cancel struct {
	c   *cell
	idx int
}

// New returns a new instance of the Reactor interface
func New() *React {
	return &React{}
}

// CreateInput is a bound method to the default React struct, creates a new input
func (r *React) CreateInput(in int) InputCell {
	return &cell{val: in}
}

// CreateCompute1 is bound to React, creates new input with one dep
func (r *React) CreateCompute1(a Cell, fun func(int) int) ComputeCell {
	return &cell{val: fun(a.Value()), wf: fun, a: a.(*cell)}
}

// CreateCompute2 is bound to React, creates new input with two deps
func (r *React) CreateCompute2(a, b Cell, fun func(int, int) int) ComputeCell {
	return &cell{val: fun(a.Value(), b.Value()), wfd: fun, a: a.(*cell), b: b.(*cell)}
}

// Value is bound to Cell, returns the value of the specific cell
func (c *cell) Value() int {
	if c.wfd != nil {
		c.val = c.wfd(c.a.Value(), c.b.Value())
	} else if c.wf != nil {
		c.val = c.wf(c.a.Value())
	}

	return c.val
}

// AddCallback appends a specific function to call when a cell changes
func (c *cell) AddCallback(fun func(int)) Canceler {
	var idx int
	tester := func() {
		t := c.val
		nt := c.Value()

		if t != nt || len(c.a.cb) > 1 {
			fun(nt)
		}
	}

	if c.a != nil {
		c.a.cb = append(c.a.cb, tester)
		idx = len(c.a.cb)
	}

	if c.b != nil {
		c.b.cb = append(c.b.cb, tester)
		idx = len(c.b.cb)
	}

	c.cb = append(c.cb, tester)

	return &Cancel{c, idx - 1}
}

// SetValue is bound to cell, updates the value for the provided cell
func (c *cell) SetValue(in int) {
	if c.val == in {
		return
	}

	c.val = in
	if len(c.cb) > 0 {
		for _, each := range c.cb {
			if each != nil {
				each()
			}
		}
	}
}

// Cancel removes a function from the callback tree
func (c Cancel) Cancel() {
	if c.c.a != nil {
		c.c.a.cb[c.idx] = nil
	}

	if c.c.b != nil {
		c.c.b.cb[c.idx] = nil
	}
}
