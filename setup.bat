@echo off
chcp 65001 > nul
REM =====================================
REM AviUtlExEdit2PluginManager セットアップ
REM =====================================

echo [INFO] プロジェクトセットアップを開始します...
echo.

REM 基本ディレクトリ構造の作成
echo [INFO] ディレクトリ構造を作成中...

if not exist "src" (
    mkdir "src"
    echo [INFO] src ディレクトリを作成しました
)

if not exist "include" (
    mkdir "include"
    echo [INFO] include ディレクトリを作成しました
)

if not exist "lib" (
    mkdir "lib"
    echo [INFO] lib ディレクトリを作成しました
)

if not exist "bin" (
    mkdir "bin"
    echo [INFO] bin ディレクトリを作成しました
)

if not exist "bin\debug" (
    mkdir "bin\debug"
    echo [INFO] bin\debug ディレクトリを作成しました
)

if not exist "bin\release" (
    mkdir "bin\release"
    echo [INFO] bin\release ディレクトリを作成しました
)

if not exist "dist" (
    mkdir "dist"
    echo [INFO] dist ディレクトリを作成しました
)

if not exist "docs" (
    mkdir "docs"
    echo [INFO] docs ディレクトリを作成しました
)

if not exist "config" (
    mkdir "config"
    echo [INFO] config ディレクトリを作成しました
)

if not exist "plugins" (
    mkdir "plugins"
    echo [INFO] plugins ディレクトリを作成しました
)

if not exist "tests" (
    mkdir "tests"
    echo [INFO] tests ディレクトリを作成しました
)

echo.

REM .gitignore ファイルの作成
if not exist ".gitignore" (
    echo [INFO] .gitignore ファイルを作成中...
    (
        echo # ビルド成果物
        echo bin/
        echo dist/
        echo *.exe
        echo *.dll
        echo *.obj
        echo *.o
        echo *.lib
        echo *.a
        echo.
        echo # Visual Studio
        echo .vs/
        echo *.vcxproj.user
        echo *.vcxproj.filters
        echo *.sdf
        echo *.opensdf
        echo *.suo
        echo *.user
        echo.
        echo # デバッグファイル
        echo *.pdb
        echo *.ilk
        echo *.log
        echo.
        echo # 一時ファイル
        echo *.tmp
        echo *.temp
        echo *~
        echo.
        echo # OS固有
        echo Thumbs.db
        echo .DS_Store
        echo.
        echo # IDE固有
        echo *.swp
        echo *.swo
        echo.
        echo # パッケージファイル
        echo *.zip
        echo *.tar.gz
        echo *.rar
    ) > ".gitignore"
    echo [INFO] .gitignore ファイルを作成しました
)

REM 基本的な設定ファイルの作成
if not exist "config\settings.ini" (
    echo [INFO] 設定ファイルを作成中...
    (
        echo [General]
        echo AppName=AviUtlExEdit2PluginManager
        echo Version=0.2.0
        echo Author=seless-yuu
        echo.
        echo [Build]
        echo OutputDir=bin
        echo DebugMode=true
        echo.
        echo [Plugins]
        echo PluginDir=plugins
        echo AutoLoad=true
    ) > "config\settings.ini"
    echo [INFO] 設定ファイルを作成しました
)

REM サンプルソースファイルの作成（C++想定）
if not exist "src\main.cpp" (
    echo [INFO] サンプルソースファイルを作成中...
    (
        echo // AviUtlExEdit2PluginManager
        echo // メインエントリーポイント
        echo.
        echo #include ^<iostream^>
        echo #include ^<string^>
        echo.
        echo int main^(int argc, char* argv[]^)
        echo {
        echo     std::cout ^<^< "AviUtlExEdit2PluginManager v0.2" ^<^< std::endl;
        echo     std::cout ^<^< "プラグインマネージャーを開始します..." ^<^< std::endl;
        echo.
        echo     // TODO: プラグインマネージャーのメイン処理を実装
        echo.
        echo     return 0;
        echo }
    ) > "src\main.cpp"
    echo [INFO] サンプルmain.cppを作成しました
)

