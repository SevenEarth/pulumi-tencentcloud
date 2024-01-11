// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package antiddos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a antiddos ddosGeoIpBlockConfig
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Antiddos"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Antiddos"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Antiddos.NewDdosGeoIpBlockConfig(ctx, "ddosGeoIpBlockConfig", &Antiddos.DdosGeoIpBlockConfigArgs{
// 			DdosGeoIpBlockConfig: &antiddos.DdosGeoIpBlockConfigDdosGeoIpBlockConfigArgs{
// 				Action: pulumi.String("drop"),
// 				AreaLists: pulumi.IntArray{
// 					pulumi.Int(100002),
// 				},
// 				RegionType: pulumi.String("customized"),
// 			},
// 			InstanceId: pulumi.String("bgp-xxxxxx"),
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
// antiddos ddos_geo_ip_block_config can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig ddos_geo_ip_block_config ${instanceId}#${configId}
// ```
type DdosGeoIpBlockConfig struct {
	pulumi.CustomResourceState

	// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
	DdosGeoIpBlockConfig DdosGeoIpBlockConfigDdosGeoIpBlockConfigOutput `pulumi:"ddosGeoIpBlockConfig"`
	// InstanceId.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewDdosGeoIpBlockConfig registers a new resource with the given unique name, arguments, and options.
func NewDdosGeoIpBlockConfig(ctx *pulumi.Context,
	name string, args *DdosGeoIpBlockConfigArgs, opts ...pulumi.ResourceOption) (*DdosGeoIpBlockConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DdosGeoIpBlockConfig == nil {
		return nil, errors.New("invalid value for required argument 'DdosGeoIpBlockConfig'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DdosGeoIpBlockConfig
	err := ctx.RegisterResource("tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosGeoIpBlockConfig gets an existing DdosGeoIpBlockConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosGeoIpBlockConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosGeoIpBlockConfigState, opts ...pulumi.ResourceOption) (*DdosGeoIpBlockConfig, error) {
	var resource DdosGeoIpBlockConfig
	err := ctx.ReadResource("tencentcloud:Antiddos/ddosGeoIpBlockConfig:DdosGeoIpBlockConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosGeoIpBlockConfig resources.
type ddosGeoIpBlockConfigState struct {
	// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
	DdosGeoIpBlockConfig *DdosGeoIpBlockConfigDdosGeoIpBlockConfig `pulumi:"ddosGeoIpBlockConfig"`
	// InstanceId.
	InstanceId *string `pulumi:"instanceId"`
}

type DdosGeoIpBlockConfigState struct {
	// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
	DdosGeoIpBlockConfig DdosGeoIpBlockConfigDdosGeoIpBlockConfigPtrInput
	// InstanceId.
	InstanceId pulumi.StringPtrInput
}

func (DdosGeoIpBlockConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosGeoIpBlockConfigState)(nil)).Elem()
}

type ddosGeoIpBlockConfigArgs struct {
	// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
	DdosGeoIpBlockConfig DdosGeoIpBlockConfigDdosGeoIpBlockConfig `pulumi:"ddosGeoIpBlockConfig"`
	// InstanceId.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a DdosGeoIpBlockConfig resource.
type DdosGeoIpBlockConfigArgs struct {
	// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
	DdosGeoIpBlockConfig DdosGeoIpBlockConfigDdosGeoIpBlockConfigInput
	// InstanceId.
	InstanceId pulumi.StringInput
}

func (DdosGeoIpBlockConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosGeoIpBlockConfigArgs)(nil)).Elem()
}

type DdosGeoIpBlockConfigInput interface {
	pulumi.Input

	ToDdosGeoIpBlockConfigOutput() DdosGeoIpBlockConfigOutput
	ToDdosGeoIpBlockConfigOutputWithContext(ctx context.Context) DdosGeoIpBlockConfigOutput
}

func (*DdosGeoIpBlockConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosGeoIpBlockConfig)(nil)).Elem()
}

func (i *DdosGeoIpBlockConfig) ToDdosGeoIpBlockConfigOutput() DdosGeoIpBlockConfigOutput {
	return i.ToDdosGeoIpBlockConfigOutputWithContext(context.Background())
}

func (i *DdosGeoIpBlockConfig) ToDdosGeoIpBlockConfigOutputWithContext(ctx context.Context) DdosGeoIpBlockConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosGeoIpBlockConfigOutput)
}

// DdosGeoIpBlockConfigArrayInput is an input type that accepts DdosGeoIpBlockConfigArray and DdosGeoIpBlockConfigArrayOutput values.
// You can construct a concrete instance of `DdosGeoIpBlockConfigArrayInput` via:
//
//          DdosGeoIpBlockConfigArray{ DdosGeoIpBlockConfigArgs{...} }
type DdosGeoIpBlockConfigArrayInput interface {
	pulumi.Input

	ToDdosGeoIpBlockConfigArrayOutput() DdosGeoIpBlockConfigArrayOutput
	ToDdosGeoIpBlockConfigArrayOutputWithContext(context.Context) DdosGeoIpBlockConfigArrayOutput
}

type DdosGeoIpBlockConfigArray []DdosGeoIpBlockConfigInput

func (DdosGeoIpBlockConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosGeoIpBlockConfig)(nil)).Elem()
}

