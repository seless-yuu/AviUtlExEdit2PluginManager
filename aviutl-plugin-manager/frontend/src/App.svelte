<script>
  import { onMount } from 'svelte';
  import { GetPlugins, GetProfiles, SaveProfile, ActivateProfile, LaunchAviUtl, AddPluginFromZip } from '../wailsjs/go/main/App.js';
  import * as runtime from '../wailsjs/runtime/runtime.js';

  let pluginsPromise;
  let profiles = [];
  let selectedProfileName = '';
  let selectedProfile = null;
  let aviutlPath = '';
  let statusMessage = '';
  let errorMessage = '';

  let saveTimeout;

  onMount(() => {
    refreshPlugins();
    loadProfiles();
  });

  function refreshPlugins() {
      pluginsPromise = GetPlugins();
  }

  function loadProfiles() {
    GetProfiles().then(result => {
      profiles = result;
      if (profiles && profiles.length > 0) {
        if (!selectedProfileName || !profiles.find(p => p.name === selectedProfileName)) {
            selectedProfileName = profiles[0].name;
        }
        updateSelectedProfile();
      }
    }).catch(handleError);
  }

  function updateSelectedProfile() {
    selectedProfile = profiles.find(p => p.name === selectedProfileName) || null;
  }

  function handlePluginToggle(pluginId) {
    if (!selectedProfile) return;

    selectedProfile.enabled_plugins[pluginId] = !selectedProfile.enabled_plugins[pluginId];

    clearTimeout(saveTimeout);
    saveTimeout = setTimeout(() => {
      SaveProfile(selectedProfile)
        .then(() => setStatus("Profile saved.", 2000))
        .catch(handleError);
    }, 500);
  }

  function activate() {
    if (!selectedProfile || !aviutlPath) {
      handleError("Please select a profile and specify the AviUtl path.");
      return;
    }
    setStatus("Activating profile...");
    ActivateProfile(selectedProfile.name, aviutlPath)
      .then(() => setStatus("Profile activated successfully!", 3000))
      .catch(handleError);
  }

  function launch() {
    if (!aviutlPath) {
        handleError("Please specify the AviUtl path.");
        return;
    }
    setStatus("Launching AviUtl...");
    LaunchAviUtl(aviutlPath)
        .then(() => setStatus("Launch command sent.", 2000))
        .catch(handleError);
  }

  function addPlugin() {
    runtime.OpenFileDialog({
        title: "Select Plugin Zip File",
        filters: [{ displayName: "Zip Archives", pattern: "*.zip" }]
    }).then(zipPath => {
        if (zipPath) {
            setStatus("Adding plugin...");
            AddPluginFromZip(zipPath)
                .then(newPlugin => {
                    setStatus(`Plugin '${newPlugin.name}' added successfully!`, 3000);
                    refreshPlugins(); // Refresh the list
                })
                .catch(handleError);
        }
    }).catch(handleError);
  }

  function setStatus(message, clearAfter = 0) {
    errorMessage = '';
    statusMessage = message;
    if (clearAfter > 0) {
      setTimeout(() => statusMessage = '', clearAfter);
    }
  }

  function handleError(error) {
    statusMessage = '';
    errorMessage = error.toString();
    setTimeout(() => errorMessage = '', 5000); // Clear error after 5s
  }

</script>

<main>
  <h1>AviUtl Plugin Manager</h1>

  <div class="status-bar">
    {#if statusMessage}
      <p class="status">{statusMessage}</p>
    {/if}
    {#if errorMessage}
      <p class="error">{errorMessage}</p>
    {/if}
  </div>

  <div class="top-controls">
    <div class="control-group">
      <label for="profile-select">Profile:</label>
      <select id="profile-select" bind:value={selectedProfileName} on:change={updateSelectedProfile}>
        {#if profiles && profiles.length > 0}
          {#each profiles as profile}
            <option value={profile.name}>{profile.name}</option>
          {/each}
        {/if}
      </select>
    </div>
    <div class="control-group">
        <label for="aviutl-path">AviUtl Directory:</label>
        <input type="text" id="aviutl-path" bind:value={aviutlPath} placeholder="C:\path\to\aviutl2 (exe is in this folder)" />
        <button on:click={activate} disabled={!selectedProfile || !aviutlPath}>Activate</button>
        <button on:click={launch} disabled={!aviutlPath}>Launch</button>
    </div>
  </div>

  <div class="list-header">
    <h2>Plugin Library</h2>
    <button on:click={addPlugin}>+ Add Plugin from Zip</button>
  </div>

  <div class="plugin-list">
    {#await pluginsPromise}
      <p>Loading plugins...</p>
    {:then plugins}
      {#if selectedProfile}
        <table>
          <thead>
            <tr>
              <th>Enabled</th>
              <th>Name</th>
              <th>Version</th>
              <th>Author</th>
            </tr>
          </thead>
          <tbody>
            {#if plugins && plugins.length > 0}
                {#each plugins as plugin (plugin.id)}
                <tr>
                    <td>
                    <input
                        type="checkbox"
                        checked={selectedProfile.enabled_plugins[plugin.id] || false}
                        on:change={() => handlePluginToggle(plugin.id)}
                    />
                    </td>
                    <td>{plugin.name || 'N/A'}</td>
                    <td>{plugin.version || 'N/A'}</td>
                    <td>{plugin.author || 'N/A'}</td>
                </tr>
                {/each}
            {:else}
                <tr>
                    <td colspan="4" style="text-align: center;">No plugins found in library.</td>
                </tr>
            {/if}
          </tbody>
        </table>
      {:else}
        <p>Select a profile to see plugins.</p>
      {/if}
    {:catch error}
      <p style="color: red;">Error loading plugins: {error}</p>
    {/await}
  </div>
</main>

<style>
  :root {
    --primary-color: #42b983;
    --secondary-color: #3498db;
    --error-color: #e74c3c;
    --status-color: #2c3e50;
    --light-gray: #f8f9fa;
    --medium-gray: #e9ecef;
    --dark-gray: #ced4da;
    --text-color: #2c3e50;
  }

  main {
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
    padding: 1.5em;
    max-width: 1000px;
    margin: 0 auto;
    color: var(--text-color);
  }

  h1 { text-align: center; margin-bottom: 1.5rem; }
  .status-bar { text-align: center; margin-bottom: 1rem; min-height: 24px; }
  .status { color: var(--status-color); font-weight: bold; }
  .error { color: var(--error-color); font-weight: bold; }

  .top-controls {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
    background-color: var(--light-gray);
    padding: 1rem;
    border-radius: 8px;
  }
  .control-group { display: flex; align-items: center; gap: 0.5rem; }
  label { font-weight: 500; }

  select, input[type="text"], button {
    padding: 0.5rem;
    border: 1px solid var(--dark-gray);
    border-radius: 4px;
    font-size: 0.95rem;
  }
  input[type="text"] { min-width: 300px; }

  button {
    background-color: var(--primary-color);
    color: white;
    border: none;
    cursor: pointer;
    transition: background-color 0.2s;
  }
  button:hover { filter: brightness(1.1); }
  button:disabled { background-color: var(--medium-gray); cursor: not-allowed; }

  .list-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;
  }
  .list-header h2 { margin: 0; }
  .list-header button { background-color: var(--secondary-color); }

  table {
    width: 100%;
    border-collapse: collapse;
  }
  th, td { border: 1px solid #dee2e6; padding: 12px; text-align: left; }
  thead { background-color: var(--light-gray); }
  tbody tr:nth-child(odd) { background-color: #fdfdfd; }
  input[type="checkbox"] { width: 1.2em; height: 1.2em; }
</style>
