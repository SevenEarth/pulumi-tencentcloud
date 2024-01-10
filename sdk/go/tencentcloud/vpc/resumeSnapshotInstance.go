// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc resumeSnapshotInstance
//
// ## Example Usage
// ### Basic example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.NewResumeSnapshotInstance(ctx, "resumeSnapshotInstance", &Vpc.ResumeSnapshotInstanceArgs{
//				InstanceId:       pulumi.String("ntrgm89v"),
//				SnapshotFileId:   pulumi.String("ssfile-emtabuwu2z"),
//				SnapshotPolicyId: pulumi.String("sspolicy-1t6cobbv"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Complete example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cos"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleSnapshotFiles, err := Vpc.GetSnapshotFiles(ctx, &vpc.GetSnapshotFilesArgs{
//				BusinessType: "securitygroup",
//				InstanceId:   "sg-902tl7t7",
//				StartDate:    "2022-10-10 00:00:00",
//				EndDate:      "2023-10-30 00:00:00",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleBucket, err := Cos.NewBucket(ctx, "exampleBucket", &Cos.BucketArgs{
//				Bucket: pulumi.String("tf-example-1308919341"),
//				Acl:    pulumi.String("private"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleSnapshotPolicy, err := Vpc.NewSnapshotPolicy(ctx, "exampleSnapshotPolicy", &Vpc.SnapshotPolicyArgs{
//				SnapshotPolicyName: pulumi.String("tf-example"),
//				BackupType:         pulumi.String("time"),
//				CosBucket:          exampleBucket.Bucket,
//				CosRegion:          pulumi.String("ap-guangzhou"),
//				CreateNewCos:       pulumi.Bool(false),
//				KeepTime:           pulumi.Int(2),
//				BackupPolicies: vpc.SnapshotPolicyBackupPolicyArray{
//					&vpc.SnapshotPolicyBackupPolicyArgs{
//						BackupDay:  pulumi.String("monday"),
//						BackupTime: pulumi.String("00:00:00"),
//					},
//					&vpc.SnapshotPolicyBackupPolicyArgs{
//						BackupDay:  pulumi.String("tuesday"),
//						BackupTime: pulumi.String("01:00:00"),
//					},
//					&vpc.SnapshotPolicyBackupPolicyArgs{
//						BackupDay:  pulumi.String("wednesday"),
//						BackupTime: pulumi.String("02:00:00"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Vpc.NewResumeSnapshotInstance(ctx, "exampleResumeSnapshotInstance", &Vpc.ResumeSnapshotInstanceArgs{
//				SnapshotPolicyId: exampleSnapshotPolicy.ID(),
//				SnapshotFileId:   pulumi.String(exampleSnapshotFiles.SnapshotFileSets[0].SnapshotFileId),
//				InstanceId:       pulumi.String("policy-1t6cob"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ResumeSnapshotInstance struct {
	pulumi.CustomResourceState

	// InstanceId.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Snapshot file Id.
	SnapshotFileId pulumi.StringOutput `pulumi:"snapshotFileId"`
	// Snapshot policy Id.
	SnapshotPolicyId pulumi.StringOutput `pulumi:"snapshotPolicyId"`
}

// NewResumeSnapshotInstance registers a new resource with the given unique name, arguments, and options.
func NewResumeSnapshotInstance(ctx *pulumi.Context,
	name string, args *ResumeSnapshotInstanceArgs, opts ...pulumi.ResourceOption) (*ResumeSnapshotInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.SnapshotFileId == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotFileId'")
	}
	if args.SnapshotPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'SnapshotPolicyId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ResumeSnapshotInstance
	err := ctx.RegisterResource("tencentcloud:Vpc/resumeSnapshotInstance:ResumeSnapshotInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResumeSnapshotInstance gets an existing ResumeSnapshotInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResumeSnapshotInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResumeSnapshotInstanceState, opts ...pulumi.ResourceOption) (*ResumeSnapshotInstance, error) {
	var resource ResumeSnapshotInstance
	err := ctx.ReadResource("tencentcloud:Vpc/resumeSnapshotInstance:ResumeSnapshotInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResumeSnapshotInstance resources.
type resumeSnapshotInstanceState struct {
	// InstanceId.
	InstanceId *string `pulumi:"instanceId"`
	// Snapshot file Id.
	SnapshotFileId *string `pulumi:"snapshotFileId"`
	// Snapshot policy Id.
	SnapshotPolicyId *string `pulumi:"snapshotPolicyId"`
}

type ResumeSnapshotInstanceState struct {
	// InstanceId.
	InstanceId pulumi.StringPtrInput
	// Snapshot file Id.
	SnapshotFileId pulumi.StringPtrInput
	// Snapshot policy Id.
	SnapshotPolicyId pulumi.StringPtrInput
}

func (ResumeSnapshotInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resumeSnapshotInstanceState)(nil)).Elem()
}

