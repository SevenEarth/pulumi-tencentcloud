// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a scf invokeFunction
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const invokeFunction = new tencentcloud.Scf.InvokeFunction("invoke_function", {
 *     functionName: "keep-1676351130",
 *     namespace: "default",
 *     qualifier: "2",
 * });
 * ```
 */
export class InvokeFunction extends pulumi.CustomResource {
    /**
     * Get an existing InvokeFunction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InvokeFunctionState, opts?: pulumi.CustomResourceOptions): InvokeFunction {
        return new InvokeFunction(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Scf/invokeFunction:InvokeFunction';

    /**
     * Returns true if the given object is an instance of InvokeFunction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InvokeFunction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InvokeFunction.__pulumiType;
    }

    /**
     * Function running parameter, which is in the JSON format. The maximum parameter size is 6 MB for synchronized invocations and 128KB for asynchronized invocations. This field corresponds to event input parameter.
     */
    public readonly clientContext!: pulumi.Output<string | undefined>;
    /**
     * Function name.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * Fill in RequestResponse for synchronized invocations (default and recommended) and Event for asychronized invocations. Note that for synchronized invocations, the max timeout period is 300s. Choose asychronized invocations if the required timeout period is longer than 300 seconds. You can also use InvokeFunction for synchronized invocations.
     */
    public readonly invocationType!: pulumi.Output<string | undefined>;
    /**
     * Null for async invocations.
     */
    public readonly logType!: pulumi.Output<string | undefined>;
    /**
     * Namespace.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The version or alias of the triggered function. It defaults to $LATEST.
     */
    public readonly qualifier!: pulumi.Output<string | undefined>;
    /**
     * Traffic routing config in json format, e.g., {k:v}. Please note that both k and v must be strings. Up to 1024 bytes allowed.
     */
    public readonly routingKey!: pulumi.Output<string | undefined>;

    /**
     * Create a InvokeFunction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InvokeFunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InvokeFunctionArgs | InvokeFunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InvokeFunctionState | undefined;
            resourceInputs["clientContext"] = state ? state.clientContext : undefined;
            resourceInputs["functionName"] = state ? state.functionName : undefined;
            resourceInputs["invocationType"] = state ? state.invocationType : undefined;
            resourceInputs["logType"] = state ? state.logType : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["qualifier"] = state ? state.qualifier : undefined;
            resourceInputs["routingKey"] = state ? state.routingKey : undefined;
        } else {
            const args = argsOrState as InvokeFunctionArgs | undefined;
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            resourceInputs["clientContext"] = args ? args.clientContext : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["invocationType"] = args ? args.invocationType : undefined;
            resourceInputs["logType"] = args ? args.logType : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["qualifier"] = args ? args.qualifier : undefined;
            resourceInputs["routingKey"] = args ? args.routingKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InvokeFunction.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InvokeFunction resources.
 */
export interface InvokeFunctionState {
    /**
     * Function running parameter, which is in the JSON format. The maximum parameter size is 6 MB for synchronized invocations and 128KB for asynchronized invocations. This field corresponds to event input parameter.
     */
    clientContext?: pulumi.Input<string>;
    /**
     * Function name.
     */
    functionName?: pulumi.Input<string>;
    /**
     * Fill in RequestResponse for synchronized invocations (default and recommended) and Event for asychronized invocations. Note that for synchronized invocations, the max timeout period is 300s. Choose asychronized invocations if the required timeout period is longer than 300 seconds. You can also use InvokeFunction for synchronized invocations.
     */
    invocationType?: pulumi.Input<string>;
    /**
     * Null for async invocations.
     */
    logType?: pulumi.Input<string>;
    /**
     * Namespace.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The version or alias of the triggered function. It defaults to $LATEST.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * Traffic routing config in json format, e.g., {k:v}. Please note that both k and v must be strings. Up to 1024 bytes allowed.
     */
    routingKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InvokeFunction resource.
 */
export interface InvokeFunctionArgs {
    /**
     * Function running parameter, which is in the JSON format. The maximum parameter size is 6 MB for synchronized invocations and 128KB for asynchronized invocations. This field corresponds to event input parameter.
     */
    clientContext?: pulumi.Input<string>;
    /**
     * Function name.
     */
    functionName: pulumi.Input<string>;
    /**
     * Fill in RequestResponse for synchronized invocations (default and recommended) and Event for asychronized invocations. Note that for synchronized invocations, the max timeout period is 300s. Choose asychronized invocations if the required timeout period is longer than 300 seconds. You can also use InvokeFunction for synchronized invocations.
     */
    invocationType?: pulumi.Input<string>;
    /**
     * Null for async invocations.
     */
    logType?: pulumi.Input<string>;
    /**
     * Namespace.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The version or alias of the triggered function. It defaults to $LATEST.
     */
    qualifier?: pulumi.Input<string>;
    /**
     * Traffic routing config in json format, e.g., {k:v}. Please note that both k and v must be strings. Up to 1024 bytes allowed.
     */
    routingKey?: pulumi.Input<string>;
}
