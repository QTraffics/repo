import json
import logging
import subprocess
import typing

from . import run

logger = logging.getLogger("recipe.gomodules")

def list_go_workspace_module(command:str="go") -> typing.List[str]:
    try:
        result = run.command_shell(f"{command} work edit -json", capture_output=True)
        load_json = json.loads(result.stdout.decode('utf-8'))
        ans = [str(ele["DiskPath"]) for ele in load_json.get("Use", [])]
        return ans
    except (subprocess.CalledProcessError, json.JSONDecodeError) as e:
        logger.error(f"failed to list go workspace modules: {str(e)}")
        return []