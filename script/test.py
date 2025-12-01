#!/usr/bin/env python3
import logging
import os
import sys
import typing

import recipe
from recipe import run, gomodules

logger = logging.getLogger("test")

class GOLintManager:
    def __init__(self,
                config:str = None,
                go_modules:typing.List[str] = None,
                lint_command:str = None):
        self.config = ".golangci.yaml" if len(config) == 0 else config
        self.go_modules = gomodules.list_go_workspace_module() if len(go_modules) == 0 else go_modules
        self.lint_command = "golangci-lint" if len(lint_command) == 0 else lint_command

    def lint(self) -> int:
        cp = run.command_shell(f"{self.lint_command} run --config {self.config} {GOLintManager.build_modules_arg(modules=self.go_modules)}")
        return cp.returncode

    def fmt(self) -> int:
        cp = run.command_shell(f"{self.lint_command} fmt --config {self.config} {GOLintManager.build_modules_arg(modules=self.go_modules)}")
        return cp.returncode

    @staticmethod
    def build_modules_arg(split:str = " ", modules:typing.List[str] = None)-> str:
        if len(modules) == 0:
            return ""
        return split.join([ x + "/..." for x in modules])

def main():
    config_file:str = os.environ.get("TEST_LINT_CONFIG",".golangci.yaml")
    modules:typing.List[str] = [m.strip() for m in os.environ.get("TEST_LINT_MODULE","").split(",") if m.strip()]
    command:str = os.environ.get("TEST_LINT_COMMAND","golangci-lint")

    lm = GOLintManager(config=config_file,go_modules=modules,lint_command=command)
    for do in sys.argv[1:]:
        logger.info(f"found action : {do}")
        match do.lower():
            case "lint":
                lm.lint()
            case "test":
                pass
            case "fmt":
                lm.fmt()
            case _:
                logger.warning(f"unknown action {do}")

if __name__ == "__main__":
    recipe.configure_logger()
    main()



