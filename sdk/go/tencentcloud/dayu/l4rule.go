// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this resource to create dayu layer 4 rule
//
// > **NOTE:** This resource only support resource Anti-DDoS of type `bgpip` and `net`
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dayu.NewL4Rule(ctx, "testRule", &Dayu.L4RuleArgs{
//				DPort:                  pulumi.Int(60),
//				HealthCheckHealthNum:   pulumi.Int(5),
//				HealthCheckInterval:    pulumi.Int(35),
//				HealthCheckSwitch:      pulumi.Bool(true),
//				HealthCheckTimeout:     pulumi.Int(30),
//				HealthCheckUnhealthNum: pulumi.Int(10),
//				Protocol:               pulumi.String("TCP"),
//				ResourceId:             pulumi.String("bgpip-00000294"),
//				ResourceType:           pulumi.String("bgpip"),
//				SPort:                  pulumi.Int(80),
//				SessionSwitch:          pulumi.Bool(false),
//				SessionTime:            pulumi.Int(300),
//				SourceLists: dayu.L4RuleSourceListArray{
//					&dayu.L4RuleSourceListArgs{
//						Source: pulumi.String("1.1.1.1"),
//						Weight: pulumi.Int(100),
//					},
//					&dayu.L4RuleSourceListArgs{
//						Source: pulumi.String("2.2.2.2"),
//						Weight: pulumi.Int(50),
//					},
//				},
//				SourceType: pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type L4Rule struct {
	pulumi.CustomResourceState

	// The destination port of the L4 rule.
	DPort pulumi.IntOutput `pulumi:"dPort"`
	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
	HealthCheckHealthNum pulumi.IntOutput `pulumi:"healthCheckHealthNum"`
	// Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
	HealthCheckInterval pulumi.IntOutput `pulumi:"healthCheckInterval"`
	// Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
	HealthCheckSwitch pulumi.BoolOutput `pulumi:"healthCheckSwitch"`
	// HTTP Status Code. The default is 26 and value range is 2-60.
	HealthCheckTimeout pulumi.IntOutput `pulumi:"healthCheckTimeout"`
	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
	HealthCheckUnhealthNum pulumi.IntOutput `pulumi:"healthCheckUnhealthNum"`
	// LB type of the rule. Valid values: `1`, `2`. `1` for weight cycling and `2` for IP hash.
	LbType pulumi.IntOutput `pulumi:"lbType"`
	// Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
	Name pulumi.StringOutput `pulumi:"name"`
	// Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// ID of the resource that the layer 4 rule works for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// ID of the layer 4 rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// The source port of the L4 rule.
	SPort pulumi.IntOutput `pulumi:"sPort"`
	// Indicate that the session will keep or not, and default value is `false`.
	SessionSwitch pulumi.BoolPtrOutput `pulumi:"sessionSwitch"`
	// Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
	SessionTime pulumi.IntOutput `pulumi:"sessionTime"`
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
	SourceLists L4RuleSourceListArrayOutput `pulumi:"sourceLists"`
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType pulumi.IntOutput `pulumi:"sourceType"`
}

// NewL4Rule registers a new resource with the given unique name, arguments, and options.
func NewL4Rule(ctx *pulumi.Context,
	name string, args *L4RuleArgs, opts ...pulumi.ResourceOption) (*L4Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DPort == nil {
		return nil, errors.New("invalid value for required argument 'DPort'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.SPort == nil {
		return nil, errors.New("invalid value for required argument 'SPort'")
	}
	if args.SourceLists == nil {
		return nil, errors.New("invalid value for required argument 'SourceLists'")
	}
	if args.SourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource L4Rule
	err := ctx.RegisterResource("tencentcloud:Dayu/l4Rule:L4Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL4Rule gets an existing L4Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL4Rule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L4RuleState, opts ...pulumi.ResourceOption) (*L4Rule, error) {
	var resource L4Rule
	err := ctx.ReadResource("tencentcloud:Dayu/l4Rule:L4Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L4Rule resources.
type l4ruleState struct {
	// The destination port of the L4 rule.
	DPort *int `pulumi:"dPort"`
	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
	HealthCheckHealthNum *int `pulumi:"healthCheckHealthNum"`
	// Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
	HealthCheckSwitch *bool `pulumi:"healthCheckSwitch"`
	// HTTP Status Code. The default is 26 and value range is 2-60.
	HealthCheckTimeout *int `pulumi:"healthCheckTimeout"`
	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
	HealthCheckUnhealthNum *int `pulumi:"healthCheckUnhealthNum"`
	// LB type of the rule. Valid values: `1`, `2`. `1` for weight cycling and `2` for IP hash.
	LbType *int `pulumi:"lbType"`
	// Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
	Name *string `pulumi:"name"`
	// Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
	Protocol *string `pulumi:"protocol"`
	// ID of the resource that the layer 4 rule works for.
	ResourceId *string `pulumi:"resourceId"`
	// Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	ResourceType *string `pulumi:"resourceType"`
	// ID of the layer 4 rule.
	RuleId *string `pulumi:"ruleId"`
	// The source port of the L4 rule.
	SPort *int `pulumi:"sPort"`
	// Indicate that the session will keep or not, and default value is `false`.
	SessionSwitch *bool `pulumi:"sessionSwitch"`
	// Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
	SessionTime *int `pulumi:"sessionTime"`
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
	SourceLists []L4RuleSourceList `pulumi:"sourceLists"`
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType *int `pulumi:"sourceType"`
}

type L4RuleState struct {
	// The destination port of the L4 rule.
	DPort pulumi.IntPtrInput
	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
	HealthCheckHealthNum pulumi.IntPtrInput
	// Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
	HealthCheckInterval pulumi.IntPtrInput
	// Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
	HealthCheckSwitch pulumi.BoolPtrInput
	// HTTP Status Code. The default is 26 and value range is 2-60.
	HealthCheckTimeout pulumi.IntPtrInput
	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
	HealthCheckUnhealthNum pulumi.IntPtrInput
	// LB type of the rule. Valid values: `1`, `2`. `1` for weight cycling and `2` for IP hash.
	LbType pulumi.IntPtrInput
	// Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
	Name pulumi.StringPtrInput
	// Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
	Protocol pulumi.StringPtrInput
	// ID of the resource that the layer 4 rule works for.
	ResourceId pulumi.StringPtrInput
	// Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	ResourceType pulumi.StringPtrInput
	// ID of the layer 4 rule.
	RuleId pulumi.StringPtrInput
	// The source port of the L4 rule.
	SPort pulumi.IntPtrInput
	// Indicate that the session will keep or not, and default value is `false`.
	SessionSwitch pulumi.BoolPtrInput
	// Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
	SessionTime pulumi.IntPtrInput
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
	SourceLists L4RuleSourceListArrayInput
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType pulumi.IntPtrInput
}

func (L4RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*l4ruleState)(nil)).Elem()
}

type l4ruleArgs struct {
	// The destination port of the L4 rule.
	DPort int `pulumi:"dPort"`
	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
	HealthCheckHealthNum *int `pulumi:"healthCheckHealthNum"`
	// Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
	HealthCheckInterval *int `pulumi:"healthCheckInterval"`
	// Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
	HealthCheckSwitch *bool `pulumi:"healthCheckSwitch"`
	// HTTP Status Code. The default is 26 and value range is 2-60.
	HealthCheckTimeout *int `pulumi:"healthCheckTimeout"`
	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
	HealthCheckUnhealthNum *int `pulumi:"healthCheckUnhealthNum"`
	// Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
	Name *string `pulumi:"name"`
	// Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
	Protocol string `pulumi:"protocol"`
	// ID of the resource that the layer 4 rule works for.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	ResourceType string `pulumi:"resourceType"`
	// The source port of the L4 rule.
	SPort int `pulumi:"sPort"`
	// Indicate that the session will keep or not, and default value is `false`.
	SessionSwitch *bool `pulumi:"sessionSwitch"`
	// Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
	SessionTime *int `pulumi:"sessionTime"`
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
	SourceLists []L4RuleSourceList `pulumi:"sourceLists"`
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType int `pulumi:"sourceType"`
}

// The set of arguments for constructing a L4Rule resource.
type L4RuleArgs struct {
	// The destination port of the L4 rule.
	DPort pulumi.IntInput
	// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
	HealthCheckHealthNum pulumi.IntPtrInput
	// Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
	HealthCheckInterval pulumi.IntPtrInput
	// Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
	HealthCheckSwitch pulumi.BoolPtrInput
	// HTTP Status Code. The default is 26 and value range is 2-60.
	HealthCheckTimeout pulumi.IntPtrInput
	// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
	HealthCheckUnhealthNum pulumi.IntPtrInput
	// Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
	Name pulumi.StringPtrInput
	// Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
	Protocol pulumi.StringInput
	// ID of the resource that the layer 4 rule works for.
	ResourceId pulumi.StringInput
	// Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
	ResourceType pulumi.StringInput
	// The source port of the L4 rule.
	SPort pulumi.IntInput
	// Indicate that the session will keep or not, and default value is `false`.
	SessionSwitch pulumi.BoolPtrInput
	// Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
	SessionTime pulumi.IntPtrInput
	// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
	SourceLists L4RuleSourceListArrayInput
	// Source type, `1` for source of host, `2` for source of IP.
	SourceType pulumi.IntInput
}

func (L4RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*l4ruleArgs)(nil)).Elem()
}

type L4RuleInput interface {
	pulumi.Input

	ToL4RuleOutput() L4RuleOutput
	ToL4RuleOutputWithContext(ctx context.Context) L4RuleOutput
}

func (*L4Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**L4Rule)(nil)).Elem()
}

