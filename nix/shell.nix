{
  inputs,
  lib,
  ...
}: {
  perSystem = {
    pkgs,
    config,
    ...
  }: {
    devShells.default = pkgs.mkShell {
      packages = with pkgs; [
        go
        golangci-lint
        protobuf
        protoc-gen-go
        protoc-gen-go-grpc
      ];
    };
  };
}
