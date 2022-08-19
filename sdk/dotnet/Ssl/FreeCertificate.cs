// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Ssl
{
    /// <summary>
    /// Provide a resource to create a Free Certificate.
    /// &gt; **NOTE:** Once certificat created, it cannot be removed within 1 hours.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Ssl.FreeCertificate("foo", new Tencentcloud.Ssl.FreeCertificateArgs
    ///         {
    ///             Alias = "my_free_cert",
    ///             ContactEmail = "foo@example.com",
    ///             ContactPhone = "12345678901",
    ///             CsrEncryptAlgo = "RSA",
    ///             CsrKeyParameter = "2048",
    ///             CsrKeyPassword = "xxxxxxxx",
    ///             Domain = "example.com",
    ///             DvAuthMethod = "DNS_AUTO",
    ///             PackageType = "2",
    ///             ValidityPeriod = "12",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// FreeCertificate instance can be imported, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Ssl/freeCertificate:FreeCertificate test free_certificate-id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ssl/freeCertificate:FreeCertificate")]
    public partial class FreeCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// Specify alias for remark.
        /// </summary>
        [Output("alias")]
        public Output<string?> Alias { get; private set; } = null!;

        /// <summary>
        /// Certificate begin time.
        /// </summary>
        [Output("certBeginTime")]
        public Output<string> CertBeginTime { get; private set; } = null!;

        /// <summary>
        /// Certificate end time.
        /// </summary>
        [Output("certEndTime")]
        public Output<string> CertEndTime { get; private set; } = null!;

        /// <summary>
        /// Certificate private key.
        /// </summary>
        [Output("certificatePrivateKey")]
        public Output<string> CertificatePrivateKey { get; private set; } = null!;

        /// <summary>
        /// Certificate public key.
        /// </summary>
        [Output("certificatePublicKey")]
        public Output<string> CertificatePublicKey { get; private set; } = null!;

        /// <summary>
        /// Email address.
        /// </summary>
        [Output("contactEmail")]
        public Output<string?> ContactEmail { get; private set; } = null!;

        /// <summary>
        /// Phone number.
        /// </summary>
        [Output("contactPhone")]
        public Output<string?> ContactPhone { get; private set; } = null!;

        /// <summary>
        /// Specify CSR encrypt algorithm, only support `RSA` for now.
        /// </summary>
        [Output("csrEncryptAlgo")]
        public Output<string?> CsrEncryptAlgo { get; private set; } = null!;

        /// <summary>
        /// Specify CSR key parameter, only support `"2048"` for now.
        /// </summary>
        [Output("csrKeyParameter")]
        public Output<string?> CsrKeyParameter { get; private set; } = null!;

        /// <summary>
        /// Specify CSR key password.
        /// </summary>
        [Output("csrKeyPassword")]
        public Output<string?> CsrKeyPassword { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the certificate deployable.
        /// </summary>
        [Output("deployable")]
        public Output<bool> Deployable { get; private set; } = null!;

        /// <summary>
        /// Specify domain name.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Specify DV authorize method, available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
        /// </summary>
        [Output("dvAuthMethod")]
        public Output<string> DvAuthMethod { get; private set; } = null!;

        /// <summary>
        /// Certificate insert time.
        /// </summary>
        [Output("insertTime")]
        public Output<string> InsertTime { get; private set; } = null!;

        /// <summary>
        /// Specify old certificate ID, used for re-apply.
        /// </summary>
        [Output("oldCertificateId")]
        public Output<string?> OldCertificateId { get; private set; } = null!;

        /// <summary>
        /// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
        /// </summary>
        [Output("packageType")]
        public Output<string?> PackageType { get; private set; } = null!;

        /// <summary>
        /// Product zh name.
        /// </summary>
        [Output("productZhName")]
        public Output<string> ProductZhName { get; private set; } = null!;

        /// <summary>
        /// ID of projects which this certification belong to.
        /// </summary>
        [Output("projectId")]
        public Output<int?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the certificate renewable.
        /// </summary>
        [Output("renewable")]
        public Output<bool> Renewable { get; private set; } = null!;

        /// <summary>
        /// Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Certificate status message.
        /// </summary>
        [Output("statusMsg")]
        public Output<string> StatusMsg { get; private set; } = null!;

        /// <summary>
        /// Certificate status name.
        /// </summary>
        [Output("statusName")]
        public Output<string> StatusName { get; private set; } = null!;

        /// <summary>
        /// Specify validity period in month, only support `"12"` months for now.
        /// </summary>
        [Output("validityPeriod")]
        public Output<string?> ValidityPeriod { get; private set; } = null!;

        /// <summary>
        /// Vulnerability status.
        /// </summary>
        [Output("vulnerabilityStatus")]
        public Output<string> VulnerabilityStatus { get; private set; } = null!;


        /// <summary>
        /// Create a FreeCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FreeCertificate(string name, FreeCertificateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/freeCertificate:FreeCertificate", name, args ?? new FreeCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FreeCertificate(string name, Input<string> id, FreeCertificateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/freeCertificate:FreeCertificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FreeCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FreeCertificate Get(string name, Input<string> id, FreeCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new FreeCertificate(name, id, state, options);
        }
    }

    public sealed class FreeCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify alias for remark.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Email address.
        /// </summary>
        [Input("contactEmail")]
        public Input<string>? ContactEmail { get; set; }

        /// <summary>
        /// Phone number.
        /// </summary>
        [Input("contactPhone")]
        public Input<string>? ContactPhone { get; set; }

        /// <summary>
        /// Specify CSR encrypt algorithm, only support `RSA` for now.
        /// </summary>
        [Input("csrEncryptAlgo")]
        public Input<string>? CsrEncryptAlgo { get; set; }

        /// <summary>
        /// Specify CSR key parameter, only support `"2048"` for now.
        /// </summary>
        [Input("csrKeyParameter")]
        public Input<string>? CsrKeyParameter { get; set; }

        /// <summary>
        /// Specify CSR key password.
        /// </summary>
        [Input("csrKeyPassword")]
        public Input<string>? CsrKeyPassword { get; set; }

        /// <summary>
        /// Specify domain name.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Specify DV authorize method, available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
        /// </summary>
        [Input("dvAuthMethod", required: true)]
        public Input<string> DvAuthMethod { get; set; } = null!;

        /// <summary>
        /// Specify old certificate ID, used for re-apply.
        /// </summary>
        [Input("oldCertificateId")]
        public Input<string>? OldCertificateId { get; set; }

        /// <summary>
        /// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// ID of projects which this certification belong to.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Specify validity period in month, only support `"12"` months for now.
        /// </summary>
        [Input("validityPeriod")]
        public Input<string>? ValidityPeriod { get; set; }

        public FreeCertificateArgs()
        {
        }
    }

    public sealed class FreeCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify alias for remark.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Certificate begin time.
        /// </summary>
        [Input("certBeginTime")]
        public Input<string>? CertBeginTime { get; set; }

        /// <summary>
        /// Certificate end time.
        /// </summary>
        [Input("certEndTime")]
        public Input<string>? CertEndTime { get; set; }

        /// <summary>
        /// Certificate private key.
        /// </summary>
        [Input("certificatePrivateKey")]
        public Input<string>? CertificatePrivateKey { get; set; }

        /// <summary>
        /// Certificate public key.
        /// </summary>
        [Input("certificatePublicKey")]
        public Input<string>? CertificatePublicKey { get; set; }

        /// <summary>
        /// Email address.
        /// </summary>
        [Input("contactEmail")]
        public Input<string>? ContactEmail { get; set; }

        /// <summary>
        /// Phone number.
        /// </summary>
        [Input("contactPhone")]
        public Input<string>? ContactPhone { get; set; }

        /// <summary>
        /// Specify CSR encrypt algorithm, only support `RSA` for now.
        /// </summary>
        [Input("csrEncryptAlgo")]
        public Input<string>? CsrEncryptAlgo { get; set; }

        /// <summary>
        /// Specify CSR key parameter, only support `"2048"` for now.
        /// </summary>
        [Input("csrKeyParameter")]
        public Input<string>? CsrKeyParameter { get; set; }

        /// <summary>
        /// Specify CSR key password.
        /// </summary>
        [Input("csrKeyPassword")]
        public Input<string>? CsrKeyPassword { get; set; }

        /// <summary>
        /// Indicates whether the certificate deployable.
        /// </summary>
        [Input("deployable")]
        public Input<bool>? Deployable { get; set; }

        /// <summary>
        /// Specify domain name.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Specify DV authorize method, available values: `DNS_AUTO` - automatic DNS auth, `DNS` - manual DNS auth, `FILE` - auth by file.
        /// </summary>
        [Input("dvAuthMethod")]
        public Input<string>? DvAuthMethod { get; set; }

        /// <summary>
        /// Certificate insert time.
        /// </summary>
        [Input("insertTime")]
        public Input<string>? InsertTime { get; set; }

        /// <summary>
        /// Specify old certificate ID, used for re-apply.
        /// </summary>
        [Input("oldCertificateId")]
        public Input<string>? OldCertificateId { get; set; }

        /// <summary>
        /// Type of package. Only support `"2"` (TrustAsia TLS RSA CA).
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// Product zh name.
        /// </summary>
        [Input("productZhName")]
        public Input<string>? ProductZhName { get; set; }

        /// <summary>
        /// ID of projects which this certification belong to.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Indicates whether the certificate renewable.
        /// </summary>
        [Input("renewable")]
        public Input<bool>? Renewable { get; set; }

        /// <summary>
        /// Certificate status. 0 = Approving, 1 = Approved, 2 = Approve failed, 3 = expired, 4 = DNS record added, 5 = OV/EV Certificate and confirm letter needed, 6 = Order canceling, 7 = Order canceled, 8 = Submitted and confirm letter needed, 9 = Revoking, 10 = Revoked, 11 = re-applying, 12 = Revoke and confirm letter needed, 13 = Free SSL and confirm letter needed.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Certificate status message.
        /// </summary>
        [Input("statusMsg")]
        public Input<string>? StatusMsg { get; set; }

        /// <summary>
        /// Certificate status name.
        /// </summary>
        [Input("statusName")]
        public Input<string>? StatusName { get; set; }

        /// <summary>
        /// Specify validity period in month, only support `"12"` months for now.
        /// </summary>
        [Input("validityPeriod")]
        public Input<string>? ValidityPeriod { get; set; }

        /// <summary>
        /// Vulnerability status.
        /// </summary>
        [Input("vulnerabilityStatus")]
        public Input<string>? VulnerabilityStatus { get; set; }

        public FreeCertificateState()
        {
        }
    }
}
