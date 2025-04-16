{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # Go
    go
    gopls        # Go language server

    # Lua
    lua
    luajit
    luarocks    # Package manager for Lua
    selene      # Lua linter
    lua-language-server

    # Etc
    git
    gnumake
    just
  ];

  shellHook = '' 
    export GOPATH="$HOME/go"
    export PATH="$GOPATH/bin:$PATH"
    
    export LUA_PATH="./?.lua;./?/init.lua;$HOME/.luarocks/share/lua/5.1/?.lua;$HOME/.luarocks/share/lua/5.1/?/init.lua"
    export LUA_CPATH="$HOME/.luarocks/lib/lua/5.1/?.so"
  '';
}
