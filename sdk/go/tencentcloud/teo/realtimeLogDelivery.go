// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package teo

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a teo teoRealtimeLogDelivery
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Teo"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Teo.NewRealtimeLogDelivery(ctx, "teoRealtimeLogDelivery", &Teo.RealtimeLogDeliveryArgs{
//				Area:           pulumi.String("overseas"),
//				DeliveryStatus: pulumi.String("disabled"),
//				EntityLists: pulumi.StringArray{
//					pulumi.String("sid-2yvhjw98uaco"),
//				},
//				Fields: pulumi.StringArray{
//					pulumi.String("ServiceID"),
//					pulumi.String("ConnectTimeStamp"),
//					pulumi.String("DisconnetTimeStamp"),
//					pulumi.String("DisconnetReason"),
//					pulumi.String("ClientRealIP"),
//					pulumi.String("ClientRegion"),
//					pulumi.String("EdgeIP"),
//					pulumi.String("ForwardProtocol"),
//					pulumi.String("ForwardPort"),
//					pulumi.String("SentBytes"),
//					pulumi.String("ReceivedBytes"),
//					pulumi.String("LogTimeStamp"),
//				},
//				LogFormat: &teo.RealtimeLogDeliveryLogFormatArgs{
//					FieldDelimiter:  pulumi.String(","),
//					FormatType:      pulumi.String("json"),
//					RecordDelimiter: pulumi.String("\n\n"),
//					RecordPrefix:    pulumi.String("{"),
//					RecordSuffix:    pulumi.String("}"),
//				},
//				LogType: pulumi.String("application"),
//				S3: &teo.RealtimeLogDeliveryS3Args{
//					AccessId:     pulumi.String("xxxxxxxxxx"),
//					AccessKey:    pulumi.String("xxxxxxxxxx"),
//					Bucket:       pulumi.String("test-1253833068"),
//					CompressType: pulumi.String("gzip"),
//					Endpoint:     pulumi.String("https://test-1253833068.cos.ap-nanjing.myqcloud.com"),
//					Region:       pulumi.String("ap-nanjing"),
//				},
//				Sample:   pulumi.Int(0),
//				TaskName: pulumi.String("test"),
//				TaskType: pulumi.String("s3"),
//				ZoneId:   pulumi.String("zone-2qtuhspy7cr6"),
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
// teo teo_realtime_log_delivery can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Teo/realtimeLogDelivery:RealtimeLogDelivery teo_realtime_log_delivery zoneId#taskId
// ```
type RealtimeLogDelivery struct {
	pulumi.CustomResourceState

	// Data delivery area, possible values are: `mainland`: within mainland China; `overseas`: worldwide (excluding mainland China).
	Area pulumi.StringOutput `pulumi:"area"`
	// CLS configuration information. This parameter is required when TaskType is cls.
	Cls RealtimeLogDeliveryClsOutput `pulumi:"cls"`
	// Customize the configuration information of the HTTP service. This parameter is required when TaskType is set to custom_endpoint.
	CustomEndpoint RealtimeLogDeliveryCustomEndpointOutput `pulumi:"customEndpoint"`
	// The list of custom fields delivered supports extracting specified field values from HTTP request headers, response headers, and cookies. Custom field names cannot be repeated and cannot exceed 200 fields.
	CustomFields RealtimeLogDeliveryCustomFieldArrayOutput `pulumi:"customFields"`
	// The filter condition for log delivery. If it is not filled, all logs will be delivered.
	DeliveryConditions RealtimeLogDeliveryDeliveryConditionArrayOutput `pulumi:"deliveryConditions"`
	// The status of the real-time log delivery task. The values are: `enabled`: enabled; `disabled`: disabled. Leave it blank to keep the original configuration. Not required when creating.
	DeliveryStatus pulumi.StringOutput `pulumi:"deliveryStatus"`
	// List of entities (seven-layer domain names or four-layer proxy instances) corresponding to real-time log delivery tasks. Example values are as follows: Seven-layer domain name: `domain.example.com`; four-layer proxy instance: sid-2s69eb5wcms7. For values, refer to: `https://cloud.tencent.com/document/api/1552/80690`, `https://cloud.tencent.com/document/api/1552/86336`.
	EntityLists pulumi.StringArrayOutput `pulumi:"entityLists"`
	// A list of preset fields for delivery.
	Fields pulumi.StringArrayOutput `pulumi:"fields"`
	// The output format of log delivery. If it is not filled, it means the default format. The default format logic is as follows: when TaskType is `customEndpoint`, the default format is an array of multiple JSON objects, each JSON object is a log; when TaskType is `s3`, the default format is JSON Lines; in particular, when TaskType is `cls`, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to pass LogFormat.
	LogFormat RealtimeLogDeliveryLogFormatOutput `pulumi:"logFormat"`
	// Data delivery type, the values are: `domain`: site acceleration log; `application`: four-layer proxy log; `web-rateLiming`: rate limit and CC attack protection log; `web-attack`: managed rule log; `web-rule`: custom rule log; `web-bot`: Bot management log.
	LogType pulumi.StringOutput `pulumi:"logType"`
	// Configuration information of AWS S3 compatible storage bucket. This parameter is required when TaskType is s3.
	S3 RealtimeLogDeliveryS3Output `pulumi:"s3"`
	// The sampling ratio is in thousandths, with a value range of 1-1000. For example, filling in 605 means the sampling ratio is 60.5%. Leaving it blank means the sampling ratio is 100%.
	Sample pulumi.IntOutput `pulumi:"sample"`
	// Real-time log delivery task ID.
	TaskId pulumi.StringOutput `pulumi:"taskId"`
	// The name of the real-time log delivery task. The format is a combination of numbers, English, -, and _. The maximum length is 200 characters.
	TaskName pulumi.StringOutput `pulumi:"taskName"`
	// The real-time log delivery task type. The possible values are: `cls`: push to Tencent Cloud CLS; `customEndpoint`: push to a custom HTTP(S) address; `s3`: push to an AWS S3 compatible storage bucket address.
	TaskType pulumi.StringOutput `pulumi:"taskType"`
	// ID of the site.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewRealtimeLogDelivery registers a new resource with the given unique name, arguments, and options.
func NewRealtimeLogDelivery(ctx *pulumi.Context,
	name string, args *RealtimeLogDeliveryArgs, opts ...pulumi.ResourceOption) (*RealtimeLogDelivery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Area == nil {
		return nil, errors.New("invalid value for required argument 'Area'")
	}
	if args.EntityLists == nil {
		return nil, errors.New("invalid value for required argument 'EntityLists'")
	}
	if args.Fields == nil {
		return nil, errors.New("invalid value for required argument 'Fields'")
	}
	if args.LogType == nil {
		return nil, errors.New("invalid value for required argument 'LogType'")
	}
	if args.Sample == nil {
		return nil, errors.New("invalid value for required argument 'Sample'")
	}
	if args.TaskName == nil {
		return nil, errors.New("invalid value for required argument 'TaskName'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RealtimeLogDelivery
	err := ctx.RegisterResource("tencentcloud:Teo/realtimeLogDelivery:RealtimeLogDelivery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealtimeLogDelivery gets an existing RealtimeLogDelivery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealtimeLogDelivery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealtimeLogDeliveryState, opts ...pulumi.ResourceOption) (*RealtimeLogDelivery, error) {
	var resource RealtimeLogDelivery
	err := ctx.ReadResource("tencentcloud:Teo/realtimeLogDelivery:RealtimeLogDelivery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealtimeLogDelivery resources.
type realtimeLogDeliveryState struct {
	// Data delivery area, possible values are: `mainland`: within mainland China; `overseas`: worldwide (excluding mainland China).
	Area *string `pulumi:"area"`
	// CLS configuration information. This parameter is required when TaskType is cls.
	Cls *RealtimeLogDeliveryCls `pulumi:"cls"`
	// Customize the configuration information of the HTTP service. This parameter is required when TaskType is set to custom_endpoint.
	CustomEndpoint *RealtimeLogDeliveryCustomEndpoint `pulumi:"customEndpoint"`
	// The list of custom fields delivered supports extracting specified field values from HTTP request headers, response headers, and cookies. Custom field names cannot be repeated and cannot exceed 200 fields.
	CustomFields []RealtimeLogDeliveryCustomField `pulumi:"customFields"`
	// The filter condition for log delivery. If it is not filled, all logs will be delivered.
	DeliveryConditions []RealtimeLogDeliveryDeliveryCondition `pulumi:"deliveryConditions"`
	// The status of the real-time log delivery task. The values are: `enabled`: enabled; `disabled`: disabled. Leave it blank to keep the original configuration. Not required when creating.
	DeliveryStatus *string `pulumi:"deliveryStatus"`
	// List of entities (seven-layer domain names or four-layer proxy instances) corresponding to real-time log delivery tasks. Example values are as follows: Seven-layer domain name: `domain.example.com`; four-layer proxy instance: sid-2s69eb5wcms7. For values, refer to: `https://cloud.tencent.com/document/api/1552/80690`, `https://cloud.tencent.com/document/api/1552/86336`.
	EntityLists []string `pulumi:"entityLists"`
	// A list of preset fields for delivery.
	Fields []string `pulumi:"fields"`
	// The output format of log delivery. If it is not filled, it means the default format. The default format logic is as follows: when TaskType is `customEndpoint`, the default format is an array of multiple JSON objects, each JSON object is a log; when TaskType is `s3`, the default format is JSON Lines; in particular, when TaskType is `cls`, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to pass LogFormat.
	LogFormat *RealtimeLogDeliveryLogFormat `pulumi:"logFormat"`
	// Data delivery type, the values are: `domain`: site acceleration log; `application`: four-layer proxy log; `web-rateLiming`: rate limit and CC attack protection log; `web-attack`: managed rule log; `web-rule`: custom rule log; `web-bot`: Bot management log.
	LogType *string `pulumi:"logType"`
	// Configuration information of AWS S3 compatible storage bucket. This parameter is required when TaskType is s3.
	S3 *RealtimeLogDeliveryS3 `pulumi:"s3"`
	// The sampling ratio is in thousandths, with a value range of 1-1000. For example, filling in 605 means the sampling ratio is 60.5%. Leaving it blank means the sampling ratio is 100%.
	Sample *int `pulumi:"sample"`
	// Real-time log delivery task ID.
	TaskId *string `pulumi:"taskId"`
	// The name of the real-time log delivery task. The format is a combination of numbers, English, -, and _. The maximum length is 200 characters.
	TaskName *string `pulumi:"taskName"`
	// The real-time log delivery task type. The possible values are: `cls`: push to Tencent Cloud CLS; `customEndpoint`: push to a custom HTTP(S) address; `s3`: push to an AWS S3 compatible storage bucket address.
	TaskType *string `pulumi:"taskType"`
	// ID of the site.
	ZoneId *string `pulumi:"zoneId"`
}

type RealtimeLogDeliveryState struct {
	// Data delivery area, possible values are: `mainland`: within mainland China; `overseas`: worldwide (excluding mainland China).
	Area pulumi.StringPtrInput
	// CLS configuration information. This parameter is required when TaskType is cls.
	Cls RealtimeLogDeliveryClsPtrInput
	// Customize the configuration information of the HTTP service. This parameter is required when TaskType is set to custom_endpoint.
	CustomEndpoint RealtimeLogDeliveryCustomEndpointPtrInput
	// The list of custom fields delivered supports extracting specified field values from HTTP request headers, response headers, and cookies. Custom field names cannot be repeated and cannot exceed 200 fields.
	CustomFields RealtimeLogDeliveryCustomFieldArrayInput
	// The filter condition for log delivery. If it is not filled, all logs will be delivered.
	DeliveryConditions RealtimeLogDeliveryDeliveryConditionArrayInput
	// The status of the real-time log delivery task. The values are: `enabled`: enabled; `disabled`: disabled. Leave it blank to keep the original configuration. Not required when creating.
	DeliveryStatus pulumi.StringPtrInput
	// List of entities (seven-layer domain names or four-layer proxy instances) corresponding to real-time log delivery tasks. Example values are as follows: Seven-layer domain name: `domain.example.com`; four-layer proxy instance: sid-2s69eb5wcms7. For values, refer to: `https://cloud.tencent.com/document/api/1552/80690`, `https://cloud.tencent.com/document/api/1552/86336`.
	EntityLists pulumi.StringArrayInput
	// A list of preset fields for delivery.
	Fields pulumi.StringArrayInput
	// The output format of log delivery. If it is not filled, it means the default format. The default format logic is as follows: when TaskType is `customEndpoint`, the default format is an array of multiple JSON objects, each JSON object is a log; when TaskType is `s3`, the default format is JSON Lines; in particular, when TaskType is `cls`, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to pass LogFormat.
	LogFormat RealtimeLogDeliveryLogFormatPtrInput
	// Data delivery type, the values are: `domain`: site acceleration log; `application`: four-layer proxy log; `web-rateLiming`: rate limit and CC attack protection log; `web-attack`: managed rule log; `web-rule`: custom rule log; `web-bot`: Bot management log.
	LogType pulumi.StringPtrInput
	// Configuration information of AWS S3 compatible storage bucket. This parameter is required when TaskType is s3.
	S3 RealtimeLogDeliveryS3PtrInput
	// The sampling ratio is in thousandths, with a value range of 1-1000. For example, filling in 605 means the sampling ratio is 60.5%. Leaving it blank means the sampling ratio is 100%.
	Sample pulumi.IntPtrInput
	// Real-time log delivery task ID.
	TaskId pulumi.StringPtrInput
	// The name of the real-time log delivery task. The format is a combination of numbers, English, -, and _. The maximum length is 200 characters.
	TaskName pulumi.StringPtrInput
	// The real-time log delivery task type. The possible values are: `cls`: push to Tencent Cloud CLS; `customEndpoint`: push to a custom HTTP(S) address; `s3`: push to an AWS S3 compatible storage bucket address.
	TaskType pulumi.StringPtrInput
	// ID of the site.
	ZoneId pulumi.StringPtrInput
}

func (RealtimeLogDeliveryState) ElementType() reflect.Type {
	return reflect.TypeOf((*realtimeLogDeliveryState)(nil)).Elem()
}

type realtimeLogDeliveryArgs struct {
	// Data delivery area, possible values are: `mainland`: within mainland China; `overseas`: worldwide (excluding mainland China).
	Area string `pulumi:"area"`
	// CLS configuration information. This parameter is required when TaskType is cls.
	Cls *RealtimeLogDeliveryCls `pulumi:"cls"`
	// Customize the configuration information of the HTTP service. This parameter is required when TaskType is set to custom_endpoint.
	CustomEndpoint *RealtimeLogDeliveryCustomEndpoint `pulumi:"customEndpoint"`
	// The list of custom fields delivered supports extracting specified field values from HTTP request headers, response headers, and cookies. Custom field names cannot be repeated and cannot exceed 200 fields.
	CustomFields []RealtimeLogDeliveryCustomField `pulumi:"customFields"`
	// The filter condition for log delivery. If it is not filled, all logs will be delivered.
	DeliveryConditions []RealtimeLogDeliveryDeliveryCondition `pulumi:"deliveryConditions"`
	// The status of the real-time log delivery task. The values are: `enabled`: enabled; `disabled`: disabled. Leave it blank to keep the original configuration. Not required when creating.
	DeliveryStatus *string `pulumi:"deliveryStatus"`
	// List of entities (seven-layer domain names or four-layer proxy instances) corresponding to real-time log delivery tasks. Example values are as follows: Seven-layer domain name: `domain.example.com`; four-layer proxy instance: sid-2s69eb5wcms7. For values, refer to: `https://cloud.tencent.com/document/api/1552/80690`, `https://cloud.tencent.com/document/api/1552/86336`.
	EntityLists []string `pulumi:"entityLists"`
	// A list of preset fields for delivery.
	Fields []string `pulumi:"fields"`
	// The output format of log delivery. If it is not filled, it means the default format. The default format logic is as follows: when TaskType is `customEndpoint`, the default format is an array of multiple JSON objects, each JSON object is a log; when TaskType is `s3`, the default format is JSON Lines; in particular, when TaskType is `cls`, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to pass LogFormat.
	LogFormat *RealtimeLogDeliveryLogFormat `pulumi:"logFormat"`
	// Data delivery type, the values are: `domain`: site acceleration log; `application`: four-layer proxy log; `web-rateLiming`: rate limit and CC attack protection log; `web-attack`: managed rule log; `web-rule`: custom rule log; `web-bot`: Bot management log.
	LogType string `pulumi:"logType"`
	// Configuration information of AWS S3 compatible storage bucket. This parameter is required when TaskType is s3.
	S3 *RealtimeLogDeliveryS3 `pulumi:"s3"`
	// The sampling ratio is in thousandths, with a value range of 1-1000. For example, filling in 605 means the sampling ratio is 60.5%. Leaving it blank means the sampling ratio is 100%.
	Sample int `pulumi:"sample"`
	// The name of the real-time log delivery task. The format is a combination of numbers, English, -, and _. The maximum length is 200 characters.
	TaskName string `pulumi:"taskName"`
	// The real-time log delivery task type. The possible values are: `cls`: push to Tencent Cloud CLS; `customEndpoint`: push to a custom HTTP(S) address; `s3`: push to an AWS S3 compatible storage bucket address.
	TaskType string `pulumi:"taskType"`
	// ID of the site.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a RealtimeLogDelivery resource.
type RealtimeLogDeliveryArgs struct {
	// Data delivery area, possible values are: `mainland`: within mainland China; `overseas`: worldwide (excluding mainland China).
	Area pulumi.StringInput
	// CLS configuration information. This parameter is required when TaskType is cls.
	Cls RealtimeLogDeliveryClsPtrInput
	// Customize the configuration information of the HTTP service. This parameter is required when TaskType is set to custom_endpoint.
	CustomEndpoint RealtimeLogDeliveryCustomEndpointPtrInput
	// The list of custom fields delivered supports extracting specified field values from HTTP request headers, response headers, and cookies. Custom field names cannot be repeated and cannot exceed 200 fields.
	CustomFields RealtimeLogDeliveryCustomFieldArrayInput
	// The filter condition for log delivery. If it is not filled, all logs will be delivered.
	DeliveryConditions RealtimeLogDeliveryDeliveryConditionArrayInput
	// The status of the real-time log delivery task. The values are: `enabled`: enabled; `disabled`: disabled. Leave it blank to keep the original configuration. Not required when creating.
	DeliveryStatus pulumi.StringPtrInput
	// List of entities (seven-layer domain names or four-layer proxy instances) corresponding to real-time log delivery tasks. Example values are as follows: Seven-layer domain name: `domain.example.com`; four-layer proxy instance: sid-2s69eb5wcms7. For values, refer to: `https://cloud.tencent.com/document/api/1552/80690`, `https://cloud.tencent.com/document/api/1552/86336`.
	EntityLists pulumi.StringArrayInput
	// A list of preset fields for delivery.
	Fields pulumi.StringArrayInput
	// The output format of log delivery. If it is not filled, it means the default format. The default format logic is as follows: when TaskType is `customEndpoint`, the default format is an array of multiple JSON objects, each JSON object is a log; when TaskType is `s3`, the default format is JSON Lines; in particular, when TaskType is `cls`, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to pass LogFormat.
	LogFormat RealtimeLogDeliveryLogFormatPtrInput
	// Data delivery type, the values are: `domain`: site acceleration log; `application`: four-layer proxy log; `web-rateLiming`: rate limit and CC attack protection log; `web-attack`: managed rule log; `web-rule`: custom rule log; `web-bot`: Bot management log.
	LogType pulumi.StringInput
	// Configuration information of AWS S3 compatible storage bucket. This parameter is required when TaskType is s3.
	S3 RealtimeLogDeliveryS3PtrInput
	// The sampling ratio is in thousandths, with a value range of 1-1000. For example, filling in 605 means the sampling ratio is 60.5%. Leaving it blank means the sampling ratio is 100%.
	Sample pulumi.IntInput
	// The name of the real-time log delivery task. The format is a combination of numbers, English, -, and _. The maximum length is 200 characters.
	TaskName pulumi.StringInput
	// The real-time log delivery task type. The possible values are: `cls`: push to Tencent Cloud CLS; `customEndpoint`: push to a custom HTTP(S) address; `s3`: push to an AWS S3 compatible storage bucket address.
	TaskType pulumi.StringInput
	// ID of the site.
	ZoneId pulumi.StringInput
}

func (RealtimeLogDeliveryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realtimeLogDeliveryArgs)(nil)).Elem()
}

type RealtimeLogDeliveryInput interface {
	pulumi.Input

	ToRealtimeLogDeliveryOutput() RealtimeLogDeliveryOutput
	ToRealtimeLogDeliveryOutputWithContext(ctx context.Context) RealtimeLogDeliveryOutput
}

func (*RealtimeLogDelivery) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeLogDelivery)(nil)).Elem()
}

func (i *RealtimeLogDelivery) ToRealtimeLogDeliveryOutput() RealtimeLogDeliveryOutput {
	return i.ToRealtimeLogDeliveryOutputWithContext(context.Background())
}

func (i *RealtimeLogDelivery) ToRealtimeLogDeliveryOutputWithContext(ctx context.Context) RealtimeLogDeliveryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeLogDeliveryOutput)
}

