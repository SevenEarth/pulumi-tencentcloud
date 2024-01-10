// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dlc storeLocationConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const storeLocationConfig = new tencentcloud.Dlc.StoreLocationConfig("store_location_config", {
 *     enable: 1,
 *     storeLocation: "cosn://bucketname/",
 * });
 * ```
 *
 * ## Import
 *
 * dlc store_location_config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dlc/storeLocationConfig:StoreLocationConfig store_location_config store_location_config_id
 * ```
 */
export class StoreLocationConfig extends pulumi.CustomResource {
    /**
     * Get an existing StoreLocationConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StoreLocationConfigState, opts?: pulumi.CustomResourceOptions): StoreLocationConfig {
        return new StoreLocationConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dlc/storeLocationConfig:StoreLocationConfig';

    /**
     * Returns true if the given object is an instance of StoreLocationConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StoreLocationConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StoreLocationConfig.__pulumiType;
    }

    /**
     * Whether to enable advanced settings: 0-no, 1-yes.
     */
    public readonly enable!: pulumi.Output<number>;
    /**
     * The calculation results are stored in the cos path, such as: cosn://bucketname/.
     */
    public readonly storeLocation!: pulumi.Output<string>;

    /**
     * Create a StoreLocationConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StoreLocationConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StoreLocationConfigArgs | StoreLocationConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StoreLocationConfigState | undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
            resourceInputs["storeLocation"] = state ? state.storeLocation : undefined;
        } else {
            const args = argsOrState as StoreLocationConfigArgs | undefined;
            if ((!args || args.enable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enable'");
            }
            if ((!args || args.storeLocation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storeLocation'");
            }
            resourceInputs["enable"] = args ? args.enable : undefined;
            resourceInputs["storeLocation"] = args ? args.storeLocation : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StoreLocationConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StoreLocationConfig resources.
 */
export interface StoreLocationConfigState {
    /**
     * Whether to enable advanced settings: 0-no, 1-yes.
     */
    enable?: pulumi.Input<number>;
    /**
     * The calculation results are stored in the cos path, such as: cosn://bucketname/.
     */
    storeLocation?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StoreLocationConfig resource.
 */
export interface StoreLocationConfigArgs {
    /**
     * Whether to enable advanced settings: 0-no, 1-yes.
     */
    enable: pulumi.Input<number>;
    /**
     * The calculation results are stored in the cos path, such as: cosn://bucketname/.
     */
    storeLocation: pulumi.Input<string>;
}
