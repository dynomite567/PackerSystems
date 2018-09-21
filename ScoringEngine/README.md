# ScoringEngine
The scoring engine from PackerSystems.

I've decided to separate out the sub-projects of PackerSystems into their own repos for easier management.
This is the scoring engine portion and I will be fleshing out documentation on it more in the coming days and weeks.

### Support this project:

I plan to keep this project completely free to make use of. If you wish to support me in some way, you can become a patron of mine on Patreon here:

https://www.patreon.com/GingerTechnology

## CyberPatriot
The CyberPatriot portion of this engine is definitely the more simple part. For the Linux side, the main "core" set of checks are currently still
some functions that called. I will be converting them to the new method soon.

The new method is having an array of a Check struct that lists the name, description, checking method, and expected output of a check.

The Linux version of the engine write a text file that gets imported into Wordpress as a new post.

The Windows version uses Pixel to have a window that displays the current scoring info.

## CCDC
The CCDC portion is much more complex. It will be scoring remote boxes, which means things like SSHing into the boxes, curling webpages, etc. which
makes it a bit harder to logic. However, since the box it will be running from shouldn't be touched once setup, I can have a JSON file of the checks,
which is what I'm doing. In addition to that, the way that I'm making the current status visible to the users will be a dotnet webapp that pulls data
from the JSON file.

Every three minutes the scoring service will iterate through the check.json file in /etc/gingertechengine/ running each one. At the end it will dump the
result into a new file called current.json, which is what gets looked at by the dotnet site. 