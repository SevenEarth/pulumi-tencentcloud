// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query the available database specifications for different regions. And a maximum of 20 requests can be initiated per second for this query.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mysql.GetZoneConfig(ctx, &mysql.GetZoneConfigArgs{
//				Region:           pulumi.StringRef("ap-guangzhou"),
//				ResultOutputFile: pulumi.StringRef("mytestpath"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetZoneConfig(ctx *pulumi.Context, args *GetZoneConfigArgs, opts ...pulumi.InvokeOption) (*GetZoneConfigResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetZoneConfigResult
	err := ctx.Invoke("tencentcloud:Mysql/getZoneConfig:getZoneConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getZoneConfig.
type GetZoneConfigArgs struct {
	// Region parameter, which is used to identify the region to which the data you want to work with belongs.
	Region *string `pulumi:"region"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getZoneConfig.
type GetZoneConfigResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of zone config. Each element contains the following attributes:
	Lists            []GetZoneConfigList `pulumi:"lists"`
	Region           *string             `pulumi:"region"`
	ResultOutputFile *string             `pulumi:"resultOutputFile"`
}

func GetZoneConfigOutput(ctx *pulumi.Context, args GetZoneConfigOutputArgs, opts ...pulumi.InvokeOption) GetZoneConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetZoneConfigResult, error) {
			args := v.(GetZoneConfigArgs)
			r, err := GetZoneConfig(ctx, &args, opts...)
			var s GetZoneConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetZoneConfigResultOutput)
}

// A collection of arguments for invoking getZoneConfig.
type GetZoneConfigOutputArgs struct {
	// Region parameter, which is used to identify the region to which the data you want to work with belongs.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetZoneConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneConfigArgs)(nil)).Elem()
}

// A collection of values returned by getZoneConfig.
type GetZoneConfigResultOutput struct{ *pulumi.OutputState }

func (GetZoneConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetZoneConfigResult)(nil)).Elem()
}

func (o GetZoneConfigResultOutput) ToGetZoneConfigResultOutput() GetZoneConfigResultOutput {
	return o
}

func (o GetZoneConfigResultOutput) ToGetZoneConfigResultOutputWithContext(ctx context.Context) GetZoneConfigResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetZoneConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetZoneConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of zone config. Each element contains the following attributes:
func (o GetZoneConfigResultOutput) Lists() GetZoneConfigListArrayOutput {
	return o.ApplyT(func(v GetZoneConfigResult) []GetZoneConfigList { return v.Lists }).(GetZoneConfigListArrayOutput)
}

func (o GetZoneConfigResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZoneConfigResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetZoneConfigResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetZoneConfigResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetZoneConfigResultOutput{})
}
