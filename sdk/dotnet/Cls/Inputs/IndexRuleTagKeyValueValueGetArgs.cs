// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cls.Inputs
{

    public sealed class IndexRuleTagKeyValueValueGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether Chinese characters are contained.
        /// </summary>
        [Input("containZH")]
        public Input<bool>? ContainZH { get; set; }

        /// <summary>
        /// Whether the analysis feature is enabled for the field.
        /// </summary>
        [Input("sqlFlag")]
        public Input<bool>? SqlFlag { get; set; }

        /// <summary>
        /// Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.
        /// </summary>
        [Input("tokenizer")]
        public Input<string>? Tokenizer { get; set; }

        /// <summary>
        /// Field type. Valid values: long, text, double.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IndexRuleTagKeyValueValueGetArgs()
        {
        }
    }
}
