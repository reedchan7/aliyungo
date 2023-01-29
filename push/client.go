package push

import "github.com/reedchan7/aliyungo/common"

const (
	PushEndPoint   = "https://cloudpush.aliyuncs.com/"
	Push           = "Push"
	PushAPIVersion = "2016-08-01"
)

type Client struct {
	common.Client
}

func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := new(Client)
	client.Init(PushEndPoint, PushAPIVersion, accessKeyId, accessKeySecret)
	return client
}
