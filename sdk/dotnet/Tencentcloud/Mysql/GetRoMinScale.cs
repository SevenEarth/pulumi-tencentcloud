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
    public static class GetRoMinScale
    {
        /// <summary>
        /// Use this data source to query detailed information of mysql ro_min_scale
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
        ///         var roMinScale = Output.Create(Tencentcloud.Mysql.GetRoMinScale.InvokeAsync(new Tencentcloud.Mysql.GetRoMinScaleArgs
        ///         {
        ///             MasterInstanceId = "cdb-fitq5t9h",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRoMinScaleResult> InvokeAsync(GetRoMinScaleArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRoMinScaleResult>("tencentcloud:Mysql/getRoMinScale:getRoMinScale", args ?? new GetRoMinScaleArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mysql ro_min_scale
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
        ///         var roMinScale = Output.Create(Tencentcloud.Mysql.GetRoMinScale.InvokeAsync(new Tencentcloud.Mysql.GetRoMinScaleArgs
        ///         {
        ///             MasterInstanceId = "cdb-fitq5t9h",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRoMinScaleResult> Invoke(GetRoMinScaleInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRoMinScaleResult>("tencentcloud:Mysql/getRoMinScale:getRoMinScale", args ?? new GetRoMinScaleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRoMinScaleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The primary instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the RoInstanceId parameter cannot be empty at the same time. Note that when the input parameter contains RoInstanceId, the return value is the minimum specification when the read-only instance is upgraded; when the input parameter only contains MasterInstanceId, the return value is the minimum specification when the read-only instance is purchased.
        /// </summary>
        [Input("masterInstanceId")]
        public string? MasterInstanceId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// The read-only instance ID, in the format: cdbro-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the MasterInstanceId parameter cannot be empty at the same time.
        /// </summary>
        [Input("roInstanceId")]
        public string? RoInstanceId { get; set; }

        public GetRoMinScaleArgs()
        {
        }
    }

    public sealed class GetRoMinScaleInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The primary instance ID, in the format: cdb-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the RoInstanceId parameter cannot be empty at the same time. Note that when the input parameter contains RoInstanceId, the return value is the minimum specification when the read-only instance is upgraded; when the input parameter only contains MasterInstanceId, the return value is the minimum specification when the read-only instance is purchased.
        /// </summary>
        [Input("masterInstanceId")]
        public Input<string>? MasterInstanceId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// The read-only instance ID, in the format: cdbro-c1nl9rpv, is the same as the instance ID displayed on the cloud database console page. This parameter and the MasterInstanceId parameter cannot be empty at the same time.
        /// </summary>
        [Input("roInstanceId")]
        public Input<string>? RoInstanceId { get; set; }

        public GetRoMinScaleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRoMinScaleResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? MasterInstanceId;
        /// <summary>
        /// Memory specification size, unit: MB.
        /// </summary>
        public readonly int Memory;
        public readonly string? ResultOutputFile;
        public readonly string? RoInstanceId;
        /// <summary>
        /// Disk specification size, unit: GB.
        /// </summary>
        public readonly int Volume;

        [OutputConstructor]
        private GetRoMinScaleResult(
            string id,

            string? masterInstanceId,

            int memory,

            string? resultOutputFile,

            string? roInstanceId,

            int volume)
        {
            Id = id;
            MasterInstanceId = masterInstanceId;
            Memory = memory;
            ResultOutputFile = resultOutputFile;
            RoInstanceId = roInstanceId;
            Volume = volume;
        }
    }
}
