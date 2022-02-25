with import <nixpkgs> {};

let
  postgresql = postgresql_10;

in 

stdenv.mkDerivation{
name = "maslow";

buildInputs =  with pkgs; [
  go
  goimports
  gopls
  fswatch
  vgo2nix
];


shellHook = ''
export GOPATH="/home/mike/dev/.go"
export GOBIN="/home/mike/dev/.go/bin"
export PATH="$PATH:/home/mike/dev/.go/bin"
source ~/dev/bin/bash-prompt.sh
 '';
}
