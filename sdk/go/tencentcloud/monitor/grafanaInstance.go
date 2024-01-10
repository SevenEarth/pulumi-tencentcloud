// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a monitor grafanaInstance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			availabilityZone := "ap-guangzhou-6"
//			if param := cfg.Get("availabilityZone"); param != "" {
//				availabilityZone = param
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				VpcId:            vpc.ID(),
//				AvailabilityZone: pulumi.String(availabilityZone),
//				CidrBlock:        pulumi.String("10.0.1.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Monitor.NewGrafanaInstance(ctx, "foo", &Monitor.GrafanaInstanceArgs{
//				InstanceName: pulumi.String("test-grafana"),
//				VpcId:        vpc.ID(),
//				SubnetIds: pulumi.StringArray{
//					subnet.ID(),
//				},
//				GrafanaInitPassword: pulumi.String("1234567890"),
//				EnableInternet:      pulumi.Bool(false),
//				IsDestroy:           pulumi.Bool(true),
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("test"),
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
// monitor grafanaInstance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Monitor/grafanaInstance:GrafanaInstance foo grafanaInstance_id
//
// ```
type GrafanaInstance struct {
	pulumi.CustomResourceState

	// Control whether grafana could be accessed by internet.
	EnableInternet pulumi.BoolOutput `pulumi:"enableInternet"`
	// Grafana server admin password.
	GrafanaInitPassword pulumi.StringOutput `pulumi:"grafanaInitPassword"`
	// Grafana instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Instance name.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
	InstanceStatus pulumi.IntOutput `pulumi:"instanceStatus"`
	// Grafana public address.
	InternalUrl pulumi.StringOutput `pulumi:"internalUrl"`
	// Grafana intranet address.
	InternetUrl pulumi.StringOutput `pulumi:"internetUrl"`
	// Whether to clean up completely, the default is false.
	IsDestroy pulumi.BoolPtrOutput `pulumi:"isDestroy"`
	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	//
	// Deprecated: It has been deprecated from version 1.81.16.
	IsDistroy pulumi.BoolPtrOutput `pulumi:"isDistroy"`
	// Grafana external url which could be accessed by user.
	RootUrl pulumi.StringOutput `pulumi:"rootUrl"`
	// Subnet Id array.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Vpc Id.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewGrafanaInstance registers a new resource with the given unique name, arguments, and options.
func NewGrafanaInstance(ctx *pulumi.Context,
	name string, args *GrafanaInstanceArgs, opts ...pulumi.ResourceOption) (*GrafanaInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource GrafanaInstance
	err := ctx.RegisterResource("tencentcloud:Monitor/grafanaInstance:GrafanaInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrafanaInstance gets an existing GrafanaInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrafanaInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrafanaInstanceState, opts ...pulumi.ResourceOption) (*GrafanaInstance, error) {
	var resource GrafanaInstance
	err := ctx.ReadResource("tencentcloud:Monitor/grafanaInstance:GrafanaInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GrafanaInstance resources.
type grafanaInstanceState struct {
	// Control whether grafana could be accessed by internet.
	EnableInternet *bool `pulumi:"enableInternet"`
	// Grafana server admin password.
	GrafanaInitPassword *string `pulumi:"grafanaInitPassword"`
	// Grafana instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Instance name.
	InstanceName *string `pulumi:"instanceName"`
	// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
	InstanceStatus *int `pulumi:"instanceStatus"`
	// Grafana public address.
	InternalUrl *string `pulumi:"internalUrl"`
	// Grafana intranet address.
	InternetUrl *string `pulumi:"internetUrl"`
	// Whether to clean up completely, the default is false.
	IsDestroy *bool `pulumi:"isDestroy"`
	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	//
	// Deprecated: It has been deprecated from version 1.81.16.
	IsDistroy *bool `pulumi:"isDistroy"`
	// Grafana external url which could be accessed by user.
	RootUrl *string `pulumi:"rootUrl"`
	// Subnet Id array.
	SubnetIds []string `pulumi:"subnetIds"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Vpc Id.
	VpcId *string `pulumi:"vpcId"`
}

type GrafanaInstanceState struct {
	// Control whether grafana could be accessed by internet.
	EnableInternet pulumi.BoolPtrInput
	// Grafana server admin password.
	GrafanaInitPassword pulumi.StringPtrInput
	// Grafana instance id.
	InstanceId pulumi.StringPtrInput
	// Instance name.
	InstanceName pulumi.StringPtrInput
	// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
	InstanceStatus pulumi.IntPtrInput
	// Grafana public address.
	InternalUrl pulumi.StringPtrInput
	// Grafana intranet address.
	InternetUrl pulumi.StringPtrInput
	// Whether to clean up completely, the default is false.
	IsDestroy pulumi.BoolPtrInput
	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	//
	// Deprecated: It has been deprecated from version 1.81.16.
	IsDistroy pulumi.BoolPtrInput
	// Grafana external url which could be accessed by user.
	RootUrl pulumi.StringPtrInput
	// Subnet Id array.
	SubnetIds pulumi.StringArrayInput
	// Tag description list.
	Tags pulumi.MapInput
	// Vpc Id.
	VpcId pulumi.StringPtrInput
}

func (GrafanaInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaInstanceState)(nil)).Elem()
}

type grafanaInstanceArgs struct {
	// Control whether grafana could be accessed by internet.
	EnableInternet *bool `pulumi:"enableInternet"`
	// Grafana server admin password.
	GrafanaInitPassword *string `pulumi:"grafanaInitPassword"`
	// Instance name.
	InstanceName string `pulumi:"instanceName"`
	// Whether to clean up completely, the default is false.
	IsDestroy *bool `pulumi:"isDestroy"`
	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	//
	// Deprecated: It has been deprecated from version 1.81.16.
	IsDistroy *bool `pulumi:"isDistroy"`
	// Subnet Id array.
	SubnetIds []string `pulumi:"subnetIds"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// Vpc Id.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a GrafanaInstance resource.
type GrafanaInstanceArgs struct {
	// Control whether grafana could be accessed by internet.
	EnableInternet pulumi.BoolPtrInput
	// Grafana server admin password.
	GrafanaInitPassword pulumi.StringPtrInput
	// Instance name.
	InstanceName pulumi.StringInput
	// Whether to clean up completely, the default is false.
	IsDestroy pulumi.BoolPtrInput
	// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
	//
	// Deprecated: It has been deprecated from version 1.81.16.
	IsDistroy pulumi.BoolPtrInput
	// Subnet Id array.
	SubnetIds pulumi.StringArrayInput
	// Tag description list.
	Tags pulumi.MapInput
	// Vpc Id.
	VpcId pulumi.StringPtrInput
}

func (GrafanaInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grafanaInstanceArgs)(nil)).Elem()
}

type GrafanaInstanceInput interface {
	pulumi.Input

	ToGrafanaInstanceOutput() GrafanaInstanceOutput
	ToGrafanaInstanceOutputWithContext(ctx context.Context) GrafanaInstanceOutput
}

func (*GrafanaInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**GrafanaInstance)(nil)).Elem()
}

func (i *GrafanaInstance) ToGrafanaInstanceOutput() GrafanaInstanceOutput {
	return i.ToGrafanaInstanceOutputWithContext(context.Background())
}

func (i *GrafanaInstance) ToGrafanaInstanceOutputWithContext(ctx context.Context) GrafanaInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaInstanceOutput)
}

