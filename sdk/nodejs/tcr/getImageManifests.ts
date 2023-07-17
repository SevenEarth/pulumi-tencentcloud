// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tcr imageManifests
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const imageManifests = pulumi.output(tencentcloud.Tcr.getImageManifests({
 *     imageVersion: "v1",
 *     namespaceName: "%s",
 *     registryId: "%s",
 *     repositoryName: "%s",
 * }));
 * ```
 */
export function getImageManifests(args: GetImageManifestsArgs, opts?: pulumi.InvokeOptions): Promise<GetImageManifestsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tencentcloud:Tcr/getImageManifests:getImageManifests", {
        "imageVersion": args.imageVersion,
        "namespaceName": args.namespaceName,
        "registryId": args.registryId,
        "repositoryName": args.repositoryName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getImageManifests.
 */
export interface GetImageManifestsArgs {
    /**
     * mirror version.
     */
    imageVersion: string;
    /**
     * namespace name.
     */
    namespaceName: string;
    /**
     * instance ID.
     */
    registryId: string;
    /**
     * mirror warehouse name.
     */
    repositoryName: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getImageManifests.
 */
export interface GetImageManifestsResult {
    /**
     * configuration information of the image.
     */
    readonly config: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly imageVersion: string;
    /**
     * Manifest information of the image.
     */
    readonly manifest: string;
    readonly namespaceName: string;
    readonly registryId: string;
    readonly repositoryName: string;
    readonly resultOutputFile?: string;
}

export function getImageManifestsOutput(args: GetImageManifestsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetImageManifestsResult> {
    return pulumi.output(args).apply(a => getImageManifests(a, opts))
}

/**
 * A collection of arguments for invoking getImageManifests.
 */
export interface GetImageManifestsOutputArgs {
    /**
     * mirror version.
     */
    imageVersion: pulumi.Input<string>;
    /**
     * namespace name.
     */
    namespaceName: pulumi.Input<string>;
    /**
     * instance ID.
     */
    registryId: pulumi.Input<string>;
    /**
     * mirror warehouse name.
     */
    repositoryName: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
