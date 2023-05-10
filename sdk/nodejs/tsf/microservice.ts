// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a tsf microservice
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const microservice = new tencentcloud.Tsf.Microservice("microservice", {
 *     microserviceDesc: "desc-microservice",
 *     microserviceName: "test-microservice",
 *     namespaceId: "namespace-vjlkzkgy",
 *     tags: {
 *         createdBy: "terraform",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * tsf microservice can be imported using the namespaceId#microserviceId, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Tsf/microservice:Microservice microservice namespace-vjlkzkgy#ms-vjeb43lw
 * ```
 */
export class Microservice extends pulumi.CustomResource {
    /**
     * Get an existing Microservice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MicroserviceState, opts?: pulumi.CustomResourceOptions): Microservice {
        return new Microservice(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Tsf/microservice:Microservice';

    /**
     * Returns true if the given object is an instance of Microservice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Microservice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Microservice.__pulumiType;
    }

    /**
     * Microservice description information.
     */
    public readonly microserviceDesc!: pulumi.Output<string | undefined>;
    /**
     * Microservice name.
     */
    public readonly microserviceName!: pulumi.Output<string>;
    /**
     * Namespace ID.
     */
    public readonly namespaceId!: pulumi.Output<string>;
    /**
     * Tag description list.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Microservice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MicroserviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MicroserviceArgs | MicroserviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MicroserviceState | undefined;
            resourceInputs["microserviceDesc"] = state ? state.microserviceDesc : undefined;
            resourceInputs["microserviceName"] = state ? state.microserviceName : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as MicroserviceArgs | undefined;
            if ((!args || args.microserviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'microserviceName'");
            }
            if ((!args || args.namespaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceId'");
            }
            resourceInputs["microserviceDesc"] = args ? args.microserviceDesc : undefined;
            resourceInputs["microserviceName"] = args ? args.microserviceName : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Microservice.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Microservice resources.
 */
export interface MicroserviceState {
    /**
     * Microservice description information.
     */
    microserviceDesc?: pulumi.Input<string>;
    /**
     * Microservice name.
     */
    microserviceName?: pulumi.Input<string>;
    /**
     * Namespace ID.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Microservice resource.
 */
export interface MicroserviceArgs {
    /**
     * Microservice description information.
     */
    microserviceDesc?: pulumi.Input<string>;
    /**
     * Microservice name.
     */
    microserviceName: pulumi.Input<string>;
    /**
     * Namespace ID.
     */
    namespaceId: pulumi.Input<string>;
    /**
     * Tag description list.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}
