// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql
{
    /// <summary>
    /// Provides a resource to create a mysql instance_encryption_operation
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zones = Tencentcloud.Availability.GetZonesByProduct.Invoke(new()
    ///     {
    ///         Product = "cdb",
    ///     });
    /// 
    ///     var vpc = new Tencentcloud.Vpc.Instance("vpc", new()
    ///     {
    ///         CidrBlock = "10.0.0.0/16",
    ///     });
    /// 
    ///     var subnet = new Tencentcloud.Subnet.Instance("subnet", new()
    ///     {
    ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[0]?.Name),
    ///         VpcId = vpc.Id,
    ///         CidrBlock = "10.0.0.0/16",
    ///         IsMulticast = false,
    ///     });
    /// 
    ///     var securityGroup = new Tencentcloud.Security.Group("securityGroup", new()
    ///     {
    ///         Description = "mysql test",
    ///     });
    /// 
    ///     var exampleInstance = new Tencentcloud.Mysql.Instance("exampleInstance", new()
    ///     {
    ///         InternetService = 1,
    ///         EngineVersion = "5.7",
    ///         ChargeType = "POSTPAID",
    ///         RootPassword = "PassWord123",
    ///         SlaveDeployMode = 0,
    ///         AvailabilityZone = zones.Apply(getZonesByProductResult =&gt; getZonesByProductResult.Zones[0]?.Name),
    ///         SlaveSyncMode = 1,
    ///         InstanceName = "tf-example-mysql",
    ///         MemSize = 4000,
    ///         VolumeSize = 200,
    ///         VpcId = vpc.Id,
    ///         SubnetId = subnet.Id,
    ///         IntranetPort = 3306,
    ///         SecurityGroups = new[]
    ///         {
    ///             securityGroup.Id,
    ///         },
    ///         Tags = 
    ///         {
    ///             { "name", "test" },
    ///         },
    ///         Parameters = 
    ///         {
    ///             { "character_set_server", "utf8" },
    ///             { "max_connections", "1000" },
    ///         },
    ///     });
    /// 
    ///     var exampleInstanceEncryptionOperation = new Tencentcloud.Mysql.InstanceEncryptionOperation("exampleInstanceEncryptionOperation", new()
    ///     {
    ///         InstanceId = exampleInstance.Id,
    ///         KeyId = "KMS-CDB",
    ///         KeyRegion = "ap-guangzhou",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Mysql/instanceEncryptionOperation:InstanceEncryptionOperation")]
    public partial class InstanceEncryptionOperation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// TencentDB instance ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
        /// </summary>
        [Output("keyId")]
        public Output<string?> KeyId { get; private set; } = null!;

        /// <summary>
        /// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
        /// </summary>
        [Output("keyRegion")]
        public Output<string?> KeyRegion { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceEncryptionOperation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceEncryptionOperation(string name, InstanceEncryptionOperationArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/instanceEncryptionOperation:InstanceEncryptionOperation", name, args ?? new InstanceEncryptionOperationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceEncryptionOperation(string name, Input<string> id, InstanceEncryptionOperationState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Mysql/instanceEncryptionOperation:InstanceEncryptionOperation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceEncryptionOperation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceEncryptionOperation Get(string name, Input<string> id, InstanceEncryptionOperationState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceEncryptionOperation(name, id, state, options);
        }
    }

    public sealed class InstanceEncryptionOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// TencentDB instance ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
        /// </summary>
        [Input("keyRegion")]
        public Input<string>? KeyRegion { get; set; }

        public InstanceEncryptionOperationArgs()
        {
        }
        public static new InstanceEncryptionOperationArgs Empty => new InstanceEncryptionOperationArgs();
    }

    public sealed class InstanceEncryptionOperationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// TencentDB instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Custom key ID, which is the unique CMK ID. If this value is empty, the key KMS-CDB auto-generated by Tencent Cloud will be used.
        /// </summary>
        [Input("keyId")]
        public Input<string>? KeyId { get; set; }

        /// <summary>
        /// Custom storage region, such as ap-guangzhou. When `KeyId` is not empty, this parameter is required.
        /// </summary>
        [Input("keyRegion")]
        public Input<string>? KeyRegion { get; set; }

        public InstanceEncryptionOperationState()
        {
        }
        public static new InstanceEncryptionOperationState Empty => new InstanceEncryptionOperationState();
    }
}
