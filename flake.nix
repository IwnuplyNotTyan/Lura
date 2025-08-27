{
  description = "Lura (not) simple game";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages.default = pkgs.buildGoModule rec {
          pname = "lura";
          version = "1.3";

          src = pkgs.fetchFromGitHub {
            owner = "iwnuplynottyan";          
            repo = "lura";          
            rev = version;                    
            sha256 = "sha256-QsmqE3PfRdKWbFLFXrqvzD0ogfTNK0A/gDBOkZVY2AU=";
          };

          vendorHash = "sha256-Ua63ON+SjMiRDAOicSRS047Kd0JjGqjgkQuUs0JNEgU=";
          # vendorHash = null;

          meta = with pkgs.lib; {
            description = "Lura (not) simple game";
            homepage = "https://github.com/iwnuplynottyan/lura";
            license = licenses.mit;
            maintainers = [ "iwnuplynottyan" ];
          };
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/lura";
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            gopls
          ];
        };
      });
}
