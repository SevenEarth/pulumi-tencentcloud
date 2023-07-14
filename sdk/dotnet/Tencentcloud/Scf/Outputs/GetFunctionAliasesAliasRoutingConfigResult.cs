// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf.Outputs
{

    [OutputType]
    public sealed class GetFunctionAliasesAliasRoutingConfigResult
    {
        /// <summary>
        /// Additional version with rule-based routing.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAliasesAliasRoutingConfigAdditionVersionMatchResult> AdditionVersionMatchs;
        /// <summary>
        /// Additional version with random weight-based routing.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionAliasesAliasRoutingConfigAdditionalVersionWeightResult> AdditionalVersionWeights;

        [OutputConstructor]
        private GetFunctionAliasesAliasRoutingConfigResult(
            ImmutableArray<Outputs.GetFunctionAliasesAliasRoutingConfigAdditionVersionMatchResult> additionVersionMatchs,

            ImmutableArray<Outputs.GetFunctionAliasesAliasRoutingConfigAdditionalVersionWeightResult> additionalVersionWeights)
        {
            AdditionVersionMatchs = additionVersionMatchs;
            AdditionalVersionWeights = additionalVersionWeights;
        }
    }
}
