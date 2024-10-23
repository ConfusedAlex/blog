---
title: "How I broke sudo on my Dads Mac and how I fixed it"
date: 2024-01-19
slug: "how-i-broke-sudo"
---

# Prologue

Currently, I'm in the process of migrating my Steam Deck to NixOS. Today finally arrived my 1 TB NVMe, so I rushed to installing it and reinstalling NixOS as I already had a working configuration. The annoying part is, that, when using [Jovian-NixOS](https://github.com/Jovian-Experiments/Jovian-NixOS) you need to compile the kernel. As this takes long and my father owns a M1 Mac, I thought about building my NixOS configuration on the mac and copying it to the Steam Deck.

# The Problem

I then installed Nix on the Mac and got everything working, but as I wanted to build my NixOS Configuration on the Mac, the build failed. It failed, because I couldn't enter the sudo password. After much browsing the web, I wanted to make sudo passwordless. So I open `visudo` and changed

```bash
%admin ALL=(ALL) ALL
```

to

```bash
%admin ALL=(ALL) :NOPASSWD ALL
```

Do you see the mistake? It's the wrong place column. But I was in a rush, saved the file, wanted to test out my new fancy passwordless sudo by running `sudo echo test` and got greeted by a `alex is not in suderos file`.

# The Journey begins

## Try 1

I tried out the [here](https://superuser.com/a/1368278) described “Finder” solution but failed. I just couldn't change the permission back. An hour later, I moved on to recovery mode.

## Try 2 — Recovery Mode, my savior

So, I finally booted into recovery mode, and logged in. Then mounted the Data disk via the disk utility program.  
Now I could open a terminal, navigate to `/foo/bar/etc` and repair the file permission while also fixing the typo.

# What I've learned

1. If visudo asks you if you really want to save the File, don't blindly type yes!
2. root is your friend

I would love to hear about your funny tech fails. Just send me an email.


