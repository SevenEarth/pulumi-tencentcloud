// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a configuration for an AS (Auto scaling) instance.
//
// > **NOTE:**  In order to ensure the integrity of customer data, if the cvm instance was destroyed due to shrinking, it will keep the cbs associate with cvm by default. If you want to destroy together, please set `deleteWithInstance` to `true`.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := Images.GetInstance(ctx, &images.GetInstanceArgs{
//				ImageTypes: []string{
//					"PUBLIC_IMAGE",
//				},
//				OsName: pulumi.StringRef("TencentOS Server 3.2 (Final)"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = As.NewScalingConfig(ctx, "exampleScalingConfig", &As.ScalingConfigArgs{
//				ConfigurationName: pulumi.String("example-launch-configuration"),
//				ImageId:           pulumi.String(exampleInstance.Images[0].ImageId),
//				InstanceTypes: pulumi.StringArray{
//					pulumi.String("SA1.SMALL1"),
//				},
//				ProjectId:      pulumi.Int(0),
//				SystemDiskType: pulumi.String("CLOUD_PREMIUM"),
//				SystemDiskSize: pulumi.Int(50),
//				DataDisks: as.ScalingConfigDataDiskArray{
//					&as.ScalingConfigDataDiskArgs{
//						DiskType: pulumi.String("CLOUD_PREMIUM"),
//						DiskSize: pulumi.Int(50),
//					},
//				},
//				InternetChargeType:      pulumi.String("TRAFFIC_POSTPAID_BY_HOUR"),
//				InternetMaxBandwidthOut: pulumi.Int(10),
//				PublicIpAssigned:        pulumi.Bool(true),
//				Password:                pulumi.String("Test@123#"),
//				EnhancedSecurityService: pulumi.Bool(false),
//				EnhancedMonitorService:  pulumi.Bool(false),
//				UserData:                pulumi.String("dGVzdA=="),
//				HostNameSettings: &as.ScalingConfigHostNameSettingsArgs{
//					HostName:      pulumi.String("host-name-test"),
//					HostNameStyle: pulumi.String("UNIQUE"),
//				},
//				InstanceTags: pulumi.Map{
//					"tag": pulumi.Any("example"),
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
// ### charge type
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := Images.GetInstance(ctx, &images.GetInstanceArgs{
//				ImageTypes: []string{
//					"PUBLIC_IMAGE",
//				},
//				OsName: pulumi.StringRef("TencentOS Server 3.2 (Final)"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = As.NewScalingConfig(ctx, "exampleScalingConfig", &As.ScalingConfigArgs{
//				ConfigurationName: pulumi.String("launch-configuration"),
//				ImageId:           pulumi.String(exampleInstance.Images[0].ImageId),
//				InstanceTypes: pulumi.StringArray{
//					pulumi.String("SA1.SMALL1"),
//				},
//				InstanceChargeType: pulumi.String("SPOTPAID"),
//				SpotInstanceType:   pulumi.String("one-time"),
//				SpotMaxPrice:       pulumi.String("1000"),
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
// AutoScaling Configuration can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:As/scalingConfig:ScalingConfig example asc-n32ymck2
// ```
type ScalingConfig struct {
	pulumi.CustomResourceState

	// CAM role name authorized to access.
	CamRoleName pulumi.StringPtrOutput `pulumi:"camRoleName"`
	// Name of a launch configuration.
	ConfigurationName pulumi.StringOutput `pulumi:"configurationName"`
	// The time when the launch configuration was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Configurations of data disk.
	DataDisks ScalingConfigDataDiskArrayOutput `pulumi:"dataDisks"`
	// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
	DiskTypePolicy pulumi.StringPtrOutput `pulumi:"diskTypePolicy"`
	// To specify whether to enable cloud monitor service. Default is `TRUE`.
	EnhancedMonitorService pulumi.BoolPtrOutput `pulumi:"enhancedMonitorService"`
	// To specify whether to enable cloud security service. Default is `TRUE`.
	EnhancedSecurityService pulumi.BoolPtrOutput `pulumi:"enhancedSecurityService"`
	// Related settings of the cloud server hostname (HostName).
	HostNameSettings ScalingConfigHostNameSettingsPtrOutput `pulumi:"hostNameSettings"`
	// An available image ID for a cvm instance.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spotInstanceType` and `spotMaxPrice` at the same time.
	InstanceChargeType pulumi.StringPtrOutput `pulumi:"instanceChargeType"`
	// The tenancy (in month) of the prepaid instance, NOTE: it only works when instanceChargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	InstanceChargeTypePrepaidPeriod pulumi.IntPtrOutput `pulumi:"instanceChargeTypePrepaidPeriod"`
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instanceChargeType is set to `PREPAID`.
	InstanceChargeTypePrepaidRenewFlag pulumi.StringOutput `pulumi:"instanceChargeTypePrepaidRenewFlag"`
	// Settings of CVM instance names.
	InstanceNameSettings ScalingConfigInstanceNameSettingsPtrOutput `pulumi:"instanceNameSettings"`
	// A list of tags used to associate different resources.
	InstanceTags pulumi.MapOutput `pulumi:"instanceTags"`
	// Specified types of CVM instances.
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType pulumi.StringPtrOutput `pulumi:"internetChargeType"`
	// Max bandwidth of Internet access in Mbps. Default is `0`.
	InternetMaxBandwidthOut pulumi.IntPtrOutput `pulumi:"internetMaxBandwidthOut"`
	// Specify whether to keep original settings of a CVM image. And it can't be used with password or keyIds together.
	KeepImageLogin pulumi.BoolPtrOutput `pulumi:"keepImageLogin"`
	// ID list of keys.
	KeyIds pulumi.StringArrayOutput `pulumi:"keyIds"`
	// Password to access.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Specifys to which project the configuration belongs.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// Specify whether to assign an Internet IP address.
	PublicIpAssigned pulumi.BoolPtrOutput `pulumi:"publicIpAssigned"`
	// Security groups to which a CVM instance belongs.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// Type of spot instance, only support `one-time` now. Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotInstanceType pulumi.StringPtrOutput `pulumi:"spotInstanceType"`
	// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotMaxPrice pulumi.StringPtrOutput `pulumi:"spotMaxPrice"`
	// Current statues of a launch configuration.
	Status pulumi.StringOutput `pulumi:"status"`
	// Volume of system disk in GB. Default is `50`.
	SystemDiskSize pulumi.IntPtrOutput `pulumi:"systemDiskSize"`
	// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when diskTypePolicy is ORIGINAL.
	SystemDiskType pulumi.StringPtrOutput `pulumi:"systemDiskType"`
	// ase64-encoded User Data text, the length limit is 16KB.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
}

// NewScalingConfig registers a new resource with the given unique name, arguments, and options.
func NewScalingConfig(ctx *pulumi.Context,
	name string, args *ScalingConfigArgs, opts ...pulumi.ResourceOption) (*ScalingConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.InstanceTypes == nil {
		return nil, errors.New("invalid value for required argument 'InstanceTypes'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ScalingConfig
	err := ctx.RegisterResource("tencentcloud:As/scalingConfig:ScalingConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingConfig gets an existing ScalingConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingConfigState, opts ...pulumi.ResourceOption) (*ScalingConfig, error) {
	var resource ScalingConfig
	err := ctx.ReadResource("tencentcloud:As/scalingConfig:ScalingConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingConfig resources.
type scalingConfigState struct {
	// CAM role name authorized to access.
	CamRoleName *string `pulumi:"camRoleName"`
	// Name of a launch configuration.
	ConfigurationName *string `pulumi:"configurationName"`
	// The time when the launch configuration was created.
	CreateTime *string `pulumi:"createTime"`
	// Configurations of data disk.
	DataDisks []ScalingConfigDataDisk `pulumi:"dataDisks"`
	// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
	DiskTypePolicy *string `pulumi:"diskTypePolicy"`
	// To specify whether to enable cloud monitor service. Default is `TRUE`.
	EnhancedMonitorService *bool `pulumi:"enhancedMonitorService"`
	// To specify whether to enable cloud security service. Default is `TRUE`.
	EnhancedSecurityService *bool `pulumi:"enhancedSecurityService"`
	// Related settings of the cloud server hostname (HostName).
	HostNameSettings *ScalingConfigHostNameSettings `pulumi:"hostNameSettings"`
	// An available image ID for a cvm instance.
	ImageId *string `pulumi:"imageId"`
	// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spotInstanceType` and `spotMaxPrice` at the same time.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The tenancy (in month) of the prepaid instance, NOTE: it only works when instanceChargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	InstanceChargeTypePrepaidPeriod *int `pulumi:"instanceChargeTypePrepaidPeriod"`
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instanceChargeType is set to `PREPAID`.
	InstanceChargeTypePrepaidRenewFlag *string `pulumi:"instanceChargeTypePrepaidRenewFlag"`
	// Settings of CVM instance names.
	InstanceNameSettings *ScalingConfigInstanceNameSettings `pulumi:"instanceNameSettings"`
	// A list of tags used to associate different resources.
	InstanceTags map[string]interface{} `pulumi:"instanceTags"`
	// Specified types of CVM instances.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Max bandwidth of Internet access in Mbps. Default is `0`.
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// Specify whether to keep original settings of a CVM image. And it can't be used with password or keyIds together.
	KeepImageLogin *bool `pulumi:"keepImageLogin"`
	// ID list of keys.
	KeyIds []string `pulumi:"keyIds"`
	// Password to access.
	Password *string `pulumi:"password"`
	// Specifys to which project the configuration belongs.
	ProjectId *int `pulumi:"projectId"`
	// Specify whether to assign an Internet IP address.
	PublicIpAssigned *bool `pulumi:"publicIpAssigned"`
	// Security groups to which a CVM instance belongs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Type of spot instance, only support `one-time` now. Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotInstanceType *string `pulumi:"spotInstanceType"`
	// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotMaxPrice *string `pulumi:"spotMaxPrice"`
	// Current statues of a launch configuration.
	Status *string `pulumi:"status"`
	// Volume of system disk in GB. Default is `50`.
	SystemDiskSize *int `pulumi:"systemDiskSize"`
	// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when diskTypePolicy is ORIGINAL.
	SystemDiskType *string `pulumi:"systemDiskType"`
	// ase64-encoded User Data text, the length limit is 16KB.
	UserData *string `pulumi:"userData"`
}

type ScalingConfigState struct {
	// CAM role name authorized to access.
	CamRoleName pulumi.StringPtrInput
	// Name of a launch configuration.
	ConfigurationName pulumi.StringPtrInput
	// The time when the launch configuration was created.
	CreateTime pulumi.StringPtrInput
	// Configurations of data disk.
	DataDisks ScalingConfigDataDiskArrayInput
	// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
	DiskTypePolicy pulumi.StringPtrInput
	// To specify whether to enable cloud monitor service. Default is `TRUE`.
	EnhancedMonitorService pulumi.BoolPtrInput
	// To specify whether to enable cloud security service. Default is `TRUE`.
	EnhancedSecurityService pulumi.BoolPtrInput
	// Related settings of the cloud server hostname (HostName).
	HostNameSettings ScalingConfigHostNameSettingsPtrInput
	// An available image ID for a cvm instance.
	ImageId pulumi.StringPtrInput
	// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spotInstanceType` and `spotMaxPrice` at the same time.
	InstanceChargeType pulumi.StringPtrInput
	// The tenancy (in month) of the prepaid instance, NOTE: it only works when instanceChargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	InstanceChargeTypePrepaidPeriod pulumi.IntPtrInput
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instanceChargeType is set to `PREPAID`.
	InstanceChargeTypePrepaidRenewFlag pulumi.StringPtrInput
	// Settings of CVM instance names.
	InstanceNameSettings ScalingConfigInstanceNameSettingsPtrInput
	// A list of tags used to associate different resources.
	InstanceTags pulumi.MapInput
	// Specified types of CVM instances.
	InstanceTypes pulumi.StringArrayInput
	// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType pulumi.StringPtrInput
	// Max bandwidth of Internet access in Mbps. Default is `0`.
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// Specify whether to keep original settings of a CVM image. And it can't be used with password or keyIds together.
	KeepImageLogin pulumi.BoolPtrInput
	// ID list of keys.
	KeyIds pulumi.StringArrayInput
	// Password to access.
	Password pulumi.StringPtrInput
	// Specifys to which project the configuration belongs.
	ProjectId pulumi.IntPtrInput
	// Specify whether to assign an Internet IP address.
	PublicIpAssigned pulumi.BoolPtrInput
	// Security groups to which a CVM instance belongs.
	SecurityGroupIds pulumi.StringArrayInput
	// Type of spot instance, only support `one-time` now. Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotInstanceType pulumi.StringPtrInput
	// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotMaxPrice pulumi.StringPtrInput
	// Current statues of a launch configuration.
	Status pulumi.StringPtrInput
	// Volume of system disk in GB. Default is `50`.
	SystemDiskSize pulumi.IntPtrInput
	// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when diskTypePolicy is ORIGINAL.
	SystemDiskType pulumi.StringPtrInput
	// ase64-encoded User Data text, the length limit is 16KB.
	UserData pulumi.StringPtrInput
}

func (ScalingConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigState)(nil)).Elem()
}

type scalingConfigArgs struct {
	// CAM role name authorized to access.
	CamRoleName *string `pulumi:"camRoleName"`
	// Name of a launch configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// Configurations of data disk.
	DataDisks []ScalingConfigDataDisk `pulumi:"dataDisks"`
	// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
	DiskTypePolicy *string `pulumi:"diskTypePolicy"`
	// To specify whether to enable cloud monitor service. Default is `TRUE`.
	EnhancedMonitorService *bool `pulumi:"enhancedMonitorService"`
	// To specify whether to enable cloud security service. Default is `TRUE`.
	EnhancedSecurityService *bool `pulumi:"enhancedSecurityService"`
	// Related settings of the cloud server hostname (HostName).
	HostNameSettings *ScalingConfigHostNameSettings `pulumi:"hostNameSettings"`
	// An available image ID for a cvm instance.
	ImageId string `pulumi:"imageId"`
	// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spotInstanceType` and `spotMaxPrice` at the same time.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The tenancy (in month) of the prepaid instance, NOTE: it only works when instanceChargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	InstanceChargeTypePrepaidPeriod *int `pulumi:"instanceChargeTypePrepaidPeriod"`
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instanceChargeType is set to `PREPAID`.
	InstanceChargeTypePrepaidRenewFlag *string `pulumi:"instanceChargeTypePrepaidRenewFlag"`
	// Settings of CVM instance names.
	InstanceNameSettings *ScalingConfigInstanceNameSettings `pulumi:"instanceNameSettings"`
	// A list of tags used to associate different resources.
	InstanceTags map[string]interface{} `pulumi:"instanceTags"`
	// Specified types of CVM instances.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// Max bandwidth of Internet access in Mbps. Default is `0`.
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// Specify whether to keep original settings of a CVM image. And it can't be used with password or keyIds together.
	KeepImageLogin *bool `pulumi:"keepImageLogin"`
	// ID list of keys.
	KeyIds []string `pulumi:"keyIds"`
	// Password to access.
	Password *string `pulumi:"password"`
	// Specifys to which project the configuration belongs.
	ProjectId *int `pulumi:"projectId"`
	// Specify whether to assign an Internet IP address.
	PublicIpAssigned *bool `pulumi:"publicIpAssigned"`
	// Security groups to which a CVM instance belongs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Type of spot instance, only support `one-time` now. Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotInstanceType *string `pulumi:"spotInstanceType"`
	// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotMaxPrice *string `pulumi:"spotMaxPrice"`
	// Volume of system disk in GB. Default is `50`.
	SystemDiskSize *int `pulumi:"systemDiskSize"`
	// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when diskTypePolicy is ORIGINAL.
	SystemDiskType *string `pulumi:"systemDiskType"`
	// ase64-encoded User Data text, the length limit is 16KB.
	UserData *string `pulumi:"userData"`
}

// The set of arguments for constructing a ScalingConfig resource.
type ScalingConfigArgs struct {
	// CAM role name authorized to access.
	CamRoleName pulumi.StringPtrInput
	// Name of a launch configuration.
	ConfigurationName pulumi.StringInput
	// Configurations of data disk.
	DataDisks ScalingConfigDataDiskArrayInput
	// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
	DiskTypePolicy pulumi.StringPtrInput
	// To specify whether to enable cloud monitor service. Default is `TRUE`.
	EnhancedMonitorService pulumi.BoolPtrInput
	// To specify whether to enable cloud security service. Default is `TRUE`.
	EnhancedSecurityService pulumi.BoolPtrInput
	// Related settings of the cloud server hostname (HostName).
	HostNameSettings ScalingConfigHostNameSettingsPtrInput
	// An available image ID for a cvm instance.
	ImageId pulumi.StringInput
	// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spotInstanceType` and `spotMaxPrice` at the same time.
	InstanceChargeType pulumi.StringPtrInput
	// The tenancy (in month) of the prepaid instance, NOTE: it only works when instanceChargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
	InstanceChargeTypePrepaidPeriod pulumi.IntPtrInput
	// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instanceChargeType is set to `PREPAID`.
	InstanceChargeTypePrepaidRenewFlag pulumi.StringPtrInput
	// Settings of CVM instance names.
	InstanceNameSettings ScalingConfigInstanceNameSettingsPtrInput
	// A list of tags used to associate different resources.
	InstanceTags pulumi.MapInput
	// Specified types of CVM instances.
	InstanceTypes pulumi.StringArrayInput
	// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
	InternetChargeType pulumi.StringPtrInput
	// Max bandwidth of Internet access in Mbps. Default is `0`.
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// Specify whether to keep original settings of a CVM image. And it can't be used with password or keyIds together.
	KeepImageLogin pulumi.BoolPtrInput
	// ID list of keys.
	KeyIds pulumi.StringArrayInput
	// Password to access.
	Password pulumi.StringPtrInput
	// Specifys to which project the configuration belongs.
	ProjectId pulumi.IntPtrInput
	// Specify whether to assign an Internet IP address.
	PublicIpAssigned pulumi.BoolPtrInput
	// Security groups to which a CVM instance belongs.
	SecurityGroupIds pulumi.StringArrayInput
	// Type of spot instance, only support `one-time` now. Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotInstanceType pulumi.StringPtrInput
	// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instanceChargeType is set to `SPOTPAID`.
	SpotMaxPrice pulumi.StringPtrInput
	// Volume of system disk in GB. Default is `50`.
	SystemDiskSize pulumi.IntPtrInput
	// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when diskTypePolicy is ORIGINAL.
	SystemDiskType pulumi.StringPtrInput
	// ase64-encoded User Data text, the length limit is 16KB.
	UserData pulumi.StringPtrInput
}

func (ScalingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigArgs)(nil)).Elem()
}

type ScalingConfigInput interface {
	pulumi.Input

	ToScalingConfigOutput() ScalingConfigOutput
	ToScalingConfigOutputWithContext(ctx context.Context) ScalingConfigOutput
}

func (*ScalingConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingConfig)(nil)).Elem()
}

func (i *ScalingConfig) ToScalingConfigOutput() ScalingConfigOutput {
	return i.ToScalingConfigOutputWithContext(context.Background())
}

func (i *ScalingConfig) ToScalingConfigOutputWithContext(ctx context.Context) ScalingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigOutput)
}

