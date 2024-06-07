// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tdmq

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tdmqRocketmq environmentRole
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tdmq"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleRocketmqCluster, err := Tdmq.NewRocketmqCluster(ctx, "exampleRocketmqCluster", &Tdmq.RocketmqClusterArgs{
//				ClusterName: pulumi.String("tf_example"),
//				Remark:      pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			exampleRocketmqRole, err := Tdmq.NewRocketmqRole(ctx, "exampleRocketmqRole", &Tdmq.RocketmqRoleArgs{
//				RoleName:  pulumi.String("tf_example_role"),
//				Remark:    pulumi.String("remark."),
//				ClusterId: exampleRocketmqCluster.ClusterId,
//			})
//			if err != nil {
//				return err
//			}
//			exampleRocketmqNamespace, err := Tdmq.NewRocketmqNamespace(ctx, "exampleRocketmqNamespace", &Tdmq.RocketmqNamespaceArgs{
//				ClusterId:     exampleRocketmqCluster.ClusterId,
//				NamespaceName: pulumi.String("tf_example_namespace"),
//				Remark:        pulumi.String("remark."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Tdmq.NewRocketmqEnvironmentRole(ctx, "exampleRocketmqEnvironmentRole", &Tdmq.RocketmqEnvironmentRoleArgs{
//				EnvironmentName: exampleRocketmqNamespace.NamespaceName,
//				RoleName:        exampleRocketmqRole.RoleName,
//				Permissions: pulumi.StringArray{
//					pulumi.String("produce"),
//					pulumi.String("consume"),
//				},
//				ClusterId: exampleRocketmqCluster.ClusterId,
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
// tdmqRocketmq environment_role can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tdmq/rocketmqEnvironmentRole:RocketmqEnvironmentRole environment_role environmentRole_id
// ```
type RocketmqEnvironmentRole struct {
	pulumi.CustomResourceState

	// Cluster ID (required).
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Environment (namespace) name.
	EnvironmentName pulumi.StringOutput `pulumi:"environmentName"`
	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions pulumi.StringArrayOutput `pulumi:"permissions"`
	// Role Name.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
}

// NewRocketmqEnvironmentRole registers a new resource with the given unique name, arguments, and options.
func NewRocketmqEnvironmentRole(ctx *pulumi.Context,
	name string, args *RocketmqEnvironmentRoleArgs, opts ...pulumi.ResourceOption) (*RocketmqEnvironmentRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RocketmqEnvironmentRole
	err := ctx.RegisterResource("tencentcloud:Tdmq/rocketmqEnvironmentRole:RocketmqEnvironmentRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRocketmqEnvironmentRole gets an existing RocketmqEnvironmentRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRocketmqEnvironmentRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RocketmqEnvironmentRoleState, opts ...pulumi.ResourceOption) (*RocketmqEnvironmentRole, error) {
	var resource RocketmqEnvironmentRole
	err := ctx.ReadResource("tencentcloud:Tdmq/rocketmqEnvironmentRole:RocketmqEnvironmentRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RocketmqEnvironmentRole resources.
type rocketmqEnvironmentRoleState struct {
	// Cluster ID (required).
	ClusterId *string `pulumi:"clusterId"`
	// Environment (namespace) name.
	EnvironmentName *string `pulumi:"environmentName"`
	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []string `pulumi:"permissions"`
	// Role Name.
	RoleName *string `pulumi:"roleName"`
}

type RocketmqEnvironmentRoleState struct {
	// Cluster ID (required).
	ClusterId pulumi.StringPtrInput
	// Environment (namespace) name.
	EnvironmentName pulumi.StringPtrInput
	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions pulumi.StringArrayInput
	// Role Name.
	RoleName pulumi.StringPtrInput
}

func (RocketmqEnvironmentRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqEnvironmentRoleState)(nil)).Elem()
}

type rocketmqEnvironmentRoleArgs struct {
	// Cluster ID (required).
	ClusterId string `pulumi:"clusterId"`
	// Environment (namespace) name.
	EnvironmentName string `pulumi:"environmentName"`
	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []string `pulumi:"permissions"`
	// Role Name.
	RoleName string `pulumi:"roleName"`
}

// The set of arguments for constructing a RocketmqEnvironmentRole resource.
type RocketmqEnvironmentRoleArgs struct {
	// Cluster ID (required).
	ClusterId pulumi.StringInput
	// Environment (namespace) name.
	EnvironmentName pulumi.StringInput
	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions pulumi.StringArrayInput
	// Role Name.
	RoleName pulumi.StringInput
}

func (RocketmqEnvironmentRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rocketmqEnvironmentRoleArgs)(nil)).Elem()
}

