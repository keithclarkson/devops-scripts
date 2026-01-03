import logging
import os
import platform
import re
import subprocess
import sys

from packaging import version

def get_os_name():
    return platform.system()

def get_os_version():
    return platform.release()

def get_python_version():
    return platform.python_version()

def get_node_version():
    return subprocess.check_output(['node', '-v']).decode('utf-8').strip()

def get_npm_version():
    return subprocess.check_output(['npm', '-v']).decode('utf-8').strip()

def get_git_branch():
    try:
        return subprocess.check_output(['git', 'rev-parse', '--abbrev-ref', 'HEAD']).decode('utf-8').strip()
    except subprocess.CalledProcessError:
        return None

def get_git_commit_hash():
    try:
        return subprocess.check_output(['git', 'rev-parse', 'HEAD']).decode('utf-8').strip()
    except subprocess.CalledProcessError:
        return None

def get_node_modules():
    return subprocess.check_output(['npm', 'ls', '--parseable', '--depth', '0']).decode('utf-8').strip()

def get_npm_install_status():
    return subprocess.check_output(['npm', 'install', '--parseable']).decode('utf-8').strip()

def get_python_packages():
    return subprocess.check_output(['pip', 'freeze', '--local']).decode('utf-8').strip()

def compare_versions(v1, v2):
    return version.parse(v1) > version.parse(v2)