// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a dlc bindWorkGroupsToUser
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dlc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dlc.NewBindWorkGroupsToUserAttachment(ctx, "bindWorkGroupsToUser", &Dlc.BindWorkGroupsToUserAttachmentArgs{
//				AddInfo: &dlc.BindWorkGroupsToUserAttachmentAddInfoArgs{
//					UserId: pulumi.String("100032772113"),
//					WorkGroupIds: pulumi.IntArray{
//						pulumi.Int(23184),
//						pulumi.Int(23181),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// dlc bind_work_groups_to_user can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Dlc/bindWorkGroupsToUserAttachment:BindWorkGroupsToUserAttachment bind_work_groups_to_user bind_work_groups_to_user_id
// ```
type BindWorkGroupsToUserAttachment struct {
	pulumi.CustomResourceState

	// Bind user and workgroup information.
	AddInfo BindWorkGroupsToUserAttachmentAddInfoOutput `pulumi:"addInfo"`
}

// NewBindWorkGroupsToUserAttachment registers a new resource with the given unique name, arguments, and options.
func NewBindWorkGroupsToUserAttachment(ctx *pulumi.Context,
	name string, args *BindWorkGroupsToUserAttachmentArgs, opts ...pulumi.ResourceOption) (*BindWorkGroupsToUserAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddInfo == nil {
		return nil, errors.New("invalid value for required argument 'AddInfo'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BindWorkGroupsToUserAttachment
	err := ctx.RegisterResource("tencentcloud:Dlc/bindWorkGroupsToUserAttachment:BindWorkGroupsToUserAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBindWorkGroupsToUserAttachment gets an existing BindWorkGroupsToUserAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBindWorkGroupsToUserAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BindWorkGroupsToUserAttachmentState, opts ...pulumi.ResourceOption) (*BindWorkGroupsToUserAttachment, error) {
	var resource BindWorkGroupsToUserAttachment
	err := ctx.ReadResource("tencentcloud:Dlc/bindWorkGroupsToUserAttachment:BindWorkGroupsToUserAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BindWorkGroupsToUserAttachment resources.
type bindWorkGroupsToUserAttachmentState struct {
	// Bind user and workgroup information.
	AddInfo *BindWorkGroupsToUserAttachmentAddInfo `pulumi:"addInfo"`
}

type BindWorkGroupsToUserAttachmentState struct {
	// Bind user and workgroup information.
	AddInfo BindWorkGroupsToUserAttachmentAddInfoPtrInput
}

func (BindWorkGroupsToUserAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*bindWorkGroupsToUserAttachmentState)(nil)).Elem()
}

type bindWorkGroupsToUserAttachmentArgs struct {
	// Bind user and workgroup information.
	AddInfo BindWorkGroupsToUserAttachmentAddInfo `pulumi:"addInfo"`
}

// The set of arguments for constructing a BindWorkGroupsToUserAttachment resource.
type BindWorkGroupsToUserAttachmentArgs struct {
	// Bind user and workgroup information.
	AddInfo BindWorkGroupsToUserAttachmentAddInfoInput
}

func (BindWorkGroupsToUserAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bindWorkGroupsToUserAttachmentArgs)(nil)).Elem()
}

type BindWorkGroupsToUserAttachmentInput interface {
	pulumi.Input

	ToBindWorkGroupsToUserAttachmentOutput() BindWorkGroupsToUserAttachmentOutput
	ToBindWorkGroupsToUserAttachmentOutputWithContext(ctx context.Context) BindWorkGroupsToUserAttachmentOutput
}

func (*BindWorkGroupsToUserAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**BindWorkGroupsToUserAttachment)(nil)).Elem()
}

func (i *BindWorkGroupsToUserAttachment) ToBindWorkGroupsToUserAttachmentOutput() BindWorkGroupsToUserAttachmentOutput {
	return i.ToBindWorkGroupsToUserAttachmentOutputWithContext(context.Background())
}

func (i *BindWorkGroupsToUserAttachment) ToBindWorkGroupsToUserAttachmentOutputWithContext(ctx context.Context) BindWorkGroupsToUserAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindWorkGroupsToUserAttachmentOutput)
}

// BindWorkGroupsToUserAttachmentArrayInput is an input type that accepts BindWorkGroupsToUserAttachmentArray and BindWorkGroupsToUserAttachmentArrayOutput values.
// You can construct a concrete instance of `BindWorkGroupsToUserAttachmentArrayInput` via:
//
//	BindWorkGroupsToUserAttachmentArray{ BindWorkGroupsToUserAttachmentArgs{...} }
type BindWorkGroupsToUserAttachmentArrayInput interface {
	pulumi.Input

	ToBindWorkGroupsToUserAttachmentArrayOutput() BindWorkGroupsToUserAttachmentArrayOutput
	ToBindWorkGroupsToUserAttachmentArrayOutputWithContext(context.Context) BindWorkGroupsToUserAttachmentArrayOutput
}

type BindWorkGroupsToUserAttachmentArray []BindWorkGroupsToUserAttachmentInput

func (BindWorkGroupsToUserAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindWorkGroupsToUserAttachment)(nil)).Elem()
}

func (i BindWorkGroupsToUserAttachmentArray) ToBindWorkGroupsToUserAttachmentArrayOutput() BindWorkGroupsToUserAttachmentArrayOutput {
	return i.ToBindWorkGroupsToUserAttachmentArrayOutputWithContext(context.Background())
}