// ScalingConfigArrayInput is an input type that accepts ScalingConfigArray and ScalingConfigArrayOutput values.
// You can construct a concrete instance of `ScalingConfigArrayInput` via:
//
//	ScalingConfigArray{ ScalingConfigArgs{...} }
type ScalingConfigArrayInput interface {
	pulumi.Input

	ToScalingConfigArrayOutput() ScalingConfigArrayOutput
	ToScalingConfigArrayOutputWithContext(context.Context) ScalingConfigArrayOutput
}

type ScalingConfigArray []ScalingConfigInput

func (ScalingConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingConfig)(nil)).Elem()
}

func (i ScalingConfigArray) ToScalingConfigArrayOutput() ScalingConfigArrayOutput {
	return i.ToScalingConfigArrayOutputWithContext(context.Background())
}

func (i ScalingConfigArray) ToScalingConfigArrayOutputWithContext(ctx context.Context) ScalingConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigArrayOutput)
}

// ScalingConfigMapInput is an input type that accepts ScalingConfigMap and ScalingConfigMapOutput values.
// You can construct a concrete instance of `ScalingConfigMapInput` via:
//
//	ScalingConfigMap{ "key": ScalingConfigArgs{...} }
type ScalingConfigMapInput interface {
	pulumi.Input

	ToScalingConfigMapOutput() ScalingConfigMapOutput
	ToScalingConfigMapOutputWithContext(context.Context) ScalingConfigMapOutput
}

