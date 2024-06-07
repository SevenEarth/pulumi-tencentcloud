// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dcdb dbSyncModeConfig
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const config = new tencentcloud.dcdb.DbSyncModeConfig("config", {
 *     instanceId: "%s",
 *     syncMode: 2,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * dcdb db_sync_mode_config can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Dcdb/dbSyncModeConfig:DbSyncModeConfig db_sync_mode_config db_sync_mode_config_id
 * ```
 */
export class DbSyncModeConfig extends pulumi.CustomResource {
    /**
     * Get an existing DbSyncModeConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DbSyncModeConfigState, opts?: pulumi.CustomResourceOptions): DbSyncModeConfig {
        return new DbSyncModeConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dcdb/dbSyncModeConfig:DbSyncModeConfig';

    /**
     * Returns true if the given object is an instance of DbSyncModeConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbSyncModeConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbSyncModeConfig.__pulumiType;
    }

    /**
     * ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
     */
    public readonly syncMode!: pulumi.Output<number>;

    /**
     * Create a DbSyncModeConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DbSyncModeConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DbSyncModeConfigArgs | DbSyncModeConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DbSyncModeConfigState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["syncMode"] = state ? state.syncMode : undefined;
        } else {
            const args = argsOrState as DbSyncModeConfigArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.syncMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'syncMode'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["syncMode"] = args ? args.syncMode : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DbSyncModeConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DbSyncModeConfig resources.
 */
export interface DbSyncModeConfigState {
    /**
     * ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
     */
    syncMode?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DbSyncModeConfig resource.
 */
export interface DbSyncModeConfigArgs {
    /**
     * ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).
     */
    syncMode: pulumi.Input<number>;
}
