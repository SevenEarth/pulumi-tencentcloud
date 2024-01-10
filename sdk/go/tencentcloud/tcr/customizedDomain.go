// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tcr customized domain
//
// ## Example Usage
// ### Create a tcr customized domain
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := Tcr.NewInstance(ctx, "exampleInstance", &Tcr.InstanceArgs{
//				InstanceType: pulumi.String("premium"),
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Tcr.NewCustomizedDomain(ctx, "exampleCustomizedDomain", &Tcr.CustomizedDomainArgs{
//				RegistryId:    exampleInstance.ID(),
//				DomainName:    pulumi.String("www.test.com"),
//				CertificateId: pulumi.String("your_cert_id"),
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("terraform"),
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
//
// ## Import
//
// tcr customized_domain can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Tcr/customizedDomain:CustomizedDomain customized_domain customized_domain_id
//
// ```
type CustomizedDomain struct {
	pulumi.CustomResourceState

	// certificate id.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// custom domain name.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// instance id.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewCustomizedDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomizedDomain(ctx *pulumi.Context,
	name string, args *CustomizedDomainArgs, opts ...pulumi.ResourceOption) (*CustomizedDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateId'")
	}
	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CustomizedDomain
	err := ctx.RegisterResource("tencentcloud:Tcr/customizedDomain:CustomizedDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomizedDomain gets an existing CustomizedDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomizedDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomizedDomainState, opts ...pulumi.ResourceOption) (*CustomizedDomain, error) {
	var resource CustomizedDomain
	err := ctx.ReadResource("tencentcloud:Tcr/customizedDomain:CustomizedDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomizedDomain resources.
type customizedDomainState struct {
	// certificate id.
	CertificateId *string `pulumi:"certificateId"`
	// custom domain name.
	DomainName *string `pulumi:"domainName"`
	// instance id.
	RegistryId *string `pulumi:"registryId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

type CustomizedDomainState struct {
	// certificate id.
	CertificateId pulumi.StringPtrInput
	// custom domain name.
	DomainName pulumi.StringPtrInput
	// instance id.
	RegistryId pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (CustomizedDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedDomainState)(nil)).Elem()
}

type customizedDomainArgs struct {
	// certificate id.
	CertificateId string `pulumi:"certificateId"`
	// custom domain name.
	DomainName string `pulumi:"domainName"`
	// instance id.
	RegistryId string `pulumi:"registryId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a CustomizedDomain resource.
type CustomizedDomainArgs struct {
	// certificate id.
	CertificateId pulumi.StringInput
	// custom domain name.
	DomainName pulumi.StringInput
	// instance id.
	RegistryId pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
}

func (CustomizedDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedDomainArgs)(nil)).Elem()
}

type CustomizedDomainInput interface {
	pulumi.Input

	ToCustomizedDomainOutput() CustomizedDomainOutput
	ToCustomizedDomainOutputWithContext(ctx context.Context) CustomizedDomainOutput
}

func (*CustomizedDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedDomain)(nil)).Elem()
}

func (i *CustomizedDomain) ToCustomizedDomainOutput() CustomizedDomainOutput {
	return i.ToCustomizedDomainOutputWithContext(context.Background())
}

func (i *CustomizedDomain) ToCustomizedDomainOutputWithContext(ctx context.Context) CustomizedDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedDomainOutput)
}

// CustomizedDomainArrayInput is an input type that accepts CustomizedDomainArray and CustomizedDomainArrayOutput values.
// You can construct a concrete instance of `CustomizedDomainArrayInput` via:
//
//	CustomizedDomainArray{ CustomizedDomainArgs{...} }
type CustomizedDomainArrayInput interface {
	pulumi.Input

	ToCustomizedDomainArrayOutput() CustomizedDomainArrayOutput
	ToCustomizedDomainArrayOutputWithContext(context.Context) CustomizedDomainArrayOutput
}

type CustomizedDomainArray []CustomizedDomainInput

func (CustomizedDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomizedDomain)(nil)).Elem()
}

