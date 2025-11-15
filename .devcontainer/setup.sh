#!/bin/bash
set -ex

cat << 'EOF' >> ~/.bashrc

# Custom aliases
alias k='kubectl'
alias kd='kubectl describe'
alias kg='kubectl get'
EOF

# mage setup
cd magefiles
mage K3D:Create