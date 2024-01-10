// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a redis upgradeCacheVersionOperation
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
//			_, err := Redis.NewUpgradeCacheVersionOperation(ctx, "upgradeCacheVersionOperation", &Redis.UpgradeCacheVersionOperationArgs{
//				CurrentRedisVersion:    pulumi.String("5.0.0"),
//				InstanceId:             pulumi.String("crs-c1nl9rpv"),
//				InstanceTypeUpgradeNow: pulumi.Int(1),
//				UpgradeRedisVersion:    pulumi.String("5.0.0"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type UpgradeCacheVersionOperation struct {
	pulumi.CustomResourceState

	// Current redis version.
	CurrentRedisVersion pulumi.StringOutput `pulumi:"currentRedisVersion"`
	// The ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
	InstanceTypeUpgradeNow pulumi.IntOutput `pulumi:"instanceTypeUpgradeNow"`
	// Upgradeable redis version.
	UpgradeRedisVersion pulumi.StringOutput `pulumi:"upgradeRedisVersion"`
}

// NewUpgradeCacheVersionOperation registers a new resource with the given unique name, arguments, and options.
func NewUpgradeCacheVersionOperation(ctx *pulumi.Context,
	name string, args *UpgradeCacheVersionOperationArgs, opts ...pulumi.ResourceOption) (*UpgradeCacheVersionOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CurrentRedisVersion == nil {
		return nil, errors.New("invalid value for required argument 'CurrentRedisVersion'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceTypeUpgradeNow == nil {
		return nil, errors.New("invalid value for required argument 'InstanceTypeUpgradeNow'")
	}
	if args.UpgradeRedisVersion == nil {
		return nil, errors.New("invalid value for required argument 'UpgradeRedisVersion'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UpgradeCacheVersionOperation
	err := ctx.RegisterResource("tencentcloud:Redis/upgradeCacheVersionOperation:UpgradeCacheVersionOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpgradeCacheVersionOperation gets an existing UpgradeCacheVersionOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpgradeCacheVersionOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpgradeCacheVersionOperationState, opts ...pulumi.ResourceOption) (*UpgradeCacheVersionOperation, error) {
	var resource UpgradeCacheVersionOperation
	err := ctx.ReadResource("tencentcloud:Redis/upgradeCacheVersionOperation:UpgradeCacheVersionOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpgradeCacheVersionOperation resources.
type upgradeCacheVersionOperationState struct {
	// Current redis version.
	CurrentRedisVersion *string `pulumi:"currentRedisVersion"`
	// The ID of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
	InstanceTypeUpgradeNow *int `pulumi:"instanceTypeUpgradeNow"`
	// Upgradeable redis version.
	UpgradeRedisVersion *string `pulumi:"upgradeRedisVersion"`
}

type UpgradeCacheVersionOperationState struct {
	// Current redis version.
	CurrentRedisVersion pulumi.StringPtrInput
	// The ID of instance.
	InstanceId pulumi.StringPtrInput
	// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
	InstanceTypeUpgradeNow pulumi.IntPtrInput
	// Upgradeable redis version.
	UpgradeRedisVersion pulumi.StringPtrInput
}

func (UpgradeCacheVersionOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeCacheVersionOperationState)(nil)).Elem()
}

type upgradeCacheVersionOperationArgs struct {
	// Current redis version.
	CurrentRedisVersion string `pulumi:"currentRedisVersion"`
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
	InstanceTypeUpgradeNow int `pulumi:"instanceTypeUpgradeNow"`
	// Upgradeable redis version.
	UpgradeRedisVersion string `pulumi:"upgradeRedisVersion"`
}

// The set of arguments for constructing a UpgradeCacheVersionOperation resource.
type UpgradeCacheVersionOperationArgs struct {
	// Current redis version.
	CurrentRedisVersion pulumi.StringInput
	// The ID of instance.
	InstanceId pulumi.StringInput
	// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
	InstanceTypeUpgradeNow pulumi.IntInput
	// Upgradeable redis version.
	UpgradeRedisVersion pulumi.StringInput
}

func (UpgradeCacheVersionOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*upgradeCacheVersionOperationArgs)(nil)).Elem()
}

type UpgradeCacheVersionOperationInput interface {
	pulumi.Input

	ToUpgradeCacheVersionOperationOutput() UpgradeCacheVersionOperationOutput
	ToUpgradeCacheVersionOperationOutputWithContext(ctx context.Context) UpgradeCacheVersionOperationOutput
}

func (*UpgradeCacheVersionOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeCacheVersionOperation)(nil)).Elem()
}

