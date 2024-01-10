// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dayu

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create dayu DDoS policy case
//
// > **NOTE:** when a dayu DDoS policy case is created, there will be a dayu DDoS policy created with the same prefix name in the same time. This resource only supports Anti-DDoS of type `bgp`, `bgp-multip` and `bgpip`. One Anti-DDoS resource can only has one DDoS policy case resource. When there is only one Anti-DDoS resource and one policy case, those two resource will be bind automatically.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dayu"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dayu.NewDdosPolicyCase(ctx, "foo", &Dayu.DdosPolicyCaseArgs{
//				AppProtocols: pulumi.StringArray{
//					pulumi.String("tcp"),
//					pulumi.String("udp"),
//				},
//				AppType:          pulumi.String("WEB"),
//				HasAbroad:        pulumi.String("yes"),
//				HasInitiateTcp:   pulumi.String("yes"),
//				HasInitiateUdp:   pulumi.String("yes"),
//				HasVpn:           pulumi.String("yes"),
//				MaxTcpPackageLen: pulumi.String("1200"),
//				MaxUdpPackageLen: pulumi.String("1200"),
//				MinTcpPackageLen: pulumi.String("1000"),
//				MinUdpPackageLen: pulumi.String("1000"),
//				PeerTcpPort:      pulumi.String("1111"),
//				PeerUdpPort:      pulumi.String("3333"),
//				PlatformTypes: pulumi.StringArray{
//					pulumi.String("PC"),
//					pulumi.String("MOBILE"),
//				},
//				ResourceType: pulumi.String("bgpip"),
//				TcpEndPort:   pulumi.String("2000"),
//				TcpFootprint: pulumi.String("511"),
//				TcpStartPort: pulumi.String("1000"),
//				UdpEndPort:   pulumi.String("4000"),
//				UdpFootprint: pulumi.String("500"),
//				UdpStartPort: pulumi.String("3000"),
//				WebApiUrls: pulumi.StringArray{
//					pulumi.String("abc.com"),
//					pulumi.String("test.cn/aaa.png"),
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
type DdosPolicyCase struct {
	pulumi.CustomResourceState

	// App protocol set of the DDoS policy case.
	AppProtocols pulumi.StringArrayOutput `pulumi:"appProtocols"`
	// App type of the DDoS policy case. Valid values: `WEB`, `GAME`, `APP` and `OTHER`.
	AppType pulumi.StringOutput `pulumi:"appType"`
	// Create time of the DDoS policy case.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicate whether the service involves overseas or not. Valid values: `no` and `yes`.
	HasAbroad pulumi.StringOutput `pulumi:"hasAbroad"`
	// Indicate whether the service actively initiates TCP requests or not. Valid values: `no` and `yes`.
	HasInitiateTcp pulumi.StringOutput `pulumi:"hasInitiateTcp"`
	// Indicate whether the actively initiate UDP requests or not. Valid values: `no` and `yes`.
	HasInitiateUdp pulumi.StringPtrOutput `pulumi:"hasInitiateUdp"`
	// Indicate whether the service involves VPN service or not. Valid values: `no` and `yes`.
	HasVpn pulumi.StringPtrOutput `pulumi:"hasVpn"`
	// The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minTcpPackageLen`.
	MaxTcpPackageLen pulumi.StringPtrOutput `pulumi:"maxTcpPackageLen"`
	// The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minUdpPackageLen`.
	MaxUdpPackageLen pulumi.StringPtrOutput `pulumi:"maxUdpPackageLen"`
	// The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.
	MinTcpPackageLen pulumi.StringPtrOutput `pulumi:"minTcpPackageLen"`
	// The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.
	MinUdpPackageLen pulumi.StringPtrOutput `pulumi:"minUdpPackageLen"`
	// Name of the DDoS policy case. Length should between 1 and 64.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port that actively initiates TCP requests. Valid value ranges: (1~65535).
	PeerTcpPort pulumi.StringPtrOutput `pulumi:"peerTcpPort"`
	// The port that actively initiates UDP requests. Valid value ranges: (1~65535).
	PeerUdpPort pulumi.StringPtrOutput `pulumi:"peerUdpPort"`
	// Platform set of the DDoS policy case.
	PlatformTypes pulumi.StringArrayOutput `pulumi:"platformTypes"`
	// Type of the resource that the DDoS policy case works for. Valid values: `bgpip`, `bgp` and `bgp-multip`.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// ID of the DDoS policy case.
	SceneId pulumi.StringOutput `pulumi:"sceneId"`
	// End port of the TCP service. Valid value ranges: (0~65535). It must be greater than `tcpStartPort`.
	TcpEndPort pulumi.StringOutput `pulumi:"tcpEndPort"`
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	TcpFootprint pulumi.StringPtrOutput `pulumi:"tcpFootprint"`
	// Start port of the TCP service. Valid value ranges: (0~65535).
	TcpStartPort pulumi.StringOutput `pulumi:"tcpStartPort"`
	// End port of the UDP service. Valid value ranges: (0~65535). It must be greater than `udpStartPort`.
	UdpEndPort pulumi.StringOutput `pulumi:"udpEndPort"`
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	UdpFootprint pulumi.StringPtrOutput `pulumi:"udpFootprint"`
	// Start port of the UDP service. Valid value ranges: (0~65535).
	UdpStartPort pulumi.StringOutput `pulumi:"udpStartPort"`
	// Web API url set.
	WebApiUrls pulumi.StringArrayOutput `pulumi:"webApiUrls"`
}

