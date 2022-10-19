// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lighthouse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceContainer struct {
	// The command to run.
	Command *string `pulumi:"command"`
	// Container image address.
	ContainerImage *string `pulumi:"containerImage"`
	// Container name.
	ContainerName *string `pulumi:"containerName"`
	// List of environment variables.
	Envs []InstanceContainerEnv `pulumi:"envs"`
	// List of mappings of container ports and host ports.
	PublishPorts []InstanceContainerPublishPort `pulumi:"publishPorts"`
	// List of container mount volumes.
	Volumes []InstanceContainerVolume `pulumi:"volumes"`
}

// InstanceContainerInput is an input type that accepts InstanceContainerArgs and InstanceContainerOutput values.
// You can construct a concrete instance of `InstanceContainerInput` via:
//
//	InstanceContainerArgs{...}
type InstanceContainerInput interface {
	pulumi.Input

	ToInstanceContainerOutput() InstanceContainerOutput
	ToInstanceContainerOutputWithContext(context.Context) InstanceContainerOutput
}

type InstanceContainerArgs struct {
	// The command to run.
	Command pulumi.StringPtrInput `pulumi:"command"`
	// Container image address.
	ContainerImage pulumi.StringPtrInput `pulumi:"containerImage"`
	// Container name.
	ContainerName pulumi.StringPtrInput `pulumi:"containerName"`
	// List of environment variables.
	Envs InstanceContainerEnvArrayInput `pulumi:"envs"`
	// List of mappings of container ports and host ports.
	PublishPorts InstanceContainerPublishPortArrayInput `pulumi:"publishPorts"`
	// List of container mount volumes.
	Volumes InstanceContainerVolumeArrayInput `pulumi:"volumes"`
}

func (InstanceContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainer)(nil)).Elem()
}

func (i InstanceContainerArgs) ToInstanceContainerOutput() InstanceContainerOutput {
	return i.ToInstanceContainerOutputWithContext(context.Background())
}

func (i InstanceContainerArgs) ToInstanceContainerOutputWithContext(ctx context.Context) InstanceContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerOutput)
}

// InstanceContainerArrayInput is an input type that accepts InstanceContainerArray and InstanceContainerArrayOutput values.
// You can construct a concrete instance of `InstanceContainerArrayInput` via:
//
//	InstanceContainerArray{ InstanceContainerArgs{...} }
type InstanceContainerArrayInput interface {
	pulumi.Input

	ToInstanceContainerArrayOutput() InstanceContainerArrayOutput
	ToInstanceContainerArrayOutputWithContext(context.Context) InstanceContainerArrayOutput
}

type InstanceContainerArray []InstanceContainerInput

func (InstanceContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainer)(nil)).Elem()
}

func (i InstanceContainerArray) ToInstanceContainerArrayOutput() InstanceContainerArrayOutput {
	return i.ToInstanceContainerArrayOutputWithContext(context.Background())
}

func (i InstanceContainerArray) ToInstanceContainerArrayOutputWithContext(ctx context.Context) InstanceContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerArrayOutput)
}

type InstanceContainerOutput struct{ *pulumi.OutputState }

func (InstanceContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainer)(nil)).Elem()
}

func (o InstanceContainerOutput) ToInstanceContainerOutput() InstanceContainerOutput {
	return o
}

func (o InstanceContainerOutput) ToInstanceContainerOutputWithContext(ctx context.Context) InstanceContainerOutput {
	return o
}

// The command to run.
func (o InstanceContainerOutput) Command() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceContainer) *string { return v.Command }).(pulumi.StringPtrOutput)
}

// Container image address.
func (o InstanceContainerOutput) ContainerImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceContainer) *string { return v.ContainerImage }).(pulumi.StringPtrOutput)
}

// Container name.
func (o InstanceContainerOutput) ContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceContainer) *string { return v.ContainerName }).(pulumi.StringPtrOutput)
}

// List of environment variables.
func (o InstanceContainerOutput) Envs() InstanceContainerEnvArrayOutput {
	return o.ApplyT(func(v InstanceContainer) []InstanceContainerEnv { return v.Envs }).(InstanceContainerEnvArrayOutput)
}

// List of mappings of container ports and host ports.
func (o InstanceContainerOutput) PublishPorts() InstanceContainerPublishPortArrayOutput {
	return o.ApplyT(func(v InstanceContainer) []InstanceContainerPublishPort { return v.PublishPorts }).(InstanceContainerPublishPortArrayOutput)
}

