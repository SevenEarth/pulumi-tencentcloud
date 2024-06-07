// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a lighthouse renewInstance
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const renewInstance = new tencentcloud.lighthouse.RenewInstance("renewInstance", {
 *     autoVoucher: false,
 *     instanceChargePrepaid: {
 *         period: 1,
 *         renewFlag: "NOTIFY_AND_MANUAL_RENEW",
 *     },
 *     instanceId: "lhins-xxxxxxx",
 *     renewDataDisk: true,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class RenewInstance extends pulumi.CustomResource {
    /**
     * Get an existing RenewInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RenewInstanceState, opts?: pulumi.CustomResourceOptions): RenewInstance {
        return new RenewInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Lighthouse/renewInstance:RenewInstance';

    /**
     * Returns true if the given object is an instance of RenewInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RenewInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RenewInstance.__pulumiType;
    }

    /**
     * Whether to automatically deduct vouchers. Valid values:
     * - true: Automatically deduct vouchers.
     * -false:Do not automatically deduct vouchers. Default value: false.
     */
    public readonly autoVoucher!: pulumi.Output<boolean | undefined>;
    /**
     * Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify attributes such as the purchase duration of the Subscription instance and whether to set automatic renewal.
     */
    public readonly instanceChargePrepaid!: pulumi.Output<outputs.Lighthouse.RenewInstanceInstanceChargePrepaid>;
    /**
     * Instance ID.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether to renew the data disk. Valid values:true: Indicates that the renewal instance also renews the data disk attached to it.false: Indicates that the instance will be renewed and the data disk attached to it will not be renewed at the same time.Default value: true.
     */
    public readonly renewDataDisk!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RenewInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RenewInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RenewInstanceArgs | RenewInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RenewInstanceState | undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["instanceChargePrepaid"] = state ? state.instanceChargePrepaid : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["renewDataDisk"] = state ? state.renewDataDisk : undefined;
        } else {
            const args = argsOrState as RenewInstanceArgs | undefined;
            if ((!args || args.instanceChargePrepaid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceChargePrepaid'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["autoVoucher"] = args ? args.autoVoucher : undefined;
            resourceInputs["instanceChargePrepaid"] = args ? args.instanceChargePrepaid : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["renewDataDisk"] = args ? args.renewDataDisk : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RenewInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RenewInstance resources.
 */
export interface RenewInstanceState {
    /**
     * Whether to automatically deduct vouchers. Valid values:
     * - true: Automatically deduct vouchers.
     * -false:Do not automatically deduct vouchers. Default value: false.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify attributes such as the purchase duration of the Subscription instance and whether to set automatic renewal.
     */
    instanceChargePrepaid?: pulumi.Input<inputs.Lighthouse.RenewInstanceInstanceChargePrepaid>;
    /**
     * Instance ID.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether to renew the data disk. Valid values:true: Indicates that the renewal instance also renews the data disk attached to it.false: Indicates that the instance will be renewed and the data disk attached to it will not be renewed at the same time.Default value: true.
     */
    renewDataDisk?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RenewInstance resource.
 */
export interface RenewInstanceArgs {
    /**
     * Whether to automatically deduct vouchers. Valid values:
     * - true: Automatically deduct vouchers.
     * -false:Do not automatically deduct vouchers. Default value: false.
     */
    autoVoucher?: pulumi.Input<boolean>;
    /**
     * Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify attributes such as the purchase duration of the Subscription instance and whether to set automatic renewal.
     */
    instanceChargePrepaid: pulumi.Input<inputs.Lighthouse.RenewInstanceInstanceChargePrepaid>;
    /**
     * Instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Whether to renew the data disk. Valid values:true: Indicates that the renewal instance also renews the data disk attached to it.false: Indicates that the instance will be renewed and the data disk attached to it will not be renewed at the same time.Default value: true.
     */
    renewDataDisk?: pulumi.Input<boolean>;
}
