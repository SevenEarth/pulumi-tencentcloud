// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of monitor grafanaPluginOverviews
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Monitor.GetGrafanaPluginOverviews(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetGrafanaPluginOverviews(ctx *pulumi.Context, args *GetGrafanaPluginOverviewsArgs, opts ...pulumi.InvokeOption) (*GetGrafanaPluginOverviewsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGrafanaPluginOverviewsResult
	err := ctx.Invoke("tencentcloud:Monitor/getGrafanaPluginOverviews:getGrafanaPluginOverviews", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGrafanaPluginOverviews.
type GetGrafanaPluginOverviewsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getGrafanaPluginOverviews.
type GetGrafanaPluginOverviewsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Plugin set.
	PluginSets       []GetGrafanaPluginOverviewsPluginSet `pulumi:"pluginSets"`
	ResultOutputFile *string                              `pulumi:"resultOutputFile"`
}

func GetGrafanaPluginOverviewsOutput(ctx *pulumi.Context, args GetGrafanaPluginOverviewsOutputArgs, opts ...pulumi.InvokeOption) GetGrafanaPluginOverviewsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGrafanaPluginOverviewsResult, error) {
			args := v.(GetGrafanaPluginOverviewsArgs)
			r, err := GetGrafanaPluginOverviews(ctx, &args, opts...)
			var s GetGrafanaPluginOverviewsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGrafanaPluginOverviewsResultOutput)
}

// A collection of arguments for invoking getGrafanaPluginOverviews.
type GetGrafanaPluginOverviewsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetGrafanaPluginOverviewsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGrafanaPluginOverviewsArgs)(nil)).Elem()
}

// A collection of values returned by getGrafanaPluginOverviews.
type GetGrafanaPluginOverviewsResultOutput struct{ *pulumi.OutputState }

func (GetGrafanaPluginOverviewsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGrafanaPluginOverviewsResult)(nil)).Elem()
}

func (o GetGrafanaPluginOverviewsResultOutput) ToGetGrafanaPluginOverviewsResultOutput() GetGrafanaPluginOverviewsResultOutput {
	return o
}

func (o GetGrafanaPluginOverviewsResultOutput) ToGetGrafanaPluginOverviewsResultOutputWithContext(ctx context.Context) GetGrafanaPluginOverviewsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetGrafanaPluginOverviewsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGrafanaPluginOverviewsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Plugin set.
func (o GetGrafanaPluginOverviewsResultOutput) PluginSets() GetGrafanaPluginOverviewsPluginSetArrayOutput {
	return o.ApplyT(func(v GetGrafanaPluginOverviewsResult) []GetGrafanaPluginOverviewsPluginSet { return v.PluginSets }).(GetGrafanaPluginOverviewsPluginSetArrayOutput)
}

func (o GetGrafanaPluginOverviewsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGrafanaPluginOverviewsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGrafanaPluginOverviewsResultOutput{})
}
