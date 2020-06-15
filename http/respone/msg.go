package respone

var ResponeMsg = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "failed",
	INVALID_PARAMS: "请求参数出错",
	USER_PASSWORD_INCORRECT: "用户帐号或登录密码不正确",
	USER_AUTHORIZATION: "登录认证失败",
	USER_CREATED_FAILED: "创建用户失败",
	USER_OLD_PASSWORD_INCORRENT: "旧密码不正确",
	USER_NOT_EXISTS: "用户帐号或登录密码不正确",
	NOT_IN_TASK:  "不属于任务执行者或任务失效",
	TASK_RECEIVE_FAILED: "任务已领取过",
	TASK_DUTY_FAILED: "执行失败，任务未领取或已执行",
	TASK_DUTY_NOT_IN_TIME: "执行失败，当前不是任务的有效执行时间",
}

const (
	SUCCESS = 200
	ERROR = 400
	INVALID_PARAMS = 4001

	USER_AUTHORIZATION = 4002
	USER_NOT_EXISTS = 4003
	USER_PASSWORD_INCORRECT = 4004

	USER_CREATED_FAILED = 4005
	USER_OLD_PASSWORD_INCORRENT = 4006

	NOT_IN_TASK = 4010
	TASK_RECEIVE_FAILED = 4011
	TASK_DUTY_FAILED = 4012
	TASK_DUTY_NOT_IN_TIME = 4013
)
