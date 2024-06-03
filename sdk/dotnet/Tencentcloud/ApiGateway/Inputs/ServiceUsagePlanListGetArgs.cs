// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Inputs
{

    public sealed class ServiceUsagePlanListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the API.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// Binding type.
        /// </summary>
        [Input("bindType")]
        public Input<string>? BindType { get; set; }

        /// <summary>
        /// ID of the usage plan.
        /// </summary>
        [Input("usagePlanId")]
        public Input<string>? UsagePlanId { get; set; }

        /// <summary>
        /// Name of the usage plan.
        /// </summary>
        [Input("usagePlanName")]
        public Input<string>? UsagePlanName { get; set; }

        public ServiceUsagePlanListGetArgs()
        {
        }
        public static new ServiceUsagePlanListGetArgs Empty => new ServiceUsagePlanListGetArgs();
    }
}