// GrafanaInstanceArrayInput is an input type that accepts GrafanaInstanceArray and GrafanaInstanceArrayOutput values.
// You can construct a concrete instance of `GrafanaInstanceArrayInput` via:
//
//	GrafanaInstanceArray{ GrafanaInstanceArgs{...} }
type GrafanaInstanceArrayInput interface {
	pulumi.Input

	ToGrafanaInstanceArrayOutput() GrafanaInstanceArrayOutput
	ToGrafanaInstanceArrayOutputWithContext(context.Context) GrafanaInstanceArrayOutput
}

type GrafanaInstanceArray []GrafanaInstanceInput

func (GrafanaInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrafanaInstance)(nil)).Elem()
}

func (i GrafanaInstanceArray) ToGrafanaInstanceArrayOutput() GrafanaInstanceArrayOutput {
	return i.ToGrafanaInstanceArrayOutputWithContext(context.Background())
}

func (i GrafanaInstanceArray) ToGrafanaInstanceArrayOutputWithContext(ctx context.Context) GrafanaInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaInstanceArrayOutput)
}

// GrafanaInstanceMapInput is an input type that accepts GrafanaInstanceMap and GrafanaInstanceMapOutput values.
// You can construct a concrete instance of `GrafanaInstanceMapInput` via:
//
//	GrafanaInstanceMap{ "key": GrafanaInstanceArgs{...} }
type GrafanaInstanceMapInput interface {
	pulumi.Input

	ToGrafanaInstanceMapOutput() GrafanaInstanceMapOutput
	ToGrafanaInstanceMapOutputWithContext(context.Context) GrafanaInstanceMapOutput
}

type GrafanaInstanceMap map[string]GrafanaInstanceInput

func (GrafanaInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrafanaInstance)(nil)).Elem()
}

