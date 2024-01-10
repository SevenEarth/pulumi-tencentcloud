// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dlc updateRowFilterOperation
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const updateRowFilterOperation = new tencentcloud.Dlc.UpdateRowFilterOperation("update_row_filter_operation", {
 *     policy: {
 *         catalog: "DataLakeCatalog",
 *         column: "",
 *         database: "test_iac_keep",
 *         function: "",
 *         mode: "SENIOR",
 *         operation: "value!=\"0\"",
 *         policyType: "ROWFILTER",
 *         reAuth: false,
 *         source: "USER",
 *         table: "test_table",
 *         view: "",
 *     },
 *     policyId: 103704,
 * });
 * ```
 */
export class UpdateRowFilterOperation extends pulumi.CustomResource {
    /**
     * Get an existing UpdateRowFilterOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UpdateRowFilterOperationState, opts?: pulumi.CustomResourceOptions): UpdateRowFilterOperation {
        return new UpdateRowFilterOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dlc/updateRowFilterOperation:UpdateRowFilterOperation';

    /**
     * Returns true if the given object is an instance of UpdateRowFilterOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UpdateRowFilterOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UpdateRowFilterOperation.__pulumiType;
    }

    /**
     * New filtering strategy.
     */
    public readonly policy!: pulumi.Output<outputs.Dlc.UpdateRowFilterOperationPolicy>;
    /**
     * The id of the row filtering policy.
     */
    public readonly policyId!: pulumi.Output<number>;

    /**
     * Create a UpdateRowFilterOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UpdateRowFilterOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UpdateRowFilterOperationArgs | UpdateRowFilterOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UpdateRowFilterOperationState | undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
        } else {
            const args = argsOrState as UpdateRowFilterOperationArgs | undefined;
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            resourceInputs["policy"] = args ? args.policy : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UpdateRowFilterOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UpdateRowFilterOperation resources.
 */
export interface UpdateRowFilterOperationState {
    /**
     * New filtering strategy.
     */
    policy?: pulumi.Input<inputs.Dlc.UpdateRowFilterOperationPolicy>;
    /**
     * The id of the row filtering policy.
     */
    policyId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a UpdateRowFilterOperation resource.
 */
export interface UpdateRowFilterOperationArgs {
    /**
     * New filtering strategy.
     */
    policy: pulumi.Input<inputs.Dlc.UpdateRowFilterOperationPolicy>;
    /**
     * The id of the row filtering policy.
     */
    policyId: pulumi.Input<number>;
}
