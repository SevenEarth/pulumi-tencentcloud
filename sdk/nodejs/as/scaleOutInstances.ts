// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a as scaleOutInstances
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
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
 *     availabilityZone: zones.then(zones => zones.zones?.[0]?.name),
 * });
 * const exampleScalingConfig = new tencentcloud.as.ScalingConfig("exampleScalingConfig", {
 *     configurationName: "tf-example",
 *     imageId: image.then(image => image.images?.[0]?.imageId),
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
 *     maxSize: 4,
 *     minSize: 0,
 *     desiredCapacity: 2,
 *     vpcId: vpc.id,
 *     subnetIds: [subnet.id],
 * });
 * const scaleOutInstances = new tencentcloud.as.ScaleOutInstances("scaleOutInstances", {
 *     autoScalingGroupId: exampleScalingGroup.id,
 *     scaleOutNumber: 2,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * as scale_out_instances can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:As/scaleOutInstances:ScaleOutInstances scale_out_instances scale_out_instances_id
 * ```
 */
export class ScaleOutInstances extends pulumi.CustomResource {
    /**
     * Get an existing ScaleOutInstances resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScaleOutInstancesState, opts?: pulumi.CustomResourceOptions): ScaleOutInstances {
        return new ScaleOutInstances(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:As/scaleOutInstances:ScaleOutInstances';

    /**
     * Returns true if the given object is an instance of ScaleOutInstances.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScaleOutInstances {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScaleOutInstances.__pulumiType;
    }

    /**
     * Scaling group ID.
     */
    public readonly autoScalingGroupId!: pulumi.Output<string>;
    /**
     * Number of instances to be added.
     */
    public readonly scaleOutNumber!: pulumi.Output<number>;

    /**
     * Create a ScaleOutInstances resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScaleOutInstancesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScaleOutInstancesArgs | ScaleOutInstancesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScaleOutInstancesState | undefined;
            resourceInputs["autoScalingGroupId"] = state ? state.autoScalingGroupId : undefined;
            resourceInputs["scaleOutNumber"] = state ? state.scaleOutNumber : undefined;
        } else {
            const args = argsOrState as ScaleOutInstancesArgs | undefined;
            if ((!args || args.autoScalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoScalingGroupId'");
            }
            if ((!args || args.scaleOutNumber === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scaleOutNumber'");
            }
            resourceInputs["autoScalingGroupId"] = args ? args.autoScalingGroupId : undefined;
            resourceInputs["scaleOutNumber"] = args ? args.scaleOutNumber : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScaleOutInstances.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScaleOutInstances resources.
 */
export interface ScaleOutInstancesState {
    /**
     * Scaling group ID.
     */
    autoScalingGroupId?: pulumi.Input<string>;
    /**
     * Number of instances to be added.
     */
    scaleOutNumber?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ScaleOutInstances resource.
 */
export interface ScaleOutInstancesArgs {
    /**
     * Scaling group ID.
     */
    autoScalingGroupId: pulumi.Input<string>;
    /**
     * Number of instances to be added.
     */
    scaleOutNumber: pulumi.Input<number>;
}
