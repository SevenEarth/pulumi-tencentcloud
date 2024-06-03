// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of mariadb logFiles
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.GetLogFiles(ctx, &mariadb.GetLogFilesArgs{
//				InstanceId: "tdsql-9vqvls95",
//				Type:       1,
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
func GetLogFiles(ctx *pulumi.Context, args *GetLogFilesArgs, opts ...pulumi.InvokeOption) (*GetLogFilesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogFilesResult
	err := ctx.Invoke("tencentcloud:Mariadb/getLogFiles:getLogFiles", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogFiles.
type GetLogFilesArgs struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
	Type int `pulumi:"type"`
}

// A collection of values returned by getLogFiles.
type GetLogFilesResult struct {
	// Information such as `uri`, `length`, and `mtime` (modification time).
	Files []GetLogFilesFile `pulumi:"files"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// For an instance in a common network, this prefix plus URI can be used as the download address.
	NormalPrefix     string  `pulumi:"normalPrefix"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Type             int     `pulumi:"type"`
	// For an instance in a VPC, this prefix plus URI can be used as the download address.
	VpcPrefix string `pulumi:"vpcPrefix"`
}

func GetLogFilesOutput(ctx *pulumi.Context, args GetLogFilesOutputArgs, opts ...pulumi.InvokeOption) GetLogFilesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLogFilesResult, error) {
			args := v.(GetLogFilesArgs)
			r, err := GetLogFiles(ctx, &args, opts...)
			var s GetLogFilesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLogFilesResultOutput)
}

// A collection of arguments for invoking getLogFiles.
type GetLogFilesOutputArgs struct {
	// Instance ID in the format of `tdsql-ow728lmc`.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Requested log type. Valid values: 1 (binlog), 2 (cold backup), 3 (errlog), 4 (slowlog).
	Type pulumi.IntInput `pulumi:"type"`
}

func (GetLogFilesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogFilesArgs)(nil)).Elem()
}

// A collection of values returned by getLogFiles.
type GetLogFilesResultOutput struct{ *pulumi.OutputState }

func (GetLogFilesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogFilesResult)(nil)).Elem()
}

func (o GetLogFilesResultOutput) ToGetLogFilesResultOutput() GetLogFilesResultOutput {
	return o
}

func (o GetLogFilesResultOutput) ToGetLogFilesResultOutputWithContext(ctx context.Context) GetLogFilesResultOutput {
	return o
}

// Information such as `uri`, `length`, and `mtime` (modification time).
func (o GetLogFilesResultOutput) Files() GetLogFilesFileArrayOutput {
	return o.ApplyT(func(v GetLogFilesResult) []GetLogFilesFile { return v.Files }).(GetLogFilesFileArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogFilesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogFilesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogFilesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogFilesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// For an instance in a common network, this prefix plus URI can be used as the download address.
func (o GetLogFilesResultOutput) NormalPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogFilesResult) string { return v.NormalPrefix }).(pulumi.StringOutput)
}

func (o GetLogFilesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLogFilesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetLogFilesResultOutput) Type() pulumi.IntOutput {
	return o.ApplyT(func(v GetLogFilesResult) int { return v.Type }).(pulumi.IntOutput)
}

// For an instance in a VPC, this prefix plus URI can be used as the download address.
func (o GetLogFilesResultOutput) VpcPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogFilesResult) string { return v.VpcPrefix }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogFilesResultOutput{})
}
