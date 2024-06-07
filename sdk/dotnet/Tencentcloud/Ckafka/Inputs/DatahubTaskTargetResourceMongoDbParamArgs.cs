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

    public sealed class DatahubTaskTargetResourceMongoDbParamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// MongoDB collection.
        /// </summary>
        [Input("collection", required: true)]
        public Input<string> Collection { get; set; } = null!;

        /// <summary>
        /// Whether to copy the stock data, the default parameter is true.
        /// </summary>
        [Input("copyExisting", required: true)]
        public Input<bool> CopyExisting { get; set; } = null!;

        /// <summary>
        /// MongoDB database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// Mongo DB connection ip.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Listening event type, if it is empty, it means select all. Values include insert, update, replace, delete, invalidate, drop, dropdatabase, rename, used between multiple types, separated by commas.
        /// </summary>
        [Input("listeningEvent")]
        public Input<string>? ListeningEvent { get; set; }

        /// <summary>
        /// MongoDB database password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// aggregation pipeline.
        /// </summary>
        [Input("pipeline")]
        public Input<string>? Pipeline { get; set; }

        /// <summary>
        /// MongoDB connection port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Master-slave priority, default master node.
        /// </summary>
        [Input("readPreference")]
        public Input<string>? ReadPreference { get; set; }

        /// <summary>
        /// resource id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Whether it is a self-built cluster.
        /// </summary>
        [Input("selfBuilt")]
        public Input<bool>? SelfBuilt { get; set; }

        /// <summary>
        /// MongoDB database user name.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public DatahubTaskTargetResourceMongoDbParamArgs()
        {
        }
        public static new DatahubTaskTargetResourceMongoDbParamArgs Empty => new DatahubTaskTargetResourceMongoDbParamArgs();
    }
}