// NewDdosPolicyCase registers a new resource with the given unique name, arguments, and options.
func NewDdosPolicyCase(ctx *pulumi.Context,
	name string, args *DdosPolicyCaseArgs, opts ...pulumi.ResourceOption) (*DdosPolicyCase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppProtocols == nil {
		return nil, errors.New("invalid value for required argument 'AppProtocols'")
	}
	if args.AppType == nil {
		return nil, errors.New("invalid value for required argument 'AppType'")
	}
	if args.HasAbroad == nil {
		return nil, errors.New("invalid value for required argument 'HasAbroad'")
	}
	if args.HasInitiateTcp == nil {
		return nil, errors.New("invalid value for required argument 'HasInitiateTcp'")
	}
	if args.PlatformTypes == nil {
		return nil, errors.New("invalid value for required argument 'PlatformTypes'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	if args.TcpEndPort == nil {
		return nil, errors.New("invalid value for required argument 'TcpEndPort'")
	}
	if args.TcpStartPort == nil {
		return nil, errors.New("invalid value for required argument 'TcpStartPort'")
	}
	if args.UdpEndPort == nil {
		return nil, errors.New("invalid value for required argument 'UdpEndPort'")
	}
	if args.UdpStartPort == nil {
		return nil, errors.New("invalid value for required argument 'UdpStartPort'")
	}
	if args.WebApiUrls == nil {
		return nil, errors.New("invalid value for required argument 'WebApiUrls'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DdosPolicyCase
	err := ctx.RegisterResource("tencentcloud:Dayu/ddosPolicyCase:DdosPolicyCase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosPolicyCase gets an existing DdosPolicyCase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosPolicyCase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosPolicyCaseState, opts ...pulumi.ResourceOption) (*DdosPolicyCase, error) {
	var resource DdosPolicyCase
	err := ctx.ReadResource("tencentcloud:Dayu/ddosPolicyCase:DdosPolicyCase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosPolicyCase resources.
type ddosPolicyCaseState struct {
	// App protocol set of the DDoS policy case.
	AppProtocols []string `pulumi:"appProtocols"`
	// App type of the DDoS policy case. Valid values: `WEB`, `GAME`, `APP` and `OTHER`.
	AppType *string `pulumi:"appType"`
	// Create time of the DDoS policy case.
	CreateTime *string `pulumi:"createTime"`
	// Indicate whether the service involves overseas or not. Valid values: `no` and `yes`.
	HasAbroad *string `pulumi:"hasAbroad"`
	// Indicate whether the service actively initiates TCP requests or not. Valid values: `no` and `yes`.
	HasInitiateTcp *string `pulumi:"hasInitiateTcp"`
	// Indicate whether the actively initiate UDP requests or not. Valid values: `no` and `yes`.
	HasInitiateUdp *string `pulumi:"hasInitiateUdp"`
	// Indicate whether the service involves VPN service or not. Valid values: `no` and `yes`.
	HasVpn *string `pulumi:"hasVpn"`
	// The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minTcpPackageLen`.
	MaxTcpPackageLen *string `pulumi:"maxTcpPackageLen"`
	// The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minUdpPackageLen`.
	MaxUdpPackageLen *string `pulumi:"maxUdpPackageLen"`
	// The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.
	MinTcpPackageLen *string `pulumi:"minTcpPackageLen"`
	// The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.
	MinUdpPackageLen *string `pulumi:"minUdpPackageLen"`
	// Name of the DDoS policy case. Length should between 1 and 64.
	Name *string `pulumi:"name"`
	// The port that actively initiates TCP requests. Valid value ranges: (1~65535).
	PeerTcpPort *string `pulumi:"peerTcpPort"`
	// The port that actively initiates UDP requests. Valid value ranges: (1~65535).
	PeerUdpPort *string `pulumi:"peerUdpPort"`
	// Platform set of the DDoS policy case.
	PlatformTypes []string `pulumi:"platformTypes"`
	// Type of the resource that the DDoS policy case works for. Valid values: `bgpip`, `bgp` and `bgp-multip`.
	ResourceType *string `pulumi:"resourceType"`
	// ID of the DDoS policy case.
	SceneId *string `pulumi:"sceneId"`
	// End port of the TCP service. Valid value ranges: (0~65535). It must be greater than `tcpStartPort`.
	TcpEndPort *string `pulumi:"tcpEndPort"`
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	TcpFootprint *string `pulumi:"tcpFootprint"`
	// Start port of the TCP service. Valid value ranges: (0~65535).
	TcpStartPort *string `pulumi:"tcpStartPort"`
	// End port of the UDP service. Valid value ranges: (0~65535). It must be greater than `udpStartPort`.
	UdpEndPort *string `pulumi:"udpEndPort"`
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	UdpFootprint *string `pulumi:"udpFootprint"`
	// Start port of the UDP service. Valid value ranges: (0~65535).
	UdpStartPort *string `pulumi:"udpStartPort"`
	// Web API url set.
	WebApiUrls []string `pulumi:"webApiUrls"`
}

type DdosPolicyCaseState struct {
	// App protocol set of the DDoS policy case.
	AppProtocols pulumi.StringArrayInput
	// App type of the DDoS policy case. Valid values: `WEB`, `GAME`, `APP` and `OTHER`.
	AppType pulumi.StringPtrInput
	// Create time of the DDoS policy case.
	CreateTime pulumi.StringPtrInput
	// Indicate whether the service involves overseas or not. Valid values: `no` and `yes`.
	HasAbroad pulumi.StringPtrInput
	// Indicate whether the service actively initiates TCP requests or not. Valid values: `no` and `yes`.
	HasInitiateTcp pulumi.StringPtrInput
	// Indicate whether the actively initiate UDP requests or not. Valid values: `no` and `yes`.
	HasInitiateUdp pulumi.StringPtrInput
	// Indicate whether the service involves VPN service or not. Valid values: `no` and `yes`.
	HasVpn pulumi.StringPtrInput
	// The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minTcpPackageLen`.
	MaxTcpPackageLen pulumi.StringPtrInput
	// The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minUdpPackageLen`.
	MaxUdpPackageLen pulumi.StringPtrInput
	// The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.
	MinTcpPackageLen pulumi.StringPtrInput
	// The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.
	MinUdpPackageLen pulumi.StringPtrInput
	// Name of the DDoS policy case. Length should between 1 and 64.
	Name pulumi.StringPtrInput
	// The port that actively initiates TCP requests. Valid value ranges: (1~65535).
	PeerTcpPort pulumi.StringPtrInput
	// The port that actively initiates UDP requests. Valid value ranges: (1~65535).
	PeerUdpPort pulumi.StringPtrInput
	// Platform set of the DDoS policy case.
	PlatformTypes pulumi.StringArrayInput
	// Type of the resource that the DDoS policy case works for. Valid values: `bgpip`, `bgp` and `bgp-multip`.
	ResourceType pulumi.StringPtrInput
	// ID of the DDoS policy case.
	SceneId pulumi.StringPtrInput
	// End port of the TCP service. Valid value ranges: (0~65535). It must be greater than `tcpStartPort`.
	TcpEndPort pulumi.StringPtrInput
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	TcpFootprint pulumi.StringPtrInput
	// Start port of the TCP service. Valid value ranges: (0~65535).
	TcpStartPort pulumi.StringPtrInput
	// End port of the UDP service. Valid value ranges: (0~65535). It must be greater than `udpStartPort`.
	UdpEndPort pulumi.StringPtrInput
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	UdpFootprint pulumi.StringPtrInput
	// Start port of the UDP service. Valid value ranges: (0~65535).
	UdpStartPort pulumi.StringPtrInput
	// Web API url set.
	WebApiUrls pulumi.StringArrayInput
}

func (DdosPolicyCaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosPolicyCaseState)(nil)).Elem()
}

type ddosPolicyCaseArgs struct {
	// App protocol set of the DDoS policy case.
	AppProtocols []string `pulumi:"appProtocols"`
	// App type of the DDoS policy case. Valid values: `WEB`, `GAME`, `APP` and `OTHER`.
	AppType string `pulumi:"appType"`
	// Indicate whether the service involves overseas or not. Valid values: `no` and `yes`.
	HasAbroad string `pulumi:"hasAbroad"`
	// Indicate whether the service actively initiates TCP requests or not. Valid values: `no` and `yes`.
	HasInitiateTcp string `pulumi:"hasInitiateTcp"`
	// Indicate whether the actively initiate UDP requests or not. Valid values: `no` and `yes`.
	HasInitiateUdp *string `pulumi:"hasInitiateUdp"`
	// Indicate whether the service involves VPN service or not. Valid values: `no` and `yes`.
	HasVpn *string `pulumi:"hasVpn"`
	// The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minTcpPackageLen`.
	MaxTcpPackageLen *string `pulumi:"maxTcpPackageLen"`
	// The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minUdpPackageLen`.
	MaxUdpPackageLen *string `pulumi:"maxUdpPackageLen"`
	// The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.
	MinTcpPackageLen *string `pulumi:"minTcpPackageLen"`
	// The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.
	MinUdpPackageLen *string `pulumi:"minUdpPackageLen"`
	// Name of the DDoS policy case. Length should between 1 and 64.
	Name *string `pulumi:"name"`
	// The port that actively initiates TCP requests. Valid value ranges: (1~65535).
	PeerTcpPort *string `pulumi:"peerTcpPort"`
	// The port that actively initiates UDP requests. Valid value ranges: (1~65535).
	PeerUdpPort *string `pulumi:"peerUdpPort"`
	// Platform set of the DDoS policy case.
	PlatformTypes []string `pulumi:"platformTypes"`
	// Type of the resource that the DDoS policy case works for. Valid values: `bgpip`, `bgp` and `bgp-multip`.
	ResourceType string `pulumi:"resourceType"`
	// End port of the TCP service. Valid value ranges: (0~65535). It must be greater than `tcpStartPort`.
	TcpEndPort string `pulumi:"tcpEndPort"`
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	TcpFootprint *string `pulumi:"tcpFootprint"`
	// Start port of the TCP service. Valid value ranges: (0~65535).
	TcpStartPort string `pulumi:"tcpStartPort"`
	// End port of the UDP service. Valid value ranges: (0~65535). It must be greater than `udpStartPort`.
	UdpEndPort string `pulumi:"udpEndPort"`
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	UdpFootprint *string `pulumi:"udpFootprint"`
	// Start port of the UDP service. Valid value ranges: (0~65535).
	UdpStartPort string `pulumi:"udpStartPort"`
	// Web API url set.
	WebApiUrls []string `pulumi:"webApiUrls"`
}

// The set of arguments for constructing a DdosPolicyCase resource.
type DdosPolicyCaseArgs struct {
	// App protocol set of the DDoS policy case.
	AppProtocols pulumi.StringArrayInput
	// App type of the DDoS policy case. Valid values: `WEB`, `GAME`, `APP` and `OTHER`.
	AppType pulumi.StringInput
	// Indicate whether the service involves overseas or not. Valid values: `no` and `yes`.
	HasAbroad pulumi.StringInput
	// Indicate whether the service actively initiates TCP requests or not. Valid values: `no` and `yes`.
	HasInitiateTcp pulumi.StringInput
	// Indicate whether the actively initiate UDP requests or not. Valid values: `no` and `yes`.
	HasInitiateUdp pulumi.StringPtrInput
	// Indicate whether the service involves VPN service or not. Valid values: `no` and `yes`.
	HasVpn pulumi.StringPtrInput
	// The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minTcpPackageLen`.
	MaxTcpPackageLen pulumi.StringPtrInput
	// The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minUdpPackageLen`.
	MaxUdpPackageLen pulumi.StringPtrInput
	// The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.
	MinTcpPackageLen pulumi.StringPtrInput
	// The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.
	MinUdpPackageLen pulumi.StringPtrInput
	// Name of the DDoS policy case. Length should between 1 and 64.
	Name pulumi.StringPtrInput
	// The port that actively initiates TCP requests. Valid value ranges: (1~65535).
	PeerTcpPort pulumi.StringPtrInput
	// The port that actively initiates UDP requests. Valid value ranges: (1~65535).
	PeerUdpPort pulumi.StringPtrInput
	// Platform set of the DDoS policy case.
	PlatformTypes pulumi.StringArrayInput
	// Type of the resource that the DDoS policy case works for. Valid values: `bgpip`, `bgp` and `bgp-multip`.
	ResourceType pulumi.StringInput
	// End port of the TCP service. Valid value ranges: (0~65535). It must be greater than `tcpStartPort`.
	TcpEndPort pulumi.StringInput
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	TcpFootprint pulumi.StringPtrInput
	// Start port of the TCP service. Valid value ranges: (0~65535).
	TcpStartPort pulumi.StringInput
	// End port of the UDP service. Valid value ranges: (0~65535). It must be greater than `udpStartPort`.
	UdpEndPort pulumi.StringInput
	// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
	UdpFootprint pulumi.StringPtrInput
	// Start port of the UDP service. Valid value ranges: (0~65535).
	UdpStartPort pulumi.StringInput
	// Web API url set.
	WebApiUrls pulumi.StringArrayInput
}

func (DdosPolicyCaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosPolicyCaseArgs)(nil)).Elem()
}

type DdosPolicyCaseInput interface {
	pulumi.Input

	ToDdosPolicyCaseOutput() DdosPolicyCaseOutput
	ToDdosPolicyCaseOutputWithContext(ctx context.Context) DdosPolicyCaseOutput
}

func (*DdosPolicyCase) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosPolicyCase)(nil)).Elem()
}

