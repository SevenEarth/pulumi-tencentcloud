// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a teo ruleEngine
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Teo.NewRuleEngine(ctx, "rule1", &Teo.RuleEngineArgs{
// 			ZoneId:   pulumi.Any(tencentcloud_teo_zone.Example.Id),
// 			RuleName: pulumi.String("test-rule"),
// 			Status:   pulumi.String("disable"),
// 			Rules: teo.RuleEngineRuleArray{
// 				&teo.RuleEngineRuleArgs{
// 					Actions: teo.RuleEngineRuleActionArray{
// 						&teo.RuleEngineRuleActionArgs{
// 							NormalAction: &teo.RuleEngineRuleActionNormalActionArgs{
// 								Action: pulumi.String("UpstreamUrlRedirect"),
// 								Parameters: teo.RuleEngineRuleActionNormalActionParameterArray{
// 									&teo.RuleEngineRuleActionNormalActionParameterArgs{
// 										Name: pulumi.String("Type"),
// 										Values: pulumi.StringArray{
// 											pulumi.String("Path"),
// 										},
// 									},
// 									&teo.RuleEngineRuleActionNormalActionParameterArgs{
// 										Name: pulumi.String("Action"),
// 										Values: pulumi.StringArray{
// 											pulumi.String("addPrefix"),
// 										},
// 									},
// 									&teo.RuleEngineRuleActionNormalActionParameterArgs{
// 										Name: pulumi.String("Value"),
// 										Values: pulumi.StringArray{
// 											pulumi.String("/sss"),
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 					Ors: teo.RuleEngineRuleOrArray{
// 						&teo.RuleEngineRuleOrArgs{
// 							Ands: teo.RuleEngineRuleOrAndArray{
// 								&teo.RuleEngineRuleOrAndArgs{
// 									Operator:   pulumi.String("equal"),
// 									Target:     pulumi.String("host"),
// 									IgnoreCase: pulumi.Bool(false),
// 									Values: pulumi.StringArray{
// 										pulumi.String("a.tf-teo-t.xyz"),
// 									},
// 								},
// 								&teo.RuleEngineRuleOrAndArgs{
// 									Operator:   pulumi.String("equal"),
// 									Target:     pulumi.String("extension"),
// 									IgnoreCase: pulumi.Bool(false),
// 									Values: pulumi.StringArray{
// 										pulumi.String("jpg"),
// 									},
// 								},
// 							},
// 						},
// 						&teo.RuleEngineRuleOrArgs{
// 							Ands: teo.RuleEngineRuleOrAndArray{
// 								&teo.RuleEngineRuleOrAndArgs{
// 									Operator:   pulumi.String("equal"),
// 									Target:     pulumi.String("filename"),
// 									IgnoreCase: pulumi.Bool(false),
// 									Values: pulumi.StringArray{
// 										pulumi.String("test.txt"),
// 									},
// 								},
// 							},
// 						},
// 					},
// 					SubRules: teo.RuleEngineRuleSubRuleArray{
// 						&teo.RuleEngineRuleSubRuleArgs{
// 							Tags: pulumi.StringArray{
// 								pulumi.String("png"),
// 							},
// 							Rules: teo.RuleEngineRuleSubRuleRuleArray{
// 								&teo.RuleEngineRuleSubRuleRuleArgs{
// 									Ors: teo.RuleEngineRuleSubRuleRuleOrArray{
// 										&teo.RuleEngineRuleSubRuleRuleOrArgs{
// 											Ands: teo.RuleEngineRuleSubRuleRuleOrAndArray{
// 												&teo.RuleEngineRuleSubRuleRuleOrAndArgs{
// 													Operator:   pulumi.String("notequal"),
// 													Target:     pulumi.String("host"),
// 													IgnoreCase: pulumi.Bool(false),
// 													Values: pulumi.StringArray{
// 														pulumi.String("a.tf-teo-t.xyz"),
// 													},
// 												},
// 												&teo.RuleEngineRuleSubRuleRuleOrAndArgs{
// 													Operator:   pulumi.String("equal"),
// 													Target:     pulumi.String("extension"),
// 													IgnoreCase: pulumi.Bool(false),
// 													Values: pulumi.StringArray{
// 														pulumi.String("png"),
// 													},
// 												},
// 											},
// 										},
// 										&teo.RuleEngineRuleSubRuleRuleOrArgs{
// 											Ands: teo.RuleEngineRuleSubRuleRuleOrAndArray{
// 												&teo.RuleEngineRuleSubRuleRuleOrAndArgs{
// 													Operator:   pulumi.String("notequal"),
// 													Target:     pulumi.String("filename"),
// 													IgnoreCase: pulumi.Bool(false),
// 													Values: pulumi.StringArray{
// 														pulumi.String("test.txt"),
// 													},
// 												},
// 											},
// 										},
// 									},
// 									Actions: teo.RuleEngineRuleSubRuleRuleActionArray{
// 										&teo.RuleEngineRuleSubRuleRuleActionArgs{
// 											NormalAction: &teo.RuleEngineRuleSubRuleRuleActionNormalActionArgs{
// 												Action: pulumi.String("UpstreamUrlRedirect"),
// 												Parameters: teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArray{
// 													&teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs{
// 														Name: pulumi.String("Type"),
// 														Values: pulumi.StringArray{
// 															pulumi.String("Path"),
// 														},
// 													},
// 													&teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs{
// 														Name: pulumi.String("Action"),
// 														Values: pulumi.StringArray{
// 															pulumi.String("addPrefix"),
// 														},
// 													},
// 													&teo.RuleEngineRuleSubRuleRuleActionNormalActionParameterArgs{
// 														Name: pulumi.String("Value"),
// 														Values: pulumi.StringArray{
// 															pulumi.String("/www"),
// 														},
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// teo rule_engine can be imported using the id#rule_id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Teo/ruleEngine:RuleEngine rule_engine zone-297z8rf93cfw#rule-ajol584a
// ```
type RuleEngine struct {
	pulumi.CustomResourceState

	// Rule ID.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Rule name.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// Rule items list.
	Rules RuleEngineRuleArrayOutput `pulumi:"rules"`
	// Status of the rule, valid value can be `enable` or `disable`.
	Status pulumi.StringOutput `pulumi:"status"`
	// rule tag list.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Site ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewRuleEngine registers a new resource with the given unique name, arguments, and options.
func NewRuleEngine(ctx *pulumi.Context,
	name string, args *RuleEngineArgs, opts ...pulumi.ResourceOption) (*RuleEngine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RuleEngine
	err := ctx.RegisterResource("tencentcloud:Teo/ruleEngine:RuleEngine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleEngine gets an existing RuleEngine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleEngine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleEngineState, opts ...pulumi.ResourceOption) (*RuleEngine, error) {
	var resource RuleEngine
	err := ctx.ReadResource("tencentcloud:Teo/ruleEngine:RuleEngine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleEngine resources.
type ruleEngineState struct {
	// Rule ID.
	RuleId *string `pulumi:"ruleId"`
	// Rule name.
	RuleName *string `pulumi:"ruleName"`
	// Rule items list.
	Rules []RuleEngineRule `pulumi:"rules"`
	// Status of the rule, valid value can be `enable` or `disable`.
	Status *string `pulumi:"status"`
	// rule tag list.
	Tags []string `pulumi:"tags"`
	// Site ID.
	ZoneId *string `pulumi:"zoneId"`
}

type RuleEngineState struct {
	// Rule ID.
	RuleId pulumi.StringPtrInput
	// Rule name.
	RuleName pulumi.StringPtrInput
	// Rule items list.
	Rules RuleEngineRuleArrayInput
	// Status of the rule, valid value can be `enable` or `disable`.
	Status pulumi.StringPtrInput
	// rule tag list.
	Tags pulumi.StringArrayInput
	// Site ID.
	ZoneId pulumi.StringPtrInput
}

func (RuleEngineState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleEngineState)(nil)).Elem()
}

type ruleEngineArgs struct {
	// Rule name.
	RuleName string `pulumi:"ruleName"`
	// Rule items list.
	Rules []RuleEngineRule `pulumi:"rules"`
	// Status of the rule, valid value can be `enable` or `disable`.
	Status string `pulumi:"status"`
	// rule tag list.
	Tags []string `pulumi:"tags"`
	// Site ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a RuleEngine resource.
type RuleEngineArgs struct {
	// Rule name.
	RuleName pulumi.StringInput
	// Rule items list.
	Rules RuleEngineRuleArrayInput
	// Status of the rule, valid value can be `enable` or `disable`.
	Status pulumi.StringInput
	// rule tag list.
	Tags pulumi.StringArrayInput
	// Site ID.
	ZoneId pulumi.StringInput
}

func (RuleEngineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleEngineArgs)(nil)).Elem()
}

type RuleEngineInput interface {
	pulumi.Input

	ToRuleEngineOutput() RuleEngineOutput
	ToRuleEngineOutputWithContext(ctx context.Context) RuleEngineOutput
}

func (*RuleEngine) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleEngine)(nil)).Elem()
}

func (i *RuleEngine) ToRuleEngineOutput() RuleEngineOutput {
	return i.ToRuleEngineOutputWithContext(context.Background())
}

func (i *RuleEngine) ToRuleEngineOutputWithContext(ctx context.Context) RuleEngineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleEngineOutput)
}

