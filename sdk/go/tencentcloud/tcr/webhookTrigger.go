// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tcr

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a tcr webhook trigger
//
// ## Example Usage
//
// ### Create a tcr webhook trigger instance
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Tcr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleInstance, err := Tcr.NewInstance(ctx, "exampleInstance", &Tcr.InstanceArgs{
//				InstanceType: pulumi.String("basic"),
//				DeleteBucket: pulumi.Bool(true),
//				Tags: pulumi.Map{
//					"test": pulumi.Any("test"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleNamespace, err := Tcr.NewNamespace(ctx, "exampleNamespace", &Tcr.NamespaceArgs{
//				InstanceId:   exampleInstance.ID(),
//				IsPublic:     pulumi.Bool(true),
//				IsAutoScan:   pulumi.Bool(true),
//				IsPreventVul: pulumi.Bool(true),
//				Severity:     pulumi.String("medium"),
//				CveWhitelistItems: tcr.NamespaceCveWhitelistItemArray{
//					&tcr.NamespaceCveWhitelistItemArgs{
//						CveId: pulumi.String("cve-xxxxx"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleNamespaces := Tcr.GetNamespacesOutput(ctx, tcr.GetNamespacesOutputArgs{
//				InstanceId: exampleNamespace.InstanceId,
//			}, nil)
//			nsId := exampleNamespaces.ApplyT(func(exampleNamespaces tcr.GetNamespacesResult) (*int, error) {
//				return &exampleNamespaces.NamespaceLists[0].Id, nil
//			}).(pulumi.IntPtrOutput)
//			_, err = Tcr.NewWebhookTrigger(ctx, "exampleWebhookTrigger", &Tcr.WebhookTriggerArgs{
//				RegistryId: exampleInstance.ID(),
//				Namespace:  exampleNamespace.Name,
//				Trigger: &tcr.WebhookTriggerTriggerArgs{
//					Name: pulumi.String("trigger-example"),
//					Targets: tcr.WebhookTriggerTriggerTargetArray{
//						&tcr.WebhookTriggerTriggerTargetArgs{
//							Address: pulumi.String("http://example.org/post"),
//							Headers: tcr.WebhookTriggerTriggerTargetHeaderArray{
//								&tcr.WebhookTriggerTriggerTargetHeaderArgs{
//									Key: pulumi.String("X-Custom-Header"),
//									Values: pulumi.StringArray{
//										pulumi.String("a"),
//									},
//								},
//							},
//						},
//					},
//					EventTypes: pulumi.StringArray{
//						pulumi.String("pushImage"),
//					},
//					Condition:   pulumi.String(".*"),
//					Enabled:     pulumi.Bool(true),
//					Description: pulumi.String("example for trigger description"),
//					NamespaceId: pulumi.Int(nsId),
//				},
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
// ## Import
//
// tcr webhook_trigger can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Tcr/webhookTrigger:WebhookTrigger example webhook_trigger_id
// ```
type WebhookTrigger struct {
	pulumi.CustomResourceState

	// namespace name.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// instance Id.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Tag description list.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// trigger parameters.
	Trigger WebhookTriggerTriggerOutput `pulumi:"trigger"`
}

// NewWebhookTrigger registers a new resource with the given unique name, arguments, and options.
func NewWebhookTrigger(ctx *pulumi.Context,
	name string, args *WebhookTriggerArgs, opts ...pulumi.ResourceOption) (*WebhookTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.Trigger == nil {
		return nil, errors.New("invalid value for required argument 'Trigger'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WebhookTrigger
	err := ctx.RegisterResource("tencentcloud:Tcr/webhookTrigger:WebhookTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhookTrigger gets an existing WebhookTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhookTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookTriggerState, opts ...pulumi.ResourceOption) (*WebhookTrigger, error) {
	var resource WebhookTrigger
	err := ctx.ReadResource("tencentcloud:Tcr/webhookTrigger:WebhookTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebhookTrigger resources.
type webhookTriggerState struct {
	// namespace name.
	Namespace *string `pulumi:"namespace"`
	// instance Id.
	RegistryId *string `pulumi:"registryId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// trigger parameters.
	Trigger *WebhookTriggerTrigger `pulumi:"trigger"`
}

type WebhookTriggerState struct {
	// namespace name.
	Namespace pulumi.StringPtrInput
	// instance Id.
	RegistryId pulumi.StringPtrInput
	// Tag description list.
	Tags pulumi.MapInput
	// trigger parameters.
	Trigger WebhookTriggerTriggerPtrInput
}

func (WebhookTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookTriggerState)(nil)).Elem()
}

type webhookTriggerArgs struct {
	// namespace name.
	Namespace string `pulumi:"namespace"`
	// instance Id.
	RegistryId string `pulumi:"registryId"`
	// Tag description list.
	Tags map[string]interface{} `pulumi:"tags"`
	// trigger parameters.
	Trigger WebhookTriggerTrigger `pulumi:"trigger"`
}

// The set of arguments for constructing a WebhookTrigger resource.
type WebhookTriggerArgs struct {
	// namespace name.
	Namespace pulumi.StringInput
	// instance Id.
	RegistryId pulumi.StringInput
	// Tag description list.
	Tags pulumi.MapInput
	// trigger parameters.
	Trigger WebhookTriggerTriggerInput
}

func (WebhookTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookTriggerArgs)(nil)).Elem()
}

type WebhookTriggerInput interface {
	pulumi.Input

	ToWebhookTriggerOutput() WebhookTriggerOutput
	ToWebhookTriggerOutputWithContext(ctx context.Context) WebhookTriggerOutput
}

func (*WebhookTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookTrigger)(nil)).Elem()
}

func (i *WebhookTrigger) ToWebhookTriggerOutput() WebhookTriggerOutput {
	return i.ToWebhookTriggerOutputWithContext(context.Background())
}

func (i *WebhookTrigger) ToWebhookTriggerOutputWithContext(ctx context.Context) WebhookTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookTriggerOutput)
}

// WebhookTriggerArrayInput is an input type that accepts WebhookTriggerArray and WebhookTriggerArrayOutput values.
// You can construct a concrete instance of `WebhookTriggerArrayInput` via:
//
//	WebhookTriggerArray{ WebhookTriggerArgs{...} }
type WebhookTriggerArrayInput interface {
	pulumi.Input

	ToWebhookTriggerArrayOutput() WebhookTriggerArrayOutput
	ToWebhookTriggerArrayOutputWithContext(context.Context) WebhookTriggerArrayOutput
}

type WebhookTriggerArray []WebhookTriggerInput

func (WebhookTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookTrigger)(nil)).Elem()
}

func (i WebhookTriggerArray) ToWebhookTriggerArrayOutput() WebhookTriggerArrayOutput {
	return i.ToWebhookTriggerArrayOutputWithContext(context.Background())
}

func (i WebhookTriggerArray) ToWebhookTriggerArrayOutputWithContext(ctx context.Context) WebhookTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookTriggerArrayOutput)
}

// WebhookTriggerMapInput is an input type that accepts WebhookTriggerMap and WebhookTriggerMapOutput values.
// You can construct a concrete instance of `WebhookTriggerMapInput` via:
//
//	WebhookTriggerMap{ "key": WebhookTriggerArgs{...} }
type WebhookTriggerMapInput interface {
	pulumi.Input

	ToWebhookTriggerMapOutput() WebhookTriggerMapOutput
	ToWebhookTriggerMapOutputWithContext(context.Context) WebhookTriggerMapOutput
}

type WebhookTriggerMap map[string]WebhookTriggerInput

func (WebhookTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookTrigger)(nil)).Elem()
}

func (i WebhookTriggerMap) ToWebhookTriggerMapOutput() WebhookTriggerMapOutput {
	return i.ToWebhookTriggerMapOutputWithContext(context.Background())
}

func (i WebhookTriggerMap) ToWebhookTriggerMapOutputWithContext(ctx context.Context) WebhookTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookTriggerMapOutput)
}

type WebhookTriggerOutput struct{ *pulumi.OutputState }

func (WebhookTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookTrigger)(nil)).Elem()
}

func (o WebhookTriggerOutput) ToWebhookTriggerOutput() WebhookTriggerOutput {
	return o
}

func (o WebhookTriggerOutput) ToWebhookTriggerOutputWithContext(ctx context.Context) WebhookTriggerOutput {
	return o
}

// namespace name.
func (o WebhookTriggerOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *WebhookTrigger) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// instance Id.
func (o WebhookTriggerOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *WebhookTrigger) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// Tag description list.
func (o WebhookTriggerOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *WebhookTrigger) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// trigger parameters.
func (o WebhookTriggerOutput) Trigger() WebhookTriggerTriggerOutput {
	return o.ApplyT(func(v *WebhookTrigger) WebhookTriggerTriggerOutput { return v.Trigger }).(WebhookTriggerTriggerOutput)
}

type WebhookTriggerArrayOutput struct{ *pulumi.OutputState }

func (WebhookTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WebhookTrigger)(nil)).Elem()
}

func (o WebhookTriggerArrayOutput) ToWebhookTriggerArrayOutput() WebhookTriggerArrayOutput {
	return o
}

func (o WebhookTriggerArrayOutput) ToWebhookTriggerArrayOutputWithContext(ctx context.Context) WebhookTriggerArrayOutput {
	return o
}

func (o WebhookTriggerArrayOutput) Index(i pulumi.IntInput) WebhookTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WebhookTrigger {
		return vs[0].([]*WebhookTrigger)[vs[1].(int)]
	}).(WebhookTriggerOutput)
}

type WebhookTriggerMapOutput struct{ *pulumi.OutputState }

func (WebhookTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WebhookTrigger)(nil)).Elem()
}

func (o WebhookTriggerMapOutput) ToWebhookTriggerMapOutput() WebhookTriggerMapOutput {
	return o
}

func (o WebhookTriggerMapOutput) ToWebhookTriggerMapOutputWithContext(ctx context.Context) WebhookTriggerMapOutput {
	return o
}

func (o WebhookTriggerMapOutput) MapIndex(k pulumi.StringInput) WebhookTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WebhookTrigger {
		return vs[0].(map[string]*WebhookTrigger)[vs[1].(string)]
	}).(WebhookTriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookTriggerInput)(nil)).Elem(), &WebhookTrigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookTriggerArrayInput)(nil)).Elem(), WebhookTriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhookTriggerMapInput)(nil)).Elem(), WebhookTriggerMap{})
	pulumi.RegisterOutputType(WebhookTriggerOutput{})
	pulumi.RegisterOutputType(WebhookTriggerArrayOutput{})
	pulumi.RegisterOutputType(WebhookTriggerMapOutput{})
}
