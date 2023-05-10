// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cbs disk_backup.
 *
 * > **NOTE:** Backup quota must greater than 1.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const diskBackup = new tencentcloud.Cbs.DiskBackup("disk_backup", {
 *     diskBackupName: "xxx",
 *     diskId: "disk-xxx",
 * });
 * ```
 *
 * ## Import
 *
 * cbs disk_backup can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cbs/diskBackup:DiskBackup disk_backup disk_backup_id
 * ```
 */
export class DiskBackup extends pulumi.CustomResource {
    /**
     * Get an existing DiskBackup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskBackupState, opts?: pulumi.CustomResourceOptions): DiskBackup {
        return new DiskBackup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cbs/diskBackup:DiskBackup';

    /**
     * Returns true if the given object is an instance of DiskBackup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiskBackup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiskBackup.__pulumiType;
    }

    /**
     * Backup point name.
     */
    public readonly diskBackupName!: pulumi.Output<string | undefined>;
    /**
     * ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
     */
    public readonly diskId!: pulumi.Output<string>;

    /**
     * Create a DiskBackup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskBackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskBackupArgs | DiskBackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DiskBackupState | undefined;
            resourceInputs["diskBackupName"] = state ? state.diskBackupName : undefined;
            resourceInputs["diskId"] = state ? state.diskId : undefined;
        } else {
            const args = argsOrState as DiskBackupArgs | undefined;
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            resourceInputs["diskBackupName"] = args ? args.diskBackupName : undefined;
            resourceInputs["diskId"] = args ? args.diskId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DiskBackup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiskBackup resources.
 */
export interface DiskBackupState {
    /**
     * Backup point name.
     */
    diskBackupName?: pulumi.Input<string>;
    /**
     * ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
     */
    diskId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DiskBackup resource.
 */
export interface DiskBackupArgs {
    /**
     * Backup point name.
     */
    diskBackupName?: pulumi.Input<string>;
    /**
     * ID of the original cloud disk of the backup point, which can be queried through the DescribeDisks API.
     */
    diskId: pulumi.Input<string>;
}
