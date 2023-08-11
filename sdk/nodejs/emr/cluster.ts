// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to create a emr cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-3";
 * const cvm4c8m = tencentcloud.Instance.getTypes({
 *     excludeSoldOut: true,
 *     cpuCoreCount: 4,
 *     memorySize: 8,
 *     filters: [
 *         {
 *             name: "instance-charge-type",
 *             values: ["POSTPAID_BY_HOUR"],
 *         },
 *         {
 *             name: "zone",
 *             values: [availabilityZone],
 *         },
 *     ],
 * });
 * const emrVpc = new tencentcloud.vpc.Instance("emrVpc", {cidrBlock: "10.0.0.0/16"});
 * const emrSubnet = new tencentcloud.subnet.Instance("emrSubnet", {
 *     availabilityZone: availabilityZone,
 *     vpcId: emrVpc.id,
 *     cidrBlock: "10.0.20.0/28",
 *     isMulticast: false,
 * });
 * const emrSg = new tencentcloud.security.Group("emrSg", {
 *     description: "emr sg",
 *     projectId: 0,
 * });
 * const emrCluster = new tencentcloud.emr.Cluster("emrCluster", {
 *     productId: 4,
 *     displayStrategy: "clusterList",
 *     vpcSettings: {
 *         vpc_id: emrVpc.id,
 *         subnet_id: emrSubnet.id,
 *     },
 *     softwares: ["zookeeper-3.6.1"],
 *     supportHa: 0,
 *     instanceName: "emr-cluster-test",
 *     resourceSpec: {
 *         masterResourceSpec: {
 *             memSize: 8192,
 *             cpu: 4,
 *             diskSize: 100,
 *             diskType: "CLOUD_PREMIUM",
 *             spec: cvm4c8m.then(cvm4c8m => `CVM.${cvm4c8m.instanceTypes?[0]?.family}`),
 *             storageType: 5,
 *             rootSize: 50,
 *         },
 *         coreResourceSpec: {
 *             memSize: 8192,
 *             cpu: 4,
 *             diskSize: 100,
 *             diskType: "CLOUD_PREMIUM",
 *             spec: cvm4c8m.then(cvm4c8m => `CVM.${cvm4c8m.instanceTypes?[0]?.family}`),
 *             storageType: 5,
 *             rootSize: 50,
 *         },
 *         masterCount: 1,
 *         coreCount: 2,
 *     },
 *     loginSettings: {
 *         password: "Tencent@cloud123",
 *     },
 *     timeSpan: 3600,
 *     timeUnit: "s",
 *     payMode: 0,
 *     placement: {
 *         zone: availabilityZone,
 *         project_id: 0,
 *     },
 *     sgId: emrSg.id,
 * });
 * ```
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Emr/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * Display strategy of EMR instance.
     */
    public readonly displayStrategy!: pulumi.Output<string>;
    /**
     * Access the external file system.
     */
    public readonly extendFsField!: pulumi.Output<string | undefined>;
    /**
     * Created EMR instance id.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Instance login settings.
     */
    public readonly loginSettings!: pulumi.Output<{[key: string]: any}>;
    /**
     * Whether to enable the cluster Master node public network. Value range:
     * - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
     * - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
     * By default, the cluster Master node internet is enabled.
     */
    public readonly needMasterWan!: pulumi.Output<string | undefined>;
    /**
     * The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
     */
    public readonly payMode!: pulumi.Output<number>;
    /**
     * The location of the instance.
     */
    public readonly placement!: pulumi.Output<{[key: string]: any}>;
    /**
     * Product ID. Different products ID represents different EMR product versions. Value range:
     * - 16: represents EMR-V2.3.0
     * - 20: indicates EMR-V2.5.0
     * - 25: represents EMR-V3.1.0
     * - 27: represents KAFKA-V1.0.0
     * - 30: indicates EMR-V2.6.0
     * - 33: represents EMR-V3.2.1
     * - 34: stands for EMR-V3.3.0
     * - 36: represents STARROCKS-V1.0.0
     * - 37: indicates EMR-V3.4.0
     * - 38: represents EMR-V2.7.0
     * - 39: stands for STARROCKS-V1.1.0
     * - 41: represents DRUID-V1.1.0.
     */
    public readonly productId!: pulumi.Output<number>;
    /**
     * Resource specification of EMR instance.
     */
    public readonly resourceSpec!: pulumi.Output<outputs.Emr.ClusterResourceSpec | undefined>;
    /**
     * The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
     */
    public readonly sgId!: pulumi.Output<string | undefined>;
    /**
     * The softwares of a EMR instance.
     */
    public readonly softwares!: pulumi.Output<string[]>;
    /**
     * The flag whether the instance support high availability.(0=>not support, 1=>support).
     */
    public readonly supportHa!: pulumi.Output<number>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;
    /**
     * The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
     * When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
     */
    public readonly timeSpan!: pulumi.Output<number>;
    /**
     * The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
     */
    public readonly timeUnit!: pulumi.Output<string>;
    /**
     * The private net config of EMR instance.
     */
    public readonly vpcSettings!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterState | undefined;
            resourceInputs["displayStrategy"] = state ? state.displayStrategy : undefined;
            resourceInputs["extendFsField"] = state ? state.extendFsField : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["loginSettings"] = state ? state.loginSettings : undefined;
            resourceInputs["needMasterWan"] = state ? state.needMasterWan : undefined;
            resourceInputs["payMode"] = state ? state.payMode : undefined;
            resourceInputs["placement"] = state ? state.placement : undefined;
            resourceInputs["productId"] = state ? state.productId : undefined;
            resourceInputs["resourceSpec"] = state ? state.resourceSpec : undefined;
            resourceInputs["sgId"] = state ? state.sgId : undefined;
            resourceInputs["softwares"] = state ? state.softwares : undefined;
            resourceInputs["supportHa"] = state ? state.supportHa : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["timeSpan"] = state ? state.timeSpan : undefined;
            resourceInputs["timeUnit"] = state ? state.timeUnit : undefined;
            resourceInputs["vpcSettings"] = state ? state.vpcSettings : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            if ((!args || args.displayStrategy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayStrategy'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.loginSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'loginSettings'");
            }
            if ((!args || args.payMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'payMode'");
            }
            if ((!args || args.placement === undefined) && !opts.urn) {
                throw new Error("Missing required property 'placement'");
            }
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            if ((!args || args.softwares === undefined) && !opts.urn) {
                throw new Error("Missing required property 'softwares'");
            }
            if ((!args || args.supportHa === undefined) && !opts.urn) {
                throw new Error("Missing required property 'supportHa'");
            }
            if ((!args || args.timeSpan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeSpan'");
            }
            if ((!args || args.timeUnit === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeUnit'");
            }
            if ((!args || args.vpcSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcSettings'");
            }
            resourceInputs["displayStrategy"] = args ? args.displayStrategy : undefined;
            resourceInputs["extendFsField"] = args ? args.extendFsField : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["loginSettings"] = args ? args.loginSettings : undefined;
            resourceInputs["needMasterWan"] = args ? args.needMasterWan : undefined;
            resourceInputs["payMode"] = args ? args.payMode : undefined;
            resourceInputs["placement"] = args ? args.placement : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["resourceSpec"] = args ? args.resourceSpec : undefined;
            resourceInputs["sgId"] = args ? args.sgId : undefined;
            resourceInputs["softwares"] = args ? args.softwares : undefined;
            resourceInputs["supportHa"] = args ? args.supportHa : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["timeSpan"] = args ? args.timeSpan : undefined;
            resourceInputs["timeUnit"] = args ? args.timeUnit : undefined;
            resourceInputs["vpcSettings"] = args ? args.vpcSettings : undefined;
            resourceInputs["instanceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Cluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * Display strategy of EMR instance.
     */
    displayStrategy?: pulumi.Input<string>;
    /**
     * Access the external file system.
     */
    extendFsField?: pulumi.Input<string>;
    /**
     * Created EMR instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Instance login settings.
     */
    loginSettings?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether to enable the cluster Master node public network. Value range:
     * - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
     * - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
     * By default, the cluster Master node internet is enabled.
     */
    needMasterWan?: pulumi.Input<string>;
    /**
     * The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
     */
    payMode?: pulumi.Input<number>;
    /**
     * The location of the instance.
     */
    placement?: pulumi.Input<{[key: string]: any}>;
    /**
     * Product ID. Different products ID represents different EMR product versions. Value range:
     * - 16: represents EMR-V2.3.0
     * - 20: indicates EMR-V2.5.0
     * - 25: represents EMR-V3.1.0
     * - 27: represents KAFKA-V1.0.0
     * - 30: indicates EMR-V2.6.0
     * - 33: represents EMR-V3.2.1
     * - 34: stands for EMR-V3.3.0
     * - 36: represents STARROCKS-V1.0.0
     * - 37: indicates EMR-V3.4.0
     * - 38: represents EMR-V2.7.0
     * - 39: stands for STARROCKS-V1.1.0
     * - 41: represents DRUID-V1.1.0.
     */
    productId?: pulumi.Input<number>;
    /**
     * Resource specification of EMR instance.
     */
    resourceSpec?: pulumi.Input<inputs.Emr.ClusterResourceSpec>;
    /**
     * The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
     */
    sgId?: pulumi.Input<string>;
    /**
     * The softwares of a EMR instance.
     */
    softwares?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The flag whether the instance support high availability.(0=>not support, 1=>support).
     */
    supportHa?: pulumi.Input<number>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
     * When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
     */
    timeUnit?: pulumi.Input<string>;
    /**
     * The private net config of EMR instance.
     */
    vpcSettings?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * Display strategy of EMR instance.
     */
    displayStrategy: pulumi.Input<string>;
    /**
     * Access the external file system.
     */
    extendFsField?: pulumi.Input<string>;
    /**
     * Name of the instance, which can contain 6 to 36 English letters, Chinese characters, digits, dashes(-), or underscores(_).
     */
    instanceName: pulumi.Input<string>;
    /**
     * Instance login settings.
     */
    loginSettings: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether to enable the cluster Master node public network. Value range:
     * - NEED_MASTER_WAN: Indicates that the cluster Master node public network is enabled.
     * - NOT_NEED_MASTER_WAN: Indicates that it is not turned on.
     * By default, the cluster Master node internet is enabled.
     */
    needMasterWan?: pulumi.Input<string>;
    /**
     * The pay mode of instance. 0 represent POSTPAID_BY_HOUR, 1 represent PREPAID.
     */
    payMode: pulumi.Input<number>;
    /**
     * The location of the instance.
     */
    placement: pulumi.Input<{[key: string]: any}>;
    /**
     * Product ID. Different products ID represents different EMR product versions. Value range:
     * - 16: represents EMR-V2.3.0
     * - 20: indicates EMR-V2.5.0
     * - 25: represents EMR-V3.1.0
     * - 27: represents KAFKA-V1.0.0
     * - 30: indicates EMR-V2.6.0
     * - 33: represents EMR-V3.2.1
     * - 34: stands for EMR-V3.3.0
     * - 36: represents STARROCKS-V1.0.0
     * - 37: indicates EMR-V3.4.0
     * - 38: represents EMR-V2.7.0
     * - 39: stands for STARROCKS-V1.1.0
     * - 41: represents DRUID-V1.1.0.
     */
    productId: pulumi.Input<number>;
    /**
     * Resource specification of EMR instance.
     */
    resourceSpec?: pulumi.Input<inputs.Emr.ClusterResourceSpec>;
    /**
     * The ID of the security group to which the instance belongs, in the form of sg-xxxxxxxx.
     */
    sgId?: pulumi.Input<string>;
    /**
     * The softwares of a EMR instance.
     */
    softwares: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The flag whether the instance support high availability.(0=>not support, 1=>support).
     */
    supportHa: pulumi.Input<number>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The length of time the instance was purchased. Use with TimeUnit.When TimeUnit is s, the parameter can only be filled in at 3600, representing a metered instance.
     * When TimeUnit is m, the number filled in by this parameter indicates the length of purchase of the monthly instance of the package year, such as 1 for one month of purchase.
     */
    timeSpan: pulumi.Input<number>;
    /**
     * The unit of time in which the instance was purchased. When PayMode is 0, TimeUnit can only take values of s(second). When PayMode is 1, TimeUnit can only take the value m(month).
     */
    timeUnit: pulumi.Input<string>;
    /**
     * The private net config of EMR instance.
     */
    vpcSettings: pulumi.Input<{[key: string]: any}>;
}
