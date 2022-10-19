// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query dayu layer 7 rules
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
//			_, err := Dayu.GetL7Rules(ctx, &dayu.GetL7RulesArgs{
//				ResourceType: tencentcloud_dayu_l7_rule.Test_rule.Resource_type,
//				ResourceId:   tencentcloud_dayu_l7_rule.Test_rule.Resource_id,
//				Domain:       pulumi.StringRef(tencentcloud_dayu_l7_rule.Test_rule.Domain),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Dayu.GetL7Rules(ctx, &dayu.GetL7RulesArgs{
//				ResourceType: tencentcloud_dayu_l7_rule.Test_rule.Resource_type,
//				ResourceId:   tencentcloud_dayu_l7_rule.Test_rule.Resource_id,
//				RuleId:       pulumi.StringRef(tencentcloud_dayu_l7_rule.Test_rule.Rule_id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetL7Rules(ctx *pulumi.Context, args *GetL7RulesArgs, opts ...pulumi.InvokeOption) (*GetL7RulesResult, error) {
	var rv GetL7RulesResult
	err := ctx.Invoke("tencentcloud:Dayu/getL7Rules:getL7Rules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getL7Rules.
type GetL7RulesArgs struct {
	// Domain of the layer 7 rule to be queried.
	Domain *string `pulumi:"domain"`
	// Id of the resource that the layer 7 rule works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType string `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Id of the layer 7 rule to be queried.
	RuleId *string `pulumi:"ruleId"`
}

// A collection of values returned by getL7Rules.
type GetL7RulesResult struct {
	// Domain that the 7 layer rule works for.
	Domain *string `pulumi:"domain"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of layer 7 rules. Each element contains the following attributes:
	Lists            []GetL7RulesList `pulumi:"lists"`
	ResourceId       string           `pulumi:"resourceId"`
	ResourceType     string           `pulumi:"resourceType"`
	ResultOutputFile *string          `pulumi:"resultOutputFile"`
	// Id of the 7 layer rule.
	RuleId *string `pulumi:"ruleId"`
}

func GetL7RulesOutput(ctx *pulumi.Context, args GetL7RulesOutputArgs, opts ...pulumi.InvokeOption) GetL7RulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetL7RulesResult, error) {
			args := v.(GetL7RulesArgs)
			r, err := GetL7Rules(ctx, &args, opts...)
			var s GetL7RulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetL7RulesResultOutput)
}

// A collection of arguments for invoking getL7Rules.
type GetL7RulesOutputArgs struct {
	// Domain of the layer 7 rule to be queried.
	Domain pulumi.StringPtrInput `pulumi:"domain"`
	// Id of the resource that the layer 7 rule works for.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Id of the layer 7 rule to be queried.
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
}

func (GetL7RulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetL7RulesArgs)(nil)).Elem()
}

// A collection of values returned by getL7Rules.
type GetL7RulesResultOutput struct{ *pulumi.OutputState }

func (GetL7RulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetL7RulesResult)(nil)).Elem()
}

func (o GetL7RulesResultOutput) ToGetL7RulesResultOutput() GetL7RulesResultOutput {
	return o
}

func (o GetL7RulesResultOutput) ToGetL7RulesResultOutputWithContext(ctx context.Context) GetL7RulesResultOutput {
	return o
}

// Domain that the 7 layer rule works for.
func (o GetL7RulesResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetL7RulesResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetL7RulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetL7RulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of layer 7 rules. Each element contains the following attributes:
func (o GetL7RulesResultOutput) Lists() GetL7RulesListArrayOutput {
	return o.ApplyT(func(v GetL7RulesResult) []GetL7RulesList { return v.Lists }).(GetL7RulesListArrayOutput)
}

func (o GetL7RulesResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetL7RulesResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o GetL7RulesResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetL7RulesResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetL7RulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetL7RulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Id of the 7 layer rule.
func (o GetL7RulesResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetL7RulesResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetL7RulesResultOutput{})
}
