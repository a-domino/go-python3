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

func TestComplex(t *testing.T) {
	Py_Initialize()

	test_real := 2.
	test_imaginary := 5.

	test_complex := PyComplex_FromDoubles(test_real, test_imaginary)
	assert.True(t, PyComplex_Check(test_complex))
	assert.True(t, PyComplex_CheckExact(test_complex))
	defer test_complex.DecRef()

	assert.Equal(t, test_real, PyComplex_RealAsDouble(test_complex))
	assert.Equal(t, test_imaginary, PyComplex_ImagAsDouble(test_complex))
}
