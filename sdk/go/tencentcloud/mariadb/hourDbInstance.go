// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mariadb hourDbInstance
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Mariadb.NewHourDbInstance(ctx, "basic", &Mariadb.HourDbInstanceArgs{
//				DbVersionId:  pulumi.String("10.0"),
//				InstanceName: pulumi.String("db-test-del"),
//				Memory:       pulumi.Int(2),
//				NodeCount:    pulumi.Int(2),
//				Storage:      pulumi.Int(10),
//				SubnetId:     pulumi.String("subnet-jdi5xn22"),
//				Tags: pulumi.AnyMap{
//					"createdBy": pulumi.Any("terraform"),
//				},
//				Vip:   pulumi.String("10.0.0.197"),
//				VpcId: pulumi.String("vpc-k1t8ickr"),
//				Zones: pulumi.StringArray{
//					pulumi.String("ap-guangzhou-6"),
//					pulumi.String("ap-guangzhou-7"),
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
// mariadb hour_db_instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Mariadb/hourDbInstance:HourDbInstance hour_db_instance tdsql-kjqih9nn
//
// ```
type HourDbInstance struct {
	pulumi.CustomResourceState

	// db engine version, default to 10.1.9.
	DbVersionId pulumi.StringOutput `pulumi:"dbVersionId"`
	// name of this instance.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// instance memory.
	Memory pulumi.IntOutput `pulumi:"memory"`
	// number of node for instance.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// project id.
	ProjectId pulumi.IntPtrOutput `pulumi:"projectId"`
	// instance disk storage.
	Storage pulumi.IntOutput `pulumi:"storage"`
	// subnet id, it&amp;#39;s required when vpcId is set.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// vip.
	Vip pulumi.StringOutput `pulumi:"vip"`
	// vpc id.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// available zone of instance.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewHourDbInstance registers a new resource with the given unique name, arguments, and options.
func NewHourDbInstance(ctx *pulumi.Context,
	name string, args *HourDbInstanceArgs, opts ...pulumi.ResourceOption) (*HourDbInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Memory == nil {
		return nil, errors.New("invalid value for required argument 'Memory'")
	}
	if args.NodeCount == nil {
		return nil, errors.New("invalid value for required argument 'NodeCount'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	if args.Zones == nil {
		return nil, errors.New("invalid value for required argument 'Zones'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource HourDbInstance
	err := ctx.RegisterResource("tencentcloud:Mariadb/hourDbInstance:HourDbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHourDbInstance gets an existing HourDbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHourDbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HourDbInstanceState, opts ...pulumi.ResourceOption) (*HourDbInstance, error) {
	var resource HourDbInstance
	err := ctx.ReadResource("tencentcloud:Mariadb/hourDbInstance:HourDbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HourDbInstance resources.
type hourDbInstanceState struct {
	// db engine version, default to 10.1.9.
	DbVersionId *string `pulumi:"dbVersionId"`
	// name of this instance.
	InstanceName *string `pulumi:"instanceName"`
	// instance memory.
	Memory *int `pulumi:"memory"`
	// number of node for instance.
	NodeCount *int `pulumi:"nodeCount"`
	// project id.
	ProjectId *int `pulumi:"projectId"`
	// instance disk storage.
	Storage *int `pulumi:"storage"`
	// subnet id, it&amp;#39;s required when vpcId is set.
	SubnetId *string `pulumi:"subnetId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// vip.
	Vip *string `pulumi:"vip"`
	// vpc id.
	VpcId *string `pulumi:"vpcId"`
	// available zone of instance.
	Zones []string `pulumi:"zones"`
}

type HourDbInstanceState struct {
	// db engine version, default to 10.1.9.
	DbVersionId pulumi.StringPtrInput
	// name of this instance.
	InstanceName pulumi.StringPtrInput
	// instance memory.
	Memory pulumi.IntPtrInput
	// number of node for instance.
	NodeCount pulumi.IntPtrInput
	// project id.
	ProjectId pulumi.IntPtrInput
	// instance disk storage.
	Storage pulumi.IntPtrInput
	// subnet id, it&amp;#39;s required when vpcId is set.
	SubnetId pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// vip.
	Vip pulumi.StringPtrInput
	// vpc id.
	VpcId pulumi.StringPtrInput
	// available zone of instance.
	Zones pulumi.StringArrayInput
}

func (HourDbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hourDbInstanceState)(nil)).Elem()
}

type hourDbInstanceArgs struct {
	// db engine version, default to 10.1.9.
	DbVersionId *string `pulumi:"dbVersionId"`
	// name of this instance.
	InstanceName *string `pulumi:"instanceName"`
	// instance memory.
	Memory int `pulumi:"memory"`
	// number of node for instance.
	NodeCount int `pulumi:"nodeCount"`
	// project id.
	ProjectId *int `pulumi:"projectId"`
	// instance disk storage.
	Storage int `pulumi:"storage"`
	// subnet id, it&amp;#39;s required when vpcId is set.
	SubnetId *string `pulumi:"subnetId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// vip.
	Vip *string `pulumi:"vip"`
	// vpc id.
	VpcId *string `pulumi:"vpcId"`
	// available zone of instance.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a HourDbInstance resource.
type HourDbInstanceArgs struct {
	// db engine version, default to 10.1.9.
	DbVersionId pulumi.StringPtrInput
	// name of this instance.
	InstanceName pulumi.StringPtrInput
	// instance memory.
	Memory pulumi.IntInput
	// number of node for instance.
	NodeCount pulumi.IntInput
	// project id.
	ProjectId pulumi.IntPtrInput
	// instance disk storage.
	Storage pulumi.IntInput
	// subnet id, it&amp;#39;s required when vpcId is set.
	SubnetId pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// vip.
	Vip pulumi.StringPtrInput
	// vpc id.
	VpcId pulumi.StringPtrInput
	// available zone of instance.
	Zones pulumi.StringArrayInput
}

func (HourDbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hourDbInstanceArgs)(nil)).Elem()
}

type HourDbInstanceInput interface {
	pulumi.Input

	ToHourDbInstanceOutput() HourDbInstanceOutput
	ToHourDbInstanceOutputWithContext(ctx context.Context) HourDbInstanceOutput
}

func (*HourDbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDbInstance)(nil)).Elem()
}

func (i *HourDbInstance) ToHourDbInstanceOutput() HourDbInstanceOutput {
	return i.ToHourDbInstanceOutputWithContext(context.Background())
}

func (i *HourDbInstance) ToHourDbInstanceOutputWithContext(ctx context.Context) HourDbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDbInstanceOutput)
}

// HourDbInstanceArrayInput is an input type that accepts HourDbInstanceArray and HourDbInstanceArrayOutput values.
// You can construct a concrete instance of `HourDbInstanceArrayInput` via:
//
//	HourDbInstanceArray{ HourDbInstanceArgs{...} }
type HourDbInstanceArrayInput interface {
	pulumi.Input

	ToHourDbInstanceArrayOutput() HourDbInstanceArrayOutput
	ToHourDbInstanceArrayOutputWithContext(context.Context) HourDbInstanceArrayOutput
}

type HourDbInstanceArray []HourDbInstanceInput

func (HourDbInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HourDbInstance)(nil)).Elem()
}

func (i HourDbInstanceArray) ToHourDbInstanceArrayOutput() HourDbInstanceArrayOutput {
	return i.ToHourDbInstanceArrayOutputWithContext(context.Background())
}

func (i HourDbInstanceArray) ToHourDbInstanceArrayOutputWithContext(ctx context.Context) HourDbInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDbInstanceArrayOutput)
}

