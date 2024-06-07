// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cam
{
    public static class GetUserPolicyAttachments
    {
        /// <summary>
        /// Use this data source to query detailed information of CAM user policy attachments
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Tencentcloud.Cam.GetUserPolicyAttachments.Invoke(new()
        ///     {
        ///         UserId = tencentcloud_cam_user.Foo.Id,
        ///     });
        /// 
        ///     var bar = Tencentcloud.Cam.GetUserPolicyAttachments.Invoke(new()
        ///     {
        ///         UserId = tencentcloud_cam_user.Foo.Id,
        ///         PolicyId = tencentcloud_cam_policy.Foo.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetUserPolicyAttachmentsResult> InvokeAsync(GetUserPolicyAttachmentsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserPolicyAttachmentsResult>("tencentcloud:Cam/getUserPolicyAttachments:getUserPolicyAttachments", args ?? new GetUserPolicyAttachmentsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of CAM user policy attachments
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var foo = Tencentcloud.Cam.GetUserPolicyAttachments.Invoke(new()
        ///     {
        ///         UserId = tencentcloud_cam_user.Foo.Id,
        ///     });
        /// 
        ///     var bar = Tencentcloud.Cam.GetUserPolicyAttachments.Invoke(new()
        ///     {
        ///         UserId = tencentcloud_cam_user.Foo.Id,
        ///         PolicyId = tencentcloud_cam_policy.Foo.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetUserPolicyAttachmentsResult> Invoke(GetUserPolicyAttachmentsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserPolicyAttachmentsResult>("tencentcloud:Cam/getUserPolicyAttachments:getUserPolicyAttachments", args ?? new GetUserPolicyAttachmentsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserPolicyAttachmentsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
        /// </summary>
        [Input("createMode")]
        public int? CreateMode { get; set; }

        /// <summary>
        /// ID of CAM policy to be queried.
        /// </summary>
        [Input("policyId")]
        public string? PolicyId { get; set; }

        /// <summary>
        /// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
        /// </summary>
        [Input("policyType")]
        public string? PolicyType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.59.6. Use `user_name` instead. ID of the attached CAM user to be queried.
        /// </summary>
        [Input("userId")]
        public string? UserId { get; set; }

        /// <summary>
        /// Name of the attached CAM user as unique key to be queried.
        /// </summary>
        [Input("userName")]
        public string? UserName { get; set; }

        public GetUserPolicyAttachmentsArgs()
        {
        }
        public static new GetUserPolicyAttachmentsArgs Empty => new GetUserPolicyAttachmentsArgs();
    }

    public sealed class GetUserPolicyAttachmentsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Mode of Creation of the CAM user policy attachment. `1` means the CAM policy attachment is created by production, and the others indicate syntax strategy ways.
        /// </summary>
        [Input("createMode")]
        public Input<int>? CreateMode { get; set; }

        /// <summary>
        /// ID of CAM policy to be queried.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// It has been deprecated from version 1.59.6. Use `user_name` instead. ID of the attached CAM user to be queried.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        /// <summary>
        /// Name of the attached CAM user as unique key to be queried.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public GetUserPolicyAttachmentsInvokeArgs()
        {
        }
        public static new GetUserPolicyAttachmentsInvokeArgs Empty => new GetUserPolicyAttachmentsInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserPolicyAttachmentsResult
    {
        /// <summary>
        /// Mode of Creation of the CAM user policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
        /// </summary>
        public readonly int? CreateMode;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of CAM user.
        /// </summary>
        public readonly string? PolicyId;
        /// <summary>
        /// Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
        /// </summary>
        public readonly string? PolicyType;
        public readonly string? ResultOutputFile;
        /// <summary>
        /// (**Deprecated**) It has been deprecated from version 1.59.6. Use `user_name` instead. ID of CAM user.
        /// </summary>
        public readonly string? UserId;
        /// <summary>
        /// Name of CAM user as unique key.
        /// </summary>
        public readonly string? UserName;
        /// <summary>
        /// A list of CAM user policy attachments. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserPolicyAttachmentsUserPolicyAttachmentListResult> UserPolicyAttachmentLists;

        [OutputConstructor]
        private GetUserPolicyAttachmentsResult(
            int? createMode,

            string id,

            string? policyId,

            string? policyType,

            string? resultOutputFile,

            string? userId,

            string? userName,

            ImmutableArray<Outputs.GetUserPolicyAttachmentsUserPolicyAttachmentListResult> userPolicyAttachmentLists)
        {
            CreateMode = createMode;
            Id = id;
            PolicyId = policyId;
            PolicyType = policyType;
            ResultOutputFile = resultOutputFile;
            UserId = userId;
            UserName = userName;
            UserPolicyAttachmentLists = userPolicyAttachmentLists;
        }
    }
}
