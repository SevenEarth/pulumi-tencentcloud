// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketWebsiteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// `Endpoint` of the static website.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// An absolute path to the document to return in case of a 4XX error.
        /// </summary>
        [Input("errorDocument")]
        public Input<string>? ErrorDocument { get; set; }

        /// <summary>
        /// COS returns this index document when requests are made to the root domain or any of the subfolders.
        /// </summary>
        [Input("indexDocument")]
        public Input<string>? IndexDocument { get; set; }

        public BucketWebsiteArgs()
        {
        }
        public static new BucketWebsiteArgs Empty => new BucketWebsiteArgs();
    }
}
