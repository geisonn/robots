# Robots - Assertive Hidden Directory Discovery Tool

Robots is a high-performance tool designed to discover hidden directories on websites by exploring information exposed in the applications' robots.txt file. Unlike traditional fuzzing tools, which generate random guesses of directory names, Robots is assertive, finding directories that are clearly exposed but commonly overlooked, increasing the efficiency of security testing.

## Overview

Robots was developed for pentesters, bug bounty hunters and security researchers who need to perform fast and accurate searches for hidden directories. Instead of performing random name searches as in fuzzing, Robots focuses directly on the information exposed in robots.txt, where many website administrators list directories and files that should not be indexed, but which, in practice, are publicly accessible. The tool is ideal for helping to better map an application's directory structure, discovering exposed directories with greater precision and speed.

## Performance

Robots is optimized for speed, using fast HTTP request algorithms to search for valid and hidden directories. The tool focuses on being assertive, avoiding the use of random fuzzing and making a more efficient exploration of websites.

In addition, Robots allows for large-scale scans, being able to process a list of URLs quickly, which makes it ideal for mass security testing.

## How to Use

### Help
```
$ robots -h
Usage:
  -h, -help           Show help
  -hf, -hidefails     Hide failed attempts
  -l, -list <file>    List of URLs to process
  -u, -url <URL>      Target URL to test

List of URLs:
$ robots -l urls.txt -hf | sort -u | tee directories.txt

Single URL:
$ robots -u <SITE>
```

## Installation

```go install github.com/geisonn/robots@latest```
