// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mysql rollbackRangeTime
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mysql.GetRollbackRangeTime(ctx, &mysql.GetRollbackRangeTimeArgs{
//				InstanceIds: []string{
//					"cdb-fitq5t9h",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRollbackRangeTime(ctx *pulumi.Context, args *GetRollbackRangeTimeArgs, opts ...pulumi.InvokeOption) (*GetRollbackRangeTimeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetRollbackRangeTimeResult
	err := ctx.Invoke("tencentcloud:Mysql/getRollbackRangeTime:getRollbackRangeTime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRollbackRangeTime.
type GetRollbackRangeTimeArgs struct {
	// If the clone instance is not in the same region as the source instance, fill in the region where the clone instance is located, for example: ap-guangzhou.
	BackupRegion *string `pulumi:"backupRegion"`
	// A list of instance IDs, the format of a single instance ID is: cdb-c1nl9rpv. Same instance ID as displayed in the ApsaraDB for Console page.
	InstanceIds []string `pulumi:"instanceIds"`
	// Whether the clone instance is in the same zone as the source instance, yes: `false`, no: `true`.
	IsRemoteZone *string `pulumi:"isRemoteZone"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getRollbackRangeTime.
type GetRollbackRangeTimeResult struct {
	BackupRegion *string `pulumi:"backupRegion"`
	// The provider-assigned unique ID for this managed resource.
	Id           string   `pulumi:"id"`
	InstanceIds  []string `pulumi:"instanceIds"`
	IsRemoteZone *string  `pulumi:"isRemoteZone"`
	// Returned parameter information.
	Items            []GetRollbackRangeTimeItem `pulumi:"items"`
	ResultOutputFile *string                    `pulumi:"resultOutputFile"`
}

func GetRollbackRangeTimeOutput(ctx *pulumi.Context, args GetRollbackRangeTimeOutputArgs, opts ...pulumi.InvokeOption) GetRollbackRangeTimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRollbackRangeTimeResult, error) {
			args := v.(GetRollbackRangeTimeArgs)
			r, err := GetRollbackRangeTime(ctx, &args, opts...)
			var s GetRollbackRangeTimeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRollbackRangeTimeResultOutput)
}

// A collection of arguments for invoking getRollbackRangeTime.
type GetRollbackRangeTimeOutputArgs struct {
	// If the clone instance is not in the same region as the source instance, fill in the region where the clone instance is located, for example: ap-guangzhou.
	BackupRegion pulumi.StringPtrInput `pulumi:"backupRegion"`
	// A list of instance IDs, the format of a single instance ID is: cdb-c1nl9rpv. Same instance ID as displayed in the ApsaraDB for Console page.
	InstanceIds pulumi.StringArrayInput `pulumi:"instanceIds"`
	// Whether the clone instance is in the same zone as the source instance, yes: `false`, no: `true`.
	IsRemoteZone pulumi.StringPtrInput `pulumi:"isRemoteZone"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetRollbackRangeTimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRollbackRangeTimeArgs)(nil)).Elem()
}

// A collection of values returned by getRollbackRangeTime.
type GetRollbackRangeTimeResultOutput struct{ *pulumi.OutputState }

func (GetRollbackRangeTimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRollbackRangeTimeResult)(nil)).Elem()
}

func (o GetRollbackRangeTimeResultOutput) ToGetRollbackRangeTimeResultOutput() GetRollbackRangeTimeResultOutput {
	return o
}

func (o GetRollbackRangeTimeResultOutput) ToGetRollbackRangeTimeResultOutputWithContext(ctx context.Context) GetRollbackRangeTimeResultOutput {
	return o
}

func (o GetRollbackRangeTimeResultOutput) BackupRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRollbackRangeTimeResult) *string { return v.BackupRegion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRollbackRangeTimeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRollbackRangeTimeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRollbackRangeTimeResultOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRollbackRangeTimeResult) []string { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

func (o GetRollbackRangeTimeResultOutput) IsRemoteZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRollbackRangeTimeResult) *string { return v.IsRemoteZone }).(pulumi.StringPtrOutput)
}

// Returned parameter information.
func (o GetRollbackRangeTimeResultOutput) Items() GetRollbackRangeTimeItemArrayOutput {
	return o.ApplyT(func(v GetRollbackRangeTimeResult) []GetRollbackRangeTimeItem { return v.Items }).(GetRollbackRangeTimeItemArrayOutput)
}

func (o GetRollbackRangeTimeResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRollbackRangeTimeResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRollbackRangeTimeResultOutput{})
}
