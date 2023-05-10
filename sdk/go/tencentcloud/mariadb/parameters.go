// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a mariadb parameters
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Mariadb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Mariadb.NewParameters(ctx, "parameters", &Mariadb.ParametersArgs{
// 			InstanceId: pulumi.String("tdsql-4pzs5b67"),
// 			Params: mariadb.ParametersParamArray{
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("auto_increment_increment"),
// 					Value: pulumi.String("1"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("auto_increment_offset"),
// 					Value: pulumi.String("1"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("autocommit"),
// 					Value: pulumi.String("ON"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("character_set_server"),
// 					Value: pulumi.String("utf8mb4"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("collation_connection"),
// 					Value: pulumi.String("utf8mb4_general_ci"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("collation_database"),
// 					Value: pulumi.String("utf8mb4_general_ci"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("collation_server"),
// 					Value: pulumi.String("utf8mb4_general_ci"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("connect_timeout"),
// 					Value: pulumi.String("10"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("default_collation_for_utf8mb4"),
// 					Value: pulumi.String("utf8mb4_general_ci"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("default_week_format"),
// 					Value: pulumi.String("0"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("delay_key_write"),
// 					Value: pulumi.String("ON"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("delayed_insert_limit"),
// 					Value: pulumi.String("100"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("delayed_insert_timeout"),
// 					Value: pulumi.String("300"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("delayed_queue_size"),
// 					Value: pulumi.String("1000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("div_precision_increment"),
// 					Value: pulumi.String("4"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("event_scheduler"),
// 					Value: pulumi.String("ON"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("group_concat_max_len"),
// 					Value: pulumi.String("1024"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_concurrency_tickets"),
// 					Value: pulumi.String("5000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_flush_log_at_trx_commit"),
// 					Value: pulumi.String("1"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_lock_wait_timeout"),
// 					Value: pulumi.String("20"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_max_dirty_pages_pct"),
// 					Value: pulumi.String("70.000000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_old_blocks_pct"),
// 					Value: pulumi.String("37"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_old_blocks_time"),
// 					Value: pulumi.String("1000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_purge_batch_size"),
// 					Value: pulumi.String("1000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_read_ahead_threshold"),
// 					Value: pulumi.String("56"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_stats_method"),
// 					Value: pulumi.String("nulls_equal"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_stats_on_metadata"),
// 					Value: pulumi.String("OFF"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_strict_mode"),
// 					Value: pulumi.String("OFF"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_table_locks"),
// 					Value: pulumi.String("ON"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("innodb_thread_concurrency"),
// 					Value: pulumi.String("0"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("interactive_timeout"),
// 					Value: pulumi.String("28800"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("key_cache_age_threshold"),
// 					Value: pulumi.String("300"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("key_cache_block_size"),
// 					Value: pulumi.String("1024"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("key_cache_division_limit"),
// 					Value: pulumi.String("100"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("local_infile"),
// 					Value: pulumi.String("OFF"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("lock_wait_timeout"),
// 					Value: pulumi.String("5"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("log_queries_not_using_indexes"),
// 					Value: pulumi.String("OFF"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("long_query_time"),
// 					Value: pulumi.String("1.000000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("low_priority_updates"),
// 					Value: pulumi.String("OFF"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("max_allowed_packet"),
// 					Value: pulumi.String("1073741824"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("max_binlog_size"),
// 					Value: pulumi.String("536870912"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("max_connect_errors"),
// 					Value: pulumi.String("2000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("max_connections"),
// 					Value: pulumi.String("10000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("max_execution_time"),
// 					Value: pulumi.String("0"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("max_prepared_stmt_count"),
// 					Value: pulumi.String("200000"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("myisam_sort_buffer_size"),
// 					Value: pulumi.String("4194304"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("net_buffer_length"),
// 					Value: pulumi.String("16384"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("net_read_timeout"),
// 					Value: pulumi.String("150"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("net_retry_count"),
// 					Value: pulumi.String("10"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("net_write_timeout"),
// 					Value: pulumi.String("300"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("query_alloc_block_size"),
// 					Value: pulumi.String("16384"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("query_prealloc_size"),
// 					Value: pulumi.String("24576"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("slow_launch_time"),
// 					Value: pulumi.String("2"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("sort_buffer_size"),
// 					Value: pulumi.String("2097152"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("sql_mode"),
// 					Value: pulumi.String("NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("sql_require_primary_key"),
// 					Value: pulumi.String("ON"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("sql_safe_updates"),
// 					Value: pulumi.String("OFF"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("sqlasyntimeout"),
// 					Value: pulumi.String("30"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("sync_binlog"),
// 					Value: pulumi.String("1"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("table_definition_cache"),
// 					Value: pulumi.String("10240"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("table_open_cache"),
// 					Value: pulumi.String("20480"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("time_zone"),
// 					Value: pulumi.String("+08:00"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("tmp_table_size"),
// 					Value: pulumi.String("33554432"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("tx_isolation"),
// 					Value: pulumi.String("READ-COMMITTED"),
// 				},
// 				&mariadb.ParametersParamArgs{
// 					Param: pulumi.String("wait_timeout"),
// 					Value: pulumi.String("28800"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// mariadb parameters can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Mariadb/parameters:Parameters parameters tdsql-4pzs5b67
// ```
type Parameters struct {
	pulumi.CustomResourceState

	// instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Number of days to keep, no more than 30.
	Params ParametersParamArrayOutput `pulumi:"params"`
}

// NewParameters registers a new resource with the given unique name, arguments, and options.
func NewParameters(ctx *pulumi.Context,
	name string, args *ParametersArgs, opts ...pulumi.ResourceOption) (*Parameters, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Params == nil {
		return nil, errors.New("invalid value for required argument 'Params'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Parameters
	err := ctx.RegisterResource("tencentcloud:Mariadb/parameters:Parameters", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameters gets an existing Parameters resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameters(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParametersState, opts ...pulumi.ResourceOption) (*Parameters, error) {
	var resource Parameters
	err := ctx.ReadResource("tencentcloud:Mariadb/parameters:Parameters", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Parameters resources.
type parametersState struct {
	// instance id.
	InstanceId *string `pulumi:"instanceId"`
	// Number of days to keep, no more than 30.
	Params []ParametersParam `pulumi:"params"`
}

type ParametersState struct {
	// instance id.
	InstanceId pulumi.StringPtrInput
	// Number of days to keep, no more than 30.
	Params ParametersParamArrayInput
}

func (ParametersState) ElementType() reflect.Type {
	return reflect.TypeOf((*parametersState)(nil)).Elem()
}

type parametersArgs struct {
	// instance id.
	InstanceId string `pulumi:"instanceId"`
	// Number of days to keep, no more than 30.
	Params []ParametersParam `pulumi:"params"`
}

// The set of arguments for constructing a Parameters resource.
type ParametersArgs struct {
	// instance id.
	InstanceId pulumi.StringInput
	// Number of days to keep, no more than 30.
	Params ParametersParamArrayInput
}

func (ParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parametersArgs)(nil)).Elem()
}

type ParametersInput interface {
	pulumi.Input

	ToParametersOutput() ParametersOutput
	ToParametersOutputWithContext(ctx context.Context) ParametersOutput
}

func (*Parameters) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameters)(nil)).Elem()
}

func (i *Parameters) ToParametersOutput() ParametersOutput {
	return i.ToParametersOutputWithContext(context.Background())
}

func (i *Parameters) ToParametersOutputWithContext(ctx context.Context) ParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersOutput)
}

// ParametersArrayInput is an input type that accepts ParametersArray and ParametersArrayOutput values.
// You can construct a concrete instance of `ParametersArrayInput` via:
//
//          ParametersArray{ ParametersArgs{...} }
type ParametersArrayInput interface {
	pulumi.Input

	ToParametersArrayOutput() ParametersArrayOutput
	ToParametersArrayOutputWithContext(context.Context) ParametersArrayOutput
}

type ParametersArray []ParametersInput

func (ParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameters)(nil)).Elem()
}

