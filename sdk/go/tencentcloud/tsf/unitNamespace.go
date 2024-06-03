// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tsf unitNamespace
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.NewUnitNamespace(ctx, "unitNamespace", &Tsf.UnitNamespaceArgs{
//				GatewayInstanceId: pulumi.String("gw-ins-lvdypq5k"),
//				NamespaceId:       pulumi.String("namespace-vwgo38wy"),
//				NamespaceName:     pulumi.String("keep-terraform-cls"),
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
//
// ## Import
//
// tsf unit_namespace can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tsf/unitNamespace:UnitNamespace unit_namespace gw-ins-lvdypq5k#namespace-vwgo38wy
// ```
type UnitNamespace struct {
	pulumi.CustomResourceState

	// Create time. Note: This field may return null, indicating that no valid value was found.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// gateway instance Id.
	GatewayInstanceId pulumi.StringOutput `pulumi:"gatewayInstanceId"`
	// namespace id.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// namespace name.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Update time. Note: This field may return null, indicating that no valid value was found.
	UpdatedTime pulumi.StringOutput `pulumi:"updatedTime"`
}

// NewUnitNamespace registers a new resource with the given unique name, arguments, and options.
func NewUnitNamespace(ctx *pulumi.Context,
	name string, args *UnitNamespaceArgs, opts ...pulumi.ResourceOption) (*UnitNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayInstanceId'")
	}
	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UnitNamespace
	err := ctx.RegisterResource("tencentcloud:Tsf/unitNamespace:UnitNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUnitNamespace gets an existing UnitNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUnitNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UnitNamespaceState, opts ...pulumi.ResourceOption) (*UnitNamespace, error) {
	var resource UnitNamespace
	err := ctx.ReadResource("tencentcloud:Tsf/unitNamespace:UnitNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UnitNamespace resources.
type unitNamespaceState struct {
	// Create time. Note: This field may return null, indicating that no valid value was found.
	CreatedTime *string `pulumi:"createdTime"`
	// gateway instance Id.
	GatewayInstanceId *string `pulumi:"gatewayInstanceId"`
	// namespace id.
	NamespaceId *string `pulumi:"namespaceId"`
	// namespace name.
	NamespaceName *string `pulumi:"namespaceName"`
	// Update time. Note: This field may return null, indicating that no valid value was found.
	UpdatedTime *string `pulumi:"updatedTime"`
}

type UnitNamespaceState struct {
	// Create time. Note: This field may return null, indicating that no valid value was found.
	CreatedTime pulumi.StringPtrInput
	// gateway instance Id.
	GatewayInstanceId pulumi.StringPtrInput
	// namespace id.
	NamespaceId pulumi.StringPtrInput
	// namespace name.
	NamespaceName pulumi.StringPtrInput
	// Update time. Note: This field may return null, indicating that no valid value was found.
	UpdatedTime pulumi.StringPtrInput
}

func (UnitNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*unitNamespaceState)(nil)).Elem()
}

type unitNamespaceArgs struct {
	// gateway instance Id.
	GatewayInstanceId string `pulumi:"gatewayInstanceId"`
	// namespace id.
	NamespaceId string `pulumi:"namespaceId"`
	// namespace name.
	NamespaceName string `pulumi:"namespaceName"`
}

// The set of arguments for constructing a UnitNamespace resource.
type UnitNamespaceArgs struct {
	// gateway instance Id.
	GatewayInstanceId pulumi.StringInput
	// namespace id.
	NamespaceId pulumi.StringInput
	// namespace name.
	NamespaceName pulumi.StringInput
}

func (UnitNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*unitNamespaceArgs)(nil)).Elem()
}

type UnitNamespaceInput interface {
	pulumi.Input

	ToUnitNamespaceOutput() UnitNamespaceOutput
	ToUnitNamespaceOutputWithContext(ctx context.Context) UnitNamespaceOutput
}

func (*UnitNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**UnitNamespace)(nil)).Elem()
}

func (i *UnitNamespace) ToUnitNamespaceOutput() UnitNamespaceOutput {
	return i.ToUnitNamespaceOutputWithContext(context.Background())
}

func (i *UnitNamespace) ToUnitNamespaceOutputWithContext(ctx context.Context) UnitNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnitNamespaceOutput)
}

// UnitNamespaceArrayInput is an input type that accepts UnitNamespaceArray and UnitNamespaceArrayOutput values.
// You can construct a concrete instance of `UnitNamespaceArrayInput` via:
//
//	UnitNamespaceArray{ UnitNamespaceArgs{...} }
type UnitNamespaceArrayInput interface {
	pulumi.Input

	ToUnitNamespaceArrayOutput() UnitNamespaceArrayOutput
	ToUnitNamespaceArrayOutputWithContext(context.Context) UnitNamespaceArrayOutput
}

type UnitNamespaceArray []UnitNamespaceInput

func (UnitNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UnitNamespace)(nil)).Elem()
}

func (i UnitNamespaceArray) ToUnitNamespaceArrayOutput() UnitNamespaceArrayOutput {
	return i.ToUnitNamespaceArrayOutputWithContext(context.Background())
}

func (i UnitNamespaceArray) ToUnitNamespaceArrayOutputWithContext(ctx context.Context) UnitNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnitNamespaceArrayOutput)
}

// UnitNamespaceMapInput is an input type that accepts UnitNamespaceMap and UnitNamespaceMapOutput values.
// You can construct a concrete instance of `UnitNamespaceMapInput` via:
//
//	UnitNamespaceMap{ "key": UnitNamespaceArgs{...} }
type UnitNamespaceMapInput interface {
	pulumi.Input

	ToUnitNamespaceMapOutput() UnitNamespaceMapOutput
	ToUnitNamespaceMapOutputWithContext(context.Context) UnitNamespaceMapOutput
}

type UnitNamespaceMap map[string]UnitNamespaceInput

func (UnitNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UnitNamespace)(nil)).Elem()
}

func (i UnitNamespaceMap) ToUnitNamespaceMapOutput() UnitNamespaceMapOutput {
	return i.ToUnitNamespaceMapOutputWithContext(context.Background())
}

func (i UnitNamespaceMap) ToUnitNamespaceMapOutputWithContext(ctx context.Context) UnitNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UnitNamespaceMapOutput)
}

type UnitNamespaceOutput struct{ *pulumi.OutputState }

func (UnitNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnitNamespace)(nil)).Elem()
}

func (o UnitNamespaceOutput) ToUnitNamespaceOutput() UnitNamespaceOutput {
	return o
}

func (o UnitNamespaceOutput) ToUnitNamespaceOutputWithContext(ctx context.Context) UnitNamespaceOutput {
	return o
}

// Create time. Note: This field may return null, indicating that no valid value was found.
func (o UnitNamespaceOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UnitNamespace) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// gateway instance Id.
func (o UnitNamespaceOutput) GatewayInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UnitNamespace) pulumi.StringOutput { return v.GatewayInstanceId }).(pulumi.StringOutput)
}

// namespace id.
func (o UnitNamespaceOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UnitNamespace) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// namespace name.
func (o UnitNamespaceOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v *UnitNamespace) pulumi.StringOutput { return v.NamespaceName }).(pulumi.StringOutput)
}

// Update time. Note: This field may return null, indicating that no valid value was found.
func (o UnitNamespaceOutput) UpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *UnitNamespace) pulumi.StringOutput { return v.UpdatedTime }).(pulumi.StringOutput)
}

type UnitNamespaceArrayOutput struct{ *pulumi.OutputState }

func (UnitNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UnitNamespace)(nil)).Elem()
}

func (o UnitNamespaceArrayOutput) ToUnitNamespaceArrayOutput() UnitNamespaceArrayOutput {
	return o
}

func (o UnitNamespaceArrayOutput) ToUnitNamespaceArrayOutputWithContext(ctx context.Context) UnitNamespaceArrayOutput {
	return o
}

func (o UnitNamespaceArrayOutput) Index(i pulumi.IntInput) UnitNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UnitNamespace {
		return vs[0].([]*UnitNamespace)[vs[1].(int)]
	}).(UnitNamespaceOutput)
}

type UnitNamespaceMapOutput struct{ *pulumi.OutputState }

func (UnitNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UnitNamespace)(nil)).Elem()
}

func (o UnitNamespaceMapOutput) ToUnitNamespaceMapOutput() UnitNamespaceMapOutput {
	return o
}

func (o UnitNamespaceMapOutput) ToUnitNamespaceMapOutputWithContext(ctx context.Context) UnitNamespaceMapOutput {
	return o
}

func (o UnitNamespaceMapOutput) MapIndex(k pulumi.StringInput) UnitNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UnitNamespace {
		return vs[0].(map[string]*UnitNamespace)[vs[1].(string)]
	}).(UnitNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UnitNamespaceInput)(nil)).Elem(), &UnitNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*UnitNamespaceArrayInput)(nil)).Elem(), UnitNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UnitNamespaceMapInput)(nil)).Elem(), UnitNamespaceMap{})
	pulumi.RegisterOutputType(UnitNamespaceOutput{})
	pulumi.RegisterOutputType(UnitNamespaceArrayOutput{})
	pulumi.RegisterOutputType(UnitNamespaceMapOutput{})
}
