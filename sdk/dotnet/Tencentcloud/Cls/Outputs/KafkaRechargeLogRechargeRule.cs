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
    public sealed class KafkaRechargeLogRechargeRule
    {
        /// <summary>
        /// default time from.
        /// </summary>
        public readonly int? DefaultTimeSrc;
        /// <summary>
        /// user default time.
        /// </summary>
        public readonly bool DefaultTimeSwitch;
        /// <summary>
        /// encoding format.
        /// </summary>
        public readonly int EncodingFormat;
        /// <summary>
        /// log key list.
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        /// <summary>
        /// log regex.
        /// </summary>
        public readonly string? LogRegex;
        /// <summary>
        /// metadata.
        /// </summary>
        public readonly ImmutableArray<string> Metadatas;
        /// <summary>
        /// recharge type.
        /// </summary>
        public readonly string RechargeType;
        /// <summary>
        /// time format.
        /// </summary>
        public readonly string? TimeFormat;
        /// <summary>
        /// time key.
        /// </summary>
        public readonly string? TimeKey;
        /// <summary>
        /// time regex.
        /// </summary>
        public readonly string? TimeRegex;
        /// <summary>
        /// time zone.
        /// </summary>
        public readonly string? TimeZone;
        /// <summary>
        /// parse failed log key.
        /// </summary>
        public readonly string? UnMatchLogKey;
        /// <summary>
        /// is push parse failed log.
        /// </summary>
        public readonly bool? UnMatchLogSwitch;
        /// <summary>
        /// parse failed log time from.
        /// </summary>
        public readonly int? UnMatchLogTimeSrc;

        [OutputConstructor]
        private KafkaRechargeLogRechargeRule(
            int? defaultTimeSrc,

            bool defaultTimeSwitch,

            int encodingFormat,

            ImmutableArray<string> keys,

            string? logRegex,

            ImmutableArray<string> metadatas,

            string rechargeType,

            string? timeFormat,

            string? timeKey,

            string? timeRegex,

            string? timeZone,

            string? unMatchLogKey,

            bool? unMatchLogSwitch,

            int? unMatchLogTimeSrc)
        {
            DefaultTimeSrc = defaultTimeSrc;
            DefaultTimeSwitch = defaultTimeSwitch;
            EncodingFormat = encodingFormat;
            Keys = keys;
            LogRegex = logRegex;
            Metadatas = metadatas;
            RechargeType = rechargeType;
            TimeFormat = timeFormat;
            TimeKey = timeKey;
            TimeRegex = timeRegex;
            TimeZone = timeZone;
            UnMatchLogKey = unMatchLogKey;
            UnMatchLogSwitch = unMatchLogSwitch;
            UnMatchLogTimeSrc = unMatchLogTimeSrc;
        }
    }
}
