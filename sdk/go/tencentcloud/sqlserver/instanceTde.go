// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sqlserver instanceTde
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "sqlserver",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[4].Name),
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				IsMulticast:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
//				Description: pulumi.String("desc."),
//			})
//			if err != nil {
//				return err
//			}
//			exampleBasicInstance, err := Sqlserver.NewBasicInstance(ctx, "exampleBasicInstance", &Sqlserver.BasicInstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[4].Name),
//				ChargeType:       pulumi.String("POSTPAID_BY_HOUR"),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				ProjectId:        pulumi.Int(0),
//				Memory:           pulumi.Int(4),
//				Storage:          pulumi.Int(100),
//				Cpu:              pulumi.Int(2),
//				MachineType:      pulumi.String("CLOUD_PREMIUM"),
//				MaintenanceWeekSets: pulumi.IntArray{
//					pulumi.Int(1),
//					pulumi.Int(2),
//					pulumi.Int(3),
//				},
//				MaintenanceStartTime: pulumi.String("09:00"),
//				MaintenanceTimeSpan:  pulumi.Int(3),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.AnyMap{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Sqlserver.NewInstanceTde(ctx, "exampleInstanceTde", &Sqlserver.InstanceTdeArgs{
//				InstanceId:             exampleBasicInstance.ID(),
//				CertificateAttribution: pulumi.String("self"),
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
// sqlserver instance_tde can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Sqlserver/instanceTde:InstanceTde example mssql-farjz9tz
//
// ```
type InstanceTde struct {
	pulumi.CustomResourceState

	// Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
	CertificateAttribution pulumi.StringOutput `pulumi:"certificateAttribution"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Other referenced main account IDs, required when CertificateAttribute is others.
	QuoteUin pulumi.StringPtrOutput `pulumi:"quoteUin"`
}

// NewInstanceTde registers a new resource with the given unique name, arguments, and options.
func NewInstanceTde(ctx *pulumi.Context,
	name string, args *InstanceTdeArgs, opts ...pulumi.ResourceOption) (*InstanceTde, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAttribution == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAttribution'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceTde
	err := ctx.RegisterResource("tencentcloud:Sqlserver/instanceTde:InstanceTde", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceTde gets an existing InstanceTde resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceTde(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceTdeState, opts ...pulumi.ResourceOption) (*InstanceTde, error) {
	var resource InstanceTde
	err := ctx.ReadResource("tencentcloud:Sqlserver/instanceTde:InstanceTde", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceTde resources.
type instanceTdeState struct {
	// Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
	CertificateAttribution *string `pulumi:"certificateAttribution"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Other referenced main account IDs, required when CertificateAttribute is others.
	QuoteUin *string `pulumi:"quoteUin"`
}

type InstanceTdeState struct {
	// Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
	CertificateAttribution pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// Other referenced main account IDs, required when CertificateAttribute is others.
	QuoteUin pulumi.StringPtrInput
}

func (InstanceTdeState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceTdeState)(nil)).Elem()
}

