// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bi

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a bi datasource
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Bi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Bi.NewDatasource(ctx, "datasource", &Bi.DatasourceArgs{
//				Charset:    pulumi.String("utf8"),
//				DbHost:     pulumi.String("bj-cdb-1lxqg5r6.sql.tencentcdb.com"),
//				DbName:     pulumi.String("tf-test"),
//				DbPort:     pulumi.Int(63694),
//				DbPwd:      pulumi.String("ABc123,,,"),
//				DbType:     pulumi.String("MYSQL"),
//				DbUser:     pulumi.String("root"),
//				ProjectId:  pulumi.Int(11015030),
//				SourceName: pulumi.String("tf-source-name"),
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
// bi datasource can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Bi/datasource:Datasource datasource datasource_id
// ```
type Datasource struct {
	pulumi.CustomResourceState

	// Catalog.
	Catalog pulumi.StringPtrOutput `pulumi:"catalog"`
	// Charset.
	Charset pulumi.StringOutput `pulumi:"charset"`
	// Third-party datasource identification, this parameter can be ignored.
	DataOrigin pulumi.StringPtrOutput `pulumi:"dataOrigin"`
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginDatasourceId pulumi.StringPtrOutput `pulumi:"dataOriginDatasourceId"`
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginProjectId pulumi.StringPtrOutput `pulumi:"dataOriginProjectId"`
	// Host.
	DbHost pulumi.StringOutput `pulumi:"dbHost"`
	// Database name.
	DbName pulumi.StringOutput `pulumi:"dbName"`
	// Port.
	DbPort pulumi.IntOutput `pulumi:"dbPort"`
	// Password.
	DbPwd pulumi.StringOutput `pulumi:"dbPwd"`
	// `MYSQL`, `MSSQL`, `POSTGRE`, `ORACLE`, `CLICKHOUSE`, `TIDB`, `HIVE`, `PRESTO`.
	DbType pulumi.StringOutput `pulumi:"dbType"`
	// User name.
	DbUser pulumi.StringOutput `pulumi:"dbUser"`
	// Project id.
	ProjectId pulumi.IntOutput `pulumi:"projectId"`
	// Own or Cloud, default: `Own`.
	ServiceType pulumi.StringPtrOutput `pulumi:"serviceType"`
	// Datasource name in BI.
	SourceName pulumi.StringOutput `pulumi:"sourceName"`
	// Tencent cloud private network unified identity.
	UniqVpcId pulumi.StringPtrOutput `pulumi:"uniqVpcId"`
	// Tencent cloud private network identity.
	VpcId pulumi.StringPtrOutput `pulumi:"vpcId"`
}

// NewDatasource registers a new resource with the given unique name, arguments, and options.
func NewDatasource(ctx *pulumi.Context,
	name string, args *DatasourceArgs, opts ...pulumi.ResourceOption) (*Datasource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Charset == nil {
		return nil, errors.New("invalid value for required argument 'Charset'")
	}
	if args.DbHost == nil {
		return nil, errors.New("invalid value for required argument 'DbHost'")
	}
	if args.DbName == nil {
		return nil, errors.New("invalid value for required argument 'DbName'")
	}
	if args.DbPort == nil {
		return nil, errors.New("invalid value for required argument 'DbPort'")
	}
	if args.DbPwd == nil {
		return nil, errors.New("invalid value for required argument 'DbPwd'")
	}
	if args.DbType == nil {
		return nil, errors.New("invalid value for required argument 'DbType'")
	}
	if args.DbUser == nil {
		return nil, errors.New("invalid value for required argument 'DbUser'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.SourceName == nil {
		return nil, errors.New("invalid value for required argument 'SourceName'")
	}
	if args.DbPwd != nil {
		args.DbPwd = pulumi.ToSecret(args.DbPwd).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"dbPwd",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Datasource
	err := ctx.RegisterResource("tencentcloud:Bi/datasource:Datasource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasource gets an existing Datasource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasourceState, opts ...pulumi.ResourceOption) (*Datasource, error) {
	var resource Datasource
	err := ctx.ReadResource("tencentcloud:Bi/datasource:Datasource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Datasource resources.
type datasourceState struct {
	// Catalog.
	Catalog *string `pulumi:"catalog"`
	// Charset.
	Charset *string `pulumi:"charset"`
	// Third-party datasource identification, this parameter can be ignored.
	DataOrigin *string `pulumi:"dataOrigin"`
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginDatasourceId *string `pulumi:"dataOriginDatasourceId"`
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginProjectId *string `pulumi:"dataOriginProjectId"`
	// Host.
	DbHost *string `pulumi:"dbHost"`
	// Database name.
	DbName *string `pulumi:"dbName"`
	// Port.
	DbPort *int `pulumi:"dbPort"`
	// Password.
	DbPwd *string `pulumi:"dbPwd"`
	// `MYSQL`, `MSSQL`, `POSTGRE`, `ORACLE`, `CLICKHOUSE`, `TIDB`, `HIVE`, `PRESTO`.
	DbType *string `pulumi:"dbType"`
	// User name.
	DbUser *string `pulumi:"dbUser"`
	// Project id.
	ProjectId *int `pulumi:"projectId"`
	// Own or Cloud, default: `Own`.
	ServiceType *string `pulumi:"serviceType"`
	// Datasource name in BI.
	SourceName *string `pulumi:"sourceName"`
	// Tencent cloud private network unified identity.
	UniqVpcId *string `pulumi:"uniqVpcId"`
	// Tencent cloud private network identity.
	VpcId *string `pulumi:"vpcId"`
}

type DatasourceState struct {
	// Catalog.
	Catalog pulumi.StringPtrInput
	// Charset.
	Charset pulumi.StringPtrInput
	// Third-party datasource identification, this parameter can be ignored.
	DataOrigin pulumi.StringPtrInput
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginDatasourceId pulumi.StringPtrInput
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginProjectId pulumi.StringPtrInput
	// Host.
	DbHost pulumi.StringPtrInput
	// Database name.
	DbName pulumi.StringPtrInput
	// Port.
	DbPort pulumi.IntPtrInput
	// Password.
	DbPwd pulumi.StringPtrInput
	// `MYSQL`, `MSSQL`, `POSTGRE`, `ORACLE`, `CLICKHOUSE`, `TIDB`, `HIVE`, `PRESTO`.
	DbType pulumi.StringPtrInput
	// User name.
	DbUser pulumi.StringPtrInput
	// Project id.
	ProjectId pulumi.IntPtrInput
	// Own or Cloud, default: `Own`.
	ServiceType pulumi.StringPtrInput
	// Datasource name in BI.
	SourceName pulumi.StringPtrInput
	// Tencent cloud private network unified identity.
	UniqVpcId pulumi.StringPtrInput
	// Tencent cloud private network identity.
	VpcId pulumi.StringPtrInput
}

func (DatasourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasourceState)(nil)).Elem()
}

type datasourceArgs struct {
	// Catalog.
	Catalog *string `pulumi:"catalog"`
	// Charset.
	Charset string `pulumi:"charset"`
	// Third-party datasource identification, this parameter can be ignored.
	DataOrigin *string `pulumi:"dataOrigin"`
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginDatasourceId *string `pulumi:"dataOriginDatasourceId"`
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginProjectId *string `pulumi:"dataOriginProjectId"`
	// Host.
	DbHost string `pulumi:"dbHost"`
	// Database name.
	DbName string `pulumi:"dbName"`
	// Port.
	DbPort int `pulumi:"dbPort"`
	// Password.
	DbPwd string `pulumi:"dbPwd"`
	// `MYSQL`, `MSSQL`, `POSTGRE`, `ORACLE`, `CLICKHOUSE`, `TIDB`, `HIVE`, `PRESTO`.
	DbType string `pulumi:"dbType"`
	// User name.
	DbUser string `pulumi:"dbUser"`
	// Project id.
	ProjectId int `pulumi:"projectId"`
	// Own or Cloud, default: `Own`.
	ServiceType *string `pulumi:"serviceType"`
	// Datasource name in BI.
	SourceName string `pulumi:"sourceName"`
	// Tencent cloud private network unified identity.
	UniqVpcId *string `pulumi:"uniqVpcId"`
	// Tencent cloud private network identity.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Datasource resource.
type DatasourceArgs struct {
	// Catalog.
	Catalog pulumi.StringPtrInput
	// Charset.
	Charset pulumi.StringInput
	// Third-party datasource identification, this parameter can be ignored.
	DataOrigin pulumi.StringPtrInput
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginDatasourceId pulumi.StringPtrInput
	// Third-party datasource project id, this parameter can be ignored.
	DataOriginProjectId pulumi.StringPtrInput
	// Host.
	DbHost pulumi.StringInput
	// Database name.
	DbName pulumi.StringInput
	// Port.
	DbPort pulumi.IntInput
	// Password.
	DbPwd pulumi.StringInput
	// `MYSQL`, `MSSQL`, `POSTGRE`, `ORACLE`, `CLICKHOUSE`, `TIDB`, `HIVE`, `PRESTO`.
	DbType pulumi.StringInput
	// User name.
	DbUser pulumi.StringInput
	// Project id.
	ProjectId pulumi.IntInput
	// Own or Cloud, default: `Own`.
	ServiceType pulumi.StringPtrInput
	// Datasource name in BI.
	SourceName pulumi.StringInput
	// Tencent cloud private network unified identity.
	UniqVpcId pulumi.StringPtrInput
	// Tencent cloud private network identity.
	VpcId pulumi.StringPtrInput
}

func (DatasourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasourceArgs)(nil)).Elem()
}

type DatasourceInput interface {
	pulumi.Input

	ToDatasourceOutput() DatasourceOutput
	ToDatasourceOutputWithContext(ctx context.Context) DatasourceOutput
}

func (*Datasource) ElementType() reflect.Type {
	return reflect.TypeOf((**Datasource)(nil)).Elem()
}

func (i *Datasource) ToDatasourceOutput() DatasourceOutput {
	return i.ToDatasourceOutputWithContext(context.Background())
}

func (i *Datasource) ToDatasourceOutputWithContext(ctx context.Context) DatasourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceOutput)
}

// DatasourceArrayInput is an input type that accepts DatasourceArray and DatasourceArrayOutput values.
// You can construct a concrete instance of `DatasourceArrayInput` via:
//
//	DatasourceArray{ DatasourceArgs{...} }
type DatasourceArrayInput interface {
	pulumi.Input

	ToDatasourceArrayOutput() DatasourceArrayOutput
	ToDatasourceArrayOutputWithContext(context.Context) DatasourceArrayOutput
}

type DatasourceArray []DatasourceInput

func (DatasourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datasource)(nil)).Elem()
}

