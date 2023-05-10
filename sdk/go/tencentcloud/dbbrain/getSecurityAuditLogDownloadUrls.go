// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dbbrain securityAuditLogDownloadUrls
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		task, err := Dbbrain.NewSecurityAuditLogExportTask(ctx, "task", &Dbbrain.SecurityAuditLogExportTaskArgs{
// 			SecAuditGroupId: pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
// 			StartTime:       pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
// 			EndTime:         pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
// 			Product:         pulumi.String("mysql"),
// 			DangerLevels: pulumi.IntArray{
// 				pulumi.Int(0),
// 				pulumi.Int(1),
// 				pulumi.Int(2),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_ = Dbbrain.GetSecurityAuditLogDownloadUrlsOutput(ctx, dbbrain.GetSecurityAuditLogDownloadUrlsOutputArgs{
// 			SecAuditGroupId: pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
// 			AsyncRequestId:  task.AsyncRequestId,
// 			Product:         pulumi.String("mysql"),
// 		}, nil)
// 		return nil
// 	})
// }
// ```
func GetSecurityAuditLogDownloadUrls(ctx *pulumi.Context, args *GetSecurityAuditLogDownloadUrlsArgs, opts ...pulumi.InvokeOption) (*GetSecurityAuditLogDownloadUrlsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecurityAuditLogDownloadUrlsResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getSecurityAuditLogDownloadUrls:getSecurityAuditLogDownloadUrls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityAuditLogDownloadUrls.
type GetSecurityAuditLogDownloadUrlsArgs struct {
	// Asynchronous task ID.
	AsyncRequestId int `pulumi:"asyncRequestId"`
	// Service product type, supported values: `mysql` - ApsaraDB for MySQL.
	Product string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Security audit group Id.
	SecAuditGroupId string `pulumi:"secAuditGroupId"`
}

// A collection of values returned by getSecurityAuditLogDownloadUrls.
type GetSecurityAuditLogDownloadUrlsResult struct {
	AsyncRequestId int `pulumi:"asyncRequestId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	Product          string  `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	SecAuditGroupId  string  `pulumi:"secAuditGroupId"`
	// List of COS links to export results. When the result set is large, it may be divided into multiple urls for download.
	Urls []string `pulumi:"urls"`
}

func GetSecurityAuditLogDownloadUrlsOutput(ctx *pulumi.Context, args GetSecurityAuditLogDownloadUrlsOutputArgs, opts ...pulumi.InvokeOption) GetSecurityAuditLogDownloadUrlsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecurityAuditLogDownloadUrlsResult, error) {
			args := v.(GetSecurityAuditLogDownloadUrlsArgs)
			r, err := GetSecurityAuditLogDownloadUrls(ctx, &args, opts...)
			var s GetSecurityAuditLogDownloadUrlsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecurityAuditLogDownloadUrlsResultOutput)
}

// A collection of arguments for invoking getSecurityAuditLogDownloadUrls.
type GetSecurityAuditLogDownloadUrlsOutputArgs struct {
	// Asynchronous task ID.
	AsyncRequestId pulumi.IntInput `pulumi:"asyncRequestId"`
	// Service product type, supported values: `mysql` - ApsaraDB for MySQL.
	Product pulumi.StringInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Security audit group Id.
	SecAuditGroupId pulumi.StringInput `pulumi:"secAuditGroupId"`
}

func (GetSecurityAuditLogDownloadUrlsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityAuditLogDownloadUrlsArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityAuditLogDownloadUrls.
type GetSecurityAuditLogDownloadUrlsResultOutput struct{ *pulumi.OutputState }

func (GetSecurityAuditLogDownloadUrlsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecurityAuditLogDownloadUrlsResult)(nil)).Elem()
}

func (o GetSecurityAuditLogDownloadUrlsResultOutput) ToGetSecurityAuditLogDownloadUrlsResultOutput() GetSecurityAuditLogDownloadUrlsResultOutput {
	return o
}

func (o GetSecurityAuditLogDownloadUrlsResultOutput) ToGetSecurityAuditLogDownloadUrlsResultOutputWithContext(ctx context.Context) GetSecurityAuditLogDownloadUrlsResultOutput {
	return o
}

func (o GetSecurityAuditLogDownloadUrlsResultOutput) AsyncRequestId() pulumi.IntOutput {
	return o.ApplyT(func(v GetSecurityAuditLogDownloadUrlsResult) int { return v.AsyncRequestId }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSecurityAuditLogDownloadUrlsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityAuditLogDownloadUrlsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSecurityAuditLogDownloadUrlsResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityAuditLogDownloadUrlsResult) string { return v.Product }).(pulumi.StringOutput)
}

func (o GetSecurityAuditLogDownloadUrlsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecurityAuditLogDownloadUrlsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSecurityAuditLogDownloadUrlsResultOutput) SecAuditGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecurityAuditLogDownloadUrlsResult) string { return v.SecAuditGroupId }).(pulumi.StringOutput)
}

// List of COS links to export results. When the result set is large, it may be divided into multiple urls for download.
func (o GetSecurityAuditLogDownloadUrlsResultOutput) Urls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecurityAuditLogDownloadUrlsResult) []string { return v.Urls }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecurityAuditLogDownloadUrlsResultOutput{})
}
