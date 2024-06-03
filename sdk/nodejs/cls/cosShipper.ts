// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cls cos shipper.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const shipper = new tencentcloud.cls.CosShipper("shipper", {
 *     bucket: "preset-scf-bucket-1308919341",
 *     compress: {
 *         format: "lzop",
 *     },
 *     content: {
 *         format: "json",
 *         json: {
 *             enableTag: true,
 *             metaFields: [
 *                 "__FILENAME__",
 *                 "__SOURCE__",
 *                 "__TIMESTAMP__",
 *             ],
 *         },
 *     },
 *     interval: 300,
 *     maxSize: 200,
 *     partition: "/%Y/%m/%d/%H/",
 *     prefix: "ap-guangzhou-fffsasad-1649734752",
 *     shipperName: "ap-guangzhou-fffsasad-1649734752",
 *     topicId: "4d07fba0-b93e-4e0b-9a7f-d58542560bbb",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * cls cos shipper can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Cls/cosShipper:CosShipper shipper 5d1b7b2a-c163-4c48-bb01-9ee00584d761
 * ```
 */
export class CosShipper extends pulumi.CustomResource {
    /**
     * Get an existing CosShipper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CosShipperState, opts?: pulumi.CustomResourceOptions): CosShipper {
        return new CosShipper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cls/cosShipper:CosShipper';

    /**
     * Returns true if the given object is an instance of CosShipper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CosShipper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CosShipper.__pulumiType;
    }

    /**
     * Destination bucket in the shipping rule to be created.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Compression configuration of shipped log.
     */
    public readonly compress!: pulumi.Output<outputs.Cls.CosShipperCompress | undefined>;
    /**
     * Format configuration of shipped log content.
     */
    public readonly content!: pulumi.Output<outputs.Cls.CosShipperContent | undefined>;
    /**
     * Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
     */
    public readonly filterRules!: pulumi.Output<outputs.Cls.CosShipperFilterRule[] | undefined>;
    /**
     * Shipping time interval in seconds. Default value: 300. Value range: 300~900.
     */
    public readonly interval!: pulumi.Output<number | undefined>;
    /**
     * Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
     */
    public readonly maxSize!: pulumi.Output<number | undefined>;
    /**
     * Partition rule of shipped log, which can be represented in strftime time format.
     */
    public readonly partition!: pulumi.Output<string | undefined>;
    /**
     * Prefix of the shipping directory in the shipping rule to be created.
     */
    public readonly prefix!: pulumi.Output<string>;
    /**
     * Shipping rule name.
     */
    public readonly shipperName!: pulumi.Output<string>;
    /**
     * ID of the log topic to which the shipping rule to be created belongs.
     */
    public readonly topicId!: pulumi.Output<string>;

    /**
     * Create a CosShipper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CosShipperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CosShipperArgs | CosShipperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CosShipperState | undefined;
            resourceInputs["bucket"] = state ? state.bucket : undefined;
            resourceInputs["compress"] = state ? state.compress : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["filterRules"] = state ? state.filterRules : undefined;
            resourceInputs["interval"] = state ? state.interval : undefined;
            resourceInputs["maxSize"] = state ? state.maxSize : undefined;
            resourceInputs["partition"] = state ? state.partition : undefined;
            resourceInputs["prefix"] = state ? state.prefix : undefined;
            resourceInputs["shipperName"] = state ? state.shipperName : undefined;
            resourceInputs["topicId"] = state ? state.topicId : undefined;
        } else {
            const args = argsOrState as CosShipperArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.prefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefix'");
            }
            if ((!args || args.shipperName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'shipperName'");
            }
            if ((!args || args.topicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topicId'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["compress"] = args ? args.compress : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["filterRules"] = args ? args.filterRules : undefined;
            resourceInputs["interval"] = args ? args.interval : undefined;
            resourceInputs["maxSize"] = args ? args.maxSize : undefined;
            resourceInputs["partition"] = args ? args.partition : undefined;
            resourceInputs["prefix"] = args ? args.prefix : undefined;
            resourceInputs["shipperName"] = args ? args.shipperName : undefined;
            resourceInputs["topicId"] = args ? args.topicId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CosShipper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CosShipper resources.
 */
export interface CosShipperState {
    /**
     * Destination bucket in the shipping rule to be created.
     */
    bucket?: pulumi.Input<string>;
    /**
     * Compression configuration of shipped log.
     */
    compress?: pulumi.Input<inputs.Cls.CosShipperCompress>;
    /**
     * Format configuration of shipped log content.
     */
    content?: pulumi.Input<inputs.Cls.CosShipperContent>;
    /**
     * Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
     */
    filterRules?: pulumi.Input<pulumi.Input<inputs.Cls.CosShipperFilterRule>[]>;
    /**
     * Shipping time interval in seconds. Default value: 300. Value range: 300~900.
     */
    interval?: pulumi.Input<number>;
    /**
     * Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
     */
    maxSize?: pulumi.Input<number>;
    /**
     * Partition rule of shipped log, which can be represented in strftime time format.
     */
    partition?: pulumi.Input<string>;
    /**
     * Prefix of the shipping directory in the shipping rule to be created.
     */
    prefix?: pulumi.Input<string>;
    /**
     * Shipping rule name.
     */
    shipperName?: pulumi.Input<string>;
    /**
     * ID of the log topic to which the shipping rule to be created belongs.
     */
    topicId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CosShipper resource.
 */
export interface CosShipperArgs {
    /**
     * Destination bucket in the shipping rule to be created.
     */
    bucket: pulumi.Input<string>;
    /**
     * Compression configuration of shipped log.
     */
    compress?: pulumi.Input<inputs.Cls.CosShipperCompress>;
    /**
     * Format configuration of shipped log content.
     */
    content?: pulumi.Input<inputs.Cls.CosShipperContent>;
    /**
     * Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
     */
    filterRules?: pulumi.Input<pulumi.Input<inputs.Cls.CosShipperFilterRule>[]>;
    /**
     * Shipping time interval in seconds. Default value: 300. Value range: 300~900.
     */
    interval?: pulumi.Input<number>;
    /**
     * Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
     */
    maxSize?: pulumi.Input<number>;
    /**
     * Partition rule of shipped log, which can be represented in strftime time format.
     */
    partition?: pulumi.Input<string>;
    /**
     * Prefix of the shipping directory in the shipping rule to be created.
     */
    prefix: pulumi.Input<string>;
    /**
     * Shipping rule name.
     */
    shipperName: pulumi.Input<string>;
    /**
     * ID of the log topic to which the shipping rule to be created belongs.
     */
    topicId: pulumi.Input<string>;
}