func (i *DdosPolicyCase) ToDdosPolicyCaseOutput() DdosPolicyCaseOutput {
	return i.ToDdosPolicyCaseOutputWithContext(context.Background())
}

func (i *DdosPolicyCase) ToDdosPolicyCaseOutputWithContext(ctx context.Context) DdosPolicyCaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyCaseOutput)
}

// DdosPolicyCaseArrayInput is an input type that accepts DdosPolicyCaseArray and DdosPolicyCaseArrayOutput values.
// You can construct a concrete instance of `DdosPolicyCaseArrayInput` via:
//
//	DdosPolicyCaseArray{ DdosPolicyCaseArgs{...} }
type DdosPolicyCaseArrayInput interface {
	pulumi.Input

	ToDdosPolicyCaseArrayOutput() DdosPolicyCaseArrayOutput
	ToDdosPolicyCaseArrayOutputWithContext(context.Context) DdosPolicyCaseArrayOutput
}

type DdosPolicyCaseArray []DdosPolicyCaseInput

func (DdosPolicyCaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosPolicyCase)(nil)).Elem()
}

func (i DdosPolicyCaseArray) ToDdosPolicyCaseArrayOutput() DdosPolicyCaseArrayOutput {
	return i.ToDdosPolicyCaseArrayOutputWithContext(context.Background())
}

