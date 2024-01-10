// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mps startFlowOperation
 *
 * ## Example Usage
 * ### Start flow
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const operation = new tencentcloud.mps.StartFlowOperation("operation", {
 *     flowId: tencentcloud_mps_flow.flow_rtp.id,
 *     start: true,
 * });
 * ```
 * ### Stop flow
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const operation = new tencentcloud.mps.StartFlowOperation("operation", {
 *     flowId: tencentcloud_mps_flow.flow_rtp.id,
 *     start: false,
 * });
 * ```
 */
export class StartFlowOperation extends pulumi.CustomResource {
    /**
     * Get an existing StartFlowOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StartFlowOperationState, opts?: pulumi.CustomResourceOptions): StartFlowOperation {
        return new StartFlowOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mps/startFlowOperation:StartFlowOperation';

    /**
     * Returns true if the given object is an instance of StartFlowOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StartFlowOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StartFlowOperation.__pulumiType;
    }

    /**
     * Flow Id.
     */
    public readonly flowId!: pulumi.Output<string>;
    /**
     * `true`: start mps stream link flow; `false`: stop.
     */
    public readonly start!: pulumi.Output<boolean>;

    /**
     * Create a StartFlowOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StartFlowOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StartFlowOperationArgs | StartFlowOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StartFlowOperationState | undefined;
            resourceInputs["flowId"] = state ? state.flowId : undefined;
            resourceInputs["start"] = state ? state.start : undefined;
        } else {
            const args = argsOrState as StartFlowOperationArgs | undefined;
            if ((!args || args.flowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowId'");
            }
            if ((!args || args.start === undefined) && !opts.urn) {
                throw new Error("Missing required property 'start'");
            }
            resourceInputs["flowId"] = args ? args.flowId : undefined;
            resourceInputs["start"] = args ? args.start : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StartFlowOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StartFlowOperation resources.
 */
export interface StartFlowOperationState {
    /**
     * Flow Id.
     */
    flowId?: pulumi.Input<string>;
    /**
     * `true`: start mps stream link flow; `false`: stop.
     */
    start?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a StartFlowOperation resource.
 */
export interface StartFlowOperationArgs {
    /**
     * Flow Id.
     */
    flowId: pulumi.Input<string>;
    /**
     * `true`: start mps stream link flow; `false`: stop.
     */
    start: pulumi.Input<boolean>;
}
