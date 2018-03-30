# GoDupes (Kata)
Super Fast Go Duplicates Finder
Find Duplicate Files ------ as fast as possible!

<img src="https://raw.githubusercontent.com/AndreiD/GoDupes/master/assets/screenshot1.JPG" alt="godupes" />

## Why:
To exercise Go

## Synopsis:
The bottleneck in this case is always the speed of the HDD and the hashing algorithm.
Since we cannot modify the HDD speed, let's choose a fast hashing algo

you can see the tests of various algos in hashhelper_test.go

#### The chosen algo is <strong>xxHash</strong>
xxHash is an Extremely fast Hash algorithm, running at RAM speed limits.

## Benchmarks:

an old i5 3470, Western Digital Black 7200rpm HDD
4.24 seconds on 151GB (24082 files)

## Todo:

- optimize it even more
- add a progress bar

## License:
<pre>
Copyright (C) 2018 me

Everyone is permitted to copy and distribute verbatim or modified
copies of this license document, and changing it is allowed as long
as the name is changed.

           DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
  TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

 0. You just DO WHAT THE FUCK YOU WANT TO.

</pre>