{ pkgs ? import <nixpkgs> {} }:

let
  version = "1.2.0";
in
pkgs.buildGoModule {
  pname = "wowforge-cli";
  inherit version;
  src = ./.;

  vendorHash = "sha256-C4KKe752MxqU/jOqKc3VXMiWalwan8hPP2RH1kgaDVI=";

  ldflags = [ "-X main.version=${version}" ];

  meta = with pkgs.lib; {
    description = "A CLI tool for managing World of Warcraft addons";
    homepage = "https://github.com/m-triassi/wowforge-cli";
    license = licenses.mit;
    mainProgram = "wowforge-cli";
  };
}
