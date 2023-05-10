// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create dayu new layer 7 rule
//
// > **NOTE:** This resource only support resource Anti-DDoS of type `bgpip`
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dayu.NewL7RuleV2(ctx, "tencentcloudDayuL7RuleV2", &Dayu.L7RuleV2Args{
// 			ResourceId:   pulumi.String("bgpip-000004xe"),
// 			ResourceIp:   pulumi.String("119.28.217.162"),
// 			ResourceType: pulumi.String("bgpip"),
// 			Rule: &dayu.L7RuleV2RuleArgs{
// 				Domain:     pulumi.String("github.com"),
// 				KeepEnable: pulumi.Int(false),
// 				Keeptime:   pulumi.Int(0),
// 				LbType:     pulumi.Int(1),
// 				Protocol:   pulumi.String("http"),
// 				SourceLists: dayu.L7RuleV2RuleSourceListArray{
// 					&dayu.L7RuleV2RuleSourceListArgs{
// 						Source: pulumi.String("1.2.3.5"),
// 						Weight: pulumi.Int(100),
// 					},
// 					&dayu.L7RuleV2RuleSourceListArgs{
// 						Source: pulumi.String("1.2.3.6"),
// 						Weight: pulumi.Int(100),
// 					},
// 				},
// 				SourceType: pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type L7RuleV2 struct {
	pulumi.CustomResourceState

	// ID of the resource that the layer 7 rule works for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Ip of the resource that the layer 7 rule works for.
	ResourceIp pulumi.StringOutput `pulumi:"resourceIp"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// A list of layer 7 rules. Each element contains the following attributes:
	Rule L7RuleV2RuleOutput `pulumi:"rule"`
}

// NewL7RuleV2 registers a new resource with the given unique name, arguments, and options.
func NewL7RuleV2(ctx *pulumi.Context,
	name string, args *L7RuleV2Args, opts ...pulumi.ResourceOption) (*L7RuleV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ResourceIp == nil {
		return nil, errors.New("invalid value for required argument 'ResourceIp'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.Rule == nil {
		return nil, errors.New("invalid value for required argument 'Rule'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource L7RuleV2
	err := ctx.RegisterResource("tencentcloud:Dayu/l7RuleV2:L7RuleV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetL7RuleV2 gets an existing L7RuleV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetL7RuleV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *L7RuleV2State, opts ...pulumi.ResourceOption) (*L7RuleV2, error) {
	var resource L7RuleV2
	err := ctx.ReadResource("tencentcloud:Dayu/l7RuleV2:L7RuleV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering L7RuleV2 resources.
type l7ruleV2State struct {
	// ID of the resource that the layer 7 rule works for.
	ResourceId *string `pulumi:"resourceId"`
	// Ip of the resource that the layer 7 rule works for.
	ResourceIp *string `pulumi:"resourceIp"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType *string `pulumi:"resourceType"`
	// A list of layer 7 rules. Each element contains the following attributes:
	Rule *L7RuleV2Rule `pulumi:"rule"`
}

type L7RuleV2State struct {
	// ID of the resource that the layer 7 rule works for.
	ResourceId pulumi.StringPtrInput
	// Ip of the resource that the layer 7 rule works for.
	ResourceIp pulumi.StringPtrInput
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType pulumi.StringPtrInput
	// A list of layer 7 rules. Each element contains the following attributes:
	Rule L7RuleV2RulePtrInput
}

func (L7RuleV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleV2State)(nil)).Elem()
}

type l7ruleV2Args struct {
	// ID of the resource that the layer 7 rule works for.
	ResourceId string `pulumi:"resourceId"`
	// Ip of the resource that the layer 7 rule works for.
	ResourceIp string `pulumi:"resourceIp"`
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType string `pulumi:"resourceType"`
	// A list of layer 7 rules. Each element contains the following attributes:
	Rule L7RuleV2Rule `pulumi:"rule"`
}

// The set of arguments for constructing a L7RuleV2 resource.
type L7RuleV2Args struct {
	// ID of the resource that the layer 7 rule works for.
	ResourceId pulumi.StringInput
	// Ip of the resource that the layer 7 rule works for.
	ResourceIp pulumi.StringInput
	// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
	ResourceType pulumi.StringInput
	// A list of layer 7 rules. Each element contains the following attributes:
	Rule L7RuleV2RuleInput
}

func (L7RuleV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*l7ruleV2Args)(nil)).Elem()
}

type L7RuleV2Input interface {
	pulumi.Input

	ToL7RuleV2Output() L7RuleV2Output
	ToL7RuleV2OutputWithContext(ctx context.Context) L7RuleV2Output
}

func (*L7RuleV2) ElementType() reflect.Type {
	return reflect.TypeOf((**L7RuleV2)(nil)).Elem()
}

func (i *L7RuleV2) ToL7RuleV2Output() L7RuleV2Output {
	return i.ToL7RuleV2OutputWithContext(context.Background())
}

func (i *L7RuleV2) ToL7RuleV2OutputWithContext(ctx context.Context) L7RuleV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(L7RuleV2Output)
}

