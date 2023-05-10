// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a vpc endPointServiceWhiteList
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const endPointServiceWhiteList = new tencentcloud.Vpc.EndPointServiceWhiteList("end_point_service_white_list", {
 *     description: "terraform for test",
 *     endPointServiceId: "vpcsvc-69y13tdb",
 *     userUin: "100020512675",
 * });
 * ```
 *
 * ## Import
 *
 * vpc end_point_service_white_list can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList end_point_service_white_list end_point_service_white_list_id
 * ```
 */
export class EndPointServiceWhiteList extends pulumi.CustomResource {
    /**
     * Get an existing EndPointServiceWhiteList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndPointServiceWhiteListState, opts?: pulumi.CustomResourceOptions): EndPointServiceWhiteList {
        return new EndPointServiceWhiteList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Vpc/endPointServiceWhiteList:EndPointServiceWhiteList';

    /**
     * Returns true if the given object is an instance of EndPointServiceWhiteList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndPointServiceWhiteList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndPointServiceWhiteList.__pulumiType;
    }

    /**
     * Create Time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of white list.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * ID of endpoint service.
     */
    public readonly endPointServiceId!: pulumi.Output<string>;
    /**
     * APPID.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * UIN.
     */
    public readonly userUin!: pulumi.Output<string>;

    /**
     * Create a EndPointServiceWhiteList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndPointServiceWhiteListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndPointServiceWhiteListArgs | EndPointServiceWhiteListState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndPointServiceWhiteListState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["endPointServiceId"] = state ? state.endPointServiceId : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["userUin"] = state ? state.userUin : undefined;
        } else {
            const args = argsOrState as EndPointServiceWhiteListArgs | undefined;
            if ((!args || args.endPointServiceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endPointServiceId'");
            }
            if ((!args || args.userUin === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userUin'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endPointServiceId"] = args ? args.endPointServiceId : undefined;
            resourceInputs["userUin"] = args ? args.userUin : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndPointServiceWhiteList.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndPointServiceWhiteList resources.
 */
export interface EndPointServiceWhiteListState {
    /**
     * Create Time.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of white list.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of endpoint service.
     */
    endPointServiceId?: pulumi.Input<string>;
    /**
     * APPID.
     */
    owner?: pulumi.Input<string>;
    /**
     * UIN.
     */
    userUin?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EndPointServiceWhiteList resource.
 */
export interface EndPointServiceWhiteListArgs {
    /**
     * Description of white list.
     */
    description?: pulumi.Input<string>;
    /**
     * ID of endpoint service.
     */
    endPointServiceId: pulumi.Input<string>;
    /**
     * UIN.
     */
    userUin: pulumi.Input<string>;
}
