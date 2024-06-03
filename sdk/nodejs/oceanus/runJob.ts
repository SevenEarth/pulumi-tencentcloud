// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a oceanus runJob
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const example = new tencentcloud.oceanus.RunJob("example", {
 *     runJobDescriptions: [{
 *         jobConfigVersion: 10,
 *         jobId: "cql-4xwincyn",
 *         runType: 1,
 *         startMode: "LATEST",
 *         useOldSystemConnector: false,
 *     }],
 *     workSpaceId: "space-2idq8wbr",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class RunJob extends pulumi.CustomResource {
    /**
     * Get an existing RunJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RunJobState, opts?: pulumi.CustomResourceOptions): RunJob {
        return new RunJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Oceanus/runJob:RunJob';

    /**
     * Returns true if the given object is an instance of RunJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RunJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RunJob.__pulumiType;
    }

    /**
     * The description information for batch job startup.
     */
    public readonly runJobDescriptions!: pulumi.Output<outputs.Oceanus.RunJobRunJobDescription[]>;
    /**
     * Workspace SerialId.
     */
    public readonly workSpaceId!: pulumi.Output<string | undefined>;

    /**
     * Create a RunJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RunJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RunJobArgs | RunJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RunJobState | undefined;
            resourceInputs["runJobDescriptions"] = state ? state.runJobDescriptions : undefined;
            resourceInputs["workSpaceId"] = state ? state.workSpaceId : undefined;
        } else {
            const args = argsOrState as RunJobArgs | undefined;
            if ((!args || args.runJobDescriptions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runJobDescriptions'");
            }
            resourceInputs["runJobDescriptions"] = args ? args.runJobDescriptions : undefined;
            resourceInputs["workSpaceId"] = args ? args.workSpaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RunJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RunJob resources.
 */
export interface RunJobState {
    /**
     * The description information for batch job startup.
     */
    runJobDescriptions?: pulumi.Input<pulumi.Input<inputs.Oceanus.RunJobRunJobDescription>[]>;
    /**
     * Workspace SerialId.
     */
    workSpaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RunJob resource.
 */
export interface RunJobArgs {
    /**
     * The description information for batch job startup.
     */
    runJobDescriptions: pulumi.Input<pulumi.Input<inputs.Oceanus.RunJobRunJobDescription>[]>;
    /**
     * Workspace SerialId.
     */
    workSpaceId?: pulumi.Input<string>;
}
