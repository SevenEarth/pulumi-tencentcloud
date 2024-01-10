// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cdn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provide a resource to invoke a Url Purge Request.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cdn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cdn.NewUrlPurge(ctx, "foo", &Cdn.UrlPurgeArgs{
//				Urls: pulumi.StringArray{
//					pulumi.String("https://www.example.com/a"),
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
// ### argument to request new purge task with same urls
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cdn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cdn.NewUrlPurge(ctx, "foo", &Cdn.UrlPurgeArgs{
//				Redo: pulumi.Int(1),
//				Urls: pulumi.StringArray{
//					pulumi.String("https://www.example.com/a"),
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
type UrlPurge struct {
	pulumi.CustomResourceState

	// Specify purge area. NOTE: only purge same area cache contents.
	Area pulumi.StringPtrOutput `pulumi:"area"`
	// logs of latest purge task.
	PurgeHistories UrlPurgePurgeHistoryArrayOutput `pulumi:"purgeHistories"`
	// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
	Redo pulumi.IntPtrOutput `pulumi:"redo"`
	// Task id of last operation.
	TaskId pulumi.StringOutput `pulumi:"taskId"`
	// Whether to encode urls, if set to `true` will auto encode instead of manual process.
	UrlEncode pulumi.BoolPtrOutput `pulumi:"urlEncode"`
	// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
	Urls pulumi.StringArrayOutput `pulumi:"urls"`
}

// NewUrlPurge registers a new resource with the given unique name, arguments, and options.
func NewUrlPurge(ctx *pulumi.Context,
	name string, args *UrlPurgeArgs, opts ...pulumi.ResourceOption) (*UrlPurge, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Urls == nil {
		return nil, errors.New("invalid value for required argument 'Urls'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UrlPurge
	err := ctx.RegisterResource("tencentcloud:Cdn/urlPurge:UrlPurge", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUrlPurge gets an existing UrlPurge resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUrlPurge(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UrlPurgeState, opts ...pulumi.ResourceOption) (*UrlPurge, error) {
	var resource UrlPurge
	err := ctx.ReadResource("tencentcloud:Cdn/urlPurge:UrlPurge", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UrlPurge resources.
type urlPurgeState struct {
	// Specify purge area. NOTE: only purge same area cache contents.
	Area *string `pulumi:"area"`
	// logs of latest purge task.
	PurgeHistories []UrlPurgePurgeHistory `pulumi:"purgeHistories"`
	// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
	Redo *int `pulumi:"redo"`
	// Task id of last operation.
	TaskId *string `pulumi:"taskId"`
	// Whether to encode urls, if set to `true` will auto encode instead of manual process.
	UrlEncode *bool `pulumi:"urlEncode"`
	// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
	Urls []string `pulumi:"urls"`
}

type UrlPurgeState struct {
	// Specify purge area. NOTE: only purge same area cache contents.
	Area pulumi.StringPtrInput
	// logs of latest purge task.
	PurgeHistories UrlPurgePurgeHistoryArrayInput
	// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
	Redo pulumi.IntPtrInput
	// Task id of last operation.
	TaskId pulumi.StringPtrInput
	// Whether to encode urls, if set to `true` will auto encode instead of manual process.
	UrlEncode pulumi.BoolPtrInput
	// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
	Urls pulumi.StringArrayInput
}

func (UrlPurgeState) ElementType() reflect.Type {
	return reflect.TypeOf((*urlPurgeState)(nil)).Elem()
}

type urlPurgeArgs struct {
	// Specify purge area. NOTE: only purge same area cache contents.
	Area *string `pulumi:"area"`
	// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
	Redo *int `pulumi:"redo"`
	// Whether to encode urls, if set to `true` will auto encode instead of manual process.
	UrlEncode *bool `pulumi:"urlEncode"`
	// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
	Urls []string `pulumi:"urls"`
}

// The set of arguments for constructing a UrlPurge resource.
type UrlPurgeArgs struct {
	// Specify purge area. NOTE: only purge same area cache contents.
	Area pulumi.StringPtrInput
	// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
	Redo pulumi.IntPtrInput
	// Whether to encode urls, if set to `true` will auto encode instead of manual process.
	UrlEncode pulumi.BoolPtrInput
	// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
	Urls pulumi.StringArrayInput
}

func (UrlPurgeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*urlPurgeArgs)(nil)).Elem()
}

type UrlPurgeInput interface {
	pulumi.Input

	ToUrlPurgeOutput() UrlPurgeOutput
	ToUrlPurgeOutputWithContext(ctx context.Context) UrlPurgeOutput
}

func (*UrlPurge) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlPurge)(nil)).Elem()
}

func (i *UrlPurge) ToUrlPurgeOutput() UrlPurgeOutput {
	return i.ToUrlPurgeOutputWithContext(context.Background())
}

func (i *UrlPurge) ToUrlPurgeOutputWithContext(ctx context.Context) UrlPurgeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlPurgeOutput)
}

// UrlPurgeArrayInput is an input type that accepts UrlPurgeArray and UrlPurgeArrayOutput values.
// You can construct a concrete instance of `UrlPurgeArrayInput` via:
//
//	UrlPurgeArray{ UrlPurgeArgs{...} }
type UrlPurgeArrayInput interface {
	pulumi.Input

	ToUrlPurgeArrayOutput() UrlPurgeArrayOutput
	ToUrlPurgeArrayOutputWithContext(context.Context) UrlPurgeArrayOutput
}

type UrlPurgeArray []UrlPurgeInput

func (UrlPurgeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UrlPurge)(nil)).Elem()
}

