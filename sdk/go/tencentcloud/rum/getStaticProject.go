// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rum

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of rum staticProject
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
//			_, err := Rum.GetStaticProject(ctx, &rum.GetStaticProjectArgs{
//				EndTime:   1625454840,
//				ProjectId: 1,
//				StartTime: 1625444040,
//				Type:      "allcount",
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
func GetStaticProject(ctx *pulumi.Context, args *GetStaticProjectArgs, opts ...pulumi.InvokeOption) (*GetStaticProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetStaticProjectResult
	err := ctx.Invoke("tencentcloud:Rum/getStaticProject:getStaticProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStaticProject.
type GetStaticProjectArgs struct {
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
	// Query Data Type. `allcount`: CostType allcount, `day`: CostType group by day, `condition`: CostType Sorting in condition, `area`: CostType query in area, `nettype`: CostType sort by nettype, `version`: CostType sort by version, `platform`: CostType sort by platform, `isp`: CostType sort by isp, `region`: CostType sort by region, `device`: CostType sort by device, `browser`: CostType sort by browser, `ext1`: CostType sort by ext1, `ext2`: CostType sort by ext2, `ext3`: CostType sort by ext3, `ret`: CostType sort by ret, `status`: CostType sort by status, `from`: CostType sort by from, `url`: CostType sort by url, `env`: CostType sort by env.
	Type string `pulumi:"type"`
	// The URL Key where the data reporting takes place.
	Url *string `pulumi:"url"`
	// The SDK version used for data reporting.
	VersionNum *string `pulumi:"versionNum"`
}

// A collection of values returned by getStaticProject.
type GetStaticProjectResult struct {
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

func GetStaticProjectOutput(ctx *pulumi.Context, args GetStaticProjectOutputArgs, opts ...pulumi.InvokeOption) GetStaticProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStaticProjectResult, error) {
			args := v.(GetStaticProjectArgs)
			r, err := GetStaticProject(ctx, &args, opts...)
			var s GetStaticProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStaticProjectResultOutput)
}

// A collection of arguments for invoking getStaticProject.
type GetStaticProjectOutputArgs struct {
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
	// Query Data Type. `allcount`: CostType allcount, `day`: CostType group by day, `condition`: CostType Sorting in condition, `area`: CostType query in area, `nettype`: CostType sort by nettype, `version`: CostType sort by version, `platform`: CostType sort by platform, `isp`: CostType sort by isp, `region`: CostType sort by region, `device`: CostType sort by device, `browser`: CostType sort by browser, `ext1`: CostType sort by ext1, `ext2`: CostType sort by ext2, `ext3`: CostType sort by ext3, `ret`: CostType sort by ret, `status`: CostType sort by status, `from`: CostType sort by from, `url`: CostType sort by url, `env`: CostType sort by env.
	Type pulumi.StringInput `pulumi:"type"`
	// The URL Key where the data reporting takes place.
	Url pulumi.StringPtrInput `pulumi:"url"`
	// The SDK version used for data reporting.
	VersionNum pulumi.StringPtrInput `pulumi:"versionNum"`
}

func (GetStaticProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticProjectArgs)(nil)).Elem()
}

// A collection of values returned by getStaticProject.
type GetStaticProjectResultOutput struct{ *pulumi.OutputState }

func (GetStaticProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStaticProjectResult)(nil)).Elem()
}

func (o GetStaticProjectResultOutput) ToGetStaticProjectResultOutput() GetStaticProjectResultOutput {
	return o
}

func (o GetStaticProjectResultOutput) ToGetStaticProjectResultOutputWithContext(ctx context.Context) GetStaticProjectResultOutput {
	return o
}

func (o GetStaticProjectResultOutput) Area() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Area }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Brand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Brand }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Browser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Browser }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) CostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.CostType }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) EndTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticProjectResult) int { return v.EndTime }).(pulumi.IntOutput)
}

func (o GetStaticProjectResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Env() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Env }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) ExtFirst() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.ExtFirst }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) ExtSecond() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.ExtSecond }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) ExtThird() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.ExtThird }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.From }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetStaticProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStaticProjectResultOutput) IsAbroad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.IsAbroad }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) NetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.NetType }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Os() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Os }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) Platform() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Platform }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticProjectResult) int { return v.ProjectId }).(pulumi.IntOutput)
}

// Return value.
func (o GetStaticProjectResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticProjectResult) string { return v.Result }).(pulumi.StringOutput)
}

func (o GetStaticProjectResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) StartTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetStaticProjectResult) int { return v.StartTime }).(pulumi.IntOutput)
}

func (o GetStaticProjectResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetStaticProjectResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetStaticProjectResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o GetStaticProjectResultOutput) VersionNum() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStaticProjectResult) *string { return v.VersionNum }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStaticProjectResultOutput{})
}