// RealtimeLogDeliveryArrayInput is an input type that accepts RealtimeLogDeliveryArray and RealtimeLogDeliveryArrayOutput values.
// You can construct a concrete instance of `RealtimeLogDeliveryArrayInput` via:
//
//	RealtimeLogDeliveryArray{ RealtimeLogDeliveryArgs{...} }
type RealtimeLogDeliveryArrayInput interface {
	pulumi.Input

	ToRealtimeLogDeliveryArrayOutput() RealtimeLogDeliveryArrayOutput
	ToRealtimeLogDeliveryArrayOutputWithContext(context.Context) RealtimeLogDeliveryArrayOutput
}

type RealtimeLogDeliveryArray []RealtimeLogDeliveryInput

func (RealtimeLogDeliveryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealtimeLogDelivery)(nil)).Elem()
}

func (i RealtimeLogDeliveryArray) ToRealtimeLogDeliveryArrayOutput() RealtimeLogDeliveryArrayOutput {
	return i.ToRealtimeLogDeliveryArrayOutputWithContext(context.Background())
}

func (i RealtimeLogDeliveryArray) ToRealtimeLogDeliveryArrayOutputWithContext(ctx context.Context) RealtimeLogDeliveryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeLogDeliveryArrayOutput)
}

// RealtimeLogDeliveryMapInput is an input type that accepts RealtimeLogDeliveryMap and RealtimeLogDeliveryMapOutput values.
// You can construct a concrete instance of `RealtimeLogDeliveryMapInput` via:
//
//	RealtimeLogDeliveryMap{ "key": RealtimeLogDeliveryArgs{...} }
type RealtimeLogDeliveryMapInput interface {
	pulumi.Input

	ToRealtimeLogDeliveryMapOutput() RealtimeLogDeliveryMapOutput
	ToRealtimeLogDeliveryMapOutputWithContext(context.Context) RealtimeLogDeliveryMapOutput
}

