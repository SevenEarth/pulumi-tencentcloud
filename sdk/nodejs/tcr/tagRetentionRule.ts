// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tcr tag retention rule.
 *
 * ## Example Usage
 *
 * ### Create a tcr tag retention rule instance
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const exampleInstance = new tencentcloud.tcr.Instance("exampleInstance", {
 *     instanceType: "basic",
 *     deleteBucket: true,
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const exampleNamespace = new tencentcloud.tcr.Namespace("exampleNamespace", {
 *     instanceId: exampleInstance.id,
 *     isPublic: true,
 *     isAutoScan: true,
 *     isPreventVul: true,
 *     severity: "medium",
 *     cveWhitelistItems: [{
 *         cveId: "cve-xxxxx",
 *     }],
 * });
 * const myRule = new tencentcloud.tcr.TagRetentionRule("myRule", {
 *     registryId: exampleInstance.id,
 *     namespaceName: exampleNamespace.name,
 *     retentionRule: {
 *         key: "nDaysSinceLastPush",
 *         value: 2,
 *     },
 *     cronSetting: "daily",
 *     disabled: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class TagRetentionRule extends pulumi.CustomResource {
    /**
     * Get an existing TagRetentionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TagRetentionRuleState, opts?: pulumi.CustomResourceOptions): TagRetentionRule {
        return new TagRetentionRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tcr/tagRetentionRule:TagRetentionRule';

    /**
     * Returns true if the given object is an instance of TagRetentionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TagRetentionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TagRetentionRule.__pulumiType;
    }

    /**
     * Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
     */
    public readonly cronSetting!: pulumi.Output<string>;
    /**
     * Whether to disable the rule, with the default value of false.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Name of the namespace.
     */
    public readonly namespaceName!: pulumi.Output<string>;
    /**
     * The main instance ID.
     */
    public readonly registryId!: pulumi.Output<string>;
    /**
     * The ID of the retention task.
     */
    public /*out*/ readonly retentionId!: pulumi.Output<number>;
    /**
     * Retention Policy.
     */
    public readonly retentionRule!: pulumi.Output<outputs.Tcr.TagRetentionRuleRetentionRule>;

    /**
     * Create a TagRetentionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TagRetentionRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TagRetentionRuleArgs | TagRetentionRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TagRetentionRuleState | undefined;
            resourceInputs["cronSetting"] = state ? state.cronSetting : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["namespaceName"] = state ? state.namespaceName : undefined;
            resourceInputs["registryId"] = state ? state.registryId : undefined;
            resourceInputs["retentionId"] = state ? state.retentionId : undefined;
            resourceInputs["retentionRule"] = state ? state.retentionRule : undefined;
        } else {
            const args = argsOrState as TagRetentionRuleArgs | undefined;
            if ((!args || args.cronSetting === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cronSetting'");
            }
            if ((!args || args.namespaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if ((!args || args.registryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registryId'");
            }
            if ((!args || args.retentionRule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retentionRule'");
            }
            resourceInputs["cronSetting"] = args ? args.cronSetting : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["namespaceName"] = args ? args.namespaceName : undefined;
            resourceInputs["registryId"] = args ? args.registryId : undefined;
            resourceInputs["retentionRule"] = args ? args.retentionRule : undefined;
            resourceInputs["retentionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TagRetentionRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TagRetentionRule resources.
 */
export interface TagRetentionRuleState {
    /**
     * Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
     */
    cronSetting?: pulumi.Input<string>;
    /**
     * Whether to disable the rule, with the default value of false.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The Name of the namespace.
     */
    namespaceName?: pulumi.Input<string>;
    /**
     * The main instance ID.
     */
    registryId?: pulumi.Input<string>;
    /**
     * The ID of the retention task.
     */
    retentionId?: pulumi.Input<number>;
    /**
     * Retention Policy.
     */
    retentionRule?: pulumi.Input<inputs.Tcr.TagRetentionRuleRetentionRule>;
}

/**
 * The set of arguments for constructing a TagRetentionRule resource.
 */
export interface TagRetentionRuleArgs {
    /**
     * Execution cycle, currently only available selections are: manual; daily; weekly; monthly.
     */
    cronSetting: pulumi.Input<string>;
    /**
     * Whether to disable the rule, with the default value of false.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The Name of the namespace.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * The main instance ID.
     */
    registryId: pulumi.Input<string>;
    /**
     * Retention Policy.
     */
    retentionRule: pulumi.Input<inputs.Tcr.TagRetentionRuleRetentionRule>;
}
