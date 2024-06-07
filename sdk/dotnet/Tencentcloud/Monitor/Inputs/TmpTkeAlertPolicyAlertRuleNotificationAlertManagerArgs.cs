// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class TmpTkeAlertPolicyAlertRuleNotificationAlertManagerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the cluster where the alertmanager is deployed. Note: This field may return null, indicating that a valid value could not be retrieved.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Alertmanager is deployed in the cluster type. Note: This field may return null, indicating that a valid value could not be retrieved.
        /// </summary>
        [Input("clusterType")]
        public Input<string>? ClusterType { get; set; }

        /// <summary>
        /// Alertmanager url.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public TmpTkeAlertPolicyAlertRuleNotificationAlertManagerArgs()
        {
        }
        public static new TmpTkeAlertPolicyAlertRuleNotificationAlertManagerArgs Empty => new TmpTkeAlertPolicyAlertRuleNotificationAlertManagerArgs();
    }
}
