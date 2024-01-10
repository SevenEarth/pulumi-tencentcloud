// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cynosdb auditLogFile
//
// ## Example Usage
//
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
//			_, err := Cynosdb.NewAuditLogFile(ctx, "auditLogFile", &Cynosdb.AuditLogFileArgs{
//				EndTime:    pulumi.String("2022-08-12 10:29:20"),
//				InstanceId: pulumi.String("cynosdbmysql-ins-afqx1hy0"),
//				StartTime:  pulumi.String("2022-07-12 10:29:20"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AuditLogFile struct {
	pulumi.CustomResourceState

	// Audit log file creation time. The format is 2019-03-20 17:09:13.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The download address of the audit logs.
	DownloadUrl pulumi.StringOutput `pulumi:"downloadUrl"`
	// End time.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Error message.
	ErrMsg pulumi.StringOutput `pulumi:"errMsg"`
	// Audit log file name.
	FileName pulumi.StringOutput `pulumi:"fileName"`
	// File size, The unit is KB.
	FileSize pulumi.IntOutput `pulumi:"fileSize"`
	// Filter condition. Logs can be filtered according to the filter conditions set.
	Filter AuditLogFileFilterPtrOutput `pulumi:"filter"`
	// The ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Sort by. Supported values are: `ASC` - ascending, `DESC` - descending.
	Order pulumi.StringPtrOutput `pulumi:"order"`
	// Sort field. supported values are:
	// `timestamp` - timestamp
	// `affectRows` - affected rows
	// `execTime` - execution time.
	OrderBy pulumi.StringPtrOutput `pulumi:"orderBy"`
	// Start time.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewAuditLogFile registers a new resource with the given unique name, arguments, and options.
func NewAuditLogFile(ctx *pulumi.Context,
	name string, args *AuditLogFileArgs, opts ...pulumi.ResourceOption) (*AuditLogFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndTime == nil {
		return nil, errors.New("invalid value for required argument 'EndTime'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AuditLogFile
	err := ctx.RegisterResource("tencentcloud:Cynosdb/auditLogFile:AuditLogFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuditLogFile gets an existing AuditLogFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuditLogFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditLogFileState, opts ...pulumi.ResourceOption) (*AuditLogFile, error) {
	var resource AuditLogFile
	err := ctx.ReadResource("tencentcloud:Cynosdb/auditLogFile:AuditLogFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuditLogFile resources.
type auditLogFileState struct {
	// Audit log file creation time. The format is 2019-03-20 17:09:13.
	CreateTime *string `pulumi:"createTime"`
	// The download address of the audit logs.
	DownloadUrl *string `pulumi:"downloadUrl"`
	// End time.
	EndTime *string `pulumi:"endTime"`
	// Error message.
	ErrMsg *string `pulumi:"errMsg"`
	// Audit log file name.
	FileName *string `pulumi:"fileName"`
	// File size, The unit is KB.
	FileSize *int `pulumi:"fileSize"`
	// Filter condition. Logs can be filtered according to the filter conditions set.
	Filter *AuditLogFileFilter `pulumi:"filter"`
	// The ID of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Sort by. Supported values are: `ASC` - ascending, `DESC` - descending.
	Order *string `pulumi:"order"`
	// Sort field. supported values are:
	// `timestamp` - timestamp
	// `affectRows` - affected rows
	// `execTime` - execution time.
	OrderBy *string `pulumi:"orderBy"`
	// Start time.
	StartTime *string `pulumi:"startTime"`
}

type AuditLogFileState struct {
	// Audit log file creation time. The format is 2019-03-20 17:09:13.
	CreateTime pulumi.StringPtrInput
	// The download address of the audit logs.
	DownloadUrl pulumi.StringPtrInput
	// End time.
	EndTime pulumi.StringPtrInput
	// Error message.
	ErrMsg pulumi.StringPtrInput
	// Audit log file name.
	FileName pulumi.StringPtrInput
	// File size, The unit is KB.
	FileSize pulumi.IntPtrInput
	// Filter condition. Logs can be filtered according to the filter conditions set.
	Filter AuditLogFileFilterPtrInput
	// The ID of instance.
	InstanceId pulumi.StringPtrInput
	// Sort by. Supported values are: `ASC` - ascending, `DESC` - descending.
	Order pulumi.StringPtrInput
	// Sort field. supported values are:
	// `timestamp` - timestamp
	// `affectRows` - affected rows
	// `execTime` - execution time.
	OrderBy pulumi.StringPtrInput
	// Start time.
	StartTime pulumi.StringPtrInput
}

func (AuditLogFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditLogFileState)(nil)).Elem()
}

type auditLogFileArgs struct {
	// End time.
	EndTime string `pulumi:"endTime"`
	// Filter condition. Logs can be filtered according to the filter conditions set.
	Filter *AuditLogFileFilter `pulumi:"filter"`
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// Sort by. Supported values are: `ASC` - ascending, `DESC` - descending.
	Order *string `pulumi:"order"`
	// Sort field. supported values are:
	// `timestamp` - timestamp
	// `affectRows` - affected rows
	// `execTime` - execution time.
	OrderBy *string `pulumi:"orderBy"`
	// Start time.
	StartTime string `pulumi:"startTime"`
}

// The set of arguments for constructing a AuditLogFile resource.
type AuditLogFileArgs struct {
	// End time.
	EndTime pulumi.StringInput
	// Filter condition. Logs can be filtered according to the filter conditions set.
	Filter AuditLogFileFilterPtrInput
	// The ID of instance.
	InstanceId pulumi.StringInput
	// Sort by. Supported values are: `ASC` - ascending, `DESC` - descending.
	Order pulumi.StringPtrInput
	// Sort field. supported values are:
	// `timestamp` - timestamp
	// `affectRows` - affected rows
	// `execTime` - execution time.
	OrderBy pulumi.StringPtrInput
	// Start time.
	StartTime pulumi.StringInput
}

func (AuditLogFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditLogFileArgs)(nil)).Elem()
}

type AuditLogFileInput interface {
	pulumi.Input

	ToAuditLogFileOutput() AuditLogFileOutput
	ToAuditLogFileOutputWithContext(ctx context.Context) AuditLogFileOutput
}

func (*AuditLogFile) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogFile)(nil)).Elem()
}

