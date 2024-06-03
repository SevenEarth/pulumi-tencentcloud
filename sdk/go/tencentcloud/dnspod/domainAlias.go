// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnspod

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a dnspod domainAlias
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dnspod"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dnspod.NewDomainAlias(ctx, "domainAlias", &Dnspod.DomainAliasArgs{
//				Domain:      pulumi.String("dnspod.cn"),
//				DomainAlias: pulumi.String("dnspod.com"),
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
// dnspod domain_alias can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Dnspod/domainAlias:DomainAlias domain_alias domain#domain_alias_id
// ```
type DomainAlias struct {
	pulumi.CustomResourceState

	// Domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Domain alias.
	DomainAlias pulumi.StringOutput `pulumi:"domainAlias"`
	// Domain alias ID.
	DomainAliasId pulumi.IntOutput `pulumi:"domainAliasId"`
}

// NewDomainAlias registers a new resource with the given unique name, arguments, and options.
func NewDomainAlias(ctx *pulumi.Context,
	name string, args *DomainAliasArgs, opts ...pulumi.ResourceOption) (*DomainAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.DomainAlias == nil {
		return nil, errors.New("invalid value for required argument 'DomainAlias'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainAlias
	err := ctx.RegisterResource("tencentcloud:Dnspod/domainAlias:DomainAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainAlias gets an existing DomainAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainAliasState, opts ...pulumi.ResourceOption) (*DomainAlias, error) {
	var resource DomainAlias
	err := ctx.ReadResource("tencentcloud:Dnspod/domainAlias:DomainAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainAlias resources.
type domainAliasState struct {
	// Domain.
	Domain *string `pulumi:"domain"`
	// Domain alias.
	DomainAlias *string `pulumi:"domainAlias"`
	// Domain alias ID.
	DomainAliasId *int `pulumi:"domainAliasId"`
}

type DomainAliasState struct {
	// Domain.
	Domain pulumi.StringPtrInput
	// Domain alias.
	DomainAlias pulumi.StringPtrInput
	// Domain alias ID.
	DomainAliasId pulumi.IntPtrInput
}

func (DomainAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAliasState)(nil)).Elem()
}

type domainAliasArgs struct {
	// Domain.
	Domain string `pulumi:"domain"`
	// Domain alias.
	DomainAlias string `pulumi:"domainAlias"`
}

// The set of arguments for constructing a DomainAlias resource.
type DomainAliasArgs struct {
	// Domain.
	Domain pulumi.StringInput
	// Domain alias.
	DomainAlias pulumi.StringInput
}

func (DomainAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainAliasArgs)(nil)).Elem()
}

type DomainAliasInput interface {
	pulumi.Input

	ToDomainAliasOutput() DomainAliasOutput
	ToDomainAliasOutputWithContext(ctx context.Context) DomainAliasOutput
}

func (*DomainAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAlias)(nil)).Elem()
}

func (i *DomainAlias) ToDomainAliasOutput() DomainAliasOutput {
	return i.ToDomainAliasOutputWithContext(context.Background())
}

func (i *DomainAlias) ToDomainAliasOutputWithContext(ctx context.Context) DomainAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAliasOutput)
}

// DomainAliasArrayInput is an input type that accepts DomainAliasArray and DomainAliasArrayOutput values.
// You can construct a concrete instance of `DomainAliasArrayInput` via:
//
//	DomainAliasArray{ DomainAliasArgs{...} }
type DomainAliasArrayInput interface {
	pulumi.Input

	ToDomainAliasArrayOutput() DomainAliasArrayOutput
	ToDomainAliasArrayOutputWithContext(context.Context) DomainAliasArrayOutput
}

type DomainAliasArray []DomainAliasInput

func (DomainAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainAlias)(nil)).Elem()
}

func (i DomainAliasArray) ToDomainAliasArrayOutput() DomainAliasArrayOutput {
	return i.ToDomainAliasArrayOutputWithContext(context.Background())
}

func (i DomainAliasArray) ToDomainAliasArrayOutputWithContext(ctx context.Context) DomainAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAliasArrayOutput)
}

// DomainAliasMapInput is an input type that accepts DomainAliasMap and DomainAliasMapOutput values.
// You can construct a concrete instance of `DomainAliasMapInput` via:
//
//	DomainAliasMap{ "key": DomainAliasArgs{...} }
type DomainAliasMapInput interface {
	pulumi.Input

	ToDomainAliasMapOutput() DomainAliasMapOutput
	ToDomainAliasMapOutputWithContext(context.Context) DomainAliasMapOutput
}

type DomainAliasMap map[string]DomainAliasInput

func (DomainAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainAlias)(nil)).Elem()
}

func (i DomainAliasMap) ToDomainAliasMapOutput() DomainAliasMapOutput {
	return i.ToDomainAliasMapOutputWithContext(context.Background())
}

func (i DomainAliasMap) ToDomainAliasMapOutputWithContext(ctx context.Context) DomainAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainAliasMapOutput)
}

type DomainAliasOutput struct{ *pulumi.OutputState }

func (DomainAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainAlias)(nil)).Elem()
}

func (o DomainAliasOutput) ToDomainAliasOutput() DomainAliasOutput {
	return o
}

func (o DomainAliasOutput) ToDomainAliasOutputWithContext(ctx context.Context) DomainAliasOutput {
	return o
}

// Domain.
func (o DomainAliasOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainAlias) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Domain alias.
func (o DomainAliasOutput) DomainAlias() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainAlias) pulumi.StringOutput { return v.DomainAlias }).(pulumi.StringOutput)
}

// Domain alias ID.
func (o DomainAliasOutput) DomainAliasId() pulumi.IntOutput {
	return o.ApplyT(func(v *DomainAlias) pulumi.IntOutput { return v.DomainAliasId }).(pulumi.IntOutput)
}

type DomainAliasArrayOutput struct{ *pulumi.OutputState }

func (DomainAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainAlias)(nil)).Elem()
}

func (o DomainAliasArrayOutput) ToDomainAliasArrayOutput() DomainAliasArrayOutput {
	return o
}

func (o DomainAliasArrayOutput) ToDomainAliasArrayOutputWithContext(ctx context.Context) DomainAliasArrayOutput {
	return o
}

func (o DomainAliasArrayOutput) Index(i pulumi.IntInput) DomainAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainAlias {
		return vs[0].([]*DomainAlias)[vs[1].(int)]
	}).(DomainAliasOutput)
}

type DomainAliasMapOutput struct{ *pulumi.OutputState }

func (DomainAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainAlias)(nil)).Elem()
}

func (o DomainAliasMapOutput) ToDomainAliasMapOutput() DomainAliasMapOutput {
	return o
}

func (o DomainAliasMapOutput) ToDomainAliasMapOutputWithContext(ctx context.Context) DomainAliasMapOutput {
	return o
}

func (o DomainAliasMapOutput) MapIndex(k pulumi.StringInput) DomainAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainAlias {
		return vs[0].(map[string]*DomainAlias)[vs[1].(string)]
	}).(DomainAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainAliasInput)(nil)).Elem(), &DomainAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainAliasArrayInput)(nil)).Elem(), DomainAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainAliasMapInput)(nil)).Elem(), DomainAliasMap{})
	pulumi.RegisterOutputType(DomainAliasOutput{})
	pulumi.RegisterOutputType(DomainAliasArrayOutput{})
	pulumi.RegisterOutputType(DomainAliasMapOutput{})
}
