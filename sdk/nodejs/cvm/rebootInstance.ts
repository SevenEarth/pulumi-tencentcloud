// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cvm rebootInstance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const rebootInstance = new tencentcloud.Cvm.RebootInstance("reboot_instance", {
 *     forceReboot: false,
 *     instanceId: "ins-xxxxx",
 * });
 * ```
 */
export class RebootInstance extends pulumi.CustomResource {
    /**
     * Get an existing RebootInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RebootInstanceState, opts?: pulumi.CustomResourceOptions): RebootInstance {
        return new RebootInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cvm/rebootInstance:RebootInstance';

    /**
     * Returns true if the given object is an instance of RebootInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RebootInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RebootInstance.__pulumiType;
    }

    /**
     * This parameter has been disused. We recommend using StopType instead. Note that ForceReboot and StopType parameters cannot be specified at the same time. Whether to forcibly restart an instance after a normal restart fails. Valid values are `TRUE` and `FALSE`. Default value: FALSE.
     */
    public readonly forceReboot!: pulumi.Output<boolean | undefined>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Shutdown type. Valid values: `SOFT`: soft shutdown; `HARD`: hard shutdown; `SOFT_FIRST`: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails. Default value: SOFT.
     */
    public readonly stopType!: pulumi.Output<string | undefined>;

    /**
     * Create a RebootInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RebootInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RebootInstanceArgs | RebootInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RebootInstanceState | undefined;
            resourceInputs["forceReboot"] = state ? state.forceReboot : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["stopType"] = state ? state.stopType : undefined;
        } else {
            const args = argsOrState as RebootInstanceArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["forceReboot"] = args ? args.forceReboot : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["stopType"] = args ? args.stopType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RebootInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RebootInstance resources.
 */
export interface RebootInstanceState {
    /**
     * This parameter has been disused. We recommend using StopType instead. Note that ForceReboot and StopType parameters cannot be specified at the same time. Whether to forcibly restart an instance after a normal restart fails. Valid values are `TRUE` and `FALSE`. Default value: FALSE.
     */
    forceReboot?: pulumi.Input<boolean>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Shutdown type. Valid values: `SOFT`: soft shutdown; `HARD`: hard shutdown; `SOFT_FIRST`: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails. Default value: SOFT.
     */
    stopType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RebootInstance resource.
 */
export interface RebootInstanceArgs {
    /**
     * This parameter has been disused. We recommend using StopType instead. Note that ForceReboot and StopType parameters cannot be specified at the same time. Whether to forcibly restart an instance after a normal restart fails. Valid values are `TRUE` and `FALSE`. Default value: FALSE.
     */
    forceReboot?: pulumi.Input<boolean>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Shutdown type. Valid values: `SOFT`: soft shutdown; `HARD`: hard shutdown; `SOFT_FIRST`: perform a soft shutdown first, and perform a hard shutdown if the soft shutdown fails. Default value: SOFT.
     */
    stopType?: pulumi.Input<string>;
}
