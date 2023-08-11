// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a ses template.
//
// ## Example Usage
// ### Create a ses html template
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ses"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ses"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Ses.NewTemplate(ctx, "example", &Ses.TemplateArgs{
// 			TemplateContent: &ses.TemplateTemplateContentArgs{
// 				Html: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "<!DOCTYPE html>\n", "<html lang=\"en\">\n", "<head>\n", "  <meta charset=\"UTF-8\">\n", "  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n", "  <title>mail title</title>\n", "</head>\n", "<body>\n", "<div class=\"container\">\n", "  <h1>Welcome to our service! </h1>\n", "  <p>Dear user,</p>\n", "  <p>Thank you for using Tencent Cloud:</p>\n", "  <p><a href=\"https://cloud.tencent.com/document/product/1653\">https://cloud.tencent.com/document/product/1653</a></p>\n", "  <p>If you did not request this email, please ignore it. </p>\n", "  <p><strong>from the iac team</strong></p>\n", "</div>\n", "</body>\n", "</html>\n", "\n")),
// 			},
// 			TemplateName: pulumi.String("tf_example_ses_temp"),
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
// ses template can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ses/template:Template example template_id
// ```
type Template struct {
	pulumi.CustomResourceState

	// Sms Template Content.
	TemplateContent TemplateTemplateContentOutput `pulumi:"templateContent"`
	// smsTemplateName, which must be required.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TemplateContent == nil {
		return nil, errors.New("invalid value for required argument 'TemplateContent'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Template
	err := ctx.RegisterResource("tencentcloud:Ses/template:Template", name, args, &resource, opts...)
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
	err := ctx.ReadResource("tencentcloud:Ses/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// Sms Template Content.
	TemplateContent *TemplateTemplateContent `pulumi:"templateContent"`
	// smsTemplateName, which must be required.
	TemplateName *string `pulumi:"templateName"`
}

type TemplateState struct {
	// Sms Template Content.
	TemplateContent TemplateTemplateContentPtrInput
	// smsTemplateName, which must be required.
	TemplateName pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// Sms Template Content.
	TemplateContent TemplateTemplateContent `pulumi:"templateContent"`
	// smsTemplateName, which must be required.
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// Sms Template Content.
	TemplateContent TemplateTemplateContentInput
	// smsTemplateName, which must be required.
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
//          TemplateArray{ TemplateArgs{...} }
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
//          TemplateMap{ "key": TemplateArgs{...} }
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

// Sms Template Content.
func (o TemplateOutput) TemplateContent() TemplateTemplateContentOutput {
	return o.ApplyT(func(v *Template) TemplateTemplateContentOutput { return v.TemplateContent }).(TemplateTemplateContentOutput)
}

// smsTemplateName, which must be required.
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
