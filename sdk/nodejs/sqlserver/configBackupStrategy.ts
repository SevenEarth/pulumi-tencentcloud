// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a sqlserver configBackupStrategy
 *
 * ## Example Usage
 *
 * Daily backup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new tencentcloud.sqlserver.ConfigBackupStrategy("config", {
 *     instanceId: local.sqlserver_id,
 *     backupType: "daily",
 *     backupTime: 0,
 *     backupDay: 1,
 *     backupModel: "master_no_pkg",
 *     backupCycles: [1],
 *     backupSaveDays: 7,
 *     regularBackupEnable: "disable",
 *     regularBackupSaveDays: 90,
 *     regularBackupStrategy: "months",
 *     regularBackupCounts: 1,
 * });
 * ```
 *
 * Weekly backup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new tencentcloud.sqlserver.ConfigBackupStrategy("config", {
 *     instanceId: local.sqlserver_id,
 *     backupType: "weekly",
 *     backupTime: 0,
 *     backupModel: "master_no_pkg",
 *     backupCycles: [
 *         1,
 *         3,
 *         5,
 *     ],
 *     backupSaveDays: 7,
 *     regularBackupEnable: "disable",
 *     regularBackupSaveDays: 90,
 *     regularBackupStrategy: "months",
 *     regularBackupCounts: 1,
 * });
 * ```
 *
 * Regular backup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new tencentcloud.sqlserver.ConfigBackupStrategy("config", {
 *     instanceId: local.sqlserver_id,
 *     backupTime: 0,
 *     backupModel: "master_no_pkg",
 *     backupCycles: [
 *         1,
 *         3,
 *     ],
 *     backupSaveDays: 7,
 *     regularBackupEnable: "enable",
 *     regularBackupSaveDays: 120,
 *     regularBackupStrategy: "months",
 *     regularBackupCounts: 1,
 *     regularBackupStartTime: `%s`,
 * });
 * ```
 *
 * ## Import
 *
 * sqlserver config_backup_strategy can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Sqlserver/configBackupStrategy:ConfigBackupStrategy config_backup_strategy config_backup_strategy_id
 * ```
 */
export class ConfigBackupStrategy extends pulumi.CustomResource {
    /**
     * Get an existing ConfigBackupStrategy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigBackupStrategyState, opts?: pulumi.CustomResourceOptions): ConfigBackupStrategy {
        return new ConfigBackupStrategy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Sqlserver/configBackupStrategy:ConfigBackupStrategy';

    /**
     * Returns true if the given object is an instance of ConfigBackupStrategy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigBackupStrategy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigBackupStrategy.__pulumiType;
    }

    /**
     * The days of the week on which backup will be performed when `BackupType` is weekly. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
     */
    public readonly backupCycles!: pulumi.Output<number[] | undefined>;
    /**
     * Backup interval in days when the BackupType is daily. The current value can only be 1.
     */
    public readonly backupDay!: pulumi.Output<number | undefined>;
    /**
     * Backup mode. Valid values: masterPkg (archive the backup files of the primary node), masterNoPkg (do not archive the backup files of the primary node), slavePkg (archive the backup files of the replica node), slaveNoPkg (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
     */
    public readonly backupModel!: pulumi.Output<string | undefined>;
    /**
     * Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
     */
    public readonly backupSaveDays!: pulumi.Output<number | undefined>;
    /**
     * Backup time. Value range: an integer from 0 to 23.
     */
    public readonly backupTime!: pulumi.Output<number | undefined>;
    /**
     * Backup type. Valid values: weekly (when length(BackupDay) <=7 && length(BackupDay) >=2), daily (when length(BackupDay)=1). Default value: daily.
     */
    public readonly backupType!: pulumi.Output<string | undefined>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The number of retained archive backups. Default value: 1.
     */
    public readonly regularBackupCounts!: pulumi.Output<number | undefined>;
    /**
     * Archive backup status. Valid values: enable (enabled); disable (disabled). Default value: disable.
     */
    public readonly regularBackupEnable!: pulumi.Output<string | undefined>;
    /**
     * Archive backup retention days. Value range: 90-3650 days. Default value: 365 days.
     */
    public readonly regularBackupSaveDays!: pulumi.Output<number | undefined>;
    /**
     * Archive backup start date in YYYY-MM-DD format, which is the current time by default.
     */
    public readonly regularBackupStartTime!: pulumi.Output<string | undefined>;
    /**
     * Archive backup policy. Valid values: years (yearly); quarters (quarterly); months(monthly); Default value: `months`.
     */
    public readonly regularBackupStrategy!: pulumi.Output<string | undefined>;

