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
    /// Provides a resource to create a scf function_alias
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
    ///         // by route
    ///         var functionAlias = new Tencentcloud.Scf.FunctionAlias("functionAlias", new Tencentcloud.Scf.FunctionAliasArgs
    ///         {
    ///             Description = "matchs for test 12312312",
    ///             FunctionName = "keep-1676351130",
    ///             FunctionVersion = "3",
    ///             Namespace = "default",
    ///             RoutingConfig = new Tencentcloud.Scf.Inputs.FunctionAliasRoutingConfigArgs
    ///             {
    ///                 AdditionalVersionMatches = 
    ///                 {
    ///                     new Tencentcloud.Scf.Inputs.FunctionAliasRoutingConfigAdditionalVersionMatchArgs
    ///                     {
    ///                         Expression = "testuser",
    ///                         Key = "invoke.headers.User",
    ///                         Method = "exact",
    ///                         Version = "2",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// scf function_alias can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Scf/functionAlias:FunctionAlias function_alias namespace#functionName#name
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Scf/functionAlias:FunctionAlias")]
    public partial class FunctionAlias : Pulumi.CustomResource
    {
        /// <summary>
        /// Alias description information.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Function name.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// Master version pointed to by the alias.
        /// </summary>
        [Output("functionVersion")]
        public Output<string> FunctionVersion { get; private set; } = null!;

        /// <summary>
        /// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Request routing configuration of alias.
        /// </summary>
        [Output("routingConfig")]
        public Output<Outputs.FunctionAliasRoutingConfig?> RoutingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionAlias(string name, FunctionAliasArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/functionAlias:FunctionAlias", name, args ?? new FunctionAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionAlias(string name, Input<string> id, FunctionAliasState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/functionAlias:FunctionAlias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FunctionAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionAlias Get(string name, Input<string> id, FunctionAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionAlias(name, id, state, options);
        }
    }

    public sealed class FunctionAliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias description information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Master version pointed to by the alias.
        /// </summary>
        [Input("functionVersion", required: true)]
        public Input<string> FunctionVersion { get; set; } = null!;

        /// <summary>
        /// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Request routing configuration of alias.
        /// </summary>
        [Input("routingConfig")]
        public Input<Inputs.FunctionAliasRoutingConfigArgs>? RoutingConfig { get; set; }

        public FunctionAliasArgs()
        {
        }
    }

    public sealed class FunctionAliasState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias description information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Function name.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// Master version pointed to by the alias.
        /// </summary>
        [Input("functionVersion")]
        public Input<string>? FunctionVersion { get; set; }

        /// <summary>
        /// Alias name, which must be unique in the function, can contain 1 to 64 letters, digits, _, and -, and must begin with a letter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Function namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Request routing configuration of alias.
        /// </summary>
        [Input("routingConfig")]
        public Input<Inputs.FunctionAliasRoutingConfigGetArgs>? RoutingConfig { get; set; }

        public FunctionAliasState()
        {
        }
    }
}
