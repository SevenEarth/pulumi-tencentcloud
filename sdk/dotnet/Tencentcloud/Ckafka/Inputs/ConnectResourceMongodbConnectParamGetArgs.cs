// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class ConnectResourceMongodbConnectParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to update to the associated Datahub task, default: false.
        /// </summary>
        [Input("isUpdate")]
        public Input<bool>? IsUpdate { get; set; }

        /// <summary>
        /// Password for the source of the Mongo DB connection.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// MongoDB port.
        /// </summary>
        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        /// <summary>
        /// Instance resource of Mongo DB connection source.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Whether the Mongo DB connection source is a self-built cluster.
        /// </summary>
        [Input("selfBuilt", required: true)]
        public Input<bool> SelfBuilt { get; set; } = null!;

        /// <summary>
        /// The instance VIP of the Mongo DB connection source, when it is a Tencent Cloud instance, it is required.
        /// </summary>
        [Input("serviceVip")]
        public Input<string>? ServiceVip { get; set; }

        /// <summary>
        /// The vpc Id of the Mongo DB connection source, which is required when it is a Tencent Cloud instance.
        /// </summary>
        [Input("uniqVpcId")]
        public Input<string>? UniqVpcId { get; set; }

        /// <summary>
        /// The username of the Mongo DB connection source.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public ConnectResourceMongodbConnectParamGetArgs()
        {
        }
    }
}
