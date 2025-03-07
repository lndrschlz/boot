/*
 * Copyright (c) 2021 boot-go
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package boot

import (
	"reflect"
	"testing"
)

func TestBootWithWire(t *testing.T) {
	setupTest()

	t1 := &testStruct1{}
	t2 := &testStruct2{}
	t3 := &testStruct3{}
	t4 := &testStruct4{}
	t5 := &testStruct5{}
	t6 := &testStruct6{}
	t7 := &testStruct7{}
	t8 := &testStruct8{}
	t9 := &testStruct9{}
	t10 := &testStruct10{}
	t11 := &testStruct11{}
	t12 := &testStruct12{}
	controls := []Component{t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12}

	registry := newRegistry()
	for _, control := range controls {
		registry.addEntry(DefaultName, false, control)
	}
	registry.addEntry("test", false, t1)

	getEntry := func(c *Component) *entry {
		cmpName := QualifiedName(*c)
		return registry.entries[cmpName][DefaultName]
	}

	tests := []struct {
		name       string
		controller Component
		expected   Component
		err        string
	}{
		{
			name:       "No injection",
			controller: t1,
			expected:   &testStruct1{},
		},
		{
			name:       "Single injection",
			controller: t2,
			expected: &testStruct2{
				F: t1,
			},
		},
		{
			name:       "Multiple injections",
			controller: t3,
			expected: &testStruct3{
				F: t1,
				G: t2,
			},
		},
		{
			name:       "Failed injection into unexported variable",
			controller: t4,
			err:        "Error dependency value cannot be set into <testStruct4.f>",
		},
		{
			name:       "Injection into interface",
			controller: t5,
			expected: &testStruct5{
				F: t1,
			},
		},
		{
			name:       "Failed injection not unique",
			controller: t6,
			err:        "Error multiple dependency values found for <default:testStruct6.F>",
		},
		{
			name:       "Failed injection for unrecognized component",
			controller: t7,
			err:        "Error dependency value not found for <default:testStruct7.F>",
		},
		{
			name:       "Failed injection non pointer receiver",
			controller: t8,
			err:        "Error dependency field is not a pointer receiver <testStruct8.F>",
		},
		{
			name:       "Single injection by name",
			controller: t9,
			expected: &testStruct9{
				F: t1,
			},
		},
		{
			name:       "Single injection by unknown name",
			controller: t10,
			err:        "Error dependency value not found for <unknown:testStruct10.F>",
		},
		{
			name:       "Single injection by unknown name",
			controller: t11,
			err:        "Error field contains unparsable tag  <testStruct11.F `wire,name:`>",
		},
		{
			name:       "Traverse injection failed",
			controller: t12,
			err:        "Error field contains unparsable tag  <testStruct11.F `wire,name:`>",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := resolveDependency(getEntry(&test.controller), registry)
			if len(test.err) == 0 {
				if err != nil {
					t.Fail()
				}
				if !reflect.DeepEqual(test.controller, test.expected) {
					t.Fail()
				}
			} else {
				if err == nil || err.Error() != test.err {
					t.Fatal(err.Error())
				}
			}
		})
	}
}

type testInterface1 interface {
	do1()
}

type testInterface2 interface {
	do2()
}

type testStruct1 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
}

func (t testStruct1) do1() {

}

func (t testStruct1) Init() {

}

func (t testStruct1) Start() {

}

func (t testStruct1) Stop() {

}

type testStruct2 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F *testStruct1 `boot:"wire"`
}

func (t testStruct2) do2() {

}

func (t testStruct2) Init() {

}

func (t testStruct2) Start() {

}

func (t testStruct2) Stop() {

}

type testStruct3 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F *testStruct1 `boot:"wire"`
	G *testStruct2 `boot:"wire"`
}

func (t testStruct3) Init() {

}

func (t testStruct3) Start() {

}

func (t testStruct3) Stop() {

}

type testStruct4 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	f *testStruct1 `boot:"wire"`
}

func (t testStruct4) Init() {

}

func (t testStruct4) Start() {

}

func (t testStruct4) Stop() {

}

type testStruct5 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F testInterface1 `boot:"wire"`
}

func (t testStruct5) Init() {

}

func (t testStruct5) Start() {

}

func (t testStruct5) Stop() {

}

type testStruct6 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F testInterface2 `boot:"wire"`
}

func (t testStruct6) do2() {

}

func (t testStruct6) Init() {

}

func (t testStruct6) Start() {

}

func (t testStruct6) Stop() {

}

type testStruct7 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F *string `boot:"wire"`
}

func (t testStruct7) Init() {

}

func (t testStruct7) Start() {

}

func (t testStruct7) Stop() {

}

type testStruct8 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F testStruct1 `boot:"wire"`
}

func (t testStruct8) Init() {

}

func (t testStruct8) Start() {

}

func (t testStruct8) Stop() {

}

type testStruct9 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F *testStruct1 `boot:"wire,name:test"`
}

func (t testStruct9) Init() {

}

type testStruct10 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F *testStruct1 `boot:"wire,name:unknown"`
}

func (t testStruct10) Init() {

}

type testStruct11 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F *testStruct1 `boot:"wire,name:"`
}

func (t testStruct11) Init() {}

type testStruct12 struct {
	a int
	B int
	c string
	d interface{}
	e []interface{}
	F *testStruct11 `boot:"wire,name:default"`
}

func (t testStruct12) Init() {}
