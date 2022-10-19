// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a datasource to query PostgreSQL Xlogs.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Postgresql.GetXlogs(ctx, &postgresql.GetXlogsArgs{
//				EndTime:    pulumi.StringRef("2022-01-07 01:02:03"),
//				InstanceId: "postgres-xxxxxxxx",
//				StartTime:  pulumi.StringRef("2022-01-01 00:00:00"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetXlogs(ctx *pulumi.Context, args *GetXlogsArgs, opts ...pulumi.InvokeOption) (*GetXlogsResult, error) {
	var rv GetXlogsResult
	err := ctx.Invoke("tencentcloud:Postgresql/getXlogs:getXlogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getXlogs.
type GetXlogsArgs struct {
	// Xlog end time, format `yyyy-MM-dd hh:mm:ss`.
	EndTime *string `pulumi:"endTime"`
	// PostgreSQL instance id.
	InstanceId string `pulumi:"instanceId"`
	// Used for save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Xlog start time, format `yyyy-MM-dd hh:mm:ss`, start time cannot before 7 days ago.
	StartTime *string `pulumi:"startTime"`
}

// A collection of values returned by getXlogs.
type GetXlogsResult struct {
	// Xlog file created end time.
	EndTime *string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// List of Xlog query result.
	Lists            []GetXlogsList `pulumi:"lists"`
	ResultOutputFile *string        `pulumi:"resultOutputFile"`
	// Xlog file created start time.
	StartTime *string `pulumi:"startTime"`
}

func GetXlogsOutput(ctx *pulumi.Context, args GetXlogsOutputArgs, opts ...pulumi.InvokeOption) GetXlogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetXlogsResult, error) {
			args := v.(GetXlogsArgs)
			r, err := GetXlogs(ctx, &args, opts...)
			var s GetXlogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetXlogsResultOutput)
}

// A collection of arguments for invoking getXlogs.
type GetXlogsOutputArgs struct {
	// Xlog end time, format `yyyy-MM-dd hh:mm:ss`.
	EndTime pulumi.StringPtrInput `pulumi:"endTime"`
	// PostgreSQL instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used for save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Xlog start time, format `yyyy-MM-dd hh:mm:ss`, start time cannot before 7 days ago.
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (GetXlogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetXlogsArgs)(nil)).Elem()
}

// A collection of values returned by getXlogs.
type GetXlogsResultOutput struct{ *pulumi.OutputState }

func (GetXlogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetXlogsResult)(nil)).Elem()
}

func (o GetXlogsResultOutput) ToGetXlogsResultOutput() GetXlogsResultOutput {
	return o
}

func (o GetXlogsResultOutput) ToGetXlogsResultOutputWithContext(ctx context.Context) GetXlogsResultOutput {
	return o
}

// Xlog file created end time.
func (o GetXlogsResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetXlogsResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetXlogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetXlogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetXlogsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetXlogsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// List of Xlog query result.
func (o GetXlogsResultOutput) Lists() GetXlogsListArrayOutput {
	return o.ApplyT(func(v GetXlogsResult) []GetXlogsList { return v.Lists }).(GetXlogsListArrayOutput)
}

func (o GetXlogsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetXlogsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Xlog file created start time.
func (o GetXlogsResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetXlogsResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetXlogsResultOutput{})
}
