// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dbbrain.Outputs
{

    [OutputType]
    public sealed class GetMysqlProcessListProcessListResult
    {
        /// <summary>
        /// The execution type of the thread, used to filter the thread list.
        /// </summary>
        public readonly string Command;
        /// <summary>
        /// The threads operations database, used to filter the thread list.
        /// </summary>
        public readonly string Db;
        /// <summary>
        /// The operating host address of the thread, used to filter the thread list.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// thread ID, used to filter the thread list.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The threads operation statement is used to filter the thread list.
        /// </summary>
        public readonly string Info;
        /// <summary>
        /// The operational state of the thread, used to filter the thread list.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The minimum value of the operation duration of a thread, in seconds, used to filter the list of threads whose operation duration is longer than this value.
        /// </summary>
        public readonly string Time;
        /// <summary>
        /// The operating account name of the thread, used to filter the thread list.
        /// </summary>
        public readonly string User;

        [OutputConstructor]
        private GetMysqlProcessListProcessListResult(
            string command,

            string db,

            string host,

            string id,

            string info,

            string state,

            string time,

            string user)
        {
            Command = command;
            Db = db;
            Host = host;
            Id = id;
            Info = info;
            State = state;
            Time = time;
            User = user;
        }
    }
}
