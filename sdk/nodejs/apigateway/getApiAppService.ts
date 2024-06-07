// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of apigateway apiAppServices
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const exampleService = new tencentcloud.apigateway.Service("exampleService", {
 *     serviceName: "tf_example",
 *     protocol: "http&https",
 *     serviceDesc: "desc.",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const exampleApi = new tencentcloud.apigateway.Api("exampleApi", {
 *     serviceId: exampleService.id,
 *     apiName: "tf_example",
 *     apiDesc: "my hello api update",
 *     authType: "APP",
 *     protocol: "HTTP",
 *     enableCors: true,
 *     requestConfigPath: "/user/info",
 *     requestConfigMethod: "POST",
 *     requestParameters: [{
 *         name: "email",
 *         position: "QUERY",
 *         type: "string",
 *         desc: "desc.",
 *         defaultValue: "test@qq.com",
 *         required: true,
 *     }],
 *     serviceConfigType: "HTTP",
 *     serviceConfigTimeout: 10,
 *     serviceConfigUrl: "http://www.tencent.com",
 *     serviceConfigPath: "/user",
 *     serviceConfigMethod: "POST",
 *     responseType: "XML",
 *     responseSuccessExample: "<note>success</note>",
 *     responseFailExample: "<note>fail</note>",
 *     responseErrorCodes: [{
 *         code: 500,
 *         msg: "system error",
 *         desc: "system error code",
 *         convertedCode: 5000,
 *         needConvert: true,
 *     }],
 * });
 * const exampleApiAppService = tencentcloud.ApiGateway.getApiAppServiceOutput({
 *     serviceId: exampleApi.serviceId,
 *     apiRegion: "ap-guangzhou",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getApiAppService(args: GetApiAppServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetApiAppServiceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getApiAppService:getApiAppService", {
        "apiRegion": args.apiRegion,
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiAppService.
 */
export interface GetApiAppServiceArgs {
    /**
     * Territory to which the service belongs.
     */
    apiRegion: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * The unique ID of the service to be queried.
     */
    serviceId: string;
}

/**
 * A collection of values returned by getApiAppService.
 */
export interface GetApiAppServiceResult {
    /**
     * API list.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly apiIdStatusSets: outputs.ApiGateway.GetApiAppServiceApiIdStatusSet[];
    readonly apiRegion: string;
    /**
     * Total number of APIs.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly apiTotalCount: number;
    /**
     * List of service environments.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly availableEnvironments: string[];
    /**
     * Use planned time.
     */
    readonly createdTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Internal network access HTTP service port number.
     */
    readonly innerHttpPort: number;
    /**
     * Internal network access https port number.
     */
    readonly innerHttpsPort: number;
    /**
     * Intranet access sub domain name.
     */
    readonly internalSubDomain: string;
    /**
     * IP version.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly ipVersion: string;
    /**
     * Use the schedule to modify the time.
     */
    readonly modifiedTime: string;
    /**
     * A list of network types, where INNER represents internal network access and OUTER represents external network access.
     */
    readonly netTypes: string[];
    /**
     * External network access sub domain name.
     */
    readonly outerSubDomain: string;
    /**
     * Service support protocol, optional values are http, https, and http&amp;amp;https.
     */
    readonly protocol: string;
    readonly resultOutputFile?: string;
    /**
     * Service description.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly serviceDesc: string;
    /**
     * Service unique ID.
     */
    readonly serviceId: string;
    /**
     * Service name.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly serviceName: string;
    /**
     * Reserved fields.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly setId: number;
    /**
     * Use a plan array.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly usagePlanLists: outputs.ApiGateway.GetApiAppServiceUsagePlanList[];
    /**
     * Total number of usage plans.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly usagePlanTotalCount: number;
    /**
     * The user type of this service.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly userType: string;
}
/**
 * Use this data source to query detailed information of apigateway apiAppServices
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const exampleService = new tencentcloud.apigateway.Service("exampleService", {
 *     serviceName: "tf_example",
 *     protocol: "http&https",
 *     serviceDesc: "desc.",
 *     netTypes: [
 *         "INNER",
 *         "OUTER",
 *     ],
 *     ipVersion: "IPv4",
 * });
 * const exampleApi = new tencentcloud.apigateway.Api("exampleApi", {
 *     serviceId: exampleService.id,
 *     apiName: "tf_example",
 *     apiDesc: "my hello api update",
 *     authType: "APP",
 *     protocol: "HTTP",
 *     enableCors: true,
 *     requestConfigPath: "/user/info",
 *     requestConfigMethod: "POST",
 *     requestParameters: [{
 *         name: "email",
 *         position: "QUERY",
 *         type: "string",
 *         desc: "desc.",
 *         defaultValue: "test@qq.com",
 *         required: true,
 *     }],
 *     serviceConfigType: "HTTP",
 *     serviceConfigTimeout: 10,
 *     serviceConfigUrl: "http://www.tencent.com",
 *     serviceConfigPath: "/user",
 *     serviceConfigMethod: "POST",
 *     responseType: "XML",
 *     responseSuccessExample: "<note>success</note>",
 *     responseFailExample: "<note>fail</note>",
 *     responseErrorCodes: [{
 *         code: 500,
 *         msg: "system error",
 *         desc: "system error code",
 *         convertedCode: 5000,
 *         needConvert: true,
 *     }],
 * });
 * const exampleApiAppService = tencentcloud.ApiGateway.getApiAppServiceOutput({
 *     serviceId: exampleApi.serviceId,
 *     apiRegion: "ap-guangzhou",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getApiAppServiceOutput(args: GetApiAppServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiAppServiceResult> {
    return pulumi.output(args).apply((a: any) => getApiAppService(a, opts))
}

/**
 * A collection of arguments for invoking getApiAppService.
 */
export interface GetApiAppServiceOutputArgs {
    /**
     * Territory to which the service belongs.
     */
    apiRegion: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * The unique ID of the service to be queried.
     */
    serviceId: pulumi.Input<string>;
}
