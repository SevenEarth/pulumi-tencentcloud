// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides an elasticsearch instance resource.
//
// ## Example Usage
//
// ### Create a basic version of elasticsearch instance paid by the hour
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Elasticsearch"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			availabilityZone, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "es",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				VpcId:            vpc.ID(),
//				AvailabilityZone: pulumi.String(availabilityZone.Zones[0].Name),
//				CidrBlock:        pulumi.String("10.0.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Elasticsearch.NewInstance(ctx, "example", &Elasticsearch.InstanceArgs{
//				InstanceName:      pulumi.String("tf_example_es"),
//				AvailabilityZone:  pulumi.String(availabilityZone.Zones[0].Name),
//				Version:           pulumi.String("7.10.1"),
//				VpcId:             vpc.ID(),
//				SubnetId:          subnet.ID(),
//				Password:          pulumi.String("Test12345"),
//				LicenseType:       pulumi.String("basic"),
//				BasicSecurityType: pulumi.Int(2),
//				WebNodeTypeInfos: elasticsearch.InstanceWebNodeTypeInfoArray{
//					&elasticsearch.InstanceWebNodeTypeInfoArgs{
//						NodeNum:  pulumi.Int(1),
//						NodeType: pulumi.String("ES.S1.MEDIUM4"),
//					},
//				},
//				NodeInfoLists: elasticsearch.InstanceNodeInfoListArray{
//					&elasticsearch.InstanceNodeInfoListArgs{
//						NodeNum:  pulumi.Int(2),
//						NodeType: pulumi.String("ES.S1.MEDIUM8"),
//						Encrypt:  pulumi.Bool(false),
//					},
//				},
//				EsAcl: &elasticsearch.InstanceEsAclArgs{
//					WhiteLists: pulumi.StringArray{
//						pulumi.String("127.0.0.1"),
//					},
//				},
//				Tags: pulumi.Map{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Create a basic version of elasticsearch instance for multi-availability zone deployment
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Elasticsearch"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			availabilityZone, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "es",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				VpcId:            vpc.ID(),
//				AvailabilityZone: pulumi.String(availabilityZone.Zones[0].Name),
//				CidrBlock:        pulumi.String("10.0.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			subnetMultiZone, err := Subnet.NewInstance(ctx, "subnetMultiZone", &Subnet.InstanceArgs{
//				VpcId:            vpc.ID(),
//				AvailabilityZone: pulumi.String(availabilityZone.Zones[1].Name),
//				CidrBlock:        pulumi.String("10.0.2.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Elasticsearch.NewInstance(ctx, "exampleMultiZone", &Elasticsearch.InstanceArgs{
//				InstanceName:      pulumi.String("tf_example_es"),
//				AvailabilityZone:  pulumi.String("-"),
//				Version:           pulumi.String("7.10.1"),
//				VpcId:             vpc.ID(),
//				SubnetId:          pulumi.String("-"),
//				Password:          pulumi.String("Test12345"),
//				LicenseType:       pulumi.String("basic"),
//				BasicSecurityType: pulumi.Int(2),
//				DeployMode:        pulumi.Int(1),
//				MultiZoneInfos: elasticsearch.InstanceMultiZoneInfoArray{
//					&elasticsearch.InstanceMultiZoneInfoArgs{
//						AvailabilityZone: pulumi.String(availabilityZone.Zones[0].Name),
//						SubnetId:         subnet.ID(),
//					},
//					&elasticsearch.InstanceMultiZoneInfoArgs{
//						AvailabilityZone: pulumi.String(availabilityZone.Zones[1].Name),
//						SubnetId:         subnetMultiZone.ID(),
//					},
//				},
//				WebNodeTypeInfos: elasticsearch.InstanceWebNodeTypeInfoArray{
//					&elasticsearch.InstanceWebNodeTypeInfoArgs{
//						NodeNum:  pulumi.Int(1),
//						NodeType: pulumi.String("ES.S1.MEDIUM4"),
//					},
//				},
//				NodeInfoLists: elasticsearch.InstanceNodeInfoListArray{
//					&elasticsearch.InstanceNodeInfoListArgs{
//						Type:     pulumi.String("dedicatedMaster"),
//						NodeNum:  pulumi.Int(3),
//						NodeType: pulumi.String("ES.S1.MEDIUM8"),
//						Encrypt:  pulumi.Bool(false),
//					},
//					&elasticsearch.InstanceNodeInfoListArgs{
//						Type:     pulumi.String("hotData"),
//						NodeNum:  pulumi.Int(2),
//						NodeType: pulumi.String("ES.S1.MEDIUM8"),
//						Encrypt:  pulumi.Bool(false),
//					},
//				},
//				EsAcl: &elasticsearch.InstanceEsAclArgs{
//					WhiteLists: pulumi.StringArray{
//						pulumi.String("127.0.0.1"),
//					},
//				},
//				Tags: pulumi.Map{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// Elasticsearch instance can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Elasticsearch/instance:Instance foo es-17634f05
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	AvailabilityZone pulumi.StringPtrOutput `pulumi:"availabilityZone"`
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	BasicSecurityType pulumi.IntPtrOutput `pulumi:"basicSecurityType"`
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
	ChargePeriod pulumi.IntPtrOutput `pulumi:"chargePeriod"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// Instance creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	DeployMode pulumi.IntPtrOutput `pulumi:"deployMode"`
	// Elasticsearch domain name.
	ElasticsearchDomain pulumi.StringOutput `pulumi:"elasticsearchDomain"`
	// Elasticsearch port.
	ElasticsearchPort pulumi.IntOutput `pulumi:"elasticsearchPort"`
	// Elasticsearch VIP.
	ElasticsearchVip pulumi.StringOutput `pulumi:"elasticsearchVip"`
	// Kibana Access Control Configuration.
	EsAcl InstanceEsAclOutput `pulumi:"esAcl"`
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// Kibana public network access status. Valid values are `OPEN` and `CLOSE`.
	KibanaPublicAccess pulumi.StringOutput `pulumi:"kibanaPublicAccess"`
	// Kibana access URL.
	KibanaUrl pulumi.StringOutput `pulumi:"kibanaUrl"`
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
	MultiZoneInfos InstanceMultiZoneInfoArrayOutput `pulumi:"multiZoneInfos"`
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	NodeInfoLists InstanceNodeInfoListArrayOutput `pulumi:"nodeInfoLists"`
	// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
	Password pulumi.StringOutput `pulumi:"password"`
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
	RenewFlag pulumi.StringPtrOutput `pulumi:"renewFlag"`
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	Version pulumi.StringOutput `pulumi:"version"`
	// The ID of a VPC network.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Visual node configuration.
	WebNodeTypeInfos InstanceWebNodeTypeInfoArrayOutput `pulumi:"webNodeTypeInfos"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeInfoLists == nil {
		return nil, errors.New("invalid value for required argument 'NodeInfoLists'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Elasticsearch/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("tencentcloud:Elasticsearch/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	BasicSecurityType *int `pulumi:"basicSecurityType"`
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
	ChargePeriod *int `pulumi:"chargePeriod"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	ChargeType *string `pulumi:"chargeType"`
	// Instance creation time.
	CreateTime *string `pulumi:"createTime"`
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	DeployMode *int `pulumi:"deployMode"`
	// Elasticsearch domain name.
	ElasticsearchDomain *string `pulumi:"elasticsearchDomain"`
	// Elasticsearch port.
	ElasticsearchPort *int `pulumi:"elasticsearchPort"`
	// Elasticsearch VIP.
	ElasticsearchVip *string `pulumi:"elasticsearchVip"`
	// Kibana Access Control Configuration.
	EsAcl *InstanceEsAcl `pulumi:"esAcl"`
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName *string `pulumi:"instanceName"`
	// Kibana public network access status. Valid values are `OPEN` and `CLOSE`.
	KibanaPublicAccess *string `pulumi:"kibanaPublicAccess"`
	// Kibana access URL.
	KibanaUrl *string `pulumi:"kibanaUrl"`
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	LicenseType *string `pulumi:"licenseType"`
	// Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
	MultiZoneInfos []InstanceMultiZoneInfo `pulumi:"multiZoneInfos"`
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	NodeInfoLists []InstanceNodeInfoList `pulumi:"nodeInfoLists"`
	// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
	Password *string `pulumi:"password"`
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
	RenewFlag *string `pulumi:"renewFlag"`
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	Tags map[string]interface{} `pulumi:"tags"`
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	Version *string `pulumi:"version"`
	// The ID of a VPC network.
	VpcId *string `pulumi:"vpcId"`
	// Visual node configuration.
	WebNodeTypeInfos []InstanceWebNodeTypeInfo `pulumi:"webNodeTypeInfos"`
}

type InstanceState struct {
	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	AvailabilityZone pulumi.StringPtrInput
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	BasicSecurityType pulumi.IntPtrInput
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
	ChargePeriod pulumi.IntPtrInput
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	ChargeType pulumi.StringPtrInput
	// Instance creation time.
	CreateTime pulumi.StringPtrInput
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	DeployMode pulumi.IntPtrInput
	// Elasticsearch domain name.
	ElasticsearchDomain pulumi.StringPtrInput
	// Elasticsearch port.
	ElasticsearchPort pulumi.IntPtrInput
	// Elasticsearch VIP.
	ElasticsearchVip pulumi.StringPtrInput
	// Kibana Access Control Configuration.
	EsAcl InstanceEsAclPtrInput
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName pulumi.StringPtrInput
	// Kibana public network access status. Valid values are `OPEN` and `CLOSE`.
	KibanaPublicAccess pulumi.StringPtrInput
	// Kibana access URL.
	KibanaUrl pulumi.StringPtrInput
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	LicenseType pulumi.StringPtrInput
	// Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
	MultiZoneInfos InstanceMultiZoneInfoArrayInput
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	NodeInfoLists InstanceNodeInfoListArrayInput
	// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
	Password pulumi.StringPtrInput
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
	RenewFlag pulumi.StringPtrInput
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	Tags pulumi.MapInput
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	Version pulumi.StringPtrInput
	// The ID of a VPC network.
	VpcId pulumi.StringPtrInput
	// Visual node configuration.
	WebNodeTypeInfos InstanceWebNodeTypeInfoArrayInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	BasicSecurityType *int `pulumi:"basicSecurityType"`
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
	ChargePeriod *int `pulumi:"chargePeriod"`
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	ChargeType *string `pulumi:"chargeType"`
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	DeployMode *int `pulumi:"deployMode"`
	// Kibana Access Control Configuration.
	EsAcl *InstanceEsAcl `pulumi:"esAcl"`
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName *string `pulumi:"instanceName"`
	// Kibana public network access status. Valid values are `OPEN` and `CLOSE`.
	KibanaPublicAccess *string `pulumi:"kibanaPublicAccess"`
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	LicenseType *string `pulumi:"licenseType"`
	// Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
	MultiZoneInfos []InstanceMultiZoneInfo `pulumi:"multiZoneInfos"`
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	NodeInfoLists []InstanceNodeInfoList `pulumi:"nodeInfoLists"`
	// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
	Password string `pulumi:"password"`
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
	RenewFlag *string `pulumi:"renewFlag"`
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	SubnetId *string `pulumi:"subnetId"`
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	Tags map[string]interface{} `pulumi:"tags"`
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	Version string `pulumi:"version"`
	// The ID of a VPC network.
	VpcId string `pulumi:"vpcId"`
	// Visual node configuration.
	WebNodeTypeInfos []InstanceWebNodeTypeInfo `pulumi:"webNodeTypeInfos"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
	AvailabilityZone pulumi.StringPtrInput
	// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
	BasicSecurityType pulumi.IntPtrInput
	// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
	ChargePeriod pulumi.IntPtrInput
	// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
	ChargeType pulumi.StringPtrInput
	// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
	DeployMode pulumi.IntPtrInput
	// Kibana Access Control Configuration.
	EsAcl InstanceEsAclPtrInput
	// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
	InstanceName pulumi.StringPtrInput
	// Kibana public network access status. Valid values are `OPEN` and `CLOSE`.
	KibanaPublicAccess pulumi.StringPtrInput
	// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
	LicenseType pulumi.StringPtrInput
	// Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
	MultiZoneInfos InstanceMultiZoneInfoArrayInput
	// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
	NodeInfoLists InstanceNodeInfoListArrayInput
	// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
	Password pulumi.StringInput
	// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
	RenewFlag pulumi.StringPtrInput
	// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
	SubnetId pulumi.StringPtrInput
	// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
	Tags pulumi.MapInput
	// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
	Version pulumi.StringInput
	// The ID of a VPC network.
	VpcId pulumi.StringInput
	// Visual node configuration.
	WebNodeTypeInfos InstanceWebNodeTypeInfoArrayInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Availability zone. When create multi-az es, this parameter must be omitted or `-`.
func (o InstanceOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// Whether to enable X-Pack security authentication in Basic Edition 6.8 and above. Valid values are `1` and `2`. `1` is disabled, `2` is enabled, and default value is `1`. Notice: this parameter is only take effect on `basic` license.
func (o InstanceOutput) BasicSecurityType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.BasicSecurityType }).(pulumi.IntPtrOutput)
}

