// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a cls export
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cls.NewExport(ctx, "export", &Cls.ExportArgs{
//				Format:   pulumi.String("json"),
//				From:     pulumi.Int(1607499107000),
//				LogCount: pulumi.Int(2),
//				Order:    pulumi.String("desc"),
//				Query:    pulumi.String("select count(*) as count"),
//				To:       pulumi.Int(1607499108000),
//				TopicId:  pulumi.String("7e34a3a7-635e-4da8-9005-88106c1fde69"),
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
// cls export can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Cls/export:Export export topic_id#export_id
// ```
type Export struct {
	pulumi.CustomResourceState

	// log export format.
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// export start time.
	From pulumi.IntOutput `pulumi:"from"`
	// export amount of log.
	LogCount pulumi.IntOutput `pulumi:"logCount"`
	// log export time sorting. desc or asc.
	Order pulumi.StringPtrOutput `pulumi:"order"`
	// export query rules.
	Query pulumi.StringOutput `pulumi:"query"`
	// export end time.
	To pulumi.IntOutput `pulumi:"to"`
	// topic id.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
}

// NewExport registers a new resource with the given unique name, arguments, and options.
func NewExport(ctx *pulumi.Context,
	name string, args *ExportArgs, opts ...pulumi.ResourceOption) (*Export, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.From == nil {
		return nil, errors.New("invalid value for required argument 'From'")
	}
	if args.LogCount == nil {
		return nil, errors.New("invalid value for required argument 'LogCount'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.To == nil {
		return nil, errors.New("invalid value for required argument 'To'")
	}
	if args.TopicId == nil {
		return nil, errors.New("invalid value for required argument 'TopicId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Export
	err := ctx.RegisterResource("tencentcloud:Cls/export:Export", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExport gets an existing Export resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExportState, opts ...pulumi.ResourceOption) (*Export, error) {
	var resource Export
	err := ctx.ReadResource("tencentcloud:Cls/export:Export", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Export resources.
type exportState struct {
	// log export format.
	Format *string `pulumi:"format"`
	// export start time.
	From *int `pulumi:"from"`
	// export amount of log.
	LogCount *int `pulumi:"logCount"`
	// log export time sorting. desc or asc.
	Order *string `pulumi:"order"`
	// export query rules.
	Query *string `pulumi:"query"`
	// export end time.
	To *int `pulumi:"to"`
	// topic id.
	TopicId *string `pulumi:"topicId"`
}

type ExportState struct {
	// log export format.
	Format pulumi.StringPtrInput
	// export start time.
	From pulumi.IntPtrInput
	// export amount of log.
	LogCount pulumi.IntPtrInput
	// log export time sorting. desc or asc.
	Order pulumi.StringPtrInput
	// export query rules.
	Query pulumi.StringPtrInput
	// export end time.
	To pulumi.IntPtrInput
	// topic id.
	TopicId pulumi.StringPtrInput
}

func (ExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*exportState)(nil)).Elem()
}

type exportArgs struct {
	// log export format.
	Format *string `pulumi:"format"`
	// export start time.
	From int `pulumi:"from"`
	// export amount of log.
	LogCount int `pulumi:"logCount"`
	// log export time sorting. desc or asc.
	Order *string `pulumi:"order"`
	// export query rules.
	Query string `pulumi:"query"`
	// export end time.
	To int `pulumi:"to"`
	// topic id.
	TopicId string `pulumi:"topicId"`
}

// The set of arguments for constructing a Export resource.
type ExportArgs struct {
	// log export format.
	Format pulumi.StringPtrInput
	// export start time.
	From pulumi.IntInput
	// export amount of log.
	LogCount pulumi.IntInput
	// log export time sorting. desc or asc.
	Order pulumi.StringPtrInput
	// export query rules.
	Query pulumi.StringInput
	// export end time.
	To pulumi.IntInput
	// topic id.
	TopicId pulumi.StringInput
}

func (ExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exportArgs)(nil)).Elem()
}

type ExportInput interface {
	pulumi.Input

	ToExportOutput() ExportOutput
	ToExportOutputWithContext(ctx context.Context) ExportOutput
}

func (*Export) ElementType() reflect.Type {
	return reflect.TypeOf((**Export)(nil)).Elem()
}

func (i *Export) ToExportOutput() ExportOutput {
	return i.ToExportOutputWithContext(context.Background())
}

func (i *Export) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportOutput)
}

// ExportArrayInput is an input type that accepts ExportArray and ExportArrayOutput values.
// You can construct a concrete instance of `ExportArrayInput` via:
//
//	ExportArray{ ExportArgs{...} }
type ExportArrayInput interface {
	pulumi.Input

	ToExportArrayOutput() ExportArrayOutput
	ToExportArrayOutputWithContext(context.Context) ExportArrayOutput
}

type ExportArray []ExportInput

func (ExportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Export)(nil)).Elem()
}

