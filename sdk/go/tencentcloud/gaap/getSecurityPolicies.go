// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query security policies of GAAP proxy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Gaap"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooProxy, err := Gaap.NewProxy(ctx, "fooProxy", &Gaap.ProxyArgs{
// 			Bandwidth:        pulumi.Int(10),
// 			Concurrent:       pulumi.Int(2),
// 			AccessRegion:     pulumi.String("SouthChina"),
// 			RealserverRegion: pulumi.String("NorthChina"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		fooSecurityPolicy, err := Gaap.NewSecurityPolicy(ctx, "fooSecurityPolicy", &Gaap.SecurityPolicyArgs{
// 			ProxyId: fooProxy.ID(),
// 			Action:  pulumi.String("ACCEPT"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = Gaap.GetSecurityPoliciesOutput(ctx, gaap.GetSecurityPoliciesOutputArgs{
// 			Id: fooSecurityPolicy.ID(),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func GetSecurityPolicies(ctx *pulumi.Context, args *GetSecurityPoliciesArgs, opts ...pulumi.InvokeOption) (*GetSecurityPoliciesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecurityPoliciesResult
	err := ctx.Invoke("tencentcloud:Gaap/getSecurityPolicies:getSecurityPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityPolicies.
type GetSecurityPoliciesArgs struct {
	// ID of the security policy to be queried.
	Id string `pulumi:"id"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getSecurityPolicies.
type GetSecurityPoliciesResult struct {
	// Default policy.
	Action string `pulumi:"action"`
	Id     string `pulumi:"id"`
	// ID of the GAAP proxy.
	ProxyId          string  `pulumi:"proxyId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Status of the security policy.
	Status string `pulumi:"status"`
}

func GetSecurityPoliciesOutput(ctx *pulumi.Context, args GetSecurityPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetSecurityPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecurityPoliciesResult, error) {
			args := v.(GetSecurityPoliciesArgs)
			r, err := GetSecurityPolicies(ctx, &args, opts...)
			var s GetSecurityPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecurityPoliciesResultOutput)
}

// A collection of arguments for invoking getSecurityPolicies.
type GetSecurityPoliciesOutputArgs struct {
	// ID of the security policy to be queried.
	Id pulumi.StringInput `pulumi:"id"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetSecurityPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityPolicies.
type GetSecurityPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetSecurityPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityPoliciesResult)(nil)).Elem()
}

func (o GetSecurityPoliciesResultOutput) ToGetSecurityPoliciesResultOutput() GetSecurityPoliciesResultOutput {
	return o
}

func (o GetSecurityPoliciesResultOutput) ToGetSecurityPoliciesResultOutputWithContext(ctx context.Context) GetSecurityPoliciesResultOutput {
	return o
}

// Default policy.
func (o GetSecurityPoliciesResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityPoliciesResult) string { return v.Action }).(pulumi.StringOutput)
}

func (o GetSecurityPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the GAAP proxy.
func (o GetSecurityPoliciesResultOutput) ProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityPoliciesResult) string { return v.ProxyId }).(pulumi.StringOutput)
}

func (o GetSecurityPoliciesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityPoliciesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Status of the security policy.
func (o GetSecurityPoliciesResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityPoliciesResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecurityPoliciesResultOutput{})
}
