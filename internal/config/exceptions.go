package config

type CodeType struct {
	status int16
	msg    string
}

// 五位，第一位：业务标识

var SuccessCode CodeType = CodeType{10000, "Request Success"}
var ErrorCode CodeType = CodeType{19999, "Request Error"}

var StudentNoRepeat CodeType = CodeType{11000, "该学号已经注册过。"}
var UserRegisterError CodeType = CodeType{11001, "用户注册失败。"}
var LoginError CodeType = CodeType{11002, "用户登陆失败，密码或者用户账号输入错误！"}
var NoneUser CodeType = CodeType{11003, "用户不存在或者用户状态异常！"}
var InvalidFileError CodeType = CodeType{11004, "上传文件错误!"}

func (code CodeType) String() string {
	return code.msg
}

func (code CodeType) Code() int16 {
	return code.status
}