func (i ExportArray) ToExportArrayOutput() ExportArrayOutput {
	return i.ToExportArrayOutputWithContext(context.Background())
}

func (i ExportArray) ToExportArrayOutputWithContext(ctx context.Context) ExportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportArrayOutput)
}

// ExportMapInput is an input type that accepts ExportMap and ExportMapOutput values.
// You can construct a concrete instance of `ExportMapInput` via:
//
//	ExportMap{ "key": ExportArgs{...} }
type ExportMapInput interface {
	pulumi.Input

	ToExportMapOutput() ExportMapOutput
	ToExportMapOutputWithContext(context.Context) ExportMapOutput
}

type ExportMap map[string]ExportInput

func (ExportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Export)(nil)).Elem()
}

func (i ExportMap) ToExportMapOutput() ExportMapOutput {
	return i.ToExportMapOutputWithContext(context.Background())
}

func (i ExportMap) ToExportMapOutputWithContext(ctx context.Context) ExportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportMapOutput)
}

type ExportOutput struct{ *pulumi.OutputState }

func (ExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Export)(nil)).Elem()
}

func (o ExportOutput) ToExportOutput() ExportOutput {
	return o
}

func (o ExportOutput) ToExportOutputWithContext(ctx context.Context) ExportOutput {
	return o
}

// log export format.
func (o ExportOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Export) pulumi.StringPtrOutput { return v.Format }).(pulumi.StringPtrOutput)
}

// export start time.
func (o ExportOutput) From() pulumi.IntOutput {
	return o.ApplyT(func(v *Export) pulumi.IntOutput { return v.From }).(pulumi.IntOutput)
}

// export amount of log.
func (o ExportOutput) LogCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Export) pulumi.IntOutput { return v.LogCount }).(pulumi.IntOutput)
}

// log export time sorting. desc or asc.
func (o ExportOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Export) pulumi.StringPtrOutput { return v.Order }).(pulumi.StringPtrOutput)
}

// export query rules.
func (o ExportOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *Export) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

// export end time.
func (o ExportOutput) To() pulumi.IntOutput {
	return o.ApplyT(func(v *Export) pulumi.IntOutput { return v.To }).(pulumi.IntOutput)
}

// topic id.
func (o ExportOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *Export) pulumi.StringOutput { return v.TopicId }).(pulumi.StringOutput)
}

type ExportArrayOutput struct{ *pulumi.OutputState }

func (ExportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Export)(nil)).Elem()
}

func (o ExportArrayOutput) ToExportArrayOutput() ExportArrayOutput {
	return o
}

func (o ExportArrayOutput) ToExportArrayOutputWithContext(ctx context.Context) ExportArrayOutput {
	return o
}

func (o ExportArrayOutput) Index(i pulumi.IntInput) ExportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Export {
		return vs[0].([]*Export)[vs[1].(int)]
	}).(ExportOutput)
}

type ExportMapOutput struct{ *pulumi.OutputState }

func (ExportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Export)(nil)).Elem()
}

func (o ExportMapOutput) ToExportMapOutput() ExportMapOutput {
	return o
}

func (o ExportMapOutput) ToExportMapOutputWithContext(ctx context.Context) ExportMapOutput {
	return o
}

func (o ExportMapOutput) MapIndex(k pulumi.StringInput) ExportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Export {
		return vs[0].(map[string]*Export)[vs[1].(string)]
	}).(ExportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExportInput)(nil)).Elem(), &Export{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportArrayInput)(nil)).Elem(), ExportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportMapInput)(nil)).Elem(), ExportMap{})
	pulumi.RegisterOutputType(ExportOutput{})
	pulumi.RegisterOutputType(ExportArrayOutput{})
	pulumi.RegisterOutputType(ExportMapOutput{})
}
