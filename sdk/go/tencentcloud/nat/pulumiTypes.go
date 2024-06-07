// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

var _ = internal.GetEnvOrDefault

type GetDcRouteNatDirectConnectGatewayRouteSet struct {
	// Create time of route.
	CreateTime string `pulumi:"createTime"`
	// IPv4 CIDR of subnet.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// Id of next-hop gateway.
	GatewayId string `pulumi:"gatewayId"`
	// Type of next-hop gateway, valid values: DIRECTCONNECT.
	GatewayType string `pulumi:"gatewayType"`
	// Update time of route.
	UpdateTime string `pulumi:"updateTime"`
}

// GetDcRouteNatDirectConnectGatewayRouteSetInput is an input type that accepts GetDcRouteNatDirectConnectGatewayRouteSetArgs and GetDcRouteNatDirectConnectGatewayRouteSetOutput values.
// You can construct a concrete instance of `GetDcRouteNatDirectConnectGatewayRouteSetInput` via:
//
//	GetDcRouteNatDirectConnectGatewayRouteSetArgs{...}
type GetDcRouteNatDirectConnectGatewayRouteSetInput interface {
	pulumi.Input

	ToGetDcRouteNatDirectConnectGatewayRouteSetOutput() GetDcRouteNatDirectConnectGatewayRouteSetOutput
	ToGetDcRouteNatDirectConnectGatewayRouteSetOutputWithContext(context.Context) GetDcRouteNatDirectConnectGatewayRouteSetOutput
}

type GetDcRouteNatDirectConnectGatewayRouteSetArgs struct {
	// Create time of route.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// IPv4 CIDR of subnet.
	DestinationCidrBlock pulumi.StringInput `pulumi:"destinationCidrBlock"`
	// Id of next-hop gateway.
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// Type of next-hop gateway, valid values: DIRECTCONNECT.
	GatewayType pulumi.StringInput `pulumi:"gatewayType"`
	// Update time of route.
	UpdateTime pulumi.StringInput `pulumi:"updateTime"`
}

func (GetDcRouteNatDirectConnectGatewayRouteSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcRouteNatDirectConnectGatewayRouteSet)(nil)).Elem()
}

func (i GetDcRouteNatDirectConnectGatewayRouteSetArgs) ToGetDcRouteNatDirectConnectGatewayRouteSetOutput() GetDcRouteNatDirectConnectGatewayRouteSetOutput {
	return i.ToGetDcRouteNatDirectConnectGatewayRouteSetOutputWithContext(context.Background())
}

func (i GetDcRouteNatDirectConnectGatewayRouteSetArgs) ToGetDcRouteNatDirectConnectGatewayRouteSetOutputWithContext(ctx context.Context) GetDcRouteNatDirectConnectGatewayRouteSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDcRouteNatDirectConnectGatewayRouteSetOutput)
}

// GetDcRouteNatDirectConnectGatewayRouteSetArrayInput is an input type that accepts GetDcRouteNatDirectConnectGatewayRouteSetArray and GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput values.
// You can construct a concrete instance of `GetDcRouteNatDirectConnectGatewayRouteSetArrayInput` via:
//
//	GetDcRouteNatDirectConnectGatewayRouteSetArray{ GetDcRouteNatDirectConnectGatewayRouteSetArgs{...} }
type GetDcRouteNatDirectConnectGatewayRouteSetArrayInput interface {
	pulumi.Input

	ToGetDcRouteNatDirectConnectGatewayRouteSetArrayOutput() GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput
	ToGetDcRouteNatDirectConnectGatewayRouteSetArrayOutputWithContext(context.Context) GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput
}

type GetDcRouteNatDirectConnectGatewayRouteSetArray []GetDcRouteNatDirectConnectGatewayRouteSetInput

func (GetDcRouteNatDirectConnectGatewayRouteSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDcRouteNatDirectConnectGatewayRouteSet)(nil)).Elem()
}

func (i GetDcRouteNatDirectConnectGatewayRouteSetArray) ToGetDcRouteNatDirectConnectGatewayRouteSetArrayOutput() GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput {
	return i.ToGetDcRouteNatDirectConnectGatewayRouteSetArrayOutputWithContext(context.Background())
}

func (i GetDcRouteNatDirectConnectGatewayRouteSetArray) ToGetDcRouteNatDirectConnectGatewayRouteSetArrayOutputWithContext(ctx context.Context) GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput)
}

type GetDcRouteNatDirectConnectGatewayRouteSetOutput struct{ *pulumi.OutputState }

