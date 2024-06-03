// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of rum staticUrl
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Rum"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Rum.GetStaticUrl(ctx, &rum.GetStaticUrlArgs{
//				EndTime:   1625454840,
//				ProjectId: 1,
//				StartTime: 1625444040,
//				Type:      "pagepv",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetStaticUrl(ctx *pulumi.Context, args *GetStaticUrlArgs, opts ...pulumi.InvokeOption) (*GetStaticUrlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStaticUrlResult
	err := ctx.Invoke("tencentcloud:Rum/getStaticUrl:getStaticUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStaticUrl.
type GetStaticUrlArgs struct {
	// The region where the data reporting takes place.
	Area *string `pulumi:"area"`
	// The mobile phone brand used for data reporting.
	Brand *string `pulumi:"brand"`
	// The browser type used for data reporting.
	Browser *string `pulumi:"browser"`
	// The method used for calculating the elapsed time `50`: 50th percentile, `75`: 75th percentile., `90`: 90th percentile., `95`: 95th percentile., `99`: 99th percentile., `99.5`: 99.5th percentile., `avg`: Mean.
	CostType *string `pulumi:"costType"`
	// The device used for data reporting.
	Device *string `pulumi:"device"`
	// End time but is represented using a timestamp in seconds.
	EndTime int `pulumi:"endTime"`
	// The browser engine used for data reporting.
	Engine *string `pulumi:"engine"`
	// The code environment where the data reporting takes place.(`production`: production env, `development`: development env, `gray`: gray env, `pre`: pre env, `daily`: daily env, `local`: local env, `others`: others env).
	Env *string `pulumi:"env"`
	// First Expansion parameter.
	ExtFirst *string `pulumi:"extFirst"`
	// Second Expansion parameter.
	ExtSecond *string `pulumi:"extSecond"`
	// Third Expansion parameter.
	ExtThird *string `pulumi:"extThird"`
	// The source page of the data reporting.
	From *string `pulumi:"from"`
	// Whether it is non-China region.`1`: yes; `0`: no.
	IsAbroad *string `pulumi:"isAbroad"`
	// The internet service provider used for data reporting.
	Isp *string `pulumi:"isp"`
	// Log level for data reporting(`1`: whitelist, `2`: normal, `4`: error, `8`: promise error, `16`: ajax request error, `32`: js resource load error, `64`: image resource load error, `128`: css resource load error, `256`: console.error, `512`: video resource load error, `1024`: request retcode error, `2048`: sdk self monitor error, `4096`: pv log, `8192`: event log).
	Level *string `pulumi:"level"`
	// The network type used for data reporting.(`1`: Wifi, `2`: 2G, `3`: 3G, `4`: 4G, `5`: 5G, `6`: 6G, `100`: Unknown).
	NetType *string `pulumi:"netType"`
	// The operating system used for data reporting.
	Os *string `pulumi:"os"`
	// The platform where the data reporting takes place.(`1`: Android, `2`: IOS, `3`: Windows, `4`: Mac, `5`: Linux, `100`: Other).
	Platform *string `pulumi:"platform"`
	// Project ID.
	ProjectId int `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Start time but is represented using a timestamp in seconds.
	StartTime int `pulumi:"startTime"`
	// Query Data Type. `pagepv`: CostType query by pagepv, `nettype`: CostType group by nettype, `version`: CostType group by version, `platform`: CostType group by platform, `isp`: CostType group by isp, `region`: CostType group by region, `device`: CostType group by device, `browser`: CostType group by browser, `ext1`: CostType group by ext1, `ext2`: CostType group by ext2, `ext3`: CostType group by ext3, `ret`: CostType group by ret, `status`: CostType group by status, `from`: CostType group by from, `url`: CostType group by url, `env`: CostType group by env.
	Type string `pulumi:"type"`
	// The URL Key where the data reporting takes place.
	Url *string `pulumi:"url"`
	// The SDK version used for data reporting.
	VersionNum *string `pulumi:"versionNum"`
}

// A collection of values returned by getStaticUrl.
type GetStaticUrlResult struct {
	Area      *string `pulumi:"area"`
	Brand     *string `pulumi:"brand"`
	Browser   *string `pulumi:"browser"`
	CostType  *string `pulumi:"costType"`
	Device    *string `pulumi:"device"`
	EndTime   int     `pulumi:"endTime"`
	Engine    *string `pulumi:"engine"`
	Env       *string `pulumi:"env"`
	ExtFirst  *string `pulumi:"extFirst"`
	ExtSecond *string `pulumi:"extSecond"`
	ExtThird  *string `pulumi:"extThird"`
	From      *string `pulumi:"from"`
	// The provider-assigned unique ID for this managed resource.
	Id        string  `pulumi:"id"`
	IsAbroad  *string `pulumi:"isAbroad"`
	Isp       *string `pulumi:"isp"`
	Level     *string `pulumi:"level"`
	NetType   *string `pulumi:"netType"`
	Os        *string `pulumi:"os"`
	Platform  *string `pulumi:"platform"`
	ProjectId int     `pulumi:"projectId"`
	// Return value.
	Result           string  `pulumi:"result"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	StartTime        int     `pulumi:"startTime"`
	Type             string  `pulumi:"type"`
	Url              *string `pulumi:"url"`
	VersionNum       *string `pulumi:"versionNum"`
}

func GetStaticUrlOutput(ctx *pulumi.Context, args GetStaticUrlOutputArgs, opts ...pulumi.InvokeOption) GetStaticUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStaticUrlResult, error) {
			args := v.(GetStaticUrlArgs)
			r, err := GetStaticUrl(ctx, &args, opts...)
			var s GetStaticUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStaticUrlResultOutput)
}

