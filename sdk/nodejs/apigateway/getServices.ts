// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query API gateway services.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const service = new tencentcloud.apigateway.Service("service", {
 *     serviceName: "niceservice",
 *     protocol: "http&https",
 *     serviceDesc: "your nice service",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const name = tencentcloud.ApiGateway.getServicesOutput({
 *     serviceName: service.serviceName,
 * });
 * const id = tencentcloud.ApiGateway.getServicesOutput({
 *     serviceId: service.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServices(args?: GetServicesArgs, opts?: pulumi.InvokeOptions): Promise<GetServicesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getServices:getServices", {
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getServices.
 */
export interface GetServicesArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Service ID for query.
     */
    serviceId?: string;
    /**
     * Service name for query.
     */
    serviceName?: string;
}

/**
 * A collection of values returned by getServices.
 */
export interface GetServicesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of services.
     */
    readonly lists: outputs.ApiGateway.GetServicesList[];
    readonly resultOutputFile?: string;
    /**
     * Custom service ID.
     */
    readonly serviceId?: string;
    /**
     * Custom service name.
     */
    readonly serviceName?: string;
}
/**
 * Use this data source to query API gateway services.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const service = new tencentcloud.apigateway.Service("service", {
 *     serviceName: "niceservice",
 *     protocol: "http&https",
 *     serviceDesc: "your nice service",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const name = tencentcloud.ApiGateway.getServicesOutput({
 *     serviceName: service.serviceName,
 * });
 * const id = tencentcloud.ApiGateway.getServicesOutput({
 *     serviceId: service.id,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getServicesOutput(args?: GetServicesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServicesResult> {
    return pulumi.output(args).apply((a: any) => getServices(a, opts))
}

/**
 * A collection of arguments for invoking getServices.
 */
export interface GetServicesOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Service ID for query.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * Service name for query.
     */
    serviceName?: pulumi.Input<string>;
}
