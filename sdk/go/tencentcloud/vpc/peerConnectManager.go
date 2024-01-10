// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc peerConnectManager
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/User"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/User"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			info, err := User.GetInfo(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			ownerUin := info.OwnerUin
//			vpc, err := Vpc.NewInstance(ctx, "vpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("10.0.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			desVpc, err := Vpc.NewInstance(ctx, "desVpc", &Vpc.InstanceArgs{
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Vpc.NewPeerConnectManager(ctx, "peerConnectManager", &Vpc.PeerConnectManagerArgs{
//				SourceVpcId:           vpc.ID(),
//				PeeringConnectionName: pulumi.String("example-iac"),
//				DestinationVpcId:      desVpc.ID(),
//				DestinationUin:        pulumi.String(ownerUin),
//				DestinationRegion:     pulumi.String("ap-guangzhou"),
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
// vpc peer_connect_manager can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Vpc/peerConnectManager:PeerConnectManager peer_connect_manager peer_connect_manager_id
//
// ```
type PeerConnectManager struct {
	pulumi.CustomResourceState

	// Bandwidth upper limit, unit Mbps.
	Bandwidth pulumi.IntPtrOutput `pulumi:"bandwidth"`
	// Billing mode, daily peak value POSTPAID_BY_DAY_MAX, monthly value 95 POSTPAID_BY_MONTH_95.
	ChargeType pulumi.StringOutput `pulumi:"chargeType"`
	// Peer region.
	DestinationRegion pulumi.StringOutput `pulumi:"destinationRegion"`
	// Peer user UIN.
	DestinationUin pulumi.StringOutput `pulumi:"destinationUin"`
	// The unique ID of the peer VPC.
	DestinationVpcId pulumi.StringOutput `pulumi:"destinationVpcId"`
	// Peer connection name.
	PeeringConnectionName pulumi.StringOutput `pulumi:"peeringConnectionName"`
	// Service classification PT, AU, AG.
	QosLevel pulumi.StringOutput `pulumi:"qosLevel"`
	// The unique ID of the local VPC.
	SourceVpcId pulumi.StringOutput `pulumi:"sourceVpcId"`
	// Interworking type, VPC_PEER interworking between VPCs; VPC_BM_PEER interworking between VPC and BM Network.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPeerConnectManager registers a new resource with the given unique name, arguments, and options.
func NewPeerConnectManager(ctx *pulumi.Context,
	name string, args *PeerConnectManagerArgs, opts ...pulumi.ResourceOption) (*PeerConnectManager, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationRegion == nil {
		return nil, errors.New("invalid value for required argument 'DestinationRegion'")
	}
	if args.DestinationUin == nil {
		return nil, errors.New("invalid value for required argument 'DestinationUin'")
	}
	if args.DestinationVpcId == nil {
		return nil, errors.New("invalid value for required argument 'DestinationVpcId'")
	}
	if args.PeeringConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringConnectionName'")
	}
	if args.SourceVpcId == nil {
		return nil, errors.New("invalid value for required argument 'SourceVpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PeerConnectManager
	err := ctx.RegisterResource("tencentcloud:Vpc/peerConnectManager:PeerConnectManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeerConnectManager gets an existing PeerConnectManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeerConnectManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeerConnectManagerState, opts ...pulumi.ResourceOption) (*PeerConnectManager, error) {
	var resource PeerConnectManager
	err := ctx.ReadResource("tencentcloud:Vpc/peerConnectManager:PeerConnectManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeerConnectManager resources.
type peerConnectManagerState struct {
	// Bandwidth upper limit, unit Mbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// Billing mode, daily peak value POSTPAID_BY_DAY_MAX, monthly value 95 POSTPAID_BY_MONTH_95.
	ChargeType *string `pulumi:"chargeType"`
	// Peer region.
	DestinationRegion *string `pulumi:"destinationRegion"`
	// Peer user UIN.
	DestinationUin *string `pulumi:"destinationUin"`
	// The unique ID of the peer VPC.
	DestinationVpcId *string `pulumi:"destinationVpcId"`
	// Peer connection name.
	PeeringConnectionName *string `pulumi:"peeringConnectionName"`
	// Service classification PT, AU, AG.
	QosLevel *string `pulumi:"qosLevel"`
	// The unique ID of the local VPC.
	SourceVpcId *string `pulumi:"sourceVpcId"`
	// Interworking type, VPC_PEER interworking between VPCs; VPC_BM_PEER interworking between VPC and BM Network.
	Type *string `pulumi:"type"`
}

type PeerConnectManagerState struct {
	// Bandwidth upper limit, unit Mbps.
	Bandwidth pulumi.IntPtrInput
	// Billing mode, daily peak value POSTPAID_BY_DAY_MAX, monthly value 95 POSTPAID_BY_MONTH_95.
	ChargeType pulumi.StringPtrInput
	// Peer region.
	DestinationRegion pulumi.StringPtrInput
	// Peer user UIN.
	DestinationUin pulumi.StringPtrInput
	// The unique ID of the peer VPC.
	DestinationVpcId pulumi.StringPtrInput
	// Peer connection name.
	PeeringConnectionName pulumi.StringPtrInput
	// Service classification PT, AU, AG.
	QosLevel pulumi.StringPtrInput
	// The unique ID of the local VPC.
	SourceVpcId pulumi.StringPtrInput
	// Interworking type, VPC_PEER interworking between VPCs; VPC_BM_PEER interworking between VPC and BM Network.
	Type pulumi.StringPtrInput
}

func (PeerConnectManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*peerConnectManagerState)(nil)).Elem()
}

type peerConnectManagerArgs struct {
	// Bandwidth upper limit, unit Mbps.
	Bandwidth *int `pulumi:"bandwidth"`
	// Billing mode, daily peak value POSTPAID_BY_DAY_MAX, monthly value 95 POSTPAID_BY_MONTH_95.
	ChargeType *string `pulumi:"chargeType"`
	// Peer region.
	DestinationRegion string `pulumi:"destinationRegion"`
	// Peer user UIN.
	DestinationUin string `pulumi:"destinationUin"`
	// The unique ID of the peer VPC.
	DestinationVpcId string `pulumi:"destinationVpcId"`
	// Peer connection name.
	PeeringConnectionName string `pulumi:"peeringConnectionName"`
	// Service classification PT, AU, AG.
	QosLevel *string `pulumi:"qosLevel"`
	// The unique ID of the local VPC.
	SourceVpcId string `pulumi:"sourceVpcId"`
	// Interworking type, VPC_PEER interworking between VPCs; VPC_BM_PEER interworking between VPC and BM Network.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a PeerConnectManager resource.
type PeerConnectManagerArgs struct {
	// Bandwidth upper limit, unit Mbps.
	Bandwidth pulumi.IntPtrInput
	// Billing mode, daily peak value POSTPAID_BY_DAY_MAX, monthly value 95 POSTPAID_BY_MONTH_95.
	ChargeType pulumi.StringPtrInput
	// Peer region.
	DestinationRegion pulumi.StringInput
	// Peer user UIN.
	DestinationUin pulumi.StringInput
	// The unique ID of the peer VPC.
	DestinationVpcId pulumi.StringInput
	// Peer connection name.
	PeeringConnectionName pulumi.StringInput
	// Service classification PT, AU, AG.
	QosLevel pulumi.StringPtrInput
	// The unique ID of the local VPC.
	SourceVpcId pulumi.StringInput
	// Interworking type, VPC_PEER interworking between VPCs; VPC_BM_PEER interworking between VPC and BM Network.
	Type pulumi.StringPtrInput
}

func (PeerConnectManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peerConnectManagerArgs)(nil)).Elem()
}

type PeerConnectManagerInput interface {
	pulumi.Input

	ToPeerConnectManagerOutput() PeerConnectManagerOutput
	ToPeerConnectManagerOutputWithContext(ctx context.Context) PeerConnectManagerOutput
}

func (*PeerConnectManager) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerConnectManager)(nil)).Elem()
}

func (i *PeerConnectManager) ToPeerConnectManagerOutput() PeerConnectManagerOutput {
	return i.ToPeerConnectManagerOutputWithContext(context.Background())
}

func (i *PeerConnectManager) ToPeerConnectManagerOutputWithContext(ctx context.Context) PeerConnectManagerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerConnectManagerOutput)
}

