import logging
import subprocess
import sys
import typing

logger = logging.getLogger("recipe.run")

def command(cmdline: typing.List[str], capture_output: bool = False, **kwargs):
    logger.info(f"running command: {' '.join(cmdline)}", **kwargs)
    return subprocess.run(
        cmdline,
        stdout=subprocess.PIPE if capture_output else sys.stdout,
        stdin=sys.stdin,
        stderr=subprocess.PIPE if capture_output else sys.stderr,
        shell=False,
        capture_output=False,  # We handle it manually
        check=True,  # Raise error on non-zero exit
        **kwargs
    )

def command_shell(cmdline: str, capture_output: bool = False, **kwargs):
    logger.info(f"running shell command: {cmdline}", **kwargs)
    return subprocess.run(
        cmdline,
        stdout=subprocess.PIPE if capture_output else sys.stdout,
        stdin=sys.stdin,
        stderr=subprocess.PIPE if capture_output else sys.stderr,
        shell=True,
        capture_output=False,  # Handled manually
        check=True,  # Raise error on non-zero exit
        **kwargs
    )