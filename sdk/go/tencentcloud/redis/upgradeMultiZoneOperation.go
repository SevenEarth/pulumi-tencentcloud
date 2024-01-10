// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a redis upgradeMultiZoneOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Redis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Redis.NewUpgradeMultiZoneOperation(ctx, "upgradeMultiZoneOperation", &Redis.UpgradeMultiZoneOperationArgs{
//				InstanceId:                 pulumi.String("crs-c1nl9rpv"),
//				UpgradeProxyAndRedisServer: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type UpgradeMultiZoneOperation struct {
	pulumi.CustomResourceState

	// The ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// After you upgrade Multi-AZ, whether the nearby access feature is supported.true: Supports nearby access.The upgrade process, which requires upgrading both the proxy version and the Redis kernel minor version, involves data migration and can take several hours.false: No need to support nearby access.Upgrading Multi-AZ only involves managing metadata migration, with no service impact, and the upgrade process typically completes within 3 minutes.
	UpgradeProxyAndRedisServer pulumi.BoolPtrOutput `pulumi:"upgradeProxyAndRedisServer"`
}

// NewUpgradeMultiZoneOperation registers a new resource with the given unique name, arguments, and options.
func NewUpgradeMultiZoneOperation(ctx *pulumi.Context,
	name string, args *UpgradeMultiZoneOperationArgs, opts ...pulumi.ResourceOption) (*UpgradeMultiZoneOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UpgradeMultiZoneOperation
	err := ctx.RegisterResource("tencentcloud:Redis/upgradeMultiZoneOperation:UpgradeMultiZoneOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpgradeMultiZoneOperation gets an existing UpgradeMultiZoneOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpgradeMultiZoneOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpgradeMultiZoneOperationState, opts ...pulumi.ResourceOption) (*UpgradeMultiZoneOperation, error) {
	var resource UpgradeMultiZoneOperation
	err := ctx.ReadResource("tencentcloud:Redis/upgradeMultiZoneOperation:UpgradeMultiZoneOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpgradeMultiZoneOperation resources.
type upgradeMultiZoneOperationState struct {
	// The ID of instance.
	InstanceId *string `pulumi:"instanceId"`
	// After you upgrade Multi-AZ, whether the nearby access feature is supported.true: Supports nearby access.The upgrade process, which requires upgrading both the proxy version and the Redis kernel minor version, involves data migration and can take several hours.false: No need to support nearby access.Upgrading Multi-AZ only involves managing metadata migration, with no service impact, and the upgrade process typically completes within 3 minutes.
	UpgradeProxyAndRedisServer *bool `pulumi:"upgradeProxyAndRedisServer"`
}

type UpgradeMultiZoneOperationState struct {
	// The ID of instance.
	InstanceId pulumi.StringPtrInput
	// After you upgrade Multi-AZ, whether the nearby access feature is supported.true: Supports nearby access.The upgrade process, which requires upgrading both the proxy version and the Redis kernel minor version, involves data migration and can take several hours.false: No need to support nearby access.Upgrading Multi-AZ only involves managing metadata migration, with no service impact, and the upgrade process typically completes within 3 minutes.
	UpgradeProxyAndRedisServer pulumi.BoolPtrInput
}

func (UpgradeMultiZoneOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeMultiZoneOperationState)(nil)).Elem()
}

type upgradeMultiZoneOperationArgs struct {
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// After you upgrade Multi-AZ, whether the nearby access feature is supported.true: Supports nearby access.The upgrade process, which requires upgrading both the proxy version and the Redis kernel minor version, involves data migration and can take several hours.false: No need to support nearby access.Upgrading Multi-AZ only involves managing metadata migration, with no service impact, and the upgrade process typically completes within 3 minutes.
	UpgradeProxyAndRedisServer *bool `pulumi:"upgradeProxyAndRedisServer"`
}

// The set of arguments for constructing a UpgradeMultiZoneOperation resource.
type UpgradeMultiZoneOperationArgs struct {
	// The ID of instance.
	InstanceId pulumi.StringInput
	// After you upgrade Multi-AZ, whether the nearby access feature is supported.true: Supports nearby access.The upgrade process, which requires upgrading both the proxy version and the Redis kernel minor version, involves data migration and can take several hours.false: No need to support nearby access.Upgrading Multi-AZ only involves managing metadata migration, with no service impact, and the upgrade process typically completes within 3 minutes.
	UpgradeProxyAndRedisServer pulumi.BoolPtrInput
}

func (UpgradeMultiZoneOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeMultiZoneOperationArgs)(nil)).Elem()
}

