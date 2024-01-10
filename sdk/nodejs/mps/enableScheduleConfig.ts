// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a mps enableScheduleConfig
 *
 * ## Example Usage
 * ### Enable the mps schedule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const object = tencentcloud.Cos.getBucketObject({
 *     bucket: `keep-bucket-${local.app_id}`,
 *     key: "/mps-test/test.mov",
 * });
 * const output = new tencentcloud.cos.Bucket("output", {
 *     bucket: `tf-bucket-mps-schedule-config-output1-${local.app_id}`,
 *     forceClean: true,
 *     acl: "public-read",
 * });
 * const example = new tencentcloud.mps.Schedule("example", {
 *     scheduleName: "tf_test_mps_schedule_config",
 *     trigger: {
 *         type: "CosFileUpload",
 *         cosFileUploadTrigger: {
 *             bucket: object.then(object => object.bucket),
 *             region: `%s`,
 *             dir: "/upload/",
 *             formats: [
 *                 "flv",
 *                 "mov",
 *             ],
 *         },
 *     },
 *     activities: [
 *         {
 *             activityType: "input",
 *             reardriveIndices: [
 *                 1,
 *                 2,
 *             ],
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [3],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [
 *                 6,
 *                 7,
 *             ],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [
 *                 4,
 *                 5,
 *             ],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [10],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [10],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [10],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [8],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [9],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "action-trans",
 *             reardriveIndices: [10],
 *             activityPara: {
 *                 transcodeTask: {
 *                     definition: 10,
 *                 },
 *             },
 *         },
 *         {
 *             activityType: "output",
 *         },
 *     ],
 *     outputStorage: {
 *         type: "COS",
 *         cosOutputStorage: {
 *             bucket: output.bucket,
 *             region: `%s`,
 *         },
 *     },
 *     outputDir: "output/",
 * });
 * const config = new tencentcloud.mps.EnableScheduleConfig("config", {
 *     scheduleId: example.id,
 *     enabled: true,
 * });
 * ```
 * ### Disable the mps schedule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const config = new tencentcloud.mps.EnableScheduleConfig("config", {
 *     scheduleId: tencentcloud_mps_schedule.example.id,
 *     enabled: false,
 * });
 * ```
 *
 * ## Import
 *
 * mps enable_schedule_config can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Mps/enableScheduleConfig:EnableScheduleConfig enable_schedule_config enable_schedule_config_id
 * ```
 */
export class EnableScheduleConfig extends pulumi.CustomResource {
    /**
     * Get an existing EnableScheduleConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EnableScheduleConfigState, opts?: pulumi.CustomResourceOptions): EnableScheduleConfig {
        return new EnableScheduleConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Mps/enableScheduleConfig:EnableScheduleConfig';

    /**
     * Returns true if the given object is an instance of EnableScheduleConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EnableScheduleConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EnableScheduleConfig.__pulumiType;
    }

    /**
     * true: enable; false: disable.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The scheme ID.
     */
    public readonly scheduleId!: pulumi.Output<number>;

    /**
     * Create a EnableScheduleConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnableScheduleConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnableScheduleConfigArgs | EnableScheduleConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EnableScheduleConfigState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["scheduleId"] = state ? state.scheduleId : undefined;
        } else {
            const args = argsOrState as EnableScheduleConfigArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.scheduleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scheduleId'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["scheduleId"] = args ? args.scheduleId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EnableScheduleConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EnableScheduleConfig resources.
 */
export interface EnableScheduleConfigState {
    /**
     * true: enable; false: disable.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The scheme ID.
     */
    scheduleId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a EnableScheduleConfig resource.
 */
export interface EnableScheduleConfigArgs {
    /**
     * true: enable; false: disable.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The scheme ID.
     */
    scheduleId: pulumi.Input<number>;
}
