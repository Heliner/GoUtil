package GoUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Foo struct {
	ValFoo int
	ValBar string
}

func TestList2Map(t *testing.T) {

	// empty arr
	arr := make([]Foo, 0)
	res := List2Map(arr, "ValFoo", Foo{}.ValFoo)
	assert.Equal(t, 0, len(res))

	// not found field
	func() {
		defer func() {
			re := recover()
			assert.NotNil(t, re)
		}()

		arr = []Foo{
			{
				ValFoo: 1,
				ValBar: "bar1",
			}}
		_ = List2Map(arr, "ValFoo1", Foo{}.ValFoo)
	}()

	// found field
	arr = []Foo{
		{
			ValFoo: 1,
			ValBar: "bar1",
		},
		{
			ValFoo: 2,
			ValBar: "bar2",
		},
	}
	res = List2Map(arr, "ValFoo", Foo{}.ValFoo)
	assert.EqualValues(t, map[int]Foo{
		1: {
			1,
			"bar1",
		},
		2: {
			2,
			"bar2",
		},
	}, res)

}