func (i BindWorkGroupsToUserAttachmentArray) ToBindWorkGroupsToUserAttachmentArrayOutputWithContext(ctx context.Context) BindWorkGroupsToUserAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindWorkGroupsToUserAttachmentArrayOutput)
}

// BindWorkGroupsToUserAttachmentMapInput is an input type that accepts BindWorkGroupsToUserAttachmentMap and BindWorkGroupsToUserAttachmentMapOutput values.
// You can construct a concrete instance of `BindWorkGroupsToUserAttachmentMapInput` via:
//
//	BindWorkGroupsToUserAttachmentMap{ "key": BindWorkGroupsToUserAttachmentArgs{...} }
type BindWorkGroupsToUserAttachmentMapInput interface {
	pulumi.Input

	ToBindWorkGroupsToUserAttachmentMapOutput() BindWorkGroupsToUserAttachmentMapOutput
	ToBindWorkGroupsToUserAttachmentMapOutputWithContext(context.Context) BindWorkGroupsToUserAttachmentMapOutput
}

type BindWorkGroupsToUserAttachmentMap map[string]BindWorkGroupsToUserAttachmentInput

func (BindWorkGroupsToUserAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindWorkGroupsToUserAttachment)(nil)).Elem()
}

func (i BindWorkGroupsToUserAttachmentMap) ToBindWorkGroupsToUserAttachmentMapOutput() BindWorkGroupsToUserAttachmentMapOutput {
	return i.ToBindWorkGroupsToUserAttachmentMapOutputWithContext(context.Background())
}

func (i BindWorkGroupsToUserAttachmentMap) ToBindWorkGroupsToUserAttachmentMapOutputWithContext(ctx context.Context) BindWorkGroupsToUserAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindWorkGroupsToUserAttachmentMapOutput)
}

type BindWorkGroupsToUserAttachmentOutput struct{ *pulumi.OutputState }

func (BindWorkGroupsToUserAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BindWorkGroupsToUserAttachment)(nil)).Elem()
}

func (o BindWorkGroupsToUserAttachmentOutput) ToBindWorkGroupsToUserAttachmentOutput() BindWorkGroupsToUserAttachmentOutput {
	return o
}

func (o BindWorkGroupsToUserAttachmentOutput) ToBindWorkGroupsToUserAttachmentOutputWithContext(ctx context.Context) BindWorkGroupsToUserAttachmentOutput {
	return o
}

// Bind user and workgroup information.
func (o BindWorkGroupsToUserAttachmentOutput) AddInfo() BindWorkGroupsToUserAttachmentAddInfoOutput {
	return o.ApplyT(func(v *BindWorkGroupsToUserAttachment) BindWorkGroupsToUserAttachmentAddInfoOutput { return v.AddInfo }).(BindWorkGroupsToUserAttachmentAddInfoOutput)
}

type BindWorkGroupsToUserAttachmentArrayOutput struct{ *pulumi.OutputState }

func (BindWorkGroupsToUserAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BindWorkGroupsToUserAttachment)(nil)).Elem()
}

func (o BindWorkGroupsToUserAttachmentArrayOutput) ToBindWorkGroupsToUserAttachmentArrayOutput() BindWorkGroupsToUserAttachmentArrayOutput {
	return o
}

func (o BindWorkGroupsToUserAttachmentArrayOutput) ToBindWorkGroupsToUserAttachmentArrayOutputWithContext(ctx context.Context) BindWorkGroupsToUserAttachmentArrayOutput {
	return o
}

func (o BindWorkGroupsToUserAttachmentArrayOutput) Index(i pulumi.IntInput) BindWorkGroupsToUserAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BindWorkGroupsToUserAttachment {
		return vs[0].([]*BindWorkGroupsToUserAttachment)[vs[1].(int)]
	}).(BindWorkGroupsToUserAttachmentOutput)
}

type BindWorkGroupsToUserAttachmentMapOutput struct{ *pulumi.OutputState }

func (BindWorkGroupsToUserAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BindWorkGroupsToUserAttachment)(nil)).Elem()
}

func (o BindWorkGroupsToUserAttachmentMapOutput) ToBindWorkGroupsToUserAttachmentMapOutput() BindWorkGroupsToUserAttachmentMapOutput {
	return o
}

func (o BindWorkGroupsToUserAttachmentMapOutput) ToBindWorkGroupsToUserAttachmentMapOutputWithContext(ctx context.Context) BindWorkGroupsToUserAttachmentMapOutput {
	return o
}

func (o BindWorkGroupsToUserAttachmentMapOutput) MapIndex(k pulumi.StringInput) BindWorkGroupsToUserAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BindWorkGroupsToUserAttachment {
		return vs[0].(map[string]*BindWorkGroupsToUserAttachment)[vs[1].(string)]
	}).(BindWorkGroupsToUserAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindWorkGroupsToUserAttachmentInput)(nil)).Elem(), &BindWorkGroupsToUserAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindWorkGroupsToUserAttachmentArrayInput)(nil)).Elem(), BindWorkGroupsToUserAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindWorkGroupsToUserAttachmentMapInput)(nil)).Elem(), BindWorkGroupsToUserAttachmentMap{})
	pulumi.RegisterOutputType(BindWorkGroupsToUserAttachmentOutput{})
	pulumi.RegisterOutputType(BindWorkGroupsToUserAttachmentArrayOutput{})
	pulumi.RegisterOutputType(BindWorkGroupsToUserAttachmentMapOutput{})
}
