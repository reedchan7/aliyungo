package ecs

import (
	"os"
	"strings"
	"testing"

	"github.com/reedchan7/aliyungo/common"
)

func TestSecurityGroups(t *testing.T) {

	client := NewTestClient()
	regions, err := client.DescribeRegions()

	t.Log("regions: ", regions, err)

	for _, region := range regions {

		arg := DescribeSecurityGroupsArgs{
			RegionId: region.RegionId,
		}

		sgs, _, err := client.DescribeSecurityGroups(&arg)
		if err != nil {
			t.Errorf("Failed to DescribeSecurityGroups for region %s: %v", region.RegionId, err)
			continue
		}
		for _, sg := range sgs {
			t.Logf("SecurityGroup: %++v", sg)

			args := DescribeSecurityGroupAttributeArgs{
				SecurityGroupId: sg.SecurityGroupId,
				RegionId:        region.RegionId,
			}
			sga, err := client.DescribeSecurityGroupAttribute(&args)
			if err != nil {
				t.Errorf("Failed to DescribeSecurityGroupAttribute %s: %v", sg.SecurityGroupId, err)
				continue
			}
			t.Logf("SecurityGroup Attribute: %++v", sga)

		}
	}
}

func TestSecurityGroups_BatchQuery(t *testing.T) {

	client := NewTestClient()
	arg := DescribeSecurityGroupsArgs{
		RegionId:         common.Region(os.Getenv("RegionId")),
		SecurityGroupIds: strings.Split(os.Getenv("SecurityGroupIDs"), ","),
	}

	sgs, _, err := client.DescribeSecurityGroups(&arg)
	if err != nil {
		t.Errorf("Failed to DescribeSecurityGroups, error:%++v", err)
	}
	for _, sg := range sgs {
		t.Logf("SecurityGroup: %++v", sg)
	}
}

func TestECSSecurityGroupCreationAndDeletion(t *testing.T) {
	client := NewTestClient()
	instance, err := client.DescribeInstanceAttribute(TestInstanceId)
	if err != nil {
		t.Fatalf("Failed to describe instance attribute %s: %v", TestInstanceId, err)
	}
	regionId := instance.RegionId

	_testECSSecurityGroupCreationAndDeletion(t, client, regionId, "")

}

func _testECSSecurityGroupCreationAndDeletion(t *testing.T, client *Client, regionId common.Region, vpcId string) {

	sgName := "test-security-group"
	args := CreateSecurityGroupArgs{
		RegionId:          regionId,
		VpcId:             vpcId,
		SecurityGroupName: sgName,
	}

	sgId, err := client.CreateSecurityGroup(&args)
	if err != nil {
		t.Fatalf("Failed to create security group %s: %v", sgName, err)
	}
	t.Logf("Security group %s is created successfully.", sgId)

	describeArgs := DescribeSecurityGroupAttributeArgs{
		SecurityGroupId: sgId,
		RegionId:        regionId,
	}
	sg, err := client.DescribeSecurityGroupAttribute(&describeArgs)
	if err != nil {
		t.Errorf("Failed to describe security group %s: %v", sgId, err)
	}
	t.Logf("Security group %s: %++v", sgId, sg)

	newName := "test-security-group-new"
	modifyArgs := ModifySecurityGroupAttributeArgs{
		SecurityGroupId:   sgId,
		RegionId:          regionId,
		SecurityGroupName: newName,
	}
	err = client.ModifySecurityGroupAttribute(&modifyArgs)
	if err != nil {
		t.Errorf("Failed to modify security group %s: %v", sgId, err)
	} else {
		sg, err := client.DescribeSecurityGroupAttribute(&describeArgs)
		if err != nil {
			t.Errorf("Failed to describe security group %s: %v", sgId, err)
		}
		t.Logf("Security group %s: %++v", sgId, sg)
		if sg.SecurityGroupName != newName {
			t.Errorf("Failed to modify security group %s with new name %s", sgId, newName)
		}
	}

	err = client.DeleteSecurityGroup(regionId, sgId)

	if err != nil {
		t.Fatalf("Failed to delete security group %s: %v", sgId, err)
	}
	t.Logf("Security group %s is deleted successfully.", sgId)
}

func TestDescribeSecurityGroupAttribute(t *testing.T) {
	client := NewTestClientForDebug()
	args := DescribeSecurityGroupAttributeArgs{
		RegionId:        common.Beijing,
		SecurityGroupId: TestSecurityGroupId,
	}
	attr, err := client.DescribeSecurityGroupAttribute(&args)
	if err != nil {
		t.Fatalf("Failed to describe securitygroup attribute %++v", err)
	}
	t.Logf("Security group attribute is %++v", attr)
}