func (i *AuditLogFile) ToAuditLogFileOutput() AuditLogFileOutput {
	return i.ToAuditLogFileOutputWithContext(context.Background())
}

func (i *AuditLogFile) ToAuditLogFileOutputWithContext(ctx context.Context) AuditLogFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditLogFileOutput)
}

// AuditLogFileArrayInput is an input type that accepts AuditLogFileArray and AuditLogFileArrayOutput values.
// You can construct a concrete instance of `AuditLogFileArrayInput` via:
//
//	AuditLogFileArray{ AuditLogFileArgs{...} }
type AuditLogFileArrayInput interface {
	pulumi.Input

	ToAuditLogFileArrayOutput() AuditLogFileArrayOutput
	ToAuditLogFileArrayOutputWithContext(context.Context) AuditLogFileArrayOutput
}

type AuditLogFileArray []AuditLogFileInput

func (AuditLogFileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditLogFile)(nil)).Elem()
}

func (i AuditLogFileArray) ToAuditLogFileArrayOutput() AuditLogFileArrayOutput {
	return i.ToAuditLogFileArrayOutputWithContext(context.Background())
}

func (i AuditLogFileArray) ToAuditLogFileArrayOutputWithContext(ctx context.Context) AuditLogFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditLogFileArrayOutput)
}

// AuditLogFileMapInput is an input type that accepts AuditLogFileMap and AuditLogFileMapOutput values.
// You can construct a concrete instance of `AuditLogFileMapInput` via:
//
//	AuditLogFileMap{ "key": AuditLogFileArgs{...} }
type AuditLogFileMapInput interface {
	pulumi.Input

	ToAuditLogFileMapOutput() AuditLogFileMapOutput
	ToAuditLogFileMapOutputWithContext(context.Context) AuditLogFileMapOutput
}

type AuditLogFileMap map[string]AuditLogFileInput

func (AuditLogFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditLogFile)(nil)).Elem()
}

