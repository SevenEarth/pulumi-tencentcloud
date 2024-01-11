// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a PrivateDns zoneVpcAttachment
//
// > **NOTE:**  If you need to bind account A to account B's VPC resources, you need to first grant role authorization to account A.
//
// ## Example Usage
// ### Append VPC associated with private dns zone
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/PrivateDns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/PrivateDns"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleZone, err := PrivateDns.NewZone(ctx, "exampleZone", &PrivateDns.ZoneArgs{
// 			Domain:             pulumi.String("domain.com"),
// 			Remark:             pulumi.String("remark."),
// 			DnsForwardStatus:   pulumi.String("DISABLED"),
// 			CnameSpeedupStatus: pulumi.String("ENABLED"),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
// 			CidrBlock: pulumi.String("10.0.0.0/16"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = PrivateDns.NewZoneVpcAttachment(ctx, "exampleZoneVpcAttachment", &PrivateDns.ZoneVpcAttachmentArgs{
// 			ZoneId: exampleZone.ID(),
// 			VpcSet: &privatedns.ZoneVpcAttachmentVpcSetArgs{
// 				UniqVpcId: vpc.ID(),
// 				Region:    pulumi.String("ap-guangzhou"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Add VPC information for associated accounts in the private dns zone
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/PrivateDns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/PrivateDns"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := PrivateDns.NewZoneVpcAttachment(ctx, "example", &PrivateDns.ZoneVpcAttachmentArgs{
// 			ZoneId: pulumi.Any(tencentcloud_private_dns_zone.Example.Id),
// 			AccountVpcSet: &privatedns.ZoneVpcAttachmentAccountVpcSetArgs{
// 				UniqVpcId: pulumi.String("vpc-82znjzn3"),
// 				Region:    pulumi.String("ap-guangzhou"),
// 				Uin:       pulumi.String("100017155920"),
// 			},
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
// PrivateDns zone_vpc_attachment can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:PrivateDns/zoneVpcAttachment:ZoneVpcAttachment example zone-6t11lof0#vpc-jdx11z0t
// ```
type ZoneVpcAttachment struct {
	pulumi.CustomResourceState

	// New add account vpc info.
	AccountVpcSet ZoneVpcAttachmentAccountVpcSetPtrOutput `pulumi:"accountVpcSet"`
	// New add vpc info.
	VpcSet ZoneVpcAttachmentVpcSetPtrOutput `pulumi:"vpcSet"`
	// PrivateZone ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewZoneVpcAttachment registers a new resource with the given unique name, arguments, and options.
func NewZoneVpcAttachment(ctx *pulumi.Context,
	name string, args *ZoneVpcAttachmentArgs, opts ...pulumi.ResourceOption) (*ZoneVpcAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ZoneVpcAttachment
	err := ctx.RegisterResource("tencentcloud:PrivateDns/zoneVpcAttachment:ZoneVpcAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneVpcAttachment gets an existing ZoneVpcAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneVpcAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneVpcAttachmentState, opts ...pulumi.ResourceOption) (*ZoneVpcAttachment, error) {
	var resource ZoneVpcAttachment
	err := ctx.ReadResource("tencentcloud:PrivateDns/zoneVpcAttachment:ZoneVpcAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneVpcAttachment resources.
type zoneVpcAttachmentState struct {
	// New add account vpc info.
	AccountVpcSet *ZoneVpcAttachmentAccountVpcSet `pulumi:"accountVpcSet"`
	// New add vpc info.
	VpcSet *ZoneVpcAttachmentVpcSet `pulumi:"vpcSet"`
	// PrivateZone ID.
	ZoneId *string `pulumi:"zoneId"`
}

type ZoneVpcAttachmentState struct {
	// New add account vpc info.
	AccountVpcSet ZoneVpcAttachmentAccountVpcSetPtrInput
	// New add vpc info.
	VpcSet ZoneVpcAttachmentVpcSetPtrInput
	// PrivateZone ID.
	ZoneId pulumi.StringPtrInput
}

func (ZoneVpcAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneVpcAttachmentState)(nil)).Elem()
}

type zoneVpcAttachmentArgs struct {
	// New add account vpc info.
	AccountVpcSet *ZoneVpcAttachmentAccountVpcSet `pulumi:"accountVpcSet"`
	// New add vpc info.
	VpcSet *ZoneVpcAttachmentVpcSet `pulumi:"vpcSet"`
	// PrivateZone ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a ZoneVpcAttachment resource.
type ZoneVpcAttachmentArgs struct {
	// New add account vpc info.
	AccountVpcSet ZoneVpcAttachmentAccountVpcSetPtrInput
	// New add vpc info.
	VpcSet ZoneVpcAttachmentVpcSetPtrInput
	// PrivateZone ID.
	ZoneId pulumi.StringInput
}

func (ZoneVpcAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneVpcAttachmentArgs)(nil)).Elem()
}

type ZoneVpcAttachmentInput interface {
	pulumi.Input

	ToZoneVpcAttachmentOutput() ZoneVpcAttachmentOutput
	ToZoneVpcAttachmentOutputWithContext(ctx context.Context) ZoneVpcAttachmentOutput
}

func (*ZoneVpcAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneVpcAttachment)(nil)).Elem()
}

func (i *ZoneVpcAttachment) ToZoneVpcAttachmentOutput() ZoneVpcAttachmentOutput {
	return i.ToZoneVpcAttachmentOutputWithContext(context.Background())
}

func (i *ZoneVpcAttachment) ToZoneVpcAttachmentOutputWithContext(ctx context.Context) ZoneVpcAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVpcAttachmentOutput)
}

