// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ssl describeHostTeoInstanceList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeHostTeoInstanceList = tencentcloud.Ssl.getDescribeHostTeoInstanceList({
 *     certificateId: "8u8DII0l",
 *     resourceType: "teo",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDescribeHostTeoInstanceList(args: GetDescribeHostTeoInstanceListArgs, opts?: pulumi.InvokeOptions): Promise<GetDescribeHostTeoInstanceListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Ssl/getDescribeHostTeoInstanceList:getDescribeHostTeoInstanceList", {
        "certificateId": args.certificateId,
        "filters": args.filters,
        "isCache": args.isCache,
        "oldCertificateId": args.oldCertificateId,
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getDescribeHostTeoInstanceList.
 */
export interface GetDescribeHostTeoInstanceListArgs {
    /**
     * Certificate ID to be deployed.
     */
    certificateId: string;
    /**
     * List of filtering parameters; Filterkey: domainmatch.
     */
    filters?: inputs.Ssl.GetDescribeHostTeoInstanceListFilter[];
    /**
     * Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
     */
    isCache?: number;
    /**
     * Deployed certificate ID.
     */
    oldCertificateId?: string;
    /**
     * Deploy resource type.
     */
    resourceType: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getDescribeHostTeoInstanceList.
 */
export interface GetDescribeHostTeoInstanceListResult {
    readonly certificateId: string;
    readonly filters?: outputs.Ssl.GetDescribeHostTeoInstanceListFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Teo instance listNote: This field may return NULL, indicating that the valid value cannot be obtained.
     */
    readonly instanceLists: outputs.Ssl.GetDescribeHostTeoInstanceListInstanceList[];
    readonly isCache?: number;
    readonly oldCertificateId?: string;
    readonly resourceType: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of ssl describeHostTeoInstanceList
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const describeHostTeoInstanceList = tencentcloud.Ssl.getDescribeHostTeoInstanceList({
 *     certificateId: "8u8DII0l",
 *     resourceType: "teo",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDescribeHostTeoInstanceListOutput(args: GetDescribeHostTeoInstanceListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDescribeHostTeoInstanceListResult> {
    return pulumi.output(args).apply((a: any) => getDescribeHostTeoInstanceList(a, opts))
}

/**
 * A collection of arguments for invoking getDescribeHostTeoInstanceList.
 */
export interface GetDescribeHostTeoInstanceListOutputArgs {
    /**
     * Certificate ID to be deployed.
     */
    certificateId: pulumi.Input<string>;
    /**
     * List of filtering parameters; Filterkey: domainmatch.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.Ssl.GetDescribeHostTeoInstanceListFilterArgs>[]>;
    /**
     * Whether to query the cache, 1: Yes; 0: No, the default is the query cache, the cache is half an hour.
     */
    isCache?: pulumi.Input<number>;
    /**
     * Deployed certificate ID.
     */
    oldCertificateId?: pulumi.Input<string>;
    /**
     * Deploy resource type.
     */
    resourceType: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
