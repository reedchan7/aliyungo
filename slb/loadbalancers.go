package slb

import (
	"fmt"
	"time"

	"github.com/reedchan7/aliyungo/common"
	"github.com/reedchan7/aliyungo/util"
)

type AddressType string

const (
	InternetAddressType = AddressType("internet")
	IntranetAddressType = AddressType("intranet")
)

type InternetChargeType string

const (
	PayByBandwidth = InternetChargeType("paybybandwidth")
	PayByTraffic   = InternetChargeType("paybytraffic")
)

type AddressIPVersionType string

const (
	IPv4 = AddressIPVersionType("ipv4")
	IPv6 = AddressIPVersionType("ipv6")
)

type LoadBalancerSpecType string

const (
	S1Small  = "slb.s1.small"
	S2Small  = "slb.s2.small"
	S2Medium = "slb.s2.medium"
	S3Small  = "slb.s3.small"
	S3Medium = "slb.s3.medium"
	S3Large  = "slb.s3.large"
)

type ModificationProtectionType string

const (
	NonProtection     = ModificationProtectionType("NonProtection")
	ConsoleProtection = ModificationProtectionType("ConsoleProtection")
)

type CreateLoadBalancerArgs struct {
	RegionId                     common.Region
	LoadBalancerName             string
	AddressType                  AddressType
	VSwitchId                    string
	InternetChargeType           InternetChargeType
	Bandwidth                    int
	ClientToken                  string
	MasterZoneId                 string
	SlaveZoneId                  string
	LoadBalancerSpec             LoadBalancerSpecType
	AddressIPVersion             AddressIPVersionType
	DeleteProtection             FlagType
	ModificationProtectionStatus ModificationProtectionType
	ModificationProtectionReason string
	ResourceGroupId              string
}

type CreateLoadBalancerResponse struct {
	common.Response
	LoadBalancerId   string
	Address          string
	NetworkType      string
	VpcId            string
	VSwitchId        string
	LoadBalancerName string
	MasterZoneId     string
	SlaveZoneId      string
	AddressIPVersion AddressIPVersionType
}

