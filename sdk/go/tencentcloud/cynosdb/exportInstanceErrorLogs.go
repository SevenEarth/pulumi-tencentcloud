// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a cynosdb exportInstanceErrorLogs
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cynosdb.NewExportInstanceErrorLogs(ctx, "exportInstanceErrorLogs", &Cynosdb.ExportInstanceErrorLogsArgs{
//				EndTime:    pulumi.String("2022-01-01 14:00:00"),
//				FileType:   pulumi.String("csv"),
//				InstanceId: pulumi.String("cynosdbmysql-ins-afqx1hy0"),
//				KeyWords: pulumi.StringArray{
//					pulumi.String("content"),
//				},
//				LogLevels: pulumi.StringArray{
//					pulumi.String("note"),
//				},
//				OrderBy:     pulumi.String("Timestamp"),
//				OrderByType: pulumi.String("ASC"),
//				StartTime:   pulumi.String("2022-01-01 12:00:00"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
type ExportInstanceErrorLogs struct {
	pulumi.CustomResourceState

	// Latest log time.
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// List of instances in the read-write instance group.
	ErrorLogItemExports ExportInstanceErrorLogsErrorLogItemExportArrayOutput `pulumi:"errorLogItemExports"`
	// File type, optional values: csv, original.
	FileType pulumi.StringPtrOutput `pulumi:"fileType"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// keyword.
	KeyWords pulumi.StringArrayOutput `pulumi:"keyWords"`
	// Log level.
	LogLevels pulumi.StringArrayOutput `pulumi:"logLevels"`
	// Optional value Timestamp.
	OrderBy pulumi.StringPtrOutput `pulumi:"orderBy"`
	// ASC or DESC.
	OrderByType pulumi.StringPtrOutput `pulumi:"orderByType"`
	// Log earliest time.
	StartTime pulumi.StringPtrOutput `pulumi:"startTime"`
}

// NewExportInstanceErrorLogs registers a new resource with the given unique name, arguments, and options.
func NewExportInstanceErrorLogs(ctx *pulumi.Context,
	name string, args *ExportInstanceErrorLogsArgs, opts ...pulumi.ResourceOption) (*ExportInstanceErrorLogs, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExportInstanceErrorLogs
	err := ctx.RegisterResource("tencentcloud:Cynosdb/exportInstanceErrorLogs:ExportInstanceErrorLogs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExportInstanceErrorLogs gets an existing ExportInstanceErrorLogs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExportInstanceErrorLogs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExportInstanceErrorLogsState, opts ...pulumi.ResourceOption) (*ExportInstanceErrorLogs, error) {
	var resource ExportInstanceErrorLogs
	err := ctx.ReadResource("tencentcloud:Cynosdb/exportInstanceErrorLogs:ExportInstanceErrorLogs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExportInstanceErrorLogs resources.
type exportInstanceErrorLogsState struct {
	// Latest log time.
	EndTime *string `pulumi:"endTime"`
	// List of instances in the read-write instance group.
	ErrorLogItemExports []ExportInstanceErrorLogsErrorLogItemExport `pulumi:"errorLogItemExports"`
	// File type, optional values: csv, original.
	FileType *string `pulumi:"fileType"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// keyword.
	KeyWords []string `pulumi:"keyWords"`
	// Log level.
	LogLevels []string `pulumi:"logLevels"`
	// Optional value Timestamp.
	OrderBy *string `pulumi:"orderBy"`
	// ASC or DESC.
	OrderByType *string `pulumi:"orderByType"`
	// Log earliest time.
	StartTime *string `pulumi:"startTime"`
}

type ExportInstanceErrorLogsState struct {
	// Latest log time.
	EndTime pulumi.StringPtrInput
	// List of instances in the read-write instance group.
	ErrorLogItemExports ExportInstanceErrorLogsErrorLogItemExportArrayInput
	// File type, optional values: csv, original.
	FileType pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// keyword.
	KeyWords pulumi.StringArrayInput
	// Log level.
	LogLevels pulumi.StringArrayInput
	// Optional value Timestamp.
	OrderBy pulumi.StringPtrInput
	// ASC or DESC.
	OrderByType pulumi.StringPtrInput
	// Log earliest time.
	StartTime pulumi.StringPtrInput
}

func (ExportInstanceErrorLogsState) ElementType() reflect.Type {
	return reflect.TypeOf((*exportInstanceErrorLogsState)(nil)).Elem()
}

type exportInstanceErrorLogsArgs struct {
	// Latest log time.
	EndTime *string `pulumi:"endTime"`
	// File type, optional values: csv, original.
	FileType *string `pulumi:"fileType"`
	// Instance ID.
	InstanceId string `pulumi:"instanceId"`
	// keyword.
	KeyWords []string `pulumi:"keyWords"`
	// Log level.
	LogLevels []string `pulumi:"logLevels"`
	// Optional value Timestamp.
	OrderBy *string `pulumi:"orderBy"`
	// ASC or DESC.
	OrderByType *string `pulumi:"orderByType"`
	// Log earliest time.
	StartTime *string `pulumi:"startTime"`
}

// The set of arguments for constructing a ExportInstanceErrorLogs resource.
type ExportInstanceErrorLogsArgs struct {
	// Latest log time.
	EndTime pulumi.StringPtrInput
	// File type, optional values: csv, original.
	FileType pulumi.StringPtrInput
	// Instance ID.
	InstanceId pulumi.StringInput
	// keyword.
	KeyWords pulumi.StringArrayInput
	// Log level.
	LogLevels pulumi.StringArrayInput
	// Optional value Timestamp.
	OrderBy pulumi.StringPtrInput
	// ASC or DESC.
	OrderByType pulumi.StringPtrInput
	// Log earliest time.
	StartTime pulumi.StringPtrInput
}

func (ExportInstanceErrorLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*exportInstanceErrorLogsArgs)(nil)).Elem()
}

