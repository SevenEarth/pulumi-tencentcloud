// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css
{
    /// <summary>
    /// Provides a resource to verify the domain ownership by specified way when DomainNeedVerifyOwner failed in domain creation.
    /// 
    /// ## Example Usage
    /// ### dnsCheck way:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var dnsCheck = new Tencentcloud.Css.AuthenticateDomainOwnerOperation("dnsCheck", new Tencentcloud.Css.AuthenticateDomainOwnerOperationArgs
    ///         {
    ///             DomainName = "your_domain_name",
    ///             VerifyType = "dnsCheck",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### fileCheck way:
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fileCheck = new Tencentcloud.Css.AuthenticateDomainOwnerOperation("fileCheck", new Tencentcloud.Css.AuthenticateDomainOwnerOperationArgs
    ///         {
    ///             DomainName = "your_domain_name",
    ///             VerifyType = "fileCheck",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Css/authenticateDomainOwnerOperation:AuthenticateDomainOwnerOperation")]
    public partial class AuthenticateDomainOwnerOperation : Pulumi.CustomResource
    {
        /// <summary>
        /// The domain name to verify.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// Authentication type. Possible values:`dnsCheck`: Immediately verify whether the resolution record of the configured dns is consistent with the content to be verified, and save the record if successful.`fileCheck`: Immediately verify whether the web file is consistent with the content to be verified, and save the record if successful.`dbCheck`: Check if authentication has been successful.
        /// </summary>
        [Output("verifyType")]
        public Output<string?> VerifyType { get; private set; } = null!;


        /// <summary>
        /// Create a AuthenticateDomainOwnerOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthenticateDomainOwnerOperation(string name, AuthenticateDomainOwnerOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/authenticateDomainOwnerOperation:AuthenticateDomainOwnerOperation", name, args ?? new AuthenticateDomainOwnerOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthenticateDomainOwnerOperation(string name, Input<string> id, AuthenticateDomainOwnerOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Css/authenticateDomainOwnerOperation:AuthenticateDomainOwnerOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthenticateDomainOwnerOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthenticateDomainOwnerOperation Get(string name, Input<string> id, AuthenticateDomainOwnerOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthenticateDomainOwnerOperation(name, id, state, options);
        }
    }

    public sealed class AuthenticateDomainOwnerOperationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name to verify.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// Authentication type. Possible values:`dnsCheck`: Immediately verify whether the resolution record of the configured dns is consistent with the content to be verified, and save the record if successful.`fileCheck`: Immediately verify whether the web file is consistent with the content to be verified, and save the record if successful.`dbCheck`: Check if authentication has been successful.
        /// </summary>
        [Input("verifyType")]
        public Input<string>? VerifyType { get; set; }

        public AuthenticateDomainOwnerOperationArgs()
        {
        }
    }

    public sealed class AuthenticateDomainOwnerOperationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The domain name to verify.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// Authentication type. Possible values:`dnsCheck`: Immediately verify whether the resolution record of the configured dns is consistent with the content to be verified, and save the record if successful.`fileCheck`: Immediately verify whether the web file is consistent with the content to be verified, and save the record if successful.`dbCheck`: Check if authentication has been successful.
        /// </summary>
        [Input("verifyType")]
        public Input<string>? VerifyType { get; set; }

        public AuthenticateDomainOwnerOperationState()
        {
        }
    }
}