func (i AuditLogFileMap) ToAuditLogFileMapOutput() AuditLogFileMapOutput {
	return i.ToAuditLogFileMapOutputWithContext(context.Background())
}

func (i AuditLogFileMap) ToAuditLogFileMapOutputWithContext(ctx context.Context) AuditLogFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditLogFileMapOutput)
}

type AuditLogFileOutput struct{ *pulumi.OutputState }

func (AuditLogFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogFile)(nil)).Elem()
}

func (o AuditLogFileOutput) ToAuditLogFileOutput() AuditLogFileOutput {
	return o
}

func (o AuditLogFileOutput) ToAuditLogFileOutputWithContext(ctx context.Context) AuditLogFileOutput {
	return o
}

// Audit log file creation time. The format is 2019-03-20 17:09:13.
func (o AuditLogFileOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The download address of the audit logs.
func (o AuditLogFileOutput) DownloadUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringOutput { return v.DownloadUrl }).(pulumi.StringOutput)
}

// End time.
func (o AuditLogFileOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Error message.
func (o AuditLogFileOutput) ErrMsg() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringOutput { return v.ErrMsg }).(pulumi.StringOutput)
}

// Audit log file name.
func (o AuditLogFileOutput) FileName() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringOutput { return v.FileName }).(pulumi.StringOutput)
}

// File size, The unit is KB.
func (o AuditLogFileOutput) FileSize() pulumi.IntOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.IntOutput { return v.FileSize }).(pulumi.IntOutput)
}

// Filter condition. Logs can be filtered according to the filter conditions set.
func (o AuditLogFileOutput) Filter() AuditLogFileFilterPtrOutput {
	return o.ApplyT(func(v *AuditLogFile) AuditLogFileFilterPtrOutput { return v.Filter }).(AuditLogFileFilterPtrOutput)
}

// The ID of instance.
func (o AuditLogFileOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Sort by. Supported values are: `ASC` - ascending, `DESC` - descending.
func (o AuditLogFileOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringPtrOutput { return v.Order }).(pulumi.StringPtrOutput)
}

// Sort field. supported values are:
// `timestamp` - timestamp
// `affectRows` - affected rows
// `execTime` - execution time.
func (o AuditLogFileOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringPtrOutput { return v.OrderBy }).(pulumi.StringPtrOutput)
}

// Start time.
func (o AuditLogFileOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AuditLogFile) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

type AuditLogFileArrayOutput struct{ *pulumi.OutputState }

func (AuditLogFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuditLogFile)(nil)).Elem()
}

func (o AuditLogFileArrayOutput) ToAuditLogFileArrayOutput() AuditLogFileArrayOutput {
	return o
}

func (o AuditLogFileArrayOutput) ToAuditLogFileArrayOutputWithContext(ctx context.Context) AuditLogFileArrayOutput {
	return o
}

func (o AuditLogFileArrayOutput) Index(i pulumi.IntInput) AuditLogFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuditLogFile {
		return vs[0].([]*AuditLogFile)[vs[1].(int)]
	}).(AuditLogFileOutput)
}

type AuditLogFileMapOutput struct{ *pulumi.OutputState }

func (AuditLogFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuditLogFile)(nil)).Elem()
}

func (o AuditLogFileMapOutput) ToAuditLogFileMapOutput() AuditLogFileMapOutput {
	return o
}

func (o AuditLogFileMapOutput) ToAuditLogFileMapOutputWithContext(ctx context.Context) AuditLogFileMapOutput {
	return o
}

func (o AuditLogFileMapOutput) MapIndex(k pulumi.StringInput) AuditLogFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuditLogFile {
		return vs[0].(map[string]*AuditLogFile)[vs[1].(string)]
	}).(AuditLogFileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogFileInput)(nil)).Elem(), &AuditLogFile{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogFileArrayInput)(nil)).Elem(), AuditLogFileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogFileMapInput)(nil)).Elem(), AuditLogFileMap{})
	pulumi.RegisterOutputType(AuditLogFileOutput{})
	pulumi.RegisterOutputType(AuditLogFileArrayOutput{})
	pulumi.RegisterOutputType(AuditLogFileMapOutput{})
}
