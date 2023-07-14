// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain
{
    public static class GetDbSpaceStatus
    {
        /// <summary>
        /// Use this data source to query detailed information of dbbrain db_space_status
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
        ///         var dbSpaceStatus = Output.Create(Tencentcloud.Dbbrain.GetDbSpaceStatus.InvokeAsync(new Tencentcloud.Dbbrain.GetDbSpaceStatusArgs
        ///         {
        ///             InstanceId = "%s",
        ///             Product = "mysql",
        ///             RangeDays = 7,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDbSpaceStatusResult> InvokeAsync(GetDbSpaceStatusArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDbSpaceStatusResult>("tencentcloud:Dbbrain/getDbSpaceStatus:getDbSpaceStatus", args ?? new GetDbSpaceStatusArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dbbrain db_space_status
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
        ///         var dbSpaceStatus = Output.Create(Tencentcloud.Dbbrain.GetDbSpaceStatus.InvokeAsync(new Tencentcloud.Dbbrain.GetDbSpaceStatusArgs
        ///         {
        ///             InstanceId = "%s",
        ///             Product = "mysql",
        ///             RangeDays = 7,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDbSpaceStatusResult> Invoke(GetDbSpaceStatusInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDbSpaceStatusResult>("tencentcloud:Dbbrain/getDbSpaceStatus:getDbSpaceStatus", args ?? new GetDbSpaceStatusInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDbSpaceStatusArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
        /// </summary>
        [Input("product")]
        public string? Product { get; set; }

        /// <summary>
        /// The number of days in the time period, the deadline is the current day, and the default is 7 days.
        /// </summary>
        [Input("rangeDays")]
        public int? RangeDays { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDbSpaceStatusArgs()
        {
        }
    }

    public sealed class GetDbSpaceStatusInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Service product type, supported values include: mysql - cloud database MySQL, cynosdb - cloud database CynosDB for MySQL, the default is mysql.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// The number of days in the time period, the deadline is the current day, and the default is 7 days.
        /// </summary>
        [Input("rangeDays")]
        public Input<int>? RangeDays { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDbSpaceStatusInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDbSpaceStatusResult
    {
        /// <summary>
        /// Estimated number of days available.
        /// </summary>
        public readonly int AvailableDays;
        /// <summary>
        /// Disk growth (MB).
        /// </summary>
        public readonly int Growth;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        public readonly string? Product;
        public readonly int? RangeDays;
        /// <summary>
        /// Disk remaining (MB).
        /// </summary>
        public readonly int Remain;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Total disk size (MB).
        /// </summary>
        public readonly int Total;

        [OutputConstructor]
        private GetDbSpaceStatusResult(
            int availableDays,

            int growth,

            string id,

            string instanceId,

            string? product,

            int? rangeDays,

            int remain,

            string? resultOutputFile,

            int total)
        {
            AvailableDays = availableDays;
            Growth = growth;
            Id = id;
            InstanceId = instanceId;
            Product = product;
            RangeDays = rangeDays;
            Remain = remain;
            ResultOutputFile = resultOutputFile;
            Total = total;
        }
    }
}
