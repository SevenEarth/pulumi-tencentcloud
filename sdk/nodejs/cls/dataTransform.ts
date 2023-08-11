// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a cls dataTransform
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const logsetSrc = new tencentcloud.cls.Logset("logsetSrc", {
 *     logsetName: "tf-example-src",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const topicSrc = new tencentcloud.cls.Topic("topicSrc", {
 *     topicName: "tf-example_src",
 *     logsetId: logsetSrc.id,
 *     autoSplit: false,
 *     maxSplitPartitions: 20,
 *     partitionCount: 1,
 *     period: 10,
 *     storageType: "hot",
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const logsetDst = new tencentcloud.cls.Logset("logsetDst", {
 *     logsetName: "tf-example-dst",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * const topicDst = new tencentcloud.cls.Topic("topicDst", {
 *     topicName: "tf-example-dst",
 *     logsetId: logsetDst.id,
 *     autoSplit: false,
 *     maxSplitPartitions: 20,
 *     partitionCount: 1,
 *     period: 10,
 *     storageType: "hot",
 *     tags: {
 *         test: "test",
 *     },
 * });
 * const dataTransform = new tencentcloud.cls.DataTransform("dataTransform", {
 *     funcType: 1,
 *     srcTopicId: topicSrc.id,
 *     etlContent: "ext_sep(\"content\", \"f1, f2, f3\", sep=\",\", quote=\"\", restrict=False, mode=\"overwrite\")fields_drop(\"content\")",
 *     taskType: 3,
 *     enableFlag: 1,
 *     dstResources: [{
 *         topicId: topicDst.id,
 *         alias: "iac-test-dst",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * cls data_transform can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Cls/dataTransform:DataTransform data_transform data_transform_id
 * ```
 */
export class DataTransform extends pulumi.CustomResource {
    /**
     * Get an existing DataTransform resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataTransformState, opts?: pulumi.CustomResourceOptions): DataTransform {
        return new DataTransform(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Cls/dataTransform:DataTransform';

    /**
     * Returns true if the given object is an instance of DataTransform.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataTransform {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataTransform.__pulumiType;
    }

    /**
     * data transform des resources.
     */
    public readonly dstResources!: pulumi.Output<outputs.Cls.DataTransformDstResource[] | undefined>;
    /**
     * task enable flag.
     */
    public readonly enableFlag!: pulumi.Output<number | undefined>;
    /**
     * data transform content.
     */
    public readonly etlContent!: pulumi.Output<string>;
    /**
     * task type.
     */
    public readonly funcType!: pulumi.Output<number>;
    /**
     * task name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * src topic id.
     */
    public readonly srcTopicId!: pulumi.Output<string>;
    /**
     * task type.
     */
    public readonly taskType!: pulumi.Output<number>;

    /**
     * Create a DataTransform resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataTransformArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataTransformArgs | DataTransformState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataTransformState | undefined;
            resourceInputs["dstResources"] = state ? state.dstResources : undefined;
            resourceInputs["enableFlag"] = state ? state.enableFlag : undefined;
            resourceInputs["etlContent"] = state ? state.etlContent : undefined;
            resourceInputs["funcType"] = state ? state.funcType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["srcTopicId"] = state ? state.srcTopicId : undefined;
            resourceInputs["taskType"] = state ? state.taskType : undefined;
        } else {
            const args = argsOrState as DataTransformArgs | undefined;
            if ((!args || args.etlContent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'etlContent'");
            }
            if ((!args || args.funcType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'funcType'");
            }
            if ((!args || args.srcTopicId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'srcTopicId'");
            }
            if ((!args || args.taskType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskType'");
            }
            resourceInputs["dstResources"] = args ? args.dstResources : undefined;
            resourceInputs["enableFlag"] = args ? args.enableFlag : undefined;
            resourceInputs["etlContent"] = args ? args.etlContent : undefined;
            resourceInputs["funcType"] = args ? args.funcType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["srcTopicId"] = args ? args.srcTopicId : undefined;
            resourceInputs["taskType"] = args ? args.taskType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataTransform.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataTransform resources.
 */
export interface DataTransformState {
    /**
     * data transform des resources.
     */
    dstResources?: pulumi.Input<pulumi.Input<inputs.Cls.DataTransformDstResource>[]>;
    /**
     * task enable flag.
     */
    enableFlag?: pulumi.Input<number>;
    /**
     * data transform content.
     */
    etlContent?: pulumi.Input<string>;
    /**
     * task type.
     */
    funcType?: pulumi.Input<number>;
    /**
     * task name.
     */
    name?: pulumi.Input<string>;
    /**
     * src topic id.
     */
    srcTopicId?: pulumi.Input<string>;
    /**
     * task type.
     */
    taskType?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DataTransform resource.
 */
export interface DataTransformArgs {
    /**
     * data transform des resources.
     */
    dstResources?: pulumi.Input<pulumi.Input<inputs.Cls.DataTransformDstResource>[]>;
    /**
     * task enable flag.
     */
    enableFlag?: pulumi.Input<number>;
    /**
     * data transform content.
     */
    etlContent: pulumi.Input<string>;
    /**
     * task type.
     */
    funcType: pulumi.Input<number>;
    /**
     * task name.
     */
    name?: pulumi.Input<string>;
    /**
     * src topic id.
     */
    srcTopicId: pulumi.Input<string>;
    /**
     * task type.
     */
    taskType: pulumi.Input<number>;
}
