@echo off
chcp 65001 > nul
REM =====================================
REM AviUtlExEdit2PluginManager クリーンアップ
REM =====================================

echo [INFO] クリーンアップを開始します...
echo.

echo 以下の項目をクリーンアップしますか？
echo [1] ビルド成果物のみ (bin/, dist/)
echo [2] 全ての生成ファイル (ビルド成果物 + 一時ファイル)
echo [3] キャンセル
echo.
set /p choice="選択してください (1-3): "

if "%choice%"=="3" (
    echo [INFO] クリーンアップをキャンセルしました。
    goto :end
)

if "%choice%"=="1" goto :clean_build
if "%choice%"=="2" goto :clean_all

echo [ERROR] 無効な選択です。
goto :end

:clean_build
echo [INFO] ビルド成果物をクリーンアップ中...

REM ビルドディレクトリの削除
if exist "bin" (
    echo [INFO] bin/ ディレクトリを削除中...
    rmdir /s /q "bin"
    if not exist "bin" echo [SUCCESS] bin/ ディレクトリを削除しました
)

if exist "dist" (
    echo [INFO] dist/ ディレクトリを削除中...
    rmdir /s /q "dist"
    if not exist "dist" echo [SUCCESS] dist/ ディレクトリを削除しました
)

REM 実行ファイルの削除
if exist "*.exe" (
    echo [INFO] 実行ファイルを削除中...
    del /q "*.exe"
    echo [SUCCESS] 実行ファイルを削除しました
)

if exist "*.dll" (
    echo [INFO] DLLファイルを削除中...
    del /q "*.dll"
    echo [SUCCESS] DLLファイルを削除しました
)

REM オブジェクトファイルの削除
if exist "*.obj" (
    del /q "*.obj"
    echo [SUCCESS] オブジェクトファイルを削除しました
)

if exist "*.o" (
    del /q "*.o"
    echo [SUCCESS] オブジェクトファイル(.o)を削除しました
)

REM ライブラリファイルの削除
if exist "*.lib" (
    del /q "*.lib"
    echo [SUCCESS] ライブラリファイルを削除しました
)

if exist "*.a" (
    del /q "*.a"
    echo [SUCCESS] ライブラリファイル(.a)を削除しました
)

echo [SUCCESS] ビルド成果物のクリーンアップが完了しました！
goto :end

:clean_all
echo [INFO] 全ての生成ファイルをクリーンアップ中...

REM ビルド成果物の削除
call :clean_build

REM デバッグファイルの削除
if exist "*.pdb" (
    echo [INFO] デバッグファイル(.pdb)を削除中...
    del /q "*.pdb"
    echo [SUCCESS] デバッグファイルを削除しました
)

if exist "*.ilk" (
    del /q "*.ilk"
    echo [SUCCESS] インクリメンタルリンクファイルを削除しました
)

REM ログファイルの削除
if exist "*.log" (
    echo [INFO] ログファイルを削除中...
    del /q "*.log"
    echo [SUCCESS] ログファイルを削除しました
)

REM 一時ファイルの削除
if exist "*.tmp" (
    echo [INFO] 一時ファイル(.tmp)を削除中...
    del /q "*.tmp"
    echo [SUCCESS] 一時ファイルを削除しました
)

if exist "*.temp" (
    del /q "*.temp"
    echo [SUCCESS] 一時ファイル(.temp)を削除しました
)

if exist "*~" (
    del /q "*~"
    echo [SUCCESS] バックアップファイルを削除しました
)

REM Visual Studio関連ファイルの削除
if exist ".vs" (
    echo [INFO] Visual Studio設定を削除中...
    rmdir /s /q ".vs"
    if not exist ".vs" echo [SUCCESS] Visual Studio設定を削除しました
)

if exist "*.vcxproj.user" (
    del /q "*.vcxproj.user"
    echo [SUCCESS] Visual Studioユーザー設定を削除しました
)

if exist "*.vcxproj.filters" (
    del /q "*.vcxproj.filters"
    echo [SUCCESS] Visual Studioフィルター設定を削除しました
)

if exist "*.sdf" (
    del /q "*.sdf"
    echo [SUCCESS] Visual Studio検索データベースを削除しました
)

if exist "*.opensdf" (
    del /q "*.opensdf"
    echo [SUCCESS] Visual Studio検索データベースを削除しました
)

if exist "*.suo" (
    del /q "*.suo"
    echo [SUCCESS] Visual Studioオプションファイルを削除しました
)

REM CMake関連ファイルの削除
if exist "CMakeCache.txt" (
    echo [INFO] CMakeキャッシュを削除中...
    del /q "CMakeCache.txt"
    echo [SUCCESS] CMakeキャッシュを削除しました
)

if exist "CMakeFiles" (
    rmdir /s /q "CMakeFiles"
    if not exist "CMakeFiles" echo [SUCCESS] CMakeファイルを削除しました
)

if exist "cmake_install.cmake" (
    del /q "cmake_install.cmake"
    echo [SUCCESS] CMakeインストールファイルを削除しました
)

REM OS固有ファイルの削除
if exist "Thumbs.db" (
    del /q "Thumbs.db"
    echo [SUCCESS] Windowsサムネイルキャッシュを削除しました
)

if exist ".DS_Store" (
    del /q ".DS_Store"
    echo [SUCCESS] macOS設定ファイルを削除しました
)

REM IDE関連ファイルの削除
if exist "*.swp" (
    del /q "*.swp"
    echo [SUCCESS] vim一時ファイルを削除しました
)

if exist "*.swo" (
    del /q "*.swo"
    echo [SUCCESS] vim一時ファイルを削除しました
)

REM パッケージファイルの削除
if exist "*.zip" (
    echo [INFO] 配布パッケージを削除中...
    del /q "*.zip"
    echo [SUCCESS] ZIPファイルを削除しました
)

if exist "*.tar.gz" (
    del /q "*.tar.gz"
    echo [SUCCESS] tar.gzファイルを削除しました
)

if exist "*.rar" (
    del /q "*.rar"
    echo [SUCCESS] RARファイルを削除しました
)

echo [SUCCESS] 全ての生成ファイルのクリーンアップが完了しました！
goto :end

:end
echo.
echo [INFO] クリーンアップスクリプトを終了します。
pause