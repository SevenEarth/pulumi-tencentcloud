// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CBS storage attachment resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cbs.NewStorageAttachment(ctx, "attachment", &Cbs.StorageAttachmentArgs{
//				InstanceId: pulumi.String("ins-jqlegd42"),
//				StorageId:  pulumi.String("disk-kdt0sq6m"),
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
// CBS storage attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cbs/storageAttachment:StorageAttachment attachment disk-41s6jwy4
//
// ```
type StorageAttachment struct {
	pulumi.CustomResourceState

	// ID of the CVM instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// ID of the mounted CBS.
	StorageId pulumi.StringOutput `pulumi:"storageId"`
}

// NewStorageAttachment registers a new resource with the given unique name, arguments, and options.
func NewStorageAttachment(ctx *pulumi.Context,
	name string, args *StorageAttachmentArgs, opts ...pulumi.ResourceOption) (*StorageAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.StorageId == nil {
		return nil, errors.New("invalid value for required argument 'StorageId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StorageAttachment
	err := ctx.RegisterResource("tencentcloud:Cbs/storageAttachment:StorageAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAttachment gets an existing StorageAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAttachmentState, opts ...pulumi.ResourceOption) (*StorageAttachment, error) {
	var resource StorageAttachment
	err := ctx.ReadResource("tencentcloud:Cbs/storageAttachment:StorageAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAttachment resources.
type storageAttachmentState struct {
	// ID of the CVM instance.
	InstanceId *string `pulumi:"instanceId"`
	// ID of the mounted CBS.
	StorageId *string `pulumi:"storageId"`
}

type StorageAttachmentState struct {
	// ID of the CVM instance.
	InstanceId pulumi.StringPtrInput
	// ID of the mounted CBS.
	StorageId pulumi.StringPtrInput
}

func (StorageAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAttachmentState)(nil)).Elem()
}

type storageAttachmentArgs struct {
	// ID of the CVM instance.
	InstanceId string `pulumi:"instanceId"`
	// ID of the mounted CBS.
	StorageId string `pulumi:"storageId"`
}

// The set of arguments for constructing a StorageAttachment resource.
type StorageAttachmentArgs struct {
	// ID of the CVM instance.
	InstanceId pulumi.StringInput
	// ID of the mounted CBS.
	StorageId pulumi.StringInput
}

func (StorageAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAttachmentArgs)(nil)).Elem()
}

type StorageAttachmentInput interface {
	pulumi.Input

	ToStorageAttachmentOutput() StorageAttachmentOutput
	ToStorageAttachmentOutputWithContext(ctx context.Context) StorageAttachmentOutput
}

func (*StorageAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAttachment)(nil)).Elem()
}

func (i *StorageAttachment) ToStorageAttachmentOutput() StorageAttachmentOutput {
	return i.ToStorageAttachmentOutputWithContext(context.Background())
}

func (i *StorageAttachment) ToStorageAttachmentOutputWithContext(ctx context.Context) StorageAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAttachmentOutput)
}

// StorageAttachmentArrayInput is an input type that accepts StorageAttachmentArray and StorageAttachmentArrayOutput values.
// You can construct a concrete instance of `StorageAttachmentArrayInput` via:
//
//	StorageAttachmentArray{ StorageAttachmentArgs{...} }
type StorageAttachmentArrayInput interface {
	pulumi.Input

	ToStorageAttachmentArrayOutput() StorageAttachmentArrayOutput
	ToStorageAttachmentArrayOutputWithContext(context.Context) StorageAttachmentArrayOutput
}

type StorageAttachmentArray []StorageAttachmentInput

func (StorageAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageAttachment)(nil)).Elem()
}

func (i StorageAttachmentArray) ToStorageAttachmentArrayOutput() StorageAttachmentArrayOutput {
	return i.ToStorageAttachmentArrayOutputWithContext(context.Background())
}

func (i StorageAttachmentArray) ToStorageAttachmentArrayOutputWithContext(ctx context.Context) StorageAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAttachmentArrayOutput)
}

// StorageAttachmentMapInput is an input type that accepts StorageAttachmentMap and StorageAttachmentMapOutput values.
// You can construct a concrete instance of `StorageAttachmentMapInput` via:
//
//	StorageAttachmentMap{ "key": StorageAttachmentArgs{...} }
type StorageAttachmentMapInput interface {
	pulumi.Input

	ToStorageAttachmentMapOutput() StorageAttachmentMapOutput
	ToStorageAttachmentMapOutputWithContext(context.Context) StorageAttachmentMapOutput
}

type StorageAttachmentMap map[string]StorageAttachmentInput

func (StorageAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageAttachment)(nil)).Elem()
}

func (i StorageAttachmentMap) ToStorageAttachmentMapOutput() StorageAttachmentMapOutput {
	return i.ToStorageAttachmentMapOutputWithContext(context.Background())
}

func (i StorageAttachmentMap) ToStorageAttachmentMapOutputWithContext(ctx context.Context) StorageAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAttachmentMapOutput)
}

type StorageAttachmentOutput struct{ *pulumi.OutputState }

func (StorageAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAttachment)(nil)).Elem()
}

func (o StorageAttachmentOutput) ToStorageAttachmentOutput() StorageAttachmentOutput {
	return o
}

func (o StorageAttachmentOutput) ToStorageAttachmentOutputWithContext(ctx context.Context) StorageAttachmentOutput {
	return o
}

// ID of the CVM instance.
func (o StorageAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// ID of the mounted CBS.
func (o StorageAttachmentOutput) StorageId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAttachment) pulumi.StringOutput { return v.StorageId }).(pulumi.StringOutput)
}

type StorageAttachmentArrayOutput struct{ *pulumi.OutputState }

func (StorageAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StorageAttachment)(nil)).Elem()
}

func (o StorageAttachmentArrayOutput) ToStorageAttachmentArrayOutput() StorageAttachmentArrayOutput {
	return o
}

func (o StorageAttachmentArrayOutput) ToStorageAttachmentArrayOutputWithContext(ctx context.Context) StorageAttachmentArrayOutput {
	return o
}

func (o StorageAttachmentArrayOutput) Index(i pulumi.IntInput) StorageAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StorageAttachment {
		return vs[0].([]*StorageAttachment)[vs[1].(int)]
	}).(StorageAttachmentOutput)
}

type StorageAttachmentMapOutput struct{ *pulumi.OutputState }

func (StorageAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StorageAttachment)(nil)).Elem()
}

func (o StorageAttachmentMapOutput) ToStorageAttachmentMapOutput() StorageAttachmentMapOutput {
	return o
}

func (o StorageAttachmentMapOutput) ToStorageAttachmentMapOutputWithContext(ctx context.Context) StorageAttachmentMapOutput {
	return o
}

func (o StorageAttachmentMapOutput) MapIndex(k pulumi.StringInput) StorageAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StorageAttachment {
		return vs[0].(map[string]*StorageAttachment)[vs[1].(string)]
	}).(StorageAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAttachmentInput)(nil)).Elem(), &StorageAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAttachmentArrayInput)(nil)).Elem(), StorageAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAttachmentMapInput)(nil)).Elem(), StorageAttachmentMap{})
	pulumi.RegisterOutputType(StorageAttachmentOutput{})
	pulumi.RegisterOutputType(StorageAttachmentArrayOutput{})
	pulumi.RegisterOutputType(StorageAttachmentMapOutput{})
}
