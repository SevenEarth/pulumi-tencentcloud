// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dnspod snapshotConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const snapshotConfig = new tencentcloud.Dnspod.SnapshotConfig("snapshot_config", {
 *     domain: "dnspod.cn",
 *     period: "hourly",
 * });
 * ```
 *
 * ## Import
 *
 * dnspod snapshot_config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dnspod/snapshotConfig:SnapshotConfig snapshot_config domain
 * ```
 */
export class SnapshotConfig extends pulumi.CustomResource {
    /**
     * Get an existing SnapshotConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotConfigState, opts?: pulumi.CustomResourceOptions): SnapshotConfig {
        return new SnapshotConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dnspod/snapshotConfig:SnapshotConfig';

    /**
     * Returns true if the given object is an instance of SnapshotConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SnapshotConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SnapshotConfig.__pulumiType;
    }

    /**
     * Domain name.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Backup interval: empty string - no backup, halfHour - every half hour, hourly - every hour, daily - every day, monthly - every month.
     */
    public readonly period!: pulumi.Output<string>;

    /**
     * Create a SnapshotConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotConfigArgs | SnapshotConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SnapshotConfigState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
        } else {
            const args = argsOrState as SnapshotConfigArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.period === undefined) && !opts.urn) {
                throw new Error("Missing required property 'period'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SnapshotConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SnapshotConfig resources.
 */
export interface SnapshotConfigState {
    /**
     * Domain name.
     */
    domain?: pulumi.Input<string>;
    /**
     * Backup interval: empty string - no backup, halfHour - every half hour, hourly - every hour, daily - every day, monthly - every month.
     */
    period?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SnapshotConfig resource.
 */
export interface SnapshotConfigArgs {
    /**
     * Domain name.
     */
    domain: pulumi.Input<string>;
    /**
     * Backup interval: empty string - no backup, halfHour - every half hour, hourly - every hour, daily - every day, monthly - every month.
     */
    period: pulumi.Input<string>;
}
