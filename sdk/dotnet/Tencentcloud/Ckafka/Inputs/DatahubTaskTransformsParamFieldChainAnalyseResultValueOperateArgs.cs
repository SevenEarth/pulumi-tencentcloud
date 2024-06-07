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

    public sealed class DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time conversion, required when TYPE=DATE.
        /// </summary>
        [Input("date")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateDateArgs>? Date { get; set; }

        /// <summary>
        /// Json Path replacement, must pass when TYPE=JSON PATH REPLACE.
        /// </summary>
        [Input("jsonPathReplace")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateJsonPathReplaceArgs>? JsonPathReplace { get; set; }

        /// <summary>
        /// Key-value secondary analysis, must be passed when TYPE=KV.
        /// </summary>
        [Input("kV")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateKVArgs>? KV { get; set; }

        /// <summary>
        /// Regular replacement, required when TYPE=REGEX REPLACE.
        /// </summary>
        [Input("regexReplace")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateRegexReplaceArgs>? RegexReplace { get; set; }

        /// <summary>
        /// replace, TYPE=REPLACE is required.
        /// </summary>
        [Input("replace")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateReplaceArgs>? Replace { get; set; }

        /// <summary>
        /// result.
        /// </summary>
        [Input("result")]
        public Input<string>? Result { get; set; }

        /// <summary>
        /// The value supports one split and multiple values, required when TYPE=SPLIT.
        /// </summary>
        [Input("split")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSplitArgs>? Split { get; set; }

        /// <summary>
        /// Substr, TYPE=SUBSTR is required.
        /// </summary>
        [Input("substr")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateSubstrArgs>? Substr { get; set; }

        /// <summary>
        /// Processing mode, REPLACE replacement, SUBSTR interception, DATE date conversion, TRIM removal of leading and trailing spaces, REGEX REPLACE regular replacement, URL DECODE, LOWERCASE conversion to lowercase.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Url parsing.
        /// </summary>
        [Input("urlDecode")]
        public Input<Inputs.DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateUrlDecodeArgs>? UrlDecode { get; set; }

        public DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateArgs()
        {
        }
        public static new DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateArgs Empty => new DatahubTaskTransformsParamFieldChainAnalyseResultValueOperateArgs();
    }
}