type RealtimeLogDeliveryMap map[string]RealtimeLogDeliveryInput

func (RealtimeLogDeliveryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealtimeLogDelivery)(nil)).Elem()
}

func (i RealtimeLogDeliveryMap) ToRealtimeLogDeliveryMapOutput() RealtimeLogDeliveryMapOutput {
	return i.ToRealtimeLogDeliveryMapOutputWithContext(context.Background())
}

func (i RealtimeLogDeliveryMap) ToRealtimeLogDeliveryMapOutputWithContext(ctx context.Context) RealtimeLogDeliveryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealtimeLogDeliveryMapOutput)
}

type RealtimeLogDeliveryOutput struct{ *pulumi.OutputState }

func (RealtimeLogDeliveryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealtimeLogDelivery)(nil)).Elem()
}

func (o RealtimeLogDeliveryOutput) ToRealtimeLogDeliveryOutput() RealtimeLogDeliveryOutput {
	return o
}

func (o RealtimeLogDeliveryOutput) ToRealtimeLogDeliveryOutputWithContext(ctx context.Context) RealtimeLogDeliveryOutput {
	return o
}

// Data delivery area, possible values are: `mainland`: within mainland China; `overseas`: worldwide (excluding mainland China).
func (o RealtimeLogDeliveryOutput) Area() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringOutput { return v.Area }).(pulumi.StringOutput)
}

