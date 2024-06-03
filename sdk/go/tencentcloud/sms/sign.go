// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a sms sign
//
// ## Example Usage
//
// ### Create a sms sign instance
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Sms"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Sms.NewSign(ctx, "example", &Sms.SignArgs{
//				DocumentType:  pulumi.Int(4),
//				International: pulumi.Int(0),
//				ProofImage:    pulumi.String("your_proof_image"),
//				SignName:      pulumi.String("tf_example_sms_sign"),
//				SignPurpose:   pulumi.Int(0),
//				SignType:      pulumi.Int(1),
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
type Sign struct {
	pulumi.CustomResourceState

	// Power of attorney, which should be submitted if SignPurpose is for use by others. You should Base64-encode the image first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter. Note: this field will take effect only when SignPurpose is 1 (for user by others).
	CommissionImage pulumi.StringPtrOutput `pulumi:"commissionImage"`
	// DocumentType is used for enterprise authentication, or website, app authentication, etc. DocumentType: 0, 1, 2, 3, 4, 5, 6, 7, 8.
	DocumentType pulumi.IntOutput `pulumi:"documentType"`
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International pulumi.IntOutput `pulumi:"international"`
	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter.
	ProofImage pulumi.StringOutput `pulumi:"proofImage"`
	// Signature application remarks.
	Remark pulumi.StringPtrOutput `pulumi:"remark"`
	// Sms sign name, unique.
	SignName pulumi.StringOutput `pulumi:"signName"`
	// Signature purpose: 0: for personal use; 1: for others.
	SignPurpose pulumi.IntOutput `pulumi:"signPurpose"`
	// Sms sign type: 0, 1, 2, 3, 4, 5, 6.
	SignType pulumi.IntOutput `pulumi:"signType"`
}

