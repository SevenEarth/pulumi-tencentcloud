// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a eb ebTransform
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"encoding/json"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Eb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Eb"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		fooEventBus, err := Eb.NewEventBus(ctx, "fooEventBus", &Eb.EventBusArgs{
// 			EventBusName: pulumi.String("tf-event_bus"),
// 			Description:  pulumi.String("event bus desc"),
// 			EnableStore:  pulumi.Bool(false),
// 			SaveDays:     pulumi.Int(1),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON0, err := json.Marshal(map[string]interface{}{
// 			"source": "apigw.cloud.tencent",
// 			"type": []string{
// 				"connector:apigw",
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json0 := string(tmpJSON0)
// 		fooEventRule, err := Eb.NewEventRule(ctx, "fooEventRule", &Eb.EventRuleArgs{
// 			EventBusId:   fooEventBus.ID(),
// 			RuleName:     pulumi.String("tf-event_rule"),
// 			Description:  pulumi.String("event rule desc"),
// 			Enable:       pulumi.Bool(true),
// 			EventPattern: pulumi.String(json0),
// 			Tags: pulumi.AnyMap{
// 				"createdBy": pulumi.Any("terraform"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		tmpJSON1, err := json.Marshal(map[string]interface{}{
// 			"Partition": 1,
// 			"msgBody":   "Hello from Ckafka again!",
// 			"msgKey":    "test",
// 			"offset":    37,
// 			"topic":     "test-topic",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		json1 := string(tmpJSON1)
// 		_, err = Eb.NewEventTransform(ctx, "fooEventTransform", &Eb.EventTransformArgs{
// 			EventBusId: fooEventBus.ID(),
// 			RuleId:     fooEventRule.RuleId,
// 			Transformations: eb.EventTransformTransformationArray{
// 				&eb.EventTransformTransformationArgs{
// 					Extraction: &eb.EventTransformTransformationExtractionArgs{
// 						ExtractionInputPath: pulumi.String("$"),
// 						Format:              pulumi.String("JSON"),
// 					},
// 					Transform: &eb.EventTransformTransformationTransformArgs{
// 						OutputStructs: eb.EventTransformTransformationTransformOutputStructArray{
// 							&eb.EventTransformTransformationTransformOutputStructArgs{
// 								Key:       pulumi.String("type"),
// 								Value:     pulumi.String("connector:ckafka"),
// 								ValueType: pulumi.String("STRING"),
// 							},
// 							&eb.EventTransformTransformationTransformOutputStructArgs{
// 								Key:       pulumi.String("source"),
// 								Value:     pulumi.String("ckafka.cloud.tencent"),
// 								ValueType: pulumi.String("STRING"),
// 							},
// 							&eb.EventTransformTransformationTransformOutputStructArgs{
// 								Key:       pulumi.String("region"),
// 								Value:     pulumi.String("ap-guangzhou"),
// 								ValueType: pulumi.String("STRING"),
// 							},
// 							&eb.EventTransformTransformationTransformOutputStructArgs{
// 								Key:       pulumi.String("datacontenttype"),
// 								Value:     pulumi.String("application/json;charset=utf-8"),
// 								ValueType: pulumi.String("STRING"),
// 							},
// 							&eb.EventTransformTransformationTransformOutputStructArgs{
// 								Key:       pulumi.String("status"),
// 								Value:     pulumi.String("-"),
// 								ValueType: pulumi.String("STRING"),
// 							},
// 							&eb.EventTransformTransformationTransformOutputStructArgs{
// 								Key:       pulumi.String("data"),
// 								Value:     pulumi.String(json1),
// 								ValueType: pulumi.String("STRING"),
// 							},
// 						},
// 					},
// 				},
// 			},
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
// eb eb_transform can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Eb/eventTransform:EventTransform eb_transform eb_transform_id
// ```
type EventTransform struct {
	pulumi.CustomResourceState

	// event bus Id.
	EventBusId pulumi.StringOutput `pulumi:"eventBusId"`
	// ruleId.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// A list of transformation rules, currently only one.
	Transformations EventTransformTransformationArrayOutput `pulumi:"transformations"`
}

// NewEventTransform registers a new resource with the given unique name, arguments, and options.
func NewEventTransform(ctx *pulumi.Context,
	name string, args *EventTransformArgs, opts ...pulumi.ResourceOption) (*EventTransform, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventBusId == nil {
		return nil, errors.New("invalid value for required argument 'EventBusId'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	if args.Transformations == nil {
		return nil, errors.New("invalid value for required argument 'Transformations'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource EventTransform
	err := ctx.RegisterResource("tencentcloud:Eb/eventTransform:EventTransform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventTransform gets an existing EventTransform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventTransform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventTransformState, opts ...pulumi.ResourceOption) (*EventTransform, error) {
	var resource EventTransform
	err := ctx.ReadResource("tencentcloud:Eb/eventTransform:EventTransform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventTransform resources.
type eventTransformState struct {
	// event bus Id.
	EventBusId *string `pulumi:"eventBusId"`
	// ruleId.
	RuleId *string `pulumi:"ruleId"`
	// A list of transformation rules, currently only one.
	Transformations []EventTransformTransformation `pulumi:"transformations"`
}

type EventTransformState struct {
	// event bus Id.
	EventBusId pulumi.StringPtrInput
	// ruleId.
	RuleId pulumi.StringPtrInput
	// A list of transformation rules, currently only one.
	Transformations EventTransformTransformationArrayInput
}

func (EventTransformState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTransformState)(nil)).Elem()
}

type eventTransformArgs struct {
	// event bus Id.
	EventBusId string `pulumi:"eventBusId"`
	// ruleId.
	RuleId string `pulumi:"ruleId"`
	// A list of transformation rules, currently only one.
	Transformations []EventTransformTransformation `pulumi:"transformations"`
}

// The set of arguments for constructing a EventTransform resource.
type EventTransformArgs struct {
	// event bus Id.
	EventBusId pulumi.StringInput
	// ruleId.
	RuleId pulumi.StringInput
	// A list of transformation rules, currently only one.
	Transformations EventTransformTransformationArrayInput
}

func (EventTransformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventTransformArgs)(nil)).Elem()
}

type EventTransformInput interface {
	pulumi.Input

	ToEventTransformOutput() EventTransformOutput
	ToEventTransformOutputWithContext(ctx context.Context) EventTransformOutput
}

func (*EventTransform) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTransform)(nil)).Elem()
}

