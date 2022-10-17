// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a security policy of GAAP proxy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const fooProxy = new tencentcloud.gaap.Proxy("fooProxy", {
 *     bandwidth: 10,
 *     concurrent: 2,
 *     accessRegion: "SouthChina",
 *     realserverRegion: "NorthChina",
 * });
 * const fooSecurityPolicy = new tencentcloud.gaap.SecurityPolicy("fooSecurityPolicy", {
 *     proxyId: fooProxy.id,
 *     action: "DROP",
 * });
 * ```
 *
 * ## Import
 *
 * GAAP security policy can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Gaap/securityPolicy:SecurityPolicy tencentcloud_gaap_security_policy.foo pl-xxxx
 * ```
 */
export class SecurityPolicy extends pulumi.CustomResource {
    /**
     * Get an existing SecurityPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityPolicyState, opts?: pulumi.CustomResourceOptions): SecurityPolicy {
        return new SecurityPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Gaap/securityPolicy:SecurityPolicy';

    /**
     * Returns true if the given object is an instance of SecurityPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityPolicy.__pulumiType;
    }

    /**
     * Default policy. Valid value: `ACCEPT` and `DROP`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Indicates whether policy is enable, default value is `true`.
     */
    public readonly enable!: pulumi.Output<boolean | undefined>;
    /**
     * ID of the GAAP proxy.
     */
    public readonly proxyId!: pulumi.Output<string>;

    /**
     * Create a SecurityPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityPolicyArgs | SecurityPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecurityPolicyState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["proxyId"] = state ? state.proxyId : undefined;
        } else {
            const args = argsOrState as SecurityPolicyArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.proxyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proxyId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["proxyId"] = args ? args.proxyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecurityPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityPolicy resources.
 */
export interface SecurityPolicyState {
    /**
     * Default policy. Valid value: `ACCEPT` and `DROP`.
     */
    action?: pulumi.Input<string>;
    /**
     * Indicates whether policy is enable, default value is `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * ID of the GAAP proxy.
     */
    proxyId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityPolicy resource.
 */
export interface SecurityPolicyArgs {
    /**
     * Default policy. Valid value: `ACCEPT` and `DROP`.
     */
    action: pulumi.Input<string>;
    /**
     * Indicates whether policy is enable, default value is `true`.
     */
    enable?: pulumi.Input<boolean>;
    /**
     * ID of the GAAP proxy.
     */
    proxyId: pulumi.Input<string>;
}
