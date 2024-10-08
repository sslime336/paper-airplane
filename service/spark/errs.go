package spark

import "errors"

var ErrorCodeMap = map[int64]error{
	10000: errors.New("升级为ws出现错误"),
	10001: errors.New("通过ws读取用户的消息 出错"),
	10002: errors.New("通过ws向用户发送消息 出错"),
	10003: errors.New("用户的消息格式有错误"),
	10004: errors.New("用户数据的schema错误"),
	10005: errors.New("用户参数值有错误"),
	10006: errors.New("用户并发错误：当前用户已连接，同一用户不能多处同时连接。"),
	10007: errors.New("用户流量受限：服务正在处理用户当前的问题，需等待处理完成后再发送新的请求。 （必须要等大模型完全回复之后，才能发送下一个问题）"),
	10008: errors.New("服务容量不足，联系服务商"),
	10009: errors.New("和引擎建立连接失败"),
	10010: errors.New("接收引擎数据的错误"),
	10011: errors.New("向引擎发送数据的错误"),
	10012: errors.New("引擎内部错误"),
	10013: errors.New("用户问题涉及敏感信息，审核不通过，拒绝处理此次请求。"),
	10014: errors.New("回复结果涉及到敏感信息，审核不通过，后续结果无法展示给用户。（建议清空当前结果，并给用户提示/警告：该答案涉及到敏感/政治/恐怖/色情/暴力等方面，不予显示/回复）"),
	10015: errors.New("appid在黑名单中"),
	10016: errors.New("appid授权类的错误。比如：未开通此功能，未开通对应版本，token不足，并发超过授权 等等。 （联系我们开通授权或提高限制）"),
	10018: errors.New("用户在5分钟内持续发送ping消息，但并没有实际请求数据，会返回该错误码并断开ws连接。短链接使用无需关注"),
	10019: errors.New("该错误码表示返回结果疑似敏感，建议拒绝用户继续交互"),
	10110: errors.New("服务忙，请稍后再试。"),
	10163: errors.New("请求引擎的参数异常 引擎的schema 检查不通过"),
	10222: errors.New("引擎网络异常"),
	10223: errors.New("LB找不到引擎节点"),
	10907: errors.New("token数量超过上限。对话历史+问题的字数太多，需要精简输入。"),
	11200: errors.New("授权错误：该appId没有相关功能的授权 或者 业务量超过限制（联系我们开通授权或提高限制）"),
	11201: errors.New("授权错误：日流控超限。超过当日最大访问量的限制。（联系我们提高限制）"),
	11202: errors.New("授权错误：秒级流控超限。秒级并发超过授权路数限制。（联系我们提高限制）"),
	11203: errors.New("授权错误：并发流控超限。并发路数超过授权路数限制。（联系我们提高限制）"),
}
