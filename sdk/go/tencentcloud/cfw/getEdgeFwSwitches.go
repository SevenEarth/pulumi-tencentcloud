// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfw

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of cfw edgeFwSwitches
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfw"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cfw.GetEdgeFwSwitches(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetEdgeFwSwitches(ctx *pulumi.Context, args *GetEdgeFwSwitchesArgs, opts ...pulumi.InvokeOption) (*GetEdgeFwSwitchesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetEdgeFwSwitchesResult
	err := ctx.Invoke("tencentcloud:Cfw/getEdgeFwSwitches:getEdgeFwSwitches", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEdgeFwSwitches.
type GetEdgeFwSwitchesArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getEdgeFwSwitches.
type GetEdgeFwSwitchesResult struct {
	// Ip switch list.
	Datas []GetEdgeFwSwitchesData `pulumi:"datas"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetEdgeFwSwitchesOutput(ctx *pulumi.Context, args GetEdgeFwSwitchesOutputArgs, opts ...pulumi.InvokeOption) GetEdgeFwSwitchesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetEdgeFwSwitchesResult, error) {
			args := v.(GetEdgeFwSwitchesArgs)
			r, err := GetEdgeFwSwitches(ctx, &args, opts...)
			var s GetEdgeFwSwitchesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetEdgeFwSwitchesResultOutput)
}

// A collection of arguments for invoking getEdgeFwSwitches.
type GetEdgeFwSwitchesOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetEdgeFwSwitchesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeFwSwitchesArgs)(nil)).Elem()
}

// A collection of values returned by getEdgeFwSwitches.
type GetEdgeFwSwitchesResultOutput struct{ *pulumi.OutputState }

func (GetEdgeFwSwitchesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEdgeFwSwitchesResult)(nil)).Elem()
}

func (o GetEdgeFwSwitchesResultOutput) ToGetEdgeFwSwitchesResultOutput() GetEdgeFwSwitchesResultOutput {
	return o
}

func (o GetEdgeFwSwitchesResultOutput) ToGetEdgeFwSwitchesResultOutputWithContext(ctx context.Context) GetEdgeFwSwitchesResultOutput {
	return o
}

// Ip switch list.
func (o GetEdgeFwSwitchesResultOutput) Datas() GetEdgeFwSwitchesDataArrayOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesResult) []GetEdgeFwSwitchesData { return v.Datas }).(GetEdgeFwSwitchesDataArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEdgeFwSwitchesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetEdgeFwSwitchesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetEdgeFwSwitchesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEdgeFwSwitchesResultOutput{})
}
