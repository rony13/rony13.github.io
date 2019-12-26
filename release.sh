#!/bin/bash
set -e
set -x
hexo c | hexo g | hexo d
