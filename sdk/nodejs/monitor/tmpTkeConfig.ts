// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tke tmpPrometheusConfig
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = new tencentcloud.Monitor.TmpTkeConfig("foo", {
 *     clusterId: "xxx",
 *     clusterType: "xxx",
 *     instanceId: "xxx",
 *     podMonitors: [{
 *         config: `apiVersion: monitoring.coreos.com/v1
 * kind: PodMonitor
 * metadata:
 *   name: pod-monitor-001
 *   namespace: mynamespace
 * `,
 *         name: "mynamespace/pod-monitor-001", // name with the specified namespace
 *     }],
 *     rawJobs: [{
 *         config: "your config for raw_jobs_001\n",
 *         name: "raw_jobs_001",
 *     }],
 *     serviceMonitors: [{
 *         config: `apiVersion: monitoring.coreos.com/v1
 * kind: ServiceMonitor
 * metadata:
 *   name: service-monitor-001
 *   namespace: kube-system
 * `,
 *         name: "kube-system/service-monitor-001", // name with default namespace kube-system
 *     }],
 * });
 * ```
 */
export class TmpTkeConfig extends pulumi.CustomResource {
    /**
     * Get an existing TmpTkeConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TmpTkeConfigState, opts?: pulumi.CustomResourceOptions): TmpTkeConfig {
        return new TmpTkeConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Monitor/tmpTkeConfig:TmpTkeConfig';

    /**
     * Returns true if the given object is an instance of TmpTkeConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TmpTkeConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TmpTkeConfig.__pulumiType;
    }

    /**
     * ID of cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Type of cluster.
     */
    public readonly clusterType!: pulumi.Output<string>;
    /**
     * Config.
     */
    public /*out*/ readonly config!: pulumi.Output<string>;
    /**
     * ID of instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Configuration of the pod monitors.
     */
    public readonly podMonitors!: pulumi.Output<outputs.Monitor.TmpTkeConfigPodMonitor[] | undefined>;
    /**
     * Configuration of the native prometheus job.
     */
    public readonly rawJobs!: pulumi.Output<outputs.Monitor.TmpTkeConfigRawJob[] | undefined>;
    /**
     * Configuration of the service monitors.
     */
    public readonly serviceMonitors!: pulumi.Output<outputs.Monitor.TmpTkeConfigServiceMonitor[] | undefined>;

    /**
     * Create a TmpTkeConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TmpTkeConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TmpTkeConfigArgs | TmpTkeConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TmpTkeConfigState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["podMonitors"] = state ? state.podMonitors : undefined;
            resourceInputs["rawJobs"] = state ? state.rawJobs : undefined;
            resourceInputs["serviceMonitors"] = state ? state.serviceMonitors : undefined;
        } else {
            const args = argsOrState as TmpTkeConfigArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.clusterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["podMonitors"] = args ? args.podMonitors : undefined;
            resourceInputs["rawJobs"] = args ? args.rawJobs : undefined;
            resourceInputs["serviceMonitors"] = args ? args.serviceMonitors : undefined;
            resourceInputs["config"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TmpTkeConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TmpTkeConfig resources.
 */
export interface TmpTkeConfigState {
    /**
     * ID of cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Type of cluster.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * Config.
     */
    config?: pulumi.Input<string>;
    /**
     * ID of instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Configuration of the pod monitors.
     */
    podMonitors?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpTkeConfigPodMonitor>[]>;
    /**
     * Configuration of the native prometheus job.
     */
    rawJobs?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpTkeConfigRawJob>[]>;
    /**
     * Configuration of the service monitors.
     */
    serviceMonitors?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpTkeConfigServiceMonitor>[]>;
}

/**
 * The set of arguments for constructing a TmpTkeConfig resource.
 */
export interface TmpTkeConfigArgs {
    /**
     * ID of cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Type of cluster.
     */
    clusterType: pulumi.Input<string>;
    /**
     * ID of instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Configuration of the pod monitors.
     */
    podMonitors?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpTkeConfigPodMonitor>[]>;
    /**
     * Configuration of the native prometheus job.
     */
    rawJobs?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpTkeConfigRawJob>[]>;
    /**
     * Configuration of the service monitors.
     */
    serviceMonitors?: pulumi.Input<pulumi.Input<inputs.Monitor.TmpTkeConfigServiceMonitor>[]>;
}
