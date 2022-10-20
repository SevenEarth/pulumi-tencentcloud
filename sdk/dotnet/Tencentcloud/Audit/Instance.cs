// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Audit
{
    /// <summary>
    /// Provides a resource to create an audit.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Audit.Instance("foo", new Tencentcloud.Audit.InstanceArgs
    ///         {
    ///             AuditSwitch = true,
    ///             CosBucket = "test",
    ///             CosRegion = "ap-hongkong",
    ///             LogFilePrefix = "test",
    ///             ReadWriteAttribute = 3,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Audit can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Audit/instance:Instance foo audit-test
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Audit/instance:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicate whether to turn on audit logging or not.
        /// </summary>
        [Output("auditSwitch")]
        public Output<bool> AuditSwitch { get; private set; } = null!;

        /// <summary>
        /// Name of the cos bucket to save audit log. Caution: the validation of existing cos bucket will not be checked by
        /// terraform.
        /// </summary>
        [Output("cosBucket")]
        public Output<string> CosBucket { get; private set; } = null!;

        /// <summary>
        /// Region of the cos bucket.
        /// </summary>
        [Output("cosRegion")]
        public Output<string> CosRegion { get; private set; } = null!;

        /// <summary>
        /// Indicate whether the log is encrypt with KMS algorithm or not.
        /// </summary>
        [Output("enableKmsEncry")]
        public Output<bool?> EnableKmsEncry { get; private set; } = null!;

        /// <summary>
        /// Existing CMK unique key. This field can be get by data source `tencentcloud.Audit.getKeyAlias`. Caution: the region of the KMS must be as same as the `cos_region`.
        /// </summary>
        [Output("keyId")]
        public Output<string?> KeyId { get; private set; } = null!;

        /// <summary>
        /// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
        /// </summary>
        [Output("logFilePrefix")]
        public Output<string> LogFilePrefix { get; private set; } = null!;

        /// <summary>
        /// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Event attribute filter. Valid values: `1`, `2`, `3`. `1` for readonly, `2` for write-only, `3` for all.
        /// </summary>
        [Output("readWriteAttribute")]
        public Output<int> ReadWriteAttribute { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Audit/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Audit/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicate whether to turn on audit logging or not.
        /// </summary>
        [Input("auditSwitch", required: true)]
        public Input<bool> AuditSwitch { get; set; } = null!;

        /// <summary>
        /// Name of the cos bucket to save audit log. Caution: the validation of existing cos bucket will not be checked by
        /// terraform.
        /// </summary>
        [Input("cosBucket", required: true)]
        public Input<string> CosBucket { get; set; } = null!;

        /// <summary>
        /// Region of the cos bucket.
        /// </summary>
        [Input("cosRegion", required: true)]
        public Input<string> CosRegion { get; set; } = null!;

        /// <summary>
        /// Indicate whether the log is encrypt with KMS algorithm or not.
        /// </summary>
        [Input("enableKmsEncry")]
        public Input<bool>? EnableKmsEncry { get; set; }

        /// <summary>
        /// Existing CMK unique key. This field can be get by data source `tencentcloud.Audit.getKeyAlias`. Caution: the region of the KMS must be as same as the `cos_region`.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
        /// </summary>
        [Input("logFilePrefix")]
        public Input<string>? LogFilePrefix { get; set; }

        /// <summary>
        /// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Event attribute filter. Valid values: `1`, `2`, `3`. `1` for readonly, `2` for write-only, `3` for all.
        /// </summary>
        [Input("readWriteAttribute", required: true)]
        public Input<int> ReadWriteAttribute { get; set; } = null!;

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicate whether to turn on audit logging or not.
        /// </summary>
        [Input("auditSwitch")]
        public Input<bool>? AuditSwitch { get; set; }

        /// <summary>
        /// Name of the cos bucket to save audit log. Caution: the validation of existing cos bucket will not be checked by
        /// terraform.
        /// </summary>
        [Input("cosBucket")]
        public Input<string>? CosBucket { get; set; }

        /// <summary>
        /// Region of the cos bucket.
        /// </summary>
        [Input("cosRegion")]
        public Input<string>? CosRegion { get; set; }

        /// <summary>
        /// Indicate whether the log is encrypt with KMS algorithm or not.
        /// </summary>
        [Input("enableKmsEncry")]
        public Input<bool>? EnableKmsEncry { get; set; }

        /// <summary>
        /// Existing CMK unique key. This field can be get by data source `tencentcloud.Audit.getKeyAlias`. Caution: the region of the KMS must be as same as the `cos_region`.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// The log file name prefix. The length ranges from 3 to 40. If not set, the account ID will be the log file prefix.
        /// </summary>
        [Input("logFilePrefix")]
        public Input<string>? LogFilePrefix { get; set; }

        /// <summary>
        /// Name of audit. Valid length ranges from 3 to 128. Only alpha character or numbers or '_' supported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Event attribute filter. Valid values: `1`, `2`, `3`. `1` for readonly, `2` for write-only, `3` for all.
        /// </summary>
        [Input("readWriteAttribute")]
        public Input<int>? ReadWriteAttribute { get; set; }

        public InstanceState()
        {
        }
    }
}
