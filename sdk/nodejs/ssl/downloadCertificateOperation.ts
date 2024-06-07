// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ssl downloadCertificate
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@tencentcloud_iac/pulumi";
 *
 * const downloadCertificate = new tencentcloud.ssl.DownloadCertificateOperation("downloadCertificate", {
 *     certificateId: "8x1eUSSl",
 *     outputPath: "./",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * ssl download_certificate can be imported using the id, e.g.
 *
 * ```sh
 * $ pulumi import tencentcloud:Ssl/downloadCertificateOperation:DownloadCertificateOperation download_certificate download_certificate_id
 * ```
 */
export class DownloadCertificateOperation extends pulumi.CustomResource {
    /**
     * Get an existing DownloadCertificateOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DownloadCertificateOperationState, opts?: pulumi.CustomResourceOptions): DownloadCertificateOperation {
        return new DownloadCertificateOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ssl/downloadCertificateOperation:DownloadCertificateOperation';

    /**
     * Returns true if the given object is an instance of DownloadCertificateOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DownloadCertificateOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DownloadCertificateOperation.__pulumiType;
    }

    /**
     * Certificate ID.
     */
    public readonly certificateId!: pulumi.Output<string>;
    /**
     * Certificate ID.
     */
    public readonly outputPath!: pulumi.Output<string>;

    /**
     * Create a DownloadCertificateOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DownloadCertificateOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DownloadCertificateOperationArgs | DownloadCertificateOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DownloadCertificateOperationState | undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["outputPath"] = state ? state.outputPath : undefined;
        } else {
            const args = argsOrState as DownloadCertificateOperationArgs | undefined;
            if ((!args || args.certificateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateId'");
            }
            if ((!args || args.outputPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'outputPath'");
            }
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["outputPath"] = args ? args.outputPath : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DownloadCertificateOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DownloadCertificateOperation resources.
 */
export interface DownloadCertificateOperationState {
    /**
     * Certificate ID.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * Certificate ID.
     */
    outputPath?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DownloadCertificateOperation resource.
 */
export interface DownloadCertificateOperationArgs {
    /**
     * Certificate ID.
     */
    certificateId: pulumi.Input<string>;
    /**
     * Certificate ID.
     */
    outputPath: pulumi.Input<string>;
}
