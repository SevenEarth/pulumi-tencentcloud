// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a scf functionAlias
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Scf.NewFunctionAlias(ctx, "functionAlias", &Scf.FunctionAliasArgs{
// 			Description:     pulumi.String("matchs for test 12312312"),
// 			FunctionName:    pulumi.String("keep-1676351130"),
// 			FunctionVersion: pulumi.String("3"),
// 			Namespace:       pulumi.String("default"),
// 			RoutingConfig: &scf.FunctionAliasRoutingConfigArgs{
// 				AdditionalVersionMatches: scf.FunctionAliasRoutingConfigAdditionalVersionMatchArray{
// 					&scf.FunctionAliasRoutingConfigAdditionalVersionMatchArgs{
// 						Expression: pulumi.String("testuser"),
// 						Key:        pulumi.String("invoke.headers.User"),
// 						Method:     pulumi.String("exact"),
// 						Version:    pulumi.String("2"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// scf function_alias can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Scf/functionAlias:FunctionAlias function_alias namespace#functionName#name
// ```
type FunctionAlias struct {
	pulumi.CustomResourceState

	// Alias description information.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Function name.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// Master version pointed to by the alias.
	FunctionVersion pulumi.StringOutput `pulumi:"functionVersion"`
	// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Function namespace.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Request routing configuration of alias.
	RoutingConfig FunctionAliasRoutingConfigPtrOutput `pulumi:"routingConfig"`
}

// NewFunctionAlias registers a new resource with the given unique name, arguments, and options.
func NewFunctionAlias(ctx *pulumi.Context,
	name string, args *FunctionAliasArgs, opts ...pulumi.ResourceOption) (*FunctionAlias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	if args.FunctionVersion == nil {
		return nil, errors.New("invalid value for required argument 'FunctionVersion'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FunctionAlias
	err := ctx.RegisterResource("tencentcloud:Scf/functionAlias:FunctionAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionAlias gets an existing FunctionAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionAliasState, opts ...pulumi.ResourceOption) (*FunctionAlias, error) {
	var resource FunctionAlias
	err := ctx.ReadResource("tencentcloud:Scf/functionAlias:FunctionAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionAlias resources.
type functionAliasState struct {
	// Alias description information.
	Description *string `pulumi:"description"`
	// Function name.
	FunctionName *string `pulumi:"functionName"`
	// Master version pointed to by the alias.
	FunctionVersion *string `pulumi:"functionVersion"`
	// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
	Name *string `pulumi:"name"`
	// Function namespace.
	Namespace *string `pulumi:"namespace"`
	// Request routing configuration of alias.
	RoutingConfig *FunctionAliasRoutingConfig `pulumi:"routingConfig"`
}

type FunctionAliasState struct {
	// Alias description information.
	Description pulumi.StringPtrInput
	// Function name.
	FunctionName pulumi.StringPtrInput
	// Master version pointed to by the alias.
	FunctionVersion pulumi.StringPtrInput
	// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
	Name pulumi.StringPtrInput
	// Function namespace.
	Namespace pulumi.StringPtrInput
	// Request routing configuration of alias.
	RoutingConfig FunctionAliasRoutingConfigPtrInput
}

func (FunctionAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionAliasState)(nil)).Elem()
}

type functionAliasArgs struct {
	// Alias description information.
	Description *string `pulumi:"description"`
	// Function name.
	FunctionName string `pulumi:"functionName"`
	// Master version pointed to by the alias.
	FunctionVersion string `pulumi:"functionVersion"`
	// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
	Name *string `pulumi:"name"`
	// Function namespace.
	Namespace *string `pulumi:"namespace"`
	// Request routing configuration of alias.
	RoutingConfig *FunctionAliasRoutingConfig `pulumi:"routingConfig"`
}

// The set of arguments for constructing a FunctionAlias resource.
type FunctionAliasArgs struct {
	// Alias description information.
	Description pulumi.StringPtrInput
	// Function name.
	FunctionName pulumi.StringInput
	// Master version pointed to by the alias.
	FunctionVersion pulumi.StringInput
	// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
	Name pulumi.StringPtrInput
	// Function namespace.
	Namespace pulumi.StringPtrInput
	// Request routing configuration of alias.
	RoutingConfig FunctionAliasRoutingConfigPtrInput
}

func (FunctionAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionAliasArgs)(nil)).Elem()
}

type FunctionAliasInput interface {
	pulumi.Input

	ToFunctionAliasOutput() FunctionAliasOutput
	ToFunctionAliasOutputWithContext(ctx context.Context) FunctionAliasOutput
}

func (*FunctionAlias) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionAlias)(nil)).Elem()
}

func (i *FunctionAlias) ToFunctionAliasOutput() FunctionAliasOutput {
	return i.ToFunctionAliasOutputWithContext(context.Background())
}

func (i *FunctionAlias) ToFunctionAliasOutputWithContext(ctx context.Context) FunctionAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAliasOutput)
}

