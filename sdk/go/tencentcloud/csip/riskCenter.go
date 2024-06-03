// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package csip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a csip riskCenter
//
// ## Example Usage
//
// ### If taskMode is 0
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Csip"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Csip.NewRiskCenter(ctx, "example", &Csip.RiskCenterArgs{
//				Assets: csip.RiskCenterAssetArray{
//					&csip.RiskCenterAssetArgs{
//						Asset:        pulumi.String("49.232.172.248"),
//						AssetName:    pulumi.String("iac-test"),
//						AssetType:    pulumi.String("PublicIp"),
//						InstanceType: pulumi.String("1"),
//						Region:       pulumi.String("ap-beijing"),
//					},
//				},
//				ScanAssetType: pulumi.Int(2),
//				ScanItems: pulumi.StringArray{
//					pulumi.String("port"),
//					pulumi.String("poc"),
//					pulumi.String("weakpass"),
//				},
//				ScanPlanContent: pulumi.String("0 0 0 */1 * * *"),
//				ScanPlanType:    pulumi.Int(0),
//				TaskMode:        pulumi.Int(0),
//				TaskName:        pulumi.String("tf_example"),
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
// ### If taskMode is 1
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Csip"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Csip.NewRiskCenter(ctx, "example", &Csip.RiskCenterArgs{
//				ScanAssetType: pulumi.Int(1),
//				ScanItems: pulumi.StringArray{
//					pulumi.String("port"),
//					pulumi.String("poc"),
//				},
//				ScanPlanType: pulumi.Int(1),
//				TaskMode:     pulumi.Int(1),
//				TaskName:     pulumi.String("tf_example"),
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
// ### If taskMode is 2
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Csip"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Csip.NewRiskCenter(ctx, "example", &Csip.RiskCenterArgs{
//				Assets: csip.RiskCenterAssetArray{
//					&csip.RiskCenterAssetArgs{
//						Asset:        pulumi.String("ins-9p3dkkwy"),
//						AssetName:    pulumi.String("sub machine of tke"),
//						AssetType:    pulumi.String("CVM"),
//						InstanceType: pulumi.String("Instance"),
//						Region:       pulumi.String("ap-guangzhou"),
//					},
//				},
//				ScanAssetType: pulumi.Int(2),
//				ScanItems: pulumi.StringArray{
//					pulumi.String("port"),
//					pulumi.String("configrisk"),
//					pulumi.String("poc"),
//					pulumi.String("weakpass"),
//				},
//				ScanPlanContent: pulumi.String("0 0 0 20 3 * 2024"),
//				ScanPlanType:    pulumi.Int(2),
//				TaskAdvanceCfg: &csip.RiskCenterTaskAdvanceCfgArgs{
//					CfgRisks: csip.RiskCenterTaskAdvanceCfgCfgRiskArray{
//						&csip.RiskCenterTaskAdvanceCfgCfgRiskArgs{
//							Enable:       pulumi.Int(1),
//							ItemId:       pulumi.String("02c9337f-a6da-49b4-8858-64663a02b79f"),
//							ResourceType: pulumi.String("cdb;rds"),
//						},
//					},
//					PortRisks: csip.RiskCenterTaskAdvanceCfgPortRiskArray{
//						&csip.RiskCenterTaskAdvanceCfgPortRiskArgs{
//							CheckType: pulumi.Int(0),
//							Detail:    pulumi.String("22、8080、80、443、3380、3389常见流量端"),
//							Enable:    pulumi.Int(1),
//							PortSets:  pulumi.String("常见端口"),
//						},
//					},
//					VulRisks: csip.RiskCenterTaskAdvanceCfgVulRiskArray{
//						&csip.RiskCenterTaskAdvanceCfgVulRiskArgs{
//							Enable: pulumi.Int(1),
//							RiskId: pulumi.String("f79e371ce5f644f0fdc72a143144c4b2"),
//						},
//					},
//					WeakPwdRisks: csip.RiskCenterTaskAdvanceCfgWeakPwdRiskArray{
//						&csip.RiskCenterTaskAdvanceCfgWeakPwdRiskArgs{
//							CheckItemId: pulumi.Int(50),
//							Enable:      pulumi.Int(1),
//						},
//					},
//				},
//				TaskMode: pulumi.Int(2),
//				TaskName: pulumi.String("tf_example"),
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
type RiskCenter struct {
	pulumi.CustomResourceState

	// Scan the asset information list.
	Assets RiskCenterAssetArrayOutput `pulumi:"assets"`
	// 0- Full scan, 1- Specify asset scan, 2- Exclude asset scan, 3- Manually fill in the scan. If 1 and 2 are required while taskMode not 1, the Assets field is required. If 3 is required, SelfDefiningAssets is required.
	ScanAssetType pulumi.IntOutput `pulumi:"scanAssetType"`
	// Request origin.
	ScanFrom pulumi.StringOutput `pulumi:"scanFrom"`
	// Scan Project. Example: port/poc/weakpass/webcontent/configrisk/exposedserver.
	ScanItems pulumi.StringArrayOutput `pulumi:"scanItems"`
	// Scan plan details.
	ScanPlanContent pulumi.StringPtrOutput `pulumi:"scanPlanContent"`
	// 0- Periodic task, 1- immediate scan, 2- periodic scan, 3- Custom; 0, 2 and 3 are required for scan_plan_content.
	ScanPlanType pulumi.IntOutput `pulumi:"scanPlanType"`
	// Ip/domain/url array.
	SelfDefiningAssets pulumi.StringArrayOutput `pulumi:"selfDefiningAssets"`
	// Advanced configuration.
	TaskAdvanceCfg RiskCenterTaskAdvanceCfgPtrOutput `pulumi:"taskAdvanceCfg"`
	// Physical examination mode, 0-standard mode, 1-fast mode, 2-advanced mode, default standard mode.
	TaskMode pulumi.IntPtrOutput `pulumi:"taskMode"`
	// Task Name.
	TaskName pulumi.StringOutput `pulumi:"taskName"`
}

