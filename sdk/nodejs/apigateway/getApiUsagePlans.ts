// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of apigateway apiUsagePlan
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@tencentcloud_iac/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const exampleUsagePlan = new tencentcloud.apigateway.UsagePlan("exampleUsagePlan", {
 *     usagePlanName: "tf_example",
 *     usagePlanDesc: "desc.",
 *     maxRequestNum: 100,
 *     maxRequestNumPreSec: 10,
 * });
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
 *     authType: "SECRET",
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
 * const exampleUsagePlanAttachment = new tencentcloud.apigateway.UsagePlanAttachment("exampleUsagePlanAttachment", {
 *     usagePlanId: exampleUsagePlan.id,
 *     serviceId: exampleService.id,
 *     environment: "release",
 *     bindType: "API",
 *     apiId: exampleApi.id,
 * });
 * const exampleApiUsagePlans = tencentcloud.ApiGateway.getApiUsagePlansOutput({
 *     serviceId: exampleUsagePlanAttachment.serviceId,
 * });
 * ```
 */
export function getApiUsagePlans(args: GetApiUsagePlansArgs, opts?: pulumi.InvokeOptions): Promise<GetApiUsagePlansResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:ApiGateway/getApiUsagePlans:getApiUsagePlans", {
        "resultOutputFile": args.resultOutputFile,
        "serviceId": args.serviceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiUsagePlans.
 */
export interface GetApiUsagePlansArgs {
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
 * A collection of values returned by getApiUsagePlans.
 */
export interface GetApiUsagePlansResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
    /**
     * API binding usage plan list.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly results: outputs.ApiGateway.GetApiUsagePlansResult[];
    /**
     * Service unique ID.Note: This field may return null, indicating that a valid value cannot be obtained.
     */
    readonly serviceId: string;
}

export function getApiUsagePlansOutput(args: GetApiUsagePlansOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApiUsagePlansResult> {
    return pulumi.output(args).apply(a => getApiUsagePlans(a, opts))
}

/**
 * A collection of arguments for invoking getApiUsagePlans.
 */
export interface GetApiUsagePlansOutputArgs {
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * The unique ID of the service to be queried.
     */
    serviceId: pulumi.Input<string>;
}