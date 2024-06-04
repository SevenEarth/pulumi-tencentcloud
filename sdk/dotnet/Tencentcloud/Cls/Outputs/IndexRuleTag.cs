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
    public sealed class IndexRuleTag
    {
        /// <summary>
        /// Case sensitivity.
        /// </summary>
        public readonly bool CaseSensitive;
        /// <summary>
        /// Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexRuleTagKeyValue> KeyValues;

        [OutputConstructor]
        private IndexRuleTag(
            bool caseSensitive,

            ImmutableArray<Outputs.IndexRuleTagKeyValue> keyValues)
        {
            CaseSensitive = caseSensitive;
            KeyValues = keyValues;
        }
    }
}
