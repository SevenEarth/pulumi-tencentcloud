// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an entry of a routing table.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const config = new pulumi.Config();
 * const availabilityZone = config.get("availabilityZone") || "na-siliconvalley-1";
 * const fooInstance = new tencentcloud.vpc.Instance("fooInstance", {cidrBlock: "10.0.0.0/16"});
 * const fooTable = new tencentcloud.route.Table("fooTable", {vpcId: fooInstance.id});
 * const fooSubnet_instanceInstance = new tencentcloud.subnet.Instance("fooSubnet/instanceInstance", {
 *     vpcId: fooInstance.id,
 *     cidrBlock: "10.0.12.0/24",
 *     availabilityZone: availabilityZone,
 *     routeTableId: fooTable.id,
 * });
 * const instance = new tencentcloud.route.TableEntry("instance", {
 *     routeTableId: fooTable.id,
 *     destinationCidrBlock: "10.4.4.0/24",
 *     nextType: "EIP",
 *     nextHub: "0",
 *     description: "ci-test-route-table-entry",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Route table entry can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Route/tableEntry:TableEntry foo 83517.rtb-mlhpg09u
 * ```
 */
export class TableEntry extends pulumi.CustomResource {
    /**
     * Get an existing TableEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableEntryState, opts?: pulumi.CustomResourceOptions): TableEntry {
        return new TableEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Route/tableEntry:TableEntry';

    /**
     * Returns true if the given object is an instance of TableEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TableEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TableEntry.__pulumiType;
    }

    /**
     * Description of the routing table entry.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Destination address block.
     */
    public readonly destinationCidrBlock!: pulumi.Output<string>;
    /**
     * Whether the entry is disabled, default is `false`.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * ID of next-hop gateway. Note: when `nextType` is EIP, `nextHub` should be `0`.
     */
    public readonly nextHub!: pulumi.Output<string>;
    /**
     * Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
     */
    public readonly nextType!: pulumi.Output<string>;
    /**
     * ID of routing table to which this entry belongs.
     */
    public readonly routeTableId!: pulumi.Output<string>;

    /**
     * Create a TableEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableEntryArgs | TableEntryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableEntryState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationCidrBlock"] = state ? state.destinationCidrBlock : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["nextHub"] = state ? state.nextHub : undefined;
            resourceInputs["nextType"] = state ? state.nextType : undefined;
            resourceInputs["routeTableId"] = state ? state.routeTableId : undefined;
        } else {
            const args = argsOrState as TableEntryArgs | undefined;
            if ((!args || args.destinationCidrBlock === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationCidrBlock'");
            }
            if ((!args || args.nextHub === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nextHub'");
            }
            if ((!args || args.nextType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nextType'");
            }
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["nextHub"] = args ? args.nextHub : undefined;
            resourceInputs["nextType"] = args ? args.nextType : undefined;
            resourceInputs["routeTableId"] = args ? args.routeTableId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TableEntry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TableEntry resources.
 */
export interface TableEntryState {
    /**
     * Description of the routing table entry.
     */
    description?: pulumi.Input<string>;
    /**
     * Destination address block.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * Whether the entry is disabled, default is `false`.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * ID of next-hop gateway. Note: when `nextType` is EIP, `nextHub` should be `0`.
     */
    nextHub?: pulumi.Input<string>;
    /**
     * Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
     */
    nextType?: pulumi.Input<string>;
    /**
     * ID of routing table to which this entry belongs.
     */
    routeTableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TableEntry resource.
 */
export interface TableEntryArgs {
    /**
     * Description of the routing table entry.
     */
    description?: pulumi.Input<string>;
    /**
     * Destination address block.
     */
    destinationCidrBlock: pulumi.Input<string>;
    /**
     * Whether the entry is disabled, default is `false`.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * ID of next-hop gateway. Note: when `nextType` is EIP, `nextHub` should be `0`.
     */
    nextHub: pulumi.Input<string>;
    /**
     * Type of next-hop. Valid values: `CVM`, `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP` and `LOCAL_GATEWAY`.
     */
    nextType: pulumi.Input<string>;
    /**
     * ID of routing table to which this entry belongs.
     */
    routeTableId: pulumi.Input<string>;
}
