// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Wedata
{
    /// <summary>
    /// Provides a resource to create a wedata script
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
    ///     var example = new Tencentcloud.Wedata.Script("example", new()
    ///     {
    ///         BucketName = "wedata-demo-1257305158",
    ///         FileExtensionType = "sql",
    ///         FilePath = "/datastudio/project/tf_example.sql",
    ///         ProjectId = "1470575647377821696",
    ///         Region = "ap-guangzhou",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// wedata script can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Wedata/script:Script example 1470575647377821696#/datastudio/project/tf_example.sql#4147824b-7ba2-432b-8a8b-7e747594c926
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Wedata/script:Script")]
    public partial class Script : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cos bucket name.
        /// </summary>
        [Output("bucketName")]
        public Output<string?> BucketName { get; private set; } = null!;

        /// <summary>
        /// File Extension Type:jar, sql, zip, py, sh, txt, di, dg, pyspark, kjb, ktr, csv.
        /// </summary>
        [Output("fileExtensionType")]
        public Output<string?> FileExtensionType { get; private set; } = null!;

        /// <summary>
        /// Cos file path:/datastudio/project/projectId/.
        /// </summary>
        [Output("filePath")]
        public Output<string?> FilePath { get; private set; } = null!;

        /// <summary>
        /// Project id.
        /// </summary>
        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Cos region.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;


        /// <summary>
        /// Create a Script resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Script(string name, ScriptArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Wedata/script:Script", name, args ?? new ScriptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Script(string name, Input<string> id, ScriptState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Wedata/script:Script", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Script resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Script Get(string name, Input<string> id, ScriptState? state = null, CustomResourceOptions? options = null)
        {
            return new Script(name, id, state, options);
        }
    }

    public sealed class ScriptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cos bucket name.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// File Extension Type:jar, sql, zip, py, sh, txt, di, dg, pyspark, kjb, ktr, csv.
        /// </summary>
        [Input("fileExtensionType")]
        public Input<string>? FileExtensionType { get; set; }

        /// <summary>
        /// Cos file path:/datastudio/project/projectId/.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// Project id.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Cos region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ScriptArgs()
        {
        }
        public static new ScriptArgs Empty => new ScriptArgs();
    }

    public sealed class ScriptState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cos bucket name.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// File Extension Type:jar, sql, zip, py, sh, txt, di, dg, pyspark, kjb, ktr, csv.
        /// </summary>
        [Input("fileExtensionType")]
        public Input<string>? FileExtensionType { get; set; }

        /// <summary>
        /// Cos file path:/datastudio/project/projectId/.
        /// </summary>
        [Input("filePath")]
        public Input<string>? FilePath { get; set; }

        /// <summary>
        /// Project id.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Cos region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public ScriptState()
        {
        }
        public static new ScriptState Empty => new ScriptState();
    }
}
