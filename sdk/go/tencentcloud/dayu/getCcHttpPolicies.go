// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query dayu CC http policies
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dayu.GetCcHttpPolicies(ctx, &dayu.GetCcHttpPoliciesArgs{
//				ResourceType: tencentcloud_dayu_cc_http_policy.Test_policy.Resource_type,
//				ResourceId:   tencentcloud_dayu_cc_http_policy.Test_policy.Resource_id,
//				PolicyId:     pulumi.StringRef(tencentcloud_dayu_cc_http_policy.Test_policy.Policy_id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Dayu.GetCcHttpPolicies(ctx, &dayu.GetCcHttpPoliciesArgs{
//				ResourceType: tencentcloud_dayu_cc_http_policy.Test_policy.Resource_type,
//				ResourceId:   tencentcloud_dayu_cc_http_policy.Test_policy.Resource_id,
//				Name:         pulumi.StringRef(tencentcloud_dayu_cc_http_policy.Test_policy.Name),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCcHttpPolicies(ctx *pulumi.Context, args *GetCcHttpPoliciesArgs, opts ...pulumi.InvokeOption) (*GetCcHttpPoliciesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCcHttpPoliciesResult
	err := ctx.Invoke("tencentcloud:Dayu/getCcHttpPolicies:getCcHttpPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCcHttpPolicies.
type GetCcHttpPoliciesArgs struct {
	// Name of the CC http policy to be queried.
	Name *string `pulumi:"name"`
	// Id of the CC http policy to be queried.
	PolicyId *string `pulumi:"policyId"`
	// ID of the resource that the CC http policy works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the CC http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType string `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCcHttpPolicies.
type GetCcHttpPoliciesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of CC http policies. Each element contains the following attributes:
	Lists []GetCcHttpPoliciesList `pulumi:"lists"`
	// Name of the CC self-define http policy.
	Name *string `pulumi:"name"`
	// ID of the CC self-define http policy.
	PolicyId *string `pulumi:"policyId"`
	// ID of the resource that the CC self-define http policy works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the CC self-define http policy works for.
	ResourceType     string  `pulumi:"resourceType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetCcHttpPoliciesOutput(ctx *pulumi.Context, args GetCcHttpPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetCcHttpPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCcHttpPoliciesResult, error) {
			args := v.(GetCcHttpPoliciesArgs)
			r, err := GetCcHttpPolicies(ctx, &args, opts...)
			var s GetCcHttpPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCcHttpPoliciesResultOutput)
}

// A collection of arguments for invoking getCcHttpPolicies.
type GetCcHttpPoliciesOutputArgs struct {
	// Name of the CC http policy to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Id of the CC http policy to be queried.
	PolicyId pulumi.StringPtrInput `pulumi:"policyId"`
	// ID of the resource that the CC http policy works for.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Type of the resource that the CC http policy works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCcHttpPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCcHttpPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getCcHttpPolicies.
type GetCcHttpPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetCcHttpPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCcHttpPoliciesResult)(nil)).Elem()
}

func (o GetCcHttpPoliciesResultOutput) ToGetCcHttpPoliciesResultOutput() GetCcHttpPoliciesResultOutput {
	return o
}

func (o GetCcHttpPoliciesResultOutput) ToGetCcHttpPoliciesResultOutputWithContext(ctx context.Context) GetCcHttpPoliciesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCcHttpPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcHttpPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of CC http policies. Each element contains the following attributes:
func (o GetCcHttpPoliciesResultOutput) Lists() GetCcHttpPoliciesListArrayOutput {
	return o.ApplyT(func(v GetCcHttpPoliciesResult) []GetCcHttpPoliciesList { return v.Lists }).(GetCcHttpPoliciesListArrayOutput)
}

// Name of the CC self-define http policy.
func (o GetCcHttpPoliciesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcHttpPoliciesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// ID of the CC self-define http policy.
func (o GetCcHttpPoliciesResultOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcHttpPoliciesResult) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

// ID of the resource that the CC self-define http policy works for.
func (o GetCcHttpPoliciesResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcHttpPoliciesResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

// Type of the resource that the CC self-define http policy works for.
func (o GetCcHttpPoliciesResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetCcHttpPoliciesResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetCcHttpPoliciesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCcHttpPoliciesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCcHttpPoliciesResultOutput{})
}
