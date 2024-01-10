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
    public static class GetCheckDataEngineConfigPairsValidity
    {
        /// <summary>
        /// Use this data source to query detailed information of dlc check_data_engine_config_pairs_validity
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
        ///         var checkDataEngineConfigPairsValidity = Output.Create(Tencentcloud.Dlc.GetCheckDataEngineConfigPairsValidity.InvokeAsync(new Tencentcloud.Dlc.GetCheckDataEngineConfigPairsValidityArgs
        ///         {
        ///             ChildImageVersionId = "d3ftghd4-9a7e-4f64-a3f4-f38507c69742",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCheckDataEngineConfigPairsValidityResult> InvokeAsync(GetCheckDataEngineConfigPairsValidityArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCheckDataEngineConfigPairsValidityResult>("tencentcloud:Dlc/getCheckDataEngineConfigPairsValidity:getCheckDataEngineConfigPairsValidity", args ?? new GetCheckDataEngineConfigPairsValidityArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of dlc check_data_engine_config_pairs_validity
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
        ///         var checkDataEngineConfigPairsValidity = Output.Create(Tencentcloud.Dlc.GetCheckDataEngineConfigPairsValidity.InvokeAsync(new Tencentcloud.Dlc.GetCheckDataEngineConfigPairsValidityArgs
        ///         {
        ///             ChildImageVersionId = "d3ftghd4-9a7e-4f64-a3f4-f38507c69742",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCheckDataEngineConfigPairsValidityResult> Invoke(GetCheckDataEngineConfigPairsValidityInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCheckDataEngineConfigPairsValidityResult>("tencentcloud:Dlc/getCheckDataEngineConfigPairsValidity:getCheckDataEngineConfigPairsValidity", args ?? new GetCheckDataEngineConfigPairsValidityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCheckDataEngineConfigPairsValidityArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Engine Image version id.
        /// </summary>
        [Input("childImageVersionId")]
        public string? ChildImageVersionId { get; set; }

        [Input("dataEngineConfigPairs")]
        private List<Inputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairArgs>? _dataEngineConfigPairs;

        /// <summary>
        /// User-defined parameters.
        /// </summary>
        public List<Inputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairArgs> DataEngineConfigPairs
        {
            get => _dataEngineConfigPairs ?? (_dataEngineConfigPairs = new List<Inputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairArgs>());
            set => _dataEngineConfigPairs = value;
        }

        /// <summary>
        /// Engine major version id. If a minor version id exists, you only need to pass in the minor version id. If it does not exist, the latest minor version id under the current major version will be obtained.
        /// </summary>
        [Input("imageVersionId")]
        public string? ImageVersionId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetCheckDataEngineConfigPairsValidityArgs()
        {
        }
    }

    public sealed class GetCheckDataEngineConfigPairsValidityInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Engine Image version id.
        /// </summary>
        [Input("childImageVersionId")]
        public Input<string>? ChildImageVersionId { get; set; }

        [Input("dataEngineConfigPairs")]
        private InputList<Inputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairInputArgs>? _dataEngineConfigPairs;

        /// <summary>
        /// User-defined parameters.
        /// </summary>
        public InputList<Inputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairInputArgs> DataEngineConfigPairs
        {
            get => _dataEngineConfigPairs ?? (_dataEngineConfigPairs = new InputList<Inputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairInputArgs>());
            set => _dataEngineConfigPairs = value;
        }

        /// <summary>
        /// Engine major version id. If a minor version id exists, you only need to pass in the minor version id. If it does not exist, the latest minor version id under the current major version will be obtained.
        /// </summary>
        [Input("imageVersionId")]
        public Input<string>? ImageVersionId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetCheckDataEngineConfigPairsValidityInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCheckDataEngineConfigPairsValidityResult
    {
        public readonly string? ChildImageVersionId;
        public readonly ImmutableArray<Outputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairResult> DataEngineConfigPairs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? ImageVersionId;
        /// <summary>
        /// Parameter validity: true: valid, false: at least one invalid parameter exists.
        /// </summary>
        public readonly bool IsAvailable;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// Invalid parameter set.
        /// </summary>
        public readonly ImmutableArray<string> UnavailableConfigs;

        [OutputConstructor]
        private GetCheckDataEngineConfigPairsValidityResult(
            string? childImageVersionId,

            ImmutableArray<Outputs.GetCheckDataEngineConfigPairsValidityDataEngineConfigPairResult> dataEngineConfigPairs,

            string id,

            string? imageVersionId,

            bool isAvailable,

            string? resultOutputFile,

            ImmutableArray<string> unavailableConfigs)
        {
            ChildImageVersionId = childImageVersionId;
            DataEngineConfigPairs = dataEngineConfigPairs;
            Id = id;
            ImageVersionId = imageVersionId;
            IsAvailable = isAvailable;
            ResultOutputFile = resultOutputFile;
            UnavailableConfigs = unavailableConfigs;
        }
    }
}
