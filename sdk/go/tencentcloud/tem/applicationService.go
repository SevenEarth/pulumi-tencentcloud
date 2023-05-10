// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tem

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a tem applicationService
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Tem"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tem"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Tem.NewApplicationService(ctx, "applicationService", &Tem.ApplicationServiceArgs{
// 			ApplicationId: pulumi.String("app-jrl3346j"),
// 			EnvironmentId: pulumi.String("en-dpxyydl5"),
// 			Service: &tem.ApplicationServiceServiceArgs{
// 				PortMappingItemLists: tem.ApplicationServiceServicePortMappingItemListArray{
// 					&tem.ApplicationServiceServicePortMappingItemListArgs{
// 						Port:       pulumi.Int(80),
// 						Protocol:   pulumi.String("TCP"),
// 						TargetPort: pulumi.Int(80),
// 					},
// 				},
// 				ServiceName: pulumi.String("test0-1"),
// 				Type:        pulumi.String("CLUSTER"),
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
// tem application_service can be imported using the environmentId#applicationId#serviceName, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Tem/applicationService:ApplicationService application_service en-dpxyydl5#app-jrl3346j#test0-1
// ```
type ApplicationService struct {
	pulumi.CustomResourceState

	// application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// environment ID.
	EnvironmentId pulumi.StringOutput `pulumi:"environmentId"`
	// service detail list.
	Service ApplicationServiceServicePtrOutput `pulumi:"service"`
}

// NewApplicationService registers a new resource with the given unique name, arguments, and options.
func NewApplicationService(ctx *pulumi.Context,
	name string, args *ApplicationServiceArgs, opts ...pulumi.ResourceOption) (*ApplicationService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ApplicationService
	err := ctx.RegisterResource("tencentcloud:Tem/applicationService:ApplicationService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationService gets an existing ApplicationService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationServiceState, opts ...pulumi.ResourceOption) (*ApplicationService, error) {
	var resource ApplicationService
	err := ctx.ReadResource("tencentcloud:Tem/applicationService:ApplicationService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationService resources.
type applicationServiceState struct {
	// application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// environment ID.
	EnvironmentId *string `pulumi:"environmentId"`
	// service detail list.
	Service *ApplicationServiceService `pulumi:"service"`
}

type ApplicationServiceState struct {
	// application ID.
	ApplicationId pulumi.StringPtrInput
	// environment ID.
	EnvironmentId pulumi.StringPtrInput
	// service detail list.
	Service ApplicationServiceServicePtrInput
}

func (ApplicationServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationServiceState)(nil)).Elem()
}

type applicationServiceArgs struct {
	// application ID.
	ApplicationId string `pulumi:"applicationId"`
	// environment ID.
	EnvironmentId string `pulumi:"environmentId"`
	// service detail list.
	Service *ApplicationServiceService `pulumi:"service"`
}

// The set of arguments for constructing a ApplicationService resource.
type ApplicationServiceArgs struct {
	// application ID.
	ApplicationId pulumi.StringInput
	// environment ID.
	EnvironmentId pulumi.StringInput
	// service detail list.
	Service ApplicationServiceServicePtrInput
}

func (ApplicationServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationServiceArgs)(nil)).Elem()
}

type ApplicationServiceInput interface {
	pulumi.Input

	ToApplicationServiceOutput() ApplicationServiceOutput
	ToApplicationServiceOutputWithContext(ctx context.Context) ApplicationServiceOutput
}

func (*ApplicationService) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationService)(nil)).Elem()
}

func (i *ApplicationService) ToApplicationServiceOutput() ApplicationServiceOutput {
	return i.ToApplicationServiceOutputWithContext(context.Background())
}

func (i *ApplicationService) ToApplicationServiceOutputWithContext(ctx context.Context) ApplicationServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationServiceOutput)
}

// ApplicationServiceArrayInput is an input type that accepts ApplicationServiceArray and ApplicationServiceArrayOutput values.
// You can construct a concrete instance of `ApplicationServiceArrayInput` via:
//
//          ApplicationServiceArray{ ApplicationServiceArgs{...} }
type ApplicationServiceArrayInput interface {
	pulumi.Input

	ToApplicationServiceArrayOutput() ApplicationServiceArrayOutput
	ToApplicationServiceArrayOutputWithContext(context.Context) ApplicationServiceArrayOutput
}

