// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dcdb switchDbInstanceHaOperation
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dcdb.NewSwitchDbInstanceHaOperation(ctx, "switchOperation", &Dcdb.SwitchDbInstanceHaOperationArgs{
// 			InstanceId: pulumi.Any(local.Dcdb_id),
// 			Zone:       pulumi.String("ap-guangzhou-4"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SwitchDbInstanceHaOperation struct {
	pulumi.CustomResourceState

	// Instance ID in the format of tdsqlshard-ow728lmc.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewSwitchDbInstanceHaOperation registers a new resource with the given unique name, arguments, and options.
func NewSwitchDbInstanceHaOperation(ctx *pulumi.Context,
	name string, args *SwitchDbInstanceHaOperationArgs, opts ...pulumi.ResourceOption) (*SwitchDbInstanceHaOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SwitchDbInstanceHaOperation
	err := ctx.RegisterResource("tencentcloud:Dcdb/switchDbInstanceHaOperation:SwitchDbInstanceHaOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSwitchDbInstanceHaOperation gets an existing SwitchDbInstanceHaOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSwitchDbInstanceHaOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SwitchDbInstanceHaOperationState, opts ...pulumi.ResourceOption) (*SwitchDbInstanceHaOperation, error) {
	var resource SwitchDbInstanceHaOperation
	err := ctx.ReadResource("tencentcloud:Dcdb/switchDbInstanceHaOperation:SwitchDbInstanceHaOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SwitchDbInstanceHaOperation resources.
type switchDbInstanceHaOperationState struct {
	// Instance ID in the format of tdsqlshard-ow728lmc.
	InstanceId *string `pulumi:"instanceId"`
	// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.
	Zone *string `pulumi:"zone"`
}

type SwitchDbInstanceHaOperationState struct {
	// Instance ID in the format of tdsqlshard-ow728lmc.
	InstanceId pulumi.StringPtrInput
	// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.
	Zone pulumi.StringPtrInput
}

func (SwitchDbInstanceHaOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*switchDbInstanceHaOperationState)(nil)).Elem()
}

type switchDbInstanceHaOperationArgs struct {
	// Instance ID in the format of tdsqlshard-ow728lmc.
	InstanceId string `pulumi:"instanceId"`
	// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a SwitchDbInstanceHaOperation resource.
type SwitchDbInstanceHaOperationArgs struct {
	// Instance ID in the format of tdsqlshard-ow728lmc.
	InstanceId pulumi.StringInput
	// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.
	Zone pulumi.StringInput
}

func (SwitchDbInstanceHaOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*switchDbInstanceHaOperationArgs)(nil)).Elem()
}

type SwitchDbInstanceHaOperationInput interface {
	pulumi.Input

	ToSwitchDbInstanceHaOperationOutput() SwitchDbInstanceHaOperationOutput
	ToSwitchDbInstanceHaOperationOutputWithContext(ctx context.Context) SwitchDbInstanceHaOperationOutput
}

func (*SwitchDbInstanceHaOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchDbInstanceHaOperation)(nil)).Elem()
}

func (i *SwitchDbInstanceHaOperation) ToSwitchDbInstanceHaOperationOutput() SwitchDbInstanceHaOperationOutput {
	return i.ToSwitchDbInstanceHaOperationOutputWithContext(context.Background())
}

func (i *SwitchDbInstanceHaOperation) ToSwitchDbInstanceHaOperationOutputWithContext(ctx context.Context) SwitchDbInstanceHaOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchDbInstanceHaOperationOutput)
}

// SwitchDbInstanceHaOperationArrayInput is an input type that accepts SwitchDbInstanceHaOperationArray and SwitchDbInstanceHaOperationArrayOutput values.
// You can construct a concrete instance of `SwitchDbInstanceHaOperationArrayInput` via:
//
//          SwitchDbInstanceHaOperationArray{ SwitchDbInstanceHaOperationArgs{...} }
type SwitchDbInstanceHaOperationArrayInput interface {
	pulumi.Input

	ToSwitchDbInstanceHaOperationArrayOutput() SwitchDbInstanceHaOperationArrayOutput
	ToSwitchDbInstanceHaOperationArrayOutputWithContext(context.Context) SwitchDbInstanceHaOperationArrayOutput
}

type SwitchDbInstanceHaOperationArray []SwitchDbInstanceHaOperationInput

func (SwitchDbInstanceHaOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchDbInstanceHaOperation)(nil)).Elem()
}

func (i SwitchDbInstanceHaOperationArray) ToSwitchDbInstanceHaOperationArrayOutput() SwitchDbInstanceHaOperationArrayOutput {
	return i.ToSwitchDbInstanceHaOperationArrayOutputWithContext(context.Background())
}

func (i SwitchDbInstanceHaOperationArray) ToSwitchDbInstanceHaOperationArrayOutputWithContext(ctx context.Context) SwitchDbInstanceHaOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchDbInstanceHaOperationArrayOutput)
}

