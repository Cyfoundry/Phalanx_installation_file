# from distutils.core import setup
# from setuptools import find_packages
from setuptools import setup, find_packages
setup(name='cusnmap',
      version='0.1',
      packages=find_packages('cusnmap'),
      package_dir={'': 'cusnmap'},
      include_package_data=False,
      package_data={'data': []},
      description='Hook Nmap to Scan',
      long_description='',
      author='jimchang,nilo',
      author_email='jimchang@cyfoundry.com,nilo@cyfoundry.com',
      url='https://cyfoundry.local.com',
      license='MIT',
      install_requires=['simplejson', 'asyncio'],
      )
