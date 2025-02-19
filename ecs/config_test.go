package ecs

import (
	"os"

	"github.com/reedchan7/aliyungo/common"
)

// Modify with your Access Key Id and Access Key Secret

var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestSecurityToken   = os.Getenv("SecurityToken")
	TestRegionID        = common.Region(os.Getenv("RegionId"))
	TestVpcId           = os.Getenv("VpcId")
	TestVswitchID       = os.Getenv("TestVswitchID")

	TestInstanceId      = os.Getenv("InstanceId")
	TestSecurityGroupId = os.Getenv("TestSecurityGroupId")
	TestResourceGroupId = os.Getenv("TestResourceGroupId")
	TestImageId         = os.Getenv("TestImageId")
	TestAccountId       = "MY_TEST_ACCOUNT_ID" // Get from https://account.console.aliyun.com
	TestInstanceType    = os.Getenv("InstanceType")
	TestVSwitchID       = "MY_TEST_VSWITCHID"

	TestIAmRich = false
	TestQuick   = false
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}

var testVpcDebugClient *Client

func NewVpcTestClientForDebug() *Client {
	if testVpcDebugClient == nil {
		testVpcDebugClient = NewVPCClient(TestAccessKeyId, TestAccessKeySecret, TestRegionID)
		testVpcDebugClient.SetDebug(true)
	}
	return testVpcDebugClient
}

var testLocationClient *Client

func NetTestLocationClientForDebug() *Client {
	if testLocationClient == nil {
		testLocationClient = NewECSClientWithSecurityToken4RegionalDomain(TestAccessKeyId, TestAccessKeySecret, TestSecurityToken, TestRegionID)
		testLocationClient.SetDebug(true)
	}

	return testLocationClient
}
