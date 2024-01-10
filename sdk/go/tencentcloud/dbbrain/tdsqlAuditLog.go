// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dbbrain tdsqlAuditLog
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dbbrain.NewTdsqlAuditLog(ctx, "myLog", &Dbbrain.TdsqlAuditLogArgs{
//				EndTime: pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
//				Filter: &dbbrain.TdsqlAuditLogFilterArgs{
//					Hosts: pulumi.StringArray{
//						pulumi.String(fmt.Sprintf("%v%v", "%", "%")),
//						pulumi.String("127.0.0.1"),
//					},
//					Users: pulumi.StringArray{
//						pulumi.String("tf_test"),
//						pulumi.String("mysql"),
//					},
//				},
//				InstanceId:      pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
//				NodeRequestType: pulumi.String("dcdb"),
//				Product:         pulumi.String("dcdb"),
//				StartTime:       pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type TdsqlAuditLog struct {
	pulumi.CustomResourceState

	// Deadline time, such as `2019-09-11 10:13:14`.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Filter conditions. Logs can be filtered according to the filter conditions set.
	Filter TdsqlAuditLogFilterPtrOutput `pulumi:"filter"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Consistent with Product. For example: dcdb, mariadb.
	NodeRequestType pulumi.StringOutput `pulumi:"nodeRequestType"`
	// Service product type, supported values include: dcdb - cloud database Tdsql, mariadb - cloud database MariaDB for MariaDB..
	Product pulumi.StringOutput `pulumi:"product"`
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewTdsqlAuditLog registers a new resource with the given unique name, arguments, and options.
func NewTdsqlAuditLog(ctx *pulumi.Context,
	name string, args *TdsqlAuditLogArgs, opts ...pulumi.ResourceOption) (*TdsqlAuditLog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndTime == nil {
		return nil, errors.New("invalid value for required argument 'EndTime'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NodeRequestType == nil {
		return nil, errors.New("invalid value for required argument 'NodeRequestType'")
	}
	if args.Product == nil {
		return nil, errors.New("invalid value for required argument 'Product'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource TdsqlAuditLog
	err := ctx.RegisterResource("tencentcloud:Dbbrain/tdsqlAuditLog:TdsqlAuditLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTdsqlAuditLog gets an existing TdsqlAuditLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTdsqlAuditLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TdsqlAuditLogState, opts ...pulumi.ResourceOption) (*TdsqlAuditLog, error) {
	var resource TdsqlAuditLog
	err := ctx.ReadResource("tencentcloud:Dbbrain/tdsqlAuditLog:TdsqlAuditLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TdsqlAuditLog resources.
type tdsqlAuditLogState struct {
	// Deadline time, such as `2019-09-11 10:13:14`.
	EndTime *string `pulumi:"endTime"`
	// Filter conditions. Logs can be filtered according to the filter conditions set.
	Filter *TdsqlAuditLogFilter `pulumi:"filter"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Consistent with Product. For example: dcdb, mariadb.
	NodeRequestType *string `pulumi:"nodeRequestType"`
	// Service product type, supported values include: dcdb - cloud database Tdsql, mariadb - cloud database MariaDB for MariaDB..
	Product *string `pulumi:"product"`
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime *string `pulumi:"startTime"`
}

type TdsqlAuditLogState struct {
	// Deadline time, such as `2019-09-11 10:13:14`.
	EndTime pulumi.StringPtrInput
	// Filter conditions. Logs can be filtered according to the filter conditions set.
	Filter TdsqlAuditLogFilterPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// Consistent with Product. For example: dcdb, mariadb.
	NodeRequestType pulumi.StringPtrInput
	// Service product type, supported values include: dcdb - cloud database Tdsql, mariadb - cloud database MariaDB for MariaDB..
	Product pulumi.StringPtrInput
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime pulumi.StringPtrInput
}

func (TdsqlAuditLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*tdsqlAuditLogState)(nil)).Elem()
}

