// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a lighthouse instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Lighthouse"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Lighthouse.NewInstance(ctx, "lighthouse", &Lighthouse.InstanceArgs{
//				BlueprintId: pulumi.String("lhbp-f1lkcd41"),
//				BundleId:    pulumi.String("bundle2022_gen_01"),
//				Containers: lighthouse.InstanceContainerArray{
//					&lighthouse.InstanceContainerArgs{
//						Command:        pulumi.String("ls -l"),
//						ContainerImage: pulumi.String("ccr.ccs.tencentyun.com/qcloud/nginx"),
//						ContainerName:  pulumi.String("nginx"),
//						Envs: lighthouse.InstanceContainerEnvArray{
//							&lighthouse.InstanceContainerEnvArgs{
//								Key:   pulumi.String("key"),
//								Value: pulumi.String("value"),
//							},
//							&lighthouse.InstanceContainerEnvArgs{
//								Key:   pulumi.String("key2"),
//								Value: pulumi.String("value2"),
//							},
//						},
//						PublishPorts: lighthouse.InstanceContainerPublishPortArray{
//							&lighthouse.InstanceContainerPublishPortArgs{
//								ContainerPort: pulumi.Int(80),
//								HostPort:      pulumi.Int(80),
//								Ip:            pulumi.String("127.0.0.1"),
//								Protocol:      pulumi.String("tcp"),
//							},
//							&lighthouse.InstanceContainerPublishPortArgs{
//								ContainerPort: pulumi.Int(36000),
//								HostPort:      pulumi.Int(36000),
//								Ip:            pulumi.String("127.0.0.1"),
//								Protocol:      pulumi.String("tcp"),
//							},
//						},
//						Volumes: lighthouse.InstanceContainerVolumeArray{
//							&lighthouse.InstanceContainerVolumeArgs{
//								ContainerPath: pulumi.String("/data"),
//								HostPath:      pulumi.String("/tmp"),
//							},
//							&lighthouse.InstanceContainerVolumeArgs{
//								ContainerPath: pulumi.String("/var"),
//								HostPath:      pulumi.String("/tmp"),
//							},
//						},
//					},
//					&lighthouse.InstanceContainerArgs{
//						Command:        pulumi.String("echo \"hello\""),
//						ContainerImage: pulumi.String("ccr.ccs.tencentyun.com/qcloud/resty"),
//						ContainerName:  pulumi.String("resty"),
//						Envs: lighthouse.InstanceContainerEnvArray{
//							&lighthouse.InstanceContainerEnvArgs{
//								Key:   pulumi.String("key2"),
//								Value: pulumi.String("value2"),
//							},
//						},
//						PublishPorts: lighthouse.InstanceContainerPublishPortArray{
//							&lighthouse.InstanceContainerPublishPortArgs{
//								ContainerPort: pulumi.Int(80),
//								HostPort:      pulumi.Int(80),
//								Ip:            pulumi.String("127.0.0.1"),
//								Protocol:      pulumi.String("udp"),
//							},
//						},
//						Volumes: lighthouse.InstanceContainerVolumeArray{
//							&lighthouse.InstanceContainerVolumeArgs{
//								ContainerPath: pulumi.String("/var"),
//								HostPath:      pulumi.String("/tmp"),
//							},
//						},
//					},
//				},
//				InstanceName: pulumi.String("hello world"),
//				Period:       pulumi.Int(1),
//				RenewFlag:    pulumi.String("NOTIFY_AND_AUTO_RENEW"),
//				Zone:         pulumi.String("ap-guangzhou-3"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// ID of the Lighthouse image.
	BlueprintId pulumi.StringOutput `pulumi:"blueprintId"`
	// ID of the Lighthouse package.
	BundleId pulumi.StringOutput `pulumi:"bundleId"`
	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken pulumi.StringPtrOutput `pulumi:"clientToken"`
	// Configuration of the containers to create.
	Containers InstanceContainerArrayOutput `pulumi:"containers"`
	// Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// The display name of the Lighthouse instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
	LoginConfiguration InstanceLoginConfigurationPtrOutput `pulumi:"loginConfiguration"`
	// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period pulumi.IntOutput `pulumi:"period"`
	// Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
	RenewFlag pulumi.StringOutput `pulumi:"renewFlag"`
	// List of availability zones. A random AZ is selected by default.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BlueprintId == nil {
		return nil, errors.New("invalid value for required argument 'BlueprintId'")
	}
	if args.BundleId == nil {
		return nil, errors.New("invalid value for required argument 'BundleId'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	if args.RenewFlag == nil {
		return nil, errors.New("invalid value for required argument 'RenewFlag'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("tencentcloud:Lighthouse/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("tencentcloud:Lighthouse/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// ID of the Lighthouse image.
	BlueprintId *string `pulumi:"blueprintId"`
	// ID of the Lighthouse package.
	BundleId *string `pulumi:"bundleId"`
	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken *string `pulumi:"clientToken"`
	// Configuration of the containers to create.
	Containers []InstanceContainer `pulumi:"containers"`
	// Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `pulumi:"dryRun"`
	// The display name of the Lighthouse instance.
	InstanceName *string `pulumi:"instanceName"`
	// Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
	LoginConfiguration *InstanceLoginConfiguration `pulumi:"loginConfiguration"`
	// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period *int `pulumi:"period"`
	// Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
	RenewFlag *string `pulumi:"renewFlag"`
	// List of availability zones. A random AZ is selected by default.
	Zone *string `pulumi:"zone"`
}

type InstanceState struct {
	// ID of the Lighthouse image.
	BlueprintId pulumi.StringPtrInput
	// ID of the Lighthouse package.
	BundleId pulumi.StringPtrInput
	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken pulumi.StringPtrInput
	// Configuration of the containers to create.
	Containers InstanceContainerArrayInput
	// Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun pulumi.BoolPtrInput
	// The display name of the Lighthouse instance.
	InstanceName pulumi.StringPtrInput
	// Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
	LoginConfiguration InstanceLoginConfigurationPtrInput
	// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period pulumi.IntPtrInput
	// Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
	RenewFlag pulumi.StringPtrInput
	// List of availability zones. A random AZ is selected by default.
	Zone pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// ID of the Lighthouse image.
	BlueprintId string `pulumi:"blueprintId"`
	// ID of the Lighthouse package.
	BundleId string `pulumi:"bundleId"`
	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken *string `pulumi:"clientToken"`
	// Configuration of the containers to create.
	Containers []InstanceContainer `pulumi:"containers"`
	// Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun *bool `pulumi:"dryRun"`
	// The display name of the Lighthouse instance.
	InstanceName string `pulumi:"instanceName"`
	// Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
	LoginConfiguration *InstanceLoginConfiguration `pulumi:"loginConfiguration"`
	// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period int `pulumi:"period"`
	// Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
	RenewFlag string `pulumi:"renewFlag"`
	// List of availability zones. A random AZ is selected by default.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// ID of the Lighthouse image.
	BlueprintId pulumi.StringInput
	// ID of the Lighthouse package.
	BundleId pulumi.StringInput
	// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
	ClientToken pulumi.StringPtrInput
	// Configuration of the containers to create.
	Containers InstanceContainerArrayInput
	// Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
	DryRun pulumi.BoolPtrInput
	// The display name of the Lighthouse instance.
	InstanceName pulumi.StringInput
	// Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
	LoginConfiguration InstanceLoginConfigurationPtrInput
	// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
	Period pulumi.IntInput
	// Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
	RenewFlag pulumi.StringInput
	// List of availability zones. A random AZ is selected by default.
	Zone pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// ID of the Lighthouse image.
func (o InstanceOutput) BlueprintId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.BlueprintId }).(pulumi.StringOutput)
}

// ID of the Lighthouse package.
func (o InstanceOutput) BundleId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.BundleId }).(pulumi.StringOutput)
}