func (i DdosPolicyCaseArray) ToDdosPolicyCaseArrayOutputWithContext(ctx context.Context) DdosPolicyCaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyCaseArrayOutput)
}

// DdosPolicyCaseMapInput is an input type that accepts DdosPolicyCaseMap and DdosPolicyCaseMapOutput values.
// You can construct a concrete instance of `DdosPolicyCaseMapInput` via:
//
//	DdosPolicyCaseMap{ "key": DdosPolicyCaseArgs{...} }
type DdosPolicyCaseMapInput interface {
	pulumi.Input

	ToDdosPolicyCaseMapOutput() DdosPolicyCaseMapOutput
	ToDdosPolicyCaseMapOutputWithContext(context.Context) DdosPolicyCaseMapOutput
}

type DdosPolicyCaseMap map[string]DdosPolicyCaseInput

func (DdosPolicyCaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosPolicyCase)(nil)).Elem()
}

func (i DdosPolicyCaseMap) ToDdosPolicyCaseMapOutput() DdosPolicyCaseMapOutput {
	return i.ToDdosPolicyCaseMapOutputWithContext(context.Background())
}

func (i DdosPolicyCaseMap) ToDdosPolicyCaseMapOutputWithContext(ctx context.Context) DdosPolicyCaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosPolicyCaseMapOutput)
}

