// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ssl completeCertificate
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ssl.NewCompleteCertificateOperation(ctx, "completeCertificate", &Ssl.CompleteCertificateOperationArgs{
//				CertificateId: pulumi.String("9Bfe1IBR"),
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
// ssl complete_certificate can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Ssl/completeCertificateOperation:CompleteCertificateOperation complete_certificate complete_certificate_id
//
// ```
type CompleteCertificateOperation struct {
	pulumi.CustomResourceState

	// Certificate ID.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
}

// NewCompleteCertificateOperation registers a new resource with the given unique name, arguments, and options.
func NewCompleteCertificateOperation(ctx *pulumi.Context,
	name string, args *CompleteCertificateOperationArgs, opts ...pulumi.ResourceOption) (*CompleteCertificateOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CompleteCertificateOperation
	err := ctx.RegisterResource("tencentcloud:Ssl/completeCertificateOperation:CompleteCertificateOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCompleteCertificateOperation gets an existing CompleteCertificateOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCompleteCertificateOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CompleteCertificateOperationState, opts ...pulumi.ResourceOption) (*CompleteCertificateOperation, error) {
	var resource CompleteCertificateOperation
	err := ctx.ReadResource("tencentcloud:Ssl/completeCertificateOperation:CompleteCertificateOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CompleteCertificateOperation resources.
type completeCertificateOperationState struct {
	// Certificate ID.
	CertificateId *string `pulumi:"certificateId"`
}

type CompleteCertificateOperationState struct {
	// Certificate ID.
	CertificateId pulumi.StringPtrInput
}

func (CompleteCertificateOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*completeCertificateOperationState)(nil)).Elem()
}

type completeCertificateOperationArgs struct {
	// Certificate ID.
	CertificateId string `pulumi:"certificateId"`
}

// The set of arguments for constructing a CompleteCertificateOperation resource.
type CompleteCertificateOperationArgs struct {
	// Certificate ID.
	CertificateId pulumi.StringInput
}

func (CompleteCertificateOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*completeCertificateOperationArgs)(nil)).Elem()
}

type CompleteCertificateOperationInput interface {
	pulumi.Input

	ToCompleteCertificateOperationOutput() CompleteCertificateOperationOutput
	ToCompleteCertificateOperationOutputWithContext(ctx context.Context) CompleteCertificateOperationOutput
}

func (*CompleteCertificateOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**CompleteCertificateOperation)(nil)).Elem()
}

func (i *CompleteCertificateOperation) ToCompleteCertificateOperationOutput() CompleteCertificateOperationOutput {
	return i.ToCompleteCertificateOperationOutputWithContext(context.Background())
}

func (i *CompleteCertificateOperation) ToCompleteCertificateOperationOutputWithContext(ctx context.Context) CompleteCertificateOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompleteCertificateOperationOutput)
}

// CompleteCertificateOperationArrayInput is an input type that accepts CompleteCertificateOperationArray and CompleteCertificateOperationArrayOutput values.
// You can construct a concrete instance of `CompleteCertificateOperationArrayInput` via:
//
//	CompleteCertificateOperationArray{ CompleteCertificateOperationArgs{...} }
type CompleteCertificateOperationArrayInput interface {
	pulumi.Input

	ToCompleteCertificateOperationArrayOutput() CompleteCertificateOperationArrayOutput
	ToCompleteCertificateOperationArrayOutputWithContext(context.Context) CompleteCertificateOperationArrayOutput
}

type CompleteCertificateOperationArray []CompleteCertificateOperationInput

func (CompleteCertificateOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompleteCertificateOperation)(nil)).Elem()
}

func (i CompleteCertificateOperationArray) ToCompleteCertificateOperationArrayOutput() CompleteCertificateOperationArrayOutput {
	return i.ToCompleteCertificateOperationArrayOutputWithContext(context.Background())
}

func (i CompleteCertificateOperationArray) ToCompleteCertificateOperationArrayOutputWithContext(ctx context.Context) CompleteCertificateOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompleteCertificateOperationArrayOutput)
}

