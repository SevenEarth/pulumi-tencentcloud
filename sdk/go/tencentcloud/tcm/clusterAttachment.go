// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tcm clusterAttachment
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tcm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcm"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tcm.NewClusterAttachment(ctx, "clusterAttachment", &Tcm.ClusterAttachmentArgs{
// 			ClusterLists: tcm.ClusterAttachmentClusterListArray{
// 				&tcm.ClusterAttachmentClusterListArgs{
// 					ClusterId: pulumi.String("cls-rc5uy6dy"),
// 					Region:    pulumi.String("ap-guangzhou"),
// 					Role:      pulumi.String("REMOTE"),
// 					SubnetId:  pulumi.String("subnet-lkyb3ayc"),
// 					Type:      pulumi.String("TKE"),
// 					VpcId:     pulumi.String("vpc-a1jycmbx"),
// 				},
// 			},
// 			MeshId: pulumi.String("mesh-b9q6vf9l"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// tcm cluster_attachment can be imported using the mesh_id#cluster_id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tcm/clusterAttachment:ClusterAttachment cluster_attachment mesh-b9q6vf9l#cls-rc5uy6dy
// ```
type ClusterAttachment struct {
	pulumi.CustomResourceState

	// Cluster list.
	ClusterLists ClusterAttachmentClusterListArrayOutput `pulumi:"clusterLists"`
	// Mesh ID.
	MeshId pulumi.StringOutput `pulumi:"meshId"`
}

// NewClusterAttachment registers a new resource with the given unique name, arguments, and options.
func NewClusterAttachment(ctx *pulumi.Context,
	name string, args *ClusterAttachmentArgs, opts ...pulumi.ResourceOption) (*ClusterAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshId == nil {
		return nil, errors.New("invalid value for required argument 'MeshId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterAttachment
	err := ctx.RegisterResource("tencentcloud:Tcm/clusterAttachment:ClusterAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterAttachment gets an existing ClusterAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterAttachmentState, opts ...pulumi.ResourceOption) (*ClusterAttachment, error) {
	var resource ClusterAttachment
	err := ctx.ReadResource("tencentcloud:Tcm/clusterAttachment:ClusterAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterAttachment resources.
type clusterAttachmentState struct {
	// Cluster list.
	ClusterLists []ClusterAttachmentClusterList `pulumi:"clusterLists"`
	// Mesh ID.
	MeshId *string `pulumi:"meshId"`
}

type ClusterAttachmentState struct {
	// Cluster list.
	ClusterLists ClusterAttachmentClusterListArrayInput
	// Mesh ID.
	MeshId pulumi.StringPtrInput
}

func (ClusterAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAttachmentState)(nil)).Elem()
}

type clusterAttachmentArgs struct {
	// Cluster list.
	ClusterLists []ClusterAttachmentClusterList `pulumi:"clusterLists"`
	// Mesh ID.
	MeshId string `pulumi:"meshId"`
}

// The set of arguments for constructing a ClusterAttachment resource.
type ClusterAttachmentArgs struct {
	// Cluster list.
	ClusterLists ClusterAttachmentClusterListArrayInput
	// Mesh ID.
	MeshId pulumi.StringInput
}

func (ClusterAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAttachmentArgs)(nil)).Elem()
}

type ClusterAttachmentInput interface {
	pulumi.Input

	ToClusterAttachmentOutput() ClusterAttachmentOutput
	ToClusterAttachmentOutputWithContext(ctx context.Context) ClusterAttachmentOutput
}

func (*ClusterAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAttachment)(nil)).Elem()
}

func (i *ClusterAttachment) ToClusterAttachmentOutput() ClusterAttachmentOutput {
	return i.ToClusterAttachmentOutputWithContext(context.Background())
}

func (i *ClusterAttachment) ToClusterAttachmentOutputWithContext(ctx context.Context) ClusterAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAttachmentOutput)
}

// ClusterAttachmentArrayInput is an input type that accepts ClusterAttachmentArray and ClusterAttachmentArrayOutput values.
// You can construct a concrete instance of `ClusterAttachmentArrayInput` via:
//
//          ClusterAttachmentArray{ ClusterAttachmentArgs{...} }
type ClusterAttachmentArrayInput interface {
	pulumi.Input

	ToClusterAttachmentArrayOutput() ClusterAttachmentArrayOutput
	ToClusterAttachmentArrayOutputWithContext(context.Context) ClusterAttachmentArrayOutput
}

