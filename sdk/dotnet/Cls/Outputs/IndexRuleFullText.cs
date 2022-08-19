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
    public sealed class IndexRuleFullText
    {
        /// <summary>
        /// Case sensitivity.
        /// </summary>
        public readonly bool CaseSensitive;
        /// <summary>
        /// Whether Chinese characters are contained.
        /// </summary>
        public readonly bool ContainZH;
        /// <summary>
        /// Full-Text index delimiter. Each character in the string represents a delimiter.
        /// </summary>
        public readonly string Tokenizer;

        [OutputConstructor]
        private IndexRuleFullText(
            bool caseSensitive,

            bool containZH,

            string tokenizer)
        {
            CaseSensitive = caseSensitive;
            ContainZH = containZH;
            Tokenizer = tokenizer;
        }
    }
}
