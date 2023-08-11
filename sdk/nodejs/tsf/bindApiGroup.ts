// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tsf bindApiGroup
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const bindApiGroup = new tencentcloud.Tsf.BindApiGroup("bind_api_group", {
 *     gatewayDeployGroupId: "group-vzd97zpy",
 *     groupId: "grp-qp0rj3zi",
 * });
 * ```
 *
 * ## Import
 *
 * tsf bind_api_group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tsf/bindApiGroup:BindApiGroup bind_api_group bind_api_group_id
 * ```
 */
export class BindApiGroup extends pulumi.CustomResource {
    /**
     * Get an existing BindApiGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BindApiGroupState, opts?: pulumi.CustomResourceOptions): BindApiGroup {
        return new BindApiGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tsf/bindApiGroup:BindApiGroup';

    /**
     * Returns true if the given object is an instance of BindApiGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BindApiGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BindApiGroup.__pulumiType;
    }

    /**
     * gateway group id.
     */
    public readonly gatewayDeployGroupId!: pulumi.Output<string>;
    /**
     * group id.
     */
    public readonly groupId!: pulumi.Output<string>;

    /**
     * Create a BindApiGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BindApiGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BindApiGroupArgs | BindApiGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BindApiGroupState | undefined;
            resourceInputs["gatewayDeployGroupId"] = state ? state.gatewayDeployGroupId : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
        } else {
            const args = argsOrState as BindApiGroupArgs | undefined;
            if ((!args || args.gatewayDeployGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gatewayDeployGroupId'");
            }
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            resourceInputs["gatewayDeployGroupId"] = args ? args.gatewayDeployGroupId : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BindApiGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BindApiGroup resources.
 */
export interface BindApiGroupState {
    /**
     * gateway group id.
     */
    gatewayDeployGroupId?: pulumi.Input<string>;
    /**
     * group id.
     */
    groupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BindApiGroup resource.
 */
export interface BindApiGroupArgs {
    /**
     * gateway group id.
     */
    gatewayDeployGroupId: pulumi.Input<string>;
    /**
     * group id.
     */
    groupId: pulumi.Input<string>;
}