// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a monitor tmpAlertRule
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			availabilityZone := "ap-guangzhou-4"
//			if param := cfg.Get("availabilityZone"); param != "" {
//				availabilityZone = param
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				VpcId:            vpc.ID(),
//				AvailabilityZone: pulumi.String(availabilityZone),
//				CidrBlock:        pulumi.String("10.0.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			fooTmpInstance, err := Monitor.NewTmpInstance(ctx, "fooTmpInstance", &Monitor.TmpInstanceArgs{
//				InstanceName:      pulumi.String("tf-tmp-instance"),
//				VpcId:             vpc.ID(),
//				SubnetId:          subnet.ID(),
//				DataRetentionTime: pulumi.Int(30),
//				Zone:              pulumi.String(availabilityZone),
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Monitor.NewTmpCvmAgent(ctx, "fooTmpCvmAgent", &Monitor.TmpCvmAgentArgs{
//				InstanceId: fooTmpInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Monitor.NewTmpAlertRule(ctx, "fooTmpAlertRule", &Monitor.TmpAlertRuleArgs{
//				Duration:   pulumi.String("2m"),
//				Expr:       pulumi.String("avg by (instance) (mysql_global_status_threads_connected) / avg by (instance) (mysql_global_variables_max_connections)  > 0.8"),
//				InstanceId: fooTmpInstance.ID(),
//				Receivers: pulumi.StringArray{
//					pulumi.String("notice-f2svbu3w"),
//				},
//				RuleName:  pulumi.String("MySQL 连接数过多"),
//				RuleState: pulumi.Int(2),
//				Type:      pulumi.String("MySQL/MySQL 连接数过多"),
//				Annotations: monitor.TmpAlertRuleAnnotationArray{
//					&monitor.TmpAlertRuleAnnotationArgs{
//						Key:   pulumi.String("description"),
//						Value: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "MySQL 连接数过多, 实例: {{", "$", "labels.instance}}，当前值: {{ ", "$", "value | humanizePercentage }}。")),
//					},
//					&monitor.TmpAlertRuleAnnotationArgs{
//						Key:   pulumi.String("summary"),
//						Value: pulumi.String(fmt.Sprintf("%v%v%v", "MySQL 连接数过多(>80", "%", ")")),
//					},
//				},
//				Labels: monitor.TmpAlertRuleLabelArray{
//					&monitor.TmpAlertRuleLabelArgs{
//						Key:   pulumi.String("severity"),
//						Value: pulumi.String("warning"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// monitor tmpAlertRule can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Monitor/tmpAlertRule:TmpAlertRule tmpAlertRule instanceId#Rule_id
//
// ```
type TmpAlertRule struct {
	pulumi.CustomResourceState

	// Rule alarm duration.
	Annotations TmpAlertRuleAnnotationArrayOutput `pulumi:"annotations"`
	// Rule alarm duration.
	Duration pulumi.StringPtrOutput `pulumi:"duration"`
	// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
	Expr pulumi.StringOutput `pulumi:"expr"`
	// Instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Rule alarm duration.
	Labels TmpAlertRuleLabelArrayOutput `pulumi:"labels"`
	// Alarm notification template id list.
	Receivers pulumi.StringArrayOutput `pulumi:"receivers"`
	// Rule name.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`
	// Rule state code.
	RuleState pulumi.IntPtrOutput `pulumi:"ruleState"`
	// Alarm Policy Template Classification.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewTmpAlertRule registers a new resource with the given unique name, arguments, and options.
func NewTmpAlertRule(ctx *pulumi.Context,
	name string, args *TmpAlertRuleArgs, opts ...pulumi.ResourceOption) (*TmpAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expr == nil {
		return nil, errors.New("invalid value for required argument 'Expr'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Receivers == nil {
		return nil, errors.New("invalid value for required argument 'Receivers'")
	}
	if args.RuleName == nil {
		return nil, errors.New("invalid value for required argument 'RuleName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TmpAlertRule
	err := ctx.RegisterResource("tencentcloud:Monitor/tmpAlertRule:TmpAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTmpAlertRule gets an existing TmpAlertRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTmpAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TmpAlertRuleState, opts ...pulumi.ResourceOption) (*TmpAlertRule, error) {
	var resource TmpAlertRule
	err := ctx.ReadResource("tencentcloud:Monitor/tmpAlertRule:TmpAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TmpAlertRule resources.
type tmpAlertRuleState struct {
	// Rule alarm duration.
	Annotations []TmpAlertRuleAnnotation `pulumi:"annotations"`
	// Rule alarm duration.
	Duration *string `pulumi:"duration"`
	// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
	Expr *string `pulumi:"expr"`
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Rule alarm duration.
	Labels []TmpAlertRuleLabel `pulumi:"labels"`
	// Alarm notification template id list.
	Receivers []string `pulumi:"receivers"`
	// Rule name.
	RuleName *string `pulumi:"ruleName"`
	// Rule state code.
	RuleState *int `pulumi:"ruleState"`
	// Alarm Policy Template Classification.
	Type *string `pulumi:"type"`
}

type TmpAlertRuleState struct {
	// Rule alarm duration.
	Annotations TmpAlertRuleAnnotationArrayInput
	// Rule alarm duration.
	Duration pulumi.StringPtrInput
	// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
	Expr pulumi.StringPtrInput
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Rule alarm duration.
	Labels TmpAlertRuleLabelArrayInput
	// Alarm notification template id list.
	Receivers pulumi.StringArrayInput
	// Rule name.
	RuleName pulumi.StringPtrInput
	// Rule state code.
	RuleState pulumi.IntPtrInput
	// Alarm Policy Template Classification.
	Type pulumi.StringPtrInput
}

func (TmpAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpAlertRuleState)(nil)).Elem()
}

type tmpAlertRuleArgs struct {
	// Rule alarm duration.
	Annotations []TmpAlertRuleAnnotation `pulumi:"annotations"`
	// Rule alarm duration.
	Duration *string `pulumi:"duration"`
	// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
	Expr string `pulumi:"expr"`
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
	// Rule alarm duration.
	Labels []TmpAlertRuleLabel `pulumi:"labels"`
	// Alarm notification template id list.
	Receivers []string `pulumi:"receivers"`
	// Rule name.
	RuleName string `pulumi:"ruleName"`
	// Rule state code.
	RuleState *int `pulumi:"ruleState"`
	// Alarm Policy Template Classification.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a TmpAlertRule resource.
type TmpAlertRuleArgs struct {
	// Rule alarm duration.
	Annotations TmpAlertRuleAnnotationArrayInput
	// Rule alarm duration.
	Duration pulumi.StringPtrInput
	// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
	Expr pulumi.StringInput
	// Instance id.
	InstanceId pulumi.StringInput
	// Rule alarm duration.
	Labels TmpAlertRuleLabelArrayInput
	// Alarm notification template id list.
	Receivers pulumi.StringArrayInput
	// Rule name.
	RuleName pulumi.StringInput
	// Rule state code.
	RuleState pulumi.IntPtrInput
	// Alarm Policy Template Classification.
	Type pulumi.StringPtrInput
}

func (TmpAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tmpAlertRuleArgs)(nil)).Elem()
}

type TmpAlertRuleInput interface {
	pulumi.Input

	ToTmpAlertRuleOutput() TmpAlertRuleOutput
	ToTmpAlertRuleOutputWithContext(ctx context.Context) TmpAlertRuleOutput
}

func (*TmpAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpAlertRule)(nil)).Elem()
}

func (i *TmpAlertRule) ToTmpAlertRuleOutput() TmpAlertRuleOutput {
	return i.ToTmpAlertRuleOutputWithContext(context.Background())
}

func (i *TmpAlertRule) ToTmpAlertRuleOutputWithContext(ctx context.Context) TmpAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpAlertRuleOutput)
}

// TmpAlertRuleArrayInput is an input type that accepts TmpAlertRuleArray and TmpAlertRuleArrayOutput values.
// You can construct a concrete instance of `TmpAlertRuleArrayInput` via:
//
//	TmpAlertRuleArray{ TmpAlertRuleArgs{...} }
type TmpAlertRuleArrayInput interface {
	pulumi.Input

	ToTmpAlertRuleArrayOutput() TmpAlertRuleArrayOutput
	ToTmpAlertRuleArrayOutputWithContext(context.Context) TmpAlertRuleArrayOutput
}

type TmpAlertRuleArray []TmpAlertRuleInput

func (TmpAlertRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpAlertRule)(nil)).Elem()
}

