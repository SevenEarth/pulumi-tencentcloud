// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Inputs
{

    public sealed class MediaVideoMontageTemplateAudioMixArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The media address of the audio track that needs to be mixed.
        /// </summary>
        [Input("audioSource", required: true)]
        public Input<string> AudioSource { get; set; } = null!;

        /// <summary>
        /// Mix Fade Configuration.
        /// </summary>
        [Input("effectConfig")]
        public Input<Inputs.MediaVideoMontageTemplateAudioMixEffectConfigArgs>? EffectConfig { get; set; }

        /// <summary>
        /// Mixing mode Repeat: background sound loop, Once: The background sound is played once.
        /// </summary>
        [Input("mixMode")]
        public Input<string>? MixMode { get; set; }

        /// <summary>
        /// Whether to replace the original audio of the Input media file with the mixed audio track media.
        /// </summary>
        [Input("replace")]
        public Input<string>? Replace { get; set; }

        public MediaVideoMontageTemplateAudioMixArgs()
        {
        }
        public static new MediaVideoMontageTemplateAudioMixArgs Empty => new MediaVideoMontageTemplateAudioMixArgs();
    }
}
