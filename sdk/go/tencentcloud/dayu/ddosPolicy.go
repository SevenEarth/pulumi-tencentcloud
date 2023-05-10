// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create dayu DDoS policy
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
// 		_, err := Dayu.NewDdosPolicy(ctx, "testPolicy", &Dayu.DdosPolicyArgs{
// 			BlackIps: pulumi.StringArray{
// 				pulumi.String("1.1.1.1"),
// 			},
// 			DropOptions: dayu.DdosPolicyDropOptionArray{
// 				&dayu.DdosPolicyDropOptionArgs{
// 					BadConnThreshold: pulumi.Int(100),
// 					CheckSyncConn:    pulumi.Bool(true),
// 					ConnTimeout:      pulumi.Int(500),
// 					DConnLimit:       pulumi.Int(100),
// 					DNewLimit:        pulumi.Int(100),
// 					DropAbroad:       pulumi.Bool(true),
// 					DropIcmp:         pulumi.Bool(true),
// 					DropOther:        pulumi.Bool(true),
// 					DropTcp:          pulumi.Bool(true),
// 					DropUdp:          pulumi.Bool(true),
// 					IcmpMbpsLimit:    pulumi.Int(100),
// 					NullConnEnable:   pulumi.Bool(true),
// 					OtherMbpsLimit:   pulumi.Int(100),
// 					SConnLimit:       pulumi.Int(100),
// 					SNewLimit:        pulumi.Int(100),
// 					SynLimit:         pulumi.Int(100),
// 					SynRate:          pulumi.Int(50),
// 					TcpMbpsLimit:     pulumi.Int(100),
// 					UdpMbpsLimit:     pulumi.Int(100),
// 				},
// 			},
// 			PacketFilters: dayu.DdosPolicyPacketFilterArray{
// 				&dayu.DdosPolicyPacketFilterArgs{
// 					Action:       pulumi.String("drop"),
// 					DEndPort:     pulumi.Int(1500),
// 					DStartPort:   pulumi.Int(1000),
// 					Depth:        pulumi.Int(1000),
// 					IsInclude:    pulumi.Bool(true),
// 					MatchBegin:   pulumi.String("begin_l5"),
// 					MatchType:    pulumi.String("pcre"),
// 					Offset:       pulumi.Int(500),
// 					PktLengthMax: pulumi.Int(1400),
// 					PktLengthMin: pulumi.Int(1000),
// 					Protocol:     pulumi.String("tcp"),
// 					SEndPort:     pulumi.Int(2500),
// 					SStartPort:   pulumi.Int(2000),
// 				},
// 			},
// 			PortFilters: dayu.DdosPolicyPortFilterArray{
// 				&dayu.DdosPolicyPortFilterArgs{
// 					Action:    pulumi.String("drop"),
// 					EndPort:   pulumi.Int(2500),
// 					Kind:      pulumi.Int(1),
// 					Protocol:  pulumi.String("all"),
// 					StartPort: pulumi.Int(2000),
// 				},
// 			},
// 			ResourceType: pulumi.String("bgpip"),
// 			WatermarkFilters: dayu.DdosPolicyWatermarkFilterArray{
// 				&dayu.DdosPolicyWatermarkFilterArgs{
// 					AutoRemove: pulumi.Bool(true),
// 					Offset:     pulumi.Int(50),
// 					OpenSwitch: pulumi.Bool(true),
// 					TcpPortLists: pulumi.StringArray{
// 						pulumi.String("2000-3000"),
// 						pulumi.String("3500-4000"),
// 					},
// 					UdpPortLists: pulumi.StringArray{
// 						pulumi.String("5000-6000"),
// 					},
// 				},
// 			},
// 			WhiteIps: pulumi.StringArray{
// 				pulumi.String("2.2.2.2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DdosPolicy struct {
	pulumi.CustomResourceState

	// Black IP list.
	BlackIps pulumi.StringArrayOutput `pulumi:"blackIps"`
	// Create time of the DDoS policy.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Option list of abnormal check of the DDos policy, should set at least one policy.
	DropOptions DdosPolicyDropOptionArrayOutput `pulumi:"dropOptions"`
	// Name of the DDoS policy. Length should between 1 and 32.
	Name pulumi.StringOutput `pulumi:"name"`
	// Message filter options list.
	PacketFilters DdosPolicyPacketFilterArrayOutput `pulumi:"packetFilters"`
	// Id of policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// Port limits of abnormal check of the DDos policy.
	PortFilters DdosPolicyPortFilterArrayOutput `pulumi:"portFilters"`
	// Type of the resource that the DDoS policy works for. Valid values: `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Id of policy case that the DDoS policy works for.
	SceneId pulumi.StringOutput `pulumi:"sceneId"`
	// Watermark policy options, and only support one watermark policy at most.
	WatermarkFilters DdosPolicyWatermarkFilterArrayOutput `pulumi:"watermarkFilters"`
	// Watermark content.
	WatermarkKeys DdosPolicyWatermarkKeyArrayOutput `pulumi:"watermarkKeys"`
	// White IP list.
	WhiteIps pulumi.StringArrayOutput `pulumi:"whiteIps"`
}

