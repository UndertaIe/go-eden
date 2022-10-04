package errcode

var (
	ErrorService = NewError(20000, "服务异常，请稍后重试")

	ErrorUserRecordNotFound       = NewError(20001, "查询用户不存在")
	ErrorPlatformRecordNotFound   = NewError(20002, "查询平台不存在")
	UserPhoneExists               = NewError(20003, "当前手机号已绑定账户")
	UserEmailExists               = NewError(20004, "当前邮箱已绑定账户")
	UserNameExists                = NewError(20005, "当前用户名已绑定账户")
	ErrorUserAuthFailed           = NewError(20006, "用户验证失败, 请输入正确的账号和密码")
	ErrorVerifyCodeNoPhoneNumbers = NewError(20007, "输入正确的手机号")
	ErrorGenerateVerifyCode       = NewError(20008, "生成验证码失败")
	ErrorSendVerifyCode           = NewError(20009, "发送验证码失败")
	ErrorCheckCode                = NewError(20010, "验证异常")
	ErrorVerifyCodeFailed         = NewError(20011, "验证失败")
	ErrorAuthLinkExpired          = NewError(20012, "认证链接失效")
	ErrorAuthLinkExists           = NewError(20013, "认证链接已存在")
	ErrorUserPhoneNotExists       = NewError(20014, "当前手机号未绑定账户")
	ErrorUserEmailNotExists       = NewError(20015, "当前邮箱未绑定账户")
	ErrorPlatformNameExists       = NewError(20016, "平台名已存在")
	ErrorPlatformAbbrExists       = NewError(20017, "平台简称已存在")
	ErrorUserNoPlatformPassword   = NewError(20018, "当前用户不存在该平台信息")
	ErrorRecordNotFound           = NewError(20019, "查询记录不存在")
	ErrorUpdateRecordNotFound     = NewError(20020, "更新记录不存在")
	ErrorDeleteRecordNotFound     = NewError(20021, "删除记录不存在")
	ErrorDuplicateEntry           = NewError(20022, "记录已存在,创建/更新失败")

	ErrorUnknownService = NewError(29999, "未知的服务异常，请联系负责人排查异常")
)