func (i *L4Rule) ToL4RuleOutput() L4RuleOutput {
	return i.ToL4RuleOutputWithContext(context.Background())
}

func (i *L4Rule) ToL4RuleOutputWithContext(ctx context.Context) L4RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L4RuleOutput)
}

// L4RuleArrayInput is an input type that accepts L4RuleArray and L4RuleArrayOutput values.
// You can construct a concrete instance of `L4RuleArrayInput` via:
//
//	L4RuleArray{ L4RuleArgs{...} }
type L4RuleArrayInput interface {
	pulumi.Input

	ToL4RuleArrayOutput() L4RuleArrayOutput
	ToL4RuleArrayOutputWithContext(context.Context) L4RuleArrayOutput
}

type L4RuleArray []L4RuleInput

func (L4RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L4Rule)(nil)).Elem()
}

func (i L4RuleArray) ToL4RuleArrayOutput() L4RuleArrayOutput {
	return i.ToL4RuleArrayOutputWithContext(context.Background())
}

func (i L4RuleArray) ToL4RuleArrayOutputWithContext(ctx context.Context) L4RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L4RuleArrayOutput)
}

// L4RuleMapInput is an input type that accepts L4RuleMap and L4RuleMapOutput values.
// You can construct a concrete instance of `L4RuleMapInput` via:
//
//	L4RuleMap{ "key": L4RuleArgs{...} }
type L4RuleMapInput interface {
	pulumi.Input

	ToL4RuleMapOutput() L4RuleMapOutput
	ToL4RuleMapOutputWithContext(context.Context) L4RuleMapOutput
}

