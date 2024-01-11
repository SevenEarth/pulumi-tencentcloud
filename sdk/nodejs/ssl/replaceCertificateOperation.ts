// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a ssl replaceCertificate
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const replaceCertificate = new tencentcloud.Ssl.ReplaceCertificateOperation("replace_certificate", {
 *     certificateId: "8L6JsWq2",
 *     csrType: "online",
 *     validType: "DNS_AUTO",
 * });
 * ```
 *
 * ## Import
 *
 * ssl replace_certificate can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Ssl/replaceCertificateOperation:ReplaceCertificateOperation replace_certificate replace_certificate_id
 * ```
 */
export class ReplaceCertificateOperation extends pulumi.CustomResource {
    /**
     * Get an existing ReplaceCertificateOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplaceCertificateOperationState, opts?: pulumi.CustomResourceOptions): ReplaceCertificateOperation {
        return new ReplaceCertificateOperation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Ssl/replaceCertificateOperation:ReplaceCertificateOperation';

    /**
     * Returns true if the given object is an instance of ReplaceCertificateOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplaceCertificateOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplaceCertificateOperation.__pulumiType;
    }

    /**
     * CSR encryption method, optional: RSA, ECC, SM2. (Selectable only if CsrType is Online), default is RSA.
     */
    public readonly certCsrEncryptAlgo!: pulumi.Output<string | undefined>;
    /**
     * CSR encryption parameter, when CsrEncryptAlgo is RSA, you can choose 2048, 4096, etc., and the default is 2048; when CsrEncryptAlgo is ECC, you can choose prime256v1, secp384r1, etc., and the default is prime256v1;.
     */
    public readonly certCsrKeyParameter!: pulumi.Output<string | undefined>;
    /**
     * Certificate ID.
     */
    public readonly certificateId!: pulumi.Output<string>;
    /**
     * CSR Content.
     */
    public readonly csrContent!: pulumi.Output<string | undefined>;
    /**
     * KEY Password.
     */
    public readonly csrKeyPassword!: pulumi.Output<string | undefined>;
    /**
     * Type, default Original. Available options: Original = original certificate CSR, Upload = manual upload, Online = online generation.
     */
    public readonly csrType!: pulumi.Output<string | undefined>;
    /**
     * Reason for reissue.
     */
    public readonly reason!: pulumi.Output<string | undefined>;
    /**
     * Verification type: DNS_AUTO = automatic DNS verification (this verification type is only supported for domain names that are resolved by Tencent Cloud and have normal resolution status), DNS = manual DNS verification, FILE = file verification.
     */
    public readonly validType!: pulumi.Output<string>;

    /**
     * Create a ReplaceCertificateOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplaceCertificateOperationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplaceCertificateOperationArgs | ReplaceCertificateOperationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplaceCertificateOperationState | undefined;
            resourceInputs["certCsrEncryptAlgo"] = state ? state.certCsrEncryptAlgo : undefined;
            resourceInputs["certCsrKeyParameter"] = state ? state.certCsrKeyParameter : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["csrContent"] = state ? state.csrContent : undefined;
            resourceInputs["csrKeyPassword"] = state ? state.csrKeyPassword : undefined;
            resourceInputs["csrType"] = state ? state.csrType : undefined;
            resourceInputs["reason"] = state ? state.reason : undefined;
            resourceInputs["validType"] = state ? state.validType : undefined;
        } else {
            const args = argsOrState as ReplaceCertificateOperationArgs | undefined;
            if ((!args || args.certificateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateId'");
            }
            if ((!args || args.validType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'validType'");
            }
            resourceInputs["certCsrEncryptAlgo"] = args ? args.certCsrEncryptAlgo : undefined;
            resourceInputs["certCsrKeyParameter"] = args ? args.certCsrKeyParameter : undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["csrContent"] = args ? args.csrContent : undefined;
            resourceInputs["csrKeyPassword"] = args ? args.csrKeyPassword : undefined;
            resourceInputs["csrType"] = args ? args.csrType : undefined;
            resourceInputs["reason"] = args ? args.reason : undefined;
            resourceInputs["validType"] = args ? args.validType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ReplaceCertificateOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplaceCertificateOperation resources.
 */
export interface ReplaceCertificateOperationState {
    /**
     * CSR encryption method, optional: RSA, ECC, SM2. (Selectable only if CsrType is Online), default is RSA.
     */
    certCsrEncryptAlgo?: pulumi.Input<string>;
    /**
     * CSR encryption parameter, when CsrEncryptAlgo is RSA, you can choose 2048, 4096, etc., and the default is 2048; when CsrEncryptAlgo is ECC, you can choose prime256v1, secp384r1, etc., and the default is prime256v1;.
     */
    certCsrKeyParameter?: pulumi.Input<string>;
    /**
     * Certificate ID.
     */
    certificateId?: pulumi.Input<string>;
    /**
     * CSR Content.
     */
    csrContent?: pulumi.Input<string>;
    /**
     * KEY Password.
     */
    csrKeyPassword?: pulumi.Input<string>;
    /**
     * Type, default Original. Available options: Original = original certificate CSR, Upload = manual upload, Online = online generation.
     */
    csrType?: pulumi.Input<string>;
    /**
     * Reason for reissue.
     */
    reason?: pulumi.Input<string>;
    /**
     * Verification type: DNS_AUTO = automatic DNS verification (this verification type is only supported for domain names that are resolved by Tencent Cloud and have normal resolution status), DNS = manual DNS verification, FILE = file verification.
     */
    validType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplaceCertificateOperation resource.
 */
export interface ReplaceCertificateOperationArgs {
    /**
     * CSR encryption method, optional: RSA, ECC, SM2. (Selectable only if CsrType is Online), default is RSA.
     */
    certCsrEncryptAlgo?: pulumi.Input<string>;
    /**
     * CSR encryption parameter, when CsrEncryptAlgo is RSA, you can choose 2048, 4096, etc., and the default is 2048; when CsrEncryptAlgo is ECC, you can choose prime256v1, secp384r1, etc., and the default is prime256v1;.
     */
    certCsrKeyParameter?: pulumi.Input<string>;
    /**
     * Certificate ID.
     */
    certificateId: pulumi.Input<string>;
    /**
     * CSR Content.
     */
    csrContent?: pulumi.Input<string>;
    /**
     * KEY Password.
     */
    csrKeyPassword?: pulumi.Input<string>;
    /**
     * Type, default Original. Available options: Original = original certificate CSR, Upload = manual upload, Online = online generation.
     */
    csrType?: pulumi.Input<string>;
    /**
     * Reason for reissue.
     */
    reason?: pulumi.Input<string>;
    /**
     * Verification type: DNS_AUTO = automatic DNS verification (this verification type is only supported for domain names that are resolved by Tencent Cloud and have normal resolution status), DNS = manual DNS verification, FILE = file verification.
     */
    validType: pulumi.Input<string>;
}