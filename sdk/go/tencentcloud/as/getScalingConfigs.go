// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query scaling configuration information.
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
//			_, err := As.GetScalingConfigs(ctx, &as.GetScalingConfigsArgs{
//				ConfigurationId:  pulumi.StringRef("asc-oqio4yyj"),
//				ResultOutputFile: pulumi.StringRef("my_test_path"),
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
func GetScalingConfigs(ctx *pulumi.Context, args *GetScalingConfigsArgs, opts ...pulumi.InvokeOption) (*GetScalingConfigsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetScalingConfigsResult
	err := ctx.Invoke("tencentcloud:As/getScalingConfigs:getScalingConfigs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScalingConfigs.
type GetScalingConfigsArgs struct {
	// Launch configuration ID.
	ConfigurationId *string `pulumi:"configurationId"`
	// Launch configuration name.
	ConfigurationName *string `pulumi:"configurationName"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getScalingConfigs.
type GetScalingConfigsResult struct {
	// Launch configuration ID.
	ConfigurationId *string `pulumi:"configurationId"`
	// A list of configuration. Each element contains the following attributes:
	ConfigurationLists []GetScalingConfigsConfigurationList `pulumi:"configurationLists"`
	// Launch configuration name.
	ConfigurationName *string `pulumi:"configurationName"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetScalingConfigsOutput(ctx *pulumi.Context, args GetScalingConfigsOutputArgs, opts ...pulumi.InvokeOption) GetScalingConfigsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScalingConfigsResult, error) {
			args := v.(GetScalingConfigsArgs)
			r, err := GetScalingConfigs(ctx, &args, opts...)
			var s GetScalingConfigsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScalingConfigsResultOutput)
}

// A collection of arguments for invoking getScalingConfigs.
type GetScalingConfigsOutputArgs struct {
	// Launch configuration ID.
	ConfigurationId pulumi.StringPtrInput `pulumi:"configurationId"`
	// Launch configuration name.
	ConfigurationName pulumi.StringPtrInput `pulumi:"configurationName"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetScalingConfigsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingConfigsArgs)(nil)).Elem()
}

// A collection of values returned by getScalingConfigs.
type GetScalingConfigsResultOutput struct{ *pulumi.OutputState }

func (GetScalingConfigsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingConfigsResult)(nil)).Elem()
}

func (o GetScalingConfigsResultOutput) ToGetScalingConfigsResultOutput() GetScalingConfigsResultOutput {
	return o
}

func (o GetScalingConfigsResultOutput) ToGetScalingConfigsResultOutputWithContext(ctx context.Context) GetScalingConfigsResultOutput {
	return o
}

// Launch configuration ID.
func (o GetScalingConfigsResultOutput) ConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingConfigsResult) *string { return v.ConfigurationId }).(pulumi.StringPtrOutput)
}

// A list of configuration. Each element contains the following attributes:
func (o GetScalingConfigsResultOutput) ConfigurationLists() GetScalingConfigsConfigurationListArrayOutput {
	return o.ApplyT(func(v GetScalingConfigsResult) []GetScalingConfigsConfigurationList { return v.ConfigurationLists }).(GetScalingConfigsConfigurationListArrayOutput)
}

// Launch configuration name.
func (o GetScalingConfigsResultOutput) ConfigurationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingConfigsResult) *string { return v.ConfigurationName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetScalingConfigsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScalingConfigsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetScalingConfigsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingConfigsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScalingConfigsResultOutput{})
}
