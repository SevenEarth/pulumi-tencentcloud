// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdcpg
{
    public static class GetInstances
    {
        /// <summary>
        /// Use this data source to query detailed information of tdcpg instances.
        /// 
        /// &gt; **NOTE:** This data source is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.
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
        ///     var instances = Tencentcloud.Tdcpg.GetInstances.Invoke(new()
        ///     {
        ///         ClusterId = "",
        ///         InstanceId = "",
        ///         InstanceName = "",
        ///         InstanceType = "",
        ///         Status = "",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetInstancesResult> InvokeAsync(GetInstancesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstancesResult>("tencentcloud:Tdcpg/getInstances:getInstances", args ?? new GetInstancesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of tdcpg instances.
        /// 
        /// &gt; **NOTE:** This data source is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.
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
        ///     var instances = Tencentcloud.Tdcpg.GetInstances.Invoke(new()
        ///     {
        ///         ClusterId = "",
        ///         InstanceId = "",
        ///         InstanceName = "",
        ///         InstanceType = "",
        ///         Status = "",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetInstancesResult> Invoke(GetInstancesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstancesResult>("tencentcloud:Tdcpg/getInstances:getInstances", args ?? new GetInstancesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstancesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// instance name.
        /// </summary>
        [Input("instanceName")]
        public string? InstanceName { get; set; }

        /// <summary>
        /// instance type.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// instance status.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        public GetInstancesArgs()
        {
        }
        public static new GetInstancesArgs Empty => new GetInstancesArgs();
    }

    public sealed class GetInstancesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// instance status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public GetInstancesInvokeArgs()
        {
        }
        public static new GetInstancesInvokeArgs Empty => new GetInstancesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstancesResult
    {
        /// <summary>
        /// cluster id.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// instance id.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// instance name.
        /// </summary>
        public readonly string? InstanceName;
        /// <summary>
        /// instance type.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// instance list.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstancesListResult> Lists;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// status.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private GetInstancesResult(
            string clusterId,

            string id,

            string? instanceId,

            string? instanceName,

            string? instanceType,

            ImmutableArray<Outputs.GetInstancesListResult> lists,

            string? resultOutputFile,

            string? status)
        {
            ClusterId = clusterId;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceType = instanceType;
            Lists = lists;
            ResultOutputFile = resultOutputFile;
            Status = status;
        }
    }
}
