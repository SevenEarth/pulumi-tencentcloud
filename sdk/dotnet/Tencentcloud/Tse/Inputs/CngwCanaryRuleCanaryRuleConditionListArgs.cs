// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class CngwCanaryRuleCanaryRuleConditionListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// delimiter. valid when operator is in or not in, reference value:`,`, `;`,`\n`.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        /// <summary>
        /// global configuration ID.
        /// </summary>
        [Input("globalConfigId")]
        public Input<string>? GlobalConfigId { get; set; }

        /// <summary>
        /// global configuration name.
        /// </summary>
        [Input("globalConfigName")]
        public Input<string>? GlobalConfigName { get; set; }

        /// <summary>
        /// parameter name.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// operator.Reference value:`le`,`eq`,`lt`,`ne`,`ge`,`gt`,`regex`,`exists`,`in`,`not in`,`prefix`,`exact`,`regex`.
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        /// <summary>
        /// type.Reference value:`path`,`method`,`query`,`header`,`cookie`,`body`,`system`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// parameter value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public CngwCanaryRuleCanaryRuleConditionListArgs()
        {
        }
    }
}
