// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package protocol

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of protocol templates.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Protocol"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Protocol"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Protocol.GetTemplates(ctx, &protocol.GetTemplatesArgs{
//				Name: pulumi.StringRef("test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetTemplates(ctx *pulumi.Context, args *GetTemplatesArgs, opts ...pulumi.InvokeOption) (*GetTemplatesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetTemplatesResult
	err := ctx.Invoke("tencentcloud:Protocol/getTemplates:getTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplates.
type GetTemplatesArgs struct {
	// ID of the protocol template to query.
	Id *string `pulumi:"id"`
	// Name of the protocol template to query.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getTemplates.
type GetTemplatesResult struct {
	// ID of the protocol template.
	Id *string `pulumi:"id"`
	// Name of protocol template.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Information list of the dedicated protocol templates.
	TemplateLists []GetTemplatesTemplateList `pulumi:"templateLists"`
}

func GetTemplatesOutput(ctx *pulumi.Context, args GetTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTemplatesResult, error) {
			args := v.(GetTemplatesArgs)
			r, err := GetTemplates(ctx, &args, opts...)
			var s GetTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTemplatesResultOutput)
}

// A collection of arguments for invoking getTemplates.
type GetTemplatesOutputArgs struct {
	// ID of the protocol template to query.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the protocol template to query.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getTemplates.
type GetTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplatesResult)(nil)).Elem()
}

func (o GetTemplatesResultOutput) ToGetTemplatesResultOutput() GetTemplatesResultOutput {
	return o
}

func (o GetTemplatesResultOutput) ToGetTemplatesResultOutputWithContext(ctx context.Context) GetTemplatesResultOutput {
	return o
}

// ID of the protocol template.
func (o GetTemplatesResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of protocol template.
func (o GetTemplatesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetTemplatesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTemplatesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Information list of the dedicated protocol templates.
func (o GetTemplatesResultOutput) TemplateLists() GetTemplatesTemplateListArrayOutput {
	return o.ApplyT(func(v GetTemplatesResult) []GetTemplatesTemplateList { return v.TemplateLists }).(GetTemplatesTemplateListArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTemplatesResultOutput{})
}
