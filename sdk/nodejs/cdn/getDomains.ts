// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query the detail information of CDN domain.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const foo = pulumi.output(tencentcloud.Cdn.getDomains({
 *     domain: "xxxx.com",
 *     fullUrlCache: false,
 *     httpsSwitch: "on",
 *     originPullProtocol: "follow",
 *     serviceType: "web",
 * }));
 * ```
 */
export function getDomains(args?: GetDomainsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Cdn/getDomains:getDomains", {
        "domain": args.domain,
        "fullUrlCache": args.fullUrlCache,
        "httpsSwitch": args.httpsSwitch,
        "originPullProtocol": args.originPullProtocol,
        "resultOutputFile": args.resultOutputFile,
        "serviceType": args.serviceType,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsArgs {
    /**
     * Acceleration domain name.
     */
    domain?: string;
    /**
     * Whether to enable full-path cache.
     */
    fullUrlCache?: boolean;
    /**
     * HTTPS configuration. Valid values: `on`, `off` and `processing`.
     */
    httpsSwitch?: string;
    /**
     * Origin-pull protocol configuration. Valid values: `http`, `https` and `follow`.
     */
    originPullProtocol?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Service type of acceleration domain name. The available value include `web`, `download` and `media`.
     */
    serviceType?: string;
}

/**
 * A collection of values returned by getDomains.
 */
export interface GetDomainsResult {
    /**
     * Acceleration domain name.
     */
    readonly domain?: string;
    /**
     * An information list of cdn domain. Each element contains the following attributes:
     */
    readonly domainLists: outputs.Cdn.GetDomainsDomainList[];
    /**
     * Whether to enable full-path cache.
     */
    readonly fullUrlCache?: boolean;
    /**
     * HTTPS configuration switch.
     */
    readonly httpsSwitch?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Origin-pull protocol configuration.
     */
    readonly originPullProtocol?: string;
    readonly resultOutputFile?: string;
    /**
     * Service type of acceleration domain name.
     */
    readonly serviceType?: string;
}

export function getDomainsOutput(args?: GetDomainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainsResult> {
    return pulumi.output(args).apply(a => getDomains(a, opts))
}

/**
 * A collection of arguments for invoking getDomains.
 */
export interface GetDomainsOutputArgs {
    /**
     * Acceleration domain name.
     */
    domain?: pulumi.Input<string>;
    /**
     * Whether to enable full-path cache.
     */
    fullUrlCache?: pulumi.Input<boolean>;
    /**
     * HTTPS configuration. Valid values: `on`, `off` and `processing`.
     */
    httpsSwitch?: pulumi.Input<string>;
    /**
     * Origin-pull protocol configuration. Valid values: `http`, `https` and `follow`.
     */
    originPullProtocol?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Service type of acceleration domain name. The available value include `web`, `download` and `media`.
     */
    serviceType?: pulumi.Input<string>;
}
