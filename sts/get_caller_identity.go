package sts

import "github.com/reedchan7/aliyungo/common"

type GetCallerIdentityRequest struct {
}

type GetCallerIdentityResponse struct {
	common.Response
	AccountId string
	UserId    string
	Arn       string
}

func (c *STSClient) GetCallerIdentity() (*GetCallerIdentityResponse, error) {
	resp := &GetCallerIdentityResponse{}
	err := c.Invoke("GetCallerIdentity", &GetCallerIdentityRequest{}, resp)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
