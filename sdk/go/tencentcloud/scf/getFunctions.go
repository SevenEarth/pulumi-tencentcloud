// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query SCF functions.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Scf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooFunction, err := Scf.NewFunction(ctx, "fooFunction", &Scf.FunctionArgs{
//				Handler:         pulumi.String("main.do_it"),
//				Runtime:         pulumi.String("Python3.6"),
//				CosBucketName:   pulumi.String("scf-code-1234567890"),
//				CosObjectName:   pulumi.String("code.zip"),
//				CosBucketRegion: pulumi.String("ap-guangzhou"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = Scf.GetFunctionsOutput(ctx, scf.GetFunctionsOutputArgs{
//				Name: fooFunction.Name,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetFunctions(ctx *pulumi.Context, args *GetFunctionsArgs, opts ...pulumi.InvokeOption) (*GetFunctionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetFunctionsResult
	err := ctx.Invoke("tencentcloud:Scf/getFunctions:getFunctions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFunctions.
type GetFunctionsArgs struct {
	// Description of the SCF function to be queried.
	Description *string `pulumi:"description"`
	// Name of the SCF function to be queried.
	Name *string `pulumi:"name"`
	// Namespace of the SCF function to be queried.
	Namespace *string `pulumi:"namespace"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Tags of the SCF function to be queried, can use up to 10 tags.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getFunctions.
type GetFunctionsResult struct {
	// Description of the SCF function.
	Description *string `pulumi:"description"`
	// An information list of functions. Each element contains the following attributes:
	Functions []GetFunctionsFunction `pulumi:"functions"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the SCF function trigger.
	Name *string `pulumi:"name"`
	// Namespace of the SCF function.
	Namespace        *string `pulumi:"namespace"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Tags of the SCF function.
	Tags map[string]interface{} `pulumi:"tags"`
}

func GetFunctionsOutput(ctx *pulumi.Context, args GetFunctionsOutputArgs, opts ...pulumi.InvokeOption) GetFunctionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFunctionsResult, error) {
			args := v.(GetFunctionsArgs)
			r, err := GetFunctions(ctx, &args, opts...)
			var s GetFunctionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFunctionsResultOutput)
}

// A collection of arguments for invoking getFunctions.
type GetFunctionsOutputArgs struct {
	// Description of the SCF function to be queried.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Name of the SCF function to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Namespace of the SCF function to be queried.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Tags of the SCF function to be queried, can use up to 10 tags.
	Tags pulumi.MapInput `pulumi:"tags"`
}

func (GetFunctionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFunctionsArgs)(nil)).Elem()
}

// A collection of values returned by getFunctions.
type GetFunctionsResultOutput struct{ *pulumi.OutputState }

func (GetFunctionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFunctionsResult)(nil)).Elem()
}

func (o GetFunctionsResultOutput) ToGetFunctionsResultOutput() GetFunctionsResultOutput {
	return o
}

func (o GetFunctionsResultOutput) ToGetFunctionsResultOutputWithContext(ctx context.Context) GetFunctionsResultOutput {
	return o
}

// Description of the SCF function.
func (o GetFunctionsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFunctionsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// An information list of functions. Each element contains the following attributes:
func (o GetFunctionsResultOutput) Functions() GetFunctionsFunctionArrayOutput {
	return o.ApplyT(func(v GetFunctionsResult) []GetFunctionsFunction { return v.Functions }).(GetFunctionsFunctionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFunctionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFunctionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the SCF function trigger.
func (o GetFunctionsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFunctionsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Namespace of the SCF function.
func (o GetFunctionsResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFunctionsResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetFunctionsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFunctionsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Tags of the SCF function.
func (o GetFunctionsResultOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetFunctionsResult) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFunctionsResultOutput{})
}
