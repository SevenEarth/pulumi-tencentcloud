// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Inputs
{

    public sealed class DomainAuthenticationTypeAGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used for calculate a signature. 6-32 characters. Only digits and letters are allowed.
        /// </summary>
        [Input("backupSecretKey")]
        public Input<string>? BackupSecretKey { get; set; }

        /// <summary>
        /// Signature expiration time in second. The maximum value is 630720000.
        /// </summary>
        [Input("expireTime", required: true)]
        public Input<int> ExpireTime { get; set; } = null!;

        [Input("fileExtensions", required: true)]
        private InputList<string>? _fileExtensions;

        /// <summary>
        /// File extension list settings determining if authentication should be performed. NOTE: If it contains an asterisk (*), this indicates all files.
        /// </summary>
        public InputList<string> FileExtensions
        {
            get => _fileExtensions ?? (_fileExtensions = new InputList<string>());
            set => _fileExtensions = value;
        }

        /// <summary>
        /// Available values: `whitelist` - all types apart from `file_extensions` are authenticated, `blacklist`: - only the types in the `file_extensions` are authenticated.
        /// </summary>
        [Input("filterType", required: true)]
        public Input<string> FilterType { get; set; } = null!;

        /// <summary>
        /// The key for signature calculation. Only digits, upper and lower-case letters are allowed. Length limit: 6-32 characters.
        /// </summary>
        [Input("secretKey", required: true)]
        public Input<string> SecretKey { get; set; } = null!;

        /// <summary>
        /// Signature parameter name. Only upper and lower-case letters, digits, and underscores (_) are allowed. It cannot start with a digit. Length limit: 1-100 characters.
        /// </summary>
        [Input("signParam", required: true)]
        public Input<string> SignParam { get; set; } = null!;

        public DomainAuthenticationTypeAGetArgs()
        {
        }
    }
}
