// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cfw natPolicy
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Cfw.NatPolicy("example", {
 *     description: "policy description.",
 *     direction: 1,
 *     enable: "true",
 *     port: "-1/-1",
 *     protocol: "TCP",
 *     ruleAction: "drop",
 *     sourceContent: "1.1.1.1/0",
 *     sourceType: "net",
 *     targetContent: "0.0.0.0/0",
 *     targetType: "net",
 * });
 * ```
 *
 * ## Import
 *
 * cfw nat_policy can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cfw/natPolicy:NatPolicy example nat_policy_id
 * ```
 */
export class NatPolicy extends pulumi.CustomResource {
    /**
     * Get an existing NatPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NatPolicyState, opts?: pulumi.CustomResourceOptions): NatPolicy {
        return new NatPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cfw/natPolicy:NatPolicy';

    /**
     * Returns true if the given object is an instance of NatPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatPolicy.__pulumiType;
    }

    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Rule direction: 1, inbound; 0, outbound.
     */
    public readonly direction!: pulumi.Output<number>;
    /**
     * Rule status, true means enabled, false means disabled. Default is true.
     */
    public readonly enable!: pulumi.Output<string | undefined>;
    /**
     * Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
     */
    public /*out*/ readonly paramTemplateId!: pulumi.Output<string>;
    /**
     * The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
     */
    public readonly ruleAction!: pulumi.Output<string>;
    /**
     * Access source example: net:IP/CIDR(192.168.0.2).
     */
    public readonly sourceContent!: pulumi.Output<string>;
    /**
     * Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
     */
    public readonly sourceType!: pulumi.Output<string>;
    /**
     * Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
     */
    public readonly targetContent!: pulumi.Output<string>;
    /**
     * Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
     */
    public readonly targetType!: pulumi.Output<string>;
    /**
     * The unique id corresponding to the rule, no need to fill in when creating the rule.
     */
    public /*out*/ readonly uuid!: pulumi.Output<number>;

    /**
     * Create a NatPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NatPolicyArgs | NatPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NatPolicyState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["paramTemplateId"] = state ? state.paramTemplateId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["ruleAction"] = state ? state.ruleAction : undefined;
            resourceInputs["sourceContent"] = state ? state.sourceContent : undefined;
            resourceInputs["sourceType"] = state ? state.sourceType : undefined;
            resourceInputs["targetContent"] = state ? state.targetContent : undefined;
            resourceInputs["targetType"] = state ? state.targetType : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as NatPolicyArgs | undefined;
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.ruleAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleAction'");
            }
            if ((!args || args.sourceContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceContent'");
            }
            if ((!args || args.sourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceType'");
            }
            if ((!args || args.targetContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetContent'");
            }
            if ((!args || args.targetType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["ruleAction"] = args ? args.ruleAction : undefined;
            resourceInputs["sourceContent"] = args ? args.sourceContent : undefined;
            resourceInputs["sourceType"] = args ? args.sourceType : undefined;
            resourceInputs["targetContent"] = args ? args.targetContent : undefined;
            resourceInputs["targetType"] = args ? args.targetType : undefined;
            resourceInputs["paramTemplateId"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NatPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NatPolicy resources.
 */
export interface NatPolicyState {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Rule direction: 1, inbound; 0, outbound.
     */
    direction?: pulumi.Input<number>;
    /**
     * Rule status, true means enabled, false means disabled. Default is true.
     */
    enable?: pulumi.Input<string>;
    /**
     * Parameter template id. Note: This field may return null, indicating that no valid value can be obtained.
     */
    paramTemplateId?: pulumi.Input<string>;
    /**
     * The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
     */
    port?: pulumi.Input<string>;
    /**
     * Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
     */
    protocol?: pulumi.Input<string>;
    /**
     * How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
     */
    ruleAction?: pulumi.Input<string>;
    /**
     * Access source example: net:IP/CIDR(192.168.0.2).
     */
    sourceContent?: pulumi.Input<string>;
    /**
     * Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
     */
    sourceType?: pulumi.Input<string>;
    /**
     * Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
     */
    targetContent?: pulumi.Input<string>;
    /**
     * Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
     */
    targetType?: pulumi.Input<string>;
    /**
     * The unique id corresponding to the rule, no need to fill in when creating the rule.
     */
    uuid?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NatPolicy resource.
 */
export interface NatPolicyArgs {
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Rule direction: 1, inbound; 0, outbound.
     */
    direction: pulumi.Input<number>;
    /**
     * Rule status, true means enabled, false means disabled. Default is true.
     */
    enable?: pulumi.Input<string>;
    /**
     * The port for the access control policy. Value: -1/-1: All ports 80: Port 80.
     */
    port: pulumi.Input<string>;
    /**
     * Protocol. If Direction=1, optional values: TCP, UDP, ANY; If Direction=0, optional values: TCP, UDP, ICMP, ANY, HTTP, HTTPS, HTTP/HTTPS, SMTP, SMTPS, SMTP/SMTPS, FTP, and DNS.
     */
    protocol: pulumi.Input<string>;
    /**
     * How the traffic set in the access control policy passes through the cloud firewall. Values: accept: allow; drop: reject; log: observe.
     */
    ruleAction: pulumi.Input<string>;
    /**
     * Access source example: net:IP/CIDR(192.168.0.2).
     */
    sourceContent: pulumi.Input<string>;
    /**
     * Access source type: for inbound rules, the type can be net, location, vendor, template; for outbound rules, it can be net, instance, tag, template, group.
     */
    sourceType: pulumi.Input<string>;
    /**
     * Example of access purpose: net: IP/CIDR(192.168.0.2) domain: domain name rules, such as *.qq.com.
     */
    targetContent: pulumi.Input<string>;
    /**
     * Access purpose type: For inbound rules, the type can be net, instance, tag, template, group; for outbound rules, it can be net, location, vendor, template.
     */
    targetType: pulumi.Input<string>;
}
