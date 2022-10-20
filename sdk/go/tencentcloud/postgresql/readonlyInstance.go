// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create postgresql readonly instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Postgresql.NewReadonlyInstance(ctx, "foo", &Postgresql.ReadonlyInstanceArgs{
//				AutoRenewFlag:      pulumi.Int(0),
//				DbVersion:          pulumi.String("10.4"),
//				InstanceChargeType: pulumi.String("POSTPAID_BY_HOUR"),
//				MasterDbInstanceId: pulumi.String("postgres-j4pm65id"),
//				Memory:             pulumi.Int(4),
//				NeedSupportIpv6:    pulumi.Int(0),
//				ProjectId:          pulumi.Int(0),
//				SecurityGroupsIds: pulumi.StringArray{
//					pulumi.String("sg-fefj5n6r"),
//				},
//				Storage:  pulumi.Int(250),
//				SubnetId: pulumi.String("subnet-enm92y0m"),
//				VpcId:    pulumi.String("vpc-86v957zb"),
//				Zone:     pulumi.String("ap-guangzhou-6"),
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
// postgresql readonly instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance foo pgro-bcqx8b9a
//
// ```
type ReadonlyInstance struct {
	pulumi.CustomResourceState

	// Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
	AutoRenewFlag pulumi.IntPtrOutput `pulumi:"autoRenewFlag"`
	// Create time of the postgresql instance.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// PostgreSQL kernel version, which must be the same as that of the primary instance.
	DbVersion pulumi.StringOutput `pulumi:"dbVersion"`
	// instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
	InstanceChargeType pulumi.StringPtrOutput `pulumi:"instanceChargeType"`
	// ID of the primary instance to which the read-only replica belongs.
	MasterDbInstanceId pulumi.StringOutput `pulumi:"masterDbInstanceId"`
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `Postgresql.getSpecinfos` provides.
	Memory pulumi.IntOutput `pulumi:"memory"`
	// Instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
	NeedSupportIpv6 pulumi.IntPtrOutput `pulumi:"needSupportIpv6"`
	// Project ID.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// ID of security group.
	SecurityGroupsIds pulumi.StringArrayOutput `pulumi:"securityGroupsIds"`
	// Instance storage capacity in GB.
	Storage pulumi.IntOutput `pulumi:"storage"`
	// VPC subnet ID.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewReadonlyInstance registers a new resource with the given unique name, arguments, and options.
func NewReadonlyInstance(ctx *pulumi.Context,
	name string, args *ReadonlyInstanceArgs, opts ...pulumi.ResourceOption) (*ReadonlyInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbVersion'")
	}
	if args.MasterDbInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'MasterDbInstanceId'")
	}
	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.SecurityGroupsIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupsIds'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ReadonlyInstance
	err := ctx.RegisterResource("tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReadonlyInstance gets an existing ReadonlyInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReadonlyInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReadonlyInstanceState, opts ...pulumi.ResourceOption) (*ReadonlyInstance, error) {
	var resource ReadonlyInstance
	err := ctx.ReadResource("tencentcloud:Postgresql/readonlyInstance:ReadonlyInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReadonlyInstance resources.
type readonlyInstanceState struct {
	// Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// Create time of the postgresql instance.
	CreateTime *string `pulumi:"createTime"`
	// PostgreSQL kernel version, which must be the same as that of the primary instance.
	DbVersion *string `pulumi:"dbVersion"`
	// instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// ID of the primary instance to which the read-only replica belongs.
	MasterDbInstanceId *string `pulumi:"masterDbInstanceId"`
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `Postgresql.getSpecinfos` provides.
	Memory *int `pulumi:"memory"`
	// Instance name.
	Name *string `pulumi:"name"`
	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
	NeedSupportIpv6 *int `pulumi:"needSupportIpv6"`
	// Project ID.
	ProjectId *int `pulumi:"projectId"`
	// ID of security group.
	SecurityGroupsIds []string `pulumi:"securityGroupsIds"`
	// Instance storage capacity in GB.
	Storage *int `pulumi:"storage"`
	// VPC subnet ID.
	SubnetId *string `pulumi:"subnetId"`
	// VPC ID.
	VpcId *string `pulumi:"vpcId"`
	// Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
	Zone *string `pulumi:"zone"`
}

type ReadonlyInstanceState struct {
	// Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
	AutoRenewFlag pulumi.IntPtrInput
	// Create time of the postgresql instance.
	CreateTime pulumi.StringPtrInput
	// PostgreSQL kernel version, which must be the same as that of the primary instance.
	DbVersion pulumi.StringPtrInput
	// instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
	InstanceChargeType pulumi.StringPtrInput
	// ID of the primary instance to which the read-only replica belongs.
	MasterDbInstanceId pulumi.StringPtrInput
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `Postgresql.getSpecinfos` provides.
	Memory pulumi.IntPtrInput
	// Instance name.
	Name pulumi.StringPtrInput
	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
	NeedSupportIpv6 pulumi.IntPtrInput
	// Project ID.
	ProjectId pulumi.IntPtrInput
	// ID of security group.
	SecurityGroupsIds pulumi.StringArrayInput
	// Instance storage capacity in GB.
	Storage pulumi.IntPtrInput
	// VPC subnet ID.
	SubnetId pulumi.StringPtrInput
	// VPC ID.
	VpcId pulumi.StringPtrInput
	// Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
	Zone pulumi.StringPtrInput
}

func (ReadonlyInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*readonlyInstanceState)(nil)).Elem()
}

type readonlyInstanceArgs struct {
	// Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
	AutoRenewFlag *int `pulumi:"autoRenewFlag"`
	// PostgreSQL kernel version, which must be the same as that of the primary instance.
	DbVersion string `pulumi:"dbVersion"`
	// instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// ID of the primary instance to which the read-only replica belongs.
	MasterDbInstanceId string `pulumi:"masterDbInstanceId"`
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `Postgresql.getSpecinfos` provides.
	Memory int `pulumi:"memory"`
	// Instance name.
	Name *string `pulumi:"name"`
	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
	NeedSupportIpv6 *int `pulumi:"needSupportIpv6"`
	// Project ID.
	ProjectId int `pulumi:"projectId"`
	// ID of security group.
	SecurityGroupsIds []string `pulumi:"securityGroupsIds"`
	// Instance storage capacity in GB.
	Storage int `pulumi:"storage"`
	// VPC subnet ID.
	SubnetId string `pulumi:"subnetId"`
	// VPC ID.
	VpcId string `pulumi:"vpcId"`
	// Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a ReadonlyInstance resource.
type ReadonlyInstanceArgs struct {
	// Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
	AutoRenewFlag pulumi.IntPtrInput
	// PostgreSQL kernel version, which must be the same as that of the primary instance.
	DbVersion pulumi.StringInput
	// instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
	InstanceChargeType pulumi.StringPtrInput
	// ID of the primary instance to which the read-only replica belongs.
	MasterDbInstanceId pulumi.StringInput
	// Memory size(in GB). Allowed value must be larger than `memory` that data source `Postgresql.getSpecinfos` provides.
	Memory pulumi.IntInput
	// Instance name.
	Name pulumi.StringPtrInput
	// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
	NeedSupportIpv6 pulumi.IntPtrInput
	// Project ID.
	ProjectId pulumi.IntInput
	// ID of security group.
	SecurityGroupsIds pulumi.StringArrayInput
	// Instance storage capacity in GB.
	Storage pulumi.IntInput
	// VPC subnet ID.
	SubnetId pulumi.StringInput
	// VPC ID.
	VpcId pulumi.StringInput
	// Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
	Zone pulumi.StringInput
}

func (ReadonlyInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*readonlyInstanceArgs)(nil)).Elem()
}

type ReadonlyInstanceInput interface {
	pulumi.Input

	ToReadonlyInstanceOutput() ReadonlyInstanceOutput
	ToReadonlyInstanceOutputWithContext(ctx context.Context) ReadonlyInstanceOutput
}

func (*ReadonlyInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadonlyInstance)(nil)).Elem()
}

func (i *ReadonlyInstance) ToReadonlyInstanceOutput() ReadonlyInstanceOutput {
	return i.ToReadonlyInstanceOutputWithContext(context.Background())
}

func (i *ReadonlyInstance) ToReadonlyInstanceOutputWithContext(ctx context.Context) ReadonlyInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyInstanceOutput)
}

// ReadonlyInstanceArrayInput is an input type that accepts ReadonlyInstanceArray and ReadonlyInstanceArrayOutput values.
// You can construct a concrete instance of `ReadonlyInstanceArrayInput` via:
//
//	ReadonlyInstanceArray{ ReadonlyInstanceArgs{...} }
type ReadonlyInstanceArrayInput interface {
	pulumi.Input

	ToReadonlyInstanceArrayOutput() ReadonlyInstanceArrayOutput
	ToReadonlyInstanceArrayOutputWithContext(context.Context) ReadonlyInstanceArrayOutput
}

type ReadonlyInstanceArray []ReadonlyInstanceInput

func (ReadonlyInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadonlyInstance)(nil)).Elem()
}

func (i ReadonlyInstanceArray) ToReadonlyInstanceArrayOutput() ReadonlyInstanceArrayOutput {
	return i.ToReadonlyInstanceArrayOutputWithContext(context.Background())
}

func (i ReadonlyInstanceArray) ToReadonlyInstanceArrayOutputWithContext(ctx context.Context) ReadonlyInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyInstanceArrayOutput)
}

// ReadonlyInstanceMapInput is an input type that accepts ReadonlyInstanceMap and ReadonlyInstanceMapOutput values.
// You can construct a concrete instance of `ReadonlyInstanceMapInput` via:
//
//	ReadonlyInstanceMap{ "key": ReadonlyInstanceArgs{...} }
type ReadonlyInstanceMapInput interface {
	pulumi.Input

	ToReadonlyInstanceMapOutput() ReadonlyInstanceMapOutput
	ToReadonlyInstanceMapOutputWithContext(context.Context) ReadonlyInstanceMapOutput
}