type ExportInstanceErrorLogsInput interface {
	pulumi.Input

	ToExportInstanceErrorLogsOutput() ExportInstanceErrorLogsOutput
	ToExportInstanceErrorLogsOutputWithContext(ctx context.Context) ExportInstanceErrorLogsOutput
}

func (*ExportInstanceErrorLogs) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportInstanceErrorLogs)(nil)).Elem()
}

func (i *ExportInstanceErrorLogs) ToExportInstanceErrorLogsOutput() ExportInstanceErrorLogsOutput {
	return i.ToExportInstanceErrorLogsOutputWithContext(context.Background())
}

func (i *ExportInstanceErrorLogs) ToExportInstanceErrorLogsOutputWithContext(ctx context.Context) ExportInstanceErrorLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportInstanceErrorLogsOutput)
}

// ExportInstanceErrorLogsArrayInput is an input type that accepts ExportInstanceErrorLogsArray and ExportInstanceErrorLogsArrayOutput values.
// You can construct a concrete instance of `ExportInstanceErrorLogsArrayInput` via:
//
//	ExportInstanceErrorLogsArray{ ExportInstanceErrorLogsArgs{...} }
type ExportInstanceErrorLogsArrayInput interface {
	pulumi.Input

	ToExportInstanceErrorLogsArrayOutput() ExportInstanceErrorLogsArrayOutput
	ToExportInstanceErrorLogsArrayOutputWithContext(context.Context) ExportInstanceErrorLogsArrayOutput
}

type ExportInstanceErrorLogsArray []ExportInstanceErrorLogsInput

func (ExportInstanceErrorLogsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExportInstanceErrorLogs)(nil)).Elem()
}

func (i ExportInstanceErrorLogsArray) ToExportInstanceErrorLogsArrayOutput() ExportInstanceErrorLogsArrayOutput {
	return i.ToExportInstanceErrorLogsArrayOutputWithContext(context.Background())
}

func (i ExportInstanceErrorLogsArray) ToExportInstanceErrorLogsArrayOutputWithContext(ctx context.Context) ExportInstanceErrorLogsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportInstanceErrorLogsArrayOutput)
}

// ExportInstanceErrorLogsMapInput is an input type that accepts ExportInstanceErrorLogsMap and ExportInstanceErrorLogsMapOutput values.
// You can construct a concrete instance of `ExportInstanceErrorLogsMapInput` via:
//
//	ExportInstanceErrorLogsMap{ "key": ExportInstanceErrorLogsArgs{...} }
type ExportInstanceErrorLogsMapInput interface {
	pulumi.Input

	ToExportInstanceErrorLogsMapOutput() ExportInstanceErrorLogsMapOutput
	ToExportInstanceErrorLogsMapOutputWithContext(context.Context) ExportInstanceErrorLogsMapOutput
}

type ExportInstanceErrorLogsMap map[string]ExportInstanceErrorLogsInput

func (ExportInstanceErrorLogsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExportInstanceErrorLogs)(nil)).Elem()
}

func (i ExportInstanceErrorLogsMap) ToExportInstanceErrorLogsMapOutput() ExportInstanceErrorLogsMapOutput {
	return i.ToExportInstanceErrorLogsMapOutputWithContext(context.Background())
}

func (i ExportInstanceErrorLogsMap) ToExportInstanceErrorLogsMapOutputWithContext(ctx context.Context) ExportInstanceErrorLogsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportInstanceErrorLogsMapOutput)
}

