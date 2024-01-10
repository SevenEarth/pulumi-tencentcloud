// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a gaap proxy group
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const proxyGroup = new tencentcloud.Gaap.ProxyGroup("proxy_group", {
 *     groupName: "tf-test-update",
 *     ipAddressVersion: "IPv4",
 *     packageType: "Thunder",
 *     projectId: 0,
 *     realServerRegion: "Beijing",
 * });
 * ```
 *
 * ## Import
 *
 * gaap proxy_group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Gaap/proxyGroup:ProxyGroup proxy_group proxy_group_id
 * ```
 */
export class ProxyGroup extends pulumi.CustomResource {
    /**
     * Get an existing ProxyGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProxyGroupState, opts?: pulumi.CustomResourceOptions): ProxyGroup {
        return new ProxyGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Gaap/proxyGroup:ProxyGroup';

    /**
     * Returns true if the given object is an instance of ProxyGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProxyGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProxyGroup.__pulumiType;
    }

    /**
     * Channel group alias.
     */
    public readonly groupName!: pulumi.Output<string>;
    /**
     * IP version, can be taken as IPv4 or IPv6 with a default value of IPv4.
     */
    public readonly ipAddressVersion!: pulumi.Output<string | undefined>;
    /**
     * Package type of channel group. Available values: Thunder and Accelerator. Default is Thunder.
     */
    public readonly packageType!: pulumi.Output<string | undefined>;
    /**
     * ID of the project to which the proxy group belongs.
     */
    public readonly projectId!: pulumi.Output<number>;
    /**
     * real server region, refer to the interface DescribeDestRegions to return the RegionId in the parameter RegionDetail.
     */
    public readonly realServerRegion!: pulumi.Output<string>;

    /**
     * Create a ProxyGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProxyGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProxyGroupArgs | ProxyGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProxyGroupState | undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["ipAddressVersion"] = state ? state.ipAddressVersion : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["realServerRegion"] = state ? state.realServerRegion : undefined;
        } else {
            const args = argsOrState as ProxyGroupArgs | undefined;
            if ((!args || args.groupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupName'");
            }
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            if ((!args || args.realServerRegion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realServerRegion'");
            }
            resourceInputs["groupName"] = args ? args.groupName : undefined;
            resourceInputs["ipAddressVersion"] = args ? args.ipAddressVersion : undefined;
            resourceInputs["packageType"] = args ? args.packageType : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["realServerRegion"] = args ? args.realServerRegion : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProxyGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProxyGroup resources.
 */
export interface ProxyGroupState {
    /**
     * Channel group alias.
     */
    groupName?: pulumi.Input<string>;
    /**
     * IP version, can be taken as IPv4 or IPv6 with a default value of IPv4.
     */
    ipAddressVersion?: pulumi.Input<string>;
    /**
     * Package type of channel group. Available values: Thunder and Accelerator. Default is Thunder.
     */
    packageType?: pulumi.Input<string>;
    /**
     * ID of the project to which the proxy group belongs.
     */
    projectId?: pulumi.Input<number>;
    /**
     * real server region, refer to the interface DescribeDestRegions to return the RegionId in the parameter RegionDetail.
     */
    realServerRegion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProxyGroup resource.
 */
export interface ProxyGroupArgs {
    /**
     * Channel group alias.
     */
    groupName: pulumi.Input<string>;
    /**
     * IP version, can be taken as IPv4 or IPv6 with a default value of IPv4.
     */
    ipAddressVersion?: pulumi.Input<string>;
    /**
     * Package type of channel group. Available values: Thunder and Accelerator. Default is Thunder.
     */
    packageType?: pulumi.Input<string>;
    /**
     * ID of the project to which the proxy group belongs.
     */
    projectId: pulumi.Input<number>;
    /**
     * real server region, refer to the interface DescribeDestRegions to return the RegionId in the parameter RegionDetail.
     */
    realServerRegion: pulumi.Input<string>;
}
