// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTransformsParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// data process.
        /// </summary>
        [Input("batchAnalyse")]
        public Input<Inputs.DatahubTaskTransformsParamBatchAnalyseGetArgs>? BatchAnalyse { get; set; }

        /// <summary>
        /// Raw data.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// fail process.
        /// </summary>
        [Input("failureParam")]
        public Input<Inputs.DatahubTaskTransformsParamFailureParamGetArgs>? FailureParam { get; set; }

        [Input("fieldChains", required: true)]
        private InputList<Inputs.DatahubTaskTransformsParamFieldChainGetArgs>? _fieldChains;

        /// <summary>
        /// processing chain.
        /// </summary>
        public InputList<Inputs.DatahubTaskTransformsParamFieldChainGetArgs> FieldChains
        {
            get => _fieldChains ?? (_fieldChains = new InputList<Inputs.DatahubTaskTransformsParamFieldChainGetArgs>());
            set => _fieldChains = value;
        }

        [Input("filterParams")]
        private InputList<Inputs.DatahubTaskTransformsParamFilterParamGetArgs>? _filterParams;

        /// <summary>
        /// filter.
        /// </summary>
        public InputList<Inputs.DatahubTaskTransformsParamFilterParamGetArgs> FilterParams
        {
            get => _filterParams ?? (_filterParams = new InputList<Inputs.DatahubTaskTransformsParamFilterParamGetArgs>());
            set => _filterParams = value;
        }

        /// <summary>
        /// Whether to keep the data source Topic metadata information (source Topic, Partition, Offset), the default is false.
        /// </summary>
        [Input("keepMetadata")]
        public Input<bool>? KeepMetadata { get; set; }

        /// <summary>
        /// output format, JSON, ROW, default JSON.
        /// </summary>
        [Input("outputFormat")]
        public Input<string>? OutputFormat { get; set; }

        /// <summary>
        /// result.
        /// </summary>
        [Input("result")]
        public Input<string>? Result { get; set; }

        /// <summary>
        /// The output format is ROW Required.
        /// </summary>
        [Input("rowParam")]
        public Input<Inputs.DatahubTaskTransformsParamRowParamGetArgs>? RowParam { get; set; }

        /// <summary>
        /// data source.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        public DatahubTaskTransformsParamGetArgs()
        {
        }
    }
}
