package Test

import (
	checkErr "2021_2_GORYACHIE_MEKSIKANSI/Errors"
	"net/http"
	"testing"
	"time"
)


func TestCheckErrorSignUp(t *testing.T) {

	testTable := []struct{
		errorInput checkErr.Errors
		errorExpected string
		resultExpected string
		codeHTTPExpected int
	} {
		{
			errorInput: checkErr.Errors{
				Text: checkErr.ErrGeneralInfoUnique,
				Time: time.Now(),
			},
			errorExpected: checkErr.ErrCheck,
			resultExpected: "{\"status\":409,\"explain\":\"Телефон или Email уже зарегистрирован\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: checkErr.Errors{
				Text: checkErr.ErrGeneralInfoUnique,
				Time: time.Now(),
			},
			errorExpected: checkErr.ErrCheck,
			resultExpected: "{\"status\":409,\"explain\":\"Телефон или Email уже зарегистрирован\"}",
			codeHTTPExpected: http.StatusOK,
		},
	}

	for _, testCase := range testTable {
		errOut, resultOut, codeHTTP := checkErr.CheckErrorSignUp(&testCase.errorInput)

		if !((errOut.Error() == testCase.errorExpected) && (string(resultOut) == testCase.resultExpected) &&
			(codeHTTP == testCase.codeHTTPExpected)) {
			t.Errorf("Expected %s, %s, %d", testCase.errorExpected, testCase.resultExpected, testCase.codeHTTPExpected)
		}
	}
}