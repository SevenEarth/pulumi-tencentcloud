// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a apigateway upstream
//
// ## Example Usage
//
// ### Create a basic VPC channel
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Availability"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Images"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Instance"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			zones, err := Availability.GetZonesByProduct(ctx, &availability.GetZonesByProductArgs{
//				Product: "cvm",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			images, err := Images.GetInstance(ctx, &images.GetInstanceArgs{
//				ImageTypes: []string{
//					"PUBLIC_IMAGE",
//				},
//				ImageNameRegex: pulumi.StringRef("Final"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			instanceTypes, err := Instance.GetTypes(ctx, &instance.GetTypesArgs{
//				Filters: []instance.GetTypesFilter{
//					{
//						Name: "instance-family",
//						Values: []string{
//							"S5",
//						},
//					},
//				},
//				CpuCoreCount:   pulumi.IntRef(2),
//				ExcludeSoldOut: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			subnet, err := Subnet.NewInstance(ctx, "subnet", &Subnet.InstanceArgs{
//				AvailabilityZone: pulumi.String(zones.Zones[3].Name),
//				VpcId:            vpc.ID(),
//				CidrBlock:        pulumi.String("10.0.0.0/16"),
//				IsMulticast:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleInstance, err := Instance.NewInstance(ctx, "exampleInstance", &Instance.InstanceArgs{
//				InstanceName:     pulumi.String("tf_example"),
//				AvailabilityZone: pulumi.String(zones.Zones[3].Name),
//				ImageId:          pulumi.String(images.Images[0].ImageId),
//				InstanceType:     pulumi.String(instanceTypes.InstanceTypes[0].InstanceType),
//				SystemDiskType:   pulumi.String("CLOUD_PREMIUM"),
//				SystemDiskSize:   pulumi.Int(50),
//				Hostname:         pulumi.String("terraform"),
//				ProjectId:        pulumi.Int(0),
//				VpcId:            vpc.ID(),
//				SubnetId:         subnet.ID(),
//				DataDisks: instance.InstanceDataDiskArray{
//					&instance.InstanceDataDiskArgs{
//						DataDiskType: pulumi.String("CLOUD_PREMIUM"),
//						DataDiskSize: pulumi.Int(50),
//						Encrypt:      pulumi.Bool(false),
//					},
//				},
//				Tags: pulumi.Map{
//					"tagKey": pulumi.Any("tagValue"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ApiGateway.NewUpstream(ctx, "exampleUpstream", &ApiGateway.UpstreamArgs{
//				Scheme:              pulumi.String("HTTP"),
//				Algorithm:           pulumi.String("ROUND-ROBIN"),
//				UniqVpcId:           vpc.ID(),
//				UpstreamName:        pulumi.String("tf_example"),
//				UpstreamDescription: pulumi.String("desc."),
//				UpstreamType:        pulumi.String("IP_PORT"),
//				Retries:             pulumi.Int(5),
//				Nodes: apigateway.UpstreamNodeArray{
//					&apigateway.UpstreamNodeArgs{
//						Host:         pulumi.String("1.1.1.1"),
//						Port:         pulumi.Int(9090),
//						Weight:       pulumi.Int(10),
//						VmInstanceId: exampleInstance.ID(),
//						Tags: pulumi.StringArray{
//							pulumi.String("tags"),
//						},
//					},
//				},
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
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
// ### Create a complete VPC channel
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ApiGateway.NewUpstream(ctx, "example", &ApiGateway.UpstreamArgs{
//				Scheme:              pulumi.String("HTTP"),
//				Algorithm:           pulumi.String("ROUND-ROBIN"),
//				UniqVpcId:           pulumi.Any(tencentcloud_vpc.Vpc.Id),
//				UpstreamName:        pulumi.String("tf_example"),
//				UpstreamDescription: pulumi.String("desc."),
//				UpstreamType:        pulumi.String("IP_PORT"),
//				Retries:             pulumi.Int(5),
//				Nodes: apigateway.UpstreamNodeArray{
//					&apigateway.UpstreamNodeArgs{
//						Host:         pulumi.String("1.1.1.1"),
//						Port:         pulumi.Int(9090),
//						Weight:       pulumi.Int(10),
//						VmInstanceId: pulumi.Any(tencentcloud_instance.Example.Id),
//						Tags: pulumi.StringArray{
//							pulumi.String("tags"),
//						},
//					},
//				},
//				HealthChecker: &apigateway.UpstreamHealthCheckerArgs{
//					EnableActiveCheck:    pulumi.Bool(true),
//					EnablePassiveCheck:   pulumi.Bool(true),
//					HealthyHttpStatus:    pulumi.String("200"),
//					UnhealthyHttpStatus:  pulumi.String("500"),
//					TcpFailureThreshold:  pulumi.Int(5),
//					TimeoutThreshold:     pulumi.Int(5),
//					HttpFailureThreshold: pulumi.Int(3),
//					ActiveCheckHttpPath:  pulumi.String("/"),
//					ActiveCheckTimeout:   pulumi.Int(5),
//					ActiveCheckInterval:  pulumi.Int(5),
//					UnhealthyTimeout:     pulumi.Int(30),
//				},
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
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
// apigateway upstream can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:ApiGateway/upstream:Upstream upstream upstream_id
// ```
type Upstream struct {
	pulumi.CustomResourceState

	// Load balancing algorithm, value range: ROUND-ROBIN.
	Algorithm pulumi.StringOutput `pulumi:"algorithm"`
	// Health check configuration, currently only supports VPC channels.
	HealthChecker UpstreamHealthCheckerPtrOutput `pulumi:"healthChecker"`
	// Configuration of K8S container service.
	K8sServices UpstreamK8sServiceArrayOutput `pulumi:"k8sServices"`
	// Backend nodes.
	Nodes UpstreamNodeArrayOutput `pulumi:"nodes"`
	// Request retry count, default to 3 times.
	Retries pulumi.IntPtrOutput `pulumi:"retries"`
	// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
	Scheme pulumi.StringOutput `pulumi:"scheme"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// VPC Unique ID.
	UniqVpcId pulumi.StringOutput `pulumi:"uniqVpcId"`
	// Backend channel description.
	UpstreamDescription pulumi.StringPtrOutput `pulumi:"upstreamDescription"`
	// Host request header forwarded by gateway to backend.
	UpstreamHost pulumi.StringPtrOutput `pulumi:"upstreamHost"`
	// Backend channel name.
	UpstreamName pulumi.StringPtrOutput `pulumi:"upstreamName"`
	// Backend access type, value range: IP_PORT, K8S.
	UpstreamType pulumi.StringPtrOutput `pulumi:"upstreamType"`
}

// NewUpstream registers a new resource with the given unique name, arguments, and options.
func NewUpstream(ctx *pulumi.Context,
	name string, args *UpstreamArgs, opts ...pulumi.ResourceOption) (*Upstream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Algorithm == nil {
		return nil, errors.New("invalid value for required argument 'Algorithm'")
	}
	if args.Scheme == nil {
		return nil, errors.New("invalid value for required argument 'Scheme'")
	}
	if args.UniqVpcId == nil {
		return nil, errors.New("invalid value for required argument 'UniqVpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Upstream
	err := ctx.RegisterResource("tencentcloud:ApiGateway/upstream:Upstream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUpstream gets an existing Upstream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUpstream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpstreamState, opts ...pulumi.ResourceOption) (*Upstream, error) {
	var resource Upstream
	err := ctx.ReadResource("tencentcloud:ApiGateway/upstream:Upstream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Upstream resources.
type upstreamState struct {
	// Load balancing algorithm, value range: ROUND-ROBIN.
	Algorithm *string `pulumi:"algorithm"`
	// Health check configuration, currently only supports VPC channels.
	HealthChecker *UpstreamHealthChecker `pulumi:"healthChecker"`
	// Configuration of K8S container service.
	K8sServices []UpstreamK8sService `pulumi:"k8sServices"`
	// Backend nodes.
	Nodes []UpstreamNode `pulumi:"nodes"`
	// Request retry count, default to 3 times.
	Retries *int `pulumi:"retries"`
	// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
	Scheme *string `pulumi:"scheme"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// VPC Unique ID.
	UniqVpcId *string `pulumi:"uniqVpcId"`
	// Backend channel description.
	UpstreamDescription *string `pulumi:"upstreamDescription"`
	// Host request header forwarded by gateway to backend.
	UpstreamHost *string `pulumi:"upstreamHost"`
	// Backend channel name.
	UpstreamName *string `pulumi:"upstreamName"`
	// Backend access type, value range: IP_PORT, K8S.
	UpstreamType *string `pulumi:"upstreamType"`
}

type UpstreamState struct {
	// Load balancing algorithm, value range: ROUND-ROBIN.
	Algorithm pulumi.StringPtrInput
	// Health check configuration, currently only supports VPC channels.
	HealthChecker UpstreamHealthCheckerPtrInput
	// Configuration of K8S container service.
	K8sServices UpstreamK8sServiceArrayInput
	// Backend nodes.
	Nodes UpstreamNodeArrayInput
	// Request retry count, default to 3 times.
	Retries pulumi.IntPtrInput
	// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
	Scheme pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// VPC Unique ID.
	UniqVpcId pulumi.StringPtrInput
	// Backend channel description.
	UpstreamDescription pulumi.StringPtrInput
	// Host request header forwarded by gateway to backend.
	UpstreamHost pulumi.StringPtrInput
	// Backend channel name.
	UpstreamName pulumi.StringPtrInput
	// Backend access type, value range: IP_PORT, K8S.
	UpstreamType pulumi.StringPtrInput
}

func (UpstreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*upstreamState)(nil)).Elem()
}