// A collection of arguments for invoking getStaticUrl.
type GetStaticUrlOutputArgs struct {
	// The region where the data reporting takes place.
	Area pulumi.StringPtrInput `pulumi:"area"`
	// The mobile phone brand used for data reporting.
	Brand pulumi.StringPtrInput `pulumi:"brand"`
	// The browser type used for data reporting.
	Browser pulumi.StringPtrInput `pulumi:"browser"`
	// The method used for calculating the elapsed time `50`: 50th percentile, `75`: 75th percentile., `90`: 90th percentile., `95`: 95th percentile., `99`: 99th percentile., `99.5`: 99.5th percentile., `avg`: Mean.
	CostType pulumi.StringPtrInput `pulumi:"costType"`
	// The device used for data reporting.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// End time but is represented using a timestamp in seconds.
	EndTime pulumi.IntInput `pulumi:"endTime"`
	// The browser engine used for data reporting.
	Engine pulumi.StringPtrInput `pulumi:"engine"`
	// The code environment where the data reporting takes place.(`production`: production env, `development`: development env, `gray`: gray env, `pre`: pre env, `daily`: daily env, `local`: local env, `others`: others env).
	Env pulumi.StringPtrInput `pulumi:"env"`
	// First Expansion parameter.
	ExtFirst pulumi.StringPtrInput `pulumi:"extFirst"`
	// Second Expansion parameter.
	ExtSecond pulumi.StringPtrInput `pulumi:"extSecond"`
	// Third Expansion parameter.
	ExtThird pulumi.StringPtrInput `pulumi:"extThird"`
	// The source page of the data reporting.
	From pulumi.StringPtrInput `pulumi:"from"`
	// Whether it is non-China region.`1`: yes; `0`: no.
	IsAbroad pulumi.StringPtrInput `pulumi:"isAbroad"`
	// The internet service provider used for data reporting.
	Isp pulumi.StringPtrInput `pulumi:"isp"`
	// Log level for data reporting(`1`: whitelist, `2`: normal, `4`: error, `8`: promise error, `16`: ajax request error, `32`: js resource load error, `64`: image resource load error, `128`: css resource load error, `256`: console.error, `512`: video resource load error, `1024`: request retcode error, `2048`: sdk self monitor error, `4096`: pv log, `8192`: event log).
	Level pulumi.StringPtrInput `pulumi:"level"`
	// The network type used for data reporting.(`1`: Wifi, `2`: 2G, `3`: 3G, `4`: 4G, `5`: 5G, `6`: 6G, `100`: Unknown).
	NetType pulumi.StringPtrInput `pulumi:"netType"`
	// The operating system used for data reporting.
	Os pulumi.StringPtrInput `pulumi:"os"`
	// The platform where the data reporting takes place.(`1`: Android, `2`: IOS, `3`: Windows, `4`: Mac, `5`: Linux, `100`: Other).
	Platform pulumi.StringPtrInput `pulumi:"platform"`
	// Project ID.
	ProjectId pulumi.IntInput `pulumi:"projectId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Start time but is represented using a timestamp in seconds.
	StartTime pulumi.IntInput `pulumi:"startTime"`
	// Query Data Type. `pagepv`: CostType query by pagepv, `nettype`: CostType group by nettype, `version`: CostType group by version, `platform`: CostType group by platform, `isp`: CostType group by isp, `region`: CostType group by region, `device`: CostType group by device, `browser`: CostType group by browser, `ext1`: CostType group by ext1, `ext2`: CostType group by ext2, `ext3`: CostType group by ext3, `ret`: CostType group by ret, `status`: CostType group by status, `from`: CostType group by from, `url`: CostType group by url, `env`: CostType group by env.
	Type pulumi.StringInput `pulumi:"type"`
	// The URL Key where the data reporting takes place.
	Url pulumi.StringPtrInput `pulumi:"url"`
	// The SDK version used for data reporting.
	VersionNum pulumi.StringPtrInput `pulumi:"versionNum"`
}

func (GetStaticUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticUrlArgs)(nil)).Elem()
}

// A collection of values returned by getStaticUrl.
type GetStaticUrlResultOutput struct{ *pulumi.OutputState }

func (GetStaticUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticUrlResult)(nil)).Elem()
}

func (o GetStaticUrlResultOutput) ToGetStaticUrlResultOutput() GetStaticUrlResultOutput {
	return o
}

func (o GetStaticUrlResultOutput) ToGetStaticUrlResultOutputWithContext(ctx context.Context) GetStaticUrlResultOutput {
	return o
}

func (o GetStaticUrlResultOutput) Area() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Area }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Brand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Brand }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Browser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Browser }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) CostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.CostType }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) EndTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticUrlResult) int { return v.EndTime }).(pulumi.IntOutput)
}

func (o GetStaticUrlResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Env() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Env }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) ExtFirst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.ExtFirst }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) ExtSecond() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.ExtSecond }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) ExtThird() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.ExtThird }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.From }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStaticUrlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticUrlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStaticUrlResultOutput) IsAbroad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.IsAbroad }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) NetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.NetType }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Os }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Platform }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticUrlResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

// Return value.
func (o GetStaticUrlResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticUrlResult) string { return v.Result }).(pulumi.StringOutput)
}

func (o GetStaticUrlResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) StartTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticUrlResult) int { return v.StartTime }).(pulumi.IntOutput)
}

func (o GetStaticUrlResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticUrlResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetStaticUrlResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o GetStaticUrlResultOutput) VersionNum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticUrlResult) *string { return v.VersionNum }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStaticUrlResultOutput{})
}
