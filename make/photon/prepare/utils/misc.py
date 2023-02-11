import os
import secrets
import string
import sys
from pathlib import Path
from functools import wraps
from g import DEFAULT_UID, DEFAULT_GID, host_root_dir

# To meet security requirement
# By default it will change file mode to 0600, and make the owner of the file to 10000:10000
def mark_file(path, mode=0o600, uid=DEFAULT_UID, gid=DEFAULT_GID):
    if mode > 0:
        os.chmod(path, mode)
    if uid > 0 and gid > 0:
        os.chown(path, uid, gid)

