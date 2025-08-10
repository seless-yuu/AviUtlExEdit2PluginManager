@echo off
chcp 65001 > nul
REM =====================================
REM AviUtlExEdit2PluginManager 本番ビルド
REM =====================================

echo [INFO] 本番ビルドを開始します...
echo.

REM 環境変数の設定
set PROJECT_NAME=AviUtlExEdit2PluginManager
set BUILD_TYPE=Release
set OUTPUT_DIR=bin\release
set DIST_DIR=dist

REM ビルドディレクトリの作成
if not exist "%OUTPUT_DIR%" (
    echo [INFO] ビルドディレクトリを作成: %OUTPUT_DIR%
    mkdir "%OUTPUT_DIR%"
)

if not exist "%DIST_DIR%" (
    echo [INFO] 配布ディレクトリを作成: %DIST_DIR%
    mkdir "%DIST_DIR%"
)

REM プロジェクト存在チェック
if not exist "src" (
    echo [WARN] srcディレクトリが見つかりません。プロジェクトを初期化してください。
    echo [INFO] setup.bat を実行してプロジェクトを初期化できます。
    goto :error
)

echo [INFO] 本番ビルドモード: %BUILD_TYPE%
echo [INFO] 出力ディレクトリ: %OUTPUT_DIR%
echo [INFO] 配布ディレクトリ: %DIST_DIR%
echo.

REM クリーンビルド
echo [INFO] 前回のビルド成果物をクリーンアップ中...
if exist "%OUTPUT_DIR%\*" del /q "%OUTPUT_DIR%\*"
if exist "%DIST_DIR%\*" del /q "%DIST_DIR%\*"

REM ビルドコマンド（プロジェクトの種類に応じて調整が必要）
echo [INFO] 最適化ビルド中...

REM C++ プロジェクトの場合
if exist "*.vcxproj" (
    echo [INFO] Visual Studio プロジェクトを検出
    msbuild *.vcxproj /p:Configuration=%BUILD_TYPE% /p:OutputPath=%OUTPUT_DIR%\ /p:Optimize=true
    if errorlevel 1 goto :error
)

REM C# プロジェクトの場合
if exist "*.csproj" (
    echo [INFO] .NET プロジェクトを検出
    dotnet build --configuration %BUILD_TYPE% --output %OUTPUT_DIR% --no-restore
    if errorlevel 1 goto :error
    
    REM 公開版の作成
    echo [INFO] 公開版を作成中...
    dotnet publish --configuration %BUILD_TYPE% --output %DIST_DIR% --no-build --self-contained false
    if errorlevel 1 goto :error
)

REM Makefileの場合
if exist "Makefile" (
    echo [INFO] Makefileを検出
    make clean && make release
    if errorlevel 1 goto :error
)

REM 配布ファイルのコピー
echo [INFO] 配布ファイルを準備中...

REM 必要なファイルを配布ディレクトリにコピー
if exist "%OUTPUT_DIR%\*.exe" copy "%OUTPUT_DIR%\*.exe" "%DIST_DIR%\"
if exist "%OUTPUT_DIR%\*.dll" copy "%OUTPUT_DIR%\*.dll" "%DIST_DIR%\"
if exist "README.md" copy "README.md" "%DIST_DIR%\"
if exist "LICENSE" copy "LICENSE" "%DIST_DIR%\"

REM 設定ファイルや依存関係のコピー
if exist "config" xcopy "config" "%DIST_DIR%\config\" /E /I /Y > nul
if exist "plugins" xcopy "plugins" "%DIST_DIR%\plugins\" /E /I /Y > nul

REM バージョン情報ファイルの作成
echo [INFO] バージョン情報を作成中...
echo %PROJECT_NAME% > "%DIST_DIR%\VERSION.txt"
echo Build Date: %date% %time% >> "%DIST_DIR%\VERSION.txt"
echo Build Type: %BUILD_TYPE% >> "%DIST_DIR%\VERSION.txt"

echo.
echo [SUCCESS] 本番ビルドが完了しました！
echo [INFO] 実行ファイル: %OUTPUT_DIR%
echo [INFO] 配布ファイル: %DIST_DIR%
echo [INFO] 最適化済みで配布準備完了です。

REM 配布パッケージの作成確認
echo.
echo 配布用ZIPファイルを作成しますか？ (y/n)
set /p choice=""
if /i "%choice%"=="y" (
    set ZIP_NAME=%PROJECT_NAME%_v%date:~0,4%%date:~5,2%%date:~8,2%.zip
    echo [INFO] ZIPファイルを作成中: !ZIP_NAME!
    
    REM PowerShellを使用してZIP作成
    powershell -command "Compress-Archive -Path '%DIST_DIR%\*' -DestinationPath '!ZIP_NAME!' -Force"
    if not errorlevel 1 (
        echo [SUCCESS] 配布パッケージが作成されました: !ZIP_NAME!
    ) else (
        echo [WARN] ZIP作成に失敗しました。手動で作成してください。
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
echo [INFO] 本番ビルドスクリプトを終了します。
pause