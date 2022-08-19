// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a SSM secret version.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Ssm"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := Ssm.NewSecret(ctx, "foo", &Ssm.SecretArgs{
// 			SecretName:           pulumi.String("test"),
// 			Description:          pulumi.String("test secret"),
// 			RecoveryWindowInDays: pulumi.Int(0),
// 			IsEnabled:            pulumi.Bool(true),
// 			Tags: pulumi.AnyMap{
// 				"test-tag": pulumi.Any("test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Ssm.NewSecretVersion(ctx, "v1", &Ssm.SecretVersionArgs{
// 			SecretName:   foo.SecretName,
// 			VersionId:    pulumi.String("v1"),
// 			SecretBinary: pulumi.String("MTIzMTIzMTIzMTIzMTIzQQ=="),
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
// SSM secret version can be imported using the secretName#versionId, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Ssm/secretVersion:SecretVersion v1 test#v1
// ```
type SecretVersion struct {
	pulumi.CustomResourceState

	// The base64-encoded binary secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretBinary pulumi.StringPtrOutput `pulumi:"secretBinary"`
	// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	SecretName pulumi.StringOutput `pulumi:"secretName"`
	// The string text of secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretString pulumi.StringPtrOutput `pulumi:"secretString"`
	// Version of secret. The maximum length is 64 bytes. The versionId can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	VersionId pulumi.StringOutput `pulumi:"versionId"`
}

// NewSecretVersion registers a new resource with the given unique name, arguments, and options.
func NewSecretVersion(ctx *pulumi.Context,
	name string, args *SecretVersionArgs, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SecretName == nil {
		return nil, errors.New("invalid value for required argument 'SecretName'")
	}
	if args.VersionId == nil {
		return nil, errors.New("invalid value for required argument 'VersionId'")
	}
	var resource SecretVersion
	err := ctx.RegisterResource("tencentcloud:Ssm/secretVersion:SecretVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretVersion gets an existing SecretVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretVersionState, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	var resource SecretVersion
	err := ctx.ReadResource("tencentcloud:Ssm/secretVersion:SecretVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretVersion resources.
type secretVersionState struct {
	// The base64-encoded binary secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretBinary *string `pulumi:"secretBinary"`
	// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	SecretName *string `pulumi:"secretName"`
	// The string text of secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretString *string `pulumi:"secretString"`
	// Version of secret. The maximum length is 64 bytes. The versionId can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	VersionId *string `pulumi:"versionId"`
}

type SecretVersionState struct {
	// The base64-encoded binary secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretBinary pulumi.StringPtrInput
	// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	SecretName pulumi.StringPtrInput
	// The string text of secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretString pulumi.StringPtrInput
	// Version of secret. The maximum length is 64 bytes. The versionId can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	VersionId pulumi.StringPtrInput
}

func (SecretVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretVersionState)(nil)).Elem()
}

type secretVersionArgs struct {
	// The base64-encoded binary secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretBinary *string `pulumi:"secretBinary"`
	// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	SecretName string `pulumi:"secretName"`
	// The string text of secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretString *string `pulumi:"secretString"`
	// Version of secret. The maximum length is 64 bytes. The versionId can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	VersionId string `pulumi:"versionId"`
}

// The set of arguments for constructing a SecretVersion resource.
type SecretVersionArgs struct {
	// The base64-encoded binary secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretBinary pulumi.StringPtrInput
	// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	SecretName pulumi.StringInput
	// The string text of secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
	SecretString pulumi.StringPtrInput
	// Version of secret. The maximum length is 64 bytes. The versionId can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	VersionId pulumi.StringInput
}

func (SecretVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretVersionArgs)(nil)).Elem()
}

type SecretVersionInput interface {
	pulumi.Input

	ToSecretVersionOutput() SecretVersionOutput
	ToSecretVersionOutputWithContext(ctx context.Context) SecretVersionOutput
}

func (*SecretVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretVersion)(nil)).Elem()
}

func (i *SecretVersion) ToSecretVersionOutput() SecretVersionOutput {
	return i.ToSecretVersionOutputWithContext(context.Background())
}

func (i *SecretVersion) ToSecretVersionOutputWithContext(ctx context.Context) SecretVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionOutput)
}

// SecretVersionArrayInput is an input type that accepts SecretVersionArray and SecretVersionArrayOutput values.
// You can construct a concrete instance of `SecretVersionArrayInput` via:
//
//          SecretVersionArray{ SecretVersionArgs{...} }
type SecretVersionArrayInput interface {
	pulumi.Input

	ToSecretVersionArrayOutput() SecretVersionArrayOutput
	ToSecretVersionArrayOutputWithContext(context.Context) SecretVersionArrayOutput
}

