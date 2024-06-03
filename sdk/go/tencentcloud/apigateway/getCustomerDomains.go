// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query API gateway domain list.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			foo, err := ApiGateway.NewCustomDomain(ctx, "foo", &ApiGateway.CustomDomainArgs{
//				ServiceId:        pulumi.String("service-ohxqslqe"),
//				SubDomain:        pulumi.String("tic-test.dnsv1.com"),
//				Protocol:         pulumi.String("http"),
//				NetType:          pulumi.String("OUTER"),
//				IsDefaultMapping: pulumi.Bool(false),
//				DefaultDomain:    pulumi.String("service-ohxqslqe-1259649581.gz.apigw.tencentcs.com"),
//				PathMappings: pulumi.StringArray{
//					pulumi.String("/good#test"),
//					pulumi.String("/root#release"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = ApiGateway.GetCustomerDomainsOutput(ctx, apigateway.GetCustomerDomainsOutputArgs{
//				ServiceId: foo.ServiceId,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetCustomerDomains(ctx *pulumi.Context, args *GetCustomerDomainsArgs, opts ...pulumi.InvokeOption) (*GetCustomerDomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCustomerDomainsResult
	err := ctx.Invoke("tencentcloud:ApiGateway/getCustomerDomains:getCustomerDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomerDomains.
type GetCustomerDomainsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The service ID.
	ServiceId string `pulumi:"serviceId"`
}

// A collection of values returned by getCustomerDomains.
type GetCustomerDomainsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Service custom domain name list.
	Lists            []GetCustomerDomainsList `pulumi:"lists"`
	ResultOutputFile *string                  `pulumi:"resultOutputFile"`
	ServiceId        string                   `pulumi:"serviceId"`
}

func GetCustomerDomainsOutput(ctx *pulumi.Context, args GetCustomerDomainsOutputArgs, opts ...pulumi.InvokeOption) GetCustomerDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCustomerDomainsResult, error) {
			args := v.(GetCustomerDomainsArgs)
			r, err := GetCustomerDomains(ctx, &args, opts...)
			var s GetCustomerDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCustomerDomainsResultOutput)
}

// A collection of arguments for invoking getCustomerDomains.
type GetCustomerDomainsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The service ID.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (GetCustomerDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getCustomerDomains.
type GetCustomerDomainsResultOutput struct{ *pulumi.OutputState }

func (GetCustomerDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCustomerDomainsResult)(nil)).Elem()
}

func (o GetCustomerDomainsResultOutput) ToGetCustomerDomainsResultOutput() GetCustomerDomainsResultOutput {
	return o
}

func (o GetCustomerDomainsResultOutput) ToGetCustomerDomainsResultOutputWithContext(ctx context.Context) GetCustomerDomainsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCustomerDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Service custom domain name list.
func (o GetCustomerDomainsResultOutput) Lists() GetCustomerDomainsListArrayOutput {
	return o.ApplyT(func(v GetCustomerDomainsResult) []GetCustomerDomainsList { return v.Lists }).(GetCustomerDomainsListArrayOutput)
}

func (o GetCustomerDomainsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCustomerDomainsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetCustomerDomainsResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCustomerDomainsResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCustomerDomainsResultOutput{})
}
