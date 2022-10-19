// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query security policy rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooProxy, err := Gaap.NewProxy(ctx, "fooProxy", &Gaap.ProxyArgs{
//				Bandwidth:        pulumi.Int(10),
//				Concurrent:       pulumi.Int(2),
//				AccessRegion:     pulumi.String("SouthChina"),
//				RealserverRegion: pulumi.String("NorthChina"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSecurityPolicy, err := Gaap.NewSecurityPolicy(ctx, "fooSecurityPolicy", &Gaap.SecurityPolicyArgs{
//				ProxyId: fooProxy.ID(),
//				Action:  pulumi.String("ACCEPT"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSecurityRule, err := Gaap.NewSecurityRule(ctx, "fooSecurityRule", &Gaap.SecurityRuleArgs{
//				PolicyId: fooSecurityPolicy.ID(),
//				CidrIp:   pulumi.String("1.1.1.1"),
//				Action:   pulumi.String("ACCEPT"),
//				Protocol: pulumi.String("TCP"),
//				Port:     pulumi.String("80"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSecurityRules(ctx *pulumi.Context, args *GetSecurityRulesArgs, opts ...pulumi.InvokeOption) (*GetSecurityRulesResult, error) {
	var rv GetSecurityRulesResult
	err := ctx.Invoke("tencentcloud:Gaap/getSecurityRules:getSecurityRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityRules.
type GetSecurityRulesArgs struct {
	// Policy of the rule to be queried.
	Action *string `pulumi:"action"`
	// A network address block of the request source to be queried.
	CidrIp *string `pulumi:"cidrIp"`
	// Name of the security policy rule to be queried.
	Name *string `pulumi:"name"`
	// ID of the security policy to be queried.
	PolicyId string `pulumi:"policyId"`
	// Port of the security policy rule to be queried.
	Port *string `pulumi:"port"`
	// Protocol of the security policy rule to be queried.
	Protocol *string `pulumi:"protocol"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// ID of the security policy rules to be queried.
	RuleId *string `pulumi:"ruleId"`
}

// A collection of values returned by getSecurityRules.
type GetSecurityRulesResult struct {
	// Policy of the rule.
	Action *string `pulumi:"action"`
	// A network address block of the request source.
	CidrIp *string `pulumi:"cidrIp"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the security policy rule.
	Name     *string `pulumi:"name"`
	PolicyId string  `pulumi:"policyId"`
	// Port of the security policy rule.
	Port *string `pulumi:"port"`
	// Protocol of the security policy rule.
	Protocol         *string `pulumi:"protocol"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	RuleId           *string `pulumi:"ruleId"`
	// An information list of security policy rule. Each element contains the following attributes:
	Rules []GetSecurityRulesRule `pulumi:"rules"`
}

func GetSecurityRulesOutput(ctx *pulumi.Context, args GetSecurityRulesOutputArgs, opts ...pulumi.InvokeOption) GetSecurityRulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecurityRulesResult, error) {
			args := v.(GetSecurityRulesArgs)
			r, err := GetSecurityRules(ctx, &args, opts...)
			var s GetSecurityRulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecurityRulesResultOutput)
}

// A collection of arguments for invoking getSecurityRules.
type GetSecurityRulesOutputArgs struct {
	// Policy of the rule to be queried.
	Action pulumi.StringPtrInput `pulumi:"action"`
	// A network address block of the request source to be queried.
	CidrIp pulumi.StringPtrInput `pulumi:"cidrIp"`
	// Name of the security policy rule to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// ID of the security policy to be queried.
	PolicyId pulumi.StringInput `pulumi:"policyId"`
	// Port of the security policy rule to be queried.
	Port pulumi.StringPtrInput `pulumi:"port"`
	// Protocol of the security policy rule to be queried.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// ID of the security policy rules to be queried.
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
}

func (GetSecurityRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityRulesArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityRules.
type GetSecurityRulesResultOutput struct{ *pulumi.OutputState }

func (GetSecurityRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityRulesResult)(nil)).Elem()
}

func (o GetSecurityRulesResultOutput) ToGetSecurityRulesResultOutput() GetSecurityRulesResultOutput {
	return o
}

func (o GetSecurityRulesResultOutput) ToGetSecurityRulesResultOutputWithContext(ctx context.Context) GetSecurityRulesResultOutput {
	return o
}

// Policy of the rule.
func (o GetSecurityRulesResultOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) *string { return v.Action }).(pulumi.StringPtrOutput)
}

// A network address block of the request source.
func (o GetSecurityRulesResultOutput) CidrIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) *string { return v.CidrIp }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecurityRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the security policy rule.
func (o GetSecurityRulesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSecurityRulesResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

// Port of the security policy rule.
func (o GetSecurityRulesResultOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) *string { return v.Port }).(pulumi.StringPtrOutput)
}

// Protocol of the security policy rule.
func (o GetSecurityRulesResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o GetSecurityRulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSecurityRulesResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

// An information list of security policy rule. Each element contains the following attributes:
func (o GetSecurityRulesResultOutput) Rules() GetSecurityRulesRuleArrayOutput {
	return o.ApplyT(func(v GetSecurityRulesResult) []GetSecurityRulesRule { return v.Rules }).(GetSecurityRulesRuleArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecurityRulesResultOutput{})
}
