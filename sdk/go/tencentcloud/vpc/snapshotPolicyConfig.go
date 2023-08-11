// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc snapshotPolicyConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleBucket, err := Cos.NewBucket(ctx, "exampleBucket", &Cos.BucketArgs{
// 			Bucket: pulumi.String("tf-example-1308919341"),
// 			Acl:    pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleSnapshotPolicy, err := Vpc.NewSnapshotPolicy(ctx, "exampleSnapshotPolicy", &Vpc.SnapshotPolicyArgs{
// 			SnapshotPolicyName: pulumi.String("tf-example"),
// 			BackupType:         pulumi.String("time"),
// 			CosBucket:          exampleBucket.Bucket,
// 			CosRegion:          pulumi.String("ap-guangzhou"),
// 			CreateNewCos:       pulumi.Bool(false),
// 			KeepTime:           pulumi.Int(2),
// 			BackupPolicies: vpc.SnapshotPolicyBackupPolicyArray{
// 				&vpc.SnapshotPolicyBackupPolicyArgs{
// 					BackupDay:  pulumi.String("monday"),
// 					BackupTime: pulumi.String("00:00:00"),
// 				},
// 				&vpc.SnapshotPolicyBackupPolicyArgs{
// 					BackupDay:  pulumi.String("tuesday"),
// 					BackupTime: pulumi.String("01:00:00"),
// 				},
// 				&vpc.SnapshotPolicyBackupPolicyArgs{
// 					BackupDay:  pulumi.String("wednesday"),
// 					BackupTime: pulumi.String("02:00:00"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Vpc.NewSnapshotPolicyConfig(ctx, "config", &Vpc.SnapshotPolicyConfigArgs{
// 			SnapshotPolicyId: exampleSnapshotPolicy.ID(),
// 			Enable:           pulumi.Bool(false),
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
// vpc snapshot_policy_config can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Vpc/snapshotPolicyConfig:SnapshotPolicyConfig snapshot_policy_config snapshot_policy_id
// ```
type SnapshotPolicyConfig struct {
	pulumi.CustomResourceState

	// If enable snapshot policy.
	Enable pulumi.BoolOutput `pulumi:"enable"`
	// Snapshot policy Id.
	SnapshotPolicyId pulumi.StringOutput `pulumi:"snapshotPolicyId"`
}

// NewSnapshotPolicyConfig registers a new resource with the given unique name, arguments, and options.
func NewSnapshotPolicyConfig(ctx *pulumi.Context,
	name string, args *SnapshotPolicyConfigArgs, opts ...pulumi.ResourceOption) (*SnapshotPolicyConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enable == nil {
		return nil, errors.New("invalid value for required argument 'Enable'")
	}
	if args.SnapshotPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotPolicyId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SnapshotPolicyConfig
	err := ctx.RegisterResource("tencentcloud:Vpc/snapshotPolicyConfig:SnapshotPolicyConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotPolicyConfig gets an existing SnapshotPolicyConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotPolicyConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotPolicyConfigState, opts ...pulumi.ResourceOption) (*SnapshotPolicyConfig, error) {
	var resource SnapshotPolicyConfig
	err := ctx.ReadResource("tencentcloud:Vpc/snapshotPolicyConfig:SnapshotPolicyConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotPolicyConfig resources.
type snapshotPolicyConfigState struct {
	// If enable snapshot policy.
	Enable *bool `pulumi:"enable"`
	// Snapshot policy Id.
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
}

type SnapshotPolicyConfigState struct {
	// If enable snapshot policy.
	Enable pulumi.BoolPtrInput
	// Snapshot policy Id.
	SnapshotPolicyId pulumi.StringPtrInput
}

func (SnapshotPolicyConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyConfigState)(nil)).Elem()
}

type snapshotPolicyConfigArgs struct {
	// If enable snapshot policy.
	Enable bool `pulumi:"enable"`
	// Snapshot policy Id.
	SnapshotPolicyId string `pulumi:"snapshotPolicyId"`
}

// The set of arguments for constructing a SnapshotPolicyConfig resource.
type SnapshotPolicyConfigArgs struct {
	// If enable snapshot policy.
	Enable pulumi.BoolInput
	// Snapshot policy Id.
	SnapshotPolicyId pulumi.StringInput
}

func (SnapshotPolicyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotPolicyConfigArgs)(nil)).Elem()
}

type SnapshotPolicyConfigInput interface {
	pulumi.Input

	ToSnapshotPolicyConfigOutput() SnapshotPolicyConfigOutput
	ToSnapshotPolicyConfigOutputWithContext(ctx context.Context) SnapshotPolicyConfigOutput
}

func (*SnapshotPolicyConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicyConfig)(nil)).Elem()
}

func (i *SnapshotPolicyConfig) ToSnapshotPolicyConfigOutput() SnapshotPolicyConfigOutput {
	return i.ToSnapshotPolicyConfigOutputWithContext(context.Background())
}

func (i *SnapshotPolicyConfig) ToSnapshotPolicyConfigOutputWithContext(ctx context.Context) SnapshotPolicyConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyConfigOutput)
}

