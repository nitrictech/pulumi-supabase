# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MigrationArgs', 'Migration']

@pulumi.input_type
class MigrationArgs:
    def __init__(__self__, *,
                 db_password: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 include_all: Optional[pulumi.Input[bool]] = None,
                 include_roles: Optional[pulumi.Input[bool]] = None,
                 include_seed: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Migration resource.
        """
        pulumi.set(__self__, "db_password", db_password)
        pulumi.set(__self__, "project_id", project_id)
        if include_all is not None:
            pulumi.set(__self__, "include_all", include_all)
        if include_roles is not None:
            pulumi.set(__self__, "include_roles", include_roles)
        if include_seed is not None:
            pulumi.set(__self__, "include_seed", include_seed)

    @property
    @pulumi.getter
    def db_password(self) -> pulumi.Input[str]:
        return pulumi.get(self, "db_password")

    @db_password.setter
    def db_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_password", value)

    @property
    @pulumi.getter
    def project_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def include_all(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_all")

    @include_all.setter
    def include_all(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_all", value)

    @property
    @pulumi.getter
    def include_roles(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_roles")

    @include_roles.setter
    def include_roles(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_roles", value)

    @property
    @pulumi.getter
    def include_seed(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_seed")

    @include_seed.setter
    def include_seed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_seed", value)


class Migration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_password: Optional[pulumi.Input[str]] = None,
                 include_all: Optional[pulumi.Input[bool]] = None,
                 include_roles: Optional[pulumi.Input[bool]] = None,
                 include_seed: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Migration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MigrationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Migration resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MigrationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MigrationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_password: Optional[pulumi.Input[str]] = None,
                 include_all: Optional[pulumi.Input[bool]] = None,
                 include_roles: Optional[pulumi.Input[bool]] = None,
                 include_seed: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MigrationArgs.__new__(MigrationArgs)

            if db_password is None and not opts.urn:
                raise TypeError("Missing required property 'db_password'")
            __props__.__dict__["db_password"] = db_password
            __props__.__dict__["include_all"] = include_all
            __props__.__dict__["include_roles"] = include_roles
            __props__.__dict__["include_seed"] = include_seed
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
        super(Migration, __self__).__init__(
            'supabase:v1:Migration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Migration':
        """
        Get an existing Migration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MigrationArgs.__new__(MigrationArgs)

        __props__.__dict__["db_password"] = None
        __props__.__dict__["include_all"] = None
        __props__.__dict__["include_roles"] = None
        __props__.__dict__["include_seed"] = None
        __props__.__dict__["project_id"] = None
        return Migration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def db_password(self) -> pulumi.Output[str]:
        return pulumi.get(self, "db_password")

    @property
    @pulumi.getter
    def include_all(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "include_all")

    @property
    @pulumi.getter
    def include_roles(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "include_roles")

    @property
    @pulumi.getter
    def include_seed(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "include_seed")

    @property
    @pulumi.getter
    def project_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_id")
