// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a redis backupDownloadRestriction
//
// ## Example Usage
// ### Modify the network information and address of the current region backup file download
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Redis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Redis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Redis.NewBackupDownloadRestriction(ctx, "foo", &Redis.BackupDownloadRestrictionArgs{
//				LimitType:           pulumi.String("Customize"),
//				VpcComparisonSymbol: pulumi.String("In"),
//				IpComparisonSymbol:  pulumi.String("In"),
//				LimitVpcs: redis.BackupDownloadRestrictionLimitVpcArray{
//					&redis.BackupDownloadRestrictionLimitVpcArgs{
//						Region: pulumi.String("ap-guangzhou"),
//						VpcLists: pulumi.StringArray{
//							pulumi.Any(_var.Vpc_id),
//						},
//					},
//				},
//				LimitIps: pulumi.StringArray{
//					pulumi.String("10.1.1.12"),
//					pulumi.String("10.1.1.13"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// redis backup_download_restriction can be imported using the region, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction foo ap-guangzhou
//
// ```
type BackupDownloadRestriction struct {
	pulumi.CustomResourceState

	// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
	IpComparisonSymbol pulumi.StringPtrOutput `pulumi:"ipComparisonSymbol"`
	// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitIps pulumi.StringArrayOutput `pulumi:"limitIps"`
	// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
	LimitType pulumi.StringOutput `pulumi:"limitType"`
	// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitVpcs BackupDownloadRestrictionLimitVpcArrayOutput `pulumi:"limitVpcs"`
	// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
	VpcComparisonSymbol pulumi.StringPtrOutput `pulumi:"vpcComparisonSymbol"`
}

// NewBackupDownloadRestriction registers a new resource with the given unique name, arguments, and options.
func NewBackupDownloadRestriction(ctx *pulumi.Context,
	name string, args *BackupDownloadRestrictionArgs, opts ...pulumi.ResourceOption) (*BackupDownloadRestriction, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LimitType == nil {
		return nil, errors.New("invalid value for required argument 'LimitType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BackupDownloadRestriction
	err := ctx.RegisterResource("tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupDownloadRestriction gets an existing BackupDownloadRestriction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupDownloadRestriction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupDownloadRestrictionState, opts ...pulumi.ResourceOption) (*BackupDownloadRestriction, error) {
	var resource BackupDownloadRestriction
	err := ctx.ReadResource("tencentcloud:Redis/backupDownloadRestriction:BackupDownloadRestriction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupDownloadRestriction resources.
type backupDownloadRestrictionState struct {
	// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
	IpComparisonSymbol *string `pulumi:"ipComparisonSymbol"`
	// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitIps []string `pulumi:"limitIps"`
	// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
	LimitType *string `pulumi:"limitType"`
	// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitVpcs []BackupDownloadRestrictionLimitVpc `pulumi:"limitVpcs"`
	// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
	VpcComparisonSymbol *string `pulumi:"vpcComparisonSymbol"`
}

type BackupDownloadRestrictionState struct {
	// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
	IpComparisonSymbol pulumi.StringPtrInput
	// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitIps pulumi.StringArrayInput
	// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
	LimitType pulumi.StringPtrInput
	// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitVpcs BackupDownloadRestrictionLimitVpcArrayInput
	// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
	VpcComparisonSymbol pulumi.StringPtrInput
}

func (BackupDownloadRestrictionState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupDownloadRestrictionState)(nil)).Elem()
}

type backupDownloadRestrictionArgs struct {
	// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
	IpComparisonSymbol *string `pulumi:"ipComparisonSymbol"`
	// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitIps []string `pulumi:"limitIps"`
	// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
	LimitType string `pulumi:"limitType"`
	// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitVpcs []BackupDownloadRestrictionLimitVpc `pulumi:"limitVpcs"`
	// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
	VpcComparisonSymbol *string `pulumi:"vpcComparisonSymbol"`
}

// The set of arguments for constructing a BackupDownloadRestriction resource.
type BackupDownloadRestrictionArgs struct {
	// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
	IpComparisonSymbol pulumi.StringPtrInput
	// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitIps pulumi.StringArrayInput
	// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
	LimitType pulumi.StringInput
	// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
	LimitVpcs BackupDownloadRestrictionLimitVpcArrayInput
	// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
	VpcComparisonSymbol pulumi.StringPtrInput
}

func (BackupDownloadRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupDownloadRestrictionArgs)(nil)).Elem()
}

type BackupDownloadRestrictionInput interface {
	pulumi.Input

	ToBackupDownloadRestrictionOutput() BackupDownloadRestrictionOutput
	ToBackupDownloadRestrictionOutputWithContext(ctx context.Context) BackupDownloadRestrictionOutput
}

func (*BackupDownloadRestriction) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupDownloadRestriction)(nil)).Elem()
}

