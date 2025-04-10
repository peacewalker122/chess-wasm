<script lang="ts">
  import { pieceType, type Board } from "$lib/type";
  import { createBoard, initWasm } from "$lib/wasm";
  import { onMount } from "svelte";
  import "../styles/reset.css";
  // Create an array representing the chess board squares
  const boardSquares: Board[] = $state([]);
  const rows = 8;
  const cols = 8;

  // Initialize game state
  let wasmInitialized = false;
  let wasmError: string | null = null;

  // Selected piece for movement
  let selectedPiece: Piece | null = null;
  let possibleMoves: string[] = [];
  let isMoving = false;

  onMount(async () => {
    try {
      // Initialize WebAssembly
      await initWasm();
      wasmInitialized = true;

      // Create the chess board
      createBoard(true); // White goes first by default
    } catch (error) {
      console.error("Failed to initialize WASM:", error);
      wasmError = "Failed to load chess engine. Using fallback implementation.";

      // Use fallback implementation
      createBoard(true);
    }
  });

  const getIcon = (isLight: boolean, piece: pieceType) => {
    if (isLight) {
      return `w${piece}.svg`;
    } else {
      return `b${piece}.svg`;
    }
  };

  // helper function to help generate the chess pieces.
  // by default white first, then black follow.
  // override should change the order reverse, black first.
  const getPiece = (
    row: number,
    col: number,
    whiteFirst: boolean = true,
  ): string | null => {
    if (row == 1) {
      return getIcon(whiteFirst, pieceType.Pawn);
    }
    if (row == 6) {
      return getIcon(!whiteFirst, pieceType.Pawn);
    }
    if (row == 0) {
      if (col == 7 || col == 0) {
        return getIcon(whiteFirst, pieceType.Rook);
      }
      if (col == 6 || col == 1) {
        return getIcon(whiteFirst, pieceType.Knight);
      }
      if (col == 5 || col == 2) {
        return getIcon(whiteFirst, pieceType.Bishop);
      }
      if (col == 4) {
        return getIcon(whiteFirst, pieceType.King);
      }
      if (col == 3) {
        return getIcon(whiteFirst, pieceType.Queen);
      }
    }
    if (row == 7) {
      if (col == 7 || col == 0) {
        return getIcon(!whiteFirst, pieceType.Rook);
      }
      if (col == 6 || col == 1) {
        return getIcon(!whiteFirst, pieceType.Knight);
      }
      if (col == 5 || col == 2) {
        return getIcon(!whiteFirst, pieceType.Bishop);
      }
      if (col == 3) {
        return getIcon(!whiteFirst, pieceType.King);
      }
      if (col == 4) {
        return getIcon(!whiteFirst, pieceType.Queen);
      }
    }

    return null;
  };

  for (let i = rows - 1; i >= 0; i--) {
    for (let j = cols - 1; j >= 0; j--) {
      // Determine if square should be light or dark
      const isLight = (i + j) % 2 === 0;
      boardSquares.push({
        row: i,
        col: j,
        isLight: isLight,
        position: String.fromCharCode(101 - i) + String.fromCharCode(j + 1),
        piece: getPiece(i, j),
      });
    }
  }
</script>

<div class="container">
  <div class="chess-board">
    {#each boardSquares as square (square.position)}
      <div
        class="square {square.isLight ? 'light' : 'dark'}"
        data-position={square.position}
      >
        {#if square.piece !== null}
          <!-- the isLight parameter should be passed from user input when they choose white / black -->
          <img src={square.piece} alt="Pawn" />
        {/if}
      </div>
    {/each}
  </div>
</div>

<style>
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    width: 100%;
  }

  .chess-board {
    display: grid;
    grid-template-columns: repeat(8, 1fr);
    grid-template-rows: repeat(8, 1fr);
    width: 400px;
    height: 400px;
    border: 2px solid #333;
  }

  .square {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .light {
    background-color: #f0d9b5;
  }

  .dark {
    background-color: #b58863;
  }
</style>
