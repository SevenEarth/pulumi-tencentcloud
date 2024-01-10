// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cls.NewCosRecharge(ctx, "cosRecharge", &Cls.CosRechargeArgs{
//				Bucket:       pulumi.String("cos-lock-1308919341"),
//				BucketRegion: pulumi.String("ap-guangzhou"),
//				ExtractRuleInfo: &cls.CosRechargeExtractRuleInfoArgs{
//					Backtracking: pulumi.Int(0),
//					FilterKeyRegexes: cls.CosRechargeExtractRuleInfoFilterKeyRegexArray{
//						&cls.CosRechargeExtractRuleInfoFilterKeyRegexArgs{
//							Key:   pulumi.String("__CONTENT__"),
//							Regex: pulumi.String("dasd"),
//						},
//					},
//					IsGbk:               pulumi.Int(0),
//					JsonStandard:        pulumi.Int(0),
//					Keys:                pulumi.StringArray{},
//					MetadataType:        pulumi.Int(0),
//					UnMatchUpLoadSwitch: pulumi.Bool(false),
//				},
//				LogType:  pulumi.String("minimalist_log"),
//				LogsetId: pulumi.String("dd426d1a-95bc-4bca-b8c2-baa169261812"),
//				Prefix:   pulumi.String("test"),
//				TopicId:  pulumi.String("7e34a3a7-635e-4da8-9005-88106c1fde69"),
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
// cls cos_recharge can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cls/cosRecharge:CosRecharge cos_recharge topic_id#cos_recharge_id
//
// ```
type CosRecharge struct {
	pulumi.CustomResourceState

	// cos bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// cos bucket region.
	BucketRegion pulumi.StringOutput `pulumi:"bucketRegion"`
	// supported gzip, lzop, snappy.
	Compress pulumi.StringPtrOutput `pulumi:"compress"`
	// extract rule info.
	ExtractRuleInfo CosRechargeExtractRuleInfoOutput `pulumi:"extractRuleInfo"`
	// log type.
	LogType pulumi.StringOutput `pulumi:"logType"`
	// logset id.
	LogsetId pulumi.StringOutput `pulumi:"logsetId"`
	// recharge name.
	Name pulumi.StringOutput `pulumi:"name"`
	// cos file prefix.
	Prefix pulumi.StringOutput `pulumi:"prefix"`
	// topic id.
	TopicId pulumi.StringOutput `pulumi:"topicId"`
}

