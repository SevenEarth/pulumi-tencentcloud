// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package css

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of css domains
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Css"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Css"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Css.GetDomains(ctx, &css.GetDomainsArgs{
// 			DomainType:  pulumi.IntRef(0),
// 			IsDelayLive: pulumi.IntRef(0),
// 			PlayType:    pulumi.IntRef(1),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDomains(ctx *pulumi.Context, args *GetDomainsArgs, opts ...pulumi.InvokeOption) (*GetDomainsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDomainsResult
	err := ctx.Invoke("tencentcloud:Css/getDomains:getDomains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomains.
type GetDomainsArgs struct {
	// domain name prefix.
	DomainPrefix *string `pulumi:"domainPrefix"`
	// domain name status filter. 0-disable, 1-enable.
	DomainStatus *int `pulumi:"domainStatus"`
	// Domain name type filtering. 0-push, 1-play.
	DomainType *int `pulumi:"domainType"`
	// 0 normal live broadcast 1 slow live broadcast default 0.
	IsDelayLive *int `pulumi:"isDelayLive"`
	// Playing area, this parameter is meaningful only when DomainType=1. 1: Domestic.2: Global.3: Overseas.
	PlayType *int `pulumi:"playType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDomains.
type GetDomainsResult struct {
	// A list of domain name details.
	DomainLists  []GetDomainsDomainList `pulumi:"domainLists"`
	DomainPrefix *string                `pulumi:"domainPrefix"`
	DomainStatus *int                   `pulumi:"domainStatus"`
	DomainType   *int                   `pulumi:"domainType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether to slow live broadcast: 0: normal live broadcast. 1: Slow live broadcast.
	IsDelayLive *int `pulumi:"isDelayLive"`
	// Playing area, this parameter is meaningful only when Type=1. 1: Domestic. 2: Global. 3: Overseas.
	PlayType         *int    `pulumi:"playType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetDomainsOutput(ctx *pulumi.Context, args GetDomainsOutputArgs, opts ...pulumi.InvokeOption) GetDomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDomainsResult, error) {
			args := v.(GetDomainsArgs)
			r, err := GetDomains(ctx, &args, opts...)
			var s GetDomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDomainsResultOutput)
}

// A collection of arguments for invoking getDomains.
type GetDomainsOutputArgs struct {
	// domain name prefix.
	DomainPrefix pulumi.StringPtrInput `pulumi:"domainPrefix"`
	// domain name status filter. 0-disable, 1-enable.
	DomainStatus pulumi.IntPtrInput `pulumi:"domainStatus"`
	// Domain name type filtering. 0-push, 1-play.
	DomainType pulumi.IntPtrInput `pulumi:"domainType"`
	// 0 normal live broadcast 1 slow live broadcast default 0.
	IsDelayLive pulumi.IntPtrInput `pulumi:"isDelayLive"`
	// Playing area, this parameter is meaningful only when DomainType=1. 1: Domestic.2: Global.3: Overseas.
	PlayType pulumi.IntPtrInput `pulumi:"playType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsArgs)(nil)).Elem()
}

// A collection of values returned by getDomains.
type GetDomainsResultOutput struct{ *pulumi.OutputState }

func (GetDomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDomainsResult)(nil)).Elem()
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutput() GetDomainsResultOutput {
	return o
}

func (o GetDomainsResultOutput) ToGetDomainsResultOutputWithContext(ctx context.Context) GetDomainsResultOutput {
	return o
}

// A list of domain name details.
func (o GetDomainsResultOutput) DomainLists() GetDomainsDomainListArrayOutput {
	return o.ApplyT(func(v GetDomainsResult) []GetDomainsDomainList { return v.DomainLists }).(GetDomainsDomainListArrayOutput)
}

func (o GetDomainsResultOutput) DomainPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.DomainPrefix }).(pulumi.StringPtrOutput)
}

func (o GetDomainsResultOutput) DomainStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *int { return v.DomainStatus }).(pulumi.IntPtrOutput)
}

func (o GetDomainsResultOutput) DomainType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *int { return v.DomainType }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether to slow live broadcast: 0: normal live broadcast. 1: Slow live broadcast.
func (o GetDomainsResultOutput) IsDelayLive() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *int { return v.IsDelayLive }).(pulumi.IntPtrOutput)
}

// Playing area, this parameter is meaningful only when Type=1. 1: Domestic. 2: Global. 3: Overseas.
func (o GetDomainsResultOutput) PlayType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *int { return v.PlayType }).(pulumi.IntPtrOutput)
}

func (o GetDomainsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDomainsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDomainsResultOutput{})
}