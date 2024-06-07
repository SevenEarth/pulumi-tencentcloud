// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tsf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tsf applicationFileConfigRelease
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tsf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tsf.NewApplicationFileConfigRelease(ctx, "applicationFileConfigRelease", &Tsf.ApplicationFileConfigReleaseArgs{
//				ConfigId:    pulumi.String("dcfg-f-123456"),
//				GroupId:     pulumi.String("group-123456"),
//				ReleaseDesc: pulumi.String("product release"),
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
// tsf applicationfile_config_release can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease application_file_config_release application_file_config_release_id
// ```
type ApplicationFileConfigRelease struct {
	pulumi.CustomResourceState

	// File config id.
	ConfigId pulumi.StringOutput `pulumi:"configId"`
	// Group Id.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// release Description.
	ReleaseDesc pulumi.StringPtrOutput `pulumi:"releaseDesc"`
}

// NewApplicationFileConfigRelease registers a new resource with the given unique name, arguments, and options.
func NewApplicationFileConfigRelease(ctx *pulumi.Context,
	name string, args *ApplicationFileConfigReleaseArgs, opts ...pulumi.ResourceOption) (*ApplicationFileConfigRelease, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationFileConfigRelease
	err := ctx.RegisterResource("tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationFileConfigRelease gets an existing ApplicationFileConfigRelease resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationFileConfigRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationFileConfigReleaseState, opts ...pulumi.ResourceOption) (*ApplicationFileConfigRelease, error) {
	var resource ApplicationFileConfigRelease
	err := ctx.ReadResource("tencentcloud:Tsf/applicationFileConfigRelease:ApplicationFileConfigRelease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationFileConfigRelease resources.
type applicationFileConfigReleaseState struct {
	// File config id.
	ConfigId *string `pulumi:"configId"`
	// Group Id.
	GroupId *string `pulumi:"groupId"`
	// release Description.
	ReleaseDesc *string `pulumi:"releaseDesc"`
}

type ApplicationFileConfigReleaseState struct {
	// File config id.
	ConfigId pulumi.StringPtrInput
	// Group Id.
	GroupId pulumi.StringPtrInput
	// release Description.
	ReleaseDesc pulumi.StringPtrInput
}

func (ApplicationFileConfigReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationFileConfigReleaseState)(nil)).Elem()
}

type applicationFileConfigReleaseArgs struct {
	// File config id.
	ConfigId string `pulumi:"configId"`
	// Group Id.
	GroupId string `pulumi:"groupId"`
	// release Description.
	ReleaseDesc *string `pulumi:"releaseDesc"`
}

// The set of arguments for constructing a ApplicationFileConfigRelease resource.
type ApplicationFileConfigReleaseArgs struct {
	// File config id.
	ConfigId pulumi.StringInput
	// Group Id.
	GroupId pulumi.StringInput
	// release Description.
	ReleaseDesc pulumi.StringPtrInput
}

func (ApplicationFileConfigReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationFileConfigReleaseArgs)(nil)).Elem()
}

type ApplicationFileConfigReleaseInput interface {
	pulumi.Input

	ToApplicationFileConfigReleaseOutput() ApplicationFileConfigReleaseOutput
	ToApplicationFileConfigReleaseOutputWithContext(ctx context.Context) ApplicationFileConfigReleaseOutput
}

func (*ApplicationFileConfigRelease) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFileConfigRelease)(nil)).Elem()
}

func (i *ApplicationFileConfigRelease) ToApplicationFileConfigReleaseOutput() ApplicationFileConfigReleaseOutput {
	return i.ToApplicationFileConfigReleaseOutputWithContext(context.Background())
}

func (i *ApplicationFileConfigRelease) ToApplicationFileConfigReleaseOutputWithContext(ctx context.Context) ApplicationFileConfigReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFileConfigReleaseOutput)
}

// ApplicationFileConfigReleaseArrayInput is an input type that accepts ApplicationFileConfigReleaseArray and ApplicationFileConfigReleaseArrayOutput values.
// You can construct a concrete instance of `ApplicationFileConfigReleaseArrayInput` via:
//
//	ApplicationFileConfigReleaseArray{ ApplicationFileConfigReleaseArgs{...} }
type ApplicationFileConfigReleaseArrayInput interface {
	pulumi.Input

	ToApplicationFileConfigReleaseArrayOutput() ApplicationFileConfigReleaseArrayOutput
	ToApplicationFileConfigReleaseArrayOutputWithContext(context.Context) ApplicationFileConfigReleaseArrayOutput
}

type ApplicationFileConfigReleaseArray []ApplicationFileConfigReleaseInput

func (ApplicationFileConfigReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationFileConfigRelease)(nil)).Elem()
}

func (i ApplicationFileConfigReleaseArray) ToApplicationFileConfigReleaseArrayOutput() ApplicationFileConfigReleaseArrayOutput {
	return i.ToApplicationFileConfigReleaseArrayOutputWithContext(context.Background())
}

