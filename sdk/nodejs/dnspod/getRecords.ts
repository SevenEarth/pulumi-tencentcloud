// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query dnspod record list.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const record = tencentcloud.Dnspod.getRecords({
 *     domain: "example.com",
 *     subdomain: "www",
 * });
 * export const result = record.then(record => record.results);
 * ```
 * ### Use verbose filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const record = tencentcloud.Dnspod.getRecords({
 *     domain: "example.com",
 *     subdomain: "www",
 *     limit: 100,
 *     recordType: "TXT",
 *     sortField: "updated_on",
 *     sortType: "DESC",
 * });
 * export const result = record.then(record => record.results);
 * ```
 */
export function getRecords(args?: GetRecordsArgs, opts?: pulumi.InvokeOptions): Promise<GetRecordsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Dnspod/getRecords:getRecords", {
        "domain": args.domain,
        "domainId": args.domainId,
        "groupId": args.groupId,
        "keyword": args.keyword,
        "limit": args.limit,
        "offset": args.offset,
        "recordLine": args.recordLine,
        "recordLineId": args.recordLineId,
        "recordType": args.recordType,
        "resultOutputFile": args.resultOutputFile,
        "sortField": args.sortField,
        "sortType": args.sortType,
        "subdomain": args.subdomain,
    }, opts);
}

/**
 * A collection of arguments for invoking getRecords.
 */
export interface GetRecordsArgs {
    /**
     * The domain for which DNS records are to be obtained.
     */
    domain?: string;
    /**
     * The ID of the domain for which DNS records are to be obtained. If DomainId is passed in, the system will omit the parameter domain.
     */
    domainId?: string;
    /**
     * The group ID.
     */
    groupId?: string;
    /**
     * The keyword for searching for DNS records. Host headers and record values are supported.
     */
    keyword?: string;
    /**
     * The limit. It defaults to 100 and can be up to 3,000.
     */
    limit?: number;
    /**
     * The offset. Default value: 0.
     */
    offset?: number;
    /**
     * The split zone name.
     */
    recordLine?: string;
    /**
     * The split zone ID. If `recordLineId` is passed in, the system will omit the parameter `recordLine`.
     */
    recordLineId?: string;
    /**
     * The type of DNS record, such as A, CNAME, NS, AAAA, explicit URL, implicit URL, CAA, or SPF record.
     */
    recordType?: string;
    /**
     * Used for store query result as JSON.
     */
    resultOutputFile?: string;
    /**
     * The sorting field. Available values: name, line, type, value, weight, mx, and ttl,updated_on.
     */
    sortField?: string;
    /**
     * The sorting type. Valid values: ASC (ascending, default), DESC (descending).
     */
    sortType?: string;
    /**
     * The host header of a DNS record. If this parameter is passed in, only the DNS record corresponding to this host header will be returned.
     */
    subdomain?: string;
}

/**
 * A collection of values returned by getRecords.
 */
export interface GetRecordsResult {
    readonly domain?: string;
    readonly domainId?: string;
    readonly groupId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyword?: string;
    readonly limit?: number;
    readonly offset?: number;
    /**
     * Count info of the queried record list.
     */
    readonly recordCountInfos: outputs.Dnspod.GetRecordsRecordCountInfo[];
    readonly recordLine?: string;
    readonly recordLineId?: string;
    readonly recordType?: string;
    readonly resultOutputFile?: string;
    /**
     * The record list result.
     */
    readonly results: outputs.Dnspod.GetRecordsResult[];
    readonly sortField?: string;
    readonly sortType?: string;
    readonly subdomain?: string;
}

export function getRecordsOutput(args?: GetRecordsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRecordsResult> {
    return pulumi.output(args).apply(a => getRecords(a, opts))
}

/**
 * A collection of arguments for invoking getRecords.
 */
export interface GetRecordsOutputArgs {
    /**
     * The domain for which DNS records are to be obtained.
     */
    domain?: pulumi.Input<string>;
    /**
     * The ID of the domain for which DNS records are to be obtained. If DomainId is passed in, the system will omit the parameter domain.
     */
    domainId?: pulumi.Input<string>;
    /**
     * The group ID.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The keyword for searching for DNS records. Host headers and record values are supported.
     */
    keyword?: pulumi.Input<string>;
    /**
     * The limit. It defaults to 100 and can be up to 3,000.
     */
    limit?: pulumi.Input<number>;
    /**
     * The offset. Default value: 0.
     */
    offset?: pulumi.Input<number>;
    /**
     * The split zone name.
     */
    recordLine?: pulumi.Input<string>;
    /**
     * The split zone ID. If `recordLineId` is passed in, the system will omit the parameter `recordLine`.
     */
    recordLineId?: pulumi.Input<string>;
    /**
     * The type of DNS record, such as A, CNAME, NS, AAAA, explicit URL, implicit URL, CAA, or SPF record.
     */
    recordType?: pulumi.Input<string>;
    /**
     * Used for store query result as JSON.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * The sorting field. Available values: name, line, type, value, weight, mx, and ttl,updated_on.
     */
    sortField?: pulumi.Input<string>;
    /**
     * The sorting type. Valid values: ASC (ascending, default), DESC (descending).
     */
    sortType?: pulumi.Input<string>;
    /**
     * The host header of a DNS record. If this parameter is passed in, only the DNS record corresponding to this host header will be returned.
     */
    subdomain?: pulumi.Input<string>;
}
