<script>
  import { GetPokemonList } from "../wailsjs/go/main/App.js";
  import { GetImageUrlByPokemon } from "../wailsjs/go/main/App.js";

  let pokemons = [];
  let photos = [];
  let selectedPokemon;
  let showRandomPhoto = false;
  let showBreedPhotos = false;

  function init() {
    getPokemonList();
  }

  init();

  function getPokemonList() {
    GetPokemonList().then((result) => (pokemons = result));
  }

  function getImageUrlsByBreed() {
    init();
    showRandomPhoto = false;
    showBreedPhotos = false;
    GetImageUrlByPokemon(selectedPokemon).then((result) => (photos = result));
    showBreedPhotos = true;
  }
</script>
<main>
<h3>PokeAPI</h3>
<p></p>
<div>
  Click on down arrow to select a pokemon
  <select class="form-select col-12" bind:value={selectedPokemon}>
    {#each pokemons as pokemon, i}
      <option value={pokemon.ID}>
        {pokemon.Name}
      </option>
    {/each}
  </select>
  <button class="btn btn-primary col-12 mt-2" on:click={getImageUrlsByBreed}>
    Fetch by this pokemon
  </button>
</div>
</main>
<br />
{#if showBreedPhotos}
  {#each photos as photo}
    <img id="pokemon-photos" src={photo} alt="No pokemon found" />
  {/each}
{/if}

<style>

  #pokemon-photos {
    width: 300px;
    height: auto;
  }

  .btn:focus {
    border-width: 3px;
  }
</style>

