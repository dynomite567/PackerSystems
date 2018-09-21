# ScoringEngine
The scoring engine from PackerSystems.

### Support this project:

I plan to keep this project completely free to make use of. If you wish to support me in some way, you can become a patron of mine on Patreon here:

https://www.patreon.com/GingerTechnology

## CyberPatriot
The CyberPatriot portion of this engine is definitely the more simple part. For the Linux side, the main "core" set of checks are currently still
some functions that called. I will be converting them to the new method soon.

The new method is having an array of a Check struct that lists the name, description, checking method, and expected output of a check.

The Linux version of the engine write a text file that gets imported into Wordpress as a new post.

The Windows version uses Pixel to have a window that displays the current scoring info.