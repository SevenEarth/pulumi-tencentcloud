// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcaplus

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create TcaplusDB table.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcaplus"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		test, err := Tcaplus.NewCluster(ctx, "test", &Tcaplus.ClusterArgs{
// 			IdlType:               pulumi.String("PROTO"),
// 			ClusterName:           pulumi.String("tf_tcaplus_cluster_test"),
// 			VpcId:                 pulumi.String("vpc-7k6gzox6"),
// 			SubnetId:              pulumi.String("subnet-akwgvfa3"),
// 			Password:              pulumi.String("1qaA2k1wgvfa3ZZZ"),
// 			OldPasswordExpireLast: pulumi.Int(3600),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tablegroup, err := Tcaplus.NewTablegroup(ctx, "tablegroup", &Tcaplus.TablegroupArgs{
// 			ClusterId:      test.ID(),
// 			TablegroupName: pulumi.String("tf_test_group_name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		main, err := Tcaplus.NewIdl(ctx, "main", &Tcaplus.IdlArgs{
// 			ClusterId:    test.ID(),
// 			TablegroupId: tablegroup.ID(),
// 			FileName:     pulumi.String("tf_idl_test_2"),
// 			FileType:     pulumi.String("PROTO"),
// 			FileExtType:  pulumi.String("proto"),
// 			FileContent:  pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "    syntax = \"proto2\";\n", "    package myTcaplusTable;\n", "    import \"tcaplusservice.optionv1.proto\";\n", "    message tb_online {\n", "       option(tcaplusservice.tcaplus_primary_key) = \"uin,name,region\";\n", "        required int64 uin = 1;\n", "        required string name = 2;\n", "        required int32 region = 3;\n", "        required int32 gamesvrid = 4;\n", "        optional int32 logintime = 5 [default = 1];\n", "        repeated int64 lockid = 6 [packed = true];\n", "        optional bool is_available = 7 [default = false];\n", "        optional pay_info pay = 8;\n", "    }\n", "\n", "    message pay_info {\n", "        required int64 pay_id = 1;\n", "        optional uint64 total_money = 2;\n", "        optional uint64 pay_times = 3;\n", "        optional pay_auth_info auth = 4;\n", "        message pay_auth_info {\n", "            required string pay_keys = 1;\n", "            optional int64 update_time = 2;\n", "        }\n", "    }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Tcaplus.NewTable(ctx, "table", &Tcaplus.TableArgs{
// 			ClusterId:       test.ID(),
// 			TablegroupId:    tablegroup.ID(),
// 			TableName:       pulumi.String("tb_online"),
// 			TableType:       pulumi.String("GENERIC"),
// 			Description:     pulumi.String("test"),
// 			IdlId:           main.ID(),
// 			TableIdlType:    pulumi.String("PROTO"),
// 			ReservedReadCu:  pulumi.Int(1000),
// 			ReservedWriteCu: pulumi.Int(20),
// 			ReservedVolume:  pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Table struct {
	pulumi.CustomResourceState

	// ID of the TcaplusDB cluster to which the table belongs.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Create time of the TcaplusDB table.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the TcaplusDB table.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Error messages for creating TcaplusDB table.
	Error pulumi.StringOutput `pulumi:"error"`
	// ID of the IDL File.
	IdlId pulumi.StringOutput `pulumi:"idlId"`
	// Reserved read capacity units of the TcaplusDB table.
	ReservedReadCu pulumi.IntOutput `pulumi:"reservedReadCu"`
	// Reserved storage capacity of the TcaplusDB table (unit: GB).
	ReservedVolume pulumi.IntOutput `pulumi:"reservedVolume"`
	// Reserved write capacity units of the TcaplusDB table.
	ReservedWriteCu pulumi.IntOutput `pulumi:"reservedWriteCu"`
	// Status of the TcaplusDB table.
	Status pulumi.StringOutput `pulumi:"status"`
	// IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
	TableIdlType pulumi.StringOutput `pulumi:"tableIdlType"`
	// Name of the TcaplusDB table.
	TableName pulumi.StringOutput `pulumi:"tableName"`
	// Size of the TcaplusDB table.
	TableSize pulumi.IntOutput `pulumi:"tableSize"`
	// Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
	TableType pulumi.StringOutput `pulumi:"tableType"`
	// ID of the table group to which the table belongs.
	TablegroupId pulumi.StringOutput `pulumi:"tablegroupId"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.IdlId == nil {
		return nil, errors.New("invalid value for required argument 'IdlId'")
	}
	if args.ReservedReadCu == nil {
		return nil, errors.New("invalid value for required argument 'ReservedReadCu'")
	}
	if args.ReservedVolume == nil {
		return nil, errors.New("invalid value for required argument 'ReservedVolume'")
	}
	if args.ReservedWriteCu == nil {
		return nil, errors.New("invalid value for required argument 'ReservedWriteCu'")
	}
	if args.TableIdlType == nil {
		return nil, errors.New("invalid value for required argument 'TableIdlType'")
	}
	if args.TableName == nil {
		return nil, errors.New("invalid value for required argument 'TableName'")
	}
	if args.TableType == nil {
		return nil, errors.New("invalid value for required argument 'TableType'")
	}
	if args.TablegroupId == nil {
		return nil, errors.New("invalid value for required argument 'TablegroupId'")
	}
	var resource Table
	err := ctx.RegisterResource("tencentcloud:Tcaplus/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("tencentcloud:Tcaplus/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// ID of the TcaplusDB cluster to which the table belongs.
	ClusterId *string `pulumi:"clusterId"`
	// Create time of the TcaplusDB table.
	CreateTime *string `pulumi:"createTime"`
	// Description of the TcaplusDB table.
	Description *string `pulumi:"description"`
	// Error messages for creating TcaplusDB table.
	Error *string `pulumi:"error"`
	// ID of the IDL File.
	IdlId *string `pulumi:"idlId"`
	// Reserved read capacity units of the TcaplusDB table.
	ReservedReadCu *int `pulumi:"reservedReadCu"`
	// Reserved storage capacity of the TcaplusDB table (unit: GB).
	ReservedVolume *int `pulumi:"reservedVolume"`
	// Reserved write capacity units of the TcaplusDB table.
	ReservedWriteCu *int `pulumi:"reservedWriteCu"`
	// Status of the TcaplusDB table.
	Status *string `pulumi:"status"`
	// IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
	TableIdlType *string `pulumi:"tableIdlType"`
	// Name of the TcaplusDB table.
	TableName *string `pulumi:"tableName"`
	// Size of the TcaplusDB table.
	TableSize *int `pulumi:"tableSize"`
	// Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
	TableType *string `pulumi:"tableType"`
	// ID of the table group to which the table belongs.
	TablegroupId *string `pulumi:"tablegroupId"`
}

type TableState struct {
	// ID of the TcaplusDB cluster to which the table belongs.
	ClusterId pulumi.StringPtrInput
	// Create time of the TcaplusDB table.
	CreateTime pulumi.StringPtrInput
	// Description of the TcaplusDB table.
	Description pulumi.StringPtrInput
	// Error messages for creating TcaplusDB table.
	Error pulumi.StringPtrInput
	// ID of the IDL File.
	IdlId pulumi.StringPtrInput
	// Reserved read capacity units of the TcaplusDB table.
	ReservedReadCu pulumi.IntPtrInput
	// Reserved storage capacity of the TcaplusDB table (unit: GB).
	ReservedVolume pulumi.IntPtrInput
	// Reserved write capacity units of the TcaplusDB table.
	ReservedWriteCu pulumi.IntPtrInput
	// Status of the TcaplusDB table.
	Status pulumi.StringPtrInput
	// IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
	TableIdlType pulumi.StringPtrInput
	// Name of the TcaplusDB table.
	TableName pulumi.StringPtrInput
	// Size of the TcaplusDB table.
	TableSize pulumi.IntPtrInput
	// Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
	TableType pulumi.StringPtrInput
	// ID of the table group to which the table belongs.
	TablegroupId pulumi.StringPtrInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// ID of the TcaplusDB cluster to which the table belongs.
	ClusterId string `pulumi:"clusterId"`
	// Description of the TcaplusDB table.
	Description *string `pulumi:"description"`
	// ID of the IDL File.
	IdlId string `pulumi:"idlId"`
	// Reserved read capacity units of the TcaplusDB table.
	ReservedReadCu int `pulumi:"reservedReadCu"`
	// Reserved storage capacity of the TcaplusDB table (unit: GB).
	ReservedVolume int `pulumi:"reservedVolume"`
	// Reserved write capacity units of the TcaplusDB table.
	ReservedWriteCu int `pulumi:"reservedWriteCu"`
	// IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
	TableIdlType string `pulumi:"tableIdlType"`
	// Name of the TcaplusDB table.
	TableName string `pulumi:"tableName"`
	// Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
	TableType string `pulumi:"tableType"`
	// ID of the table group to which the table belongs.
	TablegroupId string `pulumi:"tablegroupId"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// ID of the TcaplusDB cluster to which the table belongs.
	ClusterId pulumi.StringInput
	// Description of the TcaplusDB table.
	Description pulumi.StringPtrInput
	// ID of the IDL File.
	IdlId pulumi.StringInput
	// Reserved read capacity units of the TcaplusDB table.
	ReservedReadCu pulumi.IntInput
	// Reserved storage capacity of the TcaplusDB table (unit: GB).
	ReservedVolume pulumi.IntInput
	// Reserved write capacity units of the TcaplusDB table.
	ReservedWriteCu pulumi.IntInput
	// IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
	TableIdlType pulumi.StringInput
	// Name of the TcaplusDB table.
	TableName pulumi.StringInput
	// Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
	TableType pulumi.StringInput
	// ID of the table group to which the table belongs.
	TablegroupId pulumi.StringInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

