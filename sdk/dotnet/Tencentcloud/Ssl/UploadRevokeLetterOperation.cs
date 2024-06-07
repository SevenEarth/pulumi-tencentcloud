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
    /// Provides a resource to create a ssl upload_revoke_letter
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System;
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// 	
    /// string ReadFileBase64(string path) 
    /// {
    ///     return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)));
    /// }
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var uploadRevokeLetter = new Tencentcloud.Ssl.UploadRevokeLetterOperation("uploadRevokeLetter", new()
    ///     {
    ///         CertificateId = "8xRYdDlc",
    ///         RevokeLetter = ReadFileBase64("./c.pdf"),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ssl upload_revoke_letter can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Ssl/uploadRevokeLetterOperation:UploadRevokeLetterOperation upload_revoke_letter upload_revoke_letter_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Ssl/uploadRevokeLetterOperation:UploadRevokeLetterOperation")]
    public partial class UploadRevokeLetterOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// The format of the base64-encoded certificate confirmation letter file should be jpg, jpeg, png, or pdf, and the size should be between 1kb and 1.4M.
        /// </summary>
        [Output("revokeLetter")]
        public Output<string> RevokeLetter { get; private set; } = null!;


        /// <summary>
        /// Create a UploadRevokeLetterOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UploadRevokeLetterOperation(string name, UploadRevokeLetterOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/uploadRevokeLetterOperation:UploadRevokeLetterOperation", name, args ?? new UploadRevokeLetterOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UploadRevokeLetterOperation(string name, Input<string> id, UploadRevokeLetterOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Ssl/uploadRevokeLetterOperation:UploadRevokeLetterOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UploadRevokeLetterOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UploadRevokeLetterOperation Get(string name, Input<string> id, UploadRevokeLetterOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new UploadRevokeLetterOperation(name, id, state, options);
        }
    }

    public sealed class UploadRevokeLetterOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        /// <summary>
        /// The format of the base64-encoded certificate confirmation letter file should be jpg, jpeg, png, or pdf, and the size should be between 1kb and 1.4M.
        /// </summary>
        [Input("revokeLetter", required: true)]
        public Input<string> RevokeLetter { get; set; } = null!;

        public UploadRevokeLetterOperationArgs()
        {
        }
        public static new UploadRevokeLetterOperationArgs Empty => new UploadRevokeLetterOperationArgs();
    }

    public sealed class UploadRevokeLetterOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate ID.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// The format of the base64-encoded certificate confirmation letter file should be jpg, jpeg, png, or pdf, and the size should be between 1kb and 1.4M.
        /// </summary>
        [Input("revokeLetter")]
        public Input<string>? RevokeLetter { get; set; }

        public UploadRevokeLetterOperationState()
        {
        }
        public static new UploadRevokeLetterOperationState Empty => new UploadRevokeLetterOperationState();
    }
}
