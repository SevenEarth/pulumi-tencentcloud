// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create dayu eip rule
//
// ## Example Usage
//
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
//			_, err := Dayu.NewEip(ctx, "test", &Dayu.EipArgs{
//				BindResourceId:     pulumi.String("ins-4m0jvxic"),
//				BindResourceRegion: pulumi.String("hk"),
//				BindResourceType:   pulumi.String("cvm"),
//				Eip:                pulumi.String("162.62.163.50"),
//				ResourceId:         pulumi.String("bgpip-000004xg"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Eip struct {
	pulumi.CustomResourceState

	// Resource id to bind.
	BindResourceId pulumi.StringOutput `pulumi:"bindResourceId"`
	// Resource region to bind.
	BindResourceRegion pulumi.StringOutput `pulumi:"bindResourceRegion"`
	// Resource type to bind, value range [`clb`, `cvm`].
	BindResourceType pulumi.StringOutput `pulumi:"bindResourceType"`
	// Created time of the resource instance.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// Eip of the resource.
	Eip pulumi.StringOutput `pulumi:"eip"`
	// Eip address status of the resource instance.
	EipAddressStatus pulumi.StringOutput `pulumi:"eipAddressStatus"`
	// Eip bound rsc eni of the resource instance.
	EipBoundRscEni pulumi.StringOutput `pulumi:"eipBoundRscEni"`
	// Eip bound rsc ins of the resource instance.
	EipBoundRscIns pulumi.StringOutput `pulumi:"eipBoundRscIns"`
	// Eip bound rsc vip of the resource instance.
	EipBoundRscVip pulumi.StringOutput `pulumi:"eipBoundRscVip"`
	// Expired time of the resource instance.
	ExpiredTime pulumi.StringOutput `pulumi:"expiredTime"`
	// Modify time of the resource instance.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	// Protection status of the resource instance.
	ProtectionStatus pulumi.StringOutput `pulumi:"protectionStatus"`
	// ID of the resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Region of the resource instance.
	ResourceRegion pulumi.StringOutput `pulumi:"resourceRegion"`
}