// List of container mount volumes.
func (o InstanceContainerOutput) Volumes() InstanceContainerVolumeArrayOutput {
	return o.ApplyT(func(v InstanceContainer) []InstanceContainerVolume { return v.Volumes }).(InstanceContainerVolumeArrayOutput)
}

type InstanceContainerArrayOutput struct{ *pulumi.OutputState }

func (InstanceContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainer)(nil)).Elem()
}

func (o InstanceContainerArrayOutput) ToInstanceContainerArrayOutput() InstanceContainerArrayOutput {
	return o
}

func (o InstanceContainerArrayOutput) ToInstanceContainerArrayOutputWithContext(ctx context.Context) InstanceContainerArrayOutput {
	return o
}

func (o InstanceContainerArrayOutput) Index(i pulumi.IntInput) InstanceContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceContainer {
		return vs[0].([]InstanceContainer)[vs[1].(int)]
	}).(InstanceContainerOutput)
}

type InstanceContainerEnv struct {
	// Environment variable key.
	Key string `pulumi:"key"`
	// Environment variable value.
	Value string `pulumi:"value"`
}

// InstanceContainerEnvInput is an input type that accepts InstanceContainerEnvArgs and InstanceContainerEnvOutput values.
// You can construct a concrete instance of `InstanceContainerEnvInput` via:
//
//	InstanceContainerEnvArgs{...}
type InstanceContainerEnvInput interface {
	pulumi.Input

	ToInstanceContainerEnvOutput() InstanceContainerEnvOutput
	ToInstanceContainerEnvOutputWithContext(context.Context) InstanceContainerEnvOutput
}

type InstanceContainerEnvArgs struct {
	// Environment variable key.
	Key pulumi.StringInput `pulumi:"key"`
	// Environment variable value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (InstanceContainerEnvArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainerEnv)(nil)).Elem()
}

func (i InstanceContainerEnvArgs) ToInstanceContainerEnvOutput() InstanceContainerEnvOutput {
	return i.ToInstanceContainerEnvOutputWithContext(context.Background())
}

func (i InstanceContainerEnvArgs) ToInstanceContainerEnvOutputWithContext(ctx context.Context) InstanceContainerEnvOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerEnvOutput)
}

// InstanceContainerEnvArrayInput is an input type that accepts InstanceContainerEnvArray and InstanceContainerEnvArrayOutput values.
// You can construct a concrete instance of `InstanceContainerEnvArrayInput` via:
//
//	InstanceContainerEnvArray{ InstanceContainerEnvArgs{...} }
type InstanceContainerEnvArrayInput interface {
	pulumi.Input

	ToInstanceContainerEnvArrayOutput() InstanceContainerEnvArrayOutput
	ToInstanceContainerEnvArrayOutputWithContext(context.Context) InstanceContainerEnvArrayOutput
}

type InstanceContainerEnvArray []InstanceContainerEnvInput

func (InstanceContainerEnvArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainerEnv)(nil)).Elem()
}

func (i InstanceContainerEnvArray) ToInstanceContainerEnvArrayOutput() InstanceContainerEnvArrayOutput {
	return i.ToInstanceContainerEnvArrayOutputWithContext(context.Background())
}

func (i InstanceContainerEnvArray) ToInstanceContainerEnvArrayOutputWithContext(ctx context.Context) InstanceContainerEnvArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerEnvArrayOutput)
}

type InstanceContainerEnvOutput struct{ *pulumi.OutputState }

func (InstanceContainerEnvOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainerEnv)(nil)).Elem()
}

func (o InstanceContainerEnvOutput) ToInstanceContainerEnvOutput() InstanceContainerEnvOutput {
	return o
}

func (o InstanceContainerEnvOutput) ToInstanceContainerEnvOutputWithContext(ctx context.Context) InstanceContainerEnvOutput {
	return o
}

// Environment variable key.
func (o InstanceContainerEnvOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceContainerEnv) string { return v.Key }).(pulumi.StringOutput)
}

// Environment variable value.
func (o InstanceContainerEnvOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceContainerEnv) string { return v.Value }).(pulumi.StringOutput)
}

type InstanceContainerEnvArrayOutput struct{ *pulumi.OutputState }

func (InstanceContainerEnvArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainerEnv)(nil)).Elem()
}

