// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfw

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a cfw vpcPolicy
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfw"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Cfw.NewVpcPolicy(ctx, "example", &Cfw.VpcPolicyArgs{
// 			Description:   pulumi.String("description."),
// 			DestContent:   pulumi.String("192.168.0.2"),
// 			DestType:      pulumi.String("net"),
// 			Enable:        pulumi.String("true"),
// 			FwGroupId:     pulumi.String("ALL"),
// 			Port:          pulumi.String("-1/-1"),
// 			Protocol:      pulumi.String("ANY"),
// 			RuleAction:    pulumi.String("log"),
// 			SourceContent: pulumi.String("0.0.0.0/0"),
// 			SourceType:    pulumi.String("net"),
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
// cfw vpc_policy can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import tencentcloud:Cfw/vpcPolicy:VpcPolicy vpc_policy vpc_policy_id
// ```
type VpcPolicy struct {
	pulumi.CustomResourceState

	// Beta mission details. Note: This field may return null, indicating that no valid value can be obtained.
	BetaLists VpcPolicyBetaListArrayOutput `pulumi:"betaLists"`
	// Describe.
	Description pulumi.StringOutput `pulumi:"description"`
	// Access purpose example: net:IP/CIDR(192.168.0.2) domain:domain rule, for example*.qq.com.
	DestContent pulumi.StringOutput `pulumi:"destContent"`
	// Access purpose type, the type can be: net, template.
	DestType pulumi.StringOutput `pulumi:"destType"`
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable pulumi.StringPtrOutput `pulumi:"enable"`
	// Firewall instance ID where the rule takes effect. Default is ALL.
	FwGroupId pulumi.StringPtrOutput `pulumi:"fwGroupId"`
	// Firewall name.
	FwGroupName pulumi.StringOutput `pulumi:"fwGroupName"`
	// Uuid used internally, this field is generally not used.
	InternalUuid pulumi.IntOutput `pulumi:"internalUuid"`
	// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateId pulumi.StringOutput `pulumi:"paramTemplateId"`
	// Parameter template Name. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateName pulumi.StringOutput `pulumi:"paramTemplateName"`
	// The port for the access control policy. Value: -1/-1: All ports; 80: port 80.
	Port pulumi.StringOutput `pulumi:"port"`
	// Protocol, optional value:TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, DNS, TLS/SSL.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// How traffic set in the access control policy passes through the cloud firewall. Value: accept:accept, drop:drop, log:log.
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// Access source examplnet:IP/CIDR(192.168.0.2).
	SourceContent pulumi.StringOutput `pulumi:"sourceContent"`
	// Access source type, the type can be: net, template.
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// The unique id corresponding to the rule.
	Uuid pulumi.IntOutput `pulumi:"uuid"`
}

// NewVpcPolicy registers a new resource with the given unique name, arguments, and options.
func NewVpcPolicy(ctx *pulumi.Context,
	name string, args *VpcPolicyArgs, opts ...pulumi.ResourceOption) (*VpcPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.DestContent == nil {
		return nil, errors.New("invalid value for required argument 'DestContent'")
	}
	if args.DestType == nil {
		return nil, errors.New("invalid value for required argument 'DestType'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.RuleAction == nil {
		return nil, errors.New("invalid value for required argument 'RuleAction'")
	}
	if args.SourceContent == nil {
		return nil, errors.New("invalid value for required argument 'SourceContent'")
	}
	if args.SourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VpcPolicy
	err := ctx.RegisterResource("tencentcloud:Cfw/vpcPolicy:VpcPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPolicy gets an existing VpcPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPolicyState, opts ...pulumi.ResourceOption) (*VpcPolicy, error) {
	var resource VpcPolicy
	err := ctx.ReadResource("tencentcloud:Cfw/vpcPolicy:VpcPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPolicy resources.
type vpcPolicyState struct {
	// Beta mission details. Note: This field may return null, indicating that no valid value can be obtained.
	BetaLists []VpcPolicyBetaList `pulumi:"betaLists"`
	// Describe.
	Description *string `pulumi:"description"`
	// Access purpose example: net:IP/CIDR(192.168.0.2) domain:domain rule, for example*.qq.com.
	DestContent *string `pulumi:"destContent"`
	// Access purpose type, the type can be: net, template.
	DestType *string `pulumi:"destType"`
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable *string `pulumi:"enable"`
	// Firewall instance ID where the rule takes effect. Default is ALL.
	FwGroupId *string `pulumi:"fwGroupId"`
	// Firewall name.
	FwGroupName *string `pulumi:"fwGroupName"`
	// Uuid used internally, this field is generally not used.
	InternalUuid *int `pulumi:"internalUuid"`
	// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateId *string `pulumi:"paramTemplateId"`
	// Parameter template Name. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateName *string `pulumi:"paramTemplateName"`
	// The port for the access control policy. Value: -1/-1: All ports; 80: port 80.
	Port *string `pulumi:"port"`
	// Protocol, optional value:TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, DNS, TLS/SSL.
	Protocol *string `pulumi:"protocol"`
	// How traffic set in the access control policy passes through the cloud firewall. Value: accept:accept, drop:drop, log:log.
	RuleAction *string `pulumi:"ruleAction"`
	// Access source examplnet:IP/CIDR(192.168.0.2).
	SourceContent *string `pulumi:"sourceContent"`
	// Access source type, the type can be: net, template.
	SourceType *string `pulumi:"sourceType"`
	// The unique id corresponding to the rule.
	Uuid *int `pulumi:"uuid"`
}

type VpcPolicyState struct {
	// Beta mission details. Note: This field may return null, indicating that no valid value can be obtained.
	BetaLists VpcPolicyBetaListArrayInput
	// Describe.
	Description pulumi.StringPtrInput
	// Access purpose example: net:IP/CIDR(192.168.0.2) domain:domain rule, for example*.qq.com.
	DestContent pulumi.StringPtrInput
	// Access purpose type, the type can be: net, template.
	DestType pulumi.StringPtrInput
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable pulumi.StringPtrInput
	// Firewall instance ID where the rule takes effect. Default is ALL.
	FwGroupId pulumi.StringPtrInput
	// Firewall name.
	FwGroupName pulumi.StringPtrInput
	// Uuid used internally, this field is generally not used.
	InternalUuid pulumi.IntPtrInput
	// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateId pulumi.StringPtrInput
	// Parameter template Name. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateName pulumi.StringPtrInput
	// The port for the access control policy. Value: -1/-1: All ports; 80: port 80.
	Port pulumi.StringPtrInput
	// Protocol, optional value:TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, DNS, TLS/SSL.
	Protocol pulumi.StringPtrInput
	// How traffic set in the access control policy passes through the cloud firewall. Value: accept:accept, drop:drop, log:log.
	RuleAction pulumi.StringPtrInput
	// Access source examplnet:IP/CIDR(192.168.0.2).
	SourceContent pulumi.StringPtrInput
	// Access source type, the type can be: net, template.
	SourceType pulumi.StringPtrInput
	// The unique id corresponding to the rule.
	Uuid pulumi.IntPtrInput
}

func (VpcPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPolicyState)(nil)).Elem()
}

type vpcPolicyArgs struct {
	// Describe.
	Description string `pulumi:"description"`
	// Access purpose example: net:IP/CIDR(192.168.0.2) domain:domain rule, for example*.qq.com.
	DestContent string `pulumi:"destContent"`
	// Access purpose type, the type can be: net, template.
	DestType string `pulumi:"destType"`
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable *string `pulumi:"enable"`
	// Firewall instance ID where the rule takes effect. Default is ALL.
	FwGroupId *string `pulumi:"fwGroupId"`
	// The port for the access control policy. Value: -1/-1: All ports; 80: port 80.
	Port string `pulumi:"port"`
	// Protocol, optional value:TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, DNS, TLS/SSL.
	Protocol string `pulumi:"protocol"`
	// How traffic set in the access control policy passes through the cloud firewall. Value: accept:accept, drop:drop, log:log.
	RuleAction string `pulumi:"ruleAction"`
	// Access source examplnet:IP/CIDR(192.168.0.2).
	SourceContent string `pulumi:"sourceContent"`
	// Access source type, the type can be: net, template.
	SourceType string `pulumi:"sourceType"`
}

// The set of arguments for constructing a VpcPolicy resource.
type VpcPolicyArgs struct {
	// Describe.
	Description pulumi.StringInput
	// Access purpose example: net:IP/CIDR(192.168.0.2) domain:domain rule, for example*.qq.com.
	DestContent pulumi.StringInput
	// Access purpose type, the type can be: net, template.
	DestType pulumi.StringInput
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable pulumi.StringPtrInput
	// Firewall instance ID where the rule takes effect. Default is ALL.
	FwGroupId pulumi.StringPtrInput
	// The port for the access control policy. Value: -1/-1: All ports; 80: port 80.
	Port pulumi.StringInput
	// Protocol, optional value:TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, DNS, TLS/SSL.
	Protocol pulumi.StringInput
	// How traffic set in the access control policy passes through the cloud firewall. Value: accept:accept, drop:drop, log:log.
	RuleAction pulumi.StringInput
	// Access source examplnet:IP/CIDR(192.168.0.2).
	SourceContent pulumi.StringInput
	// Access source type, the type can be: net, template.
	SourceType pulumi.StringInput
}

func (VpcPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPolicyArgs)(nil)).Elem()
}

type VpcPolicyInput interface {
	pulumi.Input

	ToVpcPolicyOutput() VpcPolicyOutput
	ToVpcPolicyOutputWithContext(ctx context.Context) VpcPolicyOutput
}

func (*VpcPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPolicy)(nil)).Elem()
}

func (i *VpcPolicy) ToVpcPolicyOutput() VpcPolicyOutput {
	return i.ToVpcPolicyOutputWithContext(context.Background())
}

func (i *VpcPolicy) ToVpcPolicyOutputWithContext(ctx context.Context) VpcPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPolicyOutput)
}

// VpcPolicyArrayInput is an input type that accepts VpcPolicyArray and VpcPolicyArrayOutput values.
// You can construct a concrete instance of `VpcPolicyArrayInput` via:
//
//          VpcPolicyArray{ VpcPolicyArgs{...} }
type VpcPolicyArrayInput interface {
	pulumi.Input

	ToVpcPolicyArrayOutput() VpcPolicyArrayOutput
	ToVpcPolicyArrayOutputWithContext(context.Context) VpcPolicyArrayOutput
}

type VpcPolicyArray []VpcPolicyInput

func (VpcPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPolicy)(nil)).Elem()
}

func (i VpcPolicyArray) ToVpcPolicyArrayOutput() VpcPolicyArrayOutput {
	return i.ToVpcPolicyArrayOutputWithContext(context.Background())
}

func (i VpcPolicyArray) ToVpcPolicyArrayOutputWithContext(ctx context.Context) VpcPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPolicyArrayOutput)
}

// VpcPolicyMapInput is an input type that accepts VpcPolicyMap and VpcPolicyMapOutput values.
// You can construct a concrete instance of `VpcPolicyMapInput` via:
//
//          VpcPolicyMap{ "key": VpcPolicyArgs{...} }
type VpcPolicyMapInput interface {
	pulumi.Input

	ToVpcPolicyMapOutput() VpcPolicyMapOutput
	ToVpcPolicyMapOutputWithContext(context.Context) VpcPolicyMapOutput
}

type VpcPolicyMap map[string]VpcPolicyInput

func (VpcPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPolicy)(nil)).Elem()
}

func (i VpcPolicyMap) ToVpcPolicyMapOutput() VpcPolicyMapOutput {
	return i.ToVpcPolicyMapOutputWithContext(context.Background())
}

func (i VpcPolicyMap) ToVpcPolicyMapOutputWithContext(ctx context.Context) VpcPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPolicyMapOutput)
}

type VpcPolicyOutput struct{ *pulumi.OutputState }

func (VpcPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPolicy)(nil)).Elem()
}

func (o VpcPolicyOutput) ToVpcPolicyOutput() VpcPolicyOutput {
	return o
}

func (o VpcPolicyOutput) ToVpcPolicyOutputWithContext(ctx context.Context) VpcPolicyOutput {
	return o
}

// Beta mission details. Note: This field may return null, indicating that no valid value can be obtained.
func (o VpcPolicyOutput) BetaLists() VpcPolicyBetaListArrayOutput {
	return o.ApplyT(func(v *VpcPolicy) VpcPolicyBetaListArrayOutput { return v.BetaLists }).(VpcPolicyBetaListArrayOutput)
}

// Describe.
func (o VpcPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Access purpose example: net:IP/CIDR(192.168.0.2) domain:domain rule, for example*.qq.com.
func (o VpcPolicyOutput) DestContent() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.DestContent }).(pulumi.StringOutput)
}

// Access purpose type, the type can be: net, template.
func (o VpcPolicyOutput) DestType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.DestType }).(pulumi.StringOutput)
}

// Rule status, true means enabled, false means disabled. Default is true.
func (o VpcPolicyOutput) Enable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringPtrOutput { return v.Enable }).(pulumi.StringPtrOutput)
}

// Firewall instance ID where the rule takes effect. Default is ALL.
func (o VpcPolicyOutput) FwGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringPtrOutput { return v.FwGroupId }).(pulumi.StringPtrOutput)
}

// Firewall name.
func (o VpcPolicyOutput) FwGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.FwGroupName }).(pulumi.StringOutput)
}

// Uuid used internally, this field is generally not used.
func (o VpcPolicyOutput) InternalUuid() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.IntOutput { return v.InternalUuid }).(pulumi.IntOutput)
}

// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
func (o VpcPolicyOutput) ParamTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.ParamTemplateId }).(pulumi.StringOutput)
}

// Parameter template Name. Note: This field may return null, indicating that no valid value can be obtained.
func (o VpcPolicyOutput) ParamTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.ParamTemplateName }).(pulumi.StringOutput)
}

// The port for the access control policy. Value: -1/-1: All ports; 80: port 80.
func (o VpcPolicyOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// Protocol, optional value:TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, DNS, TLS/SSL.
func (o VpcPolicyOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// How traffic set in the access control policy passes through the cloud firewall. Value: accept:accept, drop:drop, log:log.
func (o VpcPolicyOutput) RuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.RuleAction }).(pulumi.StringOutput)
}

// Access source examplnet:IP/CIDR(192.168.0.2).
func (o VpcPolicyOutput) SourceContent() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.SourceContent }).(pulumi.StringOutput)
}

// Access source type, the type can be: net, template.
func (o VpcPolicyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

// The unique id corresponding to the rule.
func (o VpcPolicyOutput) Uuid() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPolicy) pulumi.IntOutput { return v.Uuid }).(pulumi.IntOutput)
}

type VpcPolicyArrayOutput struct{ *pulumi.OutputState }

func (VpcPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPolicy)(nil)).Elem()
}

func (o VpcPolicyArrayOutput) ToVpcPolicyArrayOutput() VpcPolicyArrayOutput {
	return o
}

func (o VpcPolicyArrayOutput) ToVpcPolicyArrayOutputWithContext(ctx context.Context) VpcPolicyArrayOutput {
	return o
}

func (o VpcPolicyArrayOutput) Index(i pulumi.IntInput) VpcPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPolicy {
		return vs[0].([]*VpcPolicy)[vs[1].(int)]
	}).(VpcPolicyOutput)
}

type VpcPolicyMapOutput struct{ *pulumi.OutputState }

func (VpcPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPolicy)(nil)).Elem()
}

func (o VpcPolicyMapOutput) ToVpcPolicyMapOutput() VpcPolicyMapOutput {
	return o
}

func (o VpcPolicyMapOutput) ToVpcPolicyMapOutputWithContext(ctx context.Context) VpcPolicyMapOutput {
	return o
}

func (o VpcPolicyMapOutput) MapIndex(k pulumi.StringInput) VpcPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPolicy {
		return vs[0].(map[string]*VpcPolicy)[vs[1].(string)]
	}).(VpcPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPolicyInput)(nil)).Elem(), &VpcPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPolicyArrayInput)(nil)).Elem(), VpcPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPolicyMapInput)(nil)).Elem(), VpcPolicyMap{})
	pulumi.RegisterOutputType(VpcPolicyOutput{})
	pulumi.RegisterOutputType(VpcPolicyArrayOutput{})
	pulumi.RegisterOutputType(VpcPolicyMapOutput{})
}
