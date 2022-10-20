// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcaplus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create TcaplusDB table group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcaplus"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			test, err := Tcaplus.NewCluster(ctx, "test", &Tcaplus.ClusterArgs{
//				IdlType:               pulumi.String("PROTO"),
//				ClusterName:           pulumi.String("tf_tcaplus_cluster_test"),
//				VpcId:                 pulumi.String("vpc-7k6gzox6"),
//				SubnetId:              pulumi.String("subnet-akwgvfa3"),
//				Password:              pulumi.String("1qaA2k1wgvfa3ZZZ"),
//				OldPasswordExpireLast: pulumi.Int(3600),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Tcaplus.NewTablegroup(ctx, "tablegroup", &Tcaplus.TablegroupArgs{
//				ClusterId:      test.ID(),
//				TablegroupName: pulumi.String("tf_test_group_name"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Tablegroup struct {
	pulumi.CustomResourceState

	// ID of the TcaplusDB cluster to which the table group belongs.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Create time of the TcaplusDB table group.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Number of tables.
	TableCount pulumi.IntOutput `pulumi:"tableCount"`
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	TablegroupName pulumi.StringOutput `pulumi:"tablegroupName"`
	// Total storage size (MB).
	TotalSize pulumi.IntOutput `pulumi:"totalSize"`
}

// NewTablegroup registers a new resource with the given unique name, arguments, and options.
func NewTablegroup(ctx *pulumi.Context,
	name string, args *TablegroupArgs, opts ...pulumi.ResourceOption) (*Tablegroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.TablegroupName == nil {
		return nil, errors.New("invalid value for required argument 'TablegroupName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Tablegroup
	err := ctx.RegisterResource("tencentcloud:Tcaplus/tablegroup:Tablegroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTablegroup gets an existing Tablegroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTablegroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TablegroupState, opts ...pulumi.ResourceOption) (*Tablegroup, error) {
	var resource Tablegroup
	err := ctx.ReadResource("tencentcloud:Tcaplus/tablegroup:Tablegroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tablegroup resources.
type tablegroupState struct {
	// ID of the TcaplusDB cluster to which the table group belongs.
	ClusterId *string `pulumi:"clusterId"`
	// Create time of the TcaplusDB table group.
	CreateTime *string `pulumi:"createTime"`
	// Number of tables.
	TableCount *int `pulumi:"tableCount"`
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	TablegroupName *string `pulumi:"tablegroupName"`
	// Total storage size (MB).
	TotalSize *int `pulumi:"totalSize"`
}

type TablegroupState struct {
	// ID of the TcaplusDB cluster to which the table group belongs.
	ClusterId pulumi.StringPtrInput
	// Create time of the TcaplusDB table group.
	CreateTime pulumi.StringPtrInput
	// Number of tables.
	TableCount pulumi.IntPtrInput
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	TablegroupName pulumi.StringPtrInput
	// Total storage size (MB).
	TotalSize pulumi.IntPtrInput
}

func (TablegroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*tablegroupState)(nil)).Elem()
}

type tablegroupArgs struct {
	// ID of the TcaplusDB cluster to which the table group belongs.
	ClusterId string `pulumi:"clusterId"`
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	TablegroupName string `pulumi:"tablegroupName"`
}

// The set of arguments for constructing a Tablegroup resource.
type TablegroupArgs struct {
	// ID of the TcaplusDB cluster to which the table group belongs.
	ClusterId pulumi.StringInput
	// Name of the TcaplusDB table group. Name length should be between 1 and 30.
	TablegroupName pulumi.StringInput
}

func (TablegroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tablegroupArgs)(nil)).Elem()
}

type TablegroupInput interface {
	pulumi.Input

	ToTablegroupOutput() TablegroupOutput
	ToTablegroupOutputWithContext(ctx context.Context) TablegroupOutput
}

func (*Tablegroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Tablegroup)(nil)).Elem()
}

func (i *Tablegroup) ToTablegroupOutput() TablegroupOutput {
	return i.ToTablegroupOutputWithContext(context.Background())
}

func (i *Tablegroup) ToTablegroupOutputWithContext(ctx context.Context) TablegroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablegroupOutput)
}

