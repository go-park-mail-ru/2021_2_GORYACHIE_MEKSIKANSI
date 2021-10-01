package Errors

import "time"

type ResultError struct {
	Status	int         `json:"status"`
	Explain	string		`json:"parsedJSON,omitempty"`
}

type Errors struct {
	Text string
	Time time.Time
}

func (e *Errors) Error() string {
	return e.Text
}

const(
	ERRDB  = "ERROR: database is not responding"
	ERRENCODE  = "ERROR: Encode"
	ERRUNMARSHAL  = "ERROR: unmarshal json"
	ERRAUTH  = "ERROR: authorization failed"
	ERRPHONESCAN  = "ERROR: phone not scan"
	ERRPHONEPASSQUERY  = "ERROR: password and phone error query"
	ERRPHONEPASSSCAN  = "ERROR: phone and password not scan"
	ERRMAILSCAN  = "ERROR: email not scan"
	ERRMAILPASSQUERY  = "ERROR: password or email error query"
	ERRMAILPASSSCAN   = "ERROR: email and pass not scan"
	ERRNOTLOGINORPASSWORD  = "Неправильный логин или пароль"
	ERRINFOQUERY = "ERROR: not insert info query"
	ERRINFOSCAN  = "ERROR: info not scan"
	ERRINSERTHOSTQUERY  = "ERROR: not insert host query"
	ERRINSERTCOURIERQUERY  = "ERROR: not insert courier query"
	ERRINSERTCLIENTQUERY  = "ERROR: not insert client query"
	ERRINSERTCOOKIEQUERY  = "ERROR: not insert cookie transact query"
	ERRDELETECOOKIEQUERY  = "ERROR: cookie not delete query"
	ERRINSERTLOGINCOOKIEQUERY  = "ERROR: not insert cookie query"
	ERRUNIQUE  = "Телефон или email уже зарегистрирован"
	ERRCREATEQUERY  = "ERROR: db not created"
	ERRINSERTQUERY  = "ERROR: restaurant not insert"
	ERRCOOKIEANDCSRFSCAN  = "ERROR: cookie and csrf scan"
	ERRCOOKIEQUERY  = "ERROR: cookie query"
	ERRCOOKIESCAN  = "ERROR: cookie scan"
	ERRSIDNOTFOUND  = "ERROR: id not found"
	ERRNOTCONNECT  = "ERROR: not connect"
	ERRUPDATECSRFQUERY  = "ERROR: csrf not updated"
	ERRCOOKIEEXPIRED  = "ERROR: cookie expired"
	ERRDELETEQUERY  = "ERROR: not delete query"
	ERRINSERTROOTQUERY  = "ERROR: not create root"
	ERRCLIENTSCAN  = "ERROR: check user on client not scan"
	ERRHOSTSCAN  = "ERROR: check user on host not scan"
	ERRCORIERSCAN  = "ERROR: check user on courier not scan"
	ERRGETPROFILEHOSTSCAN  = "ERROR: profile host not scan"
	ERRGETPROFILECLIENTSCAN  = "ERROR: profile client not scan"
	ERRGETBIRTHDAYSCAN  = "ERROR: birthday not scan"
	ERRGETPROFILECOURIERSCAN  = "ERROR: profile courier not scan"
	ERRQUERY  = "ERROR: restaurant not get"
	ERRSCAN  = "ERROR: restaurant scan error"
	ERRRESTNULL  = "ERROR: restaurants not found"
	ERRCOOKIEIDNOTFOUND  = "ERROR: id by cookie not found"
	ERRINSERTROOTCLIENTQUERY  = "ERROR: root not insert in client"
)
