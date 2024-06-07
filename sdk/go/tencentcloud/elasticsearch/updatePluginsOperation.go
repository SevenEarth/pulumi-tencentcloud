// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticsearch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to update elasticsearch plugins
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Elasticsearch"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Elasticsearch.NewUpdatePluginsOperation(ctx, "updatePluginsOperation", &Elasticsearch.UpdatePluginsOperationArgs{
//				ForceRestart: pulumi.Bool(false),
//				ForceUpdate:  pulumi.Bool(true),
//				InstallPluginLists: pulumi.StringArray{
//					pulumi.String("analysis-pinyin"),
//				},
//				InstanceId: pulumi.String("es-xxxxxx"),
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
type UpdatePluginsOperation struct {
	pulumi.CustomResourceState

	// Whether to force a restart. Default is false.
	ForceRestart pulumi.BoolPtrOutput `pulumi:"forceRestart"`
	// Whether to reinstall, default value false.
	ForceUpdate pulumi.BoolPtrOutput `pulumi:"forceUpdate"`
	// List of plugins that need to be installed.
	InstallPluginLists pulumi.StringArrayOutput `pulumi:"installPluginLists"`
	// Instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Plugin type. 0: system plugin.
	PluginType pulumi.IntPtrOutput `pulumi:"pluginType"`
	// List of plugins that need to be uninstalled.
	RemovePluginLists pulumi.StringArrayOutput `pulumi:"removePluginLists"`
}

// NewUpdatePluginsOperation registers a new resource with the given unique name, arguments, and options.
func NewUpdatePluginsOperation(ctx *pulumi.Context,
	name string, args *UpdatePluginsOperationArgs, opts ...pulumi.ResourceOption) (*UpdatePluginsOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UpdatePluginsOperation
	err := ctx.RegisterResource("tencentcloud:Elasticsearch/updatePluginsOperation:UpdatePluginsOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpdatePluginsOperation gets an existing UpdatePluginsOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpdatePluginsOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdatePluginsOperationState, opts ...pulumi.ResourceOption) (*UpdatePluginsOperation, error) {
	var resource UpdatePluginsOperation
	err := ctx.ReadResource("tencentcloud:Elasticsearch/updatePluginsOperation:UpdatePluginsOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UpdatePluginsOperation resources.
type updatePluginsOperationState struct {
	// Whether to force a restart. Default is false.
	ForceRestart *bool `pulumi:"forceRestart"`
	// Whether to reinstall, default value false.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// List of plugins that need to be installed.
	InstallPluginLists []string `pulumi:"installPluginLists"`
	// Instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Plugin type. 0: system plugin.
	PluginType *int `pulumi:"pluginType"`
	// List of plugins that need to be uninstalled.
	RemovePluginLists []string `pulumi:"removePluginLists"`
}

type UpdatePluginsOperationState struct {
	// Whether to force a restart. Default is false.
	ForceRestart pulumi.BoolPtrInput
	// Whether to reinstall, default value false.
	ForceUpdate pulumi.BoolPtrInput
	// List of plugins that need to be installed.
	InstallPluginLists pulumi.StringArrayInput
	// Instance id.
	InstanceId pulumi.StringPtrInput
	// Plugin type. 0: system plugin.
	PluginType pulumi.IntPtrInput
	// List of plugins that need to be uninstalled.
	RemovePluginLists pulumi.StringArrayInput
}

func (UpdatePluginsOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*updatePluginsOperationState)(nil)).Elem()
}

type updatePluginsOperationArgs struct {
	// Whether to force a restart. Default is false.
	ForceRestart *bool `pulumi:"forceRestart"`
	// Whether to reinstall, default value false.
	ForceUpdate *bool `pulumi:"forceUpdate"`
	// List of plugins that need to be installed.
	InstallPluginLists []string `pulumi:"installPluginLists"`
	// Instance id.
	InstanceId string `pulumi:"instanceId"`
	// Plugin type. 0: system plugin.
	PluginType *int `pulumi:"pluginType"`
	// List of plugins that need to be uninstalled.
	RemovePluginLists []string `pulumi:"removePluginLists"`
}

// The set of arguments for constructing a UpdatePluginsOperation resource.
type UpdatePluginsOperationArgs struct {
	// Whether to force a restart. Default is false.
	ForceRestart pulumi.BoolPtrInput
	// Whether to reinstall, default value false.
	ForceUpdate pulumi.BoolPtrInput
	// List of plugins that need to be installed.
	InstallPluginLists pulumi.StringArrayInput
	// Instance id.
	InstanceId pulumi.StringInput
	// Plugin type. 0: system plugin.
	PluginType pulumi.IntPtrInput
	// List of plugins that need to be uninstalled.
	RemovePluginLists pulumi.StringArrayInput
}

func (UpdatePluginsOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updatePluginsOperationArgs)(nil)).Elem()
}

