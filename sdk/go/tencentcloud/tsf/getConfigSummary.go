// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of tsf configSummary
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.GetConfigSummary(ctx, &tsf.GetConfigSummaryArgs{
//				ApplicationId: pulumi.StringRef("application-a24x29xv"),
//				ConfigIdLists: []string{
//					"dcfg-y54wzk3a",
//				},
//				DisableProgramAuthCheck: pulumi.BoolRef(true),
//				OrderBy:                 pulumi.StringRef("last_update_time"),
//				OrderType:               pulumi.IntRef(0),
//				SearchWord:              pulumi.StringRef("terraform"),
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
func LookupConfigSummary(ctx *pulumi.Context, args *LookupConfigSummaryArgs, opts ...pulumi.InvokeOption) (*LookupConfigSummaryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigSummaryResult
	err := ctx.Invoke("tencentcloud:Tsf/getConfigSummary:getConfigSummary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfigSummary.
type LookupConfigSummaryArgs struct {
	// Application ID. If not passed, the query will be for all.
	ApplicationId *string `pulumi:"applicationId"`
	// Config Id List.
	ConfigIdLists []string `pulumi:"configIdLists"`
	// config tag list.
	ConfigTagLists []string `pulumi:"configTagLists"`
	// Whether to disable dataset authentication.
	DisableProgramAuthCheck *bool `pulumi:"disableProgramAuthCheck"`
	// Order term. support Sort by time: creation_time; or Sort by name: config_name.
	OrderBy *string `pulumi:"orderBy"`
	// Pass 0 for ascending order and 1 for descending order.
	OrderType *int `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Query keyword, fuzzy query: application name, configuration item name. If not passed, the query will be for all.
	SearchWord *string `pulumi:"searchWord"`
}

// A collection of values returned by getConfigSummary.
type LookupConfigSummaryResult struct {
	// Application ID.Note: This field may return null, indicating that no valid value was found.
	ApplicationId           *string  `pulumi:"applicationId"`
	ConfigIdLists           []string `pulumi:"configIdLists"`
	ConfigTagLists          []string `pulumi:"configTagLists"`
	DisableProgramAuthCheck *bool    `pulumi:"disableProgramAuthCheck"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	OrderBy          *string `pulumi:"orderBy"`
	OrderType        *int    `pulumi:"orderType"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// config Page Item.
	Results    []GetConfigSummaryResult `pulumi:"results"`
	SearchWord *string                  `pulumi:"searchWord"`
}

func LookupConfigSummaryOutput(ctx *pulumi.Context, args LookupConfigSummaryOutputArgs, opts ...pulumi.InvokeOption) LookupConfigSummaryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigSummaryResult, error) {
			args := v.(LookupConfigSummaryArgs)
			r, err := LookupConfigSummary(ctx, &args, opts...)
			var s LookupConfigSummaryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigSummaryResultOutput)
}

// A collection of arguments for invoking getConfigSummary.
type LookupConfigSummaryOutputArgs struct {
	// Application ID. If not passed, the query will be for all.
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	// Config Id List.
	ConfigIdLists pulumi.StringArrayInput `pulumi:"configIdLists"`
	// config tag list.
	ConfigTagLists pulumi.StringArrayInput `pulumi:"configTagLists"`
	// Whether to disable dataset authentication.
	DisableProgramAuthCheck pulumi.BoolPtrInput `pulumi:"disableProgramAuthCheck"`
	// Order term. support Sort by time: creation_time; or Sort by name: config_name.
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
	// Pass 0 for ascending order and 1 for descending order.
	OrderType pulumi.IntPtrInput `pulumi:"orderType"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Query keyword, fuzzy query: application name, configuration item name. If not passed, the query will be for all.
	SearchWord pulumi.StringPtrInput `pulumi:"searchWord"`
}

func (LookupConfigSummaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigSummaryArgs)(nil)).Elem()
}

// A collection of values returned by getConfigSummary.
type LookupConfigSummaryResultOutput struct{ *pulumi.OutputState }

func (LookupConfigSummaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigSummaryResult)(nil)).Elem()
}

func (o LookupConfigSummaryResultOutput) ToLookupConfigSummaryResultOutput() LookupConfigSummaryResultOutput {
	return o
}

func (o LookupConfigSummaryResultOutput) ToLookupConfigSummaryResultOutputWithContext(ctx context.Context) LookupConfigSummaryResultOutput {
	return o
}

// Application ID.Note: This field may return null, indicating that no valid value was found.
func (o LookupConfigSummaryResultOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o LookupConfigSummaryResultOutput) ConfigIdLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) []string { return v.ConfigIdLists }).(pulumi.StringArrayOutput)
}

func (o LookupConfigSummaryResultOutput) ConfigTagLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) []string { return v.ConfigTagLists }).(pulumi.StringArrayOutput)
}

func (o LookupConfigSummaryResultOutput) DisableProgramAuthCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) *bool { return v.DisableProgramAuthCheck }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupConfigSummaryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupConfigSummaryResultOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

func (o LookupConfigSummaryResultOutput) OrderType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) *int { return v.OrderType }).(pulumi.IntPtrOutput)
}

func (o LookupConfigSummaryResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// config Page Item.
func (o LookupConfigSummaryResultOutput) Results() GetConfigSummaryResultArrayOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) []GetConfigSummaryResult { return v.Results }).(GetConfigSummaryResultArrayOutput)
}

func (o LookupConfigSummaryResultOutput) SearchWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigSummaryResult) *string { return v.SearchWord }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigSummaryResultOutput{})
}
