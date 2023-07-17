// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a lighthouse disk
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const disk = new tencentcloud.Lighthouse.Disk("disk", {
 *     diskChargePrepaid: {
 *         period: 1,
 *         renewFlag: "NOTIFY_AND_AUTO_RENEW",
 *         timeUnit: "m",
 *     },
 *     diskName: "test",
 *     diskSize: 20,
 *     diskType: "CLOUD_SSD",
 *     zone: "ap-hongkong-2",
 * });
 * ```
 */
export class Disk extends pulumi.CustomResource {
    /**
     * Get an existing Disk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiskState, opts?: pulumi.CustomResourceOptions): Disk {
        return new Disk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Lighthouse/disk:Disk';

    /**
     * Returns true if the given object is an instance of Disk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Disk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Disk.__pulumiType;
    }

    /**
     * Automatically mount and initialize data disks.
     */
    public readonly autoMountConfiguration!: pulumi.Output<outputs.Lighthouse.DiskAutoMountConfiguration | undefined>;
    /**
     * Whether to automatically use the voucher. Not used by default.
     */
    public readonly autoVoucher!: pulumi.Output<boolean>;
    /**
     * Specify the disk backup quota. If not uploaded, the default is no backup quota. Currently, only one disk backup quota is supported.
     */
    public readonly diskBackupQuota!: pulumi.Output<number>;
    /**
     * Disk subscription related parameter settings.
     */
    public readonly diskChargePrepaid!: pulumi.Output<outputs.Lighthouse.DiskDiskChargePrepaid>;
    /**
     * Disk count. Values: [1, 30]. Default: 1.
     */
    public readonly diskCount!: pulumi.Output<number>;
    /**
     * Disk name. Maximum length 60.
     */
    public readonly diskName!: pulumi.Output<string>;
    /**
     * Disk size, unit: GB.
     */
    public readonly diskSize!: pulumi.Output<number>;
    /**
     * Disk type. Value:CLOUD_PREMIUM, CLOUD_SSD.
     */
    public readonly diskType!: pulumi.Output<string>;
    /**
     * Availability zone.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Disk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiskArgs | DiskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DiskState | undefined;
            resourceInputs["autoMountConfiguration"] = state ? state.autoMountConfiguration : undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["diskBackupQuota"] = state ? state.diskBackupQuota : undefined;
            resourceInputs["diskChargePrepaid"] = state ? state.diskChargePrepaid : undefined;
            resourceInputs["diskCount"] = state ? state.diskCount : undefined;
            resourceInputs["diskName"] = state ? state.diskName : undefined;
            resourceInputs["diskSize"] = state ? state.diskSize : undefined;
            resourceInputs["diskType"] = state ? state.diskType : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as DiskArgs | undefined;
            if ((!args || args.diskChargePrepaid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskChargePrepaid'");
            }
            if ((!args || args.diskSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskSize'");
            }
            if ((!args || args.diskType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskType'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["autoMountConfiguration"] = args ? args.autoMountConfiguration : undefined;
            resourceInputs["autoVoucher"] = args ? args.autoVoucher : undefined;
            resourceInputs["diskBackupQuota"] = args ? args.diskBackupQuota : undefined;
            resourceInputs["diskChargePrepaid"] = args ? args.diskChargePrepaid : undefined;
            resourceInputs["diskCount"] = args ? args.diskCount : undefined;
            resourceInputs["diskName"] = args ? args.diskName : undefined;
            resourceInputs["diskSize"] = args ? args.diskSize : undefined;
            resourceInputs["diskType"] = args ? args.diskType : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Disk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Disk resources.
 */
export interface DiskState {
    /**
     * Automatically mount and initialize data disks.
     */
    autoMountConfiguration?: pulumi.Input<inputs.Lighthouse.DiskAutoMountConfiguration>;
    /**
     * Whether to automatically use the voucher. Not used by default.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * Specify the disk backup quota. If not uploaded, the default is no backup quota. Currently, only one disk backup quota is supported.
     */
    diskBackupQuota?: pulumi.Input<number>;
    /**
     * Disk subscription related parameter settings.
     */
    diskChargePrepaid?: pulumi.Input<inputs.Lighthouse.DiskDiskChargePrepaid>;
    /**
     * Disk count. Values: [1, 30]. Default: 1.
     */
    diskCount?: pulumi.Input<number>;
    /**
     * Disk name. Maximum length 60.
     */
    diskName?: pulumi.Input<string>;
    /**
     * Disk size, unit: GB.
     */
    diskSize?: pulumi.Input<number>;
    /**
     * Disk type. Value:CLOUD_PREMIUM, CLOUD_SSD.
     */
    diskType?: pulumi.Input<string>;
    /**
     * Availability zone.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Disk resource.
 */
export interface DiskArgs {
    /**
     * Automatically mount and initialize data disks.
     */
    autoMountConfiguration?: pulumi.Input<inputs.Lighthouse.DiskAutoMountConfiguration>;
    /**
     * Whether to automatically use the voucher. Not used by default.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * Specify the disk backup quota. If not uploaded, the default is no backup quota. Currently, only one disk backup quota is supported.
     */
    diskBackupQuota?: pulumi.Input<number>;
    /**
     * Disk subscription related parameter settings.
     */
    diskChargePrepaid: pulumi.Input<inputs.Lighthouse.DiskDiskChargePrepaid>;
    /**
     * Disk count. Values: [1, 30]. Default: 1.
     */
    diskCount?: pulumi.Input<number>;
    /**
     * Disk name. Maximum length 60.
     */
    diskName?: pulumi.Input<string>;
    /**
     * Disk size, unit: GB.
     */
    diskSize: pulumi.Input<number>;
    /**
     * Disk type. Value:CLOUD_PREMIUM, CLOUD_SSD.
     */
    diskType: pulumi.Input<string>;
    /**
     * Availability zone.
     */
    zone: pulumi.Input<string>;
}
