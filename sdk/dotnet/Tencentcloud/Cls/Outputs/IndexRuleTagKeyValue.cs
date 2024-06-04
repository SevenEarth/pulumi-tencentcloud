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
    public sealed class IndexRuleTagKeyValue
    {
        /// <summary>
        /// When a key value or metafield index needs to be configured for a field, the metafield Key does not need to be prefixed with __TAG__. and is consistent with the one when logs are uploaded. __TAG__. will be prefixed automatically for display in the console..
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Field index description information.
        /// </summary>
        public readonly Outputs.IndexRuleTagKeyValueValue? Value;

        [OutputConstructor]
        private IndexRuleTagKeyValue(
            string key,

            Outputs.IndexRuleTagKeyValueValue? value)
        {
            Key = key;
            Value = value;
        }
    }
}
