// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a chdfs accessGroup
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const accessGroup = new tencentcloud.chdfs.AccessGroup("accessGroup", {
 *     accessGroupName: "testAccessGroup",
 *     description: "test access group",
 *     vpcId: "vpc-4owdpnwr",
 *     vpcType: 1,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * chdfs access_group can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Chdfs/accessGroup:AccessGroup access_group access_group_id
 * ```
 */
export class AccessGroup extends pulumi.CustomResource {
    /**
     * Get an existing AccessGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccessGroupState, opts?: pulumi.CustomResourceOptions): AccessGroup {
        return new AccessGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Chdfs/accessGroup:AccessGroup';

    /**
     * Returns true if the given object is an instance of AccessGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessGroup.__pulumiType;
    }

    /**
     * Permission group name.
     */
    public readonly accessGroupName!: pulumi.Output<string>;
    /**
     * Permission group description, default empty.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * vpc network type(1:CVM, 2:BM 1.0).
     */
    public readonly vpcType!: pulumi.Output<number>;

    /**
     * Create a AccessGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccessGroupArgs | AccessGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccessGroupState | undefined;
            resourceInputs["accessGroupName"] = state ? state.accessGroupName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vpcType"] = state ? state.vpcType : undefined;
        } else {
            const args = argsOrState as AccessGroupArgs | undefined;
            if ((!args || args.accessGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessGroupName'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            if ((!args || args.vpcType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcType'");
            }
            resourceInputs["accessGroupName"] = args ? args.accessGroupName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vpcType"] = args ? args.vpcType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccessGroup resources.
 */
export interface AccessGroupState {
    /**
     * Permission group name.
     */
    accessGroupName?: pulumi.Input<string>;
    /**
     * Permission group description, default empty.
     */
    description?: pulumi.Input<string>;
    /**
     * VPC ID.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * vpc network type(1:CVM, 2:BM 1.0).
     */
    vpcType?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AccessGroup resource.
 */
export interface AccessGroupArgs {
    /**
     * Permission group name.
     */
    accessGroupName: pulumi.Input<string>;
    /**
     * Permission group description, default empty.
     */
    description?: pulumi.Input<string>;
    /**
     * VPC ID.
     */
    vpcId: pulumi.Input<string>;
    /**
     * vpc network type(1:CVM, 2:BM 1.0).
     */
    vpcType: pulumi.Input<number>;
}