// CompleteCertificateOperationMapInput is an input type that accepts CompleteCertificateOperationMap and CompleteCertificateOperationMapOutput values.
// You can construct a concrete instance of `CompleteCertificateOperationMapInput` via:
//
//	CompleteCertificateOperationMap{ "key": CompleteCertificateOperationArgs{...} }
type CompleteCertificateOperationMapInput interface {
	pulumi.Input

	ToCompleteCertificateOperationMapOutput() CompleteCertificateOperationMapOutput
	ToCompleteCertificateOperationMapOutputWithContext(context.Context) CompleteCertificateOperationMapOutput
}

type CompleteCertificateOperationMap map[string]CompleteCertificateOperationInput

func (CompleteCertificateOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompleteCertificateOperation)(nil)).Elem()
}

func (i CompleteCertificateOperationMap) ToCompleteCertificateOperationMapOutput() CompleteCertificateOperationMapOutput {
	return i.ToCompleteCertificateOperationMapOutputWithContext(context.Background())
}

func (i CompleteCertificateOperationMap) ToCompleteCertificateOperationMapOutputWithContext(ctx context.Context) CompleteCertificateOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompleteCertificateOperationMapOutput)
}

type CompleteCertificateOperationOutput struct{ *pulumi.OutputState }

func (CompleteCertificateOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompleteCertificateOperation)(nil)).Elem()
}

func (o CompleteCertificateOperationOutput) ToCompleteCertificateOperationOutput() CompleteCertificateOperationOutput {
	return o
}

func (o CompleteCertificateOperationOutput) ToCompleteCertificateOperationOutputWithContext(ctx context.Context) CompleteCertificateOperationOutput {
	return o
}

// Certificate ID.
func (o CompleteCertificateOperationOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *CompleteCertificateOperation) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

type CompleteCertificateOperationArrayOutput struct{ *pulumi.OutputState }

func (CompleteCertificateOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CompleteCertificateOperation)(nil)).Elem()
}

func (o CompleteCertificateOperationArrayOutput) ToCompleteCertificateOperationArrayOutput() CompleteCertificateOperationArrayOutput {
	return o
}

func (o CompleteCertificateOperationArrayOutput) ToCompleteCertificateOperationArrayOutputWithContext(ctx context.Context) CompleteCertificateOperationArrayOutput {
	return o
}

func (o CompleteCertificateOperationArrayOutput) Index(i pulumi.IntInput) CompleteCertificateOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CompleteCertificateOperation {
		return vs[0].([]*CompleteCertificateOperation)[vs[1].(int)]
	}).(CompleteCertificateOperationOutput)
}

type CompleteCertificateOperationMapOutput struct{ *pulumi.OutputState }

func (CompleteCertificateOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CompleteCertificateOperation)(nil)).Elem()
}

func (o CompleteCertificateOperationMapOutput) ToCompleteCertificateOperationMapOutput() CompleteCertificateOperationMapOutput {
	return o
}

func (o CompleteCertificateOperationMapOutput) ToCompleteCertificateOperationMapOutputWithContext(ctx context.Context) CompleteCertificateOperationMapOutput {
	return o
}

func (o CompleteCertificateOperationMapOutput) MapIndex(k pulumi.StringInput) CompleteCertificateOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CompleteCertificateOperation {
		return vs[0].(map[string]*CompleteCertificateOperation)[vs[1].(string)]
	}).(CompleteCertificateOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CompleteCertificateOperationInput)(nil)).Elem(), &CompleteCertificateOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompleteCertificateOperationArrayInput)(nil)).Elem(), CompleteCertificateOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CompleteCertificateOperationMapInput)(nil)).Elem(), CompleteCertificateOperationMap{})
	pulumi.RegisterOutputType(CompleteCertificateOperationOutput{})
	pulumi.RegisterOutputType(CompleteCertificateOperationArrayOutput{})
	pulumi.RegisterOutputType(CompleteCertificateOperationMapOutput{})
}
