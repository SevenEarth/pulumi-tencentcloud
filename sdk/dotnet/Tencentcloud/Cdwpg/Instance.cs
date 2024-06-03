// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cdwpg
{
    /// <summary>
    /// Provides a resource to create a cdwpg instance
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Tencentcloud.Cdwpg.Instance("instance", new()
    ///     {
    ///         AdminPassword = "xxxxxx",
    ///         ChargeProperties = new Tencentcloud.Cdwpg.Inputs.InstanceChargePropertiesArgs
    ///         {
    ///             ChargeType = "POSTPAID_BY_HOUR",
    ///             RenewFlag = 0,
    ///             TimeSpan = 1,
    ///             TimeUnit = "h",
    ///         },
    ///         InstanceName = "test_cdwpg",
    ///         Resources = new[]
    ///         {
    ///             new Tencentcloud.Cdwpg.Inputs.InstanceResourceArgs
    ///             {
    ///                 Count = 2,
    ///                 DiskSpec = new Tencentcloud.Cdwpg.Inputs.InstanceResourceDiskSpecArgs
    ///                 {
    ///                     DiskCount = 1,
    ///                     DiskSize = 200,
    ///                     DiskType = "CLOUD_HSSD",
    ///                 },
    ///                 SpecName = "S_4_16_H_CN",
    ///                 Type = "cn",
    ///             },
    ///             new Tencentcloud.Cdwpg.Inputs.InstanceResourceArgs
    ///             {
    ///                 Count = 2,
    ///                 DiskSpec = new Tencentcloud.Cdwpg.Inputs.InstanceResourceDiskSpecArgs
    ///                 {
    ///                     DiskCount = 10,
    ///                     DiskSize = 20,
    ///                     DiskType = "CLOUD_HSSD",
    ///                 },
    ///                 SpecName = "S_4_16_H_CN",
    ///                 Type = "dn",
    ///             },
    ///         },
    ///         Tags = 
    ///         {
    ///             { "tagKey", "tagValue" },
    ///         },
    ///         UserSubnetId = "subnet-xxxxxx",
    ///         UserVpcId = "vpc-xxxxxx",
    ///         Zone = "ap-guangzhou-6",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// cdwpg instance can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Cdwpg/instance:Instance instance instance_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cdwpg/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// cluster password.
        /// </summary>
        [Output("adminPassword")]
        public Output<string> AdminPassword { get; private set; } = null!;

        /// <summary>
        /// instance billing mode.
        /// </summary>
        [Output("chargeProperties")]
        public Output<Outputs.InstanceChargeProperties> ChargeProperties { get; private set; } = null!;

        /// <summary>
        /// instance name.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// resource information.
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<Outputs.InstanceResource>> Resources { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// subnet.
        /// </summary>
        [Output("userSubnetId")]
        public Output<string> UserSubnetId { get; private set; } = null!;

        /// <summary>
        /// private network.
        /// </summary>
        [Output("userVpcId")]
        public Output<string> UserVpcId { get; private set; } = null!;

        /// <summary>
        /// Availability Zone.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdwpg/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cdwpg/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
                AdditionalSecretOutputs =
                {
                    "adminPassword",
                },
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

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("adminPassword", required: true)]
        private Input<string>? _adminPassword;

        /// <summary>
        /// cluster password.
        /// </summary>
        public Input<string>? AdminPassword
        {
            get => _adminPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _adminPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// instance billing mode.
        /// </summary>
        [Input("chargeProperties", required: true)]
        public Input<Inputs.InstanceChargePropertiesArgs> ChargeProperties { get; set; } = null!;

        /// <summary>
        /// instance name.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        [Input("resources", required: true)]
        private InputList<Inputs.InstanceResourceArgs>? _resources;

        /// <summary>
        /// resource information.
        /// </summary>
        public InputList<Inputs.InstanceResourceArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.InstanceResourceArgs>());
            set => _resources = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// subnet.
        /// </summary>
        [Input("userSubnetId", required: true)]
        public Input<string> UserSubnetId { get; set; } = null!;

        /// <summary>
        /// private network.
        /// </summary>
        [Input("userVpcId", required: true)]
        public Input<string> UserVpcId { get; set; } = null!;

        /// <summary>
        /// Availability Zone.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        [Input("adminPassword")]
        private Input<string>? _adminPassword;

        /// <summary>
        /// cluster password.
        /// </summary>
        public Input<string>? AdminPassword
        {
            get => _adminPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _adminPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// instance billing mode.
        /// </summary>
        [Input("chargeProperties")]
        public Input<Inputs.InstanceChargePropertiesGetArgs>? ChargeProperties { get; set; }

        /// <summary>
        /// instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("resources")]
        private InputList<Inputs.InstanceResourceGetArgs>? _resources;

        /// <summary>
        /// resource information.
        /// </summary>
        public InputList<Inputs.InstanceResourceGetArgs> Resources
        {
            get => _resources ?? (_resources = new InputList<Inputs.InstanceResourceGetArgs>());
            set => _resources = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// subnet.
        /// </summary>
        [Input("userSubnetId")]
        public Input<string>? UserSubnetId { get; set; }

        /// <summary>
        /// private network.
        /// </summary>
        [Input("userVpcId")]
        public Input<string>? UserVpcId { get; set; }

        /// <summary>
        /// Availability Zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
