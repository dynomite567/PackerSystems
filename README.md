# PackerSystems

**Forked from GitHub repo BaileyGingerTechnology/PackerSystems**

The goal of this project is to eventually provide practice environments for things such as:

1. CyberPatriot for high school students
2. CCDC for college students

## Build instructions:

Install packer (packer-io on Arch Linux) and a hypervisor (Virtualbox, KVM, or VMware. I might support Parallels and Hyper-V in the future, but not yet).

Download the Windows 10 and Windows Server 2012r2 ISOs. The Ubuntu one will be gotten later.

Clone the repo and cd into it

cd into the directory of the environment you are building out

packer build -only=(virtualbox-iso, vmware-iso, or qemu) ./(machine).json

And then wait. A really long time. Repeat the last step for each image you want. If you have the resources, you can do several builds at once. If you do not specify -only, it will build all the platforms.

## Progress so far:

### CyberPatriot:

Ubuntu and Windows 10 are both functional and build the base of what I want. But only in Virtualbox for Windows 10, so now my goal is to get it working in VMware and KVM.

For Ubuntu:

    - https://blog.gingertechnology.net/2018/05/28/ubuntu-1-0-a-cyberpatriot-practice-image/

For Windows 10:

    - Installs updates that are the latest at the time of building

    - Runs most of the scripts from https://github.com/W4RH4WK/Debloat-Windows-10

        - This causes Windows Update to break and uber disables Windows Defender

    - Shares the whole C:\ drive to all the people

    - Puts the eicar test file (file that gets flagged as malware but actually does nothing) in SysWOW64

    - Makes a few users

    - Disables password policy

    - Sets proxy.google.com to be a web proxy

    - Routes search engines to localhost in the hosts file (except Google, which goes to ask.com's IP)

    - Disables Windows firewall

    - Installs scoop package manager for future use

For Windows Server 2012r2:

    - Does not install updates

    - Shares the whole C:\ drive to all the people

    - Puts the eicar test file (file that gets flagged as malware but actually does nothing) in system32

    - Disables password policy

    - Sets proxy.google.com to be a web proxy

    - Routes search engines to localhost in the hosts file (except Google, which goes to ask.com's IP)

    - Disables Windows firewall

    - Installs scoop package manager for future use

    - I need to research some AD PowerShell tricks to better mess up this box

### Scoring Engine:

My goal is to have a scoring engine for the CyberPatriot machines, and then another that checks whether services are running on the CCDC environment. I'll probably host the CCDC one on the attacking system that I plan to make.

The Linux CyberPatriot one is almost done. The Windows one is also functional, I just need to make the checks for it.

### Support this project:

I plan to keep this project completely free to make use of. If you wish to support me in some way, you can become a patron of mine on Patreon here:

https://www.patreon.com/GingerTechnology