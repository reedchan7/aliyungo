package cs

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/reedchan7/aliyungo/common"
)

type ServerlessCreationArgs struct {
	ClusterType           KubernetesClusterType    `json:"cluster_type"`
	Profile               KubernetesClusterProfile `json:"profile"`
	Name                  string                   `json:"name"`
	RegionId              string                   `json:"region_id"`
	VpcId                 string                   `json:"vpcid"`
	VSwitchId             string                   `json:"vswitch_id"`
	VswitchIds            []string                 `json:"vswitch_ids"`
	EndpointPublicAccess  bool                     `json:"public_slb"`
	PrivateZone           bool                     `json:"private_zone"`
	NatGateway            bool                     `json:"nat_gateway"`
	KubernetesVersion     string                   `json:"kubernetes_version"`
	DeletionProtection    bool                     `json:"deletion_protection"`
	EnableRRSA            bool                     `json:"enable_rrsa"`
	SecurityGroupId       string                   `json:"security_group_id"`
	Tags                  []Tag                    `json:"tags"`
	Addons                []Addon                  `json:"addons"`
	ResourceGroupId       string                   `json:"resource_group_id"`
	ClusterSpec           string                   `json:"cluster_spec"`
	LoadBalancerSpec      string                   `json:"load_balancer_spec"` // api server slb实例规格
	ServiceCIDR           string                   `json:"service_cidr"`
	TimeZone              string                   `json:"timezone"`
	ServiceDiscoveryTypes []string                 `json:"service_discovery_types"`
	ZoneID                string                   `json:"zone_id"`
	LoggingType           string                   `json:"logging_type"`
	SLSProjectName        string                   `json:"sls_project_name"`
	SnatEntry             bool                     `json:"snat_entry"`
}

type ServerlessClusterResponse struct {
	ClusterId          string                `json:"cluster_id"`
	Name               string                `json:"name"`
	ClusterType        KubernetesClusterType `json:"cluster_type"`
	RegionId           string                `json:"region_id"`
	State              ClusterState          `json:"state"`
	VpcId              string                `json:"vpc_id"`
	VSwitchId          string                `json:"vswitch_id"`
	SecurityGroupId    string                `json:"security_group_id"`
	Tags               []Tag                 `json:"tags"`
	Created            time.Time             `json:"created"`
	Updated            time.Time             `json:"updated"`
	InitVersion        string                `json:"init_version"`
	CurrentVersion     string                `json:"current_version"`
	PrivateZone        bool                  `json:"private_zone"`
	DeletionProtection bool                  `json:"deletion_protection"`
	EnableRRSA         bool                  `json:"enable_rrsa"`
	ResourceGroupId    string                `json:"resource_group_id"`
	ClusterSpec        string                `json:"cluster_spec"`
	Profile            string                `json:"profile"`
	MetaData           string                `json:"meta_data"`
}

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (this *ServerlessCreationArgs) Validate() error {
	if this.Name == "" || this.RegionId == "" || this.VpcId == "" {
		return common.GetCustomError("InvalidParameters", "The name,region_id,vpc_id not allowed empty")
	}
	return nil
}

// create Serverless cluster
func (client *Client) CreateServerlessKubernetesCluster(args *ServerlessCreationArgs) (*ClusterCommonResponse, error) {
	if args == nil {
		return nil, common.GetCustomError("InvalidArgs", "The args is nil ")
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	cluster := &ClusterCommonResponse{}
	path := "/clusters"
	if args.ResourceGroupId != "" {
		// 创建集群到指定资源组
		path = fmt.Sprintf("/resource_groups/%s/clusters", args.ResourceGroupId)
	}
	err := client.Invoke(common.Region(args.RegionId), http.MethodPost, path, nil, args, &cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// describe Serverless cluster
func (client *Client) DescribeServerlessKubernetesCluster(clusterId string) (*ServerlessClusterResponse, error) {
	cluster := &ServerlessClusterResponse{}
	err := client.Invoke("", http.MethodGet, "/clusters/"+clusterId, nil, nil, cluster)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

// describe Serverless clsuters
func (client *Client) DescribeServerlessKubernetesClusters() ([]*ServerlessClusterResponse, error) {
	allClusters := make([]*ServerlessClusterResponse, 0)
	askClusters := make([]*ServerlessClusterResponse, 0)

	err := client.Invoke("", http.MethodGet, "/clusters", nil, nil, &allClusters)
	if err != nil {
		return askClusters, err
	}

	for _, cluster := range allClusters {
		// Ask 1.0/2.0
		if cluster.ClusterType == ClusterTypeServerlessKubernetes {
			askClusters = append(askClusters, cluster)
		}
		// Ask 3.0
		if cluster.ClusterType == ClusterTypeManagedKubernetes && cluster.Profile == ProfileServerlessKubernetes {
			askClusters = append(askClusters, cluster)
		}
	}

	return askClusters, nil
}

// new api for get cluster kube user config
func (client *Client) DescribeClusterUserConfig(clusterId string, privateIpAddress bool) (*ClusterConfig, error) {
	config := &ClusterConfig{}
	query := url.Values{}
	query.Add("PrivateIpAddress", strconv.FormatBool(privateIpAddress))

	err := client.Invoke("", http.MethodGet, "/k8s/"+clusterId+"/user_config", query, nil, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