func (GetDcRouteNatDirectConnectGatewayRouteSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDcRouteNatDirectConnectGatewayRouteSet)(nil)).Elem()
}

func (o GetDcRouteNatDirectConnectGatewayRouteSetOutput) ToGetDcRouteNatDirectConnectGatewayRouteSetOutput() GetDcRouteNatDirectConnectGatewayRouteSetOutput {
	return o
}

func (o GetDcRouteNatDirectConnectGatewayRouteSetOutput) ToGetDcRouteNatDirectConnectGatewayRouteSetOutputWithContext(ctx context.Context) GetDcRouteNatDirectConnectGatewayRouteSetOutput {
	return o
}

// Create time of route.
func (o GetDcRouteNatDirectConnectGatewayRouteSetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcRouteNatDirectConnectGatewayRouteSet) string { return v.CreateTime }).(pulumi.StringOutput)
}

// IPv4 CIDR of subnet.
func (o GetDcRouteNatDirectConnectGatewayRouteSetOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcRouteNatDirectConnectGatewayRouteSet) string { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// Id of next-hop gateway.
func (o GetDcRouteNatDirectConnectGatewayRouteSetOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcRouteNatDirectConnectGatewayRouteSet) string { return v.GatewayId }).(pulumi.StringOutput)
}

// Type of next-hop gateway, valid values: DIRECTCONNECT.
func (o GetDcRouteNatDirectConnectGatewayRouteSetOutput) GatewayType() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcRouteNatDirectConnectGatewayRouteSet) string { return v.GatewayType }).(pulumi.StringOutput)
}

// Update time of route.
func (o GetDcRouteNatDirectConnectGatewayRouteSetOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetDcRouteNatDirectConnectGatewayRouteSet) string { return v.UpdateTime }).(pulumi.StringOutput)
}

type GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput struct{ *pulumi.OutputState }

func (GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDcRouteNatDirectConnectGatewayRouteSet)(nil)).Elem()
}

func (o GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput) ToGetDcRouteNatDirectConnectGatewayRouteSetArrayOutput() GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput {
	return o
}

func (o GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput) ToGetDcRouteNatDirectConnectGatewayRouteSetArrayOutputWithContext(ctx context.Context) GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput {
	return o
}

func (o GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput) Index(i pulumi.IntInput) GetDcRouteNatDirectConnectGatewayRouteSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDcRouteNatDirectConnectGatewayRouteSet {
		return vs[0].([]GetDcRouteNatDirectConnectGatewayRouteSet)[vs[1].(int)]
	}).(GetDcRouteNatDirectConnectGatewayRouteSetOutput)
}

