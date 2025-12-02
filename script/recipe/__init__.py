import logging
import sys

def configure_logger():

    log_handler = logging.StreamHandler(sys.stdout)

    log_formatter = logging.Formatter("[SCRIPT] [%(levelname)s] [%(asctime)s] (%(module)s:%(funcName)s) (%(name)s) %(message)s")

    log_handler.setLevel(logging.INFO)
    log_handler.setFormatter(log_formatter)

    root_logger = logging.getLogger()
    root_logger.addHandler(log_handler)
    root_logger.setLevel(logging.INFO)