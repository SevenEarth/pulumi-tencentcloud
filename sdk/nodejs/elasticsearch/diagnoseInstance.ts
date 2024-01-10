// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a elasticsearch diagnose instance
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const diagnoseInstance = new tencentcloud.Elasticsearch.DiagnoseInstance("diagnose_instance", {
 *     diagnoseIndices: "*",
 *     diagnoseJobs: ["cluster_health"],
 *     instanceId: "es-xxxxxx",
 * });
 * ```
 */
export class DiagnoseInstance extends pulumi.CustomResource {
    /**
     * Get an existing DiagnoseInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiagnoseInstanceState, opts?: pulumi.CustomResourceOptions): DiagnoseInstance {
        return new DiagnoseInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Elasticsearch/diagnoseInstance:DiagnoseInstance';

    /**
     * Returns true if the given object is an instance of DiagnoseInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DiagnoseInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DiagnoseInstance.__pulumiType;
    }

    /**
     * Indexes that need to be diagnosed. Wildcards are supported.
     */
    public readonly diagnoseIndices!: pulumi.Output<string | undefined>;
    /**
     * Diagnostic items that need to be triggered.
     */
    public readonly diagnoseJobs!: pulumi.Output<string[] | undefined>;
    /**
     * Instance id.
     */
    public readonly instanceId!: pulumi.Output<string>;

    /**
     * Create a DiagnoseInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiagnoseInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiagnoseInstanceArgs | DiagnoseInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DiagnoseInstanceState | undefined;
            resourceInputs["diagnoseIndices"] = state ? state.diagnoseIndices : undefined;
            resourceInputs["diagnoseJobs"] = state ? state.diagnoseJobs : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
        } else {
            const args = argsOrState as DiagnoseInstanceArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["diagnoseIndices"] = args ? args.diagnoseIndices : undefined;
            resourceInputs["diagnoseJobs"] = args ? args.diagnoseJobs : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DiagnoseInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DiagnoseInstance resources.
 */
export interface DiagnoseInstanceState {
    /**
     * Indexes that need to be diagnosed. Wildcards are supported.
     */
    diagnoseIndices?: pulumi.Input<string>;
    /**
     * Diagnostic items that need to be triggered.
     */
    diagnoseJobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance id.
     */
    instanceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DiagnoseInstance resource.
 */
export interface DiagnoseInstanceArgs {
    /**
     * Indexes that need to be diagnosed. Wildcards are supported.
     */
    diagnoseIndices?: pulumi.Input<string>;
    /**
     * Diagnostic items that need to be triggered.
     */
    diagnoseJobs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance id.
     */
    instanceId: pulumi.Input<string>;
}
