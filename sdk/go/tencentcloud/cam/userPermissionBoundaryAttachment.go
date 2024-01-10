// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cam userPermissionBoundary
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cam.NewUserPermissionBoundaryAttachment(ctx, "userPermissionBoundary", &Cam.UserPermissionBoundaryAttachmentArgs{
//				PolicyId:  pulumi.Int(151113272),
//				TargetUin: pulumi.Int(100032767426),
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
// cam user_permission_boundary can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cam/userPermissionBoundaryAttachment:UserPermissionBoundaryAttachment user_permission_boundary user_permission_boundary_id
//
// ```
type UserPermissionBoundaryAttachment struct {
	pulumi.CustomResourceState

	// Policy ID.
	PolicyId pulumi.IntOutput `pulumi:"policyId"`
	// Sub account Uin.
	TargetUin pulumi.IntOutput `pulumi:"targetUin"`
}

// NewUserPermissionBoundaryAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserPermissionBoundaryAttachment(ctx *pulumi.Context,
	name string, args *UserPermissionBoundaryAttachmentArgs, opts ...pulumi.ResourceOption) (*UserPermissionBoundaryAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.TargetUin == nil {
		return nil, errors.New("invalid value for required argument 'TargetUin'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserPermissionBoundaryAttachment
	err := ctx.RegisterResource("tencentcloud:Cam/userPermissionBoundaryAttachment:UserPermissionBoundaryAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPermissionBoundaryAttachment gets an existing UserPermissionBoundaryAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPermissionBoundaryAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPermissionBoundaryAttachmentState, opts ...pulumi.ResourceOption) (*UserPermissionBoundaryAttachment, error) {
	var resource UserPermissionBoundaryAttachment
	err := ctx.ReadResource("tencentcloud:Cam/userPermissionBoundaryAttachment:UserPermissionBoundaryAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPermissionBoundaryAttachment resources.
type userPermissionBoundaryAttachmentState struct {
	// Policy ID.
	PolicyId *int `pulumi:"policyId"`
	// Sub account Uin.
	TargetUin *int `pulumi:"targetUin"`
}

type UserPermissionBoundaryAttachmentState struct {
	// Policy ID.
	PolicyId pulumi.IntPtrInput
	// Sub account Uin.
	TargetUin pulumi.IntPtrInput
}

func (UserPermissionBoundaryAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPermissionBoundaryAttachmentState)(nil)).Elem()
}

type userPermissionBoundaryAttachmentArgs struct {
	// Policy ID.
	PolicyId int `pulumi:"policyId"`
	// Sub account Uin.
	TargetUin int `pulumi:"targetUin"`
}

// The set of arguments for constructing a UserPermissionBoundaryAttachment resource.
type UserPermissionBoundaryAttachmentArgs struct {
	// Policy ID.
	PolicyId pulumi.IntInput
	// Sub account Uin.
	TargetUin pulumi.IntInput
}

func (UserPermissionBoundaryAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPermissionBoundaryAttachmentArgs)(nil)).Elem()
}

type UserPermissionBoundaryAttachmentInput interface {
	pulumi.Input

	ToUserPermissionBoundaryAttachmentOutput() UserPermissionBoundaryAttachmentOutput
	ToUserPermissionBoundaryAttachmentOutputWithContext(ctx context.Context) UserPermissionBoundaryAttachmentOutput
}

func (*UserPermissionBoundaryAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPermissionBoundaryAttachment)(nil)).Elem()
}

func (i *UserPermissionBoundaryAttachment) ToUserPermissionBoundaryAttachmentOutput() UserPermissionBoundaryAttachmentOutput {
	return i.ToUserPermissionBoundaryAttachmentOutputWithContext(context.Background())
}

func (i *UserPermissionBoundaryAttachment) ToUserPermissionBoundaryAttachmentOutputWithContext(ctx context.Context) UserPermissionBoundaryAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPermissionBoundaryAttachmentOutput)
}

// UserPermissionBoundaryAttachmentArrayInput is an input type that accepts UserPermissionBoundaryAttachmentArray and UserPermissionBoundaryAttachmentArrayOutput values.
// You can construct a concrete instance of `UserPermissionBoundaryAttachmentArrayInput` via:
//
//	UserPermissionBoundaryAttachmentArray{ UserPermissionBoundaryAttachmentArgs{...} }
type UserPermissionBoundaryAttachmentArrayInput interface {
	pulumi.Input

	ToUserPermissionBoundaryAttachmentArrayOutput() UserPermissionBoundaryAttachmentArrayOutput
	ToUserPermissionBoundaryAttachmentArrayOutputWithContext(context.Context) UserPermissionBoundaryAttachmentArrayOutput
}

type UserPermissionBoundaryAttachmentArray []UserPermissionBoundaryAttachmentInput

func (UserPermissionBoundaryAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPermissionBoundaryAttachment)(nil)).Elem()
}

func (i UserPermissionBoundaryAttachmentArray) ToUserPermissionBoundaryAttachmentArrayOutput() UserPermissionBoundaryAttachmentArrayOutput {
	return i.ToUserPermissionBoundaryAttachmentArrayOutputWithContext(context.Background())
}

func (i UserPermissionBoundaryAttachmentArray) ToUserPermissionBoundaryAttachmentArrayOutputWithContext(ctx context.Context) UserPermissionBoundaryAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPermissionBoundaryAttachmentArrayOutput)
}