// TableArrayInput is an input type that accepts TableArray and TableArrayOutput values.
// You can construct a concrete instance of `TableArrayInput` via:
//
//          TableArray{ TableArgs{...} }
type TableArrayInput interface {
	pulumi.Input

	ToTableArrayOutput() TableArrayOutput
	ToTableArrayOutputWithContext(context.Context) TableArrayOutput
}

type TableArray []TableInput

func (TableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (i TableArray) ToTableArrayOutput() TableArrayOutput {
	return i.ToTableArrayOutputWithContext(context.Background())
}

func (i TableArray) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableArrayOutput)
}

// TableMapInput is an input type that accepts TableMap and TableMapOutput values.
// You can construct a concrete instance of `TableMapInput` via:
//
//          TableMap{ "key": TableArgs{...} }
type TableMapInput interface {
	pulumi.Input

	ToTableMapOutput() TableMapOutput
	ToTableMapOutputWithContext(context.Context) TableMapOutput
}

type TableMap map[string]TableInput

func (TableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (i TableMap) ToTableMapOutput() TableMapOutput {
	return i.ToTableMapOutputWithContext(context.Background())
}

func (i TableMap) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableMapOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

// ID of the TcaplusDB cluster to which the table belongs.
func (o TableOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Create time of the TcaplusDB table.
func (o TableOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the TcaplusDB table.
func (o TableOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Error messages for creating TcaplusDB table.
func (o TableOutput) Error() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Error }).(pulumi.StringOutput)
}

// ID of the IDL File.
func (o TableOutput) IdlId() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.IdlId }).(pulumi.StringOutput)
}

// Reserved read capacity units of the TcaplusDB table.
func (o TableOutput) ReservedReadCu() pulumi.IntOutput {
	return o.ApplyT(func(v *Table) pulumi.IntOutput { return v.ReservedReadCu }).(pulumi.IntOutput)
}

// Reserved storage capacity of the TcaplusDB table (unit: GB).
func (o TableOutput) ReservedVolume() pulumi.IntOutput {
	return o.ApplyT(func(v *Table) pulumi.IntOutput { return v.ReservedVolume }).(pulumi.IntOutput)
}

// Reserved write capacity units of the TcaplusDB table.
func (o TableOutput) ReservedWriteCu() pulumi.IntOutput {
	return o.ApplyT(func(v *Table) pulumi.IntOutput { return v.ReservedWriteCu }).(pulumi.IntOutput)
}

// Status of the TcaplusDB table.
func (o TableOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// IDL type of the TcaplusDB table. Valid values: `PROTO` and `TDR`.
func (o TableOutput) TableIdlType() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TableIdlType }).(pulumi.StringOutput)
}

// Name of the TcaplusDB table.
func (o TableOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TableName }).(pulumi.StringOutput)
}

// Size of the TcaplusDB table.
func (o TableOutput) TableSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Table) pulumi.IntOutput { return v.TableSize }).(pulumi.IntOutput)
}

// Type of the TcaplusDB table. Valid values are `GENERIC` and `LIST`.
func (o TableOutput) TableType() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TableType }).(pulumi.StringOutput)
}

// ID of the table group to which the table belongs.
func (o TableOutput) TablegroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.TablegroupId }).(pulumi.StringOutput)
}

type TableArrayOutput struct{ *pulumi.OutputState }

func (TableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Table)(nil)).Elem()
}

func (o TableArrayOutput) ToTableArrayOutput() TableArrayOutput {
	return o
}

func (o TableArrayOutput) ToTableArrayOutputWithContext(ctx context.Context) TableArrayOutput {
	return o
}

func (o TableArrayOutput) Index(i pulumi.IntInput) TableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Table {
		return vs[0].([]*Table)[vs[1].(int)]
	}).(TableOutput)
}

type TableMapOutput struct{ *pulumi.OutputState }

func (TableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Table)(nil)).Elem()
}

func (o TableMapOutput) ToTableMapOutput() TableMapOutput {
	return o
}

func (o TableMapOutput) ToTableMapOutputWithContext(ctx context.Context) TableMapOutput {
	return o
}

func (o TableMapOutput) MapIndex(k pulumi.StringInput) TableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Table {
		return vs[0].(map[string]*Table)[vs[1].(string)]
	}).(TableOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableArrayInput)(nil)).Elem(), TableArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableMapInput)(nil)).Elem(), TableMap{})
	pulumi.RegisterOutputType(TableOutput{})
	pulumi.RegisterOutputType(TableArrayOutput{})
	pulumi.RegisterOutputType(TableMapOutput{})
}
