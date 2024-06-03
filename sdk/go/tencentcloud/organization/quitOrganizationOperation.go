// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package organization

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a organization quitOrganizationOperation
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Organization"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Organization.NewQuitOrganizationOperation(ctx, "quitOrganizationOperation", &Organization.QuitOrganizationOperationArgs{
//				OrgId: pulumi.Int(45155),
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
// organization quit_organization_operation can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Organization/quitOrganizationOperation:QuitOrganizationOperation quit_organization_operation quit_organization_operation_id
// ```
type QuitOrganizationOperation struct {
	pulumi.CustomResourceState

	// Organization ID.
	OrgId pulumi.IntOutput `pulumi:"orgId"`
}

// NewQuitOrganizationOperation registers a new resource with the given unique name, arguments, and options.
func NewQuitOrganizationOperation(ctx *pulumi.Context,
	name string, args *QuitOrganizationOperationArgs, opts ...pulumi.ResourceOption) (*QuitOrganizationOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource QuitOrganizationOperation
	err := ctx.RegisterResource("tencentcloud:Organization/quitOrganizationOperation:QuitOrganizationOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuitOrganizationOperation gets an existing QuitOrganizationOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuitOrganizationOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuitOrganizationOperationState, opts ...pulumi.ResourceOption) (*QuitOrganizationOperation, error) {
	var resource QuitOrganizationOperation
	err := ctx.ReadResource("tencentcloud:Organization/quitOrganizationOperation:QuitOrganizationOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuitOrganizationOperation resources.
type quitOrganizationOperationState struct {
	// Organization ID.
	OrgId *int `pulumi:"orgId"`
}

type QuitOrganizationOperationState struct {
	// Organization ID.
	OrgId pulumi.IntPtrInput
}

func (QuitOrganizationOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*quitOrganizationOperationState)(nil)).Elem()
}

type quitOrganizationOperationArgs struct {
	// Organization ID.
	OrgId int `pulumi:"orgId"`
}

// The set of arguments for constructing a QuitOrganizationOperation resource.
type QuitOrganizationOperationArgs struct {
	// Organization ID.
	OrgId pulumi.IntInput
}

func (QuitOrganizationOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quitOrganizationOperationArgs)(nil)).Elem()
}

type QuitOrganizationOperationInput interface {
	pulumi.Input

	ToQuitOrganizationOperationOutput() QuitOrganizationOperationOutput
	ToQuitOrganizationOperationOutputWithContext(ctx context.Context) QuitOrganizationOperationOutput
}

func (*QuitOrganizationOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**QuitOrganizationOperation)(nil)).Elem()
}

func (i *QuitOrganizationOperation) ToQuitOrganizationOperationOutput() QuitOrganizationOperationOutput {
	return i.ToQuitOrganizationOperationOutputWithContext(context.Background())
}

func (i *QuitOrganizationOperation) ToQuitOrganizationOperationOutputWithContext(ctx context.Context) QuitOrganizationOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuitOrganizationOperationOutput)
}

// QuitOrganizationOperationArrayInput is an input type that accepts QuitOrganizationOperationArray and QuitOrganizationOperationArrayOutput values.
// You can construct a concrete instance of `QuitOrganizationOperationArrayInput` via:
//
//	QuitOrganizationOperationArray{ QuitOrganizationOperationArgs{...} }
type QuitOrganizationOperationArrayInput interface {
	pulumi.Input

	ToQuitOrganizationOperationArrayOutput() QuitOrganizationOperationArrayOutput
	ToQuitOrganizationOperationArrayOutputWithContext(context.Context) QuitOrganizationOperationArrayOutput
}

type QuitOrganizationOperationArray []QuitOrganizationOperationInput

func (QuitOrganizationOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuitOrganizationOperation)(nil)).Elem()
}

func (i QuitOrganizationOperationArray) ToQuitOrganizationOperationArrayOutput() QuitOrganizationOperationArrayOutput {
	return i.ToQuitOrganizationOperationArrayOutputWithContext(context.Background())
}

func (i QuitOrganizationOperationArray) ToQuitOrganizationOperationArrayOutputWithContext(ctx context.Context) QuitOrganizationOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuitOrganizationOperationArrayOutput)
}

