// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cynosdb clusterPasswordComplexity
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const clusterPasswordComplexity = new tencentcloud.Cynosdb.ClusterPasswordComplexity("cluster_password_complexity", {
 *     clusterId: "cynosdbmysql-cgd2gpwr",
 *     validatePasswordDictionaries: [
 *         "cccc",
 *         "xxxx",
 *     ],
 *     validatePasswordLength: 8,
 *     validatePasswordMixedCaseCount: 1,
 *     validatePasswordNumberCount: 1,
 *     validatePasswordPolicy: "STRONG",
 *     validatePasswordSpecialCharCount: 1,
 * });
 * ```
 *
 * ## Import
 *
 * cynosdb cluster_password_complexity can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cynosdb/clusterPasswordComplexity:ClusterPasswordComplexity cluster_password_complexity cluster_password_complexity_id
 * ```
 */
export class ClusterPasswordComplexity extends pulumi.CustomResource {
    /**
     * Get an existing ClusterPasswordComplexity resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterPasswordComplexityState, opts?: pulumi.CustomResourceOptions): ClusterPasswordComplexity {
        return new ClusterPasswordComplexity(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/clusterPasswordComplexity:ClusterPasswordComplexity';

    /**
     * Returns true if the given object is an instance of ClusterPasswordComplexity.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterPasswordComplexity {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterPasswordComplexity.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Data dictionary.
     */
    public readonly validatePasswordDictionaries!: pulumi.Output<string[] | undefined>;
    /**
     * Password length.
     */
    public readonly validatePasswordLength!: pulumi.Output<number>;
    /**
     * Number of uppercase and lowercase characters.
     */
    public readonly validatePasswordMixedCaseCount!: pulumi.Output<number>;
    /**
     * Number of digits.
     */
    public readonly validatePasswordNumberCount!: pulumi.Output<number>;
    /**
     * Password strength (MEDIUM, STRONG).
     */
    public readonly validatePasswordPolicy!: pulumi.Output<string>;
    /**
     * Number of special characters.
     */
    public readonly validatePasswordSpecialCharCount!: pulumi.Output<number>;

    /**
     * Create a ClusterPasswordComplexity resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterPasswordComplexityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterPasswordComplexityArgs | ClusterPasswordComplexityState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterPasswordComplexityState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["validatePasswordDictionaries"] = state ? state.validatePasswordDictionaries : undefined;
            resourceInputs["validatePasswordLength"] = state ? state.validatePasswordLength : undefined;
            resourceInputs["validatePasswordMixedCaseCount"] = state ? state.validatePasswordMixedCaseCount : undefined;
            resourceInputs["validatePasswordNumberCount"] = state ? state.validatePasswordNumberCount : undefined;
            resourceInputs["validatePasswordPolicy"] = state ? state.validatePasswordPolicy : undefined;
            resourceInputs["validatePasswordSpecialCharCount"] = state ? state.validatePasswordSpecialCharCount : undefined;
        } else {
            const args = argsOrState as ClusterPasswordComplexityArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.validatePasswordLength === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validatePasswordLength'");
            }
            if ((!args || args.validatePasswordMixedCaseCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validatePasswordMixedCaseCount'");
            }
            if ((!args || args.validatePasswordNumberCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validatePasswordNumberCount'");
            }
            if ((!args || args.validatePasswordPolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validatePasswordPolicy'");
            }
            if ((!args || args.validatePasswordSpecialCharCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validatePasswordSpecialCharCount'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["validatePasswordDictionaries"] = args ? args.validatePasswordDictionaries : undefined;
            resourceInputs["validatePasswordLength"] = args ? args.validatePasswordLength : undefined;
            resourceInputs["validatePasswordMixedCaseCount"] = args ? args.validatePasswordMixedCaseCount : undefined;
            resourceInputs["validatePasswordNumberCount"] = args ? args.validatePasswordNumberCount : undefined;
            resourceInputs["validatePasswordPolicy"] = args ? args.validatePasswordPolicy : undefined;
            resourceInputs["validatePasswordSpecialCharCount"] = args ? args.validatePasswordSpecialCharCount : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterPasswordComplexity.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterPasswordComplexity resources.
 */
export interface ClusterPasswordComplexityState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Data dictionary.
     */
    validatePasswordDictionaries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Password length.
     */
    validatePasswordLength?: pulumi.Input<number>;
    /**
     * Number of uppercase and lowercase characters.
     */
    validatePasswordMixedCaseCount?: pulumi.Input<number>;
    /**
     * Number of digits.
     */
    validatePasswordNumberCount?: pulumi.Input<number>;
    /**
     * Password strength (MEDIUM, STRONG).
     */
    validatePasswordPolicy?: pulumi.Input<string>;
    /**
     * Number of special characters.
     */
    validatePasswordSpecialCharCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ClusterPasswordComplexity resource.
 */
export interface ClusterPasswordComplexityArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Data dictionary.
     */
    validatePasswordDictionaries?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Password length.
     */
    validatePasswordLength: pulumi.Input<number>;
    /**
     * Number of uppercase and lowercase characters.
     */
    validatePasswordMixedCaseCount: pulumi.Input<number>;
    /**
     * Number of digits.
     */
    validatePasswordNumberCount: pulumi.Input<number>;
    /**
     * Password strength (MEDIUM, STRONG).
     */
    validatePasswordPolicy: pulumi.Input<string>;
    /**
     * Number of special characters.
     */
    validatePasswordSpecialCharCount: pulumi.Input<number>;
}
