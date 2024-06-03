// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of eb plateformEventTemplate
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const plateformEventTemplate = tencentcloud.Eb.getPlateformEventTemplate({
 *     eventType: "eb_platform_test:TEST:ALL",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPlateformEventTemplate(args: GetPlateformEventTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetPlateformEventTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tencentcloud:Eb/getPlateformEventTemplate:getPlateformEventTemplate", {
        "eventType": args.eventType,
        "resultOutputFile": args.resultOutputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getPlateformEventTemplate.
 */
export interface GetPlateformEventTemplateArgs {
    /**
     * Platform product event type.
     */
    eventType: string;
    /**
     * Used to save results.
     */
    resultOutputFile?: string;
}

/**
 * A collection of values returned by getPlateformEventTemplate.
 */
export interface GetPlateformEventTemplateResult {
    /**
     * Platform product event template.
     */
    readonly eventTemplate: string;
    readonly eventType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly resultOutputFile?: string;
}
/**
 * Use this data source to query detailed information of eb plateformEventTemplate
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const plateformEventTemplate = tencentcloud.Eb.getPlateformEventTemplate({
 *     eventType: "eb_platform_test:TEST:ALL",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getPlateformEventTemplateOutput(args: GetPlateformEventTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlateformEventTemplateResult> {
    return pulumi.output(args).apply((a: any) => getPlateformEventTemplate(a, opts))
}

/**
 * A collection of arguments for invoking getPlateformEventTemplate.
 */
export interface GetPlateformEventTemplateOutputArgs {
    /**
     * Platform product event type.
     */
    eventType: pulumi.Input<string>;
    /**
     * Used to save results.
     */
    resultOutputFile?: pulumi.Input<string>;
}
