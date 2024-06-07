// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mysql backupEncryptionStatus
 *
 * ## Example Usage
 *
 * ### Enable encryption
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const zones = tencentcloud.Availability.getZonesByProduct({
 *     product: "cdb",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     availabilityZone: zones.then(zones => zones.zones?.[0]?.name),
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.0.0/16",
 *     isMulticast: false,
 * });
 * const securityGroup = new tencentcloud.security.Group("securityGroup", {description: "mysql test"});
 * const exampleInstance = new tencentcloud.mysql.Instance("exampleInstance", {
 *     internetService: 1,
 *     engineVersion: "5.7",
 *     chargeType: "POSTPAID",
 *     rootPassword: "PassWord123",
 *     slaveDeployMode: 0,
 *     availabilityZone: zones.then(zones => zones.zones?.[0]?.name),
 *     slaveSyncMode: 1,
 *     instanceName: "tf-example-mysql",
 *     memSize: 4000,
 *     volumeSize: 200,
 *     vpcId: vpc.id,
 *     subnetId: subnet.id,
 *     intranetPort: 3306,
 *     securityGroups: [securityGroup.id],
 *     tags: {
 *         name: "test",
 *     },
 *     parameters: {
 *         character_set_server: "utf8",
 *         max_connections: "1000",
 *     },
 * });
 * const exampleBackupEncryptionStatus = new tencentcloud.mysql.BackupEncryptionStatus("exampleBackupEncryptionStatus", {
 *     instanceId: exampleInstance.id,
 *     encryptionStatus: "on",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Disable encryption
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.mysql.BackupEncryptionStatus("example", {
 *     instanceId: tencentcloud_mysql_instance.example.id,
 *     encryptionStatus: "off",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * mysql backup_encryption_status can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Mysql/backupEncryptionStatus:BackupEncryptionStatus backup_encryption_status backup_encryption_status_id
 * ```
 */
export class BackupEncryptionStatus extends pulumi.CustomResource {
    /**
     * Get an existing BackupEncryptionStatus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackupEncryptionStatusState, opts?: pulumi.CustomResourceOptions): BackupEncryptionStatus {
        return new BackupEncryptionStatus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mysql/backupEncryptionStatus:BackupEncryptionStatus';

    /**
     * Returns true if the given object is an instance of BackupEncryptionStatus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupEncryptionStatus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupEncryptionStatus.__pulumiType;
    }

    /**
     * Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
     */
    public readonly encryptionStatus!: pulumi.Output<string>;
    /**
     * Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a BackupEncryptionStatus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupEncryptionStatusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackupEncryptionStatusArgs | BackupEncryptionStatusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackupEncryptionStatusState | undefined;
            resourceInputs["encryptionStatus"] = state ? state.encryptionStatus : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as BackupEncryptionStatusArgs | undefined;
            if ((!args || args.encryptionStatus === undefined) && !opts.urn) {
                throw new Error("Missing required property 'encryptionStatus'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["encryptionStatus"] = args ? args.encryptionStatus : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BackupEncryptionStatus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackupEncryptionStatus resources.
 */
export interface BackupEncryptionStatusState {
    /**
     * Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
     */
    encryptionStatus?: pulumi.Input<string>;
    /**
     * Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackupEncryptionStatus resource.
 */
export interface BackupEncryptionStatusArgs {
    /**
     * Whether physical backup encryption is enabled for the instance. Possible values are `on`, `off`.
     */
    encryptionStatus: pulumi.Input<string>;
    /**
     * Instance ID, in the format: cdb-XXXX. Same instance ID as displayed in the ApsaraDB for Console page.
     */
    instanceId: pulumi.Input<string>;
}