type resumeSnapshotInstanceArgs struct {
	// InstanceId.
	InstanceId string `pulumi:"instanceId"`
	// Snapshot file Id.
	SnapshotFileId string `pulumi:"snapshotFileId"`
	// Snapshot policy Id.
	SnapshotPolicyId string `pulumi:"snapshotPolicyId"`
}

// The set of arguments for constructing a ResumeSnapshotInstance resource.
type ResumeSnapshotInstanceArgs struct {
	// InstanceId.
	InstanceId pulumi.StringInput
	// Snapshot file Id.
	SnapshotFileId pulumi.StringInput
	// Snapshot policy Id.
	SnapshotPolicyId pulumi.StringInput
}

func (ResumeSnapshotInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resumeSnapshotInstanceArgs)(nil)).Elem()
}

type ResumeSnapshotInstanceInput interface {
	pulumi.Input

	ToResumeSnapshotInstanceOutput() ResumeSnapshotInstanceOutput
	ToResumeSnapshotInstanceOutputWithContext(ctx context.Context) ResumeSnapshotInstanceOutput
}

func (*ResumeSnapshotInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ResumeSnapshotInstance)(nil)).Elem()
}

func (i *ResumeSnapshotInstance) ToResumeSnapshotInstanceOutput() ResumeSnapshotInstanceOutput {
	return i.ToResumeSnapshotInstanceOutputWithContext(context.Background())
}

func (i *ResumeSnapshotInstance) ToResumeSnapshotInstanceOutputWithContext(ctx context.Context) ResumeSnapshotInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResumeSnapshotInstanceOutput)
}

// ResumeSnapshotInstanceArrayInput is an input type that accepts ResumeSnapshotInstanceArray and ResumeSnapshotInstanceArrayOutput values.
// You can construct a concrete instance of `ResumeSnapshotInstanceArrayInput` via:
//
//	ResumeSnapshotInstanceArray{ ResumeSnapshotInstanceArgs{...} }
type ResumeSnapshotInstanceArrayInput interface {
	pulumi.Input

	ToResumeSnapshotInstanceArrayOutput() ResumeSnapshotInstanceArrayOutput
	ToResumeSnapshotInstanceArrayOutputWithContext(context.Context) ResumeSnapshotInstanceArrayOutput
}

type ResumeSnapshotInstanceArray []ResumeSnapshotInstanceInput

func (ResumeSnapshotInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResumeSnapshotInstance)(nil)).Elem()
}

func (i ResumeSnapshotInstanceArray) ToResumeSnapshotInstanceArrayOutput() ResumeSnapshotInstanceArrayOutput {
	return i.ToResumeSnapshotInstanceArrayOutputWithContext(context.Background())
}

func (i ResumeSnapshotInstanceArray) ToResumeSnapshotInstanceArrayOutputWithContext(ctx context.Context) ResumeSnapshotInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResumeSnapshotInstanceArrayOutput)
}

// ResumeSnapshotInstanceMapInput is an input type that accepts ResumeSnapshotInstanceMap and ResumeSnapshotInstanceMapOutput values.
// You can construct a concrete instance of `ResumeSnapshotInstanceMapInput` via:
//
//	ResumeSnapshotInstanceMap{ "key": ResumeSnapshotInstanceArgs{...} }
type ResumeSnapshotInstanceMapInput interface {
	pulumi.Input

	ToResumeSnapshotInstanceMapOutput() ResumeSnapshotInstanceMapOutput
	ToResumeSnapshotInstanceMapOutputWithContext(context.Context) ResumeSnapshotInstanceMapOutput
}

