// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tse cngwService
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tse"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Subnet"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tse"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			availabilityZone := "ap-guangzhou-4"
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
//			cngwGateway, err := Tse.NewCngwGateway(ctx, "cngwGateway", &Tse.CngwGatewayArgs{
//				Description:             pulumi.String("terraform test1"),
//				EnableCls:               pulumi.Bool(true),
//				EngineRegion:            pulumi.String("ap-guangzhou"),
//				FeatureVersion:          pulumi.String("STANDARD"),
//				GatewayVersion:          pulumi.String("2.5.1"),
//				IngressClassName:        pulumi.String("tse-nginx-ingress"),
//				InternetMaxBandwidthOut: pulumi.Int(0),
//				TradeType:               pulumi.Int(0),
//				Type:                    pulumi.String("kong"),
//				NodeConfig: &tse.CngwGatewayNodeConfigArgs{
//					Number:        pulumi.Int(2),
//					Specification: pulumi.String("1c2g"),
//				},
//				VpcConfig: &tse.CngwGatewayVpcConfigArgs{
//					SubnetId: subnet.ID(),
//					VpcId:    vpc.ID(),
//				},
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("terraform"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Tse.NewCngwService(ctx, "cngwService", &Tse.CngwServiceArgs{
//				GatewayId:    cngwGateway.ID(),
//				Path:         pulumi.String("/test"),
//				Protocol:     pulumi.String("http"),
//				Retries:      pulumi.Int(5),
//				Timeout:      pulumi.Int(60000),
//				UpstreamType: pulumi.String("HostIP"),
//				UpstreamInfo: &tse.CngwServiceUpstreamInfoArgs{
//					Algorithm:          pulumi.String("round-robin"),
//					AutoScalingCvmPort: pulumi.Int(0),
//					Host:               pulumi.String("arunma.cn"),
//					Port:               pulumi.Int(8012),
//					SlowStart:          pulumi.Int(0),
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
// tse cngw_service can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Tse/cngwService:CngwService cngw_service gatewayId#name
//
// ```
type CngwService struct {
	pulumi.CustomResourceState

	// gateway ID.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// path.
	Path pulumi.StringOutput `pulumi:"path"`
	// protocol. Reference value:`https`, `http`, `tcp`, `udp`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// retry times.
	Retries pulumi.IntOutput `pulumi:"retries"`
	// service id.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// Deprecate ineffective tags Tag description list.
	//
	// Deprecated: Deprecate ineffective tags
	Tags pulumi.MapOutput `pulumi:"tags"`
	// time out, unit:ms.
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// service config information.
	UpstreamInfo CngwServiceUpstreamInfoOutput `pulumi:"upstreamInfo"`
	// service type. Reference value:`Kubernetes`, `Registry`, `IPList`, `HostIP`, `Scf`.
	UpstreamType pulumi.StringOutput `pulumi:"upstreamType"`
}

// NewCngwService registers a new resource with the given unique name, arguments, and options.
func NewCngwService(ctx *pulumi.Context,
	name string, args *CngwServiceArgs, opts ...pulumi.ResourceOption) (*CngwService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Retries == nil {
		return nil, errors.New("invalid value for required argument 'Retries'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	if args.UpstreamInfo == nil {
		return nil, errors.New("invalid value for required argument 'UpstreamInfo'")
	}
	if args.UpstreamType == nil {
		return nil, errors.New("invalid value for required argument 'UpstreamType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CngwService
	err := ctx.RegisterResource("tencentcloud:Tse/cngwService:CngwService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCngwService gets an existing CngwService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCngwService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CngwServiceState, opts ...pulumi.ResourceOption) (*CngwService, error) {
	var resource CngwService
	err := ctx.ReadResource("tencentcloud:Tse/cngwService:CngwService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CngwService resources.
type cngwServiceState struct {
	// gateway ID.
	GatewayId *string `pulumi:"gatewayId"`
	// service name.
	Name *string `pulumi:"name"`
	// path.
	Path *string `pulumi:"path"`
	// protocol. Reference value:`https`, `http`, `tcp`, `udp`.
	Protocol *string `pulumi:"protocol"`
	// retry times.
	Retries *int `pulumi:"retries"`
	// service id.
	ServiceId *string `pulumi:"serviceId"`
	// Deprecate ineffective tags Tag description list.
	//
	// Deprecated: Deprecate ineffective tags
	Tags map[string]interface{} `pulumi:"tags"`
	// time out, unit:ms.
	Timeout *int `pulumi:"timeout"`
	// service config information.
	UpstreamInfo *CngwServiceUpstreamInfo `pulumi:"upstreamInfo"`
	// service type. Reference value:`Kubernetes`, `Registry`, `IPList`, `HostIP`, `Scf`.
	UpstreamType *string `pulumi:"upstreamType"`
}

type CngwServiceState struct {
	// gateway ID.
	GatewayId pulumi.StringPtrInput
	// service name.
	Name pulumi.StringPtrInput
	// path.
	Path pulumi.StringPtrInput
	// protocol. Reference value:`https`, `http`, `tcp`, `udp`.
	Protocol pulumi.StringPtrInput
	// retry times.
	Retries pulumi.IntPtrInput
	// service id.
	ServiceId pulumi.StringPtrInput
	// Deprecate ineffective tags Tag description list.
	//
	// Deprecated: Deprecate ineffective tags
	Tags pulumi.MapInput
	// time out, unit:ms.
	Timeout pulumi.IntPtrInput
	// service config information.
	UpstreamInfo CngwServiceUpstreamInfoPtrInput
	// service type. Reference value:`Kubernetes`, `Registry`, `IPList`, `HostIP`, `Scf`.
	UpstreamType pulumi.StringPtrInput
}

func (CngwServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*cngwServiceState)(nil)).Elem()
}

type cngwServiceArgs struct {
	// gateway ID.
	GatewayId string `pulumi:"gatewayId"`
	// service name.
	Name *string `pulumi:"name"`
	// path.
	Path string `pulumi:"path"`
	// protocol. Reference value:`https`, `http`, `tcp`, `udp`.
	Protocol string `pulumi:"protocol"`
	// retry times.
	Retries int `pulumi:"retries"`
	// Deprecate ineffective tags Tag description list.
	//
	// Deprecated: Deprecate ineffective tags
	Tags map[string]interface{} `pulumi:"tags"`
	// time out, unit:ms.
	Timeout int `pulumi:"timeout"`
	// service config information.
	UpstreamInfo CngwServiceUpstreamInfo `pulumi:"upstreamInfo"`
	// service type. Reference value:`Kubernetes`, `Registry`, `IPList`, `HostIP`, `Scf`.
	UpstreamType string `pulumi:"upstreamType"`
}

// The set of arguments for constructing a CngwService resource.
type CngwServiceArgs struct {
	// gateway ID.
	GatewayId pulumi.StringInput
	// service name.
	Name pulumi.StringPtrInput
	// path.
	Path pulumi.StringInput
	// protocol. Reference value:`https`, `http`, `tcp`, `udp`.
	Protocol pulumi.StringInput
	// retry times.
	Retries pulumi.IntInput
	// Deprecate ineffective tags Tag description list.
	//
	// Deprecated: Deprecate ineffective tags
	Tags pulumi.MapInput
	// time out, unit:ms.
	Timeout pulumi.IntInput
	// service config information.
	UpstreamInfo CngwServiceUpstreamInfoInput
	// service type. Reference value:`Kubernetes`, `Registry`, `IPList`, `HostIP`, `Scf`.
	UpstreamType pulumi.StringInput
}

func (CngwServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cngwServiceArgs)(nil)).Elem()
}

type CngwServiceInput interface {
	pulumi.Input

	ToCngwServiceOutput() CngwServiceOutput
	ToCngwServiceOutputWithContext(ctx context.Context) CngwServiceOutput
}

func (*CngwService) ElementType() reflect.Type {
	return reflect.TypeOf((**CngwService)(nil)).Elem()
}

func (i *CngwService) ToCngwServiceOutput() CngwServiceOutput {
	return i.ToCngwServiceOutputWithContext(context.Background())
}

func (i *CngwService) ToCngwServiceOutputWithContext(ctx context.Context) CngwServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CngwServiceOutput)
}

// CngwServiceArrayInput is an input type that accepts CngwServiceArray and CngwServiceArrayOutput values.
// You can construct a concrete instance of `CngwServiceArrayInput` via:
//
//	CngwServiceArray{ CngwServiceArgs{...} }
type CngwServiceArrayInput interface {
	pulumi.Input

	ToCngwServiceArrayOutput() CngwServiceArrayOutput
	ToCngwServiceArrayOutputWithContext(context.Context) CngwServiceArrayOutput
}

type CngwServiceArray []CngwServiceInput

func (CngwServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CngwService)(nil)).Elem()
}

func (i CngwServiceArray) ToCngwServiceArrayOutput() CngwServiceArrayOutput {
	return i.ToCngwServiceArrayOutputWithContext(context.Background())
}

func (i CngwServiceArray) ToCngwServiceArrayOutputWithContext(ctx context.Context) CngwServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CngwServiceArrayOutput)
}

