// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of waf peakPoints
 *
 * ## Example Usage
 * ### Basic Query
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Waf.getPeakPoints({
 *     fromTime: "2023-09-01 00:00:00",
 *     toTime: "2023-09-07 00:00:00",
 * }));
 * ```
 * ### Query by filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = pulumi.output(tencentcloud.Waf.getPeakPoints({
 *     domain: "domain.com",
 *     edition: "clb-waf",
 *     fromTime: "2023-09-01 00:00:00",
 *     instanceId: "waf_2kxtlbky00b2v1fn",
 *     metricName: "access",
 *     toTime: "2023-09-07 00:00:00",
 * }));
 * ```
 */
export function getPeakPoints(args: GetPeakPointsArgs, opts?: pulumi.InvokeOptions): Promise<GetPeakPointsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Waf/getPeakPoints:getPeakPoints", {
        "domain": args.domain,
        "edition": args.edition,
        "fromTime": args.fromTime,
        "instanceId": args.instanceId,
        "metricName": args.metricName,
        "resultOutputFile": args.resultOutputFile,
        "toTime": args.toTime,
    }, opts);
}

/**
 * A collection of arguments for invoking getPeakPoints.
 */
export interface GetPeakPointsArgs {
    /**
     * The domain name to be queried. If all domain name data is queried, this parameter is not filled in.
     */
    domain?: string;
    /**
     * Only support sparta-waf and clb-waf. If not passed, there will be no filtering.
     */
    edition?: string;
    /**
     * Begin time.
     */
    fromTime: string;
    /**
     * WAF instance ID, if not passed, there will be no filtering.
     */
    instanceId?: string;
    /**
     * Thirteen values are available: access-Peak qps trend chart; botAccess- bot peak qps trend chart; down-Downstream peak bandwidth trend chart; up-Upstream peak bandwidth trend chart; attack-Trend chart of total number of web attacks; cc-Trend chart of total number of CC attacks; StatusServerError-Trend chart of the number of status codes returned by WAF to the server; StatusClientError-Trend chart of the number of status codes returned by WAF to the client; StatusRedirect-Trend chart of the number of status codes returned by WAF to the client; StatusOk-Trend chart of the number of status codes returned by WAF to the client; UpstreamServerError-Trend chart of the number of status codes returned to WAF by the origin site; UpstreamClientError-Trend chart of the number of status codes returned to WAF by the origin site; UpstreamRedirect-Trend chart of the number of status codes returned to WAF by the origin site.
     */
    metricName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * End time.
     */
    toTime: string;
}

/**
 * A collection of values returned by getPeakPoints.
 */
export interface GetPeakPointsResult {
    readonly domain?: string;
    readonly edition?: string;
    readonly fromTime: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId?: string;
    readonly metricName?: string;
    /**
     * point list.
     */
    readonly points: outputs.Waf.GetPeakPointsPoint[];
    readonly resultOutputFile?: string;
    readonly toTime: string;
}

export function getPeakPointsOutput(args: GetPeakPointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPeakPointsResult> {
    return pulumi.output(args).apply(a => getPeakPoints(a, opts))
}

/**
 * A collection of arguments for invoking getPeakPoints.
 */
export interface GetPeakPointsOutputArgs {
    /**
     * The domain name to be queried. If all domain name data is queried, this parameter is not filled in.
     */
    domain?: pulumi.Input<string>;
    /**
     * Only support sparta-waf and clb-waf. If not passed, there will be no filtering.
     */
    edition?: pulumi.Input<string>;
    /**
     * Begin time.
     */
    fromTime: pulumi.Input<string>;
    /**
     * WAF instance ID, if not passed, there will be no filtering.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Thirteen values are available: access-Peak qps trend chart; botAccess- bot peak qps trend chart; down-Downstream peak bandwidth trend chart; up-Upstream peak bandwidth trend chart; attack-Trend chart of total number of web attacks; cc-Trend chart of total number of CC attacks; StatusServerError-Trend chart of the number of status codes returned by WAF to the server; StatusClientError-Trend chart of the number of status codes returned by WAF to the client; StatusRedirect-Trend chart of the number of status codes returned by WAF to the client; StatusOk-Trend chart of the number of status codes returned by WAF to the client; UpstreamServerError-Trend chart of the number of status codes returned to WAF by the origin site; UpstreamClientError-Trend chart of the number of status codes returned to WAF by the origin site; UpstreamRedirect-Trend chart of the number of status codes returned to WAF by the origin site.
     */
    metricName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * End time.
     */
    toTime: pulumi.Input<string>;
}
