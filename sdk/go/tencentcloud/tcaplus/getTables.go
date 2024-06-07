// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcaplus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query TcaplusDB tables.
func GetTables(ctx *pulumi.Context, args *GetTablesArgs, opts ...pulumi.InvokeOption) (*GetTablesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTablesResult
	err := ctx.Invoke("tencentcloud:Tcaplus/getTables:getTables", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTables.
type GetTablesArgs struct {
	// ID of the TcaplusDB cluster to be query.
	ClusterId string `pulumi:"clusterId"`
	// File for saving results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Table ID to be query.
	TableId *string `pulumi:"tableId"`
	// Table name to be query.
	TableName *string `pulumi:"tableName"`
	// ID of the table group to be query.
	TablegroupId *string `pulumi:"tablegroupId"`
}

// A collection of values returned by getTables.
type GetTablesResult struct {
	ClusterId string `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of TcaplusDB tables. Each element contains the following attributes.
	Lists            []GetTablesList `pulumi:"lists"`
	ResultOutputFile *string         `pulumi:"resultOutputFile"`
	// ID of the TcaplusDB table.
	TableId *string `pulumi:"tableId"`
	// Name of the TcaplusDB table.
	TableName *string `pulumi:"tableName"`
	// Table group id of the TcaplusDB table.
	TablegroupId *string `pulumi:"tablegroupId"`
}

func GetTablesOutput(ctx *pulumi.Context, args GetTablesOutputArgs, opts ...pulumi.InvokeOption) GetTablesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTablesResult, error) {
			args := v.(GetTablesArgs)
			r, err := GetTables(ctx, &args, opts...)
			var s GetTablesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTablesResultOutput)
}

// A collection of arguments for invoking getTables.
type GetTablesOutputArgs struct {
	// ID of the TcaplusDB cluster to be query.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// File for saving results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Table ID to be query.
	TableId pulumi.StringPtrInput `pulumi:"tableId"`
	// Table name to be query.
	TableName pulumi.StringPtrInput `pulumi:"tableName"`
	// ID of the table group to be query.
	TablegroupId pulumi.StringPtrInput `pulumi:"tablegroupId"`
}

func (GetTablesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTablesArgs)(nil)).Elem()
}

// A collection of values returned by getTables.
type GetTablesResultOutput struct{ *pulumi.OutputState }

func (GetTablesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTablesResult)(nil)).Elem()
}

func (o GetTablesResultOutput) ToGetTablesResultOutput() GetTablesResultOutput {
	return o
}

func (o GetTablesResultOutput) ToGetTablesResultOutputWithContext(ctx context.Context) GetTablesResultOutput {
	return o
}

func (o GetTablesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTablesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetTablesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTablesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of TcaplusDB tables. Each element contains the following attributes.
func (o GetTablesResultOutput) Lists() GetTablesListArrayOutput {
	return o.ApplyT(func(v GetTablesResult) []GetTablesList { return v.Lists }).(GetTablesListArrayOutput)
}

func (o GetTablesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTablesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// ID of the TcaplusDB table.
func (o GetTablesResultOutput) TableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTablesResult) *string { return v.TableId }).(pulumi.StringPtrOutput)
}

// Name of the TcaplusDB table.
func (o GetTablesResultOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTablesResult) *string { return v.TableName }).(pulumi.StringPtrOutput)
}

// Table group id of the TcaplusDB table.
func (o GetTablesResultOutput) TablegroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetTablesResult) *string { return v.TablegroupId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTablesResultOutput{})
}