REM CMakeLists.txtの作成
if not exist "CMakeLists.txt" (
    echo [INFO] CMakeLists.txtを作成中...
    (
        echo cmake_minimum_required^(VERSION 3.16^)
        echo project^(AviUtlExEdit2PluginManager^)
        echo.
        echo set^(CMAKE_CXX_STANDARD 17^)
        echo set^(CMAKE_CXX_STANDARD_REQUIRED ON^)
        echo.
        echo # ソースファイル
        echo file^(GLOB_RECURSE SOURCES "src/*.cpp" "src/*.c"^)
        echo file^(GLOB_RECURSE HEADERS "include/*.h" "include/*.hpp"^)
        echo.
        echo # インクルードディレクトリ
        echo include_directories^(include^)
        echo.
        echo # 実行ファイル
        echo add_executable^(${PROJECT_NAME} ${SOURCES} ${HEADERS}^)
        echo.
        echo # デバッグビルド設定
        echo if^(CMAKE_BUILD_TYPE STREQUAL "Debug"^)
        echo     target_compile_definitions^(${PROJECT_NAME} PRIVATE DEBUG=1^)
        echo endif^(^)
        echo.
        echo # リリースビルド設定
        echo if^(CMAKE_BUILD_TYPE STREQUAL "Release"^)
        echo     target_compile_definitions^(${PROJECT_NAME} PRIVATE NDEBUG=1^)
        echo     set_target_properties^(${PROJECT_NAME} PROPERTIES
        echo         LINK_FLAGS "/SUBSYSTEM:WINDOWS"^)
        echo endif^(^)
    ) > "CMakeLists.txt"
    echo [INFO] CMakeLists.txtを作成しました
)

REM Makefileの作成
if not exist "Makefile" (
    echo [INFO] Makefileを作成中...
    (
        echo # AviUtlExEdit2PluginManager Makefile
        echo.
        echo CXX = g++
        echo CXXFLAGS = -std=c++17 -Iinclude
        echo DEBUGFLAGS = -g -DDEBUG=1
        echo RELEASEFLAGS = -O2 -DNDEBUG=1
        echo.
        echo SRCDIR = src
        echo BINDIR = bin
        echo SOURCES = ^$^(wildcard ^$^(SRCDIR^)/*.cpp^)
        echo.
        echo TARGET = AviUtlExEdit2PluginManager
        echo.
        echo .PHONY: all debug release clean
        echo.
        echo all: debug
        echo.
        echo debug: ^$^(BINDIR^)/debug/^$^(TARGET^).exe
        echo.
        echo release: ^$^(BINDIR^)/release/^$^(TARGET^).exe
        echo.
        echo ^$^(BINDIR^)/debug/^$^(TARGET^).exe: ^$^(SOURCES^)
        echo 	@mkdir -p ^$^(BINDIR^)/debug
        echo 	^$^(CXX^) ^$^(CXXFLAGS^) ^$^(DEBUGFLAGS^) -o ^$@ ^$^
        echo.
        echo ^$^(BINDIR^)/release/^$^(TARGET^).exe: ^$^(SOURCES^)
        echo 	@mkdir -p ^$^(BINDIR^)/release
        echo 	^$^(CXX^) ^$^(CXXFLAGS^) ^$^(RELEASEFLAGS^) -o ^$@ ^$^
        echo.
        echo clean:
        echo 	rm -rf ^$^(BINDIR^)
        echo 	rm -rf dist
    ) > "Makefile"
    echo [INFO] Makefileを作成しました
)

echo.
echo [SUCCESS] プロジェクトセットアップが完了しました！
echo.
echo 作成されたディレクトリとファイル:
echo ├── src/            (ソースコード)
echo ├── include/        (ヘッダーファイル)
echo ├── lib/            (ライブラリ)
echo ├── bin/            (ビルド成果物)
echo ├── dist/           (配布ファイル)
echo ├── docs/           (ドキュメント)
echo ├── config/         (設定ファイル)
echo ├── plugins/        (プラグイン)
echo ├── tests/          (テストファイル)
echo ├── .gitignore      (Git除外設定)
echo ├── CMakeLists.txt  (CMakeビルド設定)
echo └── Makefile        (Makeビルド設定)
echo.
echo 次のステップ:
echo 1. dev-build.bat で開発ビルドを実行
echo 2. build.bat で本番ビルドを実行
echo 3. src/main.cpp を編集してコードを実装
echo.
pause