func (i *BackupDownloadRestriction) ToBackupDownloadRestrictionOutput() BackupDownloadRestrictionOutput {
	return i.ToBackupDownloadRestrictionOutputWithContext(context.Background())
}

func (i *BackupDownloadRestriction) ToBackupDownloadRestrictionOutputWithContext(ctx context.Context) BackupDownloadRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupDownloadRestrictionOutput)
}

// BackupDownloadRestrictionArrayInput is an input type that accepts BackupDownloadRestrictionArray and BackupDownloadRestrictionArrayOutput values.
// You can construct a concrete instance of `BackupDownloadRestrictionArrayInput` via:
//
//	BackupDownloadRestrictionArray{ BackupDownloadRestrictionArgs{...} }
type BackupDownloadRestrictionArrayInput interface {
	pulumi.Input

	ToBackupDownloadRestrictionArrayOutput() BackupDownloadRestrictionArrayOutput
	ToBackupDownloadRestrictionArrayOutputWithContext(context.Context) BackupDownloadRestrictionArrayOutput
}

type BackupDownloadRestrictionArray []BackupDownloadRestrictionInput

func (BackupDownloadRestrictionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupDownloadRestriction)(nil)).Elem()
}

func (i BackupDownloadRestrictionArray) ToBackupDownloadRestrictionArrayOutput() BackupDownloadRestrictionArrayOutput {
	return i.ToBackupDownloadRestrictionArrayOutputWithContext(context.Background())
}

func (i BackupDownloadRestrictionArray) ToBackupDownloadRestrictionArrayOutputWithContext(ctx context.Context) BackupDownloadRestrictionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupDownloadRestrictionArrayOutput)
}

// BackupDownloadRestrictionMapInput is an input type that accepts BackupDownloadRestrictionMap and BackupDownloadRestrictionMapOutput values.
// You can construct a concrete instance of `BackupDownloadRestrictionMapInput` via:
//
//	BackupDownloadRestrictionMap{ "key": BackupDownloadRestrictionArgs{...} }
type BackupDownloadRestrictionMapInput interface {
	pulumi.Input

	ToBackupDownloadRestrictionMapOutput() BackupDownloadRestrictionMapOutput
	ToBackupDownloadRestrictionMapOutputWithContext(context.Context) BackupDownloadRestrictionMapOutput
}

type BackupDownloadRestrictionMap map[string]BackupDownloadRestrictionInput

func (BackupDownloadRestrictionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupDownloadRestriction)(nil)).Elem()
}

func (i BackupDownloadRestrictionMap) ToBackupDownloadRestrictionMapOutput() BackupDownloadRestrictionMapOutput {
	return i.ToBackupDownloadRestrictionMapOutputWithContext(context.Background())
}

func (i BackupDownloadRestrictionMap) ToBackupDownloadRestrictionMapOutputWithContext(ctx context.Context) BackupDownloadRestrictionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupDownloadRestrictionMapOutput)
}

type BackupDownloadRestrictionOutput struct{ *pulumi.OutputState }