type SecretVersionArray []SecretVersionInput

func (SecretVersionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretVersion)(nil)).Elem()
}

func (i SecretVersionArray) ToSecretVersionArrayOutput() SecretVersionArrayOutput {
	return i.ToSecretVersionArrayOutputWithContext(context.Background())
}

func (i SecretVersionArray) ToSecretVersionArrayOutputWithContext(ctx context.Context) SecretVersionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionArrayOutput)
}

// SecretVersionMapInput is an input type that accepts SecretVersionMap and SecretVersionMapOutput values.
// You can construct a concrete instance of `SecretVersionMapInput` via:
//
//          SecretVersionMap{ "key": SecretVersionArgs{...} }
type SecretVersionMapInput interface {
	pulumi.Input

	ToSecretVersionMapOutput() SecretVersionMapOutput
	ToSecretVersionMapOutputWithContext(context.Context) SecretVersionMapOutput
}

type SecretVersionMap map[string]SecretVersionInput

func (SecretVersionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretVersion)(nil)).Elem()
}

func (i SecretVersionMap) ToSecretVersionMapOutput() SecretVersionMapOutput {
	return i.ToSecretVersionMapOutputWithContext(context.Background())
}

func (i SecretVersionMap) ToSecretVersionMapOutputWithContext(ctx context.Context) SecretVersionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretVersionMapOutput)
}

type SecretVersionOutput struct{ *pulumi.OutputState }

func (SecretVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretVersion)(nil)).Elem()
}

func (o SecretVersionOutput) ToSecretVersionOutput() SecretVersionOutput {
	return o
}

func (o SecretVersionOutput) ToSecretVersionOutputWithContext(ctx context.Context) SecretVersionOutput {
	return o
}

// The base64-encoded binary secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
func (o SecretVersionOutput) SecretBinary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringPtrOutput { return v.SecretBinary }).(pulumi.StringPtrOutput)
}

// Name of secret which cannot be repeated in the same region. The maximum length is 128 bytes. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
func (o SecretVersionOutput) SecretName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.SecretName }).(pulumi.StringOutput)
}

// The string text of secret. secretBinary and secretString must be set only one, and the maximum support is 4096 bytes. When secret status is `Disabled`, this field will not update anymore.
func (o SecretVersionOutput) SecretString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringPtrOutput { return v.SecretString }).(pulumi.StringPtrOutput)
}

// Version of secret. The maximum length is 64 bytes. The versionId can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
func (o SecretVersionOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretVersion) pulumi.StringOutput { return v.VersionId }).(pulumi.StringOutput)
}

type SecretVersionArrayOutput struct{ *pulumi.OutputState }

func (SecretVersionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretVersion)(nil)).Elem()
}

func (o SecretVersionArrayOutput) ToSecretVersionArrayOutput() SecretVersionArrayOutput {
	return o
}

func (o SecretVersionArrayOutput) ToSecretVersionArrayOutputWithContext(ctx context.Context) SecretVersionArrayOutput {
	return o
}

func (o SecretVersionArrayOutput) Index(i pulumi.IntInput) SecretVersionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretVersion {
		return vs[0].([]*SecretVersion)[vs[1].(int)]
	}).(SecretVersionOutput)
}

type SecretVersionMapOutput struct{ *pulumi.OutputState }

func (SecretVersionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretVersion)(nil)).Elem()
}

func (o SecretVersionMapOutput) ToSecretVersionMapOutput() SecretVersionMapOutput {
	return o
}

func (o SecretVersionMapOutput) ToSecretVersionMapOutputWithContext(ctx context.Context) SecretVersionMapOutput {
	return o
}

func (o SecretVersionMapOutput) MapIndex(k pulumi.StringInput) SecretVersionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretVersion {
		return vs[0].(map[string]*SecretVersion)[vs[1].(string)]
	}).(SecretVersionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionInput)(nil)).Elem(), &SecretVersion{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionArrayInput)(nil)).Elem(), SecretVersionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretVersionMapInput)(nil)).Elem(), SecretVersionMap{})
	pulumi.RegisterOutputType(SecretVersionOutput{})
	pulumi.RegisterOutputType(SecretVersionArrayOutput{})
	pulumi.RegisterOutputType(SecretVersionMapOutput{})
}
