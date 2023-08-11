// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource for an AS (Auto scaling) policy.
 *
 * ## Example Usage
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
 * const exampleScalingPolicy = new tencentcloud.as.ScalingPolicy("exampleScalingPolicy", {
 *     scalingGroupId: exampleScalingGroup.id,
 *     policyName: "tf-as-scaling-policy",
 *     adjustmentType: "EXACT_CAPACITY",
 *     adjustmentValue: 0,
 *     comparisonOperator: "GREATER_THAN",
 *     metricName: "CPU_UTILIZATION",
 *     threshold: 80,
 *     period: 300,
 *     continuousTime: 10,
 *     statistic: "AVERAGE",
 *     cooldown: 360,
 * });
 * ```
 */
export class ScalingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ScalingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingPolicyState, opts?: pulumi.CustomResourceOptions): ScalingPolicy {
        return new ScalingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:As/scalingPolicy:ScalingPolicy';

    /**
     * Returns true if the given object is an instance of ScalingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingPolicy.__pulumiType;
    }

    /**
     * Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
     */
    public readonly adjustmentType!: pulumi.Output<string>;
    /**
     * Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
     */
    public readonly adjustmentValue!: pulumi.Output<number>;
    /**
     * Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
     */
    public readonly comparisonOperator!: pulumi.Output<string>;
    /**
     * Retry times. Valid value ranges: (1~10).
     */
    public readonly continuousTime!: pulumi.Output<number>;
    /**
     * Cooldwon time in second. Default is `30`0.
     */
    public readonly cooldown!: pulumi.Output<number | undefined>;
    /**
     * Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * An ID group of users to be notified when an alarm is triggered.
     */
    public readonly notificationUserGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * Time period in second. Valid values: `60` and `300`.
     */
    public readonly period!: pulumi.Output<number>;
    /**
     * Name of a policy used to define a reaction when an alarm is triggered.
     */
    public readonly policyName!: pulumi.Output<string>;
    /**
     * ID of a scaling group.
     */
    public readonly scalingGroupId!: pulumi.Output<string>;
    /**
     * Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
     */
    public readonly statistic!: pulumi.Output<string | undefined>;
    /**
     * Alarm threshold.
     */
    public readonly threshold!: pulumi.Output<number>;

    /**
     * Create a ScalingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingPolicyArgs | ScalingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScalingPolicyState | undefined;
            resourceInputs["adjustmentType"] = state ? state.adjustmentType : undefined;
            resourceInputs["adjustmentValue"] = state ? state.adjustmentValue : undefined;
            resourceInputs["comparisonOperator"] = state ? state.comparisonOperator : undefined;
            resourceInputs["continuousTime"] = state ? state.continuousTime : undefined;
            resourceInputs["cooldown"] = state ? state.cooldown : undefined;
            resourceInputs["metricName"] = state ? state.metricName : undefined;
            resourceInputs["notificationUserGroupIds"] = state ? state.notificationUserGroupIds : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["policyName"] = state ? state.policyName : undefined;
            resourceInputs["scalingGroupId"] = state ? state.scalingGroupId : undefined;
            resourceInputs["statistic"] = state ? state.statistic : undefined;
            resourceInputs["threshold"] = state ? state.threshold : undefined;
        } else {
            const args = argsOrState as ScalingPolicyArgs | undefined;
            if ((!args || args.adjustmentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adjustmentType'");
            }
            if ((!args || args.adjustmentValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adjustmentValue'");
            }
            if ((!args || args.comparisonOperator === undefined) && !opts.urn) {
                throw new Error("Missing required property 'comparisonOperator'");
            }
            if ((!args || args.continuousTime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'continuousTime'");
            }
            if ((!args || args.metricName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metricName'");
            }
            if ((!args || args.period === undefined) && !opts.urn) {
                throw new Error("Missing required property 'period'");
            }
            if ((!args || args.policyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyName'");
            }
            if ((!args || args.scalingGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scalingGroupId'");
            }
            if ((!args || args.threshold === undefined) && !opts.urn) {
                throw new Error("Missing required property 'threshold'");
            }
            resourceInputs["adjustmentType"] = args ? args.adjustmentType : undefined;
            resourceInputs["adjustmentValue"] = args ? args.adjustmentValue : undefined;
            resourceInputs["comparisonOperator"] = args ? args.comparisonOperator : undefined;
            resourceInputs["continuousTime"] = args ? args.continuousTime : undefined;
            resourceInputs["cooldown"] = args ? args.cooldown : undefined;
            resourceInputs["metricName"] = args ? args.metricName : undefined;
            resourceInputs["notificationUserGroupIds"] = args ? args.notificationUserGroupIds : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["policyName"] = args ? args.policyName : undefined;
            resourceInputs["scalingGroupId"] = args ? args.scalingGroupId : undefined;
            resourceInputs["statistic"] = args ? args.statistic : undefined;
            resourceInputs["threshold"] = args ? args.threshold : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ScalingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingPolicy resources.
 */
export interface ScalingPolicyState {
    /**
     * Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
     */
    adjustmentValue?: pulumi.Input<number>;
    /**
     * Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
     */
    comparisonOperator?: pulumi.Input<string>;
    /**
     * Retry times. Valid value ranges: (1~10).
     */
    continuousTime?: pulumi.Input<number>;
    /**
     * Cooldwon time in second. Default is `30`0.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
     */
    metricName?: pulumi.Input<string>;
    /**
     * An ID group of users to be notified when an alarm is triggered.
     */
    notificationUserGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Time period in second. Valid values: `60` and `300`.
     */
    period?: pulumi.Input<number>;
    /**
     * Name of a policy used to define a reaction when an alarm is triggered.
     */
    policyName?: pulumi.Input<string>;
    /**
     * ID of a scaling group.
     */
    scalingGroupId?: pulumi.Input<string>;
    /**
     * Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
     */
    statistic?: pulumi.Input<string>;
    /**
     * Alarm threshold.
     */
    threshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ScalingPolicy resource.
 */
export interface ScalingPolicyArgs {
    /**
     * Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values: `CHANGE_IN_CAPACITY`, `EXACT_CAPACITY` and `PERCENT_CHANGE_IN_CAPACITY`.
     */
    adjustmentType: pulumi.Input<string>;
    /**
     * Define the number of instances by which to scale.For `CHANGE_IN_CAPACITY` type or PERCENT_CHANGE_IN_CAPACITY, a positive increment adds to the current capacity and a negative value removes from the current capacity. For `EXACT_CAPACITY` type, it defines an absolute number of the existing Auto Scaling group size.
     */
    adjustmentValue: pulumi.Input<number>;
    /**
     * Comparison operator. Valid values: `GREATER_THAN`, `GREATER_THAN_OR_EQUAL_TO`, `LESS_THAN`, `LESS_THAN_OR_EQUAL_TO`, `EQUAL_TO` and `NOT_EQUAL_TO`.
     */
    comparisonOperator: pulumi.Input<string>;
    /**
     * Retry times. Valid value ranges: (1~10).
     */
    continuousTime: pulumi.Input<number>;
    /**
     * Cooldwon time in second. Default is `30`0.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * Name of an indicator. Valid values: `CPU_UTILIZATION`, `MEM_UTILIZATION`, `LAN_TRAFFIC_OUT`, `LAN_TRAFFIC_IN`, `WAN_TRAFFIC_OUT` and `WAN_TRAFFIC_IN`.
     */
    metricName: pulumi.Input<string>;
    /**
     * An ID group of users to be notified when an alarm is triggered.
     */
    notificationUserGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Time period in second. Valid values: `60` and `300`.
     */
    period: pulumi.Input<number>;
    /**
     * Name of a policy used to define a reaction when an alarm is triggered.
     */
    policyName: pulumi.Input<string>;
    /**
     * ID of a scaling group.
     */
    scalingGroupId: pulumi.Input<string>;
    /**
     * Statistic types. Valid values: `AVERAGE`, `MAXIMUM` and `MINIMUM`. Default is `AVERAGE`.
     */
    statistic?: pulumi.Input<string>;
    /**
     * Alarm threshold.
     */
    threshold: pulumi.Input<number>;
}
