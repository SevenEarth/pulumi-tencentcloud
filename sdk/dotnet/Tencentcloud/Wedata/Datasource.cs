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
    /// Provides a resource to create a wedata datasource
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Tencentcloud.Wedata.Datasource("example", new Tencentcloud.Wedata.DatasourceArgs
    ///         {
    ///             Category = "DB",
    ///             Type = "MYSQL",
    ///             OwnerProjectId = "1612982498218618880",
    ///             OwnerProjectName = "project_demo",
    ///             OwnerProjectIdent = "体验项目",
    ///             Description = "description.",
    ///             Display = "tf_example_demo",
    ///             Status = 1,
    ///             CosBucket = "wedata-agent-sh-1257305158",
    ///             CosRegion = "ap-shanghai",
    ///             Params = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "connectType", "public" },
    ///                 { "authorityType", "true" },
    ///                 { "deployType", "CONNSTR_PUBLICDB" },
    ///                 { "url", "jdbc:mysql://1.1.1.1:8080/database" },
    ///                 { "username", "root" },
    ///                 { "password", "password" },
    ///                 { "type", "MYSQL" },
    ///             }),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Wedata/datasource:Datasource")]
    public partial class Datasource : Pulumi.CustomResource
    {
        /// <summary>
        /// BizParams.
        /// </summary>
        [Output("bizParams")]
        public Output<string?> BizParams { get; private set; } = null!;

        /// <summary>
        /// DataSource Category.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// ClusterId.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Collect.
        /// </summary>
        [Output("collect")]
        public Output<string> Collect { get; private set; } = null!;

        /// <summary>
        /// COSBucket.
        /// </summary>
        [Output("cosBucket")]
        public Output<string?> CosBucket { get; private set; } = null!;

        /// <summary>
        /// Cos region.
        /// </summary>
        [Output("cosRegion")]
        public Output<string?> CosRegion { get; private set; } = null!;

        /// <summary>
        /// Dbname.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Display.
        /// </summary>
        [Output("display")]
        public Output<string?> Display { get; private set; } = null!;

        /// <summary>
        /// Instance.
        /// </summary>
        [Output("instance")]
        public Output<string?> Instance { get; private set; } = null!;

        /// <summary>
        /// DataSource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Owner projectId.
        /// </summary>
        [Output("ownerProjectId")]
        public Output<string> OwnerProjectId { get; private set; } = null!;

        /// <summary>
        /// Owner Project Ident.
        /// </summary>
        [Output("ownerProjectIdent")]
        public Output<string> OwnerProjectIdent { get; private set; } = null!;

        /// <summary>
        /// Owner project name.
        /// </summary>
        [Output("ownerProjectName")]
        public Output<string> OwnerProjectName { get; private set; } = null!;

        /// <summary>
        /// Params.
        /// </summary>
        [Output("params")]
        public Output<string> Params { get; private set; } = null!;

        /// <summary>
        /// Params Out.
        /// </summary>
        [Output("paramsOut")]
        public Output<string> ParamsOut { get; private set; } = null!;

        /// <summary>
        /// Status.
        /// </summary>
        [Output("status")]
        public Output<int?> Status { get; private set; } = null!;

        /// <summary>
        /// DataSource Type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Datasource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Datasource(string name, DatasourceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Wedata/datasource:Datasource", name, args ?? new DatasourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Datasource(string name, Input<string> id, DatasourceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Wedata/datasource:Datasource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Datasource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Datasource Get(string name, Input<string> id, DatasourceState? state = null, CustomResourceOptions? options = null)
        {
            return new Datasource(name, id, state, options);
        }
    }

    public sealed class DatasourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// BizParams.
        /// </summary>
        [Input("bizParams")]
        public Input<string>? BizParams { get; set; }

        /// <summary>
        /// DataSource Category.
        /// </summary>
        [Input("category", required: true)]
        public Input<string> Category { get; set; } = null!;

        /// <summary>
        /// ClusterId.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Collect.
        /// </summary>
        [Input("collect")]
        public Input<string>? Collect { get; set; }

        /// <summary>
        /// COSBucket.
        /// </summary>
        [Input("cosBucket")]
        public Input<string>? CosBucket { get; set; }

        /// <summary>
        /// Cos region.
        /// </summary>
        [Input("cosRegion")]
        public Input<string>? CosRegion { get; set; }

        /// <summary>
        /// Dbname.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display.
        /// </summary>
        [Input("display")]
        public Input<string>? Display { get; set; }

        /// <summary>
        /// Instance.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// DataSource Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Owner projectId.
        /// </summary>
        [Input("ownerProjectId", required: true)]
        public Input<string> OwnerProjectId { get; set; } = null!;

        /// <summary>
        /// Owner Project Ident.
        /// </summary>
        [Input("ownerProjectIdent", required: true)]
        public Input<string> OwnerProjectIdent { get; set; } = null!;

        /// <summary>
        /// Owner project name.
        /// </summary>
        [Input("ownerProjectName", required: true)]
        public Input<string> OwnerProjectName { get; set; } = null!;

        /// <summary>
        /// Params.
        /// </summary>
        [Input("params", required: true)]
        public Input<string> Params { get; set; } = null!;

        /// <summary>
        /// Status.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// DataSource Type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatasourceArgs()
        {
        }
    }

    public sealed class DatasourceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// BizParams.
        /// </summary>
        [Input("bizParams")]
        public Input<string>? BizParams { get; set; }

        /// <summary>
        /// DataSource Category.
        /// </summary>
        [Input("category")]
        public Input<string>? Category { get; set; }

        /// <summary>
        /// ClusterId.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Collect.
        /// </summary>
        [Input("collect")]
        public Input<string>? Collect { get; set; }

        /// <summary>
        /// COSBucket.
        /// </summary>
        [Input("cosBucket")]
        public Input<string>? CosBucket { get; set; }

        /// <summary>
        /// Cos region.
        /// </summary>
        [Input("cosRegion")]
        public Input<string>? CosRegion { get; set; }

        /// <summary>
        /// Dbname.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display.
        /// </summary>
        [Input("display")]
        public Input<string>? Display { get; set; }

        /// <summary>
        /// Instance.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        /// <summary>
        /// DataSource Name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Owner projectId.
        /// </summary>
        [Input("ownerProjectId")]
        public Input<string>? OwnerProjectId { get; set; }

        /// <summary>
        /// Owner Project Ident.
        /// </summary>
        [Input("ownerProjectIdent")]
        public Input<string>? OwnerProjectIdent { get; set; }

        /// <summary>
        /// Owner project name.
        /// </summary>
        [Input("ownerProjectName")]
        public Input<string>? OwnerProjectName { get; set; }

        /// <summary>
        /// Params.
        /// </summary>
        [Input("params")]
        public Input<string>? Params { get; set; }

        /// <summary>
        /// Params Out.
        /// </summary>
        [Input("paramsOut")]
        public Input<string>? ParamsOut { get; set; }

        /// <summary>
        /// Status.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// DataSource Type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public DatasourceState()
        {
        }
    }
}
