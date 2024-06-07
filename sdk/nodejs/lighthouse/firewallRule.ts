// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a lighthouse firewall rule
 *
 * > **NOTE:**  Use an empty template to clean up the default rules before using this resource manage firewall rules.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const firewallRule = new tencentcloud.lighthouse.FirewallRule("firewallRule", {
 *     firewallRules: [
 *         {
 *             action: "ACCEPT",
 *             cidrBlock: "10.0.0.1",
 *             firewallRuleDescription: "description 1",
 *             port: "80",
 *             protocol: "TCP",
 *         },
 *         {
 *             action: "ACCEPT",
 *             cidrBlock: "10.0.0.2",
 *             firewallRuleDescription: "description 2",
 *             port: "80",
 *             protocol: "TCP",
 *         },
 *     ],
 *     instanceId: "lhins-xxxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * lighthouse firewall_rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Lighthouse/firewallRule:FirewallRule firewall_rule lighthouse_instance_id
 * ```
 */
export class FirewallRule extends pulumi.CustomResource {
    /**
     * Get an existing FirewallRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FirewallRuleState, opts?: pulumi.CustomResourceOptions): FirewallRule {
        return new FirewallRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Lighthouse/firewallRule:FirewallRule';

    /**
     * Returns true if the given object is an instance of FirewallRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FirewallRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FirewallRule.__pulumiType;
    }

    /**
     * Firewall rule list.
     */
    public readonly firewallRules!: pulumi.Output<outputs.Lighthouse.FirewallRuleFirewallRule[]>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a FirewallRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FirewallRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FirewallRuleArgs | FirewallRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FirewallRuleState | undefined;
            resourceInputs["firewallRules"] = state ? state.firewallRules : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as FirewallRuleArgs | undefined;
            if ((!args || args.firewallRules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firewallRules'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["firewallRules"] = args ? args.firewallRules : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FirewallRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FirewallRule resources.
 */
export interface FirewallRuleState {
    /**
     * Firewall rule list.
     */
    firewallRules?: pulumi.Input<pulumi.Input<inputs.Lighthouse.FirewallRuleFirewallRule>[]>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FirewallRule resource.
 */
export interface FirewallRuleArgs {
    /**
     * Firewall rule list.
     */
    firewallRules: pulumi.Input<pulumi.Input<inputs.Lighthouse.FirewallRuleFirewallRule>[]>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
}
