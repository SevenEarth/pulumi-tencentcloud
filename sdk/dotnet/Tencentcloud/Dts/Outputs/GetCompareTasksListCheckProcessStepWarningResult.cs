// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class GetCompareTasksListCheckProcessStepWarningResult
    {
        public readonly string? HelpDoc;
        /// <summary>
        /// message.
        /// </summary>
        public readonly string? Message;
        public readonly string? Solution;

        [OutputConstructor]
        private GetCompareTasksListCheckProcessStepWarningResult(
            string? helpDoc,

            string? message,

            string? solution)
        {
            HelpDoc = helpDoc;
            Message = message;
            Solution = solution;
        }
    }
}
