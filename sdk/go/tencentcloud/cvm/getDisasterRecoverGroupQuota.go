// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cvm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cvm disasterRecoverGroupQuota
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cvm.GetDisasterRecoverGroupQuota(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDisasterRecoverGroupQuota(ctx *pulumi.Context, args *GetDisasterRecoverGroupQuotaArgs, opts ...pulumi.InvokeOption) (*GetDisasterRecoverGroupQuotaResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDisasterRecoverGroupQuotaResult
	err := ctx.Invoke("tencentcloud:Cvm/getDisasterRecoverGroupQuota:getDisasterRecoverGroupQuota", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDisasterRecoverGroupQuota.
type GetDisasterRecoverGroupQuotaArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDisasterRecoverGroupQuota.
type GetDisasterRecoverGroupQuotaResult struct {
	// The number of placement groups that have been created by the current user.
	CurrentNum int `pulumi:"currentNum"`
	// Quota on instances in a physical-machine-type disaster recovery group.
	CvmInHostGroupQuota int `pulumi:"cvmInHostGroupQuota"`
	// Quota on instances in a rack-type disaster recovery group.
	CvmInRackGroupQuota int `pulumi:"cvmInRackGroupQuota"`
	// Quota on instances in a switch-type disaster recovery group.
	CvmInSwGroupQuota int `pulumi:"cvmInSwGroupQuota"`
	// The maximum number of placement groups that can be created.
	GroupQuota int `pulumi:"groupQuota"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDisasterRecoverGroupQuotaOutput(ctx *pulumi.Context, args GetDisasterRecoverGroupQuotaOutputArgs, opts ...pulumi.InvokeOption) GetDisasterRecoverGroupQuotaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDisasterRecoverGroupQuotaResult, error) {
			args := v.(GetDisasterRecoverGroupQuotaArgs)
			r, err := GetDisasterRecoverGroupQuota(ctx, &args, opts...)
			var s GetDisasterRecoverGroupQuotaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDisasterRecoverGroupQuotaResultOutput)
}

// A collection of arguments for invoking getDisasterRecoverGroupQuota.
type GetDisasterRecoverGroupQuotaOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDisasterRecoverGroupQuotaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDisasterRecoverGroupQuotaArgs)(nil)).Elem()
}

// A collection of values returned by getDisasterRecoverGroupQuota.
type GetDisasterRecoverGroupQuotaResultOutput struct{ *pulumi.OutputState }

func (GetDisasterRecoverGroupQuotaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDisasterRecoverGroupQuotaResult)(nil)).Elem()
}

func (o GetDisasterRecoverGroupQuotaResultOutput) ToGetDisasterRecoverGroupQuotaResultOutput() GetDisasterRecoverGroupQuotaResultOutput {
	return o
}

func (o GetDisasterRecoverGroupQuotaResultOutput) ToGetDisasterRecoverGroupQuotaResultOutputWithContext(ctx context.Context) GetDisasterRecoverGroupQuotaResultOutput {
	return o
}

// The number of placement groups that have been created by the current user.
func (o GetDisasterRecoverGroupQuotaResultOutput) CurrentNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetDisasterRecoverGroupQuotaResult) int { return v.CurrentNum }).(pulumi.IntOutput)
}

// Quota on instances in a physical-machine-type disaster recovery group.
func (o GetDisasterRecoverGroupQuotaResultOutput) CvmInHostGroupQuota() pulumi.IntOutput {
	return o.ApplyT(func(v GetDisasterRecoverGroupQuotaResult) int { return v.CvmInHostGroupQuota }).(pulumi.IntOutput)
}

// Quota on instances in a rack-type disaster recovery group.
func (o GetDisasterRecoverGroupQuotaResultOutput) CvmInRackGroupQuota() pulumi.IntOutput {
	return o.ApplyT(func(v GetDisasterRecoverGroupQuotaResult) int { return v.CvmInRackGroupQuota }).(pulumi.IntOutput)
}

// Quota on instances in a switch-type disaster recovery group.
func (o GetDisasterRecoverGroupQuotaResultOutput) CvmInSwGroupQuota() pulumi.IntOutput {
	return o.ApplyT(func(v GetDisasterRecoverGroupQuotaResult) int { return v.CvmInSwGroupQuota }).(pulumi.IntOutput)
}

// The maximum number of placement groups that can be created.
func (o GetDisasterRecoverGroupQuotaResultOutput) GroupQuota() pulumi.IntOutput {
	return o.ApplyT(func(v GetDisasterRecoverGroupQuotaResult) int { return v.GroupQuota }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDisasterRecoverGroupQuotaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDisasterRecoverGroupQuotaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDisasterRecoverGroupQuotaResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDisasterRecoverGroupQuotaResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDisasterRecoverGroupQuotaResultOutput{})
}
