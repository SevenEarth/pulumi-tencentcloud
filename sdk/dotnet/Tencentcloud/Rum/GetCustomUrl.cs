// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Rum
{
    public static class GetCustomUrl
    {
        /// <summary>
        /// Use this data source to query detailed information of rum custom_url
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var customUrl = Output.Create(Tencentcloud.Rum.GetCustomUrl.InvokeAsync(new Tencentcloud.Rum.GetCustomUrlArgs
        ///         {
        ///             EndTime = 1625454840,
        ///             ProjectId = 1,
        ///             StartTime = 1625444040,
        ///             Type = "top",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCustomUrlResult> InvokeAsync(GetCustomUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomUrlResult>("tencentcloud:Rum/getCustomUrl:getCustomUrl", args ?? new GetCustomUrlArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of rum custom_url
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tencentcloud = Pulumi.Tencentcloud;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var customUrl = Output.Create(Tencentcloud.Rum.GetCustomUrl.InvokeAsync(new Tencentcloud.Rum.GetCustomUrlArgs
        ///         {
        ///             EndTime = 1625454840,
        ///             ProjectId = 1,
        ///             StartTime = 1625444040,
        ///             Type = "top",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCustomUrlResult> Invoke(GetCustomUrlInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCustomUrlResult>("tencentcloud:Rum/getCustomUrl:getCustomUrl", args ?? new GetCustomUrlInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomUrlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region where the data reporting takes place.
        /// </summary>
        [Input("area")]
        public string? Area { get; set; }

        /// <summary>
        /// The mobile phone brand used for data reporting.
        /// </summary>
        [Input("brand")]
        public string? Brand { get; set; }

        /// <summary>
        /// The browser type used for data reporting.
        /// </summary>
        [Input("browser")]
        public string? Browser { get; set; }

        /// <summary>
        /// The method used for calculating the elapsed time `50`: 50th percentile, `75`: 75th percentile, `90`: 90th percentile, `95`: 95th percentile, `99`: 99th percentile, `99.5`: 99.5th percentile, `avg`: Mean.
        /// </summary>
        [Input("costType")]
        public string? CostType { get; set; }

        /// <summary>
        /// The device used for data reporting.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// End time but is represented using a timestamp in seconds.
        /// </summary>
        [Input("endTime", required: true)]
        public int EndTime { get; set; }

        /// <summary>
        /// The browser engine used for data reporting.
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        /// <summary>
        /// The code environment where the data reporting takes place.(`production`: production env, `development`: development env, `gray`: gray env, `pre`: pre env, `daily`: daily env, `local`: local env, `others`: others env).
        /// </summary>
        [Input("env")]
        public string? Env { get; set; }

        /// <summary>
        /// First Expansion parameter.
        /// </summary>
        [Input("extFirst")]
        public string? ExtFirst { get; set; }

        /// <summary>
        /// Second Expansion parameter.
        /// </summary>
        [Input("extSecond")]
        public string? ExtSecond { get; set; }

        /// <summary>
        /// Third Expansion parameter.
        /// </summary>
        [Input("extThird")]
        public string? ExtThird { get; set; }

        /// <summary>
        /// The source page of the data reporting.
        /// </summary>
        [Input("from")]
        public string? From { get; set; }

        /// <summary>
        /// Whether it is non-China region.`1`: yes; `0`: no.
        /// </summary>
        [Input("isAbroad")]
        public string? IsAbroad { get; set; }

        /// <summary>
        /// The internet service provider used for data reporting.
        /// </summary>
        [Input("isp")]
        public string? Isp { get; set; }

        /// <summary>
        /// Log level for data reporting(`1`: whitelist, `2`: normal, `4`: error, `8`: promise error, `16`: ajax request error, `32`: js resource load error, `64`: image resource load error, `128`: css resource load error, `256`: console.error, `512`: video resource load error, `1024`: request retcode error, `2048`: sdk self monitor error, `4096`: pv log, `8192`: event log).
        /// </summary>
        [Input("level")]
        public string? Level { get; set; }

        /// <summary>
        /// The network type used for data reporting.(`1`: Wifi, `2`: 2G, `3`: 3G, `4`: 4G, `5`: 5G, `6`: 6G, `100`: Unknown).
        /// </summary>
        [Input("netType")]
        public string? NetType { get; set; }

        /// <summary>
        /// The operating system used for data reporting.
        /// </summary>
        [Input("os")]
        public string? Os { get; set; }

        /// <summary>
        /// The platform where the data reporting takes place.(`1`: Android, `2`: IOS, `3`: Windows, `4`: Mac, `5`: Linux, `100`: Other).
        /// </summary>
        [Input("platform")]
        public string? Platform { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public int ProjectId { get; set; }

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public string? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time but is represented using a timestamp in seconds.
        /// </summary>
        [Input("startTime", required: true)]
        public int StartTime { get; set; }

        /// <summary>
        /// Query Data Type. `top`: CostType Sorting order by top, `nettype`: CostType Sorting order by nettype, `allcount`: CostType allcount, `day`: CostType group by day, `condition`: CostType Sorting in condition, `pagepv`: CostType Sorting order by pagepv, `area`: CostType query in area, `version`: CostType sort by version, `platform`: CostType sort by platform, `isp`: CostType sort by isp, `region`: region, `device`: CostType sort by device, `browser`: CostType sort by browser, `ext1`: CostType sort by ext1, `ext2`: CostType sort by ext2, `ext3`: CostType sort by ext3, `ret`: CostType sort by ret, `status`: CostType sort by status, `from`: CostType sort by from, `url`: CostType sort by url, `env`: CostType sort by env.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        /// <summary>
        /// The URL Key where the data reporting takes place.
        /// </summary>
        [Input("url")]
        public string? Url { get; set; }

        /// <summary>
        /// The SDK version used for data reporting.
        /// </summary>
        [Input("versionNum")]
        public string? VersionNum { get; set; }

        public GetCustomUrlArgs()
        {
        }
    }

    public sealed class GetCustomUrlInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The region where the data reporting takes place.
        /// </summary>
        [Input("area")]
        public Input<string>? Area { get; set; }

        /// <summary>
        /// The mobile phone brand used for data reporting.
        /// </summary>
        [Input("brand")]
        public Input<string>? Brand { get; set; }

        /// <summary>
        /// The browser type used for data reporting.
        /// </summary>
        [Input("browser")]
        public Input<string>? Browser { get; set; }

        /// <summary>
        /// The method used for calculating the elapsed time `50`: 50th percentile, `75`: 75th percentile, `90`: 90th percentile, `95`: 95th percentile, `99`: 99th percentile, `99.5`: 99.5th percentile, `avg`: Mean.
        /// </summary>
        [Input("costType")]
        public Input<string>? CostType { get; set; }

        /// <summary>
        /// The device used for data reporting.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// End time but is represented using a timestamp in seconds.
        /// </summary>
        [Input("endTime", required: true)]
        public Input<int> EndTime { get; set; } = null!;

        /// <summary>
        /// The browser engine used for data reporting.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The code environment where the data reporting takes place.(`production`: production env, `development`: development env, `gray`: gray env, `pre`: pre env, `daily`: daily env, `local`: local env, `others`: others env).
        /// </summary>
        [Input("env")]
        public Input<string>? Env { get; set; }

        /// <summary>
        /// First Expansion parameter.
        /// </summary>
        [Input("extFirst")]
        public Input<string>? ExtFirst { get; set; }

        /// <summary>
        /// Second Expansion parameter.
        /// </summary>
        [Input("extSecond")]
        public Input<string>? ExtSecond { get; set; }

        /// <summary>
        /// Third Expansion parameter.
        /// </summary>
        [Input("extThird")]
        public Input<string>? ExtThird { get; set; }

        /// <summary>
        /// The source page of the data reporting.
        /// </summary>
        [Input("from")]
        public Input<string>? From { get; set; }

        /// <summary>
        /// Whether it is non-China region.`1`: yes; `0`: no.
        /// </summary>
        [Input("isAbroad")]
        public Input<string>? IsAbroad { get; set; }

        /// <summary>
        /// The internet service provider used for data reporting.
        /// </summary>
        [Input("isp")]
        public Input<string>? Isp { get; set; }

        /// <summary>
        /// Log level for data reporting(`1`: whitelist, `2`: normal, `4`: error, `8`: promise error, `16`: ajax request error, `32`: js resource load error, `64`: image resource load error, `128`: css resource load error, `256`: console.error, `512`: video resource load error, `1024`: request retcode error, `2048`: sdk self monitor error, `4096`: pv log, `8192`: event log).
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// The network type used for data reporting.(`1`: Wifi, `2`: 2G, `3`: 3G, `4`: 4G, `5`: 5G, `6`: 6G, `100`: Unknown).
        /// </summary>
        [Input("netType")]
        public Input<string>? NetType { get; set; }

        /// <summary>
        /// The operating system used for data reporting.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// The platform where the data reporting takes place.(`1`: Android, `2`: IOS, `3`: Windows, `4`: Mac, `5`: Linux, `100`: Other).
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<int> ProjectId { get; set; } = null!;

        /// <summary>
        /// Used to save results.
        /// </summary>
        [Input("resultOutputFile")]
        public Input<string>? ResultOutputFile { get; set; }

        /// <summary>
        /// Start time but is represented using a timestamp in seconds.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<int> StartTime { get; set; } = null!;

        /// <summary>
        /// Query Data Type. `top`: CostType Sorting order by top, `nettype`: CostType Sorting order by nettype, `allcount`: CostType allcount, `day`: CostType group by day, `condition`: CostType Sorting in condition, `pagepv`: CostType Sorting order by pagepv, `area`: CostType query in area, `version`: CostType sort by version, `platform`: CostType sort by platform, `isp`: CostType sort by isp, `region`: region, `device`: CostType sort by device, `browser`: CostType sort by browser, `ext1`: CostType sort by ext1, `ext2`: CostType sort by ext2, `ext3`: CostType sort by ext3, `ret`: CostType sort by ret, `status`: CostType sort by status, `from`: CostType sort by from, `url`: CostType sort by url, `env`: CostType sort by env.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The URL Key where the data reporting takes place.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The SDK version used for data reporting.
        /// </summary>
        [Input("versionNum")]
        public Input<string>? VersionNum { get; set; }

        public GetCustomUrlInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomUrlResult
    {
        public readonly string? Area;
        public readonly string? Brand;
        public readonly string? Browser;
        public readonly string? CostType;
        public readonly string? Device;
        public readonly int EndTime;
        public readonly string? Engine;
        public readonly string? Env;
        public readonly string? ExtFirst;
        public readonly string? ExtSecond;
        public readonly string? ExtThird;
        public readonly string? From;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? IsAbroad;
        public readonly string? Isp;
        public readonly string? Level;
        public readonly string? NetType;
        public readonly string? Os;
        public readonly string? Platform;
        public readonly int ProjectId;
        /// <summary>
        /// Return value.
        /// </summary>
        public readonly string Result;
        public readonly string? ResultOutputFile;
        public readonly int StartTime;
        public readonly string Type;
        public readonly string? Url;
        public readonly string? VersionNum;

        [OutputConstructor]
        private GetCustomUrlResult(
            string? area,

            string? brand,

            string? browser,

            string? costType,

            string? device,

            int endTime,

            string? engine,

            string? env,

            string? extFirst,

            string? extSecond,

            string? extThird,

            string? from,

            string id,

            string? isAbroad,

            string? isp,

            string? level,

            string? netType,

            string? os,

            string? platform,

            int projectId,

            string result,

            string? resultOutputFile,

            int startTime,

            string type,

            string? url,

            string? versionNum)
        {
            Area = area;
            Brand = brand;
            Browser = browser;
            CostType = costType;
            Device = device;
            EndTime = endTime;
            Engine = engine;
            Env = env;
            ExtFirst = extFirst;
            ExtSecond = extSecond;
            ExtThird = extThird;
            From = from;
            Id = id;
            IsAbroad = isAbroad;
            Isp = isp;
            Level = level;
            NetType = netType;
            Os = os;
            Platform = platform;
            ProjectId = projectId;
            Result = result;
            ResultOutputFile = resultOutputFile;
            StartTime = startTime;
            Type = type;
            Url = url;
            VersionNum = versionNum;
        }
    }
}