type ReadonlyInstanceMap map[string]ReadonlyInstanceInput

func (ReadonlyInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadonlyInstance)(nil)).Elem()
}

func (i ReadonlyInstanceMap) ToReadonlyInstanceMapOutput() ReadonlyInstanceMapOutput {
	return i.ToReadonlyInstanceMapOutputWithContext(context.Background())
}

func (i ReadonlyInstanceMap) ToReadonlyInstanceMapOutputWithContext(ctx context.Context) ReadonlyInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReadonlyInstanceMapOutput)
}

type ReadonlyInstanceOutput struct{ *pulumi.OutputState }

func (ReadonlyInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReadonlyInstance)(nil)).Elem()
}

func (o ReadonlyInstanceOutput) ToReadonlyInstanceOutput() ReadonlyInstanceOutput {
	return o
}

func (o ReadonlyInstanceOutput) ToReadonlyInstanceOutputWithContext(ctx context.Context) ReadonlyInstanceOutput {
	return o
}

// Renewal flag. Valid values: 0 (manual renewal), 1 (auto-renewal). Default value: 0.
func (o ReadonlyInstanceOutput) AutoRenewFlag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.AutoRenewFlag }).(pulumi.IntPtrOutput)
}

// Create time of the postgresql instance.
func (o ReadonlyInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// PostgreSQL kernel version, which must be the same as that of the primary instance.
func (o ReadonlyInstanceOutput) DbVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.DbVersion }).(pulumi.StringOutput)
}

// instance billing mode. Valid values: PREPAID (monthly subscription), POSTPAID_BY_HOUR (pay-as-you-go).
func (o ReadonlyInstanceOutput) InstanceChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringPtrOutput { return v.InstanceChargeType }).(pulumi.StringPtrOutput)
}

// ID of the primary instance to which the read-only replica belongs.
func (o ReadonlyInstanceOutput) MasterDbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.MasterDbInstanceId }).(pulumi.StringOutput)
}

// Memory size(in GB). Allowed value must be larger than `memory` that data source `Postgresql.getSpecinfos` provides.
func (o ReadonlyInstanceOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

// Instance name.
func (o ReadonlyInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether to support IPv6 address access. Valid values: 1 (yes), 0 (no).
func (o ReadonlyInstanceOutput) NeedSupportIpv6() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntPtrOutput { return v.NeedSupportIpv6 }).(pulumi.IntPtrOutput)
}

// Project ID.
func (o ReadonlyInstanceOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// ID of security group.
func (o ReadonlyInstanceOutput) SecurityGroupsIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringArrayOutput { return v.SecurityGroupsIds }).(pulumi.StringArrayOutput)
}

// Instance storage capacity in GB.
func (o ReadonlyInstanceOutput) Storage() pulumi.IntOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.IntOutput { return v.Storage }).(pulumi.IntOutput)
}

// VPC subnet ID.
func (o ReadonlyInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// VPC ID.
func (o ReadonlyInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Availability zone ID, which can be obtained through the Zone field in the returned value of the DescribeZones API.
func (o ReadonlyInstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ReadonlyInstance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ReadonlyInstanceArrayOutput struct{ *pulumi.OutputState }

func (ReadonlyInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReadonlyInstance)(nil)).Elem()
}

func (o ReadonlyInstanceArrayOutput) ToReadonlyInstanceArrayOutput() ReadonlyInstanceArrayOutput {
	return o
}

func (o ReadonlyInstanceArrayOutput) ToReadonlyInstanceArrayOutputWithContext(ctx context.Context) ReadonlyInstanceArrayOutput {
	return o
}

func (o ReadonlyInstanceArrayOutput) Index(i pulumi.IntInput) ReadonlyInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReadonlyInstance {
		return vs[0].([]*ReadonlyInstance)[vs[1].(int)]
	}).(ReadonlyInstanceOutput)
}

type ReadonlyInstanceMapOutput struct{ *pulumi.OutputState }

func (ReadonlyInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReadonlyInstance)(nil)).Elem()
}

func (o ReadonlyInstanceMapOutput) ToReadonlyInstanceMapOutput() ReadonlyInstanceMapOutput {
	return o
}

func (o ReadonlyInstanceMapOutput) ToReadonlyInstanceMapOutputWithContext(ctx context.Context) ReadonlyInstanceMapOutput {
	return o
}

func (o ReadonlyInstanceMapOutput) MapIndex(k pulumi.StringInput) ReadonlyInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReadonlyInstance {
		return vs[0].(map[string]*ReadonlyInstance)[vs[1].(string)]
	}).(ReadonlyInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyInstanceInput)(nil)).Elem(), &ReadonlyInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyInstanceArrayInput)(nil)).Elem(), ReadonlyInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReadonlyInstanceMapInput)(nil)).Elem(), ReadonlyInstanceMap{})
	pulumi.RegisterOutputType(ReadonlyInstanceOutput{})
	pulumi.RegisterOutputType(ReadonlyInstanceArrayOutput{})
	pulumi.RegisterOutputType(ReadonlyInstanceMapOutput{})
}