// UserPermissionBoundaryAttachmentMapInput is an input type that accepts UserPermissionBoundaryAttachmentMap and UserPermissionBoundaryAttachmentMapOutput values.
// You can construct a concrete instance of `UserPermissionBoundaryAttachmentMapInput` via:
//
//	UserPermissionBoundaryAttachmentMap{ "key": UserPermissionBoundaryAttachmentArgs{...} }
type UserPermissionBoundaryAttachmentMapInput interface {
	pulumi.Input

	ToUserPermissionBoundaryAttachmentMapOutput() UserPermissionBoundaryAttachmentMapOutput
	ToUserPermissionBoundaryAttachmentMapOutputWithContext(context.Context) UserPermissionBoundaryAttachmentMapOutput
}

type UserPermissionBoundaryAttachmentMap map[string]UserPermissionBoundaryAttachmentInput

func (UserPermissionBoundaryAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPermissionBoundaryAttachment)(nil)).Elem()
}

func (i UserPermissionBoundaryAttachmentMap) ToUserPermissionBoundaryAttachmentMapOutput() UserPermissionBoundaryAttachmentMapOutput {
	return i.ToUserPermissionBoundaryAttachmentMapOutputWithContext(context.Background())
}

func (i UserPermissionBoundaryAttachmentMap) ToUserPermissionBoundaryAttachmentMapOutputWithContext(ctx context.Context) UserPermissionBoundaryAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPermissionBoundaryAttachmentMapOutput)
}

type UserPermissionBoundaryAttachmentOutput struct{ *pulumi.OutputState }

func (UserPermissionBoundaryAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPermissionBoundaryAttachment)(nil)).Elem()
}

func (o UserPermissionBoundaryAttachmentOutput) ToUserPermissionBoundaryAttachmentOutput() UserPermissionBoundaryAttachmentOutput {
	return o
}

func (o UserPermissionBoundaryAttachmentOutput) ToUserPermissionBoundaryAttachmentOutputWithContext(ctx context.Context) UserPermissionBoundaryAttachmentOutput {
	return o
}

// Policy ID.
func (o UserPermissionBoundaryAttachmentOutput) PolicyId() pulumi.IntOutput {
	return o.ApplyT(func(v *UserPermissionBoundaryAttachment) pulumi.IntOutput { return v.PolicyId }).(pulumi.IntOutput)
}

// Sub account Uin.
func (o UserPermissionBoundaryAttachmentOutput) TargetUin() pulumi.IntOutput {
	return o.ApplyT(func(v *UserPermissionBoundaryAttachment) pulumi.IntOutput { return v.TargetUin }).(pulumi.IntOutput)
}

type UserPermissionBoundaryAttachmentArrayOutput struct{ *pulumi.OutputState }

func (UserPermissionBoundaryAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPermissionBoundaryAttachment)(nil)).Elem()
}

func (o UserPermissionBoundaryAttachmentArrayOutput) ToUserPermissionBoundaryAttachmentArrayOutput() UserPermissionBoundaryAttachmentArrayOutput {
	return o
}

func (o UserPermissionBoundaryAttachmentArrayOutput) ToUserPermissionBoundaryAttachmentArrayOutputWithContext(ctx context.Context) UserPermissionBoundaryAttachmentArrayOutput {
	return o
}

func (o UserPermissionBoundaryAttachmentArrayOutput) Index(i pulumi.IntInput) UserPermissionBoundaryAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPermissionBoundaryAttachment {
		return vs[0].([]*UserPermissionBoundaryAttachment)[vs[1].(int)]
	}).(UserPermissionBoundaryAttachmentOutput)
}

type UserPermissionBoundaryAttachmentMapOutput struct{ *pulumi.OutputState }

func (UserPermissionBoundaryAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPermissionBoundaryAttachment)(nil)).Elem()
}

func (o UserPermissionBoundaryAttachmentMapOutput) ToUserPermissionBoundaryAttachmentMapOutput() UserPermissionBoundaryAttachmentMapOutput {
	return o
}

func (o UserPermissionBoundaryAttachmentMapOutput) ToUserPermissionBoundaryAttachmentMapOutputWithContext(ctx context.Context) UserPermissionBoundaryAttachmentMapOutput {
	return o
}

func (o UserPermissionBoundaryAttachmentMapOutput) MapIndex(k pulumi.StringInput) UserPermissionBoundaryAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPermissionBoundaryAttachment {
		return vs[0].(map[string]*UserPermissionBoundaryAttachment)[vs[1].(string)]
	}).(UserPermissionBoundaryAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPermissionBoundaryAttachmentInput)(nil)).Elem(), &UserPermissionBoundaryAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPermissionBoundaryAttachmentArrayInput)(nil)).Elem(), UserPermissionBoundaryAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPermissionBoundaryAttachmentMapInput)(nil)).Elem(), UserPermissionBoundaryAttachmentMap{})
	pulumi.RegisterOutputType(UserPermissionBoundaryAttachmentOutput{})
	pulumi.RegisterOutputType(UserPermissionBoundaryAttachmentArrayOutput{})
	pulumi.RegisterOutputType(UserPermissionBoundaryAttachmentMapOutput{})
}