// PeerConnectManagerArrayInput is an input type that accepts PeerConnectManagerArray and PeerConnectManagerArrayOutput values.
// You can construct a concrete instance of `PeerConnectManagerArrayInput` via:
//
//	PeerConnectManagerArray{ PeerConnectManagerArgs{...} }
type PeerConnectManagerArrayInput interface {
	pulumi.Input

	ToPeerConnectManagerArrayOutput() PeerConnectManagerArrayOutput
	ToPeerConnectManagerArrayOutputWithContext(context.Context) PeerConnectManagerArrayOutput
}

type PeerConnectManagerArray []PeerConnectManagerInput

func (PeerConnectManagerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeerConnectManager)(nil)).Elem()
}

func (i PeerConnectManagerArray) ToPeerConnectManagerArrayOutput() PeerConnectManagerArrayOutput {
	return i.ToPeerConnectManagerArrayOutputWithContext(context.Background())
}

func (i PeerConnectManagerArray) ToPeerConnectManagerArrayOutputWithContext(ctx context.Context) PeerConnectManagerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerConnectManagerArrayOutput)
}

// PeerConnectManagerMapInput is an input type that accepts PeerConnectManagerMap and PeerConnectManagerMapOutput values.
// You can construct a concrete instance of `PeerConnectManagerMapInput` via:
//
//	PeerConnectManagerMap{ "key": PeerConnectManagerArgs{...} }
type PeerConnectManagerMapInput interface {
	pulumi.Input

	ToPeerConnectManagerMapOutput() PeerConnectManagerMapOutput
	ToPeerConnectManagerMapOutputWithContext(context.Context) PeerConnectManagerMapOutput
}