// HourDbInstanceMapInput is an input type that accepts HourDbInstanceMap and HourDbInstanceMapOutput values.
// You can construct a concrete instance of `HourDbInstanceMapInput` via:
//
//	HourDbInstanceMap{ "key": HourDbInstanceArgs{...} }
type HourDbInstanceMapInput interface {
	pulumi.Input

	ToHourDbInstanceMapOutput() HourDbInstanceMapOutput
	ToHourDbInstanceMapOutputWithContext(context.Context) HourDbInstanceMapOutput
}

type HourDbInstanceMap map[string]HourDbInstanceInput

func (HourDbInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HourDbInstance)(nil)).Elem()
}

func (i HourDbInstanceMap) ToHourDbInstanceMapOutput() HourDbInstanceMapOutput {
	return i.ToHourDbInstanceMapOutputWithContext(context.Background())
}

func (i HourDbInstanceMap) ToHourDbInstanceMapOutputWithContext(ctx context.Context) HourDbInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HourDbInstanceMapOutput)
}

type HourDbInstanceOutput struct{ *pulumi.OutputState }

func (HourDbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HourDbInstance)(nil)).Elem()
}

func (o HourDbInstanceOutput) ToHourDbInstanceOutput() HourDbInstanceOutput {
	return o
}

func (o HourDbInstanceOutput) ToHourDbInstanceOutputWithContext(ctx context.Context) HourDbInstanceOutput {
	return o
}

// db engine version, default to 10.1.9.
func (o HourDbInstanceOutput) DbVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.StringOutput { return v.DbVersionId }).(pulumi.StringOutput)
}

// name of this instance.
func (o HourDbInstanceOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.StringPtrOutput { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// instance memory.
func (o HourDbInstanceOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

// number of node for instance.
func (o HourDbInstanceOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.IntOutput { return v.NodeCount }).(pulumi.IntOutput)
}

// project id.
func (o HourDbInstanceOutput) ProjectId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.IntPtrOutput { return v.ProjectId }).(pulumi.IntPtrOutput)
}

// instance disk storage.
func (o HourDbInstanceOutput) Storage() pulumi.IntOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.IntOutput { return v.Storage }).(pulumi.IntOutput)
}

// subnet id, it&amp;#39;s required when vpcId is set.
func (o HourDbInstanceOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Tag description list.
func (o HourDbInstanceOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// vip.
func (o HourDbInstanceOutput) Vip() pulumi.StringOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.StringOutput { return v.Vip }).(pulumi.StringOutput)
}

// vpc id.
func (o HourDbInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// available zone of instance.
func (o HourDbInstanceOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HourDbInstance) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

type HourDbInstanceArrayOutput struct{ *pulumi.OutputState }

func (HourDbInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HourDbInstance)(nil)).Elem()
}

func (o HourDbInstanceArrayOutput) ToHourDbInstanceArrayOutput() HourDbInstanceArrayOutput {
	return o
}

func (o HourDbInstanceArrayOutput) ToHourDbInstanceArrayOutputWithContext(ctx context.Context) HourDbInstanceArrayOutput {
	return o
}

func (o HourDbInstanceArrayOutput) Index(i pulumi.IntInput) HourDbInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HourDbInstance {
		return vs[0].([]*HourDbInstance)[vs[1].(int)]
	}).(HourDbInstanceOutput)
}

type HourDbInstanceMapOutput struct{ *pulumi.OutputState }

func (HourDbInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HourDbInstance)(nil)).Elem()
}

func (o HourDbInstanceMapOutput) ToHourDbInstanceMapOutput() HourDbInstanceMapOutput {
	return o
}

func (o HourDbInstanceMapOutput) ToHourDbInstanceMapOutputWithContext(ctx context.Context) HourDbInstanceMapOutput {
	return o
}

func (o HourDbInstanceMapOutput) MapIndex(k pulumi.StringInput) HourDbInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HourDbInstance {
		return vs[0].(map[string]*HourDbInstance)[vs[1].(string)]
	}).(HourDbInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HourDbInstanceInput)(nil)).Elem(), &HourDbInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*HourDbInstanceArrayInput)(nil)).Elem(), HourDbInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HourDbInstanceMapInput)(nil)).Elem(), HourDbInstanceMap{})
	pulumi.RegisterOutputType(HourDbInstanceOutput{})
	pulumi.RegisterOutputType(HourDbInstanceArrayOutput{})
	pulumi.RegisterOutputType(HourDbInstanceMapOutput{})
}
