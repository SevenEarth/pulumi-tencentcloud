// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Outputs
{

    [OutputType]
    public sealed class CngwCanaryRuleCanaryRuleConditionList
    {
        /// <summary>
        /// delimiter. valid when operator is in or not in, reference value:`,`, `;`,`\n`.
        /// </summary>
        public readonly string? Delimiter;
        /// <summary>
        /// global configuration ID.
        /// </summary>
        public readonly string? GlobalConfigId;
        /// <summary>
        /// global configuration name.
        /// </summary>
        public readonly string? GlobalConfigName;
        /// <summary>
        /// parameter name.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// operator.Reference value:`le`,`eq`,`lt`,`ne`,`ge`,`gt`,`regex`,`exists`,`in`,`not in`,`prefix`,`exact`,`regex`.
        /// </summary>
        public readonly string? Operator;
        /// <summary>
        /// type.Reference value:`path`,`method`,`query`,`header`,`cookie`,`body`,`system`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// parameter value.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private CngwCanaryRuleCanaryRuleConditionList(
            string? delimiter,

            string? globalConfigId,

            string? globalConfigName,

            string? key,

            string? @operator,

            string type,

            string? value)
        {
            Delimiter = delimiter;
            GlobalConfigId = globalConfigId;
            GlobalConfigName = globalConfigName;
            Key = key;
            Operator = @operator;
            Type = type;
            Value = value;
        }
    }
}
