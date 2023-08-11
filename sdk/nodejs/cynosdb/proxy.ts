// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cynosdb proxy
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const proxy = new tencentcloud.Cynosdb.Proxy("proxy", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     connectionPoolTimeOut: 30,
 *     connectionPoolType: "SessionConnectionPool",
 *     cpu: 2,
 *     description: "desc sample",
 *     mem: 4000,
 *     openConnectionPool: "yes",
 *     proxyZones: [{
 *         proxyNodeCount: 2,
 *         proxyNodeZone: "ap-guangzhou-7",
 *     }],
 *     securityGroupIds: ["sg-baxfiao5"],
 *     uniqueSubnetId: "subnet-jdi5xn22",
 *     uniqueVpcId: "vpc-k1t8ickr",
 * });
 * ```
 */
export class Proxy extends pulumi.CustomResource {
    /**
     * Get an existing Proxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyState, opts?: pulumi.CustomResourceOptions): Proxy {
        return new Proxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/proxy:Proxy';

    /**
     * Returns true if the given object is an instance of Proxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Proxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Proxy.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Connection pool threshold: unit (second).
     */
    public readonly connectionPoolTimeOut!: pulumi.Output<number | undefined>;
    /**
     * Connection pool type: SessionConnectionPool (session level Connection pool).
     */
    public readonly connectionPoolType!: pulumi.Output<string | undefined>;
    /**
     * Number of CPU cores.
     */
    public readonly cpu!: pulumi.Output<number>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Memory.
     */
    public readonly mem!: pulumi.Output<number>;
    /**
     * Whether to enable Connection pool, yes - enable, no - do not enable.
     */
    public readonly openConnectionPool!: pulumi.Output<string | undefined>;
    /**
     * Number of database proxy group nodes. If it is set at the same time as the `proxyZones` field, the `proxyZones` parameter shall prevail.
     */
    public readonly proxyCount!: pulumi.Output<number>;
    /**
     * Proxy Group Id.
     */
    public /*out*/ readonly proxyGroupId!: pulumi.Output<string>;
    /**
     * Database node information.
     */
    public readonly proxyZones!: pulumi.Output<outputs.Cynosdb.ProxyProxyZone[]>;
    /**
     * Read only instance list.
     */
    public /*out*/ readonly roInstances!: pulumi.Output<outputs.Cynosdb.ProxyRoInstance[]>;
    /**
     * Security Group ID Array.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * The private network subnet ID is consistent with the cluster subnet ID by default.
     */
    public readonly uniqueSubnetId!: pulumi.Output<string>;
    /**
     * Private network ID, which is consistent with the cluster private network ID by default.
     */
    public readonly uniqueVpcId!: pulumi.Output<string>;

    /**
     * Create a Proxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyArgs | ProxyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["connectionPoolTimeOut"] = state ? state.connectionPoolTimeOut : undefined;
            resourceInputs["connectionPoolType"] = state ? state.connectionPoolType : undefined;
            resourceInputs["cpu"] = state ? state.cpu : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["mem"] = state ? state.mem : undefined;
            resourceInputs["openConnectionPool"] = state ? state.openConnectionPool : undefined;
            resourceInputs["proxyCount"] = state ? state.proxyCount : undefined;
            resourceInputs["proxyGroupId"] = state ? state.proxyGroupId : undefined;
            resourceInputs["proxyZones"] = state ? state.proxyZones : undefined;
            resourceInputs["roInstances"] = state ? state.roInstances : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["uniqueSubnetId"] = state ? state.uniqueSubnetId : undefined;
            resourceInputs["uniqueVpcId"] = state ? state.uniqueVpcId : undefined;
        } else {
            const args = argsOrState as ProxyArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.cpu === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cpu'");
            }
            if ((!args || args.mem === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mem'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["connectionPoolTimeOut"] = args ? args.connectionPoolTimeOut : undefined;
            resourceInputs["connectionPoolType"] = args ? args.connectionPoolType : undefined;
            resourceInputs["cpu"] = args ? args.cpu : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["mem"] = args ? args.mem : undefined;
            resourceInputs["openConnectionPool"] = args ? args.openConnectionPool : undefined;
            resourceInputs["proxyCount"] = args ? args.proxyCount : undefined;
            resourceInputs["proxyZones"] = args ? args.proxyZones : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["uniqueSubnetId"] = args ? args.uniqueSubnetId : undefined;
            resourceInputs["uniqueVpcId"] = args ? args.uniqueVpcId : undefined;
            resourceInputs["proxyGroupId"] = undefined /*out*/;
            resourceInputs["roInstances"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Proxy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Proxy resources.
 */
export interface ProxyState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Connection pool threshold: unit (second).
     */
    connectionPoolTimeOut?: pulumi.Input<number>;
    /**
     * Connection pool type: SessionConnectionPool (session level Connection pool).
     */
    connectionPoolType?: pulumi.Input<string>;
    /**
     * Number of CPU cores.
     */
    cpu?: pulumi.Input<number>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Memory.
     */
    mem?: pulumi.Input<number>;
    /**
     * Whether to enable Connection pool, yes - enable, no - do not enable.
     */
    openConnectionPool?: pulumi.Input<string>;
    /**
     * Number of database proxy group nodes. If it is set at the same time as the `proxyZones` field, the `proxyZones` parameter shall prevail.
     */
    proxyCount?: pulumi.Input<number>;
    /**
     * Proxy Group Id.
     */
    proxyGroupId?: pulumi.Input<string>;
    /**
     * Database node information.
     */
    proxyZones?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ProxyProxyZone>[]>;
    /**
     * Read only instance list.
     */
    roInstances?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ProxyRoInstance>[]>;
    /**
     * Security Group ID Array.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The private network subnet ID is consistent with the cluster subnet ID by default.
     */
    uniqueSubnetId?: pulumi.Input<string>;
    /**
     * Private network ID, which is consistent with the cluster private network ID by default.
     */
    uniqueVpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Proxy resource.
 */
export interface ProxyArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Connection pool threshold: unit (second).
     */
    connectionPoolTimeOut?: pulumi.Input<number>;
    /**
     * Connection pool type: SessionConnectionPool (session level Connection pool).
     */
    connectionPoolType?: pulumi.Input<string>;
    /**
     * Number of CPU cores.
     */
    cpu: pulumi.Input<number>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Memory.
     */
    mem: pulumi.Input<number>;
    /**
     * Whether to enable Connection pool, yes - enable, no - do not enable.
     */
    openConnectionPool?: pulumi.Input<string>;
    /**
     * Number of database proxy group nodes. If it is set at the same time as the `proxyZones` field, the `proxyZones` parameter shall prevail.
     */
    proxyCount?: pulumi.Input<number>;
    /**
     * Database node information.
     */
    proxyZones?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ProxyProxyZone>[]>;
    /**
     * Security Group ID Array.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The private network subnet ID is consistent with the cluster subnet ID by default.
     */
    uniqueSubnetId?: pulumi.Input<string>;
    /**
     * Private network ID, which is consistent with the cluster private network ID by default.
     */
    uniqueVpcId?: pulumi.Input<string>;
}