type tdsqlAuditLogArgs struct {
	// Deadline time, such as `2019-09-11 10:13:14`.
	EndTime string `pulumi:"endTime"`
	// Filter conditions. Logs can be filtered according to the filter conditions set.
	Filter *TdsqlAuditLogFilter `pulumi:"filter"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Consistent with Product. For example: dcdb, mariadb.
	NodeRequestType string `pulumi:"nodeRequestType"`
	// Service product type, supported values include: dcdb - cloud database Tdsql, mariadb - cloud database MariaDB for MariaDB..
	Product string `pulumi:"product"`
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime string `pulumi:"startTime"`
}

// The set of arguments for constructing a TdsqlAuditLog resource.
type TdsqlAuditLogArgs struct {
	// Deadline time, such as `2019-09-11 10:13:14`.
	EndTime pulumi.StringInput
	// Filter conditions. Logs can be filtered according to the filter conditions set.
	Filter TdsqlAuditLogFilterPtrInput
	// Instance ID.
	InstanceId pulumi.StringInput
	// Consistent with Product. For example: dcdb, mariadb.
	NodeRequestType pulumi.StringInput
	// Service product type, supported values include: dcdb - cloud database Tdsql, mariadb - cloud database MariaDB for MariaDB..
	Product pulumi.StringInput
	// Start time, such as `2019-09-10 12:13:14`.
	StartTime pulumi.StringInput
}

func (TdsqlAuditLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tdsqlAuditLogArgs)(nil)).Elem()
}

type TdsqlAuditLogInput interface {
	pulumi.Input

	ToTdsqlAuditLogOutput() TdsqlAuditLogOutput
	ToTdsqlAuditLogOutputWithContext(ctx context.Context) TdsqlAuditLogOutput
}

func (*TdsqlAuditLog) ElementType() reflect.Type {
	return reflect.TypeOf((**TdsqlAuditLog)(nil)).Elem()
}

func (i *TdsqlAuditLog) ToTdsqlAuditLogOutput() TdsqlAuditLogOutput {
	return i.ToTdsqlAuditLogOutputWithContext(context.Background())
}

func (i *TdsqlAuditLog) ToTdsqlAuditLogOutputWithContext(ctx context.Context) TdsqlAuditLogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TdsqlAuditLogOutput)
}

// TdsqlAuditLogArrayInput is an input type that accepts TdsqlAuditLogArray and TdsqlAuditLogArrayOutput values.
// You can construct a concrete instance of `TdsqlAuditLogArrayInput` via:
//
//	TdsqlAuditLogArray{ TdsqlAuditLogArgs{...} }
type TdsqlAuditLogArrayInput interface {
	pulumi.Input

	ToTdsqlAuditLogArrayOutput() TdsqlAuditLogArrayOutput
	ToTdsqlAuditLogArrayOutputWithContext(context.Context) TdsqlAuditLogArrayOutput
}

type TdsqlAuditLogArray []TdsqlAuditLogInput

func (TdsqlAuditLogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TdsqlAuditLog)(nil)).Elem()
}

func (i TdsqlAuditLogArray) ToTdsqlAuditLogArrayOutput() TdsqlAuditLogArrayOutput {
	return i.ToTdsqlAuditLogArrayOutputWithContext(context.Background())
}

func (i TdsqlAuditLogArray) ToTdsqlAuditLogArrayOutputWithContext(ctx context.Context) TdsqlAuditLogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TdsqlAuditLogArrayOutput)
}

// TdsqlAuditLogMapInput is an input type that accepts TdsqlAuditLogMap and TdsqlAuditLogMapOutput values.
// You can construct a concrete instance of `TdsqlAuditLogMapInput` via:
//
//	TdsqlAuditLogMap{ "key": TdsqlAuditLogArgs{...} }
type TdsqlAuditLogMapInput interface {
	pulumi.Input

	ToTdsqlAuditLogMapOutput() TdsqlAuditLogMapOutput
	ToTdsqlAuditLogMapOutputWithContext(context.Context) TdsqlAuditLogMapOutput
}

type TdsqlAuditLogMap map[string]TdsqlAuditLogInput

func (TdsqlAuditLogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TdsqlAuditLog)(nil)).Elem()
}

func (i TdsqlAuditLogMap) ToTdsqlAuditLogMapOutput() TdsqlAuditLogMapOutput {
	return i.ToTdsqlAuditLogMapOutputWithContext(context.Background())
}

func (i TdsqlAuditLogMap) ToTdsqlAuditLogMapOutputWithContext(ctx context.Context) TdsqlAuditLogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TdsqlAuditLogMapOutput)
}

type TdsqlAuditLogOutput struct{ *pulumi.OutputState }

func (TdsqlAuditLogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TdsqlAuditLog)(nil)).Elem()
}

func (o TdsqlAuditLogOutput) ToTdsqlAuditLogOutput() TdsqlAuditLogOutput {
	return o
}

func (o TdsqlAuditLogOutput) ToTdsqlAuditLogOutputWithContext(ctx context.Context) TdsqlAuditLogOutput {
	return o
}

// Deadline time, such as `2019-09-11 10:13:14`.
func (o TdsqlAuditLogOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TdsqlAuditLog) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Filter conditions. Logs can be filtered according to the filter conditions set.
func (o TdsqlAuditLogOutput) Filter() TdsqlAuditLogFilterPtrOutput {
	return o.ApplyT(func(v *TdsqlAuditLog) TdsqlAuditLogFilterPtrOutput { return v.Filter }).(TdsqlAuditLogFilterPtrOutput)
}

// Instance ID.
func (o TdsqlAuditLogOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TdsqlAuditLog) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Consistent with Product. For example: dcdb, mariadb.
func (o TdsqlAuditLogOutput) NodeRequestType() pulumi.StringOutput {
	return o.ApplyT(func(v *TdsqlAuditLog) pulumi.StringOutput { return v.NodeRequestType }).(pulumi.StringOutput)
}

// Service product type, supported values include: dcdb - cloud database Tdsql, mariadb - cloud database MariaDB for MariaDB..
func (o TdsqlAuditLogOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v *TdsqlAuditLog) pulumi.StringOutput { return v.Product }).(pulumi.StringOutput)
}

// Start time, such as `2019-09-10 12:13:14`.
func (o TdsqlAuditLogOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TdsqlAuditLog) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

type TdsqlAuditLogArrayOutput struct{ *pulumi.OutputState }

func (TdsqlAuditLogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TdsqlAuditLog)(nil)).Elem()
}

func (o TdsqlAuditLogArrayOutput) ToTdsqlAuditLogArrayOutput() TdsqlAuditLogArrayOutput {
	return o
}

func (o TdsqlAuditLogArrayOutput) ToTdsqlAuditLogArrayOutputWithContext(ctx context.Context) TdsqlAuditLogArrayOutput {
	return o
}

func (o TdsqlAuditLogArrayOutput) Index(i pulumi.IntInput) TdsqlAuditLogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TdsqlAuditLog {
		return vs[0].([]*TdsqlAuditLog)[vs[1].(int)]
	}).(TdsqlAuditLogOutput)
}

type TdsqlAuditLogMapOutput struct{ *pulumi.OutputState }

func (TdsqlAuditLogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TdsqlAuditLog)(nil)).Elem()
}

func (o TdsqlAuditLogMapOutput) ToTdsqlAuditLogMapOutput() TdsqlAuditLogMapOutput {
	return o
}

func (o TdsqlAuditLogMapOutput) ToTdsqlAuditLogMapOutputWithContext(ctx context.Context) TdsqlAuditLogMapOutput {
	return o
}

func (o TdsqlAuditLogMapOutput) MapIndex(k pulumi.StringInput) TdsqlAuditLogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TdsqlAuditLog {
		return vs[0].(map[string]*TdsqlAuditLog)[vs[1].(string)]
	}).(TdsqlAuditLogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TdsqlAuditLogInput)(nil)).Elem(), &TdsqlAuditLog{})
	pulumi.RegisterInputType(reflect.TypeOf((*TdsqlAuditLogArrayInput)(nil)).Elem(), TdsqlAuditLogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TdsqlAuditLogMapInput)(nil)).Elem(), TdsqlAuditLogMap{})
	pulumi.RegisterOutputType(TdsqlAuditLogOutput{})
	pulumi.RegisterOutputType(TdsqlAuditLogArrayOutput{})
	pulumi.RegisterOutputType(TdsqlAuditLogMapOutput{})
}
