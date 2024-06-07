// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Use this resource to create API gateway service release.
//
// ## Import
//
// API gateway service release can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:ApiGateway/serviceRelease:ServiceRelease service service-jjt3fs3s#release#20201015121916d85fb161-eaec-4dda-a7e0-659aa5f401be
// ```
type ServiceRelease struct {
	pulumi.CustomResourceState

	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	EnvironmentName pulumi.StringOutput `pulumi:"environmentName"`
	// This release description of the API gateway service.
	ReleaseDesc pulumi.StringOutput `pulumi:"releaseDesc"`
	// The release version.
	ReleaseVersion pulumi.StringOutput `pulumi:"releaseVersion"`
	// ID of API gateway service.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
}

// NewServiceRelease registers a new resource with the given unique name, arguments, and options.
func NewServiceRelease(ctx *pulumi.Context,
	name string, args *ServiceReleaseArgs, opts ...pulumi.ResourceOption) (*ServiceRelease, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentName == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentName'")
	}
	if args.ReleaseDesc == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseDesc'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceRelease
	err := ctx.RegisterResource("tencentcloud:ApiGateway/serviceRelease:ServiceRelease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceRelease gets an existing ServiceRelease resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceReleaseState, opts ...pulumi.ResourceOption) (*ServiceRelease, error) {
	var resource ServiceRelease
	err := ctx.ReadResource("tencentcloud:ApiGateway/serviceRelease:ServiceRelease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceRelease resources.
type serviceReleaseState struct {
	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	EnvironmentName *string `pulumi:"environmentName"`
	// This release description of the API gateway service.
	ReleaseDesc *string `pulumi:"releaseDesc"`
	// The release version.
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// ID of API gateway service.
	ServiceId *string `pulumi:"serviceId"`
}

type ServiceReleaseState struct {
	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	EnvironmentName pulumi.StringPtrInput
	// This release description of the API gateway service.
	ReleaseDesc pulumi.StringPtrInput
	// The release version.
	ReleaseVersion pulumi.StringPtrInput
	// ID of API gateway service.
	ServiceId pulumi.StringPtrInput
}

func (ServiceReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceReleaseState)(nil)).Elem()
}

type serviceReleaseArgs struct {
	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	EnvironmentName string `pulumi:"environmentName"`
	// This release description of the API gateway service.
	ReleaseDesc string `pulumi:"releaseDesc"`
	// The release version.
	ReleaseVersion *string `pulumi:"releaseVersion"`
	// ID of API gateway service.
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a ServiceRelease resource.
type ServiceReleaseArgs struct {
	// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
	EnvironmentName pulumi.StringInput
	// This release description of the API gateway service.
	ReleaseDesc pulumi.StringInput
	// The release version.
	ReleaseVersion pulumi.StringPtrInput
	// ID of API gateway service.
	ServiceId pulumi.StringInput
}

func (ServiceReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceReleaseArgs)(nil)).Elem()
}

type ServiceReleaseInput interface {
	pulumi.Input

	ToServiceReleaseOutput() ServiceReleaseOutput
	ToServiceReleaseOutputWithContext(ctx context.Context) ServiceReleaseOutput
}

func (*ServiceRelease) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRelease)(nil)).Elem()
}

func (i *ServiceRelease) ToServiceReleaseOutput() ServiceReleaseOutput {
	return i.ToServiceReleaseOutputWithContext(context.Background())
}

func (i *ServiceRelease) ToServiceReleaseOutputWithContext(ctx context.Context) ServiceReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceReleaseOutput)
}

// ServiceReleaseArrayInput is an input type that accepts ServiceReleaseArray and ServiceReleaseArrayOutput values.
// You can construct a concrete instance of `ServiceReleaseArrayInput` via:
//
//	ServiceReleaseArray{ ServiceReleaseArgs{...} }
type ServiceReleaseArrayInput interface {
	pulumi.Input

	ToServiceReleaseArrayOutput() ServiceReleaseArrayOutput
	ToServiceReleaseArrayOutputWithContext(context.Context) ServiceReleaseArrayOutput
}

type ServiceReleaseArray []ServiceReleaseInput

func (ServiceReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceRelease)(nil)).Elem()
}