// NewRiskCenter registers a new resource with the given unique name, arguments, and options.
func NewRiskCenter(ctx *pulumi.Context,
	name string, args *RiskCenterArgs, opts ...pulumi.ResourceOption) (*RiskCenter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScanAssetType == nil {
		return nil, errors.New("invalid value for required argument 'ScanAssetType'")
	}
	if args.ScanItems == nil {
		return nil, errors.New("invalid value for required argument 'ScanItems'")
	}
	if args.ScanPlanType == nil {
		return nil, errors.New("invalid value for required argument 'ScanPlanType'")
	}
	if args.TaskName == nil {
		return nil, errors.New("invalid value for required argument 'TaskName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RiskCenter
	err := ctx.RegisterResource("tencentcloud:Csip/riskCenter:RiskCenter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRiskCenter gets an existing RiskCenter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRiskCenter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RiskCenterState, opts ...pulumi.ResourceOption) (*RiskCenter, error) {
	var resource RiskCenter
	err := ctx.ReadResource("tencentcloud:Csip/riskCenter:RiskCenter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RiskCenter resources.
type riskCenterState struct {
	// Scan the asset information list.
	Assets []RiskCenterAsset `pulumi:"assets"`
	// 0- Full scan, 1- Specify asset scan, 2- Exclude asset scan, 3- Manually fill in the scan. If 1 and 2 are required while taskMode not 1, the Assets field is required. If 3 is required, SelfDefiningAssets is required.
	ScanAssetType *int `pulumi:"scanAssetType"`
	// Request origin.
	ScanFrom *string `pulumi:"scanFrom"`
	// Scan Project. Example: port/poc/weakpass/webcontent/configrisk/exposedserver.
	ScanItems []string `pulumi:"scanItems"`
	// Scan plan details.
	ScanPlanContent *string `pulumi:"scanPlanContent"`
	// 0- Periodic task, 1- immediate scan, 2- periodic scan, 3- Custom; 0, 2 and 3 are required for scan_plan_content.
	ScanPlanType *int `pulumi:"scanPlanType"`
	// Ip/domain/url array.
	SelfDefiningAssets []string `pulumi:"selfDefiningAssets"`
	// Advanced configuration.
	TaskAdvanceCfg *RiskCenterTaskAdvanceCfg `pulumi:"taskAdvanceCfg"`
	// Physical examination mode, 0-standard mode, 1-fast mode, 2-advanced mode, default standard mode.
	TaskMode *int `pulumi:"taskMode"`
	// Task Name.
	TaskName *string `pulumi:"taskName"`
}

type RiskCenterState struct {
	// Scan the asset information list.
	Assets RiskCenterAssetArrayInput
	// 0- Full scan, 1- Specify asset scan, 2- Exclude asset scan, 3- Manually fill in the scan. If 1 and 2 are required while taskMode not 1, the Assets field is required. If 3 is required, SelfDefiningAssets is required.
	ScanAssetType pulumi.IntPtrInput
	// Request origin.
	ScanFrom pulumi.StringPtrInput
	// Scan Project. Example: port/poc/weakpass/webcontent/configrisk/exposedserver.
	ScanItems pulumi.StringArrayInput
	// Scan plan details.
	ScanPlanContent pulumi.StringPtrInput
	// 0- Periodic task, 1- immediate scan, 2- periodic scan, 3- Custom; 0, 2 and 3 are required for scan_plan_content.
	ScanPlanType pulumi.IntPtrInput
	// Ip/domain/url array.
	SelfDefiningAssets pulumi.StringArrayInput
	// Advanced configuration.
	TaskAdvanceCfg RiskCenterTaskAdvanceCfgPtrInput
	// Physical examination mode, 0-standard mode, 1-fast mode, 2-advanced mode, default standard mode.
	TaskMode pulumi.IntPtrInput
	// Task Name.
	TaskName pulumi.StringPtrInput
}

func (RiskCenterState) ElementType() reflect.Type {
	return reflect.TypeOf((*riskCenterState)(nil)).Elem()
}

type riskCenterArgs struct {
	// Scan the asset information list.
	Assets []RiskCenterAsset `pulumi:"assets"`
	// 0- Full scan, 1- Specify asset scan, 2- Exclude asset scan, 3- Manually fill in the scan. If 1 and 2 are required while taskMode not 1, the Assets field is required. If 3 is required, SelfDefiningAssets is required.
	ScanAssetType int `pulumi:"scanAssetType"`
	// Scan Project. Example: port/poc/weakpass/webcontent/configrisk/exposedserver.
	ScanItems []string `pulumi:"scanItems"`
	// Scan plan details.
	ScanPlanContent *string `pulumi:"scanPlanContent"`
	// 0- Periodic task, 1- immediate scan, 2- periodic scan, 3- Custom; 0, 2 and 3 are required for scan_plan_content.
	ScanPlanType int `pulumi:"scanPlanType"`
	// Ip/domain/url array.
	SelfDefiningAssets []string `pulumi:"selfDefiningAssets"`
	// Advanced configuration.
	TaskAdvanceCfg *RiskCenterTaskAdvanceCfg `pulumi:"taskAdvanceCfg"`
	// Physical examination mode, 0-standard mode, 1-fast mode, 2-advanced mode, default standard mode.
	TaskMode *int `pulumi:"taskMode"`
	// Task Name.
	TaskName string `pulumi:"taskName"`
}

// The set of arguments for constructing a RiskCenter resource.
type RiskCenterArgs struct {
	// Scan the asset information list.
	Assets RiskCenterAssetArrayInput
	// 0- Full scan, 1- Specify asset scan, 2- Exclude asset scan, 3- Manually fill in the scan. If 1 and 2 are required while taskMode not 1, the Assets field is required. If 3 is required, SelfDefiningAssets is required.
	ScanAssetType pulumi.IntInput
	// Scan Project. Example: port/poc/weakpass/webcontent/configrisk/exposedserver.
	ScanItems pulumi.StringArrayInput
	// Scan plan details.
	ScanPlanContent pulumi.StringPtrInput
	// 0- Periodic task, 1- immediate scan, 2- periodic scan, 3- Custom; 0, 2 and 3 are required for scan_plan_content.
	ScanPlanType pulumi.IntInput
	// Ip/domain/url array.
	SelfDefiningAssets pulumi.StringArrayInput
	// Advanced configuration.
	TaskAdvanceCfg RiskCenterTaskAdvanceCfgPtrInput
	// Physical examination mode, 0-standard mode, 1-fast mode, 2-advanced mode, default standard mode.
	TaskMode pulumi.IntPtrInput
	// Task Name.
	TaskName pulumi.StringInput
}

func (RiskCenterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*riskCenterArgs)(nil)).Elem()
}

type RiskCenterInput interface {
	pulumi.Input

	ToRiskCenterOutput() RiskCenterOutput
	ToRiskCenterOutputWithContext(ctx context.Context) RiskCenterOutput
}

func (*RiskCenter) ElementType() reflect.Type {
	return reflect.TypeOf((**RiskCenter)(nil)).Elem()
}

func (i *RiskCenter) ToRiskCenterOutput() RiskCenterOutput {
	return i.ToRiskCenterOutputWithContext(context.Background())
}

func (i *RiskCenter) ToRiskCenterOutputWithContext(ctx context.Context) RiskCenterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RiskCenterOutput)
}

// RiskCenterArrayInput is an input type that accepts RiskCenterArray and RiskCenterArrayOutput values.
// You can construct a concrete instance of `RiskCenterArrayInput` via:
//
//	RiskCenterArray{ RiskCenterArgs{...} }
type RiskCenterArrayInput interface {
	pulumi.Input

	ToRiskCenterArrayOutput() RiskCenterArrayOutput
	ToRiskCenterArrayOutputWithContext(context.Context) RiskCenterArrayOutput
}

type RiskCenterArray []RiskCenterInput

func (RiskCenterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RiskCenter)(nil)).Elem()
}