type ResumeSnapshotInstanceMap map[string]ResumeSnapshotInstanceInput

func (ResumeSnapshotInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResumeSnapshotInstance)(nil)).Elem()
}

func (i ResumeSnapshotInstanceMap) ToResumeSnapshotInstanceMapOutput() ResumeSnapshotInstanceMapOutput {
	return i.ToResumeSnapshotInstanceMapOutputWithContext(context.Background())
}

func (i ResumeSnapshotInstanceMap) ToResumeSnapshotInstanceMapOutputWithContext(ctx context.Context) ResumeSnapshotInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResumeSnapshotInstanceMapOutput)
}

type ResumeSnapshotInstanceOutput struct{ *pulumi.OutputState }

func (ResumeSnapshotInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResumeSnapshotInstance)(nil)).Elem()
}

func (o ResumeSnapshotInstanceOutput) ToResumeSnapshotInstanceOutput() ResumeSnapshotInstanceOutput {
	return o
}

func (o ResumeSnapshotInstanceOutput) ToResumeSnapshotInstanceOutputWithContext(ctx context.Context) ResumeSnapshotInstanceOutput {
	return o
}

// InstanceId.
func (o ResumeSnapshotInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResumeSnapshotInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Snapshot file Id.
func (o ResumeSnapshotInstanceOutput) SnapshotFileId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResumeSnapshotInstance) pulumi.StringOutput { return v.SnapshotFileId }).(pulumi.StringOutput)
}

// Snapshot policy Id.
func (o ResumeSnapshotInstanceOutput) SnapshotPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ResumeSnapshotInstance) pulumi.StringOutput { return v.SnapshotPolicyId }).(pulumi.StringOutput)
}

type ResumeSnapshotInstanceArrayOutput struct{ *pulumi.OutputState }

func (ResumeSnapshotInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResumeSnapshotInstance)(nil)).Elem()
}

func (o ResumeSnapshotInstanceArrayOutput) ToResumeSnapshotInstanceArrayOutput() ResumeSnapshotInstanceArrayOutput {
	return o
}

func (o ResumeSnapshotInstanceArrayOutput) ToResumeSnapshotInstanceArrayOutputWithContext(ctx context.Context) ResumeSnapshotInstanceArrayOutput {
	return o
}

func (o ResumeSnapshotInstanceArrayOutput) Index(i pulumi.IntInput) ResumeSnapshotInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResumeSnapshotInstance {
		return vs[0].([]*ResumeSnapshotInstance)[vs[1].(int)]
	}).(ResumeSnapshotInstanceOutput)
}

type ResumeSnapshotInstanceMapOutput struct{ *pulumi.OutputState }

func (ResumeSnapshotInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResumeSnapshotInstance)(nil)).Elem()
}

func (o ResumeSnapshotInstanceMapOutput) ToResumeSnapshotInstanceMapOutput() ResumeSnapshotInstanceMapOutput {
	return o
}

func (o ResumeSnapshotInstanceMapOutput) ToResumeSnapshotInstanceMapOutputWithContext(ctx context.Context) ResumeSnapshotInstanceMapOutput {
	return o
}

func (o ResumeSnapshotInstanceMapOutput) MapIndex(k pulumi.StringInput) ResumeSnapshotInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResumeSnapshotInstance {
		return vs[0].(map[string]*ResumeSnapshotInstance)[vs[1].(string)]
	}).(ResumeSnapshotInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResumeSnapshotInstanceInput)(nil)).Elem(), &ResumeSnapshotInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResumeSnapshotInstanceArrayInput)(nil)).Elem(), ResumeSnapshotInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResumeSnapshotInstanceMapInput)(nil)).Elem(), ResumeSnapshotInstanceMap{})
	pulumi.RegisterOutputType(ResumeSnapshotInstanceOutput{})
	pulumi.RegisterOutputType(ResumeSnapshotInstanceArrayOutput{})
	pulumi.RegisterOutputType(ResumeSnapshotInstanceMapOutput{})
}
