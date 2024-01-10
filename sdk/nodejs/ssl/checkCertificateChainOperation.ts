// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ssl checkCertificateChain
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const checkCertificateChain = new tencentcloud.Ssl.CheckCertificateChainOperation("check_certificate_chain", {
 *     certificateChain: "-----BEGIN CERTIFICATE--·····---END CERTIFICATE-----",
 * });
 * ```
 *
 * ## Import
 *
 * ssl check_certificate_chain can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ssl/checkCertificateChainOperation:CheckCertificateChainOperation check_certificate_chain check_certificate_chain_id
 * ```
 */
export class CheckCertificateChainOperation extends pulumi.CustomResource {
    /**
     * Get an existing CheckCertificateChainOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CheckCertificateChainOperationState, opts?: pulumi.CustomResourceOptions): CheckCertificateChainOperation {
        return new CheckCertificateChainOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ssl/checkCertificateChainOperation:CheckCertificateChainOperation';

    /**
     * Returns true if the given object is an instance of CheckCertificateChainOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CheckCertificateChainOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CheckCertificateChainOperation.__pulumiType;
    }

    /**
     * The certificate chain to check.
     */
    public readonly certificateChain!: pulumi.Output<string>;

    /**
     * Create a CheckCertificateChainOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CheckCertificateChainOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CheckCertificateChainOperationArgs | CheckCertificateChainOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CheckCertificateChainOperationState | undefined;
            resourceInputs["certificateChain"] = state ? state.certificateChain : undefined;
        } else {
            const args = argsOrState as CheckCertificateChainOperationArgs | undefined;
            if ((!args || args.certificateChain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateChain'");
            }
            resourceInputs["certificateChain"] = args ? args.certificateChain : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CheckCertificateChainOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CheckCertificateChainOperation resources.
 */
export interface CheckCertificateChainOperationState {
    /**
     * The certificate chain to check.
     */
    certificateChain?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CheckCertificateChainOperation resource.
 */
export interface CheckCertificateChainOperationArgs {
    /**
     * The certificate chain to check.
     */
    certificateChain: pulumi.Input<string>;
}