type DdosPolicyCaseOutput struct{ *pulumi.OutputState }

func (DdosPolicyCaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosPolicyCase)(nil)).Elem()
}

func (o DdosPolicyCaseOutput) ToDdosPolicyCaseOutput() DdosPolicyCaseOutput {
	return o
}

func (o DdosPolicyCaseOutput) ToDdosPolicyCaseOutputWithContext(ctx context.Context) DdosPolicyCaseOutput {
	return o
}

// App protocol set of the DDoS policy case.
func (o DdosPolicyCaseOutput) AppProtocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringArrayOutput { return v.AppProtocols }).(pulumi.StringArrayOutput)
}

// App type of the DDoS policy case. Valid values: `WEB`, `GAME`, `APP` and `OTHER`.
func (o DdosPolicyCaseOutput) AppType() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.AppType }).(pulumi.StringOutput)
}

// Create time of the DDoS policy case.
func (o DdosPolicyCaseOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicate whether the service involves overseas or not. Valid values: `no` and `yes`.
func (o DdosPolicyCaseOutput) HasAbroad() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.HasAbroad }).(pulumi.StringOutput)
}

// Indicate whether the service actively initiates TCP requests or not. Valid values: `no` and `yes`.
func (o DdosPolicyCaseOutput) HasInitiateTcp() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.HasInitiateTcp }).(pulumi.StringOutput)
}

