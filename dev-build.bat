@echo off
chcp 65001 > nul
REM =====================================
REM AviUtlExEdit2PluginManager 開発ビルド
REM =====================================

echo [INFO] 開発ビルドを開始します...
echo.

REM 環境変数の設定
set PROJECT_NAME=AviUtlExEdit2PluginManager
set BUILD_TYPE=Debug
set OUTPUT_DIR=bin\debug

REM ビルドディレクトリの作成
if not exist "%OUTPUT_DIR%" (
    echo [INFO] ビルドディレクトリを作成: %OUTPUT_DIR%
    mkdir "%OUTPUT_DIR%"
)

REM プロジェクト存在チェック
if not exist "src" (
    echo [WARN] srcディレクトリが見つかりません。プロジェクトを初期化してください。
    echo [INFO] setup.bat を実行してプロジェクトを初期化できます。
    goto :error
)

echo [INFO] 開発ビルドモード: %BUILD_TYPE%
echo [INFO] 出力ディレクトリ: %OUTPUT_DIR%
echo.

REM ビルドコマンド（プロジェクトの種類に応じて調整が必要）
echo [INFO] コンパイル中...

REM C++ プロジェクトの場合
if exist "*.vcxproj" (
    echo [INFO] Visual Studio プロジェクトを検出
    msbuild *.vcxproj /p:Configuration=%BUILD_TYPE% /p:OutputPath=%OUTPUT_DIR%\
    if errorlevel 1 goto :error
)

REM C# プロジェクトの場合
if exist "*.csproj" (
    echo [INFO] .NET プロジェクトを検出
    dotnet build --configuration %BUILD_TYPE% --output %OUTPUT_DIR%
    if errorlevel 1 goto :error
)

REM Makefileの場合
if exist "Makefile" (
    echo [INFO] Makefileを検出
    make debug
    if errorlevel 1 goto :error
)

REM 汎用的なビルド（具体的なビルドツールが見つからない場合）
if not exist "%OUTPUT_DIR%\*" (
    echo [WARN] 特定のビルドシステムが検出されませんでした。
    echo [INFO] 手動でビルドコマンドを実行してください。
)

echo.
echo [SUCCESS] 開発ビルドが完了しました！
echo [INFO] 出力場所: %OUTPUT_DIR%
echo [INFO] デバッグ情報が含まれています。

REM 実行可能ファイルがあれば実行オプションを提示
if exist "%OUTPUT_DIR%\*.exe" (
    echo.
    echo ビルドされた実行ファイルを今すぐ実行しますか？ (y/n)
    set /p choice=""
    if /i "%choice%"=="y" (
        echo [INFO] プログラムを実行中...
        start "" "%OUTPUT_DIR%\*.exe"
    )
)

goto :end

:error
echo.
echo [ERROR] ビルドでエラーが発生しました。
echo [INFO] ログを確認して問題を解決してください。
exit /b 1

:end
echo.
echo [INFO] 開発ビルドスクリプトを終了します。
pause