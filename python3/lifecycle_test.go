/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialization(t *testing.T) {

	Py_Initialize()
	assert.True(t, Py_IsInitialized())
	Py_Finalize()
	assert.False(t, Py_IsInitialized())

}

func TestInitializationEx(t *testing.T) {

	Py_Initialize()
	assert.True(t, Py_IsInitialized())
	assert.Zero(t, Py_FinalizeEx())
	assert.False(t, Py_IsInitialized())

}

func TestProgramName(t *testing.T) {
	Py_Finalize()

	defaultName, err := Py_GetProgramName()
	defer func(name string) {
		err = Py_SetProgramName(name)
		if err != nil {
			t.Logf("Py_SetProgramName to default name failed: %#v\n", err)
		}
	}(defaultName)

	assert.Nil(t, err)
	name := "py3é"
	if err = Py_SetProgramName(name); err != nil {
		t.Logf("Py_SetProgramName to name failed: %#v\n", err)
	}
	newName, err := Py_GetProgramName()
	assert.Nil(t, err)
	assert.Equal(t, name, newName)

}

func TestPrefix(t *testing.T) {
	prefix, err := Py_GetPrefix()
	assert.Nil(t, err)
	assert.IsType(t, "", prefix)

}

func TestExecPrefix(t *testing.T) {
	execPrefix, err := Py_GetExecPrefix()
	assert.Nil(t, err)
	assert.IsType(t, "", execPrefix)

}

func TestProgramFullPath(t *testing.T) {
	programFullPath, err := Py_GetProgramFullPath()
	assert.Nil(t, err)
	assert.IsType(t, "", programFullPath)

}

func TestPath(t *testing.T) {
	Py_Finalize()

	defaultPath, err := Py_GetPath()
	defer func(path string) {
		if err = Py_SetPath(path); err != nil {
			t.Logf("Py_SetPath to default path failed: %#v\n", err)
		}
	}(defaultPath)

	assert.Nil(t, err)
	name := "test_path"
	if err := Py_SetPath(name); err != nil {
		t.Logf("Py_SetPath to path failed: %#v\n", err)
	}
	newName, err := Py_GetPath()
	assert.Nil(t, err)
	assert.Equal(t, name, newName)

}

func TestVersion(t *testing.T) {
	version := Py_GetVersion()
	assert.IsType(t, "", version)
}

func TestPlatform(t *testing.T) {
	platform := Py_GetPlatform()
	assert.IsType(t, "", platform)
}

func TestCopyright(t *testing.T) {
	copyright := Py_GetCopyright()
	assert.IsType(t, "", copyright)
}

func TestCompiler(t *testing.T) {
	compiler := Py_GetCompiler()
	assert.IsType(t, "", compiler)
}

func TestBuildInfo(t *testing.T) {
	buildInfo := Py_GetBuildInfo()
	assert.IsType(t, "", buildInfo)
}

func TestPythonHome(t *testing.T) {
	name := "new_home"

	defaultHome, err := Py_GetPythonHome()
	defer func(home string) {
		if err = Py_SetPythonHome(home); err != nil {
			t.Logf("Py_SetPythonHome to default home failed: %#v\n", err)
		}
	}(defaultHome)

	assert.Nil(t, err)
	if err = Py_SetPythonHome(name); err != nil {
		t.Logf("Py_SetPythonHome to default home failed: %#v\n", err)
	}
	newName, err := Py_GetPythonHome()
	assert.Nil(t, err)
	assert.Equal(t, name, newName)
}

func TestSetArgv(t *testing.T) {
	Py_Initialize()

	if err := PySys_SetArgv([]string{"test.py"}); err != nil {
		t.Logf("PySys_SetArgv to test.py failed: %#v\n", err)
	}

	argv := PySys_GetObject("argv")
	assert.Equal(t, 1, PyList_Size(argv))
	assert.Equal(t, "test.py", PyUnicode_AsUTF8(PyList_GetItem(argv, 0)))

	Py_Finalize()
}

func TestSetArgvEx(t *testing.T) {
	Py_Initialize()

	if err := PySys_SetArgvEx([]string{"test.py"}, false); err != nil {
		t.Logf("PySys_SetArgvEx to test.py failed: %#v\n", err)
	}

	argv := PySys_GetObject("argv")
	assert.Equal(t, 1, PyList_Size(argv))
	assert.Equal(t, "test.py", PyUnicode_AsUTF8(PyList_GetItem(argv, 0)))

	Py_Finalize()
}
