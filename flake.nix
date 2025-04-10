{
  description = "Lura ~";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        go = pkgs.go_1_24;
        
        packageName = "lura";
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = packageName;
          version = "1.2.0";
          
          src = ./.;
          
          vendorHash = "sha256-E5RjjrBus7iECoz+09mUCMWNGuJFT8aeI+qyAHR5xKs=";
          
          proxyVendor = true;
          
          env = {
            CGO_ENABLED = "0";
          };
          
          ldflags = ["-s" "-w" "-extldflags '-static'"];
          
          # preBuild = ''
          #   go mod vendor
          # '';
        };
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
          #  gocode
            delve
            gotestsum
            git
          ];
          shellHook = ''
            export GOPATH=$HOME/go
            export PATH=$GOPATH/bin:$PATH
            export GO111MODULE=on
            
            export PROJECT_ROOT=$(pwd)
            
            echo "Go development environment ready!"
          '';
        };
      }
    );
}
