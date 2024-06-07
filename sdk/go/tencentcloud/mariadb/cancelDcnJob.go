// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a mariadb cancelDcnJob
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.NewCancelDcnJob(ctx, "cancelDcnJob", &Mariadb.CancelDcnJobArgs{
//				InstanceId: pulumi.String("tdsql-9vqvls95"),
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
type CancelDcnJob struct {
	pulumi.CustomResourceState

	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewCancelDcnJob registers a new resource with the given unique name, arguments, and options.
func NewCancelDcnJob(ctx *pulumi.Context,
	name string, args *CancelDcnJobArgs, opts ...pulumi.ResourceOption) (*CancelDcnJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CancelDcnJob
	err := ctx.RegisterResource("tencentcloud:Mariadb/cancelDcnJob:CancelDcnJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCancelDcnJob gets an existing CancelDcnJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCancelDcnJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CancelDcnJobState, opts ...pulumi.ResourceOption) (*CancelDcnJob, error) {
	var resource CancelDcnJob
	err := ctx.ReadResource("tencentcloud:Mariadb/cancelDcnJob:CancelDcnJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CancelDcnJob resources.
type cancelDcnJobState struct {
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
}

type CancelDcnJobState struct {
	// Instance ID.
	InstanceId pulumi.StringPtrInput
}

func (CancelDcnJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*cancelDcnJobState)(nil)).Elem()
}

type cancelDcnJobArgs struct {
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a CancelDcnJob resource.
type CancelDcnJobArgs struct {
	// Instance ID.
	InstanceId pulumi.StringInput
}

func (CancelDcnJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cancelDcnJobArgs)(nil)).Elem()
}

type CancelDcnJobInput interface {
	pulumi.Input

	ToCancelDcnJobOutput() CancelDcnJobOutput
	ToCancelDcnJobOutputWithContext(ctx context.Context) CancelDcnJobOutput
}

func (*CancelDcnJob) ElementType() reflect.Type {
	return reflect.TypeOf((**CancelDcnJob)(nil)).Elem()
}

func (i *CancelDcnJob) ToCancelDcnJobOutput() CancelDcnJobOutput {
	return i.ToCancelDcnJobOutputWithContext(context.Background())
}

func (i *CancelDcnJob) ToCancelDcnJobOutputWithContext(ctx context.Context) CancelDcnJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CancelDcnJobOutput)
}

// CancelDcnJobArrayInput is an input type that accepts CancelDcnJobArray and CancelDcnJobArrayOutput values.
// You can construct a concrete instance of `CancelDcnJobArrayInput` via:
//
//	CancelDcnJobArray{ CancelDcnJobArgs{...} }
type CancelDcnJobArrayInput interface {
	pulumi.Input

	ToCancelDcnJobArrayOutput() CancelDcnJobArrayOutput
	ToCancelDcnJobArrayOutputWithContext(context.Context) CancelDcnJobArrayOutput
}

type CancelDcnJobArray []CancelDcnJobInput

func (CancelDcnJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CancelDcnJob)(nil)).Elem()
}

func (i CancelDcnJobArray) ToCancelDcnJobArrayOutput() CancelDcnJobArrayOutput {
	return i.ToCancelDcnJobArrayOutputWithContext(context.Background())
}

func (i CancelDcnJobArray) ToCancelDcnJobArrayOutputWithContext(ctx context.Context) CancelDcnJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CancelDcnJobArrayOutput)
}

// CancelDcnJobMapInput is an input type that accepts CancelDcnJobMap and CancelDcnJobMapOutput values.
// You can construct a concrete instance of `CancelDcnJobMapInput` via:
//
//	CancelDcnJobMap{ "key": CancelDcnJobArgs{...} }
type CancelDcnJobMapInput interface {
	pulumi.Input

	ToCancelDcnJobMapOutput() CancelDcnJobMapOutput
	ToCancelDcnJobMapOutputWithContext(context.Context) CancelDcnJobMapOutput
}

type CancelDcnJobMap map[string]CancelDcnJobInput

func (CancelDcnJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CancelDcnJob)(nil)).Elem()
}

func (i CancelDcnJobMap) ToCancelDcnJobMapOutput() CancelDcnJobMapOutput {
	return i.ToCancelDcnJobMapOutputWithContext(context.Background())
}

func (i CancelDcnJobMap) ToCancelDcnJobMapOutputWithContext(ctx context.Context) CancelDcnJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CancelDcnJobMapOutput)
}

type CancelDcnJobOutput struct{ *pulumi.OutputState }

func (CancelDcnJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CancelDcnJob)(nil)).Elem()
}

func (o CancelDcnJobOutput) ToCancelDcnJobOutput() CancelDcnJobOutput {
	return o
}

func (o CancelDcnJobOutput) ToCancelDcnJobOutputWithContext(ctx context.Context) CancelDcnJobOutput {
	return o
}

// Instance ID.
func (o CancelDcnJobOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CancelDcnJob) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type CancelDcnJobArrayOutput struct{ *pulumi.OutputState }

func (CancelDcnJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CancelDcnJob)(nil)).Elem()
}

func (o CancelDcnJobArrayOutput) ToCancelDcnJobArrayOutput() CancelDcnJobArrayOutput {
	return o
}

func (o CancelDcnJobArrayOutput) ToCancelDcnJobArrayOutputWithContext(ctx context.Context) CancelDcnJobArrayOutput {
	return o
}

func (o CancelDcnJobArrayOutput) Index(i pulumi.IntInput) CancelDcnJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CancelDcnJob {
		return vs[0].([]*CancelDcnJob)[vs[1].(int)]
	}).(CancelDcnJobOutput)
}

type CancelDcnJobMapOutput struct{ *pulumi.OutputState }

func (CancelDcnJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CancelDcnJob)(nil)).Elem()
}

func (o CancelDcnJobMapOutput) ToCancelDcnJobMapOutput() CancelDcnJobMapOutput {
	return o
}

func (o CancelDcnJobMapOutput) ToCancelDcnJobMapOutputWithContext(ctx context.Context) CancelDcnJobMapOutput {
	return o
}

func (o CancelDcnJobMapOutput) MapIndex(k pulumi.StringInput) CancelDcnJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CancelDcnJob {
		return vs[0].(map[string]*CancelDcnJob)[vs[1].(string)]
	}).(CancelDcnJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CancelDcnJobInput)(nil)).Elem(), &CancelDcnJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*CancelDcnJobArrayInput)(nil)).Elem(), CancelDcnJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CancelDcnJobMapInput)(nil)).Elem(), CancelDcnJobMap{})
	pulumi.RegisterOutputType(CancelDcnJobOutput{})
	pulumi.RegisterOutputType(CancelDcnJobArrayOutput{})
	pulumi.RegisterOutputType(CancelDcnJobMapOutput{})
}
