// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Outputs
{

    [OutputType]
    public sealed class GetAlarmPolicyPolicyConditionRuleFilterResult
    {
        /// <summary>
        /// The alarm object list, which is a JSON string. The outer array corresponds to multiple instances, and the inner array is the dimension of an object.For example, 'CVM - Basic Monitor' can be written as: [ {Dimensions: {unInstanceId: ins-qr8d555g}}, {Dimensions: {unInstanceId: ins-qr8d555h}} ]You can also refer to the 'Example 2' below.For more information on the parameter samples of different Tencent Cloud services, see [Product Policy Type and Dimension Information](https://www.tencentcloud.com/document/product/248/39565?has_map=1).Note: If 1 is passed in for NeedCorrespondence, the relationship between a policy and an instance needs to be returned. You can pass in up to 20 alarm object dimensions to avoid request timeout.
        /// </summary>
        public readonly string Dimensions;
        /// <summary>
        /// Triggered task type. Valid value: AS (auto scaling)Note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAlarmPolicyPolicyConditionRuleFilterResult(
            string dimensions,

            string type)
        {
            Dimensions = dimensions;
            Type = type;
        }
    }
}
