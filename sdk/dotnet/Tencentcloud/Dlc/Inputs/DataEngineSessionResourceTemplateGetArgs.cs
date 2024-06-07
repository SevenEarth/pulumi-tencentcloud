// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc.Inputs
{

    public sealed class DataEngineSessionResourceTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Engine driver size specification only supports: small/medium/large/xlarge/m.small/m.medium/m.large/m.xlarge.
        /// </summary>
        [Input("driverSize")]
        public Input<string>? DriverSize { get; set; }

        /// <summary>
        /// Specify the executor max number (in a dynamic configuration scenario), the minimum value is 1, and the maximum value is less than the cluster specification (when ExecutorMaxNumbers is less than ExecutorNums, the value is set to ExecutorNums).
        /// </summary>
        [Input("executorMaxNumbers")]
        public Input<int>? ExecutorMaxNumbers { get; set; }

        /// <summary>
        /// Specify the number of executors. The minimum value is 1 and the maximum value is less than the cluster specification.
        /// </summary>
        [Input("executorNums")]
        public Input<int>? ExecutorNums { get; set; }

        /// <summary>
        /// Engine executor size specification only supports: small/medium/large/xlarge/m.small/m.medium/m.large/m.xlarge.
        /// </summary>
        [Input("executorSize")]
        public Input<string>? ExecutorSize { get; set; }

        public DataEngineSessionResourceTemplateGetArgs()
        {
        }
        public static new DataEngineSessionResourceTemplateGetArgs Empty => new DataEngineSessionResourceTemplateGetArgs();
    }
}