func (i ServiceReleaseArray) ToServiceReleaseArrayOutput() ServiceReleaseArrayOutput {
	return i.ToServiceReleaseArrayOutputWithContext(context.Background())
}

func (i ServiceReleaseArray) ToServiceReleaseArrayOutputWithContext(ctx context.Context) ServiceReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceReleaseArrayOutput)
}

// ServiceReleaseMapInput is an input type that accepts ServiceReleaseMap and ServiceReleaseMapOutput values.
// You can construct a concrete instance of `ServiceReleaseMapInput` via:
//
//	ServiceReleaseMap{ "key": ServiceReleaseArgs{...} }
type ServiceReleaseMapInput interface {
	pulumi.Input

	ToServiceReleaseMapOutput() ServiceReleaseMapOutput
	ToServiceReleaseMapOutputWithContext(context.Context) ServiceReleaseMapOutput
}

type ServiceReleaseMap map[string]ServiceReleaseInput

func (ServiceReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceRelease)(nil)).Elem()
}

func (i ServiceReleaseMap) ToServiceReleaseMapOutput() ServiceReleaseMapOutput {
	return i.ToServiceReleaseMapOutputWithContext(context.Background())
}

func (i ServiceReleaseMap) ToServiceReleaseMapOutputWithContext(ctx context.Context) ServiceReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceReleaseMapOutput)
}

type ServiceReleaseOutput struct{ *pulumi.OutputState }

func (ServiceReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceRelease)(nil)).Elem()
}

func (o ServiceReleaseOutput) ToServiceReleaseOutput() ServiceReleaseOutput {
	return o
}

func (o ServiceReleaseOutput) ToServiceReleaseOutputWithContext(ctx context.Context) ServiceReleaseOutput {
	return o
}

// API gateway service environment name to be released. Valid values: `test`, `prepub`, `release`.
func (o ServiceReleaseOutput) EnvironmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRelease) pulumi.StringOutput { return v.EnvironmentName }).(pulumi.StringOutput)
}

// This release description of the API gateway service.
func (o ServiceReleaseOutput) ReleaseDesc() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRelease) pulumi.StringOutput { return v.ReleaseDesc }).(pulumi.StringOutput)
}

// The release version.
func (o ServiceReleaseOutput) ReleaseVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRelease) pulumi.StringOutput { return v.ReleaseVersion }).(pulumi.StringOutput)
}

// ID of API gateway service.
func (o ServiceReleaseOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceRelease) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

type ServiceReleaseArrayOutput struct{ *pulumi.OutputState }

func (ServiceReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceRelease)(nil)).Elem()
}

func (o ServiceReleaseArrayOutput) ToServiceReleaseArrayOutput() ServiceReleaseArrayOutput {
	return o
}

func (o ServiceReleaseArrayOutput) ToServiceReleaseArrayOutputWithContext(ctx context.Context) ServiceReleaseArrayOutput {
	return o
}

func (o ServiceReleaseArrayOutput) Index(i pulumi.IntInput) ServiceReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceRelease {
		return vs[0].([]*ServiceRelease)[vs[1].(int)]
	}).(ServiceReleaseOutput)
}

type ServiceReleaseMapOutput struct{ *pulumi.OutputState }

func (ServiceReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceRelease)(nil)).Elem()
}

func (o ServiceReleaseMapOutput) ToServiceReleaseMapOutput() ServiceReleaseMapOutput {
	return o
}

func (o ServiceReleaseMapOutput) ToServiceReleaseMapOutputWithContext(ctx context.Context) ServiceReleaseMapOutput {
	return o
}

func (o ServiceReleaseMapOutput) MapIndex(k pulumi.StringInput) ServiceReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceRelease {
		return vs[0].(map[string]*ServiceRelease)[vs[1].(string)]
	}).(ServiceReleaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceReleaseInput)(nil)).Elem(), &ServiceRelease{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceReleaseArrayInput)(nil)).Elem(), ServiceReleaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceReleaseMapInput)(nil)).Elem(), ServiceReleaseMap{})
	pulumi.RegisterOutputType(ServiceReleaseOutput{})
	pulumi.RegisterOutputType(ServiceReleaseArrayOutput{})
	pulumi.RegisterOutputType(ServiceReleaseMapOutput{})
}
