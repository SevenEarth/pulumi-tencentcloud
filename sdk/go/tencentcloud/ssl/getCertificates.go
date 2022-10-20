// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssl

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query SSL certificate.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssl"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ssl.GetCertificates(ctx, &ssl.GetCertificatesArgs{
//				Name: pulumi.StringRef("certificate"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCertificates(ctx *pulumi.Context, args *GetCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCertificatesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetCertificatesResult
	err := ctx.Invoke("tencentcloud:Ssl/getCertificates:getCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesArgs struct {
	// ID of the SSL certificate to be queried.
	Id *string `pulumi:"id"`
	// Name of the SSL certificate to be queried.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Type of the SSL certificate to be queried. Available values includes: `CA` and `SVR`.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getCertificates.
type GetCertificatesResult struct {
	// An information list of certificate. Each element contains the following attributes:
	Certificates []GetCertificatesCertificate `pulumi:"certificates"`
	// ID of the SSL certificate.
	Id *string `pulumi:"id"`
	// Name of the SSL certificate.
	Name             *string `pulumi:"name"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Type of the SSL certificate.
	Type *string `pulumi:"type"`
}

func GetCertificatesOutput(ctx *pulumi.Context, args GetCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCertificatesResult, error) {
			args := v.(GetCertificatesArgs)
			r, err := GetCertificates(ctx, &args, opts...)
			var s GetCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCertificatesResultOutput)
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesOutputArgs struct {
	// ID of the SSL certificate to be queried.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the SSL certificate to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Type of the SSL certificate to be queried. Available values includes: `CA` and `SVR`.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GetCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getCertificates.
type GetCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesResult)(nil)).Elem()
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutput() GetCertificatesResultOutput {
	return o
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutputWithContext(ctx context.Context) GetCertificatesResultOutput {
	return o
}

// An information list of certificate. Each element contains the following attributes:
func (o GetCertificatesResultOutput) Certificates() GetCertificatesCertificateArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []GetCertificatesCertificate { return v.Certificates }).(GetCertificatesCertificateArrayOutput)
}

// ID of the SSL certificate.
func (o GetCertificatesResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the SSL certificate.
func (o GetCertificatesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetCertificatesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Type of the SSL certificate.
func (o GetCertificatesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCertificatesResultOutput{})
}
