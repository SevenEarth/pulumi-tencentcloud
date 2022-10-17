// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to create an auto scaling group for kubernetes cluster.
 *
 * > **NOTE:**  We recommend the usage of one cluster with essential worker config + node pool to manage cluster and nodes. Its a more flexible way than manage worker config with tencentcloud_kubernetes_cluster, tencentcloud.Kubernetes.ScaleWorker or exist node management of `tencentcloudKubernetesAttachment`. Cause some unchangeable parameters of `workerConfig` may cause the whole cluster resource `force new`.
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
 * const clusterCidr = config.get("clusterCidr") || "172.31.0.0/16";
 * const vpc = tencentcloud.Vpc.getSubnets({
 *     isDefault: true,
 *     availabilityZone: availabilityZone,
 * });
 * const defaultInstanceType = config.get("defaultInstanceType") || "S1.SMALL1";
 * //this is the cluster with empty worker config
 * const managedCluster = new tencentcloud.kubernetes.Cluster("managedCluster", {
 *     vpcId: vpc.then(vpc => vpc.instanceLists?[0]?.vpcId),
 *     clusterCidr: clusterCidr,
 *     clusterMaxPodNum: 32,
 *     clusterName: "tf-tke-unit-test",
 *     clusterDesc: "test cluster desc",
 *     clusterMaxServiceNum: 32,
 *     clusterVersion: "1.18.4",
 *     clusterDeployType: "MANAGED_CLUSTER",
 * });
 * //this is one example of managing node using node pool
 * const mynodepool = new tencentcloud.kubernetes.NodePool("mynodepool", {
 *     clusterId: managedCluster.id,
 *     maxSize: 6,
 *     minSize: 1,
 *     vpcId: vpc.then(vpc => vpc.instanceLists?[0]?.vpcId),
 *     subnetIds: [vpc.then(vpc => vpc.instanceLists?[0]?.subnetId)],
 *     retryPolicy: "INCREMENTAL_INTERVALS",
 *     desiredCapacity: 4,
 *     enableAutoScale: true,
 *     multiZoneSubnetPolicy: "EQUALITY",
 *     autoScalingConfig: {
 *         instanceType: defaultInstanceType,
 *         systemDiskType: "CLOUD_PREMIUM",
 *         systemDiskSize: 50,
 *         securityGroupIds: ["sg-24vswocp"],
 *         dataDisks: [{
 *             diskType: "CLOUD_PREMIUM",
 *             diskSize: 50,
 *         }],
 *         internetChargeType: "TRAFFIC_POSTPAID_BY_HOUR",
 *         internetMaxBandwidthOut: 10,
 *         publicIpAssigned: true,
 *         password: "test123#",
 *         enhancedSecurityService: false,
 *         enhancedMonitorService: false,
 *     },
 *     labels: {
 *         test1: "test1",
 *         test2: "test2",
 *     },
 *     taints: [
 *         {
 *             key: "test_taint",
 *             value: "taint_value",
 *             effect: "PreferNoSchedule",
 *         },
 *         {
 *             key: "test_taint2",
 *             value: "taint_value2",
 *             effect: "PreferNoSchedule",
 *         },
 *     ],
 *     nodeConfig: {
 *         extraArgs: ["root-dir=/var/lib/kubelet"],
 *     },
 * });
 * ```
 *
 * Using Spot CVM Instance
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const mynodepool = new tencentcloud.kubernetes.NodePool("mynodepool", {
 *     clusterId: tencentcloud_kubernetes_cluster.managed_cluster.id,
 *     maxSize: 6,
 *     minSize: 1,
 *     vpcId: data.tencentcloud_vpc_subnets.vpc.instance_list[0].vpc_id,
 *     subnetIds: [data.tencentcloud_vpc_subnets.vpc.instance_list[0].subnet_id],
 *     retryPolicy: "INCREMENTAL_INTERVALS",
 *     desiredCapacity: 4,
 *     enableAutoScale: true,
 *     multiZoneSubnetPolicy: "EQUALITY",
 *     autoScalingConfig: {
 *         instanceType: _var.default_instance_type,
 *         systemDiskType: "CLOUD_PREMIUM",
 *         systemDiskSize: 50,
 *         securityGroupIds: ["sg-24vswocp"],
 *         instanceChargeType: "SPOTPAID",
 *         spotInstanceType: "one-time",
 *         spotMaxPrice: "1000",
 *         dataDisks: [{
 *             diskType: "CLOUD_PREMIUM",
 *             diskSize: 50,
 *         }],
 *         internetChargeType: "TRAFFIC_POSTPAID_BY_HOUR",
 *         internetMaxBandwidthOut: 10,
 *         publicIpAssigned: true,
 *         password: "test123#",
 *         enhancedSecurityService: false,
 *         enhancedMonitorService: false,
 *     },
 *     labels: {
 *         test1: "test1",
 *         test2: "test2",
 *     },
 * });
 * ```
 */
export class NodePool extends pulumi.CustomResource {
    /**
     * Get an existing NodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodePoolState, opts?: pulumi.CustomResourceOptions): NodePool {
        return new NodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Kubernetes/nodePool:NodePool';

    /**
     * Returns true if the given object is an instance of NodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NodePool.__pulumiType;
    }

    /**
     * Auto scaling config parameters.
     */
    public readonly autoScalingConfig!: pulumi.Output<outputs.Kubernetes.NodePoolAutoScalingConfig>;
    /**
     * The auto scaling group ID.
     */
    public /*out*/ readonly autoScalingGroupId!: pulumi.Output<string>;
    /**
     * The total of autoscaling added node.
     */
    public /*out*/ readonly autoscalingAddedTotal!: pulumi.Output<number>;
    /**
     * ID of the cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Seconds of scaling group cool down. Default value is `300`.
     */
    public readonly defaultCooldown!: pulumi.Output<number>;
    /**
     * Indicate to keep the CVM instance when delete the node pool. Default is `true`.
     */
    public readonly deleteKeepInstance!: pulumi.Output<boolean | undefined>;
    /**
     * Desired capacity ot the node. If `enableAutoScale` is set `true`, this will be a computed parameter.
     */
    public readonly desiredCapacity!: pulumi.Output<number>;
    /**
     * Indicate whether to enable auto scaling or not.
     */
    public readonly enableAutoScale!: pulumi.Output<boolean | undefined>;
    /**
     * Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The launch config ID.
     */
    public /*out*/ readonly launchConfigId!: pulumi.Output<string>;
    /**
     * The total of manually added node.
     */
    public /*out*/ readonly manuallyAddedTotal!: pulumi.Output<number>;
    /**
     * Maximum number of node.
     */
    public readonly maxSize!: pulumi.Output<number>;
    /**
     * Minimum number of node.
     */
    public readonly minSize!: pulumi.Output<number>;
    /**
     * Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
     */
    public readonly multiZoneSubnetPolicy!: pulumi.Output<string | undefined>;
    /**
     * Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (`-`) and decimal points.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Node config.
     */
    public readonly nodeConfig!: pulumi.Output<outputs.Kubernetes.NodePoolNodeConfig | undefined>;
    /**
     * The total node count.
     */
    public /*out*/ readonly nodeCount!: pulumi.Output<number>;
    /**
     * Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1 LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new nodes, not including the existing nodes.
     */
    public readonly nodeOs!: pulumi.Output<string | undefined>;
    /**
     * The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter will only affect new nodes, not including the existing nodes.
     */
    public readonly nodeOsType!: pulumi.Output<string | undefined>;
    /**
     * Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
     */
    public readonly retryPolicy!: pulumi.Output<string | undefined>;
    /**
     * Name of relative scaling group.
     */
    public readonly scalingGroupName!: pulumi.Output<string>;
    /**
     * Project ID the scaling group belongs to.
     */
    public readonly scalingGroupProjectId!: pulumi.Output<number | undefined>;
    /**
     * Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances), `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given priority to the shutdown of the instance. If the number of instances is still lower than the expected number of instances after the startup, the instance will be created, and the method of destroying the instance will still be used for shrinking).
     */
    public readonly scalingMode!: pulumi.Output<string | undefined>;
    /**
     * Status of the node pool.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * ID list of subnet, and for VPC it is required.
     */
    public readonly subnetIds!: pulumi.Output<string[] | undefined>;
    /**
     * Taints of kubernetes node pool created nodes.
     */
    public readonly taints!: pulumi.Output<outputs.Kubernetes.NodePoolTaint[] | undefined>;
    /**
     * Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
     */
    public readonly terminationPolicies!: pulumi.Output<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    public readonly unschedulable!: pulumi.Output<number | undefined>;
    /**
     * ID of VPC network.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * List of auto scaling group available zones, for Basic network it is required.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a NodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodePoolArgs | NodePoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NodePoolState | undefined;
            resourceInputs["autoScalingConfig"] = state ? state.autoScalingConfig : undefined;
            resourceInputs["autoScalingGroupId"] = state ? state.autoScalingGroupId : undefined;
            resourceInputs["autoscalingAddedTotal"] = state ? state.autoscalingAddedTotal : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["defaultCooldown"] = state ? state.defaultCooldown : undefined;
            resourceInputs["deleteKeepInstance"] = state ? state.deleteKeepInstance : undefined;
            resourceInputs["desiredCapacity"] = state ? state.desiredCapacity : undefined;
            resourceInputs["enableAutoScale"] = state ? state.enableAutoScale : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["launchConfigId"] = state ? state.launchConfigId : undefined;
            resourceInputs["manuallyAddedTotal"] = state ? state.manuallyAddedTotal : undefined;
            resourceInputs["maxSize"] = state ? state.maxSize : undefined;
            resourceInputs["minSize"] = state ? state.minSize : undefined;
            resourceInputs["multiZoneSubnetPolicy"] = state ? state.multiZoneSubnetPolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodeConfig"] = state ? state.nodeConfig : undefined;
            resourceInputs["nodeCount"] = state ? state.nodeCount : undefined;
            resourceInputs["nodeOs"] = state ? state.nodeOs : undefined;
            resourceInputs["nodeOsType"] = state ? state.nodeOsType : undefined;
            resourceInputs["retryPolicy"] = state ? state.retryPolicy : undefined;
            resourceInputs["scalingGroupName"] = state ? state.scalingGroupName : undefined;
            resourceInputs["scalingGroupProjectId"] = state ? state.scalingGroupProjectId : undefined;
            resourceInputs["scalingMode"] = state ? state.scalingMode : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["taints"] = state ? state.taints : undefined;
            resourceInputs["terminationPolicies"] = state ? state.terminationPolicies : undefined;
            resourceInputs["unschedulable"] = state ? state.unschedulable : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as NodePoolArgs | undefined;
            if ((!args || args.autoScalingConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoScalingConfig'");
            }
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.maxSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxSize'");
            }
            if ((!args || args.minSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'minSize'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["autoScalingConfig"] = args ? args.autoScalingConfig : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["defaultCooldown"] = args ? args.defaultCooldown : undefined;
            resourceInputs["deleteKeepInstance"] = args ? args.deleteKeepInstance : undefined;
            resourceInputs["desiredCapacity"] = args ? args.desiredCapacity : undefined;
            resourceInputs["enableAutoScale"] = args ? args.enableAutoScale : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maxSize"] = args ? args.maxSize : undefined;
            resourceInputs["minSize"] = args ? args.minSize : undefined;
            resourceInputs["multiZoneSubnetPolicy"] = args ? args.multiZoneSubnetPolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nodeConfig"] = args ? args.nodeConfig : undefined;
            resourceInputs["nodeOs"] = args ? args.nodeOs : undefined;
            resourceInputs["nodeOsType"] = args ? args.nodeOsType : undefined;
            resourceInputs["retryPolicy"] = args ? args.retryPolicy : undefined;
            resourceInputs["scalingGroupName"] = args ? args.scalingGroupName : undefined;
            resourceInputs["scalingGroupProjectId"] = args ? args.scalingGroupProjectId : undefined;
            resourceInputs["scalingMode"] = args ? args.scalingMode : undefined;
            resourceInputs["subnetIds"] = args ? args.subnetIds : undefined;
            resourceInputs["taints"] = args ? args.taints : undefined;
            resourceInputs["terminationPolicies"] = args ? args.terminationPolicies : undefined;
            resourceInputs["unschedulable"] = args ? args.unschedulable : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["autoScalingGroupId"] = undefined /*out*/;
            resourceInputs["autoscalingAddedTotal"] = undefined /*out*/;
            resourceInputs["launchConfigId"] = undefined /*out*/;
            resourceInputs["manuallyAddedTotal"] = undefined /*out*/;
            resourceInputs["nodeCount"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NodePool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodePool resources.
 */
export interface NodePoolState {
    /**
     * Auto scaling config parameters.
     */
    autoScalingConfig?: pulumi.Input<inputs.Kubernetes.NodePoolAutoScalingConfig>;
    /**
     * The auto scaling group ID.
     */
    autoScalingGroupId?: pulumi.Input<string>;
    /**
     * The total of autoscaling added node.
     */
    autoscalingAddedTotal?: pulumi.Input<number>;
    /**
     * ID of the cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Seconds of scaling group cool down. Default value is `300`.
     */
    defaultCooldown?: pulumi.Input<number>;
    /**
     * Indicate to keep the CVM instance when delete the node pool. Default is `true`.
     */
    deleteKeepInstance?: pulumi.Input<boolean>;
    /**
     * Desired capacity ot the node. If `enableAutoScale` is set `true`, this will be a computed parameter.
     */
    desiredCapacity?: pulumi.Input<number>;
    /**
     * Indicate whether to enable auto scaling or not.
     */
    enableAutoScale?: pulumi.Input<boolean>;
    /**
     * Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The launch config ID.
     */
    launchConfigId?: pulumi.Input<string>;
    /**
     * The total of manually added node.
     */
    manuallyAddedTotal?: pulumi.Input<number>;
    /**
     * Maximum number of node.
     */
    maxSize?: pulumi.Input<number>;
    /**
     * Minimum number of node.
     */
    minSize?: pulumi.Input<number>;
    /**
     * Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
     */
    multiZoneSubnetPolicy?: pulumi.Input<string>;
    /**
     * Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (`-`) and decimal points.
     */
    name?: pulumi.Input<string>;
    /**
     * Node config.
     */
    nodeConfig?: pulumi.Input<inputs.Kubernetes.NodePoolNodeConfig>;
    /**
     * The total node count.
     */
    nodeCount?: pulumi.Input<number>;
    /**
     * Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1 LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new nodes, not including the existing nodes.
     */
    nodeOs?: pulumi.Input<string>;
    /**
     * The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter will only affect new nodes, not including the existing nodes.
     */
    nodeOsType?: pulumi.Input<string>;
    /**
     * Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
     */
    retryPolicy?: pulumi.Input<string>;
    /**
     * Name of relative scaling group.
     */
    scalingGroupName?: pulumi.Input<string>;
    /**
     * Project ID the scaling group belongs to.
     */
    scalingGroupProjectId?: pulumi.Input<number>;
    /**
     * Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances), `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given priority to the shutdown of the instance. If the number of instances is still lower than the expected number of instances after the startup, the instance will be created, and the method of destroying the instance will still be used for shrinking).
     */
    scalingMode?: pulumi.Input<string>;
    /**
     * Status of the node pool.
     */
    status?: pulumi.Input<string>;
    /**
     * ID list of subnet, and for VPC it is required.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Taints of kubernetes node pool created nodes.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NodePoolTaint>[]>;
    /**
     * Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
     */
    terminationPolicies?: pulumi.Input<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * ID of VPC network.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * List of auto scaling group available zones, for Basic network it is required.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a NodePool resource.
 */
export interface NodePoolArgs {
    /**
     * Auto scaling config parameters.
     */
    autoScalingConfig: pulumi.Input<inputs.Kubernetes.NodePoolAutoScalingConfig>;
    /**
     * ID of the cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Seconds of scaling group cool down. Default value is `300`.
     */
    defaultCooldown?: pulumi.Input<number>;
    /**
     * Indicate to keep the CVM instance when delete the node pool. Default is `true`.
     */
    deleteKeepInstance?: pulumi.Input<boolean>;
    /**
     * Desired capacity ot the node. If `enableAutoScale` is set `true`, this will be a computed parameter.
     */
    desiredCapacity?: pulumi.Input<number>;
    /**
     * Indicate whether to enable auto scaling or not.
     */
    enableAutoScale?: pulumi.Input<boolean>;
    /**
     * Labels of kubernetes node pool created nodes. The label key name does not exceed 63 characters, only supports English, numbers,'/','-', and does not allow beginning with ('/').
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Maximum number of node.
     */
    maxSize: pulumi.Input<number>;
    /**
     * Minimum number of node.
     */
    minSize: pulumi.Input<number>;
    /**
     * Multi-availability zone/subnet policy. Valid values: PRIORITY and EQUALITY. Default value: PRIORITY.
     */
    multiZoneSubnetPolicy?: pulumi.Input<string>;
    /**
     * Name of the node pool. The name does not exceed 25 characters, and only supports Chinese, English, numbers, underscores, separators (`-`) and decimal points.
     */
    name?: pulumi.Input<string>;
    /**
     * Node config.
     */
    nodeConfig?: pulumi.Input<inputs.Kubernetes.NodePoolNodeConfig>;
    /**
     * Operating system of the cluster, the available values include: `tlinux2.4x86_64`, `ubuntu18.04.1x86_64`, `ubuntu16.04.1 LTSx86_64`, `centos7.6.0_x64` and `centos7.2x86_64`. Default is 'tlinux2.4x86_64'. This parameter will only affect new nodes, not including the existing nodes.
     */
    nodeOs?: pulumi.Input<string>;
    /**
     * The image version of the node. Valida values are `DOCKER_CUSTOMIZE` and `GENERAL`. Default is `GENERAL`. This parameter will only affect new nodes, not including the existing nodes.
     */
    nodeOsType?: pulumi.Input<string>;
    /**
     * Available values for retry policies include `IMMEDIATE_RETRY` and `INCREMENTAL_INTERVALS`.
     */
    retryPolicy?: pulumi.Input<string>;
    /**
     * Name of relative scaling group.
     */
    scalingGroupName?: pulumi.Input<string>;
    /**
     * Project ID the scaling group belongs to.
     */
    scalingGroupProjectId?: pulumi.Input<number>;
    /**
     * Auto scaling mode. Valid values are `CLASSIC_SCALING`(scaling by create/destroy instances), `WAKE_UP_STOPPED_SCALING`(Boot priority for expansion. When expanding the capacity, the shutdown operation is given priority to the shutdown of the instance. If the number of instances is still lower than the expected number of instances after the startup, the instance will be created, and the method of destroying the instance will still be used for shrinking).
     */
    scalingMode?: pulumi.Input<string>;
    /**
     * ID list of subnet, and for VPC it is required.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Taints of kubernetes node pool created nodes.
     */
    taints?: pulumi.Input<pulumi.Input<inputs.Kubernetes.NodePoolTaint>[]>;
    /**
     * Policy of scaling group termination. Available values: `["OLDEST_INSTANCE"]`, `["NEWEST_INSTANCE"]`.
     */
    terminationPolicies?: pulumi.Input<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * ID of VPC network.
     */
    vpcId: pulumi.Input<string>;
    /**
     * List of auto scaling group available zones, for Basic network it is required.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
