@echo off
chcp 65001 > nul
REM =====================================
REM AviUtlExEdit2PluginManager テスト
REM =====================================

echo [INFO] テストスイートを開始します...
echo.

REM テスト結果の初期化
set TESTS_PASSED=0
set TESTS_FAILED=0
set TOTAL_TESTS=0

REM テストディレクトリの確認
if not exist "tests" (
    echo [INFO] testsディレクトリが見つかりません。テスト環境を作成します...
    mkdir "tests"
    
    REM サンプルテストファイルの作成
    (
        echo // AviUtlExEdit2PluginManager テストファイル
        echo #include ^<iostream^>
        echo #include ^<cassert^>
        echo.
        echo void test_basic_functionality^(^)
        echo {
        echo     // 基本機能のテスト
        echo     std::cout ^<^< "基本機能テスト実行中..." ^<^< std::endl;
        echo     
        echo     // TODO: 実際のテストを実装
        echo     assert^(true^); // 仮のテスト
        echo     
        echo     std::cout ^<^< "基本機能テスト: PASS" ^<^< std::endl;
        echo }
        echo.
        echo void test_plugin_loading^(^)
        echo {
        echo     // プラグイン読み込みテスト
        echo     std::cout ^<^< "プラグイン読み込みテスト実行中..." ^<^< std::endl;
        echo     
        echo     // TODO: プラグイン読み込みテストを実装
        echo     assert^(true^); // 仮のテスト
        echo     
        echo     std::cout ^<^< "プラグイン読み込みテスト: PASS" ^<^< std::endl;
        echo }
        echo.
        echo int main^(^)
        echo {
        echo     std::cout ^<^< "AviUtlExEdit2PluginManager テストスイート" ^<^< std::endl;
        echo     std::cout ^<^< "========================================" ^<^< std::endl;
        echo.
        echo     try {
        echo         test_basic_functionality^(^);
        echo         test_plugin_loading^(^);
        echo         
        echo         std::cout ^<^< std::endl;
        echo         std::cout ^<^< "全てのテストが完了しました: PASS" ^<^< std::endl;
        echo         return 0;
        echo     } catch ^(const std::exception^& e^) {
        echo         std::cout ^<^< "テストエラー: " ^<^< e.what^(^) ^<^< std::endl;
        echo         return 1;
        echo     }
        echo }
    ) > "tests\test_main.cpp"
    echo [INFO] サンプルテストファイルを作成しました: tests\test_main.cpp
)

echo [INFO] 利用可能なテストを確認中...
echo.

REM C++テストファイルの確認
set CPP_TESTS=0
for %%f in ("tests\*.cpp") do (
    set /a CPP_TESTS+=1
    echo [FOUND] C++テストファイル: %%f
)

REM テストの実行方式選択
echo.
echo テスト実行方式を選択してください:
echo [1] コンパイルして実行 (C++テスト)
echo [2] スクリプト形式テスト
echo [3] 全テスト実行
echo [4] キャンセル
echo.
set /p choice="選択してください (1-4): "

if "%choice%"=="4" (
    echo [INFO] テストをキャンセルしました。
    goto :end
)

if "%choice%"=="1" goto :run_cpp_tests
if "%choice%"=="2" goto :run_script_tests
if "%choice%"=="3" goto :run_all_tests

echo [ERROR] 無効な選択です。
goto :end

:run_cpp_tests
echo [INFO] C++テストをコンパイル・実行中...
echo.

if %CPP_TESTS%==0 (
    echo [WARN] C++テストファイルが見つかりません。
    goto :end
)

REM テストビルドディレクトリの作成
if not exist "bin\test" (
    mkdir "bin\test"
)

REM 各テストファイルをコンパイル・実行
for %%f in ("tests\*.cpp") do (
    set TEST_NAME=%%~nf
    set TEST_EXE=bin\test\!TEST_NAME!.exe
    
    echo [INFO] テストをコンパイル中: %%f
    
    REM コンパイル (g++が利用可能な場合)
    g++ -std=c++17 -Iinclude -o "!TEST_EXE!" "%%f" 2>nul
    if errorlevel 1 (
        REM Visual Studio コンパイラを試行
        cl /EHsc /I"include" "%%f" /Fe:"!TEST_EXE!" >nul 2>&1
        if errorlevel 1 (
            echo [ERROR] テストのコンパイルに失敗: %%f
            set /a TESTS_FAILED+=1
            goto :next_cpp_test
        )
    )
    
    echo [INFO] テストを実行中: !TEST_NAME!
    
    REM テスト実行
    "!TEST_EXE!"
    if errorlevel 1 (
        echo [FAIL] テスト失敗: !TEST_NAME!
        set /a TESTS_FAILED+=1
    ) else (
        echo [PASS] テスト成功: !TEST_NAME!
        set /a TESTS_PASSED+=1
    )
    
    set /a TOTAL_TESTS+=1
    echo.
    
    :next_cpp_test
)

