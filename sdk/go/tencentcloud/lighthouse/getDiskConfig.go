// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of lighthouse diskConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Lighthouse.GetDiskConfig(ctx, &lighthouse.GetDiskConfigArgs{
//				Filters: []lighthouse.GetDiskConfigFilter{
//					lighthouse.GetDiskConfigFilter{
//						Name: "zone",
//						Values: []string{
//							"ap-guangzhou-3",
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDiskConfig(ctx *pulumi.Context, args *GetDiskConfigArgs, opts ...pulumi.InvokeOption) (*GetDiskConfigResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDiskConfigResult
	err := ctx.Invoke("tencentcloud:Lighthouse/getDiskConfig:getDiskConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDiskConfig.
type GetDiskConfigArgs struct {
	// Filter list.zoneFilter by availability zone.Type: StringRequired: no.
	Filters []GetDiskConfigFilter `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDiskConfig.
type GetDiskConfigResult struct {
	// List of cloud disk configurations.
	DiskConfigSets []GetDiskConfigDiskConfigSet `pulumi:"diskConfigSets"`
	Filters        []GetDiskConfigFilter        `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDiskConfigOutput(ctx *pulumi.Context, args GetDiskConfigOutputArgs, opts ...pulumi.InvokeOption) GetDiskConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDiskConfigResult, error) {
			args := v.(GetDiskConfigArgs)
			r, err := GetDiskConfig(ctx, &args, opts...)
			var s GetDiskConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDiskConfigResultOutput)
}

// A collection of arguments for invoking getDiskConfig.
type GetDiskConfigOutputArgs struct {
	// Filter list.zoneFilter by availability zone.Type: StringRequired: no.
	Filters GetDiskConfigFilterArrayInput `pulumi:"filters"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDiskConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiskConfigArgs)(nil)).Elem()
}

// A collection of values returned by getDiskConfig.
type GetDiskConfigResultOutput struct{ *pulumi.OutputState }

func (GetDiskConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDiskConfigResult)(nil)).Elem()
}

func (o GetDiskConfigResultOutput) ToGetDiskConfigResultOutput() GetDiskConfigResultOutput {
	return o
}

func (o GetDiskConfigResultOutput) ToGetDiskConfigResultOutputWithContext(ctx context.Context) GetDiskConfigResultOutput {
	return o
}

// List of cloud disk configurations.
func (o GetDiskConfigResultOutput) DiskConfigSets() GetDiskConfigDiskConfigSetArrayOutput {
	return o.ApplyT(func(v GetDiskConfigResult) []GetDiskConfigDiskConfigSet { return v.DiskConfigSets }).(GetDiskConfigDiskConfigSetArrayOutput)
}

func (o GetDiskConfigResultOutput) Filters() GetDiskConfigFilterArrayOutput {
	return o.ApplyT(func(v GetDiskConfigResult) []GetDiskConfigFilter { return v.Filters }).(GetDiskConfigFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDiskConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDiskConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDiskConfigResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDiskConfigResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDiskConfigResultOutput{})
}
