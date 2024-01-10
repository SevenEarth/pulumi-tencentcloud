// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tcm prometheusAttachment
//
// > **NOTE:** Instructions for use: 1. Use Tencent Cloud Prometheus to monitor TMP, please enter `vpcId`, `subnetId`, `region` or `instanceId`, it is recommended to use an existing tmp instance; 2. To use the third-party Prometheus service, please enter `customProm`; 3. `Tcm.PrometheusAttachment` does not support modification; 4. If you use Tencent Cloud Prometheus to monitor TMP, enter `vpcId`, `subnetId`, `region` to create a new Prometheus monitoring instance, destroy will not destroy the Prometheus monitoring instance
// **NOTE:** If you use the config attribute prometheus in tencentcloud_tcm_mesh, do not use Tcm.PrometheusAttachment
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tcm"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcm"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Tcm.NewPrometheusAttachment(ctx, "prometheusAttachment", &Tcm.PrometheusAttachmentArgs{
//				MeshId: pulumi.String("mesh-rofjmxxx"),
//				Prometheus: &tcm.PrometheusAttachmentPrometheusArgs{
//					InstanceId: pulumi.String(""),
//					Region:     pulumi.String("ap-guangzhou"),
//					SubnetId:   pulumi.String("subnet-driddxxx"),
//					VpcId:      pulumi.String("vpc-pewdpxxx"),
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
// tcm prometheus_attachment can be imported using the mesh_id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Tcm/prometheusAttachment:PrometheusAttachment prometheus_attachment mesh-rofjmxxx
//
// ```
type PrometheusAttachment struct {
	pulumi.CustomResourceState

	// Mesh ID.
	MeshId pulumi.StringOutput `pulumi:"meshId"`
	// Prometheus configuration.
	Prometheus PrometheusAttachmentPrometheusOutput `pulumi:"prometheus"`
}

// NewPrometheusAttachment registers a new resource with the given unique name, arguments, and options.
func NewPrometheusAttachment(ctx *pulumi.Context,
	name string, args *PrometheusAttachmentArgs, opts ...pulumi.ResourceOption) (*PrometheusAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshId == nil {
		return nil, errors.New("invalid value for required argument 'MeshId'")
	}
	if args.Prometheus == nil {
		return nil, errors.New("invalid value for required argument 'Prometheus'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PrometheusAttachment
	err := ctx.RegisterResource("tencentcloud:Tcm/prometheusAttachment:PrometheusAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheusAttachment gets an existing PrometheusAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheusAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusAttachmentState, opts ...pulumi.ResourceOption) (*PrometheusAttachment, error) {
	var resource PrometheusAttachment
	err := ctx.ReadResource("tencentcloud:Tcm/prometheusAttachment:PrometheusAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrometheusAttachment resources.
type prometheusAttachmentState struct {
	// Mesh ID.
	MeshId *string `pulumi:"meshId"`
	// Prometheus configuration.
	Prometheus *PrometheusAttachmentPrometheus `pulumi:"prometheus"`
}

type PrometheusAttachmentState struct {
	// Mesh ID.
	MeshId pulumi.StringPtrInput
	// Prometheus configuration.
	Prometheus PrometheusAttachmentPrometheusPtrInput
}

func (PrometheusAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAttachmentState)(nil)).Elem()
}

type prometheusAttachmentArgs struct {
	// Mesh ID.
	MeshId string `pulumi:"meshId"`
	// Prometheus configuration.
	Prometheus PrometheusAttachmentPrometheus `pulumi:"prometheus"`
}

// The set of arguments for constructing a PrometheusAttachment resource.
type PrometheusAttachmentArgs struct {
	// Mesh ID.
	MeshId pulumi.StringInput
	// Prometheus configuration.
	Prometheus PrometheusAttachmentPrometheusInput
}

func (PrometheusAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusAttachmentArgs)(nil)).Elem()
}

type PrometheusAttachmentInput interface {
	pulumi.Input

	ToPrometheusAttachmentOutput() PrometheusAttachmentOutput
	ToPrometheusAttachmentOutputWithContext(ctx context.Context) PrometheusAttachmentOutput
}

func (*PrometheusAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAttachment)(nil)).Elem()
}

func (i *PrometheusAttachment) ToPrometheusAttachmentOutput() PrometheusAttachmentOutput {
	return i.ToPrometheusAttachmentOutputWithContext(context.Background())
}

func (i *PrometheusAttachment) ToPrometheusAttachmentOutputWithContext(ctx context.Context) PrometheusAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAttachmentOutput)
}

// PrometheusAttachmentArrayInput is an input type that accepts PrometheusAttachmentArray and PrometheusAttachmentArrayOutput values.
// You can construct a concrete instance of `PrometheusAttachmentArrayInput` via:
//
//	PrometheusAttachmentArray{ PrometheusAttachmentArgs{...} }
type PrometheusAttachmentArrayInput interface {
	pulumi.Input

	ToPrometheusAttachmentArrayOutput() PrometheusAttachmentArrayOutput
	ToPrometheusAttachmentArrayOutputWithContext(context.Context) PrometheusAttachmentArrayOutput
}

