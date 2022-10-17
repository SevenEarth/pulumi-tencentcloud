// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create security group rule.
 *
 * ## Example Usage
 *
 * Source is CIDR ip
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const sglab1Group = new tencentcloud.security.Group("sglab1Group", {
 *     description: "favourite sg_1",
 *     projectId: 0,
 * });
 * const sglab1GroupRule = new tencentcloud.security.GroupRule("sglab1GroupRule", {
 *     securityGroupId: sglab1Group.id,
 *     type: "ingress",
 *     cidrIp: "10.0.0.0/16",
 *     ipProtocol: "TCP",
 *     portRange: "80",
 *     policy: "ACCEPT",
 *     description: "favourite sg rule_1",
 * });
 * ```
 *
 * Source is a security group id
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const sglab2Group = new tencentcloud.security.Group("sglab2Group", {
 *     description: "favourite sg_2",
 *     projectId: 0,
 * });
 * const sglab3 = new tencentcloud.security.Group("sglab3", {
 *     description: "favourite sg_3",
 *     projectId: 0,
 * });
 * const sglab2GroupRule = new tencentcloud.security.GroupRule("sglab2GroupRule", {
 *     securityGroupId: sglab2Group.id,
 *     type: "ingress",
 *     ipProtocol: "TCP",
 *     portRange: "80",
 *     policy: "ACCEPT",
 *     sourceSgid: sglab3.id,
 *     description: "favourite sg rule_2",
 * });
 * ```
 */
export class GroupRule extends pulumi.CustomResource {
    /**
     * Get an existing GroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupRuleState, opts?: pulumi.CustomResourceOptions): GroupRule {
        return new GroupRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Security/groupRule:GroupRule';

    /**
     * Returns true if the given object is an instance of GroupRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupRule.__pulumiType;
    }

    /**
     * ID of the address template, and confilicts with `sourceSgid` and `cidrIp`.
     */
    public readonly addressTemplate!: pulumi.Output<outputs.Security.GroupRuleAddressTemplate>;
    /**
     * An IP address network or segment, and conflict with `sourceSgid` and `addressTemplate`.
     */
    public readonly cidrIp!: pulumi.Output<string | undefined>;
    /**
     * Description of the security group rule.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocolTemplate`.
     */
    public readonly ipProtocol!: pulumi.Output<string>;
    /**
     * Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocolTemplate`.
     */
    public readonly portRange!: pulumi.Output<string>;
    /**
     * ID of the address template, and conflict with `ipProtocol`, `portRange`.
     */
    public readonly protocolTemplate!: pulumi.Output<outputs.Security.GroupRuleProtocolTemplate>;
    /**
     * ID of the security group to be queried.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * ID of the nested security group, and conflicts with `cidrIp` and `addressTemplate`.
     */
    public readonly sourceSgid!: pulumi.Output<string>;
    /**
     * Type of the security group rule. Valid values: `ingress` and `egress`.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a GroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupRuleArgs | GroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupRuleState | undefined;
            resourceInputs["addressTemplate"] = state ? state.addressTemplate : undefined;
            resourceInputs["cidrIp"] = state ? state.cidrIp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipProtocol"] = state ? state.ipProtocol : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["portRange"] = state ? state.portRange : undefined;
            resourceInputs["protocolTemplate"] = state ? state.protocolTemplate : undefined;
            resourceInputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            resourceInputs["sourceSgid"] = state ? state.sourceSgid : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as GroupRuleArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.securityGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["addressTemplate"] = args ? args.addressTemplate : undefined;
            resourceInputs["cidrIp"] = args ? args.cidrIp : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipProtocol"] = args ? args.ipProtocol : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["portRange"] = args ? args.portRange : undefined;
            resourceInputs["protocolTemplate"] = args ? args.protocolTemplate : undefined;
            resourceInputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            resourceInputs["sourceSgid"] = args ? args.sourceSgid : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupRule resources.
 */
export interface GroupRuleState {
    /**
     * ID of the address template, and confilicts with `sourceSgid` and `cidrIp`.
     */
    addressTemplate?: pulumi.Input<inputs.Security.GroupRuleAddressTemplate>;
    /**
     * An IP address network or segment, and conflict with `sourceSgid` and `addressTemplate`.
     */
    cidrIp?: pulumi.Input<string>;
    /**
     * Description of the security group rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocolTemplate`.
     */
    ipProtocol?: pulumi.Input<string>;
    /**
     * Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
     */
    policy?: pulumi.Input<string>;
    /**
     * Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocolTemplate`.
     */
    portRange?: pulumi.Input<string>;
    /**
     * ID of the address template, and conflict with `ipProtocol`, `portRange`.
     */
    protocolTemplate?: pulumi.Input<inputs.Security.GroupRuleProtocolTemplate>;
    /**
     * ID of the security group to be queried.
     */
    securityGroupId?: pulumi.Input<string>;
    /**
     * ID of the nested security group, and conflicts with `cidrIp` and `addressTemplate`.
     */
    sourceSgid?: pulumi.Input<string>;
    /**
     * Type of the security group rule. Valid values: `ingress` and `egress`.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupRule resource.
 */
export interface GroupRuleArgs {
    /**
     * ID of the address template, and confilicts with `sourceSgid` and `cidrIp`.
     */
    addressTemplate?: pulumi.Input<inputs.Security.GroupRuleAddressTemplate>;
    /**
     * An IP address network or segment, and conflict with `sourceSgid` and `addressTemplate`.
     */
    cidrIp?: pulumi.Input<string>;
    /**
     * Description of the security group rule.
     */
    description?: pulumi.Input<string>;
    /**
     * Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `protocolTemplate`.
     */
    ipProtocol?: pulumi.Input<string>;
    /**
     * Rule policy of security group. Valid values: `ACCEPT` and `DROP`.
     */
    policy: pulumi.Input<string>;
    /**
     * Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and confilicts with `protocolTemplate`.
     */
    portRange?: pulumi.Input<string>;
    /**
     * ID of the address template, and conflict with `ipProtocol`, `portRange`.
     */
    protocolTemplate?: pulumi.Input<inputs.Security.GroupRuleProtocolTemplate>;
    /**
     * ID of the security group to be queried.
     */
    securityGroupId: pulumi.Input<string>;
    /**
     * ID of the nested security group, and conflicts with `cidrIp` and `addressTemplate`.
     */
    sourceSgid?: pulumi.Input<string>;
    /**
     * Type of the security group rule. Valid values: `ingress` and `egress`.
     */
    type: pulumi.Input<string>;
}