func (i DdosGeoIpBlockConfigArray) ToDdosGeoIpBlockConfigArrayOutput() DdosGeoIpBlockConfigArrayOutput {
	return i.ToDdosGeoIpBlockConfigArrayOutputWithContext(context.Background())
}

func (i DdosGeoIpBlockConfigArray) ToDdosGeoIpBlockConfigArrayOutputWithContext(ctx context.Context) DdosGeoIpBlockConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosGeoIpBlockConfigArrayOutput)
}

// DdosGeoIpBlockConfigMapInput is an input type that accepts DdosGeoIpBlockConfigMap and DdosGeoIpBlockConfigMapOutput values.
// You can construct a concrete instance of `DdosGeoIpBlockConfigMapInput` via:
//
//          DdosGeoIpBlockConfigMap{ "key": DdosGeoIpBlockConfigArgs{...} }
type DdosGeoIpBlockConfigMapInput interface {
	pulumi.Input

	ToDdosGeoIpBlockConfigMapOutput() DdosGeoIpBlockConfigMapOutput
	ToDdosGeoIpBlockConfigMapOutputWithContext(context.Context) DdosGeoIpBlockConfigMapOutput
}

type DdosGeoIpBlockConfigMap map[string]DdosGeoIpBlockConfigInput

func (DdosGeoIpBlockConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosGeoIpBlockConfig)(nil)).Elem()
}

func (i DdosGeoIpBlockConfigMap) ToDdosGeoIpBlockConfigMapOutput() DdosGeoIpBlockConfigMapOutput {
	return i.ToDdosGeoIpBlockConfigMapOutputWithContext(context.Background())
}

func (i DdosGeoIpBlockConfigMap) ToDdosGeoIpBlockConfigMapOutputWithContext(ctx context.Context) DdosGeoIpBlockConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DdosGeoIpBlockConfigMapOutput)
}

type DdosGeoIpBlockConfigOutput struct{ *pulumi.OutputState }

func (DdosGeoIpBlockConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DdosGeoIpBlockConfig)(nil)).Elem()
}

func (o DdosGeoIpBlockConfigOutput) ToDdosGeoIpBlockConfigOutput() DdosGeoIpBlockConfigOutput {
	return o
}

func (o DdosGeoIpBlockConfigOutput) ToDdosGeoIpBlockConfigOutputWithContext(ctx context.Context) DdosGeoIpBlockConfigOutput {
	return o
}

// DDoS region blocking configuration, configuration ID cannot be empty when filling in parameters.
func (o DdosGeoIpBlockConfigOutput) DdosGeoIpBlockConfig() DdosGeoIpBlockConfigDdosGeoIpBlockConfigOutput {
	return o.ApplyT(func(v *DdosGeoIpBlockConfig) DdosGeoIpBlockConfigDdosGeoIpBlockConfigOutput {
		return v.DdosGeoIpBlockConfig
	}).(DdosGeoIpBlockConfigDdosGeoIpBlockConfigOutput)
}

// InstanceId.
func (o DdosGeoIpBlockConfigOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DdosGeoIpBlockConfig) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type DdosGeoIpBlockConfigArrayOutput struct{ *pulumi.OutputState }

func (DdosGeoIpBlockConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DdosGeoIpBlockConfig)(nil)).Elem()
}

func (o DdosGeoIpBlockConfigArrayOutput) ToDdosGeoIpBlockConfigArrayOutput() DdosGeoIpBlockConfigArrayOutput {
	return o
}

func (o DdosGeoIpBlockConfigArrayOutput) ToDdosGeoIpBlockConfigArrayOutputWithContext(ctx context.Context) DdosGeoIpBlockConfigArrayOutput {
	return o
}

func (o DdosGeoIpBlockConfigArrayOutput) Index(i pulumi.IntInput) DdosGeoIpBlockConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DdosGeoIpBlockConfig {
		return vs[0].([]*DdosGeoIpBlockConfig)[vs[1].(int)]
	}).(DdosGeoIpBlockConfigOutput)
}

type DdosGeoIpBlockConfigMapOutput struct{ *pulumi.OutputState }

func (DdosGeoIpBlockConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DdosGeoIpBlockConfig)(nil)).Elem()
}

func (o DdosGeoIpBlockConfigMapOutput) ToDdosGeoIpBlockConfigMapOutput() DdosGeoIpBlockConfigMapOutput {
	return o
}

func (o DdosGeoIpBlockConfigMapOutput) ToDdosGeoIpBlockConfigMapOutputWithContext(ctx context.Context) DdosGeoIpBlockConfigMapOutput {
	return o
}

func (o DdosGeoIpBlockConfigMapOutput) MapIndex(k pulumi.StringInput) DdosGeoIpBlockConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DdosGeoIpBlockConfig {
		return vs[0].(map[string]*DdosGeoIpBlockConfig)[vs[1].(string)]
	}).(DdosGeoIpBlockConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DdosGeoIpBlockConfigInput)(nil)).Elem(), &DdosGeoIpBlockConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosGeoIpBlockConfigArrayInput)(nil)).Elem(), DdosGeoIpBlockConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DdosGeoIpBlockConfigMapInput)(nil)).Elem(), DdosGeoIpBlockConfigMap{})
	pulumi.RegisterOutputType(DdosGeoIpBlockConfigOutput{})
	pulumi.RegisterOutputType(DdosGeoIpBlockConfigArrayOutput{})
	pulumi.RegisterOutputType(DdosGeoIpBlockConfigMapOutput{})
}