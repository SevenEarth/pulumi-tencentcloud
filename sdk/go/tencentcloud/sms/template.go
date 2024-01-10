// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a sms template
//
// ## Example Usage
// ### Create a sms template
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sms"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Sms.NewTemplate(ctx, "template", &Sms.TemplateArgs{
//				International:   pulumi.Int(0),
//				Remark:          pulumi.String("terraform example"),
//				SmsType:         pulumi.Int(0),
//				TemplateContent: pulumi.String("example for sms template"),
//				TemplateName:    pulumi.String("tf_example_sms_template"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Template struct {
	pulumi.CustomResourceState

	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International pulumi.IntOutput `pulumi:"international"`
	// Template remarks, such as reason for application and use case.
	Remark pulumi.StringOutput `pulumi:"remark"`
	// SMS type. 0: regular SMS, 1: marketing SMS.
	SmsType pulumi.IntOutput `pulumi:"smsType"`
	// Message Template Content.
	TemplateContent pulumi.StringOutput `pulumi:"templateContent"`
	// Message Template name, which must be unique.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.International == nil {
		return nil, errors.New("invalid value for required argument 'International'")
	}
	if args.Remark == nil {
		return nil, errors.New("invalid value for required argument 'Remark'")
	}
	if args.SmsType == nil {
		return nil, errors.New("invalid value for required argument 'SmsType'")
	}
	if args.TemplateContent == nil {
		return nil, errors.New("invalid value for required argument 'TemplateContent'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("tencentcloud:Sms/template:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("tencentcloud:Sms/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International *int `pulumi:"international"`
	// Template remarks, such as reason for application and use case.
	Remark *string `pulumi:"remark"`
	// SMS type. 0: regular SMS, 1: marketing SMS.
	SmsType *int `pulumi:"smsType"`
	// Message Template Content.
	TemplateContent *string `pulumi:"templateContent"`
	// Message Template name, which must be unique.
	TemplateName *string `pulumi:"templateName"`
}

type TemplateState struct {
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International pulumi.IntPtrInput
	// Template remarks, such as reason for application and use case.
	Remark pulumi.StringPtrInput
	// SMS type. 0: regular SMS, 1: marketing SMS.
	SmsType pulumi.IntPtrInput
	// Message Template Content.
	TemplateContent pulumi.StringPtrInput
	// Message Template name, which must be unique.
	TemplateName pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International int `pulumi:"international"`
	// Template remarks, such as reason for application and use case.
	Remark string `pulumi:"remark"`
	// SMS type. 0: regular SMS, 1: marketing SMS.
	SmsType int `pulumi:"smsType"`
	// Message Template Content.
	TemplateContent string `pulumi:"templateContent"`
	// Message Template name, which must be unique.
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International pulumi.IntInput
	// Template remarks, such as reason for application and use case.
	Remark pulumi.StringInput
	// SMS type. 0: regular SMS, 1: marketing SMS.
	SmsType pulumi.IntInput
	// Message Template Content.
	TemplateContent pulumi.StringInput
	// Message Template name, which must be unique.
	TemplateName pulumi.StringInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

// TemplateArrayInput is an input type that accepts TemplateArray and TemplateArrayOutput values.
// You can construct a concrete instance of `TemplateArrayInput` via:
//
//	TemplateArray{ TemplateArgs{...} }
type TemplateArrayInput interface {
	pulumi.Input

	ToTemplateArrayOutput() TemplateArrayOutput
	ToTemplateArrayOutputWithContext(context.Context) TemplateArrayOutput
}

type TemplateArray []TemplateInput

func (TemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (i TemplateArray) ToTemplateArrayOutput() TemplateArrayOutput {
	return i.ToTemplateArrayOutputWithContext(context.Background())
}

func (i TemplateArray) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArrayOutput)
}

// TemplateMapInput is an input type that accepts TemplateMap and TemplateMapOutput values.
// You can construct a concrete instance of `TemplateMapInput` via:
//
//	TemplateMap{ "key": TemplateArgs{...} }
type TemplateMapInput interface {
	pulumi.Input

	ToTemplateMapOutput() TemplateMapOutput
	ToTemplateMapOutputWithContext(context.Context) TemplateMapOutput
}

type TemplateMap map[string]TemplateInput

func (TemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (i TemplateMap) ToTemplateMapOutput() TemplateMapOutput {
	return i.ToTemplateMapOutputWithContext(context.Background())
}

func (i TemplateMap) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateMapOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
func (o TemplateOutput) International() pulumi.IntOutput {
	return o.ApplyT(func(v *Template) pulumi.IntOutput { return v.International }).(pulumi.IntOutput)
}

// Template remarks, such as reason for application and use case.
func (o TemplateOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.Remark }).(pulumi.StringOutput)
}

// SMS type. 0: regular SMS, 1: marketing SMS.
func (o TemplateOutput) SmsType() pulumi.IntOutput {
	return o.ApplyT(func(v *Template) pulumi.IntOutput { return v.SmsType }).(pulumi.IntOutput)
}

// Message Template Content.
func (o TemplateOutput) TemplateContent() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateContent }).(pulumi.StringOutput)
}

// Message Template name, which must be unique.
func (o TemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *Template) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

type TemplateArrayOutput struct{ *pulumi.OutputState }

func (TemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Template)(nil)).Elem()
}

func (o TemplateArrayOutput) ToTemplateArrayOutput() TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) Index(i pulumi.IntInput) TemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Template {
		return vs[0].([]*Template)[vs[1].(int)]
	}).(TemplateOutput)
}

type TemplateMapOutput struct{ *pulumi.OutputState }

func (TemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Template)(nil)).Elem()
}

func (o TemplateMapOutput) ToTemplateMapOutput() TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) MapIndex(k pulumi.StringInput) TemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Template {
		return vs[0].(map[string]*Template)[vs[1].(string)]
	}).(TemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateInput)(nil)).Elem(), &Template{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateArrayInput)(nil)).Elem(), TemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TemplateMapInput)(nil)).Elem(), TemplateMap{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplateArrayOutput{})
	pulumi.RegisterOutputType(TemplateMapOutput{})
}
