// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Outputs
{

    [OutputType]
    public sealed class ConfigExtractRule
    {
        /// <summary>
        /// syslog system log collection specifies the address and port that the collector listens to.
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).
        /// </summary>
        public readonly int? Backtracking;
        /// <summary>
        /// First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.
        /// </summary>
        public readonly string? BeginRegex;
        /// <summary>
        /// Delimiter for delimited log, which is valid only if log_type is delimiter_log.
        /// </summary>
        public readonly string? Delimiter;
        /// <summary>
        /// Log keys to be filtered and the corresponding regex.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigExtractRuleFilterKeyRegex> FilterKeyRegexes;
        /// <summary>
        /// GBK encoding. Default 0.
        /// </summary>
        public readonly int? IsGbk;
        /// <summary>
        /// standard json. Default 0.
        /// </summary>
        public readonly int? JsonStandard;
        /// <summary>
        /// Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        /// <summary>
        /// Full log matching rule, which is valid only if log_type is fullregex_log.
        /// </summary>
        public readonly string? LogRegex;
        /// <summary>
        /// metadata tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.ConfigExtractRuleMetaTag> MetaTags;
        /// <summary>
        /// metadata type.
        /// </summary>
        public readonly int? MetadataType;
        /// <summary>
        /// parse protocol.
        /// </summary>
        public readonly string? ParseProtocol;
        /// <summary>
        /// metadata path regex.
        /// </summary>
        public readonly string? PathRegex;
        /// <summary>
        /// syslog protocol, tcp or udp.
        /// </summary>
        public readonly string? Protocol;
        /// <summary>
        /// Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.
        /// </summary>
        public readonly string? TimeFormat;
        /// <summary>
        /// Time field key name. time_key and time_format must appear in pair.
        /// </summary>
        public readonly string? TimeKey;
        /// <summary>
        /// Unmatched log key.
        /// </summary>
        public readonly string? UnMatchLogKey;
        /// <summary>
        /// Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no.
        /// </summary>
        public readonly bool? UnMatchUpLoadSwitch;

        [OutputConstructor]
        private ConfigExtractRule(
            string? address,

            int? backtracking,

            string? beginRegex,

            string? delimiter,

            ImmutableArray<Outputs.ConfigExtractRuleFilterKeyRegex> filterKeyRegexes,

            int? isGbk,

            int? jsonStandard,

            ImmutableArray<string> keys,

            string? logRegex,

            ImmutableArray<Outputs.ConfigExtractRuleMetaTag> metaTags,

            int? metadataType,

            string? parseProtocol,

            string? pathRegex,

            string? protocol,

            string? timeFormat,

            string? timeKey,

            string? unMatchLogKey,

            bool? unMatchUpLoadSwitch)
        {
            Address = address;
            Backtracking = backtracking;
            BeginRegex = beginRegex;
            Delimiter = delimiter;
            FilterKeyRegexes = filterKeyRegexes;
            IsGbk = isGbk;
            JsonStandard = jsonStandard;
            Keys = keys;
            LogRegex = logRegex;
            MetaTags = metaTags;
            MetadataType = metadataType;
            ParseProtocol = parseProtocol;
            PathRegex = pathRegex;
            Protocol = protocol;
            TimeFormat = timeFormat;
            TimeKey = timeKey;
            UnMatchLogKey = unMatchLogKey;
            UnMatchUpLoadSwitch = unMatchUpLoadSwitch;
        }
    }
}
