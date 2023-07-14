// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class DatahubTaskTransformsParamFieldChainSMT
    {
        /// <summary>
        /// KEY.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Operation, DATE system preset - timestamp, CUSTOMIZE customization, MAPPING mapping, JSONPATH.
        /// </summary>
        public readonly string Operate;
        /// <summary>
        /// OriginalValue.
        /// </summary>
        public readonly string? OriginalValue;
        /// <summary>
        /// data type, ORIGINAL, STRING, INT64, FLOAT64, BOOLEAN, MAP, ARRAY.
        /// </summary>
        public readonly string SchemeType;
        /// <summary>
        /// VALUE.
        /// </summary>
        public readonly string? Value;
        /// <summary>
        /// VALUE process.
        /// </summary>
        public readonly Outputs.DatahubTaskTransformsParamFieldChainSMTValueOperate? ValueOperate;
        /// <summary>
        /// VALUE process chain.
        /// </summary>
        public readonly ImmutableArray<Outputs.DatahubTaskTransformsParamFieldChainSMTValueOperate> ValueOperates;

        [OutputConstructor]
        private DatahubTaskTransformsParamFieldChainSMT(
            string key,

            string operate,

            string? originalValue,

            string schemeType,

            string? value,

            Outputs.DatahubTaskTransformsParamFieldChainSMTValueOperate? valueOperate,

            ImmutableArray<Outputs.DatahubTaskTransformsParamFieldChainSMTValueOperate> valueOperates)
        {
            Key = key;
            Operate = operate;
            OriginalValue = originalValue;
            SchemeType = schemeType;
            Value = value;
            ValueOperate = valueOperate;
            ValueOperates = valueOperates;
        }
    }
}