// NewEip registers a new resource with the given unique name, arguments, and options.
func NewEip(ctx *pulumi.Context,
	name string, args *EipArgs, opts ...pulumi.ResourceOption) (*Eip, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BindResourceId == nil {
		return nil, errors.New("invalid value for required argument 'BindResourceId'")
	}
	if args.BindResourceRegion == nil {
		return nil, errors.New("invalid value for required argument 'BindResourceRegion'")
	}
	if args.BindResourceType == nil {
		return nil, errors.New("invalid value for required argument 'BindResourceType'")
	}
	if args.Eip == nil {
		return nil, errors.New("invalid value for required argument 'Eip'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	var resource Eip
	err := ctx.RegisterResource("tencentcloud:Dayu/eip:Eip", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEip gets an existing Eip resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEip(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipState, opts ...pulumi.ResourceOption) (*Eip, error) {
	var resource Eip
	err := ctx.ReadResource("tencentcloud:Dayu/eip:Eip", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Eip resources.
type eipState struct {
	// Resource id to bind.
	BindResourceId *string `pulumi:"bindResourceId"`
	// Resource region to bind.
	BindResourceRegion *string `pulumi:"bindResourceRegion"`
	// Resource type to bind, value range [`clb`, `cvm`].
	BindResourceType *string `pulumi:"bindResourceType"`
	// Created time of the resource instance.
	CreatedTime *string `pulumi:"createdTime"`
	// Eip of the resource.
	Eip *string `pulumi:"eip"`
	// Eip address status of the resource instance.
	EipAddressStatus *string `pulumi:"eipAddressStatus"`
	// Eip bound rsc eni of the resource instance.
	EipBoundRscEni *string `pulumi:"eipBoundRscEni"`
	// Eip bound rsc ins of the resource instance.
	EipBoundRscIns *string `pulumi:"eipBoundRscIns"`
	// Eip bound rsc vip of the resource instance.
	EipBoundRscVip *string `pulumi:"eipBoundRscVip"`
	// Expired time of the resource instance.
	ExpiredTime *string `pulumi:"expiredTime"`
	// Modify time of the resource instance.
	ModifyTime *string `pulumi:"modifyTime"`
	// Protection status of the resource instance.
	ProtectionStatus *string `pulumi:"protectionStatus"`
	// ID of the resource.
	ResourceId *string `pulumi:"resourceId"`
	// Region of the resource instance.
	ResourceRegion *string `pulumi:"resourceRegion"`
}

type EipState struct {
	// Resource id to bind.
	BindResourceId pulumi.StringPtrInput
	// Resource region to bind.
	BindResourceRegion pulumi.StringPtrInput
	// Resource type to bind, value range [`clb`, `cvm`].
	BindResourceType pulumi.StringPtrInput
	// Created time of the resource instance.
	CreatedTime pulumi.StringPtrInput
	// Eip of the resource.
	Eip pulumi.StringPtrInput
	// Eip address status of the resource instance.
	EipAddressStatus pulumi.StringPtrInput
	// Eip bound rsc eni of the resource instance.
	EipBoundRscEni pulumi.StringPtrInput
	// Eip bound rsc ins of the resource instance.
	EipBoundRscIns pulumi.StringPtrInput
	// Eip bound rsc vip of the resource instance.
	EipBoundRscVip pulumi.StringPtrInput
	// Expired time of the resource instance.
	ExpiredTime pulumi.StringPtrInput
	// Modify time of the resource instance.
	ModifyTime pulumi.StringPtrInput
	// Protection status of the resource instance.
	ProtectionStatus pulumi.StringPtrInput
	// ID of the resource.
	ResourceId pulumi.StringPtrInput
	// Region of the resource instance.
	ResourceRegion pulumi.StringPtrInput
}

func (EipState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipState)(nil)).Elem()
}

type eipArgs struct {
	// Resource id to bind.
	BindResourceId string `pulumi:"bindResourceId"`
	// Resource region to bind.
	BindResourceRegion string `pulumi:"bindResourceRegion"`
	// Resource type to bind, value range [`clb`, `cvm`].
	BindResourceType string `pulumi:"bindResourceType"`
	// Eip of the resource.
	Eip string `pulumi:"eip"`
	// ID of the resource.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a Eip resource.
type EipArgs struct {
	// Resource id to bind.
	BindResourceId pulumi.StringInput
	// Resource region to bind.
	BindResourceRegion pulumi.StringInput
	// Resource type to bind, value range [`clb`, `cvm`].
	BindResourceType pulumi.StringInput
	// Eip of the resource.
	Eip pulumi.StringInput
	// ID of the resource.
	ResourceId pulumi.StringInput
}

func (EipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipArgs)(nil)).Elem()
}

type EipInput interface {
	pulumi.Input

	ToEipOutput() EipOutput
	ToEipOutputWithContext(ctx context.Context) EipOutput
}

func (*Eip) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (i *Eip) ToEipOutput() EipOutput {
	return i.ToEipOutputWithContext(context.Background())
}

func (i *Eip) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipOutput)
}

// EipArrayInput is an input type that accepts EipArray and EipArrayOutput values.
// You can construct a concrete instance of `EipArrayInput` via:
//
//	EipArray{ EipArgs{...} }
type EipArrayInput interface {
	pulumi.Input

	ToEipArrayOutput() EipArrayOutput
	ToEipArrayOutputWithContext(context.Context) EipArrayOutput
}

type EipArray []EipInput

func (EipArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (i EipArray) ToEipArrayOutput() EipArrayOutput {
	return i.ToEipArrayOutputWithContext(context.Background())
}

func (i EipArray) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipArrayOutput)
}

