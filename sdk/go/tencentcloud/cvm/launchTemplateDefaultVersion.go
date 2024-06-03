// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cvm

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a cvm launchTemplateDefaultVersion
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cvm"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cvm.NewLaunchTemplateDefaultVersion(ctx, "launchTemplateDefaultVersion", &Cvm.LaunchTemplateDefaultVersionArgs{
//				DefaultVersion:   pulumi.Int(2),
//				LaunchTemplateId: pulumi.String("lt-34vaef8fe"),
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
// cvm launch_template_default_version can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion launch_template_default_version launch_template_id
// ```
type LaunchTemplateDefaultVersion struct {
	pulumi.CustomResourceState

	// The number of the version that you want to set as the default version.
	DefaultVersion pulumi.IntOutput `pulumi:"defaultVersion"`
	// Instance launch template ID.
	LaunchTemplateId pulumi.StringOutput `pulumi:"launchTemplateId"`
}

// NewLaunchTemplateDefaultVersion registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplateDefaultVersion(ctx *pulumi.Context,
	name string, args *LaunchTemplateDefaultVersionArgs, opts ...pulumi.ResourceOption) (*LaunchTemplateDefaultVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultVersion == nil {
		return nil, errors.New("invalid value for required argument 'DefaultVersion'")
	}
	if args.LaunchTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'LaunchTemplateId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LaunchTemplateDefaultVersion
	err := ctx.RegisterResource("tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplateDefaultVersion gets an existing LaunchTemplateDefaultVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplateDefaultVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateDefaultVersionState, opts ...pulumi.ResourceOption) (*LaunchTemplateDefaultVersion, error) {
	var resource LaunchTemplateDefaultVersion
	err := ctx.ReadResource("tencentcloud:Cvm/launchTemplateDefaultVersion:LaunchTemplateDefaultVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplateDefaultVersion resources.
type launchTemplateDefaultVersionState struct {
	// The number of the version that you want to set as the default version.
	DefaultVersion *int `pulumi:"defaultVersion"`
	// Instance launch template ID.
	LaunchTemplateId *string `pulumi:"launchTemplateId"`
}

type LaunchTemplateDefaultVersionState struct {
	// The number of the version that you want to set as the default version.
	DefaultVersion pulumi.IntPtrInput
	// Instance launch template ID.
	LaunchTemplateId pulumi.StringPtrInput
}

func (LaunchTemplateDefaultVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateDefaultVersionState)(nil)).Elem()
}

type launchTemplateDefaultVersionArgs struct {
	// The number of the version that you want to set as the default version.
	DefaultVersion int `pulumi:"defaultVersion"`
	// Instance launch template ID.
	LaunchTemplateId string `pulumi:"launchTemplateId"`
}

// The set of arguments for constructing a LaunchTemplateDefaultVersion resource.
type LaunchTemplateDefaultVersionArgs struct {
	// The number of the version that you want to set as the default version.
	DefaultVersion pulumi.IntInput
	// Instance launch template ID.
	LaunchTemplateId pulumi.StringInput
}

func (LaunchTemplateDefaultVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateDefaultVersionArgs)(nil)).Elem()
}

type LaunchTemplateDefaultVersionInput interface {
	pulumi.Input

	ToLaunchTemplateDefaultVersionOutput() LaunchTemplateDefaultVersionOutput
	ToLaunchTemplateDefaultVersionOutputWithContext(ctx context.Context) LaunchTemplateDefaultVersionOutput
}

func (*LaunchTemplateDefaultVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplateDefaultVersion)(nil)).Elem()
}

func (i *LaunchTemplateDefaultVersion) ToLaunchTemplateDefaultVersionOutput() LaunchTemplateDefaultVersionOutput {
	return i.ToLaunchTemplateDefaultVersionOutputWithContext(context.Background())
}

func (i *LaunchTemplateDefaultVersion) ToLaunchTemplateDefaultVersionOutputWithContext(ctx context.Context) LaunchTemplateDefaultVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateDefaultVersionOutput)
}

// LaunchTemplateDefaultVersionArrayInput is an input type that accepts LaunchTemplateDefaultVersionArray and LaunchTemplateDefaultVersionArrayOutput values.
// You can construct a concrete instance of `LaunchTemplateDefaultVersionArrayInput` via:
//
//	LaunchTemplateDefaultVersionArray{ LaunchTemplateDefaultVersionArgs{...} }
type LaunchTemplateDefaultVersionArrayInput interface {
	pulumi.Input

	ToLaunchTemplateDefaultVersionArrayOutput() LaunchTemplateDefaultVersionArrayOutput
	ToLaunchTemplateDefaultVersionArrayOutputWithContext(context.Context) LaunchTemplateDefaultVersionArrayOutput
}

type LaunchTemplateDefaultVersionArray []LaunchTemplateDefaultVersionInput

func (LaunchTemplateDefaultVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LaunchTemplateDefaultVersion)(nil)).Elem()
}

func (i LaunchTemplateDefaultVersionArray) ToLaunchTemplateDefaultVersionArrayOutput() LaunchTemplateDefaultVersionArrayOutput {
	return i.ToLaunchTemplateDefaultVersionArrayOutputWithContext(context.Background())
}

