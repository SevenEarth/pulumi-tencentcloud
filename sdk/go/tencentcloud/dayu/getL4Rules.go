// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query dayu layer 4 rules
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
// 		_, err := Dayu.GetL4Rules(ctx, &dayu.GetL4RulesArgs{
// 			ResourceType: tencentcloud_dayu_l4_rule.Test_rule.Resource_type,
// 			ResourceId:   tencentcloud_dayu_l4_rule.Test_rule.Resource_id,
// 			Name:         pulumi.StringRef(tencentcloud_dayu_l4_rule.Test_rule.Name),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Dayu.GetL4Rules(ctx, &dayu.GetL4RulesArgs{
// 			ResourceType: tencentcloud_dayu_l4_rule.Test_rule.Resource_type,
// 			ResourceId:   tencentcloud_dayu_l4_rule.Test_rule.Resource_id,
// 			RuleId:       pulumi.StringRef(tencentcloud_dayu_l4_rule.Test_rule.Rule_id),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetL4Rules(ctx *pulumi.Context, args *GetL4RulesArgs, opts ...pulumi.InvokeOption) (*GetL4RulesResult, error) {
	var rv GetL4RulesResult
	err := ctx.Invoke("tencentcloud:Dayu/getL4Rules:getL4Rules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getL4Rules.
type GetL4RulesArgs struct {
	// Name of the layer 4 rule to be queried.
	Name *string `pulumi:"name"`
	// Id of the resource that the layer 4 rule works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType string `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Id of the layer 4 rule to be queried.
	RuleId *string `pulumi:"ruleId"`
}

// A collection of values returned by getL4Rules.
type GetL4RulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of layer 4 rules. Each element contains the following attributes:
	Lists []GetL4RulesList `pulumi:"lists"`
	// Name of the rule.
	Name             *string `pulumi:"name"`
	ResourceId       string  `pulumi:"resourceId"`
	ResourceType     string  `pulumi:"resourceType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the 4 layer rule.
	RuleId *string `pulumi:"ruleId"`
}

func GetL4RulesOutput(ctx *pulumi.Context, args GetL4RulesOutputArgs, opts ...pulumi.InvokeOption) GetL4RulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetL4RulesResult, error) {
			args := v.(GetL4RulesArgs)
			r, err := GetL4Rules(ctx, &args, opts...)
			var s GetL4RulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetL4RulesResultOutput)
}

// A collection of arguments for invoking getL4Rules.
type GetL4RulesOutputArgs struct {
	// Name of the layer 4 rule to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Id of the resource that the layer 4 rule works for.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Type of the resource that the layer 4 rule works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Id of the layer 4 rule to be queried.
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
}

func (GetL4RulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetL4RulesArgs)(nil)).Elem()
}

// A collection of values returned by getL4Rules.
type GetL4RulesResultOutput struct{ *pulumi.OutputState }

func (GetL4RulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetL4RulesResult)(nil)).Elem()
}

func (o GetL4RulesResultOutput) ToGetL4RulesResultOutput() GetL4RulesResultOutput {
	return o
}

func (o GetL4RulesResultOutput) ToGetL4RulesResultOutputWithContext(ctx context.Context) GetL4RulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetL4RulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetL4RulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of layer 4 rules. Each element contains the following attributes:
func (o GetL4RulesResultOutput) Lists() GetL4RulesListArrayOutput {
	return o.ApplyT(func(v GetL4RulesResult) []GetL4RulesList { return v.Lists }).(GetL4RulesListArrayOutput)
}

// Name of the rule.
func (o GetL4RulesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetL4RulesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetL4RulesResultOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetL4RulesResult) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o GetL4RulesResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetL4RulesResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GetL4RulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetL4RulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of the 4 layer rule.
func (o GetL4RulesResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetL4RulesResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetL4RulesResultOutput{})
}
