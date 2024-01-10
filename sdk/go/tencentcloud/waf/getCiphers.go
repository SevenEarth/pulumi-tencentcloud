// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of waf ciphers
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Waf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Waf.GetCiphers(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCiphers(ctx *pulumi.Context, args *GetCiphersArgs, opts ...pulumi.InvokeOption) (*GetCiphersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCiphersResult
	err := ctx.Invoke("tencentcloud:Waf/getCiphers:getCiphers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCiphers.
type GetCiphersArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getCiphers.
type GetCiphersResult struct {
	// Encryption Suite InformationNote: This field may return null, indicating that a valid value cannot be obtained.
	Ciphers []GetCiphersCipher `pulumi:"ciphers"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetCiphersOutput(ctx *pulumi.Context, args GetCiphersOutputArgs, opts ...pulumi.InvokeOption) GetCiphersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCiphersResult, error) {
			args := v.(GetCiphersArgs)
			r, err := GetCiphers(ctx, &args, opts...)
			var s GetCiphersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCiphersResultOutput)
}

// A collection of arguments for invoking getCiphers.
type GetCiphersOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetCiphersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCiphersArgs)(nil)).Elem()
}

// A collection of values returned by getCiphers.
type GetCiphersResultOutput struct{ *pulumi.OutputState }

func (GetCiphersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCiphersResult)(nil)).Elem()
}

func (o GetCiphersResultOutput) ToGetCiphersResultOutput() GetCiphersResultOutput {
	return o
}

func (o GetCiphersResultOutput) ToGetCiphersResultOutputWithContext(ctx context.Context) GetCiphersResultOutput {
	return o
}

// Encryption Suite InformationNote: This field may return null, indicating that a valid value cannot be obtained.
func (o GetCiphersResultOutput) Ciphers() GetCiphersCipherArrayOutput {
	return o.ApplyT(func(v GetCiphersResult) []GetCiphersCipher { return v.Ciphers }).(GetCiphersCipherArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCiphersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCiphersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCiphersResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCiphersResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCiphersResultOutput{})
}
