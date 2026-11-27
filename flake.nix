{
  description = "b36date: print a date as YY plus base36 month and day";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "b36date";
          version = "0.1.0";
          src = self;
          vendorHash = null; # no dependencies

          meta = {
            description = "Print a date as YY plus base36 month and day";
            homepage = "https://github.com/NWoodsman/b36date";
            mainProgram = "b36date";
          };
        };

        apps.default = utils.lib.mkApp { drv = self.packages.${system}.default; };

        devShells.default = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [ go gopls ];
        };
      });
}
