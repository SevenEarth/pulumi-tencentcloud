// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a waf ipAccessControl
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Waf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Waf.NewIpAccessControl(ctx, "example", &Waf.IpAccessControlArgs{
//				Domain:     pulumi.String("www.demo.com"),
//				Edition:    pulumi.String("sparta-waf"),
//				InstanceId: pulumi.String("waf_2kxtlbky00b3b4qz"),
//				Items: waf.IpAccessControlItemArray{
//					&waf.IpAccessControlItemArgs{
//						Action:  pulumi.Int(40),
//						Ip:      pulumi.String("1.1.1.1"),
//						Note:    pulumi.String("desc info."),
//						ValidTs: pulumi.Int(2019571199),
//					},
//					&waf.IpAccessControlItemArgs{
//						Action:  pulumi.Int(42),
//						Ip:      pulumi.String("2.2.2.2"),
//						Note:    pulumi.String("desc info."),
//						ValidTs: pulumi.Int(2019571199),
//					},
//					&waf.IpAccessControlItemArgs{
//						Action:  pulumi.Int(40),
//						Ip:      pulumi.String("3.3.3.3"),
//						Note:    pulumi.String("desc info."),
//						ValidTs: pulumi.Int(1680570420),
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
// waf ip_access_control can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Waf/ipAccessControl:IpAccessControl example waf_2kxtlbky00b3b4qz#www.demo.com#sparta-waf
//
// ```
type IpAccessControl struct {
	pulumi.CustomResourceState

	// Domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
	Edition pulumi.StringOutput `pulumi:"edition"`
	// Waf instance Id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Ip parameter list.
	Items IpAccessControlItemArrayOutput `pulumi:"items"`
}

// NewIpAccessControl registers a new resource with the given unique name, arguments, and options.
func NewIpAccessControl(ctx *pulumi.Context,
	name string, args *IpAccessControlArgs, opts ...pulumi.ResourceOption) (*IpAccessControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Edition == nil {
		return nil, errors.New("invalid value for required argument 'Edition'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpAccessControl
	err := ctx.RegisterResource("tencentcloud:Waf/ipAccessControl:IpAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpAccessControl gets an existing IpAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpAccessControlState, opts ...pulumi.ResourceOption) (*IpAccessControl, error) {
	var resource IpAccessControl
	err := ctx.ReadResource("tencentcloud:Waf/ipAccessControl:IpAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpAccessControl resources.
type ipAccessControlState struct {
	// Domain.
	Domain *string `pulumi:"domain"`
	// Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
	Edition *string `pulumi:"edition"`
	// Waf instance Id.
	InstanceId *string `pulumi:"instanceId"`
	// Ip parameter list.
	Items []IpAccessControlItem `pulumi:"items"`
}

type IpAccessControlState struct {
	// Domain.
	Domain pulumi.StringPtrInput
	// Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
	Edition pulumi.StringPtrInput
	// Waf instance Id.
	InstanceId pulumi.StringPtrInput
	// Ip parameter list.
	Items IpAccessControlItemArrayInput
}

func (IpAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAccessControlState)(nil)).Elem()
}

type ipAccessControlArgs struct {
	// Domain.
	Domain string `pulumi:"domain"`
	// Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
	Edition string `pulumi:"edition"`
	// Waf instance Id.
	InstanceId string `pulumi:"instanceId"`
	// Ip parameter list.
	Items []IpAccessControlItem `pulumi:"items"`
}

// The set of arguments for constructing a IpAccessControl resource.
type IpAccessControlArgs struct {
	// Domain.
	Domain pulumi.StringInput
	// Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
	Edition pulumi.StringInput
	// Waf instance Id.
	InstanceId pulumi.StringInput
	// Ip parameter list.
	Items IpAccessControlItemArrayInput
}

func (IpAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipAccessControlArgs)(nil)).Elem()
}

type IpAccessControlInput interface {
	pulumi.Input

	ToIpAccessControlOutput() IpAccessControlOutput
	ToIpAccessControlOutputWithContext(ctx context.Context) IpAccessControlOutput
}

func (*IpAccessControl) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAccessControl)(nil)).Elem()
}

func (i *IpAccessControl) ToIpAccessControlOutput() IpAccessControlOutput {
	return i.ToIpAccessControlOutputWithContext(context.Background())
}

func (i *IpAccessControl) ToIpAccessControlOutputWithContext(ctx context.Context) IpAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAccessControlOutput)
}