// CngwServiceMapInput is an input type that accepts CngwServiceMap and CngwServiceMapOutput values.
// You can construct a concrete instance of `CngwServiceMapInput` via:
//
//	CngwServiceMap{ "key": CngwServiceArgs{...} }
type CngwServiceMapInput interface {
	pulumi.Input

	ToCngwServiceMapOutput() CngwServiceMapOutput
	ToCngwServiceMapOutputWithContext(context.Context) CngwServiceMapOutput
}

type CngwServiceMap map[string]CngwServiceInput

func (CngwServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CngwService)(nil)).Elem()
}

func (i CngwServiceMap) ToCngwServiceMapOutput() CngwServiceMapOutput {
	return i.ToCngwServiceMapOutputWithContext(context.Background())
}

func (i CngwServiceMap) ToCngwServiceMapOutputWithContext(ctx context.Context) CngwServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CngwServiceMapOutput)
}

type CngwServiceOutput struct{ *pulumi.OutputState }

func (CngwServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CngwService)(nil)).Elem()
}

func (o CngwServiceOutput) ToCngwServiceOutput() CngwServiceOutput {
	return o
}

func (o CngwServiceOutput) ToCngwServiceOutputWithContext(ctx context.Context) CngwServiceOutput {
	return o
}

// gateway ID.
func (o CngwServiceOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwService) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// service name.
func (o CngwServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// path.
func (o CngwServiceOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwService) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// protocol. Reference value:`https`, `http`, `tcp`, `udp`.
func (o CngwServiceOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwService) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// retry times.
func (o CngwServiceOutput) Retries() pulumi.IntOutput {
	return o.ApplyT(func(v *CngwService) pulumi.IntOutput { return v.Retries }).(pulumi.IntOutput)
}

// service id.
func (o CngwServiceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwService) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// Deprecate ineffective tags Tag description list.
//
// Deprecated: Deprecate ineffective tags
func (o CngwServiceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *CngwService) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// time out, unit:ms.
func (o CngwServiceOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *CngwService) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

// service config information.
func (o CngwServiceOutput) UpstreamInfo() CngwServiceUpstreamInfoOutput {
	return o.ApplyT(func(v *CngwService) CngwServiceUpstreamInfoOutput { return v.UpstreamInfo }).(CngwServiceUpstreamInfoOutput)
}

// service type. Reference value:`Kubernetes`, `Registry`, `IPList`, `HostIP`, `Scf`.
func (o CngwServiceOutput) UpstreamType() pulumi.StringOutput {
	return o.ApplyT(func(v *CngwService) pulumi.StringOutput { return v.UpstreamType }).(pulumi.StringOutput)
}

type CngwServiceArrayOutput struct{ *pulumi.OutputState }

func (CngwServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CngwService)(nil)).Elem()
}

