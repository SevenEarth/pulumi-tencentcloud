// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of monitor alarmPolicy
func LookupAlarmPolicy(ctx *pulumi.Context, args *LookupAlarmPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAlarmPolicyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupAlarmPolicyResult
	err := ctx.Invoke("tencentcloud:Monitor/getAlarmPolicy:getAlarmPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAlarmPolicy.
type LookupAlarmPolicyArgs struct {
	// The alarm object list, which is a JSON string. The outer array corresponds to multiple instances, and the inner array is the dimension of an object.For example, 'CVM - Basic Monitor' can be written as: [ {Dimensions: {unInstanceId: ins-qr8d555g}}, {Dimensions: {unInstanceId: ins-qr8d555h}} ]You can also refer to the 'Example 2' below.For more information on the parameter samples of different Tencent Cloud services, see [Product Policy Type and Dimension Information](https://www.tencentcloud.com/document/product/248/39565?has_map=1).Note: If 1 is passed in for NeedCorrespondence, the relationship between a policy and an instance needs to be returned. You can pass in up to 20 alarm object dimensions to avoid request timeout.
	Dimensions *string `pulumi:"dimensions"`
	// Filter by alarm status. Valid values: [1]: enabled; [0]: disabled; [0, 1]: all.
	Enables []int `pulumi:"enables"`
	// Sort by field. For example, to sort by the last modification time, use Field: UpdateTime.
	Field *string `pulumi:"field"`
	// Instance group ID.
	InstanceGroupId *int `pulumi:"instanceGroupId"`
	// Value fixed at monitor.
	Module string `pulumi:"module"`
	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default.
	MonitorTypes []string `pulumi:"monitorTypes"`
	// Filter by namespace. For the values of different policy types, please see:[Poicy Type List](https://www.tencentcloud.com/document/product/248/39565?has_map=1).
	Namespaces []string `pulumi:"namespaces"`
	// Whether the relationship between a policy and the input parameter filter dimension is required. 1: Yes. 0: No. Default value: 0.
	NeedCorrespondence *int `pulumi:"needCorrespondence"`
	// Whether the returned result needs to filter policies associated with all objects. Valid values: 1 (Yes), 0 (No).
	NotBindAll *int `pulumi:"notBindAll"`
	// If 1 is passed in, alarm policies with no notification rules configured are queried. If it is left empty or other values are passed in, all alarm policies are queried.
	NotBindingNoticeRule *int `pulumi:"notBindingNoticeRule"`
	// Whether the returned result needs to filter policies associated with instance groups. Valid values: 1 (Yes), 0 (No).
	NotInstanceGroup *int `pulumi:"notInstanceGroup"`
	// List of the notification template IDs, which can be obtained by querying the notification template list.It can be queried with the API [DescribeAlarmNotices](https://www.tencentcloud.com/document/product/248/39300).
	NoticeIds []string `pulumi:"noticeIds"`
	// Filter by quick alarm policy. If this parameter is left empty, all policies are displayed. ONECLICK: Display quick alarm policies; NOT_ONECLICK: Display non-quick alarm policies.
	OneClickPolicyTypes []string `pulumi:"oneClickPolicyTypes"`
	// Sort order. Valid values: ASC (ascending), DESC (descending).
	Order *string `pulumi:"order"`
	// Fuzzy search by policy name.
	PolicyName *string `pulumi:"policyName"`
	// Filter by default policy. Valid values: DEFAULT (display default policy), NOT_DEFAULT (display non-default policies). If this parameter is left empty, all policies will be displayed.
	PolicyTypes []string `pulumi:"policyTypes"`
	// ID array of the policy project, which can be viewed on the following page: [Project Management](https://console.tencentcloud.com/project).
	ProjectIds []int `pulumi:"projectIds"`
	// ID of the TencentCloud Managed Service for Prometheus instance, which is used for customizing a metric policy.
	PromInsId *string `pulumi:"promInsId"`
	// Search by recipient group. You can get the user group list with the API [ListGroups](https://www.tencentcloud.com/document/product/598/34589?from_cn_redirect=1) in 'Cloud Access Management' or query the user group list where a sub-user is in with the API [ListGroupsForUser](https://www.tencentcloud.com/document/product/598/34588?from_cn_redirect=1). The GroupId field in the returned result should be entered here.
	ReceiverGroups []int `pulumi:"receiverGroups"`
	// Search by schedule.
	ReceiverOnCallFormIds []string `pulumi:"receiverOnCallFormIds"`
	// Search by recipient. You can get the user list with the API [ListUsers](https://www.tencentcloud.com/document/product/598/34587?from_cn_redirect=1) in 'Cloud Access Management' or query the sub-user information with the API [GetUser](https://www.tencentcloud.com/document/product/598/34590?from_cn_redirect=1). The Uid field in the returned result should be entered here.
	ReceiverUids []int `pulumi:"receiverUids"`
	// Used to save results.
	ResultOutputFile *string `pulumi:"resultOutputFile"`
	// Filter by trigger condition. Valid values: STATIC (display policies with static threshold), DYNAMIC (display policies with dynamic threshold). If this parameter is left empty, all policies will be displayed.
	RuleTypes []string `pulumi:"ruleTypes"`
	// Filter alarm policy by triggered task (such as auto scaling task). Up to 10 tasks can be specified.
	TriggerTasks []GetAlarmPolicyTriggerTask `pulumi:"triggerTasks"`
}

// A collection of values returned by getAlarmPolicy.
type LookupAlarmPolicyResult struct {
	// JSON string generated by serializing the AlarmPolicyDimension two-dimensional array. The one-dimensional arrays are in OR relationship, and the elements in a one-dimensional array are in AND relationshipNote: this field may return null, indicating that no valid values can be obtained.
	Dimensions *string `pulumi:"dimensions"`
	// Status. Valid values: 0 (disabled), 1 (enabled). Default value: 1 (enabled). This parameter can be left empty.
	Enables []int   `pulumi:"enables"`
	Field   *string `pulumi:"field"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Instance group IDNote: this field may return null, indicating that no valid values can be obtained.
	InstanceGroupId      *int     `pulumi:"instanceGroupId"`
	Module               string   `pulumi:"module"`
	MonitorTypes         []string `pulumi:"monitorTypes"`
	Namespaces           []string `pulumi:"namespaces"`
	NeedCorrespondence   *int     `pulumi:"needCorrespondence"`
	NotBindAll           *int     `pulumi:"notBindAll"`
	NotBindingNoticeRule *int     `pulumi:"notBindingNoticeRule"`
	NotInstanceGroup     *int     `pulumi:"notInstanceGroup"`
	// Notification rule ID listNote: this field may return null, indicating that no valid values can be obtained.
	NoticeIds           []string `pulumi:"noticeIds"`
	OneClickPolicyTypes []string `pulumi:"oneClickPolicyTypes"`
	Order               *string  `pulumi:"order"`
	// Policy array.
	Policies []GetAlarmPolicyPolicy `pulumi:"policies"`
	// Alarm policy nameNote: this field may return null, indicating that no valid values can be obtained.
	PolicyName            *string  `pulumi:"policyName"`
	PolicyTypes           []string `pulumi:"policyTypes"`
	ProjectIds            []int    `pulumi:"projectIds"`
	PromInsId             *string  `pulumi:"promInsId"`
	ReceiverGroups        []int    `pulumi:"receiverGroups"`
	ReceiverOnCallFormIds []string `pulumi:"receiverOnCallFormIds"`
	ReceiverUids          []int    `pulumi:"receiverUids"`
	ResultOutputFile      *string  `pulumi:"resultOutputFile"`
	RuleTypes             []string `pulumi:"ruleTypes"`
	// Triggered task listNote: this field may return null, indicating that no valid values can be obtained.
	TriggerTasks []GetAlarmPolicyTriggerTask `pulumi:"triggerTasks"`
}

func LookupAlarmPolicyOutput(ctx *pulumi.Context, args LookupAlarmPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupAlarmPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlarmPolicyResult, error) {
			args := v.(LookupAlarmPolicyArgs)
			r, err := LookupAlarmPolicy(ctx, &args, opts...)
			var s LookupAlarmPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlarmPolicyResultOutput)
}

// A collection of arguments for invoking getAlarmPolicy.
type LookupAlarmPolicyOutputArgs struct {
	// The alarm object list, which is a JSON string. The outer array corresponds to multiple instances, and the inner array is the dimension of an object.For example, 'CVM - Basic Monitor' can be written as: [ {Dimensions: {unInstanceId: ins-qr8d555g}}, {Dimensions: {unInstanceId: ins-qr8d555h}} ]You can also refer to the 'Example 2' below.For more information on the parameter samples of different Tencent Cloud services, see [Product Policy Type and Dimension Information](https://www.tencentcloud.com/document/product/248/39565?has_map=1).Note: If 1 is passed in for NeedCorrespondence, the relationship between a policy and an instance needs to be returned. You can pass in up to 20 alarm object dimensions to avoid request timeout.
	Dimensions pulumi.StringPtrInput `pulumi:"dimensions"`
	// Filter by alarm status. Valid values: [1]: enabled; [0]: disabled; [0, 1]: all.
	Enables pulumi.IntArrayInput `pulumi:"enables"`
	// Sort by field. For example, to sort by the last modification time, use Field: UpdateTime.
	Field pulumi.StringPtrInput `pulumi:"field"`
	// Instance group ID.
	InstanceGroupId pulumi.IntPtrInput `pulumi:"instanceGroupId"`
	// Value fixed at monitor.
	Module pulumi.StringInput `pulumi:"module"`
	// Filter by monitor type. Valid values: MT_QCE (Tencent Cloud service monitoring). If this parameter is left empty, all will be queried by default.
	MonitorTypes pulumi.StringArrayInput `pulumi:"monitorTypes"`
	// Filter by namespace. For the values of different policy types, please see:[Poicy Type List](https://www.tencentcloud.com/document/product/248/39565?has_map=1).
	Namespaces pulumi.StringArrayInput `pulumi:"namespaces"`
	// Whether the relationship between a policy and the input parameter filter dimension is required. 1: Yes. 0: No. Default value: 0.
	NeedCorrespondence pulumi.IntPtrInput `pulumi:"needCorrespondence"`
	// Whether the returned result needs to filter policies associated with all objects. Valid values: 1 (Yes), 0 (No).
	NotBindAll pulumi.IntPtrInput `pulumi:"notBindAll"`
	// If 1 is passed in, alarm policies with no notification rules configured are queried. If it is left empty or other values are passed in, all alarm policies are queried.
	NotBindingNoticeRule pulumi.IntPtrInput `pulumi:"notBindingNoticeRule"`
	// Whether the returned result needs to filter policies associated with instance groups. Valid values: 1 (Yes), 0 (No).
	NotInstanceGroup pulumi.IntPtrInput `pulumi:"notInstanceGroup"`
	// List of the notification template IDs, which can be obtained by querying the notification template list.It can be queried with the API [DescribeAlarmNotices](https://www.tencentcloud.com/document/product/248/39300).
	NoticeIds pulumi.StringArrayInput `pulumi:"noticeIds"`
	// Filter by quick alarm policy. If this parameter is left empty, all policies are displayed. ONECLICK: Display quick alarm policies; NOT_ONECLICK: Display non-quick alarm policies.
	OneClickPolicyTypes pulumi.StringArrayInput `pulumi:"oneClickPolicyTypes"`
	// Sort order. Valid values: ASC (ascending), DESC (descending).
	Order pulumi.StringPtrInput `pulumi:"order"`
	// Fuzzy search by policy name.
	PolicyName pulumi.StringPtrInput `pulumi:"policyName"`
	// Filter by default policy. Valid values: DEFAULT (display default policy), NOT_DEFAULT (display non-default policies). If this parameter is left empty, all policies will be displayed.
	PolicyTypes pulumi.StringArrayInput `pulumi:"policyTypes"`
	// ID array of the policy project, which can be viewed on the following page: [Project Management](https://console.tencentcloud.com/project).
	ProjectIds pulumi.IntArrayInput `pulumi:"projectIds"`
	// ID of the TencentCloud Managed Service for Prometheus instance, which is used for customizing a metric policy.
	PromInsId pulumi.StringPtrInput `pulumi:"promInsId"`
	// Search by recipient group. You can get the user group list with the API [ListGroups](https://www.tencentcloud.com/document/product/598/34589?from_cn_redirect=1) in 'Cloud Access Management' or query the user group list where a sub-user is in with the API [ListGroupsForUser](https://www.tencentcloud.com/document/product/598/34588?from_cn_redirect=1). The GroupId field in the returned result should be entered here.
	ReceiverGroups pulumi.IntArrayInput `pulumi:"receiverGroups"`
	// Search by schedule.
	ReceiverOnCallFormIds pulumi.StringArrayInput `pulumi:"receiverOnCallFormIds"`
	// Search by recipient. You can get the user list with the API [ListUsers](https://www.tencentcloud.com/document/product/598/34587?from_cn_redirect=1) in 'Cloud Access Management' or query the sub-user information with the API [GetUser](https://www.tencentcloud.com/document/product/598/34590?from_cn_redirect=1). The Uid field in the returned result should be entered here.
	ReceiverUids pulumi.IntArrayInput `pulumi:"receiverUids"`
	// Used to save results.
	ResultOutputFile pulumi.StringPtrInput `pulumi:"resultOutputFile"`
	// Filter by trigger condition. Valid values: STATIC (display policies with static threshold), DYNAMIC (display policies with dynamic threshold). If this parameter is left empty, all policies will be displayed.
	RuleTypes pulumi.StringArrayInput `pulumi:"ruleTypes"`
	// Filter alarm policy by triggered task (such as auto scaling task). Up to 10 tasks can be specified.
	TriggerTasks GetAlarmPolicyTriggerTaskArrayInput `pulumi:"triggerTasks"`
}

func (LookupAlarmPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getAlarmPolicy.
type LookupAlarmPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupAlarmPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlarmPolicyResult)(nil)).Elem()
}

func (o LookupAlarmPolicyResultOutput) ToLookupAlarmPolicyResultOutput() LookupAlarmPolicyResultOutput {
	return o
}

func (o LookupAlarmPolicyResultOutput) ToLookupAlarmPolicyResultOutputWithContext(ctx context.Context) LookupAlarmPolicyResultOutput {
	return o
}

// JSON string generated by serializing the AlarmPolicyDimension two-dimensional array. The one-dimensional arrays are in OR relationship, and the elements in a one-dimensional array are in AND relationshipNote: this field may return null, indicating that no valid values can be obtained.
func (o LookupAlarmPolicyResultOutput) Dimensions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.Dimensions }).(pulumi.StringPtrOutput)
}

// Status. Valid values: 0 (disabled), 1 (enabled). Default value: 1 (enabled). This parameter can be left empty.
func (o LookupAlarmPolicyResultOutput) Enables() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.Enables }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.Field }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAlarmPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Instance group IDNote: this field may return null, indicating that no valid values can be obtained.
func (o LookupAlarmPolicyResultOutput) InstanceGroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.InstanceGroupId }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) Module() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) string { return v.Module }).(pulumi.StringOutput)
}

func (o LookupAlarmPolicyResultOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) Namespaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.Namespaces }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) NeedCorrespondence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NeedCorrespondence }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) NotBindAll() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NotBindAll }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) NotBindingNoticeRule() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NotBindingNoticeRule }).(pulumi.IntPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) NotInstanceGroup() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *int { return v.NotInstanceGroup }).(pulumi.IntPtrOutput)
}

// Notification rule ID listNote: this field may return null, indicating that no valid values can be obtained.
func (o LookupAlarmPolicyResultOutput) NoticeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.NoticeIds }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) OneClickPolicyTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.OneClickPolicyTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.Order }).(pulumi.StringPtrOutput)
}

// Policy array.
func (o LookupAlarmPolicyResultOutput) Policies() GetAlarmPolicyPolicyArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []GetAlarmPolicyPolicy { return v.Policies }).(GetAlarmPolicyPolicyArrayOutput)
}

// Alarm policy nameNote: this field may return null, indicating that no valid values can be obtained.
func (o LookupAlarmPolicyResultOutput) PolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.PolicyName }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) PolicyTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.PolicyTypes }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ProjectIds() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.ProjectIds }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) PromInsId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.PromInsId }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) ReceiverGroups() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.ReceiverGroups }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ReceiverOnCallFormIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.ReceiverOnCallFormIds }).(pulumi.StringArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ReceiverUids() pulumi.IntArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []int { return v.ReceiverUids }).(pulumi.IntArrayOutput)
}

func (o LookupAlarmPolicyResultOutput) ResultOutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) *string { return v.ResultOutputFile }).(pulumi.StringPtrOutput)
}

func (o LookupAlarmPolicyResultOutput) RuleTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []string { return v.RuleTypes }).(pulumi.StringArrayOutput)
}

// Triggered task listNote: this field may return null, indicating that no valid values can be obtained.
func (o LookupAlarmPolicyResultOutput) TriggerTasks() GetAlarmPolicyTriggerTaskArrayOutput {
	return o.ApplyT(func(v LookupAlarmPolicyResult) []GetAlarmPolicyTriggerTask { return v.TriggerTasks }).(GetAlarmPolicyTriggerTaskArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlarmPolicyResultOutput{})
}
