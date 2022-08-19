// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Kms.Outputs
{

    [OutputType]
    public sealed class GetKeysKeyListResult
    {
        /// <summary>
        /// Name of CMK.
        /// </summary>
        public readonly string Alias;
        /// <summary>
        /// Create time of CMK.
        /// </summary>
        public readonly int CreateTime;
        /// <summary>
        /// Uin of CMK Creator.
        /// </summary>
        public readonly int CreatorUin;
        /// <summary>
        /// Delete time of CMK.
        /// </summary>
        public readonly int DeletionDate;
        /// <summary>
        /// Description of CMK.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// ID of CMK.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// Specify whether to enable key rotation.
        /// </summary>
        public readonly bool KeyRotationEnabled;
        /// <summary>
        /// Filter by state of CMK. `0` - all CMKs are queried, `1` - only Enabled CMKs are queried, `2` - only Disabled CMKs are queried, `3` - only PendingDelete CMKs are queried, `4` - only PendingImport CMKs are queried, `5` - only Archived CMKs are queried.
        /// </summary>
        public readonly string KeyState;
        /// <summary>
        /// Filter by usage of CMK. Available values include `ALL`, `ENCRYPT_DECRYPT`, `ASYMMETRIC_DECRYPT_RSA_2048`, `ASYMMETRIC_DECRYPT_SM2`, `ASYMMETRIC_SIGN_VERIFY_SM2`, `ASYMMETRIC_SIGN_VERIFY_RSA_2048`, `ASYMMETRIC_SIGN_VERIFY_ECC`. Default value is `ENCRYPT_DECRYPT`.
        /// </summary>
        public readonly string KeyUsage;
        /// <summary>
        /// Next rotate time of CMK when key_rotation_enabled is true.
        /// </summary>
        public readonly int NextRotateTime;
        /// <summary>
        /// Filter by origin of CMK. `TENCENT_KMS` - CMK created by KMS, `EXTERNAL` - CMK imported by user, `ALL` - all CMKs. Default value is `ALL`.
        /// </summary>
        public readonly string Origin;
        /// <summary>
        /// Creator of CMK.
        /// </summary>
        public readonly string Owner;
        /// <summary>
        /// Valid when origin is `EXTERNAL`, it means the effective date of the key material.
        /// </summary>
        public readonly int ValidTo;

        [OutputConstructor]
        private GetKeysKeyListResult(
            string alias,

            int createTime,

            int creatorUin,

            int deletionDate,

            string description,

            string keyId,

            bool keyRotationEnabled,

            string keyState,

            string keyUsage,

            int nextRotateTime,

            string origin,

            string owner,

            int validTo)
        {
            Alias = alias;
            CreateTime = createTime;
            CreatorUin = creatorUin;
            DeletionDate = deletionDate;
            Description = description;
            KeyId = keyId;
            KeyRotationEnabled = keyRotationEnabled;
            KeyState = keyState;
            KeyUsage = keyUsage;
            NextRotateTime = nextRotateTime;
            Origin = origin;
            Owner = owner;
            ValidTo = validTo;
        }
    }
}
