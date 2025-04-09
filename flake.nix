
{
  description = "Lura";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        packages.default = pkgs.buildGoModule {
          pname = "lura";
          version = "1.2.0";

          src = ./.;

          vendorHash = null;

          meta = {
            description = "Lura - simple cli roguelike game";
            license = pkgs.lib.licenses.mit;
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls       
            #golangci-lint
          ];

          shellHook = ''
          '';
        };
      }
    );
}

