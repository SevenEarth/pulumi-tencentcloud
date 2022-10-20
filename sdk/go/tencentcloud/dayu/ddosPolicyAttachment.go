// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dayu DDoS policy attachment.
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
//			_, err := Dayu.NewDdosPolicyAttachment(ctx, "dayuDdosPolicyAttachmentBasic", &Dayu.DdosPolicyAttachmentArgs{
//				ResourceType: pulumi.Any(tencentcloud_dayu_ddos_policy.Test_policy.Resource_type),
//				ResourceId:   pulumi.String("bgpip-00000294"),
//				PolicyId:     pulumi.Any(tencentcloud_dayu_ddos_policy.Test_policy.Policy_id),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DdosPolicyAttachment struct {
	pulumi.CustomResourceState

	// ID of the policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// ID of the attached resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
}

// NewDdosPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewDdosPolicyAttachment(ctx *pulumi.Context,
	name string, args *DdosPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*DdosPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DdosPolicyAttachment
	err := ctx.RegisterResource("tencentcloud:Dayu/ddosPolicyAttachment:DdosPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosPolicyAttachment gets an existing DdosPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosPolicyAttachmentState, opts ...pulumi.ResourceOption) (*DdosPolicyAttachment, error) {
	var resource DdosPolicyAttachment
	err := ctx.ReadResource("tencentcloud:Dayu/ddosPolicyAttachment:DdosPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosPolicyAttachment resources.
type ddosPolicyAttachmentState struct {
	// ID of the policy.
	PolicyId *string `pulumi:"policyId"`
	// ID of the attached resource.
	ResourceId *string `pulumi:"resourceId"`
	// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
	ResourceType *string `pulumi:"resourceType"`
}

type DdosPolicyAttachmentState struct {
	// ID of the policy.
	PolicyId pulumi.StringPtrInput
	// ID of the attached resource.
	ResourceId pulumi.StringPtrInput
	// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
	ResourceType pulumi.StringPtrInput
}

func (DdosPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosPolicyAttachmentState)(nil)).Elem()
}

type ddosPolicyAttachmentArgs struct {
	// ID of the policy.
	PolicyId string `pulumi:"policyId"`
	// ID of the attached resource.
	ResourceId string `pulumi:"resourceId"`
	// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
	ResourceType string `pulumi:"resourceType"`
}

// The set of arguments for constructing a DdosPolicyAttachment resource.
type DdosPolicyAttachmentArgs struct {
	// ID of the policy.
	PolicyId pulumi.StringInput
	// ID of the attached resource.
	ResourceId pulumi.StringInput
	// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
	ResourceType pulumi.StringInput
}

func (DdosPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosPolicyAttachmentArgs)(nil)).Elem()
}

type DdosPolicyAttachmentInput interface {
	pulumi.Input

	ToDdosPolicyAttachmentOutput() DdosPolicyAttachmentOutput
	ToDdosPolicyAttachmentOutputWithContext(ctx context.Context) DdosPolicyAttachmentOutput
}

func (*DdosPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosPolicyAttachment)(nil)).Elem()
}

func (i *DdosPolicyAttachment) ToDdosPolicyAttachmentOutput() DdosPolicyAttachmentOutput {
	return i.ToDdosPolicyAttachmentOutputWithContext(context.Background())
}

func (i *DdosPolicyAttachment) ToDdosPolicyAttachmentOutputWithContext(ctx context.Context) DdosPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyAttachmentOutput)
}

// DdosPolicyAttachmentArrayInput is an input type that accepts DdosPolicyAttachmentArray and DdosPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `DdosPolicyAttachmentArrayInput` via:
//
//	DdosPolicyAttachmentArray{ DdosPolicyAttachmentArgs{...} }
type DdosPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToDdosPolicyAttachmentArrayOutput() DdosPolicyAttachmentArrayOutput
	ToDdosPolicyAttachmentArrayOutputWithContext(context.Context) DdosPolicyAttachmentArrayOutput
}