// NewDdosPolicy registers a new resource with the given unique name, arguments, and options.
func NewDdosPolicy(ctx *pulumi.Context,
	name string, args *DdosPolicyArgs, opts ...pulumi.ResourceOption) (*DdosPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DropOptions == nil {
		return nil, errors.New("invalid value for required argument 'DropOptions'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DdosPolicy
	err := ctx.RegisterResource("tencentcloud:Dayu/ddosPolicy:DdosPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosPolicy gets an existing DdosPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosPolicyState, opts ...pulumi.ResourceOption) (*DdosPolicy, error) {
	var resource DdosPolicy
	err := ctx.ReadResource("tencentcloud:Dayu/ddosPolicy:DdosPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosPolicy resources.
type ddosPolicyState struct {
	// Black IP list.
	BlackIps []string `pulumi:"blackIps"`
	// Create time of the DDoS policy.
	CreateTime *string `pulumi:"createTime"`
	// Option list of abnormal check of the DDos policy, should set at least one policy.
	DropOptions []DdosPolicyDropOption `pulumi:"dropOptions"`
	// Name of the DDoS policy. Length should between 1 and 32.
	Name *string `pulumi:"name"`
	// Message filter options list.
	PacketFilters []DdosPolicyPacketFilter `pulumi:"packetFilters"`
	// Id of policy.
	PolicyId *string `pulumi:"policyId"`
	// Port limits of abnormal check of the DDos policy.
	PortFilters []DdosPolicyPortFilter `pulumi:"portFilters"`
	// Type of the resource that the DDoS policy works for. Valid values: `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType *string `pulumi:"resourceType"`
	// Id of policy case that the DDoS policy works for.
	SceneId *string `pulumi:"sceneId"`
	// Watermark policy options, and only support one watermark policy at most.
	WatermarkFilters []DdosPolicyWatermarkFilter `pulumi:"watermarkFilters"`
	// Watermark content.
	WatermarkKeys []DdosPolicyWatermarkKey `pulumi:"watermarkKeys"`
	// White IP list.
	WhiteIps []string `pulumi:"whiteIps"`
}

type DdosPolicyState struct {
	// Black IP list.
	BlackIps pulumi.StringArrayInput
	// Create time of the DDoS policy.
	CreateTime pulumi.StringPtrInput
	// Option list of abnormal check of the DDos policy, should set at least one policy.
	DropOptions DdosPolicyDropOptionArrayInput
	// Name of the DDoS policy. Length should between 1 and 32.
	Name pulumi.StringPtrInput
	// Message filter options list.
	PacketFilters DdosPolicyPacketFilterArrayInput
	// Id of policy.
	PolicyId pulumi.StringPtrInput
	// Port limits of abnormal check of the DDos policy.
	PortFilters DdosPolicyPortFilterArrayInput
	// Type of the resource that the DDoS policy works for. Valid values: `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringPtrInput
	// Id of policy case that the DDoS policy works for.
	SceneId pulumi.StringPtrInput
	// Watermark policy options, and only support one watermark policy at most.
	WatermarkFilters DdosPolicyWatermarkFilterArrayInput
	// Watermark content.
	WatermarkKeys DdosPolicyWatermarkKeyArrayInput
	// White IP list.
	WhiteIps pulumi.StringArrayInput
}

func (DdosPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosPolicyState)(nil)).Elem()
}

type ddosPolicyArgs struct {
	// Black IP list.
	BlackIps []string `pulumi:"blackIps"`
	// Option list of abnormal check of the DDos policy, should set at least one policy.
	DropOptions []DdosPolicyDropOption `pulumi:"dropOptions"`
	// Name of the DDoS policy. Length should between 1 and 32.
	Name *string `pulumi:"name"`
	// Message filter options list.
	PacketFilters []DdosPolicyPacketFilter `pulumi:"packetFilters"`
	// Port limits of abnormal check of the DDos policy.
	PortFilters []DdosPolicyPortFilter `pulumi:"portFilters"`
	// Type of the resource that the DDoS policy works for. Valid values: `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType string `pulumi:"resourceType"`
	// Watermark policy options, and only support one watermark policy at most.
	WatermarkFilters []DdosPolicyWatermarkFilter `pulumi:"watermarkFilters"`
	// White IP list.
	WhiteIps []string `pulumi:"whiteIps"`
}

// The set of arguments for constructing a DdosPolicy resource.
type DdosPolicyArgs struct {
	// Black IP list.
	BlackIps pulumi.StringArrayInput
	// Option list of abnormal check of the DDos policy, should set at least one policy.
	DropOptions DdosPolicyDropOptionArrayInput
	// Name of the DDoS policy. Length should between 1 and 32.
	Name pulumi.StringPtrInput
	// Message filter options list.
	PacketFilters DdosPolicyPacketFilterArrayInput
	// Port limits of abnormal check of the DDos policy.
	PortFilters DdosPolicyPortFilterArrayInput
	// Type of the resource that the DDoS policy works for. Valid values: `bgpip`, `bgp`, `bgp-multip` and `net`.
	ResourceType pulumi.StringInput
	// Watermark policy options, and only support one watermark policy at most.
	WatermarkFilters DdosPolicyWatermarkFilterArrayInput
	// White IP list.
	WhiteIps pulumi.StringArrayInput
}

func (DdosPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosPolicyArgs)(nil)).Elem()
}

type DdosPolicyInput interface {
	pulumi.Input

	ToDdosPolicyOutput() DdosPolicyOutput
	ToDdosPolicyOutputWithContext(ctx context.Context) DdosPolicyOutput
}

func (*DdosPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosPolicy)(nil)).Elem()
}

