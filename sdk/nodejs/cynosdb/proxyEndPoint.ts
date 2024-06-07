// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cynosdb proxyEndPoint
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const proxyEndPoint = new tencentcloud.cynosdb.ProxyEndPoint("proxyEndPoint", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     instanceWeights: [{
 *         instanceId: "cynosdbmysql-ins-afqx1hy0",
 *         weight: 1,
 *     }],
 *     uniqueSubnetId: "subnet-dwj7ipnc",
 *     uniqueVpcId: "vpc-4owdpnwr",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const proxyEndPoint = new tencentcloud.cynosdb.ProxyEndPoint("proxyEndPoint", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     instanceWeights: [{
 *         instanceId: "cynosdbmysql-ins-afqx1hy0",
 *         weight: 1,
 *     }],
 *     uniqueSubnetId: "subnet-dwj7ipnc",
 *     uniqueVpcId: "vpc-4owdpnwr",
 *     vip: "172.16.112.108",
 *     vport: 3306,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Open connection pool
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const proxyEndPoint = new tencentcloud.cynosdb.ProxyEndPoint("proxyEndPoint", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     connectionPoolTimeOut: 30,
 *     connectionPoolType: "SessionConnectionPool",
 *     instanceWeights: [{
 *         instanceId: "cynosdbmysql-ins-afqx1hy0",
 *         weight: 1,
 *     }],
 *     openConnectionPool: "yes",
 *     uniqueSubnetId: "subnet-dwj7ipnc",
 *     uniqueVpcId: "vpc-4owdpnwr",
 *     vip: "172.16.112.108",
 *     vport: 3306,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Close connection pool
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const proxyEndPoint = new tencentcloud.cynosdb.ProxyEndPoint("proxyEndPoint", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     instanceWeights: [{
 *         instanceId: "cynosdbmysql-ins-afqx1hy0",
 *         weight: 1,
 *     }],
 *     openConnectionPool: "no",
 *     uniqueSubnetId: "subnet-dwj7ipnc",
 *     uniqueVpcId: "vpc-4owdpnwr",
 *     vip: "172.16.112.108",
 *     vport: 3306,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const proxyEndPoint = new tencentcloud.cynosdb.ProxyEndPoint("proxyEndPoint", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     consistencyTimeOut: 30,
 *     consistencyType: "global",
 *     failOver: "yes",
 *     instanceWeights: [{
 *         instanceId: "cynosdbmysql-ins-afqx1hy0",
 *         weight: 1,
 *     }],
 *     openConnectionPool: "no",
 *     rwType: "READWRITE",
 *     uniqueSubnetId: "subnet-dwj7ipnc",
 *     uniqueVpcId: "vpc-4owdpnwr",
 *     vip: "172.16.112.108",
 *     vport: 3306,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const proxyEndPoint = new tencentcloud.cynosdb.ProxyEndPoint("proxyEndPoint", {
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     instanceWeights: [{
 *         instanceId: "cynosdbmysql-ins-rikr6z4o",
 *         weight: 1,
 *     }],
 *     openConnectionPool: "no",
 *     rwType: "READONLY",
 *     uniqueSubnetId: "subnet-dwj7ipnc",
 *     uniqueVpcId: "vpc-4owdpnwr",
 *     vip: "172.16.112.108",
 *     vport: 3306,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Comprehensive parameter examples
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const proxyEndPoint = new tencentcloud.cynosdb.ProxyEndPoint("proxyEndPoint", {
 *     accessMode: "nearby",
 *     autoAddRo: "yes",
 *     clusterId: "cynosdbmysql-bws8h88b",
 *     connectionPoolTimeOut: 30,
 *     connectionPoolType: "SessionConnectionPool",
 *     consistencyTimeOut: 30,
 *     consistencyType: "global",
 *     description: "desc value",
 *     failOver: "yes",
 *     instanceWeights: [{
 *         instanceId: "cynosdbmysql-ins-afqx1hy0",
 *         weight: 1,
 *     }],
 *     openConnectionPool: "yes",
 *     rwType: "READWRITE",
 *     securityGroupIds: ["sg-7kpsbxdb"],
 *     transSplit: true,
 *     uniqueSubnetId: "subnet-dwj7ipnc",
 *     uniqueVpcId: "vpc-4owdpnwr",
 *     vip: "172.16.112.118",
 *     vport: 3306,
 *     weightMode: "system",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class ProxyEndPoint extends pulumi.CustomResource {
    /**
     * Get an existing ProxyEndPoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyEndPointState, opts?: pulumi.CustomResourceOptions): ProxyEndPoint {
        return new ProxyEndPoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cynosdb/proxyEndPoint:ProxyEndPoint';

    /**
     * Returns true if the given object is an instance of ProxyEndPoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProxyEndPoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProxyEndPoint.__pulumiType;
    }

    /**
     * Connection mode: nearby, balance.
     */
    public readonly accessMode!: pulumi.Output<string>;
    /**
     * Do you want to automatically add read-only instances? Yes - Yes, no - Do not automatically add.
     */
    public readonly autoAddRo!: pulumi.Output<string>;
    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Connection pool threshold: unit (second).
     */
    public readonly connectionPoolTimeOut!: pulumi.Output<number>;
    /**
     * Connection pool type: SessionConnectionPool (session level Connection pool).
     */
    public readonly connectionPoolType!: pulumi.Output<string>;
    /**
     * Consistency timeout.
     */
    public readonly consistencyTimeOut!: pulumi.Output<number>;
    /**
     * Consistency type: event, global, session.
     */
    public readonly consistencyType!: pulumi.Output<string>;
    /**
     * Description.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Enable Failover. yes or no.
     */
    public readonly failOver!: pulumi.Output<string>;
    /**
     * Instance Group ID.
     */
    public /*out*/ readonly instanceGroupId!: pulumi.Output<string>;
    /**
     * Instance Weight.
     */
    public readonly instanceWeights!: pulumi.Output<outputs.Cynosdb.ProxyEndPointInstanceWeight[]>;
    /**
     * Whether to enable Connection pool, yes - enable, no - do not enable.
     */
    public readonly openConnectionPool!: pulumi.Output<string>;
    /**
     * Proxy Group ID.
     */
    public /*out*/ readonly proxyGroupId!: pulumi.Output<string>;
    /**
     * Read and write attributes: READWRITE, READONLY.
     */
    public readonly rwType!: pulumi.Output<string>;
    /**
     * Security Group ID Array.
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Transaction splitting.
     */
    public readonly transSplit!: pulumi.Output<boolean>;
    /**
     * The private network subnet ID is consistent with the cluster subnet ID by default.
     */
    public readonly uniqueSubnetId!: pulumi.Output<string>;
    /**
     * Private network ID, which is consistent with the cluster private network ID by default.
     */
    public readonly uniqueVpcId!: pulumi.Output<string>;
    /**
     * VIP Information.
     */
    public readonly vip!: pulumi.Output<string>;
    /**
     * Port Information.
     */
    public readonly vport!: pulumi.Output<number>;
    /**
     * Weight mode: system system allocation, custom customization.
     */
    public readonly weightMode!: pulumi.Output<string>;

    /**
     * Create a ProxyEndPoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyEndPointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyEndPointArgs | ProxyEndPointState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyEndPointState | undefined;
            resourceInputs["accessMode"] = state ? state.accessMode : undefined;
            resourceInputs["autoAddRo"] = state ? state.autoAddRo : undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["connectionPoolTimeOut"] = state ? state.connectionPoolTimeOut : undefined;
            resourceInputs["connectionPoolType"] = state ? state.connectionPoolType : undefined;
            resourceInputs["consistencyTimeOut"] = state ? state.consistencyTimeOut : undefined;
            resourceInputs["consistencyType"] = state ? state.consistencyType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["failOver"] = state ? state.failOver : undefined;
            resourceInputs["instanceGroupId"] = state ? state.instanceGroupId : undefined;
            resourceInputs["instanceWeights"] = state ? state.instanceWeights : undefined;
            resourceInputs["openConnectionPool"] = state ? state.openConnectionPool : undefined;
            resourceInputs["proxyGroupId"] = state ? state.proxyGroupId : undefined;
            resourceInputs["rwType"] = state ? state.rwType : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["transSplit"] = state ? state.transSplit : undefined;
            resourceInputs["uniqueSubnetId"] = state ? state.uniqueSubnetId : undefined;
            resourceInputs["uniqueVpcId"] = state ? state.uniqueVpcId : undefined;
            resourceInputs["vip"] = state ? state.vip : undefined;
            resourceInputs["vport"] = state ? state.vport : undefined;
            resourceInputs["weightMode"] = state ? state.weightMode : undefined;
        } else {
            const args = argsOrState as ProxyEndPointArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.uniqueSubnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uniqueSubnetId'");
            }
            if ((!args || args.uniqueVpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'uniqueVpcId'");
            }
            resourceInputs["accessMode"] = args ? args.accessMode : undefined;
            resourceInputs["autoAddRo"] = args ? args.autoAddRo : undefined;
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["connectionPoolTimeOut"] = args ? args.connectionPoolTimeOut : undefined;
            resourceInputs["connectionPoolType"] = args ? args.connectionPoolType : undefined;
            resourceInputs["consistencyTimeOut"] = args ? args.consistencyTimeOut : undefined;
            resourceInputs["consistencyType"] = args ? args.consistencyType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["failOver"] = args ? args.failOver : undefined;
            resourceInputs["instanceWeights"] = args ? args.instanceWeights : undefined;
            resourceInputs["openConnectionPool"] = args ? args.openConnectionPool : undefined;
            resourceInputs["rwType"] = args ? args.rwType : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["transSplit"] = args ? args.transSplit : undefined;
            resourceInputs["uniqueSubnetId"] = args ? args.uniqueSubnetId : undefined;
            resourceInputs["uniqueVpcId"] = args ? args.uniqueVpcId : undefined;
            resourceInputs["vip"] = args ? args.vip : undefined;
            resourceInputs["vport"] = args ? args.vport : undefined;
            resourceInputs["weightMode"] = args ? args.weightMode : undefined;
            resourceInputs["instanceGroupId"] = undefined /*out*/;
            resourceInputs["proxyGroupId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProxyEndPoint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProxyEndPoint resources.
 */
export interface ProxyEndPointState {
    /**
     * Connection mode: nearby, balance.
     */
    accessMode?: pulumi.Input<string>;
    /**
     * Do you want to automatically add read-only instances? Yes - Yes, no - Do not automatically add.
     */
    autoAddRo?: pulumi.Input<string>;
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
     * Consistency timeout.
     */
    consistencyTimeOut?: pulumi.Input<number>;
    /**
     * Consistency type: event, global, session.
     */
    consistencyType?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable Failover. yes or no.
     */
    failOver?: pulumi.Input<string>;
    /**
     * Instance Group ID.
     */
    instanceGroupId?: pulumi.Input<string>;
    /**
     * Instance Weight.
     */
    instanceWeights?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ProxyEndPointInstanceWeight>[]>;
    /**
     * Whether to enable Connection pool, yes - enable, no - do not enable.
     */
    openConnectionPool?: pulumi.Input<string>;
    /**
     * Proxy Group ID.
     */
    proxyGroupId?: pulumi.Input<string>;
    /**
     * Read and write attributes: READWRITE, READONLY.
     */
    rwType?: pulumi.Input<string>;
    /**
     * Security Group ID Array.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Transaction splitting.
     */
    transSplit?: pulumi.Input<boolean>;
    /**
     * The private network subnet ID is consistent with the cluster subnet ID by default.
     */
    uniqueSubnetId?: pulumi.Input<string>;
    /**
     * Private network ID, which is consistent with the cluster private network ID by default.
     */
    uniqueVpcId?: pulumi.Input<string>;
    /**
     * VIP Information.
     */
    vip?: pulumi.Input<string>;
    /**
     * Port Information.
     */
    vport?: pulumi.Input<number>;
    /**
     * Weight mode: system system allocation, custom customization.
     */
    weightMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProxyEndPoint resource.
 */
export interface ProxyEndPointArgs {
    /**
     * Connection mode: nearby, balance.
     */
    accessMode?: pulumi.Input<string>;
    /**
     * Do you want to automatically add read-only instances? Yes - Yes, no - Do not automatically add.
     */
    autoAddRo?: pulumi.Input<string>;
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
     * Consistency timeout.
     */
    consistencyTimeOut?: pulumi.Input<number>;
    /**
     * Consistency type: event, global, session.
     */
    consistencyType?: pulumi.Input<string>;
    /**
     * Description.
     */
    description?: pulumi.Input<string>;
    /**
     * Enable Failover. yes or no.
     */
    failOver?: pulumi.Input<string>;
    /**
     * Instance Weight.
     */
    instanceWeights?: pulumi.Input<pulumi.Input<inputs.Cynosdb.ProxyEndPointInstanceWeight>[]>;
    /**
     * Whether to enable Connection pool, yes - enable, no - do not enable.
     */
    openConnectionPool?: pulumi.Input<string>;
    /**
     * Read and write attributes: READWRITE, READONLY.
     */
    rwType?: pulumi.Input<string>;
    /**
     * Security Group ID Array.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Transaction splitting.
     */
    transSplit?: pulumi.Input<boolean>;
    /**
     * The private network subnet ID is consistent with the cluster subnet ID by default.
     */
    uniqueSubnetId: pulumi.Input<string>;
    /**
     * Private network ID, which is consistent with the cluster private network ID by default.
     */
    uniqueVpcId: pulumi.Input<string>;
    /**
     * VIP Information.
     */
    vip?: pulumi.Input<string>;
    /**
     * Port Information.
     */
    vport?: pulumi.Input<number>;
    /**
     * Weight mode: system system allocation, custom customization.
     */
    weightMode?: pulumi.Input<string>;
}
