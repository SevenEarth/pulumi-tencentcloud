// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mongodb instanceBackups
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mongodb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mongodb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mongodb.GetInstanceBackups(ctx, &mongodb.GetInstanceBackupsArgs{
//				BackupMethod: pulumi.IntRef(0),
//				InstanceId:   "cmgo-9d0p6umb",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstanceBackups(ctx *pulumi.Context, args *GetInstanceBackupsArgs, opts ...pulumi.InvokeOption) (*GetInstanceBackupsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstanceBackupsResult
	err := ctx.Invoke("tencentcloud:Mongodb/getInstanceBackups:getInstanceBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceBackups.
type GetInstanceBackupsArgs struct {
	// Backup mode, currently supported: 0-logic backup, 1-physical backup, 2-all backups.The default is logical backup.
	BackupMethod *int `pulumi:"backupMethod"`
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getInstanceBackups.
type GetInstanceBackupsResult struct {
	// backup list.
	BackupLists []GetInstanceBackupsBackupList `pulumi:"backupLists"`
	// Backup method.
	BackupMethod *int `pulumi:"backupMethod"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance ID.
	InstanceId       string  `pulumi:"instanceId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

func GetInstanceBackupsOutput(ctx *pulumi.Context, args GetInstanceBackupsOutputArgs, opts ...pulumi.InvokeOption) GetInstanceBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstanceBackupsResult, error) {
			args := v.(GetInstanceBackupsArgs)
			r, err := GetInstanceBackups(ctx, &args, opts...)
			var s GetInstanceBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstanceBackupsResultOutput)
}

// A collection of arguments for invoking getInstanceBackups.
type GetInstanceBackupsOutputArgs struct {
	// Backup mode, currently supported: 0-logic backup, 1-physical backup, 2-all backups.The default is logical backup.
	BackupMethod pulumi.IntPtrInput `pulumi:"backupMethod"`
	// Instance ID, the format is: cmgo-9d0p6umb.Same as the instance ID displayed in the cloud database console page.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetInstanceBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceBackupsArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceBackups.
type GetInstanceBackupsResultOutput struct{ *pulumi.OutputState }

func (GetInstanceBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstanceBackupsResult)(nil)).Elem()
}

func (o GetInstanceBackupsResultOutput) ToGetInstanceBackupsResultOutput() GetInstanceBackupsResultOutput {
	return o
}

func (o GetInstanceBackupsResultOutput) ToGetInstanceBackupsResultOutputWithContext(ctx context.Context) GetInstanceBackupsResultOutput {
	return o
}

// backup list.
func (o GetInstanceBackupsResultOutput) BackupLists() GetInstanceBackupsBackupListArrayOutput {
	return o.ApplyT(func(v GetInstanceBackupsResult) []GetInstanceBackupsBackupList { return v.BackupLists }).(GetInstanceBackupsBackupListArrayOutput)
}

// Backup method.
func (o GetInstanceBackupsResultOutput) BackupMethod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetInstanceBackupsResult) *int { return v.BackupMethod }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstanceBackupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceBackupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance ID.
func (o GetInstanceBackupsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstanceBackupsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetInstanceBackupsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstanceBackupsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstanceBackupsResultOutput{})
}
