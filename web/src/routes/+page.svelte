<script lang="ts">
  import {pieceType, type Board} from "$lib/type";
  import "../styles/reset.css";
  // Create an array representing the chess board squares
  const boardSquares: Board[] = $state([]);
  const rows = 8;
  const cols = 8;

  const getIcon = (isLight: boolean, piece: pieceType) => {
    if (isLight) {
      return `w${piece}.svg`;
    } else {
      return `b${piece}.svg`;
    }
  };

  // helper function to help generate the chess pieces
  const getPiece = (row: number, col: number): string | null => {
    if (row === 1) {
      return getIcon(true, pieceType.Pawn);
    } else if (row === 6) {
      return getIcon(false, pieceType.Pawn);
    } else if (row === 0 || row === 7) {
      if (col === 0 || col === 7) {
        return getIcon(row === 0, pieceType.Rook);
      } else if (col === 1 || col === 6) {
        return getIcon(row === 0, pieceType.Knight);
      } else if (col === 2 || col === 5) {
        return getIcon(row === 0, pieceType.Bishop);
      } else if (col === 3) {
        return getIcon(row === 0, pieceType.Queen);
      } else if (col === 4) {
        return getIcon(row === 0, pieceType.King);
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
    <div class="square {square.isLight ? 'light' : 'dark'}" data-position={square.position}>
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
