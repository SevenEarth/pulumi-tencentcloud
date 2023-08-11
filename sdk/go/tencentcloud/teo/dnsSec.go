// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DnsSec struct {
	pulumi.CustomResourceState

	// DNSSEC infos.
	Dnssec DnsSecDnssecOutput `pulumi:"dnssec"`
	// Last modification date.
	ModifiedOn pulumi.StringOutput `pulumi:"modifiedOn"`
	// DNSSEC status. Valid values: `enabled`, `disabled`.
	Status pulumi.StringOutput `pulumi:"status"`
	// Site ID.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewDnsSec registers a new resource with the given unique name, arguments, and options.
func NewDnsSec(ctx *pulumi.Context,
	name string, args *DnsSecArgs, opts ...pulumi.ResourceOption) (*DnsSec, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DnsSec
	err := ctx.RegisterResource("tencentcloud:Teo/dnsSec:DnsSec", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsSec gets an existing DnsSec resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsSec(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsSecState, opts ...pulumi.ResourceOption) (*DnsSec, error) {
	var resource DnsSec
	err := ctx.ReadResource("tencentcloud:Teo/dnsSec:DnsSec", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsSec resources.
type dnsSecState struct {
	// DNSSEC infos.
	Dnssec *DnsSecDnssec `pulumi:"dnssec"`
	// Last modification date.
	ModifiedOn *string `pulumi:"modifiedOn"`
	// DNSSEC status. Valid values: `enabled`, `disabled`.
	Status *string `pulumi:"status"`
	// Site ID.
	ZoneId *string `pulumi:"zoneId"`
}

type DnsSecState struct {
	// DNSSEC infos.
	Dnssec DnsSecDnssecPtrInput
	// Last modification date.
	ModifiedOn pulumi.StringPtrInput
	// DNSSEC status. Valid values: `enabled`, `disabled`.
	Status pulumi.StringPtrInput
	// Site ID.
	ZoneId pulumi.StringPtrInput
}

func (DnsSecState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsSecState)(nil)).Elem()
}

type dnsSecArgs struct {
	// DNSSEC infos.
	Dnssec *DnsSecDnssec `pulumi:"dnssec"`
	// DNSSEC status. Valid values: `enabled`, `disabled`.
	Status string `pulumi:"status"`
	// Site ID.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a DnsSec resource.
type DnsSecArgs struct {
	// DNSSEC infos.
	Dnssec DnsSecDnssecPtrInput
	// DNSSEC status. Valid values: `enabled`, `disabled`.
	Status pulumi.StringInput
	// Site ID.
	ZoneId pulumi.StringInput
}

func (DnsSecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsSecArgs)(nil)).Elem()
}

type DnsSecInput interface {
	pulumi.Input

	ToDnsSecOutput() DnsSecOutput
	ToDnsSecOutputWithContext(ctx context.Context) DnsSecOutput
}

func (*DnsSec) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsSec)(nil)).Elem()
}

func (i *DnsSec) ToDnsSecOutput() DnsSecOutput {
	return i.ToDnsSecOutputWithContext(context.Background())
}

func (i *DnsSec) ToDnsSecOutputWithContext(ctx context.Context) DnsSecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSecOutput)
}

// DnsSecArrayInput is an input type that accepts DnsSecArray and DnsSecArrayOutput values.
// You can construct a concrete instance of `DnsSecArrayInput` via:
//
//          DnsSecArray{ DnsSecArgs{...} }
type DnsSecArrayInput interface {
	pulumi.Input

	ToDnsSecArrayOutput() DnsSecArrayOutput
	ToDnsSecArrayOutputWithContext(context.Context) DnsSecArrayOutput
}

type DnsSecArray []DnsSecInput

func (DnsSecArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsSec)(nil)).Elem()
}

func (i DnsSecArray) ToDnsSecArrayOutput() DnsSecArrayOutput {
	return i.ToDnsSecArrayOutputWithContext(context.Background())
}

func (i DnsSecArray) ToDnsSecArrayOutputWithContext(ctx context.Context) DnsSecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSecArrayOutput)
}

// DnsSecMapInput is an input type that accepts DnsSecMap and DnsSecMapOutput values.
// You can construct a concrete instance of `DnsSecMapInput` via:
//
//          DnsSecMap{ "key": DnsSecArgs{...} }
type DnsSecMapInput interface {
	pulumi.Input

	ToDnsSecMapOutput() DnsSecMapOutput
	ToDnsSecMapOutputWithContext(context.Context) DnsSecMapOutput
}

type DnsSecMap map[string]DnsSecInput

func (DnsSecMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsSec)(nil)).Elem()
}

func (i DnsSecMap) ToDnsSecMapOutput() DnsSecMapOutput {
	return i.ToDnsSecMapOutputWithContext(context.Background())
}

func (i DnsSecMap) ToDnsSecMapOutputWithContext(ctx context.Context) DnsSecMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSecMapOutput)
}

type DnsSecOutput struct{ *pulumi.OutputState }

func (DnsSecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsSec)(nil)).Elem()
}

func (o DnsSecOutput) ToDnsSecOutput() DnsSecOutput {
	return o
}

func (o DnsSecOutput) ToDnsSecOutputWithContext(ctx context.Context) DnsSecOutput {
	return o
}

// DNSSEC infos.
func (o DnsSecOutput) Dnssec() DnsSecDnssecOutput {
	return o.ApplyT(func(v *DnsSec) DnsSecDnssecOutput { return v.Dnssec }).(DnsSecDnssecOutput)
}

// Last modification date.
func (o DnsSecOutput) ModifiedOn() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsSec) pulumi.StringOutput { return v.ModifiedOn }).(pulumi.StringOutput)
}

// DNSSEC status. Valid values: `enabled`, `disabled`.
func (o DnsSecOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsSec) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Site ID.
func (o DnsSecOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsSec) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type DnsSecArrayOutput struct{ *pulumi.OutputState }

func (DnsSecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsSec)(nil)).Elem()
}

func (o DnsSecArrayOutput) ToDnsSecArrayOutput() DnsSecArrayOutput {
	return o
}

func (o DnsSecArrayOutput) ToDnsSecArrayOutputWithContext(ctx context.Context) DnsSecArrayOutput {
	return o
}

func (o DnsSecArrayOutput) Index(i pulumi.IntInput) DnsSecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsSec {
		return vs[0].([]*DnsSec)[vs[1].(int)]
	}).(DnsSecOutput)
}

type DnsSecMapOutput struct{ *pulumi.OutputState }

func (DnsSecMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsSec)(nil)).Elem()
}

func (o DnsSecMapOutput) ToDnsSecMapOutput() DnsSecMapOutput {
	return o
}

func (o DnsSecMapOutput) ToDnsSecMapOutputWithContext(ctx context.Context) DnsSecMapOutput {
	return o
}

func (o DnsSecMapOutput) MapIndex(k pulumi.StringInput) DnsSecOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsSec {
		return vs[0].(map[string]*DnsSec)[vs[1].(string)]
	}).(DnsSecOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSecInput)(nil)).Elem(), &DnsSec{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSecArrayInput)(nil)).Elem(), DnsSecArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSecMapInput)(nil)).Elem(), DnsSecMap{})
	pulumi.RegisterOutputType(DnsSecOutput{})
	pulumi.RegisterOutputType(DnsSecArrayOutput{})
	pulumi.RegisterOutputType(DnsSecMapOutput{})
}
