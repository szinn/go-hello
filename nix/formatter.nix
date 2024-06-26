{inputs, ...}: {
  imports = [
    inputs.treefmt-nix.flakeModule
  ];
  perSystem = {
    config,
    pkgs,
    lib,
    ...
  }: {
    treefmt.config = {
      inherit (config.flake-root) projectRootFile;
      package = pkgs.treefmt;

      programs = {
        alejandra.enable = true;
        gofumpt.enable = true;
        prettier.enable = true;
      };

      settings.formatter.prettier.options = ["--tab-width" "4"];
    };

    formatter = config.treefmt.build.wrapper;
  };
}