func (o InstanceContainerEnvArrayOutput) ToInstanceContainerEnvArrayOutput() InstanceContainerEnvArrayOutput {
	return o
}

func (o InstanceContainerEnvArrayOutput) ToInstanceContainerEnvArrayOutputWithContext(ctx context.Context) InstanceContainerEnvArrayOutput {
	return o
}

func (o InstanceContainerEnvArrayOutput) Index(i pulumi.IntInput) InstanceContainerEnvOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceContainerEnv {
		return vs[0].([]InstanceContainerEnv)[vs[1].(int)]
	}).(InstanceContainerEnvOutput)
}

type InstanceContainerPublishPort struct {
	// Container port.
	ContainerPort int `pulumi:"containerPort"`
	// Host port.
	HostPort int `pulumi:"hostPort"`
	// External IP. It defaults to 0.0.0.0.
	Ip *string `pulumi:"ip"`
	// The protocol defaults to tcp. Valid values: tcp, udp and sctp.
	Protocol *string `pulumi:"protocol"`
}

// InstanceContainerPublishPortInput is an input type that accepts InstanceContainerPublishPortArgs and InstanceContainerPublishPortOutput values.
// You can construct a concrete instance of `InstanceContainerPublishPortInput` via:
//
//	InstanceContainerPublishPortArgs{...}
type InstanceContainerPublishPortInput interface {
	pulumi.Input

	ToInstanceContainerPublishPortOutput() InstanceContainerPublishPortOutput
	ToInstanceContainerPublishPortOutputWithContext(context.Context) InstanceContainerPublishPortOutput
}

type InstanceContainerPublishPortArgs struct {
	// Container port.
	ContainerPort pulumi.IntInput `pulumi:"containerPort"`
	// Host port.
	HostPort pulumi.IntInput `pulumi:"hostPort"`
	// External IP. It defaults to 0.0.0.0.
	Ip pulumi.StringPtrInput `pulumi:"ip"`
	// The protocol defaults to tcp. Valid values: tcp, udp and sctp.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
}

func (InstanceContainerPublishPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainerPublishPort)(nil)).Elem()
}

func (i InstanceContainerPublishPortArgs) ToInstanceContainerPublishPortOutput() InstanceContainerPublishPortOutput {
	return i.ToInstanceContainerPublishPortOutputWithContext(context.Background())
}

func (i InstanceContainerPublishPortArgs) ToInstanceContainerPublishPortOutputWithContext(ctx context.Context) InstanceContainerPublishPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerPublishPortOutput)
}

// InstanceContainerPublishPortArrayInput is an input type that accepts InstanceContainerPublishPortArray and InstanceContainerPublishPortArrayOutput values.
// You can construct a concrete instance of `InstanceContainerPublishPortArrayInput` via:
//
//	InstanceContainerPublishPortArray{ InstanceContainerPublishPortArgs{...} }
type InstanceContainerPublishPortArrayInput interface {
	pulumi.Input

	ToInstanceContainerPublishPortArrayOutput() InstanceContainerPublishPortArrayOutput
	ToInstanceContainerPublishPortArrayOutputWithContext(context.Context) InstanceContainerPublishPortArrayOutput
}

type InstanceContainerPublishPortArray []InstanceContainerPublishPortInput

func (InstanceContainerPublishPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainerPublishPort)(nil)).Elem()
}

func (i InstanceContainerPublishPortArray) ToInstanceContainerPublishPortArrayOutput() InstanceContainerPublishPortArrayOutput {
	return i.ToInstanceContainerPublishPortArrayOutputWithContext(context.Background())
}

func (i InstanceContainerPublishPortArray) ToInstanceContainerPublishPortArrayOutputWithContext(ctx context.Context) InstanceContainerPublishPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerPublishPortArrayOutput)
}

type InstanceContainerPublishPortOutput struct{ *pulumi.OutputState }

func (InstanceContainerPublishPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainerPublishPort)(nil)).Elem()
}

func (o InstanceContainerPublishPortOutput) ToInstanceContainerPublishPortOutput() InstanceContainerPublishPortOutput {
	return o
}

func (o InstanceContainerPublishPortOutput) ToInstanceContainerPublishPortOutputWithContext(ctx context.Context) InstanceContainerPublishPortOutput {
	return o
}

// Container port.
func (o InstanceContainerPublishPortOutput) ContainerPort() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceContainerPublishPort) int { return v.ContainerPort }).(pulumi.IntOutput)
}

