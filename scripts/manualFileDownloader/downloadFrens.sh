#!/bin/bash

input="ls.txt"
while IFS= read -r line
do
  wget "https://ipfs.stargaze.zone/ipfs/bafybeiaip3vwwhhgerw6gcs66clj4onubkdadquqzzpzrftvuyhgnzojse/images/$line"
done < "$input"
