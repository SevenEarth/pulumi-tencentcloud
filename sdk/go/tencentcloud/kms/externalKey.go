// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provide a resource to create a KMS external key.
//
// ## Example Usage
//
// ### Create a basic instance.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kms"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Kms.NewExternalKey(ctx, "example", &Kms.ExternalKeyArgs{
//				Alias:       pulumi.String("tf-example-kms-externalkey"),
//				Description: pulumi.String("example of kms external key"),
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
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
// <!--End PulumiCodeChooser -->
//
// ### Specify the encryption algorithm and public key.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kms"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Kms.NewExternalKey(ctx, "example", &Kms.ExternalKeyArgs{
//				Alias:             pulumi.String("tf-example-kms-externalkey"),
//				Description:       pulumi.String("example of kms external key"),
//				IsEnabled:         pulumi.Bool(true),
//				KeyMaterialBase64: pulumi.String("your_public_key_base64_encoded"),
//				Tags: pulumi.Map{
//					"createdBy": pulumi.Any("terraform"),
//				},
//				WrappingAlgorithm: pulumi.String("RSAES_PKCS1_V1_5"),
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
// ### Disable the external kms key.
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Kms"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Kms.NewExternalKey(ctx, "example", &Kms.ExternalKeyArgs{
//				Alias:             pulumi.String("tf-example-kms-externalkey"),
//				Description:       pulumi.String("example of kms external key"),
//				IsEnabled:         pulumi.Bool(false),
//				KeyMaterialBase64: pulumi.String("your_public_key_base64_encoded"),
//				Tags: pulumi.Map{
//					"test-tag": pulumi.Any("unit-test"),
//				},
//				WrappingAlgorithm: pulumi.String("RSAES_PKCS1_V1_5"),
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
// KMS external keys can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Kms/externalKey:ExternalKey example 287e8f40-7cbb-11eb-9a3a-xxxxx
// ```
type ExternalKey struct {
	pulumi.CustomResourceState

	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Description of CMK. The maximum is 1024 bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived pulumi.BoolPtrOutput `pulumi:"isArchived"`
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
	KeyMaterialBase64 pulumi.StringPtrOutput `pulumi:"keyMaterialBase64"`
	// State of CMK.
	KeyState pulumi.StringOutput `pulumi:"keyState"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays pulumi.IntPtrOutput `pulumi:"pendingDeleteWindowInDays"`
	// Tags of CMK.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
	ValidTo pulumi.IntPtrOutput `pulumi:"validTo"`
	// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
	WrappingAlgorithm pulumi.StringPtrOutput `pulumi:"wrappingAlgorithm"`
}

// NewExternalKey registers a new resource with the given unique name, arguments, and options.
func NewExternalKey(ctx *pulumi.Context,
	name string, args *ExternalKeyArgs, opts ...pulumi.ResourceOption) (*ExternalKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.KeyMaterialBase64 != nil {
		args.KeyMaterialBase64 = pulumi.ToSecret(args.KeyMaterialBase64).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keyMaterialBase64",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExternalKey
	err := ctx.RegisterResource("tencentcloud:Kms/externalKey:ExternalKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalKey gets an existing ExternalKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalKeyState, opts ...pulumi.ResourceOption) (*ExternalKey, error) {
	var resource ExternalKey
	err := ctx.ReadResource("tencentcloud:Kms/externalKey:ExternalKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalKey resources.
type externalKeyState struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias *string `pulumi:"alias"`
	// Description of CMK. The maximum is 1024 bytes.
	Description *string `pulumi:"description"`
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived *bool `pulumi:"isArchived"`
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
	KeyMaterialBase64 *string `pulumi:"keyMaterialBase64"`
	// State of CMK.
	KeyState *string `pulumi:"keyState"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays *int `pulumi:"pendingDeleteWindowInDays"`
	// Tags of CMK.
	Tags map[string]interface{} `pulumi:"tags"`
	// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
	ValidTo *int `pulumi:"validTo"`
	// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
	WrappingAlgorithm *string `pulumi:"wrappingAlgorithm"`
}

type ExternalKeyState struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias pulumi.StringPtrInput
	// Description of CMK. The maximum is 1024 bytes.
	Description pulumi.StringPtrInput
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived pulumi.BoolPtrInput
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled pulumi.BoolPtrInput
	// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
	KeyMaterialBase64 pulumi.StringPtrInput
	// State of CMK.
	KeyState pulumi.StringPtrInput
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays pulumi.IntPtrInput
	// Tags of CMK.
	Tags pulumi.MapInput
	// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
	ValidTo pulumi.IntPtrInput
	// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
	WrappingAlgorithm pulumi.StringPtrInput
}

func (ExternalKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalKeyState)(nil)).Elem()
}

type externalKeyArgs struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias string `pulumi:"alias"`
	// Description of CMK. The maximum is 1024 bytes.
	Description *string `pulumi:"description"`
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived *bool `pulumi:"isArchived"`
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
	KeyMaterialBase64 *string `pulumi:"keyMaterialBase64"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays *int `pulumi:"pendingDeleteWindowInDays"`
	// Tags of CMK.
	Tags map[string]interface{} `pulumi:"tags"`
	// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
	ValidTo *int `pulumi:"validTo"`
	// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
	WrappingAlgorithm *string `pulumi:"wrappingAlgorithm"`
}

// The set of arguments for constructing a ExternalKey resource.
type ExternalKeyArgs struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias pulumi.StringInput
	// Description of CMK. The maximum is 1024 bytes.
	Description pulumi.StringPtrInput
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived pulumi.BoolPtrInput
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled pulumi.BoolPtrInput
	// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
	KeyMaterialBase64 pulumi.StringPtrInput
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays pulumi.IntPtrInput
	// Tags of CMK.
	Tags pulumi.MapInput
	// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
	ValidTo pulumi.IntPtrInput
	// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
	WrappingAlgorithm pulumi.StringPtrInput
}

func (ExternalKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalKeyArgs)(nil)).Elem()
}

type ExternalKeyInput interface {
	pulumi.Input

	ToExternalKeyOutput() ExternalKeyOutput
	ToExternalKeyOutputWithContext(ctx context.Context) ExternalKeyOutput
}

func (*ExternalKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalKey)(nil)).Elem()
}

func (i *ExternalKey) ToExternalKeyOutput() ExternalKeyOutput {
	return i.ToExternalKeyOutputWithContext(context.Background())
}

func (i *ExternalKey) ToExternalKeyOutputWithContext(ctx context.Context) ExternalKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyOutput)
}

// ExternalKeyArrayInput is an input type that accepts ExternalKeyArray and ExternalKeyArrayOutput values.
// You can construct a concrete instance of `ExternalKeyArrayInput` via:
//
//	ExternalKeyArray{ ExternalKeyArgs{...} }
type ExternalKeyArrayInput interface {
	pulumi.Input

	ToExternalKeyArrayOutput() ExternalKeyArrayOutput
	ToExternalKeyArrayOutputWithContext(context.Context) ExternalKeyArrayOutput
}

type ExternalKeyArray []ExternalKeyInput

func (ExternalKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalKey)(nil)).Elem()
}

func (i ExternalKeyArray) ToExternalKeyArrayOutput() ExternalKeyArrayOutput {
	return i.ToExternalKeyArrayOutputWithContext(context.Background())
}

func (i ExternalKeyArray) ToExternalKeyArrayOutputWithContext(ctx context.Context) ExternalKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyArrayOutput)
}

// ExternalKeyMapInput is an input type that accepts ExternalKeyMap and ExternalKeyMapOutput values.
// You can construct a concrete instance of `ExternalKeyMapInput` via:
//
//	ExternalKeyMap{ "key": ExternalKeyArgs{...} }
type ExternalKeyMapInput interface {
	pulumi.Input

	ToExternalKeyMapOutput() ExternalKeyMapOutput
	ToExternalKeyMapOutputWithContext(context.Context) ExternalKeyMapOutput
}

type ExternalKeyMap map[string]ExternalKeyInput

func (ExternalKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalKey)(nil)).Elem()
}

func (i ExternalKeyMap) ToExternalKeyMapOutput() ExternalKeyMapOutput {
	return i.ToExternalKeyMapOutputWithContext(context.Background())
}

func (i ExternalKeyMap) ToExternalKeyMapOutputWithContext(ctx context.Context) ExternalKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalKeyMapOutput)
}

type ExternalKeyOutput struct{ *pulumi.OutputState }

func (ExternalKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalKey)(nil)).Elem()
}

func (o ExternalKeyOutput) ToExternalKeyOutput() ExternalKeyOutput {
	return o
}

func (o ExternalKeyOutput) ToExternalKeyOutputWithContext(ctx context.Context) ExternalKeyOutput {
	return o
}

// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
func (o ExternalKeyOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// Description of CMK. The maximum is 1024 bytes.
func (o ExternalKeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
func (o ExternalKeyOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.BoolPtrOutput { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
func (o ExternalKeyOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The base64-encoded key material encrypted with the public_key. For regions using the national secret version, the length of the imported key material is required to be 128 bits, and for regions using the FIPS version, the length of the imported key material is required to be 256 bits.
func (o ExternalKeyOutput) KeyMaterialBase64() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.StringPtrOutput { return v.KeyMaterialBase64 }).(pulumi.StringPtrOutput)
}

// State of CMK.
func (o ExternalKeyOutput) KeyState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.StringOutput { return v.KeyState }).(pulumi.StringOutput)
}

// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
func (o ExternalKeyOutput) PendingDeleteWindowInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.IntPtrOutput { return v.PendingDeleteWindowInDays }).(pulumi.IntPtrOutput)
}

// Tags of CMK.
func (o ExternalKeyOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// This value means the effective timestamp of the key material, 0 means it does not expire. Need to be greater than the current timestamp, the maximum support is 2147443200.
func (o ExternalKeyOutput) ValidTo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.IntPtrOutput { return v.ValidTo }).(pulumi.IntPtrOutput)
}

// The algorithm for encrypting key material. Available values include `RSAES_PKCS1_V1_5`, `RSAES_OAEP_SHA_1` and `RSAES_OAEP_SHA_256`. Default value is `RSAES_PKCS1_V1_5`.
func (o ExternalKeyOutput) WrappingAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalKey) pulumi.StringPtrOutput { return v.WrappingAlgorithm }).(pulumi.StringPtrOutput)
}

type ExternalKeyArrayOutput struct{ *pulumi.OutputState }

func (ExternalKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalKey)(nil)).Elem()
}

func (o ExternalKeyArrayOutput) ToExternalKeyArrayOutput() ExternalKeyArrayOutput {
	return o
}

func (o ExternalKeyArrayOutput) ToExternalKeyArrayOutputWithContext(ctx context.Context) ExternalKeyArrayOutput {
	return o
}

func (o ExternalKeyArrayOutput) Index(i pulumi.IntInput) ExternalKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalKey {
		return vs[0].([]*ExternalKey)[vs[1].(int)]
	}).(ExternalKeyOutput)
}

type ExternalKeyMapOutput struct{ *pulumi.OutputState }

func (ExternalKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalKey)(nil)).Elem()
}

func (o ExternalKeyMapOutput) ToExternalKeyMapOutput() ExternalKeyMapOutput {
	return o
}

func (o ExternalKeyMapOutput) ToExternalKeyMapOutputWithContext(ctx context.Context) ExternalKeyMapOutput {
	return o
}

func (o ExternalKeyMapOutput) MapIndex(k pulumi.StringInput) ExternalKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalKey {
		return vs[0].(map[string]*ExternalKey)[vs[1].(string)]
	}).(ExternalKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalKeyInput)(nil)).Elem(), &ExternalKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalKeyArrayInput)(nil)).Elem(), ExternalKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalKeyMapInput)(nil)).Elem(), ExternalKeyMap{})
	pulumi.RegisterOutputType(ExternalKeyOutput{})
	pulumi.RegisterOutputType(ExternalKeyArrayOutput{})
	pulumi.RegisterOutputType(ExternalKeyMapOutput{})
}
