#!/bin/bash
curl -s "https://acad.learn2earn.ng/assets/superhero/all.json" | jq -r '.[] | select(.id == 170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'