func (i *UpgradeCacheVersionOperation) ToUpgradeCacheVersionOperationOutput() UpgradeCacheVersionOperationOutput {
	return i.ToUpgradeCacheVersionOperationOutputWithContext(context.Background())
}

func (i *UpgradeCacheVersionOperation) ToUpgradeCacheVersionOperationOutputWithContext(ctx context.Context) UpgradeCacheVersionOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeCacheVersionOperationOutput)
}

// UpgradeCacheVersionOperationArrayInput is an input type that accepts UpgradeCacheVersionOperationArray and UpgradeCacheVersionOperationArrayOutput values.
// You can construct a concrete instance of `UpgradeCacheVersionOperationArrayInput` via:
//
//	UpgradeCacheVersionOperationArray{ UpgradeCacheVersionOperationArgs{...} }
type UpgradeCacheVersionOperationArrayInput interface {
	pulumi.Input

	ToUpgradeCacheVersionOperationArrayOutput() UpgradeCacheVersionOperationArrayOutput
	ToUpgradeCacheVersionOperationArrayOutputWithContext(context.Context) UpgradeCacheVersionOperationArrayOutput
}

type UpgradeCacheVersionOperationArray []UpgradeCacheVersionOperationInput

func (UpgradeCacheVersionOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeCacheVersionOperation)(nil)).Elem()
}

func (i UpgradeCacheVersionOperationArray) ToUpgradeCacheVersionOperationArrayOutput() UpgradeCacheVersionOperationArrayOutput {
	return i.ToUpgradeCacheVersionOperationArrayOutputWithContext(context.Background())
}

func (i UpgradeCacheVersionOperationArray) ToUpgradeCacheVersionOperationArrayOutputWithContext(ctx context.Context) UpgradeCacheVersionOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeCacheVersionOperationArrayOutput)
}

// UpgradeCacheVersionOperationMapInput is an input type that accepts UpgradeCacheVersionOperationMap and UpgradeCacheVersionOperationMapOutput values.
// You can construct a concrete instance of `UpgradeCacheVersionOperationMapInput` via:
//
//	UpgradeCacheVersionOperationMap{ "key": UpgradeCacheVersionOperationArgs{...} }
type UpgradeCacheVersionOperationMapInput interface {
	pulumi.Input

	ToUpgradeCacheVersionOperationMapOutput() UpgradeCacheVersionOperationMapOutput
	ToUpgradeCacheVersionOperationMapOutputWithContext(context.Context) UpgradeCacheVersionOperationMapOutput
}

type UpgradeCacheVersionOperationMap map[string]UpgradeCacheVersionOperationInput

func (UpgradeCacheVersionOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeCacheVersionOperation)(nil)).Elem()
}

func (i UpgradeCacheVersionOperationMap) ToUpgradeCacheVersionOperationMapOutput() UpgradeCacheVersionOperationMapOutput {
	return i.ToUpgradeCacheVersionOperationMapOutputWithContext(context.Background())
}

func (i UpgradeCacheVersionOperationMap) ToUpgradeCacheVersionOperationMapOutputWithContext(ctx context.Context) UpgradeCacheVersionOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeCacheVersionOperationMapOutput)
}

