import os
from pathlib import Path

## Const
DEFAULT_UID = 10000
DEFAULT_GID = 10000

PG_UID = 999
PG_GID = 999

REDIS_UID = 999
REDIS_GID = 999

host_root_dir = Path('/hostfs')

data_dir = Path('/data')

secret_dir = data_dir.joinpath('secret')
secret_key_dir = secret_dir.joinpath('keys')
trust_ca_dir = secret_dir.joinpath('keys', 'trust_ca')
internal_tls_dir = secret_dir.joinpath('tls')