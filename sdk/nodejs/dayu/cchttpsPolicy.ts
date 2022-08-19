// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this resource to create a dayu CC self-define https policy
 *
 * > **NOTE:** creating CC self-define https policy need a valid resource `tencentcloud.Dayu.L7Rule`; The resource only support Anti-DDoS of resource type `bgpip`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const testPolicy = new tencentcloud.dayu.CcHttpsPolicy("testPolicy", {
 *     resourceType: tencentcloud_dayu_l7_rule.test_rule.resource_type,
 *     resourceId: tencentcloud_dayu_l7_rule.test_rule.resource_id,
 *     ruleId: tencentcloud_dayu_l7_rule.test_rule.rule_id,
 *     domain: tencentcloud_dayu_l7_rule.test_rule.domain,
 *     action: "drop",
 *     "switch": true,
 *     ruleLists: [{
 *         skey: "cgi",
 *         operator: "include",
 *         value: "123",
 *     }],
 * });
 * ```
 */
export class CcHttpsPolicy extends pulumi.CustomResource {
    /**
     * Get an existing CcHttpsPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CcHttpsPolicyState, opts?: pulumi.CustomResourceOptions): CcHttpsPolicy {
        return new CcHttpsPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dayu/ccHttpsPolicy:CcHttpsPolicy';

    /**
     * Returns true if the given object is an instance of CcHttpsPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CcHttpsPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CcHttpsPolicy.__pulumiType;
    }

    /**
     * Action mode. Valid values are `alg` and `drop`.
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Create time of the CC self-define https policy.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Ip of the CC self-define https policy.
     */
    public /*out*/ readonly ipLists!: pulumi.Output<string[]>;
    /**
     * Name of the CC self-define https policy. Length should between 1 and 20.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Id of the CC self-define https policy.
     */
    public /*out*/ readonly policyId!: pulumi.Output<string>;
    /**
     * ID of the resource that the CC self-define https policy works for.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
     */
    public readonly ruleId!: pulumi.Output<string>;
    /**
     * Rule list of the CC self-define https policy.
     */
    public readonly ruleLists!: pulumi.Output<outputs.Dayu.CcHttpsPolicyRuleList[]>;
    /**
     * Indicate the CC self-define https policy takes effect or not.
     */
    public readonly switch!: pulumi.Output<boolean | undefined>;

    /**
     * Create a CcHttpsPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CcHttpsPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CcHttpsPolicyArgs | CcHttpsPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CcHttpsPolicyState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["ipLists"] = state ? state.ipLists : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["resourceId"] = state ? state.resourceId : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["ruleId"] = state ? state.ruleId : undefined;
            resourceInputs["ruleLists"] = state ? state.ruleLists : undefined;
            resourceInputs["switch"] = state ? state.switch : undefined;
        } else {
            const args = argsOrState as CcHttpsPolicyArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.ruleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleId'");
            }
            if ((!args || args.ruleLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleLists'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["ruleId"] = args ? args.ruleId : undefined;
            resourceInputs["ruleLists"] = args ? args.ruleLists : undefined;
            resourceInputs["switch"] = args ? args.switch : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["ipLists"] = undefined /*out*/;
            resourceInputs["policyId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CcHttpsPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CcHttpsPolicy resources.
 */
export interface CcHttpsPolicyState {
    /**
     * Action mode. Valid values are `alg` and `drop`.
     */
    action?: pulumi.Input<string>;
    /**
     * Create time of the CC self-define https policy.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
     */
    domain?: pulumi.Input<string>;
    /**
     * Ip of the CC self-define https policy.
     */
    ipLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the CC self-define https policy. Length should between 1 and 20.
     */
    name?: pulumi.Input<string>;
    /**
     * Id of the CC self-define https policy.
     */
    policyId?: pulumi.Input<string>;
    /**
     * ID of the resource that the CC self-define https policy works for.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
     */
    ruleId?: pulumi.Input<string>;
    /**
     * Rule list of the CC self-define https policy.
     */
    ruleLists?: pulumi.Input<pulumi.Input<inputs.Dayu.CcHttpsPolicyRuleList>[]>;
    /**
     * Indicate the CC self-define https policy takes effect or not.
     */
    switch?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a CcHttpsPolicy resource.
 */
export interface CcHttpsPolicyArgs {
    /**
     * Action mode. Valid values are `alg` and `drop`.
     */
    action?: pulumi.Input<string>;
    /**
     * Domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
     */
    domain: pulumi.Input<string>;
    /**
     * Name of the CC self-define https policy. Length should between 1 and 20.
     */
    name?: pulumi.Input<string>;
    /**
     * ID of the resource that the CC self-define https policy works for.
     */
    resourceId: pulumi.Input<string>;
    /**
     * Type of the resource that the CC self-define https policy works for, valid value is `bgpip`.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Rule id of the domain that the CC self-define https policy works for, only valid when `protocol` is `https`.
     */
    ruleId: pulumi.Input<string>;
    /**
     * Rule list of the CC self-define https policy.
     */
    ruleLists: pulumi.Input<pulumi.Input<inputs.Dayu.CcHttpsPolicyRuleList>[]>;
    /**
     * Indicate the CC self-define https policy takes effect or not.
     */
    switch?: pulumi.Input<boolean>;
}
