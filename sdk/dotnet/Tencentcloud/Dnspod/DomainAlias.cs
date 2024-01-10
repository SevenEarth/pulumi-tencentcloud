// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dnspod
{
    /// <summary>
    /// Provides a resource to create a dnspod domain_alias
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
    ///         var domainAlias = new Tencentcloud.Dnspod.DomainAlias("domainAlias", new Tencentcloud.Dnspod.DomainAliasArgs
    ///         {
    ///             Domain = "dnspod.cn",
    ///             DomainAlias = "dnspod.com",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// dnspod domain_alias can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Dnspod/domainAlias:DomainAlias domain_alias domain#domain_alias_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Dnspod/domainAlias:DomainAlias")]
    public partial class DomainAlias : Pulumi.CustomResource
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Domain alias.
        /// </summary>
        [Output("domainAlias")]
        public Output<string> DnspodDomainAlias { get; private set; } = null!;

        /// <summary>
        /// Domain alias ID.
        /// </summary>
        [Output("domainAliasId")]
        public Output<int> DomainAliasId { get; private set; } = null!;


        /// <summary>
        /// Create a DomainAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DomainAlias(string name, DomainAliasArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/domainAlias:DomainAlias", name, args ?? new DomainAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DomainAlias(string name, Input<string> id, DomainAliasState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Dnspod/domainAlias:DomainAlias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DomainAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DomainAlias Get(string name, Input<string> id, DomainAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new DomainAlias(name, id, state, options);
        }
    }

    public sealed class DomainAliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Domain alias.
        /// </summary>
        [Input("domainAlias", required: true)]
        public Input<string> DnspodDomainAlias { get; set; } = null!;

        public DomainAliasArgs()
        {
        }
    }

    public sealed class DomainAliasState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Domain alias.
        /// </summary>
        [Input("domainAlias")]
        public Input<string>? DnspodDomainAlias { get; set; }

        /// <summary>
        /// Domain alias ID.
        /// </summary>
        [Input("domainAliasId")]
        public Input<int>? DomainAliasId { get; set; }

        public DomainAliasState()
        {
        }
    }
}