type RocketmqEnvironmentRoleInput interface {
	pulumi.Input

	ToRocketmqEnvironmentRoleOutput() RocketmqEnvironmentRoleOutput
	ToRocketmqEnvironmentRoleOutputWithContext(ctx context.Context) RocketmqEnvironmentRoleOutput
}

func (*RocketmqEnvironmentRole) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqEnvironmentRole)(nil)).Elem()
}

func (i *RocketmqEnvironmentRole) ToRocketmqEnvironmentRoleOutput() RocketmqEnvironmentRoleOutput {
	return i.ToRocketmqEnvironmentRoleOutputWithContext(context.Background())
}

func (i *RocketmqEnvironmentRole) ToRocketmqEnvironmentRoleOutputWithContext(ctx context.Context) RocketmqEnvironmentRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqEnvironmentRoleOutput)
}

// RocketmqEnvironmentRoleArrayInput is an input type that accepts RocketmqEnvironmentRoleArray and RocketmqEnvironmentRoleArrayOutput values.
// You can construct a concrete instance of `RocketmqEnvironmentRoleArrayInput` via:
//
//	RocketmqEnvironmentRoleArray{ RocketmqEnvironmentRoleArgs{...} }
type RocketmqEnvironmentRoleArrayInput interface {
	pulumi.Input

	ToRocketmqEnvironmentRoleArrayOutput() RocketmqEnvironmentRoleArrayOutput
	ToRocketmqEnvironmentRoleArrayOutputWithContext(context.Context) RocketmqEnvironmentRoleArrayOutput
}

type RocketmqEnvironmentRoleArray []RocketmqEnvironmentRoleInput

func (RocketmqEnvironmentRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqEnvironmentRole)(nil)).Elem()
}

func (i RocketmqEnvironmentRoleArray) ToRocketmqEnvironmentRoleArrayOutput() RocketmqEnvironmentRoleArrayOutput {
	return i.ToRocketmqEnvironmentRoleArrayOutputWithContext(context.Background())
}

func (i RocketmqEnvironmentRoleArray) ToRocketmqEnvironmentRoleArrayOutputWithContext(ctx context.Context) RocketmqEnvironmentRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqEnvironmentRoleArrayOutput)
}

// RocketmqEnvironmentRoleMapInput is an input type that accepts RocketmqEnvironmentRoleMap and RocketmqEnvironmentRoleMapOutput values.
// You can construct a concrete instance of `RocketmqEnvironmentRoleMapInput` via:
//
//	RocketmqEnvironmentRoleMap{ "key": RocketmqEnvironmentRoleArgs{...} }
type RocketmqEnvironmentRoleMapInput interface {
	pulumi.Input

	ToRocketmqEnvironmentRoleMapOutput() RocketmqEnvironmentRoleMapOutput
	ToRocketmqEnvironmentRoleMapOutputWithContext(context.Context) RocketmqEnvironmentRoleMapOutput
}

type RocketmqEnvironmentRoleMap map[string]RocketmqEnvironmentRoleInput

func (RocketmqEnvironmentRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqEnvironmentRole)(nil)).Elem()
}