// EipMapInput is an input type that accepts EipMap and EipMapOutput values.
// You can construct a concrete instance of `EipMapInput` via:
//
//	EipMap{ "key": EipArgs{...} }
type EipMapInput interface {
	pulumi.Input

	ToEipMapOutput() EipMapOutput
	ToEipMapOutputWithContext(context.Context) EipMapOutput
}

type EipMap map[string]EipInput

func (EipMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (i EipMap) ToEipMapOutput() EipMapOutput {
	return i.ToEipMapOutputWithContext(context.Background())
}

func (i EipMap) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EipMapOutput)
}

type EipOutput struct{ *pulumi.OutputState }

func (EipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Eip)(nil)).Elem()
}

func (o EipOutput) ToEipOutput() EipOutput {
	return o
}

func (o EipOutput) ToEipOutputWithContext(ctx context.Context) EipOutput {
	return o
}

// Resource id to bind.
func (o EipOutput) BindResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.BindResourceId }).(pulumi.StringOutput)
}

// Resource region to bind.
func (o EipOutput) BindResourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.BindResourceRegion }).(pulumi.StringOutput)
}

// Resource type to bind, value range [`clb`, `cvm`].
func (o EipOutput) BindResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.BindResourceType }).(pulumi.StringOutput)
}

// Created time of the resource instance.
func (o EipOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// Eip of the resource.
func (o EipOutput) Eip() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.Eip }).(pulumi.StringOutput)
}

// Eip address status of the resource instance.
func (o EipOutput) EipAddressStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.EipAddressStatus }).(pulumi.StringOutput)
}

// Eip bound rsc eni of the resource instance.
func (o EipOutput) EipBoundRscEni() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.EipBoundRscEni }).(pulumi.StringOutput)
}

// Eip bound rsc ins of the resource instance.
func (o EipOutput) EipBoundRscIns() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.EipBoundRscIns }).(pulumi.StringOutput)
}

// Eip bound rsc vip of the resource instance.
func (o EipOutput) EipBoundRscVip() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.EipBoundRscVip }).(pulumi.StringOutput)
}

// Expired time of the resource instance.
func (o EipOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.ExpiredTime }).(pulumi.StringOutput)
}

// Modify time of the resource instance.
func (o EipOutput) ModifyTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.ModifyTime }).(pulumi.StringOutput)
}

// Protection status of the resource instance.
func (o EipOutput) ProtectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.ProtectionStatus }).(pulumi.StringOutput)
}

// ID of the resource.
func (o EipOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Region of the resource instance.
func (o EipOutput) ResourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *Eip) pulumi.StringOutput { return v.ResourceRegion }).(pulumi.StringOutput)
}

type EipArrayOutput struct{ *pulumi.OutputState }

func (EipArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eip)(nil)).Elem()
}

func (o EipArrayOutput) ToEipArrayOutput() EipArrayOutput {
	return o
}

func (o EipArrayOutput) ToEipArrayOutputWithContext(ctx context.Context) EipArrayOutput {
	return o
}

func (o EipArrayOutput) Index(i pulumi.IntInput) EipOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].([]*Eip)[vs[1].(int)]
	}).(EipOutput)
}

type EipMapOutput struct{ *pulumi.OutputState }

func (EipMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eip)(nil)).Elem()
}

func (o EipMapOutput) ToEipMapOutput() EipMapOutput {
	return o
}

func (o EipMapOutput) ToEipMapOutputWithContext(ctx context.Context) EipMapOutput {
	return o
}

func (o EipMapOutput) MapIndex(k pulumi.StringInput) EipOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Eip {
		return vs[0].(map[string]*Eip)[vs[1].(string)]
	}).(EipOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EipInput)(nil)).Elem(), &Eip{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipArrayInput)(nil)).Elem(), EipArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EipMapInput)(nil)).Elem(), EipMap{})
	pulumi.RegisterOutputType(EipOutput{})
	pulumi.RegisterOutputType(EipArrayOutput{})
	pulumi.RegisterOutputType(EipMapOutput{})
}
