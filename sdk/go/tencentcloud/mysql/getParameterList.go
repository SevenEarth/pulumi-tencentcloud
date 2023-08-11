// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a parameter group of a database instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mysql"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mysql.GetParameterList(ctx, &mysql.GetParameterListArgs{
// 			EngineVersion:    pulumi.StringRef("5.5"),
// 			MysqlId:          pulumi.StringRef("terraform-test-local-database"),
// 			ResultOutputFile: pulumi.StringRef("mytestpath"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetParameterList(ctx *pulumi.Context, args *GetParameterListArgs, opts ...pulumi.InvokeOption) (*GetParameterListResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetParameterListResult
	err := ctx.Invoke("tencentcloud:Mysql/getParameterList:getParameterList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getParameterList.
type GetParameterListArgs struct {
	// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0.
	EngineVersion *string `pulumi:"engineVersion"`
	// Instance ID.
	MysqlId *string `pulumi:"mysqlId"`
	// Used to store results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
}

// A collection of values returned by getParameterList.
type GetParameterListResult struct {
	EngineVersion *string `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	MysqlId *string `pulumi:"mysqlId"`
	// A list of parameters. Each element contains the following attributes:
	ParameterLists   []GetParameterListParameterList `pulumi:"parameterLists"`
	ResultOutputFile *string                         `pulumi:"resultOutputFile"`
}

func GetParameterListOutput(ctx *pulumi.Context, args GetParameterListOutputArgs, opts ...pulumi.InvokeOption) GetParameterListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetParameterListResult, error) {
			args := v.(GetParameterListArgs)
			r, err := GetParameterList(ctx, &args, opts...)
			var s GetParameterListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetParameterListResultOutput)
}

// A collection of arguments for invoking getParameterList.
type GetParameterListOutputArgs struct {
	// The version number of the database engine to use. Supported versions include 5.5/5.6/5.7/8.0.
	EngineVersion pulumi.StringPtrInput `pulumi:"engineVersion"`
	// Instance ID.
	MysqlId pulumi.StringPtrInput `pulumi:"mysqlId"`
	// Used to store results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
}

func (GetParameterListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetParameterListArgs)(nil)).Elem()
}

// A collection of values returned by getParameterList.
type GetParameterListResultOutput struct{ *pulumi.OutputState }

func (GetParameterListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetParameterListResult)(nil)).Elem()
}

func (o GetParameterListResultOutput) ToGetParameterListResultOutput() GetParameterListResultOutput {
	return o
}

func (o GetParameterListResultOutput) ToGetParameterListResultOutputWithContext(ctx context.Context) GetParameterListResultOutput {
	return o
}

func (o GetParameterListResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterListResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetParameterListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetParameterListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetParameterListResultOutput) MysqlId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterListResult) *string { return v.MysqlId }).(pulumi.StringPtrOutput)
}

// A list of parameters. Each element contains the following attributes:
func (o GetParameterListResultOutput) ParameterLists() GetParameterListParameterListArrayOutput {
	return o.ApplyT(func(v GetParameterListResult) []GetParameterListParameterList { return v.ParameterLists }).(GetParameterListParameterListArrayOutput)
}

func (o GetParameterListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetParameterListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetParameterListResultOutput{})
}