func (i *EventTransform) ToEventTransformOutput() EventTransformOutput {
	return i.ToEventTransformOutputWithContext(context.Background())
}

func (i *EventTransform) ToEventTransformOutputWithContext(ctx context.Context) EventTransformOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTransformOutput)
}

// EventTransformArrayInput is an input type that accepts EventTransformArray and EventTransformArrayOutput values.
// You can construct a concrete instance of `EventTransformArrayInput` via:
//
//          EventTransformArray{ EventTransformArgs{...} }
type EventTransformArrayInput interface {
	pulumi.Input

	ToEventTransformArrayOutput() EventTransformArrayOutput
	ToEventTransformArrayOutputWithContext(context.Context) EventTransformArrayOutput
}

type EventTransformArray []EventTransformInput

func (EventTransformArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventTransform)(nil)).Elem()
}

func (i EventTransformArray) ToEventTransformArrayOutput() EventTransformArrayOutput {
	return i.ToEventTransformArrayOutputWithContext(context.Background())
}

func (i EventTransformArray) ToEventTransformArrayOutputWithContext(ctx context.Context) EventTransformArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTransformArrayOutput)
}

// EventTransformMapInput is an input type that accepts EventTransformMap and EventTransformMapOutput values.
// You can construct a concrete instance of `EventTransformMapInput` via:
//
//          EventTransformMap{ "key": EventTransformArgs{...} }
type EventTransformMapInput interface {
	pulumi.Input

	ToEventTransformMapOutput() EventTransformMapOutput
	ToEventTransformMapOutputWithContext(context.Context) EventTransformMapOutput
}

type EventTransformMap map[string]EventTransformInput

func (EventTransformMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventTransform)(nil)).Elem()
}

func (i EventTransformMap) ToEventTransformMapOutput() EventTransformMapOutput {
	return i.ToEventTransformMapOutputWithContext(context.Background())
}

func (i EventTransformMap) ToEventTransformMapOutputWithContext(ctx context.Context) EventTransformMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventTransformMapOutput)
}

type EventTransformOutput struct{ *pulumi.OutputState }

func (EventTransformOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventTransform)(nil)).Elem()
}

func (o EventTransformOutput) ToEventTransformOutput() EventTransformOutput {
	return o
}

func (o EventTransformOutput) ToEventTransformOutputWithContext(ctx context.Context) EventTransformOutput {
	return o
}

// event bus Id.
func (o EventTransformOutput) EventBusId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventTransform) pulumi.StringOutput { return v.EventBusId }).(pulumi.StringOutput)
}

// ruleId.
func (o EventTransformOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventTransform) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// A list of transformation rules, currently only one.
func (o EventTransformOutput) Transformations() EventTransformTransformationArrayOutput {
	return o.ApplyT(func(v *EventTransform) EventTransformTransformationArrayOutput { return v.Transformations }).(EventTransformTransformationArrayOutput)
}

type EventTransformArrayOutput struct{ *pulumi.OutputState }

func (EventTransformArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EventTransform)(nil)).Elem()
}

func (o EventTransformArrayOutput) ToEventTransformArrayOutput() EventTransformArrayOutput {
	return o
}

func (o EventTransformArrayOutput) ToEventTransformArrayOutputWithContext(ctx context.Context) EventTransformArrayOutput {
	return o
}

func (o EventTransformArrayOutput) Index(i pulumi.IntInput) EventTransformOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EventTransform {
		return vs[0].([]*EventTransform)[vs[1].(int)]
	}).(EventTransformOutput)
}

type EventTransformMapOutput struct{ *pulumi.OutputState }

func (EventTransformMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EventTransform)(nil)).Elem()
}

func (o EventTransformMapOutput) ToEventTransformMapOutput() EventTransformMapOutput {
	return o
}

func (o EventTransformMapOutput) ToEventTransformMapOutputWithContext(ctx context.Context) EventTransformMapOutput {
	return o
}

func (o EventTransformMapOutput) MapIndex(k pulumi.StringInput) EventTransformOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EventTransform {
		return vs[0].(map[string]*EventTransform)[vs[1].(string)]
	}).(EventTransformOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventTransformInput)(nil)).Elem(), &EventTransform{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventTransformArrayInput)(nil)).Elem(), EventTransformArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventTransformMapInput)(nil)).Elem(), EventTransformMap{})
	pulumi.RegisterOutputType(EventTransformOutput{})
	pulumi.RegisterOutputType(EventTransformArrayOutput{})
	pulumi.RegisterOutputType(EventTransformMapOutput{})
}
