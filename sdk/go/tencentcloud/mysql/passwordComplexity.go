// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mysql passwordComplexity
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Security"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "cdb",
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
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				IsMulticast:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			securityGroup, err := Security.NewGroup(ctx, "securityGroup", &Security.GroupArgs{
//				Description: pulumi.String("mysql test"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := Mysql.NewInstance(ctx, "exampleInstance", &Mysql.InstanceArgs{
//				InternetService:  pulumi.Int(1),
//				EngineVersion:    pulumi.String("5.7"),
//				ChargeType:       pulumi.String("POSTPAID"),
//				RootPassword:     pulumi.String("PassWord123"),
//				SlaveDeployMode:  pulumi.Int(0),
//				AvailabilityZone: pulumi.String(zones.Zones[0].Name),
//				SlaveSyncMode:    pulumi.Int(1),
//				InstanceName:     pulumi.String("tf-example-mysql"),
//				MemSize:          pulumi.Int(4000),
//				VolumeSize:       pulumi.Int(200),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				IntranetPort:     pulumi.Int(3306),
//				SecurityGroups: pulumi.StringArray{
//					securityGroup.ID(),
//				},
//				Tags: pulumi.AnyMap{
//					"name": pulumi.Any("test"),
//				},
//				Parameters: pulumi.AnyMap{
//					"character_set_server": pulumi.Any("utf8"),
//					"max_connections":      pulumi.Any("1000"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Mysql.NewPasswordComplexity(ctx, "examplePasswordComplexity", &Mysql.PasswordComplexityArgs{
//				InstanceId: exampleInstance.ID(),
//				ParamLists: mysql.PasswordComplexityParamListArray{
//					&mysql.PasswordComplexityParamListArgs{
//						Name:         pulumi.String("validate_password_length"),
//						CurrentValue: pulumi.String("8"),
//					},
//					&mysql.PasswordComplexityParamListArgs{
//						Name:         pulumi.String("validate_password_mixed_case_count"),
//						CurrentValue: pulumi.String("2"),
//					},
//					&mysql.PasswordComplexityParamListArgs{
//						Name:         pulumi.String("validate_password_number_count"),
//						CurrentValue: pulumi.String("2"),
//					},
//					&mysql.PasswordComplexityParamListArgs{
//						Name:         pulumi.String("validate_password_special_char_count"),
//						CurrentValue: pulumi.String("2"),
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
type PasswordComplexity struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validatePasswordPolic`, `validatePasswordLengt` `validatePasswordMixedCaseCoun`, `validatePasswordNumberCoun`, `validatePasswordSpecialCharCoun`.
	ParamLists PasswordComplexityParamListArrayOutput `pulumi:"paramLists"`
}

// NewPasswordComplexity registers a new resource with the given unique name, arguments, and options.
func NewPasswordComplexity(ctx *pulumi.Context,
	name string, args *PasswordComplexityArgs, opts ...pulumi.ResourceOption) (*PasswordComplexity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PasswordComplexity
	err := ctx.RegisterResource("tencentcloud:Mysql/passwordComplexity:PasswordComplexity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPasswordComplexity gets an existing PasswordComplexity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPasswordComplexity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordComplexityState, opts ...pulumi.ResourceOption) (*PasswordComplexity, error) {
	var resource PasswordComplexity
	err := ctx.ReadResource("tencentcloud:Mysql/passwordComplexity:PasswordComplexity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PasswordComplexity resources.
type passwordComplexityState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validatePasswordPolic`, `validatePasswordLengt` `validatePasswordMixedCaseCoun`, `validatePasswordNumberCoun`, `validatePasswordSpecialCharCoun`.
	ParamLists []PasswordComplexityParamList `pulumi:"paramLists"`
}

type PasswordComplexityState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validatePasswordPolic`, `validatePasswordLengt` `validatePasswordMixedCaseCoun`, `validatePasswordNumberCoun`, `validatePasswordSpecialCharCoun`.
	ParamLists PasswordComplexityParamListArrayInput
}

func (PasswordComplexityState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordComplexityState)(nil)).Elem()
}

type passwordComplexityArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validatePasswordPolic`, `validatePasswordLengt` `validatePasswordMixedCaseCoun`, `validatePasswordNumberCoun`, `validatePasswordSpecialCharCoun`.
	ParamLists []PasswordComplexityParamList `pulumi:"paramLists"`
}

// The set of arguments for constructing a PasswordComplexity resource.
type PasswordComplexityArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
	// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validatePasswordPolic`, `validatePasswordLengt` `validatePasswordMixedCaseCoun`, `validatePasswordNumberCoun`, `validatePasswordSpecialCharCoun`.
	ParamLists PasswordComplexityParamListArrayInput
}

func (PasswordComplexityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordComplexityArgs)(nil)).Elem()
}

type PasswordComplexityInput interface {
	pulumi.Input

	ToPasswordComplexityOutput() PasswordComplexityOutput
	ToPasswordComplexityOutputWithContext(ctx context.Context) PasswordComplexityOutput
}

func (*PasswordComplexity) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordComplexity)(nil)).Elem()
}

func (i *PasswordComplexity) ToPasswordComplexityOutput() PasswordComplexityOutput {
	return i.ToPasswordComplexityOutputWithContext(context.Background())
}

func (i *PasswordComplexity) ToPasswordComplexityOutputWithContext(ctx context.Context) PasswordComplexityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordComplexityOutput)
}

