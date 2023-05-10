// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem.Inputs
{

    public sealed class WorkloadDeployStrategyConfArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// interval between batches.
        /// </summary>
        [Input("batchInterval")]
        public Input<int>? BatchInterval { get; set; }

        /// <summary>
        /// beta batch number.
        /// </summary>
        [Input("betaBatchNum")]
        public Input<int>? BetaBatchNum { get; set; }

        /// <summary>
        /// strategy type, 0 means auto, 1 means manual, 2 means manual with beta batch.
        /// </summary>
        [Input("deployStrategyType", required: true)]
        public Input<int> DeployStrategyType { get; set; } = null!;

        /// <summary>
        /// force update.
        /// </summary>
        [Input("force")]
        public Input<bool>? Force { get; set; }

        /// <summary>
        /// minimal available instances duration deployment.
        /// </summary>
        [Input("minAvailable")]
        public Input<int>? MinAvailable { get; set; }

        /// <summary>
        /// total batch number.
        /// </summary>
        [Input("totalBatchCount", required: true)]
        public Input<int> TotalBatchCount { get; set; } = null!;

        public WorkloadDeployStrategyConfArgs()
        {
        }
    }
}