func (i TmpAlertRuleArray) ToTmpAlertRuleArrayOutput() TmpAlertRuleArrayOutput {
	return i.ToTmpAlertRuleArrayOutputWithContext(context.Background())
}

func (i TmpAlertRuleArray) ToTmpAlertRuleArrayOutputWithContext(ctx context.Context) TmpAlertRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpAlertRuleArrayOutput)
}

// TmpAlertRuleMapInput is an input type that accepts TmpAlertRuleMap and TmpAlertRuleMapOutput values.
// You can construct a concrete instance of `TmpAlertRuleMapInput` via:
//
//	TmpAlertRuleMap{ "key": TmpAlertRuleArgs{...} }
type TmpAlertRuleMapInput interface {
	pulumi.Input

	ToTmpAlertRuleMapOutput() TmpAlertRuleMapOutput
	ToTmpAlertRuleMapOutputWithContext(context.Context) TmpAlertRuleMapOutput
}

type TmpAlertRuleMap map[string]TmpAlertRuleInput

func (TmpAlertRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpAlertRule)(nil)).Elem()
}

func (i TmpAlertRuleMap) ToTmpAlertRuleMapOutput() TmpAlertRuleMapOutput {
	return i.ToTmpAlertRuleMapOutputWithContext(context.Background())
}