func (i DatasourceArray) ToDatasourceArrayOutput() DatasourceArrayOutput {
	return i.ToDatasourceArrayOutputWithContext(context.Background())
}

func (i DatasourceArray) ToDatasourceArrayOutputWithContext(ctx context.Context) DatasourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceArrayOutput)
}

// DatasourceMapInput is an input type that accepts DatasourceMap and DatasourceMapOutput values.
// You can construct a concrete instance of `DatasourceMapInput` via:
//
//	DatasourceMap{ "key": DatasourceArgs{...} }
type DatasourceMapInput interface {
	pulumi.Input

	ToDatasourceMapOutput() DatasourceMapOutput
	ToDatasourceMapOutputWithContext(context.Context) DatasourceMapOutput
}

type DatasourceMap map[string]DatasourceInput

func (DatasourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datasource)(nil)).Elem()
}

func (i DatasourceMap) ToDatasourceMapOutput() DatasourceMapOutput {
	return i.ToDatasourceMapOutputWithContext(context.Background())
}

func (i DatasourceMap) ToDatasourceMapOutputWithContext(ctx context.Context) DatasourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasourceMapOutput)
}

type DatasourceOutput struct{ *pulumi.OutputState }

func (DatasourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Datasource)(nil)).Elem()
}

func (o DatasourceOutput) ToDatasourceOutput() DatasourceOutput {
	return o
}