func (i CustomizedDomainArray) ToCustomizedDomainArrayOutput() CustomizedDomainArrayOutput {
	return i.ToCustomizedDomainArrayOutputWithContext(context.Background())
}

func (i CustomizedDomainArray) ToCustomizedDomainArrayOutputWithContext(ctx context.Context) CustomizedDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedDomainArrayOutput)
}

// CustomizedDomainMapInput is an input type that accepts CustomizedDomainMap and CustomizedDomainMapOutput values.
// You can construct a concrete instance of `CustomizedDomainMapInput` via:
//
//	CustomizedDomainMap{ "key": CustomizedDomainArgs{...} }
type CustomizedDomainMapInput interface {
	pulumi.Input

	ToCustomizedDomainMapOutput() CustomizedDomainMapOutput
	ToCustomizedDomainMapOutputWithContext(context.Context) CustomizedDomainMapOutput
}

type CustomizedDomainMap map[string]CustomizedDomainInput

func (CustomizedDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomizedDomain)(nil)).Elem()
}

func (i CustomizedDomainMap) ToCustomizedDomainMapOutput() CustomizedDomainMapOutput {
	return i.ToCustomizedDomainMapOutputWithContext(context.Background())
}

func (i CustomizedDomainMap) ToCustomizedDomainMapOutputWithContext(ctx context.Context) CustomizedDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedDomainMapOutput)
}

type CustomizedDomainOutput struct{ *pulumi.OutputState }

func (CustomizedDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedDomain)(nil)).Elem()
}

func (o CustomizedDomainOutput) ToCustomizedDomainOutput() CustomizedDomainOutput {
	return o
}

func (o CustomizedDomainOutput) ToCustomizedDomainOutputWithContext(ctx context.Context) CustomizedDomainOutput {
	return o
}

// certificate id.
func (o CustomizedDomainOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedDomain) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// custom domain name.
func (o CustomizedDomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedDomain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// instance id.
func (o CustomizedDomainOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedDomain) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// Tag description list.
func (o CustomizedDomainOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *CustomizedDomain) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type CustomizedDomainArrayOutput struct{ *pulumi.OutputState }

func (CustomizedDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomizedDomain)(nil)).Elem()
}

func (o CustomizedDomainArrayOutput) ToCustomizedDomainArrayOutput() CustomizedDomainArrayOutput {
	return o
}

func (o CustomizedDomainArrayOutput) ToCustomizedDomainArrayOutputWithContext(ctx context.Context) CustomizedDomainArrayOutput {
	return o
}

func (o CustomizedDomainArrayOutput) Index(i pulumi.IntInput) CustomizedDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomizedDomain {
		return vs[0].([]*CustomizedDomain)[vs[1].(int)]
	}).(CustomizedDomainOutput)
}

type CustomizedDomainMapOutput struct{ *pulumi.OutputState }

func (CustomizedDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomizedDomain)(nil)).Elem()
}

func (o CustomizedDomainMapOutput) ToCustomizedDomainMapOutput() CustomizedDomainMapOutput {
	return o
}

func (o CustomizedDomainMapOutput) ToCustomizedDomainMapOutputWithContext(ctx context.Context) CustomizedDomainMapOutput {
	return o
}

func (o CustomizedDomainMapOutput) MapIndex(k pulumi.StringInput) CustomizedDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomizedDomain {
		return vs[0].(map[string]*CustomizedDomain)[vs[1].(string)]
	}).(CustomizedDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomizedDomainInput)(nil)).Elem(), &CustomizedDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomizedDomainArrayInput)(nil)).Elem(), CustomizedDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomizedDomainMapInput)(nil)).Elem(), CustomizedDomainMap{})
	pulumi.RegisterOutputType(CustomizedDomainOutput{})
	pulumi.RegisterOutputType(CustomizedDomainArrayOutput{})
	pulumi.RegisterOutputType(CustomizedDomainMapOutput{})
}