// RuleEngineArrayInput is an input type that accepts RuleEngineArray and RuleEngineArrayOutput values.
// You can construct a concrete instance of `RuleEngineArrayInput` via:
//
//          RuleEngineArray{ RuleEngineArgs{...} }
type RuleEngineArrayInput interface {
	pulumi.Input

	ToRuleEngineArrayOutput() RuleEngineArrayOutput
	ToRuleEngineArrayOutputWithContext(context.Context) RuleEngineArrayOutput
}

type RuleEngineArray []RuleEngineInput

func (RuleEngineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleEngine)(nil)).Elem()
}

func (i RuleEngineArray) ToRuleEngineArrayOutput() RuleEngineArrayOutput {
	return i.ToRuleEngineArrayOutputWithContext(context.Background())
}

func (i RuleEngineArray) ToRuleEngineArrayOutputWithContext(ctx context.Context) RuleEngineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleEngineArrayOutput)
}

// RuleEngineMapInput is an input type that accepts RuleEngineMap and RuleEngineMapOutput values.
// You can construct a concrete instance of `RuleEngineMapInput` via:
//
//          RuleEngineMap{ "key": RuleEngineArgs{...} }
type RuleEngineMapInput interface {
	pulumi.Input

	ToRuleEngineMapOutput() RuleEngineMapOutput
	ToRuleEngineMapOutputWithContext(context.Context) RuleEngineMapOutput
}

