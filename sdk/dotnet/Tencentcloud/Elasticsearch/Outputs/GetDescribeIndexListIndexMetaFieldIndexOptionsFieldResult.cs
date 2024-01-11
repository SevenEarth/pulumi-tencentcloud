// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class GetDescribeIndexListIndexMetaFieldIndexOptionsFieldResult
    {
        /// <summary>
        /// Expire max age.
        /// </summary>
        public readonly string ExpireMaxAge;
        /// <summary>
        /// Expire max size.
        /// </summary>
        public readonly string ExpireMaxSize;
        /// <summary>
        /// Whether to turn on dynamic scrolling.
        /// </summary>
        public readonly string RolloverDynamic;
        /// <summary>
        /// Rollover max age.
        /// </summary>
        public readonly string RolloverMaxAge;
        /// <summary>
        /// Whether to enable dynamic slicing.
        /// </summary>
        public readonly string ShardNumDynamic;
        /// <summary>
        /// Time partition field.
        /// </summary>
        public readonly string TimestampField;
        /// <summary>
        /// Write mode.
        /// </summary>
        public readonly string WriteMode;

        [OutputConstructor]
        private GetDescribeIndexListIndexMetaFieldIndexOptionsFieldResult(
            string expireMaxAge,

            string expireMaxSize,

            string rolloverDynamic,

            string rolloverMaxAge,

            string shardNumDynamic,

            string timestampField,

            string writeMode)
        {
            ExpireMaxAge = expireMaxAge;
            ExpireMaxSize = expireMaxSize;
            RolloverDynamic = rolloverDynamic;
            RolloverMaxAge = rolloverMaxAge;
            ShardNumDynamic = shardNumDynamic;
            TimestampField = timestampField;
            WriteMode = writeMode;
        }
    }
}