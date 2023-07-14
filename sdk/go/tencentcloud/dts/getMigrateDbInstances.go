// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dts migrateDbInstances
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dts"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dts.GetMigrateDbInstances(ctx, &dts.GetMigrateDbInstancesArgs{
// 			AccountMode:  pulumi.StringRef("self"),
// 			DatabaseType: "mysql",
// 			InstanceId:   pulumi.StringRef("cdb-ffulb2sg"),
// 			InstanceName: pulumi.StringRef("cdb_test"),
// 			Limit:        pulumi.IntRef(10),
// 			MigrateRole:  pulumi.StringRef("src"),
// 			Offset:       pulumi.IntRef(10),
// 			TmpSecretId:  pulumi.StringRef("AKIDvBDyVmna9TadcS4YzfBZmkU5TbX12345"),
// 			TmpSecretKey: pulumi.StringRef("ZswjGWWHm24qMeiX6QUJsELDpC12345"),
// 			TmpToken:     pulumi.StringRef("JOqqCPVuWdNZvlVDLxxx"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetMigrateDbInstances(ctx *pulumi.Context, args *GetMigrateDbInstancesArgs, opts ...pulumi.InvokeOption) (*GetMigrateDbInstancesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMigrateDbInstancesResult
	err := ctx.Invoke("tencentcloud:Dts/getMigrateDbInstances:getMigrateDbInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMigrateDbInstances.
type GetMigrateDbInstancesArgs struct {
	// The owning account of the resource is null or self(resources in the self account), other(resources in the other account).
	AccountMode *string `pulumi:"accountMode"`
	// Database type.
	DatabaseType string `pulumi:"databaseType"`
	// Database instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Database instance name.
	InstanceName *string `pulumi:"instanceName"`
	// Limit.
	Limit *int `pulumi:"limit"`
	// Whether the instance is the migration source or destination,src(for source), dst(for destination).
	MigrateRole *string `pulumi:"migrateRole"`
	// Offset.
	Offset *int `pulumi:"offset"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// temporary secret id, used across account.
	TmpSecretId *string `pulumi:"tmpSecretId"`
	// temporary secret key, used across account.
	TmpSecretKey *string `pulumi:"tmpSecretKey"`
	// temporary token, used across account.
	TmpToken *string `pulumi:"tmpToken"`
}

// A collection of values returned by getMigrateDbInstances.
type GetMigrateDbInstancesResult struct {
	AccountMode  *string `pulumi:"accountMode"`
	DatabaseType string  `pulumi:"databaseType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance Id.
	InstanceId *string `pulumi:"instanceId"`
	// Database instance name.
	InstanceName *string `pulumi:"instanceName"`
	// Instance list.
	Instances   []GetMigrateDbInstancesInstance `pulumi:"instances"`
	Limit       *int                            `pulumi:"limit"`
	MigrateRole *string                         `pulumi:"migrateRole"`
	Offset      *int                            `pulumi:"offset"`
	// Unique request id, provide this when encounter a problem.
	RequestId        string  `pulumi:"requestId"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	TmpSecretId      *string `pulumi:"tmpSecretId"`
	TmpSecretKey     *string `pulumi:"tmpSecretKey"`
	TmpToken         *string `pulumi:"tmpToken"`
}

func GetMigrateDbInstancesOutput(ctx *pulumi.Context, args GetMigrateDbInstancesOutputArgs, opts ...pulumi.InvokeOption) GetMigrateDbInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMigrateDbInstancesResult, error) {
			args := v.(GetMigrateDbInstancesArgs)
			r, err := GetMigrateDbInstances(ctx, &args, opts...)
			var s GetMigrateDbInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMigrateDbInstancesResultOutput)
}

// A collection of arguments for invoking getMigrateDbInstances.
type GetMigrateDbInstancesOutputArgs struct {
	// The owning account of the resource is null or self(resources in the self account), other(resources in the other account).
	AccountMode pulumi.StringPtrInput `pulumi:"accountMode"`
	// Database type.
	DatabaseType pulumi.StringInput `pulumi:"databaseType"`
	// Database instance id.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Database instance name.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	// Limit.
	Limit pulumi.IntPtrInput `pulumi:"limit"`
	// Whether the instance is the migration source or destination,src(for source), dst(for destination).
	MigrateRole pulumi.StringPtrInput `pulumi:"migrateRole"`
	// Offset.
	Offset pulumi.IntPtrInput `pulumi:"offset"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// temporary secret id, used across account.
	TmpSecretId pulumi.StringPtrInput `pulumi:"tmpSecretId"`
	// temporary secret key, used across account.
	TmpSecretKey pulumi.StringPtrInput `pulumi:"tmpSecretKey"`
	// temporary token, used across account.
	TmpToken pulumi.StringPtrInput `pulumi:"tmpToken"`
}

func (GetMigrateDbInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMigrateDbInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getMigrateDbInstances.
type GetMigrateDbInstancesResultOutput struct{ *pulumi.OutputState }

func (GetMigrateDbInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMigrateDbInstancesResult)(nil)).Elem()
}

func (o GetMigrateDbInstancesResultOutput) ToGetMigrateDbInstancesResultOutput() GetMigrateDbInstancesResultOutput {
	return o
}

func (o GetMigrateDbInstancesResultOutput) ToGetMigrateDbInstancesResultOutputWithContext(ctx context.Context) GetMigrateDbInstancesResultOutput {
	return o
}

func (o GetMigrateDbInstancesResultOutput) AccountMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.AccountMode }).(pulumi.StringPtrOutput)
}

func (o GetMigrateDbInstancesResultOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) string { return v.DatabaseType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMigrateDbInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance Id.
func (o GetMigrateDbInstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// Database instance name.
func (o GetMigrateDbInstancesResultOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// Instance list.
func (o GetMigrateDbInstancesResultOutput) Instances() GetMigrateDbInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) []GetMigrateDbInstancesInstance { return v.Instances }).(GetMigrateDbInstancesInstanceArrayOutput)
}

func (o GetMigrateDbInstancesResultOutput) Limit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *int { return v.Limit }).(pulumi.IntPtrOutput)
}

func (o GetMigrateDbInstancesResultOutput) MigrateRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.MigrateRole }).(pulumi.StringPtrOutput)
}

func (o GetMigrateDbInstancesResultOutput) Offset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *int { return v.Offset }).(pulumi.IntPtrOutput)
}

// Unique request id, provide this when encounter a problem.
func (o GetMigrateDbInstancesResultOutput) RequestId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) string { return v.RequestId }).(pulumi.StringOutput)
}

func (o GetMigrateDbInstancesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetMigrateDbInstancesResultOutput) TmpSecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.TmpSecretId }).(pulumi.StringPtrOutput)
}

func (o GetMigrateDbInstancesResultOutput) TmpSecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.TmpSecretKey }).(pulumi.StringPtrOutput)
}

func (o GetMigrateDbInstancesResultOutput) TmpToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMigrateDbInstancesResult) *string { return v.TmpToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMigrateDbInstancesResultOutput{})
}
