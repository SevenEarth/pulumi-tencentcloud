// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver startBackupFullMigration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const startBackupFullMigration = new tencentcloud.Sqlserver.StartBackupFullMigration("start_backup_full_migration", {
 *     backupMigrationId: "mssql-backup-migration-kpl74n9l",
 *     instanceId: "mssql-i1z41iwd",
 * });
 * ```
 */
export class StartBackupFullMigration extends pulumi.CustomResource {
    /**
     * Get an existing StartBackupFullMigration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StartBackupFullMigrationState, opts?: pulumi.CustomResourceOptions): StartBackupFullMigration {
        return new StartBackupFullMigration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/startBackupFullMigration:StartBackupFullMigration';

    /**
     * Returns true if the given object is an instance of StartBackupFullMigration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StartBackupFullMigration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StartBackupFullMigration.__pulumiType;
    }

    /**
     * Backup import task ID, returned by the CreateBackupMigration interface.
     */
    public readonly backupMigrationId!: pulumi.Output<string>;
    /**
     * ID of imported target instance.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a StartBackupFullMigration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StartBackupFullMigrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StartBackupFullMigrationArgs | StartBackupFullMigrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StartBackupFullMigrationState | undefined;
            resourceInputs["backupMigrationId"] = state ? state.backupMigrationId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as StartBackupFullMigrationArgs | undefined;
            if ((!args || args.backupMigrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupMigrationId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["backupMigrationId"] = args ? args.backupMigrationId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StartBackupFullMigration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StartBackupFullMigration resources.
 */
export interface StartBackupFullMigrationState {
    /**
     * Backup import task ID, returned by the CreateBackupMigration interface.
     */
    backupMigrationId?: pulumi.Input<string>;
    /**
     * ID of imported target instance.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StartBackupFullMigration resource.
 */
export interface StartBackupFullMigrationArgs {
    /**
     * Backup import task ID, returned by the CreateBackupMigration interface.
     */
    backupMigrationId: pulumi.Input<string>;
    /**
     * ID of imported target instance.
     */
    instanceId: pulumi.Input<string>;
}