// NewCosRecharge registers a new resource with the given unique name, arguments, and options.
func NewCosRecharge(ctx *pulumi.Context,
	name string, args *CosRechargeArgs, opts ...pulumi.ResourceOption) (*CosRecharge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.BucketRegion == nil {
		return nil, errors.New("invalid value for required argument 'BucketRegion'")
	}
	if args.LogType == nil {
		return nil, errors.New("invalid value for required argument 'LogType'")
	}
	if args.LogsetId == nil {
		return nil, errors.New("invalid value for required argument 'LogsetId'")
	}
	if args.Prefix == nil {
		return nil, errors.New("invalid value for required argument 'Prefix'")
	}
	if args.TopicId == nil {
		return nil, errors.New("invalid value for required argument 'TopicId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CosRecharge
	err := ctx.RegisterResource("tencentcloud:Cls/cosRecharge:CosRecharge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCosRecharge gets an existing CosRecharge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCosRecharge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CosRechargeState, opts ...pulumi.ResourceOption) (*CosRecharge, error) {
	var resource CosRecharge
	err := ctx.ReadResource("tencentcloud:Cls/cosRecharge:CosRecharge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CosRecharge resources.
type cosRechargeState struct {
	// cos bucket.
	Bucket *string `pulumi:"bucket"`
	// cos bucket region.
	BucketRegion *string `pulumi:"bucketRegion"`
	// supported gzip, lzop, snappy.
	Compress *string `pulumi:"compress"`
	// extract rule info.
	ExtractRuleInfo *CosRechargeExtractRuleInfo `pulumi:"extractRuleInfo"`
	// log type.
	LogType *string `pulumi:"logType"`
	// logset id.
	LogsetId *string `pulumi:"logsetId"`
	// recharge name.
	Name *string `pulumi:"name"`
	// cos file prefix.
	Prefix *string `pulumi:"prefix"`
	// topic id.
	TopicId *string `pulumi:"topicId"`
}

type CosRechargeState struct {
	// cos bucket.
	Bucket pulumi.StringPtrInput
	// cos bucket region.
	BucketRegion pulumi.StringPtrInput
	// supported gzip, lzop, snappy.
	Compress pulumi.StringPtrInput
	// extract rule info.
	ExtractRuleInfo CosRechargeExtractRuleInfoPtrInput
	// log type.
	LogType pulumi.StringPtrInput
	// logset id.
	LogsetId pulumi.StringPtrInput
	// recharge name.
	Name pulumi.StringPtrInput
	// cos file prefix.
	Prefix pulumi.StringPtrInput
	// topic id.
	TopicId pulumi.StringPtrInput
}

func (CosRechargeState) ElementType() reflect.Type {
	return reflect.TypeOf((*cosRechargeState)(nil)).Elem()
}

type cosRechargeArgs struct {
	// cos bucket.
	Bucket string `pulumi:"bucket"`
	// cos bucket region.
	BucketRegion string `pulumi:"bucketRegion"`
	// supported gzip, lzop, snappy.
	Compress *string `pulumi:"compress"`
	// extract rule info.
	ExtractRuleInfo *CosRechargeExtractRuleInfo `pulumi:"extractRuleInfo"`
	// log type.
	LogType string `pulumi:"logType"`
	// logset id.
	LogsetId string `pulumi:"logsetId"`
	// recharge name.
	Name *string `pulumi:"name"`
	// cos file prefix.
	Prefix string `pulumi:"prefix"`
	// topic id.
	TopicId string `pulumi:"topicId"`
}

// The set of arguments for constructing a CosRecharge resource.
type CosRechargeArgs struct {
	// cos bucket.
	Bucket pulumi.StringInput
	// cos bucket region.
	BucketRegion pulumi.StringInput
	// supported gzip, lzop, snappy.
	Compress pulumi.StringPtrInput
	// extract rule info.
	ExtractRuleInfo CosRechargeExtractRuleInfoPtrInput
	// log type.
	LogType pulumi.StringInput
	// logset id.
	LogsetId pulumi.StringInput
	// recharge name.
	Name pulumi.StringPtrInput
	// cos file prefix.
	Prefix pulumi.StringInput
	// topic id.
	TopicId pulumi.StringInput
}

func (CosRechargeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cosRechargeArgs)(nil)).Elem()
}

type CosRechargeInput interface {
	pulumi.Input

	ToCosRechargeOutput() CosRechargeOutput
	ToCosRechargeOutputWithContext(ctx context.Context) CosRechargeOutput
}

func (*CosRecharge) ElementType() reflect.Type {
	return reflect.TypeOf((**CosRecharge)(nil)).Elem()
}

func (i *CosRecharge) ToCosRechargeOutput() CosRechargeOutput {
	return i.ToCosRechargeOutputWithContext(context.Background())
}

func (i *CosRecharge) ToCosRechargeOutputWithContext(ctx context.Context) CosRechargeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosRechargeOutput)
}

// CosRechargeArrayInput is an input type that accepts CosRechargeArray and CosRechargeArrayOutput values.
// You can construct a concrete instance of `CosRechargeArrayInput` via:
//
//	CosRechargeArray{ CosRechargeArgs{...} }
type CosRechargeArrayInput interface {
	pulumi.Input

	ToCosRechargeArrayOutput() CosRechargeArrayOutput
	ToCosRechargeArrayOutputWithContext(context.Context) CosRechargeArrayOutput
}

type CosRechargeArray []CosRechargeInput

func (CosRechargeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CosRecharge)(nil)).Elem()
}

func (i CosRechargeArray) ToCosRechargeArrayOutput() CosRechargeArrayOutput {
	return i.ToCosRechargeArrayOutputWithContext(context.Background())
}

func (i CosRechargeArray) ToCosRechargeArrayOutputWithContext(ctx context.Context) CosRechargeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosRechargeArrayOutput)
}