// The tenancy of the prepaid instance, and uint is month. NOTE: it only works when chargeType is set to `PREPAID`.
func (o InstanceOutput) ChargePeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.ChargePeriod }).(pulumi.IntPtrOutput)
}

// The charge type of instance. Valid values are `PREPAID` and `POSTPAID_BY_HOUR`.
func (o InstanceOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// Instance creation time.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Cluster deployment mode. Valid values are `0` and `1`. `0` is single-AZ deployment, and `1` is multi-AZ deployment. Default value is `0`.
func (o InstanceOutput) DeployMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.DeployMode }).(pulumi.IntPtrOutput)
}

// Elasticsearch domain name.
func (o InstanceOutput) ElasticsearchDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ElasticsearchDomain }).(pulumi.StringOutput)
}

// Elasticsearch port.
func (o InstanceOutput) ElasticsearchPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.ElasticsearchPort }).(pulumi.IntOutput)
}

// Elasticsearch VIP.
func (o InstanceOutput) ElasticsearchVip() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ElasticsearchVip }).(pulumi.StringOutput)
}

// Kibana Access Control Configuration.
func (o InstanceOutput) EsAcl() InstanceEsAclOutput {
	return o.ApplyT(func(v *Instance) InstanceEsAclOutput { return v.EsAcl }).(InstanceEsAclOutput)
}

