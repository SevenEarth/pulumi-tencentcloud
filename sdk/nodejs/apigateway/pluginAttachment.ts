// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a apiGateway pluginAttachment
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const pluginAttachment = new tencentcloud.ApiGateway.PluginAttachment("plugin_attachment", {
 *     apiId: "api-6tfrdysk",
 *     environmentName: "test",
 *     pluginId: "plugin-ny74siyz",
 *     serviceId: "service-n1mgl0sq",
 * });
 * ```
 *
 * ## Import
 *
 * apiGateway plugin_attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:ApiGateway/pluginAttachment:PluginAttachment plugin_attachment pluginId#serviceId#environmentName#apiId
 * ```
 */
export class PluginAttachment extends pulumi.CustomResource {
    /**
     * Get an existing PluginAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PluginAttachmentState, opts?: pulumi.CustomResourceOptions): PluginAttachment {
        return new PluginAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:ApiGateway/pluginAttachment:PluginAttachment';

    /**
     * Returns true if the given object is an instance of PluginAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PluginAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PluginAttachment.__pulumiType;
    }

    /**
     * Id of API.
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * Name of Environment.
     */
    public readonly environmentName!: pulumi.Output<string>;
    /**
     * Id of Plugin.
     */
    public readonly pluginId!: pulumi.Output<string>;
    /**
     * Id of Service.
     */
    public readonly serviceId!: pulumi.Output<string>;

    /**
     * Create a PluginAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PluginAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PluginAttachmentArgs | PluginAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PluginAttachmentState | undefined;
            resourceInputs["apiId"] = state ? state.apiId : undefined;
            resourceInputs["environmentName"] = state ? state.environmentName : undefined;
            resourceInputs["pluginId"] = state ? state.pluginId : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
        } else {
            const args = argsOrState as PluginAttachmentArgs | undefined;
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.environmentName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentName'");
            }
            if ((!args || args.pluginId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pluginId'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["environmentName"] = args ? args.environmentName : undefined;
            resourceInputs["pluginId"] = args ? args.pluginId : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PluginAttachment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PluginAttachment resources.
 */
export interface PluginAttachmentState {
    /**
     * Id of API.
     */
    apiId?: pulumi.Input<string>;
    /**
     * Name of Environment.
     */
    environmentName?: pulumi.Input<string>;
    /**
     * Id of Plugin.
     */
    pluginId?: pulumi.Input<string>;
    /**
     * Id of Service.
     */
    serviceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PluginAttachment resource.
 */
export interface PluginAttachmentArgs {
    /**
     * Id of API.
     */
    apiId: pulumi.Input<string>;
    /**
     * Name of Environment.
     */
    environmentName: pulumi.Input<string>;
    /**
     * Id of Plugin.
     */
    pluginId: pulumi.Input<string>;
    /**
     * Id of Service.
     */
    serviceId: pulumi.Input<string>;
}