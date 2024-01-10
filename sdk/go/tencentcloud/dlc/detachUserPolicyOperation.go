// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dlc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DetachUserPolicyOperation struct {
	pulumi.CustomResourceState

	// Authentication policy collection.
	PolicySets DetachUserPolicyOperationPolicySetArrayOutput `pulumi:"policySets"`
	// User id, the same as the sub-user uin.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewDetachUserPolicyOperation registers a new resource with the given unique name, arguments, and options.
func NewDetachUserPolicyOperation(ctx *pulumi.Context,
	name string, args *DetachUserPolicyOperationArgs, opts ...pulumi.ResourceOption) (*DetachUserPolicyOperation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DetachUserPolicyOperation
	err := ctx.RegisterResource("tencentcloud:Dlc/detachUserPolicyOperation:DetachUserPolicyOperation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDetachUserPolicyOperation gets an existing DetachUserPolicyOperation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDetachUserPolicyOperation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DetachUserPolicyOperationState, opts ...pulumi.ResourceOption) (*DetachUserPolicyOperation, error) {
	var resource DetachUserPolicyOperation
	err := ctx.ReadResource("tencentcloud:Dlc/detachUserPolicyOperation:DetachUserPolicyOperation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DetachUserPolicyOperation resources.
type detachUserPolicyOperationState struct {
	// Authentication policy collection.
	PolicySets []DetachUserPolicyOperationPolicySet `pulumi:"policySets"`
	// User id, the same as the sub-user uin.
	UserId *string `pulumi:"userId"`
}

type DetachUserPolicyOperationState struct {
	// Authentication policy collection.
	PolicySets DetachUserPolicyOperationPolicySetArrayInput
	// User id, the same as the sub-user uin.
	UserId pulumi.StringPtrInput
}

func (DetachUserPolicyOperationState) ElementType() reflect.Type {
	return reflect.TypeOf((*detachUserPolicyOperationState)(nil)).Elem()
}

type detachUserPolicyOperationArgs struct {
	// Authentication policy collection.
	PolicySets []DetachUserPolicyOperationPolicySet `pulumi:"policySets"`
	// User id, the same as the sub-user uin.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a DetachUserPolicyOperation resource.
type DetachUserPolicyOperationArgs struct {
	// Authentication policy collection.
	PolicySets DetachUserPolicyOperationPolicySetArrayInput
	// User id, the same as the sub-user uin.
	UserId pulumi.StringInput
}

func (DetachUserPolicyOperationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*detachUserPolicyOperationArgs)(nil)).Elem()
}

type DetachUserPolicyOperationInput interface {
	pulumi.Input

	ToDetachUserPolicyOperationOutput() DetachUserPolicyOperationOutput
	ToDetachUserPolicyOperationOutputWithContext(ctx context.Context) DetachUserPolicyOperationOutput
}

func (*DetachUserPolicyOperation) ElementType() reflect.Type {
	return reflect.TypeOf((**DetachUserPolicyOperation)(nil)).Elem()
}

func (i *DetachUserPolicyOperation) ToDetachUserPolicyOperationOutput() DetachUserPolicyOperationOutput {
	return i.ToDetachUserPolicyOperationOutputWithContext(context.Background())
}

func (i *DetachUserPolicyOperation) ToDetachUserPolicyOperationOutputWithContext(ctx context.Context) DetachUserPolicyOperationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetachUserPolicyOperationOutput)
}

// DetachUserPolicyOperationArrayInput is an input type that accepts DetachUserPolicyOperationArray and DetachUserPolicyOperationArrayOutput values.
// You can construct a concrete instance of `DetachUserPolicyOperationArrayInput` via:
//
//	DetachUserPolicyOperationArray{ DetachUserPolicyOperationArgs{...} }
type DetachUserPolicyOperationArrayInput interface {
	pulumi.Input

	ToDetachUserPolicyOperationArrayOutput() DetachUserPolicyOperationArrayOutput
	ToDetachUserPolicyOperationArrayOutputWithContext(context.Context) DetachUserPolicyOperationArrayOutput
}

type DetachUserPolicyOperationArray []DetachUserPolicyOperationInput

func (DetachUserPolicyOperationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DetachUserPolicyOperation)(nil)).Elem()
}

func (i DetachUserPolicyOperationArray) ToDetachUserPolicyOperationArrayOutput() DetachUserPolicyOperationArrayOutput {
	return i.ToDetachUserPolicyOperationArrayOutputWithContext(context.Background())
}

func (i DetachUserPolicyOperationArray) ToDetachUserPolicyOperationArrayOutputWithContext(ctx context.Context) DetachUserPolicyOperationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetachUserPolicyOperationArrayOutput)
}