func (o DatasourceOutput) ToDatasourceOutputWithContext(ctx context.Context) DatasourceOutput {
	return o
}

// Catalog.
func (o DatasourceOutput) Catalog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringPtrOutput { return v.Catalog }).(pulumi.StringPtrOutput)
}

// Charset.
func (o DatasourceOutput) Charset() pulumi.StringOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringOutput { return v.Charset }).(pulumi.StringOutput)
}

// Third-party datasource identification, this parameter can be ignored.
func (o DatasourceOutput) DataOrigin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringPtrOutput { return v.DataOrigin }).(pulumi.StringPtrOutput)
}

// Third-party datasource project id, this parameter can be ignored.
func (o DatasourceOutput) DataOriginDatasourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringPtrOutput { return v.DataOriginDatasourceId }).(pulumi.StringPtrOutput)
}

// Third-party datasource project id, this parameter can be ignored.
func (o DatasourceOutput) DataOriginProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringPtrOutput { return v.DataOriginProjectId }).(pulumi.StringPtrOutput)
}

// Host.
func (o DatasourceOutput) DbHost() pulumi.StringOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringOutput { return v.DbHost }).(pulumi.StringOutput)
}

// Database name.
func (o DatasourceOutput) DbName() pulumi.StringOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringOutput { return v.DbName }).(pulumi.StringOutput)
}

// Port.
func (o DatasourceOutput) DbPort() pulumi.IntOutput {
	return o.ApplyT(func(v *Datasource) pulumi.IntOutput { return v.DbPort }).(pulumi.IntOutput)
}

// Password.
func (o DatasourceOutput) DbPwd() pulumi.StringOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringOutput { return v.DbPwd }).(pulumi.StringOutput)
}

// `MYSQL`, `MSSQL`, `POSTGRE`, `ORACLE`, `CLICKHOUSE`, `TIDB`, `HIVE`, `PRESTO`.
func (o DatasourceOutput) DbType() pulumi.StringOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringOutput { return v.DbType }).(pulumi.StringOutput)
}

// User name.
func (o DatasourceOutput) DbUser() pulumi.StringOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringOutput { return v.DbUser }).(pulumi.StringOutput)
}

// Project id.
func (o DatasourceOutput) ProjectId() pulumi.IntOutput {
	return o.ApplyT(func(v *Datasource) pulumi.IntOutput { return v.ProjectId }).(pulumi.IntOutput)
}

// Own or Cloud, default: `Own`.
func (o DatasourceOutput) ServiceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringPtrOutput { return v.ServiceType }).(pulumi.StringPtrOutput)
}

// Datasource name in BI.
func (o DatasourceOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringOutput { return v.SourceName }).(pulumi.StringOutput)
}

// Tencent cloud private network unified identity.
func (o DatasourceOutput) UniqVpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringPtrOutput { return v.UniqVpcId }).(pulumi.StringPtrOutput)
}

// Tencent cloud private network identity.
func (o DatasourceOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Datasource) pulumi.StringPtrOutput { return v.VpcId }).(pulumi.StringPtrOutput)
}

type DatasourceArrayOutput struct{ *pulumi.OutputState }

func (DatasourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Datasource)(nil)).Elem()
}

func (o DatasourceArrayOutput) ToDatasourceArrayOutput() DatasourceArrayOutput {
	return o
}

func (o DatasourceArrayOutput) ToDatasourceArrayOutputWithContext(ctx context.Context) DatasourceArrayOutput {
	return o
}

func (o DatasourceArrayOutput) Index(i pulumi.IntInput) DatasourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Datasource {
		return vs[0].([]*Datasource)[vs[1].(int)]
	}).(DatasourceOutput)
}

type DatasourceMapOutput struct{ *pulumi.OutputState }

func (DatasourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Datasource)(nil)).Elem()
}

func (o DatasourceMapOutput) ToDatasourceMapOutput() DatasourceMapOutput {
	return o
}

func (o DatasourceMapOutput) ToDatasourceMapOutputWithContext(ctx context.Context) DatasourceMapOutput {
	return o
}

func (o DatasourceMapOutput) MapIndex(k pulumi.StringInput) DatasourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Datasource {
		return vs[0].(map[string]*Datasource)[vs[1].(string)]
	}).(DatasourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatasourceInput)(nil)).Elem(), &Datasource{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasourceArrayInput)(nil)).Elem(), DatasourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatasourceMapInput)(nil)).Elem(), DatasourceMap{})
	pulumi.RegisterOutputType(DatasourceOutput{})
	pulumi.RegisterOutputType(DatasourceArrayOutput{})
	pulumi.RegisterOutputType(DatasourceMapOutput{})
}