// FunctionAliasArrayInput is an input type that accepts FunctionAliasArray and FunctionAliasArrayOutput values.
// You can construct a concrete instance of `FunctionAliasArrayInput` via:
//
//          FunctionAliasArray{ FunctionAliasArgs{...} }
type FunctionAliasArrayInput interface {
	pulumi.Input

	ToFunctionAliasArrayOutput() FunctionAliasArrayOutput
	ToFunctionAliasArrayOutputWithContext(context.Context) FunctionAliasArrayOutput
}

type FunctionAliasArray []FunctionAliasInput

func (FunctionAliasArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionAlias)(nil)).Elem()
}

func (i FunctionAliasArray) ToFunctionAliasArrayOutput() FunctionAliasArrayOutput {
	return i.ToFunctionAliasArrayOutputWithContext(context.Background())
}

func (i FunctionAliasArray) ToFunctionAliasArrayOutputWithContext(ctx context.Context) FunctionAliasArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAliasArrayOutput)
}

// FunctionAliasMapInput is an input type that accepts FunctionAliasMap and FunctionAliasMapOutput values.
// You can construct a concrete instance of `FunctionAliasMapInput` via:
//
//          FunctionAliasMap{ "key": FunctionAliasArgs{...} }
type FunctionAliasMapInput interface {
	pulumi.Input

	ToFunctionAliasMapOutput() FunctionAliasMapOutput
	ToFunctionAliasMapOutputWithContext(context.Context) FunctionAliasMapOutput
}

type FunctionAliasMap map[string]FunctionAliasInput

func (FunctionAliasMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionAlias)(nil)).Elem()
}

func (i FunctionAliasMap) ToFunctionAliasMapOutput() FunctionAliasMapOutput {
	return i.ToFunctionAliasMapOutputWithContext(context.Background())
}

func (i FunctionAliasMap) ToFunctionAliasMapOutputWithContext(ctx context.Context) FunctionAliasMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionAliasMapOutput)
}

type FunctionAliasOutput struct{ *pulumi.OutputState }

func (FunctionAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionAlias)(nil)).Elem()
}

func (o FunctionAliasOutput) ToFunctionAliasOutput() FunctionAliasOutput {
	return o
}

func (o FunctionAliasOutput) ToFunctionAliasOutputWithContext(ctx context.Context) FunctionAliasOutput {
	return o
}

// Alias description information.
func (o FunctionAliasOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionAlias) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Function name.
func (o FunctionAliasOutput) FunctionName() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionAlias) pulumi.StringOutput { return v.FunctionName }).(pulumi.StringOutput)
}

// Master version pointed to by the alias.
func (o FunctionAliasOutput) FunctionVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionAlias) pulumi.StringOutput { return v.FunctionVersion }).(pulumi.StringOutput)
}

// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
func (o FunctionAliasOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionAlias) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Function namespace.
func (o FunctionAliasOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionAlias) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Request routing configuration of alias.
func (o FunctionAliasOutput) RoutingConfig() FunctionAliasRoutingConfigPtrOutput {
	return o.ApplyT(func(v *FunctionAlias) FunctionAliasRoutingConfigPtrOutput { return v.RoutingConfig }).(FunctionAliasRoutingConfigPtrOutput)
}

type FunctionAliasArrayOutput struct{ *pulumi.OutputState }

func (FunctionAliasArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionAlias)(nil)).Elem()
}

func (o FunctionAliasArrayOutput) ToFunctionAliasArrayOutput() FunctionAliasArrayOutput {
	return o
}

func (o FunctionAliasArrayOutput) ToFunctionAliasArrayOutputWithContext(ctx context.Context) FunctionAliasArrayOutput {
	return o
}

func (o FunctionAliasArrayOutput) Index(i pulumi.IntInput) FunctionAliasOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionAlias {
		return vs[0].([]*FunctionAlias)[vs[1].(int)]
	}).(FunctionAliasOutput)
}

type FunctionAliasMapOutput struct{ *pulumi.OutputState }

func (FunctionAliasMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionAlias)(nil)).Elem()
}

func (o FunctionAliasMapOutput) ToFunctionAliasMapOutput() FunctionAliasMapOutput {
	return o
}

func (o FunctionAliasMapOutput) ToFunctionAliasMapOutputWithContext(ctx context.Context) FunctionAliasMapOutput {
	return o
}

func (o FunctionAliasMapOutput) MapIndex(k pulumi.StringInput) FunctionAliasOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionAlias {
		return vs[0].(map[string]*FunctionAlias)[vs[1].(string)]
	}).(FunctionAliasOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionAliasInput)(nil)).Elem(), &FunctionAlias{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionAliasArrayInput)(nil)).Elem(), FunctionAliasArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionAliasMapInput)(nil)).Elem(), FunctionAliasMap{})
	pulumi.RegisterOutputType(FunctionAliasOutput{})
	pulumi.RegisterOutputType(FunctionAliasArrayOutput{})
	pulumi.RegisterOutputType(FunctionAliasMapOutput{})
}
