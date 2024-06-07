// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a mysql restartDbInstancesOperation
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
//				Tags: pulumi.Map{
//					"name": pulumi.Any("test"),
//				},
//				Parameters: pulumi.Map{
//					"character_set_server": pulumi.Any("utf8"),
//					"max_connections":      pulumi.Any("1000"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Mysql.NewRestartDbInstancesOperation(ctx, "exampleRestartDbInstancesOperation", &Mysql.RestartDbInstancesOperationArgs{
//				InstanceId: exampleInstance.ID(),
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
// mysql restart_db_instances_operation can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation restart_db_instances_operation restart_db_instances_operation_id
// ```
type RestartDbInstancesOperation struct {
	pulumi.CustomResourceState

	// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Instance status.
	Status pulumi.IntOutput `pulumi:"status"`
}

// NewRestartDbInstancesOperation registers a new resource with the given unique name, arguments, and options.
func NewRestartDbInstancesOperation(ctx *pulumi.Context,
	name string, args *RestartDbInstancesOperationArgs, opts ...pulumi.ResourceOption) (*RestartDbInstancesOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RestartDbInstancesOperation
	err := ctx.RegisterResource("tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRestartDbInstancesOperation gets an existing RestartDbInstancesOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRestartDbInstancesOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RestartDbInstancesOperationState, opts ...pulumi.ResourceOption) (*RestartDbInstancesOperation, error) {
	var resource RestartDbInstancesOperation
	err := ctx.ReadResource("tencentcloud:Mysql/restartDbInstancesOperation:RestartDbInstancesOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RestartDbInstancesOperation resources.
type restartDbInstancesOperationState struct {
	// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
	InstanceId *string `pulumi:"instanceId"`
	// Instance status.
	Status *int `pulumi:"status"`
}

type RestartDbInstancesOperationState struct {
	// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringPtrInput
	// Instance status.
	Status pulumi.IntPtrInput
}

func (RestartDbInstancesOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*restartDbInstancesOperationState)(nil)).Elem()
}

type restartDbInstancesOperationArgs struct {
	// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a RestartDbInstancesOperation resource.
type RestartDbInstancesOperationArgs struct {
	// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
	InstanceId pulumi.StringInput
}

func (RestartDbInstancesOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restartDbInstancesOperationArgs)(nil)).Elem()
}

type RestartDbInstancesOperationInput interface {
	pulumi.Input

	ToRestartDbInstancesOperationOutput() RestartDbInstancesOperationOutput
	ToRestartDbInstancesOperationOutputWithContext(ctx context.Context) RestartDbInstancesOperationOutput
}

func (*RestartDbInstancesOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**RestartDbInstancesOperation)(nil)).Elem()
}

func (i *RestartDbInstancesOperation) ToRestartDbInstancesOperationOutput() RestartDbInstancesOperationOutput {
	return i.ToRestartDbInstancesOperationOutputWithContext(context.Background())
}

func (i *RestartDbInstancesOperation) ToRestartDbInstancesOperationOutputWithContext(ctx context.Context) RestartDbInstancesOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartDbInstancesOperationOutput)
}

// RestartDbInstancesOperationArrayInput is an input type that accepts RestartDbInstancesOperationArray and RestartDbInstancesOperationArrayOutput values.
// You can construct a concrete instance of `RestartDbInstancesOperationArrayInput` via:
//
//	RestartDbInstancesOperationArray{ RestartDbInstancesOperationArgs{...} }
type RestartDbInstancesOperationArrayInput interface {
	pulumi.Input

	ToRestartDbInstancesOperationArrayOutput() RestartDbInstancesOperationArrayOutput
	ToRestartDbInstancesOperationArrayOutputWithContext(context.Context) RestartDbInstancesOperationArrayOutput
}

type RestartDbInstancesOperationArray []RestartDbInstancesOperationInput

func (RestartDbInstancesOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestartDbInstancesOperation)(nil)).Elem()
}

func (i RestartDbInstancesOperationArray) ToRestartDbInstancesOperationArrayOutput() RestartDbInstancesOperationArrayOutput {
	return i.ToRestartDbInstancesOperationArrayOutputWithContext(context.Background())
}

func (i RestartDbInstancesOperationArray) ToRestartDbInstancesOperationArrayOutputWithContext(ctx context.Context) RestartDbInstancesOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartDbInstancesOperationArrayOutput)
}

