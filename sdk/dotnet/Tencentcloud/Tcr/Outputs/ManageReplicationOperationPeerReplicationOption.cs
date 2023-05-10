// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr.Outputs
{

    [OutputType]
    public sealed class ManageReplicationOperationPeerReplicationOption
    {
        /// <summary>
        /// whether to enable cross-master account instance synchronization.
        /// </summary>
        public readonly bool EnablePeerReplication;
        /// <summary>
        /// access permanent token of the instance to be synchronized.
        /// </summary>
        public readonly string PeerRegistryToken;
        /// <summary>
        /// uin of the instance to be synchronized.
        /// </summary>
        public readonly string PeerRegistryUin;

        [OutputConstructor]
        private ManageReplicationOperationPeerReplicationOption(
            bool enablePeerReplication,

            string peerRegistryToken,

            string peerRegistryUin)
        {
            EnablePeerReplication = enablePeerReplication;
            PeerRegistryToken = peerRegistryToken;
            PeerRegistryUin = peerRegistryUin;
        }
    }
}
