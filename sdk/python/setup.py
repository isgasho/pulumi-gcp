# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call

class InstallPluginCommand(install):
    def run(self):
        install.run(self)
        check_call(['pulumi', 'plugin', 'install', 'resource', 'gcp', '${PLUGIN_VERSION}'])

def readme():
    with open('README.rst') as f:
        return f.read()

setup(name='pulumi_gcp',
      version='${VERSION}',
      description='A Pulumi package for creating and managing Google Cloud Platform resources.',
      long_description=readme(),
      cmdclass={
          'install': InstallPluginCommand,
      },
      keywords='pulumi gcp',
      url='https://pulumi.io',
      project_urls={
          'Repository': 'https://github.com/pulumi/pulumi-gcp'
      },
      license='Apache 2.0',
      packages=find_packages(),
      install_requires=[
          'pulumi>=0.14.2,<0.15.0'
      ],
      zip_safe=False)
