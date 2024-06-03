// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of CBS snapshot policies.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cbs.GetSnapshotPolicies(ctx, &cbs.GetSnapshotPoliciesArgs{
//				SnapshotPolicyId:   pulumi.StringRef("snap-f3io7adt"),
//				SnapshotPolicyName: pulumi.StringRef("test"),
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
func GetSnapshotPolicies(ctx *pulumi.Context, args *GetSnapshotPoliciesArgs, opts ...pulumi.InvokeOption) (*GetSnapshotPoliciesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSnapshotPoliciesResult
	err := ctx.Invoke("tencentcloud:Cbs/getSnapshotPolicies:getSnapshotPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshotPolicies.
type GetSnapshotPoliciesArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the snapshot policy to be queried.
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
	// Name of the snapshot policy to be queried.
	SnapshotPolicyName *string `pulumi:"snapshotPolicyName"`
}

// A collection of values returned by getSnapshotPolicies.
type GetSnapshotPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the snapshot policy.
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
	// A list of snapshot policy. Each element contains the following attributes:
	SnapshotPolicyLists []GetSnapshotPoliciesSnapshotPolicyList `pulumi:"snapshotPolicyLists"`
	// Name of the snapshot policy.
	SnapshotPolicyName *string `pulumi:"snapshotPolicyName"`
}

func GetSnapshotPoliciesOutput(ctx *pulumi.Context, args GetSnapshotPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetSnapshotPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSnapshotPoliciesResult, error) {
			args := v.(GetSnapshotPoliciesArgs)
			r, err := GetSnapshotPolicies(ctx, &args, opts...)
			var s GetSnapshotPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSnapshotPoliciesResultOutput)
}

// A collection of arguments for invoking getSnapshotPolicies.
type GetSnapshotPoliciesOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of the snapshot policy to be queried.
	SnapshotPolicyId pulumi.StringPtrInput `pulumi:"snapshotPolicyId"`
	// Name of the snapshot policy to be queried.
	SnapshotPolicyName pulumi.StringPtrInput `pulumi:"snapshotPolicyName"`
}

func (GetSnapshotPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshotPolicies.
type GetSnapshotPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetSnapshotPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSnapshotPoliciesResult)(nil)).Elem()
}

func (o GetSnapshotPoliciesResultOutput) ToGetSnapshotPoliciesResultOutput() GetSnapshotPoliciesResultOutput {
	return o
}

func (o GetSnapshotPoliciesResultOutput) ToGetSnapshotPoliciesResultOutputWithContext(ctx context.Context) GetSnapshotPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSnapshotPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSnapshotPoliciesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of the snapshot policy.
func (o GetSnapshotPoliciesResultOutput) SnapshotPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesResult) *string { return v.SnapshotPolicyId }).(pulumi.StringPtrOutput)
}

// A list of snapshot policy. Each element contains the following attributes:
func (o GetSnapshotPoliciesResultOutput) SnapshotPolicyLists() GetSnapshotPoliciesSnapshotPolicyListArrayOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesResult) []GetSnapshotPoliciesSnapshotPolicyList {
		return v.SnapshotPolicyLists
	}).(GetSnapshotPoliciesSnapshotPolicyListArrayOutput)
}

// Name of the snapshot policy.
func (o GetSnapshotPoliciesResultOutput) SnapshotPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSnapshotPoliciesResult) *string { return v.SnapshotPolicyName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSnapshotPoliciesResultOutput{})
}