// Indicate whether the actively initiate UDP requests or not. Valid values: `no` and `yes`.
func (o DdosPolicyCaseOutput) HasInitiateUdp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.HasInitiateUdp }).(pulumi.StringPtrOutput)
}

// Indicate whether the service involves VPN service or not. Valid values: `no` and `yes`.
func (o DdosPolicyCaseOutput) HasVpn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.HasVpn }).(pulumi.StringPtrOutput)
}

// The max length of TCP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minTcpPackageLen`.
func (o DdosPolicyCaseOutput) MaxTcpPackageLen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.MaxTcpPackageLen }).(pulumi.StringPtrOutput)
}

// The max length of UDP message package, valid value length should be greater than 0 and less than 1500. It should be greater than `minUdpPackageLen`.
func (o DdosPolicyCaseOutput) MaxUdpPackageLen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.MaxUdpPackageLen }).(pulumi.StringPtrOutput)
}

// The minimum length of TCP message package, valid value length should be greater than 0 and less than 1500.
func (o DdosPolicyCaseOutput) MinTcpPackageLen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.MinTcpPackageLen }).(pulumi.StringPtrOutput)
}

// The minimum length of UDP message package, valid value length should be greater than 0 and less than 1500.
func (o DdosPolicyCaseOutput) MinUdpPackageLen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.MinUdpPackageLen }).(pulumi.StringPtrOutput)
}

// Name of the DDoS policy case. Length should between 1 and 64.
func (o DdosPolicyCaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port that actively initiates TCP requests. Valid value ranges: (1~65535).
func (o DdosPolicyCaseOutput) PeerTcpPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.PeerTcpPort }).(pulumi.StringPtrOutput)
}