// A unique string supplied by the client to ensure that the request is idempotent. Its maximum length is 64 ASCII characters. If this parameter is not specified, the idem-potency of the request cannot be guaranteed.
func (o InstanceOutput) ClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.ClientToken }).(pulumi.StringPtrOutput)
}

// Configuration of the containers to create.
func (o InstanceOutput) Containers() InstanceContainerArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceContainerArrayOutput { return v.Containers }).(InstanceContainerArrayOutput)
}

// Whether the request is a dry run only.true: dry run only. The request will not create instance(s). A dry run can check whether all the required parameters are specified, whether the request format is right, whether the request exceeds service limits, and whether the specified CVMs are available. If the dry run fails, the corresponding error code will be returned.If the dry run succeeds, the RequestId will be returned.false (default value): send a normal request and create instance(s) if all the requirements are met.
func (o InstanceOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// The display name of the Lighthouse instance.
func (o InstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Login password of the instance. It is only available for Windows instances. If it is not specified, it means that the user choose to set the login password after the instance creation.
func (o InstanceOutput) LoginConfiguration() InstanceLoginConfigurationPtrOutput {
	return o.ApplyT(func(v *Instance) InstanceLoginConfigurationPtrOutput { return v.LoginConfiguration }).(InstanceLoginConfigurationPtrOutput)
}

// Subscription period in months. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60.
func (o InstanceOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

// Auto-Renewal flag. Valid values: NOTIFY_AND_AUTO_RENEW: notify upon expiration and renew automatically; NOTIFY_AND_MANUAL_RENEW: notify upon expiration but do not renew automatically. You need to manually renew DISABLE_NOTIFY_AND_AUTO_RENEW: neither notify upon expiration nor renew automatically. Default value: NOTIFY_AND_MANUAL_RENEW.
func (o InstanceOutput) RenewFlag() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RenewFlag }).(pulumi.StringOutput)
}

// List of availability zones. A random AZ is selected by default.
func (o InstanceOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
