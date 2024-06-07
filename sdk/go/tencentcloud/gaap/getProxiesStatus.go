// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of gaap proxies status
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Gaap.GetProxiesStatus(ctx, &gaap.GetProxiesStatusArgs{
//				ProxyIds: []string{
//					"link-xxxxxx",
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
// <!--End PulumiCodeChooser -->
func GetProxiesStatus(ctx *pulumi.Context, args *GetProxiesStatusArgs, opts ...pulumi.InvokeOption) (*GetProxiesStatusResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProxiesStatusResult
	err := ctx.Invoke("tencentcloud:Gaap/getProxiesStatus:getProxiesStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProxiesStatus.
type GetProxiesStatusArgs struct {
	// List of Proxy IDs.
	ProxyIds []string `pulumi:"proxyIds"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getProxiesStatus.
type GetProxiesStatusResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Proxy status list.
	InstanceStatusSets []GetProxiesStatusInstanceStatusSet `pulumi:"instanceStatusSets"`
	ProxyIds           []string                            `pulumi:"proxyIds"`
	ResultOutputFile   *string                             `pulumi:"resultOutputFile"`
}

func GetProxiesStatusOutput(ctx *pulumi.Context, args GetProxiesStatusOutputArgs, opts ...pulumi.InvokeOption) GetProxiesStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProxiesStatusResult, error) {
			args := v.(GetProxiesStatusArgs)
			r, err := GetProxiesStatus(ctx, &args, opts...)
			var s GetProxiesStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProxiesStatusResultOutput)
}

// A collection of arguments for invoking getProxiesStatus.
type GetProxiesStatusOutputArgs struct {
	// List of Proxy IDs.
	ProxyIds pulumi.StringArrayInput `pulumi:"proxyIds"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetProxiesStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxiesStatusArgs)(nil)).Elem()
}

// A collection of values returned by getProxiesStatus.
type GetProxiesStatusResultOutput struct{ *pulumi.OutputState }

func (GetProxiesStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProxiesStatusResult)(nil)).Elem()
}

func (o GetProxiesStatusResultOutput) ToGetProxiesStatusResultOutput() GetProxiesStatusResultOutput {
	return o
}

func (o GetProxiesStatusResultOutput) ToGetProxiesStatusResultOutputWithContext(ctx context.Context) GetProxiesStatusResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetProxiesStatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProxiesStatusResult) string { return v.Id }).(pulumi.StringOutput)
}

// Proxy status list.
func (o GetProxiesStatusResultOutput) InstanceStatusSets() GetProxiesStatusInstanceStatusSetArrayOutput {
	return o.ApplyT(func(v GetProxiesStatusResult) []GetProxiesStatusInstanceStatusSet { return v.InstanceStatusSets }).(GetProxiesStatusInstanceStatusSetArrayOutput)
}

func (o GetProxiesStatusResultOutput) ProxyIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProxiesStatusResult) []string { return v.ProxyIds }).(pulumi.StringArrayOutput)
}

func (o GetProxiesStatusResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProxiesStatusResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProxiesStatusResultOutput{})
}
