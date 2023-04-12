package serializer

const (
	OK                 = 0
	WRONG_PASSWORD     = iota + 4000 //密码错误
	DIFF_REPASSWORD                  //两次密码不一致
	INVALID_TOKEN                    //无效的token
	REGISTED_EMAIL                   //邮箱已被注册
	FAIL_REGISTER                    //注册失败
	PARAM_VALIDATE_ERR               //参数验证失败
	USER_NOT_EXSIST                  //用户不存在
	INCORRECT_PASSWD                 //密码错误
)
