package fu

import (
	"gotest.tools/v3/assert"
	"testing"
)

type Option1 bool
type Option2 string
type Option3 int
type Option4 float64

func option1(o ...interface{}) bool {
	return Option(Option1(false), o).Bool()
}

func option2(o ...interface{}) string {
	return Option(Option2(""), o).String()
}

func option3(o ...interface{}) int {
	return int(Option(Option3(0), o).Int())
}

func Test_Option1(t *testing.T) {

	assert.Assert(t, option1(Option1(true)) == true)
	assert.Assert(t, option1(Option1(true), Option1(false)) == true)
	assert.Assert(t, option1(Option1(false), Option1(true)) == false)
	assert.Assert(t, option1(Option3(0)) == false)
	assert.Assert(t, option1() == false)

}

func Test_Option2(t *testing.T) {
	assert.Assert(t, option2(Option2("hello")) == "hello")
	assert.Assert(t, option2(Option1(false)) == "")
}

func Test_Option3(t *testing.T) {
	assert.Assert(t, option3(Option3(42)) == 42)
	assert.Assert(t, option3(Option1(false)) == 0)
}

func Test_Option4(t *testing.T) {
	opts := []interface{}{Option3(42), Option2("hello"), Option1(true), Option4(1.0)}
	assert.Assert(t, IntOption(Option3(0), opts) == 42)
	assert.Assert(t, StrOption(Option2(""), opts) == "hello")
	assert.Assert(t, FloatOption(Option4(0), opts) == 1.0)
	assert.Assert(t, BoolOption(Option1(false), opts) == true)
}