// Host port.
func (o InstanceContainerPublishPortOutput) HostPort() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceContainerPublishPort) int { return v.HostPort }).(pulumi.IntOutput)
}

// External IP. It defaults to 0.0.0.0.
func (o InstanceContainerPublishPortOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceContainerPublishPort) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

// The protocol defaults to tcp. Valid values: tcp, udp and sctp.
func (o InstanceContainerPublishPortOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceContainerPublishPort) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

type InstanceContainerPublishPortArrayOutput struct{ *pulumi.OutputState }

func (InstanceContainerPublishPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainerPublishPort)(nil)).Elem()
}

func (o InstanceContainerPublishPortArrayOutput) ToInstanceContainerPublishPortArrayOutput() InstanceContainerPublishPortArrayOutput {
	return o
}

func (o InstanceContainerPublishPortArrayOutput) ToInstanceContainerPublishPortArrayOutputWithContext(ctx context.Context) InstanceContainerPublishPortArrayOutput {
	return o
}

func (o InstanceContainerPublishPortArrayOutput) Index(i pulumi.IntInput) InstanceContainerPublishPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceContainerPublishPort {
		return vs[0].([]InstanceContainerPublishPort)[vs[1].(int)]
	}).(InstanceContainerPublishPortOutput)
}

type InstanceContainerVolume struct {
	// Container path.
	ContainerPath string `pulumi:"containerPath"`
	// Host path.
	HostPath string `pulumi:"hostPath"`
}

// InstanceContainerVolumeInput is an input type that accepts InstanceContainerVolumeArgs and InstanceContainerVolumeOutput values.
// You can construct a concrete instance of `InstanceContainerVolumeInput` via:
//
//	InstanceContainerVolumeArgs{...}
type InstanceContainerVolumeInput interface {
	pulumi.Input

	ToInstanceContainerVolumeOutput() InstanceContainerVolumeOutput
	ToInstanceContainerVolumeOutputWithContext(context.Context) InstanceContainerVolumeOutput
}

type InstanceContainerVolumeArgs struct {
	// Container path.
	ContainerPath pulumi.StringInput `pulumi:"containerPath"`
	// Host path.
	HostPath pulumi.StringInput `pulumi:"hostPath"`
}

func (InstanceContainerVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainerVolume)(nil)).Elem()
}

func (i InstanceContainerVolumeArgs) ToInstanceContainerVolumeOutput() InstanceContainerVolumeOutput {
	return i.ToInstanceContainerVolumeOutputWithContext(context.Background())
}

func (i InstanceContainerVolumeArgs) ToInstanceContainerVolumeOutputWithContext(ctx context.Context) InstanceContainerVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerVolumeOutput)
}

// InstanceContainerVolumeArrayInput is an input type that accepts InstanceContainerVolumeArray and InstanceContainerVolumeArrayOutput values.
// You can construct a concrete instance of `InstanceContainerVolumeArrayInput` via:
//
//	InstanceContainerVolumeArray{ InstanceContainerVolumeArgs{...} }
type InstanceContainerVolumeArrayInput interface {
	pulumi.Input

	ToInstanceContainerVolumeArrayOutput() InstanceContainerVolumeArrayOutput
	ToInstanceContainerVolumeArrayOutputWithContext(context.Context) InstanceContainerVolumeArrayOutput
}

type InstanceContainerVolumeArray []InstanceContainerVolumeInput

func (InstanceContainerVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainerVolume)(nil)).Elem()
}

func (i InstanceContainerVolumeArray) ToInstanceContainerVolumeArrayOutput() InstanceContainerVolumeArrayOutput {
	return i.ToInstanceContainerVolumeArrayOutputWithContext(context.Background())
}

func (i InstanceContainerVolumeArray) ToInstanceContainerVolumeArrayOutputWithContext(ctx context.Context) InstanceContainerVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceContainerVolumeArrayOutput)
}

type InstanceContainerVolumeOutput struct{ *pulumi.OutputState }

func (InstanceContainerVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceContainerVolume)(nil)).Elem()
}

func (o InstanceContainerVolumeOutput) ToInstanceContainerVolumeOutput() InstanceContainerVolumeOutput {
	return o
}

func (o InstanceContainerVolumeOutput) ToInstanceContainerVolumeOutputWithContext(ctx context.Context) InstanceContainerVolumeOutput {
	return o
}

// Container path.
func (o InstanceContainerVolumeOutput) ContainerPath() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceContainerVolume) string { return v.ContainerPath }).(pulumi.StringOutput)
}

