# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('supabase')


class _ExportableConfig(types.ModuleType):
    @property
    def token(self) -> str:
        """
        Supbase Personal Access Token for account
        """
        return __config__.get('token') or (_utilities.get_env('SUPABASE_ACCESS_TOKEN') or '')

    @property
    def version(self) -> Optional[str]:
        return __config__.get('version')

