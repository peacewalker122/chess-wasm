<script lang="ts">
  import {pieceType, type Board} from "$lib/type";
  import {createBoard, initWasm, startMove} from "$lib/wasm";
  import {onMount} from "svelte";
  import "../styles/reset.css";
  // Create an array representing the chess board squares
  let boardSquares: Board[] = $state([]);
  const rows = 8;
  const cols = 8;

  // Initialize game state
  let wasmInitialized = false;
  let wasmError: string | null = null;

  // Selected piece for movement
  let possibleMoves: string[] = [];
  let isMoving = false;
  let selectedPosition: string | null = $state(null);

  onMount(async () => {
    try {
      console.log("initializing wasm");
      // Initialize WebAssembly
      await initWasm();
      wasmInitialized = true;
      console.log("wasm initialized");

      const board = createBoard(true);
      if (board) {
        boardSquares = board;
      }

      console.log("board: ", boardSquares);
      console.log("WASM initialized successfully");
    } catch (error) {
      console.error("Failed to initialize WASM:", error);
      wasmError = "Failed to load chess engine. Using fallback implementation.";
    }
  });

  const handleSquareClick = (position: string) => {
    console.log("Square clicked:", position);
    // Find the square in our board state
    const clickedSquareIndex = boardSquares.findIndex(
      (s) => s.position === position,
    );
    if (clickedSquareIndex === -1) return;

    const clickedSquare = boardSquares[clickedSquareIndex];

    if (isMoving && selectedPosition) {
      // Move the piece from selected square to clicked square
      const selectedSquareIndex = boardSquares.findIndex(
        (s) => s.position === selectedPosition,
      );
      if (selectedSquareIndex === -1) return;

      const selectedSquare = boardSquares[selectedSquareIndex];

      // Only move if there's a piece to move
      if (selectedSquare.piece) {
        // Store the piece that's being moved
        const movingPiece = selectedSquare.piece;

        // Update board state (this will automatically update the UI)
        console.log("selectedPosition: ", selectedPosition);
        console.log("selectedSquare: ", position);
        const board = startMove(selectedPosition, position);
        if (board) {
          boardSquares = board;
        }

        // Reset movement state
        isMoving = false;
        selectedPosition = null;
      }
    } else if (clickedSquare.piece) {
      // Select this square if it has a piece
      isMoving = true;
      selectedPosition = position;
    }
  };
</script>

<div class="container">
  <div class="chess-board">
    {#each boardSquares as square (square.position)}
    <div id={square.position} class="square {square.isLight ? 'light' : 'dark'} {selectedPosition ===
        square.position
          ? 'selected'
          : ''}" data-position={square.position} onclick={()=> handleSquareClick(square.position)}
      >
      {#if square.piece !== null}
      <!-- the isLight parameter should be passed from user input when they choose white / black -->
      <img src={square.piece} alt="Chess piece" />
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
