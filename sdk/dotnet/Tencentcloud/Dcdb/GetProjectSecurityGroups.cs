// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb
{
    public static class GetProjectSecurityGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb project_security_groups
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var projectSecurityGroups = Tencentcloud.Dcdb.GetProjectSecurityGroups.Invoke(new()
        ///     {
        ///         Product = "dcdb",
        ///         ProjectId = 0,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetProjectSecurityGroupsResult> InvokeAsync(GetProjectSecurityGroupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProjectSecurityGroupsResult>("tencentcloud:Dcdb/getProjectSecurityGroups:getProjectSecurityGroups", args ?? new GetProjectSecurityGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb project_security_groups
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var projectSecurityGroups = Tencentcloud.Dcdb.GetProjectSecurityGroups.Invoke(new()
        ///     {
        ///         Product = "dcdb",
        ///         ProjectId = 0,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetProjectSecurityGroupsResult> Invoke(GetProjectSecurityGroupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProjectSecurityGroupsResult>("tencentcloud:Dcdb/getProjectSecurityGroups:getProjectSecurityGroups", args ?? new GetProjectSecurityGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectSecurityGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database engine name. Valid value: `dcdb`.
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public int? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetProjectSecurityGroupsArgs()
        {
        }
        public static new GetProjectSecurityGroupsArgs Empty => new GetProjectSecurityGroupsArgs();
    }

    public sealed class GetProjectSecurityGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Database engine name. Valid value: `dcdb`.
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetProjectSecurityGroupsInvokeArgs()
        {
        }
        public static new GetProjectSecurityGroupsInvokeArgs Empty => new GetProjectSecurityGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetProjectSecurityGroupsResult
    {
        /// <summary>
        /// Security group details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectSecurityGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Product;
        /// <summary>
        /// Project ID.
        /// </summary>
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetProjectSecurityGroupsResult(
            ImmutableArray<Outputs.GetProjectSecurityGroupsGroupResult> groups,

            string id,

            string product,

            int? projectId,

            string? resultOutputFile)
        {
            Groups = groups;
            Id = id;
            Product = product;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