type UpgradeMultiZoneOperationInput interface {
	pulumi.Input

	ToUpgradeMultiZoneOperationOutput() UpgradeMultiZoneOperationOutput
	ToUpgradeMultiZoneOperationOutputWithContext(ctx context.Context) UpgradeMultiZoneOperationOutput
}

func (*UpgradeMultiZoneOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeMultiZoneOperation)(nil)).Elem()
}

func (i *UpgradeMultiZoneOperation) ToUpgradeMultiZoneOperationOutput() UpgradeMultiZoneOperationOutput {
	return i.ToUpgradeMultiZoneOperationOutputWithContext(context.Background())
}

func (i *UpgradeMultiZoneOperation) ToUpgradeMultiZoneOperationOutputWithContext(ctx context.Context) UpgradeMultiZoneOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeMultiZoneOperationOutput)
}

// UpgradeMultiZoneOperationArrayInput is an input type that accepts UpgradeMultiZoneOperationArray and UpgradeMultiZoneOperationArrayOutput values.
// You can construct a concrete instance of `UpgradeMultiZoneOperationArrayInput` via:
//
//	UpgradeMultiZoneOperationArray{ UpgradeMultiZoneOperationArgs{...} }
type UpgradeMultiZoneOperationArrayInput interface {
	pulumi.Input

	ToUpgradeMultiZoneOperationArrayOutput() UpgradeMultiZoneOperationArrayOutput
	ToUpgradeMultiZoneOperationArrayOutputWithContext(context.Context) UpgradeMultiZoneOperationArrayOutput
}

type UpgradeMultiZoneOperationArray []UpgradeMultiZoneOperationInput

func (UpgradeMultiZoneOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeMultiZoneOperation)(nil)).Elem()
}

func (i UpgradeMultiZoneOperationArray) ToUpgradeMultiZoneOperationArrayOutput() UpgradeMultiZoneOperationArrayOutput {
	return i.ToUpgradeMultiZoneOperationArrayOutputWithContext(context.Background())
}

func (i UpgradeMultiZoneOperationArray) ToUpgradeMultiZoneOperationArrayOutputWithContext(ctx context.Context) UpgradeMultiZoneOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeMultiZoneOperationArrayOutput)
}

// UpgradeMultiZoneOperationMapInput is an input type that accepts UpgradeMultiZoneOperationMap and UpgradeMultiZoneOperationMapOutput values.
// You can construct a concrete instance of `UpgradeMultiZoneOperationMapInput` via:
//
//	UpgradeMultiZoneOperationMap{ "key": UpgradeMultiZoneOperationArgs{...} }
type UpgradeMultiZoneOperationMapInput interface {
	pulumi.Input

	ToUpgradeMultiZoneOperationMapOutput() UpgradeMultiZoneOperationMapOutput
	ToUpgradeMultiZoneOperationMapOutputWithContext(context.Context) UpgradeMultiZoneOperationMapOutput
}

type UpgradeMultiZoneOperationMap map[string]UpgradeMultiZoneOperationInput

func (UpgradeMultiZoneOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeMultiZoneOperation)(nil)).Elem()
}

func (i UpgradeMultiZoneOperationMap) ToUpgradeMultiZoneOperationMapOutput() UpgradeMultiZoneOperationMapOutput {
	return i.ToUpgradeMultiZoneOperationMapOutputWithContext(context.Background())
}