// Host path.
func (o InstanceContainerVolumeOutput) HostPath() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceContainerVolume) string { return v.HostPath }).(pulumi.StringOutput)
}

type InstanceContainerVolumeArrayOutput struct{ *pulumi.OutputState }

func (InstanceContainerVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceContainerVolume)(nil)).Elem()
}

func (o InstanceContainerVolumeArrayOutput) ToInstanceContainerVolumeArrayOutput() InstanceContainerVolumeArrayOutput {
	return o
}

func (o InstanceContainerVolumeArrayOutput) ToInstanceContainerVolumeArrayOutputWithContext(ctx context.Context) InstanceContainerVolumeArrayOutput {
	return o
}

func (o InstanceContainerVolumeArrayOutput) Index(i pulumi.IntInput) InstanceContainerVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceContainerVolume {
		return vs[0].([]InstanceContainerVolume)[vs[1].(int)]
	}).(InstanceContainerVolumeOutput)
}

type InstanceLoginConfiguration struct {
	// whether auto generate password. if false, need set password.
	AutoGeneratePassword string `pulumi:"autoGeneratePassword"`
	// Login password.
	Password *string `pulumi:"password"`
}

// InstanceLoginConfigurationInput is an input type that accepts InstanceLoginConfigurationArgs and InstanceLoginConfigurationOutput values.
// You can construct a concrete instance of `InstanceLoginConfigurationInput` via:
//
//	InstanceLoginConfigurationArgs{...}
type InstanceLoginConfigurationInput interface {
	pulumi.Input

	ToInstanceLoginConfigurationOutput() InstanceLoginConfigurationOutput
	ToInstanceLoginConfigurationOutputWithContext(context.Context) InstanceLoginConfigurationOutput
}

type InstanceLoginConfigurationArgs struct {
	// whether auto generate password. if false, need set password.
	AutoGeneratePassword pulumi.StringInput `pulumi:"autoGeneratePassword"`
	// Login password.
	Password pulumi.StringPtrInput `pulumi:"password"`
}

func (InstanceLoginConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceLoginConfiguration)(nil)).Elem()
}

func (i InstanceLoginConfigurationArgs) ToInstanceLoginConfigurationOutput() InstanceLoginConfigurationOutput {
	return i.ToInstanceLoginConfigurationOutputWithContext(context.Background())
}

func (i InstanceLoginConfigurationArgs) ToInstanceLoginConfigurationOutputWithContext(ctx context.Context) InstanceLoginConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceLoginConfigurationOutput)
}

func (i InstanceLoginConfigurationArgs) ToInstanceLoginConfigurationPtrOutput() InstanceLoginConfigurationPtrOutput {
	return i.ToInstanceLoginConfigurationPtrOutputWithContext(context.Background())
}

func (i InstanceLoginConfigurationArgs) ToInstanceLoginConfigurationPtrOutputWithContext(ctx context.Context) InstanceLoginConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceLoginConfigurationOutput).ToInstanceLoginConfigurationPtrOutputWithContext(ctx)
}

// InstanceLoginConfigurationPtrInput is an input type that accepts InstanceLoginConfigurationArgs, InstanceLoginConfigurationPtr and InstanceLoginConfigurationPtrOutput values.
// You can construct a concrete instance of `InstanceLoginConfigurationPtrInput` via:
//
//	        InstanceLoginConfigurationArgs{...}
//
//	or:
//
//	        nil
type InstanceLoginConfigurationPtrInput interface {
	pulumi.Input

	ToInstanceLoginConfigurationPtrOutput() InstanceLoginConfigurationPtrOutput
	ToInstanceLoginConfigurationPtrOutputWithContext(context.Context) InstanceLoginConfigurationPtrOutput
}

type instanceLoginConfigurationPtrType InstanceLoginConfigurationArgs

func InstanceLoginConfigurationPtr(v *InstanceLoginConfigurationArgs) InstanceLoginConfigurationPtrInput {
	return (*instanceLoginConfigurationPtrType)(v)
}

func (*instanceLoginConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceLoginConfiguration)(nil)).Elem()
}

func (i *instanceLoginConfigurationPtrType) ToInstanceLoginConfigurationPtrOutput() InstanceLoginConfigurationPtrOutput {
	return i.ToInstanceLoginConfigurationPtrOutputWithContext(context.Background())
}

