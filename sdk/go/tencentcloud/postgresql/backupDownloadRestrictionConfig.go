// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a postgresql backupDownloadRestrictionConfig
//
// ## Example Usage
//
// ### Unlimit the restriction of the backup file download.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Postgresql.NewBackupDownloadRestrictionConfig(ctx, "backupDownloadRestrictionConfig", &Postgresql.BackupDownloadRestrictionConfigArgs{
//				RestrictionType: pulumi.String("NONE"),
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
//
// ### Set the download only to allow the intranet downloads.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Postgresql.NewBackupDownloadRestrictionConfig(ctx, "backupDownloadRestrictionConfig", &Postgresql.BackupDownloadRestrictionConfigArgs{
//				RestrictionType: pulumi.String("INTRANET"),
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
//
// ### Restrict the backup file download by customizing.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Postgresql"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.NewInstance(ctx, "pgVpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.Any(_var.Vpc_cidr),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Postgresql.NewBackupDownloadRestrictionConfig(ctx, "backupDownloadRestrictionConfig", &Postgresql.BackupDownloadRestrictionConfigArgs{
//				RestrictionType:      pulumi.String("CUSTOMIZE"),
//				VpcRestrictionEffect: pulumi.String("DENY"),
//				VpcIdSets: pulumi.StringArray{
//					tencentcloud_vpc.Pg_vpc2.Id,
//				},
//				IpRestrictionEffect: pulumi.String("DENY"),
//				IpSets: pulumi.StringArray{
//					pulumi.String("192.168.0.0"),
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
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// postgresql backup_download_restriction_config can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Postgresql/backupDownloadRestrictionConfig:BackupDownloadRestrictionConfig backup_download_restriction_config backup_download_restriction_config_id
// ```
type BackupDownloadRestrictionConfig struct {
	pulumi.CustomResourceState

	// ip limit Strategy: ALLOW, DENY.
	IpRestrictionEffect pulumi.StringPtrOutput `pulumi:"ipRestrictionEffect"`
	// The list of ips that are allowed or denied to download backup files.
	IpSets pulumi.StringArrayOutput `pulumi:"ipSets"`
	// Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
	RestrictionType pulumi.StringOutput `pulumi:"restrictionType"`
	// The list of vpcIds that allow or deny downloading of backup files.
	VpcIdSets pulumi.StringArrayOutput `pulumi:"vpcIdSets"`
	// vpc limit Strategy: ALLOW, DENY.
	VpcRestrictionEffect pulumi.StringPtrOutput `pulumi:"vpcRestrictionEffect"`
}

// NewBackupDownloadRestrictionConfig registers a new resource with the given unique name, arguments, and options.
func NewBackupDownloadRestrictionConfig(ctx *pulumi.Context,
	name string, args *BackupDownloadRestrictionConfigArgs, opts ...pulumi.ResourceOption) (*BackupDownloadRestrictionConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RestrictionType == nil {
		return nil, errors.New("invalid value for required argument 'RestrictionType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BackupDownloadRestrictionConfig
	err := ctx.RegisterResource("tencentcloud:Postgresql/backupDownloadRestrictionConfig:BackupDownloadRestrictionConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupDownloadRestrictionConfig gets an existing BackupDownloadRestrictionConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupDownloadRestrictionConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupDownloadRestrictionConfigState, opts ...pulumi.ResourceOption) (*BackupDownloadRestrictionConfig, error) {
	var resource BackupDownloadRestrictionConfig
	err := ctx.ReadResource("tencentcloud:Postgresql/backupDownloadRestrictionConfig:BackupDownloadRestrictionConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupDownloadRestrictionConfig resources.
type backupDownloadRestrictionConfigState struct {
	// ip limit Strategy: ALLOW, DENY.
	IpRestrictionEffect *string `pulumi:"ipRestrictionEffect"`
	// The list of ips that are allowed or denied to download backup files.
	IpSets []string `pulumi:"ipSets"`
	// Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
	RestrictionType *string `pulumi:"restrictionType"`
	// The list of vpcIds that allow or deny downloading of backup files.
	VpcIdSets []string `pulumi:"vpcIdSets"`
	// vpc limit Strategy: ALLOW, DENY.
	VpcRestrictionEffect *string `pulumi:"vpcRestrictionEffect"`
}

type BackupDownloadRestrictionConfigState struct {
	// ip limit Strategy: ALLOW, DENY.
	IpRestrictionEffect pulumi.StringPtrInput
	// The list of ips that are allowed or denied to download backup files.
	IpSets pulumi.StringArrayInput
	// Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
	RestrictionType pulumi.StringPtrInput
	// The list of vpcIds that allow or deny downloading of backup files.
	VpcIdSets pulumi.StringArrayInput
	// vpc limit Strategy: ALLOW, DENY.
	VpcRestrictionEffect pulumi.StringPtrInput
}

func (BackupDownloadRestrictionConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupDownloadRestrictionConfigState)(nil)).Elem()
}

type backupDownloadRestrictionConfigArgs struct {
	// ip limit Strategy: ALLOW, DENY.
	IpRestrictionEffect *string `pulumi:"ipRestrictionEffect"`
	// The list of ips that are allowed or denied to download backup files.
	IpSets []string `pulumi:"ipSets"`
	// Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
	RestrictionType string `pulumi:"restrictionType"`
	// The list of vpcIds that allow or deny downloading of backup files.
	VpcIdSets []string `pulumi:"vpcIdSets"`
	// vpc limit Strategy: ALLOW, DENY.
	VpcRestrictionEffect *string `pulumi:"vpcRestrictionEffect"`
}

// The set of arguments for constructing a BackupDownloadRestrictionConfig resource.
type BackupDownloadRestrictionConfigArgs struct {
	// ip limit Strategy: ALLOW, DENY.
	IpRestrictionEffect pulumi.StringPtrInput
	// The list of ips that are allowed or denied to download backup files.
	IpSets pulumi.StringArrayInput
	// Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
	RestrictionType pulumi.StringInput
	// The list of vpcIds that allow or deny downloading of backup files.
	VpcIdSets pulumi.StringArrayInput
	// vpc limit Strategy: ALLOW, DENY.
	VpcRestrictionEffect pulumi.StringPtrInput
}

func (BackupDownloadRestrictionConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupDownloadRestrictionConfigArgs)(nil)).Elem()
}

type BackupDownloadRestrictionConfigInput interface {
	pulumi.Input

	ToBackupDownloadRestrictionConfigOutput() BackupDownloadRestrictionConfigOutput
	ToBackupDownloadRestrictionConfigOutputWithContext(ctx context.Context) BackupDownloadRestrictionConfigOutput
}

func (*BackupDownloadRestrictionConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupDownloadRestrictionConfig)(nil)).Elem()
}