type L4RuleMap map[string]L4RuleInput

func (L4RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L4Rule)(nil)).Elem()
}

func (i L4RuleMap) ToL4RuleMapOutput() L4RuleMapOutput {
	return i.ToL4RuleMapOutputWithContext(context.Background())
}

func (i L4RuleMap) ToL4RuleMapOutputWithContext(ctx context.Context) L4RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L4RuleMapOutput)
}

type L4RuleOutput struct{ *pulumi.OutputState }

func (L4RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**L4Rule)(nil)).Elem()
}

func (o L4RuleOutput) ToL4RuleOutput() L4RuleOutput {
	return o
}

func (o L4RuleOutput) ToL4RuleOutputWithContext(ctx context.Context) L4RuleOutput {
	return o
}

// The destination port of the L4 rule.
func (o L4RuleOutput) DPort() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.DPort }).(pulumi.IntOutput)
}

// Health threshold of health check, and the default is 3. If a success result is returned for the health check 3 consecutive times, indicates that the forwarding is normal. The value range is 2-10.
func (o L4RuleOutput) HealthCheckHealthNum() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.HealthCheckHealthNum }).(pulumi.IntOutput)
}

// Interval time of health check. The value range is 10-60 sec, and the default is 15 sec.
func (o L4RuleOutput) HealthCheckInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.HealthCheckInterval }).(pulumi.IntOutput)
}