func (i LaunchTemplateDefaultVersionArray) ToLaunchTemplateDefaultVersionArrayOutputWithContext(ctx context.Context) LaunchTemplateDefaultVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateDefaultVersionArrayOutput)
}

// LaunchTemplateDefaultVersionMapInput is an input type that accepts LaunchTemplateDefaultVersionMap and LaunchTemplateDefaultVersionMapOutput values.
// You can construct a concrete instance of `LaunchTemplateDefaultVersionMapInput` via:
//
//	LaunchTemplateDefaultVersionMap{ "key": LaunchTemplateDefaultVersionArgs{...} }
type LaunchTemplateDefaultVersionMapInput interface {
	pulumi.Input

	ToLaunchTemplateDefaultVersionMapOutput() LaunchTemplateDefaultVersionMapOutput
	ToLaunchTemplateDefaultVersionMapOutputWithContext(context.Context) LaunchTemplateDefaultVersionMapOutput
}

type LaunchTemplateDefaultVersionMap map[string]LaunchTemplateDefaultVersionInput

func (LaunchTemplateDefaultVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LaunchTemplateDefaultVersion)(nil)).Elem()
}

func (i LaunchTemplateDefaultVersionMap) ToLaunchTemplateDefaultVersionMapOutput() LaunchTemplateDefaultVersionMapOutput {
	return i.ToLaunchTemplateDefaultVersionMapOutputWithContext(context.Background())
}

func (i LaunchTemplateDefaultVersionMap) ToLaunchTemplateDefaultVersionMapOutputWithContext(ctx context.Context) LaunchTemplateDefaultVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LaunchTemplateDefaultVersionMapOutput)
}

type LaunchTemplateDefaultVersionOutput struct{ *pulumi.OutputState }

func (LaunchTemplateDefaultVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LaunchTemplateDefaultVersion)(nil)).Elem()
}

func (o LaunchTemplateDefaultVersionOutput) ToLaunchTemplateDefaultVersionOutput() LaunchTemplateDefaultVersionOutput {
	return o
}

func (o LaunchTemplateDefaultVersionOutput) ToLaunchTemplateDefaultVersionOutputWithContext(ctx context.Context) LaunchTemplateDefaultVersionOutput {
	return o
}

// The number of the version that you want to set as the default version.
func (o LaunchTemplateDefaultVersionOutput) DefaultVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *LaunchTemplateDefaultVersion) pulumi.IntOutput { return v.DefaultVersion }).(pulumi.IntOutput)
}

// Instance launch template ID.
func (o LaunchTemplateDefaultVersionOutput) LaunchTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *LaunchTemplateDefaultVersion) pulumi.StringOutput { return v.LaunchTemplateId }).(pulumi.StringOutput)
}

type LaunchTemplateDefaultVersionArrayOutput struct{ *pulumi.OutputState }

func (LaunchTemplateDefaultVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LaunchTemplateDefaultVersion)(nil)).Elem()
}

func (o LaunchTemplateDefaultVersionArrayOutput) ToLaunchTemplateDefaultVersionArrayOutput() LaunchTemplateDefaultVersionArrayOutput {
	return o
}

func (o LaunchTemplateDefaultVersionArrayOutput) ToLaunchTemplateDefaultVersionArrayOutputWithContext(ctx context.Context) LaunchTemplateDefaultVersionArrayOutput {
	return o
}

func (o LaunchTemplateDefaultVersionArrayOutput) Index(i pulumi.IntInput) LaunchTemplateDefaultVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LaunchTemplateDefaultVersion {
		return vs[0].([]*LaunchTemplateDefaultVersion)[vs[1].(int)]
	}).(LaunchTemplateDefaultVersionOutput)
}

type LaunchTemplateDefaultVersionMapOutput struct{ *pulumi.OutputState }

func (LaunchTemplateDefaultVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LaunchTemplateDefaultVersion)(nil)).Elem()
}

func (o LaunchTemplateDefaultVersionMapOutput) ToLaunchTemplateDefaultVersionMapOutput() LaunchTemplateDefaultVersionMapOutput {
	return o
}

func (o LaunchTemplateDefaultVersionMapOutput) ToLaunchTemplateDefaultVersionMapOutputWithContext(ctx context.Context) LaunchTemplateDefaultVersionMapOutput {
	return o
}

func (o LaunchTemplateDefaultVersionMapOutput) MapIndex(k pulumi.StringInput) LaunchTemplateDefaultVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LaunchTemplateDefaultVersion {
		return vs[0].(map[string]*LaunchTemplateDefaultVersion)[vs[1].(string)]
	}).(LaunchTemplateDefaultVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateDefaultVersionInput)(nil)).Elem(), &LaunchTemplateDefaultVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateDefaultVersionArrayInput)(nil)).Elem(), LaunchTemplateDefaultVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LaunchTemplateDefaultVersionMapInput)(nil)).Elem(), LaunchTemplateDefaultVersionMap{})
	pulumi.RegisterOutputType(LaunchTemplateDefaultVersionOutput{})
	pulumi.RegisterOutputType(LaunchTemplateDefaultVersionArrayOutput{})
	pulumi.RegisterOutputType(LaunchTemplateDefaultVersionMapOutput{})
}