// ZoneVpcAttachmentArrayInput is an input type that accepts ZoneVpcAttachmentArray and ZoneVpcAttachmentArrayOutput values.
// You can construct a concrete instance of `ZoneVpcAttachmentArrayInput` via:
//
//          ZoneVpcAttachmentArray{ ZoneVpcAttachmentArgs{...} }
type ZoneVpcAttachmentArrayInput interface {
	pulumi.Input

	ToZoneVpcAttachmentArrayOutput() ZoneVpcAttachmentArrayOutput
	ToZoneVpcAttachmentArrayOutputWithContext(context.Context) ZoneVpcAttachmentArrayOutput
}

type ZoneVpcAttachmentArray []ZoneVpcAttachmentInput

func (ZoneVpcAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneVpcAttachment)(nil)).Elem()
}

func (i ZoneVpcAttachmentArray) ToZoneVpcAttachmentArrayOutput() ZoneVpcAttachmentArrayOutput {
	return i.ToZoneVpcAttachmentArrayOutputWithContext(context.Background())
}

func (i ZoneVpcAttachmentArray) ToZoneVpcAttachmentArrayOutputWithContext(ctx context.Context) ZoneVpcAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVpcAttachmentArrayOutput)
}

// ZoneVpcAttachmentMapInput is an input type that accepts ZoneVpcAttachmentMap and ZoneVpcAttachmentMapOutput values.
// You can construct a concrete instance of `ZoneVpcAttachmentMapInput` via:
//
//          ZoneVpcAttachmentMap{ "key": ZoneVpcAttachmentArgs{...} }
type ZoneVpcAttachmentMapInput interface {
	pulumi.Input

	ToZoneVpcAttachmentMapOutput() ZoneVpcAttachmentMapOutput
	ToZoneVpcAttachmentMapOutputWithContext(context.Context) ZoneVpcAttachmentMapOutput
}

type ZoneVpcAttachmentMap map[string]ZoneVpcAttachmentInput

func (ZoneVpcAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneVpcAttachment)(nil)).Elem()
}

func (i ZoneVpcAttachmentMap) ToZoneVpcAttachmentMapOutput() ZoneVpcAttachmentMapOutput {
	return i.ToZoneVpcAttachmentMapOutputWithContext(context.Background())
}

func (i ZoneVpcAttachmentMap) ToZoneVpcAttachmentMapOutputWithContext(ctx context.Context) ZoneVpcAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneVpcAttachmentMapOutput)
}

type ZoneVpcAttachmentOutput struct{ *pulumi.OutputState }

func (ZoneVpcAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneVpcAttachment)(nil)).Elem()
}

func (o ZoneVpcAttachmentOutput) ToZoneVpcAttachmentOutput() ZoneVpcAttachmentOutput {
	return o
}

func (o ZoneVpcAttachmentOutput) ToZoneVpcAttachmentOutputWithContext(ctx context.Context) ZoneVpcAttachmentOutput {
	return o
}

// New add account vpc info.
func (o ZoneVpcAttachmentOutput) AccountVpcSet() ZoneVpcAttachmentAccountVpcSetPtrOutput {
	return o.ApplyT(func(v *ZoneVpcAttachment) ZoneVpcAttachmentAccountVpcSetPtrOutput { return v.AccountVpcSet }).(ZoneVpcAttachmentAccountVpcSetPtrOutput)
}

// New add vpc info.
func (o ZoneVpcAttachmentOutput) VpcSet() ZoneVpcAttachmentVpcSetPtrOutput {
	return o.ApplyT(func(v *ZoneVpcAttachment) ZoneVpcAttachmentVpcSetPtrOutput { return v.VpcSet }).(ZoneVpcAttachmentVpcSetPtrOutput)
}

// PrivateZone ID.
func (o ZoneVpcAttachmentOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneVpcAttachment) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ZoneVpcAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ZoneVpcAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneVpcAttachment)(nil)).Elem()
}

func (o ZoneVpcAttachmentArrayOutput) ToZoneVpcAttachmentArrayOutput() ZoneVpcAttachmentArrayOutput {
	return o
}

func (o ZoneVpcAttachmentArrayOutput) ToZoneVpcAttachmentArrayOutputWithContext(ctx context.Context) ZoneVpcAttachmentArrayOutput {
	return o
}

func (o ZoneVpcAttachmentArrayOutput) Index(i pulumi.IntInput) ZoneVpcAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZoneVpcAttachment {
		return vs[0].([]*ZoneVpcAttachment)[vs[1].(int)]
	}).(ZoneVpcAttachmentOutput)
}

type ZoneVpcAttachmentMapOutput struct{ *pulumi.OutputState }

func (ZoneVpcAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneVpcAttachment)(nil)).Elem()
}

func (o ZoneVpcAttachmentMapOutput) ToZoneVpcAttachmentMapOutput() ZoneVpcAttachmentMapOutput {
	return o
}

func (o ZoneVpcAttachmentMapOutput) ToZoneVpcAttachmentMapOutputWithContext(ctx context.Context) ZoneVpcAttachmentMapOutput {
	return o
}

func (o ZoneVpcAttachmentMapOutput) MapIndex(k pulumi.StringInput) ZoneVpcAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZoneVpcAttachment {
		return vs[0].(map[string]*ZoneVpcAttachment)[vs[1].(string)]
	}).(ZoneVpcAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneVpcAttachmentInput)(nil)).Elem(), &ZoneVpcAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneVpcAttachmentArrayInput)(nil)).Elem(), ZoneVpcAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneVpcAttachmentMapInput)(nil)).Elem(), ZoneVpcAttachmentMap{})
	pulumi.RegisterOutputType(ZoneVpcAttachmentOutput{})
	pulumi.RegisterOutputType(ZoneVpcAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ZoneVpcAttachmentMapOutput{})
}