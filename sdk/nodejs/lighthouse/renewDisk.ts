// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a lighthouse renewDisk
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const renewDisk = new tencentcloud.lighthouse.RenewDisk("renewDisk", {
 *     autoVoucher: true,
 *     diskId: "lhdisk-xxxxxx",
 *     renewDiskChargePrepaid: {
 *         period: 1,
 *         renewFlag: "NOTIFY_AND_AUTO_RENEW",
 *         timeUnit: "m",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class RenewDisk extends pulumi.CustomResource {
    /**
     * Get an existing RenewDisk resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RenewDiskState, opts?: pulumi.CustomResourceOptions): RenewDisk {
        return new RenewDisk(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Lighthouse/renewDisk:RenewDisk';

    /**
     * Returns true if the given object is an instance of RenewDisk.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RenewDisk {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RenewDisk.__pulumiType;
    }

    /**
     * Whether to automatically use the voucher. Not used by default.
     */
    public readonly autoVoucher!: pulumi.Output<boolean | undefined>;
    /**
     * List of disk ID.
     */
    public readonly diskId!: pulumi.Output<string>;
    /**
     * Renew cloud hard disk subscription related parameter settings.
     */
    public readonly renewDiskChargePrepaid!: pulumi.Output<outputs.Lighthouse.RenewDiskRenewDiskChargePrepaid>;

    /**
     * Create a RenewDisk resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RenewDiskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RenewDiskArgs | RenewDiskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RenewDiskState | undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["diskId"] = state ? state.diskId : undefined;
            resourceInputs["renewDiskChargePrepaid"] = state ? state.renewDiskChargePrepaid : undefined;
        } else {
            const args = argsOrState as RenewDiskArgs | undefined;
            if ((!args || args.diskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'diskId'");
            }
            if ((!args || args.renewDiskChargePrepaid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'renewDiskChargePrepaid'");
            }
            resourceInputs["autoVoucher"] = args ? args.autoVoucher : undefined;
            resourceInputs["diskId"] = args ? args.diskId : undefined;
            resourceInputs["renewDiskChargePrepaid"] = args ? args.renewDiskChargePrepaid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RenewDisk.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RenewDisk resources.
 */
export interface RenewDiskState {
    /**
     * Whether to automatically use the voucher. Not used by default.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * List of disk ID.
     */
    diskId?: pulumi.Input<string>;
    /**
     * Renew cloud hard disk subscription related parameter settings.
     */
    renewDiskChargePrepaid?: pulumi.Input<inputs.Lighthouse.RenewDiskRenewDiskChargePrepaid>;
}

/**
 * The set of arguments for constructing a RenewDisk resource.
 */
export interface RenewDiskArgs {
    /**
     * Whether to automatically use the voucher. Not used by default.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * List of disk ID.
     */
    diskId: pulumi.Input<string>;
    /**
     * Renew cloud hard disk subscription related parameter settings.
     */
    renewDiskChargePrepaid: pulumi.Input<inputs.Lighthouse.RenewDiskRenewDiskChargePrepaid>;
}