type PrometheusAttachmentArray []PrometheusAttachmentInput

func (PrometheusAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAttachment)(nil)).Elem()
}

func (i PrometheusAttachmentArray) ToPrometheusAttachmentArrayOutput() PrometheusAttachmentArrayOutput {
	return i.ToPrometheusAttachmentArrayOutputWithContext(context.Background())
}

func (i PrometheusAttachmentArray) ToPrometheusAttachmentArrayOutputWithContext(ctx context.Context) PrometheusAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAttachmentArrayOutput)
}

// PrometheusAttachmentMapInput is an input type that accepts PrometheusAttachmentMap and PrometheusAttachmentMapOutput values.
// You can construct a concrete instance of `PrometheusAttachmentMapInput` via:
//
//	PrometheusAttachmentMap{ "key": PrometheusAttachmentArgs{...} }
type PrometheusAttachmentMapInput interface {
	pulumi.Input

	ToPrometheusAttachmentMapOutput() PrometheusAttachmentMapOutput
	ToPrometheusAttachmentMapOutputWithContext(context.Context) PrometheusAttachmentMapOutput
}

type PrometheusAttachmentMap map[string]PrometheusAttachmentInput

func (PrometheusAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAttachment)(nil)).Elem()
}

func (i PrometheusAttachmentMap) ToPrometheusAttachmentMapOutput() PrometheusAttachmentMapOutput {
	return i.ToPrometheusAttachmentMapOutputWithContext(context.Background())
}

func (i PrometheusAttachmentMap) ToPrometheusAttachmentMapOutputWithContext(ctx context.Context) PrometheusAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusAttachmentMapOutput)
}

type PrometheusAttachmentOutput struct{ *pulumi.OutputState }

func (PrometheusAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusAttachment)(nil)).Elem()
}

func (o PrometheusAttachmentOutput) ToPrometheusAttachmentOutput() PrometheusAttachmentOutput {
	return o
}

func (o PrometheusAttachmentOutput) ToPrometheusAttachmentOutputWithContext(ctx context.Context) PrometheusAttachmentOutput {
	return o
}

// Mesh ID.
func (o PrometheusAttachmentOutput) MeshId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusAttachment) pulumi.StringOutput { return v.MeshId }).(pulumi.StringOutput)
}

// Prometheus configuration.
func (o PrometheusAttachmentOutput) Prometheus() PrometheusAttachmentPrometheusOutput {
	return o.ApplyT(func(v *PrometheusAttachment) PrometheusAttachmentPrometheusOutput { return v.Prometheus }).(PrometheusAttachmentPrometheusOutput)
}

type PrometheusAttachmentArrayOutput struct{ *pulumi.OutputState }

func (PrometheusAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrometheusAttachment)(nil)).Elem()
}

func (o PrometheusAttachmentArrayOutput) ToPrometheusAttachmentArrayOutput() PrometheusAttachmentArrayOutput {
	return o
}

func (o PrometheusAttachmentArrayOutput) ToPrometheusAttachmentArrayOutputWithContext(ctx context.Context) PrometheusAttachmentArrayOutput {
	return o
}

func (o PrometheusAttachmentArrayOutput) Index(i pulumi.IntInput) PrometheusAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrometheusAttachment {
		return vs[0].([]*PrometheusAttachment)[vs[1].(int)]
	}).(PrometheusAttachmentOutput)
}

type PrometheusAttachmentMapOutput struct{ *pulumi.OutputState }

func (PrometheusAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrometheusAttachment)(nil)).Elem()
}

func (o PrometheusAttachmentMapOutput) ToPrometheusAttachmentMapOutput() PrometheusAttachmentMapOutput {
	return o
}

func (o PrometheusAttachmentMapOutput) ToPrometheusAttachmentMapOutputWithContext(ctx context.Context) PrometheusAttachmentMapOutput {
	return o
}

func (o PrometheusAttachmentMapOutput) MapIndex(k pulumi.StringInput) PrometheusAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrometheusAttachment {
		return vs[0].(map[string]*PrometheusAttachment)[vs[1].(string)]
	}).(PrometheusAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAttachmentInput)(nil)).Elem(), &PrometheusAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAttachmentArrayInput)(nil)).Elem(), PrometheusAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusAttachmentMapInput)(nil)).Elem(), PrometheusAttachmentMap{})
	pulumi.RegisterOutputType(PrometheusAttachmentOutput{})
	pulumi.RegisterOutputType(PrometheusAttachmentArrayOutput{})
	pulumi.RegisterOutputType(PrometheusAttachmentMapOutput{})
}
