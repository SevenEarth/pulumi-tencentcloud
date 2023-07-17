// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver configInstanceRoGroup
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const configInstanceRoGroup = new tencentcloud.Sqlserver.ConfigInstanceRoGroup("config_instance_ro_group", {
 *     autoWeight: 0,
 *     instanceId: "mssql-ds1xhnt9",
 *     isOfflineDelay: 1,
 *     minReadOnlyInGroup: 1,
 *     readOnlyGroupId: "mssqlrg-cbya44fb",
 *     readOnlyGroupName: "keep-ro-group-customize",
 *     readOnlyMaxDelayTime: 10,
 *     weightPairs: [{
 *         readOnlyInstanceId: "mssqlro-o6dv2ugx",
 *         readOnlyWeight: 50,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * sqlserver config_instance_ro_group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/configInstanceRoGroup:ConfigInstanceRoGroup config_instance_ro_group config_instance_ro_group_id
 * ```
 */
export class ConfigInstanceRoGroup extends pulumi.CustomResource {
    /**
     * Get an existing ConfigInstanceRoGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigInstanceRoGroupState, opts?: pulumi.CustomResourceOptions): ConfigInstanceRoGroup {
        return new ConfigInstanceRoGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/configInstanceRoGroup:ConfigInstanceRoGroup';

    /**
     * Returns true if the given object is an instance of ConfigInstanceRoGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigInstanceRoGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigInstanceRoGroup.__pulumiType;
    }

    /**
     * 0-user-defined weight (adjusted according to WeightPairs), 1-system automatically assigns weight (WeightPairs is invalid), the default is 0.
     */
    public readonly autoWeight!: pulumi.Output<number | undefined>;
    /**
     * 0-do not rebalance the load, 1-rebalance the load, the default is 0.
     */
    public readonly balanceWeight!: pulumi.Output<number | undefined>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether to enable timeout culling function. 0- Disable the culling function. 1- Enable the culling function.
     */
    public readonly isOfflineDelay!: pulumi.Output<number | undefined>;
    /**
     * After the timeout removal function is enabled, the number of read-only copies retained by the read-only group at least, if this parameter is not filled, it will not be modified.
     */
    public readonly minReadOnlyInGroup!: pulumi.Output<number | undefined>;
    /**
     * Read-only group ID.
     */
    public readonly readOnlyGroupId!: pulumi.Output<string>;
    /**
     * Read-only group name. If this parameter is not specified, it is not modified.
     */
    public readonly readOnlyGroupName!: pulumi.Output<string | undefined>;
    /**
     * After the timeout elimination function is enabled, the timeout threshold used, if this parameter is not filled, it will not be modified.
     */
    public readonly readOnlyMaxDelayTime!: pulumi.Output<number | undefined>;
    /**
     * Read-only group instance weight modification set, if this parameter is not filled, it will not be modified.
     */
    public readonly weightPairs!: pulumi.Output<outputs.Sqlserver.ConfigInstanceRoGroupWeightPair[] | undefined>;

    /**
     * Create a ConfigInstanceRoGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigInstanceRoGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigInstanceRoGroupArgs | ConfigInstanceRoGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigInstanceRoGroupState | undefined;
            resourceInputs["autoWeight"] = state ? state.autoWeight : undefined;
            resourceInputs["balanceWeight"] = state ? state.balanceWeight : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isOfflineDelay"] = state ? state.isOfflineDelay : undefined;
            resourceInputs["minReadOnlyInGroup"] = state ? state.minReadOnlyInGroup : undefined;
            resourceInputs["readOnlyGroupId"] = state ? state.readOnlyGroupId : undefined;
            resourceInputs["readOnlyGroupName"] = state ? state.readOnlyGroupName : undefined;
            resourceInputs["readOnlyMaxDelayTime"] = state ? state.readOnlyMaxDelayTime : undefined;
            resourceInputs["weightPairs"] = state ? state.weightPairs : undefined;
        } else {
            const args = argsOrState as ConfigInstanceRoGroupArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.readOnlyGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'readOnlyGroupId'");
            }
            resourceInputs["autoWeight"] = args ? args.autoWeight : undefined;
            resourceInputs["balanceWeight"] = args ? args.balanceWeight : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isOfflineDelay"] = args ? args.isOfflineDelay : undefined;
            resourceInputs["minReadOnlyInGroup"] = args ? args.minReadOnlyInGroup : undefined;
            resourceInputs["readOnlyGroupId"] = args ? args.readOnlyGroupId : undefined;
            resourceInputs["readOnlyGroupName"] = args ? args.readOnlyGroupName : undefined;
            resourceInputs["readOnlyMaxDelayTime"] = args ? args.readOnlyMaxDelayTime : undefined;
            resourceInputs["weightPairs"] = args ? args.weightPairs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigInstanceRoGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigInstanceRoGroup resources.
 */
export interface ConfigInstanceRoGroupState {
    /**
     * 0-user-defined weight (adjusted according to WeightPairs), 1-system automatically assigns weight (WeightPairs is invalid), the default is 0.
     */
    autoWeight?: pulumi.Input<number>;
    /**
     * 0-do not rebalance the load, 1-rebalance the load, the default is 0.
     */
    balanceWeight?: pulumi.Input<number>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether to enable timeout culling function. 0- Disable the culling function. 1- Enable the culling function.
     */
    isOfflineDelay?: pulumi.Input<number>;
    /**
     * After the timeout removal function is enabled, the number of read-only copies retained by the read-only group at least, if this parameter is not filled, it will not be modified.
     */
    minReadOnlyInGroup?: pulumi.Input<number>;
    /**
     * Read-only group ID.
     */
    readOnlyGroupId?: pulumi.Input<string>;
    /**
     * Read-only group name. If this parameter is not specified, it is not modified.
     */
    readOnlyGroupName?: pulumi.Input<string>;
    /**
     * After the timeout elimination function is enabled, the timeout threshold used, if this parameter is not filled, it will not be modified.
     */
    readOnlyMaxDelayTime?: pulumi.Input<number>;
    /**
     * Read-only group instance weight modification set, if this parameter is not filled, it will not be modified.
     */
    weightPairs?: pulumi.Input<pulumi.Input<inputs.Sqlserver.ConfigInstanceRoGroupWeightPair>[]>;
}

/**
 * The set of arguments for constructing a ConfigInstanceRoGroup resource.
 */
export interface ConfigInstanceRoGroupArgs {
    /**
     * 0-user-defined weight (adjusted according to WeightPairs), 1-system automatically assigns weight (WeightPairs is invalid), the default is 0.
     */
    autoWeight?: pulumi.Input<number>;
    /**
     * 0-do not rebalance the load, 1-rebalance the load, the default is 0.
     */
    balanceWeight?: pulumi.Input<number>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Whether to enable timeout culling function. 0- Disable the culling function. 1- Enable the culling function.
     */
    isOfflineDelay?: pulumi.Input<number>;
    /**
     * After the timeout removal function is enabled, the number of read-only copies retained by the read-only group at least, if this parameter is not filled, it will not be modified.
     */
    minReadOnlyInGroup?: pulumi.Input<number>;
    /**
     * Read-only group ID.
     */
    readOnlyGroupId: pulumi.Input<string>;
    /**
     * Read-only group name. If this parameter is not specified, it is not modified.
     */
    readOnlyGroupName?: pulumi.Input<string>;
    /**
     * After the timeout elimination function is enabled, the timeout threshold used, if this parameter is not filled, it will not be modified.
     */
    readOnlyMaxDelayTime?: pulumi.Input<number>;
    /**
     * Read-only group instance weight modification set, if this parameter is not filled, it will not be modified.
     */
    weightPairs?: pulumi.Input<pulumi.Input<inputs.Sqlserver.ConfigInstanceRoGroupWeightPair>[]>;
}
