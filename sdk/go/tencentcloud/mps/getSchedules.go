// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of mps schedules
//
// ## Example Usage
//
// ### Query the enabled schedules.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mps"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mps.GetSchedules(ctx, &mps.GetSchedulesArgs{
//				Status: pulumi.StringRef("Enabled"),
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
func GetSchedules(ctx *pulumi.Context, args *GetSchedulesArgs, opts ...pulumi.InvokeOption) (*GetSchedulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSchedulesResult
	err := ctx.Invoke("tencentcloud:Mps/getSchedules:getSchedules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSchedules.
type GetSchedulesArgs struct {
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The IDs of the schemes to query. Array length limit: 100.
	ScheduleIds []int `pulumi:"scheduleIds"`
	// The scheme status. Valid values:`Enabled`, `Disabled`. If you do not specify this parameter, all schemes will be returned regardless of the status.
	Status *string `pulumi:"status"`
	// The trigger type. Valid values:`CosFileUpload`: The scheme is triggered when a file is uploaded to Tencent Cloud Object Storage (COS).`AwsS3FileUpload`: The scheme is triggered when a file is uploaded to AWS S3.If you do not specify this parameter or leave it empty, all schemes will be returned regardless of the trigger type.
	TriggerType *string `pulumi:"triggerType"`
}

// A collection of values returned by getSchedules.
type GetSchedulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	ScheduleIds      []int   `pulumi:"scheduleIds"`
	// The information of the schemes.
	ScheduleInfoSets []GetSchedulesScheduleInfoSet `pulumi:"scheduleInfoSets"`
	// The scheme status. Valid values:`Enabled``Disabled`Note: This field may return null, indicating that no valid values can be obtained.
	Status      *string `pulumi:"status"`
	TriggerType *string `pulumi:"triggerType"`
}

func GetSchedulesOutput(ctx *pulumi.Context, args GetSchedulesOutputArgs, opts ...pulumi.InvokeOption) GetSchedulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSchedulesResult, error) {
			args := v.(GetSchedulesArgs)
			r, err := GetSchedules(ctx, &args, opts...)
			var s GetSchedulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSchedulesResultOutput)
}

// A collection of arguments for invoking getSchedules.
type GetSchedulesOutputArgs struct {
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The IDs of the schemes to query. Array length limit: 100.
	ScheduleIds pulumi.IntArrayInput `pulumi:"scheduleIds"`
	// The scheme status. Valid values:`Enabled`, `Disabled`. If you do not specify this parameter, all schemes will be returned regardless of the status.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The trigger type. Valid values:`CosFileUpload`: The scheme is triggered when a file is uploaded to Tencent Cloud Object Storage (COS).`AwsS3FileUpload`: The scheme is triggered when a file is uploaded to AWS S3.If you do not specify this parameter or leave it empty, all schemes will be returned regardless of the trigger type.
	TriggerType pulumi.StringPtrInput `pulumi:"triggerType"`
}

func (GetSchedulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSchedulesArgs)(nil)).Elem()
}

// A collection of values returned by getSchedules.
type GetSchedulesResultOutput struct{ *pulumi.OutputState }

func (GetSchedulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSchedulesResult)(nil)).Elem()
}

func (o GetSchedulesResultOutput) ToGetSchedulesResultOutput() GetSchedulesResultOutput {
	return o
}

func (o GetSchedulesResultOutput) ToGetSchedulesResultOutputWithContext(ctx context.Context) GetSchedulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSchedulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSchedulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSchedulesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSchedulesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSchedulesResultOutput) ScheduleIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetSchedulesResult) []int { return v.ScheduleIds }).(pulumi.IntArrayOutput)
}

// The information of the schemes.
func (o GetSchedulesResultOutput) ScheduleInfoSets() GetSchedulesScheduleInfoSetArrayOutput {
	return o.ApplyT(func(v GetSchedulesResult) []GetSchedulesScheduleInfoSet { return v.ScheduleInfoSets }).(GetSchedulesScheduleInfoSetArrayOutput)
}

// The scheme status. Valid values:`Enabled“Disabled`Note: This field may return null, indicating that no valid values can be obtained.
func (o GetSchedulesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSchedulesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o GetSchedulesResultOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSchedulesResult) *string { return v.TriggerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSchedulesResultOutput{})
}
