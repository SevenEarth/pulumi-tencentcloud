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
    /// Provides a resource to create a cls alarm_notice
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
    ///         var alarmNotice = new Tencentcloud.Cls.AlarmNotice("alarmNotice", new Tencentcloud.Cls.AlarmNoticeArgs
    ///         {
    ///             NoticeReceivers = 
    ///             {
    ///                 new Tencentcloud.Cls.Inputs.AlarmNoticeNoticeReceiverArgs
    ///                 {
    ///                     EndTime = "23:59:59",
    ///                     Index = 0,
    ///                     ReceiverChannels = 
    ///                     {
    ///                         "Sms",
    ///                     },
    ///                     ReceiverIds = 
    ///                     {
    ///                         13478043,
    ///                         15972111,
    ///                     },
    ///                     ReceiverType = "Uin",
    ///                     StartTime = "00:00:00",
    ///                 },
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "createdBy", "terraform" },
    ///             },
    ///             Type = "All",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// cls alarm_notice can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cls/alarmNotice:AlarmNotice alarm_notice alarm_notice_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cls/alarmNotice:AlarmNotice")]
    public partial class AlarmNotice : Pulumi.CustomResource
    {
        /// <summary>
        /// alarm notice name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// notice receivers.
        /// </summary>
        [Output("noticeReceivers")]
        public Output<ImmutableArray<Outputs.AlarmNoticeNoticeReceiver>> NoticeReceivers { get; private set; } = null!;

        /// <summary>
        /// Tag description list.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// notice type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// callback info.
        /// </summary>
        [Output("webCallbacks")]
        public Output<ImmutableArray<Outputs.AlarmNoticeWebCallback>> WebCallbacks { get; private set; } = null!;


        /// <summary>
        /// Create a AlarmNotice resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlarmNotice(string name, AlarmNoticeArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/alarmNotice:AlarmNotice", name, args ?? new AlarmNoticeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlarmNotice(string name, Input<string> id, AlarmNoticeState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cls/alarmNotice:AlarmNotice", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlarmNotice resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlarmNotice Get(string name, Input<string> id, AlarmNoticeState? state = null, CustomResourceOptions? options = null)
        {
            return new AlarmNotice(name, id, state, options);
        }
    }

    public sealed class AlarmNoticeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// alarm notice name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("noticeReceivers")]
        private InputList<Inputs.AlarmNoticeNoticeReceiverArgs>? _noticeReceivers;

        /// <summary>
        /// notice receivers.
        /// </summary>
        public InputList<Inputs.AlarmNoticeNoticeReceiverArgs> NoticeReceivers
        {
            get => _noticeReceivers ?? (_noticeReceivers = new InputList<Inputs.AlarmNoticeNoticeReceiverArgs>());
            set => _noticeReceivers = value;
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
        /// notice type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("webCallbacks")]
        private InputList<Inputs.AlarmNoticeWebCallbackArgs>? _webCallbacks;

        /// <summary>
        /// callback info.
        /// </summary>
        public InputList<Inputs.AlarmNoticeWebCallbackArgs> WebCallbacks
        {
            get => _webCallbacks ?? (_webCallbacks = new InputList<Inputs.AlarmNoticeWebCallbackArgs>());
            set => _webCallbacks = value;
        }

        public AlarmNoticeArgs()
        {
        }
    }

    public sealed class AlarmNoticeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// alarm notice name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("noticeReceivers")]
        private InputList<Inputs.AlarmNoticeNoticeReceiverGetArgs>? _noticeReceivers;

        /// <summary>
        /// notice receivers.
        /// </summary>
        public InputList<Inputs.AlarmNoticeNoticeReceiverGetArgs> NoticeReceivers
        {
            get => _noticeReceivers ?? (_noticeReceivers = new InputList<Inputs.AlarmNoticeNoticeReceiverGetArgs>());
            set => _noticeReceivers = value;
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
        /// notice type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("webCallbacks")]
        private InputList<Inputs.AlarmNoticeWebCallbackGetArgs>? _webCallbacks;

        /// <summary>
        /// callback info.
        /// </summary>
        public InputList<Inputs.AlarmNoticeWebCallbackGetArgs> WebCallbacks
        {
            get => _webCallbacks ?? (_webCallbacks = new InputList<Inputs.AlarmNoticeWebCallbackGetArgs>());
            set => _webCallbacks = value;
        }

        public AlarmNoticeState()
        {
        }
    }
}
