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

    public sealed class DatahubTaskTransformsParamFieldChainAnalyseResultGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Operation, DATE system preset - timestamp, CUSTOMIZE customization, MAPPING mapping, JSONPATH.
        /// </summary>
        [Input("operate", required: true)]
        public Input<string> Operate { get; set; } = null!;

        /// <summary>
        /// OriginalValue.
        /// </summary>
        [Input("originalValue")]
        public Input<string>? OriginalValue { get; set; }

        /// <summary>
        /// data type, ORIGINAL, STRING, INT64, FLOAT64, BOOLEAN, MAP, ARRAY.
        /// </summary>
        [Input("schemeType", required: true)]
        public Input<string> SchemeType { get; set; } = null!;

        /// <summary>
        /// value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// VALUE process.
        /// </summary>
        [Input("valueOperate")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateGetArgs>? ValueOperate { get; set; }

        [Input("valueOperates")]
        private InputList<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateGetArgs>? _valueOperates;

        /// <summary>
        /// VALUE process chain.
        /// </summary>
        public InputList<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateGetArgs> ValueOperates
        {
            get => _valueOperates ?? (_valueOperates = new InputList<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateGetArgs>());
            set => _valueOperates = value;
        }

        public DatahubTaskTransformsParamFieldChainAnalyseResultGetArgs()
        {
        }
    }
}
