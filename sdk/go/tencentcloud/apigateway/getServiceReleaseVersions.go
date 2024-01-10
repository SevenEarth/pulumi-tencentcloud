// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of apiGateway serviceReleaseVersions
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ApiGateway.GetServiceReleaseVersions(ctx, &apigateway.GetServiceReleaseVersionsArgs{
//				ServiceId: "service-nxz6yync",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupServiceReleaseVersions(ctx *pulumi.Context, args *LookupServiceReleaseVersionsArgs, opts ...pulumi.InvokeOption) (*LookupServiceReleaseVersionsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupServiceReleaseVersionsResult
	err := ctx.Invoke("tencentcloud:ApiGateway/getServiceReleaseVersions:getServiceReleaseVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceReleaseVersions.
type LookupServiceReleaseVersionsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The unique ID of the service to be queried.
	ServiceId string `pulumi:"serviceId"`
}

// A collection of values returned by getServiceReleaseVersions.
type LookupServiceReleaseVersionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// List of service releases.Note: This field may return null, indicating that no valid value can be obtained.
	Results   []GetServiceReleaseVersionsResult `pulumi:"results"`
	ServiceId string                            `pulumi:"serviceId"`
}

func LookupServiceReleaseVersionsOutput(ctx *pulumi.Context, args LookupServiceReleaseVersionsOutputArgs, opts ...pulumi.InvokeOption) LookupServiceReleaseVersionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceReleaseVersionsResult, error) {
			args := v.(LookupServiceReleaseVersionsArgs)
			r, err := LookupServiceReleaseVersions(ctx, &args, opts...)
			var s LookupServiceReleaseVersionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceReleaseVersionsResultOutput)
}

// A collection of arguments for invoking getServiceReleaseVersions.
type LookupServiceReleaseVersionsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The unique ID of the service to be queried.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (LookupServiceReleaseVersionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceReleaseVersionsArgs)(nil)).Elem()
}

// A collection of values returned by getServiceReleaseVersions.
type LookupServiceReleaseVersionsResultOutput struct{ *pulumi.OutputState }

func (LookupServiceReleaseVersionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceReleaseVersionsResult)(nil)).Elem()
}

func (o LookupServiceReleaseVersionsResultOutput) ToLookupServiceReleaseVersionsResultOutput() LookupServiceReleaseVersionsResultOutput {
	return o
}

func (o LookupServiceReleaseVersionsResultOutput) ToLookupServiceReleaseVersionsResultOutputWithContext(ctx context.Context) LookupServiceReleaseVersionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupServiceReleaseVersionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceReleaseVersionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceReleaseVersionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceReleaseVersionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// List of service releases.Note: This field may return null, indicating that no valid value can be obtained.
func (o LookupServiceReleaseVersionsResultOutput) Results() GetServiceReleaseVersionsResultArrayOutput {
	return o.ApplyT(func(v LookupServiceReleaseVersionsResult) []GetServiceReleaseVersionsResult { return v.Results }).(GetServiceReleaseVersionsResultArrayOutput)
}

func (o LookupServiceReleaseVersionsResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceReleaseVersionsResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceReleaseVersionsResultOutput{})
}
