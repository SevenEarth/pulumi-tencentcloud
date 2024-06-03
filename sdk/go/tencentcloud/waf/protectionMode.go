// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a waf protectionMode
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Waf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Waf.NewProtectionMode(ctx, "example", &Waf.ProtectionModeArgs{
//				Domain:  pulumi.String("keep.qcloudwaf.com"),
//				Edition: pulumi.String("sparta-waf"),
//				Mode:    pulumi.Int(10),
//				Type:    pulumi.Int(0),
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
type ProtectionMode struct {
	pulumi.CustomResourceState

	// Domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
	Edition pulumi.StringPtrOutput `pulumi:"edition"`
	// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception; AI interception mode.
	Mode pulumi.IntOutput `pulumi:"mode"`
	// 0 is to modify the rule engine status, 1 is to modify the AI status.
	Type pulumi.IntPtrOutput `pulumi:"type"`
}

// NewProtectionMode registers a new resource with the given unique name, arguments, and options.
func NewProtectionMode(ctx *pulumi.Context,
	name string, args *ProtectionModeArgs, opts ...pulumi.ResourceOption) (*ProtectionMode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProtectionMode
	err := ctx.RegisterResource("tencentcloud:Waf/protectionMode:ProtectionMode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProtectionMode gets an existing ProtectionMode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProtectionMode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProtectionModeState, opts ...pulumi.ResourceOption) (*ProtectionMode, error) {
	var resource ProtectionMode
	err := ctx.ReadResource("tencentcloud:Waf/protectionMode:ProtectionMode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProtectionMode resources.
type protectionModeState struct {
	// Domain.
	Domain *string `pulumi:"domain"`
	// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
	Edition *string `pulumi:"edition"`
	// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception; AI interception mode.
	Mode *int `pulumi:"mode"`
	// 0 is to modify the rule engine status, 1 is to modify the AI status.
	Type *int `pulumi:"type"`
}

type ProtectionModeState struct {
	// Domain.
	Domain pulumi.StringPtrInput
	// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
	Edition pulumi.StringPtrInput
	// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception; AI interception mode.
	Mode pulumi.IntPtrInput
	// 0 is to modify the rule engine status, 1 is to modify the AI status.
	Type pulumi.IntPtrInput
}

func (ProtectionModeState) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionModeState)(nil)).Elem()
}

type protectionModeArgs struct {
	// Domain.
	Domain string `pulumi:"domain"`
	// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
	Edition *string `pulumi:"edition"`
	// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception; AI interception mode.
	Mode int `pulumi:"mode"`
	// 0 is to modify the rule engine status, 1 is to modify the AI status.
	Type *int `pulumi:"type"`
}

// The set of arguments for constructing a ProtectionMode resource.
type ProtectionModeArgs struct {
	// Domain.
	Domain pulumi.StringInput
	// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
	Edition pulumi.StringPtrInput
	// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception; AI interception mode.
	Mode pulumi.IntInput
	// 0 is to modify the rule engine status, 1 is to modify the AI status.
	Type pulumi.IntPtrInput
}

func (ProtectionModeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*protectionModeArgs)(nil)).Elem()
}

type ProtectionModeInput interface {
	pulumi.Input

	ToProtectionModeOutput() ProtectionModeOutput
	ToProtectionModeOutputWithContext(ctx context.Context) ProtectionModeOutput
}

func (*ProtectionMode) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionMode)(nil)).Elem()
}

func (i *ProtectionMode) ToProtectionModeOutput() ProtectionModeOutput {
	return i.ToProtectionModeOutputWithContext(context.Background())
}

func (i *ProtectionMode) ToProtectionModeOutputWithContext(ctx context.Context) ProtectionModeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeOutput)
}

// ProtectionModeArrayInput is an input type that accepts ProtectionModeArray and ProtectionModeArrayOutput values.
// You can construct a concrete instance of `ProtectionModeArrayInput` via:
//
//	ProtectionModeArray{ ProtectionModeArgs{...} }
type ProtectionModeArrayInput interface {
	pulumi.Input

	ToProtectionModeArrayOutput() ProtectionModeArrayOutput
	ToProtectionModeArrayOutputWithContext(context.Context) ProtectionModeArrayOutput
}

type ProtectionModeArray []ProtectionModeInput

func (ProtectionModeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProtectionMode)(nil)).Elem()
}

