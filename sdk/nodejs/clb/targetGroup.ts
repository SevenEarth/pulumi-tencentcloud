// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a CLB target group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const test = new tencentcloud.Clb.TargetGroup("test", {
 *     port: 33,
 *     targetGroupName: "test",
 * });
 * ```
 *
 * ## Import
 *
 * CLB target group can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Clb/targetGroup:TargetGroup test lbtg-3k3io0i0
 * ```
 */
export class TargetGroup extends pulumi.CustomResource {
    /**
     * Get an existing TargetGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetGroupState, opts?: pulumi.CustomResourceOptions): TargetGroup {
        return new TargetGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Clb/targetGroup:TargetGroup';

    /**
     * Returns true if the given object is an instance of TargetGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetGroup.__pulumiType;
    }

    /**
     * The default port of target group, add server after can use it.
     */
    public readonly port!: pulumi.Output<number | undefined>;
    /**
     * It has been deprecated from version 1.77.3. please use `tencentcloud.Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
     *
     * @deprecated It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
     */
    public readonly targetGroupInstances!: pulumi.Output<outputs.Clb.TargetGroupTargetGroupInstance[] | undefined>;
    /**
     * Target group name.
     */
    public readonly targetGroupName!: pulumi.Output<string | undefined>;
    /**
     * VPC ID, default is based on the network.
     */
    public readonly vpcId!: pulumi.Output<string | undefined>;

    /**
     * Create a TargetGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TargetGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetGroupArgs | TargetGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TargetGroupState | undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["targetGroupInstances"] = state ? state.targetGroupInstances : undefined;
            resourceInputs["targetGroupName"] = state ? state.targetGroupName : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as TargetGroupArgs | undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["targetGroupInstances"] = args ? args.targetGroupInstances : undefined;
            resourceInputs["targetGroupName"] = args ? args.targetGroupName : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TargetGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetGroup resources.
 */
export interface TargetGroupState {
    /**
     * The default port of target group, add server after can use it.
     */
    port?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.77.3. please use `tencentcloud.Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
     *
     * @deprecated It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
     */
    targetGroupInstances?: pulumi.Input<pulumi.Input<inputs.Clb.TargetGroupTargetGroupInstance>[]>;
    /**
     * Target group name.
     */
    targetGroupName?: pulumi.Input<string>;
    /**
     * VPC ID, default is based on the network.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetGroup resource.
 */
export interface TargetGroupArgs {
    /**
     * The default port of target group, add server after can use it.
     */
    port?: pulumi.Input<number>;
    /**
     * It has been deprecated from version 1.77.3. please use `tencentcloud.Clb.TargetGroupInstanceAttachment` instead. The backend server of target group bind.
     *
     * @deprecated It has been deprecated from version 1.77.3. please use `tencentcloud_clb_target_group_instance_attachment` instead.
     */
    targetGroupInstances?: pulumi.Input<pulumi.Input<inputs.Clb.TargetGroupTargetGroupInstance>[]>;
    /**
     * Target group name.
     */
    targetGroupName?: pulumi.Input<string>;
    /**
     * VPC ID, default is based on the network.
     */
    vpcId?: pulumi.Input<string>;
}
