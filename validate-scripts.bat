@echo off
chcp 65001 > nul
REM =====================================
REM ユーティリティスクリプト検証
REM =====================================

echo [INFO] ユーティリティスクリプトの検証を開始します...
echo.

REM 必要なスクリプトファイルのリスト
set REQUIRED_SCRIPTS=setup.bat dev-build.bat build.bat run.bat test.bat clean.bat

REM 検証結果
set SCRIPTS_FOUND=0
set SCRIPTS_MISSING=0

echo [INFO] スクリプトファイルの存在確認中...
echo.

for %%s in (%REQUIRED_SCRIPTS%) do (
    if exist "%%s" (
        echo [✓] %%s - 存在
        set /a SCRIPTS_FOUND+=1
    ) else (
        echo [✗] %%s - 見つかりません
        set /a SCRIPTS_MISSING+=1
    )
)

echo.
echo [INFO] ファイルサイズ確認中...
echo.

for %%s in (%REQUIRED_SCRIPTS%) do (
    if exist "%%s" (
        for %%f in ("%%s") do (
            echo [INFO] %%s - %%~zf bytes
        )
    )
)

echo.
echo =====================================
echo 検証結果
echo =====================================
echo 見つかったスクリプト: %SCRIPTS_FOUND%
echo 見つからないスクリプト: %SCRIPTS_MISSING%

if %SCRIPTS_MISSING%==0 (
    echo.
    echo [SUCCESS] 全てのユーティリティスクリプトが正常に配置されています！
    echo.
    echo 利用可能なスクリプト:
    echo ├── setup.bat      - プロジェクト初期化
    echo ├── dev-build.bat  - 開発ビルド
    echo ├── build.bat      - 本番ビルド  
    echo ├── run.bat        - プログラム実行
    echo ├── test.bat       - テスト実行
    echo └── clean.bat      - クリーンアップ
    echo.
    echo 使い方: まず setup.bat を実行してプロジェクトを初期化してください。
) else (
    echo.
    echo [ERROR] 一部のスクリプトが見つかりません。
    echo [INFO] プロジェクトを再度セットアップしてください。
)

echo.
pause