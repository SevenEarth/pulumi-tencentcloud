// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package as

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of scaling policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/As"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/As"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := As.GetScalingPolicies(ctx, &as.GetScalingPoliciesArgs{
//				ResultOutputFile: pulumi.StringRef("mytestpath"),
//				ScalingPolicyId:  pulumi.StringRef("asg-mvyghxu7"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetScalingPolicies(ctx *pulumi.Context, args *GetScalingPoliciesArgs, opts ...pulumi.InvokeOption) (*GetScalingPoliciesResult, error) {
	var rv GetScalingPoliciesResult
	err := ctx.Invoke("tencentcloud:As/getScalingPolicies:getScalingPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScalingPolicies.
type GetScalingPoliciesArgs struct {
	// Scaling policy name.
	PolicyName *string `pulumi:"policyName"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Scaling group ID.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// Scaling policy ID.
	ScalingPolicyId *string `pulumi:"scalingPolicyId"`
}

// A collection of values returned by getScalingPolicies.
type GetScalingPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Scaling policy name.
	PolicyName       *string `pulumi:"policyName"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Scaling policy ID.
	ScalingGroupId  *string `pulumi:"scalingGroupId"`
	ScalingPolicyId *string `pulumi:"scalingPolicyId"`
	// A list of scaling policy. Each element contains the following attributes:
	ScalingPolicyLists []GetScalingPoliciesScalingPolicyList `pulumi:"scalingPolicyLists"`
}

func GetScalingPoliciesOutput(ctx *pulumi.Context, args GetScalingPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetScalingPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScalingPoliciesResult, error) {
			args := v.(GetScalingPoliciesArgs)
			r, err := GetScalingPolicies(ctx, &args, opts...)
			var s GetScalingPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScalingPoliciesResultOutput)
}

// A collection of arguments for invoking getScalingPolicies.
type GetScalingPoliciesOutputArgs struct {
	// Scaling policy name.
	PolicyName pulumi.StringPtrInput `pulumi:"policyName"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Scaling group ID.
	ScalingGroupId pulumi.StringPtrInput `pulumi:"scalingGroupId"`
	// Scaling policy ID.
	ScalingPolicyId pulumi.StringPtrInput `pulumi:"scalingPolicyId"`
}

func (GetScalingPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getScalingPolicies.
type GetScalingPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetScalingPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScalingPoliciesResult)(nil)).Elem()
}

func (o GetScalingPoliciesResultOutput) ToGetScalingPoliciesResultOutput() GetScalingPoliciesResultOutput {
	return o
}

func (o GetScalingPoliciesResultOutput) ToGetScalingPoliciesResultOutputWithContext(ctx context.Context) GetScalingPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetScalingPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScalingPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Scaling policy name.
func (o GetScalingPoliciesResultOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingPoliciesResult) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o GetScalingPoliciesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingPoliciesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Scaling policy ID.
func (o GetScalingPoliciesResultOutput) ScalingGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingPoliciesResult) *string { return v.ScalingGroupId }).(pulumi.StringPtrOutput)
}

func (o GetScalingPoliciesResultOutput) ScalingPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScalingPoliciesResult) *string { return v.ScalingPolicyId }).(pulumi.StringPtrOutput)
}

// A list of scaling policy. Each element contains the following attributes:
func (o GetScalingPoliciesResultOutput) ScalingPolicyLists() GetScalingPoliciesScalingPolicyListArrayOutput {
	return o.ApplyT(func(v GetScalingPoliciesResult) []GetScalingPoliciesScalingPolicyList { return v.ScalingPolicyLists }).(GetScalingPoliciesScalingPolicyListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScalingPoliciesResultOutput{})
}