type DdosPolicyAttachmentArray []DdosPolicyAttachmentInput

func (DdosPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosPolicyAttachment)(nil)).Elem()
}

func (i DdosPolicyAttachmentArray) ToDdosPolicyAttachmentArrayOutput() DdosPolicyAttachmentArrayOutput {
	return i.ToDdosPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i DdosPolicyAttachmentArray) ToDdosPolicyAttachmentArrayOutputWithContext(ctx context.Context) DdosPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyAttachmentArrayOutput)
}

// DdosPolicyAttachmentMapInput is an input type that accepts DdosPolicyAttachmentMap and DdosPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `DdosPolicyAttachmentMapInput` via:
//
//	DdosPolicyAttachmentMap{ "key": DdosPolicyAttachmentArgs{...} }
type DdosPolicyAttachmentMapInput interface {
	pulumi.Input

	ToDdosPolicyAttachmentMapOutput() DdosPolicyAttachmentMapOutput
	ToDdosPolicyAttachmentMapOutputWithContext(context.Context) DdosPolicyAttachmentMapOutput
}

type DdosPolicyAttachmentMap map[string]DdosPolicyAttachmentInput

func (DdosPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosPolicyAttachment)(nil)).Elem()
}

func (i DdosPolicyAttachmentMap) ToDdosPolicyAttachmentMapOutput() DdosPolicyAttachmentMapOutput {
	return i.ToDdosPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i DdosPolicyAttachmentMap) ToDdosPolicyAttachmentMapOutputWithContext(ctx context.Context) DdosPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyAttachmentMapOutput)
}

type DdosPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (DdosPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosPolicyAttachment)(nil)).Elem()
}

func (o DdosPolicyAttachmentOutput) ToDdosPolicyAttachmentOutput() DdosPolicyAttachmentOutput {
	return o
}

func (o DdosPolicyAttachmentOutput) ToDdosPolicyAttachmentOutputWithContext(ctx context.Context) DdosPolicyAttachmentOutput {
	return o
}

// ID of the policy.
func (o DdosPolicyAttachmentOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyAttachment) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

// ID of the attached resource.
func (o DdosPolicyAttachmentOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyAttachment) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Type of the resource that the DDoS policy works for. Valid values are `bgpip`, `bgp`, `bgp-multip`, `net`.
func (o DdosPolicyAttachmentOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyAttachment) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

type DdosPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (DdosPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosPolicyAttachment)(nil)).Elem()
}

func (o DdosPolicyAttachmentArrayOutput) ToDdosPolicyAttachmentArrayOutput() DdosPolicyAttachmentArrayOutput {
	return o
}

func (o DdosPolicyAttachmentArrayOutput) ToDdosPolicyAttachmentArrayOutputWithContext(ctx context.Context) DdosPolicyAttachmentArrayOutput {
	return o
}

func (o DdosPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) DdosPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DdosPolicyAttachment {
		return vs[0].([]*DdosPolicyAttachment)[vs[1].(int)]
	}).(DdosPolicyAttachmentOutput)
}

type DdosPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (DdosPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosPolicyAttachment)(nil)).Elem()
}

func (o DdosPolicyAttachmentMapOutput) ToDdosPolicyAttachmentMapOutput() DdosPolicyAttachmentMapOutput {
	return o
}

func (o DdosPolicyAttachmentMapOutput) ToDdosPolicyAttachmentMapOutputWithContext(ctx context.Context) DdosPolicyAttachmentMapOutput {
	return o
}

func (o DdosPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) DdosPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DdosPolicyAttachment {
		return vs[0].(map[string]*DdosPolicyAttachment)[vs[1].(string)]
	}).(DdosPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyAttachmentInput)(nil)).Elem(), &DdosPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyAttachmentArrayInput)(nil)).Elem(), DdosPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyAttachmentMapInput)(nil)).Elem(), DdosPolicyAttachmentMap{})
	pulumi.RegisterOutputType(DdosPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(DdosPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(DdosPolicyAttachmentMapOutput{})
}
