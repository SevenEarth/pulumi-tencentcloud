// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc
{
    public static class GetDescribeDataEngine
    {
        /// <summary>
        /// Use this data source to query detailed information of dlc describe_data_engine
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
        ///         var describeDataEngine = Output.Create(Tencentcloud.Dlc.GetDescribeDataEngine.InvokeAsync(new Tencentcloud.Dlc.GetDescribeDataEngineArgs
        ///         {
        ///             DataEngineName = "testSpark",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDescribeDataEngineResult> InvokeAsync(GetDescribeDataEngineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDescribeDataEngineResult>("tencentcloud:Dlc/getDescribeDataEngine:getDescribeDataEngine", args ?? new GetDescribeDataEngineArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dlc describe_data_engine
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
        ///         var describeDataEngine = Output.Create(Tencentcloud.Dlc.GetDescribeDataEngine.InvokeAsync(new Tencentcloud.Dlc.GetDescribeDataEngineArgs
        ///         {
        ///             DataEngineName = "testSpark",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDescribeDataEngineResult> Invoke(GetDescribeDataEngineInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDescribeDataEngineResult>("tencentcloud:Dlc/getDescribeDataEngine:getDescribeDataEngine", args ?? new GetDescribeDataEngineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDescribeDataEngineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Engine name.
        /// </summary>
        [Input("dataEngineName", required: true)]
        public string DataEngineName { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetDescribeDataEngineArgs()
        {
        }
    }

    public sealed class GetDescribeDataEngineInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Engine name.
        /// </summary>
        [Input("dataEngineName", required: true)]
        public Input<string> DataEngineName { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetDescribeDataEngineInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDescribeDataEngineResult
    {
        /// <summary>
        /// Engine name.
        /// </summary>
        public readonly string DataEngineName;
        /// <summary>
        /// Data engine details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDescribeDataEngineDataEngineResult> DataEngines;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetDescribeDataEngineResult(
            string dataEngineName,

            ImmutableArray<Outputs.GetDescribeDataEngineDataEngineResult> dataEngines,

            string id,

            string? resultOutputFile)
        {
            DataEngineName = dataEngineName;
            DataEngines = dataEngines;
            Id = id;
            ResultOutputFile = resultOutputFile;
        }
    }
}