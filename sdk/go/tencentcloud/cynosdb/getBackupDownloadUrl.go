// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBackupDownloadUrl(ctx *pulumi.Context, args *GetBackupDownloadUrlArgs, opts ...pulumi.InvokeOption) (*GetBackupDownloadUrlResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetBackupDownloadUrlResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getBackupDownloadUrl:getBackupDownloadUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupDownloadUrl.
type GetBackupDownloadUrlArgs struct {
	BackupId         int     `pulumi:"backupId"`
	ClusterId        string  `pulumi:"clusterId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getBackupDownloadUrl.
type GetBackupDownloadUrlResult struct {
	BackupId    int    `pulumi:"backupId"`
	ClusterId   string `pulumi:"clusterId"`
	DownloadUrl string `pulumi:"downloadUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetBackupDownloadUrlOutput(ctx *pulumi.Context, args GetBackupDownloadUrlOutputArgs, opts ...pulumi.InvokeOption) GetBackupDownloadUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBackupDownloadUrlResult, error) {
			args := v.(GetBackupDownloadUrlArgs)
			r, err := GetBackupDownloadUrl(ctx, &args, opts...)
			var s GetBackupDownloadUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBackupDownloadUrlResultOutput)
}

// A collection of arguments for invoking getBackupDownloadUrl.
type GetBackupDownloadUrlOutputArgs struct {
	BackupId         pulumi.IntInput       `pulumi:"backupId"`
	ClusterId        pulumi.StringInput    `pulumi:"clusterId"`
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetBackupDownloadUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupDownloadUrlArgs)(nil)).Elem()
}

// A collection of values returned by getBackupDownloadUrl.
type GetBackupDownloadUrlResultOutput struct{ *pulumi.OutputState }

func (GetBackupDownloadUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBackupDownloadUrlResult)(nil)).Elem()
}

func (o GetBackupDownloadUrlResultOutput) ToGetBackupDownloadUrlResultOutput() GetBackupDownloadUrlResultOutput {
	return o
}

func (o GetBackupDownloadUrlResultOutput) ToGetBackupDownloadUrlResultOutputWithContext(ctx context.Context) GetBackupDownloadUrlResultOutput {
	return o
}

func (o GetBackupDownloadUrlResultOutput) BackupId() pulumi.IntOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlResult) int { return v.BackupId }).(pulumi.IntOutput)
}

func (o GetBackupDownloadUrlResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

func (o GetBackupDownloadUrlResultOutput) DownloadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlResult) string { return v.DownloadUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBackupDownloadUrlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBackupDownloadUrlResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBackupDownloadUrlResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBackupDownloadUrlResultOutput{})
}
