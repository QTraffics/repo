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
                config:str = ".golangci.yaml",
                go_modules:typing.List[str] = None,
                lint_command:str = "golangci-lint",
                 go_command:str = "go"):
        self.config = config
        self.go_modules = gomodules.list_go_workspace_module(go_command) if len(go_modules) == 0 else go_modules
        self.lint_command = lint_command
        self.go_command= go_command

    def lint(self) -> int:
        cp = run.command_shell(f"{self.lint_command} run --config {self.config} {GOLintManager.build_modules_arg(modules=self.go_modules)}")
        return cp.returncode

    def fmt(self) -> int:
        cp = run.command_shell(f"{self.lint_command} fmt --config {self.config} {GOLintManager.build_modules_arg(modules=self.go_modules)}")
        return cp.returncode
    def test(self)-> int:
        cp = run.command_shell(f'{self.go_command} test {GOLintManager.build_modules_arg(modules=self.go_modules)}')
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
        logger.info(f"found action: {do}")
        code=0
        match do.lower():
            case "lint":
                code = lm.lint()
            case "test":
                code = lm.test()
            case "fmt":
                code = lm.fmt()
            case _:
                logger.warning(f"unknown action {do}")
        if code != 0:
            logger.warning(f'subprocess quit with code {code}')
            return

if __name__ == "__main__":
    recipe.configure_logger()
    main()



