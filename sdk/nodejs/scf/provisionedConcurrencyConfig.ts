// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a scf provisionedConcurrencyConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const provisionedConcurrencyConfig = new tencentcloud.Scf.ProvisionedConcurrencyConfig("provisioned_concurrency_config", {
 *     functionName: "keep-1676351130",
 *     maxCapacity: 2,
 *     minCapacity: 1,
 *     namespace: "default",
 *     provisionedType: "Default",
 *     qualifier: "2",
 *     trackingTarget: 0.5,
 *     triggerActions: [{
 *         provisionedType: "Default",
 *         triggerCronConfig: "29 45 12 29 05 * 2023",
 *         triggerName: "test",
 *         triggerProvisionedConcurrencyNum: 2,
 *     }],
 *     versionProvisionedConcurrencyNum: 2,
 * });
 * ```
 */
export class ProvisionedConcurrencyConfig extends pulumi.CustomResource {
    /**
     * Get an existing ProvisionedConcurrencyConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProvisionedConcurrencyConfigState, opts?: pulumi.CustomResourceOptions): ProvisionedConcurrencyConfig {
        return new ProvisionedConcurrencyConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Scf/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig';

    /**
     * Returns true if the given object is an instance of ProvisionedConcurrencyConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProvisionedConcurrencyConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProvisionedConcurrencyConfig.__pulumiType;
    }

    /**
     * Name of the function for which to set the provisioned concurrency.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * The maximum number of instances.
     */
    public readonly maxCapacity!: pulumi.Output<number | undefined>;
    /**
     * The minimum number of instances. It can not be smaller than 1.
     */
    public readonly minCapacity!: pulumi.Output<number | undefined>;
    /**
     * Function namespace. Default value: default.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
     */
    public readonly provisionedType!: pulumi.Output<string | undefined>;
    /**
     * Function version number. Note: the $LATEST version does not support provisioned concurrency.
     */
    public readonly qualifier!: pulumi.Output<string>;
    /**
     * The target concurrency utilization. Range: (0,1) (two decimal places).
     */
    public readonly trackingTarget!: pulumi.Output<number | undefined>;
    /**
     * Scheduled provisioned concurrency scaling action.
     */
    public readonly triggerActions!: pulumi.Output<outputs.Scf.ProvisionedConcurrencyConfigTriggerAction[] | undefined>;
    /**
     * Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
     */
    public readonly versionProvisionedConcurrencyNum!: pulumi.Output<number>;

    /**
     * Create a ProvisionedConcurrencyConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProvisionedConcurrencyConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProvisionedConcurrencyConfigArgs | ProvisionedConcurrencyConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProvisionedConcurrencyConfigState | undefined;
            resourceInputs["functionName"] = state ? state.functionName : undefined;
            resourceInputs["maxCapacity"] = state ? state.maxCapacity : undefined;
            resourceInputs["minCapacity"] = state ? state.minCapacity : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["provisionedType"] = state ? state.provisionedType : undefined;
            resourceInputs["qualifier"] = state ? state.qualifier : undefined;
            resourceInputs["trackingTarget"] = state ? state.trackingTarget : undefined;
            resourceInputs["triggerActions"] = state ? state.triggerActions : undefined;
            resourceInputs["versionProvisionedConcurrencyNum"] = state ? state.versionProvisionedConcurrencyNum : undefined;
        } else {
            const args = argsOrState as ProvisionedConcurrencyConfigArgs | undefined;
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            if ((!args || args.qualifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'qualifier'");
            }
            if ((!args || args.versionProvisionedConcurrencyNum === undefined) && !opts.urn) {
                throw new Error("Missing required property 'versionProvisionedConcurrencyNum'");
            }
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["maxCapacity"] = args ? args.maxCapacity : undefined;
            resourceInputs["minCapacity"] = args ? args.minCapacity : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["provisionedType"] = args ? args.provisionedType : undefined;
            resourceInputs["qualifier"] = args ? args.qualifier : undefined;
            resourceInputs["trackingTarget"] = args ? args.trackingTarget : undefined;
            resourceInputs["triggerActions"] = args ? args.triggerActions : undefined;
            resourceInputs["versionProvisionedConcurrencyNum"] = args ? args.versionProvisionedConcurrencyNum : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProvisionedConcurrencyConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProvisionedConcurrencyConfig resources.
 */
export interface ProvisionedConcurrencyConfigState {
    /**
     * Name of the function for which to set the provisioned concurrency.
     */
    functionName?: pulumi.Input<string>;
    /**
     * The maximum number of instances.
     */
    maxCapacity?: pulumi.Input<number>;
    /**
     * The minimum number of instances. It can not be smaller than 1.
     */
    minCapacity?: pulumi.Input<number>;
    /**
     * Function namespace. Default value: default.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
     */
    provisionedType?: pulumi.Input<string>;
    /**
     * Function version number. Note: the $LATEST version does not support provisioned concurrency.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * The target concurrency utilization. Range: (0,1) (two decimal places).
     */
    trackingTarget?: pulumi.Input<number>;
    /**
     * Scheduled provisioned concurrency scaling action.
     */
    triggerActions?: pulumi.Input<pulumi.Input<inputs.Scf.ProvisionedConcurrencyConfigTriggerAction>[]>;
    /**
     * Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
     */
    versionProvisionedConcurrencyNum?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ProvisionedConcurrencyConfig resource.
 */
export interface ProvisionedConcurrencyConfigArgs {
    /**
     * Name of the function for which to set the provisioned concurrency.
     */
    functionName: pulumi.Input<string>;
    /**
     * The maximum number of instances.
     */
    maxCapacity?: pulumi.Input<number>;
    /**
     * The minimum number of instances. It can not be smaller than 1.
     */
    minCapacity?: pulumi.Input<number>;
    /**
     * Function namespace. Default value: default.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Specifies the provisioned concurrency type. Default: Static provisioned concurrency. ConcurrencyUtilizationTracking: Scales the concurrency automatically according to the concurrency utilization. If ConcurrencyUtilizationTracking is passed in, TrackingTarget, MinCapacity and MaxCapacity are required, and VersionProvisionedConcurrencyNum must be 0.
     */
    provisionedType?: pulumi.Input<string>;
    /**
     * Function version number. Note: the $LATEST version does not support provisioned concurrency.
     */
    qualifier: pulumi.Input<string>;
    /**
     * The target concurrency utilization. Range: (0,1) (two decimal places).
     */
    trackingTarget?: pulumi.Input<number>;
    /**
     * Scheduled provisioned concurrency scaling action.
     */
    triggerActions?: pulumi.Input<pulumi.Input<inputs.Scf.ProvisionedConcurrencyConfigTriggerAction>[]>;
    /**
     * Provisioned concurrency amount. Note: there is an upper limit for the sum of provisioned concurrency amounts of all versions, which currently is the function&amp;#39;s maximum concurrency quota minus 100.
     */
    versionProvisionedConcurrencyNum: pulumi.Input<number>;
}