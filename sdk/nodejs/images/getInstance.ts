// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query images.
 *
 * ## Example Usage
 *
 * ### Query all images
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by image ID
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     imageId: "img-9qrfy1xt",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by os name
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     osName: "TencentOS Server 3.2 (Final)",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by image name regex
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     imageNameRegex: "^TencentOS",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by image type
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     imageTypes: ["PUBLIC_IMAGE"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by instance type
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     instanceType: "S1.SMALL1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Images/getInstance:getInstance", {
        "imageId": args.imageId,
        "imageNameRegex": args.imageNameRegex,
        "imageTypes": args.imageTypes,
        "instanceType": args.instanceType,
        "osName": args.osName,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * ID of the image to be queried.
     */
    imageId?: string;
    /**
     * A regex string to apply to the image list returned by TencentCloud, conflict with 'os_name'. **NOTE**: it is not wildcard, should look like `imageNameRegex = "^CentOS\s+6\.8\s+64\w*"`.
     */
    imageNameRegex?: string;
    /**
     * A list of the image type to be queried. Valid values: 'PUBLIC_IMAGE', 'PRIVATE_IMAGE', 'SHARED_IMAGE', 'MARKET_IMAGE'.
     */
    imageTypes?: string[];
    /**
     * Instance type, such as `S1.SMALL1`.
     */
    instanceType?: string;
    /**
     * A string to apply with fuzzy match to the osName attribute on the image list returned by TencentCloud, conflict with 'image_name_regex'.
     */
    osName?: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ID of the image.
     */
    readonly imageId?: string;
    readonly imageNameRegex?: string;
    /**
     * Type of the image.
     */
    readonly imageTypes?: string[];
    /**
     * An information list of image. Each element contains the following attributes:
     */
    readonly images: outputs.Images.GetInstanceImage[];
    readonly instanceType?: string;
    /**
     * OS name of the image.
     */
    readonly osName?: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query images.
 *
 * ## Example Usage
 *
 * ### Query all images
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by image ID
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     imageId: "img-9qrfy1xt",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by os name
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     osName: "TencentOS Server 3.2 (Final)",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by image name regex
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     imageNameRegex: "^TencentOS",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by image type
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     imageTypes: ["PUBLIC_IMAGE"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Query images by instance type
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const example = tencentcloud.Images.getInstance({
 *     instanceType: "S1.SMALL1",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * ID of the image to be queried.
     */
    imageId?: pulumi.Input<string>;
    /**
     * A regex string to apply to the image list returned by TencentCloud, conflict with 'os_name'. **NOTE**: it is not wildcard, should look like `imageNameRegex = "^CentOS\s+6\.8\s+64\w*"`.
     */
    imageNameRegex?: pulumi.Input<string>;
    /**
     * A list of the image type to be queried. Valid values: 'PUBLIC_IMAGE', 'PRIVATE_IMAGE', 'SHARED_IMAGE', 'MARKET_IMAGE'.
     */
    imageTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance type, such as `S1.SMALL1`.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * A string to apply with fuzzy match to the osName attribute on the image list returned by TencentCloud, conflict with 'image_name_regex'.
     */
    osName?: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