// TablegroupArrayInput is an input type that accepts TablegroupArray and TablegroupArrayOutput values.
// You can construct a concrete instance of `TablegroupArrayInput` via:
//
//	TablegroupArray{ TablegroupArgs{...} }
type TablegroupArrayInput interface {
	pulumi.Input

	ToTablegroupArrayOutput() TablegroupArrayOutput
	ToTablegroupArrayOutputWithContext(context.Context) TablegroupArrayOutput
}

type TablegroupArray []TablegroupInput

func (TablegroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tablegroup)(nil)).Elem()
}

func (i TablegroupArray) ToTablegroupArrayOutput() TablegroupArrayOutput {
	return i.ToTablegroupArrayOutputWithContext(context.Background())
}

func (i TablegroupArray) ToTablegroupArrayOutputWithContext(ctx context.Context) TablegroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablegroupArrayOutput)
}

// TablegroupMapInput is an input type that accepts TablegroupMap and TablegroupMapOutput values.
// You can construct a concrete instance of `TablegroupMapInput` via:
//
//	TablegroupMap{ "key": TablegroupArgs{...} }
type TablegroupMapInput interface {
	pulumi.Input

	ToTablegroupMapOutput() TablegroupMapOutput
	ToTablegroupMapOutputWithContext(context.Context) TablegroupMapOutput
}

type TablegroupMap map[string]TablegroupInput

func (TablegroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tablegroup)(nil)).Elem()
}

func (i TablegroupMap) ToTablegroupMapOutput() TablegroupMapOutput {
	return i.ToTablegroupMapOutputWithContext(context.Background())
}

func (i TablegroupMap) ToTablegroupMapOutputWithContext(ctx context.Context) TablegroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TablegroupMapOutput)
}

type TablegroupOutput struct{ *pulumi.OutputState }

func (TablegroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tablegroup)(nil)).Elem()
}

func (o TablegroupOutput) ToTablegroupOutput() TablegroupOutput {
	return o
}

func (o TablegroupOutput) ToTablegroupOutputWithContext(ctx context.Context) TablegroupOutput {
	return o
}

// ID of the TcaplusDB cluster to which the table group belongs.
func (o TablegroupOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Tablegroup) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Create time of the TcaplusDB table group.
func (o TablegroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Tablegroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Number of tables.
func (o TablegroupOutput) TableCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Tablegroup) pulumi.IntOutput { return v.TableCount }).(pulumi.IntOutput)
}

// Name of the TcaplusDB table group. Name length should be between 1 and 30.
func (o TablegroupOutput) TablegroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Tablegroup) pulumi.StringOutput { return v.TablegroupName }).(pulumi.StringOutput)
}

// Total storage size (MB).
func (o TablegroupOutput) TotalSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Tablegroup) pulumi.IntOutput { return v.TotalSize }).(pulumi.IntOutput)
}

type TablegroupArrayOutput struct{ *pulumi.OutputState }

func (TablegroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tablegroup)(nil)).Elem()
}

func (o TablegroupArrayOutput) ToTablegroupArrayOutput() TablegroupArrayOutput {
	return o
}

func (o TablegroupArrayOutput) ToTablegroupArrayOutputWithContext(ctx context.Context) TablegroupArrayOutput {
	return o
}

func (o TablegroupArrayOutput) Index(i pulumi.IntInput) TablegroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tablegroup {
		return vs[0].([]*Tablegroup)[vs[1].(int)]
	}).(TablegroupOutput)
}

type TablegroupMapOutput struct{ *pulumi.OutputState }

func (TablegroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tablegroup)(nil)).Elem()
}

func (o TablegroupMapOutput) ToTablegroupMapOutput() TablegroupMapOutput {
	return o
}

func (o TablegroupMapOutput) ToTablegroupMapOutputWithContext(ctx context.Context) TablegroupMapOutput {
	return o
}

func (o TablegroupMapOutput) MapIndex(k pulumi.StringInput) TablegroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tablegroup {
		return vs[0].(map[string]*Tablegroup)[vs[1].(string)]
	}).(TablegroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TablegroupInput)(nil)).Elem(), &Tablegroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*TablegroupArrayInput)(nil)).Elem(), TablegroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TablegroupMapInput)(nil)).Elem(), TablegroupMap{})
	pulumi.RegisterOutputType(TablegroupOutput{})
	pulumi.RegisterOutputType(TablegroupArrayOutput{})
	pulumi.RegisterOutputType(TablegroupMapOutput{})
}