type ScalingConfigMap map[string]ScalingConfigInput

func (ScalingConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingConfig)(nil)).Elem()
}

func (i ScalingConfigMap) ToScalingConfigMapOutput() ScalingConfigMapOutput {
	return i.ToScalingConfigMapOutputWithContext(context.Background())
}

func (i ScalingConfigMap) ToScalingConfigMapOutputWithContext(ctx context.Context) ScalingConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigMapOutput)
}

type ScalingConfigOutput struct{ *pulumi.OutputState }

func (ScalingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingConfig)(nil)).Elem()
}

func (o ScalingConfigOutput) ToScalingConfigOutput() ScalingConfigOutput {
	return o
}

func (o ScalingConfigOutput) ToScalingConfigOutputWithContext(ctx context.Context) ScalingConfigOutput {
	return o
}

// CAM role name authorized to access.
func (o ScalingConfigOutput) CamRoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.CamRoleName }).(pulumi.StringPtrOutput)
}

// Name of a launch configuration.
func (o ScalingConfigOutput) ConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringOutput { return v.ConfigurationName }).(pulumi.StringOutput)
}

// The time when the launch configuration was created.
func (o ScalingConfigOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Configurations of data disk.
func (o ScalingConfigOutput) DataDisks() ScalingConfigDataDiskArrayOutput {
	return o.ApplyT(func(v *ScalingConfig) ScalingConfigDataDiskArrayOutput { return v.DataDisks }).(ScalingConfigDataDiskArrayOutput)
}

// Policy of cloud disk type. Valid values: `ORIGINAL` and `AUTOMATIC`. Default is `ORIGINAL`.
func (o ScalingConfigOutput) DiskTypePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.DiskTypePolicy }).(pulumi.StringPtrOutput)
}