// NewSign registers a new resource with the given unique name, arguments, and options.
func NewSign(ctx *pulumi.Context,
	name string, args *SignArgs, opts ...pulumi.ResourceOption) (*Sign, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DocumentType == nil {
		return nil, errors.New("invalid value for required argument 'DocumentType'")
	}
	if args.International == nil {
		return nil, errors.New("invalid value for required argument 'International'")
	}
	if args.ProofImage == nil {
		return nil, errors.New("invalid value for required argument 'ProofImage'")
	}
	if args.SignName == nil {
		return nil, errors.New("invalid value for required argument 'SignName'")
	}
	if args.SignPurpose == nil {
		return nil, errors.New("invalid value for required argument 'SignPurpose'")
	}
	if args.SignType == nil {
		return nil, errors.New("invalid value for required argument 'SignType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Sign
	err := ctx.RegisterResource("tencentcloud:Sms/sign:Sign", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSign gets an existing Sign resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSign(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignState, opts ...pulumi.ResourceOption) (*Sign, error) {
	var resource Sign
	err := ctx.ReadResource("tencentcloud:Sms/sign:Sign", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sign resources.
type signState struct {
	// Power of attorney, which should be submitted if SignPurpose is for use by others. You should Base64-encode the image first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter. Note: this field will take effect only when SignPurpose is 1 (for user by others).
	CommissionImage *string `pulumi:"commissionImage"`
	// DocumentType is used for enterprise authentication, or website, app authentication, etc. DocumentType: 0, 1, 2, 3, 4, 5, 6, 7, 8.
	DocumentType *int `pulumi:"documentType"`
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International *int `pulumi:"international"`
	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter.
	ProofImage *string `pulumi:"proofImage"`
	// Signature application remarks.
	Remark *string `pulumi:"remark"`
	// Sms sign name, unique.
	SignName *string `pulumi:"signName"`
	// Signature purpose: 0: for personal use; 1: for others.
	SignPurpose *int `pulumi:"signPurpose"`
	// Sms sign type: 0, 1, 2, 3, 4, 5, 6.
	SignType *int `pulumi:"signType"`
}

type SignState struct {
	// Power of attorney, which should be submitted if SignPurpose is for use by others. You should Base64-encode the image first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter. Note: this field will take effect only when SignPurpose is 1 (for user by others).
	CommissionImage pulumi.StringPtrInput
	// DocumentType is used for enterprise authentication, or website, app authentication, etc. DocumentType: 0, 1, 2, 3, 4, 5, 6, 7, 8.
	DocumentType pulumi.IntPtrInput
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International pulumi.IntPtrInput
	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter.
	ProofImage pulumi.StringPtrInput
	// Signature application remarks.
	Remark pulumi.StringPtrInput
	// Sms sign name, unique.
	SignName pulumi.StringPtrInput
	// Signature purpose: 0: for personal use; 1: for others.
	SignPurpose pulumi.IntPtrInput
	// Sms sign type: 0, 1, 2, 3, 4, 5, 6.
	SignType pulumi.IntPtrInput
}

func (SignState) ElementType() reflect.Type {
	return reflect.TypeOf((*signState)(nil)).Elem()
}

type signArgs struct {
	// Power of attorney, which should be submitted if SignPurpose is for use by others. You should Base64-encode the image first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter. Note: this field will take effect only when SignPurpose is 1 (for user by others).
	CommissionImage *string `pulumi:"commissionImage"`
	// DocumentType is used for enterprise authentication, or website, app authentication, etc. DocumentType: 0, 1, 2, 3, 4, 5, 6, 7, 8.
	DocumentType int `pulumi:"documentType"`
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International int `pulumi:"international"`
	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter.
	ProofImage string `pulumi:"proofImage"`
	// Signature application remarks.
	Remark *string `pulumi:"remark"`
	// Sms sign name, unique.
	SignName string `pulumi:"signName"`
	// Signature purpose: 0: for personal use; 1: for others.
	SignPurpose int `pulumi:"signPurpose"`
	// Sms sign type: 0, 1, 2, 3, 4, 5, 6.
	SignType int `pulumi:"signType"`
}

// The set of arguments for constructing a Sign resource.
type SignArgs struct {
	// Power of attorney, which should be submitted if SignPurpose is for use by others. You should Base64-encode the image first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter. Note: this field will take effect only when SignPurpose is 1 (for user by others).
	CommissionImage pulumi.StringPtrInput
	// DocumentType is used for enterprise authentication, or website, app authentication, etc. DocumentType: 0, 1, 2, 3, 4, 5, 6, 7, 8.
	DocumentType pulumi.IntInput
	// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
	International pulumi.IntInput
	// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter.
	ProofImage pulumi.StringInput
	// Signature application remarks.
	Remark pulumi.StringPtrInput
	// Sms sign name, unique.
	SignName pulumi.StringInput
	// Signature purpose: 0: for personal use; 1: for others.
	SignPurpose pulumi.IntInput
	// Sms sign type: 0, 1, 2, 3, 4, 5, 6.
	SignType pulumi.IntInput
}

func (SignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signArgs)(nil)).Elem()
}

type SignInput interface {
	pulumi.Input

	ToSignOutput() SignOutput
	ToSignOutputWithContext(ctx context.Context) SignOutput
}

func (*Sign) ElementType() reflect.Type {
	return reflect.TypeOf((**Sign)(nil)).Elem()
}

func (i *Sign) ToSignOutput() SignOutput {
	return i.ToSignOutputWithContext(context.Background())
}

func (i *Sign) ToSignOutputWithContext(ctx context.Context) SignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignOutput)
}

// SignArrayInput is an input type that accepts SignArray and SignArrayOutput values.
// You can construct a concrete instance of `SignArrayInput` via:
//
//	SignArray{ SignArgs{...} }
type SignArrayInput interface {
	pulumi.Input

	ToSignArrayOutput() SignArrayOutput
	ToSignArrayOutputWithContext(context.Context) SignArrayOutput
}

type SignArray []SignInput

func (SignArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sign)(nil)).Elem()
}

func (i SignArray) ToSignArrayOutput() SignArrayOutput {
	return i.ToSignArrayOutputWithContext(context.Background())
}

func (i SignArray) ToSignArrayOutputWithContext(ctx context.Context) SignArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignArrayOutput)
}

// SignMapInput is an input type that accepts SignMap and SignMapOutput values.
// You can construct a concrete instance of `SignMapInput` via:
//
//	SignMap{ "key": SignArgs{...} }
type SignMapInput interface {
	pulumi.Input

	ToSignMapOutput() SignMapOutput
	ToSignMapOutputWithContext(context.Context) SignMapOutput
}

type SignMap map[string]SignInput

