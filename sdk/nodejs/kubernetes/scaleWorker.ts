// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to increase instance to cluster
 *
 * > **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extraArgs`), you need to submit a ticket for application.
 *
 * > **NOTE:** Import Node: Currently, only one node can be imported at a time.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-3";
 * const subnet = config.get("subnet") || "subnet-pqfek0t8";
 * const scaleInstanceType = config.get("scaleInstanceType") || "S2.LARGE16";
 * const testScale = new tencentcloud.kubernetes.ScaleWorker("testScale", {
 *     clusterId: "cls-godovr32",
 *     desiredPodNum: 16,
 *     labels: {
 *         test1: "test1",
 *         test2: "test2",
 *     },
 *     workerConfig: {
 *         count: 3,
 *         availabilityZone: availabilityZone,
 *         instanceType: scaleInstanceType,
 *         subnetId: subnet,
 *         systemDiskType: "CLOUD_SSD",
 *         systemDiskSize: 50,
 *         internetChargeType: "TRAFFIC_POSTPAID_BY_HOUR",
 *         internetMaxBandwidthOut: 100,
 *         publicIpAssigned: true,
 *         dataDisks: [{
 *             diskType: "CLOUD_PREMIUM",
 *             diskSize: 50,
 *         }],
 *         enhancedSecurityService: false,
 *         enhancedMonitorService: false,
 *         userData: "dGVzdA==",
 *         password: "AABBccdd1122",
 *     },
 * });
 * ```
 * ### Use Kubelet
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "ap-guangzhou-3";
 * const subnet = config.get("subnet") || "subnet-pqfek0t8";
 * const scaleInstanceType = config.get("scaleInstanceType") || "S2.LARGE16";
 * const testScale = new tencentcloud.kubernetes.ScaleWorker("testScale", {
 *     clusterId: "cls-godovr32",
 *     extraArgs: ["root-dir=/var/lib/kubelet"],
 *     labels: {
 *         test1: "test1",
 *         test2: "test2",
 *     },
 *     workerConfig: {
 *         count: 3,
 *         availabilityZone: availabilityZone,
 *         instanceType: scaleInstanceType,
 *         subnetId: subnet,
 *         systemDiskType: "CLOUD_SSD",
 *         systemDiskSize: 50,
 *         internetChargeType: "TRAFFIC_POSTPAID_BY_HOUR",
 *         internetMaxBandwidthOut: 100,
 *         publicIpAssigned: true,
 *         dataDisks: [{
 *             diskType: "CLOUD_PREMIUM",
 *             diskSize: 50,
 *         }],
 *         enhancedSecurityService: false,
 *         enhancedMonitorService: false,
 *         userData: "dGVzdA==",
 *         password: "AABBccdd1122",
 *     },
 * });
 * ```
 */
export class ScaleWorker extends pulumi.CustomResource {
    /**
     * Get an existing ScaleWorker resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScaleWorkerState, opts?: pulumi.CustomResourceOptions): ScaleWorker {
        return new ScaleWorker(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Kubernetes/scaleWorker:ScaleWorker';

    /**
     * Returns true if the given object is an instance of ScaleWorker.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScaleWorker {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScaleWorker.__pulumiType;
    }

    /**
     * ID of the cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Configurations of data disk.
     */
    public readonly dataDisks!: pulumi.Output<outputs.Kubernetes.ScaleWorkerDataDisk[] | undefined>;
    /**
     * Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
     */
    public readonly desiredPodNum!: pulumi.Output<number | undefined>;
    /**
     * Docker graph path. Default is `/var/lib/docker`.
     */
    public readonly dockerGraphPath!: pulumi.Output<string | undefined>;
    /**
     * Custom parameter information related to the node.
     */
    public readonly extraArgs!: pulumi.Output<string[] | undefined>;
    /**
     * GPU driver parameters.
     */
    public readonly gpuArgs!: pulumi.Output<outputs.Kubernetes.ScaleWorkerGpuArgs | undefined>;
    /**
     * Labels of kubernetes scale worker created nodes.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Mount target. Default is not mounting.
     */
    public readonly mountTarget!: pulumi.Output<string | undefined>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    public readonly unschedulable!: pulumi.Output<number | undefined>;
    /**
     * Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
     */
    public readonly workerConfig!: pulumi.Output<outputs.Kubernetes.ScaleWorkerWorkerConfig>;
    /**
     * An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
     */
    public /*out*/ readonly workerInstancesLists!: pulumi.Output<outputs.Kubernetes.ScaleWorkerWorkerInstancesList[]>;

    /**
     * Create a ScaleWorker resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScaleWorkerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScaleWorkerArgs | ScaleWorkerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScaleWorkerState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["dataDisks"] = state ? state.dataDisks : undefined;
            resourceInputs["desiredPodNum"] = state ? state.desiredPodNum : undefined;
            resourceInputs["dockerGraphPath"] = state ? state.dockerGraphPath : undefined;
            resourceInputs["extraArgs"] = state ? state.extraArgs : undefined;
            resourceInputs["gpuArgs"] = state ? state.gpuArgs : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["mountTarget"] = state ? state.mountTarget : undefined;
            resourceInputs["unschedulable"] = state ? state.unschedulable : undefined;
            resourceInputs["workerConfig"] = state ? state.workerConfig : undefined;
            resourceInputs["workerInstancesLists"] = state ? state.workerInstancesLists : undefined;
        } else {
            const args = argsOrState as ScaleWorkerArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.workerConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workerConfig'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["dataDisks"] = args ? args.dataDisks : undefined;
            resourceInputs["desiredPodNum"] = args ? args.desiredPodNum : undefined;
            resourceInputs["dockerGraphPath"] = args ? args.dockerGraphPath : undefined;
            resourceInputs["extraArgs"] = args ? args.extraArgs : undefined;
            resourceInputs["gpuArgs"] = args ? args.gpuArgs : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["mountTarget"] = args ? args.mountTarget : undefined;
            resourceInputs["unschedulable"] = args ? args.unschedulable : undefined;
            resourceInputs["workerConfig"] = args ? args.workerConfig : undefined;
            resourceInputs["workerInstancesLists"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScaleWorker.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScaleWorker resources.
 */
export interface ScaleWorkerState {
    /**
     * ID of the cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Configurations of data disk.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Kubernetes.ScaleWorkerDataDisk>[]>;
    /**
     * Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
     */
    desiredPodNum?: pulumi.Input<number>;
    /**
     * Docker graph path. Default is `/var/lib/docker`.
     */
    dockerGraphPath?: pulumi.Input<string>;
    /**
     * Custom parameter information related to the node.
     */
    extraArgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GPU driver parameters.
     */
    gpuArgs?: pulumi.Input<inputs.Kubernetes.ScaleWorkerGpuArgs>;
    /**
     * Labels of kubernetes scale worker created nodes.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Mount target. Default is not mounting.
     */
    mountTarget?: pulumi.Input<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
     */
    workerConfig?: pulumi.Input<inputs.Kubernetes.ScaleWorkerWorkerConfig>;
    /**
     * An information list of kubernetes cluster 'WORKER'. Each element contains the following attributes:
     */
    workerInstancesLists?: pulumi.Input<pulumi.Input<inputs.Kubernetes.ScaleWorkerWorkerInstancesList>[]>;
}

/**
 * The set of arguments for constructing a ScaleWorker resource.
 */
export interface ScaleWorkerArgs {
    /**
     * ID of the cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Configurations of data disk.
     */
    dataDisks?: pulumi.Input<pulumi.Input<inputs.Kubernetes.ScaleWorkerDataDisk>[]>;
    /**
     * Indicate to set desired pod number in current node. Valid when the cluster enable customized pod cidr.
     */
    desiredPodNum?: pulumi.Input<number>;
    /**
     * Docker graph path. Default is `/var/lib/docker`.
     */
    dockerGraphPath?: pulumi.Input<string>;
    /**
     * Custom parameter information related to the node.
     */
    extraArgs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GPU driver parameters.
     */
    gpuArgs?: pulumi.Input<inputs.Kubernetes.ScaleWorkerGpuArgs>;
    /**
     * Labels of kubernetes scale worker created nodes.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Mount target. Default is not mounting.
     */
    mountTarget?: pulumi.Input<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * Deploy the machine configuration information of the 'WORK' service, and create <=20 units for common users.
     */
    workerConfig: pulumi.Input<inputs.Kubernetes.ScaleWorkerWorkerConfig>;
}
