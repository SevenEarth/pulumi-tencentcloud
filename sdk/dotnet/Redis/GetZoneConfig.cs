// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Redis
{
    public static class GetZoneConfig
    {
        /// <summary>
        /// Use this data source to query which instance types of Redis are available in a specific region.
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
        ///         var redislab = Output.Create(Tencentcloud.Redis.GetZoneConfig.InvokeAsync(new Tencentcloud.Redis.GetZoneConfigArgs
        ///         {
        ///             Region = "ap-hongkong",
        ///             ResultOutputFile = "/temp/mytestpath",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZoneConfigResult> InvokeAsync(GetZoneConfigArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZoneConfigResult>("tencentcloud:Redis/getZoneConfig:getZoneConfig", args ?? new GetZoneConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query which instance types of Redis are available in a specific region.
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
        ///         var redislab = Output.Create(Tencentcloud.Redis.GetZoneConfig.InvokeAsync(new Tencentcloud.Redis.GetZoneConfigArgs
        ///         {
        ///             Region = "ap-hongkong",
        ///             ResultOutputFile = "/temp/mytestpath",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZoneConfigResult> Invoke(GetZoneConfigInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetZoneConfigResult>("tencentcloud:Redis/getZoneConfig:getZoneConfig", args ?? new GetZoneConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZoneConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of a region. If this value is not set, the current region getting from provider's configuration will be used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Instance type ID.
        /// </summary>
        [Input("typeId")]
        public int? TypeId { get; set; }

        public GetZoneConfigArgs()
        {
        }
    }

    public sealed class GetZoneConfigInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of a region. If this value is not set, the current region getting from provider's configuration will be used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Instance type ID.
        /// </summary>
        [Input("typeId")]
        public Input<int>? TypeId { get; set; }

        public GetZoneConfigInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetZoneConfigResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of zone. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZoneConfigListResult> Lists;
        public readonly string? Region;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Instance type. Which redis type supports in this zone.
        /// </summary>
        public readonly int? TypeId;

        [OutputConstructor]
        private GetZoneConfigResult(
            string id,

            ImmutableArray<Outputs.GetZoneConfigListResult> lists,

            string? region,

            string? resultOutputFile,

            int? typeId)
        {
            Id = id;
            Lists = lists;
            Region = region;
            ResultOutputFile = resultOutputFile;
            TypeId = typeId;
        }
    }
}
