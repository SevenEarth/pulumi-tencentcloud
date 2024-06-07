// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf
{
    /// <summary>
    /// Provides a resource to create a scf function_version
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var functionVersion = new Tencentcloud.Scf.FunctionVersion("functionVersion", new()
    ///     {
    ///         Description = "for-terraform-test",
    ///         FunctionName = "keep-1676351130",
    ///         Namespace = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// scf function_version can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Scf/functionVersion:FunctionVersion function_version functionName#namespace#functionVersion
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Scf/functionVersion:FunctionVersion")]
    public partial class FunctionVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Function description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the released function.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// Version of the released function.
        /// </summary>
        [Output("functionVersion")]
        public Output<string> ScfFunctionVersion { get; private set; } = null!;

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionVersion(string name, FunctionVersionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/functionVersion:FunctionVersion", name, args ?? new FunctionVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionVersion(string name, Input<string> id, FunctionVersionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/functionVersion:FunctionVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionVersion Get(string name, Input<string> id, FunctionVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionVersion(name, id, state, options);
        }
    }

    public sealed class FunctionVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Function description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the released function.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public FunctionVersionArgs()
        {
        }
        public static new FunctionVersionArgs Empty => new FunctionVersionArgs();
    }

    public sealed class FunctionVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Function description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the released function.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// Version of the released function.
        /// </summary>
        [Input("functionVersion")]
        public Input<string>? ScfFunctionVersion { get; set; }

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public FunctionVersionState()
        {
        }
        public static new FunctionVersionState Empty => new FunctionVersionState();
    }
}