// RestartDbInstancesOperationMapInput is an input type that accepts RestartDbInstancesOperationMap and RestartDbInstancesOperationMapOutput values.
// You can construct a concrete instance of `RestartDbInstancesOperationMapInput` via:
//
//	RestartDbInstancesOperationMap{ "key": RestartDbInstancesOperationArgs{...} }
type RestartDbInstancesOperationMapInput interface {
	pulumi.Input

	ToRestartDbInstancesOperationMapOutput() RestartDbInstancesOperationMapOutput
	ToRestartDbInstancesOperationMapOutputWithContext(context.Context) RestartDbInstancesOperationMapOutput
}

type RestartDbInstancesOperationMap map[string]RestartDbInstancesOperationInput

func (RestartDbInstancesOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestartDbInstancesOperation)(nil)).Elem()
}

func (i RestartDbInstancesOperationMap) ToRestartDbInstancesOperationMapOutput() RestartDbInstancesOperationMapOutput {
	return i.ToRestartDbInstancesOperationMapOutputWithContext(context.Background())
}

func (i RestartDbInstancesOperationMap) ToRestartDbInstancesOperationMapOutputWithContext(ctx context.Context) RestartDbInstancesOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestartDbInstancesOperationMapOutput)
}

type RestartDbInstancesOperationOutput struct{ *pulumi.OutputState }

func (RestartDbInstancesOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestartDbInstancesOperation)(nil)).Elem()
}

func (o RestartDbInstancesOperationOutput) ToRestartDbInstancesOperationOutput() RestartDbInstancesOperationOutput {
	return o
}

func (o RestartDbInstancesOperationOutput) ToRestartDbInstancesOperationOutputWithContext(ctx context.Context) RestartDbInstancesOperationOutput {
	return o
}

// An array of instance ID in the format: cdb-c1nl9rpv, which is the same as the instance ID displayed on the cloud database console page.
func (o RestartDbInstancesOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RestartDbInstancesOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Instance status.
func (o RestartDbInstancesOperationOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *RestartDbInstancesOperation) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

type RestartDbInstancesOperationArrayOutput struct{ *pulumi.OutputState }

func (RestartDbInstancesOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RestartDbInstancesOperation)(nil)).Elem()
}

func (o RestartDbInstancesOperationArrayOutput) ToRestartDbInstancesOperationArrayOutput() RestartDbInstancesOperationArrayOutput {
	return o
}

func (o RestartDbInstancesOperationArrayOutput) ToRestartDbInstancesOperationArrayOutputWithContext(ctx context.Context) RestartDbInstancesOperationArrayOutput {
	return o
}

func (o RestartDbInstancesOperationArrayOutput) Index(i pulumi.IntInput) RestartDbInstancesOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RestartDbInstancesOperation {
		return vs[0].([]*RestartDbInstancesOperation)[vs[1].(int)]
	}).(RestartDbInstancesOperationOutput)
}

type RestartDbInstancesOperationMapOutput struct{ *pulumi.OutputState }

func (RestartDbInstancesOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RestartDbInstancesOperation)(nil)).Elem()
}

func (o RestartDbInstancesOperationMapOutput) ToRestartDbInstancesOperationMapOutput() RestartDbInstancesOperationMapOutput {
	return o
}

func (o RestartDbInstancesOperationMapOutput) ToRestartDbInstancesOperationMapOutputWithContext(ctx context.Context) RestartDbInstancesOperationMapOutput {
	return o
}

func (o RestartDbInstancesOperationMapOutput) MapIndex(k pulumi.StringInput) RestartDbInstancesOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RestartDbInstancesOperation {
		return vs[0].(map[string]*RestartDbInstancesOperation)[vs[1].(string)]
	}).(RestartDbInstancesOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestartDbInstancesOperationInput)(nil)).Elem(), &RestartDbInstancesOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestartDbInstancesOperationArrayInput)(nil)).Elem(), RestartDbInstancesOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RestartDbInstancesOperationMapInput)(nil)).Elem(), RestartDbInstancesOperationMap{})
	pulumi.RegisterOutputType(RestartDbInstancesOperationOutput{})
	pulumi.RegisterOutputType(RestartDbInstancesOperationArrayOutput{})
	pulumi.RegisterOutputType(RestartDbInstancesOperationMapOutput{})
}
