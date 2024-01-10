// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tsf pathRewrite
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.NewPathRewrite(ctx, "pathRewrite", &Tsf.PathRewriteArgs{
//				Blocked:        pulumi.String("N"),
//				GatewayGroupId: pulumi.String("group-a2j9zxpv"),
//				Order:          pulumi.Int(2),
//				Regex:          pulumi.String("/test"),
//				Replacement:    pulumi.String("/tt"),
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
// tsf path_rewrite can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Tsf/pathRewrite:PathRewrite path_rewrite rewrite-nygq33v2
//
// ```
type PathRewrite struct {
	pulumi.CustomResourceState

	// Whether to shield the mapped path, Y: Yes N: No.
	Blocked pulumi.StringOutput `pulumi:"blocked"`
	// gateway deployment group ID.
	GatewayGroupId pulumi.StringOutput `pulumi:"gatewayGroupId"`
	// rule order, the smaller the higher the priority.
	Order pulumi.IntOutput `pulumi:"order"`
	// path rewrite rule ID.
	PathRewriteId pulumi.StringOutput `pulumi:"pathRewriteId"`
	// regular expression.
	Regex pulumi.StringOutput `pulumi:"regex"`
	// content to replace.
	Replacement pulumi.StringOutput `pulumi:"replacement"`
}

// NewPathRewrite registers a new resource with the given unique name, arguments, and options.
func NewPathRewrite(ctx *pulumi.Context,
	name string, args *PathRewriteArgs, opts ...pulumi.ResourceOption) (*PathRewrite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Blocked == nil {
		return nil, errors.New("invalid value for required argument 'Blocked'")
	}
	if args.GatewayGroupId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayGroupId'")
	}
	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.Regex == nil {
		return nil, errors.New("invalid value for required argument 'Regex'")
	}
	if args.Replacement == nil {
		return nil, errors.New("invalid value for required argument 'Replacement'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PathRewrite
	err := ctx.RegisterResource("tencentcloud:Tsf/pathRewrite:PathRewrite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPathRewrite gets an existing PathRewrite resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPathRewrite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PathRewriteState, opts ...pulumi.ResourceOption) (*PathRewrite, error) {
	var resource PathRewrite
	err := ctx.ReadResource("tencentcloud:Tsf/pathRewrite:PathRewrite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PathRewrite resources.
type pathRewriteState struct {
	// Whether to shield the mapped path, Y: Yes N: No.
	Blocked *string `pulumi:"blocked"`
	// gateway deployment group ID.
	GatewayGroupId *string `pulumi:"gatewayGroupId"`
	// rule order, the smaller the higher the priority.
	Order *int `pulumi:"order"`
	// path rewrite rule ID.
	PathRewriteId *string `pulumi:"pathRewriteId"`
	// regular expression.
	Regex *string `pulumi:"regex"`
	// content to replace.
	Replacement *string `pulumi:"replacement"`
}

type PathRewriteState struct {
	// Whether to shield the mapped path, Y: Yes N: No.
	Blocked pulumi.StringPtrInput
	// gateway deployment group ID.
	GatewayGroupId pulumi.StringPtrInput
	// rule order, the smaller the higher the priority.
	Order pulumi.IntPtrInput
	// path rewrite rule ID.
	PathRewriteId pulumi.StringPtrInput
	// regular expression.
	Regex pulumi.StringPtrInput
	// content to replace.
	Replacement pulumi.StringPtrInput
}

func (PathRewriteState) ElementType() reflect.Type {
	return reflect.TypeOf((*pathRewriteState)(nil)).Elem()
}

type pathRewriteArgs struct {
	// Whether to shield the mapped path, Y: Yes N: No.
	Blocked string `pulumi:"blocked"`
	// gateway deployment group ID.
	GatewayGroupId string `pulumi:"gatewayGroupId"`
	// rule order, the smaller the higher the priority.
	Order int `pulumi:"order"`
	// regular expression.
	Regex string `pulumi:"regex"`
	// content to replace.
	Replacement string `pulumi:"replacement"`
}

// The set of arguments for constructing a PathRewrite resource.
type PathRewriteArgs struct {
	// Whether to shield the mapped path, Y: Yes N: No.
	Blocked pulumi.StringInput
	// gateway deployment group ID.
	GatewayGroupId pulumi.StringInput
	// rule order, the smaller the higher the priority.
	Order pulumi.IntInput
	// regular expression.
	Regex pulumi.StringInput
	// content to replace.
	Replacement pulumi.StringInput
}

func (PathRewriteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pathRewriteArgs)(nil)).Elem()
}

type PathRewriteInput interface {
	pulumi.Input

	ToPathRewriteOutput() PathRewriteOutput
	ToPathRewriteOutputWithContext(ctx context.Context) PathRewriteOutput
}

func (*PathRewrite) ElementType() reflect.Type {
	return reflect.TypeOf((**PathRewrite)(nil)).Elem()
}

func (i *PathRewrite) ToPathRewriteOutput() PathRewriteOutput {
	return i.ToPathRewriteOutputWithContext(context.Background())
}

func (i *PathRewrite) ToPathRewriteOutputWithContext(ctx context.Context) PathRewriteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRewriteOutput)
}

// PathRewriteArrayInput is an input type that accepts PathRewriteArray and PathRewriteArrayOutput values.
// You can construct a concrete instance of `PathRewriteArrayInput` via:
//
//	PathRewriteArray{ PathRewriteArgs{...} }
type PathRewriteArrayInput interface {
	pulumi.Input

	ToPathRewriteArrayOutput() PathRewriteArrayOutput
	ToPathRewriteArrayOutputWithContext(context.Context) PathRewriteArrayOutput
}

type PathRewriteArray []PathRewriteInput

func (PathRewriteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PathRewrite)(nil)).Elem()
}

