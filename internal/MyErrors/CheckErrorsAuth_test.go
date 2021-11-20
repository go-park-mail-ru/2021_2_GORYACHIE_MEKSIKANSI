package MyErrors

import (
	test "2021_2_GORYACHIE_MEKSIKANSI/test"
	"2021_2_GORYACHIE_MEKSIKANSI/internal/Utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"math"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func TestCheckErrorSignUp(t *testing.T) {
	reqId := Utils.RandomInteger(0, math.MaxInt64)
	loggerTest := Utils.NewLogger("./loggTest.txt")

	defer func(loggerErrWarn *zap.SugaredLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(loggerErrWarn)

	defer func(loggerInfo *zap.SugaredLogger) {
		errLogger := loggerInfo.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerInfo the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(loggerInfo)

	defer func(loggerTest *zap.SugaredLogger) {
		errLogger := loggerTest.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerTest the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(loggerTest)

	testTable := []struct {
		errorInput    Errors
		errorExpected string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Alias: AGeneralSignUpLoginNotUnique,
			},
			errorExpected:    ErrCheck,
			resultExpected:   "{\"status\":" + strconv.Itoa(http.StatusConflict) + ",\"explain\":\"" + AGeneralSignUpLoginNotUnique + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Alias: AGeneralSignUpIncorrectPhoneFormat,
			},
			errorExpected:    ErrCheck,
			resultExpected:   "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) + ",\"explain\":\"" + AGeneralSignUpIncorrectPhoneFormat + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Alias: AGeneralSignUpNotInsert,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Alias: ASignUpHostHostNotInsert,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Alias: AAddTransactionCookieNotInsert,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Alias: ASignUpCourierCourierNotInsert,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Alias: ASignUpClientClientNotInsert,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
	}

	checkError := &CheckError{
		LoggerErrWarn: loggerErrWarn,
		LoggerInfo:    loggerInfo,
		LoggerTest:    loggerTest,
		RequestId:     &reqId,
	}

	for _, testCase := range testTable {
		errOut, resultOut, codeHTTP := checkError.CheckErrorSignUp(&testCase.errorInput)
		assert.Equal(t, testCase.errorExpected, errOut.Error(),
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
	errOut, resultOut, codeHTTP := checkError.CheckErrorSignUp(err)
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

func TestCheckErrorLogin(t *testing.T) {
	testTable := []struct {
		errorInput    Errors
		errorExpected string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Alias: ALoginOrPasswordIncorrect,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) +
				",\"explain\":\"" + ALoginOrPasswordIncorrect + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Alias: ALoginNotFound,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) +
				",\"explain\":\"" + ALoginOrPasswordIncorrect + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Alias: AAddCookieCookieNotInsert,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Alias: ASaltNotSelect,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
	}

	for _, testCase := range testTable {
		errOut, resultOut, codeHTTP := CheckErrorLogin(&testCase.errorInput)
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
	errOut, resultOut, codeHTTP := CheckErrorLogin(err)
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

func TestCheckErrorLogout(t *testing.T) {
	testTable := []struct {
		errorInput    Errors
		errorExpected string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Alias: ADeleteCookieCookieNotDelete,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
	}

	for _, testCase := range testTable {
		errOut, resultOut, codeHTTP := CheckErrorLogout(&testCase.errorInput)
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
	errOut, resultOut, codeHTTP := CheckErrorLogout(err)
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

func TestCheckErrorLogoutAccess(t *testing.T) {
	testTable := []struct {
		errorInput    Errors
		errorExpected string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Alias: MCheckAccessCookieNotFound,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) +
				",\"explain\":\"" + MCheckAccessCookieNotFound + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Alias: MCheckAccessCookieNotScan,
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
	}

	for _, testCase := range testTable {
		errOut, resultOut, codeHTTP := CheckErrorAccess(&testCase.errorInput)
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
	errOut, resultOut, codeHTTP := CheckErrorAccess(err)
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
