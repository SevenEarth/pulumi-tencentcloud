// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a scf syncInvokeFunction
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const invokeFunction = new tencentcloud.scf.SyncInvokeFunction("invokeFunction", {
 *     functionName: "keep-1676351130",
 *     namespace: "default",
 *     qualifier: "2",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SyncInvokeFunction extends pulumi.CustomResource {
    /**
     * Get an existing SyncInvokeFunction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SyncInvokeFunctionState, opts?: pulumi.CustomResourceOptions): SyncInvokeFunction {
        return new SyncInvokeFunction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Scf/syncInvokeFunction:SyncInvokeFunction';

    /**
     * Returns true if the given object is an instance of SyncInvokeFunction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SyncInvokeFunction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SyncInvokeFunction.__pulumiType;
    }

    /**
     * Function running parameter, which is in the JSON format. Maximum parameter size is 6 MB. This field corresponds to event input parameter.
     */
    public readonly event!: pulumi.Output<string | undefined>;
    /**
     * Function name.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * Valid value: None (default) or Tail. If the value is Tail, log in the response will contain the corresponding function execution log (up to 4KB).
     */
    public readonly logType!: pulumi.Output<string | undefined>;
    /**
     * Namespace. default is used if it's left empty.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Version or alias of the function. It defaults to $DEFAULT.
     */
    public readonly qualifier!: pulumi.Output<string | undefined>;
    /**
     * Traffic routing config in json format, e.g., {k:v}. Please note that both k and v must be strings. Up to 1024 bytes allowed.
     */
    public readonly routingKey!: pulumi.Output<string | undefined>;

    /**
     * Create a SyncInvokeFunction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SyncInvokeFunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SyncInvokeFunctionArgs | SyncInvokeFunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SyncInvokeFunctionState | undefined;
            resourceInputs["event"] = state ? state.event : undefined;
            resourceInputs["functionName"] = state ? state.functionName : undefined;
            resourceInputs["logType"] = state ? state.logType : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["qualifier"] = state ? state.qualifier : undefined;
            resourceInputs["routingKey"] = state ? state.routingKey : undefined;
        } else {
            const args = argsOrState as SyncInvokeFunctionArgs | undefined;
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            resourceInputs["event"] = args ? args.event : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["logType"] = args ? args.logType : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["qualifier"] = args ? args.qualifier : undefined;
            resourceInputs["routingKey"] = args ? args.routingKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SyncInvokeFunction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SyncInvokeFunction resources.
 */
export interface SyncInvokeFunctionState {
    /**
     * Function running parameter, which is in the JSON format. Maximum parameter size is 6 MB. This field corresponds to event input parameter.
     */
    event?: pulumi.Input<string>;
    /**
     * Function name.
     */
    functionName?: pulumi.Input<string>;
    /**
     * Valid value: None (default) or Tail. If the value is Tail, log in the response will contain the corresponding function execution log (up to 4KB).
     */
    logType?: pulumi.Input<string>;
    /**
     * Namespace. default is used if it's left empty.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Version or alias of the function. It defaults to $DEFAULT.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * Traffic routing config in json format, e.g., {k:v}. Please note that both k and v must be strings. Up to 1024 bytes allowed.
     */
    routingKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SyncInvokeFunction resource.
 */
export interface SyncInvokeFunctionArgs {
    /**
     * Function running parameter, which is in the JSON format. Maximum parameter size is 6 MB. This field corresponds to event input parameter.
     */
    event?: pulumi.Input<string>;
    /**
     * Function name.
     */
    functionName: pulumi.Input<string>;
    /**
     * Valid value: None (default) or Tail. If the value is Tail, log in the response will contain the corresponding function execution log (up to 4KB).
     */
    logType?: pulumi.Input<string>;
    /**
     * Namespace. default is used if it's left empty.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Version or alias of the function. It defaults to $DEFAULT.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * Traffic routing config in json format, e.g., {k:v}. Please note that both k and v must be strings. Up to 1024 bytes allowed.
     */
    routingKey?: pulumi.Input<string>;
}
