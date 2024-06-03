// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cynosdb clusterDatabases
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const clusterDatabases = new tencentcloud.cynosdb.ClusterDatabases("clusterDatabases", {
 *     characterSet: "utf8",
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     collateRule: "utf8_general_ci",
 *     dbName: "terraform-test",
 *     description: "terraform test",
 *     userHostPrivileges: [{
 *         dbHost: "%",
 *         dbPrivilege: "READ_ONLY",
 *         dbUserName: "root",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * cynosdb cluster_databases can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Cynosdb/clusterDatabases:ClusterDatabases cluster_databases cluster_databases_id
 * ```
 */
export class ClusterDatabases extends pulumi.CustomResource {
    /**
     * Get an existing ClusterDatabases resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterDatabasesState, opts?: pulumi.CustomResourceOptions): ClusterDatabases {
        return new ClusterDatabases(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/clusterDatabases:ClusterDatabases';

    /**
     * Returns true if the given object is an instance of ClusterDatabases.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterDatabases {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterDatabases.__pulumiType;
    }

    /**
     * Character Set Type.
     */
    public readonly characterSet!: pulumi.Output<string>;
    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Sort Rules.
     */
    public readonly collateRule!: pulumi.Output<string>;
    /**
     * Database name.
     */
    public readonly dbName!: pulumi.Output<string>;
    /**
     * Remarks.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Authorize user host permissions.
     */
    public readonly userHostPrivileges!: pulumi.Output<outputs.Cynosdb.ClusterDatabasesUserHostPrivilege[] | undefined>;

    /**
     * Create a ClusterDatabases resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterDatabasesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterDatabasesArgs | ClusterDatabasesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterDatabasesState | undefined;
            resourceInputs["characterSet"] = state ? state.characterSet : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["collateRule"] = state ? state.collateRule : undefined;
            resourceInputs["dbName"] = state ? state.dbName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["userHostPrivileges"] = state ? state.userHostPrivileges : undefined;
        } else {
            const args = argsOrState as ClusterDatabasesArgs | undefined;
            if ((!args || args.characterSet === undefined) && !opts.urn) {
                throw new Error("Missing required property 'characterSet'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.collateRule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collateRule'");
            }
            if ((!args || args.dbName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbName'");
            }
            resourceInputs["characterSet"] = args ? args.characterSet : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["collateRule"] = args ? args.collateRule : undefined;
            resourceInputs["dbName"] = args ? args.dbName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["userHostPrivileges"] = args ? args.userHostPrivileges : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterDatabases.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterDatabases resources.
 */
export interface ClusterDatabasesState {
    /**
     * Character Set Type.
     */
    characterSet?: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Sort Rules.
     */
    collateRule?: pulumi.Input<string>;
    /**
     * Database name.
     */
    dbName?: pulumi.Input<string>;
    /**
     * Remarks.
     */
    description?: pulumi.Input<string>;
    /**
     * Authorize user host permissions.
     */
    userHostPrivileges?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterDatabasesUserHostPrivilege>[]>;
}

/**
 * The set of arguments for constructing a ClusterDatabases resource.
 */
export interface ClusterDatabasesArgs {
    /**
     * Character Set Type.
     */
    characterSet: pulumi.Input<string>;
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Sort Rules.
     */
    collateRule: pulumi.Input<string>;
    /**
     * Database name.
     */
    dbName: pulumi.Input<string>;
    /**
     * Remarks.
     */
    description?: pulumi.Input<string>;
    /**
     * Authorize user host permissions.
     */
    userHostPrivileges?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ClusterDatabasesUserHostPrivilege>[]>;
}
