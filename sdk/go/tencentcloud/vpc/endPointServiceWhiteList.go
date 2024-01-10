// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc endPointServiceWhiteList
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.NewEndPointServiceWhiteList(ctx, "endPointServiceWhiteList", &Vpc.EndPointServiceWhiteListArgs{
//				Description:       pulumi.String("terraform for test"),
//				EndPointServiceId: pulumi.String("vpcsvc-69y13tdb"),
//				UserUin:           pulumi.String("100020512675"),
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
// vpc end_point_service_white_list can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList end_point_service_white_list end_point_service_white_list_id
//
// ```
type EndPointServiceWhiteList struct {
	pulumi.CustomResourceState

	// Create Time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of white list.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// ID of endpoint service.
	EndPointServiceId pulumi.StringOutput `pulumi:"endPointServiceId"`
	// APPID.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// UIN.
	UserUin pulumi.StringOutput `pulumi:"userUin"`
}

// NewEndPointServiceWhiteList registers a new resource with the given unique name, arguments, and options.
func NewEndPointServiceWhiteList(ctx *pulumi.Context,
	name string, args *EndPointServiceWhiteListArgs, opts ...pulumi.ResourceOption) (*EndPointServiceWhiteList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndPointServiceId == nil {
		return nil, errors.New("invalid value for required argument 'EndPointServiceId'")
	}
	if args.UserUin == nil {
		return nil, errors.New("invalid value for required argument 'UserUin'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EndPointServiceWhiteList
	err := ctx.RegisterResource("tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndPointServiceWhiteList gets an existing EndPointServiceWhiteList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndPointServiceWhiteList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndPointServiceWhiteListState, opts ...pulumi.ResourceOption) (*EndPointServiceWhiteList, error) {
	var resource EndPointServiceWhiteList
	err := ctx.ReadResource("tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndPointServiceWhiteList resources.
type endPointServiceWhiteListState struct {
	// Create Time.
	CreateTime *string `pulumi:"createTime"`
	// Description of white list.
	Description *string `pulumi:"description"`
	// ID of endpoint service.
	EndPointServiceId *string `pulumi:"endPointServiceId"`
	// APPID.
	Owner *string `pulumi:"owner"`
	// UIN.
	UserUin *string `pulumi:"userUin"`
}

type EndPointServiceWhiteListState struct {
	// Create Time.
	CreateTime pulumi.StringPtrInput
	// Description of white list.
	Description pulumi.StringPtrInput
	// ID of endpoint service.
	EndPointServiceId pulumi.StringPtrInput
	// APPID.
	Owner pulumi.StringPtrInput
	// UIN.
	UserUin pulumi.StringPtrInput
}

func (EndPointServiceWhiteListState) ElementType() reflect.Type {
	return reflect.TypeOf((*endPointServiceWhiteListState)(nil)).Elem()
}

type endPointServiceWhiteListArgs struct {
	// Description of white list.
	Description *string `pulumi:"description"`
	// ID of endpoint service.
	EndPointServiceId string `pulumi:"endPointServiceId"`
	// UIN.
	UserUin string `pulumi:"userUin"`
}

// The set of arguments for constructing a EndPointServiceWhiteList resource.
type EndPointServiceWhiteListArgs struct {
	// Description of white list.
	Description pulumi.StringPtrInput
	// ID of endpoint service.
	EndPointServiceId pulumi.StringInput
	// UIN.
	UserUin pulumi.StringInput
}

func (EndPointServiceWhiteListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endPointServiceWhiteListArgs)(nil)).Elem()
}

type EndPointServiceWhiteListInput interface {
	pulumi.Input

	ToEndPointServiceWhiteListOutput() EndPointServiceWhiteListOutput
	ToEndPointServiceWhiteListOutputWithContext(ctx context.Context) EndPointServiceWhiteListOutput
}

func (*EndPointServiceWhiteList) ElementType() reflect.Type {
	return reflect.TypeOf((**EndPointServiceWhiteList)(nil)).Elem()
}

func (i *EndPointServiceWhiteList) ToEndPointServiceWhiteListOutput() EndPointServiceWhiteListOutput {
	return i.ToEndPointServiceWhiteListOutputWithContext(context.Background())
}

func (i *EndPointServiceWhiteList) ToEndPointServiceWhiteListOutputWithContext(ctx context.Context) EndPointServiceWhiteListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointServiceWhiteListOutput)
}

// EndPointServiceWhiteListArrayInput is an input type that accepts EndPointServiceWhiteListArray and EndPointServiceWhiteListArrayOutput values.
// You can construct a concrete instance of `EndPointServiceWhiteListArrayInput` via:
//
//	EndPointServiceWhiteListArray{ EndPointServiceWhiteListArgs{...} }
type EndPointServiceWhiteListArrayInput interface {
	pulumi.Input

	ToEndPointServiceWhiteListArrayOutput() EndPointServiceWhiteListArrayOutput
	ToEndPointServiceWhiteListArrayOutputWithContext(context.Context) EndPointServiceWhiteListArrayOutput
}

type EndPointServiceWhiteListArray []EndPointServiceWhiteListInput

func (EndPointServiceWhiteListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndPointServiceWhiteList)(nil)).Elem()
}

func (i EndPointServiceWhiteListArray) ToEndPointServiceWhiteListArrayOutput() EndPointServiceWhiteListArrayOutput {
	return i.ToEndPointServiceWhiteListArrayOutputWithContext(context.Background())
}

func (i EndPointServiceWhiteListArray) ToEndPointServiceWhiteListArrayOutputWithContext(ctx context.Context) EndPointServiceWhiteListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointServiceWhiteListArrayOutput)
}

