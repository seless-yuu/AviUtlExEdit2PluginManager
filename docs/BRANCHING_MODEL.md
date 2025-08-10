# ブランチモデル

このプロジェクトは、構造化され予測可能な開発プロセスを保証するために、**Git Flow**ブランチモデルに従います。

## メインブランチ

リポジトリは、無限のライフタイムを持つ2つのメインブランチを保持します：

- `main`: 公式で安定したリリース履歴を保存します。このブランチには`develop`または`hotfix`ブランチからのみマージされるべきです。
- `develop`: 新機能のための主要な統合ブランチです。すべての機能ブランチは`develop`から作成され、`develop`にマージされます。

## サポートブランチ

サポートブランチは、チームメンバー間の並行開発を支援し、機能の追跡を容易にし、リリースの準備を補助するために使用されます。メインブランチとは異なり、これらのブランチは作業が完了すると最終的に削除されるため、常に限られたライフタイムを持ちます。

使用する可能性のあるブランチの種類は次のとおりです：

### `feature/*`

- **目的:** 新機能の開発。
- **派生元:** `develop`
- **マージ先:** `develop`
- **命名規則:** `feature/<short-description>` (例: `feature/add-plugin-sorting`)

### `fix/*`

- **目的:** `develop`ブランチの重要でないバグの修正。
- **派生元:** `develop`
- **マージ先:** `develop`
- **命名規則:** `fix/<issue-description>` (例: `fix/profile-save-error`)

### `docs/*`

- **目的:** ドキュメントの追加、更新、または修正。
- **派生元:** `develop`
- **マージ先:** `develop`
- **命名規則:** `docs/<document-name>` (例: `docs/update-readme`)

### `release/*`

- **目的:** 新しい本番リリースの準備。このブランチは、最終的な修正と準備を可能にします。
- **派生元:** `develop`
- **マージ先:** `develop` と `main`
- **命名規則:** `release/vX.Y.Z` (例: `release/v0.2.0`)

### `hotfix/*`

- **目的:** 本番バージョンの重大なバグの修正。
- **派生元:** `main`
- **マージ先:** `develop` と `main`
- **命名規則:** `hotfix/<issue-description>`

## AIエージェントの作業フロー

複数のAIエージェントが同時に作業する際のコンフリクトを避けるため、各エージェントは自身の作業スペースとして`git worktree`を使用します。

1. 作業を開始する前に、`develop`ブランチから新しい機能ブランチを作成します。

```bash
git fetch origin
git branch feature/my-new-feature origin/develop
```

2.そのブランチ用の新しいワークツリーを作成します。

```bash
git worktree add ./worktrees/feature/my-new-feature feature/my-new-feature
```

3.エージェントはそのワークツリー内で作業を行います。

```bash
cd worktrees/feature/my-new-feature
# ...作業開始...
```

4.作業が完了したら、ワークツリーを削除します。

```bash
cd ../../..
git worktree remove ./worktrees/feature/my-new-feature
```
