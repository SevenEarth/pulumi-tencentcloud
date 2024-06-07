// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provide a resource to create a VPN SSL Client.
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpn.NewSslClient(ctx, "client", &Vpn.SslClientArgs{
//				SslVpnClientName: pulumi.String("hello"),
//				SslVpnServerId:   pulumi.String("vpns-aog5xcjj"),
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
// VPN SSL Client can be imported, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Vpn/sslClient:SslClient client vpn-client-id
// ```
type SslClient struct {
	pulumi.CustomResourceState

	// The name of ssl vpn client to be created.
	SslVpnClientName pulumi.StringOutput `pulumi:"sslVpnClientName"`
	// VPN ssl server id.
	SslVpnServerId pulumi.StringOutput `pulumi:"sslVpnServerId"`
}

// NewSslClient registers a new resource with the given unique name, arguments, and options.
func NewSslClient(ctx *pulumi.Context,
	name string, args *SslClientArgs, opts ...pulumi.ResourceOption) (*SslClient, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SslVpnClientName == nil {
		return nil, errors.New("invalid value for required argument 'SslVpnClientName'")
	}
	if args.SslVpnServerId == nil {
		return nil, errors.New("invalid value for required argument 'SslVpnServerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SslClient
	err := ctx.RegisterResource("tencentcloud:Vpn/sslClient:SslClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslClient gets an existing SslClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslClientState, opts ...pulumi.ResourceOption) (*SslClient, error) {
	var resource SslClient
	err := ctx.ReadResource("tencentcloud:Vpn/sslClient:SslClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslClient resources.
type sslClientState struct {
	// The name of ssl vpn client to be created.
	SslVpnClientName *string `pulumi:"sslVpnClientName"`
	// VPN ssl server id.
	SslVpnServerId *string `pulumi:"sslVpnServerId"`
}

type SslClientState struct {
	// The name of ssl vpn client to be created.
	SslVpnClientName pulumi.StringPtrInput
	// VPN ssl server id.
	SslVpnServerId pulumi.StringPtrInput
}

func (SslClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslClientState)(nil)).Elem()
}

type sslClientArgs struct {
	// The name of ssl vpn client to be created.
	SslVpnClientName string `pulumi:"sslVpnClientName"`
	// VPN ssl server id.
	SslVpnServerId string `pulumi:"sslVpnServerId"`
}

// The set of arguments for constructing a SslClient resource.
type SslClientArgs struct {
	// The name of ssl vpn client to be created.
	SslVpnClientName pulumi.StringInput
	// VPN ssl server id.
	SslVpnServerId pulumi.StringInput
}

func (SslClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslClientArgs)(nil)).Elem()
}

type SslClientInput interface {
	pulumi.Input

	ToSslClientOutput() SslClientOutput
	ToSslClientOutputWithContext(ctx context.Context) SslClientOutput
}

func (*SslClient) ElementType() reflect.Type {
	return reflect.TypeOf((**SslClient)(nil)).Elem()
}

func (i *SslClient) ToSslClientOutput() SslClientOutput {
	return i.ToSslClientOutputWithContext(context.Background())
}

func (i *SslClient) ToSslClientOutputWithContext(ctx context.Context) SslClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslClientOutput)
}

// SslClientArrayInput is an input type that accepts SslClientArray and SslClientArrayOutput values.
// You can construct a concrete instance of `SslClientArrayInput` via:
//
//	SslClientArray{ SslClientArgs{...} }
type SslClientArrayInput interface {
	pulumi.Input

	ToSslClientArrayOutput() SslClientArrayOutput
	ToSslClientArrayOutputWithContext(context.Context) SslClientArrayOutput
}

type SslClientArray []SslClientInput

func (SslClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslClient)(nil)).Elem()
}

func (i SslClientArray) ToSslClientArrayOutput() SslClientArrayOutput {
	return i.ToSslClientArrayOutputWithContext(context.Background())
}

func (i SslClientArray) ToSslClientArrayOutputWithContext(ctx context.Context) SslClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslClientArrayOutput)
}

// SslClientMapInput is an input type that accepts SslClientMap and SslClientMapOutput values.
// You can construct a concrete instance of `SslClientMapInput` via:
//
//	SslClientMap{ "key": SslClientArgs{...} }
type SslClientMapInput interface {
	pulumi.Input

	ToSslClientMapOutput() SslClientMapOutput
	ToSslClientMapOutputWithContext(context.Context) SslClientMapOutput
}

type SslClientMap map[string]SslClientInput

func (SslClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslClient)(nil)).Elem()
}

func (i SslClientMap) ToSslClientMapOutput() SslClientMapOutput {
	return i.ToSslClientMapOutputWithContext(context.Background())
}

func (i SslClientMap) ToSslClientMapOutputWithContext(ctx context.Context) SslClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SslClientMapOutput)
}

type SslClientOutput struct{ *pulumi.OutputState }

func (SslClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslClient)(nil)).Elem()
}

func (o SslClientOutput) ToSslClientOutput() SslClientOutput {
	return o
}

func (o SslClientOutput) ToSslClientOutputWithContext(ctx context.Context) SslClientOutput {
	return o
}

// The name of ssl vpn client to be created.
func (o SslClientOutput) SslVpnClientName() pulumi.StringOutput {
	return o.ApplyT(func(v *SslClient) pulumi.StringOutput { return v.SslVpnClientName }).(pulumi.StringOutput)
}

// VPN ssl server id.
func (o SslClientOutput) SslVpnServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *SslClient) pulumi.StringOutput { return v.SslVpnServerId }).(pulumi.StringOutput)
}

type SslClientArrayOutput struct{ *pulumi.OutputState }

func (SslClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SslClient)(nil)).Elem()
}

func (o SslClientArrayOutput) ToSslClientArrayOutput() SslClientArrayOutput {
	return o
}

func (o SslClientArrayOutput) ToSslClientArrayOutputWithContext(ctx context.Context) SslClientArrayOutput {
	return o
}

func (o SslClientArrayOutput) Index(i pulumi.IntInput) SslClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SslClient {
		return vs[0].([]*SslClient)[vs[1].(int)]
	}).(SslClientOutput)
}

type SslClientMapOutput struct{ *pulumi.OutputState }

func (SslClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SslClient)(nil)).Elem()
}

func (o SslClientMapOutput) ToSslClientMapOutput() SslClientMapOutput {
	return o
}

func (o SslClientMapOutput) ToSslClientMapOutputWithContext(ctx context.Context) SslClientMapOutput {
	return o
}

func (o SslClientMapOutput) MapIndex(k pulumi.StringInput) SslClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SslClient {
		return vs[0].(map[string]*SslClient)[vs[1].(string)]
	}).(SslClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SslClientInput)(nil)).Elem(), &SslClient{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslClientArrayInput)(nil)).Elem(), SslClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SslClientMapInput)(nil)).Elem(), SslClientMap{})
	pulumi.RegisterOutputType(SslClientOutput{})
	pulumi.RegisterOutputType(SslClientArrayOutput{})
	pulumi.RegisterOutputType(SslClientMapOutput{})
}
