// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to set as scalingGroup status
 *
 * ## Example Usage
 * ### Deactivate Scaling Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const zones = tencentcloud.Availability.getZonesByProduct({
 *     product: "as",
 * });
 * const image = tencentcloud.Images.getInstance({
 *     imageTypes: ["PUBLIC_IMAGE"],
 *     osName: "TencentOS Server 3.2 (Final)",
 * });
 * const vpc = new tencentcloud.vpc.Instance("vpc", {cidrBlock: "10.0.0.0/16"});
 * const subnet = new tencentcloud.subnet.Instance("subnet", {
 *     vpcId: vpc.id,
 *     cidrBlock: "10.0.0.0/16",
 *     availabilityZone: zones.then(zones => zones.zones?[0]?.name),
 * });
 * const exampleScalingConfig = new tencentcloud.as.ScalingConfig("exampleScalingConfig", {
 *     configurationName: "tf-example",
 *     imageId: image.then(image => image.images?[0]?.imageId),
 *     instanceTypes: [
 *         "SA1.SMALL1",
 *         "SA2.SMALL1",
 *         "SA2.SMALL2",
 *         "SA2.SMALL4",
 *     ],
 *     instanceNameSettings: {
 *         instanceName: "test-ins-name",
 *     },
 * });
 * const exampleScalingGroup = new tencentcloud.as.ScalingGroup("exampleScalingGroup", {
 *     scalingGroupName: "tf-example",
 *     configurationId: exampleScalingConfig.id,
 *     maxSize: 1,
 *     minSize: 0,
 *     vpcId: vpc.id,
 *     subnetIds: [subnet.id],
 * });
 * const scalingGroupStatus = new tencentcloud.as.ScalingGroupStatus("scalingGroupStatus", {
 *     autoScalingGroupId: exampleScalingGroup.id,
 *     enable: false,
 * });
 * ```
 * ### Enable Scaling Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const scalingGroupStatus = new tencentcloud.as.ScalingGroupStatus("scalingGroupStatus", {
 *     autoScalingGroupId: tencentcloud_as_scaling_group.example.id,
 *     enable: true,
 * });
 * ```
 *
 * ## Import
 *
 * as scaling_group_status can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:As/scalingGroupStatus:ScalingGroupStatus scaling_group_status scaling_group_id
 * ```
 */
export class ScalingGroupStatus extends pulumi.CustomResource {
    /**
     * Get an existing ScalingGroupStatus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingGroupStatusState, opts?: pulumi.CustomResourceOptions): ScalingGroupStatus {
        return new ScalingGroupStatus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:As/scalingGroupStatus:ScalingGroupStatus';

    /**
     * Returns true if the given object is an instance of ScalingGroupStatus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingGroupStatus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingGroupStatus.__pulumiType;
    }

    /**
     * Scaling group ID.
     */
    public readonly autoScalingGroupId!: pulumi.Output<string>;
    /**
     * If enable auto scaling group.
     */
    public readonly enable!: pulumi.Output<boolean>;

    /**
     * Create a ScalingGroupStatus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingGroupStatusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingGroupStatusArgs | ScalingGroupStatusState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScalingGroupStatusState | undefined;
            resourceInputs["autoScalingGroupId"] = state ? state.autoScalingGroupId : undefined;
            resourceInputs["enable"] = state ? state.enable : undefined;
        } else {
            const args = argsOrState as ScalingGroupStatusArgs | undefined;
            if ((!args || args.autoScalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoScalingGroupId'");
            }
            if ((!args || args.enable === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enable'");
            }
            resourceInputs["autoScalingGroupId"] = args ? args.autoScalingGroupId : undefined;
            resourceInputs["enable"] = args ? args.enable : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScalingGroupStatus.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingGroupStatus resources.
 */
export interface ScalingGroupStatusState {
    /**
     * Scaling group ID.
     */
    autoScalingGroupId?: pulumi.Input<string>;
    /**
     * If enable auto scaling group.
     */
    enable?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a ScalingGroupStatus resource.
 */
export interface ScalingGroupStatusArgs {
    /**
     * Scaling group ID.
     */
    autoScalingGroupId: pulumi.Input<string>;
    /**
     * If enable auto scaling group.
     */
    enable: pulumi.Input<boolean>;
}
