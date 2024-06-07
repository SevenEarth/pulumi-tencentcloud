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

// Provides a resource to create a cynosdb clusterResourcePackagesAttachment
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
//			_, err := Cynosdb.NewClusterResourcePackagesAttachment(ctx, "clusterResourcePackagesAttachment", &Cynosdb.ClusterResourcePackagesAttachmentArgs{
//				ClusterId: pulumi.String("cynosdbmysql-q1d8151n"),
//				PackageIds: pulumi.StringArray{
//					pulumi.String("package-hy4d2ppl"),
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
// cynosdb cluster_resource_packages_attachment can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Cynosdb/clusterResourcePackagesAttachment:ClusterResourcePackagesAttachment cluster_resource_packages_attachment cluster_resource_packages_attachment_id
// ```
type ClusterResourcePackagesAttachment struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Resource Package Unique ID.
	PackageIds pulumi.StringArrayOutput `pulumi:"packageIds"`
}

// NewClusterResourcePackagesAttachment registers a new resource with the given unique name, arguments, and options.
func NewClusterResourcePackagesAttachment(ctx *pulumi.Context,
	name string, args *ClusterResourcePackagesAttachmentArgs, opts ...pulumi.ResourceOption) (*ClusterResourcePackagesAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.PackageIds == nil {
		return nil, errors.New("invalid value for required argument 'PackageIds'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ClusterResourcePackagesAttachment
	err := ctx.RegisterResource("tencentcloud:Cynosdb/clusterResourcePackagesAttachment:ClusterResourcePackagesAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterResourcePackagesAttachment gets an existing ClusterResourcePackagesAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterResourcePackagesAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterResourcePackagesAttachmentState, opts ...pulumi.ResourceOption) (*ClusterResourcePackagesAttachment, error) {
	var resource ClusterResourcePackagesAttachment
	err := ctx.ReadResource("tencentcloud:Cynosdb/clusterResourcePackagesAttachment:ClusterResourcePackagesAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterResourcePackagesAttachment resources.
type clusterResourcePackagesAttachmentState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Resource Package Unique ID.
	PackageIds []string `pulumi:"packageIds"`
}

type ClusterResourcePackagesAttachmentState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Resource Package Unique ID.
	PackageIds pulumi.StringArrayInput
}

func (ClusterResourcePackagesAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterResourcePackagesAttachmentState)(nil)).Elem()
}

type clusterResourcePackagesAttachmentArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Resource Package Unique ID.
	PackageIds []string `pulumi:"packageIds"`
}

// The set of arguments for constructing a ClusterResourcePackagesAttachment resource.
type ClusterResourcePackagesAttachmentArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Resource Package Unique ID.
	PackageIds pulumi.StringArrayInput
}

func (ClusterResourcePackagesAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterResourcePackagesAttachmentArgs)(nil)).Elem()
}

type ClusterResourcePackagesAttachmentInput interface {
	pulumi.Input

	ToClusterResourcePackagesAttachmentOutput() ClusterResourcePackagesAttachmentOutput
	ToClusterResourcePackagesAttachmentOutputWithContext(ctx context.Context) ClusterResourcePackagesAttachmentOutput
}

func (*ClusterResourcePackagesAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourcePackagesAttachment)(nil)).Elem()
}

func (i *ClusterResourcePackagesAttachment) ToClusterResourcePackagesAttachmentOutput() ClusterResourcePackagesAttachmentOutput {
	return i.ToClusterResourcePackagesAttachmentOutputWithContext(context.Background())
}

func (i *ClusterResourcePackagesAttachment) ToClusterResourcePackagesAttachmentOutputWithContext(ctx context.Context) ClusterResourcePackagesAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePackagesAttachmentOutput)
}

// ClusterResourcePackagesAttachmentArrayInput is an input type that accepts ClusterResourcePackagesAttachmentArray and ClusterResourcePackagesAttachmentArrayOutput values.
// You can construct a concrete instance of `ClusterResourcePackagesAttachmentArrayInput` via:
//
//	ClusterResourcePackagesAttachmentArray{ ClusterResourcePackagesAttachmentArgs{...} }
type ClusterResourcePackagesAttachmentArrayInput interface {
	pulumi.Input

	ToClusterResourcePackagesAttachmentArrayOutput() ClusterResourcePackagesAttachmentArrayOutput
	ToClusterResourcePackagesAttachmentArrayOutputWithContext(context.Context) ClusterResourcePackagesAttachmentArrayOutput
}

type ClusterResourcePackagesAttachmentArray []ClusterResourcePackagesAttachmentInput

func (ClusterResourcePackagesAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterResourcePackagesAttachment)(nil)).Elem()
}

func (i ClusterResourcePackagesAttachmentArray) ToClusterResourcePackagesAttachmentArrayOutput() ClusterResourcePackagesAttachmentArrayOutput {
	return i.ToClusterResourcePackagesAttachmentArrayOutputWithContext(context.Background())
}

