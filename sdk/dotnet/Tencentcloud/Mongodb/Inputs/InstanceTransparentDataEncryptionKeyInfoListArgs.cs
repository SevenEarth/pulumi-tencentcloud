// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mongodb.Inputs
{

    public sealed class InstanceTransparentDataEncryptionKeyInfoListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Instance and key binding time.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Master Key ID.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// Master key name.
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Key origin.
        /// </summary>
        [Input("keyOrigin")]
        public Input<string>? KeyOrigin { get; set; }

        /// <summary>
        /// Purpose of the key.
        /// </summary>
        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        /// <summary>
        /// Key status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public InstanceTransparentDataEncryptionKeyInfoListArgs()
        {
        }
        public static new InstanceTransparentDataEncryptionKeyInfoListArgs Empty => new InstanceTransparentDataEncryptionKeyInfoListArgs();
    }
}
