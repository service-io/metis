// Package reply
// @author tabuyos
// @since 2023/7/19
// @description model
package reply

type Reply struct {
	Code    *uint   `json:"code"`
	State   *bool   `json:"state"`
	Message *string `json:"message"`
	Data    *any    `json:"data"`
}

func New() *Reply {
	return &Reply{}
}

func of(code uint, state bool, message string, data any) *Reply {
	return &Reply{
		Code:    &code,
		State:   &state,
		Message: &message,
		Data:    &data,
	}
}

func Ok() *Reply {
	return of(200, true, "成功", nil)
}

func OkData(data any) *Reply {
	return of(200, true, "成功", data)
}

func OkMsg(message string) *Reply {
	return of(200, true, message, nil)
}

func Failed() *Reply {
	return of(300, false, "失败", nil)
}

func FailedData(data any) *Reply {
	return of(300, false, "失败", data)
}

func FailedMsg(message string) *Reply {
	return of(300, false, message, nil)
}

func (receiver *Reply) WithCode(code uint) *Reply {
	receiver.Code = &code
	return receiver
}

func (receiver *Reply) WithState(state bool) *Reply {
	receiver.State = &state
	return receiver
}

func (receiver *Reply) WithMessage(message string) *Reply {
	receiver.Message = &message
	return receiver
}

func (receiver *Reply) WithData(data any) *Reply {
	receiver.Data = &data
	return receiver
}
