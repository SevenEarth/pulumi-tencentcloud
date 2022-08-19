// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of CAM SAML providers
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cam"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cam.GetSamlProviders(ctx, &cam.GetSamlProvidersArgs{
// 			Name: pulumi.StringRef("cam-test-provider"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSamlProviders(ctx *pulumi.Context, args *GetSamlProvidersArgs, opts ...pulumi.InvokeOption) (*GetSamlProvidersResult, error) {
	var rv GetSamlProvidersResult
	err := ctx.Invoke("tencentcloud:Cam/getSamlProviders:getSamlProviders", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSamlProviders.
type GetSamlProvidersArgs struct {
	// The description of the CAM SAML provider.
	Description *string `pulumi:"description"`
	// Name of the CAM SAML provider to be queried.
	Name *string `pulumi:"name"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getSamlProviders.
type GetSamlProvidersResult struct {
	// Description of CAM SAML provider.
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of CAM SAML provider.
	Name *string `pulumi:"name"`
	// A list of CAM SAML providers. Each element contains the following attributes:
	ProviderLists    []GetSamlProvidersProviderList `pulumi:"providerLists"`
	ResultOutputFile *string                        `pulumi:"resultOutputFile"`
}

func GetSamlProvidersOutput(ctx *pulumi.Context, args GetSamlProvidersOutputArgs, opts ...pulumi.InvokeOption) GetSamlProvidersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSamlProvidersResult, error) {
			args := v.(GetSamlProvidersArgs)
			r, err := GetSamlProviders(ctx, &args, opts...)
			var s GetSamlProvidersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSamlProvidersResultOutput)
}

// A collection of arguments for invoking getSamlProviders.
type GetSamlProvidersOutputArgs struct {
	// The description of the CAM SAML provider.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Name of the CAM SAML provider to be queried.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetSamlProvidersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSamlProvidersArgs)(nil)).Elem()
}

// A collection of values returned by getSamlProviders.
type GetSamlProvidersResultOutput struct{ *pulumi.OutputState }

func (GetSamlProvidersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSamlProvidersResult)(nil)).Elem()
}

func (o GetSamlProvidersResultOutput) ToGetSamlProvidersResultOutput() GetSamlProvidersResultOutput {
	return o
}

func (o GetSamlProvidersResultOutput) ToGetSamlProvidersResultOutputWithContext(ctx context.Context) GetSamlProvidersResultOutput {
	return o
}

// Description of CAM SAML provider.
func (o GetSamlProvidersResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSamlProvidersResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSamlProvidersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSamlProvidersResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of CAM SAML provider.
func (o GetSamlProvidersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSamlProvidersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of CAM SAML providers. Each element contains the following attributes:
func (o GetSamlProvidersResultOutput) ProviderLists() GetSamlProvidersProviderListArrayOutput {
	return o.ApplyT(func(v GetSamlProvidersResult) []GetSamlProvidersProviderList { return v.ProviderLists }).(GetSamlProvidersProviderListArrayOutput)
}

func (o GetSamlProvidersResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSamlProvidersResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSamlProvidersResultOutput{})
}
