// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create a backup config of redis.
 *
 * ## Example Usage
 * ### Set configuration for automatic backups
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zone = tencentcloud.Redis.getZoneConfig({
 *     typeId: 7,
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     vpcId: vpc.id,
 *     availabilityZone: zone.then(zone => zone.lists?[1]?.zone),
 *     cidrBlock: "10.0.1.0/24",
 * });
 * const fooInstance = new tencentcloud.redis.Instance("fooInstance", {
 *     availabilityZone: zone.then(zone => zone.lists?[1]?.zone),
 *     typeId: zone.then(zone => zone.lists?[1]?.typeId),
 *     password: "test12345789",
 *     memSize: 8192,
 *     redisShardNum: zone.then(zone => zone.lists?[1]?.redisShardNums?[0]),
 *     redisReplicasNum: zone.then(zone => zone.lists?[1]?.redisReplicasNums?[0]),
 *     port: 6379,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 * });
 * const fooBackupConfig = new tencentcloud.redis.BackupConfig("fooBackupConfig", {
 *     redisId: fooInstance.id,
 *     backupTime: "04:00-05:00",
 *     backupPeriods: ["Monday"],
 * });
 * ```
 *
 * ## Import
 *
 * Redis
 *
 * backup config can be imported, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Redis/backupConfig:BackupConfig foo redis-id
 * ```
 */
export class BackupConfig extends pulumi.CustomResource {
    /**
     * Get an existing BackupConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupConfigState, opts?: pulumi.CustomResourceOptions): BackupConfig {
        return new BackupConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Redis/backupConfig:BackupConfig';

    /**
     * Returns true if the given object is an instance of BackupConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupConfig.__pulumiType;
    }

    /**
     * It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
     *
     * @deprecated It has been deprecated from version 1.58.2. It makes no difference to online config at all
     */
    public readonly backupPeriods!: pulumi.Output<string[]>;
    /**
     * Specifys what time the backup action should take place. And the time interval should be one hour.
     */
    public readonly backupTime!: pulumi.Output<string>;
    /**
     * ID of a redis instance to which the policy will be applied.
     */
    public readonly redisId!: pulumi.Output<string>;

    /**
     * Create a BackupConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupConfigArgs | BackupConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupConfigState | undefined;
            resourceInputs["backupPeriods"] = state ? state.backupPeriods : undefined;
            resourceInputs["backupTime"] = state ? state.backupTime : undefined;
            resourceInputs["redisId"] = state ? state.redisId : undefined;
        } else {
            const args = argsOrState as BackupConfigArgs | undefined;
            if ((!args || args.backupTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupTime'");
            }
            if ((!args || args.redisId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'redisId'");
            }
            resourceInputs["backupPeriods"] = args ? args.backupPeriods : undefined;
            resourceInputs["backupTime"] = args ? args.backupTime : undefined;
            resourceInputs["redisId"] = args ? args.redisId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupConfig resources.
 */
export interface BackupConfigState {
    /**
     * It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
     *
     * @deprecated It has been deprecated from version 1.58.2. It makes no difference to online config at all
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifys what time the backup action should take place. And the time interval should be one hour.
     */
    backupTime?: pulumi.Input<string>;
    /**
     * ID of a redis instance to which the policy will be applied.
     */
    redisId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupConfig resource.
 */
export interface BackupConfigArgs {
    /**
     * It has been deprecated from version 1.58.2. It makes no difference to online config at all Specifys which day the backup action should take place. Valid values: `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday` and `Sunday`.
     *
     * @deprecated It has been deprecated from version 1.58.2. It makes no difference to online config at all
     */
    backupPeriods?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifys what time the backup action should take place. And the time interval should be one hour.
     */
    backupTime: pulumi.Input<string>;
    /**
     * ID of a redis instance to which the policy will be applied.
     */
    redisId: pulumi.Input<string>;
}