func (i *DdosPolicy) ToDdosPolicyOutput() DdosPolicyOutput {
	return i.ToDdosPolicyOutputWithContext(context.Background())
}

func (i *DdosPolicy) ToDdosPolicyOutputWithContext(ctx context.Context) DdosPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyOutput)
}

// DdosPolicyArrayInput is an input type that accepts DdosPolicyArray and DdosPolicyArrayOutput values.
// You can construct a concrete instance of `DdosPolicyArrayInput` via:
//
//          DdosPolicyArray{ DdosPolicyArgs{...} }
type DdosPolicyArrayInput interface {
	pulumi.Input

	ToDdosPolicyArrayOutput() DdosPolicyArrayOutput
	ToDdosPolicyArrayOutputWithContext(context.Context) DdosPolicyArrayOutput
}

type DdosPolicyArray []DdosPolicyInput

func (DdosPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosPolicy)(nil)).Elem()
}

func (i DdosPolicyArray) ToDdosPolicyArrayOutput() DdosPolicyArrayOutput {
	return i.ToDdosPolicyArrayOutputWithContext(context.Background())
}

func (i DdosPolicyArray) ToDdosPolicyArrayOutputWithContext(ctx context.Context) DdosPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyArrayOutput)
}

// DdosPolicyMapInput is an input type that accepts DdosPolicyMap and DdosPolicyMapOutput values.
// You can construct a concrete instance of `DdosPolicyMapInput` via:
//
//          DdosPolicyMap{ "key": DdosPolicyArgs{...} }
type DdosPolicyMapInput interface {
	pulumi.Input

	ToDdosPolicyMapOutput() DdosPolicyMapOutput
	ToDdosPolicyMapOutputWithContext(context.Context) DdosPolicyMapOutput
}

type DdosPolicyMap map[string]DdosPolicyInput

func (DdosPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosPolicy)(nil)).Elem()
}

func (i DdosPolicyMap) ToDdosPolicyMapOutput() DdosPolicyMapOutput {
	return i.ToDdosPolicyMapOutputWithContext(context.Background())
}

func (i DdosPolicyMap) ToDdosPolicyMapOutputWithContext(ctx context.Context) DdosPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyMapOutput)
}

type DdosPolicyOutput struct{ *pulumi.OutputState }

func (DdosPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosPolicy)(nil)).Elem()
}

func (o DdosPolicyOutput) ToDdosPolicyOutput() DdosPolicyOutput {
	return o
}

func (o DdosPolicyOutput) ToDdosPolicyOutputWithContext(ctx context.Context) DdosPolicyOutput {
	return o
}

// Black IP list.
func (o DdosPolicyOutput) BlackIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DdosPolicy) pulumi.StringArrayOutput { return v.BlackIps }).(pulumi.StringArrayOutput)
}

// Create time of the DDoS policy.
func (o DdosPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Option list of abnormal check of the DDos policy, should set at least one policy.
func (o DdosPolicyOutput) DropOptions() DdosPolicyDropOptionArrayOutput {
	return o.ApplyT(func(v *DdosPolicy) DdosPolicyDropOptionArrayOutput { return v.DropOptions }).(DdosPolicyDropOptionArrayOutput)
}

// Name of the DDoS policy. Length should between 1 and 32.
func (o DdosPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Message filter options list.
func (o DdosPolicyOutput) PacketFilters() DdosPolicyPacketFilterArrayOutput {
	return o.ApplyT(func(v *DdosPolicy) DdosPolicyPacketFilterArrayOutput { return v.PacketFilters }).(DdosPolicyPacketFilterArrayOutput)
}

// Id of policy.
func (o DdosPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// Port limits of abnormal check of the DDos policy.
func (o DdosPolicyOutput) PortFilters() DdosPolicyPortFilterArrayOutput {
	return o.ApplyT(func(v *DdosPolicy) DdosPolicyPortFilterArrayOutput { return v.PortFilters }).(DdosPolicyPortFilterArrayOutput)
}

// Type of the resource that the DDoS policy works for. Valid values: `bgpip`, `bgp`, `bgp-multip` and `net`.
func (o DdosPolicyOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicy) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// Id of policy case that the DDoS policy works for.
func (o DdosPolicyOutput) SceneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicy) pulumi.StringOutput { return v.SceneId }).(pulumi.StringOutput)
}

// Watermark policy options, and only support one watermark policy at most.
func (o DdosPolicyOutput) WatermarkFilters() DdosPolicyWatermarkFilterArrayOutput {
	return o.ApplyT(func(v *DdosPolicy) DdosPolicyWatermarkFilterArrayOutput { return v.WatermarkFilters }).(DdosPolicyWatermarkFilterArrayOutput)
}

// Watermark content.
func (o DdosPolicyOutput) WatermarkKeys() DdosPolicyWatermarkKeyArrayOutput {
	return o.ApplyT(func(v *DdosPolicy) DdosPolicyWatermarkKeyArrayOutput { return v.WatermarkKeys }).(DdosPolicyWatermarkKeyArrayOutput)
}

// White IP list.
func (o DdosPolicyOutput) WhiteIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DdosPolicy) pulumi.StringArrayOutput { return v.WhiteIps }).(pulumi.StringArrayOutput)
}

type DdosPolicyArrayOutput struct{ *pulumi.OutputState }

func (DdosPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosPolicy)(nil)).Elem()
}

func (o DdosPolicyArrayOutput) ToDdosPolicyArrayOutput() DdosPolicyArrayOutput {
	return o
}

func (o DdosPolicyArrayOutput) ToDdosPolicyArrayOutputWithContext(ctx context.Context) DdosPolicyArrayOutput {
	return o
}

func (o DdosPolicyArrayOutput) Index(i pulumi.IntInput) DdosPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DdosPolicy {
		return vs[0].([]*DdosPolicy)[vs[1].(int)]
	}).(DdosPolicyOutput)
}

type DdosPolicyMapOutput struct{ *pulumi.OutputState }

func (DdosPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosPolicy)(nil)).Elem()
}

func (o DdosPolicyMapOutput) ToDdosPolicyMapOutput() DdosPolicyMapOutput {
	return o
}

func (o DdosPolicyMapOutput) ToDdosPolicyMapOutputWithContext(ctx context.Context) DdosPolicyMapOutput {
	return o
}

func (o DdosPolicyMapOutput) MapIndex(k pulumi.StringInput) DdosPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DdosPolicy {
		return vs[0].(map[string]*DdosPolicy)[vs[1].(string)]
	}).(DdosPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyInput)(nil)).Elem(), &DdosPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyArrayInput)(nil)).Elem(), DdosPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyMapInput)(nil)).Elem(), DdosPolicyMap{})
	pulumi.RegisterOutputType(DdosPolicyOutput{})
	pulumi.RegisterOutputType(DdosPolicyArrayOutput{})
	pulumi.RegisterOutputType(DdosPolicyMapOutput{})
}
