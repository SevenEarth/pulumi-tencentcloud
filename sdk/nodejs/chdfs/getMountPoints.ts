// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of chdfs mountPoints
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mountPoints = tencentcloud.Chdfs.getMountPoints({
 *     fileSystemId: "f14mpfy5lh4e",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMountPoints(args?: GetMountPointsArgs, opts?: pulumi.InvokeOptions): Promise<GetMountPointsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Chdfs/getMountPoints:getMountPoints", {
        "accessGroupId": args.accessGroupId,
        "fileSystemId": args.fileSystemId,
        "ownerUin": args.ownerUin,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getMountPoints.
 */
export interface GetMountPointsArgs {
    /**
     * get mount points belongs to access group id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
     */
    accessGroupId?: string;
    /**
     * get mount points belongs to file system id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
     */
    fileSystemId?: string;
    /**
     * get mount points belongs to owner uin, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
     */
    ownerUin?: number;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getMountPoints.
 */
export interface GetMountPointsResult {
    readonly accessGroupId?: string;
    /**
     * mounted file system id.
     */
    readonly fileSystemId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * mount point list.
     */
    readonly mountPoints: outputs.Chdfs.GetMountPointsMountPoint[];
    readonly ownerUin?: number;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of chdfs mountPoints
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const mountPoints = tencentcloud.Chdfs.getMountPoints({
 *     fileSystemId: "f14mpfy5lh4e",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getMountPointsOutput(args?: GetMountPointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMountPointsResult> {
    return pulumi.output(args).apply((a: any) => getMountPoints(a, opts))
}

/**
 * A collection of arguments for invoking getMountPoints.
 */
export interface GetMountPointsOutputArgs {
    /**
     * get mount points belongs to access group id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
     */
    accessGroupId?: pulumi.Input<string>;
    /**
     * get mount points belongs to file system id, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * get mount points belongs to owner uin, only can use one of the AccessGroupId,FileSystemId,OwnerUin parameters.
     */
    ownerUin?: pulumi.Input<number>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
