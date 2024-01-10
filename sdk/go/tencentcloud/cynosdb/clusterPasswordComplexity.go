// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cynosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cynosdb clusterPasswordComplexity
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cynosdb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cynosdb.NewClusterPasswordComplexity(ctx, "clusterPasswordComplexity", &Cynosdb.ClusterPasswordComplexityArgs{
//				ClusterId: pulumi.String("cynosdbmysql-cgd2gpwr"),
//				ValidatePasswordDictionaries: pulumi.StringArray{
//					pulumi.String("cccc"),
//					pulumi.String("xxxx"),
//				},
//				ValidatePasswordLength:           pulumi.Int(8),
//				ValidatePasswordMixedCaseCount:   pulumi.Int(1),
//				ValidatePasswordNumberCount:      pulumi.Int(1),
//				ValidatePasswordPolicy:           pulumi.String("STRONG"),
//				ValidatePasswordSpecialCharCount: pulumi.Int(1),
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
// cynosdb cluster_password_complexity can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import tencentcloud:Cynosdb/clusterPasswordComplexity:ClusterPasswordComplexity cluster_password_complexity cluster_password_complexity_id
//
// ```
type ClusterPasswordComplexity struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Data dictionary.
	ValidatePasswordDictionaries pulumi.StringArrayOutput `pulumi:"validatePasswordDictionaries"`
	// Password length.
	ValidatePasswordLength pulumi.IntOutput `pulumi:"validatePasswordLength"`
	// Number of uppercase and lowercase characters.
	ValidatePasswordMixedCaseCount pulumi.IntOutput `pulumi:"validatePasswordMixedCaseCount"`
	// Number of digits.
	ValidatePasswordNumberCount pulumi.IntOutput `pulumi:"validatePasswordNumberCount"`
	// Password strength (MEDIUM, STRONG).
	ValidatePasswordPolicy pulumi.StringOutput `pulumi:"validatePasswordPolicy"`
	// Number of special characters.
	ValidatePasswordSpecialCharCount pulumi.IntOutput `pulumi:"validatePasswordSpecialCharCount"`
}

// NewClusterPasswordComplexity registers a new resource with the given unique name, arguments, and options.
func NewClusterPasswordComplexity(ctx *pulumi.Context,
	name string, args *ClusterPasswordComplexityArgs, opts ...pulumi.ResourceOption) (*ClusterPasswordComplexity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ValidatePasswordLength == nil {
		return nil, errors.New("invalid value for required argument 'ValidatePasswordLength'")
	}
	if args.ValidatePasswordMixedCaseCount == nil {
		return nil, errors.New("invalid value for required argument 'ValidatePasswordMixedCaseCount'")
	}
	if args.ValidatePasswordNumberCount == nil {
		return nil, errors.New("invalid value for required argument 'ValidatePasswordNumberCount'")
	}
	if args.ValidatePasswordPolicy == nil {
		return nil, errors.New("invalid value for required argument 'ValidatePasswordPolicy'")
	}
	if args.ValidatePasswordSpecialCharCount == nil {
		return nil, errors.New("invalid value for required argument 'ValidatePasswordSpecialCharCount'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterPasswordComplexity
	err := ctx.RegisterResource("tencentcloud:Cynosdb/clusterPasswordComplexity:ClusterPasswordComplexity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterPasswordComplexity gets an existing ClusterPasswordComplexity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterPasswordComplexity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterPasswordComplexityState, opts ...pulumi.ResourceOption) (*ClusterPasswordComplexity, error) {
	var resource ClusterPasswordComplexity
	err := ctx.ReadResource("tencentcloud:Cynosdb/clusterPasswordComplexity:ClusterPasswordComplexity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterPasswordComplexity resources.
type clusterPasswordComplexityState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Data dictionary.
	ValidatePasswordDictionaries []string `pulumi:"validatePasswordDictionaries"`
	// Password length.
	ValidatePasswordLength *int `pulumi:"validatePasswordLength"`
	// Number of uppercase and lowercase characters.
	ValidatePasswordMixedCaseCount *int `pulumi:"validatePasswordMixedCaseCount"`
	// Number of digits.
	ValidatePasswordNumberCount *int `pulumi:"validatePasswordNumberCount"`
	// Password strength (MEDIUM, STRONG).
	ValidatePasswordPolicy *string `pulumi:"validatePasswordPolicy"`
	// Number of special characters.
	ValidatePasswordSpecialCharCount *int `pulumi:"validatePasswordSpecialCharCount"`
}

type ClusterPasswordComplexityState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Data dictionary.
	ValidatePasswordDictionaries pulumi.StringArrayInput
	// Password length.
	ValidatePasswordLength pulumi.IntPtrInput
	// Number of uppercase and lowercase characters.
	ValidatePasswordMixedCaseCount pulumi.IntPtrInput
	// Number of digits.
	ValidatePasswordNumberCount pulumi.IntPtrInput
	// Password strength (MEDIUM, STRONG).
	ValidatePasswordPolicy pulumi.StringPtrInput
	// Number of special characters.
	ValidatePasswordSpecialCharCount pulumi.IntPtrInput
}

func (ClusterPasswordComplexityState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPasswordComplexityState)(nil)).Elem()
}

type clusterPasswordComplexityArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Data dictionary.
	ValidatePasswordDictionaries []string `pulumi:"validatePasswordDictionaries"`
	// Password length.
	ValidatePasswordLength int `pulumi:"validatePasswordLength"`
	// Number of uppercase and lowercase characters.
	ValidatePasswordMixedCaseCount int `pulumi:"validatePasswordMixedCaseCount"`
	// Number of digits.
	ValidatePasswordNumberCount int `pulumi:"validatePasswordNumberCount"`
	// Password strength (MEDIUM, STRONG).
	ValidatePasswordPolicy string `pulumi:"validatePasswordPolicy"`
	// Number of special characters.
	ValidatePasswordSpecialCharCount int `pulumi:"validatePasswordSpecialCharCount"`
}

// The set of arguments for constructing a ClusterPasswordComplexity resource.
type ClusterPasswordComplexityArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Data dictionary.
	ValidatePasswordDictionaries pulumi.StringArrayInput
	// Password length.
	ValidatePasswordLength pulumi.IntInput
	// Number of uppercase and lowercase characters.
	ValidatePasswordMixedCaseCount pulumi.IntInput
	// Number of digits.
	ValidatePasswordNumberCount pulumi.IntInput
	// Password strength (MEDIUM, STRONG).
	ValidatePasswordPolicy pulumi.StringInput
	// Number of special characters.
	ValidatePasswordSpecialCharCount pulumi.IntInput
}

func (ClusterPasswordComplexityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterPasswordComplexityArgs)(nil)).Elem()
}

type ClusterPasswordComplexityInput interface {
	pulumi.Input

	ToClusterPasswordComplexityOutput() ClusterPasswordComplexityOutput
	ToClusterPasswordComplexityOutputWithContext(ctx context.Context) ClusterPasswordComplexityOutput
}

func (*ClusterPasswordComplexity) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPasswordComplexity)(nil)).Elem()
}

func (i *ClusterPasswordComplexity) ToClusterPasswordComplexityOutput() ClusterPasswordComplexityOutput {
	return i.ToClusterPasswordComplexityOutputWithContext(context.Background())
}

func (i *ClusterPasswordComplexity) ToClusterPasswordComplexityOutputWithContext(ctx context.Context) ClusterPasswordComplexityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPasswordComplexityOutput)
}

// ClusterPasswordComplexityArrayInput is an input type that accepts ClusterPasswordComplexityArray and ClusterPasswordComplexityArrayOutput values.
// You can construct a concrete instance of `ClusterPasswordComplexityArrayInput` via:
//
//	ClusterPasswordComplexityArray{ ClusterPasswordComplexityArgs{...} }
type ClusterPasswordComplexityArrayInput interface {
	pulumi.Input

	ToClusterPasswordComplexityArrayOutput() ClusterPasswordComplexityArrayOutput
	ToClusterPasswordComplexityArrayOutputWithContext(context.Context) ClusterPasswordComplexityArrayOutput
}

type ClusterPasswordComplexityArray []ClusterPasswordComplexityInput

func (ClusterPasswordComplexityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterPasswordComplexity)(nil)).Elem()
}

func (i ClusterPasswordComplexityArray) ToClusterPasswordComplexityArrayOutput() ClusterPasswordComplexityArrayOutput {
	return i.ToClusterPasswordComplexityArrayOutputWithContext(context.Background())
}

func (i ClusterPasswordComplexityArray) ToClusterPasswordComplexityArrayOutputWithContext(ctx context.Context) ClusterPasswordComplexityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPasswordComplexityArrayOutput)
}

// ClusterPasswordComplexityMapInput is an input type that accepts ClusterPasswordComplexityMap and ClusterPasswordComplexityMapOutput values.
// You can construct a concrete instance of `ClusterPasswordComplexityMapInput` via:
//
//	ClusterPasswordComplexityMap{ "key": ClusterPasswordComplexityArgs{...} }
type ClusterPasswordComplexityMapInput interface {
	pulumi.Input

	ToClusterPasswordComplexityMapOutput() ClusterPasswordComplexityMapOutput
	ToClusterPasswordComplexityMapOutputWithContext(context.Context) ClusterPasswordComplexityMapOutput
}

type ClusterPasswordComplexityMap map[string]ClusterPasswordComplexityInput

func (ClusterPasswordComplexityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterPasswordComplexity)(nil)).Elem()
}

func (i ClusterPasswordComplexityMap) ToClusterPasswordComplexityMapOutput() ClusterPasswordComplexityMapOutput {
	return i.ToClusterPasswordComplexityMapOutputWithContext(context.Background())
}

func (i ClusterPasswordComplexityMap) ToClusterPasswordComplexityMapOutputWithContext(ctx context.Context) ClusterPasswordComplexityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterPasswordComplexityMapOutput)
}

type ClusterPasswordComplexityOutput struct{ *pulumi.OutputState }

func (ClusterPasswordComplexityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPasswordComplexity)(nil)).Elem()
}

func (o ClusterPasswordComplexityOutput) ToClusterPasswordComplexityOutput() ClusterPasswordComplexityOutput {
	return o
}

func (o ClusterPasswordComplexityOutput) ToClusterPasswordComplexityOutputWithContext(ctx context.Context) ClusterPasswordComplexityOutput {
	return o
}

// Cluster ID.
func (o ClusterPasswordComplexityOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPasswordComplexity) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Data dictionary.
func (o ClusterPasswordComplexityOutput) ValidatePasswordDictionaries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ClusterPasswordComplexity) pulumi.StringArrayOutput { return v.ValidatePasswordDictionaries }).(pulumi.StringArrayOutput)
}

// Password length.
func (o ClusterPasswordComplexityOutput) ValidatePasswordLength() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterPasswordComplexity) pulumi.IntOutput { return v.ValidatePasswordLength }).(pulumi.IntOutput)
}

// Number of uppercase and lowercase characters.
func (o ClusterPasswordComplexityOutput) ValidatePasswordMixedCaseCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterPasswordComplexity) pulumi.IntOutput { return v.ValidatePasswordMixedCaseCount }).(pulumi.IntOutput)
}

// Number of digits.
func (o ClusterPasswordComplexityOutput) ValidatePasswordNumberCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterPasswordComplexity) pulumi.IntOutput { return v.ValidatePasswordNumberCount }).(pulumi.IntOutput)
}

// Password strength (MEDIUM, STRONG).
func (o ClusterPasswordComplexityOutput) ValidatePasswordPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterPasswordComplexity) pulumi.StringOutput { return v.ValidatePasswordPolicy }).(pulumi.StringOutput)
}

// Number of special characters.
func (o ClusterPasswordComplexityOutput) ValidatePasswordSpecialCharCount() pulumi.IntOutput {
	return o.ApplyT(func(v *ClusterPasswordComplexity) pulumi.IntOutput { return v.ValidatePasswordSpecialCharCount }).(pulumi.IntOutput)
}

type ClusterPasswordComplexityArrayOutput struct{ *pulumi.OutputState }

func (ClusterPasswordComplexityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterPasswordComplexity)(nil)).Elem()
}

func (o ClusterPasswordComplexityArrayOutput) ToClusterPasswordComplexityArrayOutput() ClusterPasswordComplexityArrayOutput {
	return o
}

func (o ClusterPasswordComplexityArrayOutput) ToClusterPasswordComplexityArrayOutputWithContext(ctx context.Context) ClusterPasswordComplexityArrayOutput {
	return o
}

func (o ClusterPasswordComplexityArrayOutput) Index(i pulumi.IntInput) ClusterPasswordComplexityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterPasswordComplexity {
		return vs[0].([]*ClusterPasswordComplexity)[vs[1].(int)]
	}).(ClusterPasswordComplexityOutput)
}

type ClusterPasswordComplexityMapOutput struct{ *pulumi.OutputState }

func (ClusterPasswordComplexityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterPasswordComplexity)(nil)).Elem()
}

func (o ClusterPasswordComplexityMapOutput) ToClusterPasswordComplexityMapOutput() ClusterPasswordComplexityMapOutput {
	return o
}

func (o ClusterPasswordComplexityMapOutput) ToClusterPasswordComplexityMapOutputWithContext(ctx context.Context) ClusterPasswordComplexityMapOutput {
	return o
}

func (o ClusterPasswordComplexityMapOutput) MapIndex(k pulumi.StringInput) ClusterPasswordComplexityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterPasswordComplexity {
		return vs[0].(map[string]*ClusterPasswordComplexity)[vs[1].(string)]
	}).(ClusterPasswordComplexityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPasswordComplexityInput)(nil)).Elem(), &ClusterPasswordComplexity{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPasswordComplexityArrayInput)(nil)).Elem(), ClusterPasswordComplexityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterPasswordComplexityMapInput)(nil)).Elem(), ClusterPasswordComplexityMap{})
	pulumi.RegisterOutputType(ClusterPasswordComplexityOutput{})
	pulumi.RegisterOutputType(ClusterPasswordComplexityArrayOutput{})
	pulumi.RegisterOutputType(ClusterPasswordComplexityMapOutput{})
}
