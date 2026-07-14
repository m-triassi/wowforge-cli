{ pkgs ? import <nixpkgs> {} }:

let
  version = "1.4.0";
in
pkgs.buildGoModule {
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
}