// QuitOrganizationOperationMapInput is an input type that accepts QuitOrganizationOperationMap and QuitOrganizationOperationMapOutput values.
// You can construct a concrete instance of `QuitOrganizationOperationMapInput` via:
//
//	QuitOrganizationOperationMap{ "key": QuitOrganizationOperationArgs{...} }
type QuitOrganizationOperationMapInput interface {
	pulumi.Input

	ToQuitOrganizationOperationMapOutput() QuitOrganizationOperationMapOutput
	ToQuitOrganizationOperationMapOutputWithContext(context.Context) QuitOrganizationOperationMapOutput
}

type QuitOrganizationOperationMap map[string]QuitOrganizationOperationInput

func (QuitOrganizationOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuitOrganizationOperation)(nil)).Elem()
}

func (i QuitOrganizationOperationMap) ToQuitOrganizationOperationMapOutput() QuitOrganizationOperationMapOutput {
	return i.ToQuitOrganizationOperationMapOutputWithContext(context.Background())
}

func (i QuitOrganizationOperationMap) ToQuitOrganizationOperationMapOutputWithContext(ctx context.Context) QuitOrganizationOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuitOrganizationOperationMapOutput)
}

type QuitOrganizationOperationOutput struct{ *pulumi.OutputState }

func (QuitOrganizationOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuitOrganizationOperation)(nil)).Elem()
}

func (o QuitOrganizationOperationOutput) ToQuitOrganizationOperationOutput() QuitOrganizationOperationOutput {
	return o
}

func (o QuitOrganizationOperationOutput) ToQuitOrganizationOperationOutputWithContext(ctx context.Context) QuitOrganizationOperationOutput {
	return o
}

// Organization ID.
func (o QuitOrganizationOperationOutput) OrgId() pulumi.IntOutput {
	return o.ApplyT(func(v *QuitOrganizationOperation) pulumi.IntOutput { return v.OrgId }).(pulumi.IntOutput)
}

type QuitOrganizationOperationArrayOutput struct{ *pulumi.OutputState }

func (QuitOrganizationOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*QuitOrganizationOperation)(nil)).Elem()
}

func (o QuitOrganizationOperationArrayOutput) ToQuitOrganizationOperationArrayOutput() QuitOrganizationOperationArrayOutput {
	return o
}

func (o QuitOrganizationOperationArrayOutput) ToQuitOrganizationOperationArrayOutputWithContext(ctx context.Context) QuitOrganizationOperationArrayOutput {
	return o
}

func (o QuitOrganizationOperationArrayOutput) Index(i pulumi.IntInput) QuitOrganizationOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *QuitOrganizationOperation {
		return vs[0].([]*QuitOrganizationOperation)[vs[1].(int)]
	}).(QuitOrganizationOperationOutput)
}

type QuitOrganizationOperationMapOutput struct{ *pulumi.OutputState }

func (QuitOrganizationOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*QuitOrganizationOperation)(nil)).Elem()
}

func (o QuitOrganizationOperationMapOutput) ToQuitOrganizationOperationMapOutput() QuitOrganizationOperationMapOutput {
	return o
}

func (o QuitOrganizationOperationMapOutput) ToQuitOrganizationOperationMapOutputWithContext(ctx context.Context) QuitOrganizationOperationMapOutput {
	return o
}

func (o QuitOrganizationOperationMapOutput) MapIndex(k pulumi.StringInput) QuitOrganizationOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *QuitOrganizationOperation {
		return vs[0].(map[string]*QuitOrganizationOperation)[vs[1].(string)]
	}).(QuitOrganizationOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QuitOrganizationOperationInput)(nil)).Elem(), &QuitOrganizationOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuitOrganizationOperationArrayInput)(nil)).Elem(), QuitOrganizationOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*QuitOrganizationOperationMapInput)(nil)).Elem(), QuitOrganizationOperationMap{})
	pulumi.RegisterOutputType(QuitOrganizationOperationOutput{})
	pulumi.RegisterOutputType(QuitOrganizationOperationArrayOutput{})
	pulumi.RegisterOutputType(QuitOrganizationOperationMapOutput{})
}
