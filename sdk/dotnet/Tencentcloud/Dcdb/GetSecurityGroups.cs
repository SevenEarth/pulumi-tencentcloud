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
    public static class GetSecurityGroups
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb securityGroups
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
        ///     var securityGroups = Tencentcloud.Dcdb.GetSecurityGroups.Invoke(new()
        ///     {
        ///         InstanceId = "your_instance_id",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetSecurityGroupsResult> InvokeAsync(GetSecurityGroupsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupsResult>("tencentcloud:Dcdb/getSecurityGroups:getSecurityGroups", args ?? new GetSecurityGroupsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb securityGroups
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
        ///     var securityGroups = Tencentcloud.Dcdb.GetSecurityGroups.Invoke(new()
        ///     {
        ///         InstanceId = "your_instance_id",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetSecurityGroupsResult> Invoke(GetSecurityGroupsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupsResult>("tencentcloud:Dcdb/getSecurityGroups:getSecurityGroups", args ?? new GetSecurityGroupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetSecurityGroupsArgs()
        {
        }
        public static new GetSecurityGroupsArgs Empty => new GetSecurityGroupsArgs();
    }

    public sealed class GetSecurityGroupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetSecurityGroupsInvokeArgs()
        {
        }
        public static new GetSecurityGroupsInvokeArgs Empty => new GetSecurityGroupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityGroupsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// security group list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSecurityGroupsListResult> Lists;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetSecurityGroupsResult(
            string id,

            string instanceId,

            ImmutableArray<Outputs.GetSecurityGroupsListResult> lists,

            string? resultOutputFile)
        {
            Id = id;
            InstanceId = instanceId;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
        }
    }
}