func (i PathRewriteArray) ToPathRewriteArrayOutput() PathRewriteArrayOutput {
	return i.ToPathRewriteArrayOutputWithContext(context.Background())
}

func (i PathRewriteArray) ToPathRewriteArrayOutputWithContext(ctx context.Context) PathRewriteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRewriteArrayOutput)
}

// PathRewriteMapInput is an input type that accepts PathRewriteMap and PathRewriteMapOutput values.
// You can construct a concrete instance of `PathRewriteMapInput` via:
//
//	PathRewriteMap{ "key": PathRewriteArgs{...} }
type PathRewriteMapInput interface {
	pulumi.Input

	ToPathRewriteMapOutput() PathRewriteMapOutput
	ToPathRewriteMapOutputWithContext(context.Context) PathRewriteMapOutput
}

type PathRewriteMap map[string]PathRewriteInput

func (PathRewriteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PathRewrite)(nil)).Elem()
}

func (i PathRewriteMap) ToPathRewriteMapOutput() PathRewriteMapOutput {
	return i.ToPathRewriteMapOutputWithContext(context.Background())
}

func (i PathRewriteMap) ToPathRewriteMapOutputWithContext(ctx context.Context) PathRewriteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PathRewriteMapOutput)
}

type PathRewriteOutput struct{ *pulumi.OutputState }

func (PathRewriteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PathRewrite)(nil)).Elem()
}

func (o PathRewriteOutput) ToPathRewriteOutput() PathRewriteOutput {
	return o
}

func (o PathRewriteOutput) ToPathRewriteOutputWithContext(ctx context.Context) PathRewriteOutput {
	return o
}

// Whether to shield the mapped path, Y: Yes N: No.
func (o PathRewriteOutput) Blocked() pulumi.StringOutput {
	return o.ApplyT(func(v *PathRewrite) pulumi.StringOutput { return v.Blocked }).(pulumi.StringOutput)
}

// gateway deployment group ID.
func (o PathRewriteOutput) GatewayGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *PathRewrite) pulumi.StringOutput { return v.GatewayGroupId }).(pulumi.StringOutput)
}

// rule order, the smaller the higher the priority.
func (o PathRewriteOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *PathRewrite) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// path rewrite rule ID.
func (o PathRewriteOutput) PathRewriteId() pulumi.StringOutput {
	return o.ApplyT(func(v *PathRewrite) pulumi.StringOutput { return v.PathRewriteId }).(pulumi.StringOutput)
}

// regular expression.
func (o PathRewriteOutput) Regex() pulumi.StringOutput {
	return o.ApplyT(func(v *PathRewrite) pulumi.StringOutput { return v.Regex }).(pulumi.StringOutput)
}

// content to replace.
func (o PathRewriteOutput) Replacement() pulumi.StringOutput {
	return o.ApplyT(func(v *PathRewrite) pulumi.StringOutput { return v.Replacement }).(pulumi.StringOutput)
}

type PathRewriteArrayOutput struct{ *pulumi.OutputState }

func (PathRewriteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PathRewrite)(nil)).Elem()
}

func (o PathRewriteArrayOutput) ToPathRewriteArrayOutput() PathRewriteArrayOutput {
	return o
}

func (o PathRewriteArrayOutput) ToPathRewriteArrayOutputWithContext(ctx context.Context) PathRewriteArrayOutput {
	return o
}

func (o PathRewriteArrayOutput) Index(i pulumi.IntInput) PathRewriteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PathRewrite {
		return vs[0].([]*PathRewrite)[vs[1].(int)]
	}).(PathRewriteOutput)
}

type PathRewriteMapOutput struct{ *pulumi.OutputState }

func (PathRewriteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PathRewrite)(nil)).Elem()
}

func (o PathRewriteMapOutput) ToPathRewriteMapOutput() PathRewriteMapOutput {
	return o
}

func (o PathRewriteMapOutput) ToPathRewriteMapOutputWithContext(ctx context.Context) PathRewriteMapOutput {
	return o
}

func (o PathRewriteMapOutput) MapIndex(k pulumi.StringInput) PathRewriteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PathRewrite {
		return vs[0].(map[string]*PathRewrite)[vs[1].(string)]
	}).(PathRewriteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PathRewriteInput)(nil)).Elem(), &PathRewrite{})
	pulumi.RegisterInputType(reflect.TypeOf((*PathRewriteArrayInput)(nil)).Elem(), PathRewriteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PathRewriteMapInput)(nil)).Elem(), PathRewriteMap{})
	pulumi.RegisterOutputType(PathRewriteOutput{})
	pulumi.RegisterOutputType(PathRewriteArrayOutput{})
	pulumi.RegisterOutputType(PathRewriteMapOutput{})
}