// IpAccessControlArrayInput is an input type that accepts IpAccessControlArray and IpAccessControlArrayOutput values.
// You can construct a concrete instance of `IpAccessControlArrayInput` via:
//
//	IpAccessControlArray{ IpAccessControlArgs{...} }
type IpAccessControlArrayInput interface {
	pulumi.Input

	ToIpAccessControlArrayOutput() IpAccessControlArrayOutput
	ToIpAccessControlArrayOutputWithContext(context.Context) IpAccessControlArrayOutput
}

type IpAccessControlArray []IpAccessControlInput

func (IpAccessControlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpAccessControl)(nil)).Elem()
}

func (i IpAccessControlArray) ToIpAccessControlArrayOutput() IpAccessControlArrayOutput {
	return i.ToIpAccessControlArrayOutputWithContext(context.Background())
}

func (i IpAccessControlArray) ToIpAccessControlArrayOutputWithContext(ctx context.Context) IpAccessControlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAccessControlArrayOutput)
}

// IpAccessControlMapInput is an input type that accepts IpAccessControlMap and IpAccessControlMapOutput values.
// You can construct a concrete instance of `IpAccessControlMapInput` via:
//
//	IpAccessControlMap{ "key": IpAccessControlArgs{...} }
type IpAccessControlMapInput interface {
	pulumi.Input

	ToIpAccessControlMapOutput() IpAccessControlMapOutput
	ToIpAccessControlMapOutputWithContext(context.Context) IpAccessControlMapOutput
}

type IpAccessControlMap map[string]IpAccessControlInput

func (IpAccessControlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpAccessControl)(nil)).Elem()
}

func (i IpAccessControlMap) ToIpAccessControlMapOutput() IpAccessControlMapOutput {
	return i.ToIpAccessControlMapOutputWithContext(context.Background())
}

func (i IpAccessControlMap) ToIpAccessControlMapOutputWithContext(ctx context.Context) IpAccessControlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAccessControlMapOutput)
}

type IpAccessControlOutput struct{ *pulumi.OutputState }

func (IpAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpAccessControl)(nil)).Elem()
}

func (o IpAccessControlOutput) ToIpAccessControlOutput() IpAccessControlOutput {
	return o
}

func (o IpAccessControlOutput) ToIpAccessControlOutputWithContext(ctx context.Context) IpAccessControlOutput {
	return o
}

// Domain.
func (o IpAccessControlOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAccessControl) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
func (o IpAccessControlOutput) Edition() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAccessControl) pulumi.StringOutput { return v.Edition }).(pulumi.StringOutput)
}

// Waf instance Id.
func (o IpAccessControlOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpAccessControl) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Ip parameter list.
func (o IpAccessControlOutput) Items() IpAccessControlItemArrayOutput {
	return o.ApplyT(func(v *IpAccessControl) IpAccessControlItemArrayOutput { return v.Items }).(IpAccessControlItemArrayOutput)
}

type IpAccessControlArrayOutput struct{ *pulumi.OutputState }

func (IpAccessControlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpAccessControl)(nil)).Elem()
}

func (o IpAccessControlArrayOutput) ToIpAccessControlArrayOutput() IpAccessControlArrayOutput {
	return o
}

func (o IpAccessControlArrayOutput) ToIpAccessControlArrayOutputWithContext(ctx context.Context) IpAccessControlArrayOutput {
	return o
}

func (o IpAccessControlArrayOutput) Index(i pulumi.IntInput) IpAccessControlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpAccessControl {
		return vs[0].([]*IpAccessControl)[vs[1].(int)]
	}).(IpAccessControlOutput)
}

type IpAccessControlMapOutput struct{ *pulumi.OutputState }

func (IpAccessControlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpAccessControl)(nil)).Elem()
}

func (o IpAccessControlMapOutput) ToIpAccessControlMapOutput() IpAccessControlMapOutput {
	return o
}

func (o IpAccessControlMapOutput) ToIpAccessControlMapOutputWithContext(ctx context.Context) IpAccessControlMapOutput {
	return o
}

func (o IpAccessControlMapOutput) MapIndex(k pulumi.StringInput) IpAccessControlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpAccessControl {
		return vs[0].(map[string]*IpAccessControl)[vs[1].(string)]
	}).(IpAccessControlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpAccessControlInput)(nil)).Elem(), &IpAccessControl{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAccessControlArrayInput)(nil)).Elem(), IpAccessControlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpAccessControlMapInput)(nil)).Elem(), IpAccessControlMap{})
	pulumi.RegisterOutputType(IpAccessControlOutput{})
	pulumi.RegisterOutputType(IpAccessControlArrayOutput{})
	pulumi.RegisterOutputType(IpAccessControlMapOutput{})
}