// To specify whether to enable cloud monitor service. Default is `TRUE`.
func (o ScalingConfigOutput) EnhancedMonitorService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.BoolPtrOutput { return v.EnhancedMonitorService }).(pulumi.BoolPtrOutput)
}

// To specify whether to enable cloud security service. Default is `TRUE`.
func (o ScalingConfigOutput) EnhancedSecurityService() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.BoolPtrOutput { return v.EnhancedSecurityService }).(pulumi.BoolPtrOutput)
}

// Related settings of the cloud server hostname (HostName).
func (o ScalingConfigOutput) HostNameSettings() ScalingConfigHostNameSettingsPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) ScalingConfigHostNameSettingsPtrOutput { return v.HostNameSettings }).(ScalingConfigHostNameSettingsPtrOutput)
}

// An available image ID for a cvm instance.
func (o ScalingConfigOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spotInstanceType` and `spotMaxPrice` at the same time.
func (o ScalingConfigOutput) InstanceChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.InstanceChargeType }).(pulumi.StringPtrOutput)
}

// The tenancy (in month) of the prepaid instance, NOTE: it only works when instanceChargeType is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
func (o ScalingConfigOutput) InstanceChargeTypePrepaidPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.IntPtrOutput { return v.InstanceChargeTypePrepaidPeriod }).(pulumi.IntPtrOutput)
}

// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instanceChargeType is set to `PREPAID`.
func (o ScalingConfigOutput) InstanceChargeTypePrepaidRenewFlag() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringOutput { return v.InstanceChargeTypePrepaidRenewFlag }).(pulumi.StringOutput)
}

// Settings of CVM instance names.
func (o ScalingConfigOutput) InstanceNameSettings() ScalingConfigInstanceNameSettingsPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) ScalingConfigInstanceNameSettingsPtrOutput { return v.InstanceNameSettings }).(ScalingConfigInstanceNameSettingsPtrOutput)
}

// A list of tags used to associate different resources.
func (o ScalingConfigOutput) InstanceTags() pulumi.MapOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.MapOutput { return v.InstanceTags }).(pulumi.MapOutput)
}

// Specified types of CVM instances.
func (o ScalingConfigOutput) InstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringArrayOutput { return v.InstanceTypes }).(pulumi.StringArrayOutput)
}

// Charge types for network traffic. Valid values: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
func (o ScalingConfigOutput) InternetChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.InternetChargeType }).(pulumi.StringPtrOutput)
}

// Max bandwidth of Internet access in Mbps. Default is `0`.
func (o ScalingConfigOutput) InternetMaxBandwidthOut() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.IntPtrOutput { return v.InternetMaxBandwidthOut }).(pulumi.IntPtrOutput)
}

// Specify whether to keep original settings of a CVM image. And it can't be used with password or keyIds together.
func (o ScalingConfigOutput) KeepImageLogin() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.BoolPtrOutput { return v.KeepImageLogin }).(pulumi.BoolPtrOutput)
}

// ID list of keys.
func (o ScalingConfigOutput) KeyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringArrayOutput { return v.KeyIds }).(pulumi.StringArrayOutput)
}

// Password to access.
func (o ScalingConfigOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Specifys to which project the configuration belongs.
func (o ScalingConfigOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// Specify whether to assign an Internet IP address.
func (o ScalingConfigOutput) PublicIpAssigned() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.BoolPtrOutput { return v.PublicIpAssigned }).(pulumi.BoolPtrOutput)
}

// Security groups to which a CVM instance belongs.
func (o ScalingConfigOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Type of spot instance, only support `one-time` now. Note: it only works when instanceChargeType is set to `SPOTPAID`.
func (o ScalingConfigOutput) SpotInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.SpotInstanceType }).(pulumi.StringPtrOutput)
}

// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instanceChargeType is set to `SPOTPAID`.
func (o ScalingConfigOutput) SpotMaxPrice() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.SpotMaxPrice }).(pulumi.StringPtrOutput)
}

// Current statues of a launch configuration.
func (o ScalingConfigOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Volume of system disk in GB. Default is `50`.
func (o ScalingConfigOutput) SystemDiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.IntPtrOutput { return v.SystemDiskSize }).(pulumi.IntPtrOutput)
}

// Type of a CVM disk. Valid values: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`. valid when diskTypePolicy is ORIGINAL.
func (o ScalingConfigOutput) SystemDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.SystemDiskType }).(pulumi.StringPtrOutput)
}