// PasswordComplexityArrayInput is an input type that accepts PasswordComplexityArray and PasswordComplexityArrayOutput values.
// You can construct a concrete instance of `PasswordComplexityArrayInput` via:
//
//	PasswordComplexityArray{ PasswordComplexityArgs{...} }
type PasswordComplexityArrayInput interface {
	pulumi.Input

	ToPasswordComplexityArrayOutput() PasswordComplexityArrayOutput
	ToPasswordComplexityArrayOutputWithContext(context.Context) PasswordComplexityArrayOutput
}

type PasswordComplexityArray []PasswordComplexityInput

func (PasswordComplexityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordComplexity)(nil)).Elem()
}

func (i PasswordComplexityArray) ToPasswordComplexityArrayOutput() PasswordComplexityArrayOutput {
	return i.ToPasswordComplexityArrayOutputWithContext(context.Background())
}

func (i PasswordComplexityArray) ToPasswordComplexityArrayOutputWithContext(ctx context.Context) PasswordComplexityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordComplexityArrayOutput)
}

// PasswordComplexityMapInput is an input type that accepts PasswordComplexityMap and PasswordComplexityMapOutput values.
// You can construct a concrete instance of `PasswordComplexityMapInput` via:
//
//	PasswordComplexityMap{ "key": PasswordComplexityArgs{...} }
type PasswordComplexityMapInput interface {
	pulumi.Input

	ToPasswordComplexityMapOutput() PasswordComplexityMapOutput
	ToPasswordComplexityMapOutputWithContext(context.Context) PasswordComplexityMapOutput
}

type PasswordComplexityMap map[string]PasswordComplexityInput

func (PasswordComplexityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordComplexity)(nil)).Elem()
}

func (i PasswordComplexityMap) ToPasswordComplexityMapOutput() PasswordComplexityMapOutput {
	return i.ToPasswordComplexityMapOutputWithContext(context.Background())
}

func (i PasswordComplexityMap) ToPasswordComplexityMapOutputWithContext(ctx context.Context) PasswordComplexityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordComplexityMapOutput)
}

type PasswordComplexityOutput struct{ *pulumi.OutputState }

func (PasswordComplexityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordComplexity)(nil)).Elem()
}

func (o PasswordComplexityOutput) ToPasswordComplexityOutput() PasswordComplexityOutput {
	return o
}

func (o PasswordComplexityOutput) ToPasswordComplexityOutputWithContext(ctx context.Context) PasswordComplexityOutput {
	return o
}

// Instance ID.
func (o PasswordComplexityOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *PasswordComplexity) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// List of parameters to be modified. Every element is a combination of `Name` (parameter name) and `CurrentValue` (new value). Valid values for `Name` of version 8.0: `validate_password.policy`, `validate_password.lengt`, `validate_password.mixed_case_coun`, `validate_password.number_coun`, `validate_password.special_char_count`. Valid values for `Name` of version 5.6 and 5.7: `validatePasswordPolic`, `validatePasswordLengt` `validatePasswordMixedCaseCoun`, `validatePasswordNumberCoun`, `validatePasswordSpecialCharCoun`.
func (o PasswordComplexityOutput) ParamLists() PasswordComplexityParamListArrayOutput {
	return o.ApplyT(func(v *PasswordComplexity) PasswordComplexityParamListArrayOutput { return v.ParamLists }).(PasswordComplexityParamListArrayOutput)
}

type PasswordComplexityArrayOutput struct{ *pulumi.OutputState }

func (PasswordComplexityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordComplexity)(nil)).Elem()
}

func (o PasswordComplexityArrayOutput) ToPasswordComplexityArrayOutput() PasswordComplexityArrayOutput {
	return o
}

func (o PasswordComplexityArrayOutput) ToPasswordComplexityArrayOutputWithContext(ctx context.Context) PasswordComplexityArrayOutput {
	return o
}

func (o PasswordComplexityArrayOutput) Index(i pulumi.IntInput) PasswordComplexityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PasswordComplexity {
		return vs[0].([]*PasswordComplexity)[vs[1].(int)]
	}).(PasswordComplexityOutput)
}

type PasswordComplexityMapOutput struct{ *pulumi.OutputState }

func (PasswordComplexityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordComplexity)(nil)).Elem()
}

func (o PasswordComplexityMapOutput) ToPasswordComplexityMapOutput() PasswordComplexityMapOutput {
	return o
}

func (o PasswordComplexityMapOutput) ToPasswordComplexityMapOutputWithContext(ctx context.Context) PasswordComplexityMapOutput {
	return o
}

func (o PasswordComplexityMapOutput) MapIndex(k pulumi.StringInput) PasswordComplexityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PasswordComplexity {
		return vs[0].(map[string]*PasswordComplexity)[vs[1].(string)]
	}).(PasswordComplexityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordComplexityInput)(nil)).Elem(), &PasswordComplexity{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordComplexityArrayInput)(nil)).Elem(), PasswordComplexityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordComplexityMapInput)(nil)).Elem(), PasswordComplexityMap{})
	pulumi.RegisterOutputType(PasswordComplexityOutput{})
	pulumi.RegisterOutputType(PasswordComplexityArrayOutput{})
	pulumi.RegisterOutputType(PasswordComplexityMapOutput{})
}