type RuleEngineMap map[string]RuleEngineInput

func (RuleEngineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleEngine)(nil)).Elem()
}

func (i RuleEngineMap) ToRuleEngineMapOutput() RuleEngineMapOutput {
	return i.ToRuleEngineMapOutputWithContext(context.Background())
}

func (i RuleEngineMap) ToRuleEngineMapOutputWithContext(ctx context.Context) RuleEngineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleEngineMapOutput)
}

type RuleEngineOutput struct{ *pulumi.OutputState }

func (RuleEngineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleEngine)(nil)).Elem()
}

func (o RuleEngineOutput) ToRuleEngineOutput() RuleEngineOutput {
	return o
}

func (o RuleEngineOutput) ToRuleEngineOutputWithContext(ctx context.Context) RuleEngineOutput {
	return o
}

// Rule ID.
func (o RuleEngineOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleEngine) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Rule name.
func (o RuleEngineOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleEngine) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// Rule items list.
func (o RuleEngineOutput) Rules() RuleEngineRuleArrayOutput {
	return o.ApplyT(func(v *RuleEngine) RuleEngineRuleArrayOutput { return v.Rules }).(RuleEngineRuleArrayOutput)
}

// Status of the rule, valid value can be `enable` or `disable`.
func (o RuleEngineOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleEngine) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// rule tag list.
func (o RuleEngineOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RuleEngine) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Site ID.
func (o RuleEngineOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleEngine) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type RuleEngineArrayOutput struct{ *pulumi.OutputState }

func (RuleEngineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleEngine)(nil)).Elem()
}

func (o RuleEngineArrayOutput) ToRuleEngineArrayOutput() RuleEngineArrayOutput {
	return o
}

func (o RuleEngineArrayOutput) ToRuleEngineArrayOutputWithContext(ctx context.Context) RuleEngineArrayOutput {
	return o
}

func (o RuleEngineArrayOutput) Index(i pulumi.IntInput) RuleEngineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuleEngine {
		return vs[0].([]*RuleEngine)[vs[1].(int)]
	}).(RuleEngineOutput)
}

type RuleEngineMapOutput struct{ *pulumi.OutputState }

func (RuleEngineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleEngine)(nil)).Elem()
}

func (o RuleEngineMapOutput) ToRuleEngineMapOutput() RuleEngineMapOutput {
	return o
}

func (o RuleEngineMapOutput) ToRuleEngineMapOutputWithContext(ctx context.Context) RuleEngineMapOutput {
	return o
}

func (o RuleEngineMapOutput) MapIndex(k pulumi.StringInput) RuleEngineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuleEngine {
		return vs[0].(map[string]*RuleEngine)[vs[1].(string)]
	}).(RuleEngineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleEngineInput)(nil)).Elem(), &RuleEngine{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleEngineArrayInput)(nil)).Elem(), RuleEngineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleEngineMapInput)(nil)).Elem(), RuleEngineMap{})
	pulumi.RegisterOutputType(RuleEngineOutput{})
	pulumi.RegisterOutputType(RuleEngineArrayOutput{})
	pulumi.RegisterOutputType(RuleEngineMapOutput{})
}
