// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqlserver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query the list of SQL Server backups.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sqlserver"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Sqlserver.GetBackups(ctx, &sqlserver.GetBackupsArgs{
//				EndTime:    "2020-06-22 00:00:00",
//				InstanceId: "mssql-3cdq7kx5",
//				StartTime:  "2020-06-17 00:00:00",
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
func GetBackups(ctx *pulumi.Context, args *GetBackupsArgs, opts ...pulumi.InvokeOption) (*GetBackupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackupsResult
	err := ctx.Invoke("tencentcloud:Sqlserver/getBackups:getBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackups.
type GetBackupsArgs struct {
	BackupName *string `pulumi:"backupName"`
	// End time of the instance list, like yyyy-MM-dd HH:mm:ss.
	EndTime string `pulumi:"endTime"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Start time of the instance list, like yyyy-MM-dd HH:mm:ss.
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by getBackups.
type GetBackupsResult struct {
	BackupName *string `pulumi:"backupName"`
	// End time of the backup.
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// A list of SQL Server backup. Each element contains the following attributes:
	Lists            []GetBackupsList `pulumi:"lists"`
	ResultOutputFile *string          `pulumi:"resultOutputFile"`
	// Start time of the backup.
	StartTime string `pulumi:"startTime"`
}

func GetBackupsOutput(ctx *pulumi.Context, args GetBackupsOutputArgs, opts ...pulumi.InvokeOption) GetBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupsResult, error) {
			args := v.(GetBackupsArgs)
			r, err := GetBackups(ctx, &args, opts...)
			var s GetBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupsResultOutput)
}

// A collection of arguments for invoking getBackups.
type GetBackupsOutputArgs struct {
	BackupName pulumi.StringPtrInput `pulumi:"backupName"`
	// End time of the instance list, like yyyy-MM-dd HH:mm:ss.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// Instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Start time of the instance list, like yyyy-MM-dd HH:mm:ss.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (GetBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupsArgs)(nil)).Elem()
}

// A collection of values returned by getBackups.
type GetBackupsResultOutput struct{ *pulumi.OutputState }

func (GetBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupsResult)(nil)).Elem()
}

func (o GetBackupsResultOutput) ToGetBackupsResultOutput() GetBackupsResultOutput {
	return o
}

func (o GetBackupsResultOutput) ToGetBackupsResultOutputWithContext(ctx context.Context) GetBackupsResultOutput {
	return o
}

func (o GetBackupsResultOutput) BackupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.BackupName }).(pulumi.StringPtrOutput)
}

// End time of the backup.
func (o GetBackupsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance ID.
func (o GetBackupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// A list of SQL Server backup. Each element contains the following attributes:
func (o GetBackupsResultOutput) Lists() GetBackupsListArrayOutput {
	return o.ApplyT(func(v GetBackupsResult) []GetBackupsList { return v.Lists }).(GetBackupsListArrayOutput)
}

func (o GetBackupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Start time of the backup.
func (o GetBackupsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupsResultOutput{})
}