// EndPointServiceWhiteListMapInput is an input type that accepts EndPointServiceWhiteListMap and EndPointServiceWhiteListMapOutput values.
// You can construct a concrete instance of `EndPointServiceWhiteListMapInput` via:
//
//	EndPointServiceWhiteListMap{ "key": EndPointServiceWhiteListArgs{...} }
type EndPointServiceWhiteListMapInput interface {
	pulumi.Input

	ToEndPointServiceWhiteListMapOutput() EndPointServiceWhiteListMapOutput
	ToEndPointServiceWhiteListMapOutputWithContext(context.Context) EndPointServiceWhiteListMapOutput
}

type EndPointServiceWhiteListMap map[string]EndPointServiceWhiteListInput

func (EndPointServiceWhiteListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndPointServiceWhiteList)(nil)).Elem()
}

func (i EndPointServiceWhiteListMap) ToEndPointServiceWhiteListMapOutput() EndPointServiceWhiteListMapOutput {
	return i.ToEndPointServiceWhiteListMapOutputWithContext(context.Background())
}

func (i EndPointServiceWhiteListMap) ToEndPointServiceWhiteListMapOutputWithContext(ctx context.Context) EndPointServiceWhiteListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointServiceWhiteListMapOutput)
}

type EndPointServiceWhiteListOutput struct{ *pulumi.OutputState }

func (EndPointServiceWhiteListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndPointServiceWhiteList)(nil)).Elem()
}

func (o EndPointServiceWhiteListOutput) ToEndPointServiceWhiteListOutput() EndPointServiceWhiteListOutput {
	return o
}

func (o EndPointServiceWhiteListOutput) ToEndPointServiceWhiteListOutputWithContext(ctx context.Context) EndPointServiceWhiteListOutput {
	return o
}

// Create Time.
func (o EndPointServiceWhiteListOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointServiceWhiteList) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of white list.
func (o EndPointServiceWhiteListOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndPointServiceWhiteList) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// ID of endpoint service.
func (o EndPointServiceWhiteListOutput) EndPointServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointServiceWhiteList) pulumi.StringOutput { return v.EndPointServiceId }).(pulumi.StringOutput)
}

// APPID.
func (o EndPointServiceWhiteListOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointServiceWhiteList) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// UIN.
func (o EndPointServiceWhiteListOutput) UserUin() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointServiceWhiteList) pulumi.StringOutput { return v.UserUin }).(pulumi.StringOutput)
}

type EndPointServiceWhiteListArrayOutput struct{ *pulumi.OutputState }

func (EndPointServiceWhiteListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndPointServiceWhiteList)(nil)).Elem()
}

func (o EndPointServiceWhiteListArrayOutput) ToEndPointServiceWhiteListArrayOutput() EndPointServiceWhiteListArrayOutput {
	return o
}

func (o EndPointServiceWhiteListArrayOutput) ToEndPointServiceWhiteListArrayOutputWithContext(ctx context.Context) EndPointServiceWhiteListArrayOutput {
	return o
}

func (o EndPointServiceWhiteListArrayOutput) Index(i pulumi.IntInput) EndPointServiceWhiteListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndPointServiceWhiteList {
		return vs[0].([]*EndPointServiceWhiteList)[vs[1].(int)]
	}).(EndPointServiceWhiteListOutput)
}

type EndPointServiceWhiteListMapOutput struct{ *pulumi.OutputState }

func (EndPointServiceWhiteListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndPointServiceWhiteList)(nil)).Elem()
}

func (o EndPointServiceWhiteListMapOutput) ToEndPointServiceWhiteListMapOutput() EndPointServiceWhiteListMapOutput {
	return o
}

func (o EndPointServiceWhiteListMapOutput) ToEndPointServiceWhiteListMapOutputWithContext(ctx context.Context) EndPointServiceWhiteListMapOutput {
	return o
}

func (o EndPointServiceWhiteListMapOutput) MapIndex(k pulumi.StringInput) EndPointServiceWhiteListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndPointServiceWhiteList {
		return vs[0].(map[string]*EndPointServiceWhiteList)[vs[1].(string)]
	}).(EndPointServiceWhiteListOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndPointServiceWhiteListInput)(nil)).Elem(), &EndPointServiceWhiteList{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndPointServiceWhiteListArrayInput)(nil)).Elem(), EndPointServiceWhiteListArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndPointServiceWhiteListMapInput)(nil)).Elem(), EndPointServiceWhiteListMap{})
	pulumi.RegisterOutputType(EndPointServiceWhiteListOutput{})
	pulumi.RegisterOutputType(EndPointServiceWhiteListArrayOutput{})
	pulumi.RegisterOutputType(EndPointServiceWhiteListMapOutput{})
}