func (i *BackupDownloadRestrictionConfig) ToBackupDownloadRestrictionConfigOutput() BackupDownloadRestrictionConfigOutput {
	return i.ToBackupDownloadRestrictionConfigOutputWithContext(context.Background())
}

func (i *BackupDownloadRestrictionConfig) ToBackupDownloadRestrictionConfigOutputWithContext(ctx context.Context) BackupDownloadRestrictionConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupDownloadRestrictionConfigOutput)
}

// BackupDownloadRestrictionConfigArrayInput is an input type that accepts BackupDownloadRestrictionConfigArray and BackupDownloadRestrictionConfigArrayOutput values.
// You can construct a concrete instance of `BackupDownloadRestrictionConfigArrayInput` via:
//
//	BackupDownloadRestrictionConfigArray{ BackupDownloadRestrictionConfigArgs{...} }
type BackupDownloadRestrictionConfigArrayInput interface {
	pulumi.Input

	ToBackupDownloadRestrictionConfigArrayOutput() BackupDownloadRestrictionConfigArrayOutput
	ToBackupDownloadRestrictionConfigArrayOutputWithContext(context.Context) BackupDownloadRestrictionConfigArrayOutput
}

type BackupDownloadRestrictionConfigArray []BackupDownloadRestrictionConfigInput

func (BackupDownloadRestrictionConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupDownloadRestrictionConfig)(nil)).Elem()
}

func (i BackupDownloadRestrictionConfigArray) ToBackupDownloadRestrictionConfigArrayOutput() BackupDownloadRestrictionConfigArrayOutput {
	return i.ToBackupDownloadRestrictionConfigArrayOutputWithContext(context.Background())
}

func (i BackupDownloadRestrictionConfigArray) ToBackupDownloadRestrictionConfigArrayOutputWithContext(ctx context.Context) BackupDownloadRestrictionConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupDownloadRestrictionConfigArrayOutput)
}

// BackupDownloadRestrictionConfigMapInput is an input type that accepts BackupDownloadRestrictionConfigMap and BackupDownloadRestrictionConfigMapOutput values.
// You can construct a concrete instance of `BackupDownloadRestrictionConfigMapInput` via:
//
//	BackupDownloadRestrictionConfigMap{ "key": BackupDownloadRestrictionConfigArgs{...} }
type BackupDownloadRestrictionConfigMapInput interface {
	pulumi.Input

	ToBackupDownloadRestrictionConfigMapOutput() BackupDownloadRestrictionConfigMapOutput
	ToBackupDownloadRestrictionConfigMapOutputWithContext(context.Context) BackupDownloadRestrictionConfigMapOutput
}

type BackupDownloadRestrictionConfigMap map[string]BackupDownloadRestrictionConfigInput

func (BackupDownloadRestrictionConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupDownloadRestrictionConfig)(nil)).Elem()
}

func (i BackupDownloadRestrictionConfigMap) ToBackupDownloadRestrictionConfigMapOutput() BackupDownloadRestrictionConfigMapOutput {
	return i.ToBackupDownloadRestrictionConfigMapOutputWithContext(context.Background())
}

