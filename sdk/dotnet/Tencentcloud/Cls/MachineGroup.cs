// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls
{
    /// <summary>
    /// Provides a resource to create a cls machine group.
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
    ///         var @group = new Tencentcloud.Cls.MachineGroup("group", new Tencentcloud.Cls.MachineGroupArgs
    ///         {
    ///             GroupName = "group",
    ///             MachineGroupType = new Tencentcloud.Cls.Inputs.MachineGroupMachineGroupTypeArgs
    ///             {
    ///                 Type = "ip",
    ///                 Values = 
    ///                 {
    ///                     "192.168.1.1",
    ///                     "192.168.1.2",
    ///                 },
    ///             },
    ///             ServiceLogging = true,
    ///             Tags = 
    ///             {
    ///                 { "test", "test1" },
    ///             },
    ///             UpdateEndTime = "19:05:40",
    ///             UpdateStartTime = "17:05:40",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cls machine group can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cls/machineGroup:MachineGroup group caf168e7-32cd-4ac6-bf89-1950a760e09c
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cls/machineGroup:MachineGroup")]
    public partial class MachineGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to enable automatic update for the machine group.
        /// </summary>
        [Output("autoUpdate")]
        public Output<bool?> AutoUpdate { get; private set; } = null!;

        /// <summary>
        /// Machine group name, which must be unique.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// Type of the machine group to be created.
        /// </summary>
        [Output("machineGroupType")]
        public Output<Outputs.MachineGroupMachineGroupType> MachineGroupType { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log topics will be created, which will not incur fees.
        /// </summary>
        [Output("serviceLogging")]
        public Output<bool?> ServiceLogging { get; private set; } = null!;

        /// <summary>
        /// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Update end time. We recommend you update LogListener during off-peak hours.
        /// </summary>
        [Output("updateEndTime")]
        public Output<string?> UpdateEndTime { get; private set; } = null!;

        /// <summary>
        /// Update start time. We recommend you update LogListener during off-peak hours.
        /// </summary>
        [Output("updateStartTime")]
        public Output<string?> UpdateStartTime { get; private set; } = null!;


        /// <summary>
        /// Create a MachineGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineGroup(string name, MachineGroupArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/machineGroup:MachineGroup", name, args ?? new MachineGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineGroup(string name, Input<string> id, MachineGroupState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/machineGroup:MachineGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MachineGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineGroup Get(string name, Input<string> id, MachineGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new MachineGroup(name, id, state, options);
        }
    }

    public sealed class MachineGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable automatic update for the machine group.
        /// </summary>
        [Input("autoUpdate")]
        public Input<bool>? AutoUpdate { get; set; }

        /// <summary>
        /// Machine group name, which must be unique.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Type of the machine group to be created.
        /// </summary>
        [Input("machineGroupType", required: true)]
        public Input<Inputs.MachineGroupMachineGroupTypeArgs> MachineGroupType { get; set; } = null!;

        /// <summary>
        /// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log topics will be created, which will not incur fees.
        /// </summary>
        [Input("serviceLogging")]
        public Input<bool>? ServiceLogging { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Update end time. We recommend you update LogListener during off-peak hours.
        /// </summary>
        [Input("updateEndTime")]
        public Input<string>? UpdateEndTime { get; set; }

        /// <summary>
        /// Update start time. We recommend you update LogListener during off-peak hours.
        /// </summary>
        [Input("updateStartTime")]
        public Input<string>? UpdateStartTime { get; set; }

        public MachineGroupArgs()
        {
        }
    }

    public sealed class MachineGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable automatic update for the machine group.
        /// </summary>
        [Input("autoUpdate")]
        public Input<bool>? AutoUpdate { get; set; }

        /// <summary>
        /// Machine group name, which must be unique.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// Type of the machine group to be created.
        /// </summary>
        [Input("machineGroupType")]
        public Input<Inputs.MachineGroupMachineGroupTypeGetArgs>? MachineGroupType { get; set; }

        /// <summary>
        /// Whether to enable the service log to record the logs generated by the LogListener service itself. After it is enabled, the internal logset cls_service_logging and the loglistener_status, loglistener_alarm, and loglistener_business log topics will be created, which will not incur fees.
        /// </summary>
        [Input("serviceLogging")]
        public Input<bool>? ServiceLogging { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tag description list. Up to 10 tag key-value pairs are supported and must be unique.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Update end time. We recommend you update LogListener during off-peak hours.
        /// </summary>
        [Input("updateEndTime")]
        public Input<string>? UpdateEndTime { get; set; }

        /// <summary>
        /// Update start time. We recommend you update LogListener during off-peak hours.
        /// </summary>
        [Input("updateStartTime")]
        public Input<string>? UpdateStartTime { get; set; }

        public MachineGroupState()
        {
        }
    }
}