// Indicates whether health check is enabled. The default is `false`. Only valid when source list has more than one source item.
func (o L4RuleOutput) HealthCheckSwitch() pulumi.BoolOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.BoolOutput { return v.HealthCheckSwitch }).(pulumi.BoolOutput)
}

// HTTP Status Code. The default is 26 and value range is 2-60.
func (o L4RuleOutput) HealthCheckTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.HealthCheckTimeout }).(pulumi.IntOutput)
}

// Unhealthy threshold of health check, and the default is 3. If the unhealthy result is returned 3 consecutive times, indicates that the forwarding is abnormal. The value range is 2-10.
func (o L4RuleOutput) HealthCheckUnhealthNum() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.HealthCheckUnhealthNum }).(pulumi.IntOutput)
}

// LB type of the rule. Valid values: `1`, `2`. `1` for weight cycling and `2` for IP hash.
func (o L4RuleOutput) LbType() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.LbType }).(pulumi.IntOutput)
}

// Name of the rule. When the `resourceType` is `net`, this field should be set with valid domain.
func (o L4RuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Protocol of the rule. Valid values: `http`, `https`. When `sourceType` is 1(host source), the value of this field can only set with `tcp`.
func (o L4RuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// ID of the resource that the layer 4 rule works for.
func (o L4RuleOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Type of the resource that the layer 4 rule works for. Valid values: `bgpip` and `net`.
func (o L4RuleOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// ID of the layer 4 rule.
func (o L4RuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// The source port of the L4 rule.
func (o L4RuleOutput) SPort() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.SPort }).(pulumi.IntOutput)
}

// Indicate that the session will keep or not, and default value is `false`.
func (o L4RuleOutput) SessionSwitch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.BoolPtrOutput { return v.SessionSwitch }).(pulumi.BoolPtrOutput)
}

// Session keep time, only valid when `sessionSwitch` is true, the available value ranges from 1 to 300 and unit is second.
func (o L4RuleOutput) SessionTime() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.SessionTime }).(pulumi.IntOutput)
}

// Source list of the rule, it can be a set of ip sources or a set of domain sources. The number of items ranges from 1 to 20.
func (o L4RuleOutput) SourceLists() L4RuleSourceListArrayOutput {
	return o.ApplyT(func(v *L4Rule) L4RuleSourceListArrayOutput { return v.SourceLists }).(L4RuleSourceListArrayOutput)
}

// Source type, `1` for source of host, `2` for source of IP.
func (o L4RuleOutput) SourceType() pulumi.IntOutput {
	return o.ApplyT(func(v *L4Rule) pulumi.IntOutput { return v.SourceType }).(pulumi.IntOutput)
}

type L4RuleArrayOutput struct{ *pulumi.OutputState }

func (L4RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L4Rule)(nil)).Elem()
}

func (o L4RuleArrayOutput) ToL4RuleArrayOutput() L4RuleArrayOutput {
	return o
}

func (o L4RuleArrayOutput) ToL4RuleArrayOutputWithContext(ctx context.Context) L4RuleArrayOutput {
	return o
}

func (o L4RuleArrayOutput) Index(i pulumi.IntInput) L4RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *L4Rule {
		return vs[0].([]*L4Rule)[vs[1].(int)]
	}).(L4RuleOutput)
}

type L4RuleMapOutput struct{ *pulumi.OutputState }

func (L4RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L4Rule)(nil)).Elem()
}

func (o L4RuleMapOutput) ToL4RuleMapOutput() L4RuleMapOutput {
	return o
}

func (o L4RuleMapOutput) ToL4RuleMapOutputWithContext(ctx context.Context) L4RuleMapOutput {
	return o
}

func (o L4RuleMapOutput) MapIndex(k pulumi.StringInput) L4RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *L4Rule {
		return vs[0].(map[string]*L4Rule)[vs[1].(string)]
	}).(L4RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*L4RuleInput)(nil)).Elem(), &L4Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*L4RuleArrayInput)(nil)).Elem(), L4RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*L4RuleMapInput)(nil)).Elem(), L4RuleMap{})
	pulumi.RegisterOutputType(L4RuleOutput{})
	pulumi.RegisterOutputType(L4RuleArrayOutput{})
	pulumi.RegisterOutputType(L4RuleMapOutput{})
}