type GetGatewaySnatsSnatList struct {
	// Create time.
	CreateTime string `pulumi:"createTime"`
	// Description.
	Description string `pulumi:"description"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Private IPs of the instance's primary ENI, required when `resourceType` is NETWORKINTERFACE.
	InstancePrivateIpAddr *string `pulumi:"instancePrivateIpAddr"`
	// NAT gateway ID.
	NatGatewayId string `pulumi:"natGatewayId"`
	// Elastic IP address pool.
	PublicIpAddrs []string `pulumi:"publicIpAddrs"`
	// Resource type. Valid values: SUBNET, NETWORKINTERFACE.
	ResourceType string `pulumi:"resourceType"`
	// SNAT rule ID.
	SnatId string `pulumi:"snatId"`
	// The IPv4 CIDR of the subnet, required when `resourceType` is SUBNET.
	SubnetCidrBlock *string `pulumi:"subnetCidrBlock"`
	// Subnet instance ID.
	SubnetId *string `pulumi:"subnetId"`
}

// GetGatewaySnatsSnatListInput is an input type that accepts GetGatewaySnatsSnatListArgs and GetGatewaySnatsSnatListOutput values.
// You can construct a concrete instance of `GetGatewaySnatsSnatListInput` via:
//
//	GetGatewaySnatsSnatListArgs{...}
type GetGatewaySnatsSnatListInput interface {
	pulumi.Input

	ToGetGatewaySnatsSnatListOutput() GetGatewaySnatsSnatListOutput
	ToGetGatewaySnatsSnatListOutputWithContext(context.Context) GetGatewaySnatsSnatListOutput
}

type GetGatewaySnatsSnatListArgs struct {
	// Create time.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// Description.
	Description pulumi.StringInput `pulumi:"description"`
	// Instance ID.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Private IPs of the instance's primary ENI, required when `resourceType` is NETWORKINTERFACE.
	InstancePrivateIpAddr pulumi.StringPtrInput `pulumi:"instancePrivateIpAddr"`
	// NAT gateway ID.
	NatGatewayId pulumi.StringInput `pulumi:"natGatewayId"`
	// Elastic IP address pool.
	PublicIpAddrs pulumi.StringArrayInput `pulumi:"publicIpAddrs"`
	// Resource type. Valid values: SUBNET, NETWORKINTERFACE.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// SNAT rule ID.
	SnatId pulumi.StringInput `pulumi:"snatId"`
	// The IPv4 CIDR of the subnet, required when `resourceType` is SUBNET.
	SubnetCidrBlock pulumi.StringPtrInput `pulumi:"subnetCidrBlock"`
	// Subnet instance ID.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (GetGatewaySnatsSnatListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaySnatsSnatList)(nil)).Elem()
}

func (i GetGatewaySnatsSnatListArgs) ToGetGatewaySnatsSnatListOutput() GetGatewaySnatsSnatListOutput {
	return i.ToGetGatewaySnatsSnatListOutputWithContext(context.Background())
}

func (i GetGatewaySnatsSnatListArgs) ToGetGatewaySnatsSnatListOutputWithContext(ctx context.Context) GetGatewaySnatsSnatListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGatewaySnatsSnatListOutput)
}

// GetGatewaySnatsSnatListArrayInput is an input type that accepts GetGatewaySnatsSnatListArray and GetGatewaySnatsSnatListArrayOutput values.
// You can construct a concrete instance of `GetGatewaySnatsSnatListArrayInput` via:
//
//	GetGatewaySnatsSnatListArray{ GetGatewaySnatsSnatListArgs{...} }
type GetGatewaySnatsSnatListArrayInput interface {
	pulumi.Input

	ToGetGatewaySnatsSnatListArrayOutput() GetGatewaySnatsSnatListArrayOutput
	ToGetGatewaySnatsSnatListArrayOutputWithContext(context.Context) GetGatewaySnatsSnatListArrayOutput
}

type GetGatewaySnatsSnatListArray []GetGatewaySnatsSnatListInput

func (GetGatewaySnatsSnatListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGatewaySnatsSnatList)(nil)).Elem()
}

func (i GetGatewaySnatsSnatListArray) ToGetGatewaySnatsSnatListArrayOutput() GetGatewaySnatsSnatListArrayOutput {
	return i.ToGetGatewaySnatsSnatListArrayOutputWithContext(context.Background())
}

func (i GetGatewaySnatsSnatListArray) ToGetGatewaySnatsSnatListArrayOutputWithContext(ctx context.Context) GetGatewaySnatsSnatListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGatewaySnatsSnatListArrayOutput)
}

type GetGatewaySnatsSnatListOutput struct{ *pulumi.OutputState }

func (GetGatewaySnatsSnatListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaySnatsSnatList)(nil)).Elem()
}

func (o GetGatewaySnatsSnatListOutput) ToGetGatewaySnatsSnatListOutput() GetGatewaySnatsSnatListOutput {
	return o
}

func (o GetGatewaySnatsSnatListOutput) ToGetGatewaySnatsSnatListOutputWithContext(ctx context.Context) GetGatewaySnatsSnatListOutput {
	return o
}

// Create time.
func (o GetGatewaySnatsSnatListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Description.
func (o GetGatewaySnatsSnatListOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) string { return v.Description }).(pulumi.StringOutput)
}

// Instance ID.
func (o GetGatewaySnatsSnatListOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// Private IPs of the instance's primary ENI, required when `resourceType` is NETWORKINTERFACE.
func (o GetGatewaySnatsSnatListOutput) InstancePrivateIpAddr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) *string { return v.InstancePrivateIpAddr }).(pulumi.StringPtrOutput)
}

// NAT gateway ID.
func (o GetGatewaySnatsSnatListOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) string { return v.NatGatewayId }).(pulumi.StringOutput)
}

// Elastic IP address pool.
func (o GetGatewaySnatsSnatListOutput) PublicIpAddrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) []string { return v.PublicIpAddrs }).(pulumi.StringArrayOutput)
}

// Resource type. Valid values: SUBNET, NETWORKINTERFACE.
func (o GetGatewaySnatsSnatListOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) string { return v.ResourceType }).(pulumi.StringOutput)
}

// SNAT rule ID.
func (o GetGatewaySnatsSnatListOutput) SnatId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) string { return v.SnatId }).(pulumi.StringOutput)
}

// The IPv4 CIDR of the subnet, required when `resourceType` is SUBNET.
func (o GetGatewaySnatsSnatListOutput) SubnetCidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) *string { return v.SubnetCidrBlock }).(pulumi.StringPtrOutput)
}

// Subnet instance ID.
func (o GetGatewaySnatsSnatListOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaySnatsSnatList) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type GetGatewaySnatsSnatListArrayOutput struct{ *pulumi.OutputState }

func (GetGatewaySnatsSnatListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGatewaySnatsSnatList)(nil)).Elem()
}

func (o GetGatewaySnatsSnatListArrayOutput) ToGetGatewaySnatsSnatListArrayOutput() GetGatewaySnatsSnatListArrayOutput {
	return o
}

func (o GetGatewaySnatsSnatListArrayOutput) ToGetGatewaySnatsSnatListArrayOutputWithContext(ctx context.Context) GetGatewaySnatsSnatListArrayOutput {
	return o
}

func (o GetGatewaySnatsSnatListArrayOutput) Index(i pulumi.IntInput) GetGatewaySnatsSnatListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGatewaySnatsSnatList {
		return vs[0].([]GetGatewaySnatsSnatList)[vs[1].(int)]
	}).(GetGatewaySnatsSnatListOutput)
}

type GetGatewaysNat struct {
	// EIP IP address set bound to the gateway. The value of at least 1.
	AssignedEipSets []string `pulumi:"assignedEipSets"`
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.
	Bandwidth int `pulumi:"bandwidth"`
	// Create time of the NAT gateway.
	CreateTime string `pulumi:"createTime"`
	// ID of the NAT gateway.
	Id string `pulumi:"id"`
	// The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000.
	MaxConcurrent int `pulumi:"maxConcurrent"`
	// Name of the NAT gateway.
	Name string `pulumi:"name"`
	// State of the NAT gateway.
	State string `pulumi:"state"`
	// The available tags within this NAT gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// ID of the VPC.
	VpcId string `pulumi:"vpcId"`
}

// GetGatewaysNatInput is an input type that accepts GetGatewaysNatArgs and GetGatewaysNatOutput values.
// You can construct a concrete instance of `GetGatewaysNatInput` via:
//
//	GetGatewaysNatArgs{...}
type GetGatewaysNatInput interface {
	pulumi.Input

	ToGetGatewaysNatOutput() GetGatewaysNatOutput
	ToGetGatewaysNatOutputWithContext(context.Context) GetGatewaysNatOutput
}

type GetGatewaysNatArgs struct {
	// EIP IP address set bound to the gateway. The value of at least 1.
	AssignedEipSets pulumi.StringArrayInput `pulumi:"assignedEipSets"`
	// The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.
	Bandwidth pulumi.IntInput `pulumi:"bandwidth"`
	// Create time of the NAT gateway.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// ID of the NAT gateway.
	Id pulumi.StringInput `pulumi:"id"`
	// The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000.
	MaxConcurrent pulumi.IntInput `pulumi:"maxConcurrent"`
	// Name of the NAT gateway.
	Name pulumi.StringInput `pulumi:"name"`
	// State of the NAT gateway.
	State pulumi.StringInput `pulumi:"state"`
	// The available tags within this NAT gateway.
	Tags pulumi.MapInput `pulumi:"tags"`
	// ID of the VPC.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (GetGatewaysNatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaysNat)(nil)).Elem()
}

func (i GetGatewaysNatArgs) ToGetGatewaysNatOutput() GetGatewaysNatOutput {
	return i.ToGetGatewaysNatOutputWithContext(context.Background())
}

func (i GetGatewaysNatArgs) ToGetGatewaysNatOutputWithContext(ctx context.Context) GetGatewaysNatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGatewaysNatOutput)
}

// GetGatewaysNatArrayInput is an input type that accepts GetGatewaysNatArray and GetGatewaysNatArrayOutput values.
// You can construct a concrete instance of `GetGatewaysNatArrayInput` via:
//
//	GetGatewaysNatArray{ GetGatewaysNatArgs{...} }
type GetGatewaysNatArrayInput interface {
	pulumi.Input

	ToGetGatewaysNatArrayOutput() GetGatewaysNatArrayOutput
	ToGetGatewaysNatArrayOutputWithContext(context.Context) GetGatewaysNatArrayOutput
}

type GetGatewaysNatArray []GetGatewaysNatInput

func (GetGatewaysNatArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGatewaysNat)(nil)).Elem()
}

func (i GetGatewaysNatArray) ToGetGatewaysNatArrayOutput() GetGatewaysNatArrayOutput {
	return i.ToGetGatewaysNatArrayOutputWithContext(context.Background())
}

func (i GetGatewaysNatArray) ToGetGatewaysNatArrayOutputWithContext(ctx context.Context) GetGatewaysNatArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetGatewaysNatArrayOutput)
}

type GetGatewaysNatOutput struct{ *pulumi.OutputState }

func (GetGatewaysNatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaysNat)(nil)).Elem()
}

func (o GetGatewaysNatOutput) ToGetGatewaysNatOutput() GetGatewaysNatOutput {
	return o
}

func (o GetGatewaysNatOutput) ToGetGatewaysNatOutputWithContext(ctx context.Context) GetGatewaysNatOutput {
	return o
}

// EIP IP address set bound to the gateway. The value of at least 1.
func (o GetGatewaysNatOutput) AssignedEipSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGatewaysNat) []string { return v.AssignedEipSets }).(pulumi.StringArrayOutput)
}

// The maximum public network output bandwidth of NAT gateway (unit: Mbps), the available values include: 20,50,100,200,500,1000,2000,5000. Default is 100.
func (o GetGatewaysNatOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetGatewaysNat) int { return v.Bandwidth }).(pulumi.IntOutput)
}

// Create time of the NAT gateway.
func (o GetGatewaysNatOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaysNat) string { return v.CreateTime }).(pulumi.StringOutput)
}

// ID of the NAT gateway.
func (o GetGatewaysNatOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaysNat) string { return v.Id }).(pulumi.StringOutput)
}

// The upper limit of concurrent connection of NAT gateway, the available values include: 1000000,3000000,10000000. Default is 1000000.
func (o GetGatewaysNatOutput) MaxConcurrent() pulumi.IntOutput {
	return o.ApplyT(func(v GetGatewaysNat) int { return v.MaxConcurrent }).(pulumi.IntOutput)
}

// Name of the NAT gateway.
func (o GetGatewaysNatOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaysNat) string { return v.Name }).(pulumi.StringOutput)
}

// State of the NAT gateway.
func (o GetGatewaysNatOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaysNat) string { return v.State }).(pulumi.StringOutput)
}

// The available tags within this NAT gateway.
func (o GetGatewaysNatOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v GetGatewaysNat) map[string]interface{} { return v.Tags }).(pulumi.MapOutput)
}

// ID of the VPC.
func (o GetGatewaysNatOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaysNat) string { return v.VpcId }).(pulumi.StringOutput)
}

type GetGatewaysNatArrayOutput struct{ *pulumi.OutputState }

func (GetGatewaysNatArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetGatewaysNat)(nil)).Elem()
}

func (o GetGatewaysNatArrayOutput) ToGetGatewaysNatArrayOutput() GetGatewaysNatArrayOutput {
	return o
}

func (o GetGatewaysNatArrayOutput) ToGetGatewaysNatArrayOutputWithContext(ctx context.Context) GetGatewaysNatArrayOutput {
	return o
}

func (o GetGatewaysNatArrayOutput) Index(i pulumi.IntInput) GetGatewaysNatOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetGatewaysNat {
		return vs[0].([]GetGatewaysNat)[vs[1].(int)]
	}).(GetGatewaysNatOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetDcRouteNatDirectConnectGatewayRouteSetInput)(nil)).Elem(), GetDcRouteNatDirectConnectGatewayRouteSetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDcRouteNatDirectConnectGatewayRouteSetArrayInput)(nil)).Elem(), GetDcRouteNatDirectConnectGatewayRouteSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGatewaySnatsSnatListInput)(nil)).Elem(), GetGatewaySnatsSnatListArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGatewaySnatsSnatListArrayInput)(nil)).Elem(), GetGatewaySnatsSnatListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGatewaysNatInput)(nil)).Elem(), GetGatewaysNatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetGatewaysNatArrayInput)(nil)).Elem(), GetGatewaysNatArray{})
	pulumi.RegisterOutputType(GetDcRouteNatDirectConnectGatewayRouteSetOutput{})
	pulumi.RegisterOutputType(GetDcRouteNatDirectConnectGatewayRouteSetArrayOutput{})
	pulumi.RegisterOutputType(GetGatewaySnatsSnatListOutput{})
	pulumi.RegisterOutputType(GetGatewaySnatsSnatListArrayOutput{})
	pulumi.RegisterOutputType(GetGatewaysNatOutput{})
	pulumi.RegisterOutputType(GetGatewaysNatArrayOutput{})
}
