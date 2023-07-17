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
    public static class GetShardSpec
    {
        /// <summary>
        /// Use this data source to query detailed information of dcdb shard_spec
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
        ///         var shardSpec = Output.Create(Tencentcloud.Dcdb.GetShardSpec.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetShardSpecResult> InvokeAsync(GetShardSpecArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetShardSpecResult>("tencentcloud:Dcdb/getShardSpec:getShardSpec", args ?? new GetShardSpecArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dcdb shard_spec
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
        ///         var shardSpec = Output.Create(Tencentcloud.Dcdb.GetShardSpec.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetShardSpecResult> Invoke(GetShardSpecInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetShardSpecResult>("tencentcloud:Dcdb/getShardSpec:getShardSpec", args ?? new GetShardSpecInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetShardSpecArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetShardSpecArgs()
        {
        }
    }

    public sealed class GetShardSpecInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetShardSpecInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetShardSpecResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// list of instance specifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetShardSpecSpecConfigResult> SpecConfigs;

        [OutputConstructor]
        private GetShardSpecResult(
            string id,

            string? resultOutputFile,

            ImmutableArray<Outputs.GetShardSpecSpecConfigResult> specConfigs)
        {
            Id = id;
            ResultOutputFile = resultOutputFile;
            SpecConfigs = specConfigs;
        }
    }
}