type UpdatePluginsOperationInput interface {
	pulumi.Input

	ToUpdatePluginsOperationOutput() UpdatePluginsOperationOutput
	ToUpdatePluginsOperationOutputWithContext(ctx context.Context) UpdatePluginsOperationOutput
}

func (*UpdatePluginsOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdatePluginsOperation)(nil)).Elem()
}

func (i *UpdatePluginsOperation) ToUpdatePluginsOperationOutput() UpdatePluginsOperationOutput {
	return i.ToUpdatePluginsOperationOutputWithContext(context.Background())
}

func (i *UpdatePluginsOperation) ToUpdatePluginsOperationOutputWithContext(ctx context.Context) UpdatePluginsOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdatePluginsOperationOutput)
}

// UpdatePluginsOperationArrayInput is an input type that accepts UpdatePluginsOperationArray and UpdatePluginsOperationArrayOutput values.
// You can construct a concrete instance of `UpdatePluginsOperationArrayInput` via:
//
//	UpdatePluginsOperationArray{ UpdatePluginsOperationArgs{...} }
type UpdatePluginsOperationArrayInput interface {
	pulumi.Input

	ToUpdatePluginsOperationArrayOutput() UpdatePluginsOperationArrayOutput
	ToUpdatePluginsOperationArrayOutputWithContext(context.Context) UpdatePluginsOperationArrayOutput
}

type UpdatePluginsOperationArray []UpdatePluginsOperationInput

func (UpdatePluginsOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpdatePluginsOperation)(nil)).Elem()
}

func (i UpdatePluginsOperationArray) ToUpdatePluginsOperationArrayOutput() UpdatePluginsOperationArrayOutput {
	return i.ToUpdatePluginsOperationArrayOutputWithContext(context.Background())
}

func (i UpdatePluginsOperationArray) ToUpdatePluginsOperationArrayOutputWithContext(ctx context.Context) UpdatePluginsOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdatePluginsOperationArrayOutput)
}

// UpdatePluginsOperationMapInput is an input type that accepts UpdatePluginsOperationMap and UpdatePluginsOperationMapOutput values.
// You can construct a concrete instance of `UpdatePluginsOperationMapInput` via:
//
//	UpdatePluginsOperationMap{ "key": UpdatePluginsOperationArgs{...} }
type UpdatePluginsOperationMapInput interface {
	pulumi.Input

	ToUpdatePluginsOperationMapOutput() UpdatePluginsOperationMapOutput
	ToUpdatePluginsOperationMapOutputWithContext(context.Context) UpdatePluginsOperationMapOutput
}

type UpdatePluginsOperationMap map[string]UpdatePluginsOperationInput

func (UpdatePluginsOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpdatePluginsOperation)(nil)).Elem()
}

func (i UpdatePluginsOperationMap) ToUpdatePluginsOperationMapOutput() UpdatePluginsOperationMapOutput {
	return i.ToUpdatePluginsOperationMapOutputWithContext(context.Background())
}

func (i UpdatePluginsOperationMap) ToUpdatePluginsOperationMapOutputWithContext(ctx context.Context) UpdatePluginsOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdatePluginsOperationMapOutput)
}