func (i BackupDownloadRestrictionConfigMap) ToBackupDownloadRestrictionConfigMapOutputWithContext(ctx context.Context) BackupDownloadRestrictionConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupDownloadRestrictionConfigMapOutput)
}

type BackupDownloadRestrictionConfigOutput struct{ *pulumi.OutputState }

func (BackupDownloadRestrictionConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupDownloadRestrictionConfig)(nil)).Elem()
}

func (o BackupDownloadRestrictionConfigOutput) ToBackupDownloadRestrictionConfigOutput() BackupDownloadRestrictionConfigOutput {
	return o
}

func (o BackupDownloadRestrictionConfigOutput) ToBackupDownloadRestrictionConfigOutputWithContext(ctx context.Context) BackupDownloadRestrictionConfigOutput {
	return o
}

// ip limit Strategy: ALLOW, DENY.
func (o BackupDownloadRestrictionConfigOutput) IpRestrictionEffect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupDownloadRestrictionConfig) pulumi.StringPtrOutput { return v.IpRestrictionEffect }).(pulumi.StringPtrOutput)
}

// The list of ips that are allowed or denied to download backup files.
func (o BackupDownloadRestrictionConfigOutput) IpSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupDownloadRestrictionConfig) pulumi.StringArrayOutput { return v.IpSets }).(pulumi.StringArrayOutput)
}

// Backup file download restriction type: NONE:Unlimited, both internal and external networks can be downloaded. INTRANET:Only intranet downloads are allowed. CUSTOMIZE:Customize the vpc or ip that limits downloads.
func (o BackupDownloadRestrictionConfigOutput) RestrictionType() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupDownloadRestrictionConfig) pulumi.StringOutput { return v.RestrictionType }).(pulumi.StringOutput)
}

// The list of vpcIds that allow or deny downloading of backup files.
func (o BackupDownloadRestrictionConfigOutput) VpcIdSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BackupDownloadRestrictionConfig) pulumi.StringArrayOutput { return v.VpcIdSets }).(pulumi.StringArrayOutput)
}

// vpc limit Strategy: ALLOW, DENY.
func (o BackupDownloadRestrictionConfigOutput) VpcRestrictionEffect() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupDownloadRestrictionConfig) pulumi.StringPtrOutput { return v.VpcRestrictionEffect }).(pulumi.StringPtrOutput)
}

type BackupDownloadRestrictionConfigArrayOutput struct{ *pulumi.OutputState }

func (BackupDownloadRestrictionConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupDownloadRestrictionConfig)(nil)).Elem()
}

func (o BackupDownloadRestrictionConfigArrayOutput) ToBackupDownloadRestrictionConfigArrayOutput() BackupDownloadRestrictionConfigArrayOutput {
	return o
}

func (o BackupDownloadRestrictionConfigArrayOutput) ToBackupDownloadRestrictionConfigArrayOutputWithContext(ctx context.Context) BackupDownloadRestrictionConfigArrayOutput {
	return o
}

func (o BackupDownloadRestrictionConfigArrayOutput) Index(i pulumi.IntInput) BackupDownloadRestrictionConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupDownloadRestrictionConfig {
		return vs[0].([]*BackupDownloadRestrictionConfig)[vs[1].(int)]
	}).(BackupDownloadRestrictionConfigOutput)
}

type BackupDownloadRestrictionConfigMapOutput struct{ *pulumi.OutputState }

func (BackupDownloadRestrictionConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupDownloadRestrictionConfig)(nil)).Elem()
}

func (o BackupDownloadRestrictionConfigMapOutput) ToBackupDownloadRestrictionConfigMapOutput() BackupDownloadRestrictionConfigMapOutput {
	return o
}

func (o BackupDownloadRestrictionConfigMapOutput) ToBackupDownloadRestrictionConfigMapOutputWithContext(ctx context.Context) BackupDownloadRestrictionConfigMapOutput {
	return o
}

func (o BackupDownloadRestrictionConfigMapOutput) MapIndex(k pulumi.StringInput) BackupDownloadRestrictionConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupDownloadRestrictionConfig {
		return vs[0].(map[string]*BackupDownloadRestrictionConfig)[vs[1].(string)]
	}).(BackupDownloadRestrictionConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupDownloadRestrictionConfigInput)(nil)).Elem(), &BackupDownloadRestrictionConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupDownloadRestrictionConfigArrayInput)(nil)).Elem(), BackupDownloadRestrictionConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupDownloadRestrictionConfigMapInput)(nil)).Elem(), BackupDownloadRestrictionConfigMap{})
	pulumi.RegisterOutputType(BackupDownloadRestrictionConfigOutput{})
	pulumi.RegisterOutputType(BackupDownloadRestrictionConfigArrayOutput{})
	pulumi.RegisterOutputType(BackupDownloadRestrictionConfigMapOutput{})
}