// CosRechargeMapInput is an input type that accepts CosRechargeMap and CosRechargeMapOutput values.
// You can construct a concrete instance of `CosRechargeMapInput` via:
//
//	CosRechargeMap{ "key": CosRechargeArgs{...} }
type CosRechargeMapInput interface {
	pulumi.Input

	ToCosRechargeMapOutput() CosRechargeMapOutput
	ToCosRechargeMapOutputWithContext(context.Context) CosRechargeMapOutput
}

type CosRechargeMap map[string]CosRechargeInput

func (CosRechargeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CosRecharge)(nil)).Elem()
}

func (i CosRechargeMap) ToCosRechargeMapOutput() CosRechargeMapOutput {
	return i.ToCosRechargeMapOutputWithContext(context.Background())
}

func (i CosRechargeMap) ToCosRechargeMapOutputWithContext(ctx context.Context) CosRechargeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CosRechargeMapOutput)
}

type CosRechargeOutput struct{ *pulumi.OutputState }

func (CosRechargeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CosRecharge)(nil)).Elem()
}

func (o CosRechargeOutput) ToCosRechargeOutput() CosRechargeOutput {
	return o
}

func (o CosRechargeOutput) ToCosRechargeOutputWithContext(ctx context.Context) CosRechargeOutput {
	return o
}

// cos bucket.
func (o CosRechargeOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// cos bucket region.
func (o CosRechargeOutput) BucketRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringOutput { return v.BucketRegion }).(pulumi.StringOutput)
}

// supported gzip, lzop, snappy.
func (o CosRechargeOutput) Compress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringPtrOutput { return v.Compress }).(pulumi.StringPtrOutput)
}

// extract rule info.
func (o CosRechargeOutput) ExtractRuleInfo() CosRechargeExtractRuleInfoOutput {
	return o.ApplyT(func(v *CosRecharge) CosRechargeExtractRuleInfoOutput { return v.ExtractRuleInfo }).(CosRechargeExtractRuleInfoOutput)
}

// log type.
func (o CosRechargeOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringOutput { return v.LogType }).(pulumi.StringOutput)
}

// logset id.
func (o CosRechargeOutput) LogsetId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringOutput { return v.LogsetId }).(pulumi.StringOutput)
}

// recharge name.
func (o CosRechargeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// cos file prefix.
func (o CosRechargeOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

// topic id.
func (o CosRechargeOutput) TopicId() pulumi.StringOutput {
	return o.ApplyT(func(v *CosRecharge) pulumi.StringOutput { return v.TopicId }).(pulumi.StringOutput)
}

type CosRechargeArrayOutput struct{ *pulumi.OutputState }

func (CosRechargeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CosRecharge)(nil)).Elem()
}

func (o CosRechargeArrayOutput) ToCosRechargeArrayOutput() CosRechargeArrayOutput {
	return o
}

func (o CosRechargeArrayOutput) ToCosRechargeArrayOutputWithContext(ctx context.Context) CosRechargeArrayOutput {
	return o
}

func (o CosRechargeArrayOutput) Index(i pulumi.IntInput) CosRechargeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CosRecharge {
		return vs[0].([]*CosRecharge)[vs[1].(int)]
	}).(CosRechargeOutput)
}

type CosRechargeMapOutput struct{ *pulumi.OutputState }

func (CosRechargeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CosRecharge)(nil)).Elem()
}

func (o CosRechargeMapOutput) ToCosRechargeMapOutput() CosRechargeMapOutput {
	return o
}

func (o CosRechargeMapOutput) ToCosRechargeMapOutputWithContext(ctx context.Context) CosRechargeMapOutput {
	return o
}

func (o CosRechargeMapOutput) MapIndex(k pulumi.StringInput) CosRechargeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CosRecharge {
		return vs[0].(map[string]*CosRecharge)[vs[1].(string)]
	}).(CosRechargeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CosRechargeInput)(nil)).Elem(), &CosRecharge{})
	pulumi.RegisterInputType(reflect.TypeOf((*CosRechargeArrayInput)(nil)).Elem(), CosRechargeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CosRechargeMapInput)(nil)).Elem(), CosRechargeMap{})
	pulumi.RegisterOutputType(CosRechargeOutput{})
	pulumi.RegisterOutputType(CosRechargeArrayOutput{})
	pulumi.RegisterOutputType(CosRechargeMapOutput{})
}
