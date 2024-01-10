// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ssl
{
    /// <summary>
    /// Provides a resource to create a ssl replace_certificate
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var replaceCertificate = new Tencentcloud.Ssl.ReplaceCertificateOperation("replaceCertificate", new Tencentcloud.Ssl.ReplaceCertificateOperationArgs
    ///         {
    ///             CertificateId = "8L6JsWq2",
    ///             CsrType = "online",
    ///             ValidType = "DNS_AUTO",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ssl replace_certificate can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Ssl/replaceCertificateOperation:ReplaceCertificateOperation replace_certificate replace_certificate_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ssl/replaceCertificateOperation:ReplaceCertificateOperation")]
    public partial class ReplaceCertificateOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// CSR encryption method, optional: RSA, ECC, SM2. (Selectable only if CsrType is Online), default is RSA.
        /// </summary>
        [Output("certCsrEncryptAlgo")]
        public Output<string?> CertCsrEncryptAlgo { get; private set; } = null!;

        /// <summary>
        /// CSR encryption parameter, when CsrEncryptAlgo is RSA, you can choose 2048, 4096, etc., and the default is 2048; when CsrEncryptAlgo is ECC, you can choose prime256v1, secp384r1, etc., and the default is prime256v1;.
        /// </summary>
        [Output("certCsrKeyParameter")]
        public Output<string?> CertCsrKeyParameter { get; private set; } = null!;

        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// CSR Content.
        /// </summary>
        [Output("csrContent")]
        public Output<string?> CsrContent { get; private set; } = null!;

        /// <summary>
        /// KEY Password.
        /// </summary>
        [Output("csrKeyPassword")]
        public Output<string?> CsrKeyPassword { get; private set; } = null!;

        /// <summary>
        /// Type, default Original. Available options: Original = original certificate CSR, Upload = manual upload, Online = online generation.
        /// </summary>
        [Output("csrType")]
        public Output<string?> CsrType { get; private set; } = null!;

        /// <summary>
        /// Reason for reissue.
        /// </summary>
        [Output("reason")]
        public Output<string?> Reason { get; private set; } = null!;

        /// <summary>
        /// Verification type: DNS_AUTO = automatic DNS verification (this verification type is only supported for domain names that are resolved by Tencent Cloud and have normal resolution status), DNS = manual DNS verification, FILE = file verification.
        /// </summary>
        [Output("validType")]
        public Output<string> ValidType { get; private set; } = null!;


        /// <summary>
        /// Create a ReplaceCertificateOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplaceCertificateOperation(string name, ReplaceCertificateOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/replaceCertificateOperation:ReplaceCertificateOperation", name, args ?? new ReplaceCertificateOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplaceCertificateOperation(string name, Input<string> id, ReplaceCertificateOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/replaceCertificateOperation:ReplaceCertificateOperation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReplaceCertificateOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplaceCertificateOperation Get(string name, Input<string> id, ReplaceCertificateOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplaceCertificateOperation(name, id, state, options);
        }
    }

    public sealed class ReplaceCertificateOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CSR encryption method, optional: RSA, ECC, SM2. (Selectable only if CsrType is Online), default is RSA.
        /// </summary>
        [Input("certCsrEncryptAlgo")]
        public Input<string>? CertCsrEncryptAlgo { get; set; }

        /// <summary>
        /// CSR encryption parameter, when CsrEncryptAlgo is RSA, you can choose 2048, 4096, etc., and the default is 2048; when CsrEncryptAlgo is ECC, you can choose prime256v1, secp384r1, etc., and the default is prime256v1;.
        /// </summary>
        [Input("certCsrKeyParameter")]
        public Input<string>? CertCsrKeyParameter { get; set; }

        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        /// <summary>
        /// CSR Content.
        /// </summary>
        [Input("csrContent")]
        public Input<string>? CsrContent { get; set; }

        /// <summary>
        /// KEY Password.
        /// </summary>
        [Input("csrKeyPassword")]
        public Input<string>? CsrKeyPassword { get; set; }

        /// <summary>
        /// Type, default Original. Available options: Original = original certificate CSR, Upload = manual upload, Online = online generation.
        /// </summary>
        [Input("csrType")]
        public Input<string>? CsrType { get; set; }

        /// <summary>
        /// Reason for reissue.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Verification type: DNS_AUTO = automatic DNS verification (this verification type is only supported for domain names that are resolved by Tencent Cloud and have normal resolution status), DNS = manual DNS verification, FILE = file verification.
        /// </summary>
        [Input("validType", required: true)]
        public Input<string> ValidType { get; set; } = null!;

        public ReplaceCertificateOperationArgs()
        {
        }
    }

    public sealed class ReplaceCertificateOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// CSR encryption method, optional: RSA, ECC, SM2. (Selectable only if CsrType is Online), default is RSA.
        /// </summary>
        [Input("certCsrEncryptAlgo")]
        public Input<string>? CertCsrEncryptAlgo { get; set; }

        /// <summary>
        /// CSR encryption parameter, when CsrEncryptAlgo is RSA, you can choose 2048, 4096, etc., and the default is 2048; when CsrEncryptAlgo is ECC, you can choose prime256v1, secp384r1, etc., and the default is prime256v1;.
        /// </summary>
        [Input("certCsrKeyParameter")]
        public Input<string>? CertCsrKeyParameter { get; set; }

        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// CSR Content.
        /// </summary>
        [Input("csrContent")]
        public Input<string>? CsrContent { get; set; }

        /// <summary>
        /// KEY Password.
        /// </summary>
        [Input("csrKeyPassword")]
        public Input<string>? CsrKeyPassword { get; set; }

        /// <summary>
        /// Type, default Original. Available options: Original = original certificate CSR, Upload = manual upload, Online = online generation.
        /// </summary>
        [Input("csrType")]
        public Input<string>? CsrType { get; set; }

        /// <summary>
        /// Reason for reissue.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// Verification type: DNS_AUTO = automatic DNS verification (this verification type is only supported for domain names that are resolved by Tencent Cloud and have normal resolution status), DNS = manual DNS verification, FILE = file verification.
        /// </summary>
        [Input("validType")]
        public Input<string>? ValidType { get; set; }

        public ReplaceCertificateOperationState()
        {
        }
    }
}
