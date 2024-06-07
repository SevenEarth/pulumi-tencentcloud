// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a vpc bandwidthPackage
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.vpc.BandwidthPackage("example", {
 *     bandwidthPackageName: "tf-example",
 *     chargeType: "TOP5_POSTPAID_BY_MONTH",
 *     networkType: "BGP",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### PrePaid Bandwidth Package
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const bandwidthPackage = new tencentcloud.vpc.BandwidthPackage("bandwidthPackage", {
 *     bandwidthPackageName: "test-001",
 *     chargeType: "FIXED_PREPAID_BY_MONTH",
 *     internetMaxBandwidth: 100,
 *     networkType: "BGP",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 *     timeSpan: 3,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Bandwidth Package With Egress
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.vpc.BandwidthPackage("example", {
 *     bandwidthPackageName: "tf-example",
 *     chargeType: "ENHANCED95_POSTPAID_BY_MONTH",
 *     egress: "center_egress2",
 *     internetMaxBandwidth: 400,
 *     networkType: "SINGLEISP_CMCC",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * vpc bandwidth_package can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Vpc/bandwidthPackage:BandwidthPackage bandwidth_package bandwidthPackage_id
 * ```
 */
export class BandwidthPackage extends pulumi.CustomResource {
    /**
     * Get an existing BandwidthPackage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BandwidthPackageState, opts?: pulumi.CustomResourceOptions): BandwidthPackage {
        return new BandwidthPackage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/bandwidthPackage:BandwidthPackage';

    /**
     * Returns true if the given object is an instance of BandwidthPackage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BandwidthPackage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BandwidthPackage.__pulumiType;
    }

    /**
     * Bandwidth package name.
     */
    public readonly bandwidthPackageName!: pulumi.Output<string | undefined>;
    /**
     * Bandwidth package billing type, default: `TOP5_POSTPAID_BY_MONTH`. Optional value: `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid; `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid; `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction); `BANDWIDTH_POSTPAID_BY_DAY`: bandwidth billed by daily postpaid; `ENHANCED95_POSTPAID_BY_MONTH`: enhanced 95 billed monthly postpaid.
     */
    public readonly chargeType!: pulumi.Output<string | undefined>;
    /**
     * Network egress. It defaults to `centerEgress1`. If you want to try the egress feature, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
     */
    public readonly egress!: pulumi.Output<string>;
    /**
     * Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
     */
    public readonly internetMaxBandwidth!: pulumi.Output<number | undefined>;
    /**
     * Bandwidth packet type, default: `BGP`. Optional value: `BGP`: common BGP shared bandwidth package; `HIGH_QUALITY_BGP`: High Quality BGP Shared Bandwidth Package; `SINGLEISP_CMCC`: CMCC shared bandwidth package; `SINGLEISP_CTCC:`: CTCC shared bandwidth package; `SINGLEISP_CUCC`: CUCC shared bandwidth package.
     */
    public readonly networkType!: pulumi.Output<string | undefined>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The purchase duration of the prepaid monthly bandwidth package, unit: month, value range: 1~60.
     */
    public readonly timeSpan!: pulumi.Output<number | undefined>;

    /**
     * Create a BandwidthPackage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BandwidthPackageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BandwidthPackageArgs | BandwidthPackageState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BandwidthPackageState | undefined;
            resourceInputs["bandwidthPackageName"] = state ? state.bandwidthPackageName : undefined;
            resourceInputs["chargeType"] = state ? state.chargeType : undefined;
            resourceInputs["egress"] = state ? state.egress : undefined;
            resourceInputs["internetMaxBandwidth"] = state ? state.internetMaxBandwidth : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timeSpan"] = state ? state.timeSpan : undefined;
        } else {
            const args = argsOrState as BandwidthPackageArgs | undefined;
            resourceInputs["bandwidthPackageName"] = args ? args.bandwidthPackageName : undefined;
            resourceInputs["chargeType"] = args ? args.chargeType : undefined;
            resourceInputs["egress"] = args ? args.egress : undefined;
            resourceInputs["internetMaxBandwidth"] = args ? args.internetMaxBandwidth : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeSpan"] = args ? args.timeSpan : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BandwidthPackage.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BandwidthPackage resources.
 */
export interface BandwidthPackageState {
    /**
     * Bandwidth package name.
     */
    bandwidthPackageName?: pulumi.Input<string>;
    /**
     * Bandwidth package billing type, default: `TOP5_POSTPAID_BY_MONTH`. Optional value: `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid; `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid; `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction); `BANDWIDTH_POSTPAID_BY_DAY`: bandwidth billed by daily postpaid; `ENHANCED95_POSTPAID_BY_MONTH`: enhanced 95 billed monthly postpaid.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Network egress. It defaults to `centerEgress1`. If you want to try the egress feature, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
     */
    egress?: pulumi.Input<string>;
    /**
     * Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
     */
    internetMaxBandwidth?: pulumi.Input<number>;
    /**
     * Bandwidth packet type, default: `BGP`. Optional value: `BGP`: common BGP shared bandwidth package; `HIGH_QUALITY_BGP`: High Quality BGP Shared Bandwidth Package; `SINGLEISP_CMCC`: CMCC shared bandwidth package; `SINGLEISP_CTCC:`: CTCC shared bandwidth package; `SINGLEISP_CUCC`: CUCC shared bandwidth package.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The purchase duration of the prepaid monthly bandwidth package, unit: month, value range: 1~60.
     */
    timeSpan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BandwidthPackage resource.
 */
export interface BandwidthPackageArgs {
    /**
     * Bandwidth package name.
     */
    bandwidthPackageName?: pulumi.Input<string>;
    /**
     * Bandwidth package billing type, default: `TOP5_POSTPAID_BY_MONTH`. Optional value: `TOP5_POSTPAID_BY_MONTH`: TOP5 billed by monthly postpaid; `PERCENT95_POSTPAID_BY_MONTH`: 95 billed monthly postpaid; `FIXED_PREPAID_BY_MONTH`: Monthly prepaid billing (Type FIXED_PREPAID_BY_MONTH product API capability is under construction); `BANDWIDTH_POSTPAID_BY_DAY`: bandwidth billed by daily postpaid; `ENHANCED95_POSTPAID_BY_MONTH`: enhanced 95 billed monthly postpaid.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * Network egress. It defaults to `centerEgress1`. If you want to try the egress feature, please [submit a ticket](https://console.cloud.tencent.com/workorder/category).
     */
    egress?: pulumi.Input<string>;
    /**
     * Bandwidth packet speed limit size. Unit: Mbps, -1 means no speed limit.
     */
    internetMaxBandwidth?: pulumi.Input<number>;
    /**
     * Bandwidth packet type, default: `BGP`. Optional value: `BGP`: common BGP shared bandwidth package; `HIGH_QUALITY_BGP`: High Quality BGP Shared Bandwidth Package; `SINGLEISP_CMCC`: CMCC shared bandwidth package; `SINGLEISP_CTCC:`: CTCC shared bandwidth package; `SINGLEISP_CUCC`: CUCC shared bandwidth package.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The purchase duration of the prepaid monthly bandwidth package, unit: month, value range: 1~60.
     */
    timeSpan?: pulumi.Input<number>;
}