type UpdatePluginsOperationOutput struct{ *pulumi.OutputState }

func (UpdatePluginsOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdatePluginsOperation)(nil)).Elem()
}

func (o UpdatePluginsOperationOutput) ToUpdatePluginsOperationOutput() UpdatePluginsOperationOutput {
	return o
}

func (o UpdatePluginsOperationOutput) ToUpdatePluginsOperationOutputWithContext(ctx context.Context) UpdatePluginsOperationOutput {
	return o
}

// Whether to force a restart. Default is false.
func (o UpdatePluginsOperationOutput) ForceRestart() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UpdatePluginsOperation) pulumi.BoolPtrOutput { return v.ForceRestart }).(pulumi.BoolPtrOutput)
}

// Whether to reinstall, default value false.
func (o UpdatePluginsOperationOutput) ForceUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UpdatePluginsOperation) pulumi.BoolPtrOutput { return v.ForceUpdate }).(pulumi.BoolPtrOutput)
}

// List of plugins that need to be installed.
func (o UpdatePluginsOperationOutput) InstallPluginLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UpdatePluginsOperation) pulumi.StringArrayOutput { return v.InstallPluginLists }).(pulumi.StringArrayOutput)
}

// Instance id.
func (o UpdatePluginsOperationOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdatePluginsOperation) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Plugin type. 0: system plugin.
func (o UpdatePluginsOperationOutput) PluginType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UpdatePluginsOperation) pulumi.IntPtrOutput { return v.PluginType }).(pulumi.IntPtrOutput)
}

// List of plugins that need to be uninstalled.
func (o UpdatePluginsOperationOutput) RemovePluginLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UpdatePluginsOperation) pulumi.StringArrayOutput { return v.RemovePluginLists }).(pulumi.StringArrayOutput)
}

type UpdatePluginsOperationArrayOutput struct{ *pulumi.OutputState }

func (UpdatePluginsOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UpdatePluginsOperation)(nil)).Elem()
}

func (o UpdatePluginsOperationArrayOutput) ToUpdatePluginsOperationArrayOutput() UpdatePluginsOperationArrayOutput {
	return o
}

func (o UpdatePluginsOperationArrayOutput) ToUpdatePluginsOperationArrayOutputWithContext(ctx context.Context) UpdatePluginsOperationArrayOutput {
	return o
}

func (o UpdatePluginsOperationArrayOutput) Index(i pulumi.IntInput) UpdatePluginsOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UpdatePluginsOperation {
		return vs[0].([]*UpdatePluginsOperation)[vs[1].(int)]
	}).(UpdatePluginsOperationOutput)
}

type UpdatePluginsOperationMapOutput struct{ *pulumi.OutputState }

func (UpdatePluginsOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UpdatePluginsOperation)(nil)).Elem()
}

func (o UpdatePluginsOperationMapOutput) ToUpdatePluginsOperationMapOutput() UpdatePluginsOperationMapOutput {
	return o
}

func (o UpdatePluginsOperationMapOutput) ToUpdatePluginsOperationMapOutputWithContext(ctx context.Context) UpdatePluginsOperationMapOutput {
	return o
}

func (o UpdatePluginsOperationMapOutput) MapIndex(k pulumi.StringInput) UpdatePluginsOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UpdatePluginsOperation {
		return vs[0].(map[string]*UpdatePluginsOperation)[vs[1].(string)]
	}).(UpdatePluginsOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UpdatePluginsOperationInput)(nil)).Elem(), &UpdatePluginsOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpdatePluginsOperationArrayInput)(nil)).Elem(), UpdatePluginsOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpdatePluginsOperationMapInput)(nil)).Elem(), UpdatePluginsOperationMap{})
	pulumi.RegisterOutputType(UpdatePluginsOperationOutput{})
	pulumi.RegisterOutputType(UpdatePluginsOperationArrayOutput{})
	pulumi.RegisterOutputType(UpdatePluginsOperationMapOutput{})
}
