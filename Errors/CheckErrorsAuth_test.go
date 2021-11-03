package Errors

import (
	test "2021_2_GORYACHIE_MEKSIKANSI/Test"
	utils "2021_2_GORYACHIE_MEKSIKANSI/Utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"math"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestCheckErrorSignUp(t *testing.T) {
	reqId := utils.RandomInteger(0, math.MaxInt64)
	loggerErrWarn := utils.NewLogger("./loggErrWarn.txt")
	loggerInfo := utils.NewLogger("./loggInfo.txt")
	loggerTest := utils.NewLogger("./loggTest.txt")

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
		errorInput       Errors
		errorExpected    string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Text: AGeneralSignUpLoginNotUnique,
				Time: time.Now(),
			},
			errorExpected:    ErrCheck,
			resultExpected:   "{\"status\":" + strconv.Itoa(http.StatusConflict) + ",\"explain\":\"" + AGeneralSignUpLoginNotUnique + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Text: AGeneralSignUpIncorrectPhoneFormat,
				Time: time.Now(),
			},
			errorExpected:    ErrCheck,
			resultExpected:   "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) + ",\"explain\":\"" + AGeneralSignUpIncorrectPhoneFormat + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Text: AGeneralSignUpNotInsert,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Text: ASignUpHostHostNotInsert,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Text: AAddTransactionCookieNotInsert,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Text: ASignUpCourierCourierNotInsert,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Text: ASignUpClientClientNotInsert,
				Time: time.Now(),
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
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, []byte(nil), resultOut,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, HttpNil, codeHTTP,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
}

func TestCheckErrorLogin(t *testing.T) {
	testTable := []struct {
		errorInput       Errors
		errorExpected    string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Text: ALoginOrPasswordIncorrect,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) +
				",\"explain\":\"" + ALoginOrPasswordIncorrect + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Text: ALoginNotFound,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) +
				",\"explain\":\"" + ALoginOrPasswordIncorrect + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Text: AAddCookieCookieNotInsert,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusInternalServerError) +
				",\"explain\":\"" + ErrDB + "\"}",
			codeHTTPExpected: http.StatusInternalServerError,
		},
		{
			errorInput: Errors{
				Text: ASaltNotSelect,
				Time: time.Now(),
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
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, []byte(nil), resultOut,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, HttpNil, codeHTTP,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
}

func TestCheckErrorLogout(t *testing.T) {
	testTable := []struct {
		errorInput       Errors
		errorExpected    string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Text: ADeleteCookieCookieNotDelete,
				Time: time.Now(),
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
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, []byte(nil), resultOut,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, HttpNil, codeHTTP,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
}

func TestCheckErrorLogoutAccess(t *testing.T) {
	testTable := []struct {
		errorInput       Errors
		errorExpected    string
		resultExpected   string
		codeHTTPExpected int
	}{
		{
			errorInput: Errors{
				Text: MCheckAccessCookieNotFound,
				Time: time.Now(),
			},
			errorExpected: ErrCheck,
			resultExpected: "{\"status\":" + strconv.Itoa(http.StatusUnauthorized) +
				",\"explain\":\"" + MCheckAccessCookieNotFound + "\"}",
			codeHTTPExpected: http.StatusOK,
		},
		{
			errorInput: Errors{
				Text: MCheckAccessCookieNotScan,
				Time: time.Now(),
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
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, []byte(nil), resultOut,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
	assert.Equal(t, HttpNil, codeHTTP,
		fmt.Sprintf("Expected %s, %s, %d", test.NilStr, test.NilStr, HttpNil),
	)
}