func (i *instanceLoginConfigurationPtrType) ToInstanceLoginConfigurationPtrOutputWithContext(ctx context.Context) InstanceLoginConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceLoginConfigurationPtrOutput)
}

type InstanceLoginConfigurationOutput struct{ *pulumi.OutputState }

func (InstanceLoginConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceLoginConfiguration)(nil)).Elem()
}

func (o InstanceLoginConfigurationOutput) ToInstanceLoginConfigurationOutput() InstanceLoginConfigurationOutput {
	return o
}

func (o InstanceLoginConfigurationOutput) ToInstanceLoginConfigurationOutputWithContext(ctx context.Context) InstanceLoginConfigurationOutput {
	return o
}

func (o InstanceLoginConfigurationOutput) ToInstanceLoginConfigurationPtrOutput() InstanceLoginConfigurationPtrOutput {
	return o.ToInstanceLoginConfigurationPtrOutputWithContext(context.Background())
}

func (o InstanceLoginConfigurationOutput) ToInstanceLoginConfigurationPtrOutputWithContext(ctx context.Context) InstanceLoginConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceLoginConfiguration) *InstanceLoginConfiguration {
		return &v
	}).(InstanceLoginConfigurationPtrOutput)
}

// whether auto generate password. if false, need set password.
func (o InstanceLoginConfigurationOutput) AutoGeneratePassword() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceLoginConfiguration) string { return v.AutoGeneratePassword }).(pulumi.StringOutput)
}

// Login password.
func (o InstanceLoginConfigurationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceLoginConfiguration) *string { return v.Password }).(pulumi.StringPtrOutput)
}

type InstanceLoginConfigurationPtrOutput struct{ *pulumi.OutputState }

func (InstanceLoginConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceLoginConfiguration)(nil)).Elem()
}

func (o InstanceLoginConfigurationPtrOutput) ToInstanceLoginConfigurationPtrOutput() InstanceLoginConfigurationPtrOutput {
	return o
}

func (o InstanceLoginConfigurationPtrOutput) ToInstanceLoginConfigurationPtrOutputWithContext(ctx context.Context) InstanceLoginConfigurationPtrOutput {
	return o
}

func (o InstanceLoginConfigurationPtrOutput) Elem() InstanceLoginConfigurationOutput {
	return o.ApplyT(func(v *InstanceLoginConfiguration) InstanceLoginConfiguration {
		if v != nil {
			return *v
		}
		var ret InstanceLoginConfiguration
		return ret
	}).(InstanceLoginConfigurationOutput)
}

// whether auto generate password. if false, need set password.
func (o InstanceLoginConfigurationPtrOutput) AutoGeneratePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceLoginConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.AutoGeneratePassword
	}).(pulumi.StringPtrOutput)
}

// Login password.
func (o InstanceLoginConfigurationPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceLoginConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerInput)(nil)).Elem(), InstanceContainerArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerArrayInput)(nil)).Elem(), InstanceContainerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerEnvInput)(nil)).Elem(), InstanceContainerEnvArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerEnvArrayInput)(nil)).Elem(), InstanceContainerEnvArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerPublishPortInput)(nil)).Elem(), InstanceContainerPublishPortArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerPublishPortArrayInput)(nil)).Elem(), InstanceContainerPublishPortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerVolumeInput)(nil)).Elem(), InstanceContainerVolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceContainerVolumeArrayInput)(nil)).Elem(), InstanceContainerVolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceLoginConfigurationInput)(nil)).Elem(), InstanceLoginConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceLoginConfigurationPtrInput)(nil)).Elem(), InstanceLoginConfigurationArgs{})
	pulumi.RegisterOutputType(InstanceContainerOutput{})
	pulumi.RegisterOutputType(InstanceContainerArrayOutput{})
	pulumi.RegisterOutputType(InstanceContainerEnvOutput{})
	pulumi.RegisterOutputType(InstanceContainerEnvArrayOutput{})
	pulumi.RegisterOutputType(InstanceContainerPublishPortOutput{})
	pulumi.RegisterOutputType(InstanceContainerPublishPortArrayOutput{})
	pulumi.RegisterOutputType(InstanceContainerVolumeOutput{})
	pulumi.RegisterOutputType(InstanceContainerVolumeArrayOutput{})
	pulumi.RegisterOutputType(InstanceLoginConfigurationOutput{})
	pulumi.RegisterOutputType(InstanceLoginConfigurationPtrOutput{})
}
