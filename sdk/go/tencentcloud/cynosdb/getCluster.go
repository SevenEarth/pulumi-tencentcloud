// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cynosdb cluster
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cynosdb.GetCluster(ctx, &cynosdb.GetClusterArgs{
// 			ClusterId: "cynosdbmysql-bws8h88b",
// 			Database:  pulumi.StringRef("users"),
// 			Table:     pulumi.StringRef("tb_user_name"),
// 			TableType: pulumi.StringRef("all"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("tencentcloud:Cynosdb/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Database name.
	Database *string `pulumi:"database"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Data Table Name.
	Table *string `pulumi:"table"`
	// Data table type: view: only return view, base_ Table: only returns the basic table, all: returns the view and table.
	TableType *string `pulumi:"tableType"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	ClusterId string `pulumi:"clusterId"`
	// Database name note: This field may return null, indicating that a valid value cannot be obtained.
	Database *string `pulumi:"database"`
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Table            *string `pulumi:"table"`
	TableType        *string `pulumi:"tableType"`
	// Table Name List Note: This field may return null, indicating that a valid value cannot be obtained.
	Tables []GetClusterTable `pulumi:"tables"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Database name.
	Database pulumi.StringPtrInput `pulumi:"database"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Data Table Name.
	Table pulumi.StringPtrInput `pulumi:"table"`
	// Data table type: view: only return view, base_ Table: only returns the basic table, all: returns the view and table.
	TableType pulumi.StringPtrInput `pulumi:"tableType"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Database name note: This field may return null, indicating that a valid value cannot be obtained.
func (o LookupClusterResultOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Database }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o LookupClusterResultOutput) TableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupClusterResult) *string { return v.TableType }).(pulumi.StringPtrOutput)
}

// Table Name List Note: This field may return null, indicating that a valid value cannot be obtained.
func (o LookupClusterResultOutput) Tables() GetClusterTableArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterTable { return v.Tables }).(GetClusterTableArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