type ExportInstanceErrorLogsOutput struct{ *pulumi.OutputState }

func (ExportInstanceErrorLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportInstanceErrorLogs)(nil)).Elem()
}

func (o ExportInstanceErrorLogsOutput) ToExportInstanceErrorLogsOutput() ExportInstanceErrorLogsOutput {
	return o
}

func (o ExportInstanceErrorLogsOutput) ToExportInstanceErrorLogsOutputWithContext(ctx context.Context) ExportInstanceErrorLogsOutput {
	return o
}

// Latest log time.
func (o ExportInstanceErrorLogsOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringPtrOutput { return v.EndTime }).(pulumi.StringPtrOutput)
}

// List of instances in the read-write instance group.
func (o ExportInstanceErrorLogsOutput) ErrorLogItemExports() ExportInstanceErrorLogsErrorLogItemExportArrayOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) ExportInstanceErrorLogsErrorLogItemExportArrayOutput {
		return v.ErrorLogItemExports
	}).(ExportInstanceErrorLogsErrorLogItemExportArrayOutput)
}

// File type, optional values: csv, original.
func (o ExportInstanceErrorLogsOutput) FileType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringPtrOutput { return v.FileType }).(pulumi.StringPtrOutput)
}

// Instance ID.
func (o ExportInstanceErrorLogsOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// keyword.
func (o ExportInstanceErrorLogsOutput) KeyWords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringArrayOutput { return v.KeyWords }).(pulumi.StringArrayOutput)
}

// Log level.
func (o ExportInstanceErrorLogsOutput) LogLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringArrayOutput { return v.LogLevels }).(pulumi.StringArrayOutput)
}

// Optional value Timestamp.
func (o ExportInstanceErrorLogsOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringPtrOutput { return v.OrderBy }).(pulumi.StringPtrOutput)
}

// ASC or DESC.
func (o ExportInstanceErrorLogsOutput) OrderByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringPtrOutput { return v.OrderByType }).(pulumi.StringPtrOutput)
}

// Log earliest time.
func (o ExportInstanceErrorLogsOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportInstanceErrorLogs) pulumi.StringPtrOutput { return v.StartTime }).(pulumi.StringPtrOutput)
}

type ExportInstanceErrorLogsArrayOutput struct{ *pulumi.OutputState }

func (ExportInstanceErrorLogsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExportInstanceErrorLogs)(nil)).Elem()
}

func (o ExportInstanceErrorLogsArrayOutput) ToExportInstanceErrorLogsArrayOutput() ExportInstanceErrorLogsArrayOutput {
	return o
}

func (o ExportInstanceErrorLogsArrayOutput) ToExportInstanceErrorLogsArrayOutputWithContext(ctx context.Context) ExportInstanceErrorLogsArrayOutput {
	return o
}

func (o ExportInstanceErrorLogsArrayOutput) Index(i pulumi.IntInput) ExportInstanceErrorLogsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExportInstanceErrorLogs {
		return vs[0].([]*ExportInstanceErrorLogs)[vs[1].(int)]
	}).(ExportInstanceErrorLogsOutput)
}

type ExportInstanceErrorLogsMapOutput struct{ *pulumi.OutputState }

func (ExportInstanceErrorLogsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExportInstanceErrorLogs)(nil)).Elem()
}

func (o ExportInstanceErrorLogsMapOutput) ToExportInstanceErrorLogsMapOutput() ExportInstanceErrorLogsMapOutput {
	return o
}

func (o ExportInstanceErrorLogsMapOutput) ToExportInstanceErrorLogsMapOutputWithContext(ctx context.Context) ExportInstanceErrorLogsMapOutput {
	return o
}

func (o ExportInstanceErrorLogsMapOutput) MapIndex(k pulumi.StringInput) ExportInstanceErrorLogsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExportInstanceErrorLogs {
		return vs[0].(map[string]*ExportInstanceErrorLogs)[vs[1].(string)]
	}).(ExportInstanceErrorLogsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExportInstanceErrorLogsInput)(nil)).Elem(), &ExportInstanceErrorLogs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportInstanceErrorLogsArrayInput)(nil)).Elem(), ExportInstanceErrorLogsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExportInstanceErrorLogsMapInput)(nil)).Elem(), ExportInstanceErrorLogsMap{})
	pulumi.RegisterOutputType(ExportInstanceErrorLogsOutput{})
	pulumi.RegisterOutputType(ExportInstanceErrorLogsArrayOutput{})
	pulumi.RegisterOutputType(ExportInstanceErrorLogsMapOutput{})
}