type upstreamArgs struct {
	// Load balancing algorithm, value range: ROUND-ROBIN.
	Algorithm string `pulumi:"algorithm"`
	// Health check configuration, currently only supports VPC channels.
	HealthChecker *UpstreamHealthChecker `pulumi:"healthChecker"`
	// Configuration of K8S container service.
	K8sServices []UpstreamK8sService `pulumi:"k8sServices"`
	// Backend nodes.
	Nodes []UpstreamNode `pulumi:"nodes"`
	// Request retry count, default to 3 times.
	Retries *int `pulumi:"retries"`
	// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
	Scheme string `pulumi:"scheme"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// VPC Unique ID.
	UniqVpcId string `pulumi:"uniqVpcId"`
	// Backend channel description.
	UpstreamDescription *string `pulumi:"upstreamDescription"`
	// Host request header forwarded by gateway to backend.
	UpstreamHost *string `pulumi:"upstreamHost"`
	// Backend channel name.
	UpstreamName *string `pulumi:"upstreamName"`
	// Backend access type, value range: IP_PORT, K8S.
	UpstreamType *string `pulumi:"upstreamType"`
}

// The set of arguments for constructing a Upstream resource.
type UpstreamArgs struct {
	// Load balancing algorithm, value range: ROUND-ROBIN.
	Algorithm pulumi.StringInput
	// Health check configuration, currently only supports VPC channels.
	HealthChecker UpstreamHealthCheckerPtrInput
	// Configuration of K8S container service.
	K8sServices UpstreamK8sServiceArrayInput
	// Backend nodes.
	Nodes UpstreamNodeArrayInput
	// Request retry count, default to 3 times.
	Retries pulumi.IntPtrInput
	// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
	Scheme pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
	// VPC Unique ID.
	UniqVpcId pulumi.StringInput
	// Backend channel description.
	UpstreamDescription pulumi.StringPtrInput
	// Host request header forwarded by gateway to backend.
	UpstreamHost pulumi.StringPtrInput
	// Backend channel name.
	UpstreamName pulumi.StringPtrInput
	// Backend access type, value range: IP_PORT, K8S.
	UpstreamType pulumi.StringPtrInput
}

func (UpstreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*upstreamArgs)(nil)).Elem()
}

type UpstreamInput interface {
	pulumi.Input

	ToUpstreamOutput() UpstreamOutput
	ToUpstreamOutputWithContext(ctx context.Context) UpstreamOutput
}

func (*Upstream) ElementType() reflect.Type {
	return reflect.TypeOf((**Upstream)(nil)).Elem()
}

func (i *Upstream) ToUpstreamOutput() UpstreamOutput {
	return i.ToUpstreamOutputWithContext(context.Background())
}

func (i *Upstream) ToUpstreamOutputWithContext(ctx context.Context) UpstreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamOutput)
}

// UpstreamArrayInput is an input type that accepts UpstreamArray and UpstreamArrayOutput values.
// You can construct a concrete instance of `UpstreamArrayInput` via:
//
//	UpstreamArray{ UpstreamArgs{...} }
type UpstreamArrayInput interface {
	pulumi.Input

	ToUpstreamArrayOutput() UpstreamArrayOutput
	ToUpstreamArrayOutputWithContext(context.Context) UpstreamArrayOutput
}

type UpstreamArray []UpstreamInput

func (UpstreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Upstream)(nil)).Elem()
}

func (i UpstreamArray) ToUpstreamArrayOutput() UpstreamArrayOutput {
	return i.ToUpstreamArrayOutputWithContext(context.Background())
}

func (i UpstreamArray) ToUpstreamArrayOutputWithContext(ctx context.Context) UpstreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamArrayOutput)
}

// UpstreamMapInput is an input type that accepts UpstreamMap and UpstreamMapOutput values.
// You can construct a concrete instance of `UpstreamMapInput` via:
//
//	UpstreamMap{ "key": UpstreamArgs{...} }
type UpstreamMapInput interface {
	pulumi.Input

	ToUpstreamMapOutput() UpstreamMapOutput
	ToUpstreamMapOutputWithContext(context.Context) UpstreamMapOutput
}

type UpstreamMap map[string]UpstreamInput

func (UpstreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Upstream)(nil)).Elem()
}

func (i UpstreamMap) ToUpstreamMapOutput() UpstreamMapOutput {
	return i.ToUpstreamMapOutputWithContext(context.Background())
}

func (i UpstreamMap) ToUpstreamMapOutputWithContext(ctx context.Context) UpstreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamMapOutput)
}

type UpstreamOutput struct{ *pulumi.OutputState }

func (UpstreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Upstream)(nil)).Elem()
}

func (o UpstreamOutput) ToUpstreamOutput() UpstreamOutput {
	return o
}

func (o UpstreamOutput) ToUpstreamOutputWithContext(ctx context.Context) UpstreamOutput {
	return o
}

// Load balancing algorithm, value range: ROUND-ROBIN.
func (o UpstreamOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *Upstream) pulumi.StringOutput { return v.Algorithm }).(pulumi.StringOutput)
}

// Health check configuration, currently only supports VPC channels.
func (o UpstreamOutput) HealthChecker() UpstreamHealthCheckerPtrOutput {
	return o.ApplyT(func(v *Upstream) UpstreamHealthCheckerPtrOutput { return v.HealthChecker }).(UpstreamHealthCheckerPtrOutput)
}

// Configuration of K8S container service.
func (o UpstreamOutput) K8sServices() UpstreamK8sServiceArrayOutput {
	return o.ApplyT(func(v *Upstream) UpstreamK8sServiceArrayOutput { return v.K8sServices }).(UpstreamK8sServiceArrayOutput)
}

// Backend nodes.
func (o UpstreamOutput) Nodes() UpstreamNodeArrayOutput {
	return o.ApplyT(func(v *Upstream) UpstreamNodeArrayOutput { return v.Nodes }).(UpstreamNodeArrayOutput)
}

// Request retry count, default to 3 times.
func (o UpstreamOutput) Retries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Upstream) pulumi.IntPtrOutput { return v.Retries }).(pulumi.IntPtrOutput)
}

// Backend protocol, value range: HTTP, HTTPS, gRPC, gRPCs.
func (o UpstreamOutput) Scheme() pulumi.StringOutput {
	return o.ApplyT(func(v *Upstream) pulumi.StringOutput { return v.Scheme }).(pulumi.StringOutput)
}

// Tag description list.
func (o UpstreamOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Upstream) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// VPC Unique ID.
func (o UpstreamOutput) UniqVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Upstream) pulumi.StringOutput { return v.UniqVpcId }).(pulumi.StringOutput)
}

// Backend channel description.
func (o UpstreamOutput) UpstreamDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Upstream) pulumi.StringPtrOutput { return v.UpstreamDescription }).(pulumi.StringPtrOutput)
}

// Host request header forwarded by gateway to backend.
func (o UpstreamOutput) UpstreamHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Upstream) pulumi.StringPtrOutput { return v.UpstreamHost }).(pulumi.StringPtrOutput)
}

// Backend channel name.
func (o UpstreamOutput) UpstreamName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Upstream) pulumi.StringPtrOutput { return v.UpstreamName }).(pulumi.StringPtrOutput)
}

// Backend access type, value range: IP_PORT, K8S.
func (o UpstreamOutput) UpstreamType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Upstream) pulumi.StringPtrOutput { return v.UpstreamType }).(pulumi.StringPtrOutput)
}

type UpstreamArrayOutput struct{ *pulumi.OutputState }

func (UpstreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Upstream)(nil)).Elem()
}

func (o UpstreamArrayOutput) ToUpstreamArrayOutput() UpstreamArrayOutput {
	return o
}

func (o UpstreamArrayOutput) ToUpstreamArrayOutputWithContext(ctx context.Context) UpstreamArrayOutput {
	return o
}

func (o UpstreamArrayOutput) Index(i pulumi.IntInput) UpstreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Upstream {
		return vs[0].([]*Upstream)[vs[1].(int)]
	}).(UpstreamOutput)
}

type UpstreamMapOutput struct{ *pulumi.OutputState }

func (UpstreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Upstream)(nil)).Elem()
}

func (o UpstreamMapOutput) ToUpstreamMapOutput() UpstreamMapOutput {
	return o
}

func (o UpstreamMapOutput) ToUpstreamMapOutputWithContext(ctx context.Context) UpstreamMapOutput {
	return o
}

func (o UpstreamMapOutput) MapIndex(k pulumi.StringInput) UpstreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Upstream {
		return vs[0].(map[string]*Upstream)[vs[1].(string)]
	}).(UpstreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamInput)(nil)).Elem(), &Upstream{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamArrayInput)(nil)).Elem(), UpstreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpstreamMapInput)(nil)).Elem(), UpstreamMap{})
	pulumi.RegisterOutputType(UpstreamOutput{})
	pulumi.RegisterOutputType(UpstreamArrayOutput{})
	pulumi.RegisterOutputType(UpstreamMapOutput{})
}
