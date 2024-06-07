// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tsf lane
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const lane = new tencentcloud.tsf.Lane("lane", {
 *     laneGroupLists: [{
 *         entrance: true,
 *         groupId: "group-yn7j5l8a",
 *     }],
 *     laneName: "lane-name-1",
 *     remark: "lane desc1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Lane extends pulumi.CustomResource {
    /**
     * Get an existing Lane resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LaneState, opts?: pulumi.CustomResourceOptions): Lane {
        return new Lane(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tsf/lane:Lane';

    /**
     * Returns true if the given object is an instance of Lane.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Lane {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Lane.__pulumiType;
    }

    /**
     * creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<number>;
    /**
     * Whether to enter the application.
     */
    public /*out*/ readonly entrance!: pulumi.Output<boolean>;
    /**
     * Swimlane Deployment Group Information.
     */
    public readonly laneGroupLists!: pulumi.Output<outputs.Tsf.LaneLaneGroupList[]>;
    /**
     * Lane ID.
     */
    public /*out*/ readonly laneId!: pulumi.Output<string>;
    /**
     * Lane name.
     */
    public readonly laneName!: pulumi.Output<string>;
    /**
     * A list of namespaces to which the swimlane has associated deployment groups.
     */
    public /*out*/ readonly namespaceIdLists!: pulumi.Output<string[]>;
    /**
     * Program id list.
     */
    public readonly programIdLists!: pulumi.Output<string[] | undefined>;
    /**
     * Lane Remarks.
     */
    public readonly remark!: pulumi.Output<string>;
    /**
     * update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<number>;

    /**
     * Create a Lane resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LaneArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LaneArgs | LaneState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LaneState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["entrance"] = state ? state.entrance : undefined;
            resourceInputs["laneGroupLists"] = state ? state.laneGroupLists : undefined;
            resourceInputs["laneId"] = state ? state.laneId : undefined;
            resourceInputs["laneName"] = state ? state.laneName : undefined;
            resourceInputs["namespaceIdLists"] = state ? state.namespaceIdLists : undefined;
            resourceInputs["programIdLists"] = state ? state.programIdLists : undefined;
            resourceInputs["remark"] = state ? state.remark : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as LaneArgs | undefined;
            if ((!args || args.laneGroupLists === undefined) && !opts.urn) {
                throw new Error("Missing required property 'laneGroupLists'");
            }
            if ((!args || args.laneName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'laneName'");
            }
            if ((!args || args.remark === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remark'");
            }
            resourceInputs["laneGroupLists"] = args ? args.laneGroupLists : undefined;
            resourceInputs["laneName"] = args ? args.laneName : undefined;
            resourceInputs["programIdLists"] = args ? args.programIdLists : undefined;
            resourceInputs["remark"] = args ? args.remark : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["entrance"] = undefined /*out*/;
            resourceInputs["laneId"] = undefined /*out*/;
            resourceInputs["namespaceIdLists"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Lane.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Lane resources.
 */
export interface LaneState {
    /**
     * creation time.
     */
    createTime?: pulumi.Input<number>;
    /**
     * Whether to enter the application.
     */
    entrance?: pulumi.Input<boolean>;
    /**
     * Swimlane Deployment Group Information.
     */
    laneGroupLists?: pulumi.Input<pulumi.Input<inputs.Tsf.LaneLaneGroupList>[]>;
    /**
     * Lane ID.
     */
    laneId?: pulumi.Input<string>;
    /**
     * Lane name.
     */
    laneName?: pulumi.Input<string>;
    /**
     * A list of namespaces to which the swimlane has associated deployment groups.
     */
    namespaceIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Program id list.
     */
    programIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lane Remarks.
     */
    remark?: pulumi.Input<string>;
    /**
     * update time.
     */
    updateTime?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Lane resource.
 */
export interface LaneArgs {
    /**
     * Swimlane Deployment Group Information.
     */
    laneGroupLists: pulumi.Input<pulumi.Input<inputs.Tsf.LaneLaneGroupList>[]>;
    /**
     * Lane name.
     */
    laneName: pulumi.Input<string>;
    /**
     * Program id list.
     */
    programIdLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lane Remarks.
     */
    remark: pulumi.Input<string>;
}