// CreateLoadBalancer create loadbalancer
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&CreateLoadBalancer
func (client *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (response *CreateLoadBalancerResponse, err error) {
	response = &CreateLoadBalancerResponse{}
	err = client.Invoke("CreateLoadBalancer", args, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

type DeleteLoadBalancerArgs struct {
	LoadBalancerId string
}

type DeleteLoadBalancerResponse struct {
	common.Response
}

// DeleteLoadBalancer delete loadbalancer
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&DeleteLoadBalancer
func (client *Client) DeleteLoadBalancer(loadBalancerId string) (err error) {
	args := &DeleteLoadBalancerArgs{
		LoadBalancerId: loadBalancerId,
	}
	response := &DeleteLoadBalancerResponse{}
	err = client.Invoke("DeleteLoadBalancer", args, response)
	return err
}

type ModifyLoadBalancerInternetSpecArgs struct {
	LoadBalancerId     string
	InternetChargeType InternetChargeType
	Bandwidth          int
}

type ModifyLoadBalancerInternetSpecResponse struct {
	common.Response
}

// ModifyLoadBalancerInternetSpec Modify loadbalancer internet spec
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&ModifyLoadBalancerInternetSpec

func (client *Client) ModifyLoadBalancerInternetSpec(args *ModifyLoadBalancerInternetSpecArgs) (err error) {
	response := &ModifyLoadBalancerInternetSpecResponse{}
	err = client.Invoke("ModifyLoadBalancerInternetSpec", args, response)
	return err
}

type ModifyLoadBalancerInstanceSpecArgs struct {
	RegionId         common.Region
	LoadBalancerId   string
	LoadBalancerSpec LoadBalancerSpecType
}

// ModifyLoadBalancerInstanceSpec Modify loadbalancer instance spec
//
// You can read doc at https://help.aliyun.com/document_detail/53360.html

func (client *Client) ModifyLoadBalancerInstanceSpec(args *ModifyLoadBalancerInstanceSpecArgs) (err error) {
	response := &common.Response{}
	err = client.Invoke("ModifyLoadBalancerInstanceSpec", args, response)
	return err
}

type Status string

const InactiveStatus = Status("inactive")
const ActiveStatus = Status("active")
const LockedStatus = Status("locked")

type SetLoadBalancerStatusArgs struct {
	LoadBalancerId     string
	LoadBalancerStatus Status
}

type SetLoadBalancerStatusResponse struct {
	common.Response
}

// SetLoadBalancerStatus Set loadbalancer status
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&SetLoadBalancerStatus

func (client *Client) SetLoadBalancerStatus(loadBalancerId string, status Status) (err error) {
	args := &SetLoadBalancerStatusArgs{
		LoadBalancerId:     loadBalancerId,
		LoadBalancerStatus: status,
	}
	response := &SetLoadBalancerStatusResponse{}
	err = client.Invoke("SetLoadBalancerStatus", args, response)
	return err
}

type SetLoadBalancerNameArgs struct {
	LoadBalancerId   string
	LoadBalancerName string
}

type SetLoadBalancerNameResponse struct {
	common.Response
}

// SetLoadBalancerName Set loadbalancer name
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&SetLoadBalancerName

func (client *Client) SetLoadBalancerName(loadBalancerId string, name string) (err error) {
	args := &SetLoadBalancerNameArgs{
		LoadBalancerId:   loadBalancerId,
		LoadBalancerName: name,
	}
	response := &SetLoadBalancerNameResponse{}
	err = client.Invoke("SetLoadBalancerName", args, response)
	return err
}

type DescribeLoadBalancersArgs struct {
	RegionId           common.Region
	LoadBalancerId     string
	LoadBalancerName   string
	AddressType        AddressType
	NetworkType        string
	VpcId              string
	VSwitchId          string
	Address            string
	InternetChargeType InternetChargeType
	ServerId           string
	Tags               string
}

type ListenerPortAndProtocolType struct {
	ListenerPort     int
	ListenerProtocol string
	Description      string
}

type BackendServerType struct {
	ServerId string
	Weight   int
	Type     string
}

type ServiceManagedModeType string

const (
	Managed           = ServiceManagedModeType("Managed")
	Unmanaged         = ServiceManagedModeType("Unmanaged")
	DependencyManaged = ServiceManagedModeType("DependencyManaged")
)

type LoadBalancerType struct {
	LoadBalancerId               string
	ResourceGroupId              string
	LoadBalancerName             string
	LoadBalancerStatus           string
	Address                      string
	RegionId                     common.Region
	RegionIdAlias                string
	AddressType                  AddressType
	VSwitchId                    string
	VpcId                        string
	NetworkType                  string
	Bandwidth                    int
	InternetChargeType           InternetChargeType
	CreateTime                   string // Why not ISO 6801
	CreateTimeStamp              util.ISO6801Time
	DeleteProtection             FlagType
	ModificationProtectionStatus ModificationProtectionType
	ModificationProtectionReason string
	ListenerPorts                struct {
		ListenerPort []int
	}
	ListenerPortsAndProtocol struct {
		ListenerPortAndProtocol []ListenerPortAndProtocolType
	}
	BackendServers struct {
		BackendServer []BackendServerType
	}
	LoadBalancerSpec   LoadBalancerSpecType
	MasterZoneId       string
	SlaveZoneId        string
	AddressIPVersion   AddressIPVersionType
	ServiceManagedMode ServiceManagedModeType
}

type DescribeLoadBalancersResponse struct {
	common.Response
	LoadBalancers struct {
		LoadBalancer []LoadBalancerType
	}
}

// DescribeLoadBalancers Describe loadbalancers
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&DescribeLoadBalancers

func (client *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (loadBalancers []LoadBalancerType, err error) {
	response := &DescribeLoadBalancersResponse{}
	err = client.Invoke("DescribeLoadBalancers", args, response)
	if err != nil {
		return nil, err
	}
	return response.LoadBalancers.LoadBalancer, err
}

type DescribeLoadBalancerAttributeArgs struct {
	LoadBalancerId string
}

type DescribeLoadBalancerAttributeResponse struct {
	common.Response
	LoadBalancerType
}

// DescribeLoadBalancerAttribute Describe loadbalancer attribute
//
// You can read doc at http://docs.aliyun.com/#/pub/slb/api-reference/api-related-loadbalancer&DescribeLoadBalancerAttribute

func (client *Client) DescribeLoadBalancerAttribute(loadBalancerId string) (loadBalancer *LoadBalancerType, err error) {
	args := &DescribeLoadBalancersArgs{
		LoadBalancerId: loadBalancerId,
	}
	response := &DescribeLoadBalancerAttributeResponse{}
	err = client.Invoke("DescribeLoadBalancerAttribute", args, response)
	if err != nil {
		return nil, err
	}
	return &response.LoadBalancerType, err
}

type NewDescribeLoadBalancerAttributeArgs struct {
	LoadBalancerId string
	RegionId       common.Region
	MasterZoneId   string
	SlaveZoneId    string
}

// New DescribeLoadBalancerAttribute to describe loadbalancer attribute using regionId avoiding to get not found error
//
// You can read doc at https://www.alibabacloud.com/help/doc-detail/27583.htm

func (client *Client) NewDescribeLoadBalancerAttribute(args *NewDescribeLoadBalancerAttributeArgs) (loadBalancer *LoadBalancerType, err error) {

	response := &DescribeLoadBalancerAttributeResponse{}
	err = client.Invoke("DescribeLoadBalancerAttribute", args, response)
	if err != nil {
		return nil, err
	}
	return &response.LoadBalancerType, err
}

// WaitForListener waits for listener to given status
func (client *Client) WaitForLoadBalancerAsyn(loadBalancerId string, status Status, timeout int) error {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	for {
		lb, err := client.DescribeLoadBalancerAttribute(loadBalancerId)

		e, _ := err.(*common.Error)
		if e != nil {
			if e.StatusCode == 404 || e.Code == "InvalidLoadBalancerId.NotFound" {
				continue
			}
			return err
		} else if lb != nil && Status(lb.LoadBalancerStatus) == status {
			// TODO
			break
		}
		timeout = timeout - DefaultWaitForInterval
		if timeout <= 0 {
			return common.GetClientErrorFromString(fmt.Sprintf("Timeout waitting for load balacner %#v", status))
		}
		time.Sleep(DefaultWaitForInterval * time.Second)
	}
	return nil
}

type SetLoadBalancerDeleteProtectionArgs struct {
	LoadBalancerId   string
	DeleteProtection FlagType
	RegionId         common.Region
}

type SetLoadBalancerDeleteProtectionResponse struct {
	common.Response
}

// SetLoadBalancerDeleteProtection loadbalancer delete protection
//
// You can read doc at https://help.aliyun.com/document_detail/122674.html?spm=a2c4g.11186623.6.720.694f4265hwOdXQ
func (client *Client) SetLoadBalancerDeleteProtection(args *SetLoadBalancerDeleteProtectionArgs) (err error) {
	response := &SetLoadBalancerDeleteProtectionResponse{}
	err = client.Invoke("SetLoadBalancerDeleteProtection", args, response)
	return err
}

type SetLoadBalancerModificationProtectionArgs struct {
	RegionId                     common.Region
	LoadBalancerId               string
	ModificationProtectionStatus ModificationProtectionType
	ModificationProtectionReason string
}

type SetLoadBalancerModificationProtectionResponse struct {
	common.Response
}

func (client *Client) SetLoadBalancerModificationProtection(args *SetLoadBalancerModificationProtectionArgs) (err error) {
	response := &SetLoadBalancerModificationProtectionResponse{}
	err = client.Invoke("SetLoadBalancerModificationProtection", args, response)
	return err
}

type ManagedResourceType string

const (
	ManagedLoadBalancer           = ManagedResourceType("LoadBalancer")
	ManagedTls                    = ManagedResourceType("Tls")
	ManagedVServerGroup           = ManagedResourceType("VServerGroup")
	ManagedMasterSlaveServerGroup = ManagedResourceType("MasterSlaveServerGroup")
	ManagedAcl                    = ManagedResourceType("Acl")
	ManagedListener               = ManagedResourceType("Listener")
	ManagedRule                   = ManagedResourceType("Rule")
	ManagedAppRule                = ManagedResourceType("AppRule")
	ManagedDomainExtension        = ManagedResourceType("DomainExtension")
)

type ManagedResourceModel struct {
	ResourceId string
	Port       int
	Protocol   string
}

type ServiceManagedControlArgs struct {
	RegionId           common.Region
	ServiceManagedMode ServiceManagedModeType
	ResourceUid        string
	ResourceBid        string
	ResourceType       ManagedResourceType
	Resources          []ManagedResourceModel
}

type ServiceManagedControlResponse struct {
	common.Response
}

// api: https://yuque.antfin-inc.com/docs/share/63b5a2d3-6fb3-4bd7-a50e-c4b385b866fd?#
func (client *Client) ServiceManagedControl(args *ServiceManagedControlArgs) (err error) {
	response := &ServiceManagedControlResponse{}
	err = client.Invoke("ServiceManagedControl", args, response)
	return err
}