func (o CngwServiceArrayOutput) ToCngwServiceArrayOutput() CngwServiceArrayOutput {
	return o
}

func (o CngwServiceArrayOutput) ToCngwServiceArrayOutputWithContext(ctx context.Context) CngwServiceArrayOutput {
	return o
}

func (o CngwServiceArrayOutput) Index(i pulumi.IntInput) CngwServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CngwService {
		return vs[0].([]*CngwService)[vs[1].(int)]
	}).(CngwServiceOutput)
}

type CngwServiceMapOutput struct{ *pulumi.OutputState }

func (CngwServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CngwService)(nil)).Elem()
}

func (o CngwServiceMapOutput) ToCngwServiceMapOutput() CngwServiceMapOutput {
	return o
}

func (o CngwServiceMapOutput) ToCngwServiceMapOutputWithContext(ctx context.Context) CngwServiceMapOutput {
	return o
}

func (o CngwServiceMapOutput) MapIndex(k pulumi.StringInput) CngwServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CngwService {
		return vs[0].(map[string]*CngwService)[vs[1].(string)]
	}).(CngwServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CngwServiceInput)(nil)).Elem(), &CngwService{})
	pulumi.RegisterInputType(reflect.TypeOf((*CngwServiceArrayInput)(nil)).Elem(), CngwServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CngwServiceMapInput)(nil)).Elem(), CngwServiceMap{})
	pulumi.RegisterOutputType(CngwServiceOutput{})
	pulumi.RegisterOutputType(CngwServiceArrayOutput{})
	pulumi.RegisterOutputType(CngwServiceMapOutput{})
}