type ClusterAttachmentArray []ClusterAttachmentInput

func (ClusterAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAttachment)(nil)).Elem()
}

func (i ClusterAttachmentArray) ToClusterAttachmentArrayOutput() ClusterAttachmentArrayOutput {
	return i.ToClusterAttachmentArrayOutputWithContext(context.Background())
}

func (i ClusterAttachmentArray) ToClusterAttachmentArrayOutputWithContext(ctx context.Context) ClusterAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAttachmentArrayOutput)
}

// ClusterAttachmentMapInput is an input type that accepts ClusterAttachmentMap and ClusterAttachmentMapOutput values.
// You can construct a concrete instance of `ClusterAttachmentMapInput` via:
//
//          ClusterAttachmentMap{ "key": ClusterAttachmentArgs{...} }
type ClusterAttachmentMapInput interface {
	pulumi.Input

	ToClusterAttachmentMapOutput() ClusterAttachmentMapOutput
	ToClusterAttachmentMapOutputWithContext(context.Context) ClusterAttachmentMapOutput
}

type ClusterAttachmentMap map[string]ClusterAttachmentInput

func (ClusterAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAttachment)(nil)).Elem()
}

func (i ClusterAttachmentMap) ToClusterAttachmentMapOutput() ClusterAttachmentMapOutput {
	return i.ToClusterAttachmentMapOutputWithContext(context.Background())
}

func (i ClusterAttachmentMap) ToClusterAttachmentMapOutputWithContext(ctx context.Context) ClusterAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAttachmentMapOutput)
}

type ClusterAttachmentOutput struct{ *pulumi.OutputState }

func (ClusterAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAttachment)(nil)).Elem()
}

func (o ClusterAttachmentOutput) ToClusterAttachmentOutput() ClusterAttachmentOutput {
	return o
}

func (o ClusterAttachmentOutput) ToClusterAttachmentOutputWithContext(ctx context.Context) ClusterAttachmentOutput {
	return o
}

// Cluster list.
func (o ClusterAttachmentOutput) ClusterLists() ClusterAttachmentClusterListArrayOutput {
	return o.ApplyT(func(v *ClusterAttachment) ClusterAttachmentClusterListArrayOutput { return v.ClusterLists }).(ClusterAttachmentClusterListArrayOutput)
}

// Mesh ID.
func (o ClusterAttachmentOutput) MeshId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterAttachment) pulumi.StringOutput { return v.MeshId }).(pulumi.StringOutput)
}

type ClusterAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ClusterAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAttachment)(nil)).Elem()
}

func (o ClusterAttachmentArrayOutput) ToClusterAttachmentArrayOutput() ClusterAttachmentArrayOutput {
	return o
}

func (o ClusterAttachmentArrayOutput) ToClusterAttachmentArrayOutputWithContext(ctx context.Context) ClusterAttachmentArrayOutput {
	return o
}

func (o ClusterAttachmentArrayOutput) Index(i pulumi.IntInput) ClusterAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterAttachment {
		return vs[0].([]*ClusterAttachment)[vs[1].(int)]
	}).(ClusterAttachmentOutput)
}

type ClusterAttachmentMapOutput struct{ *pulumi.OutputState }

func (ClusterAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAttachment)(nil)).Elem()
}

func (o ClusterAttachmentMapOutput) ToClusterAttachmentMapOutput() ClusterAttachmentMapOutput {
	return o
}

func (o ClusterAttachmentMapOutput) ToClusterAttachmentMapOutputWithContext(ctx context.Context) ClusterAttachmentMapOutput {
	return o
}

func (o ClusterAttachmentMapOutput) MapIndex(k pulumi.StringInput) ClusterAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterAttachment {
		return vs[0].(map[string]*ClusterAttachment)[vs[1].(string)]
	}).(ClusterAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAttachmentInput)(nil)).Elem(), &ClusterAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAttachmentArrayInput)(nil)).Elem(), ClusterAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAttachmentMapInput)(nil)).Elem(), ClusterAttachmentMap{})
	pulumi.RegisterOutputType(ClusterAttachmentOutput{})
	pulumi.RegisterOutputType(ClusterAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ClusterAttachmentMapOutput{})
}