// The port that actively initiates UDP requests. Valid value ranges: (1~65535).
func (o DdosPolicyCaseOutput) PeerUdpPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.PeerUdpPort }).(pulumi.StringPtrOutput)
}

// Platform set of the DDoS policy case.
func (o DdosPolicyCaseOutput) PlatformTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringArrayOutput { return v.PlatformTypes }).(pulumi.StringArrayOutput)
}

// Type of the resource that the DDoS policy case works for. Valid values: `bgpip`, `bgp` and `bgp-multip`.
func (o DdosPolicyCaseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// ID of the DDoS policy case.
func (o DdosPolicyCaseOutput) SceneId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.SceneId }).(pulumi.StringOutput)
}

// End port of the TCP service. Valid value ranges: (0~65535). It must be greater than `tcpStartPort`.
func (o DdosPolicyCaseOutput) TcpEndPort() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.TcpEndPort }).(pulumi.StringOutput)
}

// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
func (o DdosPolicyCaseOutput) TcpFootprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.TcpFootprint }).(pulumi.StringPtrOutput)
}

// Start port of the TCP service. Valid value ranges: (0~65535).
func (o DdosPolicyCaseOutput) TcpStartPort() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.TcpStartPort }).(pulumi.StringOutput)
}

// End port of the UDP service. Valid value ranges: (0~65535). It must be greater than `udpStartPort`.
func (o DdosPolicyCaseOutput) UdpEndPort() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.UdpEndPort }).(pulumi.StringOutput)
}

// The fixed signature of TCP protocol load, valid value length is range from 1 to 512.
func (o DdosPolicyCaseOutput) UdpFootprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringPtrOutput { return v.UdpFootprint }).(pulumi.StringPtrOutput)
}

// Start port of the UDP service. Valid value ranges: (0~65535).
func (o DdosPolicyCaseOutput) UdpStartPort() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringOutput { return v.UdpStartPort }).(pulumi.StringOutput)
}

// Web API url set.
func (o DdosPolicyCaseOutput) WebApiUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DdosPolicyCase) pulumi.StringArrayOutput { return v.WebApiUrls }).(pulumi.StringArrayOutput)
}

type DdosPolicyCaseArrayOutput struct{ *pulumi.OutputState }

func (DdosPolicyCaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosPolicyCase)(nil)).Elem()
}

func (o DdosPolicyCaseArrayOutput) ToDdosPolicyCaseArrayOutput() DdosPolicyCaseArrayOutput {
	return o
}

func (o DdosPolicyCaseArrayOutput) ToDdosPolicyCaseArrayOutputWithContext(ctx context.Context) DdosPolicyCaseArrayOutput {
	return o
}

func (o DdosPolicyCaseArrayOutput) Index(i pulumi.IntInput) DdosPolicyCaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DdosPolicyCase {
		return vs[0].([]*DdosPolicyCase)[vs[1].(int)]
	}).(DdosPolicyCaseOutput)
}

type DdosPolicyCaseMapOutput struct{ *pulumi.OutputState }

func (DdosPolicyCaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosPolicyCase)(nil)).Elem()
}

func (o DdosPolicyCaseMapOutput) ToDdosPolicyCaseMapOutput() DdosPolicyCaseMapOutput {
	return o
}

func (o DdosPolicyCaseMapOutput) ToDdosPolicyCaseMapOutputWithContext(ctx context.Context) DdosPolicyCaseMapOutput {
	return o
}

func (o DdosPolicyCaseMapOutput) MapIndex(k pulumi.StringInput) DdosPolicyCaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DdosPolicyCase {
		return vs[0].(map[string]*DdosPolicyCase)[vs[1].(string)]
	}).(DdosPolicyCaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyCaseInput)(nil)).Elem(), &DdosPolicyCase{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyCaseArrayInput)(nil)).Elem(), DdosPolicyCaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosPolicyCaseMapInput)(nil)).Elem(), DdosPolicyCaseMap{})
	pulumi.RegisterOutputType(DdosPolicyCaseOutput{})
	pulumi.RegisterOutputType(DdosPolicyCaseArrayOutput{})
	pulumi.RegisterOutputType(DdosPolicyCaseMapOutput{})
}
