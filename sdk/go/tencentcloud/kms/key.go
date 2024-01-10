// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to create a KMS key.
//
// ## Example Usage
// ### Create and enable a instance.
//
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
//			_, err := Kms.NewKey(ctx, "example", &Kms.KeyArgs{
//				Alias:              pulumi.String("tf-example-kms-key"),
//				Description:        pulumi.String("example of kms key"),
//				IsEnabled:          pulumi.Bool(true),
//				KeyRotationEnabled: pulumi.Bool(false),
//				Tags: pulumi.AnyMap{
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
// ### Specify the Key Usage as an asymmetry method.
//
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
//			_, err := Kms.NewKey(ctx, "example2", &Kms.KeyArgs{
//				Alias:       pulumi.String("tf-example-kms-key"),
//				Description: pulumi.String("example of kms key"),
//				IsEnabled:   pulumi.Bool(false),
//				KeyUsage:    pulumi.String("ASYMMETRIC_DECRYPT_RSA_2048"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Disable the kms key instance.
//
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
//			_, err := Kms.NewKey(ctx, "example3", &Kms.KeyArgs{
//				Alias:              pulumi.String("tf-example-kms-key"),
//				Description:        pulumi.String("example of kms key"),
//				IsEnabled:          pulumi.Bool(false),
//				KeyRotationEnabled: pulumi.Bool(false),
//				Tags: pulumi.AnyMap{
//					"test-tag": pulumi.Any("unit-test"),
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
// KMS keys can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Kms/key:Key foo 287e8f40-7cbb-11eb-9a3a-5254004f7f94
//
// ```
type Key struct {
	pulumi.CustomResourceState

	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Description of CMK. The maximum is 1024 bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived pulumi.BoolPtrOutput `pulumi:"isArchived"`
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// Specify whether to enable key rotation, valid when keyUsage is `ENCRYPT_DECRYPT`. Default value is `false`.
	KeyRotationEnabled pulumi.BoolPtrOutput `pulumi:"keyRotationEnabled"`
	// State of CMK.
	KeyState pulumi.StringOutput `pulumi:"keyState"`
	// Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrOutput `pulumi:"keyUsage"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays pulumi.IntPtrOutput `pulumi:"pendingDeleteWindowInDays"`
	// Tags of CMK.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Key
	err := ctx.RegisterResource("tencentcloud:Kms/key:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("tencentcloud:Kms/key:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias *string `pulumi:"alias"`
	// Description of CMK. The maximum is 1024 bytes.
	Description *string `pulumi:"description"`
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived *bool `pulumi:"isArchived"`
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specify whether to enable key rotation, valid when keyUsage is `ENCRYPT_DECRYPT`. Default value is `false`.
	KeyRotationEnabled *bool `pulumi:"keyRotationEnabled"`
	// State of CMK.
	KeyState *string `pulumi:"keyState"`
	// Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
	KeyUsage *string `pulumi:"keyUsage"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays *int `pulumi:"pendingDeleteWindowInDays"`
	// Tags of CMK.
	Tags map[string]interface{} `pulumi:"tags"`
}

type KeyState struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias pulumi.StringPtrInput
	// Description of CMK. The maximum is 1024 bytes.
	Description pulumi.StringPtrInput
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived pulumi.BoolPtrInput
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled pulumi.BoolPtrInput
	// Specify whether to enable key rotation, valid when keyUsage is `ENCRYPT_DECRYPT`. Default value is `false`.
	KeyRotationEnabled pulumi.BoolPtrInput
	// State of CMK.
	KeyState pulumi.StringPtrInput
	// Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrInput
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays pulumi.IntPtrInput
	// Tags of CMK.
	Tags pulumi.MapInput
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias string `pulumi:"alias"`
	// Description of CMK. The maximum is 1024 bytes.
	Description *string `pulumi:"description"`
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived *bool `pulumi:"isArchived"`
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Specify whether to enable key rotation, valid when keyUsage is `ENCRYPT_DECRYPT`. Default value is `false`.
	KeyRotationEnabled *bool `pulumi:"keyRotationEnabled"`
	// Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
	KeyUsage *string `pulumi:"keyUsage"`
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays *int `pulumi:"pendingDeleteWindowInDays"`
	// Tags of CMK.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
	Alias pulumi.StringInput
	// Description of CMK. The maximum is 1024 bytes.
	Description pulumi.StringPtrInput
	// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsArchived pulumi.BoolPtrInput
	// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
	IsEnabled pulumi.BoolPtrInput
	// Specify whether to enable key rotation, valid when keyUsage is `ENCRYPT_DECRYPT`. Default value is `false`.
	KeyRotationEnabled pulumi.BoolPtrInput
	// Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
	KeyUsage pulumi.StringPtrInput
	// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
	PendingDeleteWindowInDays pulumi.IntPtrInput
	// Tags of CMK.
	Tags pulumi.MapInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

// KeyArrayInput is an input type that accepts KeyArray and KeyArrayOutput values.
// You can construct a concrete instance of `KeyArrayInput` via:
//
//	KeyArray{ KeyArgs{...} }
type KeyArrayInput interface {
	pulumi.Input

	ToKeyArrayOutput() KeyArrayOutput
	ToKeyArrayOutputWithContext(context.Context) KeyArrayOutput
}

type KeyArray []KeyInput

func (KeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (i KeyArray) ToKeyArrayOutput() KeyArrayOutput {
	return i.ToKeyArrayOutputWithContext(context.Background())
}

func (i KeyArray) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyArrayOutput)
}

// KeyMapInput is an input type that accepts KeyMap and KeyMapOutput values.
// You can construct a concrete instance of `KeyMapInput` via:
//
//	KeyMap{ "key": KeyArgs{...} }
type KeyMapInput interface {
	pulumi.Input

	ToKeyMapOutput() KeyMapOutput
	ToKeyMapOutputWithContext(context.Context) KeyMapOutput
}

type KeyMap map[string]KeyInput

func (KeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (i KeyMap) ToKeyMapOutput() KeyMapOutput {
	return i.ToKeyMapOutputWithContext(context.Background())
}

func (i KeyMap) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyMapOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

// Name of CMK. The name can only contain English letters, numbers, underscore and hyphen '-'. The first character must be a letter or number.
func (o KeyOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// Description of CMK. The maximum is 1024 bytes.
func (o KeyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specify whether to archive key. Default value is `false`. This field is conflict with `isEnabled`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
func (o KeyOutput) IsArchived() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.IsArchived }).(pulumi.BoolPtrOutput)
}

// Specify whether to enable key. Default value is `false`. This field is conflict with `isArchived`, valid when keyState is `Enabled`, `Disabled`, `Archived`.
func (o KeyOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Specify whether to enable key rotation, valid when keyUsage is `ENCRYPT_DECRYPT`. Default value is `false`.
func (o KeyOutput) KeyRotationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.BoolPtrOutput { return v.KeyRotationEnabled }).(pulumi.BoolPtrOutput)
}

// State of CMK.
func (o KeyOutput) KeyState() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.KeyState }).(pulumi.StringOutput)
}

// Usage of CMK. Available values include `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
func (o KeyOutput) KeyUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.KeyUsage }).(pulumi.StringPtrOutput)
}

// Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days. Defaults to 7 days.
func (o KeyOutput) PendingDeleteWindowInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.IntPtrOutput { return v.PendingDeleteWindowInDays }).(pulumi.IntPtrOutput)
}

// Tags of CMK.
func (o KeyOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Key) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

type KeyArrayOutput struct{ *pulumi.OutputState }

func (KeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Key)(nil)).Elem()
}

func (o KeyArrayOutput) ToKeyArrayOutput() KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) ToKeyArrayOutputWithContext(ctx context.Context) KeyArrayOutput {
	return o
}

func (o KeyArrayOutput) Index(i pulumi.IntInput) KeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Key {
		return vs[0].([]*Key)[vs[1].(int)]
	}).(KeyOutput)
}

type KeyMapOutput struct{ *pulumi.OutputState }

func (KeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Key)(nil)).Elem()
}

func (o KeyMapOutput) ToKeyMapOutput() KeyMapOutput {
	return o
}

func (o KeyMapOutput) ToKeyMapOutputWithContext(ctx context.Context) KeyMapOutput {
	return o
}

func (o KeyMapOutput) MapIndex(k pulumi.StringInput) KeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Key {
		return vs[0].(map[string]*Key)[vs[1].(string)]
	}).(KeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyArrayInput)(nil)).Elem(), KeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyMapInput)(nil)).Elem(), KeyMap{})
	pulumi.RegisterOutputType(KeyOutput{})
	pulumi.RegisterOutputType(KeyArrayOutput{})
	pulumi.RegisterOutputType(KeyMapOutput{})
}
