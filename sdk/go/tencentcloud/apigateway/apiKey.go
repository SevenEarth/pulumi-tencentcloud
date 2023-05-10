// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create API gateway access key.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/ApiGateway"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ApiGateway.NewApiKey(ctx, "test", &ApiGateway.ApiKeyArgs{
// 			SecretName: pulumi.String("my_api_key"),
// 			Status:     pulumi.String("on"),
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
// API gateway access key can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:ApiGateway/apiKey:ApiKey test AKIDMZwceezso9ps5p8jkro8a9fwe1e7nzF2k50B
// ```
type ApiKey struct {
	pulumi.CustomResourceState

	// Created API key.
	AccessKeySecret pulumi.StringOutput `pulumi:"accessKeySecret"`
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	// Custom key name.
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// Key status. Valid values: `on`, `off`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecretName == nil {
		return nil, errors.New("invalid value for required argument 'SecretName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ApiKey
	err := ctx.RegisterResource("tencentcloud:ApiGateway/apiKey:ApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	var resource ApiKey
	err := ctx.ReadResource("tencentcloud:ApiGateway/apiKey:ApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type apiKeyState struct {
	// Created API key.
	AccessKeySecret *string `pulumi:"accessKeySecret"`
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime *string `pulumi:"createTime"`
	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	ModifyTime *string `pulumi:"modifyTime"`
	// Custom key name.
	SecretName *string `pulumi:"secretName"`
	// Key status. Valid values: `on`, `off`.
	Status *string `pulumi:"status"`
}

type ApiKeyState struct {
	// Created API key.
	AccessKeySecret pulumi.StringPtrInput
	// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	CreateTime pulumi.StringPtrInput
	// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
	ModifyTime pulumi.StringPtrInput
	// Custom key name.
	SecretName pulumi.StringPtrInput
	// Key status. Valid values: `on`, `off`.
	Status pulumi.StringPtrInput
}

func (ApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyState)(nil)).Elem()
}

type apiKeyArgs struct {
	// Custom key name.
	SecretName string `pulumi:"secretName"`
	// Key status. Valid values: `on`, `off`.
	Status *string `pulumi:"status"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// Custom key name.
	SecretName pulumi.StringInput
	// Key status. Valid values: `on`, `off`.
	Status pulumi.StringPtrInput
}

func (ApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyArgs)(nil)).Elem()
}

type ApiKeyInput interface {
	pulumi.Input

	ToApiKeyOutput() ApiKeyOutput
	ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput
}

func (*ApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (i *ApiKey) ToApiKeyOutput() ApiKeyOutput {
	return i.ToApiKeyOutputWithContext(context.Background())
}

func (i *ApiKey) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyOutput)
}

// ApiKeyArrayInput is an input type that accepts ApiKeyArray and ApiKeyArrayOutput values.
// You can construct a concrete instance of `ApiKeyArrayInput` via:
//
//          ApiKeyArray{ ApiKeyArgs{...} }
type ApiKeyArrayInput interface {
	pulumi.Input

	ToApiKeyArrayOutput() ApiKeyArrayOutput
	ToApiKeyArrayOutputWithContext(context.Context) ApiKeyArrayOutput
}

type ApiKeyArray []ApiKeyInput

func (ApiKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiKey)(nil)).Elem()
}

func (i ApiKeyArray) ToApiKeyArrayOutput() ApiKeyArrayOutput {
	return i.ToApiKeyArrayOutputWithContext(context.Background())
}

func (i ApiKeyArray) ToApiKeyArrayOutputWithContext(ctx context.Context) ApiKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyArrayOutput)
}

// ApiKeyMapInput is an input type that accepts ApiKeyMap and ApiKeyMapOutput values.
// You can construct a concrete instance of `ApiKeyMapInput` via:
//
//          ApiKeyMap{ "key": ApiKeyArgs{...} }
type ApiKeyMapInput interface {
	pulumi.Input

	ToApiKeyMapOutput() ApiKeyMapOutput
	ToApiKeyMapOutputWithContext(context.Context) ApiKeyMapOutput
}

type ApiKeyMap map[string]ApiKeyInput

func (ApiKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiKey)(nil)).Elem()
}

func (i ApiKeyMap) ToApiKeyMapOutput() ApiKeyMapOutput {
	return i.ToApiKeyMapOutputWithContext(context.Background())
}

func (i ApiKeyMap) ToApiKeyMapOutputWithContext(ctx context.Context) ApiKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyMapOutput)
}

type ApiKeyOutput struct{ *pulumi.OutputState }

func (ApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (o ApiKeyOutput) ToApiKeyOutput() ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return o
}

// Created API key.
func (o ApiKeyOutput) AccessKeySecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.AccessKeySecret }).(pulumi.StringOutput)
}

// Creation time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
func (o ApiKeyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Last modified time in the format of YYYY-MM-DDThh:mm:ssZ according to ISO 8601 standard. UTC time is used.
func (o ApiKeyOutput) ModifyTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.ModifyTime }).(pulumi.StringOutput)
}

// Custom key name.
func (o ApiKeyOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// Key status. Valid values: `on`, `off`.
func (o ApiKeyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiKey) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

type ApiKeyArrayOutput struct{ *pulumi.OutputState }

func (ApiKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApiKey)(nil)).Elem()
}

func (o ApiKeyArrayOutput) ToApiKeyArrayOutput() ApiKeyArrayOutput {
	return o
}

func (o ApiKeyArrayOutput) ToApiKeyArrayOutputWithContext(ctx context.Context) ApiKeyArrayOutput {
	return o
}

func (o ApiKeyArrayOutput) Index(i pulumi.IntInput) ApiKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApiKey {
		return vs[0].([]*ApiKey)[vs[1].(int)]
	}).(ApiKeyOutput)
}

type ApiKeyMapOutput struct{ *pulumi.OutputState }

func (ApiKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApiKey)(nil)).Elem()
}

func (o ApiKeyMapOutput) ToApiKeyMapOutput() ApiKeyMapOutput {
	return o
}

func (o ApiKeyMapOutput) ToApiKeyMapOutputWithContext(ctx context.Context) ApiKeyMapOutput {
	return o
}

func (o ApiKeyMapOutput) MapIndex(k pulumi.StringInput) ApiKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApiKey {
		return vs[0].(map[string]*ApiKey)[vs[1].(string)]
	}).(ApiKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyInput)(nil)).Elem(), &ApiKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyArrayInput)(nil)).Elem(), ApiKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApiKeyMapInput)(nil)).Elem(), ApiKeyMap{})
	pulumi.RegisterOutputType(ApiKeyOutput{})
	pulumi.RegisterOutputType(ApiKeyArrayOutput{})
	pulumi.RegisterOutputType(ApiKeyMapOutput{})
}
