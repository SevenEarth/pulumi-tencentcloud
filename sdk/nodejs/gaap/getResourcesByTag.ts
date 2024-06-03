// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of gaap resources by tag
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const resourcesByTag = tencentcloud.Gaap.getResourcesByTag({
 *     tagKey: "tagKey",
 *     tagValue: "tagValue",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getResourcesByTag(args: GetResourcesByTagArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcesByTagResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Gaap/getResourcesByTag:getResourcesByTag", {
        "resourceType": args.resourceType,
        "resultOutputFile": args.resultOutputFile,
        "tagKey": args.tagKey,
        "tagValue": args.tagValue,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourcesByTag.
 */
export interface GetResourcesByTagArgs {
    /**
     * Resource type, where:Proxy represents the proxy;ProxyGroup represents a proxy group;RealServer represents the Real Server.If this field is not specified, all resources under the label will be queried.
     */
    resourceType?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
    /**
     * Tag key.
     */
    tagKey: string;
    /**
     * Tag value.
     */
    tagValue: string;
}

/**
 * A collection of values returned by getResourcesByTag.
 */
export interface GetResourcesByTagResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of resources corresponding to labels.
     */
    readonly resourceSets: outputs.Gaap.GetResourcesByTagResourceSet[];
    /**
     * Resource type, where:Proxy represents the proxy,ProxyGroup represents a proxy group,RealServer represents the real server.
     */
    readonly resourceType?: string;
    readonly resultOutputFile?: string;
    readonly tagKey: string;
    readonly tagValue: string;
}
/**
 * Use this data source to query detailed information of gaap resources by tag
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const resourcesByTag = tencentcloud.Gaap.getResourcesByTag({
 *     tagKey: "tagKey",
 *     tagValue: "tagValue",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getResourcesByTagOutput(args: GetResourcesByTagOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourcesByTagResult> {
    return pulumi.output(args).apply((a: any) => getResourcesByTag(a, opts))
}

/**
 * A collection of arguments for invoking getResourcesByTag.
 */
export interface GetResourcesByTagOutputArgs {
    /**
     * Resource type, where:Proxy represents the proxy;ProxyGroup represents a proxy group;RealServer represents the Real Server.If this field is not specified, all resources under the label will be queried.
     */
    resourceType?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
    /**
     * Tag key.
     */
    tagKey: pulumi.Input<string>;
    /**
     * Tag value.
     */
    tagValue: pulumi.Input<string>;
}