func (i ApplicationFileConfigReleaseArray) ToApplicationFileConfigReleaseArrayOutputWithContext(ctx context.Context) ApplicationFileConfigReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFileConfigReleaseArrayOutput)
}

// ApplicationFileConfigReleaseMapInput is an input type that accepts ApplicationFileConfigReleaseMap and ApplicationFileConfigReleaseMapOutput values.
// You can construct a concrete instance of `ApplicationFileConfigReleaseMapInput` via:
//
//	ApplicationFileConfigReleaseMap{ "key": ApplicationFileConfigReleaseArgs{...} }
type ApplicationFileConfigReleaseMapInput interface {
	pulumi.Input

	ToApplicationFileConfigReleaseMapOutput() ApplicationFileConfigReleaseMapOutput
	ToApplicationFileConfigReleaseMapOutputWithContext(context.Context) ApplicationFileConfigReleaseMapOutput
}

type ApplicationFileConfigReleaseMap map[string]ApplicationFileConfigReleaseInput

func (ApplicationFileConfigReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationFileConfigRelease)(nil)).Elem()
}

func (i ApplicationFileConfigReleaseMap) ToApplicationFileConfigReleaseMapOutput() ApplicationFileConfigReleaseMapOutput {
	return i.ToApplicationFileConfigReleaseMapOutputWithContext(context.Background())
}

func (i ApplicationFileConfigReleaseMap) ToApplicationFileConfigReleaseMapOutputWithContext(ctx context.Context) ApplicationFileConfigReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationFileConfigReleaseMapOutput)
}

type ApplicationFileConfigReleaseOutput struct{ *pulumi.OutputState }

func (ApplicationFileConfigReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationFileConfigRelease)(nil)).Elem()
}

func (o ApplicationFileConfigReleaseOutput) ToApplicationFileConfigReleaseOutput() ApplicationFileConfigReleaseOutput {
	return o
}

func (o ApplicationFileConfigReleaseOutput) ToApplicationFileConfigReleaseOutputWithContext(ctx context.Context) ApplicationFileConfigReleaseOutput {
	return o
}

// File config id.
func (o ApplicationFileConfigReleaseOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfigRelease) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// Group Id.
func (o ApplicationFileConfigReleaseOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationFileConfigRelease) pulumi.StringOutput { return v.GroupId }).(pulumi.StringOutput)
}

// release Description.
func (o ApplicationFileConfigReleaseOutput) ReleaseDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationFileConfigRelease) pulumi.StringPtrOutput { return v.ReleaseDesc }).(pulumi.StringPtrOutput)
}

type ApplicationFileConfigReleaseArrayOutput struct{ *pulumi.OutputState }

func (ApplicationFileConfigReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationFileConfigRelease)(nil)).Elem()
}

func (o ApplicationFileConfigReleaseArrayOutput) ToApplicationFileConfigReleaseArrayOutput() ApplicationFileConfigReleaseArrayOutput {
	return o
}

func (o ApplicationFileConfigReleaseArrayOutput) ToApplicationFileConfigReleaseArrayOutputWithContext(ctx context.Context) ApplicationFileConfigReleaseArrayOutput {
	return o
}

func (o ApplicationFileConfigReleaseArrayOutput) Index(i pulumi.IntInput) ApplicationFileConfigReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationFileConfigRelease {
		return vs[0].([]*ApplicationFileConfigRelease)[vs[1].(int)]
	}).(ApplicationFileConfigReleaseOutput)
}

type ApplicationFileConfigReleaseMapOutput struct{ *pulumi.OutputState }

func (ApplicationFileConfigReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationFileConfigRelease)(nil)).Elem()
}

func (o ApplicationFileConfigReleaseMapOutput) ToApplicationFileConfigReleaseMapOutput() ApplicationFileConfigReleaseMapOutput {
	return o
}

func (o ApplicationFileConfigReleaseMapOutput) ToApplicationFileConfigReleaseMapOutputWithContext(ctx context.Context) ApplicationFileConfigReleaseMapOutput {
	return o
}

func (o ApplicationFileConfigReleaseMapOutput) MapIndex(k pulumi.StringInput) ApplicationFileConfigReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationFileConfigRelease {
		return vs[0].(map[string]*ApplicationFileConfigRelease)[vs[1].(string)]
	}).(ApplicationFileConfigReleaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFileConfigReleaseInput)(nil)).Elem(), &ApplicationFileConfigRelease{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFileConfigReleaseArrayInput)(nil)).Elem(), ApplicationFileConfigReleaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationFileConfigReleaseMapInput)(nil)).Elem(), ApplicationFileConfigReleaseMap{})
	pulumi.RegisterOutputType(ApplicationFileConfigReleaseOutput{})
	pulumi.RegisterOutputType(ApplicationFileConfigReleaseArrayOutput{})
	pulumi.RegisterOutputType(ApplicationFileConfigReleaseMapOutput{})
}
