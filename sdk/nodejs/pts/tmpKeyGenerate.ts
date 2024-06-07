// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a pts tmpKey
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const tmpKey = new tencentcloud.pts.TmpKeyGenerate("tmpKey", {
 *     projectId: "project-1b0zqmhg",
 *     scenarioId: "scenario-abc",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class TmpKeyGenerate extends pulumi.CustomResource {
    /**
     * Get an existing TmpKeyGenerate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TmpKeyGenerateState, opts?: pulumi.CustomResourceOptions): TmpKeyGenerate {
        return new TmpKeyGenerate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Pts/tmpKeyGenerate:TmpKeyGenerate';

    /**
     * Returns true if the given object is an instance of TmpKeyGenerate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TmpKeyGenerate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TmpKeyGenerate.__pulumiType;
    }

    /**
     * Temporary access credentials.
     */
    public /*out*/ readonly credentials!: pulumi.Output<outputs.Pts.TmpKeyGenerateCredential[]>;
    /**
     * Timestamp of temporary access credential timeout (in seconds).
     */
    public /*out*/ readonly expiredTime!: pulumi.Output<number>;
    /**
     * Project ID.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * Scenario ID.
     */
    public readonly scenarioId!: pulumi.Output<string | undefined>;
    /**
     * The timestamp of the moment when the temporary access credential was obtained (in seconds).
     */
    public /*out*/ readonly startTime!: pulumi.Output<number>;

    /**
     * Create a TmpKeyGenerate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TmpKeyGenerateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TmpKeyGenerateArgs | TmpKeyGenerateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TmpKeyGenerateState | undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["expiredTime"] = state ? state.expiredTime : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["scenarioId"] = state ? state.scenarioId : undefined;
            resourceInputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as TmpKeyGenerateArgs | undefined;
            if ((!args || args.projectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'projectId'");
            }
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["scenarioId"] = args ? args.scenarioId : undefined;
            resourceInputs["credentials"] = undefined /*out*/;
            resourceInputs["expiredTime"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TmpKeyGenerate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TmpKeyGenerate resources.
 */
export interface TmpKeyGenerateState {
    /**
     * Temporary access credentials.
     */
    credentials?: pulumi.Input<pulumi.Input<inputs.Pts.TmpKeyGenerateCredential>[]>;
    /**
     * Timestamp of temporary access credential timeout (in seconds).
     */
    expiredTime?: pulumi.Input<number>;
    /**
     * Project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * Scenario ID.
     */
    scenarioId?: pulumi.Input<string>;
    /**
     * The timestamp of the moment when the temporary access credential was obtained (in seconds).
     */
    startTime?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a TmpKeyGenerate resource.
 */
export interface TmpKeyGenerateArgs {
    /**
     * Project ID.
     */
    projectId: pulumi.Input<string>;
    /**
     * Scenario ID.
     */
    scenarioId?: pulumi.Input<string>;
}
