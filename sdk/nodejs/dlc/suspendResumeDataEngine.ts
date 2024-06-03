// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dlc suspendResumeDataEngine
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const suspendResumeDataEngine = new tencentcloud.dlc.SuspendResumeDataEngine("suspendResumeDataEngine", {
 *     dataEngineName: "example-iac",
 *     operate: "suspend",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * dlc suspend_resume_data_engine can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Dlc/suspendResumeDataEngine:SuspendResumeDataEngine suspend_resume_data_engine suspend_resume_data_engine_id
 * ```
 */
export class SuspendResumeDataEngine extends pulumi.CustomResource {
    /**
     * Get an existing SuspendResumeDataEngine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SuspendResumeDataEngineState, opts?: pulumi.CustomResourceOptions): SuspendResumeDataEngine {
        return new SuspendResumeDataEngine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dlc/suspendResumeDataEngine:SuspendResumeDataEngine';

    /**
     * Returns true if the given object is an instance of SuspendResumeDataEngine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SuspendResumeDataEngine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SuspendResumeDataEngine.__pulumiType;
    }

    /**
     * Engine name.
     */
    public readonly dataEngineName!: pulumi.Output<string>;
    /**
     * Engine operate tye: suspend/resume.
     */
    public readonly operate!: pulumi.Output<string>;

    /**
     * Create a SuspendResumeDataEngine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SuspendResumeDataEngineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SuspendResumeDataEngineArgs | SuspendResumeDataEngineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SuspendResumeDataEngineState | undefined;
            resourceInputs["dataEngineName"] = state ? state.dataEngineName : undefined;
            resourceInputs["operate"] = state ? state.operate : undefined;
        } else {
            const args = argsOrState as SuspendResumeDataEngineArgs | undefined;
            if ((!args || args.dataEngineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataEngineName'");
            }
            if ((!args || args.operate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'operate'");
            }
            resourceInputs["dataEngineName"] = args ? args.dataEngineName : undefined;
            resourceInputs["operate"] = args ? args.operate : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SuspendResumeDataEngine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SuspendResumeDataEngine resources.
 */
export interface SuspendResumeDataEngineState {
    /**
     * Engine name.
     */
    dataEngineName?: pulumi.Input<string>;
    /**
     * Engine operate tye: suspend/resume.
     */
    operate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SuspendResumeDataEngine resource.
 */
export interface SuspendResumeDataEngineArgs {
    /**
     * Engine name.
     */
    dataEngineName: pulumi.Input<string>;
    /**
     * Engine operate tye: suspend/resume.
     */
    operate: pulumi.Input<string>;
}
