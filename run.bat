@echo off
chcp 65001 > nul
REM =====================================
REM AviUtlExEdit2PluginManager 実行
REM =====================================

echo [INFO] プログラム実行スクリプトを開始します...
echo.

REM 実行モードの選択
echo 実行モードを選択してください:
echo [1] デバッグ版実行 (bin\debug)
echo [2] リリース版実行 (bin\release)
echo [3] 配布版実行 (dist)
echo [4] 自動選択 (利用可能な最新版)
echo [5] キャンセル
echo.
set /p choice="選択してください (1-5): "

if "%choice%"=="5" (
    echo [INFO] 実行をキャンセルしました。
    goto :end
)

if "%choice%"=="1" goto :run_debug
if "%choice%"=="2" goto :run_release
if "%choice%"=="3" goto :run_dist
if "%choice%"=="4" goto :run_auto

echo [ERROR] 無効な選択です。
goto :end

:run_debug
set TARGET_DIR=bin\debug
set BUILD_TYPE=Debug
echo [INFO] デバッグ版を実行します...
goto :check_and_run

:run_release
set TARGET_DIR=bin\release
set BUILD_TYPE=Release
echo [INFO] リリース版を実行します...
goto :check_and_run

:run_dist
set TARGET_DIR=dist
set BUILD_TYPE=Distribution
echo [INFO] 配布版を実行します...
goto :check_and_run

:run_auto
echo [INFO] 利用可能なバージョンを自動選択中...

REM 優先順位: dist > bin\release > bin\debug
if exist "dist\*.exe" (
    set TARGET_DIR=dist
    set BUILD_TYPE=Distribution
    echo [INFO] 配布版を実行します (自動選択)
    goto :check_and_run
)

if exist "bin\release\*.exe" (
    set TARGET_DIR=bin\release
    set BUILD_TYPE=Release
    echo [INFO] リリース版を実行します (自動選択)
    goto :check_and_run
)

if exist "bin\debug\*.exe" (
    set TARGET_DIR=bin\debug
    set BUILD_TYPE=Debug
    echo [INFO] デバッグ版を実行します (自動選択)
    goto :check_and_run
)

echo [ERROR] 実行可能ファイルが見つかりません。
echo [INFO] 先にビルドを実行してください:
echo [INFO]   - dev-build.bat (デバッグビルド)
echo [INFO]   - build.bat (リリースビルド)
goto :error

:check_and_run
REM ディレクトリの存在確認
if not exist "%TARGET_DIR%" (
    echo [ERROR] ディレクトリが見つかりません: %TARGET_DIR%
    echo [INFO] 先にビルドを実行してください。
    goto :error
)

REM 実行ファイルの検索
set EXE_FILE=""
for %%f in ("%TARGET_DIR%\*.exe") do (
    set EXE_FILE=%%f
    goto :found_exe
)

echo [ERROR] 実行ファイルが見つかりません: %TARGET_DIR%
echo [INFO] 先にビルドを実行してください。
goto :error

:found_exe
echo [INFO] 実行ファイル: %EXE_FILE%
echo [INFO] ビルドタイプ: %BUILD_TYPE%
echo.

REM 実行前の確認
echo 実行しますか？ (y/n)
set /p run_choice=""
if /i not "%run_choice%"=="y" (
    echo [INFO] 実行をキャンセルしました。
    goto :end
)

REM 作業ディレクトリを実行ファイルのディレクトリに変更
pushd "%TARGET_DIR%"

echo [INFO] プログラムを実行中...
echo =====================================
echo.

REM プログラム実行
%EXE_FILE%
set EXIT_CODE=%errorlevel%

echo.
echo =====================================
echo [INFO] プログラムが終了しました。
echo [INFO] 終了コード: %EXIT_CODE%

REM 元のディレクトリに戻る
popd

if %EXIT_CODE% neq 0 (
    echo [WARN] プログラムがエラーで終了しました。
    echo [INFO] ログを確認してください。
) else (
    echo [SUCCESS] プログラムが正常に終了しました。
)

REM ログファイルの確認
if exist "%TARGET_DIR%\*.log" (
    echo.
    echo ログファイルを確認しますか？ (y/n)
    set /p log_choice=""
    if /i "%log_choice%"=="y" (
        echo [INFO] ログファイルを表示中...
        for %%f in ("%TARGET_DIR%\*.log") do (
            echo --- %%f ---
            type "%%f"
            echo.
        )
    )
)

goto :end

:error
echo [ERROR] 実行に失敗しました。
echo.
echo 解決方法:
echo 1. setup.bat を実行してプロジェクトを初期化
echo 2. dev-build.bat または build.bat でビルド
echo 3. 再度 run.bat を実行
exit /b 1

:end
echo.
echo [INFO] 実行スクリプトを終了します。
pause