// DetachUserPolicyOperationMapInput is an input type that accepts DetachUserPolicyOperationMap and DetachUserPolicyOperationMapOutput values.
// You can construct a concrete instance of `DetachUserPolicyOperationMapInput` via:
//
//	DetachUserPolicyOperationMap{ "key": DetachUserPolicyOperationArgs{...} }
type DetachUserPolicyOperationMapInput interface {
	pulumi.Input

	ToDetachUserPolicyOperationMapOutput() DetachUserPolicyOperationMapOutput
	ToDetachUserPolicyOperationMapOutputWithContext(context.Context) DetachUserPolicyOperationMapOutput
}

type DetachUserPolicyOperationMap map[string]DetachUserPolicyOperationInput

func (DetachUserPolicyOperationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DetachUserPolicyOperation)(nil)).Elem()
}

func (i DetachUserPolicyOperationMap) ToDetachUserPolicyOperationMapOutput() DetachUserPolicyOperationMapOutput {
	return i.ToDetachUserPolicyOperationMapOutputWithContext(context.Background())
}

func (i DetachUserPolicyOperationMap) ToDetachUserPolicyOperationMapOutputWithContext(ctx context.Context) DetachUserPolicyOperationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetachUserPolicyOperationMapOutput)
}

type DetachUserPolicyOperationOutput struct{ *pulumi.OutputState }

func (DetachUserPolicyOperationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DetachUserPolicyOperation)(nil)).Elem()
}

func (o DetachUserPolicyOperationOutput) ToDetachUserPolicyOperationOutput() DetachUserPolicyOperationOutput {
	return o
}

func (o DetachUserPolicyOperationOutput) ToDetachUserPolicyOperationOutputWithContext(ctx context.Context) DetachUserPolicyOperationOutput {
	return o
}

// Authentication policy collection.
func (o DetachUserPolicyOperationOutput) PolicySets() DetachUserPolicyOperationPolicySetArrayOutput {
	return o.ApplyT(func(v *DetachUserPolicyOperation) DetachUserPolicyOperationPolicySetArrayOutput { return v.PolicySets }).(DetachUserPolicyOperationPolicySetArrayOutput)
}

// User id, the same as the sub-user uin.
func (o DetachUserPolicyOperationOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *DetachUserPolicyOperation) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type DetachUserPolicyOperationArrayOutput struct{ *pulumi.OutputState }

func (DetachUserPolicyOperationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DetachUserPolicyOperation)(nil)).Elem()
}

func (o DetachUserPolicyOperationArrayOutput) ToDetachUserPolicyOperationArrayOutput() DetachUserPolicyOperationArrayOutput {
	return o
}

func (o DetachUserPolicyOperationArrayOutput) ToDetachUserPolicyOperationArrayOutputWithContext(ctx context.Context) DetachUserPolicyOperationArrayOutput {
	return o
}

func (o DetachUserPolicyOperationArrayOutput) Index(i pulumi.IntInput) DetachUserPolicyOperationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DetachUserPolicyOperation {
		return vs[0].([]*DetachUserPolicyOperation)[vs[1].(int)]
	}).(DetachUserPolicyOperationOutput)
}

type DetachUserPolicyOperationMapOutput struct{ *pulumi.OutputState }

func (DetachUserPolicyOperationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DetachUserPolicyOperation)(nil)).Elem()
}

func (o DetachUserPolicyOperationMapOutput) ToDetachUserPolicyOperationMapOutput() DetachUserPolicyOperationMapOutput {
	return o
}

func (o DetachUserPolicyOperationMapOutput) ToDetachUserPolicyOperationMapOutputWithContext(ctx context.Context) DetachUserPolicyOperationMapOutput {
	return o
}

func (o DetachUserPolicyOperationMapOutput) MapIndex(k pulumi.StringInput) DetachUserPolicyOperationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DetachUserPolicyOperation {
		return vs[0].(map[string]*DetachUserPolicyOperation)[vs[1].(string)]
	}).(DetachUserPolicyOperationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DetachUserPolicyOperationInput)(nil)).Elem(), &DetachUserPolicyOperation{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetachUserPolicyOperationArrayInput)(nil)).Elem(), DetachUserPolicyOperationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetachUserPolicyOperationMapInput)(nil)).Elem(), DetachUserPolicyOperationMap{})
	pulumi.RegisterOutputType(DetachUserPolicyOperationOutput{})
	pulumi.RegisterOutputType(DetachUserPolicyOperationArrayOutput{})
	pulumi.RegisterOutputType(DetachUserPolicyOperationMapOutput{})
}
