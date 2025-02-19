package nas

import (
	"github.com/reedchan7/aliyungo/common"
)

const (
	VERSION            = "2016-02-29"
	END_POINT          = "https://nasservice-inner.aliyuncs.com"
	DEFAULT_POLICY     = "readwrite"
	DEFAULT_SQUASHTYPE = "no_squash"
	DEFAULT_PRIORITY   = "1"
)

type Client struct {
	common.Client
}

// NewClient creates a new instance of NAS client
func NewClient(accessKeyId, accessKeySecret string) *Client {
	client := &Client{}
	client.Init(END_POINT, VERSION, accessKeyId, accessKeySecret)
	return client
}
