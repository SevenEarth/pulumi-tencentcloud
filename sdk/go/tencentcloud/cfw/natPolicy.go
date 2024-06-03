// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cfw

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/internal"
)

// Provides a resource to create a cfw natPolicy
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
//	"github.com/tencentcloudstack/pulumi-tencentcloud/sdk/go/tencentcloud/Cfw"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cfw.NewNatPolicy(ctx, "example", &Cfw.NatPolicyArgs{
//				Description:   pulumi.String("policy description."),
//				Direction:     pulumi.Int(1),
//				Enable:        pulumi.String("true"),
//				Port:          pulumi.String("-1/-1"),
//				Protocol:      pulumi.String("TCP"),
//				RuleAction:    pulumi.String("drop"),
//				SourceContent: pulumi.String("1.1.1.1/0"),
//				SourceType:    pulumi.String("net"),
//				TargetContent: pulumi.String("0.0.0.0/0"),
//				TargetType:    pulumi.String("net"),
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
// cfw nat_policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import tencentcloud:Cfw/natPolicy:NatPolicy example nat_policy_id
// ```
type NatPolicy struct {
	pulumi.CustomResourceState

	// Description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Rule direction: 1, inbound; 0, outbound.
	Direction pulumi.IntOutput `pulumi:"direction"`
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable pulumi.StringPtrOutput `pulumi:"enable"`
	// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateId pulumi.StringOutput `pulumi:"paramTemplateId"`
	// The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
	Port pulumi.StringOutput `pulumi:"port"`
	// Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// Access source example: net:IP/CIDR(192.168.0.2).
	SourceContent pulumi.StringOutput `pulumi:"sourceContent"`
	// Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
	TargetContent pulumi.StringOutput `pulumi:"targetContent"`
	// Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
	TargetType pulumi.StringOutput `pulumi:"targetType"`
	// The unique id corresponding to the rule, no need to fill in when creating the rule.
	Uuid pulumi.IntOutput `pulumi:"uuid"`
}

// NewNatPolicy registers a new resource with the given unique name, arguments, and options.
func NewNatPolicy(ctx *pulumi.Context,
	name string, args *NatPolicyArgs, opts ...pulumi.ResourceOption) (*NatPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
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
	if args.TargetContent == nil {
		return nil, errors.New("invalid value for required argument 'TargetContent'")
	}
	if args.TargetType == nil {
		return nil, errors.New("invalid value for required argument 'TargetType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NatPolicy
	err := ctx.RegisterResource("tencentcloud:Cfw/natPolicy:NatPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatPolicy gets an existing NatPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatPolicyState, opts ...pulumi.ResourceOption) (*NatPolicy, error) {
	var resource NatPolicy
	err := ctx.ReadResource("tencentcloud:Cfw/natPolicy:NatPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatPolicy resources.
type natPolicyState struct {
	// Description.
	Description *string `pulumi:"description"`
	// Rule direction: 1, inbound; 0, outbound.
	Direction *int `pulumi:"direction"`
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable *string `pulumi:"enable"`
	// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateId *string `pulumi:"paramTemplateId"`
	// The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
	Port *string `pulumi:"port"`
	// Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
	Protocol *string `pulumi:"protocol"`
	// How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
	RuleAction *string `pulumi:"ruleAction"`
	// Access source example: net:IP/CIDR(192.168.0.2).
	SourceContent *string `pulumi:"sourceContent"`
	// Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
	SourceType *string `pulumi:"sourceType"`
	// Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
	TargetContent *string `pulumi:"targetContent"`
	// Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
	TargetType *string `pulumi:"targetType"`
	// The unique id corresponding to the rule, no need to fill in when creating the rule.
	Uuid *int `pulumi:"uuid"`
}

type NatPolicyState struct {
	// Description.
	Description pulumi.StringPtrInput
	// Rule direction: 1, inbound; 0, outbound.
	Direction pulumi.IntPtrInput
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable pulumi.StringPtrInput
	// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
	ParamTemplateId pulumi.StringPtrInput
	// The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
	Port pulumi.StringPtrInput
	// Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
	Protocol pulumi.StringPtrInput
	// How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
	RuleAction pulumi.StringPtrInput
	// Access source example: net:IP/CIDR(192.168.0.2).
	SourceContent pulumi.StringPtrInput
	// Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
	SourceType pulumi.StringPtrInput
	// Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
	TargetContent pulumi.StringPtrInput
	// Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
	TargetType pulumi.StringPtrInput
	// The unique id corresponding to the rule, no need to fill in when creating the rule.
	Uuid pulumi.IntPtrInput
}

func (NatPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*natPolicyState)(nil)).Elem()
}

type natPolicyArgs struct {
	// Description.
	Description *string `pulumi:"description"`
	// Rule direction: 1, inbound; 0, outbound.
	Direction int `pulumi:"direction"`
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable *string `pulumi:"enable"`
	// The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
	Port string `pulumi:"port"`
	// Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
	Protocol string `pulumi:"protocol"`
	// How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
	RuleAction string `pulumi:"ruleAction"`
	// Access source example: net:IP/CIDR(192.168.0.2).
	SourceContent string `pulumi:"sourceContent"`
	// Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
	SourceType string `pulumi:"sourceType"`
	// Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
	TargetContent string `pulumi:"targetContent"`
	// Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
	TargetType string `pulumi:"targetType"`
}

// The set of arguments for constructing a NatPolicy resource.
type NatPolicyArgs struct {
	// Description.
	Description pulumi.StringPtrInput
	// Rule direction: 1, inbound; 0, outbound.
	Direction pulumi.IntInput
	// Rule status, true means enabled, false means disabled. Default is true.
	Enable pulumi.StringPtrInput
	// The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
	Port pulumi.StringInput
	// Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
	Protocol pulumi.StringInput
	// How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
	RuleAction pulumi.StringInput
	// Access source example: net:IP/CIDR(192.168.0.2).
	SourceContent pulumi.StringInput
	// Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
	SourceType pulumi.StringInput
	// Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
	TargetContent pulumi.StringInput
	// Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
	TargetType pulumi.StringInput
}

func (NatPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natPolicyArgs)(nil)).Elem()
}

type NatPolicyInput interface {
	pulumi.Input

	ToNatPolicyOutput() NatPolicyOutput
	ToNatPolicyOutputWithContext(ctx context.Context) NatPolicyOutput
}

func (*NatPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NatPolicy)(nil)).Elem()
}

func (i *NatPolicy) ToNatPolicyOutput() NatPolicyOutput {
	return i.ToNatPolicyOutputWithContext(context.Background())
}

func (i *NatPolicy) ToNatPolicyOutputWithContext(ctx context.Context) NatPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPolicyOutput)
}

// NatPolicyArrayInput is an input type that accepts NatPolicyArray and NatPolicyArrayOutput values.
// You can construct a concrete instance of `NatPolicyArrayInput` via:
//
//	NatPolicyArray{ NatPolicyArgs{...} }
type NatPolicyArrayInput interface {
	pulumi.Input

	ToNatPolicyArrayOutput() NatPolicyArrayOutput
	ToNatPolicyArrayOutputWithContext(context.Context) NatPolicyArrayOutput
}

type NatPolicyArray []NatPolicyInput

func (NatPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatPolicy)(nil)).Elem()
}

func (i NatPolicyArray) ToNatPolicyArrayOutput() NatPolicyArrayOutput {
	return i.ToNatPolicyArrayOutputWithContext(context.Background())
}

func (i NatPolicyArray) ToNatPolicyArrayOutputWithContext(ctx context.Context) NatPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPolicyArrayOutput)
}

// NatPolicyMapInput is an input type that accepts NatPolicyMap and NatPolicyMapOutput values.
// You can construct a concrete instance of `NatPolicyMapInput` via:
//
//	NatPolicyMap{ "key": NatPolicyArgs{...} }
type NatPolicyMapInput interface {
	pulumi.Input

	ToNatPolicyMapOutput() NatPolicyMapOutput
	ToNatPolicyMapOutputWithContext(context.Context) NatPolicyMapOutput
}

type NatPolicyMap map[string]NatPolicyInput

func (NatPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatPolicy)(nil)).Elem()
}

func (i NatPolicyMap) ToNatPolicyMapOutput() NatPolicyMapOutput {
	return i.ToNatPolicyMapOutputWithContext(context.Background())
}

func (i NatPolicyMap) ToNatPolicyMapOutputWithContext(ctx context.Context) NatPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatPolicyMapOutput)
}

type NatPolicyOutput struct{ *pulumi.OutputState }

func (NatPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatPolicy)(nil)).Elem()
}

func (o NatPolicyOutput) ToNatPolicyOutput() NatPolicyOutput {
	return o
}

func (o NatPolicyOutput) ToNatPolicyOutputWithContext(ctx context.Context) NatPolicyOutput {
	return o
}

// Description.
func (o NatPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Rule direction: 1, inbound; 0, outbound.
func (o NatPolicyOutput) Direction() pulumi.IntOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.IntOutput { return v.Direction }).(pulumi.IntOutput)
}

// Rule status, true means enabled, false means disabled. Default is true.
func (o NatPolicyOutput) Enable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringPtrOutput { return v.Enable }).(pulumi.StringPtrOutput)
}

// Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
func (o NatPolicyOutput) ParamTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.ParamTemplateId }).(pulumi.StringOutput)
}

// The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
func (o NatPolicyOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
func (o NatPolicyOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
func (o NatPolicyOutput) RuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.RuleAction }).(pulumi.StringOutput)
}

// Access source example: net:IP/CIDR(192.168.0.2).
func (o NatPolicyOutput) SourceContent() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.SourceContent }).(pulumi.StringOutput)
}

// Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
func (o NatPolicyOutput) SourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.SourceType }).(pulumi.StringOutput)
}

// Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
func (o NatPolicyOutput) TargetContent() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.TargetContent }).(pulumi.StringOutput)
}

// Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
func (o NatPolicyOutput) TargetType() pulumi.StringOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.StringOutput { return v.TargetType }).(pulumi.StringOutput)
}

// The unique id corresponding to the rule, no need to fill in when creating the rule.
func (o NatPolicyOutput) Uuid() pulumi.IntOutput {
	return o.ApplyT(func(v *NatPolicy) pulumi.IntOutput { return v.Uuid }).(pulumi.IntOutput)
}

type NatPolicyArrayOutput struct{ *pulumi.OutputState }

func (NatPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NatPolicy)(nil)).Elem()
}

func (o NatPolicyArrayOutput) ToNatPolicyArrayOutput() NatPolicyArrayOutput {
	return o
}

func (o NatPolicyArrayOutput) ToNatPolicyArrayOutputWithContext(ctx context.Context) NatPolicyArrayOutput {
	return o
}

func (o NatPolicyArrayOutput) Index(i pulumi.IntInput) NatPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NatPolicy {
		return vs[0].([]*NatPolicy)[vs[1].(int)]
	}).(NatPolicyOutput)
}

type NatPolicyMapOutput struct{ *pulumi.OutputState }

func (NatPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NatPolicy)(nil)).Elem()
}

func (o NatPolicyMapOutput) ToNatPolicyMapOutput() NatPolicyMapOutput {
	return o
}

func (o NatPolicyMapOutput) ToNatPolicyMapOutputWithContext(ctx context.Context) NatPolicyMapOutput {
	return o
}

func (o NatPolicyMapOutput) MapIndex(k pulumi.StringInput) NatPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NatPolicy {
		return vs[0].(map[string]*NatPolicy)[vs[1].(string)]
	}).(NatPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NatPolicyInput)(nil)).Elem(), &NatPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatPolicyArrayInput)(nil)).Elem(), NatPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NatPolicyMapInput)(nil)).Elem(), NatPolicyMap{})
	pulumi.RegisterOutputType(NatPolicyOutput{})
	pulumi.RegisterOutputType(NatPolicyArrayOutput{})
	pulumi.RegisterOutputType(NatPolicyMapOutput{})
}