type UpgradeCacheVersionOperationOutput struct{ *pulumi.OutputState }

func (UpgradeCacheVersionOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeCacheVersionOperation)(nil)).Elem()
}

func (o UpgradeCacheVersionOperationOutput) ToUpgradeCacheVersionOperationOutput() UpgradeCacheVersionOperationOutput {
	return o
}

func (o UpgradeCacheVersionOperationOutput) ToUpgradeCacheVersionOperationOutputWithContext(ctx context.Context) UpgradeCacheVersionOperationOutput {
	return o
}

// Current redis version.
func (o UpgradeCacheVersionOperationOutput) CurrentRedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *UpgradeCacheVersionOperation) pulumi.StringOutput { return v.CurrentRedisVersion }).(pulumi.StringOutput)
}

// The ID of instance.
func (o UpgradeCacheVersionOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UpgradeCacheVersionOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Switch mode:1 - Upgrade now0 - Maintenance window upgrade.
func (o UpgradeCacheVersionOperationOutput) InstanceTypeUpgradeNow() pulumi.IntOutput {
	return o.ApplyT(func(v *UpgradeCacheVersionOperation) pulumi.IntOutput { return v.InstanceTypeUpgradeNow }).(pulumi.IntOutput)
}

// Upgradeable redis version.
func (o UpgradeCacheVersionOperationOutput) UpgradeRedisVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *UpgradeCacheVersionOperation) pulumi.StringOutput { return v.UpgradeRedisVersion }).(pulumi.StringOutput)
}

type UpgradeCacheVersionOperationArrayOutput struct{ *pulumi.OutputState }

func (UpgradeCacheVersionOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpgradeCacheVersionOperation)(nil)).Elem()
}

func (o UpgradeCacheVersionOperationArrayOutput) ToUpgradeCacheVersionOperationArrayOutput() UpgradeCacheVersionOperationArrayOutput {
	return o
}

func (o UpgradeCacheVersionOperationArrayOutput) ToUpgradeCacheVersionOperationArrayOutputWithContext(ctx context.Context) UpgradeCacheVersionOperationArrayOutput {
	return o
}

func (o UpgradeCacheVersionOperationArrayOutput) Index(i pulumi.IntInput) UpgradeCacheVersionOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UpgradeCacheVersionOperation {
		return vs[0].([]*UpgradeCacheVersionOperation)[vs[1].(int)]
	}).(UpgradeCacheVersionOperationOutput)
}

type UpgradeCacheVersionOperationMapOutput struct{ *pulumi.OutputState }

func (UpgradeCacheVersionOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpgradeCacheVersionOperation)(nil)).Elem()
}

func (o UpgradeCacheVersionOperationMapOutput) ToUpgradeCacheVersionOperationMapOutput() UpgradeCacheVersionOperationMapOutput {
	return o
}

func (o UpgradeCacheVersionOperationMapOutput) ToUpgradeCacheVersionOperationMapOutputWithContext(ctx context.Context) UpgradeCacheVersionOperationMapOutput {
	return o
}

func (o UpgradeCacheVersionOperationMapOutput) MapIndex(k pulumi.StringInput) UpgradeCacheVersionOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UpgradeCacheVersionOperation {
		return vs[0].(map[string]*UpgradeCacheVersionOperation)[vs[1].(string)]
	}).(UpgradeCacheVersionOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeCacheVersionOperationInput)(nil)).Elem(), &UpgradeCacheVersionOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeCacheVersionOperationArrayInput)(nil)).Elem(), UpgradeCacheVersionOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeCacheVersionOperationMapInput)(nil)).Elem(), UpgradeCacheVersionOperationMap{})
	pulumi.RegisterOutputType(UpgradeCacheVersionOperationOutput{})
	pulumi.RegisterOutputType(UpgradeCacheVersionOperationArrayOutput{})
	pulumi.RegisterOutputType(UpgradeCacheVersionOperationMapOutput{})
}