func (i GrafanaInstanceMap) ToGrafanaInstanceMapOutput() GrafanaInstanceMapOutput {
	return i.ToGrafanaInstanceMapOutputWithContext(context.Background())
}

func (i GrafanaInstanceMap) ToGrafanaInstanceMapOutputWithContext(ctx context.Context) GrafanaInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrafanaInstanceMapOutput)
}

type GrafanaInstanceOutput struct{ *pulumi.OutputState }

func (GrafanaInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrafanaInstance)(nil)).Elem()
}

func (o GrafanaInstanceOutput) ToGrafanaInstanceOutput() GrafanaInstanceOutput {
	return o
}

func (o GrafanaInstanceOutput) ToGrafanaInstanceOutputWithContext(ctx context.Context) GrafanaInstanceOutput {
	return o
}

// Control whether grafana could be accessed by internet.
func (o GrafanaInstanceOutput) EnableInternet() pulumi.BoolOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.BoolOutput { return v.EnableInternet }).(pulumi.BoolOutput)
}

// Grafana server admin password.
func (o GrafanaInstanceOutput) GrafanaInitPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringOutput { return v.GrafanaInitPassword }).(pulumi.StringOutput)
}

// Grafana instance id.
func (o GrafanaInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Instance name.
func (o GrafanaInstanceOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Grafana instance status, 1: Creating, 2: Running, 6: Stopped.
func (o GrafanaInstanceOutput) InstanceStatus() pulumi.IntOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.IntOutput { return v.InstanceStatus }).(pulumi.IntOutput)
}

// Grafana public address.
func (o GrafanaInstanceOutput) InternalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringOutput { return v.InternalUrl }).(pulumi.StringOutput)
}

// Grafana intranet address.
func (o GrafanaInstanceOutput) InternetUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringOutput { return v.InternetUrl }).(pulumi.StringOutput)
}

// Whether to clean up completely, the default is false.
func (o GrafanaInstanceOutput) IsDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.BoolPtrOutput { return v.IsDestroy }).(pulumi.BoolPtrOutput)
}

// It has been deprecated from version 1.81.16. Whether to clean up completely, the default is false.
//
// Deprecated: It has been deprecated from version 1.81.16.
func (o GrafanaInstanceOutput) IsDistroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.BoolPtrOutput { return v.IsDistroy }).(pulumi.BoolPtrOutput)
}

// Grafana external url which could be accessed by user.
func (o GrafanaInstanceOutput) RootUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringOutput { return v.RootUrl }).(pulumi.StringOutput)
}

// Subnet Id array.
func (o GrafanaInstanceOutput) SubnetIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringArrayOutput { return v.SubnetIds }).(pulumi.StringArrayOutput)
}

// Tag description list.
func (o GrafanaInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Vpc Id.
func (o GrafanaInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrafanaInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type GrafanaInstanceArrayOutput struct{ *pulumi.OutputState }

func (GrafanaInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrafanaInstance)(nil)).Elem()
}

func (o GrafanaInstanceArrayOutput) ToGrafanaInstanceArrayOutput() GrafanaInstanceArrayOutput {
	return o
}

func (o GrafanaInstanceArrayOutput) ToGrafanaInstanceArrayOutputWithContext(ctx context.Context) GrafanaInstanceArrayOutput {
	return o
}

func (o GrafanaInstanceArrayOutput) Index(i pulumi.IntInput) GrafanaInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GrafanaInstance {
		return vs[0].([]*GrafanaInstance)[vs[1].(int)]
	}).(GrafanaInstanceOutput)
}

type GrafanaInstanceMapOutput struct{ *pulumi.OutputState }

func (GrafanaInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrafanaInstance)(nil)).Elem()
}

func (o GrafanaInstanceMapOutput) ToGrafanaInstanceMapOutput() GrafanaInstanceMapOutput {
	return o
}

func (o GrafanaInstanceMapOutput) ToGrafanaInstanceMapOutputWithContext(ctx context.Context) GrafanaInstanceMapOutput {
	return o
}

func (o GrafanaInstanceMapOutput) MapIndex(k pulumi.StringInput) GrafanaInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GrafanaInstance {
		return vs[0].(map[string]*GrafanaInstance)[vs[1].(string)]
	}).(GrafanaInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrafanaInstanceInput)(nil)).Elem(), &GrafanaInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrafanaInstanceArrayInput)(nil)).Elem(), GrafanaInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrafanaInstanceMapInput)(nil)).Elem(), GrafanaInstanceMap{})
	pulumi.RegisterOutputType(GrafanaInstanceOutput{})
	pulumi.RegisterOutputType(GrafanaInstanceArrayOutput{})
	pulumi.RegisterOutputType(GrafanaInstanceMapOutput{})
}