// CLS configuration information. This parameter is required when TaskType is cls.
func (o RealtimeLogDeliveryOutput) Cls() RealtimeLogDeliveryClsOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) RealtimeLogDeliveryClsOutput { return v.Cls }).(RealtimeLogDeliveryClsOutput)
}

// Customize the configuration information of the HTTP service. This parameter is required when TaskType is set to custom_endpoint.
func (o RealtimeLogDeliveryOutput) CustomEndpoint() RealtimeLogDeliveryCustomEndpointOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) RealtimeLogDeliveryCustomEndpointOutput { return v.CustomEndpoint }).(RealtimeLogDeliveryCustomEndpointOutput)
}

// The list of custom fields delivered supports extracting specified field values from HTTP request headers, response headers, and cookies. Custom field names cannot be repeated and cannot exceed 200 fields.
func (o RealtimeLogDeliveryOutput) CustomFields() RealtimeLogDeliveryCustomFieldArrayOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) RealtimeLogDeliveryCustomFieldArrayOutput { return v.CustomFields }).(RealtimeLogDeliveryCustomFieldArrayOutput)
}

// The filter condition for log delivery. If it is not filled, all logs will be delivered.
func (o RealtimeLogDeliveryOutput) DeliveryConditions() RealtimeLogDeliveryDeliveryConditionArrayOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) RealtimeLogDeliveryDeliveryConditionArrayOutput {
		return v.DeliveryConditions
	}).(RealtimeLogDeliveryDeliveryConditionArrayOutput)
}

// The status of the real-time log delivery task. The values are: `enabled`: enabled; `disabled`: disabled. Leave it blank to keep the original configuration. Not required when creating.
func (o RealtimeLogDeliveryOutput) DeliveryStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringOutput { return v.DeliveryStatus }).(pulumi.StringOutput)
}

// List of entities (seven-layer domain names or four-layer proxy instances) corresponding to real-time log delivery tasks. Example values are as follows: Seven-layer domain name: `domain.example.com`; four-layer proxy instance: sid-2s69eb5wcms7. For values, refer to: `https://cloud.tencent.com/document/api/1552/80690`, `https://cloud.tencent.com/document/api/1552/86336`.
func (o RealtimeLogDeliveryOutput) EntityLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringArrayOutput { return v.EntityLists }).(pulumi.StringArrayOutput)
}

// A list of preset fields for delivery.
func (o RealtimeLogDeliveryOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringArrayOutput { return v.Fields }).(pulumi.StringArrayOutput)
}

// The output format of log delivery. If it is not filled, it means the default format. The default format logic is as follows: when TaskType is `customEndpoint`, the default format is an array of multiple JSON objects, each JSON object is a log; when TaskType is `s3`, the default format is JSON Lines; in particular, when TaskType is `cls`, the value of LogFormat.FormatType can only be json, and other parameters in LogFormat will be ignored. It is recommended not to pass LogFormat.
func (o RealtimeLogDeliveryOutput) LogFormat() RealtimeLogDeliveryLogFormatOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) RealtimeLogDeliveryLogFormatOutput { return v.LogFormat }).(RealtimeLogDeliveryLogFormatOutput)
}

// Data delivery type, the values are: `domain`: site acceleration log; `application`: four-layer proxy log; `web-rateLiming`: rate limit and CC attack protection log; `web-attack`: managed rule log; `web-rule`: custom rule log; `web-bot`: Bot management log.
func (o RealtimeLogDeliveryOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringOutput { return v.LogType }).(pulumi.StringOutput)
}

// Configuration information of AWS S3 compatible storage bucket. This parameter is required when TaskType is s3.
func (o RealtimeLogDeliveryOutput) S3() RealtimeLogDeliveryS3Output {
	return o.ApplyT(func(v *RealtimeLogDelivery) RealtimeLogDeliveryS3Output { return v.S3 }).(RealtimeLogDeliveryS3Output)
}

// The sampling ratio is in thousandths, with a value range of 1-1000. For example, filling in 605 means the sampling ratio is 60.5%. Leaving it blank means the sampling ratio is 100%.
func (o RealtimeLogDeliveryOutput) Sample() pulumi.IntOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.IntOutput { return v.Sample }).(pulumi.IntOutput)
}

// Real-time log delivery task ID.
func (o RealtimeLogDeliveryOutput) TaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringOutput { return v.TaskId }).(pulumi.StringOutput)
}

// The name of the real-time log delivery task. The format is a combination of numbers, English, -, and _. The maximum length is 200 characters.
func (o RealtimeLogDeliveryOutput) TaskName() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringOutput { return v.TaskName }).(pulumi.StringOutput)
}

// The real-time log delivery task type. The possible values are: `cls`: push to Tencent Cloud CLS; `customEndpoint`: push to a custom HTTP(S) address; `s3`: push to an AWS S3 compatible storage bucket address.
func (o RealtimeLogDeliveryOutput) TaskType() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringOutput { return v.TaskType }).(pulumi.StringOutput)
}

// ID of the site.
func (o RealtimeLogDeliveryOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealtimeLogDelivery) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type RealtimeLogDeliveryArrayOutput struct{ *pulumi.OutputState }

func (RealtimeLogDeliveryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealtimeLogDelivery)(nil)).Elem()
}

func (o RealtimeLogDeliveryArrayOutput) ToRealtimeLogDeliveryArrayOutput() RealtimeLogDeliveryArrayOutput {
	return o
}

func (o RealtimeLogDeliveryArrayOutput) ToRealtimeLogDeliveryArrayOutputWithContext(ctx context.Context) RealtimeLogDeliveryArrayOutput {
	return o
}

func (o RealtimeLogDeliveryArrayOutput) Index(i pulumi.IntInput) RealtimeLogDeliveryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealtimeLogDelivery {
		return vs[0].([]*RealtimeLogDelivery)[vs[1].(int)]
	}).(RealtimeLogDeliveryOutput)
}

type RealtimeLogDeliveryMapOutput struct{ *pulumi.OutputState }

func (RealtimeLogDeliveryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealtimeLogDelivery)(nil)).Elem()
}

func (o RealtimeLogDeliveryMapOutput) ToRealtimeLogDeliveryMapOutput() RealtimeLogDeliveryMapOutput {
	return o
}

func (o RealtimeLogDeliveryMapOutput) ToRealtimeLogDeliveryMapOutputWithContext(ctx context.Context) RealtimeLogDeliveryMapOutput {
	return o
}

func (o RealtimeLogDeliveryMapOutput) MapIndex(k pulumi.StringInput) RealtimeLogDeliveryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealtimeLogDelivery {
		return vs[0].(map[string]*RealtimeLogDelivery)[vs[1].(string)]
	}).(RealtimeLogDeliveryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealtimeLogDeliveryInput)(nil)).Elem(), &RealtimeLogDelivery{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealtimeLogDeliveryArrayInput)(nil)).Elem(), RealtimeLogDeliveryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealtimeLogDeliveryMapInput)(nil)).Elem(), RealtimeLogDeliveryMap{})
	pulumi.RegisterOutputType(RealtimeLogDeliveryOutput{})
	pulumi.RegisterOutputType(RealtimeLogDeliveryArrayOutput{})
	pulumi.RegisterOutputType(RealtimeLogDeliveryMapOutput{})
}
