import { json } from "@sveltejs/kit";
import { type Board } from "$lib/type";

// src/wasm.ts
let isReady = false;

export function initWasm(): Promise<void> {
	if (isReady) return Promise.resolve();

	const go = new (window as any).Go();

	const result = WebAssembly.instantiateStreaming(
		fetch("/main.wasm"),
		go.importObject,
	)
		.then((result) => {
			go.run(result.instance);
			isReady = true;
			console.log("WebAssembly module initialized successfully");
		})
		.catch((error) => {
			console.error("Failed to initialize WebAssembly module:", error);
			throw error;
		});

	return result;
}

export function createBoard(isWhiteFirst = true): Board[] | null {
	if (!isReady) {
		console.error("WebAssembly module is not initialized yet.");
		return null;
	}

	try {
		// Call the Go WASM function to build the board
		const result = (window as any).buildBoard(isWhiteFirst);
		// console.log("Go WASM board result:", result);

		const board: Board[] = JSON.parse(result);

		// Convert the Go board representation to our TypeScript GameState
		// return convertToGameState(result, isWhiteFirst);
		return board;
	} catch (error) {
		console.error("Error creating board:", error);
		// Fallback to client-side initialization if WASM fails
		// return initializeChessGame();
		throw error;
	}
}
