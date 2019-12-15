#!/bin/bash

arrows=(

)

to_server=(
  :leftwards_arrow_with_hook:
  :arrow_right:
)

to_client=(
  :arrow_right_hook:
  :arrow_left:
)

r() {
  curl -sS 'https://microsoft.github.io/language-server-protocol/specifications/specification-3-14/' \
  | grep h4 \
  | grep 'alt="'$1'"' \
  | gsed 's!^.*href="#!!' \
  | gsed 's!".*!!' \
  | gsed 's!_!/!' \
  | sort \
  | uniq
}

for a in ${arrows[@]}; do
  echo $a
  r $a
  echo ""
done

echo "To Server"
echo ""
for a in ${to_server[@]}; do
  r $a
done | sort | uniq | sed 's/^/- [ ] /'
echo ""


echo "## To Client"
echo ""
for a in ${to_client[@]}; do
  r $a
done | sort | uniq | sed 's/^/- [ ] /'
echo ""

