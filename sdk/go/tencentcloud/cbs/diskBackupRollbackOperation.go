// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cbs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to rollback cbs disk backup.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cbs"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cbs.NewDiskBackupRollbackOperation(ctx, "operation", &Cbs.DiskBackupRollbackOperationArgs{
// 			DiskBackupId: pulumi.String("dbp-xxx"),
// 			DiskId:       pulumi.String("disk-xxx"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DiskBackupRollbackOperation struct {
	pulumi.CustomResourceState

	// Cloud disk backup point ID.
	DiskBackupId pulumi.StringOutput `pulumi:"diskBackupId"`
	// Cloud disk backup point original cloud disk ID.
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// Whether the rollback is completed. `true` meaing rollback completed, `false` meaning still rollbacking.
	IsRollbackCompleted pulumi.BoolOutput `pulumi:"isRollbackCompleted"`
}

// NewDiskBackupRollbackOperation registers a new resource with the given unique name, arguments, and options.
func NewDiskBackupRollbackOperation(ctx *pulumi.Context,
	name string, args *DiskBackupRollbackOperationArgs, opts ...pulumi.ResourceOption) (*DiskBackupRollbackOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskBackupId == nil {
		return nil, errors.New("invalid value for required argument 'DiskBackupId'")
	}
	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DiskBackupRollbackOperation
	err := ctx.RegisterResource("tencentcloud:Cbs/diskBackupRollbackOperation:DiskBackupRollbackOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskBackupRollbackOperation gets an existing DiskBackupRollbackOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskBackupRollbackOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskBackupRollbackOperationState, opts ...pulumi.ResourceOption) (*DiskBackupRollbackOperation, error) {
	var resource DiskBackupRollbackOperation
	err := ctx.ReadResource("tencentcloud:Cbs/diskBackupRollbackOperation:DiskBackupRollbackOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskBackupRollbackOperation resources.
type diskBackupRollbackOperationState struct {
	// Cloud disk backup point ID.
	DiskBackupId *string `pulumi:"diskBackupId"`
	// Cloud disk backup point original cloud disk ID.
	DiskId *string `pulumi:"diskId"`
	// Whether the rollback is completed. `true` meaing rollback completed, `false` meaning still rollbacking.
	IsRollbackCompleted *bool `pulumi:"isRollbackCompleted"`
}

type DiskBackupRollbackOperationState struct {
	// Cloud disk backup point ID.
	DiskBackupId pulumi.StringPtrInput
	// Cloud disk backup point original cloud disk ID.
	DiskId pulumi.StringPtrInput
	// Whether the rollback is completed. `true` meaing rollback completed, `false` meaning still rollbacking.
	IsRollbackCompleted pulumi.BoolPtrInput
}

func (DiskBackupRollbackOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskBackupRollbackOperationState)(nil)).Elem()
}

type diskBackupRollbackOperationArgs struct {
	// Cloud disk backup point ID.
	DiskBackupId string `pulumi:"diskBackupId"`
	// Cloud disk backup point original cloud disk ID.
	DiskId string `pulumi:"diskId"`
}

// The set of arguments for constructing a DiskBackupRollbackOperation resource.
type DiskBackupRollbackOperationArgs struct {
	// Cloud disk backup point ID.
	DiskBackupId pulumi.StringInput
	// Cloud disk backup point original cloud disk ID.
	DiskId pulumi.StringInput
}

func (DiskBackupRollbackOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskBackupRollbackOperationArgs)(nil)).Elem()
}

type DiskBackupRollbackOperationInput interface {
	pulumi.Input

	ToDiskBackupRollbackOperationOutput() DiskBackupRollbackOperationOutput
	ToDiskBackupRollbackOperationOutputWithContext(ctx context.Context) DiskBackupRollbackOperationOutput
}

func (*DiskBackupRollbackOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskBackupRollbackOperation)(nil)).Elem()
}

func (i *DiskBackupRollbackOperation) ToDiskBackupRollbackOperationOutput() DiskBackupRollbackOperationOutput {
	return i.ToDiskBackupRollbackOperationOutputWithContext(context.Background())
}

func (i *DiskBackupRollbackOperation) ToDiskBackupRollbackOperationOutputWithContext(ctx context.Context) DiskBackupRollbackOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskBackupRollbackOperationOutput)
}

// DiskBackupRollbackOperationArrayInput is an input type that accepts DiskBackupRollbackOperationArray and DiskBackupRollbackOperationArrayOutput values.
// You can construct a concrete instance of `DiskBackupRollbackOperationArrayInput` via:
//
//          DiskBackupRollbackOperationArray{ DiskBackupRollbackOperationArgs{...} }
type DiskBackupRollbackOperationArrayInput interface {
	pulumi.Input

	ToDiskBackupRollbackOperationArrayOutput() DiskBackupRollbackOperationArrayOutput
	ToDiskBackupRollbackOperationArrayOutputWithContext(context.Context) DiskBackupRollbackOperationArrayOutput
}

type DiskBackupRollbackOperationArray []DiskBackupRollbackOperationInput

func (DiskBackupRollbackOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskBackupRollbackOperation)(nil)).Elem()
}

