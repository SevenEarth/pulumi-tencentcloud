// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dbbrain sqlTemplates
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Dbbrain.GetSqlTemplates(ctx, &dbbrain.GetSqlTemplatesArgs{
// 			InstanceId: "",
// 			Product:    pulumi.StringRef(""),
// 			Schema:     "",
// 			SqlText:    "",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetSqlTemplates(ctx *pulumi.Context, args *GetSqlTemplatesArgs, opts ...pulumi.InvokeOption) (*GetSqlTemplatesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSqlTemplatesResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getSqlTemplates:getSqlTemplates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSqlTemplates.
type GetSqlTemplatesArgs struct {
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
	Product *string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// database name.
	Schema string `pulumi:"schema"`
	// SQL statements.
	SqlText string `pulumi:"sqlText"`
}

// A collection of values returned by getSqlTemplates.
type GetSqlTemplatesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string  `pulumi:"id"`
	InstanceId       string  `pulumi:"instanceId"`
	Product          *string `pulumi:"product"`
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	Schema           string  `pulumi:"schema"`
	// SQL template ID.
	SqlId int `pulumi:"sqlId"`
	// SQL template content.
	SqlTemplate string `pulumi:"sqlTemplate"`
	SqlText     string `pulumi:"sqlText"`
	// sql type.
	SqlType string `pulumi:"sqlType"`
}

func GetSqlTemplatesOutput(ctx *pulumi.Context, args GetSqlTemplatesOutputArgs, opts ...pulumi.InvokeOption) GetSqlTemplatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSqlTemplatesResult, error) {
			args := v.(GetSqlTemplatesArgs)
			r, err := GetSqlTemplates(ctx, &args, opts...)
			var s GetSqlTemplatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSqlTemplatesResultOutput)
}

// A collection of arguments for invoking getSqlTemplates.
type GetSqlTemplatesOutputArgs struct {
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
	Product pulumi.StringPtrInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// database name.
	Schema pulumi.StringInput `pulumi:"schema"`
	// SQL statements.
	SqlText pulumi.StringInput `pulumi:"sqlText"`
}

func (GetSqlTemplatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSqlTemplatesArgs)(nil)).Elem()
}

// A collection of values returned by getSqlTemplates.
type GetSqlTemplatesResultOutput struct{ *pulumi.OutputState }

func (GetSqlTemplatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSqlTemplatesResult)(nil)).Elem()
}

func (o GetSqlTemplatesResultOutput) ToGetSqlTemplatesResultOutput() GetSqlTemplatesResultOutput {
	return o
}

func (o GetSqlTemplatesResultOutput) ToGetSqlTemplatesResultOutputWithContext(ctx context.Context) GetSqlTemplatesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetSqlTemplatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetSqlTemplatesResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetSqlTemplatesResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetSqlTemplatesResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o GetSqlTemplatesResultOutput) Schema() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) string { return v.Schema }).(pulumi.StringOutput)
}

// SQL template ID.
func (o GetSqlTemplatesResultOutput) SqlId() pulumi.IntOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) int { return v.SqlId }).(pulumi.IntOutput)
}

// SQL template content.
func (o GetSqlTemplatesResultOutput) SqlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) string { return v.SqlTemplate }).(pulumi.StringOutput)
}

func (o GetSqlTemplatesResultOutput) SqlText() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) string { return v.SqlText }).(pulumi.StringOutput)
}

// sql type.
func (o GetSqlTemplatesResultOutput) SqlType() pulumi.StringOutput {
	return o.ApplyT(func(v GetSqlTemplatesResult) string { return v.SqlType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSqlTemplatesResultOutput{})
}
