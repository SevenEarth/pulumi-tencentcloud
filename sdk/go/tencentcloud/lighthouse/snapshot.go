// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a lighthouse snapshot
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Lighthouse.NewSnapshot(ctx, "snapshot", &Lighthouse.SnapshotArgs{
// 			InstanceId:   pulumi.String("lhins-acd1234"),
// 			SnapshotName: pulumi.String("snap_20200903"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// ID of the instance for which to create a snapshot.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Snapshot name, which can contain up to 60 characters.
	SnapshotName pulumi.StringPtrOutput `pulumi:"snapshotName"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Snapshot
	err := ctx.RegisterResource("tencentcloud:Lighthouse/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("tencentcloud:Lighthouse/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// ID of the instance for which to create a snapshot.
	InstanceId *string `pulumi:"instanceId"`
	// Snapshot name, which can contain up to 60 characters.
	SnapshotName *string `pulumi:"snapshotName"`
}

type SnapshotState struct {
	// ID of the instance for which to create a snapshot.
	InstanceId pulumi.StringPtrInput
	// Snapshot name, which can contain up to 60 characters.
	SnapshotName pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// ID of the instance for which to create a snapshot.
	InstanceId string `pulumi:"instanceId"`
	// Snapshot name, which can contain up to 60 characters.
	SnapshotName *string `pulumi:"snapshotName"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// ID of the instance for which to create a snapshot.
	InstanceId pulumi.StringInput
	// Snapshot name, which can contain up to 60 characters.
	SnapshotName pulumi.StringPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

// SnapshotArrayInput is an input type that accepts SnapshotArray and SnapshotArrayOutput values.
// You can construct a concrete instance of `SnapshotArrayInput` via:
//
//          SnapshotArray{ SnapshotArgs{...} }
type SnapshotArrayInput interface {
	pulumi.Input

	ToSnapshotArrayOutput() SnapshotArrayOutput
	ToSnapshotArrayOutputWithContext(context.Context) SnapshotArrayOutput
}

type SnapshotArray []SnapshotInput

func (SnapshotArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (i SnapshotArray) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return i.ToSnapshotArrayOutputWithContext(context.Background())
}

func (i SnapshotArray) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotArrayOutput)
}

// SnapshotMapInput is an input type that accepts SnapshotMap and SnapshotMapOutput values.
// You can construct a concrete instance of `SnapshotMapInput` via:
//
//          SnapshotMap{ "key": SnapshotArgs{...} }
type SnapshotMapInput interface {
	pulumi.Input

	ToSnapshotMapOutput() SnapshotMapOutput
	ToSnapshotMapOutputWithContext(context.Context) SnapshotMapOutput
}

type SnapshotMap map[string]SnapshotInput

func (SnapshotMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (i SnapshotMap) ToSnapshotMapOutput() SnapshotMapOutput {
	return i.ToSnapshotMapOutputWithContext(context.Background())
}

func (i SnapshotMap) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotMapOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// ID of the instance for which to create a snapshot.
func (o SnapshotOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Snapshot name, which can contain up to 60 characters.
func (o SnapshotOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

type SnapshotArrayOutput struct{ *pulumi.OutputState }

func (SnapshotArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snapshot)(nil)).Elem()
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutput() SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) ToSnapshotArrayOutputWithContext(ctx context.Context) SnapshotArrayOutput {
	return o
}

func (o SnapshotArrayOutput) Index(i pulumi.IntInput) SnapshotOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].([]*Snapshot)[vs[1].(int)]
	}).(SnapshotOutput)
}

type SnapshotMapOutput struct{ *pulumi.OutputState }

func (SnapshotMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snapshot)(nil)).Elem()
}

func (o SnapshotMapOutput) ToSnapshotMapOutput() SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) ToSnapshotMapOutputWithContext(ctx context.Context) SnapshotMapOutput {
	return o
}

func (o SnapshotMapOutput) MapIndex(k pulumi.StringInput) SnapshotOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snapshot {
		return vs[0].(map[string]*Snapshot)[vs[1].(string)]
	}).(SnapshotOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotArrayInput)(nil)).Elem(), SnapshotArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotMapInput)(nil)).Elem(), SnapshotMap{})
	pulumi.RegisterOutputType(SnapshotOutput{})
	pulumi.RegisterOutputType(SnapshotArrayOutput{})
	pulumi.RegisterOutputType(SnapshotMapOutput{})
}