func (i DiskBackupRollbackOperationArray) ToDiskBackupRollbackOperationArrayOutput() DiskBackupRollbackOperationArrayOutput {
	return i.ToDiskBackupRollbackOperationArrayOutputWithContext(context.Background())
}

func (i DiskBackupRollbackOperationArray) ToDiskBackupRollbackOperationArrayOutputWithContext(ctx context.Context) DiskBackupRollbackOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskBackupRollbackOperationArrayOutput)
}

// DiskBackupRollbackOperationMapInput is an input type that accepts DiskBackupRollbackOperationMap and DiskBackupRollbackOperationMapOutput values.
// You can construct a concrete instance of `DiskBackupRollbackOperationMapInput` via:
//
//          DiskBackupRollbackOperationMap{ "key": DiskBackupRollbackOperationArgs{...} }
type DiskBackupRollbackOperationMapInput interface {
	pulumi.Input

	ToDiskBackupRollbackOperationMapOutput() DiskBackupRollbackOperationMapOutput
	ToDiskBackupRollbackOperationMapOutputWithContext(context.Context) DiskBackupRollbackOperationMapOutput
}

type DiskBackupRollbackOperationMap map[string]DiskBackupRollbackOperationInput

func (DiskBackupRollbackOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskBackupRollbackOperation)(nil)).Elem()
}

func (i DiskBackupRollbackOperationMap) ToDiskBackupRollbackOperationMapOutput() DiskBackupRollbackOperationMapOutput {
	return i.ToDiskBackupRollbackOperationMapOutputWithContext(context.Background())
}

func (i DiskBackupRollbackOperationMap) ToDiskBackupRollbackOperationMapOutputWithContext(ctx context.Context) DiskBackupRollbackOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskBackupRollbackOperationMapOutput)
}

type DiskBackupRollbackOperationOutput struct{ *pulumi.OutputState }

func (DiskBackupRollbackOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskBackupRollbackOperation)(nil)).Elem()
}

func (o DiskBackupRollbackOperationOutput) ToDiskBackupRollbackOperationOutput() DiskBackupRollbackOperationOutput {
	return o
}

func (o DiskBackupRollbackOperationOutput) ToDiskBackupRollbackOperationOutputWithContext(ctx context.Context) DiskBackupRollbackOperationOutput {
	return o
}

// Cloud disk backup point ID.
func (o DiskBackupRollbackOperationOutput) DiskBackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskBackupRollbackOperation) pulumi.StringOutput { return v.DiskBackupId }).(pulumi.StringOutput)
}

// Cloud disk backup point original cloud disk ID.
func (o DiskBackupRollbackOperationOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *DiskBackupRollbackOperation) pulumi.StringOutput { return v.DiskId }).(pulumi.StringOutput)
}

// Whether the rollback is completed. `true` meaing rollback completed, `false` meaning still rollbacking.
func (o DiskBackupRollbackOperationOutput) IsRollbackCompleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *DiskBackupRollbackOperation) pulumi.BoolOutput { return v.IsRollbackCompleted }).(pulumi.BoolOutput)
}

type DiskBackupRollbackOperationArrayOutput struct{ *pulumi.OutputState }

func (DiskBackupRollbackOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DiskBackupRollbackOperation)(nil)).Elem()
}

func (o DiskBackupRollbackOperationArrayOutput) ToDiskBackupRollbackOperationArrayOutput() DiskBackupRollbackOperationArrayOutput {
	return o
}

func (o DiskBackupRollbackOperationArrayOutput) ToDiskBackupRollbackOperationArrayOutputWithContext(ctx context.Context) DiskBackupRollbackOperationArrayOutput {
	return o
}

func (o DiskBackupRollbackOperationArrayOutput) Index(i pulumi.IntInput) DiskBackupRollbackOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DiskBackupRollbackOperation {
		return vs[0].([]*DiskBackupRollbackOperation)[vs[1].(int)]
	}).(DiskBackupRollbackOperationOutput)
}

type DiskBackupRollbackOperationMapOutput struct{ *pulumi.OutputState }

func (DiskBackupRollbackOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DiskBackupRollbackOperation)(nil)).Elem()
}

func (o DiskBackupRollbackOperationMapOutput) ToDiskBackupRollbackOperationMapOutput() DiskBackupRollbackOperationMapOutput {
	return o
}

func (o DiskBackupRollbackOperationMapOutput) ToDiskBackupRollbackOperationMapOutputWithContext(ctx context.Context) DiskBackupRollbackOperationMapOutput {
	return o
}

func (o DiskBackupRollbackOperationMapOutput) MapIndex(k pulumi.StringInput) DiskBackupRollbackOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DiskBackupRollbackOperation {
		return vs[0].(map[string]*DiskBackupRollbackOperation)[vs[1].(string)]
	}).(DiskBackupRollbackOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskBackupRollbackOperationInput)(nil)).Elem(), &DiskBackupRollbackOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskBackupRollbackOperationArrayInput)(nil)).Elem(), DiskBackupRollbackOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskBackupRollbackOperationMapInput)(nil)).Elem(), DiskBackupRollbackOperationMap{})
	pulumi.RegisterOutputType(DiskBackupRollbackOperationOutput{})
	pulumi.RegisterOutputType(DiskBackupRollbackOperationArrayOutput{})
	pulumi.RegisterOutputType(DiskBackupRollbackOperationMapOutput{})
}
