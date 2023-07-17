// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ckafka aclRule
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const aclRule = new tencentcloud.Ckafka.AclRule("acl_rule", {
 *     instanceId: "ckafka-xxx",
 *     isApplied: 1,
 *     pattern: "prefix",
 *     patternType: "PREFIXED",
 *     resourceType: "Topic",
 *     ruleLists: [{
 *         host: "*",
 *         operation: "All",
 *         permissionType: "Deny",
 *         principal: "User:*",
 *     }],
 *     ruleName: "RuleName",
 * });
 * ```
 *
 * ## Import
 *
 * ckafka acl_rule can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ckafka/aclRule:AclRule acl_rule acl_rule_id
 * ```
 */
export class AclRule extends pulumi.CustomResource {
    /**
     * Get an existing AclRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclRuleState, opts?: pulumi.CustomResourceOptions): AclRule {
        return new AclRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ckafka/aclRule:AclRule';

    /**
     * Returns true if the given object is an instance of AclRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AclRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AclRule.__pulumiType;
    }

    /**
     * instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether the preset ACL rule is applied to the newly added topic.
     */
    public readonly isApplied!: pulumi.Output<number | undefined>;
    /**
     * A value representing the prefix that the prefix matches.
     */
    public readonly pattern!: pulumi.Output<string | undefined>;
    /**
     * Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
     */
    public readonly patternType!: pulumi.Output<string>;
    /**
     * Acl resource type, currently only supports Topic, enumeration value list{Topic}.
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * List of configured ACL rules.
     */
    public readonly ruleLists!: pulumi.Output<outputs.Ckafka.AclRuleRuleList[]>;
    /**
     * rule name.
     */
    public readonly ruleName!: pulumi.Output<string>;

    /**
     * Create a AclRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclRuleArgs | AclRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclRuleState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isApplied"] = state ? state.isApplied : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["patternType"] = state ? state.patternType : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["ruleLists"] = state ? state.ruleLists : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
        } else {
            const args = argsOrState as AclRuleArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.patternType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'patternType'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.ruleLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleLists'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isApplied"] = args ? args.isApplied : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["patternType"] = args ? args.patternType : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["ruleLists"] = args ? args.ruleLists : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AclRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AclRule resources.
 */
export interface AclRuleState {
    /**
     * instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether the preset ACL rule is applied to the newly added topic.
     */
    isApplied?: pulumi.Input<number>;
    /**
     * A value representing the prefix that the prefix matches.
     */
    pattern?: pulumi.Input<string>;
    /**
     * Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
     */
    patternType?: pulumi.Input<string>;
    /**
     * Acl resource type, currently only supports Topic, enumeration value list{Topic}.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * List of configured ACL rules.
     */
    ruleLists?: pulumi.Input<pulumi.Input<inputs.Ckafka.AclRuleRuleList>[]>;
    /**
     * rule name.
     */
    ruleName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AclRule resource.
 */
export interface AclRuleArgs {
    /**
     * instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Whether the preset ACL rule is applied to the newly added topic.
     */
    isApplied?: pulumi.Input<number>;
    /**
     * A value representing the prefix that the prefix matches.
     */
    pattern?: pulumi.Input<string>;
    /**
     * Match type, currently supports prefix matching and preset strategy, enumeration value list{PREFIXED/PRESET}.
     */
    patternType: pulumi.Input<string>;
    /**
     * Acl resource type, currently only supports Topic, enumeration value list{Topic}.
     */
    resourceType: pulumi.Input<string>;
    /**
     * List of configured ACL rules.
     */
    ruleLists: pulumi.Input<pulumi.Input<inputs.Ckafka.AclRuleRuleList>[]>;
    /**
     * rule name.
     */
    ruleName: pulumi.Input<string>;
}
