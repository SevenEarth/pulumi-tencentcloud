// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ssm serviceStatus
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ssm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssm"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ssm.GetServiceStatus(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServiceStatus(ctx *pulumi.Context, args *GetServiceStatusArgs, opts ...pulumi.InvokeOption) (*GetServiceStatusResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetServiceStatusResult
	err := ctx.Invoke("tencentcloud:Ssm/getServiceStatus:getServiceStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceStatus.
type GetServiceStatusArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getServiceStatus.
type GetServiceStatusResult struct {
	// True means that the user can already use the key safe hosting function, false means that the user cannot use the key safe hosting function temporarily.
	AccessKeyEscrowEnabled bool `pulumi:"accessKeyEscrowEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Service unavailability type: 0-Not purchased, 1-Normal, 2-Service suspended due to arrears, 3-Resource release.
	InvalidType      int     `pulumi:"invalidType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// True means the service has been activated, false means the service has not been activated yet.
	ServiceEnabled bool `pulumi:"serviceEnabled"`
}

func GetServiceStatusOutput(ctx *pulumi.Context, args GetServiceStatusOutputArgs, opts ...pulumi.InvokeOption) GetServiceStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceStatusResult, error) {
			args := v.(GetServiceStatusArgs)
			r, err := GetServiceStatus(ctx, &args, opts...)
			var s GetServiceStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceStatusResultOutput)
}

// A collection of arguments for invoking getServiceStatus.
type GetServiceStatusOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetServiceStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceStatusArgs)(nil)).Elem()
}

// A collection of values returned by getServiceStatus.
type GetServiceStatusResultOutput struct{ *pulumi.OutputState }

func (GetServiceStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceStatusResult)(nil)).Elem()
}

func (o GetServiceStatusResultOutput) ToGetServiceStatusResultOutput() GetServiceStatusResultOutput {
	return o
}

func (o GetServiceStatusResultOutput) ToGetServiceStatusResultOutputWithContext(ctx context.Context) GetServiceStatusResultOutput {
	return o
}

// True means that the user can already use the key safe hosting function, false means that the user cannot use the key safe hosting function temporarily.
func (o GetServiceStatusResultOutput) AccessKeyEscrowEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceStatusResult) bool { return v.AccessKeyEscrowEnabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceStatusResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceStatusResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service unavailability type: 0-Not purchased, 1-Normal, 2-Service suspended due to arrears, 3-Resource release.
func (o GetServiceStatusResultOutput) InvalidType() pulumi.IntOutput {
	return o.ApplyT(func(v GetServiceStatusResult) int { return v.InvalidType }).(pulumi.IntOutput)
}

func (o GetServiceStatusResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceStatusResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// True means the service has been activated, false means the service has not been activated yet.
func (o GetServiceStatusResultOutput) ServiceEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceStatusResult) bool { return v.ServiceEnabled }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceStatusResultOutput{})
}