func (i RiskCenterArray) ToRiskCenterArrayOutput() RiskCenterArrayOutput {
	return i.ToRiskCenterArrayOutputWithContext(context.Background())
}

func (i RiskCenterArray) ToRiskCenterArrayOutputWithContext(ctx context.Context) RiskCenterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RiskCenterArrayOutput)
}

// RiskCenterMapInput is an input type that accepts RiskCenterMap and RiskCenterMapOutput values.
// You can construct a concrete instance of `RiskCenterMapInput` via:
//
//	RiskCenterMap{ "key": RiskCenterArgs{...} }
type RiskCenterMapInput interface {
	pulumi.Input

	ToRiskCenterMapOutput() RiskCenterMapOutput
	ToRiskCenterMapOutputWithContext(context.Context) RiskCenterMapOutput
}

type RiskCenterMap map[string]RiskCenterInput

func (RiskCenterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RiskCenter)(nil)).Elem()
}

func (i RiskCenterMap) ToRiskCenterMapOutput() RiskCenterMapOutput {
	return i.ToRiskCenterMapOutputWithContext(context.Background())
}

func (i RiskCenterMap) ToRiskCenterMapOutputWithContext(ctx context.Context) RiskCenterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RiskCenterMapOutput)
}

type RiskCenterOutput struct{ *pulumi.OutputState }

