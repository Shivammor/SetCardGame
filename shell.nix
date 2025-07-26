{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go

    # X11 + GLFW dependencies
    pkgs.xorg.libX11.dev
    pkgs.xorg.libXcursor.dev
    pkgs.xorg.libXrandr.dev
    pkgs.xorg.libXi.dev
    pkgs.xorg.libXinerama.dev
    pkgs.xorg.libXxf86vm.dev

    # OpenGL headers
    pkgs.libGL
    pkgs.libGLU
    pkgs.pkg-config
    # Mesa runtime libraries (GL + GLES)
    pkgs.mesa
  ];

  shellHook = ''
    echo "⏳ Setting up LD_LIBRARY_PATH for GL libraries..."

    export LD_LIBRARY_PATH=$(cat ./.cached-libgl-paths):$LD_LIBRARY_PATH
    echo "✅ LD_LIBRARY_PATH configured."
  '';
}