func (SignMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sign)(nil)).Elem()
}

func (i SignMap) ToSignMapOutput() SignMapOutput {
	return i.ToSignMapOutputWithContext(context.Background())
}

func (i SignMap) ToSignMapOutputWithContext(ctx context.Context) SignMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignMapOutput)
}

type SignOutput struct{ *pulumi.OutputState }

func (SignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sign)(nil)).Elem()
}

func (o SignOutput) ToSignOutput() SignOutput {
	return o
}

func (o SignOutput) ToSignOutputWithContext(ctx context.Context) SignOutput {
	return o
}

// Power of attorney, which should be submitted if SignPurpose is for use by others. You should Base64-encode the image first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter. Note: this field will take effect only when SignPurpose is 1 (for user by others).
func (o SignOutput) CommissionImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sign) pulumi.StringPtrOutput { return v.CommissionImage }).(pulumi.StringPtrOutput)
}

// DocumentType is used for enterprise authentication, or website, app authentication, etc. DocumentType: 0, 1, 2, 3, 4, 5, 6, 7, 8.
func (o SignOutput) DocumentType() pulumi.IntOutput {
	return o.ApplyT(func(v *Sign) pulumi.IntOutput { return v.DocumentType }).(pulumi.IntOutput)
}

// Whether it is Global SMS: 0: Mainland China SMS; 1: Global SMS.
func (o SignOutput) International() pulumi.IntOutput {
	return o.ApplyT(func(v *Sign) pulumi.IntOutput { return v.International }).(pulumi.IntOutput)
}

// You should Base64-encode the image of the identity certificate corresponding to the signature first, remove the prefix data:image/jpeg;base64, from the resulted string, and then use it as the value of this parameter.
func (o SignOutput) ProofImage() pulumi.StringOutput {
	return o.ApplyT(func(v *Sign) pulumi.StringOutput { return v.ProofImage }).(pulumi.StringOutput)
}

// Signature application remarks.
func (o SignOutput) Remark() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sign) pulumi.StringPtrOutput { return v.Remark }).(pulumi.StringPtrOutput)
}

// Sms sign name, unique.
func (o SignOutput) SignName() pulumi.StringOutput {
	return o.ApplyT(func(v *Sign) pulumi.StringOutput { return v.SignName }).(pulumi.StringOutput)
}

// Signature purpose: 0: for personal use; 1: for others.
func (o SignOutput) SignPurpose() pulumi.IntOutput {
	return o.ApplyT(func(v *Sign) pulumi.IntOutput { return v.SignPurpose }).(pulumi.IntOutput)
}

// Sms sign type: 0, 1, 2, 3, 4, 5, 6.
func (o SignOutput) SignType() pulumi.IntOutput {
	return o.ApplyT(func(v *Sign) pulumi.IntOutput { return v.SignType }).(pulumi.IntOutput)
}

type SignArrayOutput struct{ *pulumi.OutputState }

func (SignArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sign)(nil)).Elem()
}

func (o SignArrayOutput) ToSignArrayOutput() SignArrayOutput {
	return o
}

func (o SignArrayOutput) ToSignArrayOutputWithContext(ctx context.Context) SignArrayOutput {
	return o
}

func (o SignArrayOutput) Index(i pulumi.IntInput) SignOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sign {
		return vs[0].([]*Sign)[vs[1].(int)]
	}).(SignOutput)
}

type SignMapOutput struct{ *pulumi.OutputState }

func (SignMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sign)(nil)).Elem()
}

func (o SignMapOutput) ToSignMapOutput() SignMapOutput {
	return o
}

func (o SignMapOutput) ToSignMapOutputWithContext(ctx context.Context) SignMapOutput {
	return o
}

func (o SignMapOutput) MapIndex(k pulumi.StringInput) SignOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sign {
		return vs[0].(map[string]*Sign)[vs[1].(string)]
	}).(SignOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SignInput)(nil)).Elem(), &Sign{})
	pulumi.RegisterInputType(reflect.TypeOf((*SignArrayInput)(nil)).Elem(), SignArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SignMapInput)(nil)).Elem(), SignMap{})
	pulumi.RegisterOutputType(SignOutput{})
	pulumi.RegisterOutputType(SignArrayOutput{})
	pulumi.RegisterOutputType(SignMapOutput{})
}
