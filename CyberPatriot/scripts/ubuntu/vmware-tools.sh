#!/bin/bash

set -e
set -x

if [ "$PACKER_BUILDER_TYPE" != "vmware-iso" ]; then
  exit 0
fi

apt update

apt-get -y install perl make linux-headers-$(uname -r) xserver-xorg

apt-get -y install open-vm-tools open-vm-tools-desktop

exit 0