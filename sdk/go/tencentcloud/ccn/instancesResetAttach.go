// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ccn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc ccn_instances_reset_attach, you can use this resource to reset cross-region attachment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ccn"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ccn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ccn.NewInstancesResetAttach(ctx, "ccnInstancesResetAttach", &Ccn.InstancesResetAttachArgs{
//				CcnId:  pulumi.String("ccn-39lqkygf"),
//				CcnUin: pulumi.String("100022975249"),
//				Instances: ccn.InstancesResetAttachInstanceArray{
//					&ccn.InstancesResetAttachInstanceArgs{
//						InstanceId:     pulumi.String("vpc-j9yhbzpn"),
//						InstanceRegion: pulumi.String("ap-guangzhou"),
//						InstanceType:   pulumi.String("VPC"),
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
type InstancesResetAttach struct {
	pulumi.CustomResourceState

	// CCN Instance ID.
	CcnId pulumi.StringOutput `pulumi:"ccnId"`
	// CCN Uin (root account).
	CcnUin pulumi.StringOutput `pulumi:"ccnUin"`
	// List Of Attachment Instances.
	Instances InstancesResetAttachInstanceArrayOutput `pulumi:"instances"`
}

// NewInstancesResetAttach registers a new resource with the given unique name, arguments, and options.
func NewInstancesResetAttach(ctx *pulumi.Context,
	name string, args *InstancesResetAttachArgs, opts ...pulumi.ResourceOption) (*InstancesResetAttach, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CcnId == nil {
		return nil, errors.New("invalid value for required argument 'CcnId'")
	}
	if args.CcnUin == nil {
		return nil, errors.New("invalid value for required argument 'CcnUin'")
	}
	if args.Instances == nil {
		return nil, errors.New("invalid value for required argument 'Instances'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstancesResetAttach
	err := ctx.RegisterResource("tencentcloud:Ccn/instancesResetAttach:InstancesResetAttach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstancesResetAttach gets an existing InstancesResetAttach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstancesResetAttach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstancesResetAttachState, opts ...pulumi.ResourceOption) (*InstancesResetAttach, error) {
	var resource InstancesResetAttach
	err := ctx.ReadResource("tencentcloud:Ccn/instancesResetAttach:InstancesResetAttach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstancesResetAttach resources.
type instancesResetAttachState struct {
	// CCN Instance ID.
	CcnId *string `pulumi:"ccnId"`
	// CCN Uin (root account).
	CcnUin *string `pulumi:"ccnUin"`
	// List Of Attachment Instances.
	Instances []InstancesResetAttachInstance `pulumi:"instances"`
}

type InstancesResetAttachState struct {
	// CCN Instance ID.
	CcnId pulumi.StringPtrInput
	// CCN Uin (root account).
	CcnUin pulumi.StringPtrInput
	// List Of Attachment Instances.
	Instances InstancesResetAttachInstanceArrayInput
}

func (InstancesResetAttachState) ElementType() reflect.Type {
	return reflect.TypeOf((*instancesResetAttachState)(nil)).Elem()
}

type instancesResetAttachArgs struct {
	// CCN Instance ID.
	CcnId string `pulumi:"ccnId"`
	// CCN Uin (root account).
	CcnUin string `pulumi:"ccnUin"`
	// List Of Attachment Instances.
	Instances []InstancesResetAttachInstance `pulumi:"instances"`
}

// The set of arguments for constructing a InstancesResetAttach resource.
type InstancesResetAttachArgs struct {
	// CCN Instance ID.
	CcnId pulumi.StringInput
	// CCN Uin (root account).
	CcnUin pulumi.StringInput
	// List Of Attachment Instances.
	Instances InstancesResetAttachInstanceArrayInput
}

func (InstancesResetAttachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instancesResetAttachArgs)(nil)).Elem()
}

type InstancesResetAttachInput interface {
	pulumi.Input

	ToInstancesResetAttachOutput() InstancesResetAttachOutput
	ToInstancesResetAttachOutputWithContext(ctx context.Context) InstancesResetAttachOutput
}

func (*InstancesResetAttach) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancesResetAttach)(nil)).Elem()
}

func (i *InstancesResetAttach) ToInstancesResetAttachOutput() InstancesResetAttachOutput {
	return i.ToInstancesResetAttachOutputWithContext(context.Background())
}

func (i *InstancesResetAttach) ToInstancesResetAttachOutputWithContext(ctx context.Context) InstancesResetAttachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesResetAttachOutput)
}

// InstancesResetAttachArrayInput is an input type that accepts InstancesResetAttachArray and InstancesResetAttachArrayOutput values.
// You can construct a concrete instance of `InstancesResetAttachArrayInput` via:
//
//	InstancesResetAttachArray{ InstancesResetAttachArgs{...} }
type InstancesResetAttachArrayInput interface {
	pulumi.Input

	ToInstancesResetAttachArrayOutput() InstancesResetAttachArrayOutput
	ToInstancesResetAttachArrayOutputWithContext(context.Context) InstancesResetAttachArrayOutput
}

type InstancesResetAttachArray []InstancesResetAttachInput

func (InstancesResetAttachArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancesResetAttach)(nil)).Elem()
}

func (i InstancesResetAttachArray) ToInstancesResetAttachArrayOutput() InstancesResetAttachArrayOutput {
	return i.ToInstancesResetAttachArrayOutputWithContext(context.Background())
}

func (i InstancesResetAttachArray) ToInstancesResetAttachArrayOutputWithContext(ctx context.Context) InstancesResetAttachArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesResetAttachArrayOutput)
}

// InstancesResetAttachMapInput is an input type that accepts InstancesResetAttachMap and InstancesResetAttachMapOutput values.
// You can construct a concrete instance of `InstancesResetAttachMapInput` via:
//
//	InstancesResetAttachMap{ "key": InstancesResetAttachArgs{...} }
type InstancesResetAttachMapInput interface {
	pulumi.Input

	ToInstancesResetAttachMapOutput() InstancesResetAttachMapOutput
	ToInstancesResetAttachMapOutputWithContext(context.Context) InstancesResetAttachMapOutput
}

type InstancesResetAttachMap map[string]InstancesResetAttachInput

func (InstancesResetAttachMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancesResetAttach)(nil)).Elem()
}

func (i InstancesResetAttachMap) ToInstancesResetAttachMapOutput() InstancesResetAttachMapOutput {
	return i.ToInstancesResetAttachMapOutputWithContext(context.Background())
}

func (i InstancesResetAttachMap) ToInstancesResetAttachMapOutputWithContext(ctx context.Context) InstancesResetAttachMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstancesResetAttachMapOutput)
}

type InstancesResetAttachOutput struct{ *pulumi.OutputState }

func (InstancesResetAttachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstancesResetAttach)(nil)).Elem()
}

func (o InstancesResetAttachOutput) ToInstancesResetAttachOutput() InstancesResetAttachOutput {
	return o
}

func (o InstancesResetAttachOutput) ToInstancesResetAttachOutputWithContext(ctx context.Context) InstancesResetAttachOutput {
	return o
}

// CCN Instance ID.
func (o InstancesResetAttachOutput) CcnId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancesResetAttach) pulumi.StringOutput { return v.CcnId }).(pulumi.StringOutput)
}

// CCN Uin (root account).
func (o InstancesResetAttachOutput) CcnUin() pulumi.StringOutput {
	return o.ApplyT(func(v *InstancesResetAttach) pulumi.StringOutput { return v.CcnUin }).(pulumi.StringOutput)
}

// List Of Attachment Instances.
func (o InstancesResetAttachOutput) Instances() InstancesResetAttachInstanceArrayOutput {
	return o.ApplyT(func(v *InstancesResetAttach) InstancesResetAttachInstanceArrayOutput { return v.Instances }).(InstancesResetAttachInstanceArrayOutput)
}

type InstancesResetAttachArrayOutput struct{ *pulumi.OutputState }

func (InstancesResetAttachArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstancesResetAttach)(nil)).Elem()
}

func (o InstancesResetAttachArrayOutput) ToInstancesResetAttachArrayOutput() InstancesResetAttachArrayOutput {
	return o
}

func (o InstancesResetAttachArrayOutput) ToInstancesResetAttachArrayOutputWithContext(ctx context.Context) InstancesResetAttachArrayOutput {
	return o
}

func (o InstancesResetAttachArrayOutput) Index(i pulumi.IntInput) InstancesResetAttachOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstancesResetAttach {
		return vs[0].([]*InstancesResetAttach)[vs[1].(int)]
	}).(InstancesResetAttachOutput)
}

type InstancesResetAttachMapOutput struct{ *pulumi.OutputState }

func (InstancesResetAttachMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstancesResetAttach)(nil)).Elem()
}

func (o InstancesResetAttachMapOutput) ToInstancesResetAttachMapOutput() InstancesResetAttachMapOutput {
	return o
}

func (o InstancesResetAttachMapOutput) ToInstancesResetAttachMapOutputWithContext(ctx context.Context) InstancesResetAttachMapOutput {
	return o
}

func (o InstancesResetAttachMapOutput) MapIndex(k pulumi.StringInput) InstancesResetAttachOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstancesResetAttach {
		return vs[0].(map[string]*InstancesResetAttach)[vs[1].(string)]
	}).(InstancesResetAttachOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesResetAttachInput)(nil)).Elem(), &InstancesResetAttach{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesResetAttachArrayInput)(nil)).Elem(), InstancesResetAttachArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstancesResetAttachMapInput)(nil)).Elem(), InstancesResetAttachMap{})
	pulumi.RegisterOutputType(InstancesResetAttachOutput{})
	pulumi.RegisterOutputType(InstancesResetAttachArrayOutput{})
	pulumi.RegisterOutputType(InstancesResetAttachMapOutput{})
}