    /**
     * Create a ConfigBackupStrategy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigBackupStrategyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigBackupStrategyArgs | ConfigBackupStrategyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigBackupStrategyState | undefined;
            resourceInputs["backupCycles"] = state ? state.backupCycles : undefined;
            resourceInputs["backupDay"] = state ? state.backupDay : undefined;
            resourceInputs["backupModel"] = state ? state.backupModel : undefined;
            resourceInputs["backupSaveDays"] = state ? state.backupSaveDays : undefined;
            resourceInputs["backupTime"] = state ? state.backupTime : undefined;
            resourceInputs["backupType"] = state ? state.backupType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["regularBackupCounts"] = state ? state.regularBackupCounts : undefined;
            resourceInputs["regularBackupEnable"] = state ? state.regularBackupEnable : undefined;
            resourceInputs["regularBackupSaveDays"] = state ? state.regularBackupSaveDays : undefined;
            resourceInputs["regularBackupStartTime"] = state ? state.regularBackupStartTime : undefined;
            resourceInputs["regularBackupStrategy"] = state ? state.regularBackupStrategy : undefined;
        } else {
            const args = argsOrState as ConfigBackupStrategyArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["backupCycles"] = args ? args.backupCycles : undefined;
            resourceInputs["backupDay"] = args ? args.backupDay : undefined;
            resourceInputs["backupModel"] = args ? args.backupModel : undefined;
            resourceInputs["backupSaveDays"] = args ? args.backupSaveDays : undefined;
            resourceInputs["backupTime"] = args ? args.backupTime : undefined;
            resourceInputs["backupType"] = args ? args.backupType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["regularBackupCounts"] = args ? args.regularBackupCounts : undefined;
            resourceInputs["regularBackupEnable"] = args ? args.regularBackupEnable : undefined;
            resourceInputs["regularBackupSaveDays"] = args ? args.regularBackupSaveDays : undefined;
            resourceInputs["regularBackupStartTime"] = args ? args.regularBackupStartTime : undefined;
            resourceInputs["regularBackupStrategy"] = args ? args.regularBackupStrategy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConfigBackupStrategy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigBackupStrategy resources.
 */
export interface ConfigBackupStrategyState {
    /**
     * The days of the week on which backup will be performed when `BackupType` is weekly. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
     */
    backupCycles?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Backup interval in days when the BackupType is daily. The current value can only be 1.
     */
    backupDay?: pulumi.Input<number>;
    /**
     * Backup mode. Valid values: masterPkg (archive the backup files of the primary node), masterNoPkg (do not archive the backup files of the primary node), slavePkg (archive the backup files of the replica node), slaveNoPkg (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
     */
    backupModel?: pulumi.Input<string>;
    /**
     * Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
     */
    backupSaveDays?: pulumi.Input<number>;
    /**
     * Backup time. Value range: an integer from 0 to 23.
     */
    backupTime?: pulumi.Input<number>;
    /**
     * Backup type. Valid values: weekly (when length(BackupDay) <=7 && length(BackupDay) >=2), daily (when length(BackupDay)=1). Default value: daily.
     */
    backupType?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The number of retained archive backups. Default value: 1.
     */
    regularBackupCounts?: pulumi.Input<number>;
    /**
     * Archive backup status. Valid values: enable (enabled); disable (disabled). Default value: disable.
     */
    regularBackupEnable?: pulumi.Input<string>;
    /**
     * Archive backup retention days. Value range: 90-3650 days. Default value: 365 days.
     */
    regularBackupSaveDays?: pulumi.Input<number>;
    /**
     * Archive backup start date in YYYY-MM-DD format, which is the current time by default.
     */
    regularBackupStartTime?: pulumi.Input<string>;
    /**
     * Archive backup policy. Valid values: years (yearly); quarters (quarterly); months(monthly); Default value: `months`.
     */
    regularBackupStrategy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigBackupStrategy resource.
 */
export interface ConfigBackupStrategyArgs {
    /**
     * The days of the week on which backup will be performed when `BackupType` is weekly. If data backup retention period is less than 7 days, the values will be 1-7, indicating that backup will be performed everyday by default; if data backup retention period is greater than or equal to 7 days, the values will be at least any two days, indicating that backup will be performed at least twice in a week by default.
     */
    backupCycles?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * Backup interval in days when the BackupType is daily. The current value can only be 1.
     */
    backupDay?: pulumi.Input<number>;
    /**
     * Backup mode. Valid values: masterPkg (archive the backup files of the primary node), masterNoPkg (do not archive the backup files of the primary node), slavePkg (archive the backup files of the replica node), slaveNoPkg (do not archive the backup files of the replica node). Backup files of the replica node are supported only when Always On disaster recovery is enabled.
     */
    backupModel?: pulumi.Input<string>;
    /**
     * Data (log) backup retention period. Value range: 3-1830 days, default value: 7 days.
     */
    backupSaveDays?: pulumi.Input<number>;
    /**
     * Backup time. Value range: an integer from 0 to 23.
     */
    backupTime?: pulumi.Input<number>;
    /**
     * Backup type. Valid values: weekly (when length(BackupDay) <=7 && length(BackupDay) >=2), daily (when length(BackupDay)=1). Default value: daily.
     */
    backupType?: pulumi.Input<string>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The number of retained archive backups. Default value: 1.
     */
    regularBackupCounts?: pulumi.Input<number>;
    /**
     * Archive backup status. Valid values: enable (enabled); disable (disabled). Default value: disable.
     */
    regularBackupEnable?: pulumi.Input<string>;
    /**
     * Archive backup retention days. Value range: 90-3650 days. Default value: 365 days.
     */
    regularBackupSaveDays?: pulumi.Input<number>;
    /**
     * Archive backup start date in YYYY-MM-DD format, which is the current time by default.
     */
    regularBackupStartTime?: pulumi.Input<string>;
    /**
     * Archive backup policy. Valid values: years (yearly); quarters (quarterly); months(monthly); Default value: `months`.
     */
    regularBackupStrategy?: pulumi.Input<string>;
}