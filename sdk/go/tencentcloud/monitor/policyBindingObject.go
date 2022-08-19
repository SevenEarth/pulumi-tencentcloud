// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for bind objects to a alarm policy resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Instances"
// 	"github.com/pulumi/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Instances"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Monitor"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		instances, err := Instances.GetInstance(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		policy, err := Monitor.NewAlarmPolicy(ctx, "policy", &Monitor.AlarmPolicyArgs{
// 			PolicyName:  pulumi.String("hello"),
// 			MonitorType: pulumi.String("MT_QCE"),
// 			Enable:      pulumi.Int(1),
// 			ProjectId:   pulumi.Int(1244035),
// 			Namespace:   pulumi.String("cvm_device"),
// 			Conditions: &monitor.AlarmPolicyConditionsArgs{
// 				IsUnionRule: pulumi.Int(1),
// 				Rules: monitor.AlarmPolicyConditionsRuleArray{
// 					&monitor.AlarmPolicyConditionsRuleArgs{
// 						MetricName:      pulumi.String("CpuUsage"),
// 						Period:          pulumi.Int(60),
// 						Operator:        pulumi.String("ge"),
// 						Value:           pulumi.String("89.9"),
// 						ContinuePeriod:  pulumi.Int(1),
// 						NoticeFrequency: pulumi.Int(3600),
// 						IsPowerNotice:   pulumi.Int(0),
// 					},
// 				},
// 			},
// 			EventConditions: monitor.AlarmPolicyEventConditionArray{
// 				&monitor.AlarmPolicyEventConditionArgs{
// 					MetricName: pulumi.String("ping_unreachable"),
// 				},
// 				&monitor.AlarmPolicyEventConditionArgs{
// 					MetricName: pulumi.String("guest_reboot"),
// 				},
// 			},
// 			NoticeIds: pulumi.StringArray{
// 				pulumi.String("notice-l9ziyxw6"),
// 			},
// 			TriggerTasks: monitor.AlarmPolicyTriggerTaskArray{
// 				&monitor.AlarmPolicyTriggerTaskArgs{
// 					Type:       pulumi.String("AS"),
// 					TaskConfig: pulumi.String("{\"Region\":\"ap-guangzhou\",\"Group\":\"asg-0z312312x\",\"Policy\":\"asp-ganig28\"}"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Monitor.NewPolicyBindingObject(ctx, "binding", &Monitor.PolicyBindingObjectArgs{
// 			PolicyId: policy.ID(),
// 			Dimensions: monitor.PolicyBindingObjectDimensionArray{
// 				&monitor.PolicyBindingObjectDimensionArgs{
// 					DimensionsJson: pulumi.String(fmt.Sprintf("%v%v%v", "{\"unInstanceId\":\"", instances.InstanceLists[0].InstanceId, "\"}")),
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
// Monitor Policy Binding Object can be imported, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Monitor/policyBindingObject:PolicyBindingObject binding policyId
// ```
type PolicyBindingObject struct {
	pulumi.CustomResourceState

	// A list objects. Each element contains the following attributes:
	Dimensions PolicyBindingObjectDimensionArrayOutput `pulumi:"dimensions"`
	// Alarm policy ID for binding objects.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
}

// NewPolicyBindingObject registers a new resource with the given unique name, arguments, and options.
func NewPolicyBindingObject(ctx *pulumi.Context,
	name string, args *PolicyBindingObjectArgs, opts ...pulumi.ResourceOption) (*PolicyBindingObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dimensions == nil {
		return nil, errors.New("invalid value for required argument 'Dimensions'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	var resource PolicyBindingObject
	err := ctx.RegisterResource("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyBindingObject gets an existing PolicyBindingObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyBindingObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyBindingObjectState, opts ...pulumi.ResourceOption) (*PolicyBindingObject, error) {
	var resource PolicyBindingObject
	err := ctx.ReadResource("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyBindingObject resources.
type policyBindingObjectState struct {
	// A list objects. Each element contains the following attributes:
	Dimensions []PolicyBindingObjectDimension `pulumi:"dimensions"`
	// Alarm policy ID for binding objects.
	PolicyId *string `pulumi:"policyId"`
}

type PolicyBindingObjectState struct {
	// A list objects. Each element contains the following attributes:
	Dimensions PolicyBindingObjectDimensionArrayInput
	// Alarm policy ID for binding objects.
	PolicyId pulumi.StringPtrInput
}

func (PolicyBindingObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyBindingObjectState)(nil)).Elem()
}

type policyBindingObjectArgs struct {
	// A list objects. Each element contains the following attributes:
	Dimensions []PolicyBindingObjectDimension `pulumi:"dimensions"`
	// Alarm policy ID for binding objects.
	PolicyId string `pulumi:"policyId"`
}

// The set of arguments for constructing a PolicyBindingObject resource.
type PolicyBindingObjectArgs struct {
	// A list objects. Each element contains the following attributes:
	Dimensions PolicyBindingObjectDimensionArrayInput
	// Alarm policy ID for binding objects.
	PolicyId pulumi.StringInput
}

func (PolicyBindingObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyBindingObjectArgs)(nil)).Elem()
}

type PolicyBindingObjectInput interface {
	pulumi.Input

	ToPolicyBindingObjectOutput() PolicyBindingObjectOutput
	ToPolicyBindingObjectOutputWithContext(ctx context.Context) PolicyBindingObjectOutput
}

func (*PolicyBindingObject) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyBindingObject)(nil)).Elem()
}

func (i *PolicyBindingObject) ToPolicyBindingObjectOutput() PolicyBindingObjectOutput {
	return i.ToPolicyBindingObjectOutputWithContext(context.Background())
}

func (i *PolicyBindingObject) ToPolicyBindingObjectOutputWithContext(ctx context.Context) PolicyBindingObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyBindingObjectOutput)
}

// PolicyBindingObjectArrayInput is an input type that accepts PolicyBindingObjectArray and PolicyBindingObjectArrayOutput values.
// You can construct a concrete instance of `PolicyBindingObjectArrayInput` via:
//
//          PolicyBindingObjectArray{ PolicyBindingObjectArgs{...} }
type PolicyBindingObjectArrayInput interface {
	pulumi.Input

	ToPolicyBindingObjectArrayOutput() PolicyBindingObjectArrayOutput
	ToPolicyBindingObjectArrayOutputWithContext(context.Context) PolicyBindingObjectArrayOutput
}

type PolicyBindingObjectArray []PolicyBindingObjectInput

func (PolicyBindingObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyBindingObject)(nil)).Elem()
}

func (i PolicyBindingObjectArray) ToPolicyBindingObjectArrayOutput() PolicyBindingObjectArrayOutput {
	return i.ToPolicyBindingObjectArrayOutputWithContext(context.Background())
}

func (i PolicyBindingObjectArray) ToPolicyBindingObjectArrayOutputWithContext(ctx context.Context) PolicyBindingObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyBindingObjectArrayOutput)
}

// PolicyBindingObjectMapInput is an input type that accepts PolicyBindingObjectMap and PolicyBindingObjectMapOutput values.
// You can construct a concrete instance of `PolicyBindingObjectMapInput` via:
//
//          PolicyBindingObjectMap{ "key": PolicyBindingObjectArgs{...} }
type PolicyBindingObjectMapInput interface {
	pulumi.Input

	ToPolicyBindingObjectMapOutput() PolicyBindingObjectMapOutput
	ToPolicyBindingObjectMapOutputWithContext(context.Context) PolicyBindingObjectMapOutput
}

type PolicyBindingObjectMap map[string]PolicyBindingObjectInput

func (PolicyBindingObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyBindingObject)(nil)).Elem()
}

func (i PolicyBindingObjectMap) ToPolicyBindingObjectMapOutput() PolicyBindingObjectMapOutput {
	return i.ToPolicyBindingObjectMapOutputWithContext(context.Background())
}

func (i PolicyBindingObjectMap) ToPolicyBindingObjectMapOutputWithContext(ctx context.Context) PolicyBindingObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyBindingObjectMapOutput)
}

type PolicyBindingObjectOutput struct{ *pulumi.OutputState }

func (PolicyBindingObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyBindingObject)(nil)).Elem()
}

func (o PolicyBindingObjectOutput) ToPolicyBindingObjectOutput() PolicyBindingObjectOutput {
	return o
}

func (o PolicyBindingObjectOutput) ToPolicyBindingObjectOutputWithContext(ctx context.Context) PolicyBindingObjectOutput {
	return o
}

// A list objects. Each element contains the following attributes:
func (o PolicyBindingObjectOutput) Dimensions() PolicyBindingObjectDimensionArrayOutput {
	return o.ApplyT(func(v *PolicyBindingObject) PolicyBindingObjectDimensionArrayOutput { return v.Dimensions }).(PolicyBindingObjectDimensionArrayOutput)
}

// Alarm policy ID for binding objects.
func (o PolicyBindingObjectOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicyBindingObject) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

type PolicyBindingObjectArrayOutput struct{ *pulumi.OutputState }

func (PolicyBindingObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicyBindingObject)(nil)).Elem()
}

func (o PolicyBindingObjectArrayOutput) ToPolicyBindingObjectArrayOutput() PolicyBindingObjectArrayOutput {
	return o
}

func (o PolicyBindingObjectArrayOutput) ToPolicyBindingObjectArrayOutputWithContext(ctx context.Context) PolicyBindingObjectArrayOutput {
	return o
}

func (o PolicyBindingObjectArrayOutput) Index(i pulumi.IntInput) PolicyBindingObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicyBindingObject {
		return vs[0].([]*PolicyBindingObject)[vs[1].(int)]
	}).(PolicyBindingObjectOutput)
}

type PolicyBindingObjectMapOutput struct{ *pulumi.OutputState }

func (PolicyBindingObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicyBindingObject)(nil)).Elem()
}

func (o PolicyBindingObjectMapOutput) ToPolicyBindingObjectMapOutput() PolicyBindingObjectMapOutput {
	return o
}

func (o PolicyBindingObjectMapOutput) ToPolicyBindingObjectMapOutputWithContext(ctx context.Context) PolicyBindingObjectMapOutput {
	return o
}

func (o PolicyBindingObjectMapOutput) MapIndex(k pulumi.StringInput) PolicyBindingObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicyBindingObject {
		return vs[0].(map[string]*PolicyBindingObject)[vs[1].(string)]
	}).(PolicyBindingObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyBindingObjectInput)(nil)).Elem(), &PolicyBindingObject{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyBindingObjectArrayInput)(nil)).Elem(), PolicyBindingObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyBindingObjectMapInput)(nil)).Elem(), PolicyBindingObjectMap{})
	pulumi.RegisterOutputType(PolicyBindingObjectOutput{})
	pulumi.RegisterOutputType(PolicyBindingObjectArrayOutput{})
	pulumi.RegisterOutputType(PolicyBindingObjectMapOutput{})
}