// L7RuleV2ArrayInput is an input type that accepts L7RuleV2Array and L7RuleV2ArrayOutput values.
// You can construct a concrete instance of `L7RuleV2ArrayInput` via:
//
//          L7RuleV2Array{ L7RuleV2Args{...} }
type L7RuleV2ArrayInput interface {
	pulumi.Input

	ToL7RuleV2ArrayOutput() L7RuleV2ArrayOutput
	ToL7RuleV2ArrayOutputWithContext(context.Context) L7RuleV2ArrayOutput
}

type L7RuleV2Array []L7RuleV2Input

func (L7RuleV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7RuleV2)(nil)).Elem()
}

func (i L7RuleV2Array) ToL7RuleV2ArrayOutput() L7RuleV2ArrayOutput {
	return i.ToL7RuleV2ArrayOutputWithContext(context.Background())
}

func (i L7RuleV2Array) ToL7RuleV2ArrayOutputWithContext(ctx context.Context) L7RuleV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7RuleV2ArrayOutput)
}

// L7RuleV2MapInput is an input type that accepts L7RuleV2Map and L7RuleV2MapOutput values.
// You can construct a concrete instance of `L7RuleV2MapInput` via:
//
//          L7RuleV2Map{ "key": L7RuleV2Args{...} }
type L7RuleV2MapInput interface {
	pulumi.Input

	ToL7RuleV2MapOutput() L7RuleV2MapOutput
	ToL7RuleV2MapOutputWithContext(context.Context) L7RuleV2MapOutput
}

type L7RuleV2Map map[string]L7RuleV2Input

func (L7RuleV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7RuleV2)(nil)).Elem()
}

func (i L7RuleV2Map) ToL7RuleV2MapOutput() L7RuleV2MapOutput {
	return i.ToL7RuleV2MapOutputWithContext(context.Background())
}

func (i L7RuleV2Map) ToL7RuleV2MapOutputWithContext(ctx context.Context) L7RuleV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(L7RuleV2MapOutput)
}

type L7RuleV2Output struct{ *pulumi.OutputState }

func (L7RuleV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**L7RuleV2)(nil)).Elem()
}

func (o L7RuleV2Output) ToL7RuleV2Output() L7RuleV2Output {
	return o
}

func (o L7RuleV2Output) ToL7RuleV2OutputWithContext(ctx context.Context) L7RuleV2Output {
	return o
}

// ID of the resource that the layer 7 rule works for.
func (o L7RuleV2Output) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *L7RuleV2) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Ip of the resource that the layer 7 rule works for.
func (o L7RuleV2Output) ResourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *L7RuleV2) pulumi.StringOutput { return v.ResourceIp }).(pulumi.StringOutput)
}

// Type of the resource that the layer 7 rule works for, valid value is `bgpip`.
func (o L7RuleV2Output) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *L7RuleV2) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// A list of layer 7 rules. Each element contains the following attributes:
func (o L7RuleV2Output) Rule() L7RuleV2RuleOutput {
	return o.ApplyT(func(v *L7RuleV2) L7RuleV2RuleOutput { return v.Rule }).(L7RuleV2RuleOutput)
}

type L7RuleV2ArrayOutput struct{ *pulumi.OutputState }

func (L7RuleV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*L7RuleV2)(nil)).Elem()
}

func (o L7RuleV2ArrayOutput) ToL7RuleV2ArrayOutput() L7RuleV2ArrayOutput {
	return o
}

func (o L7RuleV2ArrayOutput) ToL7RuleV2ArrayOutputWithContext(ctx context.Context) L7RuleV2ArrayOutput {
	return o
}

func (o L7RuleV2ArrayOutput) Index(i pulumi.IntInput) L7RuleV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *L7RuleV2 {
		return vs[0].([]*L7RuleV2)[vs[1].(int)]
	}).(L7RuleV2Output)
}

type L7RuleV2MapOutput struct{ *pulumi.OutputState }

func (L7RuleV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*L7RuleV2)(nil)).Elem()
}

func (o L7RuleV2MapOutput) ToL7RuleV2MapOutput() L7RuleV2MapOutput {
	return o
}

func (o L7RuleV2MapOutput) ToL7RuleV2MapOutputWithContext(ctx context.Context) L7RuleV2MapOutput {
	return o
}

func (o L7RuleV2MapOutput) MapIndex(k pulumi.StringInput) L7RuleV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *L7RuleV2 {
		return vs[0].(map[string]*L7RuleV2)[vs[1].(string)]
	}).(L7RuleV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*L7RuleV2Input)(nil)).Elem(), &L7RuleV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7RuleV2ArrayInput)(nil)).Elem(), L7RuleV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*L7RuleV2MapInput)(nil)).Elem(), L7RuleV2Map{})
	pulumi.RegisterOutputType(L7RuleV2Output{})
	pulumi.RegisterOutputType(L7RuleV2ArrayOutput{})
	pulumi.RegisterOutputType(L7RuleV2MapOutput{})
}
