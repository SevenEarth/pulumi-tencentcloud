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
    /// <summary>
    /// Provides a resource to create a CAM SAML provider.
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
    ///         var samlProviderBasic = new Tencentcloud.Cam.SamlProvider("samlProviderBasic", new Tencentcloud.Cam.SamlProviderArgs
    ///         {
    ///             Description = "test",
    ///             MetaData = "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48bWQ6RW50aXR5RGVzY3JpcHRvciBlbnRpdHlJRD0iaHR0cDovL3d3dy5va3RhLmNvbS9leGsxa3F4bWNqUW1HQURNeTM1NyIgeG1sbnM6bWQ9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDptZXRhZGF0YSI+PG1kOklEUFNTT0Rlc2NyaXB0b3IgV2FudEF1dGhuUmVxdWVzdHNTaWduZWQ9ImZhbHNlIiBwcm90b2NvbFN1cHBvcnRFbnVtZXJhdGlvbj0idXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOnByb3RvY29sIj48bWQ6S2V5RGVzY3JpcHRvciB1c2U9InNpZ25pbmciPjxkczpLZXlJbmZvIHhtbG5zOmRzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwLzA5L3htbGRzaWcjIj48ZHM6WDUwOURhdGE+PGRzOlg1MDlDZXJ0aWZpY2F0ZT5NSUlEb0RDQ0FvaWdBd0lCQWdJR0FXM0lTcExvTUEwR0NTcUdTSWIzRFFFQkN3VUFNSUdRTVFzd0NRWURWUVFHRXdKVlV6RVRNQkVHDQpBMVVFQ0F3S1EyRnNhV1p2Y201cFlURVdNQlFHQTFVRUJ3d05VMkZ1SUVaeVlXNWphWE5qYnpFTk1Bc0dBMVVFQ2d3RVQydDBZVEVVDQpNQklHQTFVRUN3d0xVMU5QVUhKdmRtbGtaWEl4RVRBUEJnTlZCQU1NQ0dsa2VIVmxkblJoTVJ3d0dnWUpLb1pJaHZjTkFRa0JGZzFwDQpibVp2UUc5cmRHRXVZMjl0TUI0WERURTVNVEF4TkRBek1qSXhNMW9YRFRJNU1UQXhOREF6TWpNeE0xb3dnWkF4Q3pBSkJnTlZCQVlUDQpBbFZUTVJNd0VRWURWUVFJREFwRFlXeHBabTl5Ym1saE1SWXdGQVlEVlFRSERBMVRZVzRnUm5KaGJtTnBjMk52TVEwd0N3WURWUVFLDQpEQVJQYTNSaE1SUXdFZ1lEVlFRTERBdFRVMDlRY205MmFXUmxjakVSTUE4R0ExVUVBd3dJYVdSNGRXVjJkR0V4SERBYUJna3Foa2lHDQo5dzBCQ1FFV0RXbHVabTlBYjJ0MFlTNWpiMjB3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRQ2g4b3dqDQpZK2dQSUM3blQvNTduLzdmeXJzcDlHMXdxa2UxdXhjMHVrTndnQXozOVNpelY3QVhLMWRReTFLaThXWjJJMzFEczJkT0FNQ1FKR2pWDQpUWWNNbnA3KzhqUzNLdmxNUkRJamk5cmxuUi9vcnBvMll1RHVWby9jVzdidlRIS2h2REo1QWZRaWxzYlNPTXdUOWM2TVlYZGhBNVBwDQpzelFsK1UrdHJmcXUrdUorSER4SVQxdlhWaVI5YlY2SUFRSzZpbWZoc2wxWmVSUytjbVFVNEpjQWlYT0xtTnFVVWM2UkpxUzhrMW1mDQpBLzhmb2VyMGc3SG4xZDVXclpCc2gyUlR2Vzh1ZVdadHQ3dmh4QTlGdE5kSVlEcXJ0eElmMlZXcXhrSHM3WFZDSm5wTnJITVovT1BRDQpGY21YSGVxNlJJMlB3Q1RlOW8zZHZpM0hqeXBaOEl4dkFnTUJBQUV3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQUFHaHk1bG9nbGtTDQoyVHg2YS90MnF5VEx0YVV5cEwrNGhySGJoMVAweVVMc0NrSnFsM2wrWG9VZDZCY2FJaFNSVGFPQk95ODViL0UzelJ4K3JzQXJwTjVVDQp5ZThuUEM4a05PYW5vTk9wWnZvYmhpTzFlMFIvYmxEcnRBL0o5UlBwMWtmdlhmS2NSTTU3TlRCWXppTURlbnFQUTRFOWN1U2lGdFFxDQpJYmpIbThaM1B1YXgwRitldkZ3U1pJMDNCWXNISGw1d1EraEJBS3hTdTJINEZRdU93Zmpnb2EveEN6Z1NKYjJ2UXdEc1MxMk9mSkNiDQpSRm1ZL1VYZXQramFhdEVORktLZStZSUJpU0J2WG1adTN0MHN5NDZTNzlPVzBacXJ0NUh2bElsT2lpTFpaN1FZamxjM1kxeG1LZ1luDQpXM2M2WGZkdmhGWHo0ZDdkbWYvTUdpNGY0enM9PC9kczpYNTA5Q2VydGlmaWNhdGU+PC9kczpYNTA5RGF0YT48L2RzOktleUluZm8+PC9tZDpLZXlEZXNjcmlwdG9yPjxtZDpOYW1lSURGb3JtYXQ+dXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6MS4xOm5hbWVpZC1mb3JtYXQ6dW5zcGVjaWZpZWQ8L21kOk5hbWVJREZvcm1hdD48bWQ6TmFtZUlERm9ybWF0PnVybjpvYXNpczpuYW1lczp0YzpTQU1MOjEuMTpuYW1laWQtZm9ybWF0OmVtYWlsQWRkcmVzczwvbWQ6TmFtZUlERm9ybWF0PjxtZDpTaW5nbGVTaWduT25TZXJ2aWNlIEJpbmRpbmc9InVybjpvYXNpczpuYW1lczp0YzpTQU1MOjIuMDpiaW5kaW5nczpIVFRQLVBPU1QiIExvY2F0aW9uPSJodHRwczovL2lkeHVldnRhLm9rdGEuY29tL2FwcC9pZHh1ZW9yZzYzNzM1OF90ZXN0XzEvZXhrMWtxeG1jalFtR0FETXkzNTcvc3NvL3NhbWwiLz48bWQ6U2luZ2xlU2lnbk9uU2VydmljZSBCaW5kaW5nPSJ1cm46b2FzaXM6bmFtZXM6dGM6U0FNTDoyLjA6YmluZGluZ3M6SFRUUC1SZWRpcmVjdCIgTG9jYXRpb249Imh0dHBzOi8vaWR4dWV2dGEub2t0YS5jb20vYXBwL2lkeHVlb3JnNjM3MzU4X3Rlc3RfMS9leGsxa3F4bWNqUW1HQURNeTM1Ny9zc28vc2FtbCIvPjwvbWQ6SURQU1NPRGVzY3JpcHRvcj48L21kOkVudGl0eURlc2NyaXB0b3I+",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CAM SAML provider can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Cam/samlProvider:SamlProvider foo cam-SAML-provider-test
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Cam/samlProvider:SamlProvider")]
    public partial class SamlProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// The create time of the CAM SAML provider.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The description of the CAM SAML provider.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The meta data document of the CAM SAML provider.
        /// </summary>
        [Output("metaData")]
        public Output<string> MetaData { get; private set; } = null!;

        /// <summary>
        /// Name of CAM SAML provider.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ARN of the CAM SAML provider.
        /// </summary>
        [Output("providerArn")]
        public Output<string> ProviderArn { get; private set; } = null!;

        /// <summary>
        /// The last update time of the CAM SAML provider.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a SamlProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SamlProvider(string name, SamlProviderArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/samlProvider:SamlProvider", name, args ?? new SamlProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SamlProvider(string name, Input<string> id, SamlProviderState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cam/samlProvider:SamlProvider", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SamlProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SamlProvider Get(string name, Input<string> id, SamlProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new SamlProvider(name, id, state, options);
        }
    }

    public sealed class SamlProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the CAM SAML provider.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The meta data document of the CAM SAML provider.
        /// </summary>
        [Input("metaData", required: true)]
        public Input<string> MetaData { get; set; } = null!;

        /// <summary>
        /// Name of CAM SAML provider.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SamlProviderArgs()
        {
        }
    }

    public sealed class SamlProviderState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The create time of the CAM SAML provider.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The description of the CAM SAML provider.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The meta data document of the CAM SAML provider.
        /// </summary>
        [Input("metaData")]
        public Input<string>? MetaData { get; set; }

        /// <summary>
        /// Name of CAM SAML provider.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ARN of the CAM SAML provider.
        /// </summary>
        [Input("providerArn")]
        public Input<string>? ProviderArn { get; set; }

        /// <summary>
        /// The last update time of the CAM SAML provider.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public SamlProviderState()
        {
        }
    }
}
