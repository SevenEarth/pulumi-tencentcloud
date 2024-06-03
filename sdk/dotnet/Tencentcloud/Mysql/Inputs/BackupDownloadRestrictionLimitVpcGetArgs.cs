// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Inputs
{

    public sealed class BackupDownloadRestrictionLimitVpcGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Restrict downloads from regions. Currently only the current region is supported.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("vpcLists", required: true)]
        private InputList<string>? _vpcLists;

        /// <summary>
        /// List of vpcs to limit downloads.
        /// </summary>
        public InputList<string> VpcLists
        {
            get => _vpcLists ?? (_vpcLists = new InputList<string>());
            set => _vpcLists = value;
        }

        public BackupDownloadRestrictionLimitVpcGetArgs()
        {
        }
        public static new BackupDownloadRestrictionLimitVpcGetArgs Empty => new BackupDownloadRestrictionLimitVpcGetArgs();
    }
}