func (i ClusterResourcePackagesAttachmentArray) ToClusterResourcePackagesAttachmentArrayOutputWithContext(ctx context.Context) ClusterResourcePackagesAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePackagesAttachmentArrayOutput)
}

// ClusterResourcePackagesAttachmentMapInput is an input type that accepts ClusterResourcePackagesAttachmentMap and ClusterResourcePackagesAttachmentMapOutput values.
// You can construct a concrete instance of `ClusterResourcePackagesAttachmentMapInput` via:
//
//	ClusterResourcePackagesAttachmentMap{ "key": ClusterResourcePackagesAttachmentArgs{...} }
type ClusterResourcePackagesAttachmentMapInput interface {
	pulumi.Input

	ToClusterResourcePackagesAttachmentMapOutput() ClusterResourcePackagesAttachmentMapOutput
	ToClusterResourcePackagesAttachmentMapOutputWithContext(context.Context) ClusterResourcePackagesAttachmentMapOutput
}

type ClusterResourcePackagesAttachmentMap map[string]ClusterResourcePackagesAttachmentInput

func (ClusterResourcePackagesAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterResourcePackagesAttachment)(nil)).Elem()
}

func (i ClusterResourcePackagesAttachmentMap) ToClusterResourcePackagesAttachmentMapOutput() ClusterResourcePackagesAttachmentMapOutput {
	return i.ToClusterResourcePackagesAttachmentMapOutputWithContext(context.Background())
}

func (i ClusterResourcePackagesAttachmentMap) ToClusterResourcePackagesAttachmentMapOutputWithContext(ctx context.Context) ClusterResourcePackagesAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePackagesAttachmentMapOutput)
}

type ClusterResourcePackagesAttachmentOutput struct{ *pulumi.OutputState }

func (ClusterResourcePackagesAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourcePackagesAttachment)(nil)).Elem()
}

func (o ClusterResourcePackagesAttachmentOutput) ToClusterResourcePackagesAttachmentOutput() ClusterResourcePackagesAttachmentOutput {
	return o
}

func (o ClusterResourcePackagesAttachmentOutput) ToClusterResourcePackagesAttachmentOutputWithContext(ctx context.Context) ClusterResourcePackagesAttachmentOutput {
	return o
}

// Cluster ID.
func (o ClusterResourcePackagesAttachmentOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterResourcePackagesAttachment) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Resource Package Unique ID.
func (o ClusterResourcePackagesAttachmentOutput) PackageIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterResourcePackagesAttachment) pulumi.StringArrayOutput { return v.PackageIds }).(pulumi.StringArrayOutput)
}

type ClusterResourcePackagesAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ClusterResourcePackagesAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterResourcePackagesAttachment)(nil)).Elem()
}

func (o ClusterResourcePackagesAttachmentArrayOutput) ToClusterResourcePackagesAttachmentArrayOutput() ClusterResourcePackagesAttachmentArrayOutput {
	return o
}

func (o ClusterResourcePackagesAttachmentArrayOutput) ToClusterResourcePackagesAttachmentArrayOutputWithContext(ctx context.Context) ClusterResourcePackagesAttachmentArrayOutput {
	return o
}

func (o ClusterResourcePackagesAttachmentArrayOutput) Index(i pulumi.IntInput) ClusterResourcePackagesAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterResourcePackagesAttachment {
		return vs[0].([]*ClusterResourcePackagesAttachment)[vs[1].(int)]
	}).(ClusterResourcePackagesAttachmentOutput)
}

type ClusterResourcePackagesAttachmentMapOutput struct{ *pulumi.OutputState }

func (ClusterResourcePackagesAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterResourcePackagesAttachment)(nil)).Elem()
}

func (o ClusterResourcePackagesAttachmentMapOutput) ToClusterResourcePackagesAttachmentMapOutput() ClusterResourcePackagesAttachmentMapOutput {
	return o
}

func (o ClusterResourcePackagesAttachmentMapOutput) ToClusterResourcePackagesAttachmentMapOutputWithContext(ctx context.Context) ClusterResourcePackagesAttachmentMapOutput {
	return o
}

func (o ClusterResourcePackagesAttachmentMapOutput) MapIndex(k pulumi.StringInput) ClusterResourcePackagesAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterResourcePackagesAttachment {
		return vs[0].(map[string]*ClusterResourcePackagesAttachment)[vs[1].(string)]
	}).(ClusterResourcePackagesAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterResourcePackagesAttachmentInput)(nil)).Elem(), &ClusterResourcePackagesAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterResourcePackagesAttachmentArrayInput)(nil)).Elem(), ClusterResourcePackagesAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterResourcePackagesAttachmentMapInput)(nil)).Elem(), ClusterResourcePackagesAttachmentMap{})
	pulumi.RegisterOutputType(ClusterResourcePackagesAttachmentOutput{})
	pulumi.RegisterOutputType(ClusterResourcePackagesAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ClusterResourcePackagesAttachmentMapOutput{})
}
