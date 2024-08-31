#!/bin/sh
# scripts/completions.sh
set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
	go run main.go completion "$sh" >"completions/bitbucket.$sh"
done


rm -rf manpages
mkdir manpages
gzip -c -9 man/bitbucket.1 >manpages/bitbucket.1.gz
