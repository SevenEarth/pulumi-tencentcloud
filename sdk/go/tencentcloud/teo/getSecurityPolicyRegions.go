// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of teo securityPolicyRegions
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Teo.GetSecurityPolicyRegions(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSecurityPolicyRegions(ctx *pulumi.Context, args *GetSecurityPolicyRegionsArgs, opts ...pulumi.InvokeOption) (*GetSecurityPolicyRegionsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecurityPolicyRegionsResult
	err := ctx.Invoke("tencentcloud:Teo/getSecurityPolicyRegions:getSecurityPolicyRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityPolicyRegions.
type GetSecurityPolicyRegionsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getSecurityPolicyRegions.
type GetSecurityPolicyRegionsResult struct {
	// Region info.
	GeoIps []GetSecurityPolicyRegionsGeoIp `pulumi:"geoIps"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetSecurityPolicyRegionsOutput(ctx *pulumi.Context, args GetSecurityPolicyRegionsOutputArgs, opts ...pulumi.InvokeOption) GetSecurityPolicyRegionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecurityPolicyRegionsResult, error) {
			args := v.(GetSecurityPolicyRegionsArgs)
			r, err := GetSecurityPolicyRegions(ctx, &args, opts...)
			var s GetSecurityPolicyRegionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecurityPolicyRegionsResultOutput)
}

// A collection of arguments for invoking getSecurityPolicyRegions.
type GetSecurityPolicyRegionsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetSecurityPolicyRegionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityPolicyRegionsArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityPolicyRegions.
type GetSecurityPolicyRegionsResultOutput struct{ *pulumi.OutputState }

func (GetSecurityPolicyRegionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityPolicyRegionsResult)(nil)).Elem()
}

func (o GetSecurityPolicyRegionsResultOutput) ToGetSecurityPolicyRegionsResultOutput() GetSecurityPolicyRegionsResultOutput {
	return o
}

func (o GetSecurityPolicyRegionsResultOutput) ToGetSecurityPolicyRegionsResultOutputWithContext(ctx context.Context) GetSecurityPolicyRegionsResultOutput {
	return o
}

// Region info.
func (o GetSecurityPolicyRegionsResultOutput) GeoIps() GetSecurityPolicyRegionsGeoIpArrayOutput {
	return o.ApplyT(func(v GetSecurityPolicyRegionsResult) []GetSecurityPolicyRegionsGeoIp { return v.GeoIps }).(GetSecurityPolicyRegionsGeoIpArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecurityPolicyRegionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityPolicyRegionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSecurityPolicyRegionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityPolicyRegionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecurityPolicyRegionsResultOutput{})
}
