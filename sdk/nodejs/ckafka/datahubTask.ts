// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ckafka datahubTask
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const datahubTask = new tencentcloud.Ckafka.DatahubTask("datahub_task", {
 *     sourceResource: {
 *         postgreSqlParam: {
 *             database: "postgres",
 *             isTableRegular: false,
 *             keyColumns: "",
 *             pluginName: "decoderbufs",
 *             recordWithSchema: false,
 *             resource: "resource-y9nxnw46",
 *             snapshotMode: "never",
 *             table: "*",
 *         },
 *         type: "POSTGRESQL",
 *     },
 *     targetResource: {
 *         topicParam: {
 *             compressionType: "none",
 *             resource: "1308726196-keep-topic",
 *             useAutoCreateTopic: false,
 *         },
 *         type: "TOPIC",
 *     },
 *     taskName: "test-task123321",
 *     taskType: "SOURCE",
 * });
 * ```
 *
 * ## Import
 *
 * ckafka datahub_task can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ckafka/datahubTask:DatahubTask datahub_task datahub_task_id
 * ```
 */
export class DatahubTask extends pulumi.CustomResource {
    /**
     * Get an existing DatahubTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatahubTaskState, opts?: pulumi.CustomResourceOptions): DatahubTask {
        return new DatahubTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ckafka/datahubTask:DatahubTask';

    /**
     * Returns true if the given object is an instance of DatahubTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatahubTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatahubTask.__pulumiType;
    }

    /**
     * SchemaId.
     */
    public readonly schemaId!: pulumi.Output<string | undefined>;
    /**
     * data resource.
     */
    public readonly sourceResource!: pulumi.Output<outputs.Ckafka.DatahubTaskSourceResource | undefined>;
    /**
     * Target Resource.
     */
    public readonly targetResource!: pulumi.Output<outputs.Ckafka.DatahubTaskTargetResource | undefined>;
    /**
     * name of the task.
     */
    public readonly taskName!: pulumi.Output<string>;
    /**
     * type of the task, SOURCE(data input), SINK(data output).
     */
    public readonly taskType!: pulumi.Output<string>;
    /**
     * Data Processing Rules.
     */
    public readonly transformParam!: pulumi.Output<outputs.Ckafka.DatahubTaskTransformParam | undefined>;
    /**
     * Data processing rules.
     */
    public readonly transformsParam!: pulumi.Output<outputs.Ckafka.DatahubTaskTransformsParam | undefined>;

    /**
     * Create a DatahubTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatahubTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatahubTaskArgs | DatahubTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatahubTaskState | undefined;
            resourceInputs["schemaId"] = state ? state.schemaId : undefined;
            resourceInputs["sourceResource"] = state ? state.sourceResource : undefined;
            resourceInputs["targetResource"] = state ? state.targetResource : undefined;
            resourceInputs["taskName"] = state ? state.taskName : undefined;
            resourceInputs["taskType"] = state ? state.taskType : undefined;
            resourceInputs["transformParam"] = state ? state.transformParam : undefined;
            resourceInputs["transformsParam"] = state ? state.transformsParam : undefined;
        } else {
            const args = argsOrState as DatahubTaskArgs | undefined;
            if ((!args || args.taskName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskName'");
            }
            if ((!args || args.taskType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taskType'");
            }
            resourceInputs["schemaId"] = args ? args.schemaId : undefined;
            resourceInputs["sourceResource"] = args ? args.sourceResource : undefined;
            resourceInputs["targetResource"] = args ? args.targetResource : undefined;
            resourceInputs["taskName"] = args ? args.taskName : undefined;
            resourceInputs["taskType"] = args ? args.taskType : undefined;
            resourceInputs["transformParam"] = args ? args.transformParam : undefined;
            resourceInputs["transformsParam"] = args ? args.transformsParam : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatahubTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatahubTask resources.
 */
export interface DatahubTaskState {
    /**
     * SchemaId.
     */
    schemaId?: pulumi.Input<string>;
    /**
     * data resource.
     */
    sourceResource?: pulumi.Input<inputs.Ckafka.DatahubTaskSourceResource>;
    /**
     * Target Resource.
     */
    targetResource?: pulumi.Input<inputs.Ckafka.DatahubTaskTargetResource>;
    /**
     * name of the task.
     */
    taskName?: pulumi.Input<string>;
    /**
     * type of the task, SOURCE(data input), SINK(data output).
     */
    taskType?: pulumi.Input<string>;
    /**
     * Data Processing Rules.
     */
    transformParam?: pulumi.Input<inputs.Ckafka.DatahubTaskTransformParam>;
    /**
     * Data processing rules.
     */
    transformsParam?: pulumi.Input<inputs.Ckafka.DatahubTaskTransformsParam>;
}

/**
 * The set of arguments for constructing a DatahubTask resource.
 */
export interface DatahubTaskArgs {
    /**
     * SchemaId.
     */
    schemaId?: pulumi.Input<string>;
    /**
     * data resource.
     */
    sourceResource?: pulumi.Input<inputs.Ckafka.DatahubTaskSourceResource>;
    /**
     * Target Resource.
     */
    targetResource?: pulumi.Input<inputs.Ckafka.DatahubTaskTargetResource>;
    /**
     * name of the task.
     */
    taskName: pulumi.Input<string>;
    /**
     * type of the task, SOURCE(data input), SINK(data output).
     */
    taskType: pulumi.Input<string>;
    /**
     * Data Processing Rules.
     */
    transformParam?: pulumi.Input<inputs.Ckafka.DatahubTaskTransformParam>;
    /**
     * Data processing rules.
     */
    transformsParam?: pulumi.Input<inputs.Ckafka.DatahubTaskTransformsParam>;
}
