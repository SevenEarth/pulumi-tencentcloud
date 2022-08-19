// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls.Outputs
{

    [OutputType]
    public sealed class IndexRuleKeyValueKeyValueValue
    {
        /// <summary>
        /// Whether Chinese characters are contained.
        /// </summary>
        public readonly bool? ContainZH;
        /// <summary>
        /// Whether the analysis feature is enabled for the field.
        /// </summary>
        public readonly bool? SqlFlag;
        /// <summary>
        /// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
        /// </summary>
        public readonly string? Tokenizer;
        /// <summary>
        /// Field type. Valid values: long, text, double.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private IndexRuleKeyValueKeyValueValue(
            bool? containZH,

            bool? sqlFlag,

            string? tokenizer,

            string type)
        {
            ContainZH = containZH;
            SqlFlag = sqlFlag;
            Tokenizer = tokenizer;
            Type = type;
        }
    }
}
