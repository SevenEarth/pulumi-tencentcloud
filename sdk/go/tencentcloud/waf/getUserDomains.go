// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of waf userDomains
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
//			_, err := Waf.GetUserDomains(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetUserDomains(ctx *pulumi.Context, args *GetUserDomainsArgs, opts ...pulumi.InvokeOption) (*GetUserDomainsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetUserDomainsResult
	err := ctx.Invoke("tencentcloud:Waf/getUserDomains:getUserDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUserDomains.
type GetUserDomainsArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getUserDomains.
type GetUserDomainsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Domain infos.
	UsersInfos []GetUserDomainsUsersInfo `pulumi:"usersInfos"`
}

func GetUserDomainsOutput(ctx *pulumi.Context, args GetUserDomainsOutputArgs, opts ...pulumi.InvokeOption) GetUserDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserDomainsResult, error) {
			args := v.(GetUserDomainsArgs)
			r, err := GetUserDomains(ctx, &args, opts...)
			var s GetUserDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUserDomainsResultOutput)
}

// A collection of arguments for invoking getUserDomains.
type GetUserDomainsOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetUserDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getUserDomains.
type GetUserDomainsResultOutput struct{ *pulumi.OutputState }

func (GetUserDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserDomainsResult)(nil)).Elem()
}

func (o GetUserDomainsResultOutput) ToGetUserDomainsResultOutput() GetUserDomainsResultOutput {
	return o
}

func (o GetUserDomainsResultOutput) ToGetUserDomainsResultOutputWithContext(ctx context.Context) GetUserDomainsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUserDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUserDomainsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserDomainsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Domain infos.
func (o GetUserDomainsResultOutput) UsersInfos() GetUserDomainsUsersInfoArrayOutput {
	return o.ApplyT(func(v GetUserDomainsResult) []GetUserDomainsUsersInfo { return v.UsersInfos }).(GetUserDomainsUsersInfoArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserDomainsResultOutput{})
}
