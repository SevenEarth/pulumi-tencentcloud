// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provide a resource to configure kubernetes cluster authentication info.
 *
 * > **NOTE:** Only available for cluster version >= 1.20
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
 * const default = tencentcloud.Images.getInstance({
 *     imageTypes: ["PUBLIC_IMAGE"],
 *     osName: "centos",
 * });
 * const vpc = tencentcloud.Vpc.getSubnets({
 *     isDefault: true,
 *     availabilityZone: availabilityZone,
 * });
 * const managedCluster = new tencentcloud.kubernetes.Cluster("managedCluster", {
 *     vpcId: vpc.then(vpc => vpc.instanceLists?[0]?.vpcId),
 *     clusterCidr: "10.31.0.0/16",
 *     clusterMaxPodNum: 32,
 *     clusterName: "keep",
 *     clusterDesc: "test cluster desc",
 *     clusterVersion: "1.20.6",
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
 * const testAuthAttach = new tencentcloud.kubernetes.AuthAttachment("testAuthAttach", {
 *     clusterId: managedCluster.id,
 *     jwksUri: pulumi.interpolate`https://${managedCluster.id}.ccs.tencent-cloud.com/openid/v1/jwks`,
 *     issuer: pulumi.interpolate`https://${managedCluster.id}.ccs.tencent-cloud.com`,
 *     autoCreateDiscoveryAnonymousAuth: true,
 * });
 * ```
 */
export class AuthAttachment extends pulumi.CustomResource {
    /**
     * Get an existing AuthAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthAttachmentState, opts?: pulumi.CustomResourceOptions): AuthAttachment {
        return new AuthAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Kubernetes/authAttachment:AuthAttachment';

    /**
     * Returns true if the given object is an instance of AuthAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthAttachment.__pulumiType;
    }

    /**
     * If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.
     */
    public readonly autoCreateDiscoveryAnonymousAuth!: pulumi.Output<boolean | undefined>;
    /**
     * ID of clusters.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Specify service-account-issuer.
     */
    public readonly issuer!: pulumi.Output<string>;
    /**
     * Specify service-account-jwks-uri.
     */
    public readonly jwksUri!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthAttachmentArgs | AuthAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthAttachmentState | undefined;
            resourceInputs["autoCreateDiscoveryAnonymousAuth"] = state ? state.autoCreateDiscoveryAnonymousAuth : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["jwksUri"] = state ? state.jwksUri : undefined;
        } else {
            const args = argsOrState as AuthAttachmentArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.issuer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuer'");
            }
            resourceInputs["autoCreateDiscoveryAnonymousAuth"] = args ? args.autoCreateDiscoveryAnonymousAuth : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["jwksUri"] = args ? args.jwksUri : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthAttachment resources.
 */
export interface AuthAttachmentState {
    /**
     * If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.
     */
    autoCreateDiscoveryAnonymousAuth?: pulumi.Input<boolean>;
    /**
     * ID of clusters.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Specify service-account-issuer.
     */
    issuer?: pulumi.Input<string>;
    /**
     * Specify service-account-jwks-uri.
     */
    jwksUri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthAttachment resource.
 */
export interface AuthAttachmentArgs {
    /**
     * If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.
     */
    autoCreateDiscoveryAnonymousAuth?: pulumi.Input<boolean>;
    /**
     * ID of clusters.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Specify service-account-issuer.
     */
    issuer: pulumi.Input<string>;
    /**
     * Specify service-account-jwks-uri.
     */
    jwksUri?: pulumi.Input<string>;
}