type instanceTdeArgs struct {
	// Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
	CertificateAttribution string `pulumi:"certificateAttribution"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Other referenced main account IDs, required when CertificateAttribute is others.
	QuoteUin *string `pulumi:"quoteUin"`
}

// The set of arguments for constructing a InstanceTde resource.
type InstanceTdeArgs struct {
	// Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
	CertificateAttribution pulumi.StringInput
	// Instance ID.
	InstanceId pulumi.StringInput
	// Other referenced main account IDs, required when CertificateAttribute is others.
	QuoteUin pulumi.StringPtrInput
}

func (InstanceTdeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceTdeArgs)(nil)).Elem()
}

type InstanceTdeInput interface {
	pulumi.Input

	ToInstanceTdeOutput() InstanceTdeOutput
	ToInstanceTdeOutputWithContext(ctx context.Context) InstanceTdeOutput
}

func (*InstanceTde) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTde)(nil)).Elem()
}

func (i *InstanceTde) ToInstanceTdeOutput() InstanceTdeOutput {
	return i.ToInstanceTdeOutputWithContext(context.Background())
}

func (i *InstanceTde) ToInstanceTdeOutputWithContext(ctx context.Context) InstanceTdeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTdeOutput)
}

// InstanceTdeArrayInput is an input type that accepts InstanceTdeArray and InstanceTdeArrayOutput values.
// You can construct a concrete instance of `InstanceTdeArrayInput` via:
//
//	InstanceTdeArray{ InstanceTdeArgs{...} }
type InstanceTdeArrayInput interface {
	pulumi.Input

	ToInstanceTdeArrayOutput() InstanceTdeArrayOutput
	ToInstanceTdeArrayOutputWithContext(context.Context) InstanceTdeArrayOutput
}

type InstanceTdeArray []InstanceTdeInput

func (InstanceTdeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceTde)(nil)).Elem()
}

func (i InstanceTdeArray) ToInstanceTdeArrayOutput() InstanceTdeArrayOutput {
	return i.ToInstanceTdeArrayOutputWithContext(context.Background())
}

func (i InstanceTdeArray) ToInstanceTdeArrayOutputWithContext(ctx context.Context) InstanceTdeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTdeArrayOutput)
}

// InstanceTdeMapInput is an input type that accepts InstanceTdeMap and InstanceTdeMapOutput values.
// You can construct a concrete instance of `InstanceTdeMapInput` via:
//
//	InstanceTdeMap{ "key": InstanceTdeArgs{...} }
type InstanceTdeMapInput interface {
	pulumi.Input

	ToInstanceTdeMapOutput() InstanceTdeMapOutput
	ToInstanceTdeMapOutputWithContext(context.Context) InstanceTdeMapOutput
}

type InstanceTdeMap map[string]InstanceTdeInput

func (InstanceTdeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceTde)(nil)).Elem()
}

func (i InstanceTdeMap) ToInstanceTdeMapOutput() InstanceTdeMapOutput {
	return i.ToInstanceTdeMapOutputWithContext(context.Background())
}

func (i InstanceTdeMap) ToInstanceTdeMapOutputWithContext(ctx context.Context) InstanceTdeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceTdeMapOutput)
}

type InstanceTdeOutput struct{ *pulumi.OutputState }

func (InstanceTdeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTde)(nil)).Elem()
}

func (o InstanceTdeOutput) ToInstanceTdeOutput() InstanceTdeOutput {
	return o
}

func (o InstanceTdeOutput) ToInstanceTdeOutputWithContext(ctx context.Context) InstanceTdeOutput {
	return o
}

// Certificate attribution. self- means to use the account's own certificate, others- means to refer to the certificate of other accounts, and the default is self.
func (o InstanceTdeOutput) CertificateAttribution() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceTde) pulumi.StringOutput { return v.CertificateAttribution }).(pulumi.StringOutput)
}

// Instance ID.
func (o InstanceTdeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceTde) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Other referenced main account IDs, required when CertificateAttribute is others.
func (o InstanceTdeOutput) QuoteUin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceTde) pulumi.StringPtrOutput { return v.QuoteUin }).(pulumi.StringPtrOutput)
}

type InstanceTdeArrayOutput struct{ *pulumi.OutputState }

func (InstanceTdeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceTde)(nil)).Elem()
}

func (o InstanceTdeArrayOutput) ToInstanceTdeArrayOutput() InstanceTdeArrayOutput {
	return o
}

func (o InstanceTdeArrayOutput) ToInstanceTdeArrayOutputWithContext(ctx context.Context) InstanceTdeArrayOutput {
	return o
}

func (o InstanceTdeArrayOutput) Index(i pulumi.IntInput) InstanceTdeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceTde {
		return vs[0].([]*InstanceTde)[vs[1].(int)]
	}).(InstanceTdeOutput)
}

type InstanceTdeMapOutput struct{ *pulumi.OutputState }

func (InstanceTdeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceTde)(nil)).Elem()
}

func (o InstanceTdeMapOutput) ToInstanceTdeMapOutput() InstanceTdeMapOutput {
	return o
}

func (o InstanceTdeMapOutput) ToInstanceTdeMapOutputWithContext(ctx context.Context) InstanceTdeMapOutput {
	return o
}

func (o InstanceTdeMapOutput) MapIndex(k pulumi.StringInput) InstanceTdeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceTde {
		return vs[0].(map[string]*InstanceTde)[vs[1].(string)]
	}).(InstanceTdeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTdeInput)(nil)).Elem(), &InstanceTde{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTdeArrayInput)(nil)).Elem(), InstanceTdeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTdeMapInput)(nil)).Elem(), InstanceTdeMap{})
	pulumi.RegisterOutputType(InstanceTdeOutput{})
	pulumi.RegisterOutputType(InstanceTdeArrayOutput{})
	pulumi.RegisterOutputType(InstanceTdeMapOutput{})
}
