// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to creating dedicated tunnels instances.
 *
 * > **NOTE:** 1. ID of the DC is queried, can only apply for this resource offline.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const dcId = config.get("dcId") || "dc-kax48sg7";
 * const dcgId = config.get("dcgId") || "dcg-dmbhf7jf";
 * const vpcId = config.get("vpcId") || "vpc-4h9v4mo3";
 * const bgpMain = new tencentcloud.dcx.Instance("bgpMain", {
 *     bandwidth: 900,
 *     dcId: dcId,
 *     dcgId: dcgId,
 *     networkType: "VPC",
 *     routeType: "BGP",
 *     vlan: 306,
 *     vpcId: vpcId,
 * });
 * const staticMain = new tencentcloud.dcx.Instance("staticMain", {
 *     bandwidth: 900,
 *     dcId: dcId,
 *     dcgId: dcgId,
 *     dcOwnerAccount: "xxxxxxxx",
 *     networkType: "VPC",
 *     routeType: "STATIC",
 *     vlan: 301,
 *     vpcId: vpcId,
 *     routeFilterPrefixes: ["10.10.10.101/32"],
 *     tencentAddress: "100.93.46.1/30",
 *     customerAddress: "100.93.46.2/30",
 * });
 * ```
 *
 * ## Import
 *
 * DCX instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dcx/instance:Instance foo dcx-cbbr1gjk
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dcx/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Bandwidth of the DC.
     */
    public readonly bandwidth!: pulumi.Output<number>;
    /**
     * BGP ASN of the user. A required field within BGP.
     */
    public readonly bgpAsn!: pulumi.Output<number>;
    /**
     * BGP key of the user.
     */
    public readonly bgpAuthKey!: pulumi.Output<string | undefined>;
    /**
     * Creation time of resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Interconnect IP of the DC within client.
     */
    public readonly customerAddress!: pulumi.Output<string>;
    /**
     * ID of the DC to be queried, application deployment offline.
     */
    public readonly dcId!: pulumi.Output<string>;
    /**
     * Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
     */
    public readonly dcOwnerAccount!: pulumi.Output<string>;
    /**
     * ID of the DC Gateway. Currently only new in the console.
     */
    public readonly dcgId!: pulumi.Output<string>;
    /**
     * Name of the dedicated tunnel.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network region.
     */
    public readonly networkRegion!: pulumi.Output<string | undefined>;
    /**
     * Type of the network. Valid value: `VPC`, `BMVPC` and `CCN`. The default value is `VPC`.
     */
    public readonly networkType!: pulumi.Output<string | undefined>;
    /**
     * Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
     */
    public readonly routeFilterPrefixes!: pulumi.Output<string[] | undefined>;
    /**
     * Type of the route, and available values include BGP and STATIC. The default value is `BGP`.
     */
    public readonly routeType!: pulumi.Output<string | undefined>;
    /**
     * State of the dedicated tunnels. Valid value: `PENDING`, `ALLOCATING`, `ALLOCATED`, `ALTERING`, `DELETING`, `DELETED`, `COMFIRMING` and `REJECTED`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Interconnect IP of the DC within Tencent.
     */
    public readonly tencentAddress!: pulumi.Output<string>;
    /**
     * Vlan of the dedicated tunnels. Valid value ranges: (0~3000). `0` means that only one tunnel can be created for the physical connect.
     */
    public readonly vlan!: pulumi.Output<number | undefined>;
    /**
     * ID of the VPC or BMVPC.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            resourceInputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["customerAddress"] = state ? state.customerAddress : undefined;
            resourceInputs["dcId"] = state ? state.dcId : undefined;
            resourceInputs["dcOwnerAccount"] = state ? state.dcOwnerAccount : undefined;
            resourceInputs["dcgId"] = state ? state.dcgId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkRegion"] = state ? state.networkRegion : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["routeFilterPrefixes"] = state ? state.routeFilterPrefixes : undefined;
            resourceInputs["routeType"] = state ? state.routeType : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["tencentAddress"] = state ? state.tencentAddress : undefined;
            resourceInputs["vlan"] = state ? state.vlan : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.dcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dcId'");
            }
            if ((!args || args.dcgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dcgId'");
            }
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            resourceInputs["bgpAuthKey"] = args ? args.bgpAuthKey : undefined;
            resourceInputs["customerAddress"] = args ? args.customerAddress : undefined;
            resourceInputs["dcId"] = args ? args.dcId : undefined;
            resourceInputs["dcOwnerAccount"] = args ? args.dcOwnerAccount : undefined;
            resourceInputs["dcgId"] = args ? args.dcgId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkRegion"] = args ? args.networkRegion : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["routeFilterPrefixes"] = args ? args.routeFilterPrefixes : undefined;
            resourceInputs["routeType"] = args ? args.routeType : undefined;
            resourceInputs["tencentAddress"] = args ? args.tencentAddress : undefined;
            resourceInputs["vlan"] = args ? args.vlan : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Bandwidth of the DC.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * BGP ASN of the user. A required field within BGP.
     */
    bgpAsn?: pulumi.Input<number>;
    /**
     * BGP key of the user.
     */
    bgpAuthKey?: pulumi.Input<string>;
    /**
     * Creation time of resource.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Interconnect IP of the DC within client.
     */
    customerAddress?: pulumi.Input<string>;
    /**
     * ID of the DC to be queried, application deployment offline.
     */
    dcId?: pulumi.Input<string>;
    /**
     * Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
     */
    dcOwnerAccount?: pulumi.Input<string>;
    /**
     * ID of the DC Gateway. Currently only new in the console.
     */
    dcgId?: pulumi.Input<string>;
    /**
     * Name of the dedicated tunnel.
     */
    name?: pulumi.Input<string>;
    /**
     * Network region.
     */
    networkRegion?: pulumi.Input<string>;
    /**
     * Type of the network. Valid value: `VPC`, `BMVPC` and `CCN`. The default value is `VPC`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
     */
    routeFilterPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the route, and available values include BGP and STATIC. The default value is `BGP`.
     */
    routeType?: pulumi.Input<string>;
    /**
     * State of the dedicated tunnels. Valid value: `PENDING`, `ALLOCATING`, `ALLOCATED`, `ALTERING`, `DELETING`, `DELETED`, `COMFIRMING` and `REJECTED`.
     */
    state?: pulumi.Input<string>;
    /**
     * Interconnect IP of the DC within Tencent.
     */
    tencentAddress?: pulumi.Input<string>;
    /**
     * Vlan of the dedicated tunnels. Valid value ranges: (0~3000). `0` means that only one tunnel can be created for the physical connect.
     */
    vlan?: pulumi.Input<number>;
    /**
     * ID of the VPC or BMVPC.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Bandwidth of the DC.
     */
    bandwidth?: pulumi.Input<number>;
    /**
     * BGP ASN of the user. A required field within BGP.
     */
    bgpAsn?: pulumi.Input<number>;
    /**
     * BGP key of the user.
     */
    bgpAuthKey?: pulumi.Input<string>;
    /**
     * Interconnect IP of the DC within client.
     */
    customerAddress?: pulumi.Input<string>;
    /**
     * ID of the DC to be queried, application deployment offline.
     */
    dcId: pulumi.Input<string>;
    /**
     * Connection owner, who is the current customer by default. The developer account ID should be entered for shared connections.
     */
    dcOwnerAccount?: pulumi.Input<string>;
    /**
     * ID of the DC Gateway. Currently only new in the console.
     */
    dcgId: pulumi.Input<string>;
    /**
     * Name of the dedicated tunnel.
     */
    name?: pulumi.Input<string>;
    /**
     * Network region.
     */
    networkRegion?: pulumi.Input<string>;
    /**
     * Type of the network. Valid value: `VPC`, `BMVPC` and `CCN`. The default value is `VPC`.
     */
    networkType?: pulumi.Input<string>;
    /**
     * Static route, the network address of the user IDC. It can be modified after setting but cannot be deleted. AN unable field within BGP.
     */
    routeFilterPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Type of the route, and available values include BGP and STATIC. The default value is `BGP`.
     */
    routeType?: pulumi.Input<string>;
    /**
     * Interconnect IP of the DC within Tencent.
     */
    tencentAddress?: pulumi.Input<string>;
    /**
     * Vlan of the dedicated tunnels. Valid value ranges: (0~3000). `0` means that only one tunnel can be created for the physical connect.
     */
    vlan?: pulumi.Input<number>;
    /**
     * ID of the VPC or BMVPC.
     */
    vpcId?: pulumi.Input<string>;
}
