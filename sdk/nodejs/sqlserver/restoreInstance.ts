// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver restoreInstance
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const restoreInstance = new tencentcloud.sqlserver.RestoreInstance("restoreInstance", {
 *     backupId: 3482091273,
 *     instanceId: "mssql-qelbzgwf",
 *     renameRestores: [{
 *         newName: "restore_keep_pubsub_db2",
 *         oldName: "keep_pubsub_db2",
 *     }],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * sqlserver restore_instance can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Sqlserver/restoreInstance:RestoreInstance restore_instance mssql-qelbzgwf#3482091273#keep_pubsub_db2#restore_keep_pubsub_db2
 * ```
 */
export class RestoreInstance extends pulumi.CustomResource {
    /**
     * Get an existing RestoreInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RestoreInstanceState, opts?: pulumi.CustomResourceOptions): RestoreInstance {
        return new RestoreInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/restoreInstance:RestoreInstance';

    /**
     * Returns true if the given object is an instance of RestoreInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestoreInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestoreInstance.__pulumiType;
    }

    /**
     * Backup file ID, which can be obtained through the Id field in the returned value of the DescribeBackups API.
     */
    public readonly backupId!: pulumi.Output<number>;
    /**
     * TDE encryption, `enable` encrypted, `disable` unencrypted.
     */
    public /*out*/ readonly encryptions!: pulumi.Output<outputs.Sqlserver.RestoreInstanceEncryption[]>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Restore the databases listed in ReNameRestoreDatabase and rename them after restoration. If this parameter is left empty, all databases will be restored and renamed in the default format.
     */
    public readonly renameRestores!: pulumi.Output<outputs.Sqlserver.RestoreInstanceRenameRestore[]>;

    /**
     * Create a RestoreInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RestoreInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RestoreInstanceArgs | RestoreInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RestoreInstanceState | undefined;
            resourceInputs["backupId"] = state ? state.backupId : undefined;
            resourceInputs["encryptions"] = state ? state.encryptions : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["renameRestores"] = state ? state.renameRestores : undefined;
        } else {
            const args = argsOrState as RestoreInstanceArgs | undefined;
            if ((!args || args.backupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.renameRestores === undefined) && !opts.urn) {
                throw new Error("Missing required property 'renameRestores'");
            }
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["renameRestores"] = args ? args.renameRestores : undefined;
            resourceInputs["encryptions"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RestoreInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RestoreInstance resources.
 */
export interface RestoreInstanceState {
    /**
     * Backup file ID, which can be obtained through the Id field in the returned value of the DescribeBackups API.
     */
    backupId?: pulumi.Input<number>;
    /**
     * TDE encryption, `enable` encrypted, `disable` unencrypted.
     */
    encryptions?: pulumi.Input<pulumi.Input<inputs.Sqlserver.RestoreInstanceEncryption>[]>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Restore the databases listed in ReNameRestoreDatabase and rename them after restoration. If this parameter is left empty, all databases will be restored and renamed in the default format.
     */
    renameRestores?: pulumi.Input<pulumi.Input<inputs.Sqlserver.RestoreInstanceRenameRestore>[]>;
}

/**
 * The set of arguments for constructing a RestoreInstance resource.
 */
export interface RestoreInstanceArgs {
    /**
     * Backup file ID, which can be obtained through the Id field in the returned value of the DescribeBackups API.
     */
    backupId: pulumi.Input<number>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Restore the databases listed in ReNameRestoreDatabase and rename them after restoration. If this parameter is left empty, all databases will be restored and renamed in the default format.
     */
    renameRestores: pulumi.Input<pulumi.Input<inputs.Sqlserver.RestoreInstanceRenameRestore>[]>;
}
