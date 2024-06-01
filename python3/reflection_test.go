package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectionBuiltins(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	test_len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(test_len))
}

func TestReflectionLocals(t *testing.T) {
	Py_Initialize()

	locals := PyEval_GetLocals()
	assert.Nil(t, locals)
}

func TestReflectionGlobals(t *testing.T) {
	Py_Initialize()

	globals := PyEval_GetGlobals()
	assert.Nil(t, globals)
}

func TestReflectionFuncName(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	test_len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(test_len))

	assert.Equal(t, "len", PyEval_GetFuncName(test_len))
}
func TestReflectionFuncDesc(t *testing.T) {
	Py_Initialize()

	builtins := PyEval_GetBuiltins()
	assert.NotNil(t, builtins)

	test_len := PyDict_GetItemString(builtins, "len")
	assert.True(t, PyCallable_Check(test_len))

	assert.Equal(t, "()", PyEval_GetFuncDesc(test_len))
}
