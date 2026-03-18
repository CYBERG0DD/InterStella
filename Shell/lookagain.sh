#!/bin/bash
# finds all finds ending with .sh, then remove the .sh extension and sort in decending order modification time
find . -type f -name "*.sh" -exec basename {} .sh \; | sort -r