// Name of the instance, which can contain 1 to 50 English letters, Chinese characters, digits, dashes(-), or underscores(_).
func (o InstanceOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// Kibana public network access status. Valid values are `OPEN` and `CLOSE`.
func (o InstanceOutput) KibanaPublicAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.KibanaPublicAccess }).(pulumi.StringOutput)
}

// Kibana access URL.
func (o InstanceOutput) KibanaUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.KibanaUrl }).(pulumi.StringOutput)
}

// License type. Valid values are `oss`, `basic` and `platinum`. The default value is `platinum`.
func (o InstanceOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// Details of AZs in multi-AZ deployment mode (which is required when deployMode is `1`).
func (o InstanceOutput) MultiZoneInfos() InstanceMultiZoneInfoArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceMultiZoneInfoArrayOutput { return v.MultiZoneInfos }).(InstanceMultiZoneInfoArrayOutput)
}

// Node information list, which is used to describe the specification information of various types of nodes in the cluster, such as node type, node quantity, node specification, disk type, and disk size.
func (o InstanceOutput) NodeInfoLists() InstanceNodeInfoListArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceNodeInfoListArrayOutput { return v.NodeInfoLists }).(InstanceNodeInfoListArrayOutput)
}

// Password to an instance, the password needs to be 8 to 16 characters, including at least two items ([a-z,A-Z], [0-9] and [-!@#$%&^*+=_:;,.?] special symbols.
func (o InstanceOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// When enabled, the instance will be renew automatically when it reach the end of the prepaid tenancy. Valid values are `RENEW_FLAG_AUTO` and `RENEW_FLAG_MANUAL`. NOTE: it only works when chargeType is set to `PREPAID`.
func (o InstanceOutput) RenewFlag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RenewFlag }).(pulumi.StringPtrOutput)
}

// The ID of a VPC subnetwork. When create multi-az es, this parameter must be omitted or `-`.
func (o InstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// A mapping of tags to assign to the instance. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).
func (o InstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Instance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Version of the instance. Valid values are `5.6.4`, `6.4.3`, `6.8.2`, `7.5.1` and `7.10.1`.
func (o InstanceOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The ID of a VPC network.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Visual node configuration.
func (o InstanceOutput) WebNodeTypeInfos() InstanceWebNodeTypeInfoArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceWebNodeTypeInfoArrayOutput { return v.WebNodeTypeInfos }).(InstanceWebNodeTypeInfoArrayOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
