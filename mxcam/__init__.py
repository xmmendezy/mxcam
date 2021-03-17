# -*- coding: utf-8 -*-

from .model import init_db
from .controller import controller_main

init_db()


def main():
    controller_main()
