# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetInstanceAuditListResult',
]

@pulumi.output_type
class GetInstanceAuditListResult(dict):
    def __init__(__self__, *,
                 audit_switch: bool,
                 cos_bucket: str,
                 id: str,
                 log_file_prefix: str,
                 name: str):
        """
        :param bool audit_switch: Indicate whether audit start logging or not.
        :param str cos_bucket: Cos bucket name where audit save logs.
        :param str id: ID of the audit.
        :param str log_file_prefix: Prefix of the log file of the audit.
        :param str name: Name of the audits.
        """
        pulumi.set(__self__, "audit_switch", audit_switch)
        pulumi.set(__self__, "cos_bucket", cos_bucket)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "log_file_prefix", log_file_prefix)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="auditSwitch")
    def audit_switch(self) -> bool:
        """
        Indicate whether audit start logging or not.
        """
        return pulumi.get(self, "audit_switch")

    @property
    @pulumi.getter(name="cosBucket")
    def cos_bucket(self) -> str:
        """
        Cos bucket name where audit save logs.
        """
        return pulumi.get(self, "cos_bucket")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the audit.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logFilePrefix")
    def log_file_prefix(self) -> str:
        """
        Prefix of the log file of the audit.
        """
        return pulumi.get(self, "log_file_prefix")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the audits.
        """
        return pulumi.get(self, "name")


