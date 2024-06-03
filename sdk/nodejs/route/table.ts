// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a VPC routing table.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const fooInstance = new tencentcloud.vpc.Instance("fooInstance", {cidrBlock: "10.0.0.0/16"});
 * const fooTable = new tencentcloud.route.Table("fooTable", {vpcId: fooInstance.id});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Vpc routetable instance can be imported, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Route/table:Table test route_table_id
 * ```
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Route/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * Creation time of the routing table.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Indicates whether it is the default routing table.
     */
    public /*out*/ readonly isDefault!: pulumi.Output<boolean>;
    /**
     * The name of routing table.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID list of the routing entries.
     */
    public /*out*/ readonly routeEntryIds!: pulumi.Output<string[]>;
    /**
     * ID list of the subnets associated with this route table.
     */
    public /*out*/ readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * The tags of routing table.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * ID of VPC to which the route table should be associated.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["isDefault"] = state ? state.isDefault : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["routeEntryIds"] = state ? state.routeEntryIds : undefined;
            resourceInputs["subnetIds"] = state ? state.subnetIds : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["isDefault"] = undefined /*out*/;
            resourceInputs["routeEntryIds"] = undefined /*out*/;
            resourceInputs["subnetIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * Creation time of the routing table.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Indicates whether it is the default routing table.
     */
    isDefault?: pulumi.Input<boolean>;
    /**
     * The name of routing table.
     */
    name?: pulumi.Input<string>;
    /**
     * ID list of the routing entries.
     */
    routeEntryIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID list of the subnets associated with this route table.
     */
    subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The tags of routing table.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of VPC to which the route table should be associated.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * The name of routing table.
     */
    name?: pulumi.Input<string>;
    /**
     * The tags of routing table.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * ID of VPC to which the route table should be associated.
     */
    vpcId: pulumi.Input<string>;
}