func (BackupDownloadRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupDownloadRestriction)(nil)).Elem()
}

func (o BackupDownloadRestrictionOutput) ToBackupDownloadRestrictionOutput() BackupDownloadRestrictionOutput {
	return o
}

func (o BackupDownloadRestrictionOutput) ToBackupDownloadRestrictionOutputWithContext(ctx context.Context) BackupDownloadRestrictionOutput {
	return o
}

// Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.
func (o BackupDownloadRestrictionOutput) IpComparisonSymbol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupDownloadRestriction) pulumi.StringPtrOutput { return v.IpComparisonSymbol }).(pulumi.StringPtrOutput)
}

// A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.
func (o BackupDownloadRestrictionOutput) LimitIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupDownloadRestriction) pulumi.StringArrayOutput { return v.LimitIps }).(pulumi.StringArrayOutput)
}

// Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.
func (o BackupDownloadRestrictionOutput) LimitType() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupDownloadRestriction) pulumi.StringOutput { return v.LimitType }).(pulumi.StringOutput)
}

// A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.
func (o BackupDownloadRestrictionOutput) LimitVpcs() BackupDownloadRestrictionLimitVpcArrayOutput {
	return o.ApplyT(func(v *BackupDownloadRestriction) BackupDownloadRestrictionLimitVpcArrayOutput { return v.LimitVpcs }).(BackupDownloadRestrictionLimitVpcArrayOutput)
}

// This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.
func (o BackupDownloadRestrictionOutput) VpcComparisonSymbol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupDownloadRestriction) pulumi.StringPtrOutput { return v.VpcComparisonSymbol }).(pulumi.StringPtrOutput)
}

type BackupDownloadRestrictionArrayOutput struct{ *pulumi.OutputState }

func (BackupDownloadRestrictionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupDownloadRestriction)(nil)).Elem()
}

func (o BackupDownloadRestrictionArrayOutput) ToBackupDownloadRestrictionArrayOutput() BackupDownloadRestrictionArrayOutput {
	return o
}

func (o BackupDownloadRestrictionArrayOutput) ToBackupDownloadRestrictionArrayOutputWithContext(ctx context.Context) BackupDownloadRestrictionArrayOutput {
	return o
}

func (o BackupDownloadRestrictionArrayOutput) Index(i pulumi.IntInput) BackupDownloadRestrictionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupDownloadRestriction {
		return vs[0].([]*BackupDownloadRestriction)[vs[1].(int)]
	}).(BackupDownloadRestrictionOutput)
}

type BackupDownloadRestrictionMapOutput struct{ *pulumi.OutputState }

func (BackupDownloadRestrictionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupDownloadRestriction)(nil)).Elem()
}

func (o BackupDownloadRestrictionMapOutput) ToBackupDownloadRestrictionMapOutput() BackupDownloadRestrictionMapOutput {
	return o
}

func (o BackupDownloadRestrictionMapOutput) ToBackupDownloadRestrictionMapOutputWithContext(ctx context.Context) BackupDownloadRestrictionMapOutput {
	return o
}

func (o BackupDownloadRestrictionMapOutput) MapIndex(k pulumi.StringInput) BackupDownloadRestrictionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupDownloadRestriction {
		return vs[0].(map[string]*BackupDownloadRestriction)[vs[1].(string)]
	}).(BackupDownloadRestrictionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupDownloadRestrictionInput)(nil)).Elem(), &BackupDownloadRestriction{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupDownloadRestrictionArrayInput)(nil)).Elem(), BackupDownloadRestrictionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupDownloadRestrictionMapInput)(nil)).Elem(), BackupDownloadRestrictionMap{})
	pulumi.RegisterOutputType(BackupDownloadRestrictionOutput{})
	pulumi.RegisterOutputType(BackupDownloadRestrictionArrayOutput{})
	pulumi.RegisterOutputType(BackupDownloadRestrictionMapOutput{})
}
