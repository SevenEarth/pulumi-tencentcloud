// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc dhcpIp
//
// ## Import
//
// vpc dhcp_ip can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Vpc/dhcpIp:DhcpIp dhcp_ip dhcp_ip_id
//
// ```
type DhcpIp struct {
	pulumi.CustomResourceState

	// `DhcpIp` name.
	DhcpIpName pulumi.StringOutput `pulumi:"dhcpIpName"`
	// Subnet `ID`.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The private network `ID`.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDhcpIp registers a new resource with the given unique name, arguments, and options.
func NewDhcpIp(ctx *pulumi.Context,
	name string, args *DhcpIpArgs, opts ...pulumi.ResourceOption) (*DhcpIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DhcpIpName == nil {
		return nil, errors.New("invalid value for required argument 'DhcpIpName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DhcpIp
	err := ctx.RegisterResource("tencentcloud:Vpc/dhcpIp:DhcpIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDhcpIp gets an existing DhcpIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDhcpIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DhcpIpState, opts ...pulumi.ResourceOption) (*DhcpIp, error) {
	var resource DhcpIp
	err := ctx.ReadResource("tencentcloud:Vpc/dhcpIp:DhcpIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DhcpIp resources.
type dhcpIpState struct {
	// `DhcpIp` name.
	DhcpIpName *string `pulumi:"dhcpIpName"`
	// Subnet `ID`.
	SubnetId *string `pulumi:"subnetId"`
	// The private network `ID`.
	VpcId *string `pulumi:"vpcId"`
}

type DhcpIpState struct {
	// `DhcpIp` name.
	DhcpIpName pulumi.StringPtrInput
	// Subnet `ID`.
	SubnetId pulumi.StringPtrInput
	// The private network `ID`.
	VpcId pulumi.StringPtrInput
}

func (DhcpIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpIpState)(nil)).Elem()
}

type dhcpIpArgs struct {
	// `DhcpIp` name.
	DhcpIpName string `pulumi:"dhcpIpName"`
	// Subnet `ID`.
	SubnetId string `pulumi:"subnetId"`
	// The private network `ID`.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a DhcpIp resource.
type DhcpIpArgs struct {
	// `DhcpIp` name.
	DhcpIpName pulumi.StringInput
	// Subnet `ID`.
	SubnetId pulumi.StringInput
	// The private network `ID`.
	VpcId pulumi.StringInput
}

func (DhcpIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dhcpIpArgs)(nil)).Elem()
}

type DhcpIpInput interface {
	pulumi.Input

	ToDhcpIpOutput() DhcpIpOutput
	ToDhcpIpOutputWithContext(ctx context.Context) DhcpIpOutput
}

func (*DhcpIp) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpIp)(nil)).Elem()
}

func (i *DhcpIp) ToDhcpIpOutput() DhcpIpOutput {
	return i.ToDhcpIpOutputWithContext(context.Background())
}

func (i *DhcpIp) ToDhcpIpOutputWithContext(ctx context.Context) DhcpIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpIpOutput)
}

// DhcpIpArrayInput is an input type that accepts DhcpIpArray and DhcpIpArrayOutput values.
// You can construct a concrete instance of `DhcpIpArrayInput` via:
//
//	DhcpIpArray{ DhcpIpArgs{...} }
type DhcpIpArrayInput interface {
	pulumi.Input

	ToDhcpIpArrayOutput() DhcpIpArrayOutput
	ToDhcpIpArrayOutputWithContext(context.Context) DhcpIpArrayOutput
}

type DhcpIpArray []DhcpIpInput

func (DhcpIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpIp)(nil)).Elem()
}

func (i DhcpIpArray) ToDhcpIpArrayOutput() DhcpIpArrayOutput {
	return i.ToDhcpIpArrayOutputWithContext(context.Background())
}

func (i DhcpIpArray) ToDhcpIpArrayOutputWithContext(ctx context.Context) DhcpIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpIpArrayOutput)
}

// DhcpIpMapInput is an input type that accepts DhcpIpMap and DhcpIpMapOutput values.
// You can construct a concrete instance of `DhcpIpMapInput` via:
//
//	DhcpIpMap{ "key": DhcpIpArgs{...} }
type DhcpIpMapInput interface {
	pulumi.Input

	ToDhcpIpMapOutput() DhcpIpMapOutput
	ToDhcpIpMapOutputWithContext(context.Context) DhcpIpMapOutput
}

type DhcpIpMap map[string]DhcpIpInput

func (DhcpIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpIp)(nil)).Elem()
}

func (i DhcpIpMap) ToDhcpIpMapOutput() DhcpIpMapOutput {
	return i.ToDhcpIpMapOutputWithContext(context.Background())
}

func (i DhcpIpMap) ToDhcpIpMapOutputWithContext(ctx context.Context) DhcpIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DhcpIpMapOutput)
}

type DhcpIpOutput struct{ *pulumi.OutputState }

func (DhcpIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DhcpIp)(nil)).Elem()
}

func (o DhcpIpOutput) ToDhcpIpOutput() DhcpIpOutput {
	return o
}

func (o DhcpIpOutput) ToDhcpIpOutputWithContext(ctx context.Context) DhcpIpOutput {
	return o
}

// `DhcpIp` name.
func (o DhcpIpOutput) DhcpIpName() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpIp) pulumi.StringOutput { return v.DhcpIpName }).(pulumi.StringOutput)
}

// Subnet `ID`.
func (o DhcpIpOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpIp) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The private network `ID`.
func (o DhcpIpOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *DhcpIp) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type DhcpIpArrayOutput struct{ *pulumi.OutputState }

func (DhcpIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DhcpIp)(nil)).Elem()
}

func (o DhcpIpArrayOutput) ToDhcpIpArrayOutput() DhcpIpArrayOutput {
	return o
}

func (o DhcpIpArrayOutput) ToDhcpIpArrayOutputWithContext(ctx context.Context) DhcpIpArrayOutput {
	return o
}

func (o DhcpIpArrayOutput) Index(i pulumi.IntInput) DhcpIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DhcpIp {
		return vs[0].([]*DhcpIp)[vs[1].(int)]
	}).(DhcpIpOutput)
}

type DhcpIpMapOutput struct{ *pulumi.OutputState }

func (DhcpIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DhcpIp)(nil)).Elem()
}

func (o DhcpIpMapOutput) ToDhcpIpMapOutput() DhcpIpMapOutput {
	return o
}

func (o DhcpIpMapOutput) ToDhcpIpMapOutputWithContext(ctx context.Context) DhcpIpMapOutput {
	return o
}

func (o DhcpIpMapOutput) MapIndex(k pulumi.StringInput) DhcpIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DhcpIp {
		return vs[0].(map[string]*DhcpIp)[vs[1].(string)]
	}).(DhcpIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpIpInput)(nil)).Elem(), &DhcpIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpIpArrayInput)(nil)).Elem(), DhcpIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DhcpIpMapInput)(nil)).Elem(), DhcpIpMap{})
	pulumi.RegisterOutputType(DhcpIpOutput{})
	pulumi.RegisterOutputType(DhcpIpArrayOutput{})
	pulumi.RegisterOutputType(DhcpIpMapOutput{})
}
