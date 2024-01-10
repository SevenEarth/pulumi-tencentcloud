// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a elasticsearch logstash pipeline
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const logstashPipeline = new tencentcloud.Elasticsearch.LogstashPipeline("logstash_pipeline", {
 *     instanceId: "ls-xxxxxx",
 *     opType: 2,
 *     pipeline: {
 *         batchDelay: 50,
 *         batchSize: 125,
 *         config: `input{
 *
 * }
 * filter{
 *
 * }
 * output{
 *
 * }
 * `,
 *         pipelineDesc: "",
 *         pipelineId: "logstash-pipeline-test",
 *         queueCheckPointWrites: 0,
 *         queueMaxBytes: "",
 *         queueType: "memory",
 *         workers: 1,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * elasticsearch logstash_pipeline can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Elasticsearch/logstashPipeline:LogstashPipeline logstash_pipeline ${instance_id}#${pipeline_id}
 * ```
 */
export class LogstashPipeline extends pulumi.CustomResource {
    /**
     * Get an existing LogstashPipeline resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LogstashPipelineState, opts?: pulumi.CustomResourceOptions): LogstashPipeline {
        return new LogstashPipeline(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Elasticsearch/logstashPipeline:LogstashPipeline';

    /**
     * Returns true if the given object is an instance of LogstashPipeline.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LogstashPipeline {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LogstashPipeline.__pulumiType;
    }

    /**
     * Logstash instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Operation type. 1: save only; 2: save and deploy.
     */
    public readonly opType!: pulumi.Output<number>;
    /**
     * Pipeline information.
     */
    public readonly pipeline!: pulumi.Output<outputs.Elasticsearch.LogstashPipelinePipeline>;

    /**
     * Create a LogstashPipeline resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LogstashPipelineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LogstashPipelineArgs | LogstashPipelineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LogstashPipelineState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["opType"] = state ? state.opType : undefined;
            resourceInputs["pipeline"] = state ? state.pipeline : undefined;
        } else {
            const args = argsOrState as LogstashPipelineArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.opType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'opType'");
            }
            if ((!args || args.pipeline === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pipeline'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["opType"] = args ? args.opType : undefined;
            resourceInputs["pipeline"] = args ? args.pipeline : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LogstashPipeline.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LogstashPipeline resources.
 */
export interface LogstashPipelineState {
    /**
     * Logstash instance id.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Operation type. 1: save only; 2: save and deploy.
     */
    opType?: pulumi.Input<number>;
    /**
     * Pipeline information.
     */
    pipeline?: pulumi.Input<inputs.Elasticsearch.LogstashPipelinePipeline>;
}

/**
 * The set of arguments for constructing a LogstashPipeline resource.
 */
export interface LogstashPipelineArgs {
    /**
     * Logstash instance id.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Operation type. 1: save only; 2: save and deploy.
     */
    opType: pulumi.Input<number>;
    /**
     * Pipeline information.
     */
    pipeline: pulumi.Input<inputs.Elasticsearch.LogstashPipelinePipeline>;
}
