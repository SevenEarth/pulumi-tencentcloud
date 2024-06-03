// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a security policy rule.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const fooProxy = new tencentcloud.gaap.Proxy("fooProxy", {
 *     bandwidth: 10,
 *     concurrent: 2,
 *     accessRegion: "SouthChina",
 *     realserverRegion: "NorthChina",
 * });
 * const fooSecurityPolicy = new tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy", {
 *     proxyId: fooProxy.id,
 *     action: "ACCEPT",
 * });
 * const fooSecurityRule = new tencentcloud.gaap.SecurityRule("fooSecurityRule", {
 *     policyId: fooSecurityPolicy.id,
 *     cidrIp: "1.1.1.1",
 *     action: "ACCEPT",
 *     protocol: "TCP",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * GAAP security rule can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Gaap/securityRule:SecurityRule tencentcloud_gaap_security_rule.foo sr-xxxxxxxx
 * ```
 */
export class SecurityRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityRuleState, opts?: pulumi.CustomResourceOptions): SecurityRule {
        return new SecurityRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Gaap/securityRule:SecurityRule';

    /**
     * Returns true if the given object is an instance of SecurityRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityRule.__pulumiType;
    }

    /**
     * Policy of the rule. Valid value: `ACCEPT` and `DROP`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * A network address block of the request source.
     */
    public readonly cidrIp!: pulumi.Output<string>;
    /**
     * Name of the security policy rule. Maximum length is 30.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of the security policy.
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
     */
    public readonly port!: pulumi.Output<string | undefined>;
    /**
     * Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
     */
    public readonly protocol!: pulumi.Output<string | undefined>;

    /**
     * Create a SecurityRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityRuleArgs | SecurityRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["cidrIp"] = state ? state.cidrIp : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
        } else {
            const args = argsOrState as SecurityRuleArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.cidrIp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cidrIp'");
            }
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["cidrIp"] = args ? args.cidrIp : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityRule resources.
 */
export interface SecurityRuleState {
    /**
     * Policy of the rule. Valid value: `ACCEPT` and `DROP`.
     */
    action?: pulumi.Input<string>;
    /**
     * A network address block of the request source.
     */
    cidrIp?: pulumi.Input<string>;
    /**
     * Name of the security policy rule. Maximum length is 30.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the security policy.
     */
    policyId?: pulumi.Input<string>;
    /**
     * Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
     */
    port?: pulumi.Input<string>;
    /**
     * Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
     */
    protocol?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityRule resource.
 */
export interface SecurityRuleArgs {
    /**
     * Policy of the rule. Valid value: `ACCEPT` and `DROP`.
     */
    action: pulumi.Input<string>;
    /**
     * A network address block of the request source.
     */
    cidrIp: pulumi.Input<string>;
    /**
     * Name of the security policy rule. Maximum length is 30.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the security policy.
     */
    policyId: pulumi.Input<string>;
    /**
     * Target port. Default value is `ALL`. Valid examples: `80`, `80,443` and `3306-20000`.
     */
    port?: pulumi.Input<string>;
    /**
     * Protocol of the security policy rule. Default value is `ALL`. Valid value: `TCP`, `UDP` and `ALL`.
     */
    protocol?: pulumi.Input<string>;
}