func (i UrlPurgeArray) ToUrlPurgeArrayOutput() UrlPurgeArrayOutput {
	return i.ToUrlPurgeArrayOutputWithContext(context.Background())
}

func (i UrlPurgeArray) ToUrlPurgeArrayOutputWithContext(ctx context.Context) UrlPurgeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlPurgeArrayOutput)
}

// UrlPurgeMapInput is an input type that accepts UrlPurgeMap and UrlPurgeMapOutput values.
// You can construct a concrete instance of `UrlPurgeMapInput` via:
//
//	UrlPurgeMap{ "key": UrlPurgeArgs{...} }
type UrlPurgeMapInput interface {
	pulumi.Input

	ToUrlPurgeMapOutput() UrlPurgeMapOutput
	ToUrlPurgeMapOutputWithContext(context.Context) UrlPurgeMapOutput
}

type UrlPurgeMap map[string]UrlPurgeInput

func (UrlPurgeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UrlPurge)(nil)).Elem()
}

func (i UrlPurgeMap) ToUrlPurgeMapOutput() UrlPurgeMapOutput {
	return i.ToUrlPurgeMapOutputWithContext(context.Background())
}

func (i UrlPurgeMap) ToUrlPurgeMapOutputWithContext(ctx context.Context) UrlPurgeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UrlPurgeMapOutput)
}

type UrlPurgeOutput struct{ *pulumi.OutputState }

func (UrlPurgeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UrlPurge)(nil)).Elem()
}

func (o UrlPurgeOutput) ToUrlPurgeOutput() UrlPurgeOutput {
	return o
}

func (o UrlPurgeOutput) ToUrlPurgeOutputWithContext(ctx context.Context) UrlPurgeOutput {
	return o
}

// Specify purge area. NOTE: only purge same area cache contents.
func (o UrlPurgeOutput) Area() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UrlPurge) pulumi.StringPtrOutput { return v.Area }).(pulumi.StringPtrOutput)
}

// logs of latest purge task.
func (o UrlPurgeOutput) PurgeHistories() UrlPurgePurgeHistoryArrayOutput {
	return o.ApplyT(func(v *UrlPurge) UrlPurgePurgeHistoryArrayOutput { return v.PurgeHistories }).(UrlPurgePurgeHistoryArrayOutput)
}

// Change to purge again. NOTE: this argument only works while resource update, if set to `0` or null will not be triggered.
func (o UrlPurgeOutput) Redo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UrlPurge) pulumi.IntPtrOutput { return v.Redo }).(pulumi.IntPtrOutput)
}

// Task id of last operation.
func (o UrlPurgeOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *UrlPurge) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

// Whether to encode urls, if set to `true` will auto encode instead of manual process.
func (o UrlPurgeOutput) UrlEncode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UrlPurge) pulumi.BoolPtrOutput { return v.UrlEncode }).(pulumi.BoolPtrOutput)
}

// List of url to purge. NOTE: urls need include protocol prefix `http://` or `https://`.
func (o UrlPurgeOutput) Urls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *UrlPurge) pulumi.StringArrayOutput { return v.Urls }).(pulumi.StringArrayOutput)
}

type UrlPurgeArrayOutput struct{ *pulumi.OutputState }

func (UrlPurgeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UrlPurge)(nil)).Elem()
}

func (o UrlPurgeArrayOutput) ToUrlPurgeArrayOutput() UrlPurgeArrayOutput {
	return o
}

func (o UrlPurgeArrayOutput) ToUrlPurgeArrayOutputWithContext(ctx context.Context) UrlPurgeArrayOutput {
	return o
}

func (o UrlPurgeArrayOutput) Index(i pulumi.IntInput) UrlPurgeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UrlPurge {
		return vs[0].([]*UrlPurge)[vs[1].(int)]
	}).(UrlPurgeOutput)
}

type UrlPurgeMapOutput struct{ *pulumi.OutputState }

func (UrlPurgeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UrlPurge)(nil)).Elem()
}

func (o UrlPurgeMapOutput) ToUrlPurgeMapOutput() UrlPurgeMapOutput {
	return o
}

func (o UrlPurgeMapOutput) ToUrlPurgeMapOutputWithContext(ctx context.Context) UrlPurgeMapOutput {
	return o
}

func (o UrlPurgeMapOutput) MapIndex(k pulumi.StringInput) UrlPurgeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UrlPurge {
		return vs[0].(map[string]*UrlPurge)[vs[1].(string)]
	}).(UrlPurgeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UrlPurgeInput)(nil)).Elem(), &UrlPurge{})
	pulumi.RegisterInputType(reflect.TypeOf((*UrlPurgeArrayInput)(nil)).Elem(), UrlPurgeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UrlPurgeMapInput)(nil)).Elem(), UrlPurgeMap{})
	pulumi.RegisterOutputType(UrlPurgeOutput{})
	pulumi.RegisterOutputType(UrlPurgeArrayOutput{})
	pulumi.RegisterOutputType(UrlPurgeMapOutput{})
}
