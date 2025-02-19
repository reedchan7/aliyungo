package sms

import (
	"net/http"

	"github.com/reedchan7/aliyungo/common"
)

// 邮件推送产品短信功能
type SingleSendSmsArgs struct {
	SignName     string
	TemplateCode string
	RecNum       string
	ParamString  string
}

func (this *Client) SingleSendSms(args *SingleSendSmsArgs) error {
	return this.InvokeByAnyMethod(http.MethodPost, SingleSendSms, "", args, &common.Response{})
}
