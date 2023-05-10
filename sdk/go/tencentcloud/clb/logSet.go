// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create an exclusive CLB Logset.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Clb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Clb.NewLogSet(ctx, "foo", &Clb.LogSetArgs{
// 			Period: pulumi.Int(7),
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
// CLB log set can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Clb/logSet:LogSet foo 4eb9e3a8-9c42-4b32-9ddf-e215e9c92764
// ```
type LogSet struct {
	pulumi.CustomResourceState

	// Logset creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Logset name, which unique and fixed `clbLogset` among all CLS logsets.
	Name pulumi.StringOutput `pulumi:"name"`
	// Logset retention period in days. Maximun value is `90`.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Number of log topics in logset.
	TopicCount pulumi.StringOutput `pulumi:"topicCount"`
}

// NewLogSet registers a new resource with the given unique name, arguments, and options.
func NewLogSet(ctx *pulumi.Context,
	name string, args *LogSetArgs, opts ...pulumi.ResourceOption) (*LogSet, error) {
	if args == nil {
		args = &LogSetArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource LogSet
	err := ctx.RegisterResource("tencentcloud:Clb/logSet:LogSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSet gets an existing LogSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSetState, opts ...pulumi.ResourceOption) (*LogSet, error) {
	var resource LogSet
	err := ctx.ReadResource("tencentcloud:Clb/logSet:LogSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSet resources.
type logSetState struct {
	// Logset creation time.
	CreateTime *string `pulumi:"createTime"`
	// Logset name, which unique and fixed `clbLogset` among all CLS logsets.
	Name *string `pulumi:"name"`
	// Logset retention period in days. Maximun value is `90`.
	Period *int `pulumi:"period"`
	// Number of log topics in logset.
	TopicCount *string `pulumi:"topicCount"`
}

type LogSetState struct {
	// Logset creation time.
	CreateTime pulumi.StringPtrInput
	// Logset name, which unique and fixed `clbLogset` among all CLS logsets.
	Name pulumi.StringPtrInput
	// Logset retention period in days. Maximun value is `90`.
	Period pulumi.IntPtrInput
	// Number of log topics in logset.
	TopicCount pulumi.StringPtrInput
}

func (LogSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSetState)(nil)).Elem()
}

type logSetArgs struct {
	// Logset retention period in days. Maximun value is `90`.
	Period *int `pulumi:"period"`
}

// The set of arguments for constructing a LogSet resource.
type LogSetArgs struct {
	// Logset retention period in days. Maximun value is `90`.
	Period pulumi.IntPtrInput
}

func (LogSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSetArgs)(nil)).Elem()
}

type LogSetInput interface {
	pulumi.Input

	ToLogSetOutput() LogSetOutput
	ToLogSetOutputWithContext(ctx context.Context) LogSetOutput
}

func (*LogSet) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSet)(nil)).Elem()
}

func (i *LogSet) ToLogSetOutput() LogSetOutput {
	return i.ToLogSetOutputWithContext(context.Background())
}

func (i *LogSet) ToLogSetOutputWithContext(ctx context.Context) LogSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSetOutput)
}

// LogSetArrayInput is an input type that accepts LogSetArray and LogSetArrayOutput values.
// You can construct a concrete instance of `LogSetArrayInput` via:
//
//          LogSetArray{ LogSetArgs{...} }
type LogSetArrayInput interface {
	pulumi.Input

	ToLogSetArrayOutput() LogSetArrayOutput
	ToLogSetArrayOutputWithContext(context.Context) LogSetArrayOutput
}

type LogSetArray []LogSetInput

func (LogSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSet)(nil)).Elem()
}

func (i LogSetArray) ToLogSetArrayOutput() LogSetArrayOutput {
	return i.ToLogSetArrayOutputWithContext(context.Background())
}

func (i LogSetArray) ToLogSetArrayOutputWithContext(ctx context.Context) LogSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSetArrayOutput)
}

// LogSetMapInput is an input type that accepts LogSetMap and LogSetMapOutput values.
// You can construct a concrete instance of `LogSetMapInput` via:
//
//          LogSetMap{ "key": LogSetArgs{...} }
type LogSetMapInput interface {
	pulumi.Input

	ToLogSetMapOutput() LogSetMapOutput
	ToLogSetMapOutputWithContext(context.Context) LogSetMapOutput
}

type LogSetMap map[string]LogSetInput

func (LogSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSet)(nil)).Elem()
}

func (i LogSetMap) ToLogSetMapOutput() LogSetMapOutput {
	return i.ToLogSetMapOutputWithContext(context.Background())
}

func (i LogSetMap) ToLogSetMapOutputWithContext(ctx context.Context) LogSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSetMapOutput)
}

type LogSetOutput struct{ *pulumi.OutputState }

func (LogSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogSet)(nil)).Elem()
}

func (o LogSetOutput) ToLogSetOutput() LogSetOutput {
	return o
}

func (o LogSetOutput) ToLogSetOutputWithContext(ctx context.Context) LogSetOutput {
	return o
}

// Logset creation time.
func (o LogSetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSet) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Logset name, which unique and fixed `clbLogset` among all CLS logsets.
func (o LogSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Logset retention period in days. Maximun value is `90`.
func (o LogSetOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LogSet) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Number of log topics in logset.
func (o LogSetOutput) TopicCount() pulumi.StringOutput {
	return o.ApplyT(func(v *LogSet) pulumi.StringOutput { return v.TopicCount }).(pulumi.StringOutput)
}

type LogSetArrayOutput struct{ *pulumi.OutputState }

func (LogSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogSet)(nil)).Elem()
}

func (o LogSetArrayOutput) ToLogSetArrayOutput() LogSetArrayOutput {
	return o
}

func (o LogSetArrayOutput) ToLogSetArrayOutputWithContext(ctx context.Context) LogSetArrayOutput {
	return o
}

func (o LogSetArrayOutput) Index(i pulumi.IntInput) LogSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogSet {
		return vs[0].([]*LogSet)[vs[1].(int)]
	}).(LogSetOutput)
}

type LogSetMapOutput struct{ *pulumi.OutputState }

func (LogSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogSet)(nil)).Elem()
}

func (o LogSetMapOutput) ToLogSetMapOutput() LogSetMapOutput {
	return o
}

func (o LogSetMapOutput) ToLogSetMapOutputWithContext(ctx context.Context) LogSetMapOutput {
	return o
}

func (o LogSetMapOutput) MapIndex(k pulumi.StringInput) LogSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogSet {
		return vs[0].(map[string]*LogSet)[vs[1].(string)]
	}).(LogSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogSetInput)(nil)).Elem(), &LogSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSetArrayInput)(nil)).Elem(), LogSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogSetMapInput)(nil)).Elem(), LogSetMap{})
	pulumi.RegisterOutputType(LogSetOutput{})
	pulumi.RegisterOutputType(LogSetArrayOutput{})
	pulumi.RegisterOutputType(LogSetMapOutput{})
}
