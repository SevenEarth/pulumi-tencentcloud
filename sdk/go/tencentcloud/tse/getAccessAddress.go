// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of tse accessAddress
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tse.GetAccessAddress(ctx, &tse.GetAccessAddressArgs{
//				EngineRegion: pulumi.StringRef("ap-guangzhou"),
//				InstanceId:   "ins-7eb7eea7",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetAccessAddress(ctx *pulumi.Context, args *GetAccessAddressArgs, opts ...pulumi.InvokeOption) (*GetAccessAddressResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccessAddressResult
	err := ctx.Invoke("tencentcloud:Tse/getAccessAddress:getAccessAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessAddress.
type GetAccessAddressArgs struct {
	// Deploy region.
	EngineRegion *string `pulumi:"engineRegion"`
	// engine instance Id.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Subnet ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
	SubnetId *string `pulumi:"subnetId"`
	// VPC ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
	VpcId *string `pulumi:"vpcId"`
	// Name of other engine components(pushgateway, polaris-limiter).
	Workload *string `pulumi:"workload"`
}

// A collection of values returned by getAccessAddress.
type GetAccessAddressResult struct {
	// Console public network access addressNote: This field may return null, indicating that a valid value is not available.
	ConsoleInternetAddress string `pulumi:"consoleInternetAddress"`
	// Console public network bandwidthNote: This field may return null, indicating that a valid value is not available.
	ConsoleInternetBandWidth int `pulumi:"consoleInternetBandWidth"`
	// Console Intranet access addressNote: This field may return null, indicating that a valid value is not available.
	ConsoleIntranetAddress string  `pulumi:"consoleIntranetAddress"`
	EngineRegion           *string `pulumi:"engineRegion"`
	// Apollo Multi-environment public ip address.
	EnvAddressInfos []GetAccessAddressEnvAddressInfo `pulumi:"envAddressInfos"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Public access address.
	InternetAddress string `pulumi:"internetAddress"`
	// Client public network bandwidthNote: This field may return null, indicating that a valid value is not available.
	InternetBandWidth int `pulumi:"internetBandWidth"`
	// VPC access IP address listNote: This field may return null, indicating that a valid value is not available.
	IntranetAddress string `pulumi:"intranetAddress"`
	// Access IP address of the Polaris traffic limiting server nodeNote: This field may return null, indicating that a valid value is not available.
	LimiterAddressInfos []GetAccessAddressLimiterAddressInfo `pulumi:"limiterAddressInfos"`
	ResultOutputFile    *string                              `pulumi:"resultOutputFile"`
	SubnetId            *string                              `pulumi:"subnetId"`
	VpcId               *string                              `pulumi:"vpcId"`
	Workload            *string                              `pulumi:"workload"`
}

func GetAccessAddressOutput(ctx *pulumi.Context, args GetAccessAddressOutputArgs, opts ...pulumi.InvokeOption) GetAccessAddressResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessAddressResult, error) {
			args := v.(GetAccessAddressArgs)
			r, err := GetAccessAddress(ctx, &args, opts...)
			var s GetAccessAddressResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccessAddressResultOutput)
}

// A collection of arguments for invoking getAccessAddress.
type GetAccessAddressOutputArgs struct {
	// Deploy region.
	EngineRegion pulumi.StringPtrInput `pulumi:"engineRegion"`
	// engine instance Id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Subnet ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// VPC ID, Zookeeper does not need to pass vpcid and subnetid; nacos and Polaris need to pass vpcid and subnetid.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// Name of other engine components(pushgateway, polaris-limiter).
	Workload pulumi.StringPtrInput `pulumi:"workload"`
}

func (GetAccessAddressOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessAddressArgs)(nil)).Elem()
}

// A collection of values returned by getAccessAddress.
type GetAccessAddressResultOutput struct{ *pulumi.OutputState }

func (GetAccessAddressResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessAddressResult)(nil)).Elem()
}

func (o GetAccessAddressResultOutput) ToGetAccessAddressResultOutput() GetAccessAddressResultOutput {
	return o
}

func (o GetAccessAddressResultOutput) ToGetAccessAddressResultOutputWithContext(ctx context.Context) GetAccessAddressResultOutput {
	return o
}

// Console public network access addressNote: This field may return null, indicating that a valid value is not available.
func (o GetAccessAddressResultOutput) ConsoleInternetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessAddressResult) string { return v.ConsoleInternetAddress }).(pulumi.StringOutput)
}

// Console public network bandwidthNote: This field may return null, indicating that a valid value is not available.
func (o GetAccessAddressResultOutput) ConsoleInternetBandWidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetAccessAddressResult) int { return v.ConsoleInternetBandWidth }).(pulumi.IntOutput)
}

// Console Intranet access addressNote: This field may return null, indicating that a valid value is not available.
func (o GetAccessAddressResultOutput) ConsoleIntranetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessAddressResult) string { return v.ConsoleIntranetAddress }).(pulumi.StringOutput)
}

func (o GetAccessAddressResultOutput) EngineRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessAddressResult) *string { return v.EngineRegion }).(pulumi.StringPtrOutput)
}

// Apollo Multi-environment public ip address.
func (o GetAccessAddressResultOutput) EnvAddressInfos() GetAccessAddressEnvAddressInfoArrayOutput {
	return o.ApplyT(func(v GetAccessAddressResult) []GetAccessAddressEnvAddressInfo { return v.EnvAddressInfos }).(GetAccessAddressEnvAddressInfoArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessAddressResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessAddressResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccessAddressResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessAddressResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Public access address.
func (o GetAccessAddressResultOutput) InternetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessAddressResult) string { return v.InternetAddress }).(pulumi.StringOutput)
}

// Client public network bandwidthNote: This field may return null, indicating that a valid value is not available.
func (o GetAccessAddressResultOutput) InternetBandWidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetAccessAddressResult) int { return v.InternetBandWidth }).(pulumi.IntOutput)
}

// VPC access IP address listNote: This field may return null, indicating that a valid value is not available.
func (o GetAccessAddressResultOutput) IntranetAddress() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessAddressResult) string { return v.IntranetAddress }).(pulumi.StringOutput)
}

// Access IP address of the Polaris traffic limiting server nodeNote: This field may return null, indicating that a valid value is not available.
func (o GetAccessAddressResultOutput) LimiterAddressInfos() GetAccessAddressLimiterAddressInfoArrayOutput {
	return o.ApplyT(func(v GetAccessAddressResult) []GetAccessAddressLimiterAddressInfo { return v.LimiterAddressInfos }).(GetAccessAddressLimiterAddressInfoArrayOutput)
}

func (o GetAccessAddressResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessAddressResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetAccessAddressResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessAddressResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o GetAccessAddressResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessAddressResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func (o GetAccessAddressResultOutput) Workload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessAddressResult) *string { return v.Workload }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessAddressResultOutput{})
}
