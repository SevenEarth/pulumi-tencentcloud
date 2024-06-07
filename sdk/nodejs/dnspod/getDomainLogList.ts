// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of dnspod domainLogList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const domainLogList = tencentcloud.Dnspod.getDomainLogList({
 *     domain: "iac-tf.cloud",
 *     domainId: 123,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDomainLogList(args: GetDomainLogListArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainLogListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Dnspod/getDomainLogList:getDomainLogList", {
        "domain": args.domain,
        "domainId": args.domainId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomainLogList.
 */
export interface GetDomainLogListArgs {
    /**
     * Domain.
     */
    domain: string;
    /**
     * Domain ID. The parameter DomainId has a higher priority than the parameter Domain. If the parameter DomainId is passed, the parameter Domain will be ignored. You can find all Domains and DomainIds through the DescribeDomainList interface.
     */
    domainId?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDomainLogList.
 */
export interface GetDomainLogListResult {
    readonly domain: string;
    readonly domainId?: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Domain Operation Log List. Note: This field may return null, indicating that no valid value can be obtained.
     */
    readonly logLists: string[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of dnspod domainLogList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const domainLogList = tencentcloud.Dnspod.getDomainLogList({
 *     domain: "iac-tf.cloud",
 *     domainId: 123,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDomainLogListOutput(args: GetDomainLogListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainLogListResult> {
    return pulumi.output(args).apply((a: any) => getDomainLogList(a, opts))
}

/**
 * A collection of arguments for invoking getDomainLogList.
 */
export interface GetDomainLogListOutputArgs {
    /**
     * Domain.
     */
    domain: pulumi.Input<string>;
    /**
     * Domain ID. The parameter DomainId has a higher priority than the parameter Domain. If the parameter DomainId is passed, the parameter Domain will be ignored. You can find all Domains and DomainIds through the DescribeDomainList interface.
     */
    domainId?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
