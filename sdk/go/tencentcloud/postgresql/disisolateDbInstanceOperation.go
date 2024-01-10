// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a postgresql disisolateDbInstanceOperation
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
//			_, err := Postgresql.NewDisisolateDbInstanceOperation(ctx, "disisolateDbInstanceOperation", &Postgresql.DisisolateDbInstanceOperationArgs{
//				DbInstanceIdSets: pulumi.StringArray{
//					pulumi.Any(local.Pgsql_id),
//				},
//				Period:      pulumi.Int(1),
//				AutoVoucher: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DisisolateDbInstanceOperation struct {
	pulumi.CustomResourceState

	// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
	AutoVoucher pulumi.BoolPtrOutput `pulumi:"autoVoucher"`
	// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
	DbInstanceIdSets pulumi.StringArrayOutput `pulumi:"dbInstanceIdSets"`
	// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Voucher ID list.
	VoucherIds pulumi.StringArrayOutput `pulumi:"voucherIds"`
}

// NewDisisolateDbInstanceOperation registers a new resource with the given unique name, arguments, and options.
func NewDisisolateDbInstanceOperation(ctx *pulumi.Context,
	name string, args *DisisolateDbInstanceOperationArgs, opts ...pulumi.ResourceOption) (*DisisolateDbInstanceOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DbInstanceIdSets == nil {
		return nil, errors.New("invalid value for required argument 'DbInstanceIdSets'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DisisolateDbInstanceOperation
	err := ctx.RegisterResource("tencentcloud:Postgresql/disisolateDbInstanceOperation:DisisolateDbInstanceOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisisolateDbInstanceOperation gets an existing DisisolateDbInstanceOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisisolateDbInstanceOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DisisolateDbInstanceOperationState, opts ...pulumi.ResourceOption) (*DisisolateDbInstanceOperation, error) {
	var resource DisisolateDbInstanceOperation
	err := ctx.ReadResource("tencentcloud:Postgresql/disisolateDbInstanceOperation:DisisolateDbInstanceOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DisisolateDbInstanceOperation resources.
type disisolateDbInstanceOperationState struct {
	// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
	AutoVoucher *bool `pulumi:"autoVoucher"`
	// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
	DbInstanceIdSets []string `pulumi:"dbInstanceIdSets"`
	// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
	Period *int `pulumi:"period"`
	// Voucher ID list.
	VoucherIds []string `pulumi:"voucherIds"`
}

type DisisolateDbInstanceOperationState struct {
	// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
	AutoVoucher pulumi.BoolPtrInput
	// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
	DbInstanceIdSets pulumi.StringArrayInput
	// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
	Period pulumi.IntPtrInput
	// Voucher ID list.
	VoucherIds pulumi.StringArrayInput
}

func (DisisolateDbInstanceOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*disisolateDbInstanceOperationState)(nil)).Elem()
}

type disisolateDbInstanceOperationArgs struct {
	// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
	AutoVoucher *bool `pulumi:"autoVoucher"`
	// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
	DbInstanceIdSets []string `pulumi:"dbInstanceIdSets"`
	// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
	Period *int `pulumi:"period"`
	// Voucher ID list.
	VoucherIds []string `pulumi:"voucherIds"`
}

// The set of arguments for constructing a DisisolateDbInstanceOperation resource.
type DisisolateDbInstanceOperationArgs struct {
	// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
	AutoVoucher pulumi.BoolPtrInput
	// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
	DbInstanceIdSets pulumi.StringArrayInput
	// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
	Period pulumi.IntPtrInput
	// Voucher ID list.
	VoucherIds pulumi.StringArrayInput
}

func (DisisolateDbInstanceOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*disisolateDbInstanceOperationArgs)(nil)).Elem()
}

type DisisolateDbInstanceOperationInput interface {
	pulumi.Input

	ToDisisolateDbInstanceOperationOutput() DisisolateDbInstanceOperationOutput
	ToDisisolateDbInstanceOperationOutputWithContext(ctx context.Context) DisisolateDbInstanceOperationOutput
}

func (*DisisolateDbInstanceOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**DisisolateDbInstanceOperation)(nil)).Elem()
}

func (i *DisisolateDbInstanceOperation) ToDisisolateDbInstanceOperationOutput() DisisolateDbInstanceOperationOutput {
	return i.ToDisisolateDbInstanceOperationOutputWithContext(context.Background())
}

func (i *DisisolateDbInstanceOperation) ToDisisolateDbInstanceOperationOutputWithContext(ctx context.Context) DisisolateDbInstanceOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisisolateDbInstanceOperationOutput)
}

// DisisolateDbInstanceOperationArrayInput is an input type that accepts DisisolateDbInstanceOperationArray and DisisolateDbInstanceOperationArrayOutput values.
// You can construct a concrete instance of `DisisolateDbInstanceOperationArrayInput` via:
//
//	DisisolateDbInstanceOperationArray{ DisisolateDbInstanceOperationArgs{...} }
type DisisolateDbInstanceOperationArrayInput interface {
	pulumi.Input

	ToDisisolateDbInstanceOperationArrayOutput() DisisolateDbInstanceOperationArrayOutput
	ToDisisolateDbInstanceOperationArrayOutputWithContext(context.Context) DisisolateDbInstanceOperationArrayOutput
}

type DisisolateDbInstanceOperationArray []DisisolateDbInstanceOperationInput

func (DisisolateDbInstanceOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DisisolateDbInstanceOperation)(nil)).Elem()
}

func (i DisisolateDbInstanceOperationArray) ToDisisolateDbInstanceOperationArrayOutput() DisisolateDbInstanceOperationArrayOutput {
	return i.ToDisisolateDbInstanceOperationArrayOutputWithContext(context.Background())
}

func (i DisisolateDbInstanceOperationArray) ToDisisolateDbInstanceOperationArrayOutputWithContext(ctx context.Context) DisisolateDbInstanceOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisisolateDbInstanceOperationArrayOutput)
}

// DisisolateDbInstanceOperationMapInput is an input type that accepts DisisolateDbInstanceOperationMap and DisisolateDbInstanceOperationMapOutput values.
// You can construct a concrete instance of `DisisolateDbInstanceOperationMapInput` via:
//
//	DisisolateDbInstanceOperationMap{ "key": DisisolateDbInstanceOperationArgs{...} }
type DisisolateDbInstanceOperationMapInput interface {
	pulumi.Input

	ToDisisolateDbInstanceOperationMapOutput() DisisolateDbInstanceOperationMapOutput
	ToDisisolateDbInstanceOperationMapOutputWithContext(context.Context) DisisolateDbInstanceOperationMapOutput
}

type DisisolateDbInstanceOperationMap map[string]DisisolateDbInstanceOperationInput

func (DisisolateDbInstanceOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DisisolateDbInstanceOperation)(nil)).Elem()
}

func (i DisisolateDbInstanceOperationMap) ToDisisolateDbInstanceOperationMapOutput() DisisolateDbInstanceOperationMapOutput {
	return i.ToDisisolateDbInstanceOperationMapOutputWithContext(context.Background())
}

func (i DisisolateDbInstanceOperationMap) ToDisisolateDbInstanceOperationMapOutputWithContext(ctx context.Context) DisisolateDbInstanceOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisisolateDbInstanceOperationMapOutput)
}

type DisisolateDbInstanceOperationOutput struct{ *pulumi.OutputState }

func (DisisolateDbInstanceOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DisisolateDbInstanceOperation)(nil)).Elem()
}

func (o DisisolateDbInstanceOperationOutput) ToDisisolateDbInstanceOperationOutput() DisisolateDbInstanceOperationOutput {
	return o
}

func (o DisisolateDbInstanceOperationOutput) ToDisisolateDbInstanceOperationOutputWithContext(ctx context.Context) DisisolateDbInstanceOperationOutput {
	return o
}

// Whether to use vouchers. Valid values:true (yes), false (no). Default value:false.
func (o DisisolateDbInstanceOperationOutput) AutoVoucher() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DisisolateDbInstanceOperation) pulumi.BoolPtrOutput { return v.AutoVoucher }).(pulumi.BoolPtrOutput)
}

// List of resource IDs. Note that currently you cannot remove multiple instances from isolation at the same time. Only one instance ID can be passed in here.
func (o DisisolateDbInstanceOperationOutput) DbInstanceIdSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DisisolateDbInstanceOperation) pulumi.StringArrayOutput { return v.DbInstanceIdSets }).(pulumi.StringArrayOutput)
}

// The valid period (in months) of the monthly-subscribed instance when removing it from isolation.
func (o DisisolateDbInstanceOperationOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DisisolateDbInstanceOperation) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Voucher ID list.
func (o DisisolateDbInstanceOperationOutput) VoucherIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DisisolateDbInstanceOperation) pulumi.StringArrayOutput { return v.VoucherIds }).(pulumi.StringArrayOutput)
}

type DisisolateDbInstanceOperationArrayOutput struct{ *pulumi.OutputState }

func (DisisolateDbInstanceOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DisisolateDbInstanceOperation)(nil)).Elem()
}

func (o DisisolateDbInstanceOperationArrayOutput) ToDisisolateDbInstanceOperationArrayOutput() DisisolateDbInstanceOperationArrayOutput {
	return o
}

func (o DisisolateDbInstanceOperationArrayOutput) ToDisisolateDbInstanceOperationArrayOutputWithContext(ctx context.Context) DisisolateDbInstanceOperationArrayOutput {
	return o
}

func (o DisisolateDbInstanceOperationArrayOutput) Index(i pulumi.IntInput) DisisolateDbInstanceOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DisisolateDbInstanceOperation {
		return vs[0].([]*DisisolateDbInstanceOperation)[vs[1].(int)]
	}).(DisisolateDbInstanceOperationOutput)
}

type DisisolateDbInstanceOperationMapOutput struct{ *pulumi.OutputState }

func (DisisolateDbInstanceOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DisisolateDbInstanceOperation)(nil)).Elem()
}

func (o DisisolateDbInstanceOperationMapOutput) ToDisisolateDbInstanceOperationMapOutput() DisisolateDbInstanceOperationMapOutput {
	return o
}

func (o DisisolateDbInstanceOperationMapOutput) ToDisisolateDbInstanceOperationMapOutputWithContext(ctx context.Context) DisisolateDbInstanceOperationMapOutput {
	return o
}

func (o DisisolateDbInstanceOperationMapOutput) MapIndex(k pulumi.StringInput) DisisolateDbInstanceOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DisisolateDbInstanceOperation {
		return vs[0].(map[string]*DisisolateDbInstanceOperation)[vs[1].(string)]
	}).(DisisolateDbInstanceOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DisisolateDbInstanceOperationInput)(nil)).Elem(), &DisisolateDbInstanceOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*DisisolateDbInstanceOperationArrayInput)(nil)).Elem(), DisisolateDbInstanceOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DisisolateDbInstanceOperationMapInput)(nil)).Elem(), DisisolateDbInstanceOperationMap{})
	pulumi.RegisterOutputType(DisisolateDbInstanceOperationOutput{})
	pulumi.RegisterOutputType(DisisolateDbInstanceOperationArrayOutput{})
	pulumi.RegisterOutputType(DisisolateDbInstanceOperationMapOutput{})
}
