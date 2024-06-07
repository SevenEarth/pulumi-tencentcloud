// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query API gateway APIs.
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
 *     serviceName: "ck",
 *     protocol: "http&https",
 *     serviceDesc: "your nice service",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const api = new tencentcloud.apigateway.Api("api", {
 *     serviceId: service.id,
 *     apiName: "hello",
 *     apiDesc: "my hello api",
 *     authType: "NONE",
 *     protocol: "HTTP",
 *     enableCors: true,
 *     requestConfigPath: "/user/info",
 *     requestConfigMethod: "GET",
 *     serviceConfigType: "HTTP",
 *     serviceConfigTimeout: 15,
 *     serviceConfigUrl: "http://www.qq.com",
 *     serviceConfigPath: "/user",
 *     serviceConfigMethod: "GET",
 *     responseType: "HTML",
 *     responseSuccessExample: "success",
 *     responseFailExample: "fail",
 * });
 * const id = tencentcloud.ApiGateway.getApisOutput({
 *     serviceId: service.id,
 *     apiId: api.id,
 * });
 * const name = tencentcloud.ApiGateway.getApisOutput({
 *     serviceId: service.id,
 *     apiName: api.apiName,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getApis(args: GetApisArgs, opts?: pulumi.InvokeOptions): Promise<GetApisResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getApis:getApis", {
        "apiId": args.apiId,
        "apiName": args.apiName,
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApis.
 */
export interface GetApisArgs {
    /**
     * Created API ID.
     */
    apiId?: string;
    /**
     * Custom API name.
     */
    apiName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Service ID for query.
     */
    serviceId: string;
}

/**
 * A collection of values returned by getApis.
 */
export interface GetApisResult {
    readonly apiId?: string;
    /**
     * Custom API name.
     */
    readonly apiName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of APIs.
     */
    readonly lists: outputs.ApiGateway.GetApisList[];
    readonly resultOutputFile?: string;
    /**
     * Which service this API belongs. Refer to resource `tencentcloud.ApiGateway.Service`.
     */
    readonly serviceId: string;
}
/**
 * Use this data source to query API gateway APIs.
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
 *     serviceName: "ck",
 *     protocol: "http&https",
 *     serviceDesc: "your nice service",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const api = new tencentcloud.apigateway.Api("api", {
 *     serviceId: service.id,
 *     apiName: "hello",
 *     apiDesc: "my hello api",
 *     authType: "NONE",
 *     protocol: "HTTP",
 *     enableCors: true,
 *     requestConfigPath: "/user/info",
 *     requestConfigMethod: "GET",
 *     serviceConfigType: "HTTP",
 *     serviceConfigTimeout: 15,
 *     serviceConfigUrl: "http://www.qq.com",
 *     serviceConfigPath: "/user",
 *     serviceConfigMethod: "GET",
 *     responseType: "HTML",
 *     responseSuccessExample: "success",
 *     responseFailExample: "fail",
 * });
 * const id = tencentcloud.ApiGateway.getApisOutput({
 *     serviceId: service.id,
 *     apiId: api.id,
 * });
 * const name = tencentcloud.ApiGateway.getApisOutput({
 *     serviceId: service.id,
 *     apiName: api.apiName,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getApisOutput(args: GetApisOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApisResult> {
    return pulumi.output(args).apply((a: any) => getApis(a, opts))
}

/**
 * A collection of arguments for invoking getApis.
 */
export interface GetApisOutputArgs {
    /**
     * Created API ID.
     */
    apiId?: pulumi.Input<string>;
    /**
     * Custom API name.
     */
    apiName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Service ID for query.
     */
    serviceId: pulumi.Input<string>;
}
