// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainErrorPagePageRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Redirect code of error page rules.
        /// </summary>
        [Input("redirectCode", required: true)]
        public Input<int> RedirectCode { get; set; } = null!;

        /// <summary>
        /// Redirect url of error page rules.
        /// </summary>
        [Input("redirectUrl", required: true)]
        public Input<string> RedirectUrl { get; set; } = null!;

        /// <summary>
        /// Status code of error page rules.
        /// </summary>
        [Input("statusCode", required: true)]
        public Input<int> StatusCode { get; set; } = null!;

        public DomainErrorPagePageRuleArgs()
        {
        }
        public static new DomainErrorPagePageRuleArgs Empty => new DomainErrorPagePageRuleArgs();
    }
}
