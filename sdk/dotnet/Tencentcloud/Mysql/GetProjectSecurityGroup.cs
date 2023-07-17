// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    public static class GetProjectSecurityGroup
    {
        /// <summary>
        /// Use this data source to query detailed information of mysql project_security_group
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var projectSecurityGroup = Output.Create(Tencentcloud.Mysql.GetProjectSecurityGroup.InvokeAsync(new Tencentcloud.Mysql.GetProjectSecurityGroupArgs
        ///         {
        ///             ProjectId = 1250480,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProjectSecurityGroupResult> InvokeAsync(GetProjectSecurityGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectSecurityGroupResult>("tencentcloud:Mysql/getProjectSecurityGroup:getProjectSecurityGroup", args ?? new GetProjectSecurityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mysql project_security_group
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var projectSecurityGroup = Output.Create(Tencentcloud.Mysql.GetProjectSecurityGroup.InvokeAsync(new Tencentcloud.Mysql.GetProjectSecurityGroupArgs
        ///         {
        ///             ProjectId = 1250480,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProjectSecurityGroupResult> Invoke(GetProjectSecurityGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProjectSecurityGroupResult>("tencentcloud:Mysql/getProjectSecurityGroup:getProjectSecurityGroup", args ?? new GetProjectSecurityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProjectSecurityGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// project id.
        /// </summary>
        [Input("projectId")]
        public int? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetProjectSecurityGroupArgs()
        {
        }
    }

    public sealed class GetProjectSecurityGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// project id.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetProjectSecurityGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectSecurityGroupResult
    {
        /// <summary>
        /// Security group details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetProjectSecurityGroupGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// project id.
        /// </summary>
        public readonly int? ProjectId;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetProjectSecurityGroupResult(
            ImmutableArray<Outputs.GetProjectSecurityGroupGroupResult> groups,

            string id,

            int? projectId,

            string? resultOutputFile)
        {
            Groups = groups;
            Id = id;
            ProjectId = projectId;
            ResultOutputFile = resultOutputFile;
        }
    }
}
