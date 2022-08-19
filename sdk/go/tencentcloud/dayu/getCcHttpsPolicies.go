// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query dayu CC https policies
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dayu.GetCcHttpsPolicies(ctx, &dayu.GetCcHttpsPoliciesArgs{
// 			ResourceType: tencentcloud_dayu_cc_https_policy.Test_policy.Resource_type,
// 			ResourceId:   tencentcloud_dayu_cc_https_policy.Test_policy.Resource_id,
// 			Name:         pulumi.StringRef(tencentcloud_dayu_cc_https_policy.Test_policy.Name),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Dayu.GetCcHttpsPolicies(ctx, &dayu.GetCcHttpsPoliciesArgs{
// 			ResourceType: tencentcloud_dayu_cc_https_policy.Test_policy.Resource_type,
// 			ResourceId:   tencentcloud_dayu_cc_https_policy.Test_policy.Resource_id,
// 			PolicyId:     pulumi.StringRef(tencentcloud_dayu_cc_https_policy.Test_policy.Policy_id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCcHttpsPolicies(ctx *pulumi.Context, args *GetCcHttpsPoliciesArgs, opts ...pulumi.InvokeOption) (*GetCcHttpsPoliciesResult, error) {
	var rv GetCcHttpsPoliciesResult
	err := ctx.Invoke("tencentcloud:Dayu/getCcHttpsPolicies:getCcHttpsPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCcHttpsPolicies.
type GetCcHttpsPoliciesArgs struct {
	// Name of the CC https policy to be queried.
	Name *string `pulumi:"name"`
	// Id of the CC https policy to be queried.
	PolicyId *string `pulumi:"policyId"`
	// Id of the resource that the CC https policy works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the CC https policy works for, valid value is `bgpip`.
	ResourceType string `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCcHttpsPolicies.
type GetCcHttpsPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CC https policies. Each element contains the following attributes:
	Lists []GetCcHttpsPoliciesList `pulumi:"lists"`
	// Name of the CC self-define https policy.
	Name *string `pulumi:"name"`
	// Id of the CC self-define https policy.
	PolicyId *string `pulumi:"policyId"`
	// ID of the resource that the CC self-define https policy works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the CC self-define https policy works for.
	ResourceType     string  `pulumi:"resourceType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetCcHttpsPoliciesOutput(ctx *pulumi.Context, args GetCcHttpsPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetCcHttpsPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCcHttpsPoliciesResult, error) {
			args := v.(GetCcHttpsPoliciesArgs)
			r, err := GetCcHttpsPolicies(ctx, &args, opts...)
			var s GetCcHttpsPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCcHttpsPoliciesResultOutput)
}

// A collection of arguments for invoking getCcHttpsPolicies.
type GetCcHttpsPoliciesOutputArgs struct {
	// Name of the CC https policy to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Id of the CC https policy to be queried.
	PolicyId pulumi.StringPtrInput `pulumi:"policyId"`
	// Id of the resource that the CC https policy works for.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Type of the resource that the CC https policy works for, valid value is `bgpip`.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCcHttpsPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCcHttpsPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getCcHttpsPolicies.
type GetCcHttpsPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetCcHttpsPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCcHttpsPoliciesResult)(nil)).Elem()
}

func (o GetCcHttpsPoliciesResultOutput) ToGetCcHttpsPoliciesResultOutput() GetCcHttpsPoliciesResultOutput {
	return o
}

func (o GetCcHttpsPoliciesResultOutput) ToGetCcHttpsPoliciesResultOutputWithContext(ctx context.Context) GetCcHttpsPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCcHttpsPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcHttpsPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of CC https policies. Each element contains the following attributes:
func (o GetCcHttpsPoliciesResultOutput) Lists() GetCcHttpsPoliciesListArrayOutput {
	return o.ApplyT(func(v GetCcHttpsPoliciesResult) []GetCcHttpsPoliciesList { return v.Lists }).(GetCcHttpsPoliciesListArrayOutput)
}

// Name of the CC self-define https policy.
func (o GetCcHttpsPoliciesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcHttpsPoliciesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Id of the CC self-define https policy.
func (o GetCcHttpsPoliciesResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcHttpsPoliciesResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// ID of the resource that the CC self-define https policy works for.
func (o GetCcHttpsPoliciesResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcHttpsPoliciesResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// Type of the resource that the CC self-define https policy works for.
func (o GetCcHttpsPoliciesResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcHttpsPoliciesResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetCcHttpsPoliciesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcHttpsPoliciesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCcHttpsPoliciesResultOutput{})
}
