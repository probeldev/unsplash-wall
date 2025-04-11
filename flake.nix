{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
      unsplashwall-package = pkgs.callPackage ./package.nix {};
    in {
      packages = rec {
        unsplashwall = unsplashwall-package;
        default = unsplashwall;
      };

      apps = rec {
        unsplashwall = flake-utils.lib.mkApp {
          drv = self.packages.${system}.unsplashwall;
        };
        default = unsplashwall;
      };

      devShells.default = pkgs.mkShell {
        packages = (with pkgs; [
          go
        ]);
      };
    });
}
