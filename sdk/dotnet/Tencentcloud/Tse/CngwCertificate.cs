// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse
{
    /// <summary>
    /// Provides a resource to create a tse cngw_certificate
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
    ///         var cngwCertificate = new Tencentcloud.Tse.CngwCertificate("cngwCertificate", new Tencentcloud.Tse.CngwCertificateArgs
    ///         {
    ///             BindDomains = 
    ///             {
    ///                 "example1.com",
    ///             },
    ///             CertId = "vYSQkJ3K",
    ///             GatewayId = "gateway-ddbb709b",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tse cngw_certificate can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tse/cngwCertificate:CngwCertificate cngw_certificate gatewayId#Id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tse/cngwCertificate:CngwCertificate")]
    public partial class CngwCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// Domains of the binding.
        /// </summary>
        [Output("bindDomains")]
        public Output<ImmutableArray<string>> BindDomains { get; private set; } = null!;

        /// <summary>
        /// Certificate ID of ssl platform.
        /// </summary>
        [Output("certId")]
        public Output<string> CertId { get; private set; } = null!;

        /// <summary>
        /// Pem format of certificate.
        /// </summary>
        [Output("crt")]
        public Output<string> Crt { get; private set; } = null!;

        /// <summary>
        /// Gateway ID.
        /// </summary>
        [Output("gatewayId")]
        public Output<string> GatewayId { get; private set; } = null!;

        /// <summary>
        /// Private key of certificate.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Certificate name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a CngwCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CngwCertificate(string name, CngwCertificateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwCertificate:CngwCertificate", name, args ?? new CngwCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CngwCertificate(string name, Input<string> id, CngwCertificateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tse/cngwCertificate:CngwCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CngwCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CngwCertificate Get(string name, Input<string> id, CngwCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new CngwCertificate(name, id, state, options);
        }
    }

    public sealed class CngwCertificateArgs : Pulumi.ResourceArgs
    {
        [Input("bindDomains", required: true)]
        private InputList<string>? _bindDomains;

        /// <summary>
        /// Domains of the binding.
        /// </summary>
        public InputList<string> BindDomains
        {
            get => _bindDomains ?? (_bindDomains = new InputList<string>());
            set => _bindDomains = value;
        }

        /// <summary>
        /// Certificate ID of ssl platform.
        /// </summary>
        [Input("certId", required: true)]
        public Input<string> CertId { get; set; } = null!;

        /// <summary>
        /// Gateway ID.
        /// </summary>
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        /// <summary>
        /// Certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CngwCertificateArgs()
        {
        }
    }

    public sealed class CngwCertificateState : Pulumi.ResourceArgs
    {
        [Input("bindDomains")]
        private InputList<string>? _bindDomains;

        /// <summary>
        /// Domains of the binding.
        /// </summary>
        public InputList<string> BindDomains
        {
            get => _bindDomains ?? (_bindDomains = new InputList<string>());
            set => _bindDomains = value;
        }

        /// <summary>
        /// Certificate ID of ssl platform.
        /// </summary>
        [Input("certId")]
        public Input<string>? CertId { get; set; }

        /// <summary>
        /// Pem format of certificate.
        /// </summary>
        [Input("crt")]
        public Input<string>? Crt { get; set; }

        /// <summary>
        /// Gateway ID.
        /// </summary>
        [Input("gatewayId")]
        public Input<string>? GatewayId { get; set; }

        /// <summary>
        /// Private key of certificate.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public CngwCertificateState()
        {
        }
    }
}