// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a vpc endPointService
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Vpc"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Vpc.NewEndPointService(ctx, "endPointService", &Vpc.EndPointServiceArgs{
// 			AutoAcceptFlag:      pulumi.Bool(false),
// 			EndPointServiceName: pulumi.String("terraform-endpoint-service"),
// 			ServiceInstanceId:   pulumi.String("lb-o5f6x7ke"),
// 			ServiceType:         pulumi.String("CLB"),
// 			VpcId:               pulumi.String("vpc-391sv4w3"),
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
// vpc end_point_service can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Vpc/endPointService:EndPointService end_point_service end_point_service_id
// ```
type EndPointService struct {
	pulumi.CustomResourceState

	// Whether to automatically accept.
	AutoAcceptFlag pulumi.BoolOutput `pulumi:"autoAcceptFlag"`
	// Create Time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Count of end point.
	EndPointCount pulumi.IntOutput `pulumi:"endPointCount"`
	// Name of end point service.
	EndPointServiceName pulumi.StringOutput `pulumi:"endPointServiceName"`
	// Id of service instance, like lb-xxx.
	ServiceInstanceId pulumi.StringOutput `pulumi:"serviceInstanceId"`
	// APPID.
	ServiceOwner pulumi.StringOutput `pulumi:"serviceOwner"`
	// Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// VIP of backend service.
	ServiceVip pulumi.StringOutput `pulumi:"serviceVip"`
	// ID of vpc instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewEndPointService registers a new resource with the given unique name, arguments, and options.
func NewEndPointService(ctx *pulumi.Context,
	name string, args *EndPointServiceArgs, opts ...pulumi.ResourceOption) (*EndPointService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutoAcceptFlag == nil {
		return nil, errors.New("invalid value for required argument 'AutoAcceptFlag'")
	}
	if args.EndPointServiceName == nil {
		return nil, errors.New("invalid value for required argument 'EndPointServiceName'")
	}
	if args.ServiceInstanceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceInstanceId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EndPointService
	err := ctx.RegisterResource("tencentcloud:Vpc/endPointService:EndPointService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndPointService gets an existing EndPointService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndPointService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndPointServiceState, opts ...pulumi.ResourceOption) (*EndPointService, error) {
	var resource EndPointService
	err := ctx.ReadResource("tencentcloud:Vpc/endPointService:EndPointService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndPointService resources.
type endPointServiceState struct {
	// Whether to automatically accept.
	AutoAcceptFlag *bool `pulumi:"autoAcceptFlag"`
	// Create Time.
	CreateTime *string `pulumi:"createTime"`
	// Count of end point.
	EndPointCount *int `pulumi:"endPointCount"`
	// Name of end point service.
	EndPointServiceName *string `pulumi:"endPointServiceName"`
	// Id of service instance, like lb-xxx.
	ServiceInstanceId *string `pulumi:"serviceInstanceId"`
	// APPID.
	ServiceOwner *string `pulumi:"serviceOwner"`
	// Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
	ServiceType *string `pulumi:"serviceType"`
	// VIP of backend service.
	ServiceVip *string `pulumi:"serviceVip"`
	// ID of vpc instance.
	VpcId *string `pulumi:"vpcId"`
}

type EndPointServiceState struct {
	// Whether to automatically accept.
	AutoAcceptFlag pulumi.BoolPtrInput
	// Create Time.
	CreateTime pulumi.StringPtrInput
	// Count of end point.
	EndPointCount pulumi.IntPtrInput
	// Name of end point service.
	EndPointServiceName pulumi.StringPtrInput
	// Id of service instance, like lb-xxx.
	ServiceInstanceId pulumi.StringPtrInput
	// APPID.
	ServiceOwner pulumi.StringPtrInput
	// Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
	ServiceType pulumi.StringPtrInput
	// VIP of backend service.
	ServiceVip pulumi.StringPtrInput
	// ID of vpc instance.
	VpcId pulumi.StringPtrInput
}

func (EndPointServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*endPointServiceState)(nil)).Elem()
}

type endPointServiceArgs struct {
	// Whether to automatically accept.
	AutoAcceptFlag bool `pulumi:"autoAcceptFlag"`
	// Name of end point service.
	EndPointServiceName string `pulumi:"endPointServiceName"`
	// Id of service instance, like lb-xxx.
	ServiceInstanceId string `pulumi:"serviceInstanceId"`
	// Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
	ServiceType *string `pulumi:"serviceType"`
	// ID of vpc instance.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a EndPointService resource.
type EndPointServiceArgs struct {
	// Whether to automatically accept.
	AutoAcceptFlag pulumi.BoolInput
	// Name of end point service.
	EndPointServiceName pulumi.StringInput
	// Id of service instance, like lb-xxx.
	ServiceInstanceId pulumi.StringInput
	// Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
	ServiceType pulumi.StringPtrInput
	// ID of vpc instance.
	VpcId pulumi.StringInput
}

func (EndPointServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endPointServiceArgs)(nil)).Elem()
}

type EndPointServiceInput interface {
	pulumi.Input

	ToEndPointServiceOutput() EndPointServiceOutput
	ToEndPointServiceOutputWithContext(ctx context.Context) EndPointServiceOutput
}

func (*EndPointService) ElementType() reflect.Type {
	return reflect.TypeOf((**EndPointService)(nil)).Elem()
}

func (i *EndPointService) ToEndPointServiceOutput() EndPointServiceOutput {
	return i.ToEndPointServiceOutputWithContext(context.Background())
}

func (i *EndPointService) ToEndPointServiceOutputWithContext(ctx context.Context) EndPointServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointServiceOutput)
}

// EndPointServiceArrayInput is an input type that accepts EndPointServiceArray and EndPointServiceArrayOutput values.
// You can construct a concrete instance of `EndPointServiceArrayInput` via:
//
//          EndPointServiceArray{ EndPointServiceArgs{...} }
type EndPointServiceArrayInput interface {
	pulumi.Input

	ToEndPointServiceArrayOutput() EndPointServiceArrayOutput
	ToEndPointServiceArrayOutputWithContext(context.Context) EndPointServiceArrayOutput
}

type EndPointServiceArray []EndPointServiceInput

func (EndPointServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndPointService)(nil)).Elem()
}

func (i EndPointServiceArray) ToEndPointServiceArrayOutput() EndPointServiceArrayOutput {
	return i.ToEndPointServiceArrayOutputWithContext(context.Background())
}

func (i EndPointServiceArray) ToEndPointServiceArrayOutputWithContext(ctx context.Context) EndPointServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointServiceArrayOutput)
}

// EndPointServiceMapInput is an input type that accepts EndPointServiceMap and EndPointServiceMapOutput values.
// You can construct a concrete instance of `EndPointServiceMapInput` via:
//
//          EndPointServiceMap{ "key": EndPointServiceArgs{...} }
type EndPointServiceMapInput interface {
	pulumi.Input

	ToEndPointServiceMapOutput() EndPointServiceMapOutput
	ToEndPointServiceMapOutputWithContext(context.Context) EndPointServiceMapOutput
}

type EndPointServiceMap map[string]EndPointServiceInput

func (EndPointServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndPointService)(nil)).Elem()
}

func (i EndPointServiceMap) ToEndPointServiceMapOutput() EndPointServiceMapOutput {
	return i.ToEndPointServiceMapOutputWithContext(context.Background())
}

func (i EndPointServiceMap) ToEndPointServiceMapOutputWithContext(ctx context.Context) EndPointServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndPointServiceMapOutput)
}

type EndPointServiceOutput struct{ *pulumi.OutputState }

func (EndPointServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndPointService)(nil)).Elem()
}

func (o EndPointServiceOutput) ToEndPointServiceOutput() EndPointServiceOutput {
	return o
}

func (o EndPointServiceOutput) ToEndPointServiceOutputWithContext(ctx context.Context) EndPointServiceOutput {
	return o
}

// Whether to automatically accept.
func (o EndPointServiceOutput) AutoAcceptFlag() pulumi.BoolOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.BoolOutput { return v.AutoAcceptFlag }).(pulumi.BoolOutput)
}

// Create Time.
func (o EndPointServiceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Count of end point.
func (o EndPointServiceOutput) EndPointCount() pulumi.IntOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.IntOutput { return v.EndPointCount }).(pulumi.IntOutput)
}

// Name of end point service.
func (o EndPointServiceOutput) EndPointServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.StringOutput { return v.EndPointServiceName }).(pulumi.StringOutput)
}

// Id of service instance, like lb-xxx.
func (o EndPointServiceOutput) ServiceInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.StringOutput { return v.ServiceInstanceId }).(pulumi.StringOutput)
}

// APPID.
func (o EndPointServiceOutput) ServiceOwner() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.StringOutput { return v.ServiceOwner }).(pulumi.StringOutput)
}

// Type of service instance, like `CLB`, `CDB`, `CRS`, default is `CLB`.
func (o EndPointServiceOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.StringOutput { return v.ServiceType }).(pulumi.StringOutput)
}

// VIP of backend service.
func (o EndPointServiceOutput) ServiceVip() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.StringOutput { return v.ServiceVip }).(pulumi.StringOutput)
}

// ID of vpc instance.
func (o EndPointServiceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *EndPointService) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type EndPointServiceArrayOutput struct{ *pulumi.OutputState }

func (EndPointServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndPointService)(nil)).Elem()
}

func (o EndPointServiceArrayOutput) ToEndPointServiceArrayOutput() EndPointServiceArrayOutput {
	return o
}

func (o EndPointServiceArrayOutput) ToEndPointServiceArrayOutputWithContext(ctx context.Context) EndPointServiceArrayOutput {
	return o
}

func (o EndPointServiceArrayOutput) Index(i pulumi.IntInput) EndPointServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndPointService {
		return vs[0].([]*EndPointService)[vs[1].(int)]
	}).(EndPointServiceOutput)
}

type EndPointServiceMapOutput struct{ *pulumi.OutputState }

func (EndPointServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndPointService)(nil)).Elem()
}

func (o EndPointServiceMapOutput) ToEndPointServiceMapOutput() EndPointServiceMapOutput {
	return o
}

func (o EndPointServiceMapOutput) ToEndPointServiceMapOutputWithContext(ctx context.Context) EndPointServiceMapOutput {
	return o
}

func (o EndPointServiceMapOutput) MapIndex(k pulumi.StringInput) EndPointServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndPointService {
		return vs[0].(map[string]*EndPointService)[vs[1].(string)]
	}).(EndPointServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndPointServiceInput)(nil)).Elem(), &EndPointService{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndPointServiceArrayInput)(nil)).Elem(), EndPointServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndPointServiceMapInput)(nil)).Elem(), EndPointServiceMap{})
	pulumi.RegisterOutputType(EndPointServiceOutput{})
	pulumi.RegisterOutputType(EndPointServiceArrayOutput{})
	pulumi.RegisterOutputType(EndPointServiceMapOutput{})
}
