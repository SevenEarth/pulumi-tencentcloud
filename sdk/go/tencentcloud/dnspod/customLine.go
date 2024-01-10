// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dnspod

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
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
//			_, err := Dnspod.NewCustomLine(ctx, "customLine", &Dnspod.CustomLineArgs{
//				Area:   pulumi.String("6.6.6.1-6.6.6.2"),
//				Domain: pulumi.String("dnspod.com"),
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
// dnspod custom_line can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Dnspod/customLine:CustomLine custom_line domain#name
//
// ```
type CustomLine struct {
	pulumi.CustomResourceState

	// The IP segment of custom line, split with `-`.
	Area pulumi.StringOutput `pulumi:"area"`
	// Domain.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The Name of custom line.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewCustomLine registers a new resource with the given unique name, arguments, and options.
func NewCustomLine(ctx *pulumi.Context,
	name string, args *CustomLineArgs, opts ...pulumi.ResourceOption) (*CustomLine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Area == nil {
		return nil, errors.New("invalid value for required argument 'Area'")
	}
	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CustomLine
	err := ctx.RegisterResource("tencentcloud:Dnspod/customLine:CustomLine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomLine gets an existing CustomLine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomLine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomLineState, opts ...pulumi.ResourceOption) (*CustomLine, error) {
	var resource CustomLine
	err := ctx.ReadResource("tencentcloud:Dnspod/customLine:CustomLine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomLine resources.
type customLineState struct {
	// The IP segment of custom line, split with `-`.
	Area *string `pulumi:"area"`
	// Domain.
	Domain *string `pulumi:"domain"`
	// The Name of custom line.
	Name *string `pulumi:"name"`
}

type CustomLineState struct {
	// The IP segment of custom line, split with `-`.
	Area pulumi.StringPtrInput
	// Domain.
	Domain pulumi.StringPtrInput
	// The Name of custom line.
	Name pulumi.StringPtrInput
}

func (CustomLineState) ElementType() reflect.Type {
	return reflect.TypeOf((*customLineState)(nil)).Elem()
}

type customLineArgs struct {
	// The IP segment of custom line, split with `-`.
	Area string `pulumi:"area"`
	// Domain.
	Domain string `pulumi:"domain"`
	// The Name of custom line.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a CustomLine resource.
type CustomLineArgs struct {
	// The IP segment of custom line, split with `-`.
	Area pulumi.StringInput
	// Domain.
	Domain pulumi.StringInput
	// The Name of custom line.
	Name pulumi.StringPtrInput
}

func (CustomLineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customLineArgs)(nil)).Elem()
}

type CustomLineInput interface {
	pulumi.Input

	ToCustomLineOutput() CustomLineOutput
	ToCustomLineOutputWithContext(ctx context.Context) CustomLineOutput
}

func (*CustomLine) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLine)(nil)).Elem()
}

func (i *CustomLine) ToCustomLineOutput() CustomLineOutput {
	return i.ToCustomLineOutputWithContext(context.Background())
}

func (i *CustomLine) ToCustomLineOutputWithContext(ctx context.Context) CustomLineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLineOutput)
}

// CustomLineArrayInput is an input type that accepts CustomLineArray and CustomLineArrayOutput values.
// You can construct a concrete instance of `CustomLineArrayInput` via:
//
//	CustomLineArray{ CustomLineArgs{...} }
type CustomLineArrayInput interface {
	pulumi.Input

	ToCustomLineArrayOutput() CustomLineArrayOutput
	ToCustomLineArrayOutputWithContext(context.Context) CustomLineArrayOutput
}

type CustomLineArray []CustomLineInput

func (CustomLineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomLine)(nil)).Elem()
}

func (i CustomLineArray) ToCustomLineArrayOutput() CustomLineArrayOutput {
	return i.ToCustomLineArrayOutputWithContext(context.Background())
}

func (i CustomLineArray) ToCustomLineArrayOutputWithContext(ctx context.Context) CustomLineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLineArrayOutput)
}

// CustomLineMapInput is an input type that accepts CustomLineMap and CustomLineMapOutput values.
// You can construct a concrete instance of `CustomLineMapInput` via:
//
//	CustomLineMap{ "key": CustomLineArgs{...} }
type CustomLineMapInput interface {
	pulumi.Input

	ToCustomLineMapOutput() CustomLineMapOutput
	ToCustomLineMapOutputWithContext(context.Context) CustomLineMapOutput
}

type CustomLineMap map[string]CustomLineInput

func (CustomLineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomLine)(nil)).Elem()
}

func (i CustomLineMap) ToCustomLineMapOutput() CustomLineMapOutput {
	return i.ToCustomLineMapOutputWithContext(context.Background())
}

func (i CustomLineMap) ToCustomLineMapOutputWithContext(ctx context.Context) CustomLineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLineMapOutput)
}

type CustomLineOutput struct{ *pulumi.OutputState }

func (CustomLineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLine)(nil)).Elem()
}

func (o CustomLineOutput) ToCustomLineOutput() CustomLineOutput {
	return o
}

func (o CustomLineOutput) ToCustomLineOutputWithContext(ctx context.Context) CustomLineOutput {
	return o
}

// The IP segment of custom line, split with `-`.
func (o CustomLineOutput) Area() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomLine) pulumi.StringOutput { return v.Area }).(pulumi.StringOutput)
}

// Domain.
func (o CustomLineOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomLine) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The Name of custom line.
func (o CustomLineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomLine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type CustomLineArrayOutput struct{ *pulumi.OutputState }

func (CustomLineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomLine)(nil)).Elem()
}

func (o CustomLineArrayOutput) ToCustomLineArrayOutput() CustomLineArrayOutput {
	return o
}

func (o CustomLineArrayOutput) ToCustomLineArrayOutputWithContext(ctx context.Context) CustomLineArrayOutput {
	return o
}

func (o CustomLineArrayOutput) Index(i pulumi.IntInput) CustomLineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomLine {
		return vs[0].([]*CustomLine)[vs[1].(int)]
	}).(CustomLineOutput)
}

type CustomLineMapOutput struct{ *pulumi.OutputState }

func (CustomLineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomLine)(nil)).Elem()
}

func (o CustomLineMapOutput) ToCustomLineMapOutput() CustomLineMapOutput {
	return o
}

func (o CustomLineMapOutput) ToCustomLineMapOutputWithContext(ctx context.Context) CustomLineMapOutput {
	return o
}

func (o CustomLineMapOutput) MapIndex(k pulumi.StringInput) CustomLineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomLine {
		return vs[0].(map[string]*CustomLine)[vs[1].(string)]
	}).(CustomLineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLineInput)(nil)).Elem(), &CustomLine{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLineArrayInput)(nil)).Elem(), CustomLineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomLineMapInput)(nil)).Elem(), CustomLineMap{})
	pulumi.RegisterOutputType(CustomLineOutput{})
	pulumi.RegisterOutputType(CustomLineArrayOutput{})
	pulumi.RegisterOutputType(CustomLineMapOutput{})
}
