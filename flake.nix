{
  description = "wowforge-cli - A CLI tool for managing World of Warcraft addons";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages = {
          wowforge-cli = let
            version = "1.4.1";
          in pkgs.buildGoModule {
            pname = "wowforge-cli";
            inherit version;
            src = ./.;

            vendorHash = "sha256-Z6d+YfqGq+zz8DoxZ0kkepn4GH4O0DB9qNJaaYCWjYI=";

            ldflags = [ "-X main.version=${version}" ];

            meta = with pkgs.lib; {
              description = "A CLI tool for managing World of Warcraft addons";
              homepage = "https://github.com/m-triassi/wowforge-cli";
              license = licenses.mit;
              mainProgram = "wowforge-cli";
            };
          };
          default = self.packages.${system}.wowforge-cli;
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/wowforge-cli";
        };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [ go gopls ];
        };
      }
    );
}
