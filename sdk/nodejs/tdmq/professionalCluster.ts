// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tdmq professionalCluster
 *
 * ## Example Usage
 * ### single-zone Professional Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const professionalCluster = new tencentcloud.Tdmq.ProfessionalCluster("professional_cluster", {
 *     autoRenewFlag: 1,
 *     clusterName: "single_zone_cluster",
 *     productName: "PULSAR.P1.MINI2",
 *     storageSize: 600,
 *     tags: {
 *         createby: "terrafrom",
 *     },
 *     vpc: {
 *         subnetId: "subnet-xxxx",
 *         vpcId: "vpc-xxxx",
 *     },
 *     zoneIds: [100004],
 * });
 * ```
 * ### Multi-zone Professional Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const professionalCluster = new tencentcloud.Tdmq.ProfessionalCluster("professional_cluster", {
 *     autoRenewFlag: 1,
 *     clusterName: "multi_zone_cluster",
 *     productName: "PULSAR.P1.MINI2",
 *     storageSize: 200,
 *     tags: {
 *         key: "value1",
 *         key2: "value2",
 *     },
 *     vpc: {
 *         subnetId: "subnet-xxxx",
 *         vpcId: "vpc-xxxx",
 *     },
 *     zoneIds: [
 *         330001,
 *         330002,
 *         330003,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * tdmq professional_cluster can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tdmq/professionalCluster:ProfessionalCluster professional_cluster professional_cluster_id
 * ```
 */
export class ProfessionalCluster extends pulumi.CustomResource {
    /**
     * Get an existing ProfessionalCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProfessionalClusterState, opts?: pulumi.CustomResourceOptions): ProfessionalCluster {
        return new ProfessionalCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tdmq/professionalCluster:ProfessionalCluster';

    /**
     * Returns true if the given object is an instance of ProfessionalCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProfessionalCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProfessionalCluster.__pulumiType;
    }

    /**
     * Whether to turn on automatic monthly renewal. `1`: turn on, `0`: turn off.
     */
    public readonly autoRenewFlag!: pulumi.Output<number>;
    /**
     * Whether to automatically select vouchers. `1`: Yes, `0`: No. Default is `0`.
     */
    public readonly autoVoucher!: pulumi.Output<number>;
    /**
     * Name of cluster. It does not support Chinese characters and special characters except dashes and underscores and cannot exceed 64 characters.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * Cluster specification code. Reference[Professional Cluster Specifications](https://cloud.tencent.com/document/product/1179/83705).
     */
    public readonly productName!: pulumi.Output<string>;
    /**
     * Storage specifications. Reference[Professional Cluster Specifications](https://cloud.tencent.com/document/product/1179/83705).
     */
    public readonly storageSize!: pulumi.Output<number>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Purchase duration, value range: 1~50. Default: 1.
     */
    public readonly timeSpan!: pulumi.Output<number>;
    /**
     * Label of VPC network.
     */
    public readonly vpc!: pulumi.Output<outputs.Tdmq.ProfessionalClusterVpc | undefined>;
    /**
     * Multi-AZ deployment select three Availability Zones, like: [200002,200003,200004]. Single availability zone deployment selects an availability zone, like [200002].
     */
    public readonly zoneIds!: pulumi.Output<number[]>;

    /**
     * Create a ProfessionalCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProfessionalClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProfessionalClusterArgs | ProfessionalClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProfessionalClusterState | undefined;
            resourceInputs["autoRenewFlag"] = state ? state.autoRenewFlag : undefined;
            resourceInputs["autoVoucher"] = state ? state.autoVoucher : undefined;
            resourceInputs["clusterName"] = state ? state.clusterName : undefined;
            resourceInputs["productName"] = state ? state.productName : undefined;
            resourceInputs["storageSize"] = state ? state.storageSize : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timeSpan"] = state ? state.timeSpan : undefined;
            resourceInputs["vpc"] = state ? state.vpc : undefined;
            resourceInputs["zoneIds"] = state ? state.zoneIds : undefined;
        } else {
            const args = argsOrState as ProfessionalClusterArgs | undefined;
            if ((!args || args.autoRenewFlag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoRenewFlag'");
            }
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.productName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productName'");
            }
            if ((!args || args.storageSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageSize'");
            }
            if ((!args || args.zoneIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneIds'");
            }
            resourceInputs["autoRenewFlag"] = args ? args.autoRenewFlag : undefined;
            resourceInputs["autoVoucher"] = args ? args.autoVoucher : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["productName"] = args ? args.productName : undefined;
            resourceInputs["storageSize"] = args ? args.storageSize : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeSpan"] = args ? args.timeSpan : undefined;
            resourceInputs["vpc"] = args ? args.vpc : undefined;
            resourceInputs["zoneIds"] = args ? args.zoneIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProfessionalCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProfessionalCluster resources.
 */
export interface ProfessionalClusterState {
    /**
     * Whether to turn on automatic monthly renewal. `1`: turn on, `0`: turn off.
     */
    autoRenewFlag?: pulumi.Input<number>;
    /**
     * Whether to automatically select vouchers. `1`: Yes, `0`: No. Default is `0`.
     */
    autoVoucher?: pulumi.Input<number>;
    /**
     * Name of cluster. It does not support Chinese characters and special characters except dashes and underscores and cannot exceed 64 characters.
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Cluster specification code. Reference[Professional Cluster Specifications](https://cloud.tencent.com/document/product/1179/83705).
     */
    productName?: pulumi.Input<string>;
    /**
     * Storage specifications. Reference[Professional Cluster Specifications](https://cloud.tencent.com/document/product/1179/83705).
     */
    storageSize?: pulumi.Input<number>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Purchase duration, value range: 1~50. Default: 1.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * Label of VPC network.
     */
    vpc?: pulumi.Input<inputs.Tdmq.ProfessionalClusterVpc>;
    /**
     * Multi-AZ deployment select three Availability Zones, like: [200002,200003,200004]. Single availability zone deployment selects an availability zone, like [200002].
     */
    zoneIds?: pulumi.Input<pulumi.Input<number>[]>;
}

/**
 * The set of arguments for constructing a ProfessionalCluster resource.
 */
export interface ProfessionalClusterArgs {
    /**
     * Whether to turn on automatic monthly renewal. `1`: turn on, `0`: turn off.
     */
    autoRenewFlag: pulumi.Input<number>;
    /**
     * Whether to automatically select vouchers. `1`: Yes, `0`: No. Default is `0`.
     */
    autoVoucher?: pulumi.Input<number>;
    /**
     * Name of cluster. It does not support Chinese characters and special characters except dashes and underscores and cannot exceed 64 characters.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Cluster specification code. Reference[Professional Cluster Specifications](https://cloud.tencent.com/document/product/1179/83705).
     */
    productName: pulumi.Input<string>;
    /**
     * Storage specifications. Reference[Professional Cluster Specifications](https://cloud.tencent.com/document/product/1179/83705).
     */
    storageSize: pulumi.Input<number>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Purchase duration, value range: 1~50. Default: 1.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * Label of VPC network.
     */
    vpc?: pulumi.Input<inputs.Tdmq.ProfessionalClusterVpc>;
    /**
     * Multi-AZ deployment select three Availability Zones, like: [200002,200003,200004]. Single availability zone deployment selects an availability zone, like [200002].
     */
    zoneIds: pulumi.Input<pulumi.Input<number>[]>;
}