func (i ProtectionModeArray) ToProtectionModeArrayOutput() ProtectionModeArrayOutput {
	return i.ToProtectionModeArrayOutputWithContext(context.Background())
}

func (i ProtectionModeArray) ToProtectionModeArrayOutputWithContext(ctx context.Context) ProtectionModeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeArrayOutput)
}

// ProtectionModeMapInput is an input type that accepts ProtectionModeMap and ProtectionModeMapOutput values.
// You can construct a concrete instance of `ProtectionModeMapInput` via:
//
//	ProtectionModeMap{ "key": ProtectionModeArgs{...} }
type ProtectionModeMapInput interface {
	pulumi.Input

	ToProtectionModeMapOutput() ProtectionModeMapOutput
	ToProtectionModeMapOutputWithContext(context.Context) ProtectionModeMapOutput
}

type ProtectionModeMap map[string]ProtectionModeInput

func (ProtectionModeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProtectionMode)(nil)).Elem()
}

func (i ProtectionModeMap) ToProtectionModeMapOutput() ProtectionModeMapOutput {
	return i.ToProtectionModeMapOutputWithContext(context.Background())
}

func (i ProtectionModeMap) ToProtectionModeMapOutputWithContext(ctx context.Context) ProtectionModeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionModeMapOutput)
}

type ProtectionModeOutput struct{ *pulumi.OutputState }

func (ProtectionModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionMode)(nil)).Elem()
}

func (o ProtectionModeOutput) ToProtectionModeOutput() ProtectionModeOutput {
	return o
}

func (o ProtectionModeOutput) ToProtectionModeOutputWithContext(ctx context.Context) ProtectionModeOutput {
	return o
}

// Domain.
func (o ProtectionModeOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *ProtectionMode) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf, default is sparta-waf.
func (o ProtectionModeOutput) Edition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) pulumi.StringPtrOutput { return v.Edition }).(pulumi.StringPtrOutput)
}

// Protection status:10: Rule observation; AI off mode, 11: Rule observation; AI observation mode, 12: Rule observation; AI interception mode20: Rule interception; AI off mode, 21: Rule interception; AI observation mode, 22: Rule interception; AI interception mode.
func (o ProtectionModeOutput) Mode() pulumi.IntOutput {
	return o.ApplyT(func(v *ProtectionMode) pulumi.IntOutput { return v.Mode }).(pulumi.IntOutput)
}

// 0 is to modify the rule engine status, 1 is to modify the AI status.
func (o ProtectionModeOutput) Type() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProtectionMode) pulumi.IntPtrOutput { return v.Type }).(pulumi.IntPtrOutput)
}

type ProtectionModeArrayOutput struct{ *pulumi.OutputState }

func (ProtectionModeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProtectionMode)(nil)).Elem()
}

func (o ProtectionModeArrayOutput) ToProtectionModeArrayOutput() ProtectionModeArrayOutput {
	return o
}

func (o ProtectionModeArrayOutput) ToProtectionModeArrayOutputWithContext(ctx context.Context) ProtectionModeArrayOutput {
	return o
}

func (o ProtectionModeArrayOutput) Index(i pulumi.IntInput) ProtectionModeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProtectionMode {
		return vs[0].([]*ProtectionMode)[vs[1].(int)]
	}).(ProtectionModeOutput)
}

type ProtectionModeMapOutput struct{ *pulumi.OutputState }

func (ProtectionModeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProtectionMode)(nil)).Elem()
}

func (o ProtectionModeMapOutput) ToProtectionModeMapOutput() ProtectionModeMapOutput {
	return o
}

func (o ProtectionModeMapOutput) ToProtectionModeMapOutputWithContext(ctx context.Context) ProtectionModeMapOutput {
	return o
}

func (o ProtectionModeMapOutput) MapIndex(k pulumi.StringInput) ProtectionModeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProtectionMode {
		return vs[0].(map[string]*ProtectionMode)[vs[1].(string)]
	}).(ProtectionModeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionModeInput)(nil)).Elem(), &ProtectionMode{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionModeArrayInput)(nil)).Elem(), ProtectionModeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionModeMapInput)(nil)).Elem(), ProtectionModeMap{})
	pulumi.RegisterOutputType(ProtectionModeOutput{})
	pulumi.RegisterOutputType(ProtectionModeArrayOutput{})
	pulumi.RegisterOutputType(ProtectionModeMapOutput{})
}