// SnapshotPolicyConfigArrayInput is an input type that accepts SnapshotPolicyConfigArray and SnapshotPolicyConfigArrayOutput values.
// You can construct a concrete instance of `SnapshotPolicyConfigArrayInput` via:
//
//          SnapshotPolicyConfigArray{ SnapshotPolicyConfigArgs{...} }
type SnapshotPolicyConfigArrayInput interface {
	pulumi.Input

	ToSnapshotPolicyConfigArrayOutput() SnapshotPolicyConfigArrayOutput
	ToSnapshotPolicyConfigArrayOutputWithContext(context.Context) SnapshotPolicyConfigArrayOutput
}

type SnapshotPolicyConfigArray []SnapshotPolicyConfigInput

func (SnapshotPolicyConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotPolicyConfig)(nil)).Elem()
}

func (i SnapshotPolicyConfigArray) ToSnapshotPolicyConfigArrayOutput() SnapshotPolicyConfigArrayOutput {
	return i.ToSnapshotPolicyConfigArrayOutputWithContext(context.Background())
}

func (i SnapshotPolicyConfigArray) ToSnapshotPolicyConfigArrayOutputWithContext(ctx context.Context) SnapshotPolicyConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyConfigArrayOutput)
}

// SnapshotPolicyConfigMapInput is an input type that accepts SnapshotPolicyConfigMap and SnapshotPolicyConfigMapOutput values.
// You can construct a concrete instance of `SnapshotPolicyConfigMapInput` via:
//
//          SnapshotPolicyConfigMap{ "key": SnapshotPolicyConfigArgs{...} }
type SnapshotPolicyConfigMapInput interface {
	pulumi.Input

	ToSnapshotPolicyConfigMapOutput() SnapshotPolicyConfigMapOutput
	ToSnapshotPolicyConfigMapOutputWithContext(context.Context) SnapshotPolicyConfigMapOutput
}

type SnapshotPolicyConfigMap map[string]SnapshotPolicyConfigInput

func (SnapshotPolicyConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotPolicyConfig)(nil)).Elem()
}

func (i SnapshotPolicyConfigMap) ToSnapshotPolicyConfigMapOutput() SnapshotPolicyConfigMapOutput {
	return i.ToSnapshotPolicyConfigMapOutputWithContext(context.Background())
}

func (i SnapshotPolicyConfigMap) ToSnapshotPolicyConfigMapOutputWithContext(ctx context.Context) SnapshotPolicyConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotPolicyConfigMapOutput)
}

type SnapshotPolicyConfigOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotPolicyConfig)(nil)).Elem()
}

func (o SnapshotPolicyConfigOutput) ToSnapshotPolicyConfigOutput() SnapshotPolicyConfigOutput {
	return o
}

func (o SnapshotPolicyConfigOutput) ToSnapshotPolicyConfigOutputWithContext(ctx context.Context) SnapshotPolicyConfigOutput {
	return o
}

// If enable snapshot policy.
func (o SnapshotPolicyConfigOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v *SnapshotPolicyConfig) pulumi.BoolOutput { return v.Enable }).(pulumi.BoolOutput)
}

// Snapshot policy Id.
func (o SnapshotPolicyConfigOutput) SnapshotPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnapshotPolicyConfig) pulumi.StringOutput { return v.SnapshotPolicyId }).(pulumi.StringOutput)
}

type SnapshotPolicyConfigArrayOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnapshotPolicyConfig)(nil)).Elem()
}

func (o SnapshotPolicyConfigArrayOutput) ToSnapshotPolicyConfigArrayOutput() SnapshotPolicyConfigArrayOutput {
	return o
}

func (o SnapshotPolicyConfigArrayOutput) ToSnapshotPolicyConfigArrayOutputWithContext(ctx context.Context) SnapshotPolicyConfigArrayOutput {
	return o
}

func (o SnapshotPolicyConfigArrayOutput) Index(i pulumi.IntInput) SnapshotPolicyConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnapshotPolicyConfig {
		return vs[0].([]*SnapshotPolicyConfig)[vs[1].(int)]
	}).(SnapshotPolicyConfigOutput)
}

type SnapshotPolicyConfigMapOutput struct{ *pulumi.OutputState }

func (SnapshotPolicyConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnapshotPolicyConfig)(nil)).Elem()
}

func (o SnapshotPolicyConfigMapOutput) ToSnapshotPolicyConfigMapOutput() SnapshotPolicyConfigMapOutput {
	return o
}

func (o SnapshotPolicyConfigMapOutput) ToSnapshotPolicyConfigMapOutputWithContext(ctx context.Context) SnapshotPolicyConfigMapOutput {
	return o
}

func (o SnapshotPolicyConfigMapOutput) MapIndex(k pulumi.StringInput) SnapshotPolicyConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnapshotPolicyConfig {
		return vs[0].(map[string]*SnapshotPolicyConfig)[vs[1].(string)]
	}).(SnapshotPolicyConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyConfigInput)(nil)).Elem(), &SnapshotPolicyConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyConfigArrayInput)(nil)).Elem(), SnapshotPolicyConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotPolicyConfigMapInput)(nil)).Elem(), SnapshotPolicyConfigMap{})
	pulumi.RegisterOutputType(SnapshotPolicyConfigOutput{})
	pulumi.RegisterOutputType(SnapshotPolicyConfigArrayOutput{})
	pulumi.RegisterOutputType(SnapshotPolicyConfigMapOutput{})
}