func (i TmpAlertRuleMap) ToTmpAlertRuleMapOutputWithContext(ctx context.Context) TmpAlertRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TmpAlertRuleMapOutput)
}

type TmpAlertRuleOutput struct{ *pulumi.OutputState }

func (TmpAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TmpAlertRule)(nil)).Elem()
}

func (o TmpAlertRuleOutput) ToTmpAlertRuleOutput() TmpAlertRuleOutput {
	return o
}

func (o TmpAlertRuleOutput) ToTmpAlertRuleOutputWithContext(ctx context.Context) TmpAlertRuleOutput {
	return o
}

// Rule alarm duration.
func (o TmpAlertRuleOutput) Annotations() TmpAlertRuleAnnotationArrayOutput {
	return o.ApplyT(func(v *TmpAlertRule) TmpAlertRuleAnnotationArrayOutput { return v.Annotations }).(TmpAlertRuleAnnotationArrayOutput)
}

// Rule alarm duration.
func (o TmpAlertRuleOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TmpAlertRule) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

// Rule expression, reference documentation: `https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/`.
func (o TmpAlertRuleOutput) Expr() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpAlertRule) pulumi.StringOutput { return v.Expr }).(pulumi.StringOutput)
}

// Instance id.
func (o TmpAlertRuleOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpAlertRule) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Rule alarm duration.
func (o TmpAlertRuleOutput) Labels() TmpAlertRuleLabelArrayOutput {
	return o.ApplyT(func(v *TmpAlertRule) TmpAlertRuleLabelArrayOutput { return v.Labels }).(TmpAlertRuleLabelArrayOutput)
}

// Alarm notification template id list.
func (o TmpAlertRuleOutput) Receivers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TmpAlertRule) pulumi.StringArrayOutput { return v.Receivers }).(pulumi.StringArrayOutput)
}

// Rule name.
func (o TmpAlertRuleOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v *TmpAlertRule) pulumi.StringOutput { return v.RuleName }).(pulumi.StringOutput)
}

// Rule state code.
func (o TmpAlertRuleOutput) RuleState() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TmpAlertRule) pulumi.IntPtrOutput { return v.RuleState }).(pulumi.IntPtrOutput)
}

// Alarm Policy Template Classification.
func (o TmpAlertRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TmpAlertRule) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

type TmpAlertRuleArrayOutput struct{ *pulumi.OutputState }

func (TmpAlertRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TmpAlertRule)(nil)).Elem()
}

func (o TmpAlertRuleArrayOutput) ToTmpAlertRuleArrayOutput() TmpAlertRuleArrayOutput {
	return o
}

func (o TmpAlertRuleArrayOutput) ToTmpAlertRuleArrayOutputWithContext(ctx context.Context) TmpAlertRuleArrayOutput {
	return o
}

func (o TmpAlertRuleArrayOutput) Index(i pulumi.IntInput) TmpAlertRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TmpAlertRule {
		return vs[0].([]*TmpAlertRule)[vs[1].(int)]
	}).(TmpAlertRuleOutput)
}

type TmpAlertRuleMapOutput struct{ *pulumi.OutputState }

func (TmpAlertRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TmpAlertRule)(nil)).Elem()
}

func (o TmpAlertRuleMapOutput) ToTmpAlertRuleMapOutput() TmpAlertRuleMapOutput {
	return o
}

func (o TmpAlertRuleMapOutput) ToTmpAlertRuleMapOutputWithContext(ctx context.Context) TmpAlertRuleMapOutput {
	return o
}

func (o TmpAlertRuleMapOutput) MapIndex(k pulumi.StringInput) TmpAlertRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TmpAlertRule {
		return vs[0].(map[string]*TmpAlertRule)[vs[1].(string)]
	}).(TmpAlertRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TmpAlertRuleInput)(nil)).Elem(), &TmpAlertRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpAlertRuleArrayInput)(nil)).Elem(), TmpAlertRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TmpAlertRuleMapInput)(nil)).Elem(), TmpAlertRuleMap{})
	pulumi.RegisterOutputType(TmpAlertRuleOutput{})
	pulumi.RegisterOutputType(TmpAlertRuleArrayOutput{})
	pulumi.RegisterOutputType(TmpAlertRuleMapOutput{})
}