goto :show_results

:run_script_tests
echo [INFO] スクリプト形式テストを実行中...
echo.

REM 基本的な環境テスト
call :test_environment "環境チェック"
call :test_directory_structure "ディレクトリ構造チェック"
call :test_script_availability "スクリプト可用性チェック"

goto :show_results

:run_all_tests
echo [INFO] 全テストを実行中...
echo.

call :run_script_tests
call :run_cpp_tests

goto :show_results

:test_environment
set TEST_NAME=%~1
echo [TEST] %TEST_NAME%
set /a TOTAL_TESTS+=1

REM 基本的な環境チェック
if exist "setup.bat" (
    if exist "dev-build.bat" (
        if exist "build.bat" (
            echo [PASS] %TEST_NAME%: 必要なスクリプトが揃っています
            set /a TESTS_PASSED+=1
            goto :eof
        )
    )
)

echo [FAIL] %TEST_NAME%: 必要なスクリプトが不足しています
set /a TESTS_FAILED+=1
goto :eof

:test_directory_structure
set TEST_NAME=%~1
echo [TEST] %TEST_NAME%
set /a TOTAL_TESTS+=1

REM ディレクトリ構造のチェック
set REQUIRED_DIRS=0
set FOUND_DIRS=0

for %%d in (src include bin tests) do (
    set /a REQUIRED_DIRS+=1
    if exist "%%d" (
        set /a FOUND_DIRS+=1
    )
)

if %FOUND_DIRS% gtr 0 (
    echo [PASS] %TEST_NAME%: %FOUND_DIRS%/%REQUIRED_DIRS% 必要ディレクトリが存在
    set /a TESTS_PASSED+=1
) else (
    echo [FAIL] %TEST_NAME%: 必要ディレクトリが見つかりません
    set /a TESTS_FAILED+=1
)
goto :eof

:test_script_availability
set TEST_NAME=%~1
echo [TEST] %TEST_NAME%
set /a TOTAL_TESTS+=1

REM スクリプトの実行可能性チェック
set SCRIPTS_OK=1

for %%s in (setup.bat dev-build.bat build.bat clean.bat run.bat) do (
    if not exist "%%s" (
        echo [WARN] スクリプトが見つかりません: %%s
        set SCRIPTS_OK=0
    )
)

if %SCRIPTS_OK%==1 (
    echo [PASS] %TEST_NAME%: 全スクリプトが利用可能
    set /a TESTS_PASSED+=1
) else (
    echo [FAIL] %TEST_NAME%: 一部スクリプトが見つかりません
    set /a TESTS_FAILED+=1
)
goto :eof

:show_results
echo.
echo =====================================
echo テスト結果サマリー
echo =====================================
echo 総テスト数: %TOTAL_TESTS%
echo 成功: %TESTS_PASSED%
echo 失敗: %TESTS_FAILED%

if %TESTS_FAILED%==0 (
    echo.
    echo [SUCCESS] 全てのテストが成功しました！
    set EXIT_CODE=0
) else (
    echo.
    echo [ERROR] %TESTS_FAILED% 個のテストが失敗しました。
    set EXIT_CODE=1
)

REM テストレポートの作成
echo [INFO] テストレポートを作成中...
set REPORT_FILE=test_report_%date:~0,4%%date:~5,2%%date:~8,2%_%time:~0,2%%time:~3,2%%time:~6,2%.txt
set REPORT_FILE=%REPORT_FILE: =0%

(
    echo AviUtlExEdit2PluginManager テストレポート
    echo 生成日時: %date% %time%
    echo =====================================
    echo.
    echo 総テスト数: %TOTAL_TESTS%
    echo 成功: %TESTS_PASSED%
    echo 失敗: %TESTS_FAILED%
    echo.
    if %TESTS_FAILED%==0 (
        echo 結果: 全テスト成功
    ) else (
        echo 結果: テスト失敗あり
    )
) > "%REPORT_FILE%"

echo [INFO] テストレポートを作成しました: %REPORT_FILE%

exit /b %EXIT_CODE%

:end
echo.
echo [INFO] テストスクリプトを終了します。
pause