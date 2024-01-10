// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a waf autoDenyRules
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = new tencentcloud.Waf.AutoDenyRules("example", {
 *     attackThreshold: 20,
 *     denyTimeThreshold: 5,
 *     domain: "demo.waf.com",
 *     timeThreshold: 12,
 * });
 * ```
 *
 * ## Import
 *
 * waf auto_deny_rules can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Waf/autoDenyRules:AutoDenyRules example demo.waf.com
 * ```
 */
export class AutoDenyRules extends pulumi.CustomResource {
    /**
     * Get an existing AutoDenyRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoDenyRulesState, opts?: pulumi.CustomResourceOptions): AutoDenyRules {
        return new AutoDenyRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Waf/autoDenyRules:AutoDenyRules';

    /**
     * Returns true if the given object is an instance of AutoDenyRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoDenyRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoDenyRules.__pulumiType;
    }

    /**
     * The threshold number of attacks that triggers IP autodeny, ranging from 2 to 100 times.
     */
    public readonly attackThreshold!: pulumi.Output<number>;
    /**
     * The IP autodeny time after triggering the IP autodeny, ranging from 5 to 360 minutes.
     */
    public readonly denyTimeThreshold!: pulumi.Output<number>;
    /**
     * Domain.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * IP autodeny statistical time, ranging from 1-60 minutes.
     */
    public readonly timeThreshold!: pulumi.Output<number>;

    /**
     * Create a AutoDenyRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoDenyRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoDenyRulesArgs | AutoDenyRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoDenyRulesState | undefined;
            resourceInputs["attackThreshold"] = state ? state.attackThreshold : undefined;
            resourceInputs["denyTimeThreshold"] = state ? state.denyTimeThreshold : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["timeThreshold"] = state ? state.timeThreshold : undefined;
        } else {
            const args = argsOrState as AutoDenyRulesArgs | undefined;
            if ((!args || args.attackThreshold === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attackThreshold'");
            }
            if ((!args || args.denyTimeThreshold === undefined) && !opts.urn) {
                throw new Error("Missing required property 'denyTimeThreshold'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.timeThreshold === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeThreshold'");
            }
            resourceInputs["attackThreshold"] = args ? args.attackThreshold : undefined;
            resourceInputs["denyTimeThreshold"] = args ? args.denyTimeThreshold : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["timeThreshold"] = args ? args.timeThreshold : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutoDenyRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoDenyRules resources.
 */
export interface AutoDenyRulesState {
    /**
     * The threshold number of attacks that triggers IP autodeny, ranging from 2 to 100 times.
     */
    attackThreshold?: pulumi.Input<number>;
    /**
     * The IP autodeny time after triggering the IP autodeny, ranging from 5 to 360 minutes.
     */
    denyTimeThreshold?: pulumi.Input<number>;
    /**
     * Domain.
     */
    domain?: pulumi.Input<string>;
    /**
     * IP autodeny statistical time, ranging from 1-60 minutes.
     */
    timeThreshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AutoDenyRules resource.
 */
export interface AutoDenyRulesArgs {
    /**
     * The threshold number of attacks that triggers IP autodeny, ranging from 2 to 100 times.
     */
    attackThreshold: pulumi.Input<number>;
    /**
     * The IP autodeny time after triggering the IP autodeny, ranging from 5 to 360 minutes.
     */
    denyTimeThreshold: pulumi.Input<number>;
    /**
     * Domain.
     */
    domain: pulumi.Input<string>;
    /**
     * IP autodeny statistical time, ranging from 1-60 minutes.
     */
    timeThreshold: pulumi.Input<number>;
}
