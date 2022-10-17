// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provide a resource to attach an existing  cvm to kubernetes cluster.
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
 * const clusterCidr = config.get("clusterCidr") || "172.16.0.0/16";
 * const defaultInstanceType = config.get("defaultInstanceType") || "S1.SMALL1";
 * const defaultInstance = tencentcloud.Images.getInstance({
 *     imageTypes: ["PUBLIC_IMAGE"],
 *     osName: "centos",
 * });
 * const vpc = tencentcloud.Vpc.getSubnets({
 *     isDefault: true,
 *     availabilityZone: availabilityZone,
 * });
 * const defaultTypes = tencentcloud.Instance.getTypes({
 *     filters: [{
 *         name: "instance-family",
 *         values: ["SA2"],
 *     }],
 *     cpuCoreCount: 8,
 *     memorySize: 16,
 * });
 * const foo = new tencentcloud.instance.Instance("foo", {
 *     instanceName: "tf-auto-test-1-1",
 *     availabilityZone: availabilityZone,
 *     imageId: defaultInstance.then(defaultInstance => defaultInstance.images?[0]?.imageId),
 *     instanceType: defaultInstanceType,
 *     systemDiskType: "CLOUD_PREMIUM",
 *     systemDiskSize: 50,
 * });
 * const managedCluster = new tencentcloud.kubernetes.Cluster("managedCluster", {
 *     vpcId: vpc.then(vpc => vpc.instanceLists?[0]?.vpcId),
 *     clusterCidr: "10.1.0.0/16",
 *     clusterMaxPodNum: 32,
 *     clusterName: "keep",
 *     clusterDesc: "test cluster desc",
 *     clusterMaxServiceNum: 32,
 *     workerConfigs: [{
 *         count: 1,
 *         availabilityZone: availabilityZone,
 *         instanceType: defaultInstanceType,
 *         systemDiskType: "CLOUD_SSD",
 *         systemDiskSize: 60,
 *         internetChargeType: "TRAFFIC_POSTPAID_BY_HOUR",
 *         internetMaxBandwidthOut: 100,
 *         publicIpAssigned: true,
 *         subnetId: vpc.then(vpc => vpc.instanceLists?[0]?.subnetId),
 *         dataDisks: [{
 *             diskType: "CLOUD_PREMIUM",
 *             diskSize: 50,
 *         }],
 *         enhancedSecurityService: false,
 *         enhancedMonitorService: false,
 *         userData: "dGVzdA==",
 *         password: "ZZXXccvv1212",
 *     }],
 *     clusterDeployType: "MANAGED_CLUSTER",
 * });
 * const testAttach = new tencentcloud.kubernetes.ClusterAttachment("testAttach", {
 *     clusterId: managedCluster.id,
 *     instanceId: foo.id,
 *     password: "Lo4wbdit",
 *     labels: {
 *         test1: "test1",
 *         test2: "test2",
 *     },
 *     workerConfigOverrides: {
 *         desiredPodNum: 8,
 *     },
 * });
 * ```
 */
export class ClusterAttachment extends pulumi.CustomResource {
    /**
     * Get an existing ClusterAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterAttachmentState, opts?: pulumi.CustomResourceOptions): ClusterAttachment {
        return new ClusterAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Kubernetes/clusterAttachment:ClusterAttachment';

    /**
     * Returns true if the given object is an instance of ClusterAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterAttachment.__pulumiType;
    }

    /**
     * ID of the cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
     */
    public readonly hostname!: pulumi.Output<string | undefined>;
    /**
     * ID of the CVM instance, this cvm will reinstall the system.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
     */
    public readonly keyIds!: pulumi.Output<string | undefined>;
    /**
     * Labels of tke attachment exits CVM.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Password to access, should be set if `keyIds` not set.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * A list of security group IDs after attach to cluster.
     */
    public /*out*/ readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * State of the node.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    public readonly unschedulable!: pulumi.Output<number | undefined>;
    /**
     * Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
     */
    public readonly workerConfig!: pulumi.Output<outputs.Kubernetes.ClusterAttachmentWorkerConfig | undefined>;
    /**
     * Override variable worker_config, commonly used to attach existing instances.
     */
    public readonly workerConfigOverrides!: pulumi.Output<outputs.Kubernetes.ClusterAttachmentWorkerConfigOverrides | undefined>;

    /**
     * Create a ClusterAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterAttachmentArgs | ClusterAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterAttachmentState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["keyIds"] = state ? state.keyIds : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["securityGroups"] = state ? state.securityGroups : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["unschedulable"] = state ? state.unschedulable : undefined;
            resourceInputs["workerConfig"] = state ? state.workerConfig : undefined;
            resourceInputs["workerConfigOverrides"] = state ? state.workerConfigOverrides : undefined;
        } else {
            const args = argsOrState as ClusterAttachmentArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["keyIds"] = args ? args.keyIds : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["unschedulable"] = args ? args.unschedulable : undefined;
            resourceInputs["workerConfig"] = args ? args.workerConfig : undefined;
            resourceInputs["workerConfigOverrides"] = args ? args.workerConfigOverrides : undefined;
            resourceInputs["securityGroups"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterAttachment resources.
 */
export interface ClusterAttachmentState {
    /**
     * ID of the cluster.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
     */
    hostname?: pulumi.Input<string>;
    /**
     * ID of the CVM instance, this cvm will reinstall the system.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
     */
    keyIds?: pulumi.Input<string>;
    /**
     * Labels of tke attachment exits CVM.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Password to access, should be set if `keyIds` not set.
     */
    password?: pulumi.Input<string>;
    /**
     * A list of security group IDs after attach to cluster.
     */
    securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * State of the node.
     */
    state?: pulumi.Input<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
     */
    workerConfig?: pulumi.Input<inputs.Kubernetes.ClusterAttachmentWorkerConfig>;
    /**
     * Override variable worker_config, commonly used to attach existing instances.
     */
    workerConfigOverrides?: pulumi.Input<inputs.Kubernetes.ClusterAttachmentWorkerConfigOverrides>;
}

/**
 * The set of arguments for constructing a ClusterAttachment resource.
 */
export interface ClusterAttachmentArgs {
    /**
     * ID of the cluster.
     */
    clusterId: pulumi.Input<string>;
    /**
     * The host name of the attached instance. Dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows example: The length of the name character is [2, 15], letters (capitalization is not restricted), numbers and dashes (-) are allowed, dots (.) are not supported, and not all numbers are allowed. Examples of other types (Linux, etc.): The character length is [2, 60], and multiple dots are allowed. There is a segment between the dots. Each segment allows letters (with no limitation on capitalization), numbers and dashes (-).
     */
    hostname?: pulumi.Input<string>;
    /**
     * ID of the CVM instance, this cvm will reinstall the system.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The key pair to use for the instance, it looks like skey-16jig7tx, it should be set if `password` not set.
     */
    keyIds?: pulumi.Input<string>;
    /**
     * Labels of tke attachment exits CVM.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * Password to access, should be set if `keyIds` not set.
     */
    password?: pulumi.Input<string>;
    /**
     * Sets whether the joining node participates in the schedule. Default is '0'. Participate in scheduling.
     */
    unschedulable?: pulumi.Input<number>;
    /**
     * Deploy the machine configuration information of the 'WORKER', commonly used to attach existing instances.
     */
    workerConfig?: pulumi.Input<inputs.Kubernetes.ClusterAttachmentWorkerConfig>;
    /**
     * Override variable worker_config, commonly used to attach existing instances.
     */
    workerConfigOverrides?: pulumi.Input<inputs.Kubernetes.ClusterAttachmentWorkerConfigOverrides>;
}
