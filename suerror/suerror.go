package suerror

/*
错误码设计 五位 00000
第一位 表示错误级别, 1 为系统错误, 2 为普通错误
第二三位 表示服务模块代码
第四五位 表示具体错误代码
*/

//
//可读性一定要高
//兼容性要好，要能支持常见的两种错误码类型，字符串和数值型都需要提供
//易于维护，错误码要便于维护
//易用性，错误码的定义要对开发人员友好，最好开发人员不需要关心错误码的值
//灵活可控，错误码可以手动指定，也可以自动进行维护，而且支持错误码自动分段和手动分段

type GlobeError struct {
	Code int
	Msg  string
}

var (
	SUCCESS     = GlobeError{Code: 200, Msg: "Success"}
	PARAMNOTFUND     = GlobeError{Code: 400, Msg: "PARAM_NOT_FUND"}
	NOTFUND     = GlobeError{Code: 404, Msg: "404"}
	SYSTEMERROR = GlobeError{Code: -1, Msg: "SYSTEM_ERROR"}
	NOPARAMETER = GlobeError{Code: -1, Msg: "NO_PARAMETER"}
	TOKENERROR  = GlobeError{Code: 40002, Msg: "TOKEN_ERROR"}
	SIGNERROR   = GlobeError{Code: 40002, Msg: "SIGN_ERROR"}
)