func (RiskCenterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RiskCenter)(nil)).Elem()
}

func (o RiskCenterOutput) ToRiskCenterOutput() RiskCenterOutput {
	return o
}

func (o RiskCenterOutput) ToRiskCenterOutputWithContext(ctx context.Context) RiskCenterOutput {
	return o
}

// Scan the asset information list.
func (o RiskCenterOutput) Assets() RiskCenterAssetArrayOutput {
	return o.ApplyT(func(v *RiskCenter) RiskCenterAssetArrayOutput { return v.Assets }).(RiskCenterAssetArrayOutput)
}

// 0- Full scan, 1- Specify asset scan, 2- Exclude asset scan, 3- Manually fill in the scan. If 1 and 2 are required while taskMode not 1, the Assets field is required. If 3 is required, SelfDefiningAssets is required.
func (o RiskCenterOutput) ScanAssetType() pulumi.IntOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.IntOutput { return v.ScanAssetType }).(pulumi.IntOutput)
}

// Request origin.
func (o RiskCenterOutput) ScanFrom() pulumi.StringOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.StringOutput { return v.ScanFrom }).(pulumi.StringOutput)
}

// Scan Project. Example: port/poc/weakpass/webcontent/configrisk/exposedserver.
func (o RiskCenterOutput) ScanItems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.StringArrayOutput { return v.ScanItems }).(pulumi.StringArrayOutput)
}

