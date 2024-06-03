// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts
{
    /// <summary>
    /// Provides a resource to create a pts alert_channel
    /// 
    /// &gt; **NOTE:** Modification is not currently supported, please go to the console to modify.
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
    ///     var example = new Tencentcloud.Monitor.AlarmNotice("example", new()
    ///     {
    ///         NoticeType = "ALL",
    ///         NoticeLanguage = "zh-CN",
    ///         UserNotices = new[]
    ///         {
    ///             new Tencentcloud.Monitor.Inputs.AlarmNoticeUserNoticeArgs
    ///             {
    ///                 ReceiverType = "USER",
    ///                 StartTime = 0,
    ///                 EndTime = 1,
    ///                 NoticeWays = new[]
    ///                 {
    ///                     "EMAIL",
    ///                     "SMS",
    ///                     "WECHAT",
    ///                 },
    ///                 UserIds = new[]
    ///                 {
    ///                     10001,
    ///                 },
    ///                 GroupIds = new() { },
    ///                 PhoneOrders = new[]
    ///                 {
    ///                     10001,
    ///                 },
    ///                 PhoneCircleTimes = 2,
    ///                 PhoneCircleInterval = 50,
    ///                 PhoneInnerInterval = 60,
    ///                 NeedPhoneArriveNotice = 1,
    ///                 PhoneCallType = "CIRCLE",
    ///                 Weekdays = new[]
    ///                 {
    ///                     1,
    ///                     2,
    ///                     3,
    ///                     4,
    ///                     5,
    ///                     6,
    ///                     7,
    ///                 },
    ///             },
    ///         },
    ///         UrlNotices = new[]
    ///         {
    ///             new Tencentcloud.Monitor.Inputs.AlarmNoticeUrlNoticeArgs
    ///             {
    ///                 Url = "https://www.mytest.com/validate",
    ///                 EndTime = 0,
    ///                 StartTime = 1,
    ///                 Weekdays = new[]
    ///                 {
    ///                     1,
    ///                     2,
    ///                     3,
    ///                     4,
    ///                     5,
    ///                     6,
    ///                     7,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var project = new Tencentcloud.Pts.Project("project", new()
    ///     {
    ///         Description = "desc",
    ///         Tags = new[]
    ///         {
    ///             new Tencentcloud.Pts.Inputs.ProjectTagArgs
    ///             {
    ///                 TagKey = "createdBy",
    ///                 TagValue = "terraform",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var alertChannel = new Tencentcloud.Pts.AlertChannel("alertChannel", new()
    ///     {
    ///         NoticeId = example.Id,
    ///         ProjectId = project.Id,
    ///         AmpConsumerId = "Consumer-vvy1xxxxxx",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// pts alert_channel can be imported using the project_id#notice_id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Pts/alertChannel:AlertChannel alert_channel project-kww5v8se#notice-kl66t6y9
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Pts/alertChannel:AlertChannel")]
    public partial class AlertChannel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AMP Consumer ID.
        /// </summary>
        [Output("ampConsumerId")]
        public Output<string?> AmpConsumerId { get; private set; } = null!;

        /// <summary>
        /// App ID Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Output("appId")]
        public Output<int> AppId { get; private set; } = null!;

        /// <summary>
        /// Creation time Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Notice ID.
        /// </summary>
        [Output("noticeId")]
        public Output<string> NoticeId { get; private set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Status Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Sub-user ID Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Output("subAccountUin")]
        public Output<string> SubAccountUin { get; private set; } = null!;

        /// <summary>
        /// User ID Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Output("uin")]
        public Output<string> Uin { get; private set; } = null!;

        /// <summary>
        /// Update time Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a AlertChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertChannel(string name, AlertChannelArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/alertChannel:AlertChannel", name, args ?? new AlertChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertChannel(string name, Input<string> id, AlertChannelState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/alertChannel:AlertChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertChannel Get(string name, Input<string> id, AlertChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertChannel(name, id, state, options);
        }
    }

    public sealed class AlertChannelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AMP Consumer ID.
        /// </summary>
        [Input("ampConsumerId")]
        public Input<string>? AmpConsumerId { get; set; }

        /// <summary>
        /// Notice ID.
        /// </summary>
        [Input("noticeId", required: true)]
        public Input<string> NoticeId { get; set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        public AlertChannelArgs()
        {
        }
        public static new AlertChannelArgs Empty => new AlertChannelArgs();
    }

    public sealed class AlertChannelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AMP Consumer ID.
        /// </summary>
        [Input("ampConsumerId")]
        public Input<string>? AmpConsumerId { get; set; }

        /// <summary>
        /// App ID Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("appId")]
        public Input<int>? AppId { get; set; }

        /// <summary>
        /// Creation time Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Notice ID.
        /// </summary>
        [Input("noticeId")]
        public Input<string>? NoticeId { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Status Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Sub-user ID Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("subAccountUin")]
        public Input<string>? SubAccountUin { get; set; }

        /// <summary>
        /// User ID Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("uin")]
        public Input<string>? Uin { get; set; }

        /// <summary>
        /// Update time Note: this field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public AlertChannelState()
        {
        }
        public static new AlertChannelState Empty => new AlertChannelState();
    }
}
