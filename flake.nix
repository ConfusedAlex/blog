{
  description = "A basic flake with a shell";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gopls
            gotools
            # golangci-lint

            templ
            just
            self.packages.${system}.confusedblog
          ];
        };
        packages = {
          confusedblog = pkgs.buildGoModule {
            name = "confusedblog";
            src = ./.;
            vendorHash = "sha256-thS9n35sbBhwlpbTkDicmMR2kc6wDMu7ALP4crJFWNA=";
            # subPackages = [ "cmd" ];
          };
        };
        defaultPackage = self.packages.${system}.confusedblog;
      }
    );
}
