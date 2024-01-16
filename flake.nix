{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    templ.url = "github:a-h/templ";
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, flake-utils, nixpkgs, gitignore, templ }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        templ-pkg = templ.packages.${system}.templ;
      in {
        packages = {
          compare = pkgs.buildGo121Module {
            name = "compare";
            src = gitignore.lib.gitignoreSource ./.;
            vendorHash = "sha256-lrSHBSARK/p9QglkhJdg8gRg4tvKfz7KrtOK4wOsrjs=";

            nativeBuildInputs = [ pkgs.pkg-config ];
            buildInputs = [ pkgs.vips ];
            doCheck = false;

            preBuild = ''
              ${templ-pkg}/bin/templ generate
            '';
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            tailwindcss
            nodejs_20
            templ-pkg
            lsof
            sqlboiler
            vips
            pkg-config
          ];
        };
      });
}
