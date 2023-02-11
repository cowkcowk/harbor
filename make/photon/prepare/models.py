import os
import logging
from pathlib import Path
from shutil import copytree, rmtree

from g import internal_tls_dir

class InternalTLS:

    harbor_certs_filename = {
        'harbor_internal_ca.crt',
    }

    def __init__(self, tls_enabled=False, verify_client_cert=False, tls_dir='', data_volume='', **kwargs):
        self.data_volume = data_volume
        self.verify_client_cert = verify_client_cert
        self.enabled = tls_enabled
        self.tls_dir = tls_dir
        if self.enabled:
            self.required_filenames = self.harbor_certs_filename
            if kwargs.get('with_notary'):
                self.required_filenames.update(self.notary_certs_filename)
            if kwargs.get('with_chartmuseum'):
                self.required_filenames.update(self.chart_museum_filename)
            if kwargs.get('with_trivy'):
                self.required_filenames.update(self.trivy_certs_filename)
            if not kwargs.get('external_database'):
                self.required_filenames.update(self.db_certs_filename)

    def __getattribute__(self, name: str):
        """
        Make the call like 'internal_tls.core_crt_path' possible
        """
        # only handle when enabled tls and name ends with 'path'
        if name.endswith('_path'):
            if not (self.enabled):
                return object.__getattribute__(self, name)

            name_parts = name.split('_')
            if len(name_parts) < 3:
                return object.__getattribute__(self, name)

            filename = '{}.{}'.format('_'.join(name_parts[:-2]), name_parts[-2])

            if filename in self.required_filenames:
                return os.path.join(self.data_volume, 'secret', 'tls', filename)

        return object.__getattribute__(self, name)

    def _check(self, filename: str):
        """
        Check cert and key files are correct
        """

        path = Path(os.path.join(internal_tls_dir, filename))

        if not path.exists:
            if filename == 'harbor_internal_ca.crt':
                return
            raise Exception('File {} not exist'.format(filename))

        if not path.is_file:
            raise Exception('invalid {}'.format(filename))

        # check key file permission
        if filename.endswith('.key') and not check_permission(path, mode=0o600):
            raise Exception('key file {} permission is not 600'.format(filename))

        # check certificate file
        if filename.endswith('.crt'):
            if not owner_can_read(path.stat().st_mode):
                # check owner can read cert file
                raise Exception('File {} should readable by owner'.format(filename))
            if not san_existed(path):
                # check SAN included
                if filename == 'harbor_internal_ca.crt':
                    return
                raise Exception('cert file {} should include SAN'.format(filename))

    def validate(self):
        if not self.enabled:
            # pass the validation if not enabled
            return

        if not internal_tls_dir.exists():
            raise Exception('Internal dir for tls {} not exist'.format(internal_tls_dir))

        for filename in self.required_filenames:
            self._check(filename)

    