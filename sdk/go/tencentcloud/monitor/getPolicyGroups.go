// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query monitor policy groups (There is a lot of data and it is recommended to output to a file)
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
//			_, err := Monitor.GetPolicyGroups(ctx, &monitor.GetPolicyGroupsArgs{
//				PolicyViewNames: []string{
//					"REDIS-CLUSTER",
//					"cvm_device",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Monitor.GetPolicyGroups(ctx, &monitor.GetPolicyGroupsArgs{
//				Name: pulumi.StringRef("test"),
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
func GetPolicyGroups(ctx *pulumi.Context, args *GetPolicyGroupsArgs, opts ...pulumi.InvokeOption) (*GetPolicyGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetPolicyGroupsResult
	err := ctx.Invoke("tencentcloud:Monitor/getPolicyGroups:getPolicyGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicyGroups.
type GetPolicyGroupsArgs struct {
	// Policy group name for query.
	Name *string `pulumi:"name"`
	// The policy view for query.
	PolicyViewNames []string `pulumi:"policyViewNames"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getPolicyGroups.
type GetPolicyGroupsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list policy groups. Each element contains the following attributes:
	Lists            []GetPolicyGroupsList `pulumi:"lists"`
	Name             *string               `pulumi:"name"`
	PolicyViewNames  []string              `pulumi:"policyViewNames"`
	ResultOutputFile *string               `pulumi:"resultOutputFile"`
}

func GetPolicyGroupsOutput(ctx *pulumi.Context, args GetPolicyGroupsOutputArgs, opts ...pulumi.InvokeOption) GetPolicyGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPolicyGroupsResult, error) {
			args := v.(GetPolicyGroupsArgs)
			r, err := GetPolicyGroups(ctx, &args, opts...)
			var s GetPolicyGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPolicyGroupsResultOutput)
}

// A collection of arguments for invoking getPolicyGroups.
type GetPolicyGroupsOutputArgs struct {
	// Policy group name for query.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The policy view for query.
	PolicyViewNames pulumi.StringArrayInput `pulumi:"policyViewNames"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetPolicyGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyGroupsArgs)(nil)).Elem()
}

// A collection of values returned by getPolicyGroups.
type GetPolicyGroupsResultOutput struct{ *pulumi.OutputState }

func (GetPolicyGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPolicyGroupsResult)(nil)).Elem()
}

func (o GetPolicyGroupsResultOutput) ToGetPolicyGroupsResultOutput() GetPolicyGroupsResultOutput {
	return o
}

func (o GetPolicyGroupsResultOutput) ToGetPolicyGroupsResultOutputWithContext(ctx context.Context) GetPolicyGroupsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPolicyGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list policy groups. Each element contains the following attributes:
func (o GetPolicyGroupsResultOutput) Lists() GetPolicyGroupsListArrayOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) []GetPolicyGroupsList { return v.Lists }).(GetPolicyGroupsListArrayOutput)
}

func (o GetPolicyGroupsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetPolicyGroupsResultOutput) PolicyViewNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) []string { return v.PolicyViewNames }).(pulumi.StringArrayOutput)
}

func (o GetPolicyGroupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPolicyGroupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPolicyGroupsResultOutput{})
}
