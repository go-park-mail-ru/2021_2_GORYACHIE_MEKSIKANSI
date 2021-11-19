package Errors

import (
	test "2021_2_GORYACHIE_MEKSIKANSI/test"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestCheckErrorRestaurant(t *testing.T) {
	testTable := []struct {
		errorInput    Errors
		errorExpected string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Alias: RGetRestaurantsRestaurantsNotFound,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusNotFound) +
				",\"explain\":\"" + RGetRestaurantsRestaurantsNotFound + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Alias: RGetRestaurantsRestaurantsNotScan,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
	}

	for _, testCase := range testTable {
		errOut, resultOut, codeHTTP := CheckErrorRestaurant(&testCase.errorInput)
		assert.Equal(t, testCase.errorExpected, errOut.Error(),
			codeHTTP, testCase.codeHTTPExpected,
			fmt.Sprintf("Expected %s, %s, %d", testCase.errorExpected, testCase.resultExpected, testCase.codeHTTPExpected),
		)
		assert.Equal(t, testCase.resultExpected, string(resultOut),
			fmt.Sprintf("Expected %s, %s, %d", testCase.errorExpected, testCase.resultExpected, testCase.codeHTTPExpected),
		)
		assert.Equal(t, testCase.codeHTTPExpected, codeHTTP,
			fmt.Sprintf("Expected %s, %s, %d", testCase.errorExpected, testCase.resultExpected, testCase.codeHTTPExpected),
		)
	}

	var err error
	errOut, resultOut, codeHTTP := CheckErrorRestaurant(err)
	assert.Equal(t, nil, errOut,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, IntNil),
	)
	assert.Equal(t, []byte(nil), resultOut,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, IntNil),
	)
	assert.Equal(t, IntNil, codeHTTP,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, IntNil),
	)
}