func (i UpgradeMultiZoneOperationMap) ToUpgradeMultiZoneOperationMapOutputWithContext(ctx context.Context) UpgradeMultiZoneOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeMultiZoneOperationMapOutput)
}

type UpgradeMultiZoneOperationOutput struct{ *pulumi.OutputState }

func (UpgradeMultiZoneOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeMultiZoneOperation)(nil)).Elem()
}

func (o UpgradeMultiZoneOperationOutput) ToUpgradeMultiZoneOperationOutput() UpgradeMultiZoneOperationOutput {
	return o
}

func (o UpgradeMultiZoneOperationOutput) ToUpgradeMultiZoneOperationOutputWithContext(ctx context.Context) UpgradeMultiZoneOperationOutput {
	return o
}

// The ID of instance.
func (o UpgradeMultiZoneOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UpgradeMultiZoneOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// After you upgrade Multi-AZ, whether the nearby access feature is supported.true: Supports nearby access.The upgrade process, which requires upgrading both the proxy version and the Redis kernel minor version, involves data migration and can take several hours.false: No need to support nearby access.Upgrading Multi-AZ only involves managing metadata migration, with no service impact, and the upgrade process typically completes within 3 minutes.
func (o UpgradeMultiZoneOperationOutput) UpgradeProxyAndRedisServer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UpgradeMultiZoneOperation) pulumi.BoolPtrOutput { return v.UpgradeProxyAndRedisServer }).(pulumi.BoolPtrOutput)
}

type UpgradeMultiZoneOperationArrayOutput struct{ *pulumi.OutputState }

func (UpgradeMultiZoneOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeMultiZoneOperation)(nil)).Elem()
}

func (o UpgradeMultiZoneOperationArrayOutput) ToUpgradeMultiZoneOperationArrayOutput() UpgradeMultiZoneOperationArrayOutput {
	return o
}

func (o UpgradeMultiZoneOperationArrayOutput) ToUpgradeMultiZoneOperationArrayOutputWithContext(ctx context.Context) UpgradeMultiZoneOperationArrayOutput {
	return o
}

func (o UpgradeMultiZoneOperationArrayOutput) Index(i pulumi.IntInput) UpgradeMultiZoneOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UpgradeMultiZoneOperation {
		return vs[0].([]*UpgradeMultiZoneOperation)[vs[1].(int)]
	}).(UpgradeMultiZoneOperationOutput)
}

type UpgradeMultiZoneOperationMapOutput struct{ *pulumi.OutputState }

func (UpgradeMultiZoneOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeMultiZoneOperation)(nil)).Elem()
}

func (o UpgradeMultiZoneOperationMapOutput) ToUpgradeMultiZoneOperationMapOutput() UpgradeMultiZoneOperationMapOutput {
	return o
}

func (o UpgradeMultiZoneOperationMapOutput) ToUpgradeMultiZoneOperationMapOutputWithContext(ctx context.Context) UpgradeMultiZoneOperationMapOutput {
	return o
}

func (o UpgradeMultiZoneOperationMapOutput) MapIndex(k pulumi.StringInput) UpgradeMultiZoneOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UpgradeMultiZoneOperation {
		return vs[0].(map[string]*UpgradeMultiZoneOperation)[vs[1].(string)]
	}).(UpgradeMultiZoneOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeMultiZoneOperationInput)(nil)).Elem(), &UpgradeMultiZoneOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeMultiZoneOperationArrayInput)(nil)).Elem(), UpgradeMultiZoneOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeMultiZoneOperationMapInput)(nil)).Elem(), UpgradeMultiZoneOperationMap{})
	pulumi.RegisterOutputType(UpgradeMultiZoneOperationOutput{})
	pulumi.RegisterOutputType(UpgradeMultiZoneOperationArrayOutput{})
	pulumi.RegisterOutputType(UpgradeMultiZoneOperationMapOutput{})
}
