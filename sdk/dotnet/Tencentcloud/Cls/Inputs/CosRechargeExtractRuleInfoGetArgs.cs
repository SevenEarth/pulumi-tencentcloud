// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Inputs
{

    public sealed class CosRechargeExtractRuleInfoGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// syslog address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// backtracking data volume in incremental acquisition mode.
        /// </summary>
        [Input("backtracking")]
        public Input<int>? Backtracking { get; set; }

        /// <summary>
        /// begin line regex.
        /// </summary>
        [Input("beginRegex")]
        public Input<string>? BeginRegex { get; set; }

        /// <summary>
        /// log delimiter.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("filterKeyRegexes")]
        private InputList<Inputs.CosRechargeExtractRuleInfoFilterKeyRegexGetArgs>? _filterKeyRegexes;

        /// <summary>
        /// rules that need to filter logs.
        /// </summary>
        public InputList<Inputs.CosRechargeExtractRuleInfoFilterKeyRegexGetArgs> FilterKeyRegexes
        {
            get => _filterKeyRegexes ?? (_filterKeyRegexes = new InputList<Inputs.CosRechargeExtractRuleInfoFilterKeyRegexGetArgs>());
            set => _filterKeyRegexes = value;
        }

        /// <summary>
        /// gbk encoding.
        /// </summary>
        [Input("isGbk")]
        public Input<int>? IsGbk { get; set; }

        /// <summary>
        /// is standard json.
        /// </summary>
        [Input("jsonStandard")]
        public Input<int>? JsonStandard { get; set; }

        [Input("keys")]
        private InputList<string>? _keys;

        /// <summary>
        /// key list.
        /// </summary>
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        /// <summary>
        /// log regex.
        /// </summary>
        [Input("logRegex")]
        public Input<string>? LogRegex { get; set; }

        [Input("metaTags")]
        private InputList<Inputs.CosRechargeExtractRuleInfoMetaTagGetArgs>? _metaTags;

        /// <summary>
        /// metadata tag list.
        /// </summary>
        public InputList<Inputs.CosRechargeExtractRuleInfoMetaTagGetArgs> MetaTags
        {
            get => _metaTags ?? (_metaTags = new InputList<Inputs.CosRechargeExtractRuleInfoMetaTagGetArgs>());
            set => _metaTags = value;
        }

        /// <summary>
        /// metadata type.
        /// </summary>
        [Input("metadataType")]
        public Input<int>? MetadataType { get; set; }

        /// <summary>
        /// parse protocol.
        /// </summary>
        [Input("parseProtocol")]
        public Input<string>? ParseProtocol { get; set; }

        /// <summary>
        /// metadata path regex.
        /// </summary>
        [Input("pathRegex")]
        public Input<string>? PathRegex { get; set; }

        /// <summary>
        /// syslog protocol.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// time format.
        /// </summary>
        [Input("timeFormat")]
        public Input<string>? TimeFormat { get; set; }

        /// <summary>
        /// time key.
        /// </summary>
        [Input("timeKey")]
        public Input<string>? TimeKey { get; set; }

        /// <summary>
        /// parsing failure log key.
        /// </summary>
        [Input("unMatchLogKey")]
        public Input<string>? UnMatchLogKey { get; set; }

        /// <summary>
        /// whether to upload the parsing failure log.
        /// </summary>
        [Input("unMatchUpLoadSwitch")]
        public Input<bool>? UnMatchUpLoadSwitch { get; set; }

        public CosRechargeExtractRuleInfoGetArgs()
        {
        }
    }
}