func (i ParametersArray) ToParametersArrayOutput() ParametersArrayOutput {
	return i.ToParametersArrayOutputWithContext(context.Background())
}

func (i ParametersArray) ToParametersArrayOutputWithContext(ctx context.Context) ParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersArrayOutput)
}

// ParametersMapInput is an input type that accepts ParametersMap and ParametersMapOutput values.
// You can construct a concrete instance of `ParametersMapInput` via:
//
//          ParametersMap{ "key": ParametersArgs{...} }
type ParametersMapInput interface {
	pulumi.Input

	ToParametersMapOutput() ParametersMapOutput
	ToParametersMapOutputWithContext(context.Context) ParametersMapOutput
}

type ParametersMap map[string]ParametersInput

func (ParametersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameters)(nil)).Elem()
}

func (i ParametersMap) ToParametersMapOutput() ParametersMapOutput {
	return i.ToParametersMapOutputWithContext(context.Background())
}

func (i ParametersMap) ToParametersMapOutputWithContext(ctx context.Context) ParametersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParametersMapOutput)
}

type ParametersOutput struct{ *pulumi.OutputState }

func (ParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameters)(nil)).Elem()
}

func (o ParametersOutput) ToParametersOutput() ParametersOutput {
	return o
}

func (o ParametersOutput) ToParametersOutputWithContext(ctx context.Context) ParametersOutput {
	return o
}

// instance id.
func (o ParametersOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Parameters) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Number of days to keep, no more than 30.
func (o ParametersOutput) Params() ParametersParamArrayOutput {
	return o.ApplyT(func(v *Parameters) ParametersParamArrayOutput { return v.Params }).(ParametersParamArrayOutput)
}

type ParametersArrayOutput struct{ *pulumi.OutputState }

func (ParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameters)(nil)).Elem()
}

func (o ParametersArrayOutput) ToParametersArrayOutput() ParametersArrayOutput {
	return o
}

func (o ParametersArrayOutput) ToParametersArrayOutputWithContext(ctx context.Context) ParametersArrayOutput {
	return o
}

func (o ParametersArrayOutput) Index(i pulumi.IntInput) ParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Parameters {
		return vs[0].([]*Parameters)[vs[1].(int)]
	}).(ParametersOutput)
}

type ParametersMapOutput struct{ *pulumi.OutputState }

func (ParametersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameters)(nil)).Elem()
}

func (o ParametersMapOutput) ToParametersMapOutput() ParametersMapOutput {
	return o
}

func (o ParametersMapOutput) ToParametersMapOutputWithContext(ctx context.Context) ParametersMapOutput {
	return o
}

func (o ParametersMapOutput) MapIndex(k pulumi.StringInput) ParametersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Parameters {
		return vs[0].(map[string]*Parameters)[vs[1].(string)]
	}).(ParametersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersInput)(nil)).Elem(), &Parameters{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersArrayInput)(nil)).Elem(), ParametersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParametersMapInput)(nil)).Elem(), ParametersMap{})
	pulumi.RegisterOutputType(ParametersOutput{})
	pulumi.RegisterOutputType(ParametersArrayOutput{})
	pulumi.RegisterOutputType(ParametersMapOutput{})
}
