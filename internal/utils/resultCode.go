package utils

/**
错误码规则:
(1)错误码需>0
(2)错误码为5位数:
---------------------------------------------------------------
		第1位			第2,3位				第4,5位
---------------------------------------------------------------
	服务级错误码     	模块级错误码	 	具体错误码
---------------------------------------------------------------

**/

var (
	// OK
	OK  = response(0, "success") // 通用成功
	Err = response(1, "failed")  // 通用错误

	//系统服务错误码
	ErrParam     = response(10001, "参数有误")
	ErrSignParam = response(10002, "签名参数有误")

	// 模块级错误码 - 用户模块
	ErrUserService         = response(20100, "用户服务异常")
	ErrUserPhone           = response(20101, "用户手机号不合法")
	ErrUserCaptcha         = response(20102, "用户验证码有误")
	ErrUserAccountOrPasswd = response(20101, "账号或密码有误")
)
