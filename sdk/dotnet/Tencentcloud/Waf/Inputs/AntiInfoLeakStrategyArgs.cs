// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf.Inputs
{

    public sealed class AntiInfoLeakStrategyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Matching Content. If field is returncode support: 400, 403, 404, 4xx, 500, 501, 502, 504, 5xx; If field is information support: idcard, phone, bankcard; If field is keywords users input matching content themselves.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Matching Fields. support: returncode, keywords, information.
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        public AntiInfoLeakStrategyArgs()
        {
        }
    }
}