func (i RocketmqEnvironmentRoleMap) ToRocketmqEnvironmentRoleMapOutput() RocketmqEnvironmentRoleMapOutput {
	return i.ToRocketmqEnvironmentRoleMapOutputWithContext(context.Background())
}

func (i RocketmqEnvironmentRoleMap) ToRocketmqEnvironmentRoleMapOutputWithContext(ctx context.Context) RocketmqEnvironmentRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RocketmqEnvironmentRoleMapOutput)
}

type RocketmqEnvironmentRoleOutput struct{ *pulumi.OutputState }

func (RocketmqEnvironmentRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RocketmqEnvironmentRole)(nil)).Elem()
}

func (o RocketmqEnvironmentRoleOutput) ToRocketmqEnvironmentRoleOutput() RocketmqEnvironmentRoleOutput {
	return o
}

func (o RocketmqEnvironmentRoleOutput) ToRocketmqEnvironmentRoleOutputWithContext(ctx context.Context) RocketmqEnvironmentRoleOutput {
	return o
}

// Cluster ID (required).
func (o RocketmqEnvironmentRoleOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqEnvironmentRole) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Environment (namespace) name.
func (o RocketmqEnvironmentRoleOutput) EnvironmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqEnvironmentRole) pulumi.StringOutput { return v.EnvironmentName }).(pulumi.StringOutput)
}

// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
func (o RocketmqEnvironmentRoleOutput) Permissions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RocketmqEnvironmentRole) pulumi.StringArrayOutput { return v.Permissions }).(pulumi.StringArrayOutput)
}

// Role Name.
func (o RocketmqEnvironmentRoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *RocketmqEnvironmentRole) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

type RocketmqEnvironmentRoleArrayOutput struct{ *pulumi.OutputState }

func (RocketmqEnvironmentRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RocketmqEnvironmentRole)(nil)).Elem()
}

func (o RocketmqEnvironmentRoleArrayOutput) ToRocketmqEnvironmentRoleArrayOutput() RocketmqEnvironmentRoleArrayOutput {
	return o
}

func (o RocketmqEnvironmentRoleArrayOutput) ToRocketmqEnvironmentRoleArrayOutputWithContext(ctx context.Context) RocketmqEnvironmentRoleArrayOutput {
	return o
}

func (o RocketmqEnvironmentRoleArrayOutput) Index(i pulumi.IntInput) RocketmqEnvironmentRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RocketmqEnvironmentRole {
		return vs[0].([]*RocketmqEnvironmentRole)[vs[1].(int)]
	}).(RocketmqEnvironmentRoleOutput)
}

type RocketmqEnvironmentRoleMapOutput struct{ *pulumi.OutputState }

func (RocketmqEnvironmentRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RocketmqEnvironmentRole)(nil)).Elem()
}

func (o RocketmqEnvironmentRoleMapOutput) ToRocketmqEnvironmentRoleMapOutput() RocketmqEnvironmentRoleMapOutput {
	return o
}

func (o RocketmqEnvironmentRoleMapOutput) ToRocketmqEnvironmentRoleMapOutputWithContext(ctx context.Context) RocketmqEnvironmentRoleMapOutput {
	return o
}

func (o RocketmqEnvironmentRoleMapOutput) MapIndex(k pulumi.StringInput) RocketmqEnvironmentRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RocketmqEnvironmentRole {
		return vs[0].(map[string]*RocketmqEnvironmentRole)[vs[1].(string)]
	}).(RocketmqEnvironmentRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqEnvironmentRoleInput)(nil)).Elem(), &RocketmqEnvironmentRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqEnvironmentRoleArrayInput)(nil)).Elem(), RocketmqEnvironmentRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RocketmqEnvironmentRoleMapInput)(nil)).Elem(), RocketmqEnvironmentRoleMap{})
	pulumi.RegisterOutputType(RocketmqEnvironmentRoleOutput{})
	pulumi.RegisterOutputType(RocketmqEnvironmentRoleArrayOutput{})
	pulumi.RegisterOutputType(RocketmqEnvironmentRoleMapOutput{})
}