// ase64-encoded User Data text, the length limit is 16KB.
func (o ScalingConfigOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfig) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

type ScalingConfigArrayOutput struct{ *pulumi.OutputState }

func (ScalingConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingConfig)(nil)).Elem()
}

func (o ScalingConfigArrayOutput) ToScalingConfigArrayOutput() ScalingConfigArrayOutput {
	return o
}

func (o ScalingConfigArrayOutput) ToScalingConfigArrayOutputWithContext(ctx context.Context) ScalingConfigArrayOutput {
	return o
}

func (o ScalingConfigArrayOutput) Index(i pulumi.IntInput) ScalingConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingConfig {
		return vs[0].([]*ScalingConfig)[vs[1].(int)]
	}).(ScalingConfigOutput)
}

type ScalingConfigMapOutput struct{ *pulumi.OutputState }

func (ScalingConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingConfig)(nil)).Elem()
}

func (o ScalingConfigMapOutput) ToScalingConfigMapOutput() ScalingConfigMapOutput {
	return o
}

func (o ScalingConfigMapOutput) ToScalingConfigMapOutputWithContext(ctx context.Context) ScalingConfigMapOutput {
	return o
}

func (o ScalingConfigMapOutput) MapIndex(k pulumi.StringInput) ScalingConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingConfig {
		return vs[0].(map[string]*ScalingConfig)[vs[1].(string)]
	}).(ScalingConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigInput)(nil)).Elem(), &ScalingConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigArrayInput)(nil)).Elem(), ScalingConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigMapInput)(nil)).Elem(), ScalingConfigMap{})
	pulumi.RegisterOutputType(ScalingConfigOutput{})
	pulumi.RegisterOutputType(ScalingConfigArrayOutput{})
	pulumi.RegisterOutputType(ScalingConfigMapOutput{})
}
