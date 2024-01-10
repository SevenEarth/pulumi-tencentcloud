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
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 *
 * const examplePlugin = new tencentcloud.apigateway.Plugin("examplePlugin", {
 *     pluginName: "tf-example",
 *     pluginType: "IPControl",
 *     pluginData: JSON.stringify({
 *         type: "white_list",
 *         blocks: "1.1.1.1",
 *     }),
 *     description: "desc.",
 * });
 * const exampleService = new tencentcloud.apigateway.Service("exampleService", {
 *     serviceName: "tf_example_service",
 *     protocol: "http&https",
 *     serviceDesc: "your nice service",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const exampleApi = new tencentcloud.apigateway.Api("exampleApi", {
 *     serviceId: exampleService.id,
 *     apiName: "tf_example_api",
 *     apiDesc: "desc.",
 *     authType: "APP",
 *     protocol: "HTTP",
 *     enableCors: true,
 *     requestConfigPath: "/user/info",
 *     requestConfigMethod: "GET",
 *     requestParameters: [{
 *         name: "name",
 *         position: "QUERY",
 *         type: "string",
 *         desc: "desc.",
 *         defaultValue: "terraform",
 *         required: true,
 *     }],
 *     serviceConfigType: "HTTP",
 *     serviceConfigTimeout: 15,
 *     serviceConfigUrl: "https://www.qq.com",
 *     serviceConfigPath: "/user",
 *     serviceConfigMethod: "GET",
 *     responseType: "HTML",
 *     responseSuccessExample: "success",
 *     responseFailExample: "fail",
 *     responseErrorCodes: [{
 *         code: 400,
 *         msg: "system error msg.",
 *         desc: "system error desc.",
 *         convertedCode: 407,
 *         needConvert: true,
 *     }],
 * });
 * const exampleServiceRelease = new tencentcloud.apigateway.ServiceRelease("exampleServiceRelease", {
 *     serviceId: exampleApi.serviceId,
 *     environmentName: "release",
 *     releaseDesc: "desc.",
 * });
 * const examplePluginAttachment = new tencentcloud.apigateway.PluginAttachment("examplePluginAttachment", {
 *     pluginId: examplePlugin.id,
 *     serviceId: exampleServiceRelease.serviceId,
 *     apiId: exampleApi.id,
 *     environmentName: "release",
 * });
 * ```
 *
 * ## Import
 *
 * apiGateway plugin_attachment can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:ApiGateway/pluginAttachment:PluginAttachment example plugin-hnqntalp#service-q3f533ja#release#api-62ud9woa
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