type ApplicationServiceArray []ApplicationServiceInput

func (ApplicationServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationService)(nil)).Elem()
}

func (i ApplicationServiceArray) ToApplicationServiceArrayOutput() ApplicationServiceArrayOutput {
	return i.ToApplicationServiceArrayOutputWithContext(context.Background())
}

func (i ApplicationServiceArray) ToApplicationServiceArrayOutputWithContext(ctx context.Context) ApplicationServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationServiceArrayOutput)
}

// ApplicationServiceMapInput is an input type that accepts ApplicationServiceMap and ApplicationServiceMapOutput values.
// You can construct a concrete instance of `ApplicationServiceMapInput` via:
//
//          ApplicationServiceMap{ "key": ApplicationServiceArgs{...} }
type ApplicationServiceMapInput interface {
	pulumi.Input

	ToApplicationServiceMapOutput() ApplicationServiceMapOutput
	ToApplicationServiceMapOutputWithContext(context.Context) ApplicationServiceMapOutput
}

type ApplicationServiceMap map[string]ApplicationServiceInput

func (ApplicationServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationService)(nil)).Elem()
}

func (i ApplicationServiceMap) ToApplicationServiceMapOutput() ApplicationServiceMapOutput {
	return i.ToApplicationServiceMapOutputWithContext(context.Background())
}

func (i ApplicationServiceMap) ToApplicationServiceMapOutputWithContext(ctx context.Context) ApplicationServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationServiceMapOutput)
}

type ApplicationServiceOutput struct{ *pulumi.OutputState }

func (ApplicationServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationService)(nil)).Elem()
}

func (o ApplicationServiceOutput) ToApplicationServiceOutput() ApplicationServiceOutput {
	return o
}

func (o ApplicationServiceOutput) ToApplicationServiceOutputWithContext(ctx context.Context) ApplicationServiceOutput {
	return o
}

// application ID.
func (o ApplicationServiceOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationService) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

// environment ID.
func (o ApplicationServiceOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationService) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// service detail list.
func (o ApplicationServiceOutput) Service() ApplicationServiceServicePtrOutput {
	return o.ApplyT(func(v *ApplicationService) ApplicationServiceServicePtrOutput { return v.Service }).(ApplicationServiceServicePtrOutput)
}

type ApplicationServiceArrayOutput struct{ *pulumi.OutputState }

func (ApplicationServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationService)(nil)).Elem()
}

func (o ApplicationServiceArrayOutput) ToApplicationServiceArrayOutput() ApplicationServiceArrayOutput {
	return o
}

func (o ApplicationServiceArrayOutput) ToApplicationServiceArrayOutputWithContext(ctx context.Context) ApplicationServiceArrayOutput {
	return o
}

func (o ApplicationServiceArrayOutput) Index(i pulumi.IntInput) ApplicationServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationService {
		return vs[0].([]*ApplicationService)[vs[1].(int)]
	}).(ApplicationServiceOutput)
}

type ApplicationServiceMapOutput struct{ *pulumi.OutputState }

func (ApplicationServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationService)(nil)).Elem()
}

func (o ApplicationServiceMapOutput) ToApplicationServiceMapOutput() ApplicationServiceMapOutput {
	return o
}

func (o ApplicationServiceMapOutput) ToApplicationServiceMapOutputWithContext(ctx context.Context) ApplicationServiceMapOutput {
	return o
}

func (o ApplicationServiceMapOutput) MapIndex(k pulumi.StringInput) ApplicationServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationService {
		return vs[0].(map[string]*ApplicationService)[vs[1].(string)]
	}).(ApplicationServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationServiceInput)(nil)).Elem(), &ApplicationService{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationServiceArrayInput)(nil)).Elem(), ApplicationServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationServiceMapInput)(nil)).Elem(), ApplicationServiceMap{})
	pulumi.RegisterOutputType(ApplicationServiceOutput{})
	pulumi.RegisterOutputType(ApplicationServiceArrayOutput{})
	pulumi.RegisterOutputType(ApplicationServiceMapOutput{})
}
