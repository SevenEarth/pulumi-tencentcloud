// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class WorkflowAiAnalysisTaskGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Video Content Analysis Template ID.
        /// </summary>
        [Input("definition", required: true)]
        public Input<int> Definition { get; set; } = null!;

        /// <summary>
        /// Extension parameter whose value is a serialized json string.Note: This parameter is a customized demand parameter, which requires offline docking.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("extendedParameter")]
        public Input<string>? ExtendedParameter { get; set; }

        public WorkflowAiAnalysisTaskGetArgs()
        {
        }
        public static new WorkflowAiAnalysisTaskGetArgs Empty => new WorkflowAiAnalysisTaskGetArgs();
    }
}
