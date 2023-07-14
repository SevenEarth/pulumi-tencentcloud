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
    public sealed class DatahubTaskTransformParamMapParam
    {
        /// <summary>
        /// key.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Type, DEFAULT default, DATE system default - timestamp, CUSTOMIZE custom, MAPPING mapping.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private DatahubTaskTransformParamMapParam(
            string key,

            string? type,

            string? value)
        {
            Key = key;
            Type = type;
            Value = value;
        }
    }
}
