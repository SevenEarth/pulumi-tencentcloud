// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of gaap listener real servers
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const listenerRealServers = tencentcloud.Gaap.getListenerRealServers({
 *     listenerId: "listener-xxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getListenerRealServers(args: GetListenerRealServersArgs, opts?: pulumi.InvokeOptions): Promise<GetListenerRealServersResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Gaap/getListenerRealServers:getListenerRealServers", {
        "listenerId": args.listenerId,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getListenerRealServers.
 */
export interface GetListenerRealServersArgs {
    /**
     * listener ID.
     */
    listenerId: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getListenerRealServers.
 */
export interface GetListenerRealServersResult {
    /**
     * Bound real server Information List.
     */
    readonly bindRealServerSets: outputs.Gaap.GetListenerRealServersBindRealServerSet[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly listenerId: string;
    /**
     * Real Server Set.
     */
    readonly realServerSets: outputs.Gaap.GetListenerRealServersRealServerSet[];
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of gaap listener real servers
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const listenerRealServers = tencentcloud.Gaap.getListenerRealServers({
 *     listenerId: "listener-xxxxxx",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getListenerRealServersOutput(args: GetListenerRealServersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetListenerRealServersResult> {
    return pulumi.output(args).apply((a: any) => getListenerRealServers(a, opts))
}

/**
 * A collection of arguments for invoking getListenerRealServers.
 */
export interface GetListenerRealServersOutputArgs {
    /**
     * listener ID.
     */
    listenerId: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
