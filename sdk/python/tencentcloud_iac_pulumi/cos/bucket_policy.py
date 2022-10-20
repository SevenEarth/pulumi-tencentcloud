# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BucketPolicyArgs', 'BucketPolicy']

@pulumi.input_type
class BucketPolicyArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 policy: pulumi.Input[str]):
        """
        The set of arguments for constructing a BucketPolicy resource.
        :param pulumi.Input[str] bucket: The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        :param pulumi.Input[str] policy: The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)


@pulumi.input_type
class _BucketPolicyState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BucketPolicy resources.
        :param pulumi.Input[str] bucket: The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        :param pulumi.Input[str] policy: The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


class BucketPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a COS resource to create a COS bucket policy and set its attributes.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cos_policy = tencentcloud.cos.BucketPolicy("cosPolicy",
            bucket="mycos-1258798060",
            policy=\"\"\"{
          "version": "2.0",
          "Statement": [
            {
              "Principal": {
                "qcs": [
                  "qcs::cam::uin/<your-account-id>:uin/<your-account-id>"
                ]
              },
              "Action": [
                "name/cos:DeleteBucket",
                "name/cos:PutBucketACL"
              ],
              "Effect": "allow",
              "Resource": [
                "qcs::cos:<bucket region>:uid/<your-account-id>:<bucket name>/*"
              ]
            }
          ]
        }

        \"\"\")
        ```

        ## Import

        COS bucket policy can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:Cos/bucketPolicy:BucketPolicy bucket bucket-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        :param pulumi.Input[str] policy: The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a COS resource to create a COS bucket policy and set its attributes.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        cos_policy = tencentcloud.cos.BucketPolicy("cosPolicy",
            bucket="mycos-1258798060",
            policy=\"\"\"{
          "version": "2.0",
          "Statement": [
            {
              "Principal": {
                "qcs": [
                  "qcs::cam::uin/<your-account-id>:uin/<your-account-id>"
                ]
              },
              "Action": [
                "name/cos:DeleteBucket",
                "name/cos:PutBucketACL"
              ],
              "Effect": "allow",
              "Resource": [
                "qcs::cos:<bucket region>:uid/<your-account-id>:<bucket name>/*"
              ]
            }
          ]
        }

        \"\"\")
        ```

        ## Import

        COS bucket policy can be imported, e.g.

        ```sh
         $ pulumi import tencentcloud:Cos/bucketPolicy:BucketPolicy bucket bucket-name
        ```

        :param str resource_name: The name of the resource.
        :param BucketPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketPolicyArgs.__new__(BucketPolicyArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
        super(BucketPolicy, __self__).__init__(
            'tencentcloud:Cos/bucketPolicy:BucketPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'BucketPolicy':
        """
        Get an existing BucketPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        :param pulumi.Input[str] policy: The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketPolicyState.__new__(_BucketPolicyState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["policy"] = policy
        return BucketPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycos-1258798060`.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).
        """
        return pulumi.get(self, "policy")

