// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query the detail information of CFS access rule.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cfs.GetAccessRules(ctx, &cfs.GetAccessRulesArgs{
//				AccessGroupId: "pgroup-7nx89k7l",
//				AccessRuleId:  pulumi.StringRef("rule-qcndbqzj"),
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
func GetAccessRules(ctx *pulumi.Context, args *GetAccessRulesArgs, opts ...pulumi.InvokeOption) (*GetAccessRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAccessRulesResult
	err := ctx.Invoke("tencentcloud:Cfs/getAccessRules:getAccessRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessRules.
type GetAccessRulesArgs struct {
	// A specified access group ID used to query.
	AccessGroupId string `pulumi:"accessGroupId"`
	// A specified access rule ID used to query.
	AccessRuleId *string `pulumi:"accessRuleId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getAccessRules.
type GetAccessRulesResult struct {
	AccessGroupId string `pulumi:"accessGroupId"`
	// ID of the access rule.
	AccessRuleId *string `pulumi:"accessRuleId"`
	// An information list of CFS access rule. Each element contains the following attributes:
	AccessRuleLists []GetAccessRulesAccessRuleList `pulumi:"accessRuleLists"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetAccessRulesOutput(ctx *pulumi.Context, args GetAccessRulesOutputArgs, opts ...pulumi.InvokeOption) GetAccessRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessRulesResult, error) {
			args := v.(GetAccessRulesArgs)
			r, err := GetAccessRules(ctx, &args, opts...)
			var s GetAccessRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAccessRulesResultOutput)
}

// A collection of arguments for invoking getAccessRules.
type GetAccessRulesOutputArgs struct {
	// A specified access group ID used to query.
	AccessGroupId pulumi.StringInput `pulumi:"accessGroupId"`
	// A specified access rule ID used to query.
	AccessRuleId pulumi.StringPtrInput `pulumi:"accessRuleId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetAccessRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessRulesArgs)(nil)).Elem()
}

// A collection of values returned by getAccessRules.
type GetAccessRulesResultOutput struct{ *pulumi.OutputState }

func (GetAccessRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessRulesResult)(nil)).Elem()
}

func (o GetAccessRulesResultOutput) ToGetAccessRulesResultOutput() GetAccessRulesResultOutput {
	return o
}

func (o GetAccessRulesResultOutput) ToGetAccessRulesResultOutputWithContext(ctx context.Context) GetAccessRulesResultOutput {
	return o
}

func (o GetAccessRulesResultOutput) AccessGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessRulesResult) string { return v.AccessGroupId }).(pulumi.StringOutput)
}

// ID of the access rule.
func (o GetAccessRulesResultOutput) AccessRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessRulesResult) *string { return v.AccessRuleId }).(pulumi.StringPtrOutput)
}

// An information list of CFS access rule. Each element contains the following attributes:
func (o GetAccessRulesResultOutput) AccessRuleLists() GetAccessRulesAccessRuleListArrayOutput {
	return o.ApplyT(func(v GetAccessRulesResult) []GetAccessRulesAccessRuleList { return v.AccessRuleLists }).(GetAccessRulesAccessRuleListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAccessRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAccessRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessRulesResultOutput{})
}
