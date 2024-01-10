// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dcdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a dcdb dbParameters
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Dcdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dcdb.NewDbParameters(ctx, "dbParameters", &Dcdb.DbParametersArgs{
//				InstanceId: pulumi.String(fmt.Sprintf("%v%v", "%", "s")),
//				Params: &dcdb.DbParametersParamsArgs{
//					Param: pulumi.String("max_connections"),
//					Value: pulumi.String("9999"),
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
// dcdb db_parameters can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Dcdb/dbParameters:DbParameters db_parameters instanceId#paramName
//
// ```
type DbParameters struct {
	pulumi.CustomResourceState

	// The ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Parameter list, each element is a combination of Param and Value.
	Params DbParametersParamsOutput `pulumi:"params"`
}

// NewDbParameters registers a new resource with the given unique name, arguments, and options.
func NewDbParameters(ctx *pulumi.Context,
	name string, args *DbParametersArgs, opts ...pulumi.ResourceOption) (*DbParameters, error) {
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
	var resource DbParameters
	err := ctx.RegisterResource("tencentcloud:Dcdb/dbParameters:DbParameters", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDbParameters gets an existing DbParameters resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDbParameters(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DbParametersState, opts ...pulumi.ResourceOption) (*DbParameters, error) {
	var resource DbParameters
	err := ctx.ReadResource("tencentcloud:Dcdb/dbParameters:DbParameters", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DbParameters resources.
type dbParametersState struct {
	// The ID of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Parameter list, each element is a combination of Param and Value.
	Params *DbParametersParams `pulumi:"params"`
}

type DbParametersState struct {
	// The ID of instance.
	InstanceId pulumi.StringPtrInput
	// Parameter list, each element is a combination of Param and Value.
	Params DbParametersParamsPtrInput
}

func (DbParametersState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbParametersState)(nil)).Elem()
}

type dbParametersArgs struct {
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
	// Parameter list, each element is a combination of Param and Value.
	Params DbParametersParams `pulumi:"params"`
}

// The set of arguments for constructing a DbParameters resource.
type DbParametersArgs struct {
	// The ID of instance.
	InstanceId pulumi.StringInput
	// Parameter list, each element is a combination of Param and Value.
	Params DbParametersParamsInput
}

func (DbParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbParametersArgs)(nil)).Elem()
}

type DbParametersInput interface {
	pulumi.Input

	ToDbParametersOutput() DbParametersOutput
	ToDbParametersOutputWithContext(ctx context.Context) DbParametersOutput
}

func (*DbParameters) ElementType() reflect.Type {
	return reflect.TypeOf((**DbParameters)(nil)).Elem()
}

func (i *DbParameters) ToDbParametersOutput() DbParametersOutput {
	return i.ToDbParametersOutputWithContext(context.Background())
}

func (i *DbParameters) ToDbParametersOutputWithContext(ctx context.Context) DbParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbParametersOutput)
}

// DbParametersArrayInput is an input type that accepts DbParametersArray and DbParametersArrayOutput values.
// You can construct a concrete instance of `DbParametersArrayInput` via:
//
//	DbParametersArray{ DbParametersArgs{...} }
type DbParametersArrayInput interface {
	pulumi.Input

	ToDbParametersArrayOutput() DbParametersArrayOutput
	ToDbParametersArrayOutputWithContext(context.Context) DbParametersArrayOutput
}

type DbParametersArray []DbParametersInput

func (DbParametersArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbParameters)(nil)).Elem()
}

func (i DbParametersArray) ToDbParametersArrayOutput() DbParametersArrayOutput {
	return i.ToDbParametersArrayOutputWithContext(context.Background())
}

func (i DbParametersArray) ToDbParametersArrayOutputWithContext(ctx context.Context) DbParametersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbParametersArrayOutput)
}

// DbParametersMapInput is an input type that accepts DbParametersMap and DbParametersMapOutput values.
// You can construct a concrete instance of `DbParametersMapInput` via:
//
//	DbParametersMap{ "key": DbParametersArgs{...} }
type DbParametersMapInput interface {
	pulumi.Input

	ToDbParametersMapOutput() DbParametersMapOutput
	ToDbParametersMapOutputWithContext(context.Context) DbParametersMapOutput
}

type DbParametersMap map[string]DbParametersInput

func (DbParametersMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbParameters)(nil)).Elem()
}

func (i DbParametersMap) ToDbParametersMapOutput() DbParametersMapOutput {
	return i.ToDbParametersMapOutputWithContext(context.Background())
}

func (i DbParametersMap) ToDbParametersMapOutputWithContext(ctx context.Context) DbParametersMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DbParametersMapOutput)
}

type DbParametersOutput struct{ *pulumi.OutputState }

func (DbParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DbParameters)(nil)).Elem()
}

func (o DbParametersOutput) ToDbParametersOutput() DbParametersOutput {
	return o
}

func (o DbParametersOutput) ToDbParametersOutputWithContext(ctx context.Context) DbParametersOutput {
	return o
}

// The ID of instance.
func (o DbParametersOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DbParameters) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Parameter list, each element is a combination of Param and Value.
func (o DbParametersOutput) Params() DbParametersParamsOutput {
	return o.ApplyT(func(v *DbParameters) DbParametersParamsOutput { return v.Params }).(DbParametersParamsOutput)
}

type DbParametersArrayOutput struct{ *pulumi.OutputState }

func (DbParametersArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DbParameters)(nil)).Elem()
}

func (o DbParametersArrayOutput) ToDbParametersArrayOutput() DbParametersArrayOutput {
	return o
}

func (o DbParametersArrayOutput) ToDbParametersArrayOutputWithContext(ctx context.Context) DbParametersArrayOutput {
	return o
}

func (o DbParametersArrayOutput) Index(i pulumi.IntInput) DbParametersOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DbParameters {
		return vs[0].([]*DbParameters)[vs[1].(int)]
	}).(DbParametersOutput)
}

type DbParametersMapOutput struct{ *pulumi.OutputState }

func (DbParametersMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DbParameters)(nil)).Elem()
}

func (o DbParametersMapOutput) ToDbParametersMapOutput() DbParametersMapOutput {
	return o
}

func (o DbParametersMapOutput) ToDbParametersMapOutputWithContext(ctx context.Context) DbParametersMapOutput {
	return o
}

func (o DbParametersMapOutput) MapIndex(k pulumi.StringInput) DbParametersOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DbParameters {
		return vs[0].(map[string]*DbParameters)[vs[1].(string)]
	}).(DbParametersOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DbParametersInput)(nil)).Elem(), &DbParameters{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbParametersArrayInput)(nil)).Elem(), DbParametersArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DbParametersMapInput)(nil)).Elem(), DbParametersMap{})
	pulumi.RegisterOutputType(DbParametersOutput{})
	pulumi.RegisterOutputType(DbParametersArrayOutput{})
	pulumi.RegisterOutputType(DbParametersMapOutput{})
}
