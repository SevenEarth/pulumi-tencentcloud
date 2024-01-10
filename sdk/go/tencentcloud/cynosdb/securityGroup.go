// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cynosdb securityGroup
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cynosdb.NewSecurityGroup(ctx, "test", &Cynosdb.SecurityGroupArgs{
//				ClusterId:         pulumi.String("cynosdbmysql-bws8h88b"),
//				InstanceGroupType: pulumi.String("RO"),
//				SecurityGroupIds: pulumi.StringArray{
//					pulumi.String("sg-baxfiao5"),
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
//
// ## Import
//
// cynosdb security_group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cynosdb/securityGroup:SecurityGroup security_group ${cluster_id}#${instance_group_type}
//
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// Cluster id.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Instance group type. Available values:
	// -`HA` - HA group;
	// -`RO` - Read-only group;
	// -`ALL` - HA and RO group.
	InstanceGroupType pulumi.StringOutput `pulumi:"instanceGroupType"`
	// A list of security group IDs to be modified, an array of one or more security group IDs.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceGroupType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceGroupType'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SecurityGroup
	err := ctx.RegisterResource("tencentcloud:Cynosdb/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("tencentcloud:Cynosdb/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// Cluster id.
	ClusterId *string `pulumi:"clusterId"`
	// Instance group type. Available values:
	// -`HA` - HA group;
	// -`RO` - Read-only group;
	// -`ALL` - HA and RO group.
	InstanceGroupType *string `pulumi:"instanceGroupType"`
	// A list of security group IDs to be modified, an array of one or more security group IDs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
}

type SecurityGroupState struct {
	// Cluster id.
	ClusterId pulumi.StringPtrInput
	// Instance group type. Available values:
	// -`HA` - HA group;
	// -`RO` - Read-only group;
	// -`ALL` - HA and RO group.
	InstanceGroupType pulumi.StringPtrInput
	// A list of security group IDs to be modified, an array of one or more security group IDs.
	SecurityGroupIds pulumi.StringArrayInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// Cluster id.
	ClusterId string `pulumi:"clusterId"`
	// Instance group type. Available values:
	// -`HA` - HA group;
	// -`RO` - Read-only group;
	// -`ALL` - HA and RO group.
	InstanceGroupType string `pulumi:"instanceGroupType"`
	// A list of security group IDs to be modified, an array of one or more security group IDs.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// Cluster id.
	ClusterId pulumi.StringInput
	// Instance group type. Available values:
	// -`HA` - HA group;
	// -`RO` - Read-only group;
	// -`ALL` - HA and RO group.
	InstanceGroupType pulumi.StringInput
	// A list of security group IDs to be modified, an array of one or more security group IDs.
	SecurityGroupIds pulumi.StringArrayInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (*SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (i *SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

// SecurityGroupArrayInput is an input type that accepts SecurityGroupArray and SecurityGroupArrayOutput values.
// You can construct a concrete instance of `SecurityGroupArrayInput` via:
//
//	SecurityGroupArray{ SecurityGroupArgs{...} }
type SecurityGroupArrayInput interface {
	pulumi.Input

	ToSecurityGroupArrayOutput() SecurityGroupArrayOutput
	ToSecurityGroupArrayOutputWithContext(context.Context) SecurityGroupArrayOutput
}

type SecurityGroupArray []SecurityGroupInput

func (SecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return i.ToSecurityGroupArrayOutputWithContext(context.Background())
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupArrayOutput)
}

// SecurityGroupMapInput is an input type that accepts SecurityGroupMap and SecurityGroupMapOutput values.
// You can construct a concrete instance of `SecurityGroupMapInput` via:
//
//	SecurityGroupMap{ "key": SecurityGroupArgs{...} }
type SecurityGroupMapInput interface {
	pulumi.Input

	ToSecurityGroupMapOutput() SecurityGroupMapOutput
	ToSecurityGroupMapOutputWithContext(context.Context) SecurityGroupMapOutput
}

type SecurityGroupMap map[string]SecurityGroupInput

func (SecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupMap) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return i.ToSecurityGroupMapOutputWithContext(context.Background())
}

func (i SecurityGroupMap) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupMapOutput)
}

type SecurityGroupOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

// Cluster id.
func (o SecurityGroupOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Instance group type. Available values:
// -`HA` - HA group;
// -`RO` - Read-only group;
// -`ALL` - HA and RO group.
func (o SecurityGroupOutput) InstanceGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.InstanceGroupType }).(pulumi.StringOutput)
}

// A list of security group IDs to be modified, an array of one or more security group IDs.
func (o SecurityGroupOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

type SecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) Index(i pulumi.IntInput) SecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].([]*SecurityGroup)[vs[1].(int)]
	}).(SecurityGroupOutput)
}

type SecurityGroupMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].(map[string]*SecurityGroup)[vs[1].(string)]
	}).(SecurityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInput)(nil)).Elem(), &SecurityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupArrayInput)(nil)).Elem(), SecurityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupMapInput)(nil)).Elem(), SecurityGroupMap{})
	pulumi.RegisterOutputType(SecurityGroupOutput{})
	pulumi.RegisterOutputType(SecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupMapOutput{})
}
