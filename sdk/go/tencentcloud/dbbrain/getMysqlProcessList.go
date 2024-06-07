// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbbrain

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this data source to query detailed information of dbbrain mysqlProcessList
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dbbrain"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dbbrain.GetMysqlProcessList(ctx, &dbbrain.GetMysqlProcessListArgs{
//				InstanceId: local.Mysql_id,
//				Product:    pulumi.StringRef("mysql"),
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
func GetMysqlProcessList(ctx *pulumi.Context, args *GetMysqlProcessListArgs, opts ...pulumi.InvokeOption) (*GetMysqlProcessListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMysqlProcessListResult
	err := ctx.Invoke("tencentcloud:Dbbrain/getMysqlProcessList:getMysqlProcessList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMysqlProcessList.
type GetMysqlProcessListArgs struct {
	// The execution type of the thread, used to filter the thread list.
	Command *string `pulumi:"command"`
	// The threads operations database, used to filter the thread list.
	Db *string `pulumi:"db"`
	// The operating host address of the thread, used to filter the thread list.
	Host *string `pulumi:"host"`
	// thread ID, used to filter the thread list.
	Id *int `pulumi:"id"`
	// The threads operation statement is used to filter the thread list.
	Info *string `pulumi:"info"`
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Service product type, supported values: `mysql` - cloud database MySQL; `cynosdb` - cloud database TDSQL-C for MySQL, the default is `mysql`.
	Product *string `pulumi:"product"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// The operational state of the thread, used to filter the thread list.
	State *string `pulumi:"state"`
	// The minimum value of the operation duration of a thread, in seconds, used to filter the list of threads whose operation duration is longer than this value.
	Time *int `pulumi:"time"`
	// The operating account name of the thread, used to filter the thread list.
	User *string `pulumi:"user"`
}

// A collection of values returned by getMysqlProcessList.
type GetMysqlProcessListResult struct {
	// The execution type of the thread.
	Command *string `pulumi:"command"`
	// The thread that operates the database.
	Db *string `pulumi:"db"`
	// The operating host address of the thread.
	Host *string `pulumi:"host"`
	// thread ID.
	Id *int `pulumi:"id"`
	// The operation statement for the thread.
	Info       *string `pulumi:"info"`
	InstanceId string  `pulumi:"instanceId"`
	// Live thread list.
	ProcessLists     []GetMysqlProcessListProcessList `pulumi:"processLists"`
	Product          *string                          `pulumi:"product"`
	ResultOutputFile *string                          `pulumi:"resultOutputFile"`
	// The operational state of the thread.
	State *string `pulumi:"state"`
	// The operation duration of the thread, in seconds.
	Time *int `pulumi:"time"`
	// The operating account name of the thread.
	User *string `pulumi:"user"`
}

func GetMysqlProcessListOutput(ctx *pulumi.Context, args GetMysqlProcessListOutputArgs, opts ...pulumi.InvokeOption) GetMysqlProcessListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMysqlProcessListResult, error) {
			args := v.(GetMysqlProcessListArgs)
			r, err := GetMysqlProcessList(ctx, &args, opts...)
			var s GetMysqlProcessListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMysqlProcessListResultOutput)
}

// A collection of arguments for invoking getMysqlProcessList.
type GetMysqlProcessListOutputArgs struct {
	// The execution type of the thread, used to filter the thread list.
	Command pulumi.StringPtrInput `pulumi:"command"`
	// The threads operations database, used to filter the thread list.
	Db pulumi.StringPtrInput `pulumi:"db"`
	// The operating host address of the thread, used to filter the thread list.
	Host pulumi.StringPtrInput `pulumi:"host"`
	// thread ID, used to filter the thread list.
	Id pulumi.IntPtrInput `pulumi:"id"`
	// The threads operation statement is used to filter the thread list.
	Info pulumi.StringPtrInput `pulumi:"info"`
	// instance id.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// Service product type, supported values: `mysql` - cloud database MySQL; `cynosdb` - cloud database TDSQL-C for MySQL, the default is `mysql`.
	Product pulumi.StringPtrInput `pulumi:"product"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// The operational state of the thread, used to filter the thread list.
	State pulumi.StringPtrInput `pulumi:"state"`
	// The minimum value of the operation duration of a thread, in seconds, used to filter the list of threads whose operation duration is longer than this value.
	Time pulumi.IntPtrInput `pulumi:"time"`
	// The operating account name of the thread, used to filter the thread list.
	User pulumi.StringPtrInput `pulumi:"user"`
}

func (GetMysqlProcessListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlProcessListArgs)(nil)).Elem()
}

// A collection of values returned by getMysqlProcessList.
type GetMysqlProcessListResultOutput struct{ *pulumi.OutputState }

func (GetMysqlProcessListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlProcessListResult)(nil)).Elem()
}

func (o GetMysqlProcessListResultOutput) ToGetMysqlProcessListResultOutput() GetMysqlProcessListResultOutput {
	return o
}

func (o GetMysqlProcessListResultOutput) ToGetMysqlProcessListResultOutputWithContext(ctx context.Context) GetMysqlProcessListResultOutput {
	return o
}

// The execution type of the thread.
func (o GetMysqlProcessListResultOutput) Command() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.Command }).(pulumi.StringPtrOutput)
}

// The thread that operates the database.
func (o GetMysqlProcessListResultOutput) Db() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.Db }).(pulumi.StringPtrOutput)
}

// The operating host address of the thread.
func (o GetMysqlProcessListResultOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.Host }).(pulumi.StringPtrOutput)
}

// thread ID.
func (o GetMysqlProcessListResultOutput) Id() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *int { return v.Id }).(pulumi.IntPtrOutput)
}

// The operation statement for the thread.
func (o GetMysqlProcessListResultOutput) Info() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.Info }).(pulumi.StringPtrOutput)
}

func (o GetMysqlProcessListResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Live thread list.
func (o GetMysqlProcessListResultOutput) ProcessLists() GetMysqlProcessListProcessListArrayOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) []GetMysqlProcessListProcessList { return v.ProcessLists }).(GetMysqlProcessListProcessListArrayOutput)
}

func (o GetMysqlProcessListResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o GetMysqlProcessListResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

// The operational state of the thread.
func (o GetMysqlProcessListResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// The operation duration of the thread, in seconds.
func (o GetMysqlProcessListResultOutput) Time() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *int { return v.Time }).(pulumi.IntPtrOutput)
}

// The operating account name of the thread.
func (o GetMysqlProcessListResultOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlProcessListResult) *string { return v.User }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMysqlProcessListResultOutput{})
}
