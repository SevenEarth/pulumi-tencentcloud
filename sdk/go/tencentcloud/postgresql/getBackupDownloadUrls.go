// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of postgresql backupDownloadUrls
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// logBackups, err := Postgresql.GetLogBackups(ctx, &postgresql.GetLogBackupsArgs{
// MinFinishTime: pulumi.StringRef("%s"),
// MaxFinishTime: pulumi.StringRef("%s"),
// Filters: []postgresql.GetLogBackupsFilter{
// {
// Name: pulumi.StringRef("db-instance-id"),
// Values: interface{}{
// local.Pgsql_id,
// },
// },
// },
// OrderBy: pulumi.StringRef("StartTime"),
// OrderByType: pulumi.StringRef("desc"),
// }, nil);
// if err != nil {
// return err
// }
// _, err = Postgresql.GetBackupDownloadUrls(ctx, &postgresql.GetBackupDownloadUrlsArgs{
// DbInstanceId: local.Pgsql_id,
// BackupType: "LogBackup",
// BackupId: logBackups.LogBackupSets[0].Id,
// UrlExpireTime: pulumi.IntRef(12),
// BackupDownloadRestriction: postgresql.GetBackupDownloadUrlsBackupDownloadRestriction{
// RestrictionType: pulumi.StringRef("NONE"),
// VpcRestrictionEffect: pulumi.StringRef("ALLOW"),
// VpcIdSets: interface{}{
// local.Vpc_id,
// },
// IpRestrictionEffect: pulumi.StringRef("ALLOW"),
// IpSets: []string{
// "0.0.0.0",
// },
// },
// }, nil);
// if err != nil {
// return err
// }
// return nil
// })
// }
// ```
// <!--End PulumiCodeChooser -->
func GetBackupDownloadUrls(ctx *pulumi.Context, args *GetBackupDownloadUrlsArgs, opts ...pulumi.InvokeOption) (*GetBackupDownloadUrlsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBackupDownloadUrlsResult
	err := ctx.Invoke("tencentcloud:Postgresql/getBackupDownloadUrls:getBackupDownloadUrls", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupDownloadUrls.
type GetBackupDownloadUrlsArgs struct {
	// Backup download restriction.
	BackupDownloadRestriction *GetBackupDownloadUrlsBackupDownloadRestriction `pulumi:"backupDownloadRestriction"`
	// Unique backup ID.
	BackupId string `pulumi:"backupId"`
	// Backup type. Valid values: `LogBackup`, `BaseBackup`.
	BackupType string `pulumi:"backupType"`
	// Instance ID.
	DbInstanceId string `pulumi:"dbInstanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Validity period of a URL, which is 12 hours by default.
	UrlExpireTime *int `pulumi:"urlExpireTime"`
}

// A collection of values returned by getBackupDownloadUrls.
type GetBackupDownloadUrlsResult struct {
	BackupDownloadRestriction *GetBackupDownloadUrlsBackupDownloadRestriction `pulumi:"backupDownloadRestriction"`
	// Backup download URL.
	BackupDownloadUrl string `pulumi:"backupDownloadUrl"`
	BackupId          string `pulumi:"backupId"`
	BackupType        string `pulumi:"backupType"`
	DbInstanceId      string `pulumi:"dbInstanceId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	UrlExpireTime    *int    `pulumi:"urlExpireTime"`
}

func GetBackupDownloadUrlsOutput(ctx *pulumi.Context, args GetBackupDownloadUrlsOutputArgs, opts ...pulumi.InvokeOption) GetBackupDownloadUrlsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupDownloadUrlsResult, error) {
			args := v.(GetBackupDownloadUrlsArgs)
			r, err := GetBackupDownloadUrls(ctx, &args, opts...)
			var s GetBackupDownloadUrlsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupDownloadUrlsResultOutput)
}

// A collection of arguments for invoking getBackupDownloadUrls.
type GetBackupDownloadUrlsOutputArgs struct {
	// Backup download restriction.
	BackupDownloadRestriction GetBackupDownloadUrlsBackupDownloadRestrictionPtrInput `pulumi:"backupDownloadRestriction"`
	// Unique backup ID.
	BackupId pulumi.StringInput `pulumi:"backupId"`
	// Backup type. Valid values: `LogBackup`, `BaseBackup`.
	BackupType pulumi.StringInput `pulumi:"backupType"`
	// Instance ID.
	DbInstanceId pulumi.StringInput `pulumi:"dbInstanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Validity period of a URL, which is 12 hours by default.
	UrlExpireTime pulumi.IntPtrInput `pulumi:"urlExpireTime"`
}

func (GetBackupDownloadUrlsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupDownloadUrlsArgs)(nil)).Elem()
}

// A collection of values returned by getBackupDownloadUrls.
type GetBackupDownloadUrlsResultOutput struct{ *pulumi.OutputState }

func (GetBackupDownloadUrlsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupDownloadUrlsResult)(nil)).Elem()
}

func (o GetBackupDownloadUrlsResultOutput) ToGetBackupDownloadUrlsResultOutput() GetBackupDownloadUrlsResultOutput {
	return o
}

func (o GetBackupDownloadUrlsResultOutput) ToGetBackupDownloadUrlsResultOutputWithContext(ctx context.Context) GetBackupDownloadUrlsResultOutput {
	return o
}

func (o GetBackupDownloadUrlsResultOutput) BackupDownloadRestriction() GetBackupDownloadUrlsBackupDownloadRestrictionPtrOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) *GetBackupDownloadUrlsBackupDownloadRestriction {
		return v.BackupDownloadRestriction
	}).(GetBackupDownloadUrlsBackupDownloadRestrictionPtrOutput)
}

// Backup download URL.
func (o GetBackupDownloadUrlsResultOutput) BackupDownloadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) string { return v.BackupDownloadUrl }).(pulumi.StringOutput)
}

func (o GetBackupDownloadUrlsResultOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) string { return v.BackupId }).(pulumi.StringOutput)
}

func (o GetBackupDownloadUrlsResultOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) string { return v.BackupType }).(pulumi.StringOutput)
}

func (o GetBackupDownloadUrlsResultOutput) DbInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) string { return v.DbInstanceId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupDownloadUrlsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBackupDownloadUrlsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetBackupDownloadUrlsResultOutput) UrlExpireTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlsResult) *int { return v.UrlExpireTime }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupDownloadUrlsResultOutput{})
}