// Scan plan details.
func (o RiskCenterOutput) ScanPlanContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.StringPtrOutput { return v.ScanPlanContent }).(pulumi.StringPtrOutput)
}

// 0- Periodic task, 1- immediate scan, 2- periodic scan, 3- Custom; 0, 2 and 3 are required for scan_plan_content.
func (o RiskCenterOutput) ScanPlanType() pulumi.IntOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.IntOutput { return v.ScanPlanType }).(pulumi.IntOutput)
}

// Ip/domain/url array.
func (o RiskCenterOutput) SelfDefiningAssets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.StringArrayOutput { return v.SelfDefiningAssets }).(pulumi.StringArrayOutput)
}

// Advanced configuration.
func (o RiskCenterOutput) TaskAdvanceCfg() RiskCenterTaskAdvanceCfgPtrOutput {
	return o.ApplyT(func(v *RiskCenter) RiskCenterTaskAdvanceCfgPtrOutput { return v.TaskAdvanceCfg }).(RiskCenterTaskAdvanceCfgPtrOutput)
}

// Physical examination mode, 0-standard mode, 1-fast mode, 2-advanced mode, default standard mode.
func (o RiskCenterOutput) TaskMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.IntPtrOutput { return v.TaskMode }).(pulumi.IntPtrOutput)
}

// Task Name.
func (o RiskCenterOutput) TaskName() pulumi.StringOutput {
	return o.ApplyT(func(v *RiskCenter) pulumi.StringOutput { return v.TaskName }).(pulumi.StringOutput)
}

type RiskCenterArrayOutput struct{ *pulumi.OutputState }

func (RiskCenterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RiskCenter)(nil)).Elem()
}

func (o RiskCenterArrayOutput) ToRiskCenterArrayOutput() RiskCenterArrayOutput {
	return o
}

func (o RiskCenterArrayOutput) ToRiskCenterArrayOutputWithContext(ctx context.Context) RiskCenterArrayOutput {
	return o
}

func (o RiskCenterArrayOutput) Index(i pulumi.IntInput) RiskCenterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RiskCenter {
		return vs[0].([]*RiskCenter)[vs[1].(int)]
	}).(RiskCenterOutput)
}

type RiskCenterMapOutput struct{ *pulumi.OutputState }

func (RiskCenterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RiskCenter)(nil)).Elem()
}

func (o RiskCenterMapOutput) ToRiskCenterMapOutput() RiskCenterMapOutput {
	return o
}

func (o RiskCenterMapOutput) ToRiskCenterMapOutputWithContext(ctx context.Context) RiskCenterMapOutput {
	return o
}

func (o RiskCenterMapOutput) MapIndex(k pulumi.StringInput) RiskCenterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RiskCenter {
		return vs[0].(map[string]*RiskCenter)[vs[1].(string)]
	}).(RiskCenterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RiskCenterInput)(nil)).Elem(), &RiskCenter{})
	pulumi.RegisterInputType(reflect.TypeOf((*RiskCenterArrayInput)(nil)).Elem(), RiskCenterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RiskCenterMapInput)(nil)).Elem(), RiskCenterMap{})
	pulumi.RegisterOutputType(RiskCenterOutput{})
	pulumi.RegisterOutputType(RiskCenterArrayOutput{})
	pulumi.RegisterOutputType(RiskCenterMapOutput{})
}
