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
    public static class GetInstanceInfo
    {
        /// <summary>
        /// Use this data source to query detailed information of mysql instance_info
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
        ///     var instanceInfo = Tencentcloud.Mysql.GetInstanceInfo.Invoke(new()
        ///     {
        ///         InstanceId = "cdb-fitq5t9h",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetInstanceInfoResult> InvokeAsync(GetInstanceInfoArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceInfoResult>("tencentcloud:Mysql/getInstanceInfo:getInstanceInfo", args ?? new GetInstanceInfoArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of mysql instance_info
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
        ///     var instanceInfo = Tencentcloud.Mysql.GetInstanceInfo.Invoke(new()
        ///     {
        ///         InstanceId = "cdb-fitq5t9h",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetInstanceInfoResult> Invoke(GetInstanceInfoInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceInfoResult>("tencentcloud:Mysql/getInstanceInfo:getInstanceInfo", args ?? new GetInstanceInfoInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceInfoArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        public GetInstanceInfoArgs()
        {
        }
        public static new GetInstanceInfoArgs Empty => new GetInstanceInfoArgs();
    }

    public sealed class GetInstanceInfoInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// instance id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        public GetInstanceInfoInvokeArgs()
        {
        }
        public static new GetInstanceInfoInvokeArgs Empty => new GetInstanceInfoInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceInfoResult
    {
        /// <summary>
        /// The default region of the KMS service used by the current CDB backend service.
        /// </summary>
        public readonly string DefaultKmsRegion;
        /// <summary>
        /// Whether to enable encryption, YES is enabled, NO is not enabled.
        /// </summary>
        public readonly string Encryption;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceId;
        /// <summary>
        /// instance name.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The key ID used for encryption.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// The region where the key is located.
        /// </summary>
        public readonly string KeyRegion;
        public readonly string? ResultOutputFile;

        [OutputConstructor]
        private GetInstanceInfoResult(
            string defaultKmsRegion,

            string encryption,

            string id,

            string instanceId,

            string instanceName,

            string keyId,

            string keyRegion,

            string? resultOutputFile)
        {
            DefaultKmsRegion = defaultKmsRegion;
            Encryption = encryption;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            KeyId = keyId;
            KeyRegion = keyRegion;
            ResultOutputFile = resultOutputFile;
        }
    }
}