type PeerConnectManagerMap map[string]PeerConnectManagerInput

func (PeerConnectManagerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeerConnectManager)(nil)).Elem()
}

func (i PeerConnectManagerMap) ToPeerConnectManagerMapOutput() PeerConnectManagerMapOutput {
	return i.ToPeerConnectManagerMapOutputWithContext(context.Background())
}

func (i PeerConnectManagerMap) ToPeerConnectManagerMapOutputWithContext(ctx context.Context) PeerConnectManagerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeerConnectManagerMapOutput)
}

type PeerConnectManagerOutput struct{ *pulumi.OutputState }

func (PeerConnectManagerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeerConnectManager)(nil)).Elem()
}

func (o PeerConnectManagerOutput) ToPeerConnectManagerOutput() PeerConnectManagerOutput {
	return o
}

func (o PeerConnectManagerOutput) ToPeerConnectManagerOutputWithContext(ctx context.Context) PeerConnectManagerOutput {
	return o
}

// Bandwidth upper limit, unit Mbps.
func (o PeerConnectManagerOutput) Bandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.IntPtrOutput { return v.Bandwidth }).(pulumi.IntPtrOutput)
}

// Billing mode, daily peak value POSTPAID_BY_DAY_MAX, monthly value 95 POSTPAID_BY_MONTH_95.
func (o PeerConnectManagerOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.ChargeType }).(pulumi.StringOutput)
}

// Peer region.
func (o PeerConnectManagerOutput) DestinationRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.DestinationRegion }).(pulumi.StringOutput)
}

// Peer user UIN.
func (o PeerConnectManagerOutput) DestinationUin() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.DestinationUin }).(pulumi.StringOutput)
}

// The unique ID of the peer VPC.
func (o PeerConnectManagerOutput) DestinationVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.DestinationVpcId }).(pulumi.StringOutput)
}

// Peer connection name.
func (o PeerConnectManagerOutput) PeeringConnectionName() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.PeeringConnectionName }).(pulumi.StringOutput)
}

// Service classification PT, AU, AG.
func (o PeerConnectManagerOutput) QosLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.QosLevel }).(pulumi.StringOutput)
}

// The unique ID of the local VPC.
func (o PeerConnectManagerOutput) SourceVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.SourceVpcId }).(pulumi.StringOutput)
}

// Interworking type, VPC_PEER interworking between VPCs; VPC_BM_PEER interworking between VPC and BM Network.
func (o PeerConnectManagerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PeerConnectManager) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type PeerConnectManagerArrayOutput struct{ *pulumi.OutputState }

func (PeerConnectManagerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PeerConnectManager)(nil)).Elem()
}

func (o PeerConnectManagerArrayOutput) ToPeerConnectManagerArrayOutput() PeerConnectManagerArrayOutput {
	return o
}

func (o PeerConnectManagerArrayOutput) ToPeerConnectManagerArrayOutputWithContext(ctx context.Context) PeerConnectManagerArrayOutput {
	return o
}

func (o PeerConnectManagerArrayOutput) Index(i pulumi.IntInput) PeerConnectManagerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PeerConnectManager {
		return vs[0].([]*PeerConnectManager)[vs[1].(int)]
	}).(PeerConnectManagerOutput)
}

type PeerConnectManagerMapOutput struct{ *pulumi.OutputState }

func (PeerConnectManagerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PeerConnectManager)(nil)).Elem()
}

func (o PeerConnectManagerMapOutput) ToPeerConnectManagerMapOutput() PeerConnectManagerMapOutput {
	return o
}

func (o PeerConnectManagerMapOutput) ToPeerConnectManagerMapOutputWithContext(ctx context.Context) PeerConnectManagerMapOutput {
	return o
}

func (o PeerConnectManagerMapOutput) MapIndex(k pulumi.StringInput) PeerConnectManagerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PeerConnectManager {
		return vs[0].(map[string]*PeerConnectManager)[vs[1].(string)]
	}).(PeerConnectManagerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PeerConnectManagerInput)(nil)).Elem(), &PeerConnectManager{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeerConnectManagerArrayInput)(nil)).Elem(), PeerConnectManagerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PeerConnectManagerMapInput)(nil)).Elem(), PeerConnectManagerMap{})
	pulumi.RegisterOutputType(PeerConnectManagerOutput{})
	pulumi.RegisterOutputType(PeerConnectManagerArrayOutput{})
	pulumi.RegisterOutputType(PeerConnectManagerMapOutput{})
}