// SwitchDbInstanceHaOperationMapInput is an input type that accepts SwitchDbInstanceHaOperationMap and SwitchDbInstanceHaOperationMapOutput values.
// You can construct a concrete instance of `SwitchDbInstanceHaOperationMapInput` via:
//
//          SwitchDbInstanceHaOperationMap{ "key": SwitchDbInstanceHaOperationArgs{...} }
type SwitchDbInstanceHaOperationMapInput interface {
	pulumi.Input

	ToSwitchDbInstanceHaOperationMapOutput() SwitchDbInstanceHaOperationMapOutput
	ToSwitchDbInstanceHaOperationMapOutputWithContext(context.Context) SwitchDbInstanceHaOperationMapOutput
}

type SwitchDbInstanceHaOperationMap map[string]SwitchDbInstanceHaOperationInput

func (SwitchDbInstanceHaOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchDbInstanceHaOperation)(nil)).Elem()
}

func (i SwitchDbInstanceHaOperationMap) ToSwitchDbInstanceHaOperationMapOutput() SwitchDbInstanceHaOperationMapOutput {
	return i.ToSwitchDbInstanceHaOperationMapOutputWithContext(context.Background())
}

func (i SwitchDbInstanceHaOperationMap) ToSwitchDbInstanceHaOperationMapOutputWithContext(ctx context.Context) SwitchDbInstanceHaOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SwitchDbInstanceHaOperationMapOutput)
}

type SwitchDbInstanceHaOperationOutput struct{ *pulumi.OutputState }

func (SwitchDbInstanceHaOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SwitchDbInstanceHaOperation)(nil)).Elem()
}

func (o SwitchDbInstanceHaOperationOutput) ToSwitchDbInstanceHaOperationOutput() SwitchDbInstanceHaOperationOutput {
	return o
}

func (o SwitchDbInstanceHaOperationOutput) ToSwitchDbInstanceHaOperationOutputWithContext(ctx context.Context) SwitchDbInstanceHaOperationOutput {
	return o
}

// Instance ID in the format of tdsqlshard-ow728lmc.
func (o SwitchDbInstanceHaOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchDbInstanceHaOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.
func (o SwitchDbInstanceHaOperationOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *SwitchDbInstanceHaOperation) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type SwitchDbInstanceHaOperationArrayOutput struct{ *pulumi.OutputState }

func (SwitchDbInstanceHaOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SwitchDbInstanceHaOperation)(nil)).Elem()
}

func (o SwitchDbInstanceHaOperationArrayOutput) ToSwitchDbInstanceHaOperationArrayOutput() SwitchDbInstanceHaOperationArrayOutput {
	return o
}

func (o SwitchDbInstanceHaOperationArrayOutput) ToSwitchDbInstanceHaOperationArrayOutputWithContext(ctx context.Context) SwitchDbInstanceHaOperationArrayOutput {
	return o
}

func (o SwitchDbInstanceHaOperationArrayOutput) Index(i pulumi.IntInput) SwitchDbInstanceHaOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SwitchDbInstanceHaOperation {
		return vs[0].([]*SwitchDbInstanceHaOperation)[vs[1].(int)]
	}).(SwitchDbInstanceHaOperationOutput)
}

type SwitchDbInstanceHaOperationMapOutput struct{ *pulumi.OutputState }

func (SwitchDbInstanceHaOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SwitchDbInstanceHaOperation)(nil)).Elem()
}

func (o SwitchDbInstanceHaOperationMapOutput) ToSwitchDbInstanceHaOperationMapOutput() SwitchDbInstanceHaOperationMapOutput {
	return o
}

func (o SwitchDbInstanceHaOperationMapOutput) ToSwitchDbInstanceHaOperationMapOutputWithContext(ctx context.Context) SwitchDbInstanceHaOperationMapOutput {
	return o
}

func (o SwitchDbInstanceHaOperationMapOutput) MapIndex(k pulumi.StringInput) SwitchDbInstanceHaOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SwitchDbInstanceHaOperation {
		return vs[0].(map[string]*SwitchDbInstanceHaOperation)[vs[1].(string)]
	}).(SwitchDbInstanceHaOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchDbInstanceHaOperationInput)(nil)).Elem(), &SwitchDbInstanceHaOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchDbInstanceHaOperationArrayInput)(nil)).Elem(), SwitchDbInstanceHaOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SwitchDbInstanceHaOperationMapInput)(nil)).Elem(), SwitchDbInstanceHaOperationMap{})
	pulumi.RegisterOutputType(SwitchDbInstanceHaOperationOutput{})
	pulumi.RegisterOutputType(SwitchDbInstanceHaOperationArrayOutput{})
	pulumi.RegisterOutputType(SwitchDbInstanceHaOperationMapOutput{})
}
