// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of as limits
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
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := As.GetLimits(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetLimits(ctx *pulumi.Context, args *GetLimitsArgs, opts ...pulumi.InvokeOption) (*GetLimitsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLimitsResult
	err := ctx.Invoke("tencentcloud:As/getLimits:getLimits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLimits.
type GetLimitsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getLimits.
type GetLimitsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Maximum number of auto scaling groups allowed for creation by the user account.
	MaxNumberOfAutoScalingGroups int `pulumi:"maxNumberOfAutoScalingGroups"`
	// Maximum number of launch configurations allowed for creation by the user account.
	MaxNumberOfLaunchConfigurations int `pulumi:"maxNumberOfLaunchConfigurations"`
	// Current number of auto scaling groups under the user account.
	NumberOfAutoScalingGroups int `pulumi:"numberOfAutoScalingGroups"`
	// Current number of launch configurations under the user account.
	NumberOfLaunchConfigurations int     `pulumi:"numberOfLaunchConfigurations"`
	ResultOutputFile             *string `pulumi:"resultOutputFile"`
}

func GetLimitsOutput(ctx *pulumi.Context, args GetLimitsOutputArgs, opts ...pulumi.InvokeOption) GetLimitsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLimitsResult, error) {
			args := v.(GetLimitsArgs)
			r, err := GetLimits(ctx, &args, opts...)
			var s GetLimitsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLimitsResultOutput)
}

// A collection of arguments for invoking getLimits.
type GetLimitsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetLimitsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLimitsArgs)(nil)).Elem()
}

// A collection of values returned by getLimits.
type GetLimitsResultOutput struct{ *pulumi.OutputState }

func (GetLimitsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLimitsResult)(nil)).Elem()
}

func (o GetLimitsResultOutput) ToGetLimitsResultOutput() GetLimitsResultOutput {
	return o
}

func (o GetLimitsResultOutput) ToGetLimitsResultOutputWithContext(ctx context.Context) GetLimitsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLimitsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLimitsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Maximum number of auto scaling groups allowed for creation by the user account.
func (o GetLimitsResultOutput) MaxNumberOfAutoScalingGroups() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsResult) int { return v.MaxNumberOfAutoScalingGroups }).(pulumi.IntOutput)
}

// Maximum number of launch configurations allowed for creation by the user account.
func (o GetLimitsResultOutput) MaxNumberOfLaunchConfigurations() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsResult) int { return v.MaxNumberOfLaunchConfigurations }).(pulumi.IntOutput)
}

// Current number of auto scaling groups under the user account.
func (o GetLimitsResultOutput) NumberOfAutoScalingGroups() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsResult) int { return v.NumberOfAutoScalingGroups }).(pulumi.IntOutput)
}

// Current number of launch configurations under the user account.
func (o GetLimitsResultOutput) NumberOfLaunchConfigurations() pulumi.IntOutput {
	return o.ApplyT(func(v GetLimitsResult) int { return v.NumberOfLaunchConfigurations }).(pulumi.IntOutput)
}

func (o GetLimitsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLimitsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLimitsResultOutput{})
}
