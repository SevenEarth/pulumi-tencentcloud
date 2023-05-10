// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dcdb databaseObjects
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dcdb.GetDatabaseObjects(ctx, &dcdb.GetDatabaseObjectsArgs{
// 			DbName: "",
// 			Gt: []map[string]interface{}{
// 				nil,
// 			},
// 			InstanceId: "dcdbt-ow7t8lmc",
// 			Lt: []map[string]interface{}{
// 				nil,
// 			},
// 			Nil: []map[string]interface{}{
// 				nil,
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDatabaseObjects(ctx *pulumi.Context, args *GetDatabaseObjectsArgs, opts ...pulumi.InvokeOption) (*GetDatabaseObjectsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabaseObjectsResult
	err := ctx.Invoke("tencentcloud:Dcdb/getDatabaseObjects:getDatabaseObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseObjects.
type GetDatabaseObjectsArgs struct {
	// Database name, obtained through the DescribeDatabases api.
	DbName string `pulumi:"dbName"`
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getDatabaseObjects.
type GetDatabaseObjectsResult struct {
	DbName string `pulumi:"dbName"`
	// Function list.
	Funcs []GetDatabaseObjectsFunc `pulumi:"funcs"`
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Procedure list.
	Procs            []GetDatabaseObjectsProc `pulumi:"procs"`
	ResultOutputFile *string                  `pulumi:"resultOutputFile"`
	// Table list.
	Tables []GetDatabaseObjectsTable `pulumi:"tables"`
	// View list.
	Views []GetDatabaseObjectsView `pulumi:"views"`
}

func GetDatabaseObjectsOutput(ctx *pulumi.Context, args GetDatabaseObjectsOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseObjectsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseObjectsResult, error) {
			args := v.(GetDatabaseObjectsArgs)
			r, err := GetDatabaseObjects(ctx, &args, opts...)
			var s GetDatabaseObjectsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseObjectsResultOutput)
}

// A collection of arguments for invoking getDatabaseObjects.
type GetDatabaseObjectsOutputArgs struct {
	// Database name, obtained through the DescribeDatabases api.
	DbName pulumi.StringInput `pulumi:"dbName"`
	// The ID of instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetDatabaseObjectsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseObjectsArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseObjects.
type GetDatabaseObjectsResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseObjectsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseObjectsResult)(nil)).Elem()
}

func (o GetDatabaseObjectsResultOutput) ToGetDatabaseObjectsResultOutput() GetDatabaseObjectsResultOutput {
	return o
}

func (o GetDatabaseObjectsResultOutput) ToGetDatabaseObjectsResultOutputWithContext(ctx context.Context) GetDatabaseObjectsResultOutput {
	return o
}

func (o GetDatabaseObjectsResultOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) string { return v.DbName }).(pulumi.StringOutput)
}

// Function list.
func (o GetDatabaseObjectsResultOutput) Funcs() GetDatabaseObjectsFuncArrayOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) []GetDatabaseObjectsFunc { return v.Funcs }).(GetDatabaseObjectsFuncArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabaseObjectsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatabaseObjectsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Procedure list.
func (o GetDatabaseObjectsResultOutput) Procs() GetDatabaseObjectsProcArrayOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) []GetDatabaseObjectsProc { return v.Procs }).(GetDatabaseObjectsProcArrayOutput)
}

func (o GetDatabaseObjectsResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// Table list.
func (o GetDatabaseObjectsResultOutput) Tables() GetDatabaseObjectsTableArrayOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) []GetDatabaseObjectsTable { return v.Tables }).(GetDatabaseObjectsTableArrayOutput)
}

// View list.
func (o GetDatabaseObjectsResultOutput) Views() GetDatabaseObjectsViewArrayOutput {
	return o.ApplyT(func(v GetDatabaseObjectsResult) []GetDatabaseObjectsView { return v.Views }).(GetDatabaseObjectsViewArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseObjectsResultOutput{})